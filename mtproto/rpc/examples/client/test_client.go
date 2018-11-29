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

package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
)

var (
	confPath string
	Conf     *zrpc.ZRpcClientConfig
)

type TestClientInsance struct {
	client *zrpc.ZRpcClient
}

func (inst *TestClientInsance) Initialize() error {
	InitializeConfig()
	inst.client = zrpc.NewZRpcClient("brpc", Conf, inst)
	return nil
}

func (inst *TestClientInsance) RunLoop() {
	// glog.Info("runLoop")
	inst.client.Serve()
	// glog.Info("endRunLoop")
}

func (inst *TestClientInsance) Destroy() {
	inst.client.Stop()
}

func (inst *TestClientInsance) OnNewClient(client *net2.TcpClient) {
	glog.Info("onNewClient")
}

func (inst *TestClientInsance) OnClientMessageArrived(client *net2.TcpClient, cntl *zrpc.ZRpcController, msg proto.Message) error {
	glog.Info("onClientMessageArrived")
	return nil
}

func (inst *TestClientInsance) OnClientClosed(client *net2.TcpClient) {
	glog.Info("onClientClosed")
}

func (inst *TestClientInsance) OnClientTimer(client *net2.TcpClient) {
	glog.Info("onClientTimer")
}

func init() {
	flag.StringVar(&confPath, "conf", "./test_client.toml", "config path")
}

func InitializeConfig() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	if err != nil {
		err = fmt.Errorf("decode file %s error: %v", confPath, err)
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
func main() {
	instance := &TestClientInsance{}
	util.DoMainAppInstance(instance)
}
