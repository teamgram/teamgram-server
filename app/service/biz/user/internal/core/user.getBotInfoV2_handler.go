// Copyright 2025 Teamgram Authors
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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// UserGetBotInfoV2
// user.getBotInfoV2 bot_id:long = BotInfoData;
func (c *UserCore) UserGetBotInfoV2(in *user.TLUserGetBotInfoV2) (*user.BotInfoData, error) {
	botsDO, err := c.svcCtx.BotsDAO.Select(c.ctx, in.BotId)
	if err != nil {
		c.Logger.Errorf("user.getBotInfo - error: %v", err)
		return nil, err
	} else if botsDO == nil {
		return nil, mtproto.ErrUserIdInvalid
	}

	botInfo := mtproto.MakeTLBotInfo(&mtproto.BotInfo{
		HasPreviewMedias:       botsDO.HasPreviewMedias,
		UserId_INT64:           in.BotId,
		UserId_FLAGINT64:       mtproto.MakeFlagsInt64(in.BotId),
		Description_STRING:     botsDO.Description,
		Description_FLAGSTRING: mtproto.MakeFlagsString(botsDO.Description),
		DescriptionPhoto:       nil,
		DescriptionDocument:    nil,
		Commands:               []*mtproto.BotCommand{},
		MenuButton:             nil,
		PrivacyPolicyUrl:       mtproto.MakeFlagsString(botsDO.PrivacyPolicyUrl),
		AppSettings:            nil,
	}).To_BotInfo()

	// TODO: HasPreviewMedias

	// Commands
	_, _ = c.svcCtx.Dao.BotCommandsDAO.SelectListWithCB(
		c.ctx,
		in.BotId,
		func(sz, i int, v *dataobject.BotCommandsDO) {
			botInfo.Commands = append(botInfo.Commands, mtproto.MakeTLBotCommand(&mtproto.BotCommand{
				Command:     v.Command,
				Description: v.Description,
			}).To_BotCommand())
		})

	// MenuButton
	if botsDO.HasMenuButton {
		botInfo.MenuButton = mtproto.MakeTLBotMenuButton(&mtproto.BotMenuButton{
			Text: botsDO.MenuButtonText,
			Url:  botsDO.MenuButtonUrl,
		}).To_BotMenuButton()
	}

	// DescriptionPhoto
	if botsDO.DescriptionPhotoId != 0 {
		botInfo.DescriptionPhoto, _ = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &media.TLMediaGetPhoto{
			PhotoId: botsDO.DescriptionPhotoId,
		})
	}

	// AppSettings
	if botsDO.HasAppSettings {
		botInfo.AppSettings = mtproto.MakeTLBotAppSettings(&mtproto.BotAppSettings{
			PlaceholderPath:     nil, // TODO: botsDO.PlaceholderPath,
			BackgroundColor:     mtproto.MakeFlagsInt32(botsDO.BackgroundColor),
			BackgroundDarkColor: mtproto.MakeFlagsInt32(botsDO.BackgroundDarkColor),
			HeaderColor:         mtproto.MakeFlagsInt32(botsDO.HeaderColor),
			HeaderDarkColor:     mtproto.MakeFlagsInt32(botsDO.HeaderDarkColor),
		}).To_BotAppSettings()
	}

	return user.MakeTLBotInfoData(&user.BotInfoData{
		BotInfo:    botInfo,
		MainAppUrl: mtproto.MakeFlagsString(botsDO.MainAppUrl),
		BotInline:  botsDO.BotInlinePlaceholder != "",
		Token:      botsDO.Token,
		BotId:      botsDO.BotId,
	}).To_BotInfoData(), nil
}
