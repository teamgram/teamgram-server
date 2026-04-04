// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/interface/session/session"
	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// SyncPushRpcResult
// sync.pushRpcResult user_id:long auth_key_id:long perm_auth_key_id:long server_id:string session_id:long client_req_msg_id:long rpc_result:bytes = Void;
func (c *SyncCore) SyncPushRpcResult(in *sync.TLSyncPushRpcResult) (*tg.Void, error) {
	if c.svcCtx != nil && c.svcCtx.SessionClient != nil {
		_, err := c.svcCtx.SessionClient.SessionPushRpcResultData(c.ctx, &session.TLSessionPushRpcResultData{
			PermAuthKeyId:  in.PermAuthKeyId,
			AuthKeyId:      in.AuthKeyId,
			SessionId:      in.SessionId,
			ClientReqMsgId: in.ClientReqMsgId,
			RpcResultData:  in.RpcResult,
		})
		if err != nil {
			return nil, err
		}
	}

	return tg.MakeTLVoid(&tg.TLVoid{}).ToVoid(), nil
}
