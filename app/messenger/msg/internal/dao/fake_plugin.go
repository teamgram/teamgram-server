// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/teamgram/proto/mtproto"
)

func (d *Dao) ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error {
	return mtproto.ErrMethodNotImpl
}

func (d *Dao) UsernameResolveUsername(ctx context.Context, in *user.TLUserResolveUsername) (*mtproto.Peer, error) {
	return d.UserClient.UserResolveUsername(ctx, in)
}

func (d *Dao) GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error) {
	return nil, mtproto.ErrMethodNotImpl
}
