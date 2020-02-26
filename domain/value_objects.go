package domain

import (
	"errors"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

var (
	// ErrApplyEventIntoAggregate ...
	ErrApplyEventIntoAggregate error = errors.New("could not apply event into aggregate")
)

const (
	// AccountCreatedEvent ...
	AccountCreatedEvent infra.EventName = "AccountCreated"
)

// AccountCreatedPayload ...
type AccountCreatedPayload struct {
}
