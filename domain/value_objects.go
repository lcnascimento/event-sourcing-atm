package domain

import "errors"

var (
	// ErrApplyEventIntoAggregate ...
	ErrApplyEventIntoAggregate error = errors.New("could not apply event into aggregate")
)

// EventName ...
type EventName string

// EventRowIDPattern ...
type EventRowIDPattern string

// AggregateID ...
type AggregateID string

const (
	// AccountCreatedEvent ...
	AccountCreatedEvent EventName = "AccountCreated"
)

// AccountCreatedPayload ...
type AccountCreatedPayload struct {
}
