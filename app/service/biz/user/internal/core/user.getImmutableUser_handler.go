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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// UserGetImmutableUser
// user.getImmutableUser id:long = ImmutableUser;
func (c *UserCore) UserGetImmutableUser(in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	userDO, _ := c.svcCtx.Dao.UsersDAO.SelectById(c.ctx, in.Id)
	if userDO == nil {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("user.getImmutableUser - error: %v", err)
		return nil, err
	}

	imUser := user.MakeTLImmutableUser(&user.ImmutableUser{
		User:             nil,
		LastSeenAt:       0,
		Contacts:         nil,
		KeysPrivacyRules: nil,
	}).To_ImmutableUser()

	// deleted
	if userDO.Deleted {
		imUser.User = user.MakeTLUserData(&user.UserData{
			Id:         userDO.Id,
			AccessHash: userDO.AccessHash,
			Deleted:    true,
		}).To_UserData()
	}

	imUser.User = user.MakeTLUserData(&user.UserData{
		Id:                userDO.Id,
		AccessHash:        userDO.AccessHash,
		UserType:          userDO.UserType,
		SceretKeyId:       userDO.SecretKeyId,
		FirstName:         userDO.FirstName,
		LastName:          userDO.LastName,
		Username:          userDO.Username,
		Phone:             userDO.Phone,
		ProfilePhoto:      nil, //
		Bot:               nil,
		CountryCode:       userDO.CountryCode,
		Verified:          userDO.Verified,
		Support:           userDO.Support,
		Scam:              userDO.Scam,
		Fake:              userDO.Fake,
		About:             mtproto.MakeFlagsString(userDO.About),
		Restricted:        userDO.Restricted,
		RestrictionReason: nil,
		ContactsVersion:   1,
		PrivaciesVersion:  1,
		Deleted:           false,
	}).To_UserData()

	//// 1. bot
	if userDO.IsBot {
		botDO, _ := c.svcCtx.Dao.BotsDAO.Select(c.ctx, in.Id)
		if botDO != nil {
			// userData.Bot
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

			//
		}
	}

	// 2. ProfilePhoto
	// TODO: ProfilePhoto
	// (userDO.RestrictionReason)

	imUser.User.ProfilePhoto, _ = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &media.TLMediaGetPhoto{
		PhotoId: userDO.PhotoId,
	})

	// 3. RestrictionReason
	// TODO: RestrictionReason
	if imUser.User.Restricted {
		json.Unmarshal([]byte(userDO.RestrictionReason), &imUser.User.RestrictionReason)
	}

	// 4. LastSeenAt
	lastSeenDO, _ := c.svcCtx.Dao.UserPresencesDAO.Select(c.ctx, in.Id)
	if lastSeenDO != nil {
		imUser.LastSeenAt = lastSeenDO.LastSeenAt
	}

	return imUser, nil
}
