package transaction

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// ServiceInput ...
type ServiceInput struct {
	Log infra.LogProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}

// Deposit ...
func (s Service) Deposit(ctx context.Context, accID infra.ObjectID, amount float32) *infra.Error {
	const opName infra.OpName = "command.transaction.Deposit"

	s.in.Log.InfoCustomData(ctx, opName, "Making a deposit", infra.CustomData{
		"accountID": accID,
		"amount":    amount,
	})
	return nil
}

// Withdraw ...
func (s Service) Withdraw(ctx context.Context, accID infra.ObjectID, amount float32) *infra.Error {
	const opName infra.OpName = "command.transaction.Withdraw"

	s.in.Log.InfoCustomData(ctx, opName, "Making a withdraw", infra.CustomData{
		"accountID": accID,
		"amount":    amount,
	})
	return nil
}

// Transfer ...
func (s Service) Transfer(ctx context.Context, from infra.ObjectID, to infra.ObjectID, amount float32) *infra.Error {
	const opName infra.OpName = "command.transaction.Transfer"

	s.in.Log.InfoCustomData(ctx, opName, "Making a withdraw", infra.CustomData{
		"fromAccountID": from,
		"toAccountID":   to,
		"amount":        amount,
	})
	return nil
}

// Loan ...
func (s Service) Loan(ctx context.Context, accID infra.ObjectID, amount float32) *infra.Error {
	const opName infra.OpName = "command.transaction.Loan"

	s.in.Log.InfoCustomData(ctx, opName, "Making a withdraw", infra.CustomData{
		"accountID": accID,
		"amount":    amount,
	})
	return nil
}
