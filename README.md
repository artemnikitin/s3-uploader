# s3-uploader
Simple uploader for AWS S3 written in Go

##### Dependency

Only depends on AWS SDK. Install it via    
```go get -u github.com/aws/aws-sdk-go/...```

##### AWS Credentials

Currently assumes that you will have credentials settled as environmental variables.
```
export AWS_ACCESS_KEY_ID=<key>
export AWS_SECRET_ACCESS_KEY=<secret>
```

##### Running
Get it via    
``` go get github.com/artemnikitin/s3-uploader ```    
To upload file use:       
``` s3-uploader -path=/path/to/file -bucket=bucket_name ```    
You can set up region of S3 bucket:   
``` s3-uploader -path=/path/to/file -bucket=bucket_name -region=region-name ```    
By default, region will be set to "us-east-1".   
You can specify parameter ```-log=true``` for logging AWS requests and responses.

##### TODO  
1. Set custom key (like folder structure)
2. Uploading whole directory with saving structure of it
3. Alternative ways to authenticate in AWS
