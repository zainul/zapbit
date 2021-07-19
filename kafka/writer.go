package kafka

import (
	"context"
	"time"

	"github.com/zainul/ark/kafka"
)

type Writer struct {
	kafkaEng kafka.Kafka
}

type Config struct {
	URL string
}

func NewWriter(cfg Config) (*Writer, error) {
	kk , err := kafka.GetKafka(kafka.Config{
		Brokers:      nil,
		GroupID:      "zapbit-log",
		Topic:        "zapbit-log",
		URL:          cfg.URL,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	return &Writer{
		kafkaEng: kk,
	}, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	err := w.kafkaEng.WriteMessage(context.Background(), "zapbit-kafka", string(data))

	if err != nil {
		return -1, err
	}

	return 0, err
}

func (w *Writer) Close() error {
	return nil
}