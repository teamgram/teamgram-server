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
	"bufio"
	"github.com/golang/glog"
	"io"
)

type TestCodec struct {
	*bufio.Reader
	io.Writer
	io.Closer
	mt string
}

func (c *TestCodec) Send(msg interface{}) error {
	buf := []byte(msg.(string))
	if _, err := c.Writer.Write(buf); err != nil {
		return err
	}

	return nil
}

func (c *TestCodec) Receive() (interface{}, error) {
	line, err := c.Reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return line, err
}

func (c *TestCodec) Close() error {
	return c.Closer.Close()
}

func (c *TestCodec) ClearSendChan(ic <-chan interface{}) {
	glog.Info(`TestCodec ClearSendChan, `, ic)
}

//////////////////////////////////////////////////////////////////////////////////////////
type TestProto struct {
}

func (b *TestProto) NewCodec(rw io.ReadWriter) (cc Codec, err error) {
	c := &TestCodec{
		Reader: bufio.NewReader(rw),
		Writer: rw.(io.Writer),
		Closer: rw.(io.Closer),
	}
	return c, nil
}
