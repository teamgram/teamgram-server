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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// SyncPushUpdates
// sync.pushUpdates user_id:long updates:Updates = Void;
func (c *SyncCore) SyncPushUpdates(in *sync.TLSyncPushUpdates) (*mtproto.Void, error) {
	var (
		userId  = in.GetUserId()
		updates = in.GetUpdates()
	)

	notification, err := c.processUpdates(syncTypeUser, userId, false, updates)
	if err != nil {
		c.Logger.Errorf("sync.updatesNotMe - error: %v", err)
		return nil, err
	}

	c.pushUpdatesToSession(syncTypeUser, userId, 0, nil, nil, nil, updates, notification)

	return mtproto.EmptyVoid, nil
}
