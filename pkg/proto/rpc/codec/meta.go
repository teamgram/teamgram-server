// Copyright 2021 CloudWeGo Authors
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
//

package codec

import (
	"github.com/apache/thrift/lib/go/thrift"
)

/*
message RpcMeta {
    optional RpcRequestMeta request = 1;
    optional RpcResponseMeta response = 2;
    optional int32 compress_type = 3;
    optional int64 correlation_id = 4;
    optional int32 attachment_size = 5;
    optional ChunkInfo chunk_info = 6;
    optional bytes authentication_data = 7;
    optional StreamSettings stream_settings = 8;
    optional MTProtoMeta mtproto_meta = 9;
}

message RpcRequestMeta {
    required string service_name = 1;
    required string method_name = 2;
    optional int64 log_id = 3;
    optional int64 trace_id = 4;
    optional int64 span_id = 5;
    optional int64 parent_span_id = 6;
}

message RpcResponseMeta {
    optional int32 error_code = 1;
    optional string error_text = 2;
}
*/

type Meta struct {
	ServiceName string
	MethodName  string
	SeqID       int32
	MsgType     uint32
	Payload     []byte
}

type Exception struct {
	TypeID  int32
	Message string
}

func (e Exception) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return thrift.NewTApplicationException(e.TypeID, "").Error()
}
