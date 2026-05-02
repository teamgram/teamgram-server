package mtproto

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

type PlainMessage struct {
	MsgId int64
	Body  []byte
}

func DecodePlainMessage(payload []byte) (PlainMessage, error) {
	d := bin.NewDecoder(payload)
	authKeyId, err := d.Int64()
	if err != nil {
		return PlainMessage{}, fmt.Errorf("decode plain message auth_key_id: %w", err)
	}
	if authKeyId != 0 {
		return PlainMessage{}, fmt.Errorf("decode plain message: auth_key_id = %d, want 0", authKeyId)
	}
	msgId, err := d.Int64()
	if err != nil {
		return PlainMessage{}, fmt.Errorf("decode plain message msg_id: %w", err)
	}
	bodyLen, err := d.Int32()
	if err != nil {
		return PlainMessage{}, fmt.Errorf("decode plain message body length: %w", err)
	}
	if bodyLen < 0 || int(bodyLen) > d.Remaining() {
		return PlainMessage{}, fmt.Errorf("decode plain message: invalid body length %d", bodyLen)
	}
	body := make([]byte, bodyLen)
	if err := d.ConsumeN(body, int(bodyLen)); err != nil {
		return PlainMessage{}, fmt.Errorf("decode plain message body: %w", err)
	}

	return PlainMessage{MsgId: msgId, Body: body}, nil
}

func EncodePlainMessage(msg PlainMessage) ([]byte, error) {
	if msg.Body == nil {
		return nil, fmt.Errorf("encode plain message: body is nil")
	}
	x := bin.NewEncoder()
	defer x.End()
	x.PutInt64(0)
	x.PutInt64(msg.MsgId)
	x.PutInt32(int32(len(msg.Body)))
	x.PutRaw(msg.Body)
	return append([]byte(nil), x.Bytes()...), nil
}
