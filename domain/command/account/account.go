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
		Agency:  1,
		Balance: 0,
		Number:  s.nextAccountNumber,
		Owner:   user,
	}
	event, err := s.in.Repository.Insert(ctx, account)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	s.nextAccountNumber++

	return event, nil
}
