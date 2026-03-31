// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"encoding/binary"
	"io"
	"math"
)

var ErrUnexpectedEOF = io.ErrUnexpectedEOF

// Decoder 解码
type Decoder struct {
	buf []byte
	off int
}

// NewDecoder NewDecoder
func NewDecoder(b []byte) *Decoder {
	return &Decoder{buf: b}
}

// ResetTo sets internal buffer exactly to provided value.
func (d *Decoder) ResetTo(buf []byte) {
	d.Reset(buf)
}

// Reset replaces the decoder input and rewinds offset.
func (d *Decoder) Reset(buf []byte) {
	d.buf = buf
	d.off = 0
}

// Raw returns internal remaining byte slice.
func (d *Decoder) Raw() []byte {
	return d.buf[d.off:]
}

// RawRemaining returns internal remaining byte slice.
func (d *Decoder) RawRemaining() []byte {
	return d.Raw()
}

// Len returns remaining length of internal buffer.
func (d *Decoder) Len() int {
	return d.Remaining()
}

// Remaining returns remaining unread bytes.
func (d *Decoder) Remaining() int {
	return len(d.buf) - d.off
}

// Offset returns current read offset.
func (d *Decoder) Offset() int {
	return d.off
}

// Skip moves cursor for next n bytes.
func (d *Decoder) Skip(n int) error {
	if n < 0 || d.Remaining() < n {
		return ErrUnexpectedEOF
	}
	d.off += n
	return nil
}

// PeekClazzID returns next type id in Buffer, but does not consume it.
func (d *Decoder) PeekClazzID() (uint32, error) {
	if d.Remaining() < WordLen {
		return 0, ErrUnexpectedEOF
	}
	return binary.LittleEndian.Uint32(d.buf[d.off:]), nil
}

// PeekN returns n bytes from Buffer to target, but does not consume it.
func (d *Decoder) PeekN(target []byte, n int) error {
	if n < 0 || len(target) < n || d.Remaining() < n {
		return ErrUnexpectedEOF
	}
	copy(target, d.buf[d.off:d.off+n])
	return nil
}

// ClazzID decodes type id from Buffer.
func (d *Decoder) ClazzID() (uint32, error) {
	return d.Uint32()
}

// Uint32 decodes unsigned 32-bit integer from Buffer.
func (d *Decoder) Uint32() (uint32, error) {
	if d.Remaining() < WordLen {
		return 0, ErrUnexpectedEOF
	}
	v := binary.LittleEndian.Uint32(d.buf[d.off:])
	d.off += WordLen
	return v, nil
}

// Int64 decodes 64-bit signed integer from Buffer.
func (d *Decoder) Int64() (int64, error) {
	return d.Long()
}

// Uint64 decodes 64-bit unsigned integer from Buffer.
func (d *Decoder) Uint64() (uint64, error) {
	const size = WordLen * 2
	if d.Remaining() < size {
		return 0, ErrUnexpectedEOF
	}
	v := binary.LittleEndian.Uint64(d.buf[d.off:])
	d.off += size
	return v, nil
}

// Ulong decodes 64-bit unsigned integer from Buffer.
func (d *Decoder) Ulong() (uint64, error) {
	return d.Uint64()
}

// Int32 decodes signed 32-bit integer from Buffer.
func (d *Decoder) Int32() (int32, error) {
	v, err := d.Uint32()
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

// ConsumeN consumes n bytes from buffer, writing them to target.
func (d *Decoder) ConsumeN(target []byte, n int) error {
	if err := d.PeekN(target, n); err != nil {
		return err
	}
	d.off += n
	return nil
}

// ConsumeClazzID decodes type id from Buffer. If id differs from provided,
// buffer will not be consumed.
func (d *Decoder) ConsumeClazzID(id uint32) error {
	v, err := d.PeekClazzID()
	if err != nil {
		return err
	}
	if v != id {
		return NewUnexpectedClazzID(id, v, d.off)
	}
	d.off += WordLen
	return nil
}

// VectorHeader decodes a TL vector header and item count.
func (d *Decoder) VectorHeader() (int32, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return 0, err
	}
	n, err := d.Int32()
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, &InvalidLengthError{Type: "vector", Length: int(n), Offset: d.off - WordLen}
	}
	return n, nil
}

// String decodes string from Buffer.
func (d *Decoder) String() (string, error) {
	n, v, err := decodeString(d.buf[d.off:])
	if err != nil {
		return "", err
	}
	if d.Remaining() < n {
		return "", ErrUnexpectedEOF
	}
	d.off += n
	return v, nil
}

// Bytes decodes byte slice from Buffer.
func (d *Decoder) Bytes() ([]byte, error) {
	v, err := d.BytesView()
	if err != nil {
		return nil, err
	}
	return append([]byte(nil), v...), nil
}

// BytesView decodes a borrowed byte slice from Buffer without copying.
func (d *Decoder) BytesView() ([]byte, error) {
	n, v, err := decodeBytes(d.buf[d.off:])
	if err != nil {
		return nil, err
	}
	if d.Remaining() < n {
		return nil, ErrUnexpectedEOF
	}
	d.off += n
	return v, nil
}

// Int decodes integer from Buffer.
func (d *Decoder) Int() (int, error) {
	v, err := d.Int32()
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// Double decodes 64-bit floating point from Buffer.
func (d *Decoder) Double() (float64, error) {
	v, err := d.Long()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(uint64(v)), nil
}

// Int53 decodes 53-bit signed integer from Buffer.
func (d *Decoder) Int53() (int64, error) {
	return d.Long()
}

// Long decodes 64-bit signed integer from Buffer.
func (d *Decoder) Long() (int64, error) {
	v, err := d.Uint64()
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

// Int128 decodes 128-bit signed integer from Buffer.
func (d *Decoder) Int128() (Int128, error) {
	v := Int128{}
	size := len(v)
	if d.Remaining() < size {
		return Int128{}, ErrUnexpectedEOF
	}
	copy(v[:], d.buf[d.off:d.off+size])
	d.off += size
	return v, nil
}

// Int256 decodes 256-bit signed integer from Buffer.
func (d *Decoder) Int256() (Int256, error) {
	v := Int256{}
	size := len(v)
	if d.Remaining() < size {
		return Int256{}, ErrUnexpectedEOF
	}
	copy(v[:], d.buf[d.off:d.off+size])
	d.off += size
	return v, nil
}
