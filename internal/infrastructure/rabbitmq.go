package infrastructure

import (
	"m-banking/internal/adapters/repository"
	"m-banking/pkg/database/sql"
	"m-banking/pkg/rabbitmq"
)

func RabbitMqConsumer() {

	queueNames := []string{"deposit", "cashout", "transfer"}
	for _, queueName := range queueNames {
		go func(queueName string) {
			consumer := NewConsumer(rabbitmq.RabbitMQChannel, repository.NewTransactionRepository(sql.DB))
			consumer.ConsumeMessage(queueName)
		}(queueName)
	}

}
