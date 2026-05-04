package event

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/IBM/sarama"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverKafkaRecord = repository.ReceiverKafkaRecord

type ReceiverKafkaHandler interface {
	HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error
}

type ReceiverPartitionClaimer interface {
	ClaimPartitionOwner(ctx context.Context, partitionID int32) (int64, error)
}

type receiverConsumerGroup interface {
	Consume(context.Context, []string, sarama.ConsumerGroupHandler) error
	Errors() <-chan error
	Close() error
	Pause(map[string][]int32)
	Resume(map[string][]int32)
	PauseAll()
	ResumeAll()
}

type ReceiverConsumerCounters interface {
	IncConsumeLoopError()
	IncConsumeBackoff()
	IncRebalanceCount()
	IncMessageSuccess()
	IncMessageRetryable()
	IncMessageCommit()
	IncPanicRecovered()
}

type noopReceiverConsumerCounters struct{}

func (noopReceiverConsumerCounters) IncConsumeLoopError() {}
func (noopReceiverConsumerCounters) IncConsumeBackoff()   {}
func (noopReceiverConsumerCounters) IncRebalanceCount()   {}
func (noopReceiverConsumerCounters) IncMessageSuccess()   {}
func (noopReceiverConsumerCounters) IncMessageRetryable() {}
func (noopReceiverConsumerCounters) IncMessageCommit()    {}
func (noopReceiverConsumerCounters) IncPanicRecovered()   {}

type ReceiverConsumer struct {
	group            receiverConsumerGroup
	topics           []string
	handler          ReceiverKafkaHandler
	partitionClaimer ReceiverPartitionClaimer
	counters         ReceiverConsumerCounters
	mu               sync.Mutex
	attempt          int
	backoffBase      time.Duration
	backoffMax       time.Duration
	jitter           func(time.Duration) time.Duration
}

func NewReceiverConsumer(c *kafka.KafkaConsumerConf, handler ReceiverKafkaHandler) (*ReceiverConsumer, error) {
	conf, err := kafka.BuildConsumerGroupConfig(c, sarama.OffsetOldest, false)
	if err != nil {
		return nil, err
	}
	group, err := kafka.NewConsumerGroup(conf, c.Brokers, c.Group)
	if err != nil {
		return nil, err
	}
	claimer, _ := handler.(ReceiverPartitionClaimer)
	return &ReceiverConsumer{
		group:            group,
		topics:           c.Topics,
		handler:          handler,
		partitionClaimer: claimer,
		counters:         noopReceiverConsumerCounters{},
		backoffBase:      100 * time.Millisecond,
		backoffMax:       5 * time.Second,
		jitter:           defaultReceiverBackoffJitter,
	}, nil
}

func (c *ReceiverConsumer) WithCounters(counters ReceiverConsumerCounters) *ReceiverConsumer {
	if counters == nil {
		c.counters = noopReceiverConsumerCounters{}
		return c
	}
	c.counters = counters
	return c
}

func (c *ReceiverConsumer) ensureDefaults() {
	if c.counters == nil {
		c.counters = noopReceiverConsumerCounters{}
	}
	if c.backoffBase <= 0 {
		c.backoffBase = 100 * time.Millisecond
	}
	if c.backoffMax <= 0 {
		c.backoffMax = 5 * time.Second
	}
	if c.jitter == nil {
		c.jitter = defaultReceiverBackoffJitter
	}
}

func defaultReceiverBackoffJitter(delay time.Duration) time.Duration {
	if delay <= 0 {
		return 0
	}
	min := float64(delay) * 0.8
	max := float64(delay) * 1.2
	return time.Duration(min + rand.Float64()*(max-min))
}

func (c *ReceiverConsumer) nextBackoff() time.Duration {
	c.ensureDefaults()
	c.mu.Lock()
	defer c.mu.Unlock()
	delay := c.backoffBase
	for i := 0; i < c.attempt; i++ {
		delay *= 2
		if delay >= c.backoffMax {
			delay = c.backoffMax
			break
		}
	}
	c.attempt++
	return c.jitter(delay)
}

