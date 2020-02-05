package main

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
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
		GoEnv: "development",
	})

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
