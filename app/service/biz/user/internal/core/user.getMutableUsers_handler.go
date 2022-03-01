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
	"encoding/json"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// UserGetMutableUsers
// user.getMutableUsers id:Vector<int> = Vector<ImmutableUser>;
func (c *UserCore) UserGetMutableUsers(in *user.TLUserGetMutableUsers) (*user.Vector_ImmutableUser, error) {
	vUser := &user.Vector_ImmutableUser{
		Datas: make([]*user.ImmutableUser, 0, len(in.Id)),
	}

	if len(in.Id) == 0 {
		return vUser, nil
	}

	doList, _ := c.svcCtx.Dao.UsersDAO.SelectUsersByIdListWithCB(c.ctx, in.Id, func(i int, v *dataobject.UsersDO) {
		imUser := user.MakeTLImmutableUser(&user.ImmutableUser{
			User:             nil,
			LastSeenAt:       0,
			Contacts:         nil,
			KeysPrivacyRules: nil,
		}).To_ImmutableUser()

		if v.Deleted {
			// deleted user
			imUser.User = user.MakeTLUserData(&user.UserData{
				Id:         v.Id,
				AccessHash: v.AccessHash,
				Deleted:    true,
			}).To_UserData()
		} else {
			imUser.User = user.MakeTLUserData(&user.UserData{
				Id:                v.Id,
				AccessHash:        v.AccessHash,
				UserType:          v.UserType,
				SceretKeyId:       v.SecretKeyId,
				FirstName:         v.FirstName,
				LastName:          v.LastName,
				Username:          v.Username,
				Phone:             v.Phone,
				ProfilePhoto:      nil, //
				Bot:               nil,
				CountryCode:       v.CountryCode,
				Verified:          v.Verified,
				Support:           v.Support,
				Scam:              v.Scam,
				Fake:              v.Fake,
				About:             mtproto.MakeFlagsString(v.About),
				Restricted:        v.Restricted,
				RestrictionReason: nil,
				ContactsVersion:   1,
				PrivaciesVersion:  1,
				Deleted:           false,
			}).To_UserData()

			//// 1. bot
			if v.IsBot {
				botDO, _ := c.svcCtx.Dao.BotsDAO.Select(c.ctx, v.Id)
				if botDO != nil {
					imUser.User.Bot = user.MakeTLBotData(&user.BotData{
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
			}

			// 2. ProfilePhoto
			// TODO: ProfilePhoto
			// (userDO.RestrictionReason)
			imUser.User.ProfilePhoto, _ = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &media.TLMediaGetPhoto{
				PhotoId: v.PhotoId,
			})

			// 3. RestrictionReason
			// TODO: RestrictionReason
			if imUser.User.Restricted {
				json.Unmarshal([]byte(v.RestrictionReason), &imUser.User.RestrictionReason)
			}

			// 4. LastSeenAt
			lastSeenDO, _ := c.svcCtx.Dao.UserPresencesDAO.Select(c.ctx, v.Id)
			if lastSeenDO != nil {
				imUser.LastSeenAt = lastSeenDO.LastSeenAt
			}

			//
			if int(v.UserType) == user.UserTypeRegular {
				c.svcCtx.Dao.UserContactsDAO.SelectUserContactsWithCB(c.ctx, v.Id, func(i int, v2 *dataobject.UserContactsDO) {
					imUser.Contacts = append(imUser.Contacts, user.MakeTLContactData(&user.ContactData{
						UserId:        v2.OwnerUserId,
						ContactUserId: v2.ContactUserId,
						FirstName:     mtproto.MakeFlagsString(v2.ContactFirstName),
						LastName:      mtproto.MakeFlagsString(v2.ContactLastName),
						MutualContact: v2.Mutual,
					}).To_ContactData())
				})

				c.svcCtx.Dao.UserPrivaciesDAO.SelectPrivacyListWithCB(c.ctx, v.Id, []int32{
					user.STATUS_TIMESTAMP,
					user.PROFILE_PHOTO,
					user.PHONE_NUMBER,
				}, func(i int, v2 *dataobject.UserPrivaciesDO) {
					r := user.MakeTLPrivacyKeyRules(&user.PrivacyKeyRules{
						Key:   v2.KeyType,
						Rules: nil,
					}).To_PrivacyKeyRules()
					json.Unmarshal([]byte(v2.Rules), &r.Rules)
					imUser.KeysPrivacyRules = append(imUser.KeysPrivacyRules, r)
				})
			}
			vUser.Datas = append(vUser.Datas, imUser)
		}
	})

	if len(doList) == 0 {
		c.Logger.Errorf("getContactsByIdr(%v) - not found", in.Id)
	}

	return vUser, nil
}
