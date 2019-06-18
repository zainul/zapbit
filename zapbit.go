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

const ExchangeName = "zapbit_logs"