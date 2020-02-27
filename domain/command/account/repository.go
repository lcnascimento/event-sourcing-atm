package account

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/domain/command"
	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// RepositoryInput ...
type RepositoryInput struct {
	Log          infra.LogProvider
	EventStore   infra.EventStoreProvider
	EventHandler domain.EventHandlerProvider
}

// Repository ...
type Repository struct {
	in RepositoryInput
}

// NewRepository ...
func NewRepository(in RepositoryInput) (*Repository, *infra.Error) {
	return &Repository{in: in}, nil
}

// Insert ...
func (r Repository) Insert(ctx context.Context, acc command.Account) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.repository.Insert"

	rowID := fmt.Sprintf("%s:AccountAggregate", acc.Owner.CPF)

	event, err := r.in.EventStore.Insert(ctx, infra.Event{
		RowID:       infra.EventRowID(rowID),
		AggregateID: infra.AggregateID(strconv.Itoa(acc.Number)),
		Type:        command.AccountCreatedEvent,
		Payload: infra.EventPayload{
			Data: command.AccountCreatedPayload{
				Account: acc,
			},
		},
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}

// ListByCPF ...
func (r Repository) ListByCPF(ctx context.Context, cpf command.CPF) ([]*command.Account, *infra.Error) {
	const opName infra.OpName = "account.repository.ListByCPF"

	rowID := infra.EventRowID(fmt.Sprintf("%s:AccountAggregate", cpf))

	events, err := r.in.EventStore.List(ctx, infra.ListEventsInput{
		RowID: &rowID,
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	aggMap := map[infra.AggregateID]*command.Account{}
	for _, event := range events {
		if aggMap[event.AggregateID] == nil {
			aggMap[event.AggregateID] = &command.Account{}
		}

		aggMap[event.AggregateID].Events = append(aggMap[event.AggregateID].Events, event)
	}

	accounts := []*command.Account{}
	for _, account := range aggMap {
		projection, err := r.in.EventHandler.Project(ctx, account, account.Events)
		if err != nil {
			return nil, errors.New(opName, err)
		}

		acc := projection.(command.Account)
		accounts = append(accounts, &acc)
	}

	return accounts, nil
}
