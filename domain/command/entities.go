package command

import (
	"context"
	"time"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// User ...
type User struct {
	CPF      CPF
	Name     string
	Email    string
	Birthday time.Time
}

// Account ...
type Account struct {
	Number  AccountNumber `json:"number"`
	Agency  AccountAgency `json:"agency"`
	Balance float64       `json:"balance"`
	Owner   User          `json:"owner"`
	Events  []infra.Event `json:"-"`
}

// Name ...
func (Account) Name() string {
	return "AccountAggregate"
}

// Apply ...
func (a *Account) Apply(ctx context.Context, e infra.Event) *infra.Error {
	const opName infra.OpName = "command.Account.Apply"
	defer func() { a.Events = append(a.Events, e) }()

	switch e.Type {
	case AccountCreatedEvent:
		payload := AccountCreatedPayload{}
		if err := e.Payload.Cast(&payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Agency = payload.Account.Agency
		a.Number = payload.Account.Number
		a.Balance = payload.Account.Balance
		a.Owner = payload.Account.Owner

		return nil
	case MoneyCreditedIntoAccountEvent:
		payload := MoneyCreditedIntoAccountPayload{}
		if err := e.Payload.Cast(&payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Balance += payload.Value

		return nil
	case MoneyDebitedFromAccountEvent:
		payload := MoneyDebitedFromAccountPayload{}
		if err := e.Payload.Cast(&payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Balance -= payload.Value

		return nil
	default:
		return errors.New(ctx, opName, domain.ErrApplyEventIntoAggregate, infra.KindNotFound, infra.SeverityWarning)
	}
}
