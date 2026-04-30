package event

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/IBM/sarama"
)

type fakeReceiverHandler struct {
	err     error
	calls   int
	records []ReceiverKafkaRecord
}

func (h *fakeReceiverHandler) HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error {
	h.calls++
	h.records = append(h.records, record)
	return h.err
}

type fakeSession struct {
	ctx           context.Context
	marked        int
	committed     int
	markedMessage *sarama.ConsumerMessage
}

func newFakeSession() *fakeSession {
	return &fakeSession{ctx: context.Background()}
}

func (s *fakeSession) Claims() map[string][]int32               { return map[string][]int32{"topic": {0}} }
func (s *fakeSession) MemberID() string                         { return "member-1" }
func (s *fakeSession) GenerationID() int32                      { return 1 }
func (s *fakeSession) MarkOffset(string, int32, int64, string)  {}
func (s *fakeSession) Commit()                                  { s.committed++ }
func (s *fakeSession) ResetOffset(string, int32, int64, string) {}
func (s *fakeSession) MarkMessage(msg *sarama.ConsumerMessage, metadata string) {
	s.marked++
	s.markedMessage = msg
}
func (s *fakeSession) Context() context.Context { return s.ctx }

type fakeClaim struct {
	topic     string
	partition int32
	offset    int64
	messages  chan *sarama.ConsumerMessage
}

func newFakeClaim(messages ...*sarama.ConsumerMessage) *fakeClaim {
	ch := make(chan *sarama.ConsumerMessage, len(messages))
	for _, message := range messages {
		ch <- message
	}
	close(ch)
	return &fakeClaim{
		topic:     "topic",
		partition: 0,
		offset:    0,
		messages:  ch,
	}
}

func (c *fakeClaim) Topic() string                            { return c.topic }
func (c *fakeClaim) Partition() int32                         { return c.partition }
func (c *fakeClaim) InitialOffset() int64                     { return c.offset }
func (c *fakeClaim) HighWaterMarkOffset() int64               { return c.offset + int64(len(c.messages)) }
func (c *fakeClaim) Messages() <-chan *sarama.ConsumerMessage { return c.messages }

func TestReceiverConsumerCommitsAfterHandlerSuccess(t *testing.T) {
	handler := &fakeReceiverHandler{}
	consumer := &ReceiverConsumer{handler: handler}
	session := newFakeSession()
	message := &sarama.ConsumerMessage{Topic: "topic", Partition: 3, Offset: 42, Key: []byte("k"), Value: []byte("v")}

	if err := consumer.ConsumeClaim(session, newFakeClaim(message)); err != nil {
		t.Fatalf("ConsumeClaim() error = %v", err)
	}
	if handler.calls != 1 {
		t.Fatalf("handler calls = %d, want 1", handler.calls)
	}
	if session.marked != 1 || session.committed != 1 {
		t.Fatalf("marked=%d committed=%d, want 1/1", session.marked, session.committed)
	}
	if got := handler.records[0]; got.Topic != "topic" || got.Partition != 3 || got.Offset != 42 {
		t.Fatalf("record = %+v", got)
	}
}

func TestReceiverConsumerDoesNotCommitRetryableError(t *testing.T) {
	handler := &fakeReceiverHandler{err: errors.New("db timeout")}
	consumer := &ReceiverConsumer{handler: handler}
	session := newFakeSession()
	message := &sarama.ConsumerMessage{Topic: "topic", Partition: 3, Offset: 42, Value: []byte("v")}

	err := consumer.ConsumeClaim(session, newFakeClaim(message))
	if err == nil {
		t.Fatal("ConsumeClaim() expected error")
	}
	if !strings.Contains(err.Error(), "topic=topic") || !strings.Contains(err.Error(), "partition=3") || !strings.Contains(err.Error(), "offset=42") {
		t.Fatalf("error lacks kafka context: %v", err)
	}
	if session.marked != 0 || session.committed != 0 {
		t.Fatalf("marked=%d committed=%d, want 0/0", session.marked, session.committed)
	}
}

func TestReceiverConsumerContinuesAfterHandlerSwallowsTerminal(t *testing.T) {
	handler := &fakeReceiverHandler{}
	consumer := &ReceiverConsumer{handler: handler}
	session := newFakeSession()
	message := &sarama.ConsumerMessage{Topic: "topic", Partition: 3, Offset: 42, Value: []byte("terminal already in dlq")}

	if err := consumer.ConsumeClaim(session, newFakeClaim(message)); err != nil {
		t.Fatalf("ConsumeClaim() error = %v", err)
	}
	if session.marked != 1 || session.committed != 1 {
		t.Fatalf("marked=%d committed=%d, want 1/1", session.marked, session.committed)
	}
}
