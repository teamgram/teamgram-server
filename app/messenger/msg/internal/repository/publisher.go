package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type ReceiverOperation = payload.ReceiverOperationEnvelopeV1

type ReceiverOperationPublisher interface {
	Publish(ctx context.Context, op ReceiverOperation) error
}

type InMemoryReceiverOperationPublisher struct {
	OnPublish func(ctx context.Context, op ReceiverOperation) error
}

func (p *InMemoryReceiverOperationPublisher) Publish(ctx context.Context, op ReceiverOperation) error {
	if p == nil || p.OnPublish == nil {
		return nil
	}
	return p.OnPublish(ctx, op)
}
