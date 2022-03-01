/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package auth

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
	"TLAuthExportLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_exportLoginToken", func() interface{} { return new(mtproto.Auth_LoginToken) }},
	"TLAuthImportLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_importLoginToken", func() interface{} { return new(mtproto.Auth_LoginToken) }},
	"TLAuthAcceptLoginToken": RPCContextTuple{"/mtproto.RPCAuth/auth_acceptLoginToken", func() interface{} { return new(mtproto.Authorization) }},
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
