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
	"net"
)

type BufferedConn struct {
	r        *bufio.Reader
	net.Conn // So that most methods are embedded
}

func NewBufferedConn(c net.Conn) *BufferedConn {
	return &BufferedConn{bufio.NewReader(c), c}
}

func NewBufferedConnSize(c net.Conn, n int) *BufferedConn {
	return &BufferedConn{bufio.NewReaderSize(c, n), c}
}

func (b *BufferedConn) Peek(n int) ([]byte, error) {
	return b.r.Peek(n)
}

func (b *BufferedConn) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b *BufferedConn) Discard(n int) (int, error) {
	return b.r.Discard(n)
}

func (b *BufferedConn) BufioReader() *bufio.Reader {
	return b.r
}
