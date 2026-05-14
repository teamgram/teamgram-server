package transport

import (
	"context"
	"errors"
	"fmt"
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

func TestGatewayTransportDriverPassesDecodedFramesAndWritesResponsesInOrder(t *testing.T) {
	codec := AbridgedCodec{}
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	handler := &transportDriverHandler{
		onFrame: func(ctx context.Context, conn Connection, frame []byte) error {
			if string(frame) == "one1" {
				if err := conn.WriteFrame(ctx, []byte("r001")); err != nil {
					return err
				}
			}
			if err := conn.WriteFrame(ctx, append([]byte("e"), frame[:3]...)); err != nil {
				return err
			}
			return nil
		},
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- newNetDriver(serverConn, handler).Serve(context.Background())
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	if err := codec.WriteFrame(clientConn, []byte("one1")); err != nil {
		t.Fatalf("WriteFrame(one) error = %v", err)
	}
	for i, want := range [][]byte{[]byte("r001"), []byte("eone")} {
		got, err := codec.ReadFrame(clientConn)
		if err != nil {
			t.Fatalf("ReadFrame(%d) error = %v", i, err)
		}
		if string(got) != string(want) {
			t.Fatalf("frame %d = %q, want %q", i, got, want)
		}
	}
	if err := codec.WriteFrame(clientConn, []byte("two2")); err != nil {
		t.Fatalf("WriteFrame(two) error = %v", err)
	}
	got, err := codec.ReadFrame(clientConn)
	if err != nil {
		t.Fatalf("ReadFrame(2) error = %v", err)
	}
	if string(got) != "etwo" {
		t.Fatalf("frame 2 = %q, want etwo", got)
	}
	handler.mu.Lock()
	received := append([]string(nil), handler.frames...)
	handler.mu.Unlock()
	if fmt.Sprint(received) != "[one1 two2]" {
		t.Fatalf("handler frames = %v, want [one1 two2]", received)
	}

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("driver did not exit after client close")
	}
	if got := handler.closeCount(); got != 1 {
		t.Fatalf("OnClose calls = %d, want 1", got)
	}
}

func TestGatewayTransportDriverOnFrameErrorClosesOnce(t *testing.T) {
	codec := AbridgedCodec{}
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	frameErr := errors.New("frame failed")
	handler := &transportDriverHandler{
		onFrame: func(ctx context.Context, conn Connection, frame []byte) error {
			return frameErr
		},
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- newNetDriver(serverConn, handler).Serve(context.Background())
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	if err := codec.WriteFrame(clientConn, []byte("boom")); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	select {
	case err := <-errCh:
		if !errors.Is(err, frameErr) {
			t.Fatalf("Serve() error = %v, want %v", err, frameErr)
		}
	case <-time.After(time.Second):
		t.Fatal("driver did not exit after OnFrame error")
	}
	if got := handler.closeCount(); got != 1 {
		t.Fatalf("OnClose calls = %d, want 1", got)
	}
	if handler.closeErr == nil || !errors.Is(handler.closeErr, frameErr) {
		t.Fatalf("OnClose error = %v, want %v", handler.closeErr, frameErr)
	}
}

func TestGatewayTransportDriverConnectionCloseClosesOnce(t *testing.T) {
	codec := AbridgedCodec{}
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	closed := make(chan struct{})
	handler := &transportDriverHandler{
		onOpen: func(ctx context.Context, conn Connection) {
			if err := conn.Close(); err != nil {
				t.Errorf("Close() error = %v", err)
			}
			close(closed)
		},
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- newNetDriver(serverConn, handler).Serve(context.Background())
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("handler did not close connection")
	}
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("driver did not exit after Connection.Close")
	}
	if got := handler.closeCount(); got != 1 {
		t.Fatalf("OnClose calls = %d, want 1", got)
	}
	if err := codec.WriteFrame(clientConn, []byte("xxxx")); err == nil {
		t.Fatal("client WriteFrame after server close succeeded, want error")
	}
}

func TestGatewayTransportDriverWriteFrameAfterCloseOrCanceledContextReturnsClearError(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	connCh := make(chan Connection, 1)
	handler := &transportDriverHandler{
		onOpen: func(ctx context.Context, conn Connection) {
			connCh <- conn
		},
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- newNetDriver(serverConn, handler).Serve(context.Background())
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	var conn Connection
	select {
	case conn = <-connCh:
	case <-time.After(time.Second):
		t.Fatal("OnOpen did not receive connection")
	}
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := conn.WriteFrame(canceled, []byte("payload")); !errors.Is(err, context.Canceled) {
		t.Fatalf("WriteFrame(canceled) error = %v, want context.Canceled", err)
	}
	if err := conn.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if err := conn.WriteFrame(context.Background(), []byte("payload")); !errors.Is(err, errConnectionClosed) {
		t.Fatalf("WriteFrame(closed) error = %v, want %v", err, errConnectionClosed)
	}
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("driver did not exit after Connection.Close")
	}
	if got := handler.closeCount(); got != 1 {
		t.Fatalf("OnClose calls = %d, want 1", got)
	}
}

func TestGatewayTransportDriverWriteCountOnlyIncludesNilWriteFrame(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()
	acceptedCh := make(chan int, 1)
	handler := &transportDriverHandler{
		onOpen: func(ctx context.Context, conn Connection) {
			accepted := 0
			if err := conn.WriteFrame(ctx, []byte("accepted")); err == nil {
				accepted++
			}
			_ = conn.Close()
			if err := conn.WriteFrame(ctx, []byte("rejected")); err == nil {
				accepted++
			}
			acceptedCh <- accepted
		},
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- newNetDriver(serverConn, handler).Serve(context.Background())
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write transport flag: %v", err)
	}
	codec := AbridgedCodec{}
	frame, err := codec.ReadFrame(clientConn)
	if err != nil {
		t.Fatalf("ReadFrame() error = %v", err)
	}
	if string(frame) != "accepted" {
		t.Fatalf("frame = %q, want accepted", frame)
	}
	select {
	case count := <-acceptedCh:
		if count != 1 {
			t.Fatalf("accepted writes = %d, want 1", count)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for accepted write count")
	}
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("driver did not exit after Connection.Close")
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
	request := encodeTransportTLLayer(&tg.TLHelpGetConfig{}, 224)
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

func TestGatewayTransportRegistersMainUpdatesFlag(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	writer := push.NewLocalWriter()
	server := NewServer(
		"",
		"gateway-test",
		"127.0.0.1:20110",
		"generation-test",
		nil,
		sessionstate.NewProcessor(transportAuthKeyStore{
			key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
			clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
				AuthKeyId: serverKey.AuthKeyId(),
				Layer:     224,
			}).ToClientSession(),
		}, transportDispatcher{}),
		writer,
		nil,
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
	reqMsgID := int64(150)
	payload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     reqMsgID,
		SeqNo:     1,
		Body:      encodeTransportTLLayer(&tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}, 224),
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

	pushErr := make(chan error, 1)
	go func() {
		count, err := writer.WriteUpdates(context.Background(), serverKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
		if err == nil && count != 1 {
			err = fmt.Errorf("WriteUpdates count = %d, want 1", count)
		}
		pushErr <- err
	}()
	pushedFrame, err := readTransportFrameForTest(t, codec, clientConn)
	if err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-pushErr:
		if err != nil {
			t.Fatalf("WriteUpdates() error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for WriteUpdates")
	}
	pushed, err := gmtproto.DecodeEncryptedMessage(pushedFrame, clientKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage() error = %v", err)
	}
	if pushed.AuthKeyId != serverKey.AuthKeyId() || pushed.SessionId != 77 {
		t.Fatalf("pushed envelope = %#v", pushed)
	}

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after client close")
	}
}

func TestGatewayTransportRefreshesMainUpdatesAfterSelectorSuccess(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	writer := push.NewLocalWriter()
	server := NewServer(
		"",
		"gateway-test",
		"127.0.0.1:20110",
		"generation-test",
		nil,
		sessionstate.NewProcessor(transportAuthKeyStore{
			key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
			clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
				AuthKeyId: serverKey.AuthKeyId(),
				Layer:     224,
			}).ToClientSession(),
		}, transportDispatcher{}),
		writer,
		nil,
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
		MsgId:     160,
		SeqNo:     1,
		Body:      encodeTransportTLLayer(&tg.TLUpdatesGetState{}, 224),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage() error = %v", err)
	}
	if err := codec.WriteFrame(clientConn, payload); err != nil {
		t.Fatalf("WriteFrame() error = %v", err)
	}
	if _, err := codec.ReadFrame(clientConn); err != nil {
		t.Fatalf("read selector response: %v", err)
	}

	pushErr := make(chan error, 1)
	go func() {
		count, err := writer.WriteUpdates(context.Background(), serverKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
		if err == nil && count != 1 {
			err = fmt.Errorf("WriteUpdates count = %d, want 1", count)
		}
		pushErr <- err
	}()
	pushedFrame, err := readTransportFrameForTest(t, codec, clientConn)
	if err != nil {
		t.Fatal(err)
	}
	if err := <-pushErr; err != nil {
		t.Fatalf("WriteUpdates() error = %v", err)
	}
	pushed, err := gmtproto.DecodeEncryptedMessage(pushedFrame, clientKey)
	if err != nil {
		t.Fatalf("DecodeEncryptedMessage() error = %v", err)
	}
	if pushed.AuthKeyId != serverKey.AuthKeyId() || pushed.SessionId != 77 {
		t.Fatalf("pushed envelope = %#v", pushed)
	}

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("ServeConn did not exit after client close")
	}
}

