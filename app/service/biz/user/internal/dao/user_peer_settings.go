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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/dal/dataobject"
)

const (
	userPeerSettingsKeyPrefix = "user_peer_settings"
)

func genUserPeerSettingsCacheKey(id int64, peerType int32, peerId int64) string {
	return fmt.Sprintf("%s_%d_%d_%d", userPeerSettingsKeyPrefix, id, peerType, peerId)
}

func (d *Dao) GetUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64) (*tg.PeerSettings, error) {
	settings := &tg.TLPeerSettings{
		ReportSpam:             false,
		AddContact:             false,
		BlockContact:           false,
		ShareContact:           false,
		NeedContactsException:  false,
		ReportGeo:              false,
		Autoarchived:           false,
		InviteMembers:          false,
		RequestChatBroadcast:   false,
		BusinessBotPaused:      false,
		BusinessBotCanReply:    false,
		GeoDistance:            nil,
		RequestChatTitle:       nil,
		RequestChatDate:        nil,
		BusinessBotId:          nil,
		BusinessBotManageUrl:   nil,
		ChargePaidMessageStars: nil,
		RegistrationMonth:      nil,
		PhoneCountry:           nil,
		NameChangeDate:         nil,
		PhotoChangeDate:        nil,
	}

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
				peerSettings = v.(*tg.TLPeerSettings)
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
					peerSettings.GeoDistance = tg.MakeFlagsInt32(settingsDO.GeoDistance)
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

	return settings.ToPeerSettings(), nil
}

func (d *Dao) SetUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64, settings *tg.PeerSettings) error {
	settings2, _ := settings.ToPeerSettings()
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			geoDistance := int32(0)
			if settings2.GeoDistance != nil {
				geoDistance = *settings2.GeoDistance
			}
			return d.UserPeerSettingsDAO.InsertOrUpdate(
				ctx,
				&dataobject.UserPeerSettingsDO{
					UserId:                id,
					PeerType:              peerType,
					PeerId:                peerId,
					Hide:                  false,
					ReportSpam:            settings2.ReportSpam,
					AddContact:            settings2.AddContact,
					BlockContact:          settings2.BlockContact,
					ShareContact:          settings2.ShareContact,
					NeedContactsException: settings2.NeedContactsException,
					ReportGeo:             settings2.ReportGeo,
					Autoarchived:          settings2.Autoarchived,
					GeoDistance:           geoDistance,
				})
		},
		genUserPeerSettingsCacheKey(id, peerType, peerId))

	return err
}

func (d *Dao) DeleteUserPeerSettings(ctx context.Context, id int64, peerType int32, peerId int64) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err := d.UserPeerSettingsDAO.Delete(ctx, id, peerType, peerId)
			return 0, affected, err
		},
		genUserPeerSettingsCacheKey(id, peerType, peerId))

	return err
}
