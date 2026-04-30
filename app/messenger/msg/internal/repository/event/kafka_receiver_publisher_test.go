package event

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/IBM/sarama"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type fakeSyncProducer struct {
	msg       *sarama.ProducerMessage
	partition int32
	offset    int64
	err       error
}

func (p *fakeSyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	p.msg = msg
	return p.partition, p.offset, p.err
}

func (p *fakeSyncProducer) SendMessages([]*sarama.ProducerMessage) error { return nil }

func (p *fakeSyncProducer) Close() error { return nil }

func TestKafkaReceiverOperationPublisherUsesExplicitPartition(t *testing.T) {
	producer := &fakeSyncProducer{partition: 7, offset: 42}
	pub := NewKafkaReceiverOperationPublisherForTest("userupdates.receiver_operations.v1", producer)
	op := repository.ReceiverOperation{
		UserID:       1001,
		BucketID:     7,
		PartitionID:  7,
		OperationID:  "v1:msg:2001:receiver:1001:in",
		OpType:       payload.OpTypeSendMessage,
		PayloadCodec: payload.PayloadCodecJSON,
		Payload:      []byte(`{"schema_version":1}`),
		PayloadHash:  payload.HashBytes([]byte(`{"schema_version":1}`)),
	}
	ack, err := pub.Publish(context.Background(), op)
	if err != nil {
		t.Fatalf("Publish() error = %v", err)
	}
	if ack.Topic != "userupdates.receiver_operations.v1" || ack.Partition != 7 || ack.Offset != 42 {
		t.Fatalf("ack = %+v", ack)
	}
	if producer.msg.Partition != 7 {
		t.Fatalf("producer partition = %d", producer.msg.Partition)
	}
	if producer.msg.Key == nil {
		t.Fatal("expected kafka key")
	}
	var body payload.ReceiverKafkaMessageV1
	if err := json.Unmarshal(producer.msg.Value.(sarama.ByteEncoder), &body); err != nil {
		t.Fatalf("unmarshal produced value: %v", err)
	}
	if body.Operation.OperationID != op.OperationID {
		t.Fatalf("operation id = %q", body.Operation.OperationID)
	}
}

func TestKafkaReceiverOperationPublisherReturnsProducerError(t *testing.T) {
	producer := &fakeSyncProducer{err: errors.New("broker unavailable")}
	pub := NewKafkaReceiverOperationPublisherForTest("topic", producer)
	_, err := pub.Publish(context.Background(), repository.ReceiverOperation{PartitionID: 1, OperationID: "op"})
	if err == nil {
		t.Fatal("expected publish error")
	}
}
