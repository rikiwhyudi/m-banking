package rabbitmq

import (
	"m-banking/pkg/rabbitmq"
)

func RabbitMqConsumer() {

	queueNames := []string{"deposit", "cashout", "transfer"}
	for _, queueName := range queueNames {
		go func(queueName string) {
			consumer := NewConsumerImpl(rabbitmq.RabbitMQChannel)
			consumer.ConsumeMessage(queueName)
		}(queueName)
	}

}
