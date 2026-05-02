package sessionstate

import (
	"context"
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
