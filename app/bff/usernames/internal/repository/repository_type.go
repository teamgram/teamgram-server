package repository

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	// MinUsernameLen is the minimum length for a username.
	MinUsernameLen = 5
)

// Type aliases for convenience in the Logic layer.
type (
	ImmutableUser = tg.ImmutableUser
	MutableUsers  = tg.MutableUsers
	ChatClazz     = tg.ChatClazz
)
