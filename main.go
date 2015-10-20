package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	logging  = flag.Bool("log", false, "Enable logging")
	filepath = flag.String("path", "", "Path to file")
	bucket   = flag.String("bucket", "", "Specify S3 bucket")
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
		fmt.Println("Failed to open a file.", err)
		os.Exit(1)
	}

	service := s3manager.NewUploader(&s3manager.UploadOptions{
		S3: s3.New(createConfig()),
	})
	resp, err := service.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String("/" + getFileName(*filepath)),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Failed to upload a file.", err)
		os.Exit(1)
	}

	fmt.Println("---------------------")
	fmt.Println("File was successfully uploaded!")
	fmt.Println("Location:", resp.Location)
}

func getFileName(filepath string) string {
	index := strings.LastIndex(filepath, "/")
	if index != -1 {
		return filepath[index+1:]
	} else {
		return ""
	}
}

func createConfig() *aws.Config {
	config := aws.NewConfig()
	config.WithCredentials(credentials.NewEnvCredentials())
	config.WithRegion("us-east-1")
	if *logging {
		config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}
	return config
}
