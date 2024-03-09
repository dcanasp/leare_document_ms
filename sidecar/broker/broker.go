package broker

import (
	"fmt"
	"global/pkg/awsConfig"
	"global/utils"
	"log"
	"os"

	logs "global/logging"

	amqp "github.com/rabbitmq/amqp091-go"
)

type BrokerClient struct {
	Credentials string
	ServerIp    string
}

func Start() (*BrokerClient, error) {
	var MQCredentials string = os.Getenv("rabbitMQCredentials")
	var MQServerIp string = os.Getenv("rabbitMQServerIP")
	brokerClient := BrokerClient{Credentials: MQCredentials, ServerIp: MQServerIp}
	return &brokerClient, nil
}

func (client *BrokerClient) Connect() {

	// amqp://${user}@${url}:5672
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s@%s:5672", client.Credentials, client.ServerIp))
	utils.ErrorFail(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.ErrorFail(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.ErrorFail(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.ErrorFail(err, "Failed to register a consumer")

	s3FullLink := os.Getenv("s3FullLink")
	dynamoClient := *awsConfig.DynamoClient
	s3Client := *awsConfig.S3Client
	var forever chan struct{}

	go func() {
		for d := range msgs {
			//TODO: vea si limpia el codigo.
			bodyBroker, err := utils.BrokerBytesToBody(d.Body)
			utils.Error(err, "No se pudo convertir al body")

			log.Printf("Received a message: %s", bodyBroker)

			filePath, err := s3Client.Upload(bodyBroker.UserId, bodyBroker.VideoId)
			if err != nil {
				logs.X.Print(bodyBroker.UserId + "," + bodyBroker.VideoId)
				continue
			}

			dynamoEntry, _ := utils.BrokerToDynamo(bodyBroker, s3FullLink+filePath) //also ads filePath
			dynamoBytes, err := utils.DynamoBodyToBytes(dynamoEntry)
			utils.Error(err, "no se pudo convertir el body de dynamo")
			dynamoClient.AddEntry(bodyBroker.VideoId, string(dynamoBytes))

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
