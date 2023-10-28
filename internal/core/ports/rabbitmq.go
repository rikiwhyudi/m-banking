package ports

import "context"

type Publisher interface {
	PublisherMessage(ctx context.Context, accountNumberID int, transactionCode string, amount float64, queueName string) error
}

type Consumer interface {
	ConsumeMessage(queueName string)
}
