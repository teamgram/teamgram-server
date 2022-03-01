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

// DialogGetDialogsByOffsetDate
// dialog.getDialogsByOffsetDate user_id:long exclude_pinned:Bool offset_date:int limit:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetDialogsByOffsetDate(in *dialog.TLDialogGetDialogsByOffsetDate) (*dialog.Vector_DialogExt, error) {
	// TODO: not impl
	c.Logger.Errorf("dialog.getDialogsByOffsetDate - error: method DialogGetDialogsByOffsetDate not impl")

	return nil, mtproto.ErrMethodNotImpl
}
