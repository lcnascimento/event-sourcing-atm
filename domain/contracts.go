package domain

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// Aggregate ...
type Aggregate interface {
	Name() string
	Apply(context.Context, infra.Event) (Aggregate, *infra.Error)
}

// EventHandlerProvider ...
type EventHandlerProvider interface {
	Project(context.Context, Aggregate, []infra.Event) (Aggregate, *infra.Error)
}

// Executable ...
type Executable interface {
	Run(context.Context) <-chan *infra.Error
}
