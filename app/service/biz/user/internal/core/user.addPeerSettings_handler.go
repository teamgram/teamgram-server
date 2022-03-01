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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserAddPeerSettings
// user.addPeerSettings user_id:int peer_type:int peer_id:int settings:PeerSettings = Bool;
func (c *UserCore) UserAddPeerSettings(in *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	settings := in.GetSettings()

	_, _, err := c.svcCtx.Dao.UserPeerSettingsDAO.InsertIgnore(c.ctx, &dataobject.UserPeerSettingsDO{
		UserId:                in.UserId,
		PeerType:              in.PeerType,
		PeerId:                in.PeerId,
		Hide:                  false,
		ReportSpam:            settings.ReportSpam,
		AddContact:            settings.AddContact,
		BlockContact:          settings.BlockContact,
		ShareContact:          settings.ShareContact,
		NeedContactsException: settings.NeedContactsException,
		ReportGeo:             settings.ReportGeo,
		Autoarchived:          settings.Autoarchived,
		GeoDistance:           settings.GetGeoDistance().GetValue(),
	})
	if err != nil {
		c.Logger.Errorf("user.addPeerSettings - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
