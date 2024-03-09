package database

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MyDynamoClient struct {
	client    *dynamodb.Client
	TableName string
}

func Start() (*MyDynamoClient, error) {

	var region string = os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		_ = fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := MyDynamoClient{client: dynamodb.NewFromConfig(cfg)}
	return &svc, nil
}

func (dynamoX *MyDynamoClient) SetTable(tableName string) error {
	dynamoX.TableName = tableName
	return nil
}
