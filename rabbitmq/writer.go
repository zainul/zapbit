package rabbitmq

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewWriter is get initial setup of rabbit mq
func NewWriter(conf Config, queue string) (*Writer, error) {

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
		true,  // no-wait
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

// Write is interface of log write
func (w *Writer) Write(data []byte) (int, error) {
	if w.Channel == nil {
		return 0, errors.New("got nil channel")
	}

	err := w.Channel.Publish(
		ExchangeName, // exchange
		"logging",    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(data),
			DeliveryMode: amqp.Persistent,
		})

	if err != nil {
		return -1, err
	}

	return 0, err
}

// GetCore is get zapcore of log engine
func (w *Writer) GetCore() zapcore.Core {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	topicLog := zapcore.AddSync(w)
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	return zapcore.NewTee(
		zapcore.NewCore(jsonEnc, topicLog, highPriority),
		zapcore.NewCore(jsonEnc, topicLog, lowPriority),
	)
}

// Close is use for closing the channel mq
func (w *Writer) Close() error {
	if w.Channel == nil {
		return errors.New("got nil channel")
	}

	err := w.Channel.Close()
	err = w.Conn.Close()
	return err
}
