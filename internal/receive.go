package internal

import (
	"context"
	transactiondto "e-wallet/dto/mutasi"
	"e-wallet/models"
	"e-wallet/pkg/postgresql"
	"e-wallet/repositories"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumerImpl(ch *amqp.Channel) MessageBroker {
	return &amqpChannel{ch, repositories.NewRepositoryTransactionImpl(postgresql.DB)}
}

func (c *amqpChannel) ConsumeMessage(queueName string) {

	_, err := c.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf("failed to declare queue: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgs, err := c.ch.ConsumeWithContext(ctx,
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Printf("failed to consume queue: %s\n", err)
	}

	forever := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic: %v\n", r)
			}
		}()

		for msg := range msgs {
			var mutasi transactiondto.TransactionRequest
			if err := json.Unmarshal(msg.Body, &mutasi); err != nil {
				fmt.Printf("failed to parsing msg from RabbitMQ: %s\n", err)
			}

			transaction := models.Transaction{
				AccountNumberID: mutasi.AccountNumberID,
				TransactionCode: mutasi.TransactionCode,
				Amount:          mutasi.Amount,
				Date:            mutasi.Date,
			}

			response, err := c.transactionRepository.CreateTransactionReposity(transaction)
			if err != nil {
				fmt.Printf("failed save transaction to database: %s\n", err)
			}

			fmt.Printf("message from RabbitMQ: %+v\n", response)
		}

	}()

	<-forever
}