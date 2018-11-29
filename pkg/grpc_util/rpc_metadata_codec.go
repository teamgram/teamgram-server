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
	"encoding/base64"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/nebula-chat/chatengine/mtproto"
	"google.golang.org/grpc/metadata"
)

var (
	headerRpcMetadata = "rpc_metadata"
)

func RpcMetadataFromMD(md metadata.MD) (*RpcMetadata, error) {
	val := metautils.NiceMD(md).Get(headerRpcMetadata)
	if val == "" {
		return nil, nil
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error, rpc_metadata: %s, error: %v", val, err)
		//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2),
		//	fmt.Sprintf("Base64 decode error, rpc_metadata: %s", val)))
	}

	rpcMetadata := &RpcMetadata{}
	err = proto.Unmarshal(buf, rpcMetadata)
	if err != nil {
		return nil, fmt.Errorf("RpcMetadata unmarshal error, rpc_metadata: %s, error: %v", val, err)
		//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2),
		//	fmt.Sprintf("Unmarshal error, rpc_metadata: %s", val)))
	}

	return rpcMetadata, nil
}

func RpcMetadataFromIncoming(ctx context.Context) *RpcMetadata {
	md2, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	md, err := RpcMetadataFromMD(md2)
	if err != nil {
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2), fmt.Sprintf("%s", err)))
	}

	return md
}

func RpcMetadataToOutgoing(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		glog.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs(headerRpcMetadata,
		base64.StdEncoding.EncodeToString(buf))), nil
}
