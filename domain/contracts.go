package domain

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// Executable ...
type Executable interface {
	Run(context.Context) <-chan *infra.Error
}
