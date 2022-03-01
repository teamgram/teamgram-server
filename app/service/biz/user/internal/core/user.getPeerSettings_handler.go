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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetPeerSettings
// user.getPeerSettings user_id:int peer_type:int peer_id:int = PeerSettings;
func (c *UserCore) UserGetPeerSettings(in *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error) {
	peerSettingsDO, err := c.svcCtx.Dao.UserPeerSettingsDAO.Select(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("user.getPeerSettings - error: %v", err)
		return nil, err
	}

	var (
		peerSettings *mtproto.PeerSettings
	)

	if peerSettingsDO != nil {
		peerSettings = &mtproto.PeerSettings{
			ReportSpam:            peerSettingsDO.ReportSpam,
			AddContact:            peerSettingsDO.AddContact,
			BlockContact:          peerSettingsDO.BlockContact,
			ShareContact:          peerSettingsDO.ShareContact,
			NeedContactsException: peerSettingsDO.NeedContactsException,
			ReportGeo:             peerSettingsDO.ReportGeo,
			Autoarchived:          peerSettingsDO.Autoarchived,
			GeoDistance:           nil,
		}

		if peerSettingsDO.GeoDistance != 0 {
			peerSettings.GeoDistance = &types.Int32Value{Value: peerSettingsDO.GeoDistance}
		}
	}

	return mtproto.MakeTLPeerSettings(peerSettings).To_PeerSettings(), nil
}
