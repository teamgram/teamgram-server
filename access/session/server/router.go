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

package server

import (
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"reflect"
	"github.com/golang/glog"
)

var routerTable = map[string]*grpc_util.RPCClient{}
var defaultClient *grpc_util.RPCClient
// var routerTable = map[string]string{}

type RouterTable struct {
	Method string
	Module string
}

func InstallRouter(rpcClients map[string]*grpc_util.RPCClient, tbl []RouterTable) {
	for k, v := range rpcClients {
		if k == "biz" {
			defaultClient = v
		}
	}

	//if defaultClient == nil {
	//	panic("not biz in rpcClients")
	//}

	for _, v := range tbl {
		if c, ok := rpcClients[v.Module]; ok {
			routerTable[v.Method] = c
		} else {
			glog.Error("duplicate method: ", v)
		}
	}
}

func getRpcClientByRequest(t interface{}) *grpc_util.RPCClient {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if c, ok := routerTable[rt.Name()]; ok {
		return c
	}
	return defaultClient
}
