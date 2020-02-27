package command

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// AccountsService ...
type AccountsService interface {
	Create(context.Context, User) (*infra.Event, *infra.Error)
}

// AccountsRepository ...
type AccountsRepository interface {
	Insert(context.Context, Account) (*infra.Event, *infra.Error)
	ListByCPF(context.Context, CPF) ([]*Account, *infra.Error)
}
