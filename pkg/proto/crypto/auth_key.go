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
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	serverSide = 0
	clientSide = 1
)

type AuthKey struct {
	authKeyId int64
	authKey   []byte
	side      int // client or server
}

func calcAuthKeyId(keyData []byte) int64 {
	sha1 := Sha1Digest(keyData)
	// Lower 64 bits = 8 bytes of 20 byte SHA1 hash.
	return int64(binary.LittleEndian.Uint64(sha1[12:]))
}

func NewAuthKey(keyId int64, keyData []byte) *AuthKey {
	// TODO(@benqi): check keyData len

	// check keyId
	if keyId == 0 {
		keyId = calcAuthKeyId(keyData)
	}
	return &AuthKey{
		authKeyId: keyId,
		authKey:   keyData,
		side:      serverSide,
	}
}

// CreateAuthKey
/*
## android client, pushAuthKey algo:
- authKey
	if (SharedConfig.pushAuthKey == null) {
		SharedConfig.pushAuthKey = new byte[256];
		Utilities.random.nextBytes(SharedConfig.pushAuthKey);
		SharedConfig.saveConfig();
	}

- calcAuthKeyId
	SharedConfig.pushAuthKeyId = new byte[8];
	byte[] authKeyHash = Utilities.computeSHA1(SharedConfig.pushAuthKey);
	System.arraycopy(authKeyHash, authKeyHash.length - 8, SharedConfig.pushAuthKeyId, 0, 8);

*/
// for bots...
func CreateAuthKey() *AuthKey {
	key := new(AuthKey)

	key.authKey = GenerateNonce(256)
	key.authKeyId = calcAuthKeyId(key.authKey)
	key.side = serverSide

	return key
}

func NewClientAuthKey(keyId int64, keyData []byte) *AuthKey {
	// check keyId
	// check keyId
	if keyId == 0 {
		keyId = calcAuthKeyId(keyData)
	}

	return &AuthKey{
		authKeyId: keyId,
		authKey:   keyData,
		side:      clientSide,
	}
}

func (k *AuthKey) AuthKeyId() int64 {
	return k.authKeyId
}

func (k *AuthKey) AuthKey() []byte {
	return k.authKey
}

func (k *AuthKey) Equals(o *AuthKey) bool {
	return bytes.Equal(k.authKey, o.authKey)
}

//func (k *AuthKey) CalcAuthKeyId() int64 {
//	return calcAuthKeyId(k.authKey)
//}

func (k *AuthKey) calcX(incoming bool) int {
	var x = 0
	if k.side == serverSide {
		if incoming {
			x = 8
		}
	} else {
		if !incoming {
			x = 8
		}
	}
	return x
}

func (k *AuthKey) prepareAESV1(msgKey []byte, incoming bool) ([]byte, []byte) {
	x := k.calcX(incoming)

	aesKey := make([]byte, 32)
	aesIV := make([]byte, 32)

	dataA := make([]byte, 16+32)
	copy(dataA, msgKey[:16])
	copy(dataA[16:], k.authKey[x:x+32])
	sha1A := Sha1Digest(dataA)

	dataB := make([]byte, 16+16+16)
	copy(dataB, k.authKey[32+x:32+x+16])
	copy(dataB[16:], msgKey[:16])
	copy(dataB[32:], k.authKey[48+x:48+x+16])
	sha1B := Sha1Digest(dataB)

	dataC := make([]byte, 32+16)
	copy(dataC, k.authKey[64+x:64+x+32])
	copy(dataC[32:], msgKey[:16])
	sha1C := Sha1Digest(dataC)

	dataD := make([]byte, 16+32)
	copy(dataD, msgKey[:16])
	copy(dataD[16:], k.authKey[96+x:96+x+32])
	sha1D := Sha1Digest(dataD)

	copy(aesKey, sha1A[:8])
	copy(aesKey[8:], sha1B[8:8+12])
	copy(aesKey[8+12:], sha1C[4:4+12])
	copy(aesIV, sha1A[8:8+12])
	copy(aesIV[12:], sha1B[:8])
	copy(aesIV[12+8:], sha1C[16:16+4])
	copy(aesIV[12+8+4:], sha1D[:8])

	return aesKey, aesIV
}

func (k *AuthKey) prepareAES(msgKey []byte, incoming bool) ([]byte, []byte) {
	x := k.calcX(incoming)

	aesKey := make([]byte, 32)
	aesIV := make([]byte, 32)

	dataA := make([]byte, 16+36)
	copy(dataA, msgKey[:16])
	copy(dataA[16:], k.authKey[x:x+36])
	sha256A := Sha256Digest(dataA)

	dataB := make([]byte, 36+16)
	copy(dataB, k.authKey[40+x:40+x+36])
	copy(dataB[36:], msgKey[:16])
	sha256B := Sha256Digest(dataB)

	copy(aesKey, sha256A[:8])
	copy(aesKey[8:], sha256B[8:8+16])
	copy(aesKey[8+16:], sha256A[24:24+8])
	copy(aesIV, sha256B[:8])
	copy(aesIV[8:], sha256A[8:8+16])
	copy(aesIV[8+16:], sha256B[24:24+8])

	return aesKey, aesIV
}

