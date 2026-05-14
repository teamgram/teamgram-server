package sessionstate

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDispatcher struct {
	payloads [][]byte
	md       []*metadata.RpcMetadata
	result   []byte
	err      error
	onInvoke func(md *metadata.RpcMetadata, payload []byte)
}

func (f *fakeDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	f.md = append(f.md, md)
	f.payloads = append(f.payloads, append([]byte(nil), payload...))
	if f.onInvoke != nil {
		f.onInvoke(md, payload)
	}
	if f.err != nil {
		return nil, f.err
	}
	return f.result, nil
}

type fakeRPCError struct {
	err *tg.TLRpcError
}

func (e fakeRPCError) Error() string {
	return e.err.Error()
}

func (e fakeRPCError) RPCError() *tg.TLRpcError {
	return e.err
}

func TestSessionDispatchesRawRPCWithMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	requestBody := encodeTL(t, &tg.TLHelpGetConfig{})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 100, requestBody)
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	rpcResult := decodeBodyAs[*mt.TLRpcResult](t, decoded.Body)
	if rpcResult.ReqMsgId != 100 {
		t.Fatalf("rpc_result req_msg_id = %d, want 100", rpcResult.ReqMsgId)
	}
	if len(dispatch.payloads) != 1 || !bytes.Equal(dispatch.payloads[0], requestBody) {
		t.Fatalf("dispatch payloads = %x", dispatch.payloads)
	}
	if got := dispatch.md[0]; got.AuthId != serverKey.AuthKeyId() || got.PermAuthKeyId != serverKey.AuthKeyId() || got.SessionId != 77 || got.ClientMsgId != 100 {
		t.Fatalf("metadata = %#v", got)
	}
}

func TestSessionReturnsBadServerSaltForZeroSaltRequest(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{
		key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
		futureSalts: tg.MakeTLFutureSalts(&tg.TLFutureSalts{
			Salts: []*tg.TLFutureSalt{
				tg.MakeTLFutureSalt(&tg.TLFutureSalt{ValidSince: 1, ValidUntil: 2000000000, Salt: 777}),
			},
		}),
	}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      0,
		SessionId: 77,
		MsgId:     100,
		SeqNo:     1,
		Body:      encodeTL(t, &mt.TLPing{PingId: 1}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}

	resp, err := processor.HandleEncrypted(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "127.0.0.1:1"}, payload)
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	if decoded.Salt != 777 {
		t.Fatalf("response salt = %d, want current future salt 777", decoded.Salt)
	}
	badSalt := decodeBodyAs[*mt.TLBadServerSalt](t, decoded.Body)
	if badSalt.BadMsgId != 100 || badSalt.BadMsgSeqno != 1 || badSalt.ErrorCode != 48 || badSalt.NewServerSalt != 777 {
		t.Fatalf("bad_server_salt = %#v", badSalt)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch count = %d, want 0", len(dispatch.payloads))
	}
}

func TestSessionUnwrapsInitConnectionMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &tg.TLHelpGetConfig{})
	initConn := encodeTL(t, &tg.TLInitConnection{
		ApiId:          1,
		DeviceModel:    "tdesktop",
		SystemVersion:  "macOS",
		AppVersion:     "5.0",
		SystemLangCode: "en",
		LangPack:       "tdesktop",
		LangCode:       "en",
		Query:          inner,
	})
	wrapped := encodeTL(t, &tg.TLInvokeWithLayer{Layer: 224, Query: initConn})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 101, wrapped)
	if len(dispatch.payloads) != 1 || !bytes.Equal(dispatch.payloads[0], inner) {
		t.Fatalf("dispatch payloads = %x", dispatch.payloads)
	}
	if got := dispatch.md[0]; got.Layer != 224 || got.Client != "tdesktop macOS 5.0" || got.Langpack != "tdesktop" || got.LangCode != "en" {
		t.Fatalf("metadata = %#v", got)
	}
}

func TestSessionReusesClientLayerAfterInitConnection(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &tg.TLHelpGetConfig{})
	initConn := encodeTL(t, &tg.TLInitConnection{
		ApiId:          1,
		DeviceModel:    "tdesktop",
		SystemVersion:  "macOS",
		AppVersion:     "5.0",
		SystemLangCode: "en",
		LangPack:       "tdesktop",
		LangCode:       "en",
		Query:          inner,
	})
	wrapped := encodeTL(t, &tg.TLInvokeWithLayer{Layer: 223, Query: initConn})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 106, wrapped)
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 107, encodeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}))

	if len(dispatch.md) != 2 {
		t.Fatalf("dispatch count = %d, want 2", len(dispatch.md))
	}
	if got := dispatch.md[1]; got.Layer != 223 || got.Client != "tdesktop macOS 5.0" || got.Langpack != "tdesktop" || got.LangCode != "en" {
		t.Fatalf("metadata = %#v", got)
	}
}

