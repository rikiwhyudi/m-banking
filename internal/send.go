package internal

import (
	"context"
	transactiondto "e-wallet/dto/mutasi"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewPublisherImpl(ch *amqp.Channel) MessageBroker {
	return &amqpChannel{ch, nil}
}

func (p *amqpChannel) PublishMessage(accountNumber int, TransactionCode string, amount float64, queueName string) error {

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
		AccountNumberID: accountNumber,
		TransactionCode: TransactionCode,
		Amount:          amount,
		Date:            time.Now(),
	}

	mutasiMessage, err := json.Marshal(mutasi)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
