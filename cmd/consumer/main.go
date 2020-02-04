package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

func main() {
	config := kafka.NewConfig()
	config.ClientID = "consumer2"
	config.Version = kafka.V2_4_0_0

	client, err := kafka.NewClient([]string{
		"localhost:9092",
		"localhost:9093",
		"localhost:9094",
	}, config)
	if err != nil {
		fmt.Println("Deu ruim:", err)
		return
	}

	consumer, err := kafka.NewConsumerFromClient(client)
	if err != nil {
		fmt.Println("Deu ruim:", err)
		return
	}
	defer consumer.Close()

	pcons, err := consumer.ConsumePartition("accounts", 0, 0)
	if err != nil {
		fmt.Println("Deu ruimm:", err)
		return
	}

	for msg := range pcons.Messages() {
		fmt.Println(msg.Offset, string(msg.Value))
	}
}
