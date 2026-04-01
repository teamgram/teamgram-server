// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"math"
	"math/big"
	"sync"
)

const (
	defaultEncoderCap = 256
	maxPooledCap      = 64 << 10
)

var (
	encoderPool = sync.Pool{
		New: func() any {
			return &Encoder{buf: make([]byte, 0, defaultEncoderCap)}
		},
	}
	zeroPad = [4]byte{}
)

type Encoder struct {
	buf      []byte
	released bool
}

func (e *Encoder) ensureWritable() {
	if e == nil || e.released {
		panic("bin: encoder used after release")
	}
}

func AcquireEncoder() *Encoder {
	return AcquireEncoderCap(defaultEncoderCap)
}

func AcquireEncoderCap(capHint int) *Encoder {
	e := encoderPool.Get().(*Encoder)
	if capHint <= 0 {
		capHint = defaultEncoderCap
	}
	if cap(e.buf) < capHint {
		e.buf = make([]byte, 0, capHint)
	} else {
		e.buf = e.buf[:0]
	}
	e.released = false
	return e
}

func NewEncoder() *Encoder {
	return AcquireEncoder()
}

// Clone returns new copy of buffer.
func (e *Encoder) Clone() []byte {
	e.ensureWritable()
	return append([]byte(nil), e.buf...)
}

// Bytes returns the internal encoded bytes view.
func (e *Encoder) Bytes() []byte {
	e.ensureWritable()
	return e.buf
}

// Len returns encoded bytes length.
func (e *Encoder) Len() int {
	e.ensureWritable()
	return len(e.buf)
}

// Cap returns current backing buffer capacity.
func (e *Encoder) Cap() int {
	e.ensureWritable()
	return cap(e.buf)
}

// Grow reserves at least n bytes for future writes.
func (e *Encoder) Grow(n int) {
	e.ensureWritable()
	if n <= 0 {
		return
	}
	required := len(e.buf) + n
	if required <= cap(e.buf) {
		return
	}
	buf := make([]byte, len(e.buf), required)
	copy(buf, e.buf)
	e.buf = buf
}

// PutClazzID serializes type definition id, like a8509bda.
func (e *Encoder) PutClazzID(id uint32) {
	e.PutUint32(id)
}

// Put appends raw bytes to buffer.
func (e *Encoder) Put(raw []byte) {
	e.ensureWritable()
	e.buf = append(e.buf, raw...)
}

// PutRaw appends raw bytes to buffer.
func (e *Encoder) PutRaw(raw []byte) {
	e.Put(raw)
}

// PutString serializes bare string.
func (e *Encoder) PutString(s string) {
	e.ensureWritable()
	e.encodeString(s)
}

// PutBytes serializes bare byte string.
func (e *Encoder) PutBytes(v []byte) {
	e.ensureWritable()
	e.encodeBytes(v)
}

// PutVectorHeader serializes a TL vector header and item count.
func (e *Encoder) PutVectorHeader(length int32) {
	e.PutClazzID(ClazzID_vector)
	e.PutInt32(length)
}

// PutInt serializes v as a signed 32-bit integer.
// Values outside the int32 range are truncated.
// Prefer PutInt32 or PutLong when the width must be explicit.
func (e *Encoder) PutInt(v int) {
	e.PutInt32(int32(v))
}

// PutUint serializes v as unsigned 32-bit integer.
func (e *Encoder) PutUint(v uint32) {
	e.PutUint32(v)
}

// PutUint16 serializes unsigned 16-bit integer.
func (e *Encoder) PutUint16(v uint16) {
	e.ensureWritable()
	e.buf = append(e.buf, byte(v), byte(v>>8))
}

// PutInt32 serializes signed 32-bit integer.
func (e *Encoder) PutInt32(v int32) {
	e.PutUint32(uint32(v))
}

// PutUint32 serializes unsigned 32-bit integer.
func (e *Encoder) PutUint32(v uint32) {
	e.ensureWritable()
	e.buf = append(e.buf,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)
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

// PutUlong serializes v as unsigned 64-bit integer.
func (e *Encoder) PutUlong(v uint64) {
	e.PutUint64(v)
}

// PutUint64 serializes unsigned 64-bit integer.
func (e *Encoder) PutUint64(v uint64) {
	e.ensureWritable()
	e.buf = append(e.buf,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56),
	)
}

// PutDouble serializes v as 64-bit floating point.
func (e *Encoder) PutDouble(v float64) {
	e.PutUint64(math.Float64bits(v))
}

// PutInt128 serializes v as 128-bit signed integer.
func (e *Encoder) PutInt128(v Int128) {
	e.ensureWritable()
	e.buf = append(e.buf, v[:]...)
}

// PutInt256 serializes v as 256-bit signed integer.
func (e *Encoder) PutInt256(v Int256) {
	e.ensureWritable()
	e.buf = append(e.buf, v[:]...)
}

// BigInt serializes s.Bytes() as a TL bytes value.
// It does not encode a dedicated TL big integer type.
// The resulting bytes use big.Int's big-endian magnitude representation.
func (e *Encoder) BigInt(s *big.Int) {
	e.ensureWritable()
	e.PutBytes(s.Bytes())
}

func (e *Encoder) Reset() {
	e.ensureWritable()
	e.buf = e.buf[:0]
}

func (e *Encoder) Release() {
	if e == nil || e.released {
		return
	}
	if cap(e.buf) > maxPooledCap {
		e.buf = make([]byte, 0, defaultEncoderCap)
	} else {
		e.buf = e.buf[:0]
	}
	e.buf = nil
	e.released = true
	encoderPool.Put(e)
}

func (e *Encoder) End() {
	e.Release()
}