func TestGatewayTransportCloseClearsMainUpdatesRole(t *testing.T) {
	serverKey, clientKey := transportTestKeys()
	writer := push.NewLocalWriter()
	processor := sessionstate.NewProcessor(transportAuthKeyStore{
		key: tg.NewAuthKeyInfo(serverKey.AuthKeyId(), serverKey.AuthKey(), tg.AuthKeyTypePerm),
		clientSession: authsession.MakeTLClientSession(&authsession.TLClientSession{
			AuthKeyId: serverKey.AuthKeyId(),
			Layer:     224,
		}).ToClientSession(),
	}, transportDispatcher{})
	server := NewServer("", "gateway-test", "127.0.0.1:20110", "generation-test", nil, processor, writer, nil)

	clientConn, serverConn := net.Pipe()
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ServeConn(context.Background(), serverConn)
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write first transport flag: %v", err)
	}
	codec := AbridgedCodec{}
	promotePayload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     170,
		SeqNo:     1,
		Body:      encodeTransportTLLayer(&tg.TLAccountUpdateStatus{Offline: tg.BoolFalseClazz}, 224),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage(promote) error = %v", err)
	}
	if err := codec.WriteFrame(clientConn, promotePayload); err != nil {
		t.Fatalf("WriteFrame(promote) error = %v", err)
	}
	if _, err := codec.ReadFrame(clientConn); err != nil {
		t.Fatalf("read promote response: %v", err)
	}
	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("first ServeConn did not exit after client close")
	}

	clientConn, serverConn = net.Pipe()
	defer clientConn.Close()
	errCh = make(chan error, 1)
	go func() {
		errCh <- server.ServeConn(context.Background(), serverConn)
	}()
	if _, err := clientConn.Write([]byte{abridgedFlag}); err != nil {
		t.Fatalf("write second transport flag: %v", err)
	}
	checkPayload, err := gmtproto.EncodeEncryptedMessage(gmtproto.EncryptedMessage{
		AuthKeyId: clientKey.AuthKeyId(),
		Salt:      55,
		SessionId: 77,
		MsgId:     171,
		SeqNo:     1,
		Body:      encodeTransportTLLayer(&tg.TLHelpGetConfig{}, 224),
	}, clientKey)
	if err != nil {
		t.Fatalf("EncodeEncryptedMessage(check) error = %v", err)
	}
	if err := codec.WriteFrame(clientConn, checkPayload); err != nil {
		t.Fatalf("WriteFrame(check) error = %v", err)
	}
	if _, err := codec.ReadFrame(clientConn); err != nil {
		t.Fatalf("read check response: %v", err)
	}

	pushErr := make(chan error, 1)
	go func() {
		count, err := writer.WriteUpdates(context.Background(), serverKey.AuthKeyId(), tg.MakeTLUpdatesTooLong(&tg.TLUpdatesTooLong{}))
		if err == nil && count != 0 {
			err = fmt.Errorf("WriteUpdates count = %d, want 0 after main session close", count)
		}
		pushErr <- err
	}()
	readCh := make(chan []byte, 1)
	readErr := make(chan error, 1)
	go func() {
		frame, err := codec.ReadFrame(clientConn)
		if err != nil {
			readErr <- err
			return
		}
		readCh <- frame
	}()
	select {
	case err := <-pushErr:
		if err != nil {
			t.Fatalf("WriteUpdates() error = %v", err)
		}
	case frame := <-readCh:
		t.Fatalf("unexpected pushed frame after main session close: %x", frame)
	case err := <-readErr:
		t.Fatalf("read pushed frame: %v", err)
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for WriteUpdates after main session close")
	}

	_ = clientConn.Close()
	select {
	case <-errCh:
	case <-time.After(time.Second):
		t.Fatal("second ServeConn did not exit after client close")
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
		Body:      encodeTransportTLLayer(&tg.TLHelpGetConfig{}, 224),
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

func readTransportFrameForTest(t *testing.T, codec Codec, conn net.Conn) ([]byte, error) {
	t.Helper()
	pushedFrameCh := make(chan []byte, 1)
	readErr := make(chan error, 1)
	go func() {
		frame, err := codec.ReadFrame(conn)
		if err != nil {
			readErr <- err
			return
		}
		pushedFrameCh <- frame
	}()
	select {
	case frame := <-pushedFrameCh:
		return frame, nil
	case err := <-readErr:
		return nil, fmt.Errorf("read pushed updates frame: %w", err)
	case <-time.After(time.Second):
		return nil, fmt.Errorf("timed out waiting for pushed updates frame")
	}
}

type transportDriverHandler struct {
	mu       sync.Mutex
	frames   []string
	closed   int
	closeErr error
	onOpen   func(context.Context, Connection)
	onFrame  func(context.Context, Connection, []byte) error
}

func (h *transportDriverHandler) OnOpen(ctx context.Context, conn Connection) {
	if h.onOpen != nil {
		h.onOpen(ctx, conn)
	}
}

func (h *transportDriverHandler) OnFrame(ctx context.Context, conn Connection, frame []byte) error {
	h.mu.Lock()
	h.frames = append(h.frames, string(frame))
	h.mu.Unlock()
	if h.onFrame != nil {
		return h.onFrame(ctx, conn, frame)
	}
	return nil
}

func (h *transportDriverHandler) OnClose(ctx context.Context, conn Connection, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.closed++
	h.closeErr = err
}

func (h *transportDriverHandler) closeCount() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.closed
}

func encodeTransportTL(obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	return encodeTransportTLLayer(obj, 0)
}

func encodeTransportTLLayer(obj interface {
	Encode(*bin.Encoder, int32) error
}, layer int32) []byte {
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, layer); err != nil {
		panic(err)
	}
	return append([]byte(nil), x.Bytes()...)
}
