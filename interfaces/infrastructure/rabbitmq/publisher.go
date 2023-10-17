package rabbitmq

import "context"

type Publisher interface {
	PublisherMessage(ctx context.Context, accountNumberID int, transactionCode string, amount float64, queueName string) error
}
