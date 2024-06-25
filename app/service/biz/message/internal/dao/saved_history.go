// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/zeromicro/go-zero/core/logx"
)

// GetOffsetIdBackwardSavedHistoryMessages offset
func (d *Dao) GetOffsetIdBackwardSavedHistoryMessages(ctx context.Context, userId int64, savedPeerId *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch savedPeerId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		rList, _ := d.MessagesDAO.SelectBackwardSavedByOffsetIdLimitWithCB(
			ctx,
			userId,
			savedPeerId.PeerType,
			savedPeerId.PeerId,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
		// logx.WithContext(ctx).Infof("GetOffsetIdBackwardHistoryMessages: %v", rList)
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetIdForwardSavedHistoryMessages(ctx context.Context, userId int64, savedPeerId *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch savedPeerId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		rList, _ := d.MessagesDAO.SelectForwardSavedByOffsetIdLimitWithCB(
			ctx,
			userId,
			savedPeerId.PeerType,
			savedPeerId.PeerId,
			offsetId,
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateBackwardSavedHistoryMessages(ctx context.Context, userId int64, savedPeerId *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch savedPeerId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		rList, _ := d.MessagesDAO.SelectBackwardSavedByOffsetDateLimitWithCB(
			ctx,
			userId,
			savedPeerId.PeerType,
			savedPeerId.PeerId,
			int64(offsetDate),
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateForwardSavedHistoryMessages(ctx context.Context, userId int64, savedPeerId *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch savedPeerId.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		rList, _ := d.MessagesDAO.SelectForwardSavedByOffsetDateLimitWithCB(
			ctx,
			userId,
			savedPeerId.PeerType,
			savedPeerId.PeerId,
			int64(offsetDate),
			limit,
			func(sz, i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://teamgram.net required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}
