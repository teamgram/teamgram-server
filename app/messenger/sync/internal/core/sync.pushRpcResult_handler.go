// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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

package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// SyncPushRpcResult
// sync.pushRpcResult user_id:long perm_auth_key_id:long auth_key_id:long gateway_id:string gateway_generation:string gateway_rpc_addr:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (c *SyncCore) SyncPushRpcResult(in *sync.TLSyncPushRpcResult) (*sync.Void, error) {
	const method = "sync.pushRpcResult"
	if in == nil {
		return nil, sync.ErrSyncInvalidArgument
	}
	if err := c.requireCaller(method); err != nil {
		return nil, err
	}
	if err := validateUserID(method, in.UserId); err != nil {
		return nil, err
	}
	for field, value := range map[string]int64{
		"perm_auth_key_id": in.PermAuthKeyId,
		"auth_key_id":      in.AuthKeyId,
		"session_id":       in.SessionId,
	} {
		if err := validateNonZeroID(method, field, value); err != nil {
			return nil, err
		}
	}
	if err := validatePositiveID(method, "client_req_msg_id", in.ClientReqMsgId); err != nil {
		return nil, err
	}
	if in.GatewayId == "" || in.GatewayGeneration == "" || in.GatewayRpcAddr == "" {
		return nil, fmt.Errorf("%w: %s gateway route is incomplete", sync.ErrSyncInvalidArgument, method)
	}
	if len(in.RpcResult) == 0 {
		return nil, fmt.Errorf("%w: %s rpc_result is empty", sync.ErrSyncInvalidArgument, method)
	}
	if err := c.svcCtx.Repo.PushRpcResult(c.ctx, repository.RpcResultRoute{
		UserID:            in.UserId,
		PermAuthKeyID:     in.PermAuthKeyId,
		AuthKeyID:         in.AuthKeyId,
		SessionID:         in.SessionId,
		ClientReqMsgID:    in.ClientReqMsgId,
		GatewayID:         in.GatewayId,
		GatewayGeneration: in.GatewayGeneration,
		GatewayRPCAddr:    in.GatewayRpcAddr,
		RPCResult:         in.RpcResult,
	}); err != nil {
		return nil, err
	}
	return tg.VoidValue, nil
}
