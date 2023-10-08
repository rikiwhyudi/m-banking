package internal

type Consumer interface {
	ConsumeMessage(queueName string)
}
