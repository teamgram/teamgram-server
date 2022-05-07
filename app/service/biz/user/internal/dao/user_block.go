// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/teamgram/proto/mtproto"
)

var (
	contactsBlockPeerPrefix = "user_block_peer"
)

type CachedPeerBlocked struct {
	PeerBlocked *mtproto.PeerBlocked `json:"peer_blocked"`
}

func (c *CachedPeerBlocked) IsEmpty() bool {
	if c == nil {
		return true
	}

	return c.PeerBlocked == nil
}

func genContactsBlockPeerCacheKey(id, blockedId int64) string {
	return fmt.Sprintf("%s_%d_%d", contactsBlockPeerPrefix, id, blockedId)
}

func (d *Dao) CheckBlocked(ctx context.Context, id, blockedId int64) bool {
	var (
		blocked = new(CachedPeerBlocked)
	)
	d.CachedConn.QueryRow(
		ctx,
		blocked,
		genContactsBlockPeerCacheKey(id, blockedId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do, err := d.UserPeerBlocksDAO.Select(ctx, id, mtproto.PEER_USER, blockedId)
			if err != nil {
				return err
			}
			if do != nil {
				v.(*CachedPeerBlocked).PeerBlocked = mtproto.MakeTLPeerBlocked(&mtproto.PeerBlocked{
					PeerId: mtproto.MakePeerUser(do.PeerId),
					Date:   int32(do.Date),
				}).To_PeerBlocked()
			}

			return nil
		},
	)

	return !blocked.IsEmpty()
}

func (d *Dao) BlockUser(ctx context.Context, id, blockId int64) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			return d.UserPeerBlocksDAO.InsertOrUpdate(ctx, &dataobject.UserPeerBlocksDO{
				UserId:   id,
				PeerType: mtproto.PEER_USER,
				PeerId:   blockId,
				Date:     time.Now().Unix(),
			})
		},
		genContactsBlockPeerCacheKey(id, blockId))

	return err == nil
}

func (d *Dao) UnBlockUser(ctx context.Context, id, unblockId int64) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err := d.UserPeerBlocksDAO.Delete(
				ctx,
				id,
				mtproto.PEER_USER,
				unblockId)
			return 0, affected, err
		},
		genContactsBlockPeerCacheKey(id, unblockId))

	return err == nil
}
