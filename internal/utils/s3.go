package utils

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func InitS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		),
	))
	if err != nil {
		panic("Failed to load AWS config")
	}

	s3Client = s3.NewFromConfig(cfg)
}

func UploadToS3(file multipart.File, fileName string) (string, error) {
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName), nil
}
