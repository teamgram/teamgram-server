package event

import (
	"context"
	"testing"
)

type fakePushTaskHandler struct {
	record PushTaskKafkaRecord
	calls  int
}

func (h *fakePushTaskHandler) HandlePushTaskKafkaRecord(ctx context.Context, record PushTaskKafkaRecord) error {
	h.calls++
	h.record = record
	return nil
}

func TestPushTaskReceiverAdapterConvertsKafkaRecord(t *testing.T) {
	handler := &fakePushTaskHandler{}
	adapter := pushTaskReceiverAdapter{handler: handler}

	err := adapter.HandleReceiverKafkaRecord(context.Background(), ReceiverKafkaRecord{
		Topic:     "userupdates.push_tasks.v1",
		Partition: 7,
		Offset:    8,
		Key:       []byte("k"),
		Value:     []byte("v"),
	})
	if err != nil {
		t.Fatalf("HandleReceiverKafkaRecord() error = %v", err)
	}
	if handler.calls != 1 {
		t.Fatalf("handler calls = %d, want 1", handler.calls)
	}
	if handler.record.Topic != "userupdates.push_tasks.v1" || handler.record.Partition != 7 || handler.record.Offset != 8 || string(handler.record.Value) != "v" {
		t.Fatalf("record = %#v", handler.record)
	}
}
