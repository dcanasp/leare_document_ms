package routes

import (
	"encoding/json"
	"fmt"
	logs "global/logging"
	"net/http"
	configs "stream/pkg/awsConfig"
	"stream/pkg/routes/res"

	"github.com/gorilla/mux"
)

func Get(r *mux.Router) {
	//GetFile
	//CreateFile
	//editFile
	//Delete file

	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		//^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$

		vars := mux.Vars(r)

		videoId := vars["id"]

		logs.I.Print(videoId)

		dynamoClient := *configs.DynamoClient
		response, err := dynamoClient.ReadEntry(videoId)
		if err != nil {
			logs.E.Printf("Dynamo entry could not be set %v", err)
			res.Response_Error(w, "Parameter not found")
			return
		}

		res.Response_S_V(w, true, response)

	}).Methods("GET")

	type requestBody struct {
		Content string `json:"content"`
	}
	r.HandleFunc("/addVideo/", func(w http.ResponseWriter, r *http.Request) {
		var body requestBody
		errBody := json.NewDecoder(r.Body).Decode(&body)
		if errBody != nil {
			http.Error(w, `{"success": false, "error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"success":true}`)

	}).Methods("POST")
}
