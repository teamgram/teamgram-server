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

// SyncPushUpdatesIfNot
// sync.pushUpdatesIfNot user_id:long excludes:Vector<int64> updates:Updates = Void;
func (c *SyncCore) SyncPushUpdatesIfNot(in *sync.TLSyncPushUpdatesIfNot) (*mtproto.Void, error) {
	// TODO: not impl
	c.Logger.Errorf("sync.pushUpdatesIfNot - error: method SyncPushUpdatesIfNot not impl")

	return nil, mtproto.ErrMethodNotImpl
}
