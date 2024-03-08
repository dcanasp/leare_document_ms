package fileStorage

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadFile reads from a file and puts the data into an object in a bucket.
func (client S3FullClient) UploadBuffer(folderPath string, uuid string) error {

	files, err := os.ReadDir("temp") // Read files from the "temp" folder
	if err != nil {
		return fmt.Errorf("couldn't read files from temp folder: %w", err)
	}

	for _, file := range files {
		objectKey := file.Name()
		buffer, err := os.ReadFile("temp/" + objectKey)
		if err != nil {
			return fmt.Errorf("Error getting buffer for file %s: %v\n", objectKey, err)
			// continue // Skip to the next file if there's an error
		}

		fullObjectKey := folderPath + "/" + uuid + "/" + objectKey
		_, err = client.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(client.Data.BucketName),
			Key:    aws.String(fullObjectKey),
			Body:   bytes.NewReader(buffer),
		})
		if err != nil {
			return fmt.Errorf("Couldn't upload buffer to %v:%v. Here's why: %v\n", client.Data.BucketName, fullObjectKey, err)
		}
	}

	return nil
}

func (client S3FullClient) DeleteItem(folderPath string, uuid string) error {

	files, err := os.ReadDir("temp") // Read files from the "temp" folder
	if err != nil {
		return fmt.Errorf("couldn't read files from temp folder: %w", err)
	}

	for _, file := range files {
		objectKey := file.Name()
		buffer, err := os.ReadFile("temp/" + objectKey)
		if err != nil {
			return fmt.Errorf("Error getting buffer for file %s: %v\n", objectKey, err)
			// continue // Skip to the next file if there's an error
		}

		fullObjectKey := folderPath + "/" + uuid + "/" + objectKey
		_, err = client.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(client.Data.BucketName),
			Key:    aws.String(fullObjectKey),
			Body:   bytes.NewReader(buffer),
		})
		if err != nil {
			return fmt.Errorf("Couldn't upload buffer to %v:%v. Here's why: %v\n", client.Data.BucketName, fullObjectKey, err)
		}
	}

	return nil
}
