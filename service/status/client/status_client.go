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

package status_client

import (
	"fmt"
	"github.com/nebula-chat/chatengine/service/status/proto"
)

const (
	ONLINE_TIMEOUT       = 60 // 15秒
	CHECK_ONLINE_TIMEOUT = 70 // 17秒, 15+2秒的误差
	// onlineKeyPrefix = "online"		//
)

type StatusClient interface {
	Initialize(config string) error
	SetSessionOnline(userId int32, authKeyId int64, serverId, layer int32) error
	SetSessionOffline(userId int32, serverId int32, authKeyId int64) error
	GetUserOnlineSessions(userId int32) (*status.SessionEntryList, error)
	GetUsersOnlineSessionsList(userIdList []int32) (*status.UsersSessionEntryList, error)
}

type Instance func() StatusClient

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("status_client: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("status_client: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewStatusClient(adapterName, config string) (adapter StatusClient, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("status_client: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.Initialize(config)
	if err != nil {
		adapter = nil
	}
	return
}
