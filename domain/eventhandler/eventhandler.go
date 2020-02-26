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
func (s Service) Project(ctx context.Context, agg domain.Aggregate, events []domain.Event) (domain.Aggregate, *infra.Error) {
	const opName infra.OpName = "eventhandler.Project"

	for _, event := range events {
		model, err := agg.Apply(ctx, event)

		if err == domain.ErrApplyEventIntoAggregate {
			s.in.Log.WarningCustomData(ctx, opName, err.Error(), infra.CustomData{
				"aggregate_name": agg.Name(),
				"event":          event,
			})
			continue
		}

		if err != nil {
			return nil, errors.New(ctx, opName, err)
		}

		agg = model
	}

	return agg, nil
}
