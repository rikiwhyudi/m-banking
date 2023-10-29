package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(ctx context.Context, queueName string, message interface{}) error {
	queue, err := RabbitMQChannel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = RabbitMQChannel.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json;",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
