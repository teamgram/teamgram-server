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

// DialogGetUserPinnedMessage
// dialog.getUserPinnedMessage user_id:long peer_type:int peer_id:long = Int32;
func (c *DialogCore) DialogGetUserPinnedMessage(in *dialog.TLDialogGetUserPinnedMessage) (*mtproto.Int32, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.getUserPinnedMessage - error: method DialogGetUserPinnedMessage not impl")

	return nil, mtproto.ErrMethodNotImpl
}
