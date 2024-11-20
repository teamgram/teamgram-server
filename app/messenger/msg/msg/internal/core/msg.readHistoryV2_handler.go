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
	"context"
	"math"

	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
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
		c.Logger.Errorf("msg.readHistoryV2 - error: invalid peer %v", err)
		return nil, mtproto.ErrInternalServerError
	} else if dlg == nil {
		c.Logger.Errorf("msg.readHistoryV2 - error: not found dialog, request: %s", in)
		return nil, mtproto.ErrPeerIdInvalid
	}

	if maxId == 0 || maxId == math.MaxInt32 {
		maxId = dlg.TopMessage
	}

	if in.PeerType == mtproto.PEER_SELF || in.PeerType == mtproto.PEER_USER && in.PeerId == in.UserId {
		maxId = 0
	}

	// inbox readed
	// if dlg.ReadInboxMaxId >= maxId || dlg.UnreadCount == 0 {
	if dlg.ReadInboxMaxId >= maxId {
		c.Logger.Debugf("dlg: %#v", dlg)
		pts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
		ptsCount = 0
		_ = ptsCount

		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      pts,
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	if kUseV3 {
		return c.msgReadHistoryV3(c.ctx, in, maxId, dlg)
	} else {
		return c.msgReadHistoryV2(c.ctx, in, maxId, dlg)
	}
}

func (c *MsgCore) msgReadHistoryV2(ctx context.Context, in *msg.TLMsgReadHistoryV2, maxId int32, dlg *dataobject.DialogsDO) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		// maxId         = in.MaxId
	)

	maxInboxMsg, err3 := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(ctx, in.UserId, maxId)
	if err3 != nil {
		c.Logger.Errorf("msg.readHistoryV2 - error: not found dialog(%d,%d), error is %v", in.UserId, maxId, err3)
		return nil, mtproto.ErrInternalServerError
	} else if maxInboxMsg == nil {
		c.Logger.Errorf("msg.readHistoryV2 - error: not found dialog(%d,%d)", in.UserId, maxId)
		return nil, mtproto.ErrMsgIdInvalid
	}

	pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(ctx, in.UserId)
	ptsCount = 1

	_, _ = c.svcCtx.Dao.InboxClient.InboxReadInboxHistory(
		ctx,
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
		_, _ = c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
			ctx,
			&inbox.TLInboxReadOutboxHistory{
				UserId:             maxInboxMsg.SenderUserId,
				PeerType:           in.PeerType,
				PeerId:             in.UserId,
				MaxDialogMessageId: maxInboxMsg.DialogMessageId,
			})
	case mtproto.PEER_CHAT:
		_, _ = c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
			ctx,
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

func (c *MsgCore) msgReadHistoryV3(ctx context.Context, in *msg.TLMsgReadHistoryV2, maxId int32, dlg *dataobject.DialogsDO) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		// maxId         = in.MaxId
	)

	maxInboxMsg, err3 := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(ctx, in.UserId, maxId)
	if err3 != nil {
		c.Logger.Errorf("msg.readHistoryV2 - error: not found dialog(%d,%d), error is %v", in.UserId, maxId, err3)
		return nil, mtproto.ErrInternalServerError
	} else if maxInboxMsg == nil {
		c.Logger.Errorf("msg.readHistoryV2 - error: not found dialog(%d,%d)", in.UserId, maxId)
		return nil, mtproto.ErrMsgIdInvalid
	}

	return threading2.WrapperGoFunc(
		ctx,
		(*mtproto.Messages_AffectedMessages)(nil),
		func(ctx context.Context) {
			pts = 0
			ptsCount = 1

			_, _ = c.svcCtx.Dao.InboxClient.InboxReadInboxHistory(
				ctx,
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
					Layer:          mtproto.MakeFlagsInt32(c.MD.Layer),
					ServerId:       mtproto.MakeFlagsString(c.MD.ServerId),
					SessionId:      mtproto.MakeFlagsInt64(c.MD.SessionId),
					ClientReqMsgId: mtproto.MakeFlagsInt64(c.MD.ClientMsgId),
				})

			switch in.PeerType {
			case mtproto.PEER_USER:
				_, _ = c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
					ctx,
					&inbox.TLInboxReadOutboxHistory{
						UserId:             maxInboxMsg.SenderUserId,
						PeerType:           in.PeerType,
						PeerId:             in.UserId,
						MaxDialogMessageId: maxInboxMsg.DialogMessageId,
					})
			case mtproto.PEER_CHAT:
				_, _ = c.svcCtx.Dao.InboxClient.InboxReadOutboxHistory(
					ctx,
					&inbox.TLInboxReadOutboxHistory{
						UserId:             maxInboxMsg.SenderUserId,
						PeerType:           in.PeerType,
						PeerId:             in.PeerId,
						MaxDialogMessageId: maxInboxMsg.DialogMessageId,
					})
			}
		}).(*mtproto.Messages_AffectedMessages), mtproto.ErrPushRpcClient
}
