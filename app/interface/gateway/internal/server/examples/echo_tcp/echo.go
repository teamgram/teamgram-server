// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on [%s] (multi-cores: %t, loops: %d)\n",
		srv.AddrsString(), srv.Multicore, srv.NumEventLoop)
	return
}

// OnOpened fires when a new connection has been opened.
// The parameter:c has information about the connection such as it's local and remote address.
// Parameter:out is the return value which is going to be sent back to the client.
func (es *echoServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("OnOpened: %s\n", c.DebugString())
	return
}

// OnClosed fires when a connection has been closed.
// The parameter:err is the last known connection error.
func (es *echoServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("OnClosed: %s", c.DebugString())
	return
}

func (es *echoServer) React(frame interface{}, c gnet.Conn) (out interface{}, action gnet.Action) {
	log.Printf("React: %s", frame)
	// Echo synchronously.
	out = frame
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port, port2 int
	var multicore, reuseport bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 10444, "--port 10444")
	flag.IntVar(&port2, "port2", 10445, "--port2 10445")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.BoolVar(&reuseport, "reuseport", false, "--reuseport true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo,
		[]string{fmt.Sprintf("tcp://:%d", port), fmt.Sprintf("tcp://:%d", port2)},
		gnet.WithLockOSThread(true),
		gnet.WithMulticore(multicore),
		gnet.WithReusePort(reuseport),
		gnet.WithSocketRecvBuffer(4096),
		gnet.WithSocketSendBuffer(4096)))
}
