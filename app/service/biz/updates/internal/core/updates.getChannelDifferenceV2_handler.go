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
)

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (c *UpdatesCore) UpdatesGetChannelDifferenceV2(in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	var (
		rDiff = updates.MakeTLChannelDifference(&updates.ChannelDifference{
			Final:        false,
			Pts:          in.Pts,
			NewMessages:  nil,
			OtherUpdates: nil,
		}).To_ChannelDifference()
	)

	// TODO: check channelDifferenceTooLong
	rList, _ := c.svcCtx.Dao.ChannelPtsUpdatesDAO.SelectByGtPtsWithCB(
		c.ctx,
		in.ChannelId,
		in.Pts,
		in.Limit,
		func(i int, v *dataobject.ChannelPtsUpdatesDO) {
			var (
				update *mtproto.Update
			)

			if v.Pts > rDiff.Pts {
				rDiff.Pts = v.Pts
			}

			err := jsonx.UnmarshalFromString(v.UpdateData, &update)
			if err != nil {
				c.Logger.Errorf("unmarshal pts's update(%d)error: %v", v.Id, err)
				return
			} else if update == nil {
				c.Logger.Errorf("unmarshal pts's update(%d)error: update is nil", v.Id)
				return
			}

			updateType := mtproto.GetUpdateType(update)
			if updateType != v.UpdateType {
				c.Logger.Errorf("update data error.")
				return
			}

			switch v.UpdateType {
			case mtproto.PTS_UPDATE_NEW_CHANNEL_MESSAGE:
				newMessage := update.GetMessage_MESSAGE()

				// TODO: add sender_id in channel_pts_updates
				if newMessage.FromId == nil {
					newMessage.Out = false
				} else {
					from := mtproto.FromPeer(newMessage.GetFromId())
					if from.GetPeerId() == in.UserId {
						newMessage.Out = true
					} else {
						newMessage.Out = false
					}
				}
				rDiff.NewMessages = append(rDiff.NewMessages, newMessage)
			default:
				//	err := jsonx.UnmarshalFromString(v.UpdateData, update)
				//	if err != nil {
				//		c.Logger.Errorf("unmarshal pts's update(%d)error: %v", v.Id, err)
				//		return
				//	}
				//
				//	updateType := mtproto.GetUpdateType(update)
				//	if updateType != v.UpdateType {
				//		c.Logger.Errorf("update data error.")
				//		return
				//	}
				rDiff.OtherUpdates = append(rDiff.OtherUpdates, update)
			}
		})

	if len(rList) < int(in.Limit) {
		rDiff.Final = true
	}

	return rDiff, nil
}
