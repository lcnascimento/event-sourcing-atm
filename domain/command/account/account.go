package account

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"

	"github.com/lcnascimento/event-sourcing-atm/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Log        infra.LogProvider
	EventStore domain.EventStoreProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}

// Insert ...
func (s Service) Insert(ctx context.Context, acc domain.Account) *infra.Error {
	const opName infra.OpName = "command.account.Insert"

	s.in.EventStore.Insert(ctx, domain.Event{
		RowID:   infra.ObjectID("userid:AccountAggregate"),
		Type:    domain.AccountCreatedEvent,
		Payload: domain.AccountCreatedPayload{},
	})

	return nil
}

// Remove ...
func (s Service) Remove(ctx context.Context, accID infra.ObjectID) *infra.Error {
	const opName infra.OpName = "command.account.Remove"

	s.in.Log.Infof(ctx, opName, "Removing account %s", accID)
	return nil
}
