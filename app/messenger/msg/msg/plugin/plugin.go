// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package plugin

import (
	"context"

	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/teamgram/proto/mtproto"
)

type MsgPlugin interface {
	ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error
	UsernameResolveUsername(ctx context.Context, in *user.TLUserResolveUsername) (*mtproto.Peer, error)
	GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error)
}
