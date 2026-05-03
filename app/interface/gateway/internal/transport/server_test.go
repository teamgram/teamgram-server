package transport

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/sessionstate"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type transportAuthKeyStore struct {
	key *tg.AuthKeyInfo
}

func (s transportAuthKeyStore) QueryAuthKey(ctx context.Context, authKeyId int64) (*tg.AuthKeyInfo, error) {
	return s.key, nil
}

func (s transportAuthKeyStore) SetAuthKey(ctx context.Context, authKey *tg.AuthKeyInfo, futureSalt *tg.FutureSalt, expiresIn int32) error {
	return nil
}

func (s transportAuthKeyStore) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.FutureSalts, error) {
	return nil, nil
}

func (s transportAuthKeyStore) GetUserId(ctx context.Context, authKeyId int64) (int64, error) {
	return 0, nil
}

func (s transportAuthKeyStore) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) error {
	return nil
}

func (s transportAuthKeyStore) SetLayer(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	return nil
}

func (s transportAuthKeyStore) GetClientSession(ctx context.Context, authKeyId int64) (*authsession.ClientSession, error) {
	return nil, nil
}

type transportDispatcher struct{}

func (d transportDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	return encodeTransportTL(&mt.TLPong{MsgId: md.ClientMsgId, PingId: 1}), nil
}

func TestGatewayTransportServerRegistersPushWriter(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	writer := push.NewLocalWriter()
	server := NewServer(
		"",
		"gateway-test",
		nil,
		sessionstate.NewProcessor(transportAuthKeyStore{key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm)}, transportDispatcher{}),
		writer,
	)

	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ServeConn(context.Background(), serverConn)
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}

	codec := AbridgedCodec{}
	reqMsgID := int64(100)
	request := encodeTransportTL(&mt.TLPing{PingId: 9})
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     reqMsgID,
		SeqNo:     1,
		Body:      request,
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	if err := codec.WriteFrame(clientConn, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	syncFrame, err := codec.ReadFrame(clientConn)
	if err != nil {
		t.Fatalf("read encrypted response: %v", err)
	}
	syncResp, err := gmtproto.DecodeEncryptedMessage(syncFrame, clientKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage(sync) error = %v", err)
	}
	if syncResp.SeqNo != 1 {
		t.Fatalf("sync seq_no = %d, want 1", syncResp.SeqNo)
	}

	body := encodeTransportTL(&mt.TLPong{MsgId: reqMsgID, PingId: 10})
	pushErr := make(chan error, 1)
	go func() {
		ok, err := writer.WriteRPCResult(context.Background(), serverKey.AuthKeyId(), 77, reqMsgID, body)
		if err == nil && !ok {
			err = errPushTargetMissing
		}
		pushErr <- err
	}()
	pushedFrame, err := codec.ReadFrame(clientConn)
	if err != nil {
		t.Fatalf("read pushed frame: %v", err)
	}
	if err := <-pushErr; err != nil {
		t.Fatalf("WriteRPCResult() error = %v", err)
	}
	pushed, err := gmtproto.DecodeEncryptedMessage(pushedFrame, clientKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage() error = %v", err)
	}
	if pushed.AuthKeyId != serverKey.AuthKeyId() || pushed.SessionId != 77 || pushed.Salt != 55 {
		t.Fatalf("pushed envelope = %#v", pushed)
	}
	if pushed.SeqNo != 3 {
		t.Fatalf("pushed seq_no = %d, want 3 after synchronous response seq_no 1", pushed.SeqNo)
	}

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after client close")
	}
}

func TestGatewayTransportStopClosesActiveConnections(t *testing.T) {
	server := NewServer("", "gateway-test", nil, nil, nil)
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ServeConn(context.Background(), serverConn)
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	if err := server.Stop(); err != nil {
		t.Fatalf("Stop() error = %v", err)
	}
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after Stop")
	}
}

func TestGatewayTransportServeConnAfterStopClosesImmediately(t *testing.T) {
	server := NewServer("", "gateway-test", nil, nil, nil)
	if err := server.Stop(); err != nil {
		t.Fatalf("Stop() error = %v", err)
	}
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	done := make(chan error, 1)
	go func() {
		done <- server.ServeConn(context.Background(), serverConn)
	}()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("ServeConn() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after Stop")
	}
}

func TestGatewayTransportReportsDetectCodecFailure(t *testing.T) {
	server := NewServer("", "gateway-test", nil, nil, nil)
	events := make(chan transportEvent, 1)
	server.eventSink = func(event transportEvent) {
		events <- event
	}

	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	done := make(chan error, 1)
	go func() {
		done <- server.ServeConn(context.Background(), serverConn)
	}()
	if _, err := clientConn.Write([]byte("GET / HTTP/1.1\r\n\r\n")); err != nil {
		t.Fatalf("write http probe: %v", err)
	}

	var serveErr error
	select {
	case serveErr = <-done:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after bad transport")
	}
	if !errors.Is(serveErr, ErrUnsupportedTransport) {
		t.Fatalf("ServeConn() error = %v, want ErrUnsupportedTransport", serveErr)
	}

	select {
	case event := <-events:
		if event.Phase != "detect_codec" {
			t.Fatalf("event phase = %q, want detect_codec", event.Phase)
		}
		if !errors.Is(event.Err, ErrUnsupportedTransport) {
			t.Fatalf("event err = %v, want ErrUnsupportedTransport", event.Err)
		}
	case <-time.After(time.Second):
		t.Fatal("missing transport event")
	}
}

var errPushTargetMissing = &testError{"push target missing"}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}

func transportTestKeys() (*crypto.AuthKey, *crypto.AuthKey) {
	keyData := make([]byte, 256)
	for i := range keyData {
		keyData[i] = byte(255 - i)
	}
	return crypto.NewAuthKey(0, keyData), crypto.NewClientAuthKey(0, keyData)
}

func encodeTransportTL(obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 0); err != nil {
		panic(err)
	}
	return append([]byte(nil), x.Bytes()...)
}
