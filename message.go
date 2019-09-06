package main

import (
    "strconv"
)

type EventParams struct {
   Time string `json:"time"`
}
type EventStruct struct {
   Name string `json:"name"`
   Parameters EventParams `json:"parameters"`
}

type Message struct {
    MessageType string `json:"type"`
    Payload EventStruct `json:"payload"`
}

func message(incTime int64) Message {
    return Message {
        MessageType: "event",
        Payload: EventStruct {
            Name: "time",
            Parameters: EventParams{
                Time: strconv.FormatInt(incTime, 10),
            },
        },
    }
}