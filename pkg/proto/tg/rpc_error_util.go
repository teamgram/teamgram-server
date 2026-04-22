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
//
// Author: teamgramio (teamgram.io@gmail.com)

package tg

import (
	"errors"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
	"github.com/teamgram/teamgram-server/v2/pkg/xerr"
)

type (
	TLRpcError = mt.TLRpcError
	RpcError   = mt.RpcError
)

var (
	DecodeRpcErrorClazz = mt.DecodeRpcErrorClazz
)

func NewRpcError(e error) *TLRpcError {
	if e == nil {
		return &TLRpcError{
			ErrorCode:    ErrInternal,
			ErrorMessage: "INTERNAL_SERVER_ERROR",
		}
	}

	var rpcErr *TLRpcError
	if errors.As(e, &rpcErr) && rpcErr != nil {
		return rpcErr
	}

	var (
		err xerr.CodeError
	)
	ok := errors.As(e, &err)
	if ok {
		return &TLRpcError{
			ErrorCode:    int32(err.Code()),
			ErrorMessage: err.Msg(),
		}
	} else {
		return &TLRpcError{
			ErrorCode:    ErrInternal,
			ErrorMessage: e.Error(),
		}
	}
}
