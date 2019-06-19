package main

import (
	"log"

	"github.com/zainul/zapbit"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	if err := ch.ExchangeDeclare(zapbit.ExchangeName, "topic", true, false, false, true, nil); err != nil {
		log.Println("exchange.declare destination: ", err)
	}

	q, err := ch.QueueDeclare(
		"logging_queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if err := ch.QueueBind(q.Name, "logging", zapbit.ExchangeName, false, nil); err != nil {
		log.Println("queue.bind source: ", err)
	}

	msgs, err := ch.Consume(
		q.Name,    // queue
		"logging", // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
