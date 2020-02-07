package kafka

import (
	"context"

	kafka "github.com/Shopify/sarama"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"
)

// ServiceInput ...
type ServiceInput struct {
	Log infra.LogProvider

	Hosts []string
}

// ProducerService ...
type ProducerService struct {
	in   ServiceInput
	conn kafka.AsyncProducer
}

// NewProducerService ...
func NewProducerService(in ServiceInput) (*ProducerService, *infra.Error) {
	const opName infra.OpName = "kafka.NewService"

	config := kafka.NewConfig()
	config.Version = kafka.V2_4_0_0
	config.Producer.RequiredAcks = kafka.RequiredAcks(1)

	client, err := kafka.NewClient(in.Hosts, config)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	conn, err := kafka.NewAsyncProducerFromClient(client)
	if err != nil {
		return nil, errors.New(opName, err)
	}

	return &ProducerService{in: in, conn: conn}, nil
}

// Publish ...
func (s ProducerService) Publish(ctx context.Context, topic infra.EventStreamTopic, data []byte) *infra.Error {
	const opName infra.OpName = "kafka.Publish"

	msg := &kafka.ProducerMessage{
		Topic: string(topic),
		Value: kafka.ByteEncoder(data),
	}

	s.in.Log.DebugCustomData(ctx, opName, "Publishing data into stream", infra.CustomData{
		"topic": topic,
		"msg":   msg,
	})

	s.conn.Input() <- msg
	return nil
}

// CloseConnection ...
func (s ProducerService) CloseConnection() {
	s.conn.Close()
}
