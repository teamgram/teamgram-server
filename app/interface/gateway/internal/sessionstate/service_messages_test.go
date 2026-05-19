package sessionstate

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestServiceMessagePingReturnsPong(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{}
	processor := NewProcessor(store, dispatch)

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1000, encodeTL(t, &mt.TLPing{PingId: 9}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	if decoded.MsgId <= 1000 || decoded.MsgId&3 != 1 {
		t.Fatalf("response msg_id = %d, want valid server msg_id > request", decoded.MsgId)
	}
	pong := decodeBodyAs[*mt.TLPong](t, decoded.Body)
	if pong.MsgId != 1000 || pong.PingId != 9 {
		t.Fatalf("pong = %#v", pong)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch calls = %d, want 0", len(dispatch.payloads))
	}
}

func TestServiceMessagePingWithObserverDoesNotRequireUserMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{
		key:     tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
		userErr: errors.New("metadata unavailable"),
	}
	processor := NewProcessor(store, &fakeDispatcher{})
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     1000,
		SeqNo:     1,
		Body:      encodeTL(t, &mt.TLPing{PingId: 9}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	observed := false

	resp, err := processor.HandleEncryptedWithSession(context.Background(), ConnInfo{}, payload, func(session ActiveSession) SeqNoAllocator {
		observed = true
		return nil
	})
	if err != nil {
		t.Fatalf("HandleEncryptedWithSession() error = %v", err)
	}
	if observed {
		t.Fatal("observer called for service message")
	}
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	pong := decodeBodyAs[*mt.TLPong](t, decoded.Body)
	if pong.MsgId != 1000 || pong.PingId != 9 {
		t.Fatalf("pong = %#v", pong)
	}
}

func TestServiceMessagePingDelayDisconnect(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1001, encodeTL(t, &mt.TLPingDelayDisconnect{PingId: 10, DisconnectDelay: 30}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	pong := decodeBodyAs[*mt.TLPong](t, decoded.Body)
	if pong.PingId != 10 {
		t.Fatalf("pong = %#v", pong)
	}
}

func TestServiceMessagePingDelayDisconnectConcurrent(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = handleEncryptedForTest(t, processor, clientKey, serverKey, int64(2000+i), encodeTL(t, &mt.TLPingDelayDisconnect{
				PingId:          int64(i),
				DisconnectDelay: 30,
			}))
		}(i)
	}
	wg.Wait()
}

func TestServiceMessageMsgsAckNoDispatch(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{}
	processor := NewProcessor(store, dispatch)

	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     1002,
		SeqNo:     1,
		Body:      encodeTL(t, &mt.TLMsgsAck{MsgIds: []int64{1, 2}}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	resp, err := processor.HandleEncrypted(context.Background(), ConnInfo{}, payload)
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	if resp != nil {
		t.Fatalf("HandleEncrypted() response = %x, want nil", resp)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch calls = %d, want 0", len(dispatch.payloads))
	}
}

func TestServiceMessageMsgsAckRecordsOutboundAckState(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})
	key := runtimeSessionKey{authKeyId: serverKey.AuthKeyId(), authKeyType: tg.AuthKeyTypePerm, sessionId: 77}
	processor.runtime.recordOutbound(key, 9001)

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1004, encodeTL(t, &mt.TLMsgsAck{MsgIds: []int64{9001}}))
	if resp != nil {
		t.Fatalf("HandleEncrypted() response = %x, want nil", resp)
	}
	if processor.runtime.hasOutboundUnacked(key, 9001) {
		t.Fatal("outbound msg_id 9001 is still unacked after msgs_ack")
	}
}

