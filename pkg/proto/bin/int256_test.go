// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
