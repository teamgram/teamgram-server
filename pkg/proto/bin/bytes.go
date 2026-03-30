// Copyright 2024 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"io"
)

// encodeBytes is same as encodeString, but for bytes.
func (e *Encoder) encodeBytes(v []byte) {
	l := len(v)
	if l <= maxSmallStringLength {
		e.buf = append(e.buf, byte(l))
		e.buf = append(e.buf, v...)
		currentLen := l + 1
		pad := nearestPaddedValueLength(currentLen) - currentLen
		e.buf = append(e.buf, zeroPad[:pad]...)
	} else {
		e.buf = append(e.buf, firstLongStringByte, byte(l), byte(l>>8), byte(l>>16))
		e.buf = append(e.buf, v...)
		currentLen := l + 4
		pad := nearestPaddedValueLength(currentLen) - currentLen
		e.buf = append(e.buf, zeroPad[:pad]...)
	}
}

// decodeBytes is same as decodeString, but for bytes.
//
// NB: v is slice of b.
func decodeBytes(b []byte) (n int, v []byte, err error) {
	if len(b) == 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}
	if b[0] == firstLongStringByte {
		if len(b) < 4 {
			return 0, nil, io.ErrUnexpectedEOF
		}
		strLen := uint32(b[1]) | uint32(b[2])<<8 | uint32(b[3])<<16
		if len(b) < (int(strLen) + 4) {
			return 0, nil, io.ErrUnexpectedEOF
		}
		return nearestPaddedValueLength(int(strLen) + 4), b[4 : strLen+4], nil
	}
	strLen := int(b[0])
	if len(b) < (strLen + 1) {
		return 0, nil, io.ErrUnexpectedEOF
	}
	if strLen > maxSmallStringLength {
		return 0, nil, &InvalidLengthError{
			Type:   "bytes",
			Length: strLen,
		}
	}
	return nearestPaddedValueLength(strLen + 1), b[1 : strLen+1], nil
}
