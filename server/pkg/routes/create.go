package routes

import (
	"global/globalTypes"
	logs "global/logging"
	"io"
	"net/http"
	"os"
	"stream/pkg/publish"
	"stream/pkg/routes/res"

	"github.com/gorilla/mux"
)

func Create(r *mux.Router) {

	//no estoy usando el body pero es este
	type requestBody struct {
		VideoId  string `json:"videoId"`
		Content  string `json:"content"`
		FileName string `json:"fileName"`
		DataType string `json:"dataType"`
		UserId   string `json:"userId"`
	}

	r.HandleFunc("/addVideo/", func(w http.ResponseWriter, r *http.Request) {

		MultiFormFile, _, err := r.FormFile("content")
		if err != nil {
			res.Response_Error(w, err.Error())
			return
		}
		defer MultiFormFile.Close()

		body := globalTypes.BrokerEntry{VideoId: r.FormValue("video_id"), FileName: r.FormValue("file_name"), FileType: r.FormValue("data_type"), UserId: r.FormValue("user_id")}
		// fileName := r.FormValue("file_name")

		// Create a new file on the server's filesystem
		filePath := "../temp/" + body.VideoId
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {

			res.Response_Error(w, err.Error())
			return
		}
		defer f.Close()

		// Copy the uploaded file to the new file
		_, err = io.Copy(f, MultiFormFile)
		if err != nil {
			res.Response_Error(w, err.Error())
			return
		}

		//aqui va el broker
		brokerClient, err := publish.Start()
		if err != nil {
			logs.E.Printf("%s: %s", "iniciarBrokerFailed", err)
			res.Response_Error(w, "Broker sending starting failed")
		}
		brokerClient.Connect(body)

		res.Response_Success(w)
		// dn := *awsConfig.DynamoClient
		// dn.AddEntry("prueba", "david")

	}).Methods("POST")
}
