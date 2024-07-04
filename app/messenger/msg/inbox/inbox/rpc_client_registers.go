/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

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
	"TLInboxEditUserMessageToInbox":     RPCContextTuple{"/mtproto.RPCInbox/inbox_editUserMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxEditChatMessageToInbox":     RPCContextTuple{"/mtproto.RPCInbox/inbox_editChatMessageToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteMessagesToInbox":      RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteMessagesToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteUserHistoryToInbox":   RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteUserHistoryToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxDeleteChatHistoryToInbox":   RPCContextTuple{"/mtproto.RPCInbox/inbox_deleteChatHistoryToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadUserMediaUnreadToInbox": RPCContextTuple{"/mtproto.RPCInbox/inbox_readUserMediaUnreadToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadChatMediaUnreadToInbox": RPCContextTuple{"/mtproto.RPCInbox/inbox_readChatMediaUnreadToInbox", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUpdateHistoryReaded":        RPCContextTuple{"/mtproto.RPCInbox/inbox_updateHistoryReaded", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUpdatePinnedMessage":        RPCContextTuple{"/mtproto.RPCInbox/inbox_updatePinnedMessage", func() interface{} { return new(mtproto.Void) }},
	"TLInboxUnpinAllMessages":           RPCContextTuple{"/mtproto.RPCInbox/inbox_unpinAllMessages", func() interface{} { return new(mtproto.Void) }},
	"TLInboxSendUserMessageToInboxV2":   RPCContextTuple{"/mtproto.RPCInbox/inbox_sendUserMessageToInboxV2", func() interface{} { return new(mtproto.Void) }},
	"TLInboxEditMessageToInboxV2":       RPCContextTuple{"/mtproto.RPCInbox/inbox_editMessageToInboxV2", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadInboxHistory":           RPCContextTuple{"/mtproto.RPCInbox/inbox_readInboxHistory", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadOutboxHistory":          RPCContextTuple{"/mtproto.RPCInbox/inbox_readOutboxHistory", func() interface{} { return new(mtproto.Void) }},
	"TLInboxReadMediaUnreadToInboxV2":   RPCContextTuple{"/mtproto.RPCInbox/inbox_readMediaUnreadToInboxV2", func() interface{} { return new(mtproto.Void) }},
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
