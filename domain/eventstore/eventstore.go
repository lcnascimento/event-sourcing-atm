package eventstore

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"

	"github.com/lcnascimento/event-sourcing-atm/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Log infra.LogProvider
	Db  infra.DatabaseProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}

// Insert ...
func (s Service) Insert(ctx context.Context, e domain.Event) *infra.Error {
	const opName infra.OpName = "eventstore.Insert"

	if err := s.in.Db.Insert(ctx, e); err != nil {
		return errors.New(ctx, opName, err)
	}

	return nil
}
