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

// DialogUpdateReadOutbox
// dialog.updateReadOutbox user_id:long peer_type:int peer_id:long read_outbox_id:int = Bool;
func (c *DialogCore) DialogUpdateReadOutbox(in *dialog.TLDialogUpdateReadOutbox) (*mtproto.Bool, error) {
	c.DialogInsertOrUpdateDialog(&dialog.TLDialogInsertOrUpdateDialog{
		UserId:          in.UserId,
		PeerType:        in.PeerType,
		PeerId:          in.PeerId,
		TopMessage:      nil,
		ReadOutboxMaxId: &wrapperspb.Int32Value{Value: in.ReadOutboxId},
		ReadInboxMaxId:  nil,
		UnreadCount:     nil,
		UnreadMark:      false,
	})

	return mtproto.BoolTrue, nil
}
