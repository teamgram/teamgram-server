// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package main

import (
	"encoding/hex"
	"fmt"
	"github.com/teamgram/proto/mtproto/crypto"
)

var (
	sData = `ef554a47be44b8d00f0b33af20667a8d5020374d2af39e7b03f0c651c18177305fff832f238409cb9691c9fa6d8003362b4805f71722b6408fa8bd910b5e4b6493fd5a5e4f522f3c16fe15405b385179d1146cd6485efae597c4df15c07d635669ef18363a8f70f21506e80109412659470d8358262bea6d3089207a05e8f159bb6fa57a24c25d820c21617566f26f1c911211ca826c2be5548b0bd0336b3e551951ffdc1fc48c174bbed9488b33020c9ad9b926061baca4d78c95fdb3fe32652ec4e63fb6d15ee483d9849b741904037bc5864a9dc0f1a61c804df3e0bcb3412bdf9d8bacda28d8f0abdd7f006c34b6d99beb0e991013806e9edc01b948eeb2b599e732cccc312bfe2dd042009b0288d6c2b7b9dff818e6c560b2d60c3abdd0750457855a11b2a280b37509971f99e674c6a04d3bb1ed4602608997cc50436cb77553cd20a06687bcf861a11a031b918db16a917e3b7a695be47821315b160b67aac624c29b1da1e1f4f36ca29783f6f95532b568189a99fa86aeef88a2a12a0bb782b553fcfe4dd63f5801cef93441b07b77ee0c620a9f8577aaf0d823fe803112f26f6a9d15e2c55a67983dfe75cdbc1ddb1fafc178f47ccfd928ab2ab226900ac77aca0aea208eb05a57f41be264db0709829a4a84cb9cad947fda54447dd10999f168b71915e80c78af2d0a511a8eee10e1e311981c3c810f1fd14d17a33ce5ff7c5cedb268e76d464f3c8816c2006d48f6b7ba79b08c952d370f5781bd48a5bf3eba6d34135f265ab2b0b3dfb8a161500232334091f6185d74a801ff148c7ecec64ae804d1ff9f21d6af5a0a7a81a32c2af6c5aa410d2f20ce948cb751cf0854f3516e07f390249f51`
)

func main() {
	tData, _ := hex.DecodeString(sData)
	//idx := 0
	//_ = idx
	//fmt.Println(int64(binary.LittleEndian.Uint64(tData[1:])))
	//fmt.Println(hex.EncodeToString(tData[:1]))
	//tData = tData[1:]
	//fmt.Println(int(tData[0]))
	//n := int(tData[0] & 0x7f)
	//fmt.Println(n)
	//if n < 0x7f {
	//	n = n << 2
	//	tData = tData[1:]
	//	// idx = 1
	//} else {
	//	n = (int(tData[1]) | int(tData[2])<<8 | int(tData[3])<<16) << 2
	//	tData = tData[4:]
	//}
	//fmt.Println(n)
	//
	//fmt.Println(int64(binary.LittleEndian.Uint64(tData)))

	var (
		tmp           [64]byte
		obfuscatedBuf = tData[:64]
	)

	// 生成decrypt_key
	for i := 0; i < 48; i++ {
		tmp[i] = obfuscatedBuf[55-i]
	}

	e, err := crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48])
	_ = e
	if err != nil {
		fmt.Println("crypto.NewAesCTR128Encrypt(tmp[:32], tmp[32:48]): ", err)
		return
	}

	d, err := crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56])
	if err != nil {
		fmt.Println("crypto.NewAesCTR128Encrypt(obfuscatedBuf[8:40], obfuscatedBuf[40:56]): ", err)
	}

	d.Encrypt(obfuscatedBuf)
	fmt.Println(hex.EncodeToString(obfuscatedBuf[:]))

	//protocolType := binary.BigEndian.Uint32(obfuscatedBuf[56:])
	//if protocolType != ABRIDGED_INT32_FLAG &&
	//	protocolType != INTERMEDIATE_FLAG &&
	//	protocolType != PADDED_INTERMEDIATE_FLAG {
	//	return nil, fmt.Errorf("conn(%s) mtproto buf[56:60]'s byte != 0xef, received: %s",
	//		conn,
	//		hex.EncodeToString(obfuscatedBuf[56:60]))
	//}
	//
	//dcId := int16(binary.BigEndian.Uint16(obfuscatedBuf[60:]))
	// TODO: check dcId

}
