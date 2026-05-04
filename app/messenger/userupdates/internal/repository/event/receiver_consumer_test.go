package event

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
)

type fakeReceiverHandler struct {
	err     error
	calls   int
	records []ReceiverKafkaRecord
}

type fakeReceiverCounters struct {
	consumeLoopError int
	consumeBackoff   int
	rebalanceCount   int
	messageSuccess   int
	messageRetryable int
	messageCommit    int
	panicRecovered   int
}

func (c *fakeReceiverCounters) IncConsumeLoopError() { c.consumeLoopError++ }
func (c *fakeReceiverCounters) IncConsumeBackoff()   { c.consumeBackoff++ }
func (c *fakeReceiverCounters) IncRebalanceCount()   { c.rebalanceCount++ }
func (c *fakeReceiverCounters) IncMessageSuccess()   { c.messageSuccess++ }
func (c *fakeReceiverCounters) IncMessageRetryable() { c.messageRetryable++ }
func (c *fakeReceiverCounters) IncMessageCommit()    { c.messageCommit++ }
func (c *fakeReceiverCounters) IncPanicRecovered()   { c.panicRecovered++ }

type panicReceiverHandler struct{}

func (panicReceiverHandler) HandleReceiverKafkaRecord(context.Context, ReceiverKafkaRecord) error {
	panic("receiver handler panic")
}

func (h *fakeReceiverHandler) HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error {
	h.calls++
	h.records = append(h.records, record)
	return h.err
}

type fakeSession struct {
	ctx           context.Context
	claims        map[string][]int32
	marked        int
	committed     int
	markedMessage *sarama.ConsumerMessage
}

func newFakeSession() *fakeSession {
	return &fakeSession{
		ctx:    context.Background(),
		claims: map[string][]int32{"topic": {0}},
	}
}

func (s *fakeSession) Claims() map[string][]int32               { return s.claims }
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

func TestReceiverConsumerRecoverPanicWithoutCommit(t *testing.T) {
	counters := &fakeReceiverCounters{}
	consumer := (&ReceiverConsumer{handler: panicReceiverHandler{}}).WithCounters(counters)
	session := newFakeSession()
	message := &sarama.ConsumerMessage{Topic: "topic", Partition: 3, Offset: 42, Value: []byte("v")}

	err := consumer.ConsumeClaim(session, newFakeClaim(message))
	if err == nil {
		t.Fatal("ConsumeClaim() expected panic recovery error")
	}
	if !strings.Contains(err.Error(), "receiver consumer panic") {
		t.Fatalf("error = %v, want panic recovery context", err)
	}
	if !strings.Contains(err.Error(), "offset=42") {
		t.Fatalf("error = %v, want current message offset", err)
	}
	if session.marked != 0 || session.committed != 0 {
		t.Fatalf("marked=%d committed=%d, want 0/0", session.marked, session.committed)
	}
	if counters.panicRecovered != 1 {
		t.Fatalf("panic counter = %d, want 1", counters.panicRecovered)
	}
}

func TestReceiverConsumerCountersForSuccessAndRetryable(t *testing.T) {
	successCounters := &fakeReceiverCounters{}
	successConsumer := (&ReceiverConsumer{handler: &fakeReceiverHandler{}}).WithCounters(successCounters)
	successSession := newFakeSession()
	message := &sarama.ConsumerMessage{Topic: "topic", Partition: 3, Offset: 42, Value: []byte("v")}
	if err := successConsumer.ConsumeClaim(successSession, newFakeClaim(message)); err != nil {
		t.Fatalf("success ConsumeClaim() error = %v", err)
	}
	if successCounters.messageSuccess != 1 || successCounters.messageCommit != 1 {
		t.Fatalf("success counters success=%d commit=%d, want 1/1", successCounters.messageSuccess, successCounters.messageCommit)
	}

	retryCounters := &fakeReceiverCounters{}
	retryConsumer := (&ReceiverConsumer{handler: &fakeReceiverHandler{err: errors.New("db timeout")}}).WithCounters(retryCounters)
	retrySession := newFakeSession()
	err := retryConsumer.ConsumeClaim(retrySession, newFakeClaim(message))
	if err == nil {
		t.Fatal("retryable ConsumeClaim() expected error")
	}
	if retryCounters.messageRetryable != 1 || retryCounters.messageCommit != 0 {
		t.Fatalf("retry counters retryable=%d commit=%d, want 1/0", retryCounters.messageRetryable, retryCounters.messageCommit)
	}
}

