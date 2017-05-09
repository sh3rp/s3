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
)

type S3Service struct {
	s3     *s3.S3
	bucket string
}

func GetService(region string, bucket string) *S3Service {
	access_key := os.Getenv("AWS_ACCESS_KEY")
	secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	token := ""
	creds := credentials.NewStaticCredentials(access_key, secret_access_key, token)
	config := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess := session.Must(session.NewSession())
	svc := s3.New(sess, config)
	return &S3Service{s3: svc, bucket: bucket}
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

func (svc *S3Service) S3GetObject(key string, toFile string) {
	result, err := svc.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(svc.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal("Failed to get object", err)
	}

	file, err := os.Create(toFile)
	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	if _, err := io.Copy(file, result.Body); err != nil {
		log.Fatal("Failed to copy object to file", err)
	}
	result.Body.Close()
	file.Close()
}
