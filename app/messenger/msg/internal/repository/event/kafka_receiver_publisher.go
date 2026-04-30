package event

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type syncProducer interface {
	SendMessage(*sarama.ProducerMessage) (int32, int64, error)
	Close() error
}

type KafkaReceiverOperationPublisher struct {
	topic    string
	producer syncProducer
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
	return &KafkaReceiverOperationPublisher{topic: c.Topic, producer: producer}, nil
}

func NewKafkaReceiverOperationPublisherForTest(topic string, producer syncProducer) *KafkaReceiverOperationPublisher {
	return &KafkaReceiverOperationPublisher{topic: topic, producer: producer}
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
		return repository.KafkaAck{}, err
	}
	if err := ctx.Err(); err != nil {
		return repository.KafkaAck{}, err
	}
	return repository.KafkaAck{Topic: p.topic, Partition: partition, Offset: offset}, nil
}

func (p *KafkaReceiverOperationPublisher) Close() error {
	if p == nil || p.producer == nil {
		return nil
	}
	return p.producer.Close()
}
