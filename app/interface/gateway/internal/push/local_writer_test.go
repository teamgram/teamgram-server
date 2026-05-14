package push

import (
	"bytes"
	"context"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
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
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: tempKey, Layer: 223, Writer: matching, MainUpdates: true})
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

func TestGatewayPushLocalWriterWriteUpdatesUsesMainUpdatesTarget(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	oldTempKey := crypto.CreateAuthKey()
	mainTempKey := crypto.CreateAuthKey()
	oldSession := &fakeSessionWriter{}
	mainSession := &fakeSessionWriter{}

	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: oldTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 10, AuthKey: oldTempKey, Layer: 223, Writer: oldSession})
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: mainTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 20, AuthKey: mainTempKey, Layer: 223, Writer: mainSession, MainUpdates: true})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("WriteUpdates() count = %d, want 1", count)
	}
	if oldSession.seq != 0 || mainSession.seq != 1 {
		t.Fatalf("old seq=%d main seq=%d, want only main target", oldSession.seq, mainSession.seq)
	}
}

func TestGatewayPushLocalWriterRegisterMainUpdatesDemotesPreviousMain(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	firstTempKey := crypto.CreateAuthKey()
	secondTempKey := crypto.CreateAuthKey()
	firstSession := &fakeSessionWriter{}
	secondSession := &fakeSessionWriter{}

	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: firstTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 10, AuthKey: firstTempKey, Layer: 223, Writer: firstSession, MainUpdates: true})
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: secondTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 20, AuthKey: secondTempKey, Layer: 223, Writer: secondSession, MainUpdates: true})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("WriteUpdates() count = %d, want 1", count)
	}
	if firstSession.seq != 0 || secondSession.seq != 1 {
		t.Fatalf("first seq=%d second seq=%d, want only latest main target", firstSession.seq, secondSession.seq)
	}
}

func TestGatewayPushLocalWriterWriteUpdatesSkipsMediaEvenIfMarkedMain(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	mediaTempKey := crypto.CreateAuthKey()
	normalTempKey := crypto.CreateAuthKey()
	mediaSession := &fakeSessionWriter{}
	normalSession := &fakeSessionWriter{}

	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: mediaTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeMediaTemp, SessionId: 10, AuthKey: mediaTempKey, Layer: 223, Writer: mediaSession, MainUpdates: true})
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: normalTempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 20, AuthKey: normalTempKey, Layer: 223, Writer: normalSession, MainUpdates: true})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("WriteUpdates() count = %d, want 1", count)
	}
	if mediaSession.seq != 0 || normalSession.seq != 1 {
		t.Fatalf("media seq=%d normal seq=%d, want only normal target", mediaSession.seq, normalSession.seq)
	}
}

func TestGatewayPushLocalWriterWriteUpdatesDoesNotFallbackToNewestTarget(t *testing.T) {
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
	if count != 0 {
		t.Fatalf("WriteUpdates() count = %d, want 0", count)
	}
	if first.seq != 0 || second.seq != 0 {
		t.Fatalf("first seq=%d second seq=%d, want no writes", first.seq, second.seq)
	}
}

type fakeGenericUpdatesPolicy struct {
	called bool
}

func (p *fakeGenericUpdatesPolicy) HandleGenericUpdatesWithWriter(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz, write GenericUpdatesWriter) (int, error) {
	p.called = true
	return write(ctx, permAuthKeyId, tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
}

func TestGatewayPushLocalWriterWriteUpdatesDelegatesQueuePolicy(t *testing.T) {
	writer := NewLocalWriter()
	permKey := crypto.CreateAuthKey()
	tempKey := crypto.CreateAuthKey()
	session := &fakeSessionWriter{}
	policy := &fakeGenericUpdatesPolicy{}
	writer.SetGenericUpdatesPolicy(policy)
	writer.Register(LocalTarget{PermAuthKeyId: permKey.AuthKeyId(), AuthKeyId: tempKey.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: tempKey, Layer: 223, Writer: session, MainUpdates: true})

	count, err := writer.WriteUpdates(context.Background(), permKey.AuthKeyId(), tg.MakeTLUpdates(&tg.TLUpdates{Seq: 1}))
	if err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	if !policy.called {
		t.Fatal("generic updates policy was not called")
	}
	if count != 1 || session.seq != 1 {
		t.Fatalf("count=%d seq=%d, want one policy-directed write", count, session.seq)
	}
	decoded, err := iface.DecodeObject(bin.NewDecoder(session.msg.Body))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	if updates, ok := decoded.(tg.UpdatesClazz); !ok || updates.UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("decoded updates = %#v, want updatesTooLong", decoded)
	}
}

