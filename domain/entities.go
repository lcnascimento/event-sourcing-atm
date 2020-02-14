package domain

import (
	"fmt"
	"time"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// Account ...
type Account struct {
}

// Event ...
type Event struct {
	RowID       infra.ObjectID
	AggregateID infra.ObjectID
	Timestamp   time.Time
	Payload     interface{}
}

// ToSQL ...
func (e Event) ToSQL() string {
	return fmt.Sprintf(`
		INSERT INTO events (row_id, aggregate_id, event_time)
		VALUES ('%s', '%s', toTimestamp(NOW()))
	`, e.RowID, e.AggregateID)
}

// ListEventsQuery ...
type ListEventsQuery struct {
	RowID       infra.ObjectID
	AggregateID infra.ObjectID
	Timestamp   time.Time
}
