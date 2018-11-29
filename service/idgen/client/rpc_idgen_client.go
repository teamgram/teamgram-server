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

package idgen

import (
	"context"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/service/idgen/proto"
	"google.golang.org/grpc"
	"fmt"
)

type RpcIDGenClient struct {
	conn *grpc.ClientConn
}

func rpcUUIDGenClientInstance() UUIDGen {
	cli := &RpcIDGenClient{}
	return cli
}

func rpcSeqIDGenClientInstance() SeqIDGen {
	cli := &RpcIDGenClient{}
	return cli
}

func NewRpcIDGenClient(discovery *service_discovery.ServiceDiscoveryClientConfig) *RpcIDGenClient {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	return &RpcIDGenClient{conn}
}

func (c *RpcIDGenClient) Initialize(config string) error {
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

func (c *RpcIDGenClient) GetUUID() (int64, error) {
	// TODO(@benqi): check c.conn

	cli := seqsvr.NewRPCIDGenClient(c.conn)

	var id int64 = 0
	res, err := cli.GetUUID(context.Background(), &seqsvr.Void{})
	if err != nil {
		glog.Error(err)
	} else {
		id = res.V
	}
	return id, err
}

func (c *RpcIDGenClient) GetCurrentSeqID(key string) (int64, error) {
	// TODO(@benqi): check c.conn

	cli := seqsvr.NewRPCIDGenClient(c.conn)

	var id int64 = 0
	res, err := cli.GetCurrentSeqID(context.Background(), &seqsvr.String{key})
	if err != nil {
		glog.Error(err)
	} else {
		id = res.V
	}
	return id, err
}

func (c *RpcIDGenClient) GetNextSeqID(key string) (int64, error) {
	// TODO(@benqi): check c.conn

	cli := seqsvr.NewRPCIDGenClient(c.conn)

	var id int64 = 0
	res, err := cli.GetNextSeqID(context.Background(), &seqsvr.String{key})
	if err != nil {
		glog.Error(err)
	} else {
		id = res.V
	}
	return id, err
}

func (c *RpcIDGenClient) GetNextNSeqID(key string, n int) (int64, error) {
	return 0, fmt.Errorf("not impl RpcIDGenClient.GetNextNSeqID")
}

func init() {
	// uuid
	UUIDGenRegister("idgen", rpcUUIDGenClientInstance)
	// seqid
	SeqIDGenRegister("idgen", rpcSeqIDGenClientInstance)
}
