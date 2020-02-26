package account

import (
	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// RepositoryInput ...
type RepositoryInput struct {
	Log        infra.LogProvider
	EventStore domain.EventStoreProvider
}

// Repository ...
type Repository struct {
	in ServiceInput
}

// NewRepository ...
func NewRepository(in ServiceInput) (*Repository, *infra.Error) {
	return &Repository{in: in}, nil
}
