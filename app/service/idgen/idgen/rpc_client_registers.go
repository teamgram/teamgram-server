/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package idgen

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
	"TLIdgenNextId":              RPCContextTuple{"/mtproto.RPCIdgen/idgen_nextId", func() interface{} { return new(mtproto.Int64) }},
	"TLIdgenNextIds":             RPCContextTuple{"/mtproto.RPCIdgen/idgen_nextIds", func() interface{} { return new(Vector_Long) }},
	"TLIdgenGetCurrentSeqId":     RPCContextTuple{"/mtproto.RPCIdgen/idgen_getCurrentSeqId", func() interface{} { return new(mtproto.Int64) }},
	"TLIdgenSetCurrentSeqId":     RPCContextTuple{"/mtproto.RPCIdgen/idgen_setCurrentSeqId", func() interface{} { return new(mtproto.Bool) }},
	"TLIdgenGetNextSeqId":        RPCContextTuple{"/mtproto.RPCIdgen/idgen_getNextSeqId", func() interface{} { return new(mtproto.Int64) }},
	"TLIdgenGetNextNSeqId":       RPCContextTuple{"/mtproto.RPCIdgen/idgen_getNextNSeqId", func() interface{} { return new(mtproto.Int64) }},
	"TLIdgenGetNextIdValList":    RPCContextTuple{"/mtproto.RPCIdgen/idgen_getNextIdValList", func() interface{} { return new(Vector_IdVal) }},
	"TLIdgenGetCurrentSeqIdList": RPCContextTuple{"/mtproto.RPCIdgen/idgen_getCurrentSeqIdList", func() interface{} { return new(Vector_IdVal) }},
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
