package netserver

import (
	"context"
	"net"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gnetway/gnetway"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func makeTestAuthKey(id int64) *authKeyUtil {
	authKey := make([]byte, 256)
	for i := range authKey {
		authKey[i] = byte((int(id) + i) % 251)
	}
	return newAuthKeyUtil(&tg.TLAuthKeyInfo{
		AuthKeyId:     id,
		AuthKey:       authKey,
		AuthKeyType:   1,
		PermAuthKeyId: id,
	})
}

func makeTestServerAndConn(t *testing.T, authKey *authKeyUtil, sessionID, connID int64) (*Server, *connection, func()) {
	t.Helper()

	clientConn, serverConn := net.Pipe()

	s := &Server{
		authSessionMgr: NewAuthSessionManager(),
		connMgr:        newConnectionManager(),
	}

	conn := s.connMgr.newConnection(connID, serverConn, true, false)
	if authKey != nil {
		conn.putAuthKey(authKey)
		s.authSessionMgr.AddNewSession(authKey, sessionID, connID)
	}

	cleanup := func() {
		_ = conn.Close()
		_ = clientConn.Close()
	}

	return s, conn, cleanup
}

func TestGnetwaySendDataToGatewayReturnsFalseWhenConnectionAuthKeyMissing(t *testing.T) {
	authKey := makeTestAuthKey(101)
	s, conn, cleanup := makeTestServerAndConn(t, nil, 2001, 3001)
	defer cleanup()

	s.authSessionMgr.AddNewSession(authKey, 2001, conn.ID())

	reply, err := s.GnetwaySendDataToGateway(context.Background(), &gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: authKey.AuthKeyId(),
		SessionId: 2001,
		Payload:   []byte("hello"),
	})
	if err != nil {
		t.Fatalf("GnetwaySendDataToGateway returned error: %v", err)
	}
	if reply != tg.BoolFalse {
		t.Fatalf("expected BoolFalse when connection auth key is missing, got %v", reply)
	}
}

func TestGnetwaySendDataToGatewayReturnsFalseWhenNoConnectionMatchesAuthKey(t *testing.T) {
	authKey := makeTestAuthKey(102)
	s, _, cleanup := makeTestServerAndConn(t, makeTestAuthKey(999), 2002, 3002)
	defer cleanup()

	s.authSessionMgr.AddNewSession(authKey, 2002, 3002)

	reply, err := s.GnetwaySendDataToGateway(context.Background(), &gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: authKey.AuthKeyId(),
		SessionId: 2002,
		Payload:   []byte("hello"),
	})
	if err != nil {
		t.Fatalf("GnetwaySendDataToGateway returned error: %v", err)
	}
	if reply != tg.BoolFalse {
		t.Fatalf("expected BoolFalse when all matching connections have the wrong auth key, got %v", reply)
	}
}

func TestGnetwaySendDataToGatewayReturnsTrueWhenAtLeastOneConnectionMatches(t *testing.T) {
	authKey := makeTestAuthKey(103)
	s, _, cleanup1 := makeTestServerAndConn(t, authKey, 2003, 3003)
	defer cleanup1()

	clientConn2, serverConn2 := net.Pipe()
	defer clientConn2.Close()
	defer serverConn2.Close()

	badConn := s.connMgr.newConnection(3004, serverConn2, true, false)
	badConn.putAuthKey(makeTestAuthKey(1003))
	s.authSessionMgr.AddNewSession(authKey, 2003, badConn.ID())

	reply, err := s.GnetwaySendDataToGateway(context.Background(), &gnetway.TLGnetwaySendDataToGateway{
		AuthKeyId: authKey.AuthKeyId(),
		SessionId: 2003,
		Payload:   []byte("hello"),
	})
	if err != nil {
		t.Fatalf("GnetwaySendDataToGateway returned error: %v", err)
	}
	if reply != tg.BoolTrue {
		t.Fatalf("expected BoolTrue when at least one connection matches, got %v", reply)
	}
}
