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

package util

import (
	"encoding/hex"
	"github.com/nebula-chat/chatengine/pkg/hack"
)

func WriteString(x *BufferOutput, s string) {
	WriteBytes(x, hack.Bytes(s))
}

func WriteBytes(x *BufferOutput, b []byte) {
	x.UInt32(uint32(len(b)))
	x.Bytes(b)
}

func ReadString(dbuf *BufferInput) (s string, err error) {
	var b []byte
	if b, err = ReadBytes(dbuf); err != nil {
		s = hack.String(b)
	}
	return
}

func ReadBytes(dbuf *BufferInput) (b []byte, err error) {
	var n = dbuf.UInt32()
	b = dbuf.Bytes(int(n))
	err = dbuf.Error()
	return
}

func DumpSize(size int, buf []byte) string {
	if size > len(buf) {
		size = len(buf)
	}
	return hex.Dump(buf[:size])
}

func Dump(buf []byte) string {
	return DumpSize(128, buf)
}

func HexDumpSize(size int, buf []byte) string {
	if size > len(buf) {
		size = len(buf)
	}
	return hex.EncodeToString(buf[:size])
}

func HexDump(buf []byte) string {
	return HexDumpSize(128, buf)
}
