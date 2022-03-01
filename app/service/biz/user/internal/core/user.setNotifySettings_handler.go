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

func makeDOByPeerNotifySettings(settings *mtproto.PeerNotifySettings) (doMap map[string]interface{}) {
	doMap = map[string]interface{}{}

	if settings.ShowPreviews != nil {
		if mtproto.FromBool(settings.ShowPreviews) {
			doMap["show_previews"] = 1
		} else {
			doMap["show_previews"] = 0
		}
	} else {
		doMap["show_previews"] = -1
	}

	if settings.Silent != nil {
		if mtproto.FromBool(settings.Silent) {
			doMap["silent"] = 1
		} else {
			doMap["silent"] = 0
		}
	} else {
		doMap["silent"] = -1
	}

	if settings.MuteUntil != nil {
		doMap["mute_until"] = settings.MuteUntil.Value
	} else {
		doMap["mute_until"] = -1
	}

	if settings.Sound != nil {
		doMap["sound"] = settings.Sound.Value
	} else {
		doMap["sound"] = "-1"
	}

	return
}

// UserSetNotifySettings
// user.setNotifySettings user_id:int peer_type:int peer_id:int settings:PeerNotifySettings = Bool;
func (c *UserCore) UserSetNotifySettings(in *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	cMap := makeDOByPeerNotifySettings(in.Settings)
	if _, _, err := c.svcCtx.Dao.UserNotifySettingsDAO.InsertOrUpdateExt(c.ctx, in.UserId, in.PeerType, in.PeerId, cMap); err != nil {
		c.Logger.Errorf("user.setNotifySettings - error: %v", err)
		return nil, err
	}

	// putCache
	// m.Dao.Redis.SetPeerNotifySettings(ctx, userId, peer, settings)
	return mtproto.BoolTrue, nil
}
