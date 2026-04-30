//go:build integration && kafka

package repository

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"
)

func TestReceiverKafkaReplayUsesOperationResultIdempotency(t *testing.T) {
	if os.Getenv("TEAMGRAM_TEST_KAFKA_BROKERS") == "" {
		t.Skip("TEAMGRAM_TEST_KAFKA_BROKERS is empty")
	}
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 140_000}, "local-userupdates")
	in := buildApplyInput(t, base+6001, base+6002, base+6001, false, "replay")
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	record := buildReceiverKafkaRecord(t, in, "test.receiver.ops", 11, base+6)

	if err := repo.HandleReceiverKafkaRecord(ctx, record); err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() first error = %v", err)
	}
	if err := repo.HandleReceiverKafkaRecord(ctx, record); err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() replay error = %v", err)
	}
	diff, err := repo.GetDifference(ctx, GetDifferenceInput{UserID: in.UserID, Pts: 0, Limit: 10})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.Events) != 1 {
		t.Fatalf("events length = %d, want 1", len(diff.Events))
	}
	opResult, err := repo.GetOperationResult(ctx, in.UserID, in.OperationID)
	if err != nil {
		t.Fatalf("GetOperationResult() error = %v", err)
	}
	if opResult.Pts != 1 || !bytes.Equal(opResult.PayloadHash, in.PayloadHash) {
		t.Fatalf("operation result mismatch: %+v", opResult)
	}
}
