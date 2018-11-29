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
	"encoding/binary"
	"math"
)

// TODO(@benqi): Add write offset data

type BufferOutput struct {
	buf []byte
}

func NewBufferOutput(cap int) *BufferOutput {
	return &BufferOutput{make([]byte, 0, cap)}
}

func (e *BufferOutput) Buf() []byte {
	return e.buf
}

func (e *BufferOutput) Len() int {
	return len(e.buf)
}

func (e *BufferOutput) Byte(s byte) {
	e.buf = append(e.buf, s)
}

func (e *BufferOutput) Int16(s int16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], uint16(s))
}

func (e *BufferOutput) UInt16(s uint16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], s)
}

func (e *BufferOutput) Int32(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *BufferOutput) UInt32(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *BufferOutput) Int64(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *BufferOutput) UInt64(s uint64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], s)
}

func (e *BufferOutput) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *BufferOutput) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *BufferOutput) ByteSize() int {
	return len(e.buf)
}
