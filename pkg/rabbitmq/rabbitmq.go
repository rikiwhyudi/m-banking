package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RMQ *amqp.Connection

func RabbitMqInit() {
	var err error

	RABBITMQ_HOST := os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT := os.Getenv("RABBITMQ_PORT")
	RABBITMQ_DEFAULT_USER := os.Getenv("RABBITMQ_DEFAULT_USER")
	RABBITMQ_DEFAULT_PASS := os.Getenv("RABBITMQ_DEFAULT_PASS")
	RABBITMQ_DEFAULT_VHOST := os.Getenv("RABBITMQ_DEFAULT_VHOST")

	if RMQ == nil {
		dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", RABBITMQ_DEFAULT_USER, RABBITMQ_DEFAULT_PASS, RABBITMQ_HOST, RABBITMQ_PORT, RABBITMQ_DEFAULT_VHOST)
		RMQ, err = amqp.Dial(dsn)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("connected to RabbitMQ Management...")

}

func GetRabbitMqChannel() (*amqp.Channel, error) {
	ch, err := RMQ.Channel()
	if err != nil {
		return nil, err
	}

	return ch, err
}
