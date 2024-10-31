// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"flag"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/teamgram/teamgram-server/pkg/net2/brpc"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
)

type simpleServer struct {
	gnet.BuiltinEventEngine
	eng          gnet.Engine
	network      string
	addr         string
	multicore    bool
	connected    int32
	disconnected int32
}

func (s *simpleServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logging.Infof("running server on %s with multi-core=%t",
		fmt.Sprintf("%s://%s", s.network, s.addr), s.multicore)
	s.eng = eng
	return
}

func (s *simpleServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(brpc.NewBaiduRpcCodec())
	atomic.AddInt32(&s.connected, 1)
	return
}

func (s *simpleServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		logging.Infof("error occurred on connection=%s, %v", c.RemoteAddr().String(), err)
	}
	//disconnected := atomic.AddInt32(&s.disconnected, 1)
	//connected := atomic.AddInt32(&s.connected, -1)
	//if connected == 0 {
	//	logging.Infof("all %d connections are closed, shut it down", disconnected)
	//	action = gnet.Shutdown
	//}
	return
}

func (s *simpleServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	codec := c.Context().(*brpc.BaiduRpcCodec)

	for {
		msg, err := codec.Decode(c)
		if errors.Is(err, brpc.ErrIncompletePacket) {
			break
		}
		if err != nil {
			logging.Errorf("invalid packet: %v", err)
			return gnet.Close
		}

		fmt.Println("connId: ", c.ConnId(), ", ", msg.String())
		vList, _ := msg.Encode()
		_, _ = c.Writev(vList)
	}

	return
}

func main() {
	var port int
	var multicore bool

	// Example command: go run server.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9300, "--port 9300")
	flag.BoolVar(&multicore, "multicore", false, "--multicore=true")
	flag.Parse()
	ss := &simpleServer{
		network:   "tcp",
		addr:      fmt.Sprintf(":%d", port),
		multicore: multicore,
	}
	err := gnet.Run(
		ss,
		ss.network+"://"+ss.addr,
		gnet.WithMulticore(multicore),
		gnet.WithReuseAddr(true),
		gnet.WithReusePort(true))
	logging.Infof("server exits with error: %v", err)
	time.Sleep(time.Hour)
}
