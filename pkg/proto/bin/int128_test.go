// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt128_Encode(t *testing.T) {
	a := require.New(t)

	v := Int128{1, 2, 3, 0, 134, 45}
	x := NewEncoder()
	defer x.End()
	v.Encode(x, 0)

	var decoded Int128
	d := NewDecoder(x.Bytes())

	a.NoError(decoded.Decode(d))
	a.Equal(v, decoded)
	a.Error(decoded.Decode(NewDecoder(nil)))
}

func BenchmarkBuffer_PutInt128(b *testing.B) {
	v := Int128{10, 15}
	x := NewEncoder()
	defer x.End()
	for i := 0; i < b.N; i++ {
		x.PutInt128(v)
		x.Reset()
	}
}
