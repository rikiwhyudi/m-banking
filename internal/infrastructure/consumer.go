package infrastructure

import (
	"encoding/json"
	"fmt"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumerImpl struct {
	ch                    *amqp.Channel
	transactionRepository ports.TransactionRepository
}

func NewConsumer(ch *amqp.Channel, transactionRepository ports.TransactionRepository) ports.Consumer {
	return &consumerImpl{ch, transactionRepository}
}

func (c *consumerImpl) ConsumeMessage(queueName string) {

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

	msgs, err := c.ch.Consume(
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
			var mutasi dto.TransactionRequest
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
