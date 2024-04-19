package database

import (
	"context"
	"fmt"
	logs "global/logging"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Gets tables names from dynamo
func (dynamoX *MyDynamoClient) ListTables() ([]string, error) {

	var tableLimit int32 = 5
	resp, err := dynamoX.Client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: &tableLimit,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list tables, %v", err)
	}

	return resp.TableNames, nil
}

// Adds a entry to dynamo
func (dynamoX *MyDynamoClient) AddEntry(videoId string, value string) error {
	// Prepare the input for the PutItem operation
	input := &dynamodb.PutItemInput{
		TableName: &dynamoX.TableName,
		Item: map[string]types.AttributeValue{
			"videoId": &types.AttributeValueMemberS{Value: videoId},
			"Value":   &types.AttributeValueMemberS{Value: value},
		},
	}

	// Execute the PutItem operation
	_, err := dynamoX.Client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to add entry on dynamo, %v", err)
	}

	logs.I.Println("Entry added successfully.")
	return nil
}

// Reads the content of an entry from dynamo
func (dynamoX *MyDynamoClient) ReadEntry(partitionKey string) (string, error) {
	// Prepare the input for the GetItem operation
	input := &dynamodb.GetItemInput{
		TableName: &dynamoX.TableName,
		Key: map[string]types.AttributeValue{
			"videoId": &types.AttributeValueMemberS{Value: partitionKey},
		},
	}

	// Execute the GetItem operation
	result, err := dynamoX.Client.GetItem(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to read entry, %v", err)
	}

	if valueAttr, ok := result.Item["Value"].(*types.AttributeValueMemberS); ok {
		//SE HACE UNA ASertion
		//todo Aprender a hacer eso
		// Now valueAttr.Value contains the clean string
		logs.I.Println("Entry:", valueAttr.Value)
		return valueAttr.Value, nil
	} else {
		// Handle the case where the value is not of the expected type or is missing
		return "", fmt.Errorf("Entry value is missing or not a string")

	}

}

func (dynamoX *MyDynamoClient) DeleteEntry(partitionKey string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &dynamoX.TableName,
		Key: map[string]types.AttributeValue{
			"videoId": &types.AttributeValueMemberS{Value: partitionKey},
		},
	}

	_, err := dynamoX.Client.DeleteItem(context.TODO(), input)
	if err != nil {
		logs.E.Printf("Got error calling DeleteItem: %s", err)
		return err
	}
	return nil
}
