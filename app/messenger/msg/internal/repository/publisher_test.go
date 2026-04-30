package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestInMemoryReceiverOperationPublisherWaitsForCallback(t *testing.T) {
	wantErr := errors.New("processor failed")
	var called bool
	publisher := &InMemoryReceiverOperationPublisher{
		OnPublish: func(_ context.Context, op payload.ReceiverOperationEnvelopeV1) error {
			called = true
			if op.OperationID != "op1" {
				t.Fatalf("unexpected operation: %+v", op)
			}
			return wantErr
		},
	}

	ack, err := publisher.Publish(context.Background(), payload.ReceiverOperationEnvelopeV1{OperationID: "op1", PartitionID: 7})
	if !errors.Is(err, wantErr) {
		t.Fatalf("Publish() error = %v, want %v", err, wantErr)
	}
	if ack != (KafkaAck{}) {
		t.Fatalf("Publish() ack = %+v, want zero", ack)
	}
	if !called {
		t.Fatal("callback was not called")
	}
}

func TestInMemoryReceiverOperationPublisherReturnsAckAfterCallback(t *testing.T) {
	publisher := &InMemoryReceiverOperationPublisher{
		OnPublish: func(_ context.Context, op payload.ReceiverOperationEnvelopeV1) error {
			if op.OperationID != "op1" {
				t.Fatalf("unexpected operation: %+v", op)
			}
			return nil
		},
	}

	ack, err := publisher.Publish(context.Background(), payload.ReceiverOperationEnvelopeV1{OperationID: "op1", PartitionID: 7})
	if err != nil {
		t.Fatalf("Publish() error = %v", err)
	}
	if ack.Topic != "memory" || ack.Partition != 7 || ack.Offset != 0 {
		t.Fatalf("Publish() ack = %+v", ack)
	}
}
