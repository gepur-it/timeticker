package main

import (
    "encoding/json"
    "github.com/streadway/amqp"
	"strconv"
	"log"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

type Writer struct {
    outChan chan int64
}

func writer(outChan chan int64) *Writer {
  return &Writer{
      outChan: outChan,
  }
}

type EventParams struct {
   Time string `json:"time"`
}
type EventStruct struct {
   Name string `json:"name"`
   Parameters *EventParams `json:"parameters"`
}

type Message struct {
    MessageType string `json:"type"`
    Payload *EventStruct `json:"payload"`
}

func message(incTime int64) *Message {
    return &Message {
        MessageType: "event",
        Payload: &EventStruct {
            Name: "time",
            Parameters: &EventParams{
                Time: strconv.FormatInt(incTime, 10),
            },
        },
    }
}


func (wrtr *Writer) run() {
    conn, err := amqp.Dial("amqp://gepur_erp:gepur_erp@127.0.0.1:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    for {
        currentTime := <- wrtr.outChan
        message := message(currentTime)
        body, err := json.Marshal(message)
        if err != nil {
            panic(err)
        }
        err = ch.Publish(
            "cets",     // exchange
            "cets", // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing {
                ContentType: "application/json",
                Body:        body,
            },
        )
        failOnError(err, "Failed to publish a message")
    }
}