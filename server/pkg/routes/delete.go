package routes

import (
	"encoding/json"
	logs "global/logging"
	"global/pkg/awsConfig"
	"net/http"
	"stream/pkg/routes/res"
	"sync"

	"github.com/gorilla/mux"
)

func Update(r *mux.Router) {

	r.HandleFunc("/addVideo/", func(w http.ResponseWriter, r *http.Request) {
		//nunca se llama porque un update es lo mismo que un create
	}).Methods("POST")
}

func Delete(r *mux.Router) {

	//CreateFile

	r.HandleFunc("/{userId}/{videoId}", func(w http.ResponseWriter, r *http.Request) {
		//como delete de dynamo y s3
		vars := mux.Vars(r)
		userId := vars["userId"]
		videoId := vars["videoId"]

		dynamoClient := *awsConfig.DynamoClient
		s3Client := *awsConfig.S3Client
		err := dynamoClient.DeleteEntry(videoId)
		if err != nil {
			logs.E.Printf("Dynamo entry could not be deleted set %v", err)
			res.Response_Error(w, "Parameter not found")
			return
		}
		err = s3Client.DeleteItem(userId, videoId)
		if err != nil {
			logs.E.Print(err)
			res.Response_Error(w, "could not delete item on s3")
			return
		}
		res.Response_Success(w)

	}).Methods("DELETE")

	type batchDeleteBody struct {
		UserId   string   `json:"userIds"`
		VideoIds []string `json:"videoIds"`
	}
	r.HandleFunc("/Batch/", func(w http.ResponseWriter, r *http.Request) {
		//como delete de dynamo y s3

		var body batchDeleteBody
		errBody := json.NewDecoder(r.Body).Decode(&body)
		if errBody != nil {
			res.Response_Error(w, "Invalid request body")
			return
		}
		dynamoClient := *awsConfig.DynamoClient
		s3Client := *awsConfig.S3Client

		deleteSuccess := make(chan bool)
		var deleteWg sync.WaitGroup

		// Deleting items from DynamoDB and S3 concurrently
		for _, videoId := range body.VideoIds {
			deleteWg.Add(1)
			go func(userId, videoId string) {
				defer deleteWg.Done()

				// Delete from DynamoDB
				if err := dynamoClient.DeleteEntry(videoId); err != nil {
					logs.E.Printf("Dynamo entry could not be deleted: %v", err)
					return
				}

				// Delete from S3
				if err := s3Client.DeleteItem(body.UserId, videoId); err != nil {
					logs.E.Printf("Item could not be deleted from S3: %v", err)
					return
				}

				deleteSuccess <- true
			}(body.UserId, videoId)
		}
		// Wait for all deletion operations to finish
		go func() {
			deleteWg.Wait()
			close(deleteSuccess)
		}()

		for range deleteSuccess {
			//si quiero verificar que llegen
		}

		res.Response_Success(w)

	}).Methods("DELETE")
}
