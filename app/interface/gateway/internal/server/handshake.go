// Copyright 2022 Teamgram Authors
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

package server

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/hex2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	SHA_DIGEST_LENGTH = 20
)

var (
	// TODO(@benqi): 预先计算出fingerprint
	// 这里直接使用了0xc3b42b026ce86b21
	// fingerprint uint64 = 12240908862933197005

	// TODO(@benqi): 使用算法生成PQ
	// 这里直接指定了PQ值: {0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81}
	pq = string([]byte{0x17, 0xED, 0x48, 0x94, 0x1A, 0x08, 0xF9, 0x81})

	// TODO(@benqi): 直接指定了p和q
	p = []byte{0x49, 0x4C, 0x55, 0x3B}
	q = []byte{0x53, 0x91, 0x10, 0x73}

	// TODO(@benqi): 直接指定了dh2048_p和dh2048_g!!!
	// android client 指定的good prime
	//
	// static const char *goodPrime = "
	//
	// c71caeb9c6b1c9048e6c522f
	// 70f13f73980d40238e3e21c1
	// 4934d037563d930f48198a0a
	// a7c14058229493d22530f4db
	// fa336f6e0ac925139543aed4
	// 4cce7c3720fd51f69458705a
	// c68cd4fe6b6b13abdc974651
	// 2969328454f18faf8c595f64
	// 2477fe96bb2a941d5bcd1d4a
	// c8cc49880708fa9b378e3c4f
	// 3a9060bee67cf9a4a4a69581
	// 1051907e162753b56b0f6b41
	// 0dba74d8a84b2a14b3144e0e
	// f1284754fd17ed950d5965b4
	// b9dd46582db1178d169c6bc4
	// 65b0d6ff9ca3928fef5b9ae4
	// e418fc15e83ebea0f87fa9ff
	// 5eed70050ded2849f47bf959
	// d956850ce929851f0d8115f6
	// 35b105ee2e4e15d04b2454bf
	// 6f4fadf034b10403119cd8e3
	// b92fcc5b";
	//
	dh2048P = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}

	dh2048G = []byte{0x03}

	zeroIV = []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
)

var (
	gBigIntDH2048P *big.Int
	gBigIntDH2048G *big.Int
)

func init() {
	gBigIntDH2048P = new(big.Int).SetBytes(dh2048P)
	gBigIntDH2048G = new(big.Int).SetBytes(dh2048G)
}

type keyCreatedF func(ctx context.Context, keyInfo *mtproto.AuthKeyInfo, salt *mtproto.FutureSalt, expiresIn int32) error

type handshake struct {
	rsa            *crypto.RSACryptor
	keyFingerprint uint64
	dh2048p        []byte
	dh2048g        []byte
	keyCreatedF
}

func newHandshake(keyFile string, keyFingerprint uint64, createdCB keyCreatedF) (*handshake, error) {
	rsa, err := crypto.NewRSACryptor(keyFile)
	if err != nil {
		return nil, err
	}
	return &handshake{
		rsa:            rsa,
		keyFingerprint: keyFingerprint,
		dh2048p:        dh2048P,
		dh2048g:        dh2048G,
		keyCreatedF:    createdCB,
	}, nil
}

