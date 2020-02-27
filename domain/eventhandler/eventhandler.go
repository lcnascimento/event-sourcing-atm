package eventhandler

import (
	"context"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// ServiceInput ...
type ServiceInput struct {
	Log infra.LogProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}

// Project ...
func (s Service) Project(ctx context.Context, state domain.Aggregate, events []infra.Event) (domain.Aggregate, *infra.Error) {
	const opName infra.OpName = "eventhandler.Project"

	for _, event := range events {
		model, err := state.Apply(ctx, event)

		if err == domain.ErrApplyEventIntoAggregate {
			s.in.Log.WarningCustomData(ctx, opName, err.Error(), infra.CustomData{
				"aggregate_name": state.Name(),
				"event":          event,
			})
			continue
		}

		if err != nil {
			return nil, errors.New(ctx, opName, err)
		}

		state = model
	}

	return state, nil
}
