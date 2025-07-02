/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialog

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_updateDraftMessage                       = "updateDraftMessage"
	ClazzName_dialogExt                                = "dialogExt"
	ClazzName_dialogPinnedExt                          = "dialogPinnedExt"
	ClazzName_dialogFilterExt                          = "dialogFilterExt"
	ClazzName_simpleDialogsData                        = "simpleDialogsData"
	ClazzName_savedDialogList                          = "savedDialogList"
	ClazzName_dialog_saveDraftMessage                  = "dialog_saveDraftMessage"
	ClazzName_dialog_clearDraftMessage                 = "dialog_clearDraftMessage"
	ClazzName_dialog_getAllDrafts                      = "dialog_getAllDrafts"
	ClazzName_dialog_clearAllDrafts                    = "dialog_clearAllDrafts"
	ClazzName_dialog_markDialogUnread                  = "dialog_markDialogUnread"
	ClazzName_dialog_toggleDialogPin                   = "dialog_toggleDialogPin"
	ClazzName_dialog_getDialogUnreadMarkList           = "dialog_getDialogUnreadMarkList"
	ClazzName_dialog_getDialogsByOffsetDate            = "dialog_getDialogsByOffsetDate"
	ClazzName_dialog_getDialogs                        = "dialog_getDialogs"
	ClazzName_dialog_getDialogsByIdList                = "dialog_getDialogsByIdList"
	ClazzName_dialog_getDialogsCount                   = "dialog_getDialogsCount"
	ClazzName_dialog_getPinnedDialogs                  = "dialog_getPinnedDialogs"
	ClazzName_dialog_reorderPinnedDialogs              = "dialog_reorderPinnedDialogs"
	ClazzName_dialog_getDialogById                     = "dialog_getDialogById"
	ClazzName_dialog_getTopMessage                     = "dialog_getTopMessage"
	ClazzName_dialog_insertOrUpdateDialog              = "dialog_insertOrUpdateDialog"
	ClazzName_dialog_deleteDialog                      = "dialog_deleteDialog"
	ClazzName_dialog_getUserPinnedMessage              = "dialog_getUserPinnedMessage"
	ClazzName_dialog_updateUserPinnedMessage           = "dialog_updateUserPinnedMessage"
	ClazzName_dialog_insertOrUpdateDialogFilter        = "dialog_insertOrUpdateDialogFilter"
	ClazzName_dialog_deleteDialogFilter                = "dialog_deleteDialogFilter"
	ClazzName_dialog_updateDialogFiltersOrder          = "dialog_updateDialogFiltersOrder"
	ClazzName_dialog_getDialogFilters                  = "dialog_getDialogFilters"
	ClazzName_dialog_getDialogFolder                   = "dialog_getDialogFolder"
	ClazzName_dialog_editPeerFolders                   = "dialog_editPeerFolders"
	ClazzName_dialog_getChannelMessageReadParticipants = "dialog_getChannelMessageReadParticipants"
	ClazzName_dialog_setChatTheme                      = "dialog_setChatTheme"
	ClazzName_dialog_setHistoryTTL                     = "dialog_setHistoryTTL"
	ClazzName_dialog_getMyDialogsData                  = "dialog_getMyDialogsData"
	ClazzName_dialog_getSavedDialogs                   = "dialog_getSavedDialogs"
	ClazzName_dialog_getPinnedSavedDialogs             = "dialog_getPinnedSavedDialogs"
	ClazzName_dialog_toggleSavedDialogPin              = "dialog_toggleSavedDialogPin"
	ClazzName_dialog_reorderPinnedSavedDialogs         = "dialog_reorderPinnedSavedDialogs"
	ClazzName_dialog_getDialogFilter                   = "dialog_getDialogFilter"
	ClazzName_dialog_getDialogFilterBySlug             = "dialog_getDialogFilterBySlug"
	ClazzName_dialog_createDialogFilter                = "dialog_createDialogFilter"
	ClazzName_dialog_updateUnreadCount                 = "dialog_updateUnreadCount"
	ClazzName_dialog_toggleDialogFilterTags            = "dialog_toggleDialogFilterTags"
	ClazzName_dialog_getDialogFilterTags               = "dialog_getDialogFilterTags"
	ClazzName_dialog_setChatWallpaper                  = "dialog_setChatWallpaper"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_updateDraftMessage, 0, 0xf6bdc4b2)                       // f6bdc4b2
	iface.RegisterClazzName(ClazzName_dialogExt, 0, 0x730ba93f)                                // 730ba93f
	iface.RegisterClazzName(ClazzName_dialogPinnedExt, 0, 0xea7222c)                           // ea7222c
	iface.RegisterClazzName(ClazzName_dialogFilterExt, 0, 0xa6d498fe)                          // a6d498fe
	iface.RegisterClazzName(ClazzName_simpleDialogsData, 0, 0x1d59b45d)                        // 1d59b45d
	iface.RegisterClazzName(ClazzName_savedDialogList, 0, 0x778fe85a)                          // 778fe85a
	iface.RegisterClazzName(ClazzName_dialog_saveDraftMessage, 0, 0x4ecad99a)                  // 4ecad99a
	iface.RegisterClazzName(ClazzName_dialog_clearDraftMessage, 0, 0xfb70b29a)                 // fb70b29a
	iface.RegisterClazzName(ClazzName_dialog_getAllDrafts, 0, 0xacde4fe6)                      // acde4fe6
	iface.RegisterClazzName(ClazzName_dialog_clearAllDrafts, 0, 0x41b890fc)                    // 41b890fc
	iface.RegisterClazzName(ClazzName_dialog_markDialogUnread, 0, 0x4532910e)                  // 4532910e
	iface.RegisterClazzName(ClazzName_dialog_toggleDialogPin, 0, 0x867ee52f)                   // 867ee52f
	iface.RegisterClazzName(ClazzName_dialog_getDialogUnreadMarkList, 0, 0xcabc38f4)           // cabc38f4
	iface.RegisterClazzName(ClazzName_dialog_getDialogsByOffsetDate, 0, 0x9d7e8604)            // 9d7e8604
	iface.RegisterClazzName(ClazzName_dialog_getDialogs, 0, 0x860b1e16)                        // 860b1e16
	iface.RegisterClazzName(ClazzName_dialog_getDialogsByIdList, 0, 0xad258871)                // ad258871
	iface.RegisterClazzName(ClazzName_dialog_getDialogsCount, 0, 0xe039b465)                   // e039b465
	iface.RegisterClazzName(ClazzName_dialog_getPinnedDialogs, 0, 0xa8c21bb5)                  // a8c21bb5
	iface.RegisterClazzName(ClazzName_dialog_reorderPinnedDialogs, 0, 0xfee33567)              // fee33567
	iface.RegisterClazzName(ClazzName_dialog_getDialogById, 0, 0xa15f3bf5)                     // a15f3bf5
	iface.RegisterClazzName(ClazzName_dialog_getTopMessage, 0, 0xfa7db272)                     // fa7db272
	iface.RegisterClazzName(ClazzName_dialog_insertOrUpdateDialog, 0, 0x5d2b8822)              // 5d2b8822
	iface.RegisterClazzName(ClazzName_dialog_deleteDialog, 0, 0x1b31de3)                       // 1b31de3
	iface.RegisterClazzName(ClazzName_dialog_getUserPinnedMessage, 0, 0x8f9bc2b1)              // 8f9bc2b1
	iface.RegisterClazzName(ClazzName_dialog_updateUserPinnedMessage, 0, 0x1622f22a)           // 1622f22a
	iface.RegisterClazzName(ClazzName_dialog_insertOrUpdateDialogFilter, 0, 0xaa8a384)         // aa8a384
	iface.RegisterClazzName(ClazzName_dialog_deleteDialogFilter, 0, 0x1dd3e97)                 // 1dd3e97
	iface.RegisterClazzName(ClazzName_dialog_updateDialogFiltersOrder, 0, 0xb13c0b3f)          // b13c0b3f
	iface.RegisterClazzName(ClazzName_dialog_getDialogFilters, 0, 0x6c676c3c)                  // 6c676c3c
	iface.RegisterClazzName(ClazzName_dialog_getDialogFolder, 0, 0x411b8eb5)                   // 411b8eb5
	iface.RegisterClazzName(ClazzName_dialog_editPeerFolders, 0, 0x2446869a)                   // 2446869a
	iface.RegisterClazzName(ClazzName_dialog_getChannelMessageReadParticipants, 0, 0x28bd4d3b) // 28bd4d3b
	iface.RegisterClazzName(ClazzName_dialog_setChatTheme, 0, 0xe9aea22a)                      // e9aea22a
	iface.RegisterClazzName(ClazzName_dialog_setHistoryTTL, 0, 0x9d9b8ac)                      // 9d9b8ac
	iface.RegisterClazzName(ClazzName_dialog_getMyDialogsData, 0, 0x7ee08f03)                  // 7ee08f03
	iface.RegisterClazzName(ClazzName_dialog_getSavedDialogs, 0, 0x38c1d668)                   // 38c1d668
	iface.RegisterClazzName(ClazzName_dialog_getPinnedSavedDialogs, 0, 0x40a3b7e7)             // 40a3b7e7
	iface.RegisterClazzName(ClazzName_dialog_toggleSavedDialogPin, 0, 0x44f317d9)              // 44f317d9
	iface.RegisterClazzName(ClazzName_dialog_reorderPinnedSavedDialogs, 0, 0xd85ccbd2)         // d85ccbd2
	iface.RegisterClazzName(ClazzName_dialog_getDialogFilter, 0, 0xf388061c)                   // f388061c
	iface.RegisterClazzName(ClazzName_dialog_getDialogFilterBySlug, 0, 0x4e457fef)             // 4e457fef
	iface.RegisterClazzName(ClazzName_dialog_createDialogFilter, 0, 0xc6cb636f)                // c6cb636f
	iface.RegisterClazzName(ClazzName_dialog_updateUnreadCount, 0, 0x2bac334d)                 // 2bac334d
	iface.RegisterClazzName(ClazzName_dialog_toggleDialogFilterTags, 0, 0xa0cd6d89)            // a0cd6d89
	iface.RegisterClazzName(ClazzName_dialog_getDialogFilterTags, 0, 0xfaf0fa97)               // faf0fa97
	iface.RegisterClazzName(ClazzName_dialog_setChatWallpaper, 0, 0xb551db12)                  // b551db12

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_updateDraftMessage, 0xf6bdc4b2)                       // f6bdc4b2
	iface.RegisterClazzIDName(ClazzName_dialogExt, 0x730ba93f)                                // 730ba93f
	iface.RegisterClazzIDName(ClazzName_dialogPinnedExt, 0xea7222c)                           // ea7222c
	iface.RegisterClazzIDName(ClazzName_dialogFilterExt, 0xa6d498fe)                          // a6d498fe
	iface.RegisterClazzIDName(ClazzName_simpleDialogsData, 0x1d59b45d)                        // 1d59b45d
	iface.RegisterClazzIDName(ClazzName_savedDialogList, 0x778fe85a)                          // 778fe85a
	iface.RegisterClazzIDName(ClazzName_dialog_saveDraftMessage, 0x4ecad99a)                  // 4ecad99a
	iface.RegisterClazzIDName(ClazzName_dialog_clearDraftMessage, 0xfb70b29a)                 // fb70b29a
	iface.RegisterClazzIDName(ClazzName_dialog_getAllDrafts, 0xacde4fe6)                      // acde4fe6
	iface.RegisterClazzIDName(ClazzName_dialog_clearAllDrafts, 0x41b890fc)                    // 41b890fc
	iface.RegisterClazzIDName(ClazzName_dialog_markDialogUnread, 0x4532910e)                  // 4532910e
	iface.RegisterClazzIDName(ClazzName_dialog_toggleDialogPin, 0x867ee52f)                   // 867ee52f
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogUnreadMarkList, 0xcabc38f4)           // cabc38f4
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogsByOffsetDate, 0x9d7e8604)            // 9d7e8604
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogs, 0x860b1e16)                        // 860b1e16
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogsByIdList, 0xad258871)                // ad258871
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogsCount, 0xe039b465)                   // e039b465
	iface.RegisterClazzIDName(ClazzName_dialog_getPinnedDialogs, 0xa8c21bb5)                  // a8c21bb5
	iface.RegisterClazzIDName(ClazzName_dialog_reorderPinnedDialogs, 0xfee33567)              // fee33567
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogById, 0xa15f3bf5)                     // a15f3bf5
	iface.RegisterClazzIDName(ClazzName_dialog_getTopMessage, 0xfa7db272)                     // fa7db272
	iface.RegisterClazzIDName(ClazzName_dialog_insertOrUpdateDialog, 0x5d2b8822)              // 5d2b8822
	iface.RegisterClazzIDName(ClazzName_dialog_deleteDialog, 0x1b31de3)                       // 1b31de3
	iface.RegisterClazzIDName(ClazzName_dialog_getUserPinnedMessage, 0x8f9bc2b1)              // 8f9bc2b1
	iface.RegisterClazzIDName(ClazzName_dialog_updateUserPinnedMessage, 0x1622f22a)           // 1622f22a
	iface.RegisterClazzIDName(ClazzName_dialog_insertOrUpdateDialogFilter, 0xaa8a384)         // aa8a384
	iface.RegisterClazzIDName(ClazzName_dialog_deleteDialogFilter, 0x1dd3e97)                 // 1dd3e97
	iface.RegisterClazzIDName(ClazzName_dialog_updateDialogFiltersOrder, 0xb13c0b3f)          // b13c0b3f
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogFilters, 0x6c676c3c)                  // 6c676c3c
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogFolder, 0x411b8eb5)                   // 411b8eb5
	iface.RegisterClazzIDName(ClazzName_dialog_editPeerFolders, 0x2446869a)                   // 2446869a
	iface.RegisterClazzIDName(ClazzName_dialog_getChannelMessageReadParticipants, 0x28bd4d3b) // 28bd4d3b
	iface.RegisterClazzIDName(ClazzName_dialog_setChatTheme, 0xe9aea22a)                      // e9aea22a
	iface.RegisterClazzIDName(ClazzName_dialog_setHistoryTTL, 0x9d9b8ac)                      // 9d9b8ac
	iface.RegisterClazzIDName(ClazzName_dialog_getMyDialogsData, 0x7ee08f03)                  // 7ee08f03
	iface.RegisterClazzIDName(ClazzName_dialog_getSavedDialogs, 0x38c1d668)                   // 38c1d668
	iface.RegisterClazzIDName(ClazzName_dialog_getPinnedSavedDialogs, 0x40a3b7e7)             // 40a3b7e7
	iface.RegisterClazzIDName(ClazzName_dialog_toggleSavedDialogPin, 0x44f317d9)              // 44f317d9
	iface.RegisterClazzIDName(ClazzName_dialog_reorderPinnedSavedDialogs, 0xd85ccbd2)         // d85ccbd2
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogFilter, 0xf388061c)                   // f388061c
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogFilterBySlug, 0x4e457fef)             // 4e457fef
	iface.RegisterClazzIDName(ClazzName_dialog_createDialogFilter, 0xc6cb636f)                // c6cb636f
	iface.RegisterClazzIDName(ClazzName_dialog_updateUnreadCount, 0x2bac334d)                 // 2bac334d
	iface.RegisterClazzIDName(ClazzName_dialog_toggleDialogFilterTags, 0xa0cd6d89)            // a0cd6d89
	iface.RegisterClazzIDName(ClazzName_dialog_getDialogFilterTags, 0xfaf0fa97)               // faf0fa97
	iface.RegisterClazzIDName(ClazzName_dialog_setChatWallpaper, 0xb551db12)                  // b551db12
}