func TestSessionObserverReceivesPermAuthKeyAndLayer(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo, userID: 12345}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &tg.TLHelpGetConfig{})
	wrapped := encodeTL(t, &tg.TLInvokeWithLayer{Layer: 223, Query: encodeTL(t, &tg.TLInitConnection{
		ApiId:       2040,
		DeviceModel: "tdesktop",
		Query:       inner,
	})})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 108, wrapped)
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     109,
		SeqNo:     1,
		Body:      encodeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	var active ActiveSession
	_, err = processor.HandleEncryptedWithSession(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "127.0.0.1:1"}, payload, func(session ActiveSession) SeqNoAllocator {
		active = session
		return nil
	})
	if err != nil {
		t.Fatalf("HandleEncryptedWithSession() error = %v", err)
	}
	if active.UserId != 12345 || active.PermAuthKeyId != 4242 || active.Layer != 223 || active.Client != "tdesktop" {
		t.Fatalf("active session = %#v", active)
	}
}

func TestSessionSelectorPromotesMainUpdatesOnlyAfterDispatchSuccess(t *testing.T) {
	tests := []struct {
		name string
		body []byte
	}{
		{name: "getState", body: encodeTL(t, &tg.TLUpdatesGetState{})},
		{name: "getDifference", body: encodeTL(t, &tg.TLUpdatesGetDifference{Pts: 1, Date: 2, Qts: 3})},
		{name: "getChannelDifference", body: encodeTL(t, &tg.TLUpdatesGetChannelDifference{
			Channel: tg.MakeTLInputChannelEmpty(&tg.TLInputChannelEmpty{}),
			Filter:  tg.MakeTLChannelMessagesFilterEmpty(&tg.TLChannelMessagesFilterEmpty{}),
			Pts:     1,
			Limit:   10,
		})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverKey, clientKey := sessionTestKeys()
			keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
			keyInfo.PermAuthKeyId = 4242
			store := &fakeAuthKeyStore{key: keyInfo}
			dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
			processor := NewProcessor(store, dispatch)

			var duringSelector ActiveSession
			handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 120, 77, tt.body, func(session ActiveSession) SeqNoAllocator {
				duringSelector = session
				return nil
			})
			if duringSelector.MainUpdates {
				t.Fatalf("selector observer MainUpdates = true, want false until dispatch succeeds")
			}

			var afterSelector ActiveSession
			handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 121, 77, encodeTL(t, &tg.TLHelpGetConfig{}), func(session ActiveSession) SeqNoAllocator {
				afterSelector = session
				return nil
			})
			if !afterSelector.MainUpdates {
				t.Fatalf("post-selector observer MainUpdates = false, want true after dispatch succeeds")
			}
		})
	}
}

func TestSessionUpdateStatusMarksMainUpdatesImmediately(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, tg.BoolTrueClazz)}
	processor := NewProcessor(store, dispatch)

	var active ActiveSession
	handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 130, 77, encodeTL(t, &tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}), func(session ActiveSession) SeqNoAllocator {
		active = session
		return nil
	})
	if !active.MainUpdates {
		t.Fatalf("account.updateStatus observer MainUpdates = false, want immediate promotion")
	}
}

func TestSessionMediaTempCannotBePromotedToMainUpdates(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeMediaTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, tg.BoolTrueClazz)}
	processor := NewProcessor(store, dispatch)

	var active ActiveSession
	handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 140, 77, encodeTL(t, &tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}), func(session ActiveSession) SeqNoAllocator {
		active = session
		return nil
	})
	if active.MainUpdates {
		t.Fatalf("media temp account.updateStatus MainUpdates = true, want false")
	}

	handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 141, 77, encodeTL(t, &tg.TLUpdatesGetState{}), func(session ActiveSession) SeqNoAllocator {
		active = session
		return nil
	})
	if active.MainUpdates {
		t.Fatalf("media temp updates.getState MainUpdates = true, want false")
	}
}

