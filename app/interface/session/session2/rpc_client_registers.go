/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package session

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
	"TLSessionQueryAuthKey":           RPCContextTuple{"/mtproto.RPCSession/session_queryAuthKey", func() interface{} { return new(mtproto.AuthKeyInfo) }},
	"TLSessionSetAuthKey":             RPCContextTuple{"/mtproto.RPCSession/session_setAuthKey", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionCreateSession":          RPCContextTuple{"/mtproto.RPCSession/session_createSession", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionSendDataToSession":      RPCContextTuple{"/mtproto.RPCSession/session_sendDataToSession", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionSendHttpDataToSession":  RPCContextTuple{"/mtproto.RPCSession/session_sendHttpDataToSession", func() interface{} { return new(HttpSessionData) }},
	"TLSessionCloseSession":           RPCContextTuple{"/mtproto.RPCSession/session_closeSession", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionPushUpdatesData":        RPCContextTuple{"/mtproto.RPCSession/session_pushUpdatesData", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionPushSessionUpdatesData": RPCContextTuple{"/mtproto.RPCSession/session_pushSessionUpdatesData", func() interface{} { return new(mtproto.Bool) }},
	"TLSessionPushRpcResultData":      RPCContextTuple{"/mtproto.RPCSession/session_pushRpcResultData", func() interface{} { return new(mtproto.Bool) }},
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
