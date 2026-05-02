package sessionstate

import (
	"bytes"
	"context"
	"sync"
	"sync/atomic"
	"testing"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
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
}

func (f *fakeDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	f.md = append(f.md, md)
	f.payloads = append(f.payloads, append([]byte(nil), payload...))
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
	requestBody := encodeTL(t, &mt.TLGetFutureSalts{Num: 1})

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

func TestSessionUnwrapsInitConnectionMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	inner := encodeTL(t, &mt.TLGetFutureSalts{Num: 1})
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

func TestSessionKeepsCachedAuthKeyInfoMetadata(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	keyInfo := tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)
	keyInfo.PermAuthKeyId = 4242
	store := &fakeAuthKeyStore{key: keyInfo}
	dispatch := &fakeDispatcher{result: encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})}
	processor := NewProcessor(store, dispatch)
	requestBody := encodeTL(t, &mt.TLGetFutureSalts{Num: 1})

	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 103, requestBody)
	_ = handleEncryptedForTest(t, processor, clientKey, serverKey, 104, requestBody)
	if len(dispatch.md) != 2 {
		t.Fatalf("dispatch count = %d, want 2", len(dispatch.md))
	}
	if got := dispatch.md[1].PermAuthKeyId; got != 4242 {
		t.Fatalf("cached PermAuthKeyId = %d, want 4242", got)
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

func TestSessionWrapsDispatchRPCError(t *testing.T) {
	serverKey, clientKey := sessionTestKeys()
	store := &fakeAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}
	dispatch := &fakeDispatcher{err: fakeRPCError{err: mt.MakeTLRpcError(&mt.TLRpcError{
		ErrorCode:    400,
		ErrorMessage: "PHONE_NUMBER_UNOCCUPIED",
	})}}
	processor := NewProcessor(store, dispatch)

	resp := handleEncryptedForTest(t, processor, clientKey, serverKey, 102, encodeTL(t, &mt.TLGetFutureSalts{Num: 1}))
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
		{MsgId: 201, Seqno: 1, Object: &mt.TLGetFutureSalts{Num: 1}},
		{MsgId: 202, Seqno: 3, Object: &mt.TLGetFutureSalts{Num: 2}},
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
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     msgID,
		SeqNo:     1,
		Body:      body,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	resp, err := processor.HandleEncrypted(context.Background(), ConnInfo{GatewayId: "gateway-test", ClientAddr: "127.0.0.1:1"}, payload)
	if err != nil {
		t.Fatalf("HandleEncrypted() error = %v", err)
	}
	return resp
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
