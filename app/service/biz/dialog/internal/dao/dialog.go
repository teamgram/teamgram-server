// Copyright 2023 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func (d *Dao) MakeDialog(dialogDO *dataobject.DialogsDO) *dialog.DialogExt {
	dialog2 := mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:               false,
		UnreadMark:           dialogDO.UnreadMark, // TODO(@benqi)
		ViewForumAsMessages:  false,
		Peer:                 mtproto.MakePeer(dialogDO.PeerType, dialogDO.PeerId),
		TopMessage:           dialogDO.TopMessage,
		ReadInboxMaxId:       dialogDO.ReadInboxMaxId,
		ReadOutboxMaxId:      dialogDO.ReadOutboxMaxId,
		UnreadCount:          dialogDO.UnreadCount,
		UnreadMentionsCount:  dialogDO.UnreadMentionsCount,
		UnreadReactionsCount: dialogDO.UnreadReactionsCount,
		NotifySettings:       nil,
		Pts:                  nil,
		Draft:                nil,
		FolderId:             mtproto.MakeFlagsInt32(dialogDO.FolderId),
		TtlPeriod:            nil,
	}).To_Dialog()

	// fix unreadCount
	if dialog2.UnreadMentionsCount < 0 {
		dialog2.UnreadMentionsCount = 0
	}
	if dialogDO.UnreadReactionsCount < 0 {
		dialog2.UnreadReactionsCount = 0
	}

	// pinned
	if dialogDO.FolderId == 0 {
		dialog2.Pinned = dialogDO.Pinned > 0
	} else {
		dialog2.Pinned = dialogDO.FolderPinned > 0
	}

	// draft message.
	if dialogDO.DraftType == 2 {
		draft := &mtproto.DraftMessage{}
		err := jsonx.UnmarshalFromString(dialogDO.DraftMessageData, &draft)
		if err == nil {
			dialog2.Draft = draft
		} else {
			dialog2.Draft = mtproto.MakeTLDraftMessageEmpty(draft).To_DraftMessage()
		}
	} else if dialogDO.DraftType == 1 {
		dialog2.Draft = mtproto.MakeTLDraftMessageEmpty(nil).To_DraftMessage()
	}

	// NotifySettings
	dialog2.NotifySettings = mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{
		//
	}).To_PeerNotifySettings()

	return &dialog.DialogExt{
		Order:          dialogDO.Date2,
		Dialog:         dialog2,
		AvailableMinId: 0,
		Date:           dialogDO.Date2,
	}
}

func (d *Dao) GetDialog(ctx context.Context, userId int64, peerType int32, peerId int64) (*dialog.DialogExt, error) {
	return d.GetDialogByPeerDialogId(ctx, userId, mtproto.MakePeerDialogId(peerType, peerId))
}

func (d *Dao) GetDialogByPeerDialogId(ctx context.Context, userId, peerDialogId int64) (*dialog.DialogExt, error) {
	var (
		dlgExt *dialog.DialogExt
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&dlgExt,
		dialog.GenDialogCacheKey(userId, peerDialogId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			dialogDO, err := d.DialogsDAO.SelectByPeerDialogId(ctx, userId, peerDialogId)
			if err != nil {
				return err
			} else if dialogDO == nil {
				err = sqlc.ErrNotFound
				return err
			}

			*v.(**dialog.DialogExt) = d.MakeDialog(dialogDO)
			return nil
		})

	if err != nil {
		logx.WithContext(ctx).Errorf("dialog.getDialogById - error: %v", err)
		if errors.Is(err, sqlc.ErrNotFound) {
			err = mtproto.ErrPeerIdInvalid
		}
		return nil, err
	}

	return dlgExt, nil
}
