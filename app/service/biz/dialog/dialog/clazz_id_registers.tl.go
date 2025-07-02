/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package dialog

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(0x730ba93f, func() iface.TLObject { return &TLDialogExt{ClazzID: 0x730ba93f} })          // 0x730ba93f
	iface.RegisterClazzID(0xa6d498fe, func() iface.TLObject { return &TLDialogFilterExt{ClazzID: 0xa6d498fe} })    // 0xa6d498fe
	iface.RegisterClazzID(0xea7222c, func() iface.TLObject { return &TLDialogPinnedExt{ClazzID: 0xea7222c} })      // 0xea7222c
	iface.RegisterClazzID(0x1d59b45d, func() iface.TLObject { return &TLSimpleDialogsData{ClazzID: 0x1d59b45d} })  // 0x1d59b45d
	iface.RegisterClazzID(0xf6bdc4b2, func() iface.TLObject { return &TLUpdateDraftMessage{ClazzID: 0xf6bdc4b2} }) // 0xf6bdc4b2
	iface.RegisterClazzID(0x778fe85a, func() iface.TLObject { return &TLSavedDialogList{ClazzID: 0x778fe85a} })    // 0x778fe85a

	// Method
	iface.RegisterClazzID(0x4ecad99a, func() iface.TLObject { return &TLDialogSaveDraftMessage{ClazzID: 0x4ecad99a} })                  // 0x4ecad99a
	iface.RegisterClazzID(0xfb70b29a, func() iface.TLObject { return &TLDialogClearDraftMessage{ClazzID: 0xfb70b29a} })                 // 0xfb70b29a
	iface.RegisterClazzID(0xacde4fe6, func() iface.TLObject { return &TLDialogGetAllDrafts{ClazzID: 0xacde4fe6} })                      // 0xacde4fe6
	iface.RegisterClazzID(0x41b890fc, func() iface.TLObject { return &TLDialogClearAllDrafts{ClazzID: 0x41b890fc} })                    // 0x41b890fc
	iface.RegisterClazzID(0x4532910e, func() iface.TLObject { return &TLDialogMarkDialogUnread{ClazzID: 0x4532910e} })                  // 0x4532910e
	iface.RegisterClazzID(0x867ee52f, func() iface.TLObject { return &TLDialogToggleDialogPin{ClazzID: 0x867ee52f} })                   // 0x867ee52f
	iface.RegisterClazzID(0xcabc38f4, func() iface.TLObject { return &TLDialogGetDialogUnreadMarkList{ClazzID: 0xcabc38f4} })           // 0xcabc38f4
	iface.RegisterClazzID(0x9d7e8604, func() iface.TLObject { return &TLDialogGetDialogsByOffsetDate{ClazzID: 0x9d7e8604} })            // 0x9d7e8604
	iface.RegisterClazzID(0x860b1e16, func() iface.TLObject { return &TLDialogGetDialogs{ClazzID: 0x860b1e16} })                        // 0x860b1e16
	iface.RegisterClazzID(0xad258871, func() iface.TLObject { return &TLDialogGetDialogsByIdList{ClazzID: 0xad258871} })                // 0xad258871
	iface.RegisterClazzID(0xe039b465, func() iface.TLObject { return &TLDialogGetDialogsCount{ClazzID: 0xe039b465} })                   // 0xe039b465
	iface.RegisterClazzID(0xa8c21bb5, func() iface.TLObject { return &TLDialogGetPinnedDialogs{ClazzID: 0xa8c21bb5} })                  // 0xa8c21bb5
	iface.RegisterClazzID(0xfee33567, func() iface.TLObject { return &TLDialogReorderPinnedDialogs{ClazzID: 0xfee33567} })              // 0xfee33567
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
	iface.RegisterClazzID(0x2446869a, func() iface.TLObject { return &TLDialogEditPeerFolders{ClazzID: 0x2446869a} })                   // 0x2446869a
	iface.RegisterClazzID(0x28bd4d3b, func() iface.TLObject { return &TLDialogGetChannelMessageReadParticipants{ClazzID: 0x28bd4d3b} }) // 0x28bd4d3b
	iface.RegisterClazzID(0xe9aea22a, func() iface.TLObject { return &TLDialogSetChatTheme{ClazzID: 0xe9aea22a} })                      // 0xe9aea22a
	iface.RegisterClazzID(0x9d9b8ac, func() iface.TLObject { return &TLDialogSetHistoryTTL{ClazzID: 0x9d9b8ac} })                       // 0x9d9b8ac
	iface.RegisterClazzID(0x7ee08f03, func() iface.TLObject { return &TLDialogGetMyDialogsData{ClazzID: 0x7ee08f03} })                  // 0x7ee08f03
	iface.RegisterClazzID(0x38c1d668, func() iface.TLObject { return &TLDialogGetSavedDialogs{ClazzID: 0x38c1d668} })                   // 0x38c1d668
	iface.RegisterClazzID(0x40a3b7e7, func() iface.TLObject { return &TLDialogGetPinnedSavedDialogs{ClazzID: 0x40a3b7e7} })             // 0x40a3b7e7
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
