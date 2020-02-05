package query

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// AccountsProvider ...
type AccountsProvider interface {
	Insert(context.Context, domain.Account) *infra.Error
	Remove(context.Context, infra.ObjectID) *infra.Error
}

// TransactionsProvider ...
type TransactionsProvider interface {
	Deposit(context.Context, infra.ObjectID, float32) *infra.Error
	Withdraw(context.Context, infra.ObjectID, float32) *infra.Error
	Transfer(context.Context, infra.ObjectID, infra.ObjectID, float32) *infra.Error
	Loan(context.Context, infra.ObjectID, float32) *infra.Error
}
