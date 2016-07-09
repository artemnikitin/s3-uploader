package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	logging    = flag.Bool("log", false, "Enable logging")
	region     = flag.String("region", "us-east-1", "Set S3 region")
	bucket     = flag.String("bucket", "", "Specify S3 bucket")
	filesPath  = flag.String("path", "", "Path to file")
	rename     = flag.String("rename", "", "Set a new name for file")
	uploadPath = flag.String("uploadto", "", "Set a specific path for a file inside S3 bucket")
)

func main() {
	flag.Parse()
	if *filesPath == "" || *bucket == "" {
		fmt.Println("Please specify correct parameters!")
		fmt.Println("You should specify:")
		fmt.Println("-path with path to file you want to upload")
		fmt.Println("-bucket name of bucket in S3 where you want to upload")
		os.Exit(1)
	}

	file, err := os.Open(*filesPath)
	if err != nil {
		log.Fatal("Failed to open a file with an error: ", err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get info about file with an error: ", err)
	}

	session := session.New(createConfig())
	service := s3manager.NewUploader(session)

	switch mode := info.Mode(); {
	case mode.IsDir():
		uploadDirectory(*service, *file)
	case mode.IsRegular():
		uploadFile(*service, *uploadPath+getFileName(*filesPath), file)
	}
}

func uploadFile(uploader s3manager.Uploader, key string, file io.Reader) {
	resp, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Println("Failed to upload a file because of: ", err)
		return
	}
	log.Println("File was successfully uploaded! Location:", resp.Location)
}

func uploadDirectory(uploader s3manager.Uploader, file os.File) {
	var wg sync.WaitGroup
	err := filepath.Walk(*filesPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(path)
			if err == nil {
				path := getPathInsideFolder(path, getFolderName(*filesPath))
				wg.Add(1)
				go func() {
					uploadFile(uploader, createKey(path), file)
					wg.Done()
					defer file.Close()
				}()
			} else {
				log.Println("Can't open a file because of: ", err)
			}
		}
		return nil
	})
	wg.Wait()
	if err != nil {
		log.Println("Can't process directory because of:", err)
		return
	}
	log.Println("Directory was successfully uploaded!")
}

func createConfig() *aws.Config {
	config := aws.NewConfig()
	config.WithCredentials(credentials.NewEnvCredentials())
	config.WithRegion(*region)
	if *logging {
		config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}
	return config
}
