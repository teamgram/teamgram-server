package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (r *Repository) insertAffectedOutboxesTx(ctx context.Context, txModels *model.TxModels, rows []AffectedOutbox) error {
	if len(rows) == 0 {
		return nil
	}
	for i := range rows {
		row := rows[i]
		if err := validateAffectedOutbox(row); err != nil {
			return err
		}
		outboxID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next affected operation outbox id", err)
		}
		_, affected, err := txModels.AffectedOperationOutboxModel.InsertIgnore(&model.AffectedOperationOutbox{
			OutboxId:          outboxID,
			UserId:            row.TargetUserID,
			RequesterUserId:   row.RequesterUserID,
			OperationId:       row.OperationID,
			OpType:            row.OpType,
			OperationKind:     row.OperationKind,
			PeerType:          row.PeerType,
			PeerId:            row.PeerID,
			PayloadCodec:      row.PayloadCodec,
			PayloadHash:       row.PayloadHash,
			Payload:           row.Payload,
			DeliveryPolicy:    DeliveryPolicyDurableAsync,
			Status:            AffectedOutboxStatusPending,
			RetryCount:        0,
			AvailableAt:       unixNow(),
			BucketId:          row.TargetBucketID,
			PartitionId:       row.TargetPartitionID,
			OwnerTokenPayload: row.OwnerTokenPayload,
		})
		if err != nil {
			return storageError("insert affected operation outbox", err)
		}
		if affected != 0 {
			continue
		}
		existing, err := txModels.AffectedOperationOutboxModel.SelectByUserOperation(row.TargetUserID, row.OperationID)
		if err != nil {
			return storageError("select affected operation outbox duplicate", err)
		}
		if !bytes.Equal(existing.PayloadHash, row.PayloadHash) {
			return userupdates.ErrOperationPayloadConflict
		}
	}
	return nil
}

func validateAffectedOutbox(row AffectedOutbox) error {
	switch {
	case row.TargetUserID <= 0:
		return fmt.Errorf("%w: affected outbox target user id is required", userupdates.ErrOperationTerminal)
	case row.OperationID == "":
		return fmt.Errorf("%w: affected outbox operation id is required", userupdates.ErrOperationTerminal)
	case row.OperationKind == "":
		return fmt.Errorf("%w: affected outbox operation kind is required", userupdates.ErrOperationTerminal)
	case row.DeliveryPolicy != DeliveryPolicyDurableAsync:
		return fmt.Errorf("%w: affected outbox delivery policy=%d", userupdates.ErrOperationTerminal, row.DeliveryPolicy)
	case row.PayloadCodec != PayloadCodecJSON:
		return fmt.Errorf("%w: affected outbox payload codec=%d", userupdates.ErrOperationTerminal, row.PayloadCodec)
	case !bytes.Equal(row.PayloadHash, payload.HashBytes(row.Payload)):
		return userupdates.ErrOperationPayloadConflict
	default:
		return nil
	}
}
