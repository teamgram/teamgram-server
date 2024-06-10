// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
)

func (d *Dao) ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error {
	return mtproto.ErrMethodNotImpl
}

func (d *Dao) UsernameResolveUsername(ctx context.Context, in *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	return d.UsernameClient.UsernameResolveUsername(ctx, in)
}

func (d *Dao) GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error) {
	return nil, mtproto.ErrMethodNotImpl
}
