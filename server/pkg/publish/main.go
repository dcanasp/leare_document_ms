package publish

import (
	"context"
	"encoding/json"
	"fmt"
	"global/globalTypes"
	logs "global/logging"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logs.E.Panicf("%s: %s", msg, err)
	}
}

func convertBodyToBytes(body globalTypes.BrokerEntry) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

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
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	binaryBody, err := convertBodyToBytes(body)
	if err != nil {
		failOnError(err, "could not convert to binary")
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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

// func ListenToQueue() {
// 	queueName := os.Getenv("queueName")
// 	MQCredentials := os.Getenv("rabbitMQCredentials")
// 	MQServerIp := os.Getenv("rabbitMQServerIP")

// 	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s@%s:5672", MQCredentials, MQServerIp))
// 	// amqp://${user}@${url}:5672
// 	if err != nil {
// 		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("Failed to open a channel: %s", err)
// 	}
// 	defer ch.Close()

// 	msgs, err := ch.Consume(
// 		queueName, // queue
// 		"",        // consumer
// 		true,      // auto-ack
// 		false,     // exclusive
// 		false,     // no-local
// 		false,     // no-wait
// 		nil,       // args
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to register a consumer: %s", err)
// 	}

// 	forever := make(chan bool)

// 	go func() {
// 		for d := range msgs {
// 			var msg QueueMessage
// 			if err := json.Unmarshal(d.Body, &msg); err != nil {
// 				log.Printf("Error decoding JSON: %s", err)
// 				continue
// 			}

// 			videoBuffer, err := base64.StdEncoding.DecodeString(msg.VideoData)
// 			if err != nil {
// 				log.Printf("Error decoding base64 video data: %s", err)
// 				continue
// 			}

// 			fmt.Printf("Received a message: %s with video buffer size: %d\n", msg.UUID, len(videoBuffer))
// 			//create files
// 			err = ffmpeg.Create(videoBuffer)
// 			if err != nil {
// 				log.Fatal("Video creation failed ", err)
// 			}
// 			//UPLOAD TO S3
// 			BucketName := os.Getenv("S3_BUCKET")
// 			s3CreationErr := bucketBasics.UploadBuffer(BucketName, msg.UserName, msg.UUID)
// 			if s3CreationErr != nil {
// 				log.Printf("Error uploading to S3: %v", err)
// 			}
// 			//Send to server that upload is completed
// 			ffmpeg.Delete()
// 			resp, err := MarkCompleted(msg.UUID)
// 			if err != nil {
// 				log.Fatalf("Error when marking as completed: %v", err)
// 			}
// 			fmt.Printf("MarkCompleted response: %v\n", resp)

// 		}
// 	}()

// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }
