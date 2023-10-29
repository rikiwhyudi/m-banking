package infrastructure

import (
	"encoding/json"
	"fmt"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/dto"
	"m-banking/pkg/database/sql"
	"m-banking/pkg/rabbitmq"
)

func RabbitMqConsumer() {

	repo := repository.NewTransactionRepository(sql.DB)
	queueNames := []string{"deposit", "cashout", "transfer"}

	for _, queueName := range queueNames {
		go func(queueName string) {
			msgs, err := rabbitmq.ConsumeMessage(queueName)
			if err != nil {
				fmt.Printf("failed to get queue: %s\n", err)
				return
			}

			forever := make(chan bool)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("panic: %v\n", r)
					}
				}()

				for msg := range msgs {
					var mutasi dto.TransactionRequest
					if err := json.Unmarshal(msg.Body, &mutasi); err != nil {
						fmt.Printf("failed to parse msg from RabbitMQ: %s\n", err)
						continue
					}

					transaction := models.Transaction{
						AccountNumberID: mutasi.AccountNumberID,
						TransactionCode: mutasi.TransactionCode,
						Amount:          mutasi.Amount,
						Date:            mutasi.Date,
					}

					response, err := repo.CreateTransactionReposity(transaction)
					if err != nil {
						fmt.Printf("failed to save transaction to database: %s\n", err)
						continue
					}

					fmt.Printf("message from RabbitMQ: %+v\n", response)
					msg.Ack(false)
				}

			}()

			<-forever

		}(queueName)
	}
}
