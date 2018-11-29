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
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc/metadata"
)

var (
	headerRpcError = "rpc_error"
)

// Server To Client
func RpcErrorFromMD(md metadata.MD) (rpcErr *mtproto.TLRpcError) {
	glog.Info("rpc error from md: ", md)
	val := metautils.NiceMD(md).Get(headerRpcError)
	if val == "" {
		// TODO(@benqi): 未设置rpc_error
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("Unknown error"))
		glog.Errorf("%v", rpcErr)
		return
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("Base64 decode error, rpc_error: %s, error: %v", val, err))
		glog.Errorf("%v", rpcErr)
		return
	}

	rpcErr = &mtproto.TLRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL),
			fmt.Sprintf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err))
		glog.Errorf("%v", rpcErr)
		return
	}

	// glog.Errorf("%v", rpcErr)
	return rpcErr
}

func RpcErrorToMD(md *mtproto.TLRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		glog.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRpcError, base64.StdEncoding.EncodeToString(buf)), nil
}
