package command

import (
	"errors"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

var (
	// ErrAccountAlreadyExists ...
	ErrAccountAlreadyExists error = errors.New("could not create a new account because the given user already has one")
	// ErrAccountDeleted ...
	ErrAccountDeleted error = errors.New("could not perform the operation because the given account in already deleted")
	// ErrDecodeEventPayload ...
	ErrDecodeEventPayload error = errors.New("could not decode event payload")
	// ErrUnexpectedEventInAggregateState ...
	ErrUnexpectedEventInAggregateState error = errors.New("could not project event into aggregate because its actual state does not expect it")
	// ErrInsufficientFounds ...
	ErrInsufficientFounds error = errors.New("could not perform the operation due to insufficient founds in account")
)

// CPF ...
type CPF string

// AccountNumber ...
type AccountNumber int

// AccountAgency ...
type AccountAgency int

const (
	// AccountCreatedEvent ...
	AccountCreatedEvent infra.EventName = "AccountCreated"
	// AccountDeletedEvent ...
	AccountDeletedEvent infra.EventName = "AccountDeleted"
	// MoneyCreditedIntoAccountEvent ...
	MoneyCreditedIntoAccountEvent infra.EventName = "MoneyCreditedIntoAccount"
	// MoneyDebitedFromAccountEvent ...
	MoneyDebitedFromAccountEvent infra.EventName = "MoneyDebitedFromAccount"
)

// AccountCreatedPayload ...
type AccountCreatedPayload struct {
	Account *Account `json:"account"`
}

// AccountDeletedPayload ...
type AccountDeletedPayload struct {
	AccountOwnerCPF CPF           `json:"cpf"`
	AccountNumber   AccountNumber `json:"accountNumber"`
}

// MoneyCreditedIntoAccountPayload ...
type MoneyCreditedIntoAccountPayload struct {
	AccountOwnerCPF CPF           `json:"cpf"`
	AccountNumber   AccountNumber `json:"accountNumber"`
	Value           float64       `json:"value"`
}

// MoneyDebitedFromAccountPayload ...
type MoneyDebitedFromAccountPayload struct {
	AccountOwnerCPF CPF           `json:"cpf"`
	AccountNumber   AccountNumber `json:"accountNumber"`
	Value           float64       `json:"value"`
}
