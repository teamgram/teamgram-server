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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

//func makePeerNotifySettingsByDO(do *dataobject.UserNotifySettingsDO) (settings *mtproto.PeerNotifySettings) {
//	settings = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
//	if do.ShowPreviews != -1 {
//		settings.ShowPreviews = mtproto.ToBool(do.ShowPreviews == 1)
//	}
//	if do.Silent != -1 {
//		settings.Silent = mtproto.ToBool(do.Silent == 1)
//	}
//	if do.MuteUntil != -1 {
//		settings.MuteUntil = &types.Int32Value{Value: do.MuteUntil}
//	}
//	if do.Sound != "-1" {
//		settings.Sound = &types.StringValue{Value: do.Sound}
//	}
//	return
//}

// UserGetNotifySettings
// user.getNotifySettings user_id:int peer_type:int peer_id:int = PeerNotifySettings;
func (c *UserCore) UserGetNotifySettings(in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	settings, err := c.svcCtx.Dao.GetUserNotifySettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetUserId())

	if err != nil {
		c.Logger.Errorf("user.getNotifySettings - error: %v", err)
		return nil, err
	}

	return settings, nil
}
