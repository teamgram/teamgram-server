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
	"github.com/gogo/protobuf/types"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

func makePeerNotifySettingsByDO(do *dataobject.UserNotifySettingsDO) (settings *mtproto.PeerNotifySettings) {
	settings = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	if do.ShowPreviews != -1 {
		settings.ShowPreviews = mtproto.ToBool(do.ShowPreviews == 1)
	}
	if do.Silent != -1 {
		settings.Silent = mtproto.ToBool(do.Silent == 1)
	}
	if do.MuteUntil != -1 {
		settings.MuteUntil = &types.Int32Value{Value: do.MuteUntil}
	}
	if do.Sound != "-1" {
		settings.Sound = &types.StringValue{Value: do.Sound}
	}
	return
}

// UserGetNotifySettings
// user.getNotifySettings user_id:int peer_type:int peer_id:int = PeerNotifySettings;
func (c *UserCore) UserGetNotifySettings(in *user.TLUserGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	var (
		settings *mtproto.PeerNotifySettings
	)

	// miss or redis error
	do, err := c.svcCtx.Dao.UserNotifySettingsDAO.Select(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("user.getNotifySettings - error: %v", err)
		return nil, err
	}

	if do == nil {
		switch in.PeerType {
		case mtproto.PEER_USERS,
			mtproto.PEER_CHATS,
			mtproto.PEER_BROADCASTS:
			settings = mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{
				ShowPreviews: mtproto.BoolTrue,
				Silent:       mtproto.BoolFalse,
				MuteUntil:    &types.Int32Value{Value: 0},
				Sound:        &types.StringValue{Value: "default"},
			}).To_PeerNotifySettings()
		default:
			settings = mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{
				ShowPreviews: nil,
				Silent:       nil,
				MuteUntil:    nil,
				Sound:        nil,
			}).To_PeerNotifySettings()
		}
	} else {
		settings = makePeerNotifySettingsByDO(do)
	}

	return settings, nil
}
