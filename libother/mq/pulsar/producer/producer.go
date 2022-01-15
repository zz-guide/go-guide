package main

import (
	"context"
	"fmt"
	"strconv"
)

import "github.com/apache/pulsar-client-go/pulsar"

var topic = "bingo"
var URL = "pulsar://localhost:6650"

func main() {
	SendMessage()
}

func SendMessage() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: URL,
	})

	defer client.Close()
	if err != nil {
		fmt.Println("connect Fail ", err)
	} else {
		fmt.Println("connect Success ")
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	defer producer.Close()

	for i := 0; i < 1000000; i++ {
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte("{'name':'许磊','id':'" + strconv.Itoa(i) + "','age':12}"),
		})
	}

	if err != nil {
		fmt.Println("Failed to publish message", err)
	} else {
		fmt.Println("Published message")
	}
}
