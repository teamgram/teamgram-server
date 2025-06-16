// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"github.com/teamgram/proto/v2/crypto"
)

type ObfuscatedCodec struct {
	Codec
	dc int16
}

func newMTProtoObfuscatedCodec(d, e *crypto.AesCTR128Encrypt, protocolType uint32, dc int16) *ObfuscatedCodec {
	codec := new(ObfuscatedCodec)
	codec.dc = dc

	switch protocolType {
	case ABRIDGED_INT32_FLAG:
		codec.Codec = newMTProtoAbridgedCodec(newAesCTR128Crypto(d, e))
	case INTERMEDIATE_FLAG:
		codec.Codec = newMTProtoIntermediateCodec(newAesCTR128Crypto(d, e))
	case PADDED_INTERMEDIATE_FLAG:
		codec.Codec = newMTProtoPaddedIntermediateCodec(newAesCTR128Crypto(d, e))
	}

	return codec
}
