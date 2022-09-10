// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerNotifySettings>;
func (c *UserCore) UserGetNotifySettingsList(in *user.TLUserGetNotifySettingsList) (*user.Vector_PeerPeerNotifySettings, error) {
	var (
		settings = &user.Vector_PeerPeerNotifySettings{
			Datas: []*user.PeerPeerNotifySettings{},
		}
	)

	if _, err := c.svcCtx.Dao.UserNotifySettingsDAO.SelectListWithCB(c.ctx,
		in.UserId,
		in.Peers,
		func(i int, v *dataobject.UserNotifySettingsDO) {
			settings.Datas = append(settings.Datas, user.MakeTLPeerPeerNotifySettings(&user.PeerPeerNotifySettings{
				PeerType: v.PeerType,
				PeerId:   v.PeerId,
				Settings: makePeerNotifySettingsByDO(v),
			}).To_PeerPeerNotifySettings())
		}); err != nil {

		c.Logger.Errorf("user.getNotifySettingsList - error: %v", err)
		return nil, err
	}

	return settings, nil
}
