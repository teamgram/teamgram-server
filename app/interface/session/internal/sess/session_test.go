package sess

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
)

func TestSessionConnNewRefreshesGatewayIDWhileOnline(t *testing.T) {
	s := newSession(1, &SessionList{})

	s.onSessionConnNew(context.Background(), "gateway-a")
	if got := s.getGatewayId(); got != "gateway-a" {
		t.Fatalf("expected initial gateway id gateway-a, got %q", got)
	}
	if s.connState != kStateOnline {
		t.Fatalf("expected session to be online after first connection, got %d", s.connState)
	}

	s.onSessionConnNew(context.Background(), "gateway-b")
	if got := s.getGatewayId(); got != "gateway-b" {
		t.Fatalf("expected gateway id to refresh to gateway-b, got %q", got)
	}
}

func TestSessionCloseIgnoresStaleGatewayAfterSwitch(t *testing.T) {
	s := newSession(1, &SessionList{})

	s.onSessionConnNew(context.Background(), "gateway-a")
	s.onSessionConnNew(context.Background(), "gateway-b")
	s.onSessionConnClose(context.Background(), "gateway-a")

	if s.connState != kStateOnline {
		t.Fatalf("expected stale gateway close to keep session online, got %d", s.connState)
	}
	if got := s.getGatewayId(); got != "gateway-b" {
		t.Fatalf("expected active gateway to remain gateway-b, got %q", got)
	}
}

func TestOnSyncRpcResultDataRemovesPendingAndQueuesResult(t *testing.T) {
	s := newSession(1, &SessionList{})

	const reqMsgID int64 = 1001
	s.pendingQueue.Add(reqMsgID)
	s.onSyncRpcResultData(context.Background(), reqMsgID, []byte{1, 2, 3})

	if got := s.pendingQueue.q.Len(); got != 0 {
		t.Fatalf("expected pending queue to be empty, got %d", got)
	}

	oMsg := s.outQueue.Lookup(reqMsgID)
	if oMsg == nil {
		t.Fatal("expected rpc result to be queued")
	}
	if oMsg.msg == nil {
		t.Fatal("expected raw message payload, got nil")
	}
	if oMsg.msg.Bytes != 3 {
		t.Fatalf("expected raw message bytes=3, got %d", oMsg.msg.Bytes)
	}
	if string(oMsg.msg.Body) != string([]byte{1, 2, 3}) {
		t.Fatalf("expected rpc result body to match input")
	}
	if _, ok := any(oMsg.msg).(*mt.TLMessageRawData); !ok {
		t.Fatal("expected queued message to be TLMessageRawData")
	}
}
