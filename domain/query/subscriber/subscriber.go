package subscriber

import (
	"context"

	"github.com/golang/protobuf/proto"

	pb "github.com/lcnascimento/event-sourcing-atm/proto/impl"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
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

// Run ...
func (s Service) Run(ctx context.Context) <-chan *infra.Error {
	const opName infra.OpName = "query.subscriber.Run"

	ch := make(chan *infra.Error)

	s.in.Log.Info(ctx, opName, "Initializing subscriber")
	stream, err := s.in.Stream.Subscribe(ctx, infra.EventStreamTopic("accounts"))
	if err != nil {
		ch <- errors.New(opName, err)
		close(ch)
		return ch
	}

	for msg := range stream {
		event := &pb.Event{}
		if err := proto.Unmarshal(msg, event); err != nil {
			ch <- errors.New(ctx, opName, err)
			continue
		}

		data := &pb.AccountCreated{}
		if err := proto.Unmarshal(event.Data, data); err != nil {
			ch <- errors.New(ctx, opName, err)
			continue
		}

		s.in.Log.InfoCustomData(ctx, opName, "Processing new event", infra.CustomData{
			"data": data,
		})
	}

	return ch
}
