package command

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// AccountsService ...
type AccountsService interface {
	Create(context.Context, User) (*infra.Event, *infra.Error)
	Delete(context.Context, CPF, AccountNumber) (*infra.Event, *infra.Error)
	Credit(context.Context, CPF, AccountNumber, float64) (*infra.Event, *infra.Error)
	Debit(context.Context, CPF, AccountNumber, float64) (*infra.Event, *infra.Error)
}

// AccountsRepository ...
type AccountsRepository interface {
	Insert(context.Context, Account) (*infra.Event, *infra.Error)
	Remove(context.Context, Account) (*infra.Event, *infra.Error)
	ListByCPF(context.Context, CPF) ([]*Account, *infra.Error)
	Find(context.Context, CPF, AccountNumber) (*Account, *infra.Error)
	PutMoney(context.Context, Account, float64) (*infra.Event, *infra.Error)
	RemoveMoney(context.Context, Account, float64) (*infra.Event, *infra.Error)
}
