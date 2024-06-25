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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (c *DialogCore) DialogToggleDialogPin(in *dialog.TLDialogToggleDialogPin) (*mtproto.Int32, error) {
	var (
		peerDialogId = mtproto.MakePeerDialogId(in.PeerType, in.PeerId)
		pinned       int64
		folderId     int32
		dialogDO     *dataobject.DialogsDO
	)

	_, err := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogListWithCB(c.ctx,
		in.UserId,
		[]int64{peerDialogId},
		func(sz, i int, v *dataobject.DialogsDO) {
			dialogDO = v
			folderId = v.FolderId
		})
	if err != nil {
		c.Logger.Errorf("dialog.toggleDialogPin - error: %v", err)
		return nil, err
	}

	if dialogDO == nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("dialog.toggleDialogPin - error: %v", err)
		return nil, err
	}

	if mtproto.FromBool(in.Pinned) {
		pinned = time.Now().Unix() << 32
	} else {
		pinned = 0
	}

	if folderId == 0 {
		c.svcCtx.Dao.DialogsDAO.UpdatePeerDialogListPinned(c.ctx, pinned, in.UserId, []int64{peerDialogId})
	} else {
		c.svcCtx.Dao.DialogsDAO.UpdateFolderPeerDialogListPinned(c.ctx, pinned, in.UserId, []int64{peerDialogId})
	}

	return &mtproto.Int32{
		V: folderId,
	}, nil
}
