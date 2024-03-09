package routes

import (
	logs "global/logging"
	"global/pkg/awsConfig"
	"net/http"
	"stream/pkg/routes/res"

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
}
