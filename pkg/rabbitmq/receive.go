package rabbitmq

import (
	"e-wallet/internal"
)

func RabbitMqConsumer() {

	queueNames := []string{"deposit", "cashout"}
	for _, queueName := range queueNames {
		go func(queueName string) {
			consumer := internal.NewConsumerImpl(RabbitMQChannel)
			consumer.ConsumeMessage(queueName)
		}(queueName)
	}

}
