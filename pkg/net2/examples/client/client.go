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
	"fmt"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/pkg/net2/brpc"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
)

type clientEvents struct {
	*gnet.BuiltinEventEngine
	conns sync.Map
}

func (ev *clientEvents) OnBoot(e gnet.Engine) gnet.Action {
	return gnet.None
}

func (ev *clientEvents) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(brpc.NewBaiduRpcCodec())
	return
}

func (ev *clientEvents) OnClose(gnet.Conn, error) gnet.Action {
	return gnet.None
}

func (ev *clientEvents) OnTraffic(c gnet.Conn) (action gnet.Action) {
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

		fmt.Println(msg)
		// vList, _ := codec.Encode(c, msg)
		// _, _ = c.Writev(vList)
	}

	return
}

func (ev *clientEvents) OnTick() (delay time.Duration, action gnet.Action) {
	ev.conns.Range(func(key, value any) bool {
		c := value.(gnet.Conn)
		msg := &brpc.BaiduRpcMessage{
			Meta:       new(brpc.RpcMeta),
			Payload:    []byte("12345"),
			Attachment: []byte("0123456789"),
		}
		msg.Meta.CorrelationId = 1
		msg.Meta.CompressType = 1

		vList, _ := msg.Encode()
		c.AsyncWritev(vList, func(c gnet.Conn, err error) error {
			fmt.Println("cb: ", c.ConnId())
			return nil
		})

		return true
	})

	delay = time.Second
	return
}

func (ev *clientEvents) OnShutdown(e gnet.Engine) {
	//fd, err := e.Dup()
	//require.ErrorIsf(ev.tester, err, errorx.ErrEmptyEngine, "expected error: %v, but got: %v",
	//	errorx.ErrUnsupportedOp, err)
	//assert.EqualValuesf(ev.tester, fd, -1, "expected -1, but got: %d", fd)
}

func main() {
	clientEV := &clientEvents{}
	client, _ := gnet.NewClient(
		clientEV,
		gnet.WithLogLevel(logging.DebugLevel),
		gnet.WithLockOSThread(true),
		gnet.WithTicker(true),
	)

	client.Start()
	defer client.Stop() //nolint:errcheck

	c, _ := client.Dial("tcp", "127.0.0.1:9300")
	clientEV.conns.Store("127.0.0.1:9300", c)
	defer c.Close()
	// c.Writev(vList)

	time.Sleep(100 * time.Second)
}
