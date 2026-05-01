package mtproto

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestPlainMessageRoundTrip(t *testing.T) {
	body := encodeTL(t, &mt.TLPing{PingId: 7})
	wire, err := EncodePlainMessage(PlainMessage{MsgId: 123456789, Body: body})
	if err != nil {
		t.Fatalf("EncodePlainMessage() error = %v", err)
	}

	if got := binary.LittleEndian.Uint64(wire[:8]); got != 0 {
		t.Fatalf("auth_key_id = %d, want 0", got)
	}
	msg, err := DecodePlainMessage(wire)
	if err != nil {
		t.Fatalf("DecodePlainMessage() error = %v", err)
	}
	if msg.MsgId != 123456789 || !bytes.Equal(msg.Body, body) {
		t.Fatalf("DecodePlainMessage() = %#v, body %x", msg, msg.Body)
	}
}

func TestEncryptedMessageRoundTrip(t *testing.T) {
	serverKey, clientKey := testAuthKeys()
	body := encodeTL(t, &mt.TLPing{PingId: 8})
	wire, err := EncodeEncryptedMessage(EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      11,
		SessionId: 22,
		MsgId:     33,
		SeqNo:     1,
		Body:      body,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}

	msg, err := DecodeEncryptedMessage(wire, serverKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage() error = %v", err)
	}
	if msg.AuthKeyId != serverKey.AuthKeyId() || msg.Salt != 11 || msg.SessionId != 22 || msg.MsgId != 33 || msg.SeqNo != 1 {
		t.Fatalf("DecodeEncryptedMessage() = %#v", msg)
	}
	if !bytes.Equal(msg.Body, body) {
		t.Fatalf("DecodeEncryptedMessage().Body = %x, want %x", msg.Body, body)
	}
}

func TestEncryptedMessageRejectsAuthKeyMismatch(t *testing.T) {
	serverKey, clientKey := testAuthKeys()
	wire, err := EncodeEncryptedMessage(EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      1,
		SessionId: 2,
		MsgId:     3,
		SeqNo:     4,
		Body:      encodeTL(t, &mt.TLPing{PingId: 9}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	binary.LittleEndian.PutUint64(wire[:8], uint64(serverKey.AuthKeyId()+1))

	if _, err := DecodeEncryptedMessage(wire, serverKey); err == nil {
		t.Fatal("DecodeEncryptedMessage() error is nil")
	}
}

func TestNextServerMsgIdMonotonic(t *testing.T) {
	after := int64(1<<62 - 4)
	first := NextServerMsgId(after)
	second := NextServerMsgId(first)
	if first <= after {
		t.Fatalf("first msg_id = %d, want > %d", first, after)
	}
	if second <= first {
		t.Fatalf("second msg_id = %d, want > %d", second, first)
	}
	if first&3 != 1 || second&3 != 1 {
		t.Fatalf("server msg_id low bits = %d, %d, want 1", first&3, second&3)
	}
}

func TestWrapRPCResult(t *testing.T) {
	result := encodeTL(t, &mt.TLPong{MsgId: 1, PingId: 2})
	wrapped, err := WrapRPCResult(99, result)
	if err != nil {
		t.Fatalf("WrapRPCResult() error = %v", err)
	}

	d := bin.NewDecoder(wrapped)
	clazz, err := d.ClazzID()
	if err != nil {
		t.Fatalf("ClazzID() error = %v", err)
	}
	if clazz != mt.ClazzID_rpc_result {
		t.Fatalf("clazz = %#x, want rpc_result", clazz)
	}
	reqMsgId, err := d.Int64()
	if err != nil {
		t.Fatalf("Int64() error = %v", err)
	}
	if reqMsgId != 99 {
		t.Fatalf("req_msg_id = %d, want 99", reqMsgId)
	}
	if !bytes.Equal(d.Raw(), result) {
		t.Fatalf("result = %x, want %x", d.Raw(), result)
	}
}

func TestWrapRPCError(t *testing.T) {
	wrapped, err := WrapRPCError(100, 400, "MESSAGE_ID_INVALID")
	if err != nil {
		t.Fatalf("WrapRPCError() error = %v", err)
	}

	d := bin.NewDecoder(wrapped)
	if clazz, _ := d.ClazzID(); clazz != mt.ClazzID_rpc_result {
		t.Fatalf("clazz = %#x, want rpc_result", clazz)
	}
	if reqMsgId, _ := d.Int64(); reqMsgId != 100 {
		t.Fatalf("req_msg_id = %d, want 100", reqMsgId)
	}
	errClazz, err := d.ClazzID()
	if err != nil {
		t.Fatalf("rpc_error clazz error = %v", err)
	}
	rpcErr := &mt.TLRpcError{ClazzID: errClazz}
	if err := rpcErr.Decode(d); err != nil {
		t.Fatalf("rpc_error Decode() error = %v", err)
	}
	if rpcErr.ErrorCode != 400 || rpcErr.ErrorMessage != "MESSAGE_ID_INVALID" {
		t.Fatalf("rpc_error = %#v", rpcErr)
	}
}

func TestUnwrapClientRPCInvokeWithLayerAndInitConnection(t *testing.T) {
	inner := encodeTL(t, &mt.TLPing{PingId: 10})
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
	payload := encodeTL(t, &tg.TLInvokeWithLayer{
		Layer: 224,
		Query: initConn,
	})

	gotInner, md, err := UnwrapClientRPC(payload)
	if err != nil {
		t.Fatalf("UnwrapClientRPC() error = %v", err)
	}
	if !bytes.Equal(gotInner, inner) {
		t.Fatalf("inner = %x, want %x", gotInner, inner)
	}
	if md.Layer != 224 || md.Client != "tdesktop macOS 5.0" || md.Langpack != "tdesktop" || md.LangCode != "en" {
		t.Fatalf("metadata = %#v", md)
	}
}

func TestUnwrapClientRPCLeavesBusinessMethodRaw(t *testing.T) {
	payload := encodeTL(t, &mt.TLPing{PingId: 11})
	got, md, err := UnwrapClientRPC(payload)
	if err != nil {
		t.Fatalf("UnwrapClientRPC() error = %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("payload = %x, want %x", got, payload)
	}
	if md != (WrapperMetadata{}) {
		t.Fatalf("metadata = %#v, want zero", md)
	}
}

func encodeTL(t *testing.T, obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	t.Helper()
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 224); err != nil {
		t.Fatalf("Encode(%T) error = %v", obj, err)
	}
	return append([]byte(nil), x.Bytes()...)
}

func testAuthKeys() (*crypto.AuthKey, *crypto.AuthKey) {
	keyData := make([]byte, 256)
	for i := range keyData {
		keyData[i] = byte(i)
	}
	return crypto.NewAuthKey(0, keyData), crypto.NewClientAuthKey(0, keyData)
}