func TestSessionSelectorStaleCompletionDoesNotMutateRegistryQueueOrPendingTooLong(t *testing.T) {
	tests := []struct {
		name        string
		queueWrites int
		wantQueued  int
		wantTooLong bool
		completeOK  bool
	}{
		{name: "queued", queueWrites: 1, wantQueued: 1, completeOK: true},
		{name: "pendingTooLong", queueWrites: 65, wantTooLong: true, completeOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := newRoleRegistry()
			first := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
			second := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1002, authKeyType: tg.AuthKeyTypeTemp, sessionId: 20}

			oldTransition, ok := registry.beginTransition(tg.ClazzName_updates_getState, first)
			if !ok {
				t.Fatal("first selector did not create transition")
			}
			for i := 0; i < tt.queueWrites; i++ {
				registry.enqueueSelectorUpdate(first.permAuthKeyId)
			}
			newTransition, ok := registry.beginTransition(tg.ClazzName_updates_getDifference, second)
			if !ok {
				t.Fatal("second selector did not create transition")
			}

			result := registry.completeSelector(oldTransition.token, tt.completeOK)
			if !result.stale {
				t.Fatalf("old completion stale = false, want true")
			}
			if registry.isMain(first) || registry.isMain(second) {
				t.Fatalf("stale completion promoted main target: first=%t second=%t", registry.isMain(first), registry.isMain(second))
			}
			snapshot := registry.snapshot(first.permAuthKeyId)
			if snapshot.candidateToken != newTransition.token || snapshot.candidateKey != second {
				t.Fatalf("candidate snapshot = %#v, want token %d key %#v", snapshot, newTransition.token, second)
			}
			if snapshot.queued != tt.wantQueued || snapshot.pendingTooLong != tt.wantTooLong {
				t.Fatalf("queue snapshot queued=%d tooLong=%t, want queued=%d tooLong=%t", snapshot.queued, snapshot.pendingTooLong, tt.wantQueued, tt.wantTooLong)
			}
		})
	}
}

func TestSessionSelectorQueueBuffersGenericUpdatesUntilSuccess(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	queuedUpdate := tg.MakeTLUpdates(&tg.TLUpdates{Seq: 10})
	var flushed []tg.UpdatesClazz
	var processor *Processor
	processor = NewProcessor(store, &fakeDispatcher{
		result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2}),
		onInvoke: func(md *metadata.RpcMetadata, payload []byte) {
			count, err := processor.HandleGenericUpdates(context.Background(), 4242, queuedUpdate)
			if err != nil {
				t.Fatalf("HandleGenericUpdates() error = %v", err)
			}
			if count != 0 {
				t.Fatalf("HandleGenericUpdates() count = %d while selector in flight, want 0", count)
			}
		},
	})
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 320, 77, encodeTL(t, &tg.TLUpdatesGetState{}), nil)
	if len(flushed) != 1 || flushed[0] != queuedUpdate {
		t.Fatalf("flushed updates = %#v, want queued update", flushed)
	}
}

func TestSessionSelectorQueueLimitCoalescesToUpdatesTooLong(t *testing.T) {
	processor := NewProcessor(&fakeAuthKeyStore{}, &fakeDispatcher{})
	key := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
	processor.roles.beginSelector(tg.ClazzName_updates_getState, key)

	for i := 0; i < selectorQueueLimit; i++ {
		if count, err := processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: int32(i)})); err != nil || count != 0 {
			t.Fatalf("HandleGenericUpdates(%d) count=%d err=%v, want queued", i, count, err)
		}
	}
	snapshot := processor.roles.snapshot(4242)
	if snapshot.queued != selectorQueueLimit || snapshot.pendingTooLong {
		t.Fatalf("snapshot after limit = %#v, want %d queued and no updatesTooLong", snapshot, selectorQueueLimit)
	}
	if count, err := processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: 65})); err != nil || count != 0 {
		t.Fatalf("overflow HandleGenericUpdates() count=%d err=%v, want coalesced", count, err)
	}
	snapshot = processor.roles.snapshot(4242)
	if snapshot.queued != 0 || !snapshot.pendingTooLong {
		t.Fatalf("snapshot after overflow = %#v, want pending updatesTooLong only", snapshot)
	}
}

func TestSessionSelectorQueueFlushesUpdatesTooLongAfterOverflow(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	var flushed []tg.UpdatesClazz
	var processor *Processor
	processor = NewProcessor(store, &fakeDispatcher{
		result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2}),
		onInvoke: func(md *metadata.RpcMetadata, payload []byte) {
			for i := 0; i <= selectorQueueLimit; i++ {
				if _, err := processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: int32(i)})); err != nil {
					t.Fatalf("HandleGenericUpdates(%d) error = %v", i, err)
				}
			}
		},
	})
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 321, 77, encodeTL(t, &tg.TLUpdatesGetState{}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed updates = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionSelectorFailureAndCloseCreatePendingUpdatesTooLong(t *testing.T) {
	tests := []struct {
		name string
		done func(*Processor, roleTransition, roleSessionKey)
	}{
		{name: "failure", done: func(p *Processor, transition roleTransition, key roleSessionKey) {
			_ = p.roles.completeSelector(transition.token, false)
		}},
		{name: "close", done: func(p *Processor, transition roleTransition, key roleSessionKey) {
			p.UnregisterSession(ActiveSession{PermAuthKeyId: key.permAuthKeyId, AuthKeyId: key.authKeyId, AuthKeyType: key.authKeyType, SessionId: key.sessionId})
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processor := NewProcessor(&fakeAuthKeyStore{}, &fakeDispatcher{})
			key := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
			transition := processor.roles.beginSelector(tg.ClazzName_updates_getState, key)
			_, _ = processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: 1}))

			tt.done(processor, transition, key)
			snapshot := processor.roles.snapshot(4242)
			if snapshot.candidateToken != 0 || snapshot.queued != 0 || !snapshot.pendingTooLong {
				t.Fatalf("snapshot = %#v, want candidate cleared with pending updatesTooLong", snapshot)
			}
		})
	}
}

