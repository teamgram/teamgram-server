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

	err := publisher.Publish(context.Background(), payload.ReceiverOperationEnvelopeV1{OperationID: "op1"})
	if !errors.Is(err, wantErr) {
		t.Fatalf("Publish() error = %v, want %v", err, wantErr)
	}
	if !called {
		t.Fatal("callback was not called")
	}
}
