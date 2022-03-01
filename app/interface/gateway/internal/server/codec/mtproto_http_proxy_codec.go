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

package codec

import (
	"github.com/panjf2000/gnet"
)

type HttpProxyCodec struct {
	canSend bool
}

func newMTProtoHttpProxyCodec() *HttpProxyCodec {
	return &HttpProxyCodec{
		canSend: false,
	}
}

// Encode encodes frames upon server responses into TCP stream.
func (c *HttpProxyCodec) Encode(conn gnet.Conn, msg interface{}) ([]byte, error) {
	//message, ok := msg.(*MTPRawMessage)
	//if !ok {
	//	err := fmt.Errorf("msg type error, only MTPRawMessage, msg: {%v}", msg)
	//	log.Error(err.Error())
	//	// conn.Close()
	//	return err
	//}
	//log.Debugf("Send: %s", c.conn)
	//
	//b := message.Encode()
	//
	//rsp := http.Response{
	//	StatusCode: 200,
	//	ProtoMajor: 1,
	//	ProtoMinor: 1,
	//	Request:    &http.Request{Method: "POST"},
	//	Header: http.Header{
	//		"Access-Control-Allow-Headers": {"origin, content-type"},
	//		"Access-Control-Allow-Methods": {"POST, OPTIONS"},
	//		"Access-Control-Allow-Origin":  {"*"},
	//		"Access-Control-Max-Age":       {"1728000"},
	//		"Cache-control":                {"no-store"},
	//		"Connection":                   {"keep-alive"},
	//		"Content-type":                 {"application/octet-stream"},
	//		"Pragma":                       {"no-cache"},
	//		"Strict-Transport-Security":    {"max-age=15768000"},
	//	},
	//	ContentLength: int64(len(b)),
	//	Body:          ioutil.NopCloser(bytes.NewReader(b)),
	//	Close:         false,
	//}
	//
	//err := rsp.Write(c.conn)
	//if err != nil {
	//	log.Error(err.Error())
	//}
	//
	//c.canSend = false
	//return err
	return nil, nil
}

// Decode decodes frames from TCP stream via specific implementation.
func (c *HttpProxyCodec) Decode(conn gnet.Conn) (interface{}, error) {

	//c.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	//req, err := http.ReadRequest(c.conn.(*net2.BufferedConn).BufioReader())
	//if err != nil {
	//	log.Debugf("Receive: error - %s, %s", c.conn, err.Error())
	//	log.Error(err.Error())
	//	return nil, err
	//}
	//
	//log.Debugf("Receive: %s", c.conn)
	//body, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	log.Error(err.Error())
	//	return nil, err
	//}
	//
	//if len(body) < 8 {
	//	err = fmt.Errorf("not enough uint64 len error - %d", len(body))
	//	log.Error(err.Error())
	//	return nil, err
	//}
	//
	//authKeyId := int64(binary.LittleEndian.Uint64(body))
	//msg := NewMTPRawMessage(authKeyId, 0, TRANSPORT_HTTP)
	//err = msg.Decode(body)
	//if err != nil {
	//	log.Error(err.Error())
	//	return nil, err
	//}
	//
	//c.canSend = true
	//return msg, nil
	return nil, nil
}

// Clone ...
func (c *HttpProxyCodec) Clone() gnet.ICodec {
	return c
}

// Release ...
func (c *HttpProxyCodec) Release() {
}
