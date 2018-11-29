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
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/gogo/protobuf/proto"
)

var (
	confPath string
	Conf     *zrpc.ZRpcServerConfig
)

type TestServerInsance struct {
	server *zrpc.ZRpcServer
}

func (s *TestServerInsance) Initialize() error {
	InitializeConfig()
	s.server = zrpc.NewZRpcServer(Conf, s)
	return nil
}

func (s *TestServerInsance) RunLoop() {
	// go this.server.httpServer.Serve(this.server.httpListener)
	s.server.Serve()
	// this.client.Serve()
}

func (s *TestServerInsance) Destroy() {
	s.server.Stop()
}

func (s *TestServerInsance) OnServerNewConnection(conn *net2.TcpConnection) {

}

func (s *TestServerInsance) OnServerMessageDataArrived(c *net2.TcpConnection, cntl *zrpc.ZRpcController, msg proto.Message) error {
	return nil
}

func (s *TestServerInsance) OnServerConnectionClosed(c *net2.TcpConnection) {
}

func init() {
	flag.StringVar(&confPath, "conf", "./test_server.toml", "config path")
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
	instance := &TestServerInsance{}
	util.DoMainAppInstance(instance)
}
