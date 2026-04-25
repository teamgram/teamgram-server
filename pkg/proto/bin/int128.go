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
	"encoding/hex"
	"math/big"
)

// Int128 represents signed 128-bit integer.
type Int128 [16]byte

// Decode implements bin.Decoder.
func (i *Int128) Decode(d *Decoder) error {
	v, err := d.Int128()
	if err != nil {
		return err
	}
	*i = v

	return nil
}

// Encode implements bin.Encoder.
func (i *Int128) Encode(x *Encoder, layer int) {
	x.PutInt128(*i)
}

// BigInt interprets the raw 16 bytes as a big-endian unsigned magnitude.
// This is only one possible interpretation of the stored bytes.
func (i *Int128) BigInt() *big.Int {
	return big.NewInt(0).SetBytes(i[:])
}

func (i *Int128) Zero() bool {
	if i == nil {
		return true
	}
	if !(*i == Int128{}) {
		return false
	}

	return true
}

func (i *Int128) ToHex() string {
	return hex.EncodeToString(i[:])
}
