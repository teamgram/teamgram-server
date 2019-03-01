// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/sync2"
	"time"
)

var msgIdSeq = sync2.NewAtomicInt64(0)

func GenerateMessageId() int64 {
	unixnano := time.Now().UnixNano()
	ts := unixnano / 1e9
	ms := (unixnano % 1e9) / 1e6
	sid := msgIdSeq.Add(1) & 0x1ffff
	msgIdSeq.CompareAndSwap(0x1ffff, 0)
	last := 1
	//if !isRpc {
	//	last = 3
	//}
	msgId := int64(ts<<32) | int64(ms)<<21 | int64(sid)<<3 | int64(last)
	return msgId
}

//func nextMessageId(isRpc bool) int64 {
//	unixnano := time.Now().UnixNano()
//	ts := unixnano / 1e9
//	ms := (unixnano % 1e9) / 1e6
//	sid := msgIdSeq.Add(1) & 0x1ffff
//	msgIdSeq.CompareAndSwap(0x1ffff, 0)
//	last := 1
//	if !isRpc {
//		last = 3
//	}
//	msgId := int64(ts<<32) | int64(ms)<<21 | int64(sid)<<3 | int64(last)
//	return msgId
//}
//

/*
	uint32_t x = incoming ? 8 : 0;
	static uint8_t sha[68];

	SHA256_Init(&sha256Ctx);
	SHA256_Update(&sha256Ctx, messageKey, 16);
	SHA256_Update(&sha256Ctx, authKey + x, 36);
	SHA256_Final(sha, &sha256Ctx);

	SHA256_Init(&sha256Ctx);
	SHA256_Update(&sha256Ctx, authKey + 40 + x, 36);
	SHA256_Update(&sha256Ctx, messageKey, 16);
	SHA256_Final(sha + 32, &sha256Ctx);

	memcpy(result, sha, 8);
	memcpy(result + 8, sha + 32 + 8, 16);
	memcpy(result + 8 + 16, sha + 24, 8);

	memcpy(result + 32, sha + 32, 8);
	memcpy(result + 32 + 8, sha + 8, 16);
	memcpy(result + 32 + 8 + 16, sha + 32 + 24, 8);

*/
func generateMessageKey(msgKey, authKey []byte, incoming bool) (aesKey, aesIV []byte) {
	var x = 0
	if incoming {
		x = 8
	}

	switch MTPROTO_VERSION {
	case 2:
		t_a := make([]byte, 0, 52)
		t_a = append(t_a, msgKey[:16]...)
		t_a = append(t_a, authKey[x:x+36]...)
		sha256_a := crypto.Sha256Digest(t_a)

		t_b := make([]byte, 0, 52)
		t_b = append(t_b, authKey[40+x:40+x+36]...)
		t_b = append(t_b, msgKey[:16]...)
		sha256_b := crypto.Sha256Digest(t_b)

		aesKey = make([]byte, 0, 32)
		aesKey = append(aesKey, sha256_a[:8]...)
		aesKey = append(aesKey, sha256_b[8:8+16]...)
		aesKey = append(aesKey, sha256_a[24:24+8]...)

		aesIV = make([]byte, 0, 32)
		aesIV = append(aesIV, sha256_b[:8]...)
		aesIV = append(aesIV, sha256_a[8:8+16]...)
		aesIV = append(aesIV, sha256_b[24:24+8]...)

	default:
		aesKey = make([]byte, 0, 32)
		aesIV = make([]byte, 0, 32)
		t_a := make([]byte, 0, 48)
		t_b := make([]byte, 0, 48)
		t_c := make([]byte, 0, 48)
		t_d := make([]byte, 0, 48)

		t_a = append(t_a, msgKey...)
		t_a = append(t_a, authKey[x:x+32]...)

		t_b = append(t_b, authKey[32+x:32+x+16]...)
		t_b = append(t_b, msgKey...)
		t_b = append(t_b, authKey[48+x:48+x+16]...)

		t_c = append(t_c, authKey[64+x:64+x+32]...)
		t_c = append(t_c, msgKey...)

		t_d = append(t_d, msgKey...)
		t_d = append(t_d, authKey[96+x:96+x+32]...)

		sha1_a := crypto.Sha1Digest(t_a)
		sha1_b := crypto.Sha1Digest(t_b)
		sha1_c := crypto.Sha1Digest(t_c)
		sha1_d := crypto.Sha1Digest(t_d)

		aesKey = append(aesKey, sha1_a[0:8]...)
		aesKey = append(aesKey, sha1_b[8:8+12]...)
		aesKey = append(aesKey, sha1_c[4:4+12]...)

		aesIV = append(aesIV, sha1_a[8:8+12]...)
		aesIV = append(aesIV, sha1_b[0:8]...)
		aesIV = append(aesIV, sha1_c[16:16+4]...)
		aesIV = append(aesIV, sha1_d[0:8]...)
	}

	return
}
