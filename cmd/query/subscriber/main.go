package main

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
	"github.com/lcnascimento/event-sourcing-atm/infra/kafka"
	"github.com/lcnascimento/event-sourcing-atm/infra/log"

	"github.com/lcnascimento/event-sourcing-atm/domain/query/subscriber"
)

func main() {
	const opName infra.OpName = "main"

	ctx := context.Background()

	log := log.NewClient(log.ClientInput{
		Level: "Info",
		GoEnv: "development",
	})

	kafka, err := kafka.NewConsumerService(kafka.ServiceInput{
		Log: log,
		Hosts: []string{
			"localhost:9092",
			"localhost:9093",
			"localhost:9094",
		},
	})
	if err != nil {
		errors.Log(log, err)
		return
	}
	defer kafka.CloseConnection()

	subscriber, err := subscriber.NewService(subscriber.ServiceInput{
		Log:    log,
		Stream: kafka,
	})
	if err != nil {
		errors.Log(log, err)
		return
	}

	ch := subscriber.Run(ctx)
	for err := range ch {
		errors.Log(log, err)
	}
}
