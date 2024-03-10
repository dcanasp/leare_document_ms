package utils

import (
	"encoding/json"
	"global/globalTypes"
	"time"
)

func BrokerBodyToBytes(body globalTypes.BrokerEntry) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func BrokerBytesToBody(body []byte) (globalTypes.BrokerEntry, error) {
	var brokerEntry globalTypes.BrokerEntry
	if err := json.Unmarshal(body, &brokerEntry); err != nil {
		return globalTypes.BrokerEntry{}, err
	}
	return brokerEntry, nil
}

func BrokerToDynamo(brokerEntry globalTypes.BrokerEntry, filepath string) (globalTypes.DynamoEntry, error) {
	var dynamoEntry globalTypes.DynamoEntry
	dynamoEntry.FilePath = filepath
	dynamoEntry.UserId = brokerEntry.UserId
	dynamoEntry.VideoId = brokerEntry.VideoId
	dynamoEntry.FileName = brokerEntry.FileName
	dynamoEntry.FileType = brokerEntry.FileType
	dynamoEntry.Date = time.Now().Unix()

	return dynamoEntry, nil
}

func DynamoBodyToBytes(body globalTypes.DynamoEntry) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}