func TestReceiverConsumerSetupCleanupCounters(t *testing.T) {
	counters := &fakeReceiverCounters{}
	consumer := (&ReceiverConsumer{}).WithCounters(counters)
	session := newFakeSession()

	if err := consumer.Setup(session); err != nil {
		t.Fatalf("Setup() error = %v", err)
	}
	if err := consumer.Cleanup(session); err != nil {
		t.Fatalf("Cleanup() error = %v", err)
	}
	if counters.rebalanceCount != 1 {
		t.Fatalf("rebalance counter = %d, want 1", counters.rebalanceCount)
	}
}

type fakePartitionClaimer struct {
	claimed []int32
	err     error
}

func (c *fakePartitionClaimer) ClaimPartitionOwner(_ context.Context, partitionID int32) (int64, error) {
	c.claimed = append(c.claimed, partitionID)
	return int64(len(c.claimed)), c.err
}

func TestReceiverConsumerSetupClaimsAssignedPartitions(t *testing.T) {
	claimer := &fakePartitionClaimer{}
	consumer := &ReceiverConsumer{partitionClaimer: claimer}
	session := newFakeSession()
	session.claims = map[string][]int32{
		"userupdates.receiver_operations.v1": {35, 58},
	}

	if err := consumer.Setup(session); err != nil {
		t.Fatalf("Setup() error = %v", err)
	}
	if len(claimer.claimed) != 2 || claimer.claimed[0] != 35 || claimer.claimed[1] != 58 {
		t.Fatalf("claimed partitions = %v, want [35 58]", claimer.claimed)
	}
}

type fakeConsumerGroup struct {
	mu           sync.Mutex
	consumeErrs  []error
	consumeCalls int
	closed       bool
}

func (g *fakeConsumerGroup) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	g.mu.Lock()
	g.consumeCalls++
	idx := g.consumeCalls - 1
	g.mu.Unlock()
	if idx < len(g.consumeErrs) {
		return g.consumeErrs[idx]
	}
	return sarama.ErrClosedConsumerGroup
}

func (g *fakeConsumerGroup) Errors() <-chan error { return nil }

func (g *fakeConsumerGroup) Close() error {
	g.closed = true
	return nil
}

func (g *fakeConsumerGroup) Pause(map[string][]int32)  {}
func (g *fakeConsumerGroup) Resume(map[string][]int32) {}
func (g *fakeConsumerGroup) PauseAll()                 {}
func (g *fakeConsumerGroup) ResumeAll()                {}

func TestReceiverConsumerStartBacksOffAndRetriesConsumeErrors(t *testing.T) {
	counters := &fakeReceiverCounters{}
	group := &fakeConsumerGroup{consumeErrs: []error{errors.New("temporary kafka failure")}}
	consumer := (&ReceiverConsumer{
		group:       group,
		topics:      []string{"topic"},
		handler:     &fakeReceiverHandler{},
		backoffBase: time.Millisecond,
		backoffMax:  2 * time.Millisecond,
		jitter:      func(time.Duration) time.Duration { return 0 },
	}).WithCounters(counters)

	err := consumer.Start(context.Background())
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}
	if group.consumeCalls != 2 {
		t.Fatalf("consume calls = %d, want 2", group.consumeCalls)
	}
	if counters.consumeLoopError != 1 || counters.consumeBackoff != 1 {
		t.Fatalf("counters loop_error=%d backoff=%d, want 1/1", counters.consumeLoopError, counters.consumeBackoff)
	}
}

func TestReceiverConsumerBackoffSequenceAndReset(t *testing.T) {
	consumer := &ReceiverConsumer{
		backoffBase: 100 * time.Millisecond,
		backoffMax:  350 * time.Millisecond,
		jitter:      func(delay time.Duration) time.Duration { return delay },
	}

	if got := consumer.nextBackoff(); got != 100*time.Millisecond {
		t.Fatalf("first backoff = %s, want 100ms", got)
	}
	if got := consumer.nextBackoff(); got != 200*time.Millisecond {
		t.Fatalf("second backoff = %s, want 200ms", got)
	}
	if got := consumer.nextBackoff(); got != 350*time.Millisecond {
		t.Fatalf("third backoff = %s, want capped 350ms", got)
	}
	consumer.resetBackoff()
	if got := consumer.nextBackoff(); got != 100*time.Millisecond {
		t.Fatalf("reset backoff = %s, want 100ms", got)
	}
}
