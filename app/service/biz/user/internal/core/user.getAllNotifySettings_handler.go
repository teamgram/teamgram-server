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
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:int = Vector<PeerNotifySettings>;
func (c *UserCore) UserGetAllNotifySettings(in *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	// TODO(@benqi): GetAll settings?
	var (
		settings = &user.Vector_PeerPeerNotifySettings{
			Datas: []*user.PeerPeerNotifySettings{},
		}
	)

	if _, err := c.svcCtx.Dao.UserNotifySettingsDAO.SelectAllWithCB(c.ctx,
		in.UserId,
		func(i int, v *dataobject.UserNotifySettingsDO) {
			settings.Datas = append(settings.Datas, user.MakeTLPeerPeerNotifySettings(&user.PeerPeerNotifySettings{
				PeerType: v.PeerType,
				PeerId:   v.PeerId,
				Settings: makePeerNotifySettingsByDO(v),
			}).To_PeerPeerNotifySettings())
		}); err != nil {

		c.Logger.Errorf("user.getAllNotifySettings - error: %v", err)
		return nil, err
	}

	return settings, nil
}
