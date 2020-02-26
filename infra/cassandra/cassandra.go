package cassandra

import (
	"context"
	cql "github.com/gocql/gocql"

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
func (c Client) Insert(ctx context.Context, query infra.DatabaseQuery) *infra.Error {
	const opName infra.OpName = "cassandra.Insert"

	if err := c.sess.Query(query.ToSQL()).Exec(); err != nil {
		return errors.New(ctx, opName, err)
	}

	return nil
}
