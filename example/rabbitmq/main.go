package main

import (
	"log"
	"time"

	"github.com/zainul/zapbit"
	"go.uber.org/zap"
)

func main() {

	writer, err := zapbit.NewWriter(zapbit.RabbitMQConfig{
		Address:  "localhost",
		Password: "root",
		User:     "root",
		Port:     5672,
	}, "logging_queue")

	if err != nil {
		log.Println("error connect to rabbit mq streamer", err)
	}

	logger := zap.New(writer.GetCore())

	defer logger.Sync()
	defer writer.Close()

	for {
		logger.Info("Hello From zap")
		time.Sleep(10 * time.Millisecond)
	}

}
