// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package repository

import (
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// Type aliases for convenience in the Logic layer.
type (
	UserClient               = userclient.UserClient
	ChatClient               = chatclient.ChatClient
	GetNotifySettings        = userpb.TLUserGetNotifySettings
	GetNotifySettingsList    = userpb.TLUserGetNotifySettingsList
	SetNotifySettings        = userpb.TLUserSetNotifySettings
	ResetNotifySettings      = userpb.TLUserResetNotifySettings
	NotifySettingsListResult = userpb.VectorPeerPeerNotifySettings
	GetMutableChat           = chatpb.TLChatGetMutableChat
	GetChatListByIdList      = chatpb.TLChatGetChatListByIdList
	MutableChat              = tg.MutableChat
	MutableChatClazz         = tg.MutableChatClazz
)