func TestSessionSelectorTimeoutCreatesPendingUpdatesTooLongWithoutQueuedPayloads(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{err: context.DeadlineExceeded}
	processor := NewProcessor(store, dispatch)
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	_, err := handleEncryptedErrorForTest(t, processor, clientKey, serverKey, 323, 77, encodeTL(t, &tg.TLUpdatesGetState{}))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("selector error = %v, want context deadline exceeded", err)
	}
	snapshot := processor.roles.snapshot(4242)
	if snapshot.candidateToken != 0 || snapshot.queued != 0 || !snapshot.pendingTooLong {
		t.Fatalf("snapshot after timeout = %#v, want candidate cleared with pending updatesTooLong", snapshot)
	}
	if len(flushed) != 0 {
		t.Fatalf("flushed after timeout = %#v, want no immediate flush", flushed)
	}

	dispatch.err = nil
	dispatch.result = encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})
	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 324, 77, encodeTL(t, &tg.TLUpdatesGetState{}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed after next selector = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionSelectorFailureWithoutQueuedPayloadsFlushesUpdatesTooLongOnNextSuccess(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{err: context.Canceled}
	processor := NewProcessor(store, dispatch)
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	_, err := handleEncryptedErrorForTest(t, processor, clientKey, serverKey, 325, 77, encodeTL(t, &tg.TLUpdatesGetState{}))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("selector error = %v, want context canceled", err)
	}
	snapshot := processor.roles.snapshot(4242)
	if snapshot.candidateToken != 0 || snapshot.queued != 0 || !snapshot.pendingTooLong {
		t.Fatalf("snapshot after failure = %#v, want candidate cleared with pending updatesTooLong", snapshot)
	}

	dispatch.err = nil
	dispatch.result = encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})
	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 326, 77, encodeTL(t, &tg.TLUpdatesGetDifference{Pts: 1, Date: 2, Qts: 3}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed after next selector = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionSelectorCloseWithoutQueuedPayloadsFlushesUpdatesTooLongOnNextSuccess(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	processor := NewProcessor(store, &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})})
	first := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
	processor.roles.beginSelector(tg.ClazzName_updates_getState, first)
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	processor.UnregisterSession(ActiveSession{PermAuthKeyId: first.permAuthKeyId, AuthKeyId: first.authKeyId, AuthKeyType: first.authKeyType, SessionId: first.sessionId})
	snapshot := processor.roles.snapshot(4242)
	if snapshot.candidateToken != 0 || snapshot.queued != 0 || !snapshot.pendingTooLong {
		t.Fatalf("snapshot after close = %#v, want candidate cleared with pending updatesTooLong", snapshot)
	}

	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 327, 77, encodeTL(t, &tg.TLUpdatesGetState{}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed after next selector = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionGenericUpdatesNoMainNoSelectorColdStartDoesNotWriteOrQueue(t *testing.T) {
	processor := NewProcessor(&fakeAuthKeyStore{}, &fakeDispatcher{})
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		t.Fatal("generic update writer called without main or selector")
		return 0, nil
	})

	count, err := processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: 1}))
	if err != nil {
		t.Fatalf("HandleGenericUpdates() error = %v", err)
	}
	if count != 0 {
		t.Fatalf("HandleGenericUpdates() count = %d, want 0", count)
	}
	if snapshot := processor.roles.snapshot(4242); snapshot.queued != 0 || snapshot.pendingTooLong || snapshot.candidateToken != 0 {
		t.Fatalf("snapshot = %#v, want no cold-start queue", snapshot)
	}
}

func TestSessionUpdateStatusDrainsQueuedUpdatesImmediately(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	processor := NewProcessor(store, &fakeDispatcher{result: encodeTL(t, tg.BoolTrueClazz)})
	key := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
	processor.roles.beginSelector(tg.ClazzName_updates_getState, key)
	queuedUpdate := tg.MakeTLUpdates(&tg.TLUpdates{Seq: 99})
	_, _ = processor.HandleGenericUpdates(context.Background(), 4242, queuedUpdate)
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		flushed = append(flushed, updates)
		return 1, nil
	})

	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 322, 77, encodeTL(t, &tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}), nil)
	if len(flushed) != 1 || flushed[0] != queuedUpdate {
		t.Fatalf("flushed updates = %#v, want queued update", flushed)
	}
	snapshot := processor.roles.snapshot(4242)
	if snapshot.queued != 0 || snapshot.pendingTooLong || snapshot.candidateToken != 0 {
		t.Fatalf("snapshot = %#v, want drained queue and cleared candidate", snapshot)
	}
}

