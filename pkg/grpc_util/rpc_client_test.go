// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package grpc_util

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

type newRPCReplyFunc func() interface{}

type RPCContextTuple struct {
	Method       string
	NewReplyFunc newRPCReplyFunc
}

var rpcContextRegisters = map[string]RPCContextTuple{
	"TLAuthSendCode": RPCContextTuple{
		"/rpc.Auth/auth_sentCode", func() interface{} {
			return new(mtproto.Auth_SentCode)
		},
	},
}

func FindRPCContextTuple(t interface{}) *RPCContextTuple {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	fmt.Println(rt.Name())

	m, ok := rpcContextRegisters[rt.Name()]
	if !ok {
		glog.Error("Can't find name: ", rt.Name())
		return nil
	}
	return &m
}

func TestReflectTLObject(t *testing.T) {
	authSendCode := &mtproto.TLAuthSendCode{}

	rt := reflect.TypeOf(authSendCode)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	fmt.Println(rt.Name())
}

func TestRPCClient(t *testing.T) {
	fmt.Println("TestRPCClient...")
	conn, err := grpc.Dial("127.0.0.1:10001", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	// client := NewAuthClient(conn)
	authSendCode := &mtproto.TLAuthSendCode{}
	// glog.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	// auth_SentCode, err := client.AuthSentCode(context.Background(), authSendCode)

	t2 := FindRPCContextTuple(authSendCode)
	if t2 == nil {
		return
	}

	fmt.Printf("1. %v\n", t2)

	o2 := t2.NewReplyFunc()
	err = grpc.Invoke(context.Background(), t2.Method, authSendCode, o2, conn)

	if err != nil {
		fmt.Errorf("%v.AuthSentCode(_) = _, %v: \n", conn, err)
	}
	fmt.Printf("2. %v\n", o2)
}