func (k *AuthKey) partForMsgKey(incoming bool) []byte {
	x := k.calcX(incoming)
	return k.authKey[88+x : 88+x+32]
}

// AesIgeEncryptV1
/*
| salt <br> int64	| `session_id` <br> int64 | `message_id` <br> int64 | `seq_no` <br> int32 |`message_data_length` <br> int32	| `message_data` <br> bytes | padding12..1024 <br> bytes|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
*/
func (k *AuthKey) AesIgeEncryptV1(rawData []byte) ([]byte, []byte, error) {
	var additionalSize = len(rawData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}

	// var tmpData []byte
	// if additionalSize >
	tmpData := make([]byte, 0, len(rawData)+additionalSize)
	tmpData = append(tmpData, rawData...)
	if additionalSize > 0 {
		tmpData = append(tmpData, GenerateNonce(additionalSize)...)
	}

	// calc msg_key
	msgKey := make([]byte, 32)
	copy(msgKey[4:], Sha1Digest(rawData))

	aesKey, aesIV := k.prepareAESV1(msgKey[8:8+16], true)
	e := NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(tmpData)
	if err != nil {
		// log.Errorf("aesIgeEncrypt data error: %v", err)
		return nil, nil, err
	}

	return msgKey[8 : 8+16], x, nil
}

func (k *AuthKey) AesIgeEncrypt(rawData []byte) ([]byte, []byte, error) {
	var additionalSize = len(rawData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}

	if additionalSize < 12 {
		additionalSize += 16
	}

	// var tmpData []byte
	// if additionalSize >
	tmpData := make([]byte, 0, len(rawData)+additionalSize)
	tmpData = append(tmpData, rawData...)
	if additionalSize > 0 {
		tmpData = append(tmpData, GenerateNonce(additionalSize)...)
	}

	// calc msg_key
	msgKey := make([]byte, 32)
	sha256Dig := sha256.New()
	sha256Dig.Write(k.partForMsgKey(true))
	sha256Dig.Write(tmpData)
	copy(msgKey, sha256Dig.Sum(nil))

	aesKey, aesIV := k.prepareAES(msgKey[8:8+16], true)
	e := NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(tmpData)
	if err != nil {
		// log.Errorf("aesIgeEncrypt data error: %v", err)
		return nil, nil, err
	}

	return msgKey[8 : 8+16], x, nil
}

func (k *AuthKey) AesIgeDecryptV1(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAESV1(msgKey, false)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		// log.Errorf("aesIgeDecrypt data error: %v", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	var dataLen = uint32(len(rawData))
	messageLen := binary.LittleEndian.Uint32(x[28:])
	if messageLen+32 > dataLen {
		err = fmt.Errorf("aesIgeDecrypt data(%d) error - Wrong message length %d", dataLen, messageLen)
		return nil, err
	}

	calcMsgKey := make([]byte, 96)
	copy(calcMsgKey[4:], Sha1Digest(x[:32+messageLen]))

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		return nil, err
	}

	return x, nil
}

func (k *AuthKey) AesIgeDecrypt(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAES(msgKey, false)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		// log.Errorf("aesIgeDecrypt data error: %v", err)
		return nil, err
	}

	//// 校验解密后的数据合法性
	var dataLen = uint32(len(rawData))

	// dBuf := mtproto.NewDecodeBuf(rawData)

	//salt := binary.LittleEndian.Uint64(x[:8])
	//sessionId := binary.LittleEndian.Uint64(x[8:])
	//msgId := binary.LittleEndian.Uint64(x[16:])
	//seq := binary.LittleEndian.Uint32(x[24:])
	messageLen := binary.LittleEndian.Uint32(x[28:])
	//c := binary.LittleEndian.Uint32(x[32:])
	//fmt.Printf("decrypt: {salt: %d, session_id: %d, msg_id: %d, seq: %d, bytes: %d, crc32: 0x%x}\n",
	//	int64(salt),
	//	int64(sessionId),
	//	int64(msgId),
	//	seq,
	//	messageLen,
	//	c)

	//
	//messageLen := binary.LittleEndian.Uint32(x[28:])
	//sessionId := int64(binary.LittleEndian.Uint64(x[8:]))

	// log.Info("decrypt - dataLen = %d, messageLen = ", dataLen, messageLen)
	// log.Debugf("decrypt - dataLen = %d, messageLen = %d, sessionId = %d", dataLen, messageLen, sessionId)

	if messageLen+32 > dataLen {
		// 	return fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		err = fmt.Errorf("aesIgeDecrypt data(%d) error - Wrong message length %d", dataLen, messageLen)
		// log.Error(err.Error())
		return nil, err
	}

	calcMsgKey := make([]byte, 96)
	sha256Dig := sha256.New()
	sha256Dig.Write(k.partForMsgKey(false))
	sha256Dig.Write(x[:dataLen])
	copy(calcMsgKey, sha256Dig.Sum(nil))

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		//calcMsgKey := make([]byte, 96)
		//sha256Dig := sha256.New()
		//sha256Dig.Write(k.partForMsgKey(false))
		//sha256Dig.Write(x[:messageLen+32])
		//copy(calcMsgKey, sha256Dig.Sum(nil))
		//if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		// log.Error(err.Error())
		return nil, err
		//}
	}

	return x, nil
}
