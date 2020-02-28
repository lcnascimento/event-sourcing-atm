package account

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lcnascimento/event-sourcing-atm/domain/command"
	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// RepositoryInput ...
type RepositoryInput struct {
	Log        infra.LogProvider
	EventStore infra.EventStoreProvider
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
		AggregateID: infra.AggregateID(strconv.Itoa(int(acc.Number))),
		Type:        command.AccountCreatedEvent,
		Payload: infra.EventPayload{
			Data: command.AccountCreatedPayload{
				Account: &acc,
			},
		},
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}

// Remove ...
func (r Repository) Remove(ctx context.Context, acc command.Account) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.repository.Remove"

	rowID := fmt.Sprintf("%s:AccountAggregate", acc.Owner.CPF)

	event, err := r.in.EventStore.Insert(ctx, infra.Event{
		RowID:       infra.EventRowID(rowID),
		AggregateID: infra.AggregateID(strconv.Itoa(int(acc.Number))),
		Type:        command.AccountDeletedEvent,
		Payload: infra.EventPayload{
			Data: command.AccountDeletedPayload{
				AccountOwnerCPF: acc.Owner.CPF,
				AccountNumber:   acc.Number,
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
		account := aggMap[event.AggregateID]
		if account == nil && event.Type == command.AccountCreatedEvent {
			account = &command.Account{}
			aggMap[event.AggregateID] = account
		}

		if account == nil {
			return nil, errors.New(ctx, opName, command.ErrUnexpectedEventInAggregateState, infra.SeverityCritical)
		}

		if event.Type == command.AccountDeletedEvent {
			aggMap[event.AggregateID] = nil
			continue
		}

		if err := account.Apply(ctx, event); err != nil && err.Severity != infra.SeverityWarning {
			return nil, errors.New(opName, err)
		}
	}

	accounts := []*command.Account{}
	for _, account := range aggMap {
		if account != nil {
			accounts = append(accounts, account)
		}
	}

	return accounts, nil
}

// Find ...
func (r Repository) Find(ctx context.Context, cpf command.CPF, accNum command.AccountNumber) (*command.Account, *infra.Error) {
	const opName infra.OpName = "account.repository.Find"

	rowID := infra.EventRowID(fmt.Sprintf("%s:AccountAggregate", cpf))
	aggID := infra.AggregateID(strconv.Itoa(int(accNum)))

	events, err := r.in.EventStore.List(ctx, infra.ListEventsInput{
		RowID:       &rowID,
		AggregateID: &aggID,
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	if len(events) == 0 {
		return nil, errors.New(ctx, opName, infra.ErrEntityNotFound, infra.KindNotFound)
	}

	account := &command.Account{}
	for _, event := range events {
		if event.Type == command.AccountDeletedEvent {
			return nil, errors.New(ctx, opName, command.ErrAccountDeleted, infra.KindBadRequest)
		}

		if err := account.Apply(ctx, event); err != nil && err.Severity != infra.SeverityWarning {
			return nil, errors.New(opName, err)
		}
	}

	return account, nil
}

// PutMoney ...
func (r Repository) PutMoney(ctx context.Context, acc command.Account, value float64) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.repository.PutMoney"

	rowID := infra.EventRowID(fmt.Sprintf("%s:AccountAggregate", acc.Owner.CPF))

	event, err := r.in.EventStore.Insert(ctx, infra.Event{
		RowID:       infra.EventRowID(rowID),
		AggregateID: infra.AggregateID(strconv.Itoa(int(acc.Number))),
		Type:        command.MoneyCreditedIntoAccountEvent,
		Payload: infra.EventPayload{
			Data: command.MoneyCreditedIntoAccountPayload{
				AccountOwnerCPF: acc.Owner.CPF,
				AccountNumber:   acc.Number,
				Value:           value,
			},
		},
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}

// RemoveMoney ...
func (r Repository) RemoveMoney(ctx context.Context, acc command.Account, value float64) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.repository.RemoveMoney"

	rowID := infra.EventRowID(fmt.Sprintf("%s:AccountAggregate", acc.Owner.CPF))

	event, err := r.in.EventStore.Insert(ctx, infra.Event{
		RowID:       infra.EventRowID(rowID),
		AggregateID: infra.AggregateID(strconv.Itoa(int(acc.Number))),
		Type:        command.MoneyDebitedFromAccountEvent,
		Payload: infra.EventPayload{
			Data: command.MoneyDebitedFromAccountPayload{
				AccountOwnerCPF: acc.Owner.CPF,
				AccountNumber:   acc.Number,
				Value:           value,
			},
		},
	})
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}
