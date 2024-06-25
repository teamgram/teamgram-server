// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/mr"
)

// DialogGetMyDialogsData
// dialog.getMyDialogsData flags:# user:flags.0?true chat:flags.1?true channel:flags.2?true = Vector<PeerUtil>;
func (c *DialogCore) DialogGetMyDialogsData(in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	var (
		fns      []func() error
		uIdList  []int64
		cIdList  []int64
		chIdList []int64
	)

	if in.User {
		fns = append(fns, func() error {
			err2 := c.svcCtx.Dao.CachedConn.QueryRow(
				c.ctx,
				&uIdList,
				dialog.GenConversationsCacheKey(in.UserId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := c.svcCtx.Dao.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						in.UserId,
						[]int32{mtproto.PEER_USER},
						func(sz, i int, v *dataobject.DialogsDO) {
							idList = append(idList, v.PeerId)
						})
					if err2 != nil {
						// TODO: log
						return err2
					}

					*v.(*[]int64) = idList

					return nil
				})
			return err2
		})
	}

	if in.Chat {
		fns = append(fns, func() error {
			err2 := c.svcCtx.Dao.CachedConn.QueryRow(
				c.ctx,
				&cIdList,
				dialog.GenChatsCacheKey(in.UserId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := c.svcCtx.Dao.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						in.UserId,
						[]int32{mtproto.PEER_CHAT},
						func(sz, i int, v *dataobject.DialogsDO) {
							idList = append(idList, v.PeerId)
						})
					if err2 != nil {
						// TODO: log
						return err2
					}

					*v.(*[]int64) = idList

					return nil
				})
			return err2
		})
	}

	if in.Channel {
		fns = append(fns, func() error {
			err2 := c.svcCtx.Dao.CachedConn.QueryRow(
				c.ctx,
				&chIdList,
				dialog.GenChannelsCacheKey(in.UserId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := c.svcCtx.Dao.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						in.UserId,
						[]int32{mtproto.PEER_CHANNEL},
						func(sz, i int, v *dataobject.DialogsDO) {
							idList = append(idList, v.PeerId)
						})
					if err2 != nil {
						// TODO: log
						return err2
					}

					*v.(*[]int64) = idList

					return nil
				})
			return err2
		})
	}

	if err := mr.Finish(fns...); err != nil {
		c.Logger.Errorf("dialog.getMyDialogsData - error: %v", err)
		return nil, err
	}

	return dialog.MakeTLSimpleDialogsData(&dialog.DialogsData{
		Users:    uIdList,
		Chats:    cIdList,
		Channels: chIdList,
	}).To_DialogsData(), nil
}
