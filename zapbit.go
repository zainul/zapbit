package zapbit

import "github.com/streadway/amqp"

type RabbitMQConfig struct {
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
	KafkaMessanger  = 1
	RabbitMessanger = 2
	ExchangeName    = "zapbit_logs"
)
