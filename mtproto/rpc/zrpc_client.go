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

package zrpc

import (
	"fmt"
	"math/rand"
	"github.com/coreos/etcd/clientv3"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util/load_balancer"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/net2/watcher2"
	"github.com/nebula-chat/chatengine/mtproto/rpc/brpc"
	"github.com/gogo/protobuf/proto"
	"github.com/nebula-chat/chatengine/mtproto"
)

func init() {
	proto.RegisterType((*mtproto.TLPing)(nil), "mtproto.TLPing")
	proto.RegisterType((*mtproto.Pong)(nil), "mtproto.Pong")
}

type ZPpcClientCallBack interface {
	OnNewClient(client *net2.TcpClient)
	OnClientMessageArrived(client *net2.TcpClient, cntl *ZRpcController, msg proto.Message) error
	OnClientClosed(client *net2.TcpClient)
	OnClientTimer(client *net2.TcpClient)
}

type ZRpcClientConfig struct {
	Clients []net2.ClientConfig
}

type Watcher struct {
	name    string
	watcher *watcher2.ClientWatcher
	ketama  *load_balancer.Ketama
}

type ZRpcClient struct {
	watchers []*Watcher
	clients  *net2.TcpClientGroupManager
	callback ZPpcClientCallBack
}

func NewZRpcClient(protoName string, conf *ZRpcClientConfig, cb ZPpcClientCallBack) *ZRpcClient {
	clients := map[string][]string{}

	c := &ZRpcClient{
		callback: cb,
	}

	c.clients = net2.NewTcpClientGroupManager(protoName, clients, c)

	// Check name
	for i := 0; i < len(conf.Clients); i++ {
		// service discovery
		etcdConfg := clientv3.Config{
			Endpoints: conf.Clients[i].EtcdAddrs,
		}
		watcher := &Watcher{
			name: conf.Clients[i].Name,
		}
		watcher.watcher, _ = watcher2.NewClientWatcher("/nebulaim", conf.Clients[i].Name, etcdConfg, c.clients)
		if conf.Clients[i].Balancer == "ketama" {
			watcher.ketama = load_balancer.NewKetama(10, nil)
		}
		c.watchers = append(c.watchers, watcher)
	}

	return c
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (c *ZRpcClient) Serve() {
	for _, w := range c.watchers {
		if w.ketama != nil {
			go w.watcher.WatchClients(func(etype, addr string) {
				switch etype {
				case "add":
					w.ketama.Add(addr)
				case "delete":
					w.ketama.Remove(addr)
				}
			})
		} else {
			go w.watcher.WatchClients(nil)
		}
	}
}

func (c *ZRpcClient) Stop() {
	c.clients.Stop()
}

func (c *ZRpcClient) Pause() {
	// s.clients.Pause()
}

func (c *ZRpcClient) selectKetama(name string) *load_balancer.Ketama {
	for _, w := range c.watchers {
		if w.name == name && w.ketama != nil {
			return w.ketama
		}
	}

	return nil
}

func (c *ZRpcClient) SendKetamaMessage(name, key string, cntl *ZRpcController, msg proto.Message, f func(addr string)) error {
	ketama := c.selectKetama(name)
	if ketama == nil {
		err := fmt.Errorf("not found ketama by name: %s", name)
		glog.Error(err)
		return err
	}

	if kaddr, ok := ketama.Get(key); ok {
		if f != nil {
			f(kaddr)
		}
		return c.SendMessageToAddress(name, kaddr, cntl, msg)
	} else {
		err := fmt.Errorf("not found kaddr by key: %s", key)
		glog.Error(err)
		return err
	}
}

func (c *ZRpcClient) SendMessage(name string, cntl *ZRpcController, msg proto.Message) error {
	bmsg := MakeBaiduRpcMessage(cntl, msg)
	return c.clients.SendData(name, bmsg)
}

func (c *ZRpcClient) SendMessageToAddress(name, addr string, cntl *ZRpcController, msg proto.Message) error {
	bmsg := MakeBaiduRpcMessage(cntl, msg)
	return c.clients.SendDataToAddress(name, addr, bmsg)
}

///////////////////////////////////////////////////////////////////////////////////////////////
func (c *ZRpcClient) OnNewClient(client *net2.TcpClient) {
	// glog.Info("onNewClient - client: ", client, ", conn: ", client.GetConnection())

	codec := client.GetConnection().Codec()
	glog.Info("codec: ", codec)

	client.StartTimer()
	c.SendPing(client)

	if c.callback != nil {
		c.callback.OnNewClient(client)
	}
}

func (c *ZRpcClient) OnClientDataArrived(client *net2.TcpClient, msg interface{}) error {
	bmsg, ok := msg.(*brpc.BaiduRpcMessage)
	if !ok {
		return fmt.Errorf("recv invalid bmsg - {%v}", bmsg)
	}

	if bmsg.Meta.GetRequest() == nil {
		return fmt.Errorf("invalid meta - {%v}", bmsg)
	}

	cntl, zmsg, err := SplitBaiduRpcMessage(bmsg)
	if err != nil {
		glog.Error(err)
		return err
	}

	switch zmsg.(type) {
	case *mtproto.Pong:
		glog.Info("recv pong: ", zmsg)
		return nil
	default:
		if c.callback != nil {
			return c.callback.OnClientMessageArrived(client, cntl, zmsg)
		} else {
			err = fmt.Errorf("callback is nil")
			return err
		}
	}
}

func (c *ZRpcClient) OnClientClosed(client *net2.TcpClient) {
	// glog.Infof("onClientClosed - peer(%s)", client.GetConnection())

	if c.callback != nil {
		c.callback.OnClientClosed(client)
	}

	if client.AutoReconnect() {
		client.Reconnect()
	}
}

func (c *ZRpcClient) OnClientTimer(client *net2.TcpClient) {
	c.SendPing(client)

	if c.callback != nil {
		c.callback.OnClientTimer(client)
	}
}

func (c *ZRpcClient) SendPing(client *net2.TcpClient) {
	ping := &mtproto.TLPing{
		PingId: rand.Int63(),
	}

	cntl := NewController()
	cntl.SetServiceName("zrpc")
	cntl.SetMethodName(proto.MessageName(ping))

	glog.Info("sendPing: ", ping)
	SendMessageByClient(client, cntl, ping)
}

///////////////////////////////////////////////////////////////////////////////////////////////
func SendMessageByClient(client *net2.TcpClient, cntl *ZRpcController, msg proto.Message) error {
	bmsg := MakeBaiduRpcMessage(cntl, msg)
	glog.Info(bmsg)
	return client.Send(bmsg)
}
