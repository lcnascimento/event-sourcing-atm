package command

import (
	"context"
	"encoding/json"
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
		// TODO: Handle this logic during JSONUnmarshling of Event
		b, err := json.Marshal(e.Payload.Data)
		if err != nil {
			return errors.New(ctx, opName, err)
		}

		payload := AccountCreatedPayload{}
		if err := json.Unmarshal(b, &payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Agency = payload.Account.Agency
		a.Number = payload.Account.Number
		a.Balance = payload.Account.Balance
		a.Owner = payload.Account.Owner

		return nil
	case MoneyCreditedIntoAccountEvent:
		// TODO: Handle this logic during JSONUnmarshling of Event
		b, err := json.Marshal(e.Payload.Data)
		if err != nil {
			return errors.New(ctx, opName, err)
		}

		payload := MoneyCreditedIntoAccountPayload{}
		if err := json.Unmarshal(b, &payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Balance += payload.Value

		return nil
	case MoneyDebitedFromAccountEvent:
		// TODO: Handle this logic during JSONUnmarshling of Event
		b, err := json.Marshal(e.Payload.Data)
		if err != nil {
			return errors.New(ctx, opName, err)
		}

		payload := MoneyDebitedFromAccountPayload{}
		if err := json.Unmarshal(b, &payload); err != nil {
			return errors.New(ctx, opName, err)
		}

		a.Balance -= payload.Value

		return nil
	default:
		return errors.New(ctx, opName, domain.ErrApplyEventIntoAggregate, infra.KindNotFound, infra.SeverityWarning)
	}
}
