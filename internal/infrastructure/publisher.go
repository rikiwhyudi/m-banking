package infrastructure

import (
	"context"
	"encoding/json"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type publisherImpl struct {
	ch *amqp.Channel
}

func NewPublisher(ch *amqp.Channel) ports.Publisher {
	return &publisherImpl{ch}
}

func (p *publisherImpl) PublisherMessage(ctx context.Context, accountNumberID int, transactionCode string, amount float64, queueName string) error {

	queue, err := p.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	mutasi := dto.TransactionRequest{
		AccountNumberID: accountNumberID,
		TransactionCode: transactionCode,
		Amount:          amount,
		Date:            time.Now(),
	}

	mutasiMessage, err := json.Marshal(mutasi)
	if err != nil {
		return err
	}

	err = p.ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json;",
			Body:        mutasiMessage,
		},
	)

	if err != nil {
		return err
	}

	return err
}
