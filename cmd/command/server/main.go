package main

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
	"github.com/lcnascimento/event-sourcing-atm/infra/kafka"
	"github.com/lcnascimento/event-sourcing-atm/infra/log"

	account "github.com/lcnascimento/event-sourcing-atm/domain/command/account"
	server "github.com/lcnascimento/event-sourcing-atm/domain/command/server"
	transaction "github.com/lcnascimento/event-sourcing-atm/domain/command/transaction"
)

func main() {
	const opName infra.OpName = "main"

	ctx := context.Background()

	log := log.NewClient(log.ClientInput{
		Level: "Info",
		GoEnv: "test",
	})

	kafka, err := kafka.NewProducerService(kafka.ServiceInput{
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

	account, err := account.NewService(account.ServiceInput{
		Log: log,
	})
	if err != nil {
		errors.Log(log, err)
		return
	}

	transaction, err := transaction.NewService(transaction.ServiceInput{
		Log: log,
	})
	if err != nil {
		errors.Log(log, err)
		return
	}

	server, err := server.NewService(server.ServiceInput{
		Stream:       kafka,
		Accounts:     account,
		Transactions: transaction,
	})
	if err != nil {
		errors.Log(log, err)
		return
	}

	ch := server.Run(ctx)
	for err := range ch {
		errors.Log(log, err)
	}
}
