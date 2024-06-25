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

	"google.golang.org/protobuf/types/known/wrapperspb"
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
		settings.MuteUntil = &wrapperspb.Int32Value{Value: do.MuteUntil}
	}
	if do.Sound != "-1" {
		settings.Sound = &wrapperspb.StringValue{Value: do.Sound}
	}
	return
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:int = Vector<PeerNotifySettings>;
func (c *UserCore) UserGetAllNotifySettings(in *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	var (
		settings = &user.Vector_PeerPeerNotifySettings{
			Datas: []*user.PeerPeerNotifySettings{},
		}
	)

	if _, err := c.svcCtx.Dao.UserNotifySettingsDAO.SelectAllWithCB(c.ctx,
		in.UserId,
		func(sz, i int, v *dataobject.UserNotifySettingsDO) {
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
