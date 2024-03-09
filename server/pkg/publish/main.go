package publish

import (
	"context"
	"fmt"
	"global/globalTypes"
	"global/utils"
	"log"
	"os"
	"time"

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

func (client *BrokerClient) Connect(body globalTypes.BrokerEntry) {

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	binaryBody, err := utils.BrokerBodyToBytes(body)
	if err != nil {
		utils.ErrorFail(err, "could not convert to binary")
	}
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(binaryBody),
		})
	utils.ErrorFail(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
