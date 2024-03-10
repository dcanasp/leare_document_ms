package awsConfig

import (
	logs "global/logging"
	"global/pkg/database"
	"global/pkg/fileStorage"
)

var (
	S3Client     *fileStorage.S3FullClient
	DynamoClient *database.MyDynamoClient
)

func Main() {

	logs.I.Print("ENTRA A la configuracion de aws")
	//Configure aws
	cfg, err := Session()
	if err != nil {
		logs.E.Fatalf("Aws not configured %v", err)
	}

	//configure s3 client
	S3Client, err = fileStorage.SetS3(*cfg)
	//Todo estos son punteros
	if err != nil {
		logs.E.Fatalf("s3 could not be started %v", err)
	}

	//configure Dynamo client
	DynamoClient, err = database.Start(*cfg)
	if err != nil {
		logs.E.Fatalf("Dynamo could not be started %v", err)
	}

	tablesNames, err := DynamoClient.ListTables()
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("No tables could be found %v", err)
	}
	err = DynamoClient.SetTable(tablesNames[0])
	if err != nil || len(tablesNames) == 0 {
		logs.E.Fatalf("table name could not be set %v", err)
	}

}
