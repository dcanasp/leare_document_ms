package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Stream(r *mux.Router) {

	r.HandleFunc("/findVideo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		videoId, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"success" : false,"err" : "not allowed string for comments}`)

			return
		}

		// var responseComments []dto.Comment
		// for _, c := range comments {
		// 	responseComments = append(responseComments, dto.Comment{
		// 		Id:             c.IdComments,
		// 		Username:       c.Username,
		// 		Date:           c.Date.Time.Format("2006-01-02"), // Format date as needed
		// 		UserProfilePic: c.Profile_picture_url.String,
		// 		Content:        c.Content,
		// 	})
		// }

		// response := dto.CommentResponse{
		// 	Success:  true,
		// 	Comments: responseComments,
		// }
		response := videoId
		w.Header().Set("Content-Type", "application/json")
		prettyResult := json.NewEncoder(w)
		errJson := prettyResult.Encode(response)
		if errJson != nil {
			log.Printf("failed json")
		}

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
