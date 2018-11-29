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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/mtproto"
)

type sessionClient struct {
	client *net2.TcpClient
}

func (s *sessionClient) OnNewClient(client *net2.TcpClient) {
	glog.Infof("OnNewConnection")

	req_pq := &mtproto.TLReqPq{
		Nonce: crypto.GenerateNonce(16),
	}
	_ = req_pq

	//authKeyMD := &mtproto.AuthKeyMetadata{}
	//state := &zproto.HandshakeState{
	//	State:    zproto.STATE_pq,
	//	ResState: zproto.RES_STATE_NONE,
	//}
	//state.Ctx, _ = proto.Marshal(authKeyMD)
	//
	//smsg := &zproto.SessionHandshakeMessage{
	//	State: state,
	//	MTPMessage: &mtproto.UnencryptedMessage{
	//		MessageId: 0,
	//		Object:    req_pq,
	//	},
	//}
	//
	//zmsg := &zproto.ZProtoMessage{
	//	// SessionId: 0,
	//	// SeqNum: 0,
	//	Metadata: &zproto.ZProtoMetadata{},
	//	Message: &zproto.ZProtoRawPayload{
	//		Payload: smsg.Encode(),
	//	},
	//}
	//
	//client.Send(zmsg)
}

func (s *sessionClient) OnClientDataArrived(client *net2.TcpClient, msg interface{}) error {
	glog.Infof("OnDataArrived - recv data: %v", msg)
	return nil
}

func (s *sessionClient) OnClientClosed(client *net2.TcpClient) {
	glog.Infof("OnConnectionClosed")
}

func (s *sessionClient) OnClientTimer(client *net2.TcpClient) {
	glog.Infof("OnTimer")
}

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	client := &sessionClient{}
	client.client = net2.NewTcpClient("session", 1024, "zproto", "127.0.0.1:10000", client)
	client.client.Serve()
	select {}
}
