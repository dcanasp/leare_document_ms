package fileStorage

import (
	"bytes"
	"context"
	"fmt"
	logs "global/logging"
	"global/utils"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (client S3FullClient) Upload(userId string, videoId string, fileType string) (string, error) {

	files, err := os.ReadDir("../temp") // Read files from the "temp" folder
	if err != nil {
		return "", fmt.Errorf("couldn't read files from temp folder: %w", err)
	}

	var videoBuffer []byte
	var found bool = false
	for _, file := range files {
		objectKey := file.Name()
		if objectKey == videoId {
			videoBuffer, err = os.ReadFile("../temp/" + objectKey)
			if err != nil {
				return "", fmt.Errorf("Error getting buffer for file %s: %v\n", objectKey, err)
			}
			found = true
			break
		}
	}
	if found == false {
		return "", fmt.Errorf("Could find the video")
	}
	fullObjectKey := userId + "/" + videoId + "." + fileType
	_, err = client.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(client.Data.BucketName),
		Key:    aws.String(fullObjectKey),
		Body:   bytes.NewReader(videoBuffer),
	})
	if err != nil {
		return "", fmt.Errorf("Couldn't upload buffer to %v:%v. Here's why: %v\n", client.Data.BucketName, fullObjectKey, err)
	}

	utils.DeleteFileFromTemp(videoId)
	//function that deletes the item in my folder temp (the var is files) the file with the name fileName
	return fullObjectKey, nil
}
func (client S3FullClient) DeleteItem(userId string, videoId string) error {
	fullObjectKey := userId + "/" + videoId
	_, err := client.S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(client.Data.BucketName),
		Key:    aws.String(fullObjectKey),
	})
	if err != nil {
		return fmt.Errorf("couldn't delete %v:%v. Here's why: %v", client.Data.BucketName, fullObjectKey, err)
	}

	return nil
}

// UNUSED UploadFile reads from a file and puts the data into an object in a bucket.
func (client S3FullClient) UploadBuffer(folderPath string, uuid string) error {

	files, err := os.ReadDir("../temp") // Read files from the "temp" folder
	if err != nil {
		return fmt.Errorf("couldn't read files from temp folder: %w", err)
	}

	for _, file := range files {
		objectKey := file.Name()
		buffer, err := os.ReadFile("../temp/" + objectKey)
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

func (client S3FullClient) ListBuckets() ([]types.Bucket, error) {
	result, err := client.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []types.Bucket
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	} else {
		buckets = result.Buckets
	}
	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}
	return buckets, err
}

// BucketExists checks whether a bucket exists in the current account.
func (client S3FullClient) BucketExists(bucketName string) (bool, error) {
	_, err := client.S3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	logs.I.Printf("RETORNA DE LA FUNCION HEADBUCKET %v", err)
	exists := true
	if err != nil {
		exists = false
		// var apiError
		// logs.I.Println("before the catch")
		// if errors.As(err, &apiError) {
		// 	logs.I.Println("i beleive this is the problem, incorrect apiError type")
		// 	switch apiError.(type) {
		// 	case *types.NotFound:
		// 		logs.I.Printf("Bucket %v is available.\n", bucketName)
		// 		exists = false
		// 		err = nil
		// 	default:
		// 		exists = false
		// 		logs.E.Printf("Either you don't have access to bucket %v or another error occurred. "+
		// 			"Here's what happened: %v\n", bucketName, err)
		// 	}
		// }
	} else {
		logs.I.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}

// CreateBucket creates a bucket with the specified name in the specified Region.
func (client S3FullClient) CreateBucket(name string, region string) error {
	_, err := client.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name),
		// CreateBucketConfiguration: &types.CreateBucketConfiguration{
		// 	LocationConstraint: types.BucketLocationConstraint(region),
		// },
	})
	if err != nil {
		logs.E.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
			name, region, err)
		return err
	}
	// Set the bucket policy after the bucket is successfully created
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Sid": "AllowPublicRead",
				"Effect": "Allow",
				"Principal": "*",
				"Action": "s3:GetObject",
				"Resource": "arn:aws:s3:::` + name + `/*"
			},
			{
				"Sid": "AllowPublicWrite",
				"Effect": "Allow",
				"Principal": "*",
				"Action": [
					"s3:PutObject",
					"s3:DeleteObject"
				],
				"Resource": "arn:aws:s3:::` + name + `/*"
			}
		]
	}`
	_, err = client.S3Client.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
		Bucket: aws.String(name),
		Policy: aws.String(policy),
	})
	if err != nil {
		logs.E.Printf("Failed to set bucket policy on %v. Here's why: %v\n", name, err)
		return err
	}

	logs.I.Printf("Bucket %v created and policy set successfully in Region %v\n", name, region)

	//cors
	// Set CORS configuration
	corsConfiguration := &s3.PutBucketCorsInput{
		Bucket: aws.String(name),
		CORSConfiguration: &types.CORSConfiguration{
			CORSRules: []types.CORSRule{
				{
					AllowedHeaders: []string{"*"},
					AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
					AllowedOrigins: []string{"http://localhost:8080", "http://localhost:7777", "http://35.215.20.21:7777", "http://35.215.20.21:8080"},
					MaxAgeSeconds:  aws.Int32(3000),
				},
			},
		},
	}
	_, err = client.S3Client.PutBucketCors(context.TODO(), corsConfiguration)
	if err != nil {
		logs.E.Printf("Failed to set CORS configuration on bucket %v. Here's why: %v\n", name, err)
		return err
	}

	logs.I.Printf("Bucket %v created, policy and CORS settings set successfully in Region %v\n", name, region)

	return nil

}
