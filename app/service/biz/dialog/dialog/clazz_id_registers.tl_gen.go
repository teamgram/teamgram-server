/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

// ConstructorList
// RequestList

package dialog

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x5b146008, func() iface.TLObject { return &TLDialogCursor{ClazzID: 0x5b146008} })       // 0x5b146008
	iface.RegisterClazzID(0x730ba93f, func() iface.TLObject { return &TLDialogExt{ClazzID: 0x730ba93f} })          // 0x730ba93f
	iface.RegisterClazzID(0x7c9d7c44, func() iface.TLObject { return &TLDialogExtV2{ClazzID: 0x7c9d7c44} })        // 0x7c9d7c44
	iface.RegisterClazzID(0x7a69125, func() iface.TLObject { return &TLDialogExtras{ClazzID: 0x7a69125} })         // 0x7a69125
	iface.RegisterClazzID(0xa6d498fe, func() iface.TLObject { return &TLDialogFilterExt{ClazzID: 0xa6d498fe} })    // 0xa6d498fe
	iface.RegisterClazzID(0x92dbd5aa, func() iface.TLObject { return &TLDialogPage{ClazzID: 0x92dbd5aa} })         // 0x92dbd5aa
	iface.RegisterClazzID(0xb7789d79, func() iface.TLObject { return &TLDialogPeer{ClazzID: 0xb7789d79} })         // 0xb7789d79
	iface.RegisterClazzID(0xea7222c, func() iface.TLObject { return &TLDialogPinnedExt{ClazzID: 0xea7222c} })      // 0xea7222c
	iface.RegisterClazzID(0x1d59b45d, func() iface.TLObject { return &TLSimpleDialogsData{ClazzID: 0x1d59b45d} })  // 0x1d59b45d
	iface.RegisterClazzID(0xf6bdc4b2, func() iface.TLObject { return &TLUpdateDraftMessage{ClazzID: 0xf6bdc4b2} }) // 0xf6bdc4b2
	iface.RegisterClazzID(0x778fe85a, func() iface.TLObject { return &TLSavedDialogList{ClazzID: 0x778fe85a} })    // 0x778fe85a

	// Method
	iface.RegisterClazzID(0x3bfb9d31, func() iface.TLObject { return &TLDialogSaveDraftMessage{ClazzID: 0x3bfb9d31} })                  // 0x3bfb9d31
	iface.RegisterClazzID(0x9390695, func() iface.TLObject { return &TLDialogClearDraftMessage{ClazzID: 0x9390695} })                   // 0x9390695
	iface.RegisterClazzID(0xda48bb0b, func() iface.TLObject { return &TLDialogClearDraftAfterSend{ClazzID: 0xda48bb0b} })               // 0xda48bb0b
	iface.RegisterClazzID(0xacde4fe6, func() iface.TLObject { return &TLDialogGetAllDrafts{ClazzID: 0xacde4fe6} })                      // 0xacde4fe6
	iface.RegisterClazzID(0x8432b418, func() iface.TLObject { return &TLDialogClearAllDrafts{ClazzID: 0x8432b418} })                    // 0x8432b418
	iface.RegisterClazzID(0x4532910e, func() iface.TLObject { return &TLDialogMarkDialogUnread{ClazzID: 0x4532910e} })                  // 0x4532910e
	iface.RegisterClazzID(0x6ad45bb4, func() iface.TLObject { return &TLDialogToggleDialogPin{ClazzID: 0x6ad45bb4} })                   // 0x6ad45bb4
	iface.RegisterClazzID(0xcabc38f4, func() iface.TLObject { return &TLDialogGetDialogUnreadMarkList{ClazzID: 0xcabc38f4} })           // 0xcabc38f4
	iface.RegisterClazzID(0x9d7e8604, func() iface.TLObject { return &TLDialogGetDialogsByOffsetDate{ClazzID: 0x9d7e8604} })            // 0x9d7e8604
	iface.RegisterClazzID(0x860b1e16, func() iface.TLObject { return &TLDialogGetDialogs{ClazzID: 0x860b1e16} })                        // 0x860b1e16
	iface.RegisterClazzID(0xad258871, func() iface.TLObject { return &TLDialogGetDialogsByIdList{ClazzID: 0xad258871} })                // 0xad258871
	iface.RegisterClazzID(0xe039b465, func() iface.TLObject { return &TLDialogGetDialogsCount{ClazzID: 0xe039b465} })                   // 0xe039b465
	iface.RegisterClazzID(0xa8c21bb5, func() iface.TLObject { return &TLDialogGetPinnedDialogs{ClazzID: 0xa8c21bb5} })                  // 0xa8c21bb5
	iface.RegisterClazzID(0x3aff2348, func() iface.TLObject { return &TLDialogReorderPinnedDialogs{ClazzID: 0x3aff2348} })              // 0x3aff2348
	iface.RegisterClazzID(0xa15f3bf5, func() iface.TLObject { return &TLDialogGetDialogById{ClazzID: 0xa15f3bf5} })                     // 0xa15f3bf5
	iface.RegisterClazzID(0xfa7db272, func() iface.TLObject { return &TLDialogGetTopMessage{ClazzID: 0xfa7db272} })                     // 0xfa7db272
	iface.RegisterClazzID(0x5d2b8822, func() iface.TLObject { return &TLDialogInsertOrUpdateDialog{ClazzID: 0x5d2b8822} })              // 0x5d2b8822
	iface.RegisterClazzID(0x1b31de3, func() iface.TLObject { return &TLDialogDeleteDialog{ClazzID: 0x1b31de3} })                        // 0x1b31de3
	iface.RegisterClazzID(0x8f9bc2b1, func() iface.TLObject { return &TLDialogGetUserPinnedMessage{ClazzID: 0x8f9bc2b1} })              // 0x8f9bc2b1
	iface.RegisterClazzID(0x1622f22a, func() iface.TLObject { return &TLDialogUpdateUserPinnedMessage{ClazzID: 0x1622f22a} })           // 0x1622f22a
	iface.RegisterClazzID(0xaa8a384, func() iface.TLObject { return &TLDialogInsertOrUpdateDialogFilter{ClazzID: 0xaa8a384} })          // 0xaa8a384
	iface.RegisterClazzID(0x1dd3e97, func() iface.TLObject { return &TLDialogDeleteDialogFilter{ClazzID: 0x1dd3e97} })                  // 0x1dd3e97
	iface.RegisterClazzID(0xb13c0b3f, func() iface.TLObject { return &TLDialogUpdateDialogFiltersOrder{ClazzID: 0xb13c0b3f} })          // 0xb13c0b3f
	iface.RegisterClazzID(0x6c676c3c, func() iface.TLObject { return &TLDialogGetDialogFilters{ClazzID: 0x6c676c3c} })                  // 0x6c676c3c
	iface.RegisterClazzID(0x411b8eb5, func() iface.TLObject { return &TLDialogGetDialogFolder{ClazzID: 0x411b8eb5} })                   // 0x411b8eb5
	iface.RegisterClazzID(0xfbe6f2f, func() iface.TLObject { return &TLDialogEditPeerFolders{ClazzID: 0xfbe6f2f} })                     // 0xfbe6f2f
	iface.RegisterClazzID(0x10dbef02, func() iface.TLObject { return &TLDialogGetDialogsV2{ClazzID: 0x10dbef02} })                      // 0x10dbef02
	iface.RegisterClazzID(0xee61c04d, func() iface.TLObject { return &TLDialogGetPeerDialogsV2{ClazzID: 0xee61c04d} })                  // 0xee61c04d
	iface.RegisterClazzID(0x74909bab, func() iface.TLObject { return &TLDialogGetPinnedDialogsV2{ClazzID: 0x74909bab} })                // 0x74909bab
	iface.RegisterClazzID(0xfb112ce3, func() iface.TLObject { return &TLDialogGetDialogByPeerV2{ClazzID: 0xfb112ce3} })                 // 0xfb112ce3
	iface.RegisterClazzID(0xa393a92e, func() iface.TLObject { return &TLDialogBatchGetDialogExtras{ClazzID: 0xa393a92e} })              // 0xa393a92e
	iface.RegisterClazzID(0x28bd4d3b, func() iface.TLObject { return &TLDialogGetChannelMessageReadParticipants{ClazzID: 0x28bd4d3b} }) // 0x28bd4d3b
	iface.RegisterClazzID(0xe9aea22a, func() iface.TLObject { return &TLDialogSetChatTheme{ClazzID: 0xe9aea22a} })                      // 0xe9aea22a
	iface.RegisterClazzID(0x9d9b8ac, func() iface.TLObject { return &TLDialogSetHistoryTTL{ClazzID: 0x9d9b8ac} })                       // 0x9d9b8ac
	iface.RegisterClazzID(0x7ee08f03, func() iface.TLObject { return &TLDialogGetMyDialogsData{ClazzID: 0x7ee08f03} })                  // 0x7ee08f03
	iface.RegisterClazzID(0x38c1d668, func() iface.TLObject { return &TLDialogGetSavedDialogs{ClazzID: 0x38c1d668} })                   // 0x38c1d668
	iface.RegisterClazzID(0x40a3b7e7, func() iface.TLObject { return &TLDialogGetPinnedSavedDialogs{ClazzID: 0x40a3b7e7} })             // 0x40a3b7e7
	iface.RegisterClazzID(0x9e3ab43c, func() iface.TLObject { return &TLDialogUpsertSavedDialogFromMessage{ClazzID: 0x9e3ab43c} })      // 0x9e3ab43c
	iface.RegisterClazzID(0x44f317d9, func() iface.TLObject { return &TLDialogToggleSavedDialogPin{ClazzID: 0x44f317d9} })              // 0x44f317d9
	iface.RegisterClazzID(0xd85ccbd2, func() iface.TLObject { return &TLDialogReorderPinnedSavedDialogs{ClazzID: 0xd85ccbd2} })         // 0xd85ccbd2
	iface.RegisterClazzID(0xf388061c, func() iface.TLObject { return &TLDialogGetDialogFilter{ClazzID: 0xf388061c} })                   // 0xf388061c
	iface.RegisterClazzID(0x4e457fef, func() iface.TLObject { return &TLDialogGetDialogFilterBySlug{ClazzID: 0x4e457fef} })             // 0x4e457fef
	iface.RegisterClazzID(0xc6cb636f, func() iface.TLObject { return &TLDialogCreateDialogFilter{ClazzID: 0xc6cb636f} })                // 0xc6cb636f
	iface.RegisterClazzID(0x2bac334d, func() iface.TLObject { return &TLDialogUpdateUnreadCount{ClazzID: 0x2bac334d} })                 // 0x2bac334d
	iface.RegisterClazzID(0xa0cd6d89, func() iface.TLObject { return &TLDialogToggleDialogFilterTags{ClazzID: 0xa0cd6d89} })            // 0xa0cd6d89
	iface.RegisterClazzID(0xfaf0fa97, func() iface.TLObject { return &TLDialogGetDialogFilterTags{ClazzID: 0xfaf0fa97} })               // 0xfaf0fa97
	iface.RegisterClazzID(0xb551db12, func() iface.TLObject { return &TLDialogSetChatWallpaper{ClazzID: 0xb551db12} })                  // 0xb551db12
}
