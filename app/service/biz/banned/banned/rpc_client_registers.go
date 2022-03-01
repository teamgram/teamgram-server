/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package banned

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
	"TLBannedCheckPhoneNumberBanned": RPCContextTuple{"/mtproto.RPCBanned/banned_checkPhoneNumberBanned", func() interface{} { return new(mtproto.Bool) }},
	"TLBannedGetBannedByPhoneList":   RPCContextTuple{"/mtproto.RPCBanned/banned_getBannedByPhoneList", func() interface{} { return new(Vector_String) }},
	"TLBannedBan":                    RPCContextTuple{"/mtproto.RPCBanned/banned_ban", func() interface{} { return new(mtproto.Bool) }},
	"TLBannedUnBan":                  RPCContextTuple{"/mtproto.RPCBanned/banned_unBan", func() interface{} { return new(mtproto.Bool) }},
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
