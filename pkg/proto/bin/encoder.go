// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"math"
	"math/big"

	"github.com/valyala/bytebufferpool"
)

var (
	pool = bytebufferpool.Pool{}
)

//type EncodeOptions struct {
//	Cap int
//}
//
//func NewEncodeOptions() *EncodeOptions {
//	return &EncodeOptions{
//		Cap: 100,
//	}
//}
//
//type EncodeOption func(*EncodeOptions)
//
//func EncodeWithCap(cap int) EncodeOption {
//	return func(o *EncodeOptions) {
//		o.Cap = cap
//	}
//}

type Encoder struct {
	w *bytebufferpool.ByteBuffer
}

func NewEncoder() *Encoder {
	return &Encoder{
		w: pool.Get(),
	}
}

// Clone returns new copy of buffer.
func (e *Encoder) Clone() []byte {
	return append([]byte{}, e.w.Bytes()...)
}

// Bytes Bytes
func (e *Encoder) Bytes() []byte {
	return e.w.Bytes()
}

// Len Len
func (e *Encoder) Len() int {
	return e.w.Len()
}

// PutClazzID serializes type definition id, like a8509bda.
func (e *Encoder) PutClazzID(id uint32) {
	e.PutUint32(id)
}

// Put appends raw bytes to buffer.
//
// Buffer does not retain raw.
func (e *Encoder) Put(raw []byte) {
	_, _ = e.w.Write(raw)
}

// PutString serializes bare string.
func (e *Encoder) PutString(s string) {
	e.encodeString(s)
}

// PutBytes serializes bare byte string.
func (e *Encoder) PutBytes(v []byte) {
	e.encodeBytes(v)
}

//// PutVectorHeader serializes vector header with provided length.
//func (e *Encoder) PutVectorHeader(length int) {
//	e.PutClazzID(ClazzID_vector)
//	e.PutInt32(int32(length))
//}

// PutInt serializes v as signed 32-bit integer.
//
// If v is bigger than 32-bit, `behavior` is undefined.
func (e *Encoder) PutInt(v int) {
	e.PutUint32(uint32(v))
}

//// PutBool serializes bare boolean.
//func (e *Encoder) PutBool(v bool) {
//	switch v {
//	case true:
//		e.PutClazzID(ClazzID_boolTrue)
//	case false:
//		e.PutClazzID(ClazzID_boolFalse)
//	}
//}

// PutUint16 serializes unsigned 16-bit integer.
func (e *Encoder) PutUint16(v uint16) {
	_, _ = e.w.Write([]byte{
		byte(v),
		byte(v >> 8)})
}

// PutInt32 serializes signed 32-bit integer.
func (e *Encoder) PutInt32(v int32) {
	e.PutUint32(uint32(v))
}

// PutUint32 serializes unsigned 32-bit integer.
func (e *Encoder) PutUint32(v uint32) {
	_, _ = e.w.Write([]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	})
}

// PutInt53 serializes v as signed integer.
func (e *Encoder) PutInt53(v int64) {
	e.PutLong(v)
}

// PutLong serializes v as signed integer.
func (e *Encoder) PutLong(v int64) {
	e.PutUint64(uint64(v))
}

// PutInt64 serializes v as signed integer.
func (e *Encoder) PutInt64(v int64) {
	e.PutUint64(uint64(v))
}

// PutUint64 serializes v as unsigned 64-bit integer.
func (e *Encoder) PutUint64(v uint64) {
	_, _ = e.w.Write([]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	})
}

// PutDouble serializes v as 64-bit floating point.
func (e *Encoder) PutDouble(v float64) {
	e.PutUint64(math.Float64bits(v))
}

// PutInt128 serializes v as 128-bit signed integer.
func (e *Encoder) PutInt128(v Int128) {
	_, _ = e.w.Write(v[:])
}

// PutInt256 serializes v as 256-bit signed integer.
func (e *Encoder) PutInt256(v Int256) {
	_, _ = e.w.Write(v[:])
}

func (e *Encoder) BigInt(s *big.Int) {
	e.PutBytes(s.Bytes())
}

func (e *Encoder) Reset() {
	e.w.Reset()
}

func (e *Encoder) End() {
	bytebufferpool.Put(e.w)
}
