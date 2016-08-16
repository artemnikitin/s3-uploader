# s3-uploader 
[![Go Report Card](https://goreportcard.com/badge/artemnikitin/s3-uploader)](https://goreportcard.com/report/artemnikitin/s3-uploader)   [![codebeat badge](https://codebeat.co/badges/67984735-0e55-4d39-aa38-6213b14ed456)](https://codebeat.co/projects/github-com-artemnikitin-s3-uploader)   [![Build Status](https://travis-ci.org/artemnikitin/s3-uploader.svg?branch=master)](https://travis-ci.org/artemnikitin/s3-uploader)    
Simple uploader for AWS S3, written in Go. Development continues [here](https://github.com/artemnikitin/s3-tool)   
##### Dependency

Only depends on AWS SDK. Install it via    
```
go get github.com/aws/aws-sdk-go/...
```

##### AWS Credentials

Currently assumes that you will have credentials settled as environmental variables.   
```
export AWS_ACCESS_KEY_ID=<key>
export AWS_SECRET_ACCESS_KEY=<secret>
```

##### Running
Get it via    
``` 
go get github.com/artemnikitin/s3-uploader 
``` 
   
Required parameters:          
``` 
s3-uploader -path=/path/to/file -bucket=bucket_name 
```   
```path``` can be specified as path to a specific file or an entire folder. In case of folder, all content of the folder will be uploaded with respect to structure of files in the folder.   
   
Additional optional parameters:   
- ```region``` set S3 region, by default region will be set to ```us-east-1```       
Example:    
``` 
s3-uploader -path=/path/to/file -bucket=bucket_name -region=region-name 
```    
- ```rename``` gives an ability to rename file for upload      
Example:   
``` 
s3-uploader -path=/path/to/file -bucket=bucket_name -rename=newname.file
```   
- ```uploadto``` create specific key (like folder structure inside bucket) for S3 bucket, by default equal to ```/```   
Example:   
``` 
s3-uploader -path=/path/to/file -bucket=bucket_name -uploadto=/path/inside/S3/bucket/
```  

You can specify parameter ```-log``` for logging AWS requests and responses.
