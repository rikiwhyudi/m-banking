package rabbitmq

type Consumer interface {
	ConsumeMessage(queueName string)
}
