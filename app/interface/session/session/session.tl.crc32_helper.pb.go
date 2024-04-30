/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

const (
	CRC32_UNKNOWN                        TLConstructor = 0
	CRC32_sessionClientEvent             TLConstructor = -739769057  // 0xd3e8051f
	CRC32_sessionClientData              TLConstructor = 825806990   // 0x3138d08e
	CRC32_httpSessionData                TLConstructor = -606579889  // 0xdbd8534f
	CRC32_session_queryAuthKey           TLConstructor = 1798174801  // 0x6b2df851
	CRC32_session_setAuthKey             TLConstructor = 487672075   // 0x1d11490b
	CRC32_session_createSession          TLConstructor = 1091351053  // 0x410cb20d
	CRC32_session_sendDataToSession      TLConstructor = -2023019028 // 0x876b2dec
	CRC32_session_sendHttpDataToSession  TLConstructor = -1142152274 // 0xbbec23ae
	CRC32_session_closeSession           TLConstructor = 393200211   // 0x176fc253
	CRC32_session_pushUpdatesData        TLConstructor = 1075152191  // 0x4015853f
	CRC32_session_pushSessionUpdatesData TLConstructor = 106898165   // 0x65f22f5
	CRC32_session_pushRpcResultData      TLConstructor = 556344000   // 0x212922c0
)
