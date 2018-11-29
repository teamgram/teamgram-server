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

package server

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"github.com/nebula-chat/chatengine/pkg/net2/codec"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	kStateUnknown = iota
	kStateConnected
	kStateData
)

//func init() {
//	net2.RegisterProtocol("multi_proto", NewMultiProtoProxy())
//}

type MultiProtoProxy struct {
	httpServer   *http.Server
	httpListener *HttpListener
}

func NewMultiProtoProxy() *MultiProtoProxy {
	proto := &MultiProtoProxy{
		httpServer: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			ConnState:      ConnStateListener,
			Handler:        nil,
		},
		httpListener: &HttpListener{
			// helper:       listener,
			closeChan:  make(chan struct{}),
			acceptChan: make(chan net.Conn, 1000),
		},
	}

	http.HandleFunc("/apiw1", onMTProtoHttpApiw1)
	//http.HandleFunc("/apiw1", func(w http.ResponseWriter, r *http.Request) {
	//})

	go proto.httpServer.Serve(proto.httpListener)
	return proto
}

func (m *MultiProtoProxy) NewCodec(rw io.ReadWriter) (net2.Codec, error) {
	codec := &MultiProtoProxyCodec{
		conn:  rw.(net.Conn),
		State: kStateConnected,
		proto: m,
	}

	return codec, nil
}

type MultiProtoProxyCodec struct {
	codecType int // codec type
	conn      net.Conn
	State     int
	codec     net2.Codec
	proto     *MultiProtoProxy
}

func (c *MultiProtoProxyCodec) Receive() (interface{}, error) {
	//if c.codec == nil {
	//	err := fmt.Errorf("codec is nil")
	//	glog.Error(err)
	//	return nil, err
	//}
	//
	if c.State == kStateConnected {
		conn := NewTcpConnWrapper(c.conn)
		b, e := conn.Peek(4)
		if e != nil {
			fmt.Println(string(b))
			// conn2.Close()
			return nil, e
		}
		fmt.Println(string(b))

		if !bytes.Equal(b, []byte("POST")) {
			// conn2.Close()
			// return
			c.codec, _ = codec.NewLengthBasedFrame(1024).NewCodec(conn)
		} else {
			c.conn = conn
			c.proto.httpListener.acceptChan <- conn

			//go func() {
			//	select {
			//	case c.proto.httpListener.acceptChan <- conn:
			//	// case <-c.proto.httpListener.closeChan:
			//	}
			//	glog.Info("close conn...")
			//}()

			c.codec = NewMultiProtoHttpCodec(conn)
		}
		c.State = kStateData
	}

	return c.codec.Receive()
}

func (c *MultiProtoProxyCodec) Send(msg interface{}) error {
	if c.codec == nil {
		err := fmt.Errorf("codec is nil")
		glog.Error(err)
		return err
	}
	return c.codec.Send(msg)
}

func (c *MultiProtoProxyCodec) Close() error {
	glog.Info("Close")
	if c.codec == nil {
		err := fmt.Errorf("codec is nil")
		glog.Error(err)
		return err
	}
	return c.codec.Close()
}