func TestServiceMessageGetFutureSaltsNoDispatch(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &recordingFutureSaltsStore{
		fakeAuthKeyStore: fakeAuthKeyStore{
			key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
			futureSalts: tg.MakeTLFutureSalts(&tg.TLFutureSalts{
				Now: 123,
				Salts: []*tg.TLFutureSalt{
					tg.MakeTLFutureSalt(&tg.TLFutureSalt{ValidSince: 1, ValidUntil: 200, Salt: 555}),
				},
			}),
		},
	}
	dispatch := &fakeDispatcher{}
	processor := NewProcessor(store, dispatch)

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1005, encodeTL(t, &mt.TLGetFutureSalts{Num: 3}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	salts := decodeBodyAs[*mt.TLFutureSalts](t, decoded.Body)
	if store.futureSaltsAuthKeyId != serverKey.AuthKeyId() || store.futureSaltsNum != 3 {
		t.Fatalf("GetFutureSalts auth_key_id=%d num=%d, want %d/3", store.futureSaltsAuthKeyId, store.futureSaltsNum, serverKey.AuthKeyId())
	}
	if salts.ReqMsgId != 1005 || salts.Now != 123 || len(salts.Salts) != 1 || salts.Salts[0].Salt != 555 {
		t.Fatalf("future_salts = %#v", salts)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch calls = %d, want 0", len(dispatch.payloads))
	}
}

func TestServiceMessageGetFutureSaltsNormalizesCount(t *testing.T) {
	tests := []struct {
		name string
		num  int32
		want int32
	}{
		{name: "zero defaults", num: 0, want: defaultFutureSaltsCount},
		{name: "negative defaults", num: -1, want: defaultFutureSaltsCount},
		{name: "too large clamps", num: maxFutureSaltsCount + 1, want: maxFutureSaltsCount},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverKey, clientKey := sessionTestKeys()
			store := &recordingFutureSaltsStore{
				fakeAuthKeyStore: fakeAuthKeyStore{
					key:         tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
					futureSalts: tg.MakeTLFutureSalts(&tg.TLFutureSalts{Now: 123}),
				},
			}
			processor := NewProcessor(store, &fakeDispatcher{})

			_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 1005, encodeTL(t, &mt.TLGetFutureSalts{Num: tt.num}))
			if store.futureSaltsNum != tt.want {
				t.Fatalf("GetFutureSalts num = %d, want %d", store.futureSaltsNum, tt.want)
			}
		})
	}
}

func TestServiceMessageDestroySessionMarksRuntimeAndReturnsResult(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})
	target := runtimeSessionKey{authKeyId: serverKey.AuthKeyId(), authKeyType: tg.AuthKeyTypePerm, sessionId: 88}

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1006, encodeTL(t, &mt.TLDestroySession{SessionId: 88}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	ok := decodeBodyAs[*mt.TLDestroySessionOk](t, decoded.Body)
	if ok.SessionId != 88 {
		t.Fatalf("destroy_session_ok = %#v", ok)
	}
	if !processor.runtime.isDestroyed(target) {
		t.Fatal("destroy_session did not mark target session destroyed")
	}

	resp = handleEncryptedForTest(t, processor, clientKey, serverKey, 1007, encodeTL(t, &mt.TLDestroySession{SessionId: 88}))
	decoded = decodeEncryptedForTest(t, clientKey, resp)
	none := decodeBodyAs[*mt.TLDestroySessionNone](t, decoded.Body)
	if none.SessionId != 88 {
		t.Fatalf("destroy_session_none = %#v", none)
	}
}

func TestServiceMessageDestroyedSessionDoesNotDispatchLaterMessages(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 1006, encodeTL(t, &mt.TLDestroySession{SessionId: 88}))
	resp, err := handleEncryptedErrorForTest(t, processor, clientKey, serverKey, 1007, 88, encodeTL(t, &tg.TLHelpGetConfig{}))
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	if resp != nil {
		t.Fatalf("destroyed session response = %x, want nil", resp)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch calls = %d, want 0", len(dispatch.payloads))
	}
}

func TestServiceMessageContainerStaysInSessionstate(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{}
	processor := NewProcessor(store, dispatch)
	container := &mt.TLMsgContainer{Messages: []*mt.TLMessage2{
		{MsgId: 301, Seqno: 1, Object: &mt.TLPing{PingId: 12}},
		{MsgId: 302, Seqno: 1, Object: &mt.TLMsgsAck{MsgIds: []int64{9002}}},
	}}

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 300, encodeTL(t, container))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	pong := decodeBodyAs[*mt.TLPong](t, decoded.Body)
	if pong.MsgId != 301 || pong.PingId != 12 {
		t.Fatalf("response = %#v, want pong for 301/12", pong)
	}
	if len(dispatch.payloads) != 0 {
		t.Fatalf("dispatch calls = %d, want 0", len(dispatch.payloads))
	}
}

