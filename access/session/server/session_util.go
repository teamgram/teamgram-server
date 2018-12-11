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

package server

import (
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
)

func sendSessionDataByConnID(connID uint64, cntl *zrpc.ZRpcController, msg proto.Message) error {
	return util.GAppInstance.(*SessionServer).server.SendMessageByConnID(connID, cntl, msg)
}

func getBizRPCClient() (*grpc_util.RPCClient, error) {
	return util.GAppInstance.(*SessionServer).bizRpcClient, nil
}

func getNbfsRPCClient() (*grpc_util.RPCClient, error) {
	return util.GAppInstance.(*SessionServer).nbfsRpcClient, nil
}

//func getSyncRPCClient() (mtproto.RPCSyncClient, error) {
//	return util.GAppInstance.(*SessionServer).syncRpcClient, nil
//}
//
//func getAuthSessionRPCClient() (mtproto.RPCSessionClient, error) {
//	return util.GAppInstance.(*SessionServer).authSessionRpcClient, nil
//}

func deleteClientSessionManager(authKeyID int64) {
	util.GAppInstance.(*SessionServer).onCloseSessionClientManager(authKeyID)
}

func getServerID() int32 {
	return Conf.ServerId
}

func getUUID() int64 {
	uuid, _ := util.GAppInstance.(*SessionServer).idgen.GetUUID()
	return uuid
}

func setOnline(userId int32, authKeyId int64, serverId, layer int32) {
	util.GAppInstance.(*SessionServer).status.SetSessionOnline(userId, authKeyId, serverId, layer)
}

func setOnlineTTL(userId int32, authKeyId int64, serverId, layer, ttl int32) {
	glog.Infof("setOnlineTTL - {user_id: %d, auth_key_id: %d, server_id: %d, layer: %d, ttl: %d}",
		userId,
		authKeyId,
		serverId,
		layer,
		ttl)
	util.GAppInstance.(*SessionServer).status.SetSessionOnlineTTL(userId, authKeyId, serverId, layer, ttl)
}

func setOffline(userId int32, authKeyId int64, serverId int32) {
	util.GAppInstance.(*SessionServer).status.SetSessionOffline(userId, serverId, authKeyId)
}

func setOfflineTTL(userId int32, authKeyId int64, serverId int32) {
	glog.Infof("setOfflineTTL - {user_id: %d, auth_key_id: %d, server_id: %d}",
		userId,
		authKeyId,
		serverId)
	util.GAppInstance.(*SessionServer).status.SetSessionOfflineTTL(userId, serverId, authKeyId)
}

