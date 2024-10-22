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
		meId   = in.GetUserId()
		idList = in.GetIdList()
	)

	dlgExtList, err := c.svcCtx.Dao.GetDialogListByIdList(c.ctx, meId, idList)
	if err != nil {
		c.Logger.Errorf("dialog.getDialogsByIdList - error: %v", err)
		return nil, err
	}

	for _, id := range idList {
		found := false
		for _, dlgExt := range dlgExtList {
			peer := mtproto.FromPeer(dlgExt.GetDialog().GetPeer())
			if mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId) == id {
				found = true
				break
			}
		}
		if !found {
			peerType, peerId := mtproto.GetPeerUtilByPeerDialogId(id)
			dlgExtList = append(dlgExtList, c.svcCtx.Dao.MakeDialog(&dataobject.DialogsDO{
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
		Datas: dlgExtList,
	}, nil
}
