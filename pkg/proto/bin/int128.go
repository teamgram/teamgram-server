// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

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

// BigInt returns corresponding big.Int value.
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
