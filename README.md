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
   
Required parameters:          
``` s3-uploader -path=/path/to/file -bucket=bucket_name ```   
```path``` can be specified as path to specific file or entire folder. In case of folder, all content of folder will be uploaded with respect to structure of files in folder.   
   
Additional optional parameters:   
- ```region``` set S3 region, by default region will be set to ```us-east-1```       
Example:    
``` s3-uploader -path=/path/to/file -bucket=bucket_name -region=region-name ```    
- ```rename``` gives ability to rename file for upload      
Example:   
``` s3-uploader -path=/path/to/file -bucket=bucket_name -rename=newname.file```   
- ```uploadto``` create specific key (like folder structure), by default equla to ```/```   
Example:   
``` s3-uploader -path=/path/to/file -bucket=bucket_name -uploadto=/path/inside/S3/bucket/```  

You can specify parameter ```-log=true``` for logging AWS requests and responses.

##### TODO  
1. Managing additional files details (like permissions, storage class, etc)
2. Alternative ways to authenticate in AWS
