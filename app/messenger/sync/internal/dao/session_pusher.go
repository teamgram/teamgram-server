// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/teamgram-server/app/interface/session/session"
)

// SessionPusher abstracts push operations to a session node.
// Implemented by both Session (unary RPC) and StreamingSession (bidi stream).
type SessionPusher interface {
	PushUpdates(ctx context.Context, msg *session.TLSessionPushUpdatesData) error
	PushSessionUpdates(ctx context.Context, msg *session.TLSessionPushSessionUpdatesData) error
	PushRpcResult(ctx context.Context, msg *session.TLSessionPushRpcResultData) error
	Close() error
}
