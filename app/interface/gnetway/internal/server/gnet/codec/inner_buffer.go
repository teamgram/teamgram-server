// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package codec

import (
	"errors"
)

// innerBuffer is a light wrapper around a byte slice that supports
// sequential fixed‑length reads. It is used only inside MTProto codecs
// to avoid extra allocations when slicing data returned by Peek.
type (
	innerBuffer []byte
)

func (in *innerBuffer) readN(n int) (buf []byte, err error) {
	if n <= 0 {
		return nil, errors.New("zero or negative length is invalid")
	} else if n > len(*in) {
		return nil, errors.New("exceeding buffer length")
	}
	buf = (*in)[:n]
	*in = (*in)[n:]

	return
}
