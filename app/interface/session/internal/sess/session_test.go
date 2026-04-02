package sess

import (
	"context"
	"testing"
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
