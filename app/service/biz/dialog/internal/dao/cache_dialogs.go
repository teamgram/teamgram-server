// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/mr"
)

func (d *Dao) GetMyDialogsData(ctx context.Context, userId int64, conversation, chat, channel bool) (dialogs *dialog.DialogsData, err error) {
	var (
		fns      []func() error
		uIdList  []int64
		cIdList  []int64
		chIdList []int64
	)

	if conversation {
		fns = append(fns, func() error {
			err2 := d.CachedConn.QueryRow(
				ctx,
				&uIdList,
				dialog.GenConversationsCacheKey(userId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := d.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						userId,
						[]int32{mtproto.PEER_USER},
						func(i int, v *dataobject.DialogsDO) {
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

	if chat {
		fns = append(fns, func() error {
			err2 := d.CachedConn.QueryRow(
				ctx,
				&cIdList,
				dialog.GenChatsCacheKey(userId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := d.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						userId,
						[]int32{mtproto.PEER_CHAT},
						func(i int, v *dataobject.DialogsDO) {
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

	if channel {
		fns = append(fns, func() error {
			err2 := d.CachedConn.QueryRow(
				ctx,
				&chIdList,
				dialog.GenConversationsCacheKey(userId),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					var (
						idList []int64
					)
					_, err2 := d.DialogsDAO.SelectDialogsByPeerTypeWithCB(
						ctx,
						userId,
						[]int32{mtproto.PEER_CHANNEL},
						func(i int, v *dataobject.DialogsDO) {
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

	err = mr.Finish(fns...)
	if err != nil {
		return
	}

	dialogs = dialog.MakeTLSimpleDialogsData(&dialog.DialogsData{
		Users:    uIdList,
		Chats:    cIdList,
		Channels: chIdList,
	}).To_DialogsData()

	return
}
