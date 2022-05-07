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
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/media/media"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

const (
	userDataKeyPrefix = "user_data"
)

func genUserDataCacheKey(id int64) string {
	return fmt.Sprintf("%s_%d", userDataKeyPrefix, id)
}

func (d *Dao) GetUserData(ctx context.Context, id int64) *user.UserData {
	userData := user.MakeTLUserData(nil).To_UserData()

	err := d.CachedConn.QueryRow(
		ctx,
		userData,
		genUserDataCacheKey(id),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do, err := d.UsersDAO.SelectById(ctx, id)
			if err != nil {
				return err
			}
			if do == nil {
				return sqlc.ErrNotFound
			}
			userData2 := v.(*user.UserData)
			d.setUserDataByDO(userData2, do)

			if do.IsBot && do.PhotoId != 0 {
				mr.FinishVoid(
					func() {
						userData2.Bot = d.getBotData(ctx, do.Id)
					},
					func() {
						userData2.ProfilePhoto, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
							PhotoId: do.PhotoId,
						})
					})
			} else {
				if do.IsBot {
					userData2.Bot = d.getBotData(ctx, do.Id)
				}
				if do.PhotoId != 0 {
					userData2.ProfilePhoto, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
						PhotoId: do.PhotoId,
					})
				}
			}

			if do.Restricted {
				jsonx.UnmarshalFromString(do.RestrictionReason, &userData2.RestrictionReason)
			}
			return nil
		})
	if err != nil {
		return nil
	}

	return userData
}

func (d *Dao) setUserDataByDO(userData *user.UserData, userDO *dataobject.UsersDO) {
	// deleted
	if userDO.Deleted {
		userData.Id = userDO.Id
		userData.AccessHash = userDO.AccessHash
		userData.Deleted = true
	} else {
		userData.Id = userDO.Id
		userData.AccessHash = userDO.AccessHash
		userData.UserType = userDO.UserType
		userData.SceretKeyId = userDO.SecretKeyId
		userData.FirstName = userDO.FirstName
		userData.LastName = userDO.LastName
		userData.Username = userDO.Username
		userData.Phone = userDO.Phone
		userData.ProfilePhoto = nil //
		userData.Bot = nil
		userData.CountryCode = userDO.CountryCode
		userData.Verified = userDO.Verified
		userData.Support = userDO.Support
		userData.Scam = userDO.Scam
		userData.Fake = userDO.Fake
		userData.About = mtproto.MakeFlagsString(userDO.About)
		userData.Restricted = userDO.Restricted
		userData.RestrictionReason = nil
		userData.ContactsVersion = 1
		userData.PrivaciesVersion = 1
		userData.Deleted = false
	}
}

func (d *Dao) getBotData(ctx context.Context, botId int64) *user.BotData {
	var (
		botData *user.BotData
	)

	botDO, _ := d.BotsDAO.Select(ctx, botId)
	if botDO != nil {
		// userData.Bot
		botData = user.MakeTLBotData(&user.BotData{
			Id:                   botDO.BotId,
			BotType:              botDO.BotType,
			Creator:              botDO.CreatorUserId,
			Token:                botDO.Token,
			Description:          botDO.Description,
			BotChatHistory:       botDO.BotChatHistory,
			BotNochats:           botDO.BotNochats,
			BotInlineGeo:         botDO.BotInlineGeo,
			BotInfoVersion:       botDO.BotInfoVersion,
			BotInlinePlaceholder: mtproto.MakeFlagsString(botDO.BotInlinePlaceholder),
		}).To_BotData()
	}

	return botData
}

func (d *Dao) UpdateUserFirstAndLastName(ctx context.Context, id int64, firstName, lastName string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"first_name": firstName,
				"last_name":  lastName,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genUserDataCacheKey(id))
	if err != nil {
		return false
	}

	return true
}

func (d *Dao) UpdateUserAbout(ctx context.Context, id int64, about string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"about": about,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genUserDataCacheKey(id))
	if err != nil {
		return false
	}

	return true
}

func (d *Dao) UpdateUserUsername(ctx context.Context, id int64, username string) bool {
	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err := d.UsersDAO.UpdateUser(ctx, map[string]interface{}{
				"username": username,
			}, id)

			if err != nil {
				return 0, 0, err
			}

			return 0, rowsAffected, nil
		},
		genUserDataCacheKey(id))
	if err != nil {
		return false
	}

	return true
}
