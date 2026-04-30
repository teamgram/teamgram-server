package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type ReceiverOperation = payload.ReceiverOperationEnvelopeV1

type ReceiverOperationPublisher interface {
	Publish(ctx context.Context, op ReceiverOperation) (KafkaAck, error)
}

type InMemoryReceiverOperationPublisher struct {
	OnPublish func(ctx context.Context, op ReceiverOperation) error
}

func (p *InMemoryReceiverOperationPublisher) Publish(ctx context.Context, op ReceiverOperation) (KafkaAck, error) {
	if p == nil || p.OnPublish == nil {
		return KafkaAck{}, nil
	}
	if err := p.OnPublish(ctx, op); err != nil {
		return KafkaAck{}, err
	}
	return KafkaAck{Topic: "memory", Partition: op.PartitionID, Offset: 0}, nil
}
