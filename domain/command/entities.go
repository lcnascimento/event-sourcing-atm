package command

import (
	"context"
	"time"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// CPF ...
type CPF string

// User ...
type User struct {
	CPF      CPF
	Name     string
	Email    string
	Birthday time.Time
}

// Account ...
type Account struct {
	Number  int           `json:"number"`
	Agency  int           `json:"agency"`
	Balance float64       `json:"balance"`
	Owner   User          `json:"owner"`
	Events  []infra.Event `json:"-"`
}

// Name ...
func (Account) Name() string {
	return "AccountAggregate"
}

// Apply ...
func (a Account) Apply(ctx context.Context, e infra.Event) (domain.Aggregate, *infra.Error) {
	const opName infra.OpName = "command.Account.Apply"

	switch e.Type {
	case AccountCreatedEvent:
		return a, nil
	default:
		return nil, errors.New(ctx, opName, domain.ErrApplyEventIntoAggregate, infra.KindNotFound, infra.SeverityWarning)
	}
}
