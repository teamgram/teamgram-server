// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iface

import (
	"reflect"
	"strings"
	"unicode"
)

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var (
	rpcContextRegisters = map[string]RPCContextTuple{}
)

func RegisterRPCContextTuple(name string, method string, newReplyFunc newRPCReplyFunc) {
	if _, ok := rpcContextRegisters[name]; ok {
		// log.Errorf("Already registered name: %s", rt.Name())
		return
	}

	rpcContextRegisters[name] = RPCContextTuple{
		Method:       method,
		NewReplyFunc: newReplyFunc,
	}
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

func FindRPCContextTupleByClazzID(clazzID uint32) *RPCContextTuple {
	clazzName := GetClazzNameByID(clazzID)
	if clazzName == "" {
		return nil
	}

	m, ok := rpcContextRegisters[tlObjectNameFromClazzName(clazzName)]
	if !ok {
		return nil
	}
	return &m
}

func (m RPCContextTuple) ServiceName() string {
	method := strings.TrimPrefix(m.Method, "/tg.")
	if idx := strings.IndexByte(method, '/'); idx >= 0 {
		return method[:idx]
	}
	return ""
}

func GetRPCContextRegisters() map[string]RPCContextTuple {
	return rpcContextRegisters
}

func tlObjectNameFromClazzName(clazzName string) string {
	var b strings.Builder
	b.WriteString("TL")
	upperNext := true
	for _, r := range clazzName {
		if r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			b.WriteRune(unicode.ToUpper(r))
			upperNext = false
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
