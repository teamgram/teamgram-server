// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) AddSeqToUpdatesQueue(ctx context.Context, authId, userId int64, updateType int32, updateData []byte) int32 {
	seq := int32(d.NextSeqId(ctx, authId))
	do := &dataobject.AuthSeqUpdatesDO{
		AuthId:     authId,
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: string(updateData),
		Date2:      time.Now().Unix(),
		Seq:        seq,
	}

	i, _, _ := d.AuthSeqUpdatesDAO.Insert(ctx, do)
	return int32(i)
}

func (d *Dao) AddToPtsQueue(ctx context.Context, userId int64, pts, ptsCount int32, update *mtproto.Update) int32 {
	// TODO(@benqi): check error
	updateData, _ := jsonx.Marshal(update)

	do := &dataobject.UserPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: mtproto.GetUpdateType(update),
		UpdateData: string(updateData),
		Date2:      time.Now().Unix(),
	}

	i, _, _ := d.UserPtsUpdatesDAO.Insert(ctx, do)
	return int32(i)
}
