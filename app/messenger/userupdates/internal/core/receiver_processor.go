package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type ReceiverProcessor struct {
	repo svc.UserUpdatesRepository
}

func NewReceiverProcessor(repo svc.UserUpdatesRepository) *ReceiverProcessor {
	return &ReceiverProcessor{repo: repo}
}

func (p *ReceiverProcessor) Process(ctx context.Context, op payload.ReceiverOperationEnvelopeV1) error {
	_, err := p.repo.ApplyUserOperation(ctx, repository.ApplyUserOperationInput{
		UserID:        op.UserID,
		OperationID:   op.OperationID,
		OpType:        op.OpType,
		PeerType:      op.PeerType,
		PeerID:        op.PeerID,
		PayloadCodec:  op.PayloadCodec,
		Payload:       op.Payload,
		PayloadHash:   op.PayloadHash,
		BucketID:      op.BucketID,
		PartitionID:   op.PartitionID,
		DependencyPts: op.DependencyPts,
	})
	return err
}
