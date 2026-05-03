package core

import (
	"bytes"
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type coreFakeSessionWriter struct {
	seq int32
	msg gmtproto.EncryptedMessage
}

func (w *coreFakeSessionWriter) NextSeqNo(contentRelated bool) int32 {
	w.seq++
	return w.seq
}

func (w *coreFakeSessionWriter) WriteEncrypted(ctx context.Context, msg gmtproto.EncryptedMessage) error {
	w.msg = msg
	return nil
}

func TestGatewayPushUpdatesDataWritesLocalUpdate(t *testing.T) {
	local := push.NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	tempKey := crypto.CreateAuthKey()
	writer := &coreFakeSessionWriter{}
	local.Register(push.LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, Layer: 223, AuthKey: tempKey, Writer: writer})
	core := New(context.Background(), &svc.ServiceContext{Push: local})
	updates := tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})

	ok, err := core.GatewayPushUpdatesData(&gateway.TLGatewayPushUpdatesData{PermAuthKeyId: permKey.AuthKeyId(), Updates: updates})
	if err != nil {
		t.Fatalf("GatewayPushUpdatesData() error = %v", err)
	}
	if !tg.FromBool(ok) {
		t.Fatal("GatewayPushUpdatesData() = false")
	}
	wantBody, err := iface.EncodeObject(updates, 223)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	if writer.msg.AuthKeyId != tempKey.AuthKeyId() || writer.msg.SessionId != 22 || !bytes.Equal(writer.msg.Body, wantBody) {
		t.Fatalf("written msg = %#v", writer.msg)
	}
}

func TestGatewayPushSessionUpdatesDataWritesLocalUpdate(t *testing.T) {
	local := push.NewLocalWriter()
	key := crypto.CreateAuthKey()
	writer := &coreFakeSessionWriter{}
	local.Register(push.LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, Layer: 223, AuthKey: key, Writer: writer})
	core := New(context.Background(), &svc.ServiceContext{Push: local})
	updates := tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})

	ok, err := core.GatewayPushSessionUpdatesData(&gateway.TLGatewayPushSessionUpdatesData{AuthKeyId: key.AuthKeyId(), SessionId: 22, Updates: updates})
	if err != nil {
		t.Fatalf("GatewayPushSessionUpdatesData() error = %v", err)
	}
	if !tg.FromBool(ok) {
		t.Fatal("GatewayPushSessionUpdatesData() = false")
	}
	if writer.msg.AuthKeyId != key.AuthKeyId() || writer.msg.SessionId != 22 {
		t.Fatalf("written msg = %#v", writer.msg)
	}
}
