package main

import (
	"global"
	logs "global/logging"
	"global/pkg/awsConfig"

	"stream/pkg/routes"
)

func main() {
	//on the logging side
	//there is a init function on the logging package. That is executed before our main function
	//there are 3 loggers, on the logs package, use them accordingly
	//there is a global logger. if you use the normal log. it will go to the info and console
	//there is also a
	err := global.SetENV()
	if err != nil {
		logs.E.Fatalf("Could not start the ENV %v", err)
	}
	awsConfig.Main()

	routes.Main(3012)
	// go routes.Main(3012) // Start the server in a new goroutine
	// select {}

	// Block the main function from exiting
	/*
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

	*/

	//instance webservice

}
