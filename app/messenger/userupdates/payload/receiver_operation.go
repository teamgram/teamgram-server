package payload

import (
	"encoding/json"
	"fmt"
)

const ReceiverKafkaMessageSchemaVersion = 1

type ReceiverOperationEnvelopeV1 struct {
	UserID        int64
	BucketID      int32
	PartitionID   int32
	OperationID   string
	OpType        int32
	PeerType      int32
	PeerID        int64
	PayloadCodec  int32
	Payload       []byte
	PayloadHash   []byte
	DependencyPts []int64
}

type ReceiverKafkaMessageV1 struct {
	SchemaVersion     int32                       `json:"schema_version"`
	Operation         ReceiverOperationEnvelopeV1 `json:"operation"`
	Attempt           int32                       `json:"attempt"`
	OriginalTopic     string                      `json:"original_topic,omitempty"`
	OriginalPartition int32                       `json:"original_partition,omitempty"`
	OriginalOffset    int64                       `json:"original_offset,omitempty"`
}

func MarshalReceiverKafkaMessage(msg ReceiverKafkaMessageV1) ([]byte, error) {
	if msg.SchemaVersion == 0 {
		msg.SchemaVersion = ReceiverKafkaMessageSchemaVersion
	}
	return json.Marshal(msg)
}

func UnmarshalReceiverKafkaMessage(body []byte) (*ReceiverKafkaMessageV1, error) {
	var msg ReceiverKafkaMessageV1
	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, err
	}
	if msg.SchemaVersion != ReceiverKafkaMessageSchemaVersion {
		return nil, fmt.Errorf("unknown receiver kafka message schema_version=%d", msg.SchemaVersion)
	}
	return &msg, nil
}

const PushTaskKafkaMessageSchemaVersion = 1

type PushTaskKafkaMessageV1 struct {
	SchemaVersion int32  `json:"schema_version"`
	TaskID        int64  `json:"task_id"`
	UserID        int64  `json:"user_id"`
	Pts           int64  `json:"pts"`
	PushType      int32  `json:"push_type"`
	PeerType      int32  `json:"peer_type"`
	PeerID        int64  `json:"peer_id"`
	OperationID   string `json:"operation_id"`
	Payload       []byte `json:"payload"`
}

func MarshalPushTaskKafkaMessage(msg PushTaskKafkaMessageV1) ([]byte, error) {
	if msg.SchemaVersion == 0 {
		msg.SchemaVersion = PushTaskKafkaMessageSchemaVersion
	}
	return json.Marshal(msg)
}

func UnmarshalPushTaskKafkaMessage(body []byte) (*PushTaskKafkaMessageV1, error) {
	var msg PushTaskKafkaMessageV1
	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, err
	}
	if msg.SchemaVersion != PushTaskKafkaMessageSchemaVersion {
		return nil, fmt.Errorf("unknown push task kafka message schema_version=%d", msg.SchemaVersion)
	}
	return &msg, nil
}
