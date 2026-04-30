package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/zeromicro/go-zero/core/logx"
)

func (r *Repository) HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error {
	msg, err := payload.UnmarshalReceiverKafkaMessage(record.Value)
	if err != nil {
		if dlqErr := r.recordReceiverTerminal(ctx, record, nil, "payload_decode_failed", err); dlqErr != nil {
			return fmt.Errorf("receiver delivery dlq failed: topic=%s partition=%d offset=%d: %w", record.Topic, record.Partition, record.Offset, dlqErr)
		}
		logx.WithContext(ctx).Errorf("receiver delivery terminal: topic=%s partition=%d offset=%d code=payload_decode_failed", record.Topic, record.Partition, record.Offset)
		return nil
	}

	_, err = r.ApplyUserOperation(ctx, ApplyUserOperationInput{
		UserID:        msg.Operation.UserID,
		OperationID:   msg.Operation.OperationID,
		OpType:        msg.Operation.OpType,
		PeerType:      msg.Operation.PeerType,
		PeerID:        msg.Operation.PeerID,
		PayloadCodec:  msg.Operation.PayloadCodec,
		Payload:       msg.Operation.Payload,
		PayloadHash:   msg.Operation.PayloadHash,
		BucketID:      msg.Operation.BucketID,
		PartitionID:   msg.Operation.PartitionID,
		DependencyPts: msg.Operation.DependencyPts,
	})
	if err == nil {
		return nil
	}
	if errors.Is(err, userupdates.ErrOperationTerminal) || errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		if dlqErr := r.recordReceiverTerminal(ctx, record, &msg.Operation, "operation_terminal", err); dlqErr != nil {
			return fmt.Errorf("receiver delivery dlq failed: operation_id=%s topic=%s partition=%d offset=%d: %w", msg.Operation.OperationID, record.Topic, record.Partition, record.Offset, dlqErr)
		}
		logx.WithContext(ctx).Errorf("receiver delivery terminal: operation_id=%s topic=%s partition=%d offset=%d code=operation_terminal", msg.Operation.OperationID, record.Topic, record.Partition, record.Offset)
		return nil
	}
	return fmt.Errorf("receiver delivery retryable: operation_id=%s topic=%s partition=%d offset=%d: %w", msg.Operation.OperationID, record.Topic, record.Partition, record.Offset, err)
}

func (r *Repository) recordReceiverTerminal(ctx context.Context, record ReceiverKafkaRecord, op *payload.ReceiverOperationEnvelopeV1, code string, cause error) error {
	in := RecordDeliveryFailureInput{
		KafkaTopic:      record.Topic,
		KafkaPartition:  record.Partition,
		KafkaOffset:     record.Offset,
		PayloadHash:     payload.HashBytes(record.Value),
		FailureCategory: FailureCategoryCorruption,
		FailureCode:     code,
		FailureMessage:  cause.Error(),
	}
	if op != nil {
		in.UserId = op.UserID
		in.OperationId = op.OperationID
		in.OpType = op.OpType
		in.BucketId = op.BucketID
		in.PayloadHash = op.PayloadHash
	}
	return r.RecordDeliveryFailure(ctx, in)
}
