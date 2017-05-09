package aws

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Service struct {
	s3            *s3.S3
	uploadManager *s3manager.Uploader
	bucket        string
}

func GetService(region string, bucket string) *S3Service {
	access_key := os.Getenv("AWS_ACCESS_KEY")
	secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	token := ""
	creds := credentials.NewStaticCredentials(access_key, secret_access_key, token)
	config := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess := session.Must(session.NewSession(config))
	svc := s3.New(sess)
	uploader := s3manager.NewUploader(sess)
	return &S3Service{s3: svc, uploadManager: uploader, bucket: bucket}
}

func (svc *S3Service) S3ListBucket() {
	resp, err := svc.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: &svc.bucket,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%-30s %s\n", "Name", "Size")
	for _, obj := range resp.Contents {
		fmt.Printf("%-30s %d\n", *obj.Key, *obj.Size)
	}
}

func (svc *S3Service) S3PutObject(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %q, %v", filename, err)
	}

	result, err := svc.uploadManager.Upload(&s3manager.UploadInput{
		Bucket: aws.String(svc.bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}
	fmt.Printf("File uploaded to %s\n", result.Location)
}

func (svc *S3Service) S3GetObject(key string, toFile string) {
	result, err := svc.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(svc.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal("Failed to get object", err)
	}
	defer result.Body.Close()

	file, err := os.Create(toFile)
	defer file.Close()

	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	if _, err := io.Copy(file, result.Body); err != nil {
		log.Fatal("Failed to copy object to file", err)
	}
}
