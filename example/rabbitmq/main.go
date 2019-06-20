package main

import (
	"log"
	"os"

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
		hostName, _ := os.Hostname()

		logger.Info(
			"logging",
			zap.String("request_id", "somerequestid"),
			zap.String("hostname", hostName),
			zap.String("service", "someservicename"),
			zap.ByteString("request_body", []byte("12 12 121")),
			zap.ByteString("response_body", []byte("asd")),
		)
		// time.Sleep(10 * time.Millisecond)
	}

}
