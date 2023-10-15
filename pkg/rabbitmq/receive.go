package rabbitmq

import (
	"m-banking/internal"
)

func RabbitMqConsumer() {

	queueNames := []string{"deposit", "cashout", "transfer"}
	for _, queueName := range queueNames {
		go func(queueName string) {
			consumer := internal.NewConsumerImpl(RabbitMQChannel)
			consumer.ConsumeMessage(queueName)
		}(queueName)
	}

}
