package internal

import (
	"m-banking/repositories"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker interface {
	PublishMessage(accountNumberID int, transactionCode string, amount float64, queueName string) error
	ConsumeMessage(queueName string)
}

type amqpChannel struct {
	ch                    *amqp.Channel
	transactionRepository repositories.TransactionRepository
}