func TestSessionUpdateStatusFlushErrorDoesNotFailRPCAndLeavesUpdatesTooLong(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, tg.BoolTrueClazz)}
	processor := NewProcessor(store, dispatch)
	key := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
	processor.roles.beginSelector(tg.ClazzName_updates_getState, key)
	_, _ = processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: 99}))
	errFlush := errors.New("flush failed")
	failFlush := true
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		if failFlush {
			failFlush = false
			return 0, errFlush
		}
		flushed = append(flushed, updates)
		return 1, nil
	})

	resp, err := handleEncryptedErrorForTest(t, processor, clientKey, serverKey, 328, 77, encodeTL(t, &tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}))
	if err != nil {
		t.Fatalf("account.updateStatus returned error = %v, want nil", err)
	}
	if resp == nil || len(dispatch.payloads) != 1 {
		t.Fatalf("account.updateStatus resp nil=%t dispatches=%d, want reply and dispatch", resp == nil, len(dispatch.payloads))
	}
	if snapshot := processor.roles.snapshot(4242); !snapshot.pendingTooLong {
		t.Fatalf("snapshot after failed drain = %#v, want pending updatesTooLong", snapshot)
	}

	dispatch.result = encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})
	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 329, 77, encodeTL(t, &tg.TLUpdatesGetState{}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed after recovery selector = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionSelectorSuccessFlushErrorDoesNotFailRPCAndLeavesUpdatesTooLong(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	var processor *Processor
	queuedOnce := false
	dispatch := &fakeDispatcher{
		result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2}),
		onInvoke: func(md *metadata.RpcMetadata, payload []byte) {
			if rawRPCMethodName(payload) == tg.ClazzName_updates_getState && !queuedOnce {
				queuedOnce = true
				_, _ = processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: 100}))
			}
		},
	}
	processor = NewProcessor(store, dispatch)
	errFlush := errors.New("selector flush failed")
	failFlush := true
	var flushed []tg.UpdatesClazz
	processor.SetGenericUpdatesWriter(func(ctx context.Context, permAuthKeyId int64, updates tg.UpdatesClazz) (int, error) {
		if failFlush {
			failFlush = false
			return 0, errFlush
		}
		flushed = append(flushed, updates)
		return 1, nil
	})

	resp, err := handleEncryptedErrorForTest(t, processor, clientKey, serverKey, 330, 77, encodeTL(t, &tg.TLUpdatesGetState{}))
	if err != nil {
		t.Fatalf("selector returned error = %v, want nil", err)
	}
	if resp == nil {
		t.Fatal("selector response is nil, want foreground RPC reply")
	}
	if snapshot := processor.roles.snapshot(4242); !snapshot.pendingTooLong {
		t.Fatalf("snapshot after failed selector flush = %#v, want pending updatesTooLong", snapshot)
	}

	_ = handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, 331, 77, encodeTL(t, &tg.TLUpdatesGetDifference{Pts: 1, Date: 2, Qts: 3}), nil)
	if len(flushed) != 1 || flushed[0].UpdatesClazzName() != tg.ClazzName_updatesTooLong {
		t.Fatalf("flushed after recovery selector = %#v, want one updatesTooLong", flushed)
	}
}

func TestSessionCandidateReplacementPreservesPendingUpdatesTooLong(t *testing.T) {
	processor := NewProcessor(&fakeAuthKeyStore{}, &fakeDispatcher{})
	first := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1001, authKeyType: tg.AuthKeyTypeTemp, sessionId: 10}
	second := roleSessionKey{permAuthKeyId: 4242, authKeyId: 1002, authKeyType: tg.AuthKeyTypeTemp, sessionId: 20}
	processor.roles.beginSelector(tg.ClazzName_updates_getState, first)
	for i := 0; i <= selectorQueueLimit; i++ {
		_, _ = processor.HandleGenericUpdates(context.Background(), 4242, tg.MakeTLUpdates(&tg.TLUpdates{Seq: int32(i)}))
	}

	processor.roles.beginSelector(tg.ClazzName_updates_getDifference, second)
	snapshot := processor.roles.snapshot(4242)
	if snapshot.candidateKey != second || !snapshot.pendingTooLong || snapshot.queued != 0 {
		t.Fatalf("snapshot = %#v, want replacement candidate with pending updatesTooLong preserved", snapshot)
	}
}

func TestSessionDispatchesGatewayRouteMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     109,
		SeqNo:     1,
		Body:      encodeTL(t, &tg.TLHelpGetConfig{}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}

	_, err = processor.HandleEncrypted(context.Background(), ConnInfo{
		GatewayId:         "gateway-test",
		GatewayRpcAddr:    "127.0.0.1:20110",
		GatewayGeneration: "generation-test",
		ClientAddr:        "127.0.0.1:1",
	}, payload)
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	if got := dispatch.md[0]; got.GatewayRpcAddr != "127.0.0.1:20110" || got.GatewayGeneration != "generation-test" {
		t.Fatalf("gateway route metadata = %#v", got)
	}
}

func TestSessionPersistsClientMetadataFromInitConnection(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &tg.TLHelpGetConfig{})
	wrapped := encodeTL(t, &tg.TLInvokeWithLayer{Layer: 223, Query: encodeTL(t, &tg.TLInitConnection{
		ApiId:          2040,
		DeviceModel:    "tdesktop",
		SystemVersion:  "macOS",
		AppVersion:     "5.13",
		SystemLangCode: "en-US",
		LangPack:       "tdesktop",
		LangCode:       "en",
		Query:          inner,
	})})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 301, wrapped)
	if len(store.setClientSessions) != 1 {
		t.Fatalf("SetClientSessionInfo calls = %d, want 1", len(store.setClientSessions))
	}
	got := store.setClientSessions[0]
	if got.AuthKeyId != serverKey.AuthKeyId() || got.Layer != 223 || got.ApiId != 2040 ||
		got.DeviceModel != "tdesktop" || got.SystemVersion != "macOS" ||
		got.AppVersion != "5.13" || got.SystemLangCode != "en-US" ||
		got.LangPack != "tdesktop" || got.LangCode != "en" {
		t.Fatalf("persisted client session = %#v", got)
	}
}

func TestSessionRestoresClientMetadataForBareRPC(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{
		key: keyInfo,
		clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
			AuthKeyId:   serverKey.AuthKeyId(),
			Layer:       223,
			DeviceModel: "tdesktop",
			LangPack:    "tdesktop",
			LangCode:    "en",
		}).ToClientSession(),
	}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 302, encodeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}))
	if got := dispatch.md[0]; got.Layer != 223 || got.Langpack != "tdesktop" || got.LangCode != "en" {
		t.Fatalf("metadata = %#v", got)
	}
	if len(store.setClientSessions) != 0 {
		t.Fatalf("bare RPC persisted client sessions = %d, want 0", len(store.setClientSessions))
	}
}

func TestSessionSkipsDuplicateClientMetadataWrites(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &tg.TLHelpGetConfig{})
	wrapped := encodeTL(t, &tg.TLInvokeWithLayer{Layer: 223, Query: encodeTL(t, &tg.TLInitConnection{
		ApiId:       2040,
		DeviceModel: "tdesktop",
		LangPack:    "tdesktop",
		LangCode:    "en",
		Query:       inner,
	})})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 303, wrapped)
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 304, wrapped)
	if len(store.setClientSessions) != 1 {
		t.Fatalf("SetClientSessionInfo calls = %d, want 1", len(store.setClientSessions))
	}
}

func TestSessionBareRPCDoesNotPersistOnlyForIPChange(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	store := &fakeAuthKeyStore{
		key: keyInfo,
		clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
			AuthKeyId: serverKey.AuthKeyId(),
			Ip:        "127.0.0.1",
			Layer:     223,
			LangPack:  "tdesktop",
			LangCode:  "en",
		}).ToClientSession(),
	}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	body := encodeTL(t, &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz})
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     305,
		SeqNo:     1,
		Body:      body,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	_, err = processor.HandleEncrypted(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "10.0.0.2:44321"}, payload)
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	if len(store.setClientSessions) != 0 {
		t.Fatalf("SetClientSessionInfo calls = %d, want 0", len(store.setClientSessions))
	}
}

func TestSessionKeepsCachedAuthKeyInfoMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	requestBody := encodeTL(t, &tg.TLHelpGetConfig{})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 103, requestBody)
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 104, requestBody)
	if len(dispatch.md) != 2 {
		t.Fatalf("dispatch count = %d, want 2", len(dispatch.md))
	}
	if got := dispatch.md[1].PermAuthKeyId; got != 4242 {
		t.Fatalf("cached PermAuthKeyId = %d, want 4242", got)
	}
}

