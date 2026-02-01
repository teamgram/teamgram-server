// Copyright 2024 Teamgram Authors
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

// Decoder 解码
type Decoder struct {
	buf []byte
}

// NewDecoder NewDecoder
func NewDecoder(b []byte) *Decoder {
	return &Decoder{
		buf: b,
	}
}

// ResetTo sets internal buffer exactly to provided value.
//
// Buffer will retain buf, so user should not modify or read it
// concurrently.
func (d *Decoder) ResetTo(buf []byte) {
	d.buf = buf
}

// Raw returns internal byte slice.
func (d *Decoder) Raw() []byte {
	return d.buf
}

// Len returns length of internal buffer.
func (d *Decoder) Len() int {
	return len(d.buf)
}

// Skip moves cursor for next n bytes.
func (d *Decoder) Skip(n int) {
	d.buf = d.buf[n:]
}

// PeekClazzID returns next type id in Buffer, but does not consume it.
func (d *Decoder) PeekClazzID() (uint32, error) {
	if len(d.buf) < WordLen {
		return 0, io.ErrUnexpectedEOF
	}
	v := binary.LittleEndian.Uint32(d.buf)
	return v, nil
}

// PeekN returns n bytes from Buffer to target, but does not consume it.
//
// Returns io.ErrUnexpectedEOF if buffer contains less that n bytes.
// Expects that len(target) >= n.
func (d *Decoder) PeekN(target []byte, n int) error {
	if len(d.buf) < n {
		return io.ErrUnexpectedEOF
	}
	copy(target, d.buf[:n])
	return nil
}

// ClazzID decodes type id from Buffer.
func (d *Decoder) ClazzID() (uint32, error) {
	return d.Uint32()
}

// Uint32 decodes unsigned 32-bit integer from Buffer.
func (d *Decoder) Uint32() (uint32, error) {
	if len(d.buf) < WordLen {
		return 0, io.ErrUnexpectedEOF
	}
	v := binary.LittleEndian.Uint32(d.buf)
	d.buf = d.buf[WordLen:]
	return v, nil
}

// Int64 decodes 64-bit signed integer from Buffer.
func (d *Decoder) Int64() (int64, error) {
	v, err := d.Uint64()
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

// Uint64 decodes 64-bit unsigned integer from Buffer.
func (d *Decoder) Uint64() (uint64, error) {
	const size = WordLen * 2
	if len(d.buf) < size {
		return 0, io.ErrUnexpectedEOF
	}
	v := binary.LittleEndian.Uint64(d.buf)
	d.buf = d.buf[size:]
	return v, nil
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
//
// Returns io.ErrUnexpectedEOF if buffer contains less that n bytes.
// Expects that len(target) >= n.
func (d *Decoder) ConsumeN(target []byte, n int) error {
	if err := d.PeekN(target, n); err != nil {
		return err
	}
	d.buf = d.buf[n:]
	return nil
}

//// Bool decodes bare boolean from Buffer.
//func (d *Decoder) Bool() (bool, error) {
//	v, err := d.PeekClazzID()
//	if err != nil {
//		return false, err
//	}
//	switch v {
//	case ClazzID_boolTrue:
//		d.buf = d.buf[WordLen:]
//		return true, nil
//	case ClazzID_boolFalse:
//		d.buf = d.buf[WordLen:]
//		return false, nil
//	default:
//		return false, NewUnexpectedClazzID(v)
//	}
//}

// ConsumeClazzID decodes type id from Buffer. If id differs from provided,
// the *UnexpectedIDErr{ID: gotID} will be returned and buffer will be
// not consumed.
func (d *Decoder) ConsumeClazzID(id uint32) error {
	v, err := d.PeekClazzID()
	if err != nil {
		return err
	}
	if v != id {
		return NewUnexpectedClazzID(v)
	}
	d.buf = d.buf[WordLen:]
	return nil
}

//// VectorHeader decodes vector length from Buffer.
//func (d *Decoder) VectorHeader() (int, error) {
//	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
//		return 0, err
//	}
//	n, err := d.Int()
//	if err != nil {
//		return 0, err
//	}
//	if n < 0 {
//		return 0, &InvalidLengthError{
//			Length: int(n),
//			Where:  "vector",
//		}
//	}
//	return int(n), nil
//}

// String decodes string from Buffer.
func (d *Decoder) String() (string, error) {
	n, v, err := decodeString(d.buf)
	if err != nil {
		return "", err
	}
	if len(d.buf) < n {
		return "", io.ErrUnexpectedEOF
	}
	d.buf = d.buf[n:]
	return v, nil
}

// Bytes decodes byte slice from Buffer.
//
// NB: returning value is a copy, it's safe to modify it.
func (d *Decoder) Bytes() ([]byte, error) {
	n, v, err := decodeBytes(d.buf)
	if err != nil {
		return nil, err
	}
	if len(d.buf) < n {
		return nil, io.ErrUnexpectedEOF
	}
	d.buf = d.buf[n:]
	return append([]byte(nil), v...), nil
}

// Int decodes integer from Buffer.
func (d *Decoder) Int() (int, error) {
	v, err := d.Uint32()
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
	if len(d.buf) < size {
		return Int128{}, io.ErrUnexpectedEOF
	}
	copy(v[:size], d.buf[:size])
	d.buf = d.buf[size:]
	return v, nil
}

// Int256 decodes 256-bit signed integer from Buffer.
func (d *Decoder) Int256() (Int256, error) {
	v := Int256{}
	size := len(v)
	if len(d.buf) < size {
		return Int256{}, io.ErrUnexpectedEOF
	}
	copy(v[:size], d.buf[:size])
	d.buf = d.buf[size:]
	return v, nil
}
