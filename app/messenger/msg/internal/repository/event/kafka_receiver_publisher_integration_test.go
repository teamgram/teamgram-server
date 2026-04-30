//go:build integration && kafka

package event

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestKafkaReceiverOperationPublisherIntegration(t *testing.T) {
	brokers := os.Getenv("TEAMGRAM_TEST_KAFKA_BROKERS")
	if brokers == "" {
		t.Skip("TEAMGRAM_TEST_KAFKA_BROKERS is empty")
	}
	topic := os.Getenv("TEAMGRAM_TEST_KAFKA_RECEIVER_TOPIC")
	if topic == "" {
		topic = "teamgram.test.receiver_operations"
	}
	publisher, err := NewKafkaReceiverOperationPublisher(&kafka.KafkaProducerConf{
		Brokers:     strings.Split(brokers, ","),
		Topic:       topic,
		ProducerAck: "wait_for_all",
	})
	if err != nil {
		t.Fatalf("NewKafkaReceiverOperationPublisher() error = %v", err)
	}
	defer publisher.Close()

	userID := time.Now().UnixNano() % 1_000_000_000
	route := payload.RouteUser(userID)
	body := []byte(`{"schema_version":1}`)
	op := repository.ReceiverOperation{
		UserID:       userID,
		BucketID:     int32(route.BucketID),
		PartitionID:  int32(route.ReceiverPartitionID),
		OperationID:  payload.ReceiverOperationID(userID*10+1, userID),
		OpType:       payload.OpTypeSendMessage,
		PeerType:     payload.PeerTypeUser,
		PeerID:       userID,
		PayloadCodec: payload.PayloadCodecJSON,
		Payload:      body,
		PayloadHash:  payload.HashBytes(body),
	}
	ack, err := publisher.Publish(context.Background(), op)
	if err != nil {
		t.Fatalf("Publish() error = %v", err)
	}
	if ack.Topic != topic || ack.Partition != op.PartitionID || ack.Offset < 0 {
		t.Fatalf("ack = %+v, want topic=%s partition=%d", ack, topic, op.PartitionID)
	}
}
