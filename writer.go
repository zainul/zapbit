package zapbit

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func NewWriter(conf RabbitMQConfig, queue string) (*Writer, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.User, conf.Password, conf.Address, conf.Port))

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return nil, err
	}

	return &Writer{
		Queue:   q,
		Conn:    conn,
		Channel: ch,
	}, nil

}

func (w *Writer) Write(data []byte) (int, error) {
	if w.Channel == nil {
		return 0, errors.New("got nil channel")
	}

	err := w.Channel.Publish(
		ExchangeName, // exchange
		w.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	return 0, err
}

func (w *Writer) Close() error {
	if w.Channel == nil {
		return errors.New("got nil channel")
	}

	err := w.Channel.Close()
	err = w.Conn.Close()
	return err
}
