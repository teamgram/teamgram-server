/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialog

import (
	"reflect"

	"github.com/teamgram/proto/mtproto"
)

var _ *mtproto.Bool

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLDialogSaveDraftMessage":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_saveDraftMessage", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogClearDraftMessage":                 RPCContextTuple{"/mtproto.RPCDialog/dialog_clearDraftMessage", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetAllDrafts":                      RPCContextTuple{"/mtproto.RPCDialog/dialog_getAllDrafts", func() interface{} { return new(Vector_PeerWithDraftMessage) }},
	"TLDialogClearAllDrafts":                    RPCContextTuple{"/mtproto.RPCDialog/dialog_clearAllDrafts", func() interface{} { return new(Vector_PeerWithDraftMessage) }},
	"TLDialogMarkDialogUnread":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_markDialogUnread", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogToggleDialogPin":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_toggleDialogPin", func() interface{} { return new(mtproto.Int32) }},
	"TLDialogGetDialogUnreadMarkList":           RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogUnreadMarkList", func() interface{} { return new(Vector_DialogPeer) }},
	"TLDialogGetDialogsByOffsetDate":            RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogsByOffsetDate", func() interface{} { return new(Vector_DialogExt) }},
	"TLDialogGetDialogs":                        RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogs", func() interface{} { return new(Vector_DialogExt) }},
	"TLDialogGetDialogsByIdList":                RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogsByIdList", func() interface{} { return new(Vector_DialogExt) }},
	"TLDialogGetDialogsCount":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogsCount", func() interface{} { return new(mtproto.Int32) }},
	"TLDialogGetPinnedDialogs":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_getPinnedDialogs", func() interface{} { return new(Vector_DialogExt) }},
	"TLDialogReorderPinnedDialogs":              RPCContextTuple{"/mtproto.RPCDialog/dialog_reorderPinnedDialogs", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetDialogById":                     RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogById", func() interface{} { return new(DialogExt) }},
	"TLDialogGetTopMessage":                     RPCContextTuple{"/mtproto.RPCDialog/dialog_getTopMessage", func() interface{} { return new(mtproto.Int32) }},
	"TLDialogUpdateReadInbox":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_updateReadInbox", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogUpdateReadOutbox":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_updateReadOutbox", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogInsertOrUpdateDialog":              RPCContextTuple{"/mtproto.RPCDialog/dialog_insertOrUpdateDialog", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogDeleteDialog":                      RPCContextTuple{"/mtproto.RPCDialog/dialog_deleteDialog", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetUserPinnedMessage":              RPCContextTuple{"/mtproto.RPCDialog/dialog_getUserPinnedMessage", func() interface{} { return new(mtproto.Int32) }},
	"TLDialogUpdateUserPinnedMessage":           RPCContextTuple{"/mtproto.RPCDialog/dialog_updateUserPinnedMessage", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogInsertOrUpdateDialogFilter":        RPCContextTuple{"/mtproto.RPCDialog/dialog_insertOrUpdateDialogFilter", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogDeleteDialogFilter":                RPCContextTuple{"/mtproto.RPCDialog/dialog_deleteDialogFilter", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogUpdateDialogFiltersOrder":          RPCContextTuple{"/mtproto.RPCDialog/dialog_updateDialogFiltersOrder", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetDialogFilters":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogFilters", func() interface{} { return new(Vector_DialogFilterExt) }},
	"TLDialogGetDialogFolder":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogFolder", func() interface{} { return new(Vector_DialogExt) }},
	"TLDialogEditPeerFolders":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_editPeerFolders", func() interface{} { return new(Vector_DialogPinnedExt) }},
	"TLDialogGetChannelMessageReadParticipants": RPCContextTuple{"/mtproto.RPCDialog/dialog_getChannelMessageReadParticipants", func() interface{} { return new(Vector_Long) }},
	"TLDialogSetChatTheme":                      RPCContextTuple{"/mtproto.RPCDialog/dialog_setChatTheme", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogSetHistoryTTL":                     RPCContextTuple{"/mtproto.RPCDialog/dialog_setHistoryTTL", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetMyDialogsData":                  RPCContextTuple{"/mtproto.RPCDialog/dialog_getMyDialogsData", func() interface{} { return new(DialogsData) }},
	"TLDialogGetSavedDialogs":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_getSavedDialogs", func() interface{} { return new(SavedDialogList) }},
	"TLDialogGetPinnedSavedDialogs":             RPCContextTuple{"/mtproto.RPCDialog/dialog_getPinnedSavedDialogs", func() interface{} { return new(SavedDialogList) }},
	"TLDialogToggleSavedDialogPin":              RPCContextTuple{"/mtproto.RPCDialog/dialog_toggleSavedDialogPin", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogReorderPinnedSavedDialogs":         RPCContextTuple{"/mtproto.RPCDialog/dialog_reorderPinnedSavedDialogs", func() interface{} { return new(mtproto.Bool) }},
	"TLDialogGetDialogFilter":                   RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogFilter", func() interface{} { return new(DialogFilterExt) }},
	"TLDialogGetDialogFilterBySlug":             RPCContextTuple{"/mtproto.RPCDialog/dialog_getDialogFilterBySlug", func() interface{} { return new(DialogFilterExt) }},
	"TLDialogCreateDialogFilter":                RPCContextTuple{"/mtproto.RPCDialog/dialog_createDialogFilter", func() interface{} { return new(DialogFilterExt) }},
	"TLDialogUpdateUnreadCount":                 RPCContextTuple{"/mtproto.RPCDialog/dialog_updateUnreadCount", func() interface{} { return new(mtproto.Bool) }},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		// log.Errorf("Can't find name: %s", rt.Name())
		return nil
	}
	return &m
}

func GetRPCContextRegisters() map[string]RPCContextTuple {
	return rpcContextRegisters
}
