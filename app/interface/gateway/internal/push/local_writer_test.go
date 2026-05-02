package push

import (
	"bytes"
	"context"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
)

type fakeSessionWriter struct {
	seq int32
	msg gmtproto.EncryptedMessage
}

func (w *fakeSessionWriter) NextSeqNo(contentRelated bool) int32 {
	w.seq++
	return w.seq
}

func (w *fakeSessionWriter) WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error {
	w.msg = msg
	return nil
}

func TestGatewayPushLocalWriterWriteRPCResult(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	sw := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), SessionId: 22, AuthKey: key, Writer: sw})
	body := []byte{1, 2, 3, 4}

	ok, err := writer.WriteRPCResult(context.Background(), key.AuthKeyId(), 22, 100, body)
	if err != nil {
		t.Fatalf("WriteRPCResult() error = %v", err)
	}
	if !ok {
		t.Fatal("WriteRPCResult() ok = false")
	}
	if sw.msg.AuthKeyId != key.AuthKeyId() || sw.msg.SessionId != 22 || sw.msg.MsgId <= 100 || sw.msg.SeqNo != 1 || !bytes.Equal(sw.msg.Body, body) {
		t.Fatalf("written msg = %#v", sw.msg)
	}
}

func TestGatewayPushLocalWriterUnregister(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), SessionId: 22, AuthKey: key, Writer: &fakeSessionWriter{}})
	writer.Unregister(key.AuthKeyId(), 22)

	ok, err := writer.WriteRPCResult(context.Background(), key.AuthKeyId(), 22, 100, []byte{1})
	if err != nil {
		t.Fatalf("WriteRPCResult() error = %v", err)
	}
	if ok {
		t.Fatal("WriteRPCResult() ok = true, want false")
	}
}
