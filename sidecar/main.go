package main

import (
	"global"
	logs "global/logging"
	"global/pkg/awsConfig"
	"sidecar/broker"
)

func main() {
	err := global.SetENV()
	if err != nil {
		logs.E.Fatalf("Could not start the ENV %v", err)
	}
	awsConfig.Main()

	brokerClient, err := broker.Start()
	if err != nil {
		logs.E.Fatalf("Could not connect to mq on receive")
	}
	brokerClient.Connect()

}
