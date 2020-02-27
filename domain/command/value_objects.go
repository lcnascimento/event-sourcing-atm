package command

import (
	"errors"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

var (
	// ErrAccountAlreadyExists ...
	ErrAccountAlreadyExists error = errors.New("could not create a new account because the given user already has one")
)

const (
	// AccountCreatedEvent ...
	AccountCreatedEvent infra.EventName = "AccountCreated"
)

// AccountCreatedPayload ...
type AccountCreatedPayload struct {
	Account Account
}
