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

func TestInt256_Encode(t *testing.T) {
	a := require.New(t)

	v := Int256{4, 3, 1, 2}
	x := NewEncoder()
	defer x.End()

	v.Encode(x, 0)

	var decoded Int256
	d := NewDecoder(x.Bytes())
	a.NoError(decoded.Decode(d))
	a.Equal(v, decoded)

	a.Error(decoded.Decode(NewDecoder(nil)))
}

func BenchmarkBuffer_PutInt256(b *testing.B) {
	b.ReportAllocs()
	v := Int256{1, 4, 4, 6}

	x := NewEncoder()
	defer x.End()
	for i := 0; i < b.N; i++ {
		x.PutInt256(v)
		x.Reset()
	}
}
