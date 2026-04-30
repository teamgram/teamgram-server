//go:build integration

package repository

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestRecordDeliveryFailureInsertsDLQRow(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 60_000}, "local-userupdates")

	in := RecordDeliveryFailureInput{
		UserId:          base + 1001,
		OperationId:     "v1:test:dlq:insert",
		OpType:          OpTypeSendMessage,
		BucketId:        17,
		KafkaTopic:      "test.receiver.dlq",
		KafkaPartition:  3,
		KafkaOffset:     base + 1,
		PayloadHash:     payload.HashBytes([]byte("bad-payload")),
		FailureCategory: FailureCategoryCorruption,
		FailureCode:     "payload_decode_failed",
		FailureMessage:  strings.Repeat("x", 300) + "\nraw details",
	}
	if err := repo.RecordDeliveryFailure(ctx, in); err != nil {
		t.Fatalf("RecordDeliveryFailure() error = %v", err)
	}
	row, err := repo.models.DeliveryFailedOperationsModel.SelectByKafkaOffset(ctx, in.KafkaTopic, in.KafkaPartition, in.KafkaOffset)
	if err != nil {
		t.Fatalf("SelectByKafkaOffset() error = %v", err)
	}
	if row.UserId != in.UserId || row.OperationId != in.OperationId || row.BucketId != in.BucketId || !bytes.Equal(row.PayloadHash, in.PayloadHash) {
		t.Fatalf("dlq row mismatch: %+v", row)
	}
	if row.Status != DeliveryFailedOperationStatusOpen || row.FailureCategory != FailureCategoryCorruption {
		t.Fatalf("dlq status/category mismatch: %+v", row)
	}
	if strings.Contains(row.FailureMessage, "\n") || len(row.FailureMessage) > 256 {
		t.Fatalf("failure message was not sanitized: %q", row.FailureMessage)
	}
}

func TestRecordDeliveryFailureIsIdempotentByKafkaOffset(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 70_000}, "local-userupdates")
	in := RecordDeliveryFailureInput{
		UserId:          base + 2001,
		OperationId:     "v1:test:dlq:idempotent",
		OpType:          OpTypeSendMessage,
		BucketId:        23,
		KafkaTopic:      "test.receiver.dlq",
		KafkaPartition:  4,
		KafkaOffset:     base + 2,
		PayloadHash:     payload.HashBytes([]byte("terminal")),
		FailureCategory: FailureCategoryCorruption,
		FailureCode:     "operation_terminal",
		FailureMessage:  "terminal",
	}
	if err := repo.RecordDeliveryFailure(ctx, in); err != nil {
		t.Fatalf("RecordDeliveryFailure() first error = %v", err)
	}
	if err := repo.RecordDeliveryFailure(ctx, in); err != nil {
		t.Fatalf("RecordDeliveryFailure() duplicate error = %v", err)
	}
	row, err := repo.models.DeliveryFailedOperationsModel.SelectByKafkaOffset(ctx, in.KafkaTopic, in.KafkaPartition, in.KafkaOffset)
	if err != nil {
		t.Fatalf("SelectByKafkaOffset() error = %v", err)
	}
	if row.OperationId != in.OperationId {
		t.Fatalf("row operation_id = %q", row.OperationId)
	}
}
