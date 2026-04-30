package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/IBM/sarama"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiverKafkaRecord = repository.ReceiverKafkaRecord

type ReceiverKafkaHandler interface {
	HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error
}

type ReceiverConsumer struct {
	group   sarama.ConsumerGroup
	topics  []string
	handler ReceiverKafkaHandler
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
	return &ReceiverConsumer{group: group, topics: c.Topics, handler: handler}, nil
}

func (c *ReceiverConsumer) Start(ctx context.Context) error {
	for {
		if err := c.group.Consume(ctx, c.topics, c); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return err
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

func (c *ReceiverConsumer) Setup(sarama.ConsumerGroupSession) error { return nil }

func (c *ReceiverConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c *ReceiverConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		record := ReceiverKafkaRecord{
			Topic:     message.Topic,
			Partition: message.Partition,
			Offset:    message.Offset,
			Key:       message.Key,
			Value:     message.Value,
		}
		if err := c.handler.HandleReceiverKafkaRecord(session.Context(), record); err != nil {
			wrapped := fmt.Errorf("receiver consumer: topic=%s partition=%d offset=%d: %w", message.Topic, message.Partition, message.Offset, err)
			logx.WithContext(session.Context()).Errorf("%v", wrapped)
			return wrapped
		}
		session.MarkMessage(message, "")
		session.Commit()
		logx.WithContext(session.Context()).Infow("receiver operation offset committed",
			logx.Field("topic", message.Topic),
			logx.Field("partition", message.Partition),
			logx.Field("offset", message.Offset))
	}
	return nil
}
