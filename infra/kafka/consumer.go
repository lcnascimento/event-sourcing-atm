package kafka

import (
	"context"

	kafka "github.com/Shopify/sarama"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// ConsumerService ...
type ConsumerService struct {
	in   ServiceInput
	conn kafka.Consumer
}

// NewConsumerService ...
func NewConsumerService(in ServiceInput) (*ConsumerService, *infra.Error) {
	const opName infra.OpName = "kafka.NewService"

	config := kafka.NewConfig()
	config.Version = kafka.V2_4_0_0
	config.ClientID = "1"

	client, err := kafka.NewClient(in.Hosts, config)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	conn, err := kafka.NewConsumerFromClient(client)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return &ConsumerService{in: in, conn: conn}, nil
}

// Subscribe ...
func (s ConsumerService) Subscribe(ctx context.Context, topic infra.EventStreamTopic) (<-chan []byte, *infra.Error) {
	const opName infra.OpName = "kafka.Subscribe"

	cons, err := s.conn.ConsumePartition(string(topic), 0, 0)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	ch := make(chan []byte)
	go func() {
		for msg := range cons.Messages() {
			ch <- msg.Value
		}
	}()

	return ch, nil
}

// CloseConnection ...
func (s ConsumerService) CloseConnection() {
	s.conn.Close()
}
