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

// Int256 represents signed 256-bit integer.
type (
	Int256 [32]byte
)

// Decode implements bin.Decoder.
func (i *Int256) Decode(d *Decoder) error {
	v, err := d.Int256()
	if err != nil {
		return err
	}
	*i = v
	return nil
}

// Encode implements bin.Encoder.
func (i *Int256) Encode(x *Encoder, _ int) {
	x.PutInt256(*i)
}

// BigInt returns corresponding big.Int value.
func (i *Int256) BigInt() *big.Int {
	return big.NewInt(0).SetBytes(i[:])
}

func (i *Int256) Zero() bool {
	if i == nil {
		return true
	}
	if !(*i == Int256{}) {
		return false
	}

	return true
}

func (i *Int256) ToHex() string {
	return hex.EncodeToString(i[:])
}
