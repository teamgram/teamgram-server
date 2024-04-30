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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	userNotifySettingsKeyPrefix = "user_notify_settings"
)

func genUserNotifySettingsCacheKey(id int64, peerType int32, peerId int64) string {
	return fmt.Sprintf("%s_%d_%d_%d", userNotifySettingsKeyPrefix, id, peerType, peerId)
}

func (d *Dao) GetUserNotifySettings(ctx context.Context, id int64, peerType int32, peerId int64) (*mtproto.PeerNotifySettings, error) {
	settings := mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()

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
				notifySettings = v.(*mtproto.PeerNotifySettings)
			)

			if do == nil {
				switch peerType {
				case mtproto.PEER_USERS,
					mtproto.PEER_CHATS,
					mtproto.PEER_BROADCASTS:

					notifySettings.ShowPreviews = mtproto.BoolTrue
					notifySettings.Silent = mtproto.BoolFalse
					notifySettings.MuteUntil = &wrapperspb.Int32Value{Value: 0}
					notifySettings.Sound = &wrapperspb.StringValue{Value: "default"}
				}
			} else {
				setPeerNotifySettingsByDO(notifySettings, do)
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func setPeerNotifySettingsByDO(settings *mtproto.PeerNotifySettings, do *dataobject.UserNotifySettingsDO) {
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
}

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

func (d *Dao) SetUserPeerNotifySettings(ctx context.Context, id int64, peerType int32, peerId int64, settings *mtproto.PeerNotifySettings) error {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			cMap := makeDOByPeerNotifySettings(settings)
			return d.UserNotifySettingsDAO.InsertOrUpdateExt(ctx, id, peerType, peerId, cMap)
		},
		genUserNotifySettingsCacheKey(id, peerType, peerId))

	return err
}
