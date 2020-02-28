package infra

import (
	"encoding/json"
	"time"

	"github.com/gocql/gocql"
)

// EventPayload ...
type EventPayload struct {
	Data interface{}
}

// Cast ...
func (p EventPayload) Cast(out interface{}) error {
	const opName OpName = "entity.Cast"

	b, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, out); err != nil {
		return err
	}

	return nil
}

// MarshalCQL ...
func (p EventPayload) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return json.Marshal(p.Data)
}

// UnmarshalCQL ...
func (p *EventPayload) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	return json.Unmarshal(data, &p.Data)
}

// Event ...
type Event struct {
	RowID       EventRowID   `db:"row_id"`
	AggregateID AggregateID  `db:"aggregate_id"`
	Timestamp   *time.Time   `db:"event_time"`
	Payload     EventPayload `db:"payload"`
	Type        EventName    `db:"event_type"`
}
