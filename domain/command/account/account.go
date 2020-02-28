package account

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"

	"github.com/lcnascimento/event-sourcing-atm/domain/command"
)

// ServiceInput ...
type ServiceInput struct {
	Log        infra.LogProvider
	Repository command.AccountsRepository
}

// Service ...
type Service struct {
	in                ServiceInput
	nextAccountNumber int // TODO: This must be changed in the future!
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in, nextAccountNumber: 1}, nil
}

// Create ...
func (s *Service) Create(ctx context.Context, user command.User) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.Create"

	accounts, err := s.in.Repository.ListByCPF(ctx, user.CPF)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	if len(accounts) > 0 {
		return nil, errors.New(ctx, opName, command.ErrAccountAlreadyExists, infra.KindBadRequest)
	}

	account := command.Account{
		Agency:  command.AccountAgency(1),
		Number:  command.AccountNumber(s.nextAccountNumber),
		Balance: 0,
		Owner:   user,
	}
	event, err := s.in.Repository.Insert(ctx, account)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	s.nextAccountNumber++

	return event, nil
}

// Delete ...
func (s Service) Delete(ctx context.Context, cpf command.CPF, accNumber command.AccountNumber) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.Delete"

	account, err := s.in.Repository.Find(ctx, cpf, accNumber)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	event, err := s.in.Repository.Remove(ctx, *account)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}

// Credit ...
func (s Service) Credit(ctx context.Context, cpf command.CPF, accNumber command.AccountNumber, value float64) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.Credit"

	account, err := s.in.Repository.Find(ctx, cpf, accNumber)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	event, err := s.in.Repository.PutMoney(ctx, *account, value)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}

// Debit ...
func (s Service) Debit(ctx context.Context, cpf command.CPF, accNumber command.AccountNumber, value float64) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "account.Debit"

	account, err := s.in.Repository.Find(ctx, cpf, accNumber)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	if account.Balance < value {
		return nil, errors.New(ctx, opName, command.ErrInsufficientFounds, infra.KindBadRequest)
	}

	event, err := s.in.Repository.RemoveMoney(ctx, *account, value)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return event, nil
}
