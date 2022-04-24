// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package plugin

import (
	"context"
)

type MsgPlugin interface {
	ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error
}
