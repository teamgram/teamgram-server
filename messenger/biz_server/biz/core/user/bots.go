// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package user

import "github.com/nebula-chat/chatengine/mtproto"

func (m *UserModel) GetBotInfo(botId int32) *mtproto.BotInfo {
	botsDO := m.dao.BotsDAO.Select(botId)
	if botsDO == nil {
		return nil
	}
	botInfo := &mtproto.TLBotInfo{Data2: &mtproto.BotInfo_Data{
		UserId: botId,
		Description: botsDO.Description,
	}}

	botCommandsDOList := m.dao.BotCommandsDAO.SelectList(botId)
	for i := 0; i < len(botCommandsDOList); i++ {
		botCommand := &mtproto.TLBotCommand{Data2: &mtproto.BotCommand_Data{
			Command:     botCommandsDOList[i].Command,
			Description: botCommandsDOList[i].Description,
		}}
		botInfo.Data2.Commands = append(botInfo.Data2.Commands, botCommand.To_BotCommand())
	}

	return botInfo.To_BotInfo()
}
