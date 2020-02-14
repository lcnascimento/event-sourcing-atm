package domain

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// EventStoreProvider ...
type EventStoreProvider interface {
	Insert(context.Context, Event) *infra.Error
	List(context.Context, ListEventsQuery) ([]*Event, *infra.Error)
}

// Executable ...
type Executable interface {
	Run(context.Context) <-chan *infra.Error
}
