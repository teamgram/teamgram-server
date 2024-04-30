/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package code

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
	"TLCodeCreatePhoneCode":     RPCContextTuple{"/mtproto.RPCCode/code_createPhoneCode", func() interface{} { return new(PhoneCodeTransaction) }},
	"TLCodeGetPhoneCode":        RPCContextTuple{"/mtproto.RPCCode/code_getPhoneCode", func() interface{} { return new(PhoneCodeTransaction) }},
	"TLCodeDeletePhoneCode":     RPCContextTuple{"/mtproto.RPCCode/code_deletePhoneCode", func() interface{} { return new(mtproto.Bool) }},
	"TLCodeUpdatePhoneCodeData": RPCContextTuple{"/mtproto.RPCCode/code_updatePhoneCodeData", func() interface{} { return new(mtproto.Bool) }},
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
