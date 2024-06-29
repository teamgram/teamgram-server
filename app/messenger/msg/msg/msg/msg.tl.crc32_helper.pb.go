/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package msg

const (
	CRC32_UNKNOWN                    TLConstructor = 0
	CRC32_sender                     TLConstructor = 1513645242  // 0x5a3864ba
	CRC32_outboxMessage              TLConstructor = 1402283185  // 0x539524b1
	CRC32_contentMessage             TLConstructor = -1301595468 // 0xb26b3ab4
	CRC32_msg_pushUserMessage        TLConstructor = 902887962   // 0x35d0fa1a
	CRC32_msg_readMessageContents    TLConstructor = 673481940   // 0x282484d4
	CRC32_msg_sendMessageV2          TLConstructor = -188056380  // 0xf4ca7cc4
	CRC32_msg_editMessage            TLConstructor = -2129725231 // 0x810ef8d1
	CRC32_msg_deleteMessages         TLConstructor = 568855069   // 0x21e80a1d
	CRC32_msg_deleteHistory          TLConstructor = 1975576778  // 0x75c0e8ca
	CRC32_msg_deletePhoneCallHistory TLConstructor = 649568574   // 0x26b7a13e
	CRC32_msg_deleteChatHistory      TLConstructor = -283155749  // 0xef1f62db
	CRC32_msg_readHistory            TLConstructor = 1510960658  // 0x5a0f6e12
	CRC32_msg_updatePinnedMessage    TLConstructor = -441560663  // 0xe5ae51a9
	CRC32_msg_unpinAllMessages       TLConstructor = -1199153371 // 0xb8865f25
)
