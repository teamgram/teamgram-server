/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package sync

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
	"TLSyncUpdatesMe":        RPCContextTuple{"/mtproto.RPCSync/sync_updatesMe", func() interface{} { return new(mtproto.Void) }},
	"TLSyncUpdatesNotMe":     RPCContextTuple{"/mtproto.RPCSync/sync_updatesNotMe", func() interface{} { return new(mtproto.Void) }},
	"TLSyncPushUpdates":      RPCContextTuple{"/mtproto.RPCSync/sync_pushUpdates", func() interface{} { return new(mtproto.Void) }},
	"TLSyncPushUpdatesIfNot": RPCContextTuple{"/mtproto.RPCSync/sync_pushUpdatesIfNot", func() interface{} { return new(mtproto.Void) }},
	"TLSyncPushBotUpdates":   RPCContextTuple{"/mtproto.RPCSync/sync_pushBotUpdates", func() interface{} { return new(mtproto.Void) }},
	"TLSyncPushRpcResult":    RPCContextTuple{"/mtproto.RPCSync/sync_pushRpcResult", func() interface{} { return new(mtproto.Void) }},
	"TLSyncBroadcastUpdates": RPCContextTuple{"/mtproto.RPCSync/sync_broadcastUpdates", func() interface{} { return new(mtproto.Void) }},
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
