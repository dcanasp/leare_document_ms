package database

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MyDynamoClient struct {
	client    *dynamodb.Client
	TableName string
}

func Start(cfg aws.Config) (*MyDynamoClient, error) {
	// Using the Config value, create the DynamoDB client
	svc := MyDynamoClient{client: dynamodb.NewFromConfig(cfg)}
	return &svc, nil
}

func (dynamoX *MyDynamoClient) SetTable(tableName string) error {
	dynamoX.TableName = tableName
	return nil
}
