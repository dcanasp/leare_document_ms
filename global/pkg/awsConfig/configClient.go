package awsConfig

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func Session() (*aws.Config, error) {

	var region string = os.Getenv("AWS_REGION")
	var aws_access_key_id string = os.Getenv("aws_access_key_id")
	var aws_secret_access_key string = os.Getenv("aws_secret_access_key")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(aws_access_key_id, aws_secret_access_key, "")),
	)

	if err != nil {
		return nil, fmt.Errorf("aws connection failed %v", err)
	}

	return &cfg, nil
}
