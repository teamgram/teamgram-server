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
	"flag"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/config"
	msg_client "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/client"
	sync_client "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	authsession_client "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
	chat_client "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	user_client "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	username_client "github.com/teamgram/teamgram-server/v2/app/service/biz/username/client"
	status_client "github.com/teamgram/teamgram-server/v2/app/service/status/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"

	"github.com/oschwald/geoip2-golang"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

var (
	mmdb string
)

func init() {
	flag.StringVar(&mmdb, "mmdb", "./GeoLite2-City.mmdb", "mmdb")
}

type Dao struct {
	kv   kv.Store
	MMDB *geoip2.Reader
	authsession_client.AuthsessionClient
	user_client.UserClient
	sync_client.SyncClient
	chat_client.ChatClient
	status_client.StatusClient
	msg_client.MsgClient
	username_client.UsernameClient
}

func New(c config.Config) *Dao {
	MMDB, err := geoip2.Open(mmdb)
	if err != nil {
		// panic(err)
	}
	return &Dao{
		kv:                kv.NewStore(c.KV),
		MMDB:              MMDB,
		UserClient:        user_client.NewUserClient(kitex.GetCachedKitexClient(c.UserClient)),
		AuthsessionClient: authsession_client.NewAuthsessionClient(kitex.GetCachedKitexClient(c.AuthsessionClient)),
		ChatClient:        chat_client.NewChatClient(kitex.GetCachedKitexClient(c.ChatClient)),
		StatusClient:      status_client.NewStatusClient(kitex.GetCachedKitexClient(c.StatusClient)),
		MsgClient:         msg_client.NewMsgClient(kitex.GetCachedKitexClient(c.MsgClient)),
		UsernameClient:    username_client.NewUsernameClient(kitex.GetCachedKitexClient(c.UsernameClient)),
		// SyncClient:        sync_client.NewSyncMqClient(kitex.GetCachedKitexClient(c.SyncClient)),
	}
}
