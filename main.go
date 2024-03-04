package main

import (
	"log"
	"os"
	logs "stream/pkg/utils/logging"

	"stream/pkg/awsConfig"
	"stream/pkg/database"
	"stream/pkg/fileStorage"
	"stream/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	//on the logging side
	//there is a init function on the logging package. That is executed before our main function
	//there are 3 loggers, on the logs package, use them accordingly
	//there is a global logger. if you use the normal log. it will go to the info and console
	//there is also a
	err := SetENV()
	if err != nil {
		logs.E.Fatalf("Could not start the ENV %v", err)
	}

	//Configure aws
	_, err = awsConfig.Session()
	if err != nil {
		logs.E.Fatalf("Aws not configured %v", err)
	}

	//configure s3 client
	s3Client, err := fileStorage.SetS3()
	if err != nil {
		logs.E.Fatalf("s3 could not be started %v", err)
	}
	err = s3Client.UploadBuffer("a", "vafbs")
	if err != nil {
		logs.E.Printf("s3 could not be started %v", err)
	}

	//configure Dynamo client
	dynamoClient, err := database.Start()
	if err != nil {
		logs.E.Fatalf("Dynamo could not be started %v", err)
	}

	//todo: other file
	tablesNames, err := dynamoClient.ListTables()
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("No tables could be found %v", err)
	}
	err = dynamoClient.SetTable(tablesNames[0])
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("table name could not be set %v", err)
	}
	err = dynamoClient.AddEntry()
	if err != nil || len(tablesNames) == 0 {
		logs.E.Printf("Dynamo entry could not be set %v", err)
	}
	err = dynamoClient.ReadEntry("test")
	if err != nil || len(tablesNames) == 0 {
		logs.E.Printf("Dynamo entry not found %v", err)
	}
	// //instance db
	// err := db.InitDB()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize the database: %v", err)
	// }
	// //instance webservice
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	s := r.PathPrefix("/stream").Subrouter()
	// k := r.PathPrefix("/comments").Subrouter()
	// l := r.PathPrefix("/subscriptions").Subrouter()
	// m := r.PathPrefix("/videos").Subrouter()
	// // "/products/"
	// // s.HandleFunc("/", )
	routes.Stream(s)
	// routes.Comments(k)
	// routes.Subscriptions(l)
	// routes.Videos(m)

	// http.ListenAndServe(":3012", r)
}

func SetupGlobalLogger() {
	// Ensure the directory exists (MkdirAll is no-op if directory already exists)
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open or create the log file for appending, create it if it doesn't exist
	logFile, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set the global log output to the file
	log.SetOutput(logFile)

	// Optional: Set the log to also output the date and time
	log.SetFlags(log.LstdFlags)
}
