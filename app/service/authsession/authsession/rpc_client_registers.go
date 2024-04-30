/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package authsession

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
	"TLAuthsessionGetAuthorizations":    RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getAuthorizations", func() interface{} { return new(mtproto.Account_Authorizations) }},
	"TLAuthsessionResetAuthorization":   RPCContextTuple{"/mtproto.RPCAuthsession/authsession_resetAuthorization", func() interface{} { return new(Vector_Long) }},
	"TLAuthsessionGetLayer":             RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getLayer", func() interface{} { return new(mtproto.Int32) }},
	"TLAuthsessionGetLangPack":          RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getLangPack", func() interface{} { return new(mtproto.String) }},
	"TLAuthsessionGetClient":            RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getClient", func() interface{} { return new(mtproto.String) }},
	"TLAuthsessionGetLangCode":          RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getLangCode", func() interface{} { return new(mtproto.String) }},
	"TLAuthsessionGetUserId":            RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getUserId", func() interface{} { return new(mtproto.Int64) }},
	"TLAuthsessionGetPushSessionId":     RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getPushSessionId", func() interface{} { return new(mtproto.Int64) }},
	"TLAuthsessionGetFutureSalts":       RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getFutureSalts", func() interface{} { return new(mtproto.FutureSalts) }},
	"TLAuthsessionQueryAuthKey":         RPCContextTuple{"/mtproto.RPCAuthsession/authsession_queryAuthKey", func() interface{} { return new(mtproto.AuthKeyInfo) }},
	"TLAuthsessionSetAuthKey":           RPCContextTuple{"/mtproto.RPCAuthsession/authsession_setAuthKey", func() interface{} { return new(mtproto.Bool) }},
	"TLAuthsessionBindAuthKeyUser":      RPCContextTuple{"/mtproto.RPCAuthsession/authsession_bindAuthKeyUser", func() interface{} { return new(mtproto.Int64) }},
	"TLAuthsessionUnbindAuthKeyUser":    RPCContextTuple{"/mtproto.RPCAuthsession/authsession_unbindAuthKeyUser", func() interface{} { return new(mtproto.Bool) }},
	"TLAuthsessionGetPermAuthKeyId":     RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getPermAuthKeyId", func() interface{} { return new(mtproto.Int64) }},
	"TLAuthsessionBindTempAuthKey":      RPCContextTuple{"/mtproto.RPCAuthsession/authsession_bindTempAuthKey", func() interface{} { return new(mtproto.Bool) }},
	"TLAuthsessionSetClientSessionInfo": RPCContextTuple{"/mtproto.RPCAuthsession/authsession_setClientSessionInfo", func() interface{} { return new(mtproto.Bool) }},
	"TLAuthsessionGetAuthorization":     RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getAuthorization", func() interface{} { return new(mtproto.Authorization) }},
	"TLAuthsessionGetAuthStateData":     RPCContextTuple{"/mtproto.RPCAuthsession/authsession_getAuthStateData", func() interface{} { return new(AuthKeyStateData) }},
	"TLAuthsessionSetLayer":             RPCContextTuple{"/mtproto.RPCAuthsession/authsession_setLayer", func() interface{} { return new(mtproto.Bool) }},
	"TLAuthsessionSetInitConnection":    RPCContextTuple{"/mtproto.RPCAuthsession/authsession_setInitConnection", func() interface{} { return new(mtproto.Bool) }},
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
