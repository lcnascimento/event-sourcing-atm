package domain

import (
	"context"
	"time"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// Event ...
type Event struct {
	RowID       infra.ObjectID
	AggregateID infra.ObjectID
	Timestamp   *time.Time
	Payload     interface{}
	Type        EventName
}

// Account ...
type Account struct {
	Number  int
	Agency  int
	Balance float64
}

// Name ...
func (Account) Name() string {
	return "AccountAggregate"
}

// Apply ...
func (a Account) Apply(ctx context.Context, e Event) (*Account, *infra.Error) {
	const opName infra.OpName = "command.Account.Apply"

	switch e.Type {
	default:
		return nil, errors.New(ctx, opName, ErrApplyEventIntoAggregate, infra.KindNotFound, infra.SeverityWarning)
	}
}
