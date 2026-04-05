package sess

import (
	"errors"
	"testing"

	"github.com/teamgram/proto/mtproto"
)

func TestRPCErrorForAuthStateUsesRestartForTransientStates(t *testing.T) {
	for _, state := range []int{
		mtproto.AuthStateNew,
		mtproto.AuthStateWaitInit,
		mtproto.AuthStateUnauthorized,
	} {
		err := rpcErrorForAuthState(state)
		if !errors.Is(err, mtproto.ErrAuthRestart) {
			t.Fatalf("state %d: got %v, want AUTH_RESTART", state, err)
		}
	}
}

func TestRPCErrorForAuthStateUsesAuthKeyUnregisteredForTerminalStates(t *testing.T) {
	for _, state := range []int{
		mtproto.AuthStateUnknown,
		mtproto.AuthStateLogout,
		mtproto.AuthStateDeleted,
	} {
		err := rpcErrorForAuthState(state)
		if !errors.Is(err, mtproto.ErrAuthKeyUnregistered) {
			t.Fatalf("state %d: got %v, want AUTH_KEY_UNREGISTERED", state, err)
		}
	}
}

func TestShouldRefreshAuthState(t *testing.T) {
	cases := []struct {
		state int
		want  bool
	}{
		{mtproto.AuthStateNew, true},
		{mtproto.AuthStateWaitInit, true},
		{mtproto.AuthStateUnauthorized, true},
		{mtproto.AuthStateNeedPassword, false},
		{mtproto.AuthStateNormal, false},
		{mtproto.AuthStateLogout, false},
	}

	for _, c := range cases {
		if got := shouldRefreshAuthState(c.state); got != c.want {
			t.Fatalf("state %d: got %v, want %v", c.state, got, c.want)
		}
	}
}
