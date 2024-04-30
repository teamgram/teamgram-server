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

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// DialogUpdateReadInbox
// dialog.updateReadInbox user_id:long peer_type:int peer_id:long read_inbox_id:int = Bool;
func (c *DialogCore) DialogUpdateReadInbox(in *dialog.TLDialogUpdateReadInbox) (*mtproto.Bool, error) {
	c.DialogInsertOrUpdateDialog(&dialog.TLDialogInsertOrUpdateDialog{
		UserId:          in.UserId,
		PeerType:        in.PeerType,
		PeerId:          in.PeerId,
		TopMessage:      nil,
		ReadOutboxMaxId: nil,
		ReadInboxMaxId:  &wrapperspb.Int32Value{Value: in.ReadInboxId},
		UnreadCount:     nil,
		UnreadMark:      false,
	})

	return mtproto.BoolTrue, nil
}
