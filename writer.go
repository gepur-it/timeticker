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
    Connection *amqp.Connection
    Channel *amqp.Channel
    QueueName string
}

func (wrtr *Writer) send(Time time.Time) {
    message := message(Time.Unix())
    body, err := json.Marshal(message)
    failOnError(err, "Failed encode message")
    err = wrtr.Channel.Publish(
        wrtr.QueueName,     // exchange
        wrtr.QueueName, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "application/json",
            Body:        body,
        },
    )
    failOnError(err, "Failed to publish a message")
}

func (wrtr *Writer) connect(Config Configuration) {
    wrtr.QueueName = Config.QueueName
    conn, err := amqp.Dial(Config.ConnectionString)
    failOnError(err, "Failed to connect to RabbitMQ")
    wrtr.Connection = conn
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    wrtr.Channel = ch
}

func (wrtr *Writer) run() {
    defer wrtr.Connection.Close()
    defer wrtr.Channel.Close()
    ticker := time.NewTicker(time.Second)
    for {
        Time := <- ticker.C
        wrtr.send(Time)
    }
}