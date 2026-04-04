package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthLogOutReturnsLoggedOutPlaceholder(t *testing.T) {
	core, _, _ := newAuthorizationCoreForAuthTest(t, nil)

	resp, err := core.AuthLogOut(&tg.TLAuthLogOut{})
	if err != nil {
		t.Fatalf("AuthLogOut returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected logged out response, got nil")
	}
	if len(resp.FutureAuthToken) != 0 {
		t.Fatalf("expected empty future auth token placeholder, got %d bytes", len(resp.FutureAuthToken))
	}
}