func (c *ReceiverConsumer) resetBackoff() {
	c.mu.Lock()
	c.attempt = 0
	c.mu.Unlock()
}

func (c *ReceiverConsumer) Start(ctx context.Context) error {
	c.ensureDefaults()
	for {
		if err := c.group.Consume(ctx, c.topics, c); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}
			c.counters.IncConsumeLoopError()
			delay := c.nextBackoff()
			c.counters.IncConsumeBackoff()
			logx.WithContext(ctx).Errorf("receiver consumer consume failed; retrying after %s: %v", delay, err)
			timer := time.NewTimer(delay)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
			continue
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *ReceiverConsumer) Close() error {
	if c == nil || c.group == nil {
		return nil
	}
	return c.group.Close()
}

func (c *ReceiverConsumer) Run(ctx context.Context) {
	if err := c.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		logx.WithContext(ctx).Errorf("receiver consumer stopped: %v", err)
	}
}

func (c *ReceiverConsumer) Stop() {
	_ = c.Close()
}

func (c *ReceiverConsumer) Setup(session sarama.ConsumerGroupSession) error {
	c.ensureDefaults()
	c.resetBackoff()
	c.counters.IncRebalanceCount()
	if c.partitionClaimer != nil {
		for topic, partitions := range session.Claims() {
			for _, partition := range partitions {
				epoch, err := c.partitionClaimer.ClaimPartitionOwner(session.Context(), partition)
				if err != nil {
					return fmt.Errorf("claim receiver partition owner: topic=%s partition=%d: %w", topic, partition, err)
				}
				logx.WithContext(session.Context()).Infow("receiver consumer partition claimed",
					logx.Field("topic", topic),
					logx.Field("partition", partition),
					logx.Field("owner_epoch", epoch))
			}
		}
	}
	logx.WithContext(session.Context()).Infow("receiver consumer session setup",
		logx.Field("member_id", session.MemberID()),
		logx.Field("generation_id", session.GenerationID()),
		logx.Field("claims", session.Claims()))
	return nil
}

func (c *ReceiverConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	logx.WithContext(session.Context()).Infow("receiver consumer session cleanup",
		logx.Field("member_id", session.MemberID()),
		logx.Field("generation_id", session.GenerationID()),
		logx.Field("claims", session.Claims()))
	return nil
}

func (c *ReceiverConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	c.ensureDefaults()
	var current *sarama.ConsumerMessage
	defer func() {
		if recovered := recover(); recovered != nil {
			c.counters.IncPanicRecovered()
			if current != nil {
				err = fmt.Errorf("receiver consumer panic: topic=%s partition=%d offset=%d: %v", current.Topic, current.Partition, current.Offset, recovered)
			} else {
				err = fmt.Errorf("receiver consumer panic: topic=%s partition=%d offset=%d: %v", claim.Topic(), claim.Partition(), claim.InitialOffset(), recovered)
			}
			logx.WithContext(session.Context()).Errorf("%v", err)
		}
	}()
	for message := range claim.Messages() {
		current = message
		record := ReceiverKafkaRecord{
			Topic:     message.Topic,
			Partition: message.Partition,
			Offset:    message.Offset,
			Key:       message.Key,
			Value:     message.Value,
		}
		if err := c.handler.HandleReceiverKafkaRecord(session.Context(), record); err != nil {
			c.counters.IncMessageRetryable()
			wrapped := fmt.Errorf("receiver consumer: topic=%s partition=%d offset=%d: %w", message.Topic, message.Partition, message.Offset, err)
			logx.WithContext(session.Context()).Errorf("%v", wrapped)
			return wrapped
		}
		c.counters.IncMessageSuccess()
		session.MarkMessage(message, "")
		session.Commit()
		c.counters.IncMessageCommit()
		logx.WithContext(session.Context()).Infow("receiver operation offset committed",
			logx.Field("topic", message.Topic),
			logx.Field("partition", message.Partition),
			logx.Field("offset", message.Offset))
	}
	return nil
}
