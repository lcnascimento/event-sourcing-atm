package atmcli

import (
	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// ServiceInput ...
type ServiceInput struct {
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}
