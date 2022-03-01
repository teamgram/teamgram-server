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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

// DialogGetDialogsByIdList
// dialog.getDialogsByIdList user_id:long id_list:Vector<long> = Vector<DialogExt>;
func (c *DialogCore) DialogGetDialogsByIdList(in *dialog.TLDialogGetDialogsByIdList) (*dialog.Vector_DialogExt, error) {
	var (
		dList dialog.DialogExtList
		meId  = in.GetUserId()
		// peerId := mtproto.MakePeerDialogId()
	)

	doList, _ := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogList(
		c.ctx,
		meId,
		in.IdList)

	for _, id := range in.IdList {
		found := false
		for i := 0; i < len(doList); i++ {
			if doList[i].PeerDialogId == id {
				found = true
				dList = append(dList, makeDialog(&doList[i]))
				break
			}
		}
		if !found {
			peerType, peerId := mtproto.GetPeerUtilByPeerDialogId(id)
			dList = append(dList, makeDialog(&dataobject.DialogsDO{
				UserId:           in.UserId,
				PeerType:         peerType,
				PeerId:           peerId,
				PeerDialogId:     id,
				Pinned:           0,
				TopMessage:       0,
				PinnedMsgId:      0,
				ReadInboxMaxId:   0,
				ReadOutboxMaxId:  0,
				UnreadCount:      0,
				UnreadMark:       false,
				DraftType:        0,
				DraftMessageData: "null",
				FolderId:         0,
				FolderPinned:     0,
				HasScheduled:     false,
				Date2:            0,
			}))
		}
	}

	return &dialog.Vector_DialogExt{
		Datas: dList,
	}, nil
}
