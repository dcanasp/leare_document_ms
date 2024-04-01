package routes

import (
	"encoding/json"
	"global/globalTypes"
	logs "global/logging"
	configs "global/pkg/awsConfig"
	"global/utils"
	"net/http"
	"stream/pkg/routes/res"
	"sync"

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

		jsonResponse, err := utils.DynamoStringToJson(response)
		if err != nil {
			res.Response_S_Dynamo(w, false, globalTypes.DynamoEntry{})
		}
		res.Response_S_Dynamo(w, true, jsonResponse)

	}).Methods("GET")

	type batchBody struct {
		IDs []string `json:"ids"`
	}
	r.HandleFunc("/Batch", func(w http.ResponseWriter, r *http.Request) {
		//^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$
		var body batchBody
		errBody := json.NewDecoder(r.Body).Decode(&body)
		if errBody != nil {
			res.Response_Error(w, "Invalid request body")
			return
		}

		dynamoClient := *configs.DynamoClient
		responses := make(chan map[string]string, len(body.IDs))
		var wg sync.WaitGroup
		wg.Add(len(body.IDs))

		for _, id := range body.IDs {
			go func(id string) {
				defer wg.Done()
				response, err := dynamoClient.ReadEntry(id)
				if err != nil {
					logs.E.Printf("Dynamo entry could not be set %v", err)
					response = ""
				}
				responses <- map[string]string{id: response}
			}(id)
		}

		go func() {
			wg.Wait()
			close(responses)
		}()

		result := make(map[string]globalTypes.DynamoEntry)
		for resp := range responses {
			for id, response := range resp {
				result[id], _ = utils.DynamoStringToJson(response) // aca se convierte el response
			}
		}

		res.Response_S_Structure(w, true, result)

	}).Methods("POST")

	// type requestBody struct {
	// 	Content string `json:"content"`
	// }
	// r.HandleFunc("/addVideo/", func(w http.ResponseWriter, r *http.Request) {
	// 	var body requestBody
	// 	errBody := json.NewDecoder(r.Body).Decode(&body)
	// 	if errBody != nil {
	// 		http.Error(w, `{"success": false, "error": "Invalid request body"}`, http.StatusBadRequest)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")

	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Fprintf(w, `{"success":true}`)

	// }).Methods("POST")

}
