// Package plugin defines the UsernamesPlugin interface for
// enterprise features.
package plugin

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UsernamesPlugin provides enterprise username features.
type UsernamesPlugin interface {
	GetChannelListByIdList(ctx context.Context, selfId int64, id ...int64) []tg.ChatClazz
}
