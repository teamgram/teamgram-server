/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msgtransfer

const (
	CRC32_UNKNOWN                         TLConstructor = 0
	CRC32_sender                          TLConstructor = 1513645242  // 0x5a3864ba
	CRC32_outboxMessage                   TLConstructor = 1763737728  // 0x69208080
	CRC32_contentMessage                  TLConstructor = -1922780877 // 0x8d64b133
	CRC32_inboxMessageData                TLConstructor = 1002286548  // 0x3bbdadd4
	CRC32_inboxMessageId                  TLConstructor = -963460705  // 0xc692c19f
	CRC32_msgtransfer_sendMessageToOutbox TLConstructor = -508367556  // 0xe1b2ed3c
	CRC32_msgtransfer_sendMessageToInbox  TLConstructor = -750661413  // 0xd341d0db
)
