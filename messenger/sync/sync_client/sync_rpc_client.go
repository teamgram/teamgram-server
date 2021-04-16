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

package sync_client

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
)

type syncClient struct {
	client mtproto.RPCSyncClient
}

var (
	syncInstance = &syncClient{}
)

func GetSyncClient() *syncClient {
	return syncInstance
}

func InstallSyncClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	syncInstance.client = mtproto.NewRPCSyncClient(conn)
}

// sync.syncUpdates#3a077679 flags:# layer:int user_id:int auth_key_id:long server_id:flags.1?int not_me:flags.0?true updates:Updates = Bool;
func (c *syncClient) SyncUpdatesMe(userId int32, authKeyId int64, serverId int32, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncSyncUpdates{
		UserId:    userId,
		AuthKeyId: authKeyId,
		ServerId:  serverId,
		Updates:   updates,
	}

	r, err := c.client.SyncSyncUpdates(context.Background(), m)
	return mtproto.FromBool(r), err
}

// sync.syncUpdates#3a077679 flags:# layer:int user_id:int auth_key_id:long server_id:flags.1?int not_me:flags.0?true updates:Updates = Bool;
func (c *syncClient) SyncUpdatesNotMe(userId int32, authKeyId int64, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncSyncUpdates{
		UserId:    userId,
		AuthKeyId: authKeyId,
		Updates:   updates,
	}

	r, err := c.client.SyncSyncUpdates(context.Background(), m)
	return mtproto.FromBool(r), err
}

// sync.pushUpdates#5c612649 user_id:int updates:Updates = Bool;
func (c *syncClient) PushUpdates(userId int32, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncPushUpdates{
		UserId:  userId,
		Updates: updates,
	}

	r, err := c.client.SyncPushUpdates(context.Background(), m)
	return mtproto.FromBool(r), err
}

func (c *syncClient) SyncChannelUpdatesMe(channelId int32, participantId int32, authKeyId int64, serverId int32, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncSyncChannelUpdates{
		ChannelId: channelId,
		UserId:    participantId,
		AuthKeyId: authKeyId,
		ServerId:  serverId,
		Updates:   updates,
	}

	r, err := c.client.SyncSyncChannelUpdates(context.Background(), m)
	return mtproto.FromBool(r), err
}

func (c *syncClient) SyncChannelUpdatesNotMe(channelId int32, participantId int32, authKeyId int64, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncSyncChannelUpdates{
		ChannelId: channelId,
		UserId:    participantId,
		AuthKeyId: authKeyId,
		Updates:   updates,
	}

	r, err := c.client.SyncSyncChannelUpdates(context.Background(), m)
	glog.Infoln("This is err", err)
	return mtproto.FromBool(r), err
}

func (c *syncClient) PushChannelUpdates(channelId, userId int32, updates *mtproto.Updates) (bool, error) {
	m := &mtproto.TLSyncPushChannelUpdates{
		ChannelId: channelId,
		UserId:    userId,
		Updates:   updates,
	}

	r, err := c.client.SyncPushChannelUpdates(context.Background(), m)
	return mtproto.FromBool(r), err
}

// sync.pushRpcResult#1bf9b15e auth_key_id:long req_msg_id:long result:bytes = Bool;
func (c *syncClient) SyncPushRpcResult(authKeyId int64, serverId int32, clientReqMsgId int64, result []byte) (bool, error) {
	m := &mtproto.TLSyncPushRpcResult{
		AuthKeyId: authKeyId,
		ServerId:  serverId,
		ReqMsgId:  clientReqMsgId,
		Result:    result,
	}

	r, err := c.client.SyncPushRpcResult(context.Background(), m)
	return mtproto.FromBool(r), err
}

// sync.getState auth_key_id:long user_id:int = updates.State;
func (c *syncClient) SyncGetState(authKeyId int64, userId int32) (*mtproto.Updates_State, error) {
	req := &mtproto.TLSyncGetState{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}

	state, err := c.client.SyncGetState(context.Background(), req)
	return state, err
}

// sync.getDifference flags:# auth_key_id:long user_id:int pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (c *syncClient) SyncGetDifference(authKeyId int64, userId, pts int32) (*mtproto.Updates_Difference, error) {
	req := &mtproto.TLSyncGetDifference{
		AuthKeyId: authKeyId,
		UserId:    userId,
		Pts:       pts,
		Date:      int32(time.Now().Unix()),
		Qts:       0,
	}

	difference, err := c.client.SyncGetDifference(context.Background(), req)
	return difference, err
}

func (c *syncClient) SyncGetChannelDifference(authKeyId int64, userId, pts int32, channel *mtproto.InputChannel) (*mtproto.Updates_ChannelDifference, error) {
	req := &mtproto.TLSyncGetChannelDifference{
		AuthKeyId: authKeyId,
		UserId:    userId,
		Pts:       pts,
		Channel:   channel,
		// Date:      int32(time.Now().Unix()),
		// Qts:       0,
	}

	difference, err := c.client.SyncGetChannelDifference(context.Background(), req)
	return difference, err
}
