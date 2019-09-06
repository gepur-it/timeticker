package main

import (
    "os"
    "fmt"
    "encoding/json"
)

type Configuration struct {
    ConnectionString string `json:"rabbit_connection_string"`
    QueueName string `json:"rabbit_cets_queue"`
}

func main() {
    file, Err := os.Open("config.json")
    if Err != nil {
        fmt.Println("error:", Err)
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }
//      writer := writer(configuration)
    writer := &Writer{}
    writer.connect(configuration)
    go writer.run()
    for {
    }
}