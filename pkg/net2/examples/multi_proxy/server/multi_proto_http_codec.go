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
	"fmt"
	"github.com/golang/glog"
	"net"
	"net/http"
	"os"
	"reflect"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////////
var _ net.Listener = &HttpListener{}

type HttpListener struct {
	base       net.Listener
	acceptChan chan net.Conn
	closed     bool
	closeOnce  sync.Once
	closeChan  chan struct{}
}

func (l *HttpListener) Addr() net.Addr {
	return l.base.Addr()
}

func (l *HttpListener) Close() error {
	glog.Info("HttpListener.Close()")
	l.closeOnce.Do(func() {
		l.closed = true
		close(l.closeChan)
	})
	return l.base.Close()
}

func (l *HttpListener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.acceptChan:
		return conn, nil
	case <-l.closeChan:
	}
	return nil, os.ErrInvalid
}

///////////////////////////////////////////////////////////////////////////////////
func onMTProtoHttpApiw1(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)

	connPtr := writerToConnPtr(w)
	connMutex.Lock()
	defer connMutex.Unlock()

	// Requests can access connections by pointer from the responseWriter object
	conn, ok := conns[connPtr]
	if !ok {
		glog.Info("error: no matching connection found")
		return
	} else {
		// defer conn.Close()
		conn.(*TcpConnWrapper).RecvChan <- req
		// _ = conn
		// _, _ = w.Write([]byte(req.RequestURI + "\n"))
		msgData, _ := <-conn.(*TcpConnWrapper).SendChan
		glog.Info(msgData)
		w.Write([]byte(msgData.(*http.Request).RequestURI + "\n"))
		// close(conn.(*TcpConnWrapper).RecvChan)
	}
}

// Connection array indexed by connection address
var conns = make(map[uintptr]net.Conn)
var connMutex = sync.Mutex{}

// writerToConnPrt converts an http.ResponseWriter to a pointer for indexing
func writerToConnPtr(w http.ResponseWriter) uintptr {
	ptrVal := reflect.ValueOf(w)
	val := reflect.Indirect(ptrVal)

	// http.conn
	valconn := val.FieldByName("conn")
	val1 := reflect.Indirect(valconn)

	// net.TCPConn
	ptrRwc := val1.FieldByName("rwc").Elem()
	rwc := reflect.Indirect(ptrRwc)

	// net.Conn
	val1conn := rwc.FieldByName("conn")
	val2 := reflect.Indirect(val1conn)

	return val2.Addr().Pointer()
}

// connToPtr converts a net.Conn into a pointer for indexing
func connToPtr(c net.Conn) uintptr {
	ptrVal := reflect.ValueOf(c)
	return ptrVal.Pointer()
}

// ConnStateListener bound to server and maintains a list of connections by pointer
func ConnStateListener(c net.Conn, cs http.ConnState) {
	connPtr := connToPtr(c)
	connMutex.Lock()
	defer connMutex.Unlock()

	switch cs {
	case http.StateNew:
		glog.Infof("CONN Opened: 0x%x\n", connPtr)
		conns[connPtr] = c

	case http.StateClosed:
		glog.Infof("CONN Closed: 0x%x\n", connPtr)
		delete(conns, connPtr)
	}
}

///////////////////////////////////////////////////////////////////////////////////

type MultiProtoHttpCodec struct {
	conn *TcpConnWrapper
}

func NewMultiProtoHttpCodec(conn *TcpConnWrapper) *MultiProtoHttpCodec {
	return &MultiProtoHttpCodec{conn}
}

func (m *MultiProtoHttpCodec) Receive() (interface{}, error) {
	msgData, ok := <-m.conn.RecvChan
	glog.Info(msgData, ",  ok: ", ok)
	// fmt.Errorf()
	if !ok {
		return nil, fmt.Errorf("chan closed")
	}
	return msgData, nil
}

func (m *MultiProtoHttpCodec) Send(msg interface{}) error {
	m.conn.SendChan <- msg
	return nil
}

func (m *MultiProtoHttpCodec) Close() error {
	return m.conn.Close()
}
