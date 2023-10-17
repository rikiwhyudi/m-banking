package internal

import (
	"context"
	"encoding/json"
	transactiondto "m-banking/dto/mutasi"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewPublisherImpl(ch *amqp.Channel) MessageBroker {
	return &amqpChannel{ch, nil}
}

func (p *amqpChannel) PublishMessage(ctx context.Context, accountNumberID int, transactionCode string, amount float64, queueName string) error {

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

	mutasi := transactiondto.TransactionRequest{
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
