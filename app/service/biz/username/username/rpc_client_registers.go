/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package username

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
	"TLUsernameGetAccountUsername":    RPCContextTuple{"/mtproto.RPCUsername/username_getAccountUsername", func() interface{} { return new(UsernameData) }},
	"TLUsernameCheckAccountUsername":  RPCContextTuple{"/mtproto.RPCUsername/username_checkAccountUsername", func() interface{} { return new(UsernameExist) }},
	"TLUsernameGetChannelUsername":    RPCContextTuple{"/mtproto.RPCUsername/username_getChannelUsername", func() interface{} { return new(UsernameData) }},
	"TLUsernameCheckChannelUsername":  RPCContextTuple{"/mtproto.RPCUsername/username_checkChannelUsername", func() interface{} { return new(UsernameExist) }},
	"TLUsernameUpdateUsernameByPeer":  RPCContextTuple{"/mtproto.RPCUsername/username_updateUsernameByPeer", func() interface{} { return new(mtproto.Bool) }},
	"TLUsernameCheckUsername":         RPCContextTuple{"/mtproto.RPCUsername/username_checkUsername", func() interface{} { return new(UsernameExist) }},
	"TLUsernameUpdateUsername":        RPCContextTuple{"/mtproto.RPCUsername/username_updateUsername", func() interface{} { return new(mtproto.Bool) }},
	"TLUsernameDeleteUsername":        RPCContextTuple{"/mtproto.RPCUsername/username_deleteUsername", func() interface{} { return new(mtproto.Bool) }},
	"TLUsernameResolveUsername":       RPCContextTuple{"/mtproto.RPCUsername/username_resolveUsername", func() interface{} { return new(mtproto.Peer) }},
	"TLUsernameGetListByUsernameList": RPCContextTuple{"/mtproto.RPCUsername/username_getListByUsernameList", func() interface{} { return new(Vector_UsernameData) }},
	"TLUsernameDeleteUsernameByPeer":  RPCContextTuple{"/mtproto.RPCUsername/username_deleteUsernameByPeer", func() interface{} { return new(mtproto.Bool) }},
	"TLUsernameSearch":                RPCContextTuple{"/mtproto.RPCUsername/username_search", func() interface{} { return new(Vector_UsernameData) }},
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