func TestGatewayPushLocalWriterWriteSessionUpdates(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	sw := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: key, Layer: 223, Writer: sw})
	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
		Date:   1,
	})

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

func TestGatewayPushLocalWriterExactSessionAllowsWhitelistedUpdatesBeforeMainUpdates(t *testing.T) {
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	sw := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: tg.AuthKeyTypeTemp, SessionId: 22, AuthKey: key, Layer: 223, Writer: sw})
	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
		Date:   1,
	})

	ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, updates)
	if err != nil {
		t.Fatalf("WriteSessionUpdates() error = %v", err)
	}
	if !ok {
		t.Fatal("WriteSessionUpdates() ok = false")
	}
	if sw.seq != 1 {
		t.Fatalf("seq = %d, want one exact-session write", sw.seq)
	}
}

func TestGatewayPushLocalWriterExactSessionUpdateShortWhitelist(t *testing.T) {
	tests := []struct {
		name string
		nest tg.UpdateClazz
		want bool
	}{
		{
			name: "allowed nested update",
			nest: tg.MakeTLUpdateConfig(&tg.TLUpdateConfig{}),
			want: true,
		},
		{
			name: "rejected nested update",
			nest: tg.MakeTLUpdateContactsReset(&tg.TLUpdateContactsReset{}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer, key, sw := newExactSessionLocalWriter(t, tg.AuthKeyTypeTemp)
			updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{Update: tt.nest, Date: 1})

			ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, updates)
			if err != nil {
				t.Fatalf("WriteSessionUpdates() error = %v", err)
			}
			if ok != tt.want {
				t.Fatalf("WriteSessionUpdates() ok = %t, want %t", ok, tt.want)
			}
			if wrote := sw.seq != 0; wrote != tt.want {
				t.Fatalf("wrote = %t, want %t", wrote, tt.want)
			}
		})
	}
}

func TestGatewayPushLocalWriterExactSessionUpdatesWhitelistRequiresEveryNestedUpdate(t *testing.T) {
	tests := []struct {
		name    string
		updates tg.UpdatesClazz
		want    bool
	}{
		{
			name: "updates all allowed",
			updates: tg.MakeTLUpdates(&tg.TLUpdates{
				Updates: []tg.UpdateClazz{
					tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
					tg.MakeTLUpdateConfig(&tg.TLUpdateConfig{}),
				},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
				Date:  1,
			}),
			want: true,
		},
		{
			name: "updates has rejected nested update",
			updates: tg.MakeTLUpdates(&tg.TLUpdates{
				Updates: []tg.UpdateClazz{
					tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
					tg.MakeTLUpdateContactsReset(&tg.TLUpdateContactsReset{}),
				},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
				Date:  1,
			}),
			want: false,
		},
		{
			name: "updatesCombined all allowed",
			updates: tg.MakeTLUpdatesCombined(&tg.TLUpdatesCombined{
				Updates: []tg.UpdateClazz{
					tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
					tg.MakeTLUpdateLangPackTooLong(&tg.TLUpdateLangPackTooLong{LangCode: "en"}),
				},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
				Date:  1,
			}),
			want: true,
		},
		{
			name: "updatesCombined has rejected nested update",
			updates: tg.MakeTLUpdatesCombined(&tg.TLUpdatesCombined{
				Updates: []tg.UpdateClazz{
					tg.MakeTLUpdateLangPackTooLong(&tg.TLUpdateLangPackTooLong{LangCode: "en"}),
					tg.MakeTLUpdateContactsReset(&tg.TLUpdateContactsReset{}),
				},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
				Date:  1,
			}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer, key, sw := newExactSessionLocalWriter(t, tg.AuthKeyTypeTemp)

			ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, tt.updates)
			if err != nil {
				t.Fatalf("WriteSessionUpdates() error = %v", err)
			}
			if ok != tt.want {
				t.Fatalf("WriteSessionUpdates() ok = %t, want %t", ok, tt.want)
			}
			if wrote := sw.seq != 0; wrote != tt.want {
				t.Fatalf("wrote = %t, want %t", wrote, tt.want)
			}
		})
	}
}

