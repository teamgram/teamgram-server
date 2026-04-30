//go:build integration

package repository

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestReceiverDeliveryAppliesOperation(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 80_000}, "local-userupdates")
	in := buildApplyInput(t, base+3001, base+3002, base+3001, false, "receiver delivery")
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	record := buildReceiverKafkaRecord(t, in, "test.receiver.ops", 0, base+3)

	if err := repo.HandleReceiverKafkaRecord(ctx, record); err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() error = %v", err)
	}
	diff, err := repo.GetDifference(ctx, GetDifferenceInput{UserID: in.UserID, Pts: 0, Limit: 10})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.Events) != 1 || diff.Events[0].OperationID != in.OperationID {
		t.Fatalf("events mismatch: %+v", diff.Events)
	}
}

func TestReceiverDeliveryRecordsDLQForDecodeFailure(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 90_000}, "local-userupdates")
	record := ReceiverKafkaRecord{Topic: "test.receiver.ops", Partition: 8, Offset: base + 4, Value: []byte("{bad-json")}

	if err := repo.HandleReceiverKafkaRecord(ctx, record); err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() error = %v", err)
	}
	row, err := repo.models.DeliveryFailedOperationsModel.SelectByKafkaOffset(ctx, record.Topic, record.Partition, record.Offset)
	if err != nil {
		t.Fatalf("SelectByKafkaOffset() error = %v", err)
	}
	if row.FailureCode != "payload_decode_failed" || row.Status != DeliveryFailedOperationStatusOpen {
		t.Fatalf("dlq row mismatch: %+v", row)
	}
}

func TestReceiverDeliveryRecordsDLQForTerminalOperation(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 100_000}, "local-userupdates")
	in := buildApplyInput(t, base+4001, base+4002, base+4001, false, "terminal")
	in.DependencyPts = []int64{1}
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	record := buildReceiverKafkaRecord(t, in, "test.receiver.ops", 9, base+5)

	if err := repo.HandleReceiverKafkaRecord(ctx, record); err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() error = %v", err)
	}
	row, err := repo.models.DeliveryFailedOperationsModel.SelectByKafkaOffset(ctx, record.Topic, record.Partition, record.Offset)
	if err != nil {
		t.Fatalf("SelectByKafkaOffset() error = %v", err)
	}
	if row.OperationId != in.OperationID || row.FailureCode != "operation_terminal" {
		t.Fatalf("dlq row mismatch: %+v", row)
	}
}

func TestReceiverDeliveryWrapsRetryableErrorWithKafkaContext(t *testing.T) {
	ctx := context.Background()
	repo := NewForTest(nil, &testIDGenerator{next: 1}, "local-userupdates")
	in := buildApplyInput(t, 5001, 5002, 5001, false, "retryable")
	record := buildReceiverKafkaRecord(t, in, "test.receiver.ops", 10, 99)

	err := repo.HandleReceiverKafkaRecord(ctx, record)
	if err == nil {
		t.Fatal("HandleReceiverKafkaRecord() expected retryable error")
	}
	if !strings.Contains(err.Error(), "operation_id="+in.OperationID) ||
		!strings.Contains(err.Error(), "topic=test.receiver.ops") ||
		!strings.Contains(err.Error(), "partition=10") ||
		!strings.Contains(err.Error(), "offset=99") {
		t.Fatalf("error lacks kafka context: %v", err)
	}
}

func buildReceiverKafkaRecord(t *testing.T, in ApplyUserOperationInput, topic string, partition int32, offset int64) ReceiverKafkaRecord {
	t.Helper()
	body, err := payload.MarshalReceiverKafkaMessage(payload.ReceiverKafkaMessageV1{
		SchemaVersion: payload.ReceiverKafkaMessageSchemaVersion,
		Operation: payload.ReceiverOperationEnvelopeV1{
			UserID:        in.UserID,
			BucketID:      in.BucketID,
			PartitionID:   in.PartitionID,
			OperationID:   in.OperationID,
			OpType:        in.OpType,
			PeerType:      in.PeerType,
			PeerID:        in.PeerID,
			PayloadCodec:  in.PayloadCodec,
			Payload:       in.Payload,
			PayloadHash:   in.PayloadHash,
			DependencyPts: in.DependencyPts,
		},
	})
	if err != nil {
		t.Fatalf("MarshalReceiverKafkaMessage() error = %v", err)
	}
	return ReceiverKafkaRecord{Topic: topic, Partition: partition, Offset: offset, Value: body}
}
