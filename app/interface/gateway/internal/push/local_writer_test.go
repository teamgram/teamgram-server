package push

import (
	"bytes"
	"context"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
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
	writer.Unregister(key.AuthKeyId(), 0, 22)

	ok, err := writer.WriteRPCResult(context.Background(), key.AuthKeyId(), 22, 100, []byte{1})
	if err != nil {
		t.Fatalf("WriteRPCResult() error = %v", err)
	}
	if ok {
		t.Fatal("WriteRPCResult() ok = true, want false")
	}
}

func TestGatewayPushLocalWriterSeparatesAuthKeyType(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	normal := &fakeSessionWriter{}
	media := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: key, Writer: normal})
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeMediaTemp, SessionId: 22, AuthKey: key, Writer: media})

	ok, err := writer.WriteRPCResult(context.Background(), key.AuthKeyId(), 22, 100, []byte{9})
	if err != nil {
		t.Fatalf("WriteRPCResult() error = %v", err)
	}
	if ok {
		t.Fatal("ambiguous WriteRPCResult() ok = true, want false")
	}
	writer.Unregister(key.AuthKeyId(), tg.AuthKeyTypeMediaTemp, 22)
	ok, err = writer.WriteRPCResult(context.Background(), key.AuthKeyId(), 22, 102, []byte{8})
	if err != nil {
		t.Fatalf("normal WriteRPCResult() error = %v", err)
	}
	if !ok {
		t.Fatal("normal WriteRPCResult() ok = false after media unregister")
	}
	if normal.seq != 1 || media.seq != 0 {
		t.Fatalf("normal seq=%d media seq=%d", normal.seq, media.seq)
	}
}

func TestGatewayPushLocalWriterWriteUpdatesByPermAuthKey(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	tempKey := crypto.CreateAuthKey()
	matching := &fakeSessionWriter{}
	media := &fakeSessionWriter{}
	other := &fakeSessionWriter{}
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: tempKey, Layer: 223, Writer: matching})
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeMediaTemp, SessionId: 23, AuthKey: tempKey, Layer: 223, Writer: media})
	writer.Register(LocalTarget{PermAuthKeyId: 999, AuthKeyId: 999, AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 24, AuthKey: crypto.CreateAuthKey(), Layer: 223, Writer: other})
	updates := tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), updates)
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("WriteUpdates() count = %d, want 1", count)
	}
	wantBody, err := iface.EncodeObject(updates, 223)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	if matching.msg.AuthKeyId != tempKey.AuthKeyId() || matching.msg.SessionId != 22 || matching.msg.SeqNo != 1 || !bytes.Equal(matching.msg.Body, wantBody) {
		t.Fatalf("matching written msg = %#v", matching.msg)
	}
	if media.seq != 0 || other.seq != 0 {
		t.Fatalf("media seq=%d other seq=%d, want no writes", media.seq, other.seq)
	}
}

func TestGatewayPushLocalWriterWriteUpdatesDeduplicatesSessionsByAuthKey(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	tempKey := crypto.CreateAuthKey()
	first := &fakeSessionWriter{}
	second := &fakeSessionWriter{}
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: tempKey, Layer: 223, Writer: first})
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 23, AuthKey: tempKey, Layer: 223, Writer: second})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("WriteUpdates() count = %d, want 1", count)
	}
	if first.seq+second.seq != 1 {
		t.Fatalf("first seq=%d second seq=%d, want one write", first.seq, second.seq)
	}
}

func TestGatewayPushLocalWriterWriteSessionUpdates(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	sw := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: key, Layer: 223, Writer: sw})
	updates := tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})

	ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, updates)
	if err != nil {
		t.Fatalf("WriteSessionUpdates() error = %v", err)
	}
	if !ok {
		t.Fatal("WriteSessionUpdates() ok = false")
	}
	wantBody, err := iface.EncodeObject(updates, 223)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	if sw.msg.AuthKeyId != key.AuthKeyId() || sw.msg.SessionId != 22 || sw.msg.SeqNo != 1 || !bytes.Equal(sw.msg.Body, wantBody) {
		t.Fatalf("written msg = %#v", sw.msg)
	}
}