func TestSessionRefreshesAuthKeyInfoAfterBindTempAuthKey(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	tempAuthKeyID := serverKey.AuthKeyId()
	permAuthKeyID := int64(4242)
	staleKeyInfo := tg.NewAuthKeyInfo(tempAuthKeyID, serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	staleKeyInfo.PermAuthKeyId = tempAuthKeyID
	refreshedKeyInfo := tg.NewAuthKeyInfo(tempAuthKeyID, serverKey.AuthKey(), tg.AuthKeyTypeTemp)
	refreshedKeyInfo.PermAuthKeyId = permAuthKeyID
	store := &fakeAuthKeyStore{key: staleKeyInfo}
	dispatch := &fakeDispatcher{
		result: encodeTL(t, tg.BoolTrueClazz),
		onInvoke: func(md *metadata.RpcMetadata, payload []byte) {
			if rawRPCMethodName(payload) == tg.ClazzName_auth_bindTempAuthKey {
				store.key = refreshedKeyInfo
			}
		},
	}
	processor := NewProcessor(store, dispatch)

	bind := encodeTL(t, &tg.TLAuthBindTempAuthKey{
		PermAuthKeyId:    permAuthKeyID,
		Nonce:            1,
		ExpiresAt:        2,
		EncryptedMessage: []byte("bind"),
	})
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 108, bind)
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 109, encodeTL(t, &tg.TLHelpGetConfig{}))

	if len(dispatch.md) != 2 {
		t.Fatalf("dispatch count = %d, want 2", len(dispatch.md))
	}
	if got := dispatch.md[0].PermAuthKeyId; got != tempAuthKeyID {
		t.Fatalf("bind PermAuthKeyId = %d, want initial temp auth key %d", got, tempAuthKeyID)
	}
	if got := dispatch.md[1].PermAuthKeyId; got != permAuthKeyID {
		t.Fatalf("post-bind PermAuthKeyId = %d, want refreshed perm auth key %d", got, permAuthKeyID)
	}
}

func TestSessionDispatchesBoundUserMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo, userID: 1001}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 105, encodeTL(t, &tg.TLHelpGetConfig{}))
	if len(dispatch.md) != 1 {
		t.Fatalf("dispatch count = %d, want 1", len(dispatch.md))
	}
	if got := dispatch.md[0].UserId; got != 1001 {
		t.Fatalf("metadata UserId = %d, want 1001", got)
	}
	if got := store.userKeyID; got != serverKey.AuthKeyId() {
		t.Fatalf("GetUserId auth key = %d, want current auth key %d", got, serverKey.AuthKeyId())
	}
}

func TestSessionAuthKeyCacheConcurrent(t *testing.T) {
	serverKey, _ := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	keyInfo.PermAuthKeyId = 4242
	store := &countingAuthKeyStore{key: keyInfo}
	processor := NewProcessor(store, &fakeDispatcher{})

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key, info, err := processor.authKey(context.Background(), serverKey.AuthKeyId())
			if err != nil {
				t.Errorf("authKey() error = %v", err)
				return
			}
			if key.AuthKeyId() != serverKey.AuthKeyId() || info == nil || info.PermAuthKeyId != 4242 {
				t.Errorf("authKey() = key %v info %#v", key.AuthKeyId(), info)
			}
		}()
	}
	wg.Wait()
	if got := store.calls.Load(); got != 1 {
		t.Fatalf("QueryAuthKey calls = %d, want 1", got)
	}
}

func TestSessionSeqNoSeparatesAuthKeyType(t *testing.T) {
	processor := NewProcessor(&fakeAuthKeyStore{}, &fakeDispatcher{})
	authKeyID := int64(1001)
	sessionID := int64(77)

	if got := processor.nextSeqNo(authKeyID, tg.AuthKeyTypeTemp, sessionID, true, nil); got != 1 {
		t.Fatalf("temp first seq = %d, want 1", got)
	}
	if got := processor.nextSeqNo(authKeyID, tg.AuthKeyTypeMediaTemp, sessionID, true, nil); got != 1 {
		t.Fatalf("media first seq = %d, want 1", got)
	}
	if got := processor.nextSeqNo(authKeyID, tg.AuthKeyTypeTemp, sessionID, true, nil); got != 3 {
		t.Fatalf("temp second seq = %d, want 3", got)
	}
}

