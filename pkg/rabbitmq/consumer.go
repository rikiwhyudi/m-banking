package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage(queueName string) (<-chan amqp.Delivery, error) {

	_, err := RabbitMQChannel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := RabbitMQChannel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
