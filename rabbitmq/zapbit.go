package rabbitmq

import "github.com/streadway/amqp"

type Config struct {
	User     string
	Password string
	Address  string
	Port     int
}

type Writer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// ExchangeName is exchange name rabbit mq
const (
	ExchangeName    = "zapbit_logs"
)
