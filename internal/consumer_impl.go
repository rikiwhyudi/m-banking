package internal

import (
	transactiondto "e-wallet/dto/mutasi"
	"e-wallet/models"
	"e-wallet/repositories"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	ch                    *amqp.Channel
	transactionRepository repositories.TransactionRepository
}

func NewConsumerImpl(ch *amqp.Channel, transactionRepository repositories.TransactionRepository) Consumer {
	return &consumer{ch, transactionRepository}
}

func (c *consumer) ConsumeMessage(queueName string) {
	queue, err := c.ch.QueueDeclare(
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
		queue.Name,
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
		for msg := range msgs {
			go func(msg amqp.Delivery) {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("panic: %v\n", r)
					}

				}()

				var mutasi transactiondto.TransactionRequest
				if err := json.Unmarshal(msg.Body, &mutasi); err != nil {
					fmt.Printf("failed parsing msg from RabbitMQ: %s\n", err)
				}

				transaction := models.Transaction{
					AccountNumberID: mutasi.AccountNumberID,
					TransactionCode: mutasi.TransactionCode,
					Amount:          mutasi.Amount,
					Date:            mutasi.Date,
				}

				_, err := c.transactionRepository.CreateTransactionReposity(transaction)
				if err != nil {
					fmt.Printf("failed save transaction to database: %s\n", err)
				}

				fmt.Printf("message from RabbitMQ: %v\n", mutasi)

			}(msg)
		}

	}()

	<-forever
}
