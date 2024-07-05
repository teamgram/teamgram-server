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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/svc"
)

type DialogCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *DialogCore {
	return &DialogCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

// ////////////////////////////////////////////////////////////////////////////////////
func makeDialog(dialogDO *dataobject.DialogsDO) *dialog.DialogExt {
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

	order := dialogDO.Date2
	// pinned
	if dialogDO.FolderId == 0 {
		dialog2.Pinned = dialogDO.Pinned > 0
		if dialog2.Pinned {
			order = dialogDO.Pinned
		}
	} else {
		dialog2.Pinned = dialogDO.FolderPinned > 0
		if dialog2.Pinned {
			order = dialogDO.FolderPinned
		}
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

	return dialog.MakeTLDialogExt(&dialog.DialogExt{
		Order:          order,
		Dialog:         dialog2,
		AvailableMinId: 0,
		Date:           dialogDO.Date2,
		ThemeEmoticon:  dialogDO.ThemeEmoticon,
		TtlPeriod:      dialogDO.TtlPeriod,
	}).To_DialogExt()
}
