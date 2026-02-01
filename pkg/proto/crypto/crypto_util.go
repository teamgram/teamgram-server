// Copyright 2024 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"math/big"
	"math/rand"

	"golang.org/x/crypto/pbkdf2"
)

var bigIntZero = big.NewInt(0)

//// modPow uses right-to-left binary method
//func modPow(a, n *big.Int) *big.Int {
//	b, e, m, r := new(big.Int), new(big.Int), new(big.Int), big.NewInt(1)
//	b.Set(a)
//	e.Set(n)
//	m.Set(n)
//
//	b.Mod(b, m)
//	zero := big.NewInt(0)
//	for e.Cmp(zero) == 1 {
//		if e.Bit(0) == 1 {
//			r.Mul(r, b).Mod(r, m)
//		}
//		e.Rsh(e, 1)
//		b.Mul(b, b).Mod(b, m)
//	}
//	return r
//}

// public static boolean i
func isGoodGaAndGb(gA, p *big.Int) bool {
	return !(gA.Cmp(big.NewInt(1)) <= 0 || gA.Cmp(new(big.Int).Sub(p, big.NewInt(1))) >= 0)
}

func isGoodPrime(prime []byte, g int) bool {
	if !(g >= 2 && g <= 7) {
		return false
	}

	if len(prime) != 256 || int8(prime[0]) >= 0 {
		return false
	}

	dhBI := new(big.Int).SetBytes(prime)

	if g == 2 { // p mod 8 = 7 for g = 2;
		res := new(big.Int).Mod(dhBI, new(big.Int).SetUint64(8))
		if res.Int64() != 7 {
			return false
		}
	} else if g == 3 { // p mod 3 = 2 for g = 3;
		res := new(big.Int).Mod(dhBI, new(big.Int).SetUint64(3))
		if res.Int64() != 2 {
			return false
		}
	} else if g == 5 { // p mod 5 = 1 or 4 for g = 5;
		res := new(big.Int).Mod(dhBI, new(big.Int).SetUint64(5))
		val := res.Int64()
		if val != 1 && val != 4 {
			return false
		}
	} else if g == 6 { // p mod 24 = 19 or 23 for g = 6;
		res := new(big.Int).Mod(dhBI, new(big.Int).SetUint64(24))
		val := res.Int64()
		if val != 19 && val != 23 {
			return false
		}
	} else if g == 7 { // p mod 7 = 3, 5 or 6 for g = 7.
		res := new(big.Int).Mod(dhBI, new(big.Int).SetUint64(7))
		val := res.Int64()
		if val != 3 && val != 5 && val != 6 {
			return false
		}
	}

	hex := hex.EncodeToString(prime)
	if hex == "C71CAEB9C6B1C9048E6C522F70F13F73980D40238E3E21C14934D037563D930F48198A0AA7C14058229493D22530F4DBFA336F6E0AC925139543AED44CCE7C3720FD51F69458705AC68CD4FE6B6B13ABDC9746512969328454F18FAF8C595F642477FE96BB2A941D5BCD1D4AC8CC49880708FA9B378E3C4F3A9060BEE67CF9A4A4A695811051907E162753B56B0F6B410DBA74D8A84B2A14B3144E0EF1284754FD17ED950D5965B4B9DD46582DB1178D169C6BC465B0D6FF9CA3928FEF5B9AE4E418FC15E83EBEA0F87FA9FF5EED70050DED2849F47BF959D956850CE929851F0D8115F635B105EE2E4E15D04B2454BF6F4FADF034B10403119CD8E3B92FCC5B" {
		return true
	}

	dhBI2 := new(big.Int).Div(new(big.Int).Sub(dhBI, big.NewInt(1)), big.NewInt(2))
	return !(!dhBI.ProbablyPrime(30) || !dhBI2.ProbablyPrime(30))
}

func calcSHA256(args ...[]byte) []byte {
	hash := sha256.New()
	for _, a := range args {
		hash.Write(a)
	}
	return hash.Sum(nil)
}

/*
	 from android client:
		byte[] dst = new byte[64];
		Utilities.pbkdf2(password, salt, dst, 100000);
		return dst;

		int result = PKCS5_PBKDF2_HMAC((char *) passwordBuff, passwordLength, (uint8_t *) saltBuff, saltLength, (unsigned int) iterations, EVP_sha512(), dstLength, (uint8_t *) dstBuff);
*/
func calcPBKDF2(password []byte, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 100000, 64, sha512.New)
}

func getBigIntegerBytes(value *big.Int) []byte {
	bytes := value.Bytes()
	if len(bytes) > 256 {
		correctedAuth := make([]byte, 256)
		copy(correctedAuth, bytes[1:256+1])
		return correctedAuth
	} else if len(bytes) < 256 {
		correctedAuth := make([]byte, 256)
		copy(correctedAuth[256-len(bytes):], bytes)
		for i := 0; i < 256-len(bytes); i++ {
			correctedAuth[i] = 0
		}
		return correctedAuth
	}
	return bytes
}

func RandomBytes(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func RandomString(size int) string {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func Sha256Digest(data []byte) []byte {
	r := sha256.Sum256(data)
	return r[:]
}

func Sha1Digest(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateStringNonce(size int) string {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
