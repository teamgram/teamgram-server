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

package net2

import (
	"fmt"
	"github.com/golang/glog"
	"go.uber.org/atomic"
	"net"
	"testing"
)

type TestTcpPingClientGroup struct {
	client   *TcpClientGroupManager
	ready    atomic.Int32
	receiver chan interface{}
}

func NewTestClientGroup(protoName string, clients map[string][]string, receiver chan interface{}) *TestTcpPingClientGroup {
	c := &TestTcpPingClientGroup{}
	c.client = NewTcpClientGroupManager(protoName, clients, c)
	c.receiver = receiver
	return c
}

func (c *TestTcpPingClientGroup) Serve() {
	c.client.Serve()
}

func (c *TestTcpPingClientGroup) Stop() {
	c.client.Stop()
}

func (c *TestTcpPingClientGroup) Send(name string, msg interface{}) error {
	return c.client.SendData(name, msg)
}

func (c *TestTcpPingClientGroup) OnNewClient(client *TcpClient) {
	glog.Infof("client group OnNewConnection %v", client.GetRemoteName())
	c.ready.Add(1)
}

func (c *TestTcpPingClientGroup) OnClientDataArrived(client *TcpClient, msg interface{}) error {
	glog.Infof("client group OnClientDataArrived - client: %s, receive data: %v", client.GetRemoteName(), msg)
	c.receiver <- msg
	return nil
}

func (c *TestTcpPingClientGroup) OnClientClosed(client *TcpClient) {
	glog.Infof("client group OnConnectionClosed %s", client.GetRemoteName())
}

func (c *TestTcpPingClientGroup) OnClientTimer(client *TcpClient) {
	glog.Infof("client group OnTimer")
}

//////////////////////////////////////////////////////////////////////////////////////////
type TestClientGroupSimulation struct {
	protoName string
	// ----- server -----
	serverCnt      int
	serverChanSize int
	serverMaxConn  int
	// ----- client -----
	clientMsgCnt      int
	clientChanSizeCnt int

	receivedChan chan interface{}

	servers []*TestPingPongServer
	client  *TestTcpPingClientGroup
}

func (ts *TestClientGroupSimulation) simulate() (result []string, e error) {
	serviceName := "test_service"
	services := make(map[string][]string)
	servers := make([]string, 0)
	ts.servers = make([]*TestPingPongServer, 0)

	for i := 0; i < ts.serverCnt; i++ {
		listener, err := net.Listen("tcp", "0.0.0.0:0")
		if err != nil {
			glog.Errorf("listen error: %v", err)
			return nil, err
		}
		server := NewTestServer(listener, fmt.Sprintf(`TestPingPongServer%d`, i), ts.protoName, ts.serverChanSize, ts.serverMaxConn)
		go server.Serve()
		for { // wait to ready
			if server.isReady() {
				break
			}
		}
		ts.servers = append(ts.servers, server)
		servers = append(servers, listener.Addr().String())
	}

	services[serviceName] = servers
	ts.client = NewTestClientGroup(ts.protoName, services, ts.receivedChan)
	ts.client.Serve()
	for { // wait to ready
		if int(ts.client.ready.Load()) == ts.serverCnt {
			break
		}
	}

	errChan := make(chan error)

	for i := 0; i < ts.clientMsgCnt; i++ {

		go func(c *TestTcpPingClientGroup, errorChan chan error, n int) {
			if err := c.Send(serviceName, fmt.Sprintf("client_group(ping%d)\n", n)); err != nil {
				errorChan <- err
			}
		}(ts.client, errChan, i)
	}

	result = make([]string, 0)

	for cnt := 0; cnt < ts.clientMsgCnt; cnt++ {
		select {
		case msg, _ := <-ts.receivedChan:
			result = append(result, msg.(string))
		case m, _ := <-errChan:
			result = append(result, m.Error())
		}
	}

	return
}

func (ts *TestClientGroupSimulation) stop() {
	for _, s := range ts.servers {
		s.Stop()
	}

	ts.client.Stop()
}

func TestClientGroupServer(t *testing.T) {
	protoName := "TestProto"

	RegisterProtocol(protoName, &TestProto{})

	testTable := []struct {
		serverCnt         int
		serverChanSize    int
		serverMaxConn     int
		clientMsgCnt      int
		clientChanSizeCnt int
	}{
		{
			// n server, 1 client
			clientMsgCnt:      1000,
			clientChanSizeCnt: 1024,
			serverCnt:         10,
			serverChanSize:    1024,
			serverMaxConn:     1024,
		},
	}

	for _, data := range testTable {

		simulation := &TestClientGroupSimulation{
			protoName:         protoName,
			serverChanSize:    data.serverChanSize,
			serverMaxConn:     data.serverMaxConn,
			serverCnt:         data.serverCnt,
			clientMsgCnt:      data.clientMsgCnt,
			clientChanSizeCnt: data.clientChanSizeCnt,
			receivedChan:      make(chan interface{}, 0),
		}

		result, err := simulation.simulate()
		if err != nil {
			t.Error(err)
			return
		}

		simulation.stop()

		expected := make(map[string]bool)
		for i := 0; i < simulation.clientMsgCnt; i++ {
			expected[fmt.Sprintf("pong --> client_group(ping%d)\n", i)] = false
		}

		for _, v := range result {
			expected[v] = true
		}

		for k, v := range expected {
			if v == false {
				t.Errorf("expected received msg: %s", k)
			}
		}

		/*j := 0
		for _, v := range simulation.servers {
			fmt.Println(v.serverName, ", -->", v.workLoadCnt)
			j += v.workLoadCnt
		}
		fmt.Println(j)*/
	}
}

func BenchmarkClientGroupServerDistribution(b *testing.B) {
	result := make(map[string]int)
	RegisterProtocol("TestProto", &TestProto{})
	simulation := &TestClientGroupSimulation{
		protoName:         "TestProto",
		serverChanSize:    1024,
		serverMaxConn:     1024,
		serverCnt:         10,
		clientMsgCnt:      1000,
		clientChanSizeCnt: 1024,
		receivedChan:      make(chan interface{}, 0),
	}

	for i := 0; i < b.N; i++ {
		_, err := simulation.simulate()
		if err != nil {
			b.Error(err)
		}
		simulation.stop()

		for _, v := range simulation.servers {
			result[v.serverName] = result[v.serverName] + 1
		}

	}

	b.Log(result)
}

//func init() {
//	flag.Set("alsologtostderr", "true")
//	flag.Set("log_dir", "false")
//}
