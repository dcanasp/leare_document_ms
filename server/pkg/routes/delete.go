package routes

import (
	logs "global/logging"
	"net/http"
	"stream/pkg/awsConfig"
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

	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		//como delete de dynamo y s3
		vars := mux.Vars(r)
		videoId := vars["id"]

		dynamoClient := *awsConfig.DynamoClient
		err := dynamoClient.DeleteEntry(videoId)
		if err != nil {
			logs.E.Printf("Dynamo entry could not be set %v", err)
			res.Response_Error(w, "Parameter not found")
			return
		}

		res.Response_Success(w)

	}).Methods("DELETE")
}
