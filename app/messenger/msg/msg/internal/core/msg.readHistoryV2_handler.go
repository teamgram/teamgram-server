// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"math"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
)

// MsgReadHistoryV2
// msg.readHistoryV2 user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (c *MsgCore) MsgReadHistoryV2(in *msg.TLMsgReadHistoryV2) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		maxId         = in.MaxId
	)

	dlg, err := c.svcCtx.Dao.DialogsDAO.SelectDialog(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("messages.readHistory - error: invalid peer %v", err)
		return nil, mtproto.ErrInternalServerError
	} else if dlg == nil {
		c.Logger.Errorf("messages.readHistory - error: not found dialog, request: %s", in)
		return nil, mtproto.ErrPeerIdInvalid
	}

	if maxId == 0 || maxId == math.MaxInt32 {
		maxId = dlg.TopMessage
	}

	if in.PeerType == mtproto.PEER_SELF || in.PeerType == mtproto.PEER_USER && in.PeerId == in.UserId {
		maxId = 0
	}

	// inbox readed
	if dlg.ReadInboxMaxId >= maxId || dlg.UnreadCount == 0 {
		pts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
		ptsCount = 0
		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      pts,
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	maxInboxMsg, err3 := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, maxId)
	if err3 != nil {
		c.Logger.Errorf("messages.readHistory - error: not found dialog(%d,%d), error is %v", in.UserId, maxId, err3)
		return nil, mtproto.ErrInternalServerError
	} else if maxInboxMsg == nil {
		c.Logger.Errorf("messages.readHistory - error: not found dialog(%d,%d)", in.UserId, maxId)
		return nil, mtproto.ErrMsgIdInvalid
	}

	pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId)
	ptsCount = 1

	c.svcCtx.Dao.InboxClient.InboxReadInboxHistory(
		c.ctx,
		&inbox.TLInboxReadInboxHistory{
			UserId:         in.UserId,
			AuthKeyId:      in.AuthKeyId,
			PeerType:       in.PeerType,
			PeerId:         in.PeerId,
			Pts:            pts,
			PtsCount:       ptsCount,
			UnreadCount:    dlg.UnreadCount,
			ReadInboxMaxId: dlg.ReadInboxMaxId,
			MaxId:          maxId,
		})

	switch in.PeerType {
	case mtproto.PEER_USER:
		c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
			c.ctx,
			&inbox.TLInboxReadOutboxHistory{
				UserId:             maxInboxMsg.SenderUserId,
				PeerType:           in.PeerType,
				PeerId:             in.UserId,
				MaxDialogMessageId: maxInboxMsg.DialogMessageId,
			})
	case mtproto.PEER_CHAT:
		c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
			c.ctx,
			&inbox.TLInboxReadOutboxHistory{
				UserId:             maxInboxMsg.SenderUserId,
				PeerType:           in.PeerType,
				PeerId:             in.PeerId,
				MaxDialogMessageId: maxInboxMsg.DialogMessageId,
			})
	}

	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages(), nil
}
