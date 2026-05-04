package transport

import (
	"context"
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	gmtproto "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/mtproto"
	gatewaypresence "github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/presence"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/push"
	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/internal/sessionstate"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	presencepb "github.com/teamgram/teamgram-server/v2/app/service/presence/presence"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type transportAuthKeyStore struct {
	key           *tg.AuthKeyInfo
	userID        int64
	clientSession *authsession.ClientSession
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
	return s.userID, nil
}

func (s transportAuthKeyStore) SetClientSessionInfo(ctx context.Context, session *authsession.ClientSession) error {
	return nil
}

func (s transportAuthKeyStore) SetLayer(ctx context.Context, authKeyId int64, ip string, layer int32) error {
	return nil
}

func (s transportAuthKeyStore) GetClientSession(ctx context.Context, authKeyId int64) (*authsession.ClientSession, error) {
	return s.clientSession, nil
}

type transportDispatcher struct{}

func (d transportDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	return encodeTransportTL(&mt.TLPong{MsgId: md.ClientMsgId, PingId: 1}), nil
}

type transportPresenceClient struct {
	mu      sync.Mutex
	online  []*presencepb.OnlineSession
	offline []offlineCall
}

type offlineCall struct {
	userID    int64
	authKeyID int64
	sessionID int64
}

func (c *transportPresenceClient) SetSessionOnline(ctx context.Context, session *presencepb.OnlineSession) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.online = append(c.online, session)
	return nil
}

func (c *transportPresenceClient) SetSessionOffline(ctx context.Context, userID, authKeyID, sessionID int64) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.offline = append(c.offline, offlineCall{userID: userID, authKeyID: authKeyID, sessionID: sessionID})
	return nil
}

func (c *transportPresenceClient) onlineCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.online)
}

func (c *transportPresenceClient) offlineCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.offline)
}

func (c *transportPresenceClient) waitOnline(t *testing.T, want int) {
	t.Helper()
	deadline := time.After(time.Second)
	for {
		if c.onlineCount() >= want {
			return
		}
		select {
		case <-deadline:
			t.Fatalf("presence online calls = %d, want at least %d", c.onlineCount(), want)
		case <-time.After(time.Millisecond):
		}
	}
}

func TestGatewayTransportServerRegistersPushWriter(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	writer := push.NewLocalWriter()
	presenceClient := &transportPresenceClient{}
	registrar := gatewaypresence.NewRegistrar(gatewaypresence.Config{
		GatewayID:                  "gateway-test",
		GatewayGeneration:          "generation-test",
		GatewayRPCAddr:             "127.0.0.1:20110",
		RefreshMinInterval:         time.Second,
		RefreshRetryMinInterval:    time.Second,
		RefreshScanInterval:        time.Second,
		ShutdownOfflineDeadline:    time.Second,
		ShutdownOfflineMaxSessions: 100,
	}, presenceClient, time.Now)
	server := NewServer(
		"",
		"gateway-test",
		"127.0.0.1:20110",
		"generation-test",
		nil,
		sessionstate.NewProcessor(transportAuthKeyStore{
			key:    tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
			userID: 12345,
			clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
				AuthKeyId:   serverKey.AuthKeyId(),
				Layer:       224,
				DeviceModel: "tdesktop",
			}).ToClientSession(),
		}, transportDispatcher{}),
		writer,
		registrar,
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
	request := encodeTransportTL(&mt.TLGetFutureSalts{Num: 1})
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
	presenceClient.waitOnline(t, 1)
	if presenceClient.onlineCount() != 1 {
		t.Fatalf("presence online calls = %d, want 1", presenceClient.onlineCount())
	}
	presenceClient.mu.Lock()
	if got := presenceClient.online[0]; got.UserId != 12345 || got.AuthKeyId != serverKey.AuthKeyId() || got.SessionId != 77 || got.GatewayRpcAddr != "127.0.0.1:20110" || got.GatewayGeneration != "generation-test" {
		presenceClient.mu.Unlock()
		t.Fatalf("presence session = %#v", got)
	}
	presenceClient.mu.Unlock()

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

func TestGatewayTransportUnregistersPresenceWhenPushWriterNil(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	presenceClient := &transportPresenceClient{}
	registrar := gatewaypresence.NewRegistrar(gatewaypresence.Config{
		GatewayID:                  "gateway-test",
		GatewayGeneration:          "generation-test",
		GatewayRPCAddr:             "127.0.0.1:20110",
		RefreshMinInterval:         time.Second,
		RefreshRetryMinInterval:    time.Second,
		RefreshScanInterval:        time.Second,
		ShutdownOfflineDeadline:    time.Second,
		ShutdownOfflineMaxSessions: 100,
	}, presenceClient, time.Now)
	server := NewServer(
		"",
		"gateway-test",
		"127.0.0.1:20110",
		"generation-test",
		nil,
		sessionstate.NewProcessor(transportAuthKeyStore{
			key:    tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
			userID: 12345,
			clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
				AuthKeyId:   serverKey.AuthKeyId(),
				Layer:       224,
				DeviceModel: "tdesktop",
			}).ToClientSession(),
		}, transportDispatcher{}),
		nil,
		registrar,
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
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     100,
		SeqNo:     1,
		Body:      encodeTransportTL(&mt.TLGetFutureSalts{Num: 1}),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	if err := codec.WriteFrame(clientConn, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	if _, err := codec.ReadFrame(clientConn); err != nil {
		t.Fatalf("read encrypted response: %v", err)
	}
	presenceClient.waitOnline(t, 1)

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after client close")
	}
	if presenceClient.offlineCount() != 1 {
		t.Fatalf("presence offline calls = %d, want 1", presenceClient.offlineCount())
	}
}

func TestGatewayTransportStopClosesActiveConnections(t *testing.T) {
	server := NewServer("", "gateway-test", "127.0.0.1:20110", "generation-test", nil, nil, nil, nil)
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
	server := NewServer("", "gateway-test", "127.0.0.1:20110", "generation-test", nil, nil, nil, nil)
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
	server := NewServer("", "gateway-test", "127.0.0.1:20110", "generation-test", nil, nil, nil, nil)
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
