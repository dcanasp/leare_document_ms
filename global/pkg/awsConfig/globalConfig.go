package awsConfig

import (
	"context"
	"fmt"
	logs "global/logging"
	"global/pkg/database"
	"global/pkg/fileStorage"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	S3Client     *fileStorage.S3FullClient
	DynamoClient *database.MyDynamoClient
)

func Main() {

	logs.I.Print("ENTRA A la configuracion de aws")
	//Configure aws
	cfg, err := Session()
	if err != nil {
		logs.E.Fatalf("Aws not configured %v", err)
	}

	//configure s3 client
	S3Client, err = fileStorage.SetS3(*cfg)
	//Todo estos son punteros
	if err != nil {
		logs.E.Fatalf("s3 could not be started %v", err)
	}

	//configure Dynamo client
	DynamoClient, err = database.Start(*cfg)
	if err != nil {
		logs.E.Fatalf("Dynamo could not be started %v", err)
	}

	tableName := "streams"
	err = createTableIfNotExists(DynamoClient, tableName)

	tablesNames, err := DynamoClient.ListTables()
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("No tables could be found %v", err)
	}
	err = DynamoClient.SetTable(tablesNames[0])
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("table name could not be set %v", err)
	}

}

func createTableIfNotExists(svc *database.MyDynamoClient, tableName string) error {
	// Check if the table already exists
	_, err := svc.Client.DescribeTable(context.Background(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err == nil {
		return nil // Table already exists
	}

	// Table does not exist, so create it
	_, err = svc.Client.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("videoId"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("videoId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
