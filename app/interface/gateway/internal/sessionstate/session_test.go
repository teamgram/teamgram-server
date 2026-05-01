package sessionstate

import (
	"bytes"
	"context"
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
}

func (f *fakeDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	f.md = append(f.md, md)
	f.payloads = append(f.payloads, append([]byte(nil), payload...))
	return f.result, nil
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