func TestSessionWrapsDispatchRPCError(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{err: fakeRPCError{err: mt.MakeTLRpcError(&mt.TLRpcError{
		ErrorCode:    400,
		ErrorMessage: "PHONE_NUMBER_UNOCCUPIED",
	})}}
	processor := NewProcessor(store, dispatch)

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 102, encodeTL(t, &tg.TLHelpGetConfig{}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	rpcResult := decodeBodyAs[*mt.TLRpcResult](t, decoded.Body)
	errObj, ok := rpcResult.Result.(*mt.TLRpcError)
	if !ok {
		t.Fatalf("rpc_result result = %T, want *mt.TLRpcError", rpcResult.Result)
	}
	if errObj.ErrorCode != 400 || errObj.ErrorMessage != "PHONE_NUMBER_UNOCCUPIED" {
		t.Fatalf("rpc_error = %#v", errObj)
	}
}

func TestSessionContainerReturnsAllRPCResponses(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	container := &mt.TLMsgContainer{Messages: []*mt.TLMessage2{
		{MsgId: 201, Seqno: 1, Object: &tg.TLHelpGetConfig{}},
		{MsgId: 202, Seqno: 3, Object: &tg.TLUsersGetFullUser{Id: tg.InputUserSelfClazz}},
	}}

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 200, encodeTL(t, container))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	respContainer := decodeBodyAs[*mt.TLMsgContainer](t, decoded.Body)
	if len(respContainer.Messages) != 2 {
		t.Fatalf("response count = %d, want 2", len(respContainer.Messages))
	}
	for i, wantReqID := range []int64{201, 202} {
		rpcResult, ok := respContainer.Messages[i].Object.(*mt.TLRpcResult)
		if !ok {
			t.Fatalf("response %d = %T, want *mt.TLRpcResult", i, respContainer.Messages[i].Object)
		}
		if rpcResult.ReqMsgId != wantReqID {
			t.Fatalf("response %d req_msg_id = %d, want %d", i, rpcResult.ReqMsgId, wantReqID)
		}
	}
}

func TestSessionContainerPreservesRawLayeredRPCPayload(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	container := &mt.TLMsgContainer{Messages: []*mt.TLMessage2{
		{MsgId: 203, Seqno: 1, Object: &tg.TLHelpGetConfig{}},
	}}
	wantPayload := encodeTL(t, &tg.TLHelpGetConfig{})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 200, encodeTL(t, container))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	respContainer := decodeBodyAs[*mt.TLMsgContainer](t, decoded.Body)
	if len(respContainer.Messages) != 1 {
		t.Fatalf("response count = %d, want 1", len(respContainer.Messages))
	}
	if len(dispatch.payloads) != 1 || !bytes.Equal(dispatch.payloads[0], wantPayload) {
		t.Fatalf("dispatch payloads = %x, want %x", dispatch.payloads, wantPayload)
	}
	if got := dispatch.md[0].ClientMsgId; got != 203 {
		t.Fatalf("metadata ClientMsgId = %d, want 203", got)
	}
}

func handleEncryptedForTest(t *testing.T, processor *Processor, clientKey, serverKey *crypto.AuthKey, msgID int64, body []byte) []byte {
	t.Helper()
	return handleEncryptedWithObserverForTest(t, processor, clientKey, serverKey, msgID, 77, body, nil)
}

func handleEncryptedWithObserverForTest(t *testing.T, processor *Processor, clientKey, serverKey *crypto.AuthKey, msgID int64, sessionID int64, body []byte, observe SessionObserver) []byte {
	t.Helper()
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: sessionID,
		MsgId:     msgID,
		SeqNo:     1,
		Body:      body,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	resp, err := processor.HandleEncryptedWithSession(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "127.0.0.1:1"}, payload, observe)
	if err != nil {
		t.Fatalf("HandleEncryptedWithSession() error = %v", err)
	}
	return resp
}

func handleEncryptedErrorForTest(t *testing.T, processor *Processor, clientKey, serverKey *crypto.AuthKey, msgID int64, sessionID int64, body []byte) ([]byte, error) {
	t.Helper()
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: sessionID,
		MsgId:     msgID,
		SeqNo:     1,
		Body:      body,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	return processor.HandleEncrypted(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "127.0.0.1:1"}, payload)
}

func decodeEncryptedForTest(t *testing.T, clientKey *crypto.AuthKey, payload []byte) gmtproto.EncryptedMessage {
	t.Helper()
	msg, err := gmtproto.DecodeEncryptedMessage(payload, clientKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage() error = %v", err)
	}
	return msg
}

func sessionTestKeys() (*crypto.AuthKey, *crypto.AuthKey) {
	keyData := make([]byte, 256)
	for i := range keyData {
		keyData[i] = byte(255 - i)
	}
	return crypto.NewAuthKey(0, keyData), crypto.NewClientAuthKey(0, keyData)
}

var _ = bin.WordLen

type countingAuthKeyStore struct {
	key   *tg.AuthKeyInfo
	calls atomic.Int32
}

func (s *countingAuthKeyStore) QueryAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	s.calls.Add(1)
	return s.key, nil
}

func (s *countingAuthKeyStore) SetAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, futureSalt *tg.FutureSalt, expiresIn int32) error {
	return nil
}

func (s *countingAuthKeyStore) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.FutureSalts, error) {
	return nil, nil
}

func (s *countingAuthKeyStore) GetUserId(ctx context.Context, authKeyId int64) (int64, error) {
	return 0, nil
}

func (s *countingAuthKeyStore) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) error {
	return nil
}

func (s *countingAuthKeyStore) SetLayer(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	return nil
}

func (s *countingAuthKeyStore) GetClientSession(ctx context.Context, authKeyId int64) (*authsession.ClientSession, error) {
	return nil, nil
}
