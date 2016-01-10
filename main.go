package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	logging    = flag.Bool("log", false, "Enable logging")
	filepath   = flag.String("path", "", "Path to file")
	bucket     = flag.String("bucket", "", "Specify S3 bucket")
	region     = flag.String("region", "us-east-1", "Set S3 region")
	rename     = flag.String("rename", "", "Set a new name for file")
	uploadpath = flag.String("uploadto", "/", "Set a specific path for a file inside S3 bucket")
)

func main() {
	flag.Parse()
	if *filepath == "" || *bucket == "" {
		fmt.Println("Please specify correct parameters!")
		fmt.Println("You should specify:")
		fmt.Println("-path with path to file you want to upload")
		fmt.Println("-bucket name of bucket in S3 where you want to upload")
		os.Exit(1)
	}

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatal("Failed to open a file with an error: ", err)
	}

	session := session.New(createConfig())
	service := s3manager.NewUploader(session)

	uploadFile(*service, *bucket, *uploadpath+getFileName(*filepath), file)
}

func uploadFile(uploader s3manager.Uploader, bucket, key string, file io.Reader) {
	resp, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Println("Failed to upload a file because of: ", err)
		return
	}
	fmt.Println("File was successfully uploaded! Location: ", resp.Location)
}

func getFileName(filepath string) string {
	if *rename != "" {
		return *rename
	} else {
		index := strings.LastIndex(filepath, "/")
		if index != -1 {
			return filepath[index+1:]
		} else {
			return ""
		}
	}
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
