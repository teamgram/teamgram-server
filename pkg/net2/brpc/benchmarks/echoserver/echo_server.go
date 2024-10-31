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
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teamgram/teamgram-server/pkg/net2/brpc"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/zeromicro/go-zero/core/logx"
)

type echoServer struct {
	gnet.BuiltinEventEngine
	eng       gnet.Engine
	addr      string
	multicore bool
}

func (s *echoServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logx.Infof("running server on %s with multi-core=%t", s.addr, s.multicore)
	s.eng = eng
	return
}

func (s *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(brpc.NewBaiduRpcCodec())
	// atomic.AddInt32(&s.connected, 1)
	return
}

func (s *echoServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		logx.Errorf("error occurred on connection=%s, %v", c.RemoteAddr().String(), err)
	}
	return
}

func (s *echoServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	codec := c.Context().(*brpc.BaiduRpcCodec)

	for {
		msg, err := codec.Decode(c)
		if errors.Is(err, brpc.ErrIncompletePacket) {
			break
		}
		if err != nil {
			logx.Errorf("invalid packet: %v", err)
			return gnet.Close
		}

		// fmt.Println("connId: ", c.ConnId(), ", ", msg.String())
		vList, _ := msg.Encode()
		_, _ = c.Writev(vList)
	}

	return
}

func (s *echoServer) OnTick() (delay time.Duration, action gnet.Action) {
	logx.Infof("conn count: %d", s.eng.CountConnections())
	delay = time.Second

	return
}

func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9300, "--port 9300")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()

	//logx.SetUp(logx.LogConf{
	//	ServiceName:         "",
	//	Mode:                "",
	//	Encoding:            "",
	//	TimeFormat:          "",
	//	Path:                "",
	//	Level:               "debug",
	//	MaxContentLength:    0,
	//	Compress:            false,
	//	Stat:                false,
	//	KeepDays:            0,
	//	StackCooldownMillis: 0,
	//	MaxBackups:          0,
	//	MaxSize:             0,
	//	Rotation:            "",
	//	FileTimeFormat:      "",
	//})

	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	go func() {
		logx.Info("gnet server exits:", echo.addr)
		gnet.Run(
			echo,
			echo.addr,
			gnet.WithMulticore(multicore),
			gnet.WithLockOSThread(true),
			gnet.WithTicker(true),
			gnet.WithLogLevel(logging.DebugLevel),
			gnet.WithLogger(brpc.NewLogger()))
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	echo.eng.Stop(context.TODO())
}
