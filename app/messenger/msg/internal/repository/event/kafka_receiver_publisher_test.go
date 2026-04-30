package event

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
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
	closeErr  error
	onSend    func()
}

func (p *fakeSyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	p.msg = msg
	if p.onSend != nil {
		p.onSend()
	}
	return p.partition, p.offset, p.err
}

func (p *fakeSyncProducer) SendMessages([]*sarama.ProducerMessage) error { return nil }

func (p *fakeSyncProducer) Close() error { return p.closeErr }

type fakeProducerCounters struct {
	publishSuccess int
	publishError   int
	closeError     int
}

func (c *fakeProducerCounters) IncPublishSuccess() { c.publishSuccess++ }
func (c *fakeProducerCounters) IncPublishError()   { c.publishError++ }
func (c *fakeProducerCounters) IncCloseError()     { c.closeError++ }

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

func TestKafkaReceiverOperationPublisherCountsPublishSuccess(t *testing.T) {
	producer := &fakeSyncProducer{partition: 7, offset: 42}
	counters := &fakeProducerCounters{}
	pub := NewKafkaReceiverOperationPublisherForTest("topic", producer).WithCounters(counters)

	_, err := pub.Publish(context.Background(), repository.ReceiverOperation{
		PartitionID: 1,
		BucketID:    1,
		OperationID: "op-success",
	})
	if err != nil {
		t.Fatalf("Publish() error = %v", err)
	}
	if counters.publishSuccess != 1 || counters.publishError != 0 {
		t.Fatalf("counters success=%d error=%d, want 1/0", counters.publishSuccess, counters.publishError)
	}
}

func TestKafkaReceiverOperationPublisherCountsAndSanitizesProducerError(t *testing.T) {
	producer := &fakeSyncProducer{err: errors.New("broker 10.0.0.7 unavailable")}
	counters := &fakeProducerCounters{}
	pub := NewKafkaReceiverOperationPublisherForTest("topic", producer).WithCounters(counters)

	_, err := pub.Publish(context.Background(), repository.ReceiverOperation{
		PartitionID: 1,
		BucketID:    1,
		OperationID: "op-error",
	})
	if err == nil {
		t.Fatal("expected publish error")
	}
	if strings.Contains(err.Error(), "broker 10.0.0.7 unavailable") {
		t.Fatalf("error leaked raw kafka detail: %v", err)
	}
	if !strings.Contains(err.Error(), "receiver operation publish failed") {
		t.Fatalf("error = %v, want sanitized publish failure", err)
	}
	if counters.publishSuccess != 0 || counters.publishError != 1 {
		t.Fatalf("counters success=%d error=%d, want 0/1", counters.publishSuccess, counters.publishError)
	}
}

func TestKafkaReceiverOperationPublisherKeepsAckWhenContextCanceledAfterSend(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	producer := &fakeSyncProducer{
		partition: 7,
		offset:    42,
		onSend:    cancel,
	}
	pub := NewKafkaReceiverOperationPublisherForTest("topic", producer)

	ack, err := pub.Publish(ctx, repository.ReceiverOperation{
		PartitionID: 7,
		BucketID:    7,
		OperationID: "op-acked",
	})
	if err != nil {
		t.Fatalf("Publish() error = %v", err)
	}
	if ack.Topic != "topic" || ack.Partition != 7 || ack.Offset != 42 {
		t.Fatalf("ack = %+v", ack)
	}
}

func TestKafkaReceiverOperationPublisherCountsCloseError(t *testing.T) {
	producer := &fakeSyncProducer{closeErr: errors.New("close failed")}
	counters := &fakeProducerCounters{}
	pub := NewKafkaReceiverOperationPublisherForTest("topic", producer).WithCounters(counters)

	err := pub.Close()
	if err == nil {
		t.Fatal("expected close error")
	}
	if counters.closeError != 1 {
		t.Fatalf("close error counter = %d, want 1", counters.closeError)
	}
}