//func (s *handshake) onHandshake(state *HandshakeStateCtx, mtpBuf []byte) ([]byte, error) {
//	var (
//		err error
//		res mtproto.TLObject
//	)
//
//	if len(mtpBuf) < 8 {
//		err = fmt.Errorf("invalid data len < 8")
//		logx.Error(err.Error())
//		return nil, err
//	}
//
//	_, obj, err := ParseFromIncomingMessage(mtpBuf[8:])
//
//	// TODO(@benqi): check msgId
//	// _ = msgId
//
//	//mtpMessage := &mtproto.UnencryptedMessage{}
//	//err = mtpMessage.Decode(mtpBuf[8:])
//	if err != nil {
//		logx.Error(err.Error())
//		return nil, err
//	}
//
//	switch obj.(type) {
//	case *mtproto.TLReqPq:
//		res, err = s.onReqPq(obj.(*mtproto.TLReqPq))
//		state.State = STATE_pq_res
//	case *mtproto.TLReqPqMulti:
//		res, err = s.onReqPqMulti(state, obj.(*mtproto.TLReqPqMulti))
//		state.State = STATE_pq_res
//	case *mtproto.TLReq_DHParams:
//		res, err = s.onReqDHParams(state, obj.(*mtproto.TLReq_DHParams))
//		state.State = STATE_DH_params_res
//	case *mtproto.TLSetClient_DHParams:
//		res, err = s.onSetClientDHParams(state, obj.(*mtproto.TLSetClient_DHParams))
//		state.State = STATE_dh_gen_res
//	case *mtproto.TLMsgsAck:
//		err = s.onMsgsAck(state, obj.(*mtproto.TLMsgsAck))
//		return nil, err
//	default:
//		err = fmt.Errorf("invalid handshake type: %v", state)
//		return nil, err
//	}
//
//	if err != nil {
//		state.ResState = RES_STATE_ERROR
//		return nil, err
//	} else {
//		state.ResState = RES_STATE_OK
//		return SerializeToBuffer(mtproto.GenerateMessageId(), res), nil
//	}
//}

// req_pq#60469778 nonce:int128 = ResPQ;
func (s *handshake) onReqPq(request *mtproto.TLReqPq) (*mtproto.TLResPQ, error) {
	logx.Infof("req_pq#60469778 - {\"request\":%s", request)

	// check State and ResState

	// 检查数据是否合法
	if request.GetNonce() == nil || len(request.GetNonce()) != 16 {
		err := fmt.Errorf("onReqPq - invalid nonce: %v", request)
		logx.Error(err.Error())
		return nil, err
	}

	resPQ := &mtproto.TLResPQ{Data2: &mtproto.ResPQ{
		Nonce:                       request.Nonce,
		ServerNonce:                 crypto.GenerateNonce(16),
		Pq:                          pq,
		ServerPublicKeyFingerprints: []int64{int64(s.keyFingerprint)},
	}}

	// 缓存客户端Nonce
	// ctx.Nonce = request.GetNonce()
	// ctx.ServerNonce = resPQ.GetServerNonce()

	// PutCacheState(ctx.Nonce, ctx.ServerNonce, ctx)
	logx.Infof("req_pq#60469778 reply - {\"resPQ\":%s", resPQ)
	return resPQ, nil
}

// req_pq_multi#be7e8ef1 nonce:int128 = ResPQ;
func (s *handshake) onReqPqMulti(request *mtproto.TLReqPqMulti) (*mtproto.TLResPQ, error) {
	logx.Infof("req_pq_multi#be7e8ef1 - {\"request\":%s", request)

	// check State and ResState

	// 检查数据是否合法
	if request.GetNonce() == nil || len(request.GetNonce()) != 16 {
		err := fmt.Errorf("onReqPq - invalid nonce: %v", request)
		logx.Error(err.Error())
		return nil, err
	}

	resPQ := &mtproto.TLResPQ{Data2: &mtproto.ResPQ{
		Nonce:                       request.Nonce,
		ServerNonce:                 crypto.GenerateNonce(16),
		Pq:                          pq,
		ServerPublicKeyFingerprints: []int64{int64(s.keyFingerprint)},
	}}

	// 缓存客户端Nonce
	// ctx := state.GetCtx()
	//ctx.Nonce = request.GetNonce()
	//ctx.ServerNonce = resPQ.GetServerNonce()

	// PutCacheState(ctx.Nonce, ctx.ServerNonce, ctx)

	logx.Infof("req_pq_multi#be7e8ef1 - reply: %s", resPQ)
	return resPQ, nil
}

// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
func (s *handshake) onReqDHParams(ctx *HandshakeStateCtx, request *mtproto.TLReq_DHParams) (*mtproto.Server_DH_Params, error) {
	logx.Infof("req_DH_params#d712e4be - state: {%s}, request: %s", ctx, request)

	var (
		err error
	)

	// 客户端传输数据解析
	// check Nonce
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		err = fmt.Errorf("onReq_DHParams - Invalid Nonce, req: %s, back: %s",
			hex2.HexDump(request.Nonce),
			hex2.HexDump(ctx.Nonce))
		logx.Error(err.Error())
		return nil, err
	}

	// check ServerNonce
	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		err = fmt.Errorf("onReq_DHParams - Wrong ServerNonce, req: %s, back: %s",
			hex2.HexDump(request.ServerNonce),
			hex2.HexDump(ctx.ServerNonce))
		logx.Error(err.Error())
		return nil, err
	}

	// check P
	if !bytes.Equal([]byte(request.P), p) {
		err = fmt.Errorf("onReq_DHParams - Invalid p valuee")
		logx.Error(err.Error())
		return nil, err
	}

	// check Q
	if !bytes.Equal([]byte(request.Q), q) {
		err = fmt.Errorf("onReq_DHParams - Invalid q value")
		logx.Error(err.Error())
		return nil, err
	}

	if request.PublicKeyFingerprint != int64(s.keyFingerprint) {
		err = fmt.Errorf("onReq_DHParams - Invalid PublicKeyFingerprint value")
		logx.Error(err.Error())
		return nil, err
	}

	/*
		### 4.1) RSA_PAD(data, server_public_key) mentioned above is implemented as follows:

		- data_with_padding := data + random_padding_bytes; — where random_padding_bytes are chosen so that the resulting length of data_with_padding is precisely 192 bytes, and data is the TL-serialized data to be encrypted as before. One has to check that data is not longer than 144 bytes.
		- data_pad_reversed := BYTE_REVERSE(data_with_padding); — is obtained from data_with_padding by reversing the byte order.
		- a random 32-byte temp_key is generated.
		- data_with_hash := data_pad_reversed + SHA256(temp_key + data_with_padding); — after this assignment, data_with_hash is exactly 224 bytes long.
		- aes_encrypted := AES256_IGE(data_with_hash, temp_key, 0); — AES256-IGE encryption with zero IV.
		- temp_key_xor := temp_key XOR SHA256(aes_encrypted); — adjusted key, 32 bytes
		- key_aes_encrypted := temp_key_xor + aes_encrypted; — exactly 256 bytes (2048 bits) long
		- The value of key_aes_encrypted is compared with the RSA-modulus of server_pubkey as a big-endian 2048-bit (256-byte) unsigned integer. If key_aes_encrypted turns out to be greater than or equal to the RSA modulus, the previous steps starting from the generation of new random temp_key are repeated. Otherwise the final step is performed:
		- encrypted_data := RSA(key_aes_encrypted, server_pubkey); — 256-byte big-endian integer is elevated to the requisite power from the RSA public key modulo the RSA modulus, and the result is stored as a big-endian integer consisting of exactly 256 bytes (with leading zero bytes if required).
	*/

	// encryptedData := []byte(request.EncryptedData)
	//
	logx.Infof("EncryptedData: len = %d, data: %s", len(request.EncryptedData), hex.EncodeToString(hack.Bytes(request.EncryptedData)))

	if len(request.EncryptedData) < 256 {
		logx.Error("need len(encryptedPQInnerData) < 256")
		return nil, fmt.Errorf("process Req_DHParams - len(encryptedPQInnerData) != 256")
	}

	//
	// 1. 解密
	innerData := s.rsa.Decrypt([]byte(request.EncryptedData))
	if len(innerData) != 256 {
		logx.Error("need len(encryptedPQInnerData) < 256")
		return nil, fmt.Errorf("process Req_DHParams - len(encryptedPQInnerData) != 256")
	}

	// void Datacenter::aesIgeEncryption(uint8_t *buffer, uint8_t *key, uint8_t *iv, bool encrypt, bool changeIv, uint32_t length) {
	// Datacenter::aesIgeEncryption(
	//	innerDataBuffer->bytes() + keySize,
	//	innerDataBuffer->bytes(),
	//	innerDataBuffer->bytes() + encryptedDataSize + paddedDataSize,
	//	true,
	//	true,
	//	paddedDataSize + SHA256_DIGEST_LENGTH);
	//

	key := innerData[:32]
	logx.Infof("key1: %s", hex.EncodeToString(key))
	hash := crypto.Sha256Digest(innerData[32:])
	for i := 0; i < 32; i++ {
		key[i] = key[i] ^ hash[i]
	}
	logx.Infof("key2: %s", hex.EncodeToString(key))

	paddedDataWithHash, err := crypto.NewAES256IGECryptor(key, zeroIV).Decrypt(innerData[32:])
	if err != nil {
		err = fmt.Errorf("onReq_DHParams - decode P_Q_inner_data error")
		logx.Error(err.Error())
		return nil, err
	}

	logx.Infof("paddedDataWithHash1: %s", hex.EncodeToString(paddedDataWithHash))

	//if !bytes.Equal(crypto.Sha256Digest(paddedDataWithHash[:192]), paddedDataWithHash[192:]) {
	//	logx.Error("process Req_DHParams - sha1Check error")
	//	return nil, fmt.Errorf("process Req_DHParams - sha1Check error")
	//}

	for i, j := 0, 191; i < j; i, j = i+1, j-1 {
		paddedDataWithHash[i], paddedDataWithHash[j] = paddedDataWithHash[j], paddedDataWithHash[i]
	}
	logx.Infof("paddedDataWithHash2: %s", hex.EncodeToString(paddedDataWithHash))

	// TODO
	//if !checkSha1(encryptedPQInnerData, 256-SHA_DIGEST_LENGTH) {
	//	logx.Error("process Req_DHParams - sha1Check error")
	//	return nil, fmt.Errorf("process Req_DHParams - sha1Check error")
	//}

	// 2. 反序列化出pqInnerData
	dbuf := mtproto.NewDecodeBuf(paddedDataWithHash)
	o := dbuf.Object()
	if dbuf.GetError() != nil {
		err = fmt.Errorf("onReq_DHParams - decode P_Q_inner_data error")
		logx.Error(err.Error())
		return nil, err
	}

	var pqInnerData *mtproto.P_QInnerData
	// TODO(@benqi):
	switch innerData := o.(type) {
	case *mtproto.TLPQInnerData:
		ctx.handshakeType = mtproto.AuthKeyTypePerm
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataDc:
		ctx.handshakeType = mtproto.AuthKeyTypePerm
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTemp:
		ctx.handshakeType = mtproto.AuthKeyTypeTemp
		ctx.ExpiresIn = innerData.GetExpiresIn()
		pqInnerData = innerData.To_P_QInnerData()
	case *mtproto.TLPQInnerDataTempDc:
		if innerData.GetDc() < 0 {
			ctx.handshakeType = mtproto.AuthKeyTypeMediaTemp
		} else {
			ctx.handshakeType = mtproto.AuthKeyTypeTemp
		}
		ctx.ExpiresIn = innerData.GetExpiresIn()
		pqInnerData = innerData.To_P_QInnerData()
	default:
		err = fmt.Errorf("onReq_DHParams - decode P_Q_inner_data error")
		logx.Error(err.Error())
		return nil, err
	}

	// 2. 再检查一遍p_q_inner_data里的pq, p, q, nonce, server_nonce合法性
	// 客户端传输数据解析
	// PQ
	if !bytes.Equal([]byte(pqInnerData.GetPq()), []byte(pq)) {
		logx.Error("process Req_DHParams - Invalid p_q_inner_data.pq value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.pq value")
	}

	// P
	if !bytes.Equal([]byte(pqInnerData.GetP()), p) {
		logx.Error("process Req_DHParams - Invalid p_q_inner_data.p value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.p value")
	}

	// Q
	if !bytes.Equal([]byte(pqInnerData.GetQ()), q) {
		logx.Error("process Req_DHParams - Invalid p_q_inner_data.q value")
		return nil, fmt.Errorf("process Req_DHParams - Invalid p_q_inner_data.q value")
	}

	// Nonce
	if !bytes.Equal(pqInnerData.GetNonce(), ctx.Nonce) {
		logx.Error("process Req_DHParams - Invalid Nonce")
		return nil, fmt.Errorf("process Req_DHParams - InvalidNonce")
	}

	// ServerNonce
	if !bytes.Equal(pqInnerData.GetServerNonce(), ctx.ServerNonce) {
		logx.Error("process Req_DHParams - Wrong ServerNonce")
		return nil, fmt.Errorf("process Req_DHParams - Wrong ServerNonce")
	}

	// TODO(@benqi): check dc
	// logx.Info("processReq_DHParams - pqInnerData Decode sucess: ", pqInnerData.String())

	// 检查NewNonce的长度(int256)
	// 缓存NewNonce
	ctx.NewNonce = pqInnerData.GetNewNonce()
	A := crypto.GenerateNonce(256)
	ctx.A = A
	ctx.P = s.dh2048p

	bigIntA := new(big.Int).SetBytes(A)

	// 服务端计算GA = g^a mod p
	gA := new(big.Int).Exp(gBigIntDH2048G, bigIntA, gBigIntDH2048P)

	// ServerNonce
	serverDHInnerData := &mtproto.TLServer_DHInnerData{Data2: &mtproto.Server_DHInnerData{
		Nonce:       ctx.Nonce,
		ServerNonce: ctx.ServerNonce,
		G:           int32(s.dh2048g[0]),
		GA:          string(gA.Bytes()),
		DhPrime:     string(s.dh2048p),
		ServerTime:  int32(time.Now().Unix()),
	}}

	x := mtproto.NewEncodeBuf(500)
	serverDHInnerData.Encode(x, 0)
	serverDHInnerDataBuf := x.GetBuf()
	// server_DHInnerData_buf_sha1 := sha1.Sum(server_DHInnerData_buf)

	// 创建aes和iv key
	tmpAesKeyAndIV := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce, ctx.ServerNonce...))
	sha1B := sha1.Sum(append(ctx.ServerNonce, ctx.NewNonce...))
	sha1C := sha1.Sum(append(ctx.NewNonce, ctx.NewNonce...))
	copy(tmpAesKeyAndIV, sha1A[:])
	copy(tmpAesKeyAndIV[20:], sha1B[:])
	copy(tmpAesKeyAndIV[40:], sha1C[:])
	copy(tmpAesKeyAndIV[60:], ctx.NewNonce[:4])

	tmpLen := 20 + len(serverDHInnerDataBuf)
	if tmpLen%16 > 0 {
		tmpLen = (tmpLen/16 + 1) * 16
	} else {
		tmpLen = 20 + len(serverDHInnerDataBuf)
	}

	tmpEncryptedAnswer := make([]byte, tmpLen)
	sha1Tmp := sha1.Sum(serverDHInnerDataBuf)
	copy(tmpEncryptedAnswer, sha1Tmp[:])
	copy(tmpEncryptedAnswer[20:], serverDHInnerDataBuf)

	e := crypto.NewAES256IGECryptor(tmpAesKeyAndIV[:32], tmpAesKeyAndIV[32:64])
	tmpEncryptedAnswer, _ = e.Encrypt(tmpEncryptedAnswer)

	serverDHParamsOk := &mtproto.TLServer_DHParamsOk{Data2: &mtproto.Server_DH_Params{
		Nonce:           ctx.Nonce,
		ServerNonce:     ctx.ServerNonce,
		EncryptedAnswer: hack.String(tmpEncryptedAnswer),
	}}

	logx.Infof("onReq_DHParams - state: {%s}, reply: %s", ctx, serverDHParamsOk)

	return serverDHParamsOk.To_Server_DH_Params(), nil
}

// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
func (s *handshake) onSetClientDHParams(ctx *HandshakeStateCtx, request *mtproto.TLSetClient_DHParams) (*mtproto.SetClient_DHParamsAnswer, error) {
	logx.Infof("set_client_DH_params#f5045f1f - state: {%s}, request: %s", ctx, request)

	// TODO(@benqi): Impl SetClient_DHParams logic
	// 客户端传输数据解析
	// Nonce
	if !bytes.Equal(request.Nonce, ctx.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong Nonce")
		logx.Error(err.Error())
		return nil, err
	}

	// ServerNonce
	if !bytes.Equal(request.ServerNonce, ctx.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong ServerNonce")
		logx.Error(err.Error())
		return nil, err
	}

	bEncryptedData := []byte(request.EncryptedData)

	// 创建aes和iv key
	tmpAesKeyAndIv := make([]byte, 64)
	sha1A := sha1.Sum(append(ctx.NewNonce, ctx.ServerNonce...))
	sha1B := sha1.Sum(append(ctx.ServerNonce, ctx.NewNonce...))
	sha1C := sha1.Sum(append(ctx.NewNonce, ctx.NewNonce...))
	copy(tmpAesKeyAndIv, sha1A[:])
	copy(tmpAesKeyAndIv[20:], sha1B[:])
	copy(tmpAesKeyAndIv[40:], sha1C[:])
	copy(tmpAesKeyAndIv[60:], ctx.NewNonce[:4])

	d := crypto.NewAES256IGECryptor(tmpAesKeyAndIv[:32], tmpAesKeyAndIv[32:64])
	decryptedData, err := d.Decrypt(bEncryptedData)
	if err != nil {
		err := fmt.Errorf("process SetClient_DHParams - AES256IGECryptor descrypt error")
		logx.Error(err.Error())
		return nil, err
	}

	// TODO(@benqi): 检查签名是否合法
	dBuf := mtproto.NewDecodeBuf(decryptedData[20:])
	clientDHInnerData := mtproto.MakeTLClient_DHInnerData(nil)
	clientDHInnerData.Data2.Constructor = mtproto.TLConstructor(dBuf.Int())
	err = clientDHInnerData.Decode(dBuf)
	if err != nil {
		logx.Errorf("processSetClient_DHParams - TLClient_DHInnerData decode error: %s", err)
		return nil, err
	}

	logx.Infof("processSetClient_DHParams - client_DHInnerData: %#v", clientDHInnerData.String())

	//
	if !bytes.Equal(clientDHInnerData.GetNonce(), ctx.Nonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's Nonce")
		logx.Error(err.Error())
		return nil, err
	}

	// ServerNonce
	if !bytes.Equal(clientDHInnerData.GetServerNonce(), ctx.ServerNonce) {
		err := fmt.Errorf("process SetClient_DHParams - Wrong client_DHInnerData's ServerNonce")
		logx.Error(err.Error())
		return nil, err
	}

	bigIntA := new(big.Int).SetBytes(ctx.A)
	// bigIntP := new(big.Int).SetBytes(authKeyMD.P)

	// hash_key
	authKeyNum := new(big.Int)
	authKeyNum.Exp(new(big.Int).SetBytes([]byte(clientDHInnerData.GetGB())), bigIntA, gBigIntDH2048P)

	authKey := make([]byte, 256)

	// TODO(@benqi): dhGenRetry and dhGenFail
	copy(authKey[256-len(authKeyNum.Bytes()):], authKeyNum.Bytes())

	authKeyAuxHash := make([]byte, len(ctx.NewNonce))
	copy(authKeyAuxHash, ctx.NewNonce)
	authKeyAuxHash = append(authKeyAuxHash, byte(0x01))
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)

	// 至此key已经创建成功
	authKeyId := int64(binary.LittleEndian.Uint64(authKeyAuxHash[len(ctx.NewNonce)+1+12 : len(ctx.NewNonce)+1+12+8]))

	// TODO(@benqi): authKeyId生成后要检查在数据库里是否已经存在，有非常小的概率会碰撞
	// 如果碰撞让客户端重新再来一轮

	// state.Ctx, _ = proto.Marshal(authKeyMD)
	if s.saveAuthKeyInfo(ctx, mtproto.NewAuthKeyInfo(authKeyId, authKey, ctx.handshakeType)) {
		dhGenOk := &mtproto.TLDhGenOk{Data2: &mtproto.SetClient_DHParamsAnswer{
			Nonce:         ctx.Nonce,
			ServerNonce:   ctx.ServerNonce,
			NewNonceHash1: calcNewNonceHash(ctx.NewNonce, authKey, 0x01),
		}}

		//ctx.AuthKeyId = authKeyId
		//ctx.AuthKey = authKey

		logx.Infof("onSetClient_DHParams - ctx: {%s}, reply: %s", ctx, dhGenOk)
		return dhGenOk.To_SetClient_DHParamsAnswer(), nil
	} else {
		// TODO(@benqi): dhGenFail
		dhGenRetry := &mtproto.TLDhGenRetry{Data2: &mtproto.SetClient_DHParamsAnswer{
			Nonce:         ctx.Nonce,
			ServerNonce:   ctx.ServerNonce,
			NewNonceHash2: calcNewNonceHash(ctx.NewNonce, authKey, 0x02),
		}}

		logx.Infof("onSetClient_DHParams - ctx: {%v}, reply: %s", ctx, dhGenRetry)
		return dhGenRetry.To_SetClient_DHParamsAnswer(), nil
	}
}

