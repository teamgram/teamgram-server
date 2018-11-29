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

package mtproto

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net"
	"net/http"
	// "time"
	"github.com/nebula-chat/chatengine/pkg/net2"
	// "strings"
	"bytes"
	"time"
)

type MTProtoHttpProxyCodec struct {
	// conn  *MTProtoHttpProxyConn
	conn net.Conn
}

func NewMTProtoHttpProxyCodec(conn net.Conn) *MTProtoHttpProxyCodec {
	// return conn.tcpConn.SetReadDeadline(time.Now().Add(tcpHeartbeat * 2))
	conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	// .(*net2.BufferedConn).Conn.(*net.TCPConn).SetReadDeadline(time.Now().Add())
	return &MTProtoHttpProxyCodec{
		conn: conn,
	}
}

func (c *MTProtoHttpProxyCodec) Receive() (interface{}, error) {
	req, err := http.ReadRequest(c.conn.(*net2.BufferedConn).BufioReader())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if len(body) < 8 {
		err = fmt.Errorf("not enough uint64 len error - %d", len(body))
		glog.Error(err)
		return nil, err
	}

	authKeyId := int64(binary.LittleEndian.Uint64(body))
	msg := NewMTPRawMessage(authKeyId, 0, TRANSPORT_HTTP)
	err = msg.Decode(body)
	if err != nil {
		glog.Error(err)
		// conn.Close()
		return nil, err
	}

	return msg, nil
}

func (c *MTProtoHttpProxyCodec) Send(msg interface{}) error {
	// SendToHttpReply(msg, w)
	message, ok := msg.(*MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		glog.Error(err)
		// conn.Close()
		return err
	}

	b := message.Encode()

	rsp := http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    &http.Request{Method: "POST"},
		Header: http.Header{
			"Access-Control-Allow-Headers": {"origin, content-type"},
			"Access-Control-Allow-Methods": {"POST, OPTIONS"},
			"Access-Control-Allow-Origin":  {"*"},
			"Access-Control-Max-Age":       {"1728000"},
			"Cache-control":                {"no-store"},
			"Connection":                   {"keep-alive"},
			"Content-type":                 {"application/octet-stream"},
			"Pragma":                       {"no-cache"},
			"Strict-Transport-Security":    {"max-age=15768000"},
		},
		ContentLength: int64(len(b)),
		Body:          ioutil.NopCloser(bytes.NewReader(b)),
		Close:         false,
	}

	err := rsp.Write(c.conn)
	if err != nil {
		glog.Error(err)
	}

	return err
}

func (c *MTProtoHttpProxyCodec) Close() error {
	return c.conn.Close()
}
