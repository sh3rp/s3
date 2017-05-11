package aws

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Service struct {
	s3                *s3.S3
	uploadManager     *s3manager.Uploader
	Bucket            string
	Access_key        string
	Secret_access_key string
	Region            string
}

func GetService(region string, bucket string) *S3Service {
	access_key := os.Getenv("S3_ACCESS_KEY")
	secret_access_key := os.Getenv("S3_SECRET_ACCESS_KEY")
	if bucket == "" {
		bucket = os.Getenv("S3_BUCKET")
	}
	if region == "" {
		region = os.Getenv("S3_REGION")
	}
	token := ""
	creds := credentials.NewStaticCredentials(access_key, secret_access_key, token)
	config := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	sess := session.Must(session.NewSession(config))
	svc := s3.New(sess)
	uploader := s3manager.NewUploader(sess)
	return &S3Service{
		s3:                svc,
		uploadManager:     uploader,
		Bucket:            bucket,
		Access_key:        access_key,
		Secret_access_key: secret_access_key,
		Region:            region,
	}
}

func (svc *S3Service) S3ListBucket(key string) {
	objects := svc.S3GetObjects(key)
	for _, obj := range objects {
		fmt.Printf("%-30s %d\n", *obj.Key, *obj.Size)
	}
}

func (svc *S3Service) S3PutObject(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %q, %v", filename, err)
	}

	result, err := svc.uploadManager.Upload(&s3manager.UploadInput{
		Bucket: aws.String(svc.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}
	fmt.Printf("File uploaded to %s\n", result.Location)
}

func (svc *S3Service) S3GetObject(key string) {
	result, err := svc.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(svc.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal("Failed to get object", err)
	}
	defer result.Body.Close()

	if strings.Contains(key, "/") {
		dir := key[0:strings.LastIndex(key, "/")]
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			} else {
				log.Fatal("Error for directory: %v", err)
			}
		}
	}

	file, err := os.Create(key)
	defer file.Close()

	if err != nil {
		log.Fatal("Failed to create file", err)
	}
	if _, err := io.Copy(file, result.Body); err != nil {
		log.Fatal("Failed to copy object to file", err)
	}
}

func (svc S3Service) S3RawListObject(key string) {
	resp, err := svc.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: &svc.Bucket,
		Prefix: aws.String(key),
	})

	if err != nil {
		log.Fatal("Raw list failed")
	}

	fmt.Println(resp)
}

func (svc S3Service) S3RemoveObject(key string) {
	_, err := svc.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &svc.Bucket,
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Removed %s\n", key)
	}
}

func (svc S3Service) S3GetObjects(key string) []*s3.Object {
	resp, err := svc.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: &svc.Bucket,
		Prefix: aws.String(key),
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var objects []*s3.Object

	for _, obj := range resp.Contents {
		objects = append(objects, obj)
	}

	return objects
}