// msgs_ack#62d6b459 msg_ids:Vector<long> = MsgsAck;
func (s *handshake) onMsgsAck(state *HandshakeStateCtx, request *mtproto.TLMsgsAck) error {
	logx.Infof("msgs_ack#62d6b459 - state: {%s}, request: %s", state, request)

	switch state.State {
	case STATE_pq_res:
		state.State = STATE_pq_ack
	case STATE_DH_params_res:
		state.State = STATE_DH_params_ack
	case STATE_dh_gen_res:
		state.State = STATE_dh_gen_ack
	default:
		return fmt.Errorf("invalid state: %v", state)
	}

	return nil
}

func (s *handshake) saveAuthKeyInfo(ctx *HandshakeStateCtx, key *mtproto.AuthKeyInfo) bool {
	var (
		salt       = int64(0)
		serverSalt *mtproto.TLFutureSalt
		now        = int32(time.Now().Unix())
	)

	for a := 7; a >= 0; a-- {
		salt <<= 8
		salt |= int64(ctx.NewNonce[a] ^ ctx.ServerNonce[a])
	}

	serverSalt = &mtproto.TLFutureSalt{Data2: &mtproto.FutureSalt{
		ValidSince: now,
		ValidUntil: now + 30*60,
		Salt:       salt,
	}}

	keyInfo := &mtproto.AuthKeyInfo{
		AuthKeyId:          key.AuthKeyId,
		AuthKey:            key.AuthKey,
		AuthKeyType:        int32(key.AuthKeyType),
		PermAuthKeyId:      key.PermAuthKeyId,
		TempAuthKeyId:      key.TempAuthKeyId,
		MediaTempAuthKeyId: key.MediaTempAuthKeyId,
	}

	// Fix by @wuyun9527, 2018-12-21
	err := s.keyCreatedF(context.Background(), keyInfo, serverSalt.To_FutureSalt(), ctx.ExpiresIn)
	if err != nil {
		logx.Errorf("saveAuthKeyInfo not successful - auth_key_id:%d, err:%v", key.AuthKeyId, err)
		return false
	}
	return true
}

func calcNewNonceHash(newNonce, authKey []byte, b byte) []byte {
	authKeyAuxHash := make([]byte, len(newNonce))
	copy(authKeyAuxHash, newNonce)
	authKeyAuxHash = append(authKeyAuxHash, b)
	sha1D := sha1.Sum(authKey)
	authKeyAuxHash = append(authKeyAuxHash, sha1D[:]...)
	sha1E := sha1.Sum(authKeyAuxHash[:len(authKeyAuxHash)-12])
	authKeyAuxHash = append(authKeyAuxHash, sha1E[:]...)
	return authKeyAuxHash[len(authKeyAuxHash)-16:]
}

func checkSha1(data []byte, maxPaddingLen int) bool {
	var (
		dataLen  = len(data)
		sha1Data = data[:SHA_DIGEST_LENGTH]
	)

	for paddingLen := 0; paddingLen < maxPaddingLen; paddingLen++ {
		sha1Check := sha1.Sum(data[SHA_DIGEST_LENGTH : dataLen-paddingLen])
		if bytes.Equal(sha1Check[:], sha1Data) {
			return true
		}
	}
	return false
}
