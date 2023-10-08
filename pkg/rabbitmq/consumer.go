package rabbitmq

import (
	"e-wallet/internal"
	"e-wallet/pkg/postgresql"
	"e-wallet/repositories"
	"fmt"
)

func RabbitMqConsumerInit() {

	transactionRepository := repositories.NewRepositoryTransactionImpl(postgresql.DB)
	queueNames := []string{"deposit", "cashout"}

	ch, err := GetRabbitMqChannel()
	if err != nil {
		fmt.Println("failed to open channel: ", err)
	}

	for _, queueName := range queueNames {
		go func(queueName string) {
			consumer := internal.NewConsumerImpl(ch, transactionRepository)
			consumer.ConsumeMessage(queueName)
		}(queueName)
	}

}
