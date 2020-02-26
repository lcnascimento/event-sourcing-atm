package eventstore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

type insertEventQuery struct {
	event domain.Event
}

// ToSQL ...
func (q insertEventQuery) ToSQL() string {
	bytes, err := json.Marshal(q.event.Payload)
	if err != nil {
		fmt.Println("Deu ruim demais!")
		return ""
	}

	return fmt.Sprintf(`
		INSERT INTO events (row_id, aggregate_id, event_time, event_type, payload)
		VALUES ('%s', '%s', toTimestamp(NOW()), '%s', '%s')
	`, q.event.RowID, q.event.AggregateID, q.event.Type, bytes)
}

// Insert ...
func (s Service) Insert(ctx context.Context, e domain.Event) *infra.Error {
	const opName infra.OpName = "eventstore.Insert"

	if e.Timestamp == nil {
		now := time.Now()
		e.Timestamp = &now
	}

	if err := s.in.Db.Insert(ctx, insertEventQuery{event: e}); err != nil {
		return errors.New(ctx, opName, err)
	}

	return nil
}

type findEventsQuery struct {
	event domain.Event
}

// ToSQL ...
func (q findEventsQuery) ToSQL() string {
	return fmt.Sprintf(`
		SELECT row_id, aggregate_id, event_time, event_type, payload
		FROM events
		WHERE 1 = 1
		AND row_id = '%s'
	`, q.event.RowID)
}

// Find ...
func (s Service) Find(
	ctx context.Context,
	rowID *domain.EventRowIDPattern,
	aggID *domain.AggregateID,
) ([]*domain.Event, *infra.Error) {
	return nil, nil
}
