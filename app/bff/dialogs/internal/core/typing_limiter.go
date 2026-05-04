package core

import (
	"time"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
)

func newTypingLimiter(interval time.Duration) svc.TypingLimiter {
	return svc.NewTypingLimiter(interval)
}
