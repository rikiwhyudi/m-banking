package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQChannel *amqp.Channel

func RabbitMqInit() {
	var connection *amqp.Connection
	var err error

	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")
	RABBITMQ_DEFAULT_USER := os.Getenv("RABBITMQ_DEFAULT_USER")
	RABBITMQ_DEFAULT_PASS := os.Getenv("RABBITMQ_DEFAULT_PASS")
	RABBITMQ_DEFAULT_VHOST := os.Getenv("RABBITMQ_DEFAULT_VHOST")

	if connection == nil {
		dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", RABBITMQ_DEFAULT_USER, RABBITMQ_DEFAULT_PASS, RABBITMQ_HOST, RABBITMQ_PORT, RABBITMQ_DEFAULT_VHOST)
		connection, err = amqp.Dial(dsn)
		if err != nil {
			panic(err)
		}

		fmt.Println("connected to RabbitMQ Management...")
	}

	RabbitMQChannel, err = connection.Channel()
	if err != nil {
		fmt.Printf("failed to get RabbitMQ channel: %v", err)
	}

}
