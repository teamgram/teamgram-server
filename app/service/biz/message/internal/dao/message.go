// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) MakeMessageBox(ctx context.Context, selfUserId int64, do *dataobject.MessagesDO) (box *mtproto.MessageBox) {
	box = mtproto.MakeTLMessageBox(&mtproto.MessageBox{
		UserId:            do.UserId,
		MessageId:         do.UserMessageBoxId,
		SenderUserId:      do.SenderUserId,
		PeerType:          do.PeerType,
		PeerId:            do.PeerId,
		RandomId:          do.RandomId,
		DialogId1:         do.DialogId1,
		DialogId2:         do.DialogId2,
		DialogMessageId:   do.DialogMessageId,
		MessageFilterType: do.MessageFilterType,
		Message:           nil,
		Mentioned:         do.Mentioned,
		MediaUnread:       do.MediaUnread,
		Pinned:            do.Pinned, // TODO
		Reaction:          do.Reaction,
	}).To_MessageBox()
	// Message
	_ = jsonx.UnmarshalFromString(do.MessageData, &box.Message)
	box.Message = box.Message.FixData()

	// poll
	if d.Plugin != nil {
		pollId, _ := mtproto.GetPollIdByMessage(box.Message.Media)
		if pollId != 0 {
			box.Message.Media, _ = d.Plugin.GetMessageMediaPoll(ctx, selfUserId, pollId)
		}
	}

	return
}
