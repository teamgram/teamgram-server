// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/teamgram/marmota/pkg/net2"
	"github.com/teamgram/proto/mtproto"

	log "github.com/zeromicro/go-zero/core/logx"
)

type HttpProxyCodec struct {
	conn    net.Conn
	canSend bool
}

func NewMTProtoHttpProxyCodec(conn net.Conn) *HttpProxyCodec {
	return &HttpProxyCodec{
		conn:    conn,
		canSend: false,
	}
}

func (c *HttpProxyCodec) Receive() (interface{}, error) {
	c.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	req, err := http.ReadRequest(c.conn.(*net2.BufferedConn).BufioReader())
	if err != nil {
		log.Infof("Receive: error - %s, %s", c.conn, err.Error())
		log.Error(err.Error())
		return nil, err
	}

	log.Infof("Receive: %s", c.conn)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if len(body) < 8 {
		err = fmt.Errorf("not enough uint64 len error - %d", len(body))
		log.Error(err.Error())
		return nil, err
	}

	authKeyId := int64(binary.LittleEndian.Uint64(body))
	msg := mtproto.NewMTPRawMessage(authKeyId, 0, TRANSPORT_HTTP)
	err = msg.Decode(body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	c.canSend = true
	return msg, nil
}

func (c *HttpProxyCodec) Send(msg interface{}) error {
	message, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
		log.Error(err.Error())
		// conn.Close()
		return err
	}
	log.Infof("Send: %s", c.conn)

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
		log.Error(err.Error())
	}

	c.canSend = false
	return err
}

func (c *HttpProxyCodec) Close() error {
	c.canSend = false
	return c.conn.Close()
}

func (c *HttpProxyCodec) Context() interface{} {
	return ""
}
