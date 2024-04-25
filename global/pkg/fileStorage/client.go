package fileStorage

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type bucketDetails struct {
	BucketName string
	Region     string
}

type S3FullClient struct {
	S3Client *s3.Client
	Data     bucketDetails
}

func SetS3(cfg aws.Config) (*S3FullClient, error) {

	var region string = os.Getenv("AWS_REGION")
	var fileStorageIP string = os.Getenv("file_storage_ip")

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fileStorageIP)
		o.UsePathStyle = true
	})
	// file-storage:4566
	var bucketName string = os.Getenv("bucketName")
	S3FullClient := S3FullClient{S3Client: s3Client, Data: bucketDetails{Region: region, BucketName: bucketName}}

	// S3FullClient.data.Region = region
	// S3FullClient.data.BucketName = bucketName
	return &S3FullClient, nil
}

// package main

// func main() {
// 	var region string = os.Getenv("AWS_REGION")
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion(region), // Replace with your region
// 	)
// 	if err != nil {
// 		log.Fatalf("Unable to load SDK config, %v", err)
// 	}

// 	s3Client := s3.NewFromConfig(cfg)

// 	bucketBasics := s3Connect.BucketBasics{S3Client: s3Client}
// 	queue.ListenToQueue(bucketBasics)
// }
