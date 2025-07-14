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
	userNotifySettingsKeyPrefix = "user_notify_settings"
)

func genUserNotifySettingsCacheKey(id int64, peerType int32, peerId int64) string {
	return fmt.Sprintf("%s_%d_%d_%d", userNotifySettingsKeyPrefix, id, peerType, peerId)
}

func (d *Dao) GetUserNotifySettings(ctx context.Context, id int64, peerType int32, peerId int64) (*tg.PeerNotifySettings, error) {
	settings := &tg.TLPeerNotifySettings{}

	err := d.CachedConn.QueryRow(
		ctx,
		settings,
		genUserNotifySettingsCacheKey(id, peerType, peerId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do, err := d.UserNotifySettingsDAO.Select(ctx, id, peerType, peerId)
			if err != nil {
				// c.Logger.Errorf("user.getNotifySettings - error: %v", err)
				return err
			}

			var (
				notifySettings = v.(*tg.TLPeerNotifySettings)
			)

			if do == nil {
				switch peerType {
				case tg.PEER_USERS,
					tg.PEER_CHATS,
					tg.PEER_BROADCASTS:

					notifySettings.ShowPreviews = tg.BoolTrueClazz
					notifySettings.Silent = tg.BoolFalseClazz
					notifySettings.MuteUntil = tg.MakeFlagsZeroInt32()
					notifySettings.IosSound = nil
				}
			} else {
				setPeerNotifySettingsByDO(notifySettings.ToPeerNotifySettings(), do)
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return settings.ToPeerNotifySettings(), nil
}

func setPeerNotifySettingsByDO(settings *tg.PeerNotifySettings, do *dataobject.UserNotifySettingsDO) {
	settings.Match(
		func(c *tg.TLPeerNotifySettings) interface{} {
			if do.ShowPreviews != -1 {
				c.ShowPreviews = tg.ToBoolClazz(do.ShowPreviews == 1)
			}
			if do.Silent != -1 {
				c.Silent = tg.ToBoolClazz(do.Silent == 1)
			}
			if do.MuteUntil != -1 {
				c.MuteUntil = tg.MakeFlagsInt32(do.MuteUntil)
			}
			if do.Sound != "-1" {
				// c.Sound = &wrapperspb.StringValue{Value: do.Sound}
			}

			return nil
		})
}

func makeDOByPeerNotifySettings(settings *tg.PeerNotifySettings) (doMap map[string]interface{}) {
	doMap = map[string]interface{}{}

	settings.Match(
		func(c *tg.TLPeerNotifySettings) interface{} {
			if c.ShowPreviews != nil {
				if tg.FromBoolClazz(c.ShowPreviews) {
					doMap["show_previews"] = 1
				} else {
					doMap["show_previews"] = 0
				}
			} else {
				doMap["show_previews"] = -1
			}

			if c.Silent != nil {
				if tg.FromBoolClazz(c.Silent) {
					doMap["silent"] = 1
				} else {
					doMap["silent"] = 0
				}
			} else {
				doMap["silent"] = -1
			}

			if c.MuteUntil != nil {
				doMap["mute_until"] = tg.GetFlagsInt32(c.MuteUntil)
			} else {
				doMap["mute_until"] = -1
			}

			//if c.Sound != nil {
			//	doMap["sound"] = settings.Sound.Value
			//} else {
			//	doMap["sound"] = "-1"
			//}

			return nil
		})

	return
}

func (d *Dao) SetUserPeerNotifySettings(ctx context.Context, id int64, peerType int32, peerId int64, settings *tg.PeerNotifySettings) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			cMap := makeDOByPeerNotifySettings(settings)
			return d.UserNotifySettingsDAO.InsertOrUpdateExt(ctx, id, peerType, peerId, cMap)
		},
		genUserNotifySettingsCacheKey(id, peerType, peerId))

	return err
}
