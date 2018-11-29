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
	"context"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/service/status/proto"
	"google.golang.org/grpc"
	"time"
)

type rpcStatusClient struct {
	conn *grpc.ClientConn
}

func rpcStatusClientInstance() StatusClient {
	cli := &rpcStatusClient{}
	return cli
}

func NewRpcStatusClient(discovery *service_discovery.ServiceDiscoveryClientConfig) *rpcStatusClient {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	return &rpcStatusClient{conn}
}

func (c *rpcStatusClient) Initialize(config string) error {
	var err error

	discovery := &service_discovery.ServiceDiscoveryClientConfig{}
	err = json.Unmarshal([]byte(config), &discovery)
	if err != nil {
		glog.Error(err)
		return err
	}

	c.conn, err = grpc_util.NewRPCClientByServiceDiscovery(discovery)
	if err != nil {
		glog.Error(err)
	}

	return err
}

func (c *rpcStatusClient) SetSessionOnline(userId int32, authKeyId int64, serverId, layer int32) error {
	cli := status.NewRPCStatusClient(c.conn)
	session := &status.SessionEntry{
		UserId:    userId,
		ServerId:  serverId,
		AuthKeyId: authKeyId,
		Expired:   time.Now().Unix() + 120,
		Layer:     layer,
	}
	_, err := cli.SetSessionOnline(context.Background(), session)
	return err
}

func (c *rpcStatusClient) SetSessionOffline(userId int32, serverId int32, authKeyId int64) error {
	cli := status.NewRPCStatusClient(c.conn)
	session := &status.SessionEntry{
		UserId:    userId,
		ServerId:  serverId,
		AuthKeyId: authKeyId,
		Expired:   0,
	}
	_, err := cli.SetSessionOffline(context.Background(), session)
	return err
}

func (c *rpcStatusClient) GetUserOnlineSessions(userId int32) (*status.SessionEntryList, error) {
	cli := status.NewRPCStatusClient(c.conn)
	return cli.GetUserOnlineSessions(context.Background(), &status.Int32{V: userId})
}

func (c *rpcStatusClient) GetUsersOnlineSessionsList(userIdList []int32) (*status.UsersSessionEntryList, error) {
	cli := status.NewRPCStatusClient(c.conn)
	return cli.GetUsersOnlineSessionsList(context.Background(), &status.Int32List{Vlist: userIdList})
}

func init() {
	Register("rpc", rpcStatusClientInstance)
}
