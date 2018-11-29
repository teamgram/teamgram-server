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

package grpc_util

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/load_balancer"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/service_discovery/etcd3"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

const (
	random          = "random"
	round_robin     = "round_robin"
	consistent_hash = "consistent_hash"
)

func NewRPCClientByServiceDiscovery(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *grpc.ClientConn, err error) {

	//etcdConfg := clientv3.Config{
	//	Endpoints: []string{"http://127.0.0.1:2379"},
	//}
	//r := etcd3.NewResolver("/nebulaim", "auth_key", etcdConfg)
	//b := load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	//c, err := grpc.Dial("", grpc.WithInsecure(),  grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	//if err != nil {
	//	log.Printf("grpc dial: %s", err)
	//	return
	//}
	//defer c.Close()
	//
	//client := mtproto.NewRPCAuthKeyClient(c)

	etcdConfg := clientv3.Config{
		Endpoints: discovery.EtcdAddrs,
	}
	r := etcd3.NewResolver("/nebulaim", discovery.ServiceName, etcdConfg)
	var b grpc.Balancer
	switch discovery.Balancer {
	case "random":
		b = load_balancer.NewBalancer(r, load_balancer.NewRandomSelector())
	case "round_robin":
		b = load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	case "consistent_hash":
		b = load_balancer.NewBalancer(r, load_balancer.NewKetamaSelector(load_balancer.DefaultKetamaKey))
	default:
		b = load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	}

	c, err = grpc.Dial("", grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	return
}

type RPCClient struct {
	conn *grpc.ClientConn
}

func NewRPCClient(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *RPCClient, err error) {
	conn, err := NewRPCClientByServiceDiscovery(discovery)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	c = &RPCClient{
		conn: conn,
	}
	return
}

func (c *RPCClient) GetClientConn() *grpc.ClientConn {
	return c.conn
}

// 通用grpc转发器
func (c *RPCClient) Invoke(rpcMetaData *RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		glog.Error(err)
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		// return nil, err
	}

	glog.Infof("Invoke - method: {%s}", t.Method)
	r := t.NewReplyFunc()
	glog.Infof("Invoke - NewReplyFunc: {%v}, t: {%v}", r, reflect.TypeOf(r))

	var header, trailer metadata.MD

	// ctx := context.Background()
	// glog.Infof("Invoke - NewReplyFunc: {%v}\n", r)
	ctx, _ := RpcMetadataToOutgoing(context.Background(), rpcMetaData)
	glog.Infof("Invoke - NewReplyFunc: {%v}\n", r)
	err := c.conn.Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))

	glog.Infof("header: {%v}, trailer: {%v}", header, trailer)

	// TODO(@benqi): process header from serverF
	// grpc.Header(&header)
	// glog.Infof("Invoke - error: {%v}", err)

	if err != nil {
		glog.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: \n", t.Method, c.conn, err)
		// TODO(@benqi): 哪些情况需要断开客户端连接
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			// TODO(@benqi): Rpc error, trailer has rpc_error metadata
			case codes.Unknown:
				return nil, RpcErrorFromMD(trailer)
			}
		}
		return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
	} else {
		// glog.Infof("Invoke - Invoke reply: {%v}\n", r)
		reply, ok := r.(mtproto.TLObject)

		glog.Infof("Invoke %s time: %d", t.Method, (time.Now().Unix() - rpcMetaData.ReceiveTime))

		if !ok {
			err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
			glog.Error(err)
			return nil, mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		}

		return reply, nil
	}
}
