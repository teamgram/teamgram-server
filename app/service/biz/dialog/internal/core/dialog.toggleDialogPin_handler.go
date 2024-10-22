/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
)

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (c *DialogCore) DialogToggleDialogPin(in *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error) {
	var (
		peerDialogId = mtproto.MakePeerDialogId(in.PeerType, in.PeerId)
		pinned       int64
	)

	dlgExt, err := c.svcCtx.Dao.GetDialogByPeerDialogId(c.ctx, in.GetUserId(), peerDialogId)
	if err != nil {
		c.Logger.Errorf("dialog.toggleDialogPin - error: %v", err)
		return nil, err
	}

	folderId := dlgExt.GetDialog().GetFolderId().GetValue()

	if mtproto.FromBool(in.Pinned) {
		pinned = time.Now().Unix() << 32
	} else {
		pinned = 0
	}

	if folderId == 0 {
		c.svcCtx.Dao.CachedConn.Exec(
			c.ctx,
			func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
				c.svcCtx.Dao.DialogsDAO.UpdatePeerDialogListPinned(c.ctx, pinned, in.UserId, []int64{peerDialogId})
				return 0, 0, nil
			},
			dialog.GetDialogCacheKey(in.GetUserId(), peerDialogId),
			dialog.GetPinnedDialogIdListCacheKey(in.GetUserId()))
	} else {
		c.svcCtx.Dao.CachedConn.Exec(
			c.ctx,
			func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
				c.svcCtx.Dao.DialogsDAO.UpdateFolderPeerDialogListPinned(c.ctx, pinned, in.UserId, []int64{peerDialogId})
				return 0, 0, nil
			},
			dialog.GetDialogCacheKey(in.GetUserId(), peerDialogId),
			dialog.GetFolderPinnedDialogIdListCacheKey(in.GetUserId()))
	}

	return &mtproto.Int32{
		V: folderId,
	}, nil
}
