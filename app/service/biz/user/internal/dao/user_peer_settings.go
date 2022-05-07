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

package dao

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
)

const (
	userPeerSettingsKeyPrefix = "user_peer_settings"
)

func genUserPeerSettingsCacheKey(id int64, peerType int32, peerId int64) string {
	return fmt.Sprintf("%s_%d_%d_%d", userPeerSettingsKeyPrefix, id, peerType, peerId)
}

func (d *Dao) GetUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64) (*mtproto.PeerSettings, error) {
	settings := mtproto.MakeTLPeerSettings(nil).To_PeerSettings()

	err := d.CachedConn.QueryRow(
		ctx,
		settings,
		genUserPeerSettingsCacheKey(id, peerType, peerId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			settingsDO, err := d.UserPeerSettingsDAO.Select(ctx, id, peerType, peerId)
			if err != nil {
				return err
			}

			var (
				peerSettings = v.(*mtproto.PeerSettings)
			)

			if settingsDO != nil {
				peerSettings.ReportSpam = settingsDO.ReportSpam
				peerSettings.AddContact = settingsDO.AddContact
				peerSettings.BlockContact = settingsDO.BlockContact
				peerSettings.ShareContact = settingsDO.ShareContact
				peerSettings.NeedContactsException = settingsDO.NeedContactsException
				peerSettings.ReportGeo = settingsDO.ReportGeo
				peerSettings.Autoarchived = settingsDO.Autoarchived
				peerSettings.GeoDistance = nil

				if settingsDO.GeoDistance != 0 {
					peerSettings.GeoDistance = &types.Int32Value{Value: settingsDO.GeoDistance}
				}
			} else {
				peerSettings.ReportSpam = false
				peerSettings.AddContact = false
				peerSettings.BlockContact = false
				peerSettings.ShareContact = false
				peerSettings.NeedContactsException = false
				peerSettings.ReportGeo = false
				peerSettings.Autoarchived = false
				peerSettings.GeoDistance = nil
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return settings, nil
}
