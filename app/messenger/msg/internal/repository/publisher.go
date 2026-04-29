package repository

import "context"

type ReceiverOperation struct {
	UserID        int64
	BucketID      int32
	PartitionID   int32
	OperationID   string
	OpType        int32
	PeerType      int32
	PeerID        int64
	PayloadCodec  int32
	Payload       []byte
	PayloadHash   string
	DependencyPts []int64
}

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
