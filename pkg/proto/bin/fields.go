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
	"fmt"
	"strconv"
)

// Fields represent a bitfield value that compactly encodes
// information about provided conditional fields, e.g. says
// that fields "1", "5" and "10" were set.
type Fields uint32

func validateBitIndex(n int) {
	if n < 0 || n >= 32 {
		panic(fmt.Sprintf("bin: invalid fields bit index %d", n))
	}
}

// Zero returns true, if all bits are equal to zero.
func (f *Fields) Zero() bool {
	return *f == 0
}

// String implement fmt.Stringer
func (f *Fields) String() string {
	return strconv.FormatUint(uint64(*f), 2)
}

// Decode implements Decoder.
func (f *Fields) Decode(d *Decoder) error {
	v, err := d.Int32()
	if err != nil {
		return err
	}
	*f = Fields(v)
	return nil
}

// Encode implements Encoder.
func (f *Fields) Encode(e *Encoder, layer int) {
	e.PutUint32(uint32(*f))
}

// Has reports whether field with index n was set.
func (f *Fields) Has(n int) bool {
	validateBitIndex(n)
	return *f&(1<<n) != 0
}

// Unset unsets field with index n.
func (f *Fields) Unset(n int) {
	validateBitIndex(n)
	*f &= ^(1 << n)
}

// Set sets field with index n.
func (f *Fields) Set(n int) {
	validateBitIndex(n)
	*f |= 1 << n
}
