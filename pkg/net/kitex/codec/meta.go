// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
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
// Author: teamgramio (teamgram.io@gmail.com)

package codec

import (
	"github.com/apache/thrift/lib/go/thrift"
)

type Meta struct {
	ServiceName string
	MethodName  string
	SeqID       int32
	MsgType     uint32
	Payload     []byte
	Metadata    map[string]string
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
