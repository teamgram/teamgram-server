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
	"net"
	"testing"
	"time"
)

type TestTcpPingClient struct {
	client   *TcpClient
	ready    chan bool
	receiver chan interface{}
}

func NewTestClient(RemoteName, protoName, remoteAddress string, receiver chan interface{}, chanSize int) *TestTcpPingClient {
	c := &TestTcpPingClient{}
	c.client = NewTcpClient(RemoteName, chanSize, protoName, remoteAddress, c)
	c.ready = make(chan bool)
	c.receiver = receiver
	return c
}

func (c *TestTcpPingClient) getName() string {
	return c.client.GetRemoteName()
}

func (c *TestTcpPingClient) Serve() {
	c.client.Serve()
}

func (c *TestTcpPingClient) Stop() {
	c.client.Stop()
}

func (c *TestTcpPingClient) Send(msg interface{}) error {
	return c.client.Send(msg)
}

func (c *TestTcpPingClient) OnNewClient(client *TcpClient) {
	glog.Infof("client OnNewConnection %v", client.GetRemoteName())
	c.ready <- true
}

func (c *TestTcpPingClient) OnClientDataArrived(client *TcpClient, msg interface{}) error {
	glog.Infof("client OnClientDataArrived - client: %s, receive data: %v", client.GetRemoteName(), msg)
	c.receiver <- msg
	return nil
}

func (c *TestTcpPingClient) OnClientClosed(client *TcpClient) {
	glog.Infof("client OnConnectionClosed %s", client.GetRemoteName())
}

func (c *TestTcpPingClient) startTimer() {
	c.client.SetTimer(time.Second)
	c.client.StartTimer()
}

func (c *TestTcpPingClient) OnClientTimer(client *TcpClient) {
	fmt.Println("OnTimer")
	c.Send("timer_ping")
}

//////////////////////////////////////////////////////////////////////////////////////////
type TestTcpClientSimulation struct {
	protoName string
	// ----- server -----
	serverChanSize int
	serverMaxConn  int
	// ----- client -----
	clientNo          int
	clientMsgCnt      int
	clientChanSizeCnt int

	receivedChan chan interface{}

	server  *TestPingPongServer
	clients []*TestTcpPingClient
}

func (ts *TestTcpClientSimulation) simulate() (result []string, e error) {
	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		glog.Errorf("listen error: %v", err)
		return nil, err
	}

	ts.server = NewTestServer(listener, `TestServer0`, ts.protoName, ts.serverChanSize, ts.serverMaxConn)
	go ts.server.Serve()

	// clients
	var clients []*TestTcpPingClient
	for i := 0; i < ts.clientNo; i++ {
		client := NewTestClient(fmt.Sprintf(`client%d`, i), ts.protoName, listener.Addr().String(), ts.receivedChan, ts.clientChanSizeCnt)
		clients = append(clients, client)
		go client.Serve()
		<-client.ready
	}
	ts.clients = append(ts.clients, clients...)

	errChan := make(chan error)

	for _, c := range clients {
		for i := 0; i < ts.clientMsgCnt; i++ {
			go func(tc *TestTcpPingClient, errorChan chan error, n int) {
				if err = tc.Send(fmt.Sprintf("%s(ping%d)\n", tc.getName(), n)); err != nil {
					errorChan <- err
				}
			}(c, errChan, i)
		}
	}

	result = make([]string, 0)

	for cnt := 0; cnt < ts.clientNo*ts.clientMsgCnt; cnt++ {
		select {
		case msg, _ := <-ts.receivedChan:
			result = append(result, msg.(string))
		case m, _ := <-errChan:
			result = append(result, m.Error())
		}
	}

	return
}

func (ts *TestTcpClientSimulation) stop() {
	ts.server.Stop()

	for _, c := range ts.clients {
		c.Stop()
	}
}

//////////////////////////////////////////////////////////////////////////////////////////
// net2.tcpClient 																      net2.tcpServer
//     |     																	      |				|
//   --------------	   															-----------------	|
//  | net2.tcpCnn  |											  			   |   net2.tcpCnn   |	|
//   --------------   															-----------------	|
//   		|   															          |				|
//      -----------------------											   --------------------		|
//     |   chanSize (buffer)   |										  |  chanSize (buffer) |	|
//      -----------------------									           --------------------		|
//                |________________________net.TCPConn____________________________|					|
//																									|
//     .																		      				|
//     .																		     				|
//     .																		      				|
//																				                    |
// net2.tcpClient 																                    |
//     |     																	                    |
//   --------------	   															-----------------   |
//  | net2.tcpCnn  |											  			   |   net2.tcpCnn   |--
//   --------------   															-----------------
//   		|   															          |
//      -----------------------											   --------------------
//     |   chanSize (buffer)   |										  |  chanSize (buffer) |
//      -----------------------									           --------------------
//                |________________________net.TCPConn____________________________|
//
func TestClientServer(t *testing.T) {
	protoName := "TestProto"

	RegisterProtocol(protoName, &TestProto{})

	testTable := []struct {
		serverChanSize    int
		serverMaxConn     int
		clientNo          int
		clientMsgCnt      int
		clientChanSizeCnt int
	}{
		{
			// 1 server, 1 client
			clientNo:          1,
			clientMsgCnt:      1024,
			clientChanSizeCnt: 1024,
			serverChanSize:    1024,
			serverMaxConn:     1024,
		},
		{
			// 1 server, n client
			clientNo:          1000,
			clientMsgCnt:      1024,
			clientChanSizeCnt: 1024,
			serverChanSize:    1024,
			serverMaxConn:     1024,
		},
	}

	for _, data := range testTable {

		simulation := &TestTcpClientSimulation{
			protoName:         protoName,
			serverChanSize:    data.serverChanSize,
			serverMaxConn:     data.serverMaxConn,
			clientNo:          data.clientNo,
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
		for c := 0; c < simulation.clientNo; c++ {
			for i := 0; i < simulation.clientMsgCnt; i++ {
				expected[fmt.Sprintf("pong --> client%d(ping%d)\n", c, i)] = false
			}
		}

		for _, v := range result {
			expected[v] = true
		}

		for k, v := range expected {
			if v == false {
				t.Errorf("expected received msg: %s", k)
			}
		}
	}
}

// todo(yumcoder): timer test
