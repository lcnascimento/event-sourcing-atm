package domain

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// Aggregate ...
type Aggregate interface {
	Name() string
	Apply(context.Context, Event) (Aggregate, *infra.Error)
}

// EventHandlerProvider ...
type EventHandlerProvider interface {
	Project(context.Context, Aggregate, []Event) (Aggregate, *infra.Error)
}

// EventStoreProvider ...
type EventStoreProvider interface {
	Insert(context.Context, Event) *infra.Error
	List(context.Context, *EventRowIDPattern, *AggregateID) ([]*Event, *infra.Error)
}

// Executable ...
type Executable interface {
	Run(context.Context) <-chan *infra.Error
}
