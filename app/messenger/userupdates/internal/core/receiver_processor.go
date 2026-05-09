package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type ReceiverProcessor struct {
	repo svc.UserUpdatesRepository
}

func NewReceiverProcessor(repo svc.UserUpdatesRepository) *ReceiverProcessor {
	return &ReceiverProcessor{repo: repo}
}

func (p *ReceiverProcessor) Process(ctx context.Context, op payload.ReceiverOperationEnvelopeV1) error {
	if err := validateReceiverOperationPayload(op.Payload); err != nil {
		return err
	}
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

func validateReceiverOperationPayload(body []byte) error {
	var envelope struct {
		SchemaVersion int    `json:"schema_version"`
		OperationKind string `json:"operation_kind"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return fmt.Errorf("%w: decode receiver operation envelope: %v", userupdates.ErrOperationTerminal, err)
	}
	switch envelope.SchemaVersion {
	case payload.MessageOperationSchemaVersionV1, payload.MessageOperationSchemaVersion:
		var op payload.MessageOperationV1
		if err := json.Unmarshal(body, &op); err != nil {
			return fmt.Errorf("%w: decode receiver message operation: %v", userupdates.ErrOperationTerminal, err)
		}
	case payload.MessageOperationSchemaVersionV3:
		var op payload.MessageOperationV3
		if err := json.Unmarshal(body, &op); err != nil {
			return fmt.Errorf("%w: decode receiver v3 message operation: %v", userupdates.ErrOperationTerminal, err)
		}
	default:
		return fmt.Errorf("%w: unsupported receiver operation schema=%d kind=%s", userupdates.ErrOperationTerminal, envelope.SchemaVersion, envelope.OperationKind)
	}
	return nil
}
