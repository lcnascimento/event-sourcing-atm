package cassandra

import (
	"context"
	"time"

	cql "github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// ClientInput ...
type ClientInput struct {
	Log      infra.LogProvider
	Hosts    []string
	Keyspace string
}

// Client ...
type Client struct {
	in   ClientInput
	sess *cql.Session
}

// NewClient ...
func NewClient(in ClientInput) (*Client, *infra.Error) {
	const opName infra.OpName = "cassandra.NewClient"

	cluster := cql.NewCluster(in.Hosts...)
	cluster.Keyspace = in.Keyspace
	cluster.Consistency = cql.Quorum
	cluster.ProtoVersion = 3

	sess, err := cluster.CreateSession()
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return &Client{in: in, sess: sess}, nil
}

// CloseSession ...
func (c Client) CloseSession() {
	c.sess.Close()
}

// Insert ...
func (c Client) Insert(ctx context.Context, event infra.Event) (*infra.Event, *infra.Error) {
	const opName infra.OpName = "cassandra.Insert"

	if event.Timestamp == nil {
		now := time.Now()
		event.Timestamp = &now
	}

	stmt, names := qb.Insert("events").Columns("row_id", "aggregate_id", "event_time", "event_type", "payload").ToCql()
	query := gocqlx.Query(c.sess.Query(stmt), names).BindStruct(event)

	if err := query.ExecRelease(); err != nil {
		return nil, errors.New(ctx, opName, err)
	}

	return &event, nil
}

// List ...
func (c Client) List(ctx context.Context, in infra.ListEventsInput) ([]infra.Event, *infra.Error) {
	const opName infra.OpName = "cassandra.List"

	whereStmts := []qb.Cmp{}
	bindMap := qb.M{}

	if in.RowID != nil {
		whereStmts = append(whereStmts, qb.Eq("row_id"))
		bindMap["row_id"] = string(*in.RowID)
	}

	if in.AggregateID != nil {
		whereStmts = append(whereStmts, qb.Eq("aggregate_id"))
		bindMap["aggregate_id"] = string(*in.AggregateID)
	}

	stmt, names := qb.Select("events").Where(whereStmts...).ToCql()
	query := gocqlx.Query(c.sess.Query(stmt), names).BindMap(bindMap)

	var events []infra.Event
	if err := query.SelectRelease(&events); err != nil {
		return nil, errors.New(ctx, opName, err)
	}

	return events, nil
}
