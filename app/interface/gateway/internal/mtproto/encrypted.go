package mtproto

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

var lastServerMsgId atomic.Int64

const minEncryptedPayloadLen = 8 + 16 + 32

type EncryptedMessage struct {
	AuthKeyId int64
	Salt      int64
	SessionId int64
	MsgId     int64
	SeqNo     int32
	Body      []byte
}

func DecodeEncryptedMessage(payload []byte, key *crypto.AuthKey) (EncryptedMessage, error) {
	if key == nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message: auth key is nil")
	}
	if len(payload) < minEncryptedPayloadLen {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message: payload too short %d", len(payload))
	}

	d := bin.NewDecoder(payload)
	authKeyId, err := d.Int64()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message auth_key_id: %w", err)
	}
	if authKeyId != key.AuthKeyId() {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message: auth_key_id = %d, want %d", authKeyId, key.AuthKeyId())
	}

	msgKey := make([]byte, 16)
	if err := d.ConsumeN(msgKey, len(msgKey)); err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message msg_key: %w", err)
	}
	encrypted := d.Raw()
	if len(encrypted)%16 != 0 {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message: encrypted payload length %d is not aes block aligned", len(encrypted))
	}
	raw, err := key.AesIgeDecrypt(msgKey, encrypted)
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message decrypt: %w", err)
	}

	inner := bin.NewDecoder(raw)
	salt, err := inner.Int64()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message salt: %w", err)
	}
	sessionId, err := inner.Int64()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message session_id: %w", err)
	}
	msgId, err := inner.Int64()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message msg_id: %w", err)
	}
	seqNo, err := inner.Int32()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message seq_no: %w", err)
	}
	bodyLen, err := inner.Int32()
	if err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message body length: %w", err)
	}
	if bodyLen < 0 || int(bodyLen) > inner.Remaining() {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message: invalid body length %d", bodyLen)
	}
	body := make([]byte, bodyLen)
	if err := inner.ConsumeN(body, int(bodyLen)); err != nil {
		return EncryptedMessage{}, fmt.Errorf("decode encrypted message body: %w", err)
	}

	return EncryptedMessage{
		AuthKeyId: authKeyId,
		Salt:      salt,
		SessionId: sessionId,
		MsgId:     msgId,
		SeqNo:     seqNo,
		Body:      body,
	}, nil
}

func EncodeEncryptedMessage(msg EncryptedMessage, key *crypto.AuthKey) ([]byte, error) {
	if key == nil {
		return nil, fmt.Errorf("encode encrypted message: auth key is nil")
	}
	if msg.AuthKeyId != key.AuthKeyId() {
		return nil, fmt.Errorf("encode encrypted message: auth_key_id = %d, want %d", msg.AuthKeyId, key.AuthKeyId())
	}
	if msg.Body == nil {
		return nil, fmt.Errorf("encode encrypted message: body is nil")
	}

	x := bin.NewEncoder()
	defer x.End()
	x.PutInt64(msg.Salt)
	x.PutInt64(msg.SessionId)
	x.PutInt64(msg.MsgId)
	x.PutInt32(msg.SeqNo)
	x.PutInt32(int32(len(msg.Body)))
	x.PutRaw(msg.Body)

	msgKey, encrypted, err := key.AesIgeEncrypt(x.Bytes())
	if err != nil {
		return nil, fmt.Errorf("encode encrypted message encrypt: %w", err)
	}

	out := bin.NewEncoder()
	defer out.End()
	out.PutInt64(msg.AuthKeyId)
	out.PutRaw(msgKey)
	out.PutRaw(encrypted)
	return append([]byte(nil), out.Bytes()...), nil
}

func NextServerMsgId(after int64) int64 {
	now := time.Now()
	candidate := (now.Unix() << 32) + ((int64(now.Nanosecond()) << 32) / int64(time.Second))
	candidate = (candidate &^ 3) | 1
	if candidate <= after {
		candidate = (after &^ 3) | 1
		if candidate <= after {
			candidate += 4
		}
	}

	for {
		last := lastServerMsgId.Load()
		next := candidate
		if next <= last {
			next = (last &^ 3) + 5
		}
		if lastServerMsgId.CompareAndSwap(last, next) {
			return next
		}
	}
}