func TestServiceMessageGzipPackedReentersNormalHandling(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})
	unpackedBody := encodeTL(t, &mt.TLPing{PingId: 15})
	packedBody := encodeTL(t, &mt.TLGzipPacked{PackedData: unpackedBody})

	unpackedResp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1008, unpackedBody)
	packedResp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1008, packedBody)
	unpacked := decodeEncryptedForTest(t, clientKey, unpackedResp)
	packed := decodeEncryptedForTest(t, clientKey, packedResp)
	if string(unpacked.Body) != string(packed.Body) {
		t.Fatalf("gzip response body = %x, want %x", packed.Body, unpacked.Body)
	}
}

func TestServiceMessageGzipPackedDispatchesNestedRPCPayload(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	rpcBody := encodeTL(t, &tg.TLHelpGetConfig{})
	packedBody := encodeTL(t, &mt.TLGzipPacked{PackedData: rpcBody})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1009, packedBody)
	if resp == nil {
		t.Fatal("gzip nested RPC response is nil")
	}
	if len(dispatch.payloads) != 1 || !bytes.Equal(dispatch.payloads[0], rpcBody) {
		t.Fatalf("dispatch payloads = %x, want %x", dispatch.payloads, rpcBody)
	}
}

func TestServiceMessageGzipPackedContainerPreservesRawMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	container := &mt.TLMsgContainer{Messages: []*mt.TLMessage2{
		{MsgId: 401, Seqno: 1, Object: &tg.TLHelpGetConfig{}},
		{MsgId: 402, Seqno: 1, Object: &mt.TLPing{PingId: 40}},
	}}
	containerBody := encodeTL(t, container)
	packedBody := encodeTL(t, &mt.TLGzipPacked{PackedData: containerBody})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 400, packedBody)
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	respContainer := decodeBodyAs[*mt.TLMsgContainer](t, decoded.Body)
	if len(respContainer.Messages) != 2 {
		t.Fatalf("response count = %d, want 2", len(respContainer.Messages))
	}
	rpcResult, ok := respContainer.Messages[0].Object.(*mt.TLRpcResult)
	if !ok || rpcResult.ReqMsgId != 401 {
		t.Fatalf("first response = %#v, want rpc_result for 401", respContainer.Messages[0].Object)
	}
	pong, ok := respContainer.Messages[1].Object.(*mt.TLPong)
	if !ok || pong.MsgId != 402 || pong.PingId != 40 {
		t.Fatalf("second response = %#v, want pong for 402/40", respContainer.Messages[1].Object)
	}
}

func TestServiceMessageMsgsStateReqReturnsMsgsStateInfo(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	processor := NewProcessor(store, &fakeDispatcher{})

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 1003, encodeTL(t, &mt.TLMsgsStateReq{MsgIds: []int64{11, 12, 13}}))
	decoded := decodeEncryptedForTest(t, clientKey, resp)
	info := decodeBodyAs[*mt.TLMsgsStateInfo](t, decoded.Body)
	if info.ReqMsgId != 1003 {
		t.Fatalf("ReqMsgId = %d, want 1003", info.ReqMsgId)
	}
	if len(info.Info) != 3 {
		t.Fatalf("Info length = %d, want 3", len(info.Info))
	}
}

type recordingFutureSaltsStore struct {
	fakeAuthKeyStore
	futureSaltsAuthKeyId int64
	futureSaltsNum       int32
}

func (s *recordingFutureSaltsStore) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.FutureSalts, error) {
	s.futureSaltsAuthKeyId = authKeyId
	s.futureSaltsNum = num
	return s.fakeAuthKeyStore.GetFutureSalts(ctx, authKeyId, num)
}
