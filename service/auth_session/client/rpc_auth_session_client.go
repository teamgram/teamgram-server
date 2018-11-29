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

package auth_session_client

import (
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

type authSessionClient struct {
	client mtproto.RPCSessionClient
}

var (
	authSessionInstance = &authSessionClient{}
)

func InstallAuthSessionClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	authSessionInstance.client = mtproto.NewRPCSessionClient(conn)
}

func BindAuthKeyUser(authKeyId int64, userId int32) bool {
	request := &mtproto.TLSessionBindAuthKeyUser{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}

	_, err := authSessionInstance.client.SessionBindAuthKeyUser(context.Background(), request)

	if err != nil {
		glog.Error(err)
		return false
	}

	return true
}

func UnbindAuthKeyUser(authKeyId int64, userId int32) bool {
	request := &mtproto.TLSessionUnbindAuthKeyUser{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}

	_, err := authSessionInstance.client.SessionUnbindAuthKeyUser(context.Background(), request)

	if err != nil {
		glog.Error(err)
		return false
	}

	return true
}

func GetLangCode(authKeyId int64) string {
	request := &mtproto.TLSessionGetLangCode{
		AuthKeyId: authKeyId,
	}

	langCode, err := authSessionInstance.client.SessionGetLangCode(context.Background(), request)

	if err != nil {
		glog.Error(err)
		return "en"
	}

	return langCode.GetData2().GetV()
}
