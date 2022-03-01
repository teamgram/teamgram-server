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
)

// DialogMarkDialogUnread
// dialog.markDialogUnread user_id:long peer_type:int peer_id:long unread_mark:Bool = Bool;
func (c *DialogCore) DialogMarkDialogUnread(in *dialog.TLDialogMarkDialogUnread) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.markDialogUnread - error: method DialogMarkDialogUnread not impl")

	return nil, mtproto.ErrMethodNotImpl
}
