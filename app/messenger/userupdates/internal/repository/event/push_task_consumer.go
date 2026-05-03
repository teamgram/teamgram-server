package event

import (
	"context"

	kafka "github.com/teamgram/marmota/pkg/mq"
)

type PushTaskKafkaHandler interface {
	HandlePushTaskKafkaRecord(ctx context.Context, record PushTaskKafkaRecord) error
}

type pushTaskReceiverAdapter struct {
	handler PushTaskKafkaHandler
}

func NewPushTaskConsumer(c *kafka.KafkaConsumerConf, handler PushTaskKafkaHandler) (*ReceiverConsumer, error) {
	return NewReceiverConsumer(c, pushTaskReceiverAdapter{handler: handler})
}

func (a pushTaskReceiverAdapter) HandleReceiverKafkaRecord(ctx context.Context, record ReceiverKafkaRecord) error {
	return a.handler.HandlePushTaskKafkaRecord(ctx, PushTaskKafkaRecord(record))
}
