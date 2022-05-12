/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  NebulaChat Studio (https://nebula.chat).
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package message

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
	"TLMessageGetUserMessage":                 RPCContextTuple{"/mtproto.RPCMessage/message_getUserMessage", func() interface{} { return new(mtproto.MessageBox) }},
	"TLMessageGetUserMessageList":             RPCContextTuple{"/mtproto.RPCMessage/message_getUserMessageList", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetUserMessageListByDataIdList": RPCContextTuple{"/mtproto.RPCMessage/message_getUserMessageListByDataIdList", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetHistoryMessages":             RPCContextTuple{"/mtproto.RPCMessage/message_getHistoryMessages", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetHistoryMessagesCount":        RPCContextTuple{"/mtproto.RPCMessage/message_getHistoryMessagesCount", func() interface{} { return new(mtproto.Int32) }},
	"TLMessageGetPeerUserMessageId":           RPCContextTuple{"/mtproto.RPCMessage/message_getPeerUserMessageId", func() interface{} { return new(mtproto.Int32) }},
	"TLMessageGetPeerUserMessage":             RPCContextTuple{"/mtproto.RPCMessage/message_getPeerUserMessage", func() interface{} { return new(mtproto.MessageBox) }},
	"TLMessageGetPeerChatMessageIdList":       RPCContextTuple{"/mtproto.RPCMessage/message_getPeerChatMessageIdList", func() interface{} { return new(Vector_PeerMessageId) }},
	"TLMessageGetPeerChatMessageList":         RPCContextTuple{"/mtproto.RPCMessage/message_getPeerChatMessageList", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageSearchByMediaType":              RPCContextTuple{"/mtproto.RPCMessage/message_searchByMediaType", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageSearch":                         RPCContextTuple{"/mtproto.RPCMessage/message_search", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageSearchGlobal":                   RPCContextTuple{"/mtproto.RPCMessage/message_searchGlobal", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageSearchByPinned":                 RPCContextTuple{"/mtproto.RPCMessage/message_searchByPinned", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetSearchCounter":               RPCContextTuple{"/mtproto.RPCMessage/message_getSearchCounter", func() interface{} { return new(mtproto.Int32) }},
	"TLMessageSearchV2":                       RPCContextTuple{"/mtproto.RPCMessage/message_searchV2", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetLastTwoPinnedMessageId":      RPCContextTuple{"/mtproto.RPCMessage/message_getLastTwoPinnedMessageId", func() interface{} { return new(Vector_Int) }},
	"TLMessageUpdatePinnedMessageId":          RPCContextTuple{"/mtproto.RPCMessage/message_updatePinnedMessageId", func() interface{} { return new(mtproto.Bool) }},
	"TLMessageGetPinnedMessageIdList":         RPCContextTuple{"/mtproto.RPCMessage/message_getPinnedMessageIdList", func() interface{} { return new(Vector_Int) }},
	"TLMessageUnPinAllMessages":               RPCContextTuple{"/mtproto.RPCMessage/message_unPinAllMessages", func() interface{} { return new(Vector_Int) }},
	"TLMessageGetUnreadMentions":              RPCContextTuple{"/mtproto.RPCMessage/message_getUnreadMentions", func() interface{} { return new(Vector_MessageBox) }},
	"TLMessageGetUnreadMentionsCount":         RPCContextTuple{"/mtproto.RPCMessage/message_getUnreadMentionsCount", func() interface{} { return new(mtproto.Int32) }},
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
