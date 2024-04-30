/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package chat

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
	"TLChatGetMutableChat":                   RPCContextTuple{"/mtproto.RPCChat/chat_getMutableChat", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatGetChatListByIdList":              RPCContextTuple{"/mtproto.RPCChat/chat_getChatListByIdList", func() interface{} { return new(Vector_MutableChat) }},
	"TLChatGetChatBySelfId":                  RPCContextTuple{"/mtproto.RPCChat/chat_getChatBySelfId", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatCreateChat2":                      RPCContextTuple{"/mtproto.RPCChat/chat_createChat2", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatDeleteChat":                       RPCContextTuple{"/mtproto.RPCChat/chat_deleteChat", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatDeleteChatUser":                   RPCContextTuple{"/mtproto.RPCChat/chat_deleteChatUser", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatEditChatTitle":                    RPCContextTuple{"/mtproto.RPCChat/chat_editChatTitle", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatEditChatAbout":                    RPCContextTuple{"/mtproto.RPCChat/chat_editChatAbout", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatEditChatPhoto":                    RPCContextTuple{"/mtproto.RPCChat/chat_editChatPhoto", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatEditChatAdmin":                    RPCContextTuple{"/mtproto.RPCChat/chat_editChatAdmin", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatEditChatDefaultBannedRights":      RPCContextTuple{"/mtproto.RPCChat/chat_editChatDefaultBannedRights", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatAddChatUser":                      RPCContextTuple{"/mtproto.RPCChat/chat_addChatUser", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatGetMutableChatByLink":             RPCContextTuple{"/mtproto.RPCChat/chat_getMutableChatByLink", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatToggleNoForwards":                 RPCContextTuple{"/mtproto.RPCChat/chat_toggleNoForwards", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatMigratedToChannel":                RPCContextTuple{"/mtproto.RPCChat/chat_migratedToChannel", func() interface{} { return new(mtproto.Bool) }},
	"TLChatGetChatParticipantIdList":         RPCContextTuple{"/mtproto.RPCChat/chat_getChatParticipantIdList", func() interface{} { return new(Vector_Long) }},
	"TLChatGetUsersChatIdList":               RPCContextTuple{"/mtproto.RPCChat/chat_getUsersChatIdList", func() interface{} { return new(Vector_UserChatIdList) }},
	"TLChatGetMyChatList":                    RPCContextTuple{"/mtproto.RPCChat/chat_getMyChatList", func() interface{} { return new(Vector_MutableChat) }},
	"TLChatExportChatInvite":                 RPCContextTuple{"/mtproto.RPCChat/chat_exportChatInvite", func() interface{} { return new(mtproto.ExportedChatInvite) }},
	"TLChatGetAdminsWithInvites":             RPCContextTuple{"/mtproto.RPCChat/chat_getAdminsWithInvites", func() interface{} { return new(Vector_ChatAdminWithInvites) }},
	"TLChatGetExportedChatInvite":            RPCContextTuple{"/mtproto.RPCChat/chat_getExportedChatInvite", func() interface{} { return new(mtproto.ExportedChatInvite) }},
	"TLChatGetExportedChatInvites":           RPCContextTuple{"/mtproto.RPCChat/chat_getExportedChatInvites", func() interface{} { return new(Vector_ExportedChatInvite) }},
	"TLChatCheckChatInvite":                  RPCContextTuple{"/mtproto.RPCChat/chat_checkChatInvite", func() interface{} { return new(ChatInviteExt) }},
	"TLChatImportChatInvite":                 RPCContextTuple{"/mtproto.RPCChat/chat_importChatInvite", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatGetChatInviteImporters":           RPCContextTuple{"/mtproto.RPCChat/chat_getChatInviteImporters", func() interface{} { return new(Vector_ChatInviteImporter) }},
	"TLChatDeleteExportedChatInvite":         RPCContextTuple{"/mtproto.RPCChat/chat_deleteExportedChatInvite", func() interface{} { return new(mtproto.Bool) }},
	"TLChatDeleteRevokedExportedChatInvites": RPCContextTuple{"/mtproto.RPCChat/chat_deleteRevokedExportedChatInvites", func() interface{} { return new(mtproto.Bool) }},
	"TLChatEditExportedChatInvite":           RPCContextTuple{"/mtproto.RPCChat/chat_editExportedChatInvite", func() interface{} { return new(Vector_ExportedChatInvite) }},
	"TLChatSetChatAvailableReactions":        RPCContextTuple{"/mtproto.RPCChat/chat_setChatAvailableReactions", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatSetHistoryTTL":                    RPCContextTuple{"/mtproto.RPCChat/chat_setHistoryTTL", func() interface{} { return new(mtproto.MutableChat) }},
	"TLChatSearch":                           RPCContextTuple{"/mtproto.RPCChat/chat_search", func() interface{} { return new(Vector_MutableChat) }},
	"TLChatGetRecentChatInviteRequesters":    RPCContextTuple{"/mtproto.RPCChat/chat_getRecentChatInviteRequesters", func() interface{} { return new(RecentChatInviteRequesters) }},
	"TLChatHideChatJoinRequests":             RPCContextTuple{"/mtproto.RPCChat/chat_hideChatJoinRequests", func() interface{} { return new(RecentChatInviteRequesters) }},
	"TLChatImportChatInvite2":                RPCContextTuple{"/mtproto.RPCChat/chat_importChatInvite2", func() interface{} { return new(ChatInviteImported) }},
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
