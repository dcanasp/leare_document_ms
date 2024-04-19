package database

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MyDynamoClient struct {
	Client    *dynamodb.Client
	TableName string
}

func Start(cfg aws.Config) (*MyDynamoClient, error) {
	// Using the Config value, create the DynamoDB client
	var dynamoEndpoint string = os.Getenv("DYNAMO_ENDPOINT")
	svc := MyDynamoClient{Client: dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(dynamoEndpoint)
	})}
	return &svc, nil
}

func (dynamoX *MyDynamoClient) SetTable(tableName string) error {
	dynamoX.TableName = tableName
	return nil
}
