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

package codec

import (
	"bufio"
	"github.com/nebula-chat/chatengine/pkg/net2"
	"io"
)

func init() {
	net2.RegisterProtocol("length_based_frame", NewLengthBasedFrame(kDefaultReadBufferSize))
}

const (
	kDefaultReadBufferSize = 1024
)

func NewLengthBasedFrame(readBuf int) net2.Protocol {
	if readBuf <= 0 {
		readBuf = kDefaultReadBufferSize
	}

	return &LengthBasedFrame{
		readBuf: readBuf,
	}
}

type LengthBasedFrame struct {
	readBuf int
}

func (b *LengthBasedFrame) NewCodec(rw io.ReadWriter) (cc net2.Codec, err error) {
	codec := new(LengthBasedFrameCodec)

	codec.stream.w = rw.(io.Writer)
	codec.stream.r = bufio.NewReaderSize(rw, b.readBuf)
	codec.stream.c = rw.(io.Closer)

	return codec, nil
}

type LengthBasedFrameStream struct {
	w io.Writer
	r *bufio.Reader
	c io.Closer
}

func (s *LengthBasedFrameStream) close() error {
	if s.c != nil {
		return s.c.Close()
	}
	return nil
}

type LengthBasedFrameCodec struct {
	stream LengthBasedFrameStream
}

func (c *LengthBasedFrameCodec) Send(msg interface{}) error {
	buf := []byte(msg.(string))

	if _, err := c.stream.w.Write(buf); err != nil {
		return err
	}

	return nil
}

func (c *LengthBasedFrameCodec) Receive() (interface{}, error) {
	line, err := c.stream.r.ReadString('\n')
	return line, err
}

func (c *LengthBasedFrameCodec) Close() error {
	return c.stream.close()
}
