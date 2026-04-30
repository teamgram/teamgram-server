package event

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

type syncProducer interface {
	SendMessage(*sarama.ProducerMessage) (int32, int64, error)
	Close() error
}

type PushTaskPublisher struct {
	topic    string
	producer syncProducer
}

func NewPushTaskPublisher(c *kafka.KafkaProducerConf) (*PushTaskPublisher, error) {
	if c == nil || c.Topic == "" || len(c.Brokers) == 0 {
		return nil, fmt.Errorf("push tasks kafka config is empty")
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
	return &PushTaskPublisher{topic: c.Topic, producer: producer}, nil
}

func NewPushTaskPublisherForTest(topic string, producer syncProducer) *PushTaskPublisher {
	return &PushTaskPublisher{topic: topic, producer: producer}
}

func (p *PushTaskPublisher) Publish(ctx context.Context, task repository.PushTask) (repository.KafkaAck, error) {
	if err := ctx.Err(); err != nil {
		return repository.KafkaAck{}, err
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        task.TaskID,
		UserID:        task.UserID,
		Pts:           task.Pts,
		PushType:      task.PushType,
		PeerType:      task.PeerType,
		PeerID:        task.PeerID,
		OperationID:   task.OperationID,
		Payload:       task.TaskPayload,
	})
	if err != nil {
		return repository.KafkaAck{}, err
	}
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(strconv.FormatInt(task.UserID, 10)),
		Value:     sarama.ByteEncoder(body),
		Partition: task.PushPartitionID,
	})
	if err != nil {
		return repository.KafkaAck{}, err
	}
	if err := ctx.Err(); err != nil {
		return repository.KafkaAck{}, err
	}
	return repository.KafkaAck{Topic: p.topic, Partition: partition, Offset: offset}, nil
}

func (p *PushTaskPublisher) Close() error {
	if p == nil || p.producer == nil {
		return nil
	}
	return p.producer.Close()
}
