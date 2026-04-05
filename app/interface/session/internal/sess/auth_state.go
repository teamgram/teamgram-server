package sess

import "github.com/teamgram/proto/mtproto"

func shouldRefreshAuthState(state int) bool {
	switch state {
	case mtproto.AuthStateNew,
		mtproto.AuthStateWaitInit,
		mtproto.AuthStateUnauthorized:
		return true
	default:
		return false
	}
}

func rpcErrorForAuthState(state int) error {
	switch state {
	case mtproto.AuthStateNew,
		mtproto.AuthStateWaitInit,
		mtproto.AuthStateUnauthorized:
		return mtproto.ErrAuthRestart
	default:
		return mtproto.ErrAuthKeyUnregistered
	}
}
