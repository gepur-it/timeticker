package main

import (
    "encoding/json"
    "github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

type Writer struct {
    configuration Configuration
}

func connect(configuration Configuration) (*amqp.Connection, *amqp.Channel) {
      conn, err := amqp.Dial(configuration.ConnectionString)
      failOnError(err, "Failed to connect to RabbitMQ")
      ch, err := conn.Channel()
      failOnError(err, "Failed to open a channel")
      return conn, ch
}

func (wrtr *Writer) send(Time time.Time) {
    conn, ch := connect(wrtr.configuration)
    defer conn.Close()
    defer ch.Close()

    message := message(Time.Unix())
    body, err := json.Marshal(message)
    failOnError(err, "Failed encode message")
    err = ch.Publish(
        wrtr.configuration.QueueName,     // exchange
        wrtr.configuration.QueueName, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "application/json",
            Body:        body,
        },
    )
    failOnError(err, "Failed to publish a message")
}

func (wrtr *Writer) run() {
    ticker := time.NewTicker(time.Second)
    for {
        Time := <- ticker.C
        wrtr.send(Time)
    }
}