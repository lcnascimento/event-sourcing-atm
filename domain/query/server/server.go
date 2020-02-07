package server

import (
	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// ServiceInput ...
type ServiceInput struct {
	Log    infra.LogProvider
	Stream infra.EventStreamSubscriber
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}
