package plugin

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type ChatPlugin interface {
	GetChatCallActiveAndNotEmpty(ctx context.Context, userID int64, chatID int64) (bool, bool)
	GetChatGroupCall(ctx context.Context, userID int64, chatID int64) tg.InputGroupCallClazz
}
