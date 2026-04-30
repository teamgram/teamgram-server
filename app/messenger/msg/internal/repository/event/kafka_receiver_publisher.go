package event

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/zeromicro/go-zero/core/logx"
)

type syncProducer interface {
	SendMessage(*sarama.ProducerMessage) (int32, int64, error)
	Close() error
}

type ProducerCounters interface {
	IncPublishSuccess()
	IncPublishError()
	IncCloseError()
}

type noopProducerCounters struct{}

func (noopProducerCounters) IncPublishSuccess() {}
func (noopProducerCounters) IncPublishError()   {}
func (noopProducerCounters) IncCloseError()     {}

type KafkaReceiverOperationPublisher struct {
	topic    string
	producer syncProducer
	counters ProducerCounters
}

func NewKafkaReceiverOperationPublisher(c *kafka.KafkaProducerConf) (*KafkaReceiverOperationPublisher, error) {
	if c == nil || c.Topic == "" || len(c.Brokers) == 0 {
		return nil, fmt.Errorf("receiver operations kafka config is empty")
	}
	conf, err := kafka.BuildProducerConfig(*c)
	if err != nil {
		return nil, err
	}
	conf.Producer.Partitioner = sarama.NewManualPartitioner
	producer, err := kafka.NewProducer(conf, c.Brokers)
	if err != nil {
		return nil, err
	}
	return &KafkaReceiverOperationPublisher{topic: c.Topic, producer: producer, counters: noopProducerCounters{}}, nil
}

func NewKafkaReceiverOperationPublisherForTest(topic string, producer syncProducer) *KafkaReceiverOperationPublisher {
	return &KafkaReceiverOperationPublisher{topic: topic, producer: producer, counters: noopProducerCounters{}}
}

func (p *KafkaReceiverOperationPublisher) WithCounters(counters ProducerCounters) *KafkaReceiverOperationPublisher {
	if counters == nil {
		p.counters = noopProducerCounters{}
		return p
	}
	p.counters = counters
	return p
}

func (p *KafkaReceiverOperationPublisher) Publish(ctx context.Context, op repository.ReceiverOperation) (repository.KafkaAck, error) {
	if err := ctx.Err(); err != nil {
		return repository.KafkaAck{}, err
	}
	body, err := payload.MarshalReceiverKafkaMessage(payload.ReceiverKafkaMessageV1{
		SchemaVersion: payload.ReceiverKafkaMessageSchemaVersion,
		Operation:     op,
		Attempt:       0,
	})
	if err != nil {
		return repository.KafkaAck{}, err
	}
	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(strconv.FormatInt(int64(op.BucketID), 10)),
		Value:     sarama.ByteEncoder(body),
		Partition: op.PartitionID,
	}
	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		p.counters.IncPublishError()
		logx.WithContext(ctx).Errorf("receiver operation publish failed: operation_id=%s topic=%s partition=%d err=%v", op.OperationID, p.topic, op.PartitionID, err)
		return repository.KafkaAck{}, fmt.Errorf("receiver operation publish failed")
	}
	p.counters.IncPublishSuccess()
	return repository.KafkaAck{Topic: p.topic, Partition: partition, Offset: offset}, nil
}

func (p *KafkaReceiverOperationPublisher) Close() error {
	if p == nil || p.producer == nil {
		return nil
	}
	err := p.producer.Close()
	if err != nil {
		p.counters.IncCloseError()
	}
	return err
}
