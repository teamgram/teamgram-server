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
	"github.com/teamgram/teamgram-server/app/service/biz/updates/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/updates/updates"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
)

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (c *UpdatesCore) UpdatesGetDifferenceV2(in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	limit := in.GetPtsTotalLimit().GetValue()
	// check pts_total_limit
	if limit <= 0 || limit > 5000 {
		limit = 5000
	}

	updateList, lastPts, lastSeq, lastDate := c.GetMergedUpdateList(in.UserId, in.AuthKeyId, in.Pts, int32(in.Date), limit)
	if lastDate == 0 {
		lastDate = time.Now().Unix()
	}

	if len(updateList) == 0 {
		// 1. updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
		if lastSeq == 0 {
			lastSeq = c.svcCtx.Dao.IDGenClient2.CurrentSeqId(c.ctx, in.AuthKeyId)
		}

		return updates.MakeTLDifferenceEmpty(&updates.Difference{
			State: mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
				Pts:         lastPts,
				Date:        int32(lastDate),
				UnreadCount: 0,
				Seq:         lastSeq,
			}).To_Updates_State(),
		}).To_Difference(), nil
	} else if len(updateList) >= int(limit) {
		// 2. updates.difference#f49ca0
		//	new_messages:Vector<Message>
		//	new_encrypted_messages:Vector<EncryptedMessage>
		//	other_updates:Vector<Update>
		//	chats:Vector<Chat>
		//	users:Vector<User>
		//	state:updates.State = updates.Difference;
		// 2. 1 updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
		//

		var (
			totalUpdatesCount = int32(len(updateList))
			maxPts            = lastPts
		)
		if lastPts > in.Pts {
			maxPts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
			totalUpdatesCount += maxPts - lastPts
		}

		if lastSeq > 0 {
			totalUpdatesCount += c.svcCtx.Dao.IDGenClient2.CurrentSeqId(c.ctx, in.AuthKeyId) - lastSeq
		}

		//if totalUpdatesCount >= 1000 {
		//	return updates.MakeTLDifferenceTooLong(&updates.Difference{
		//		Pts: maxPts,
		//	}).To_Difference(), nil
		//}
	}

	merge := newMergeUpdatesHelper()
	for _, update := range updateList {
		merge.merge(update, lastPts)
	}

	if lastSeq == 0 {
		lastSeq = c.svcCtx.Dao.IDGenClient2.CurrentSeqId(c.ctx, in.AuthKeyId)
	}
	state := mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
		Pts:         lastPts,
		Date:        int32(lastDate),
		UnreadCount: 0,
		Seq:         lastSeq,
	}).To_Updates_State()

	if len(updateList) >= int(limit) {
		// 2.2 updates.differenceSlice#a8fb1981
		//	new_messages:Vector<Message>
		//	new_encrypted_messages:Vector<EncryptedMessage>
		//	other_updates:Vector<Update>
		//	chats:Vector<Chat>
		//	users:Vector<User>
		//	intermediate_state:updates.State = updates.Difference;
		//
		return updates.MakeTLDifferenceSlice(&updates.Difference{
			NewMessages:       merge.newMessages,
			OtherUpdates:      merge.toUpdates(),
			IntermediateState: state,
		}).To_Difference(), nil
	} else {
		// 2. updates.difference#f49ca0
		//	new_messages:Vector<Message>
		//	new_encrypted_messages:Vector<EncryptedMessage>
		//	other_updates:Vector<Update>
		//	chats:Vector<Chat>
		//	users:Vector<User>
		//	state:updates.State = updates.Difference;
		return updates.MakeTLDifference(&updates.Difference{
			NewMessages:  merge.newMessages,
			OtherUpdates: merge.toUpdates(),
			State:        state,
		}).To_Difference(), nil
	}

}

func (c *UpdatesCore) addPtsUpdate(updates []*mtproto.Update, do *dataobject.UserPtsUpdatesDO) []*mtproto.Update {
	update := &mtproto.Update{}
	err := jsonx.UnmarshalFromString(do.UpdateData, update)
	if err != nil {
		c.Logger.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
		return updates
	}
	if mtproto.GetUpdateType(update) != do.UpdateType {
		c.Logger.Errorf("update data error.")
		return updates
	}
	updates = append(updates, update.FixData())
	return updates
}

func (c *UpdatesCore) addSeqUpdate(updates []*mtproto.Update, do *dataobject.AuthSeqUpdatesDO) []*mtproto.Update {
	update := &mtproto.Update{}
	err := jsonx.UnmarshalFromString(do.UpdateData, update)
	if err != nil {
		c.Logger.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
		return updates
	}
	updates = append(updates, update.FixData())
	return updates
}

func (c *UpdatesCore) GetMergedUpdateList(userId int64, authId int64, pts, date, limit int32) ([]*mtproto.Update, int32, int32, int64) {
	// ptsDOList, _ := m.ChannelPtsUpdatesDAO.SelectByGtPts(ctx, channelId, pts)
	a, _ := c.svcCtx.Dao.UserPtsUpdatesDAO.SelectByGtPts(c.ctx, userId, pts, limit)
	b, _ := c.svcCtx.Dao.AuthSeqUpdatesDAO.SelectByGtDate(c.ctx, authId, userId, int64(date), limit)

	//创建结果集数组
	var (
		lastDate int64 = 0
		lastSeq  int32 = 0
		c2             = make([]*mtproto.Update, 0, len(a)+len(b))
	)

	alen := len(a)
	blen := len(b)

	//默认数组从0开始
	aIndex := 0
	bIndex := 0
	cIndex := 0

	//开始遍历数组
	for aIndex < alen && bIndex < blen {
		if a[aIndex].Date2 < b[bIndex].Date2 {
			c2 = c.addPtsUpdate(c2, &a[aIndex])
			if a[aIndex].Pts > pts {
				pts = a[aIndex].Pts
			}
			// if a[aIndex].Date2 > lastDate {
			// 	lastDate = a[aIndex].Date2
			// }
			cIndex++
			aIndex++
		} else {
			c2 = c.addSeqUpdate(c2, &b[bIndex])
			if b[bIndex].Seq > lastSeq {
				lastSeq = b[bIndex].Seq
			}
			if b[bIndex].Date2 > lastDate {
				lastDate = b[bIndex].Date2
			}
			cIndex++
			bIndex++
		}
	}
	//当其中一个数组遍历完,另一个数组还没遍历完应.当将另一个数组遍历完
	for aIndex < alen {
		c2 = c.addPtsUpdate(c2, &a[aIndex])
		if a[aIndex].Pts > pts {
			pts = a[aIndex].Pts
		}
		// if a[aIndex].Date2 > lastDate {
		// 	lastDate = a[aIndex].Date2
		// }
		cIndex++
		aIndex++
	}
	for bIndex < blen {
		c2 = c.addSeqUpdate(c2, &b[bIndex])
		if b[bIndex].Seq > lastSeq {
			lastSeq = b[bIndex].Seq
		}
		if b[bIndex].Date2 > lastDate {
			lastDate = b[bIndex].Date2
		}
		cIndex++
		bIndex++
	}

	for i := 0; i < len(b); i++ {
		if b[i].Date2 > lastDate {
			lastDate = b[i].Date2
		}
	}

	return c2, pts, lastSeq, lastDate
}
