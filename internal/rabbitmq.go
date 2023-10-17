package internal

import (
	"context"
	"m-banking/domain/repository"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker interface {
	PublishMessage(ctx context.Context, accountNumberID int, transactionCode string, amount float64, queueName string) error
	ConsumeMessage(queueName string)
}

type amqpChannel struct {
	ch                    *amqp.Channel
	transactionRepository repository.TransactionRepository
}
