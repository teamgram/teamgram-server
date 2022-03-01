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

//////////////////////////////////////////////////////////////////////////////////////
func makeDialog(dialogDO *dataobject.DialogsDO) *dialog.DialogExt {
	dialog2 := mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:              false,
		UnreadMark:          dialogDO.UnreadMark, // TODO(@benqi)
		Peer:                mtproto.MakePeer(dialogDO.PeerType, dialogDO.PeerId),
		TopMessage:          dialogDO.TopMessage,
		ReadInboxMaxId:      dialogDO.ReadInboxMaxId,
		ReadOutboxMaxId:     dialogDO.ReadOutboxMaxId,
		UnreadCount:         dialogDO.UnreadCount,
		UnreadMentionsCount: 0,
		NotifySettings:      nil,
		Pts:                 nil,
		Draft:               nil,
		FolderId:            mtproto.MakeFlagsInt32(dialogDO.FolderId),
	}).To_Dialog()

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
