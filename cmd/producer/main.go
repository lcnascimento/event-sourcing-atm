package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

func main() {
	config := kafka.NewConfig()
	config.Version = kafka.V2_4_0_0
	config.Producer.RequiredAcks = kafka.RequiredAcks(1)

	client, err := kafka.NewClient([]string{
		"localhost:9092",
		"localhost:9093",
		"localhost:9094",
	}, config)
	if err != nil {
		fmt.Println("Deu ruim:", err)
		return
	}

	producer, err := kafka.NewAsyncProducerFromClient(client)
	if err != nil {
		fmt.Println("Deu ruim:", err)
		return
	}
	defer producer.Close()

	msg := &kafka.ProducerMessage{
		Topic: "accounts",
		Value: kafka.StringEncoder("Hello World"),
	}

	producer.Input() <- msg
	fmt.Println("Message sent!")
}
