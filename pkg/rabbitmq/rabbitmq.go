package rabbitmq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var rabbitMQConnection *amqp.Connection

func GetRabbitMQChannel() (*amqp.Channel, error) {
	var err error

	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")
	RABBITMQ_USERNAME := os.Getenv("RABBITMQ_DEFAULT_USER")
	RABBITMQ_PASSWORD := os.Getenv("RABBITMQ_DEFAULT_PASS")

	if rabbitMQConnection == nil {

		// create a new RabbitMQ connection if it doesnt exist
		dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", RABBITMQ_USERNAME, RABBITMQ_PASSWORD, RABBITMQ_HOST, RABBITMQ_PORT)
		rabbitMQConnection, err = amqp.Dial(dsn)
		if err != nil {
			fmt.Printf("failed to connect to RabbitMQ: %v", err)
			return nil, err
		}

	}

	// create and return RabbitMQ channel
	ch, err := rabbitMQConnection.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}