func TestGatewayPushLocalWriterExactSessionRejectsNonWhitelistedUpdates(t *testing.T) {
	tests := []struct {
		name    string
		updates tg.UpdatesClazz
	}{
		{name: "updatesTooLong", updates: tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})},
		{name: "updateShortMessage", updates: tg.MakeTLUpdateShortMessage(&tg.TLUpdateShortMessage{})},
		{name: "updateShortChatMessage", updates: tg.MakeTLUpdateShortChatMessage(&tg.TLUpdateShortChatMessage{})},
		{name: "updateShortSentMessage", updates: tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{})},
		{name: "updateAccountResetAuthorization", updates: tg.MakeTLUpdateAccountResetAuthorization(&tg.TLUpdateAccountResetAuthorization{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer, key, sw := newExactSessionLocalWriter(t, tg.AuthKeyTypeTemp)

			ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, tt.updates)
			if err != nil {
				t.Fatalf("WriteSessionUpdates() error = %v", err)
			}
			if ok {
				t.Fatal("WriteSessionUpdates() ok = true, want false")
			}
			if sw.seq != 0 {
				t.Fatalf("seq = %d, want no write", sw.seq)
			}
		})
	}
}

func TestGatewayPushLocalWriterExactSessionRejectReportsPolicyReason(t *testing.T) {
	writer, key, sw := newExactSessionLocalWriter(t, tg.AuthKeyTypeTemp)
	updates := tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{})

	result, err := writer.WriteSessionUpdatesDetailed(context.Background(), key.AuthKeyId(), 22, updates)
	if err != nil {
		t.Fatalf("WriteSessionUpdatesDetailed() error = %v", err)
	}
	if result.OK {
		t.Fatal("WriteSessionUpdatesDetailed() OK = true, want false")
	}
	if result.Reason != "exact_session_update_not_allowed" {
		t.Fatalf("Reason = %q, want exact_session_update_not_allowed", result.Reason)
	}
	if result.PermAuthKeyId != key.AuthKeyId() || result.AuthKeyId != key.AuthKeyId() || result.AuthKeyType != tg.AuthKeyTypeTemp || result.SessionId != 22 {
		t.Fatalf("result target fields = %+v, want registered exact session target", result)
	}
	if result.UpdatesClass != tg.ClazzName_updatesTooLong {
		t.Fatalf("UpdatesClass = %q, want %q", result.UpdatesClass, tg.ClazzName_updatesTooLong)
	}
	if sw.seq != 0 {
		t.Fatalf("seq = %d, want no write", sw.seq)
	}
}

func TestGatewayPushLocalWriterMediaExactSessionRejectsNonRPCUpdates(t *testing.T) {
	writer, key, sw := newExactSessionLocalWriter(t, tg.AuthKeyTypeMediaTemp)
	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateLoginToken(&tg.TLUpdateLoginToken{}),
		Date:   1,
	})

	ok, err := writer.WriteSessionUpdates(context.Background(), key.AuthKeyId(), 22, updates)
	if err != nil {
		t.Fatalf("WriteSessionUpdates() error = %v", err)
	}
	if ok {
		t.Fatal("WriteSessionUpdates() ok = true, want false")
	}
	if sw.seq != 0 {
		t.Fatalf("seq = %d, want no write", sw.seq)
	}
}

func newExactSessionLocalWriter(t *testing.T, authKeyType int32) (*LocalWriter, *crypto.AuthKey, *fakeSessionWriter) {
	t.Helper()
	writer := NewLocalWriter()
	key := crypto.CreateAuthKey()
	sw := &fakeSessionWriter{}
	writer.Register(LocalTarget{AuthKeyId: key.AuthKeyId(), AuthKeyType: authKeyType, SessionId: 22, AuthKey: key, Layer: 223, Writer: sw})
	return writer, key, sw
}
