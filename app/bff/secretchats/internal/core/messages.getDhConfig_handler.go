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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
)

// MessagesGetDhConfig
// messages.getDhConfig#26cf8950 version:int random_length:int = messages.DhConfig;
func (c *SecretChatsCore) MessagesGetDhConfig(in *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.getDhConfig - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check in.RandomLength?
	//// 400	RANDOM_LENGTH_INVALID	Random length invalid
	//if request.RandomLength != 256 {
	//	err := mtproto.ErrRandomLengthInvalid
	//	log.Errorf("messages.getDhConfig - error: %v", err)
	//	return nil, err
	//}

	var (
		dhConfig *mtproto.Messages_DhConfig
	)

	if in.Version == 0 {
		// TODO(@benqi): 直接设定P和G
		dhConfig = mtproto.MakeTLMessagesDhConfig(&mtproto.Messages_DhConfig{
			G:       c.svcCtx.Config.G,
			P:       c.svcCtx.Config.P,
			Version: 3,
			Random:  crypto.RandomBytes(int(in.RandomLength)),
		}).To_Messages_DhConfig()
	} else {
		// TODO(@benqi): check version and return messages.dhConfigNotModified
		dhConfig = mtproto.MakeTLMessagesDhConfig(&mtproto.Messages_DhConfig{
			G:       c.svcCtx.Config.G,
			P:       c.svcCtx.Config.P,
			Version: 3,
			Random:  crypto.RandomBytes(int(in.RandomLength)),
		}).To_Messages_DhConfig()
	}

	return dhConfig, nil
}
