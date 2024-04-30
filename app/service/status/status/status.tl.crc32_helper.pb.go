/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package status

const (
	CRC32_UNKNOWN                           TLConstructor = 0
	CRC32_sessionEntry                      TLConstructor = 392473649   // 0x1764ac31
	CRC32_userSessionEntryList              TLConstructor = -269700200  // 0xefecb398
	CRC32_status_setSessionOnline           TLConstructor = 1381075919  // 0x52518bcf
	CRC32_status_setSessionOffline          TLConstructor = 631663196   // 0x25a66a5c
	CRC32_status_getUserOnlineSessions      TLConstructor = -406788659  // 0xe7c0e5cd
	CRC32_status_getUsersOnlineSessionsList TLConstructor = -2009385532 // 0x883b35c4
	CRC32_status_getChannelOnlineUsers      TLConstructor = 1166257237  // 0x4583ac55
	CRC32_status_setUserChannelsOnline      TLConstructor = -851901363  // 0xcd39044d
	CRC32_status_setUserChannelsOffline     TLConstructor = 1822646698  // 0x6ca361aa
	CRC32_status_setChannelUserOffline      TLConstructor = -997471364  // 0xc48bcb7c
	CRC32_status_setChannelUsersOnline      TLConstructor = -1499734793 // 0xa69bdcf7
	CRC32_status_setChannelOffline          TLConstructor = 1266112245  // 0x4b7756f5
)
