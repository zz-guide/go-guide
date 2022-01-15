package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

var topic = "bingo"
var URL = "pulsar://localhost:6650"

func main() {
	Consumption()
}

func Consumption() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: URL,
	})

	defer client.Close()
	if err != nil {
		fmt.Println("connect Fail ", err)
	} else {
		fmt.Println("connect Success ")
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "sub1",
		Type:             pulsar.Shared,
	})
	if err != nil {
		fmt.Println("connect Fail ", err)
	} else {
		fmt.Println("connect Success ")
	}

	defer consumer.Close()

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID().EntryID(), string(msg.Payload()))
		consumer.Ack(msg)
	}
}
