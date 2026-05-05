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

package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesAppendDialogPtsSideEffect
// userupdates.appendDialogPtsSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserPtsAppendResult;
func (c *UserupdatesCore) UserupdatesAppendDialogPtsSideEffect(in *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: missing append pts side effect request", userupdates.ErrOperationTerminal)
	}
	if err := validateDialogSideEffectAppend(dialogSideEffectAppendInput{
		UserID:              in.UserId,
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		TargetAuthPolicy:    in.TargetAuthPolicy,
		Payload:             in.Payload,
		PayloadHash:         in.PayloadHash,
	}); err != nil {
		return nil, err
	}

	result, err := c.svcCtx.Repo.AppendDialogPtsSideEffect(c.ctx, repository.DialogSideEffectAppendInput{
		UserID:               in.UserId,
		SourcePermAuthKeyID:  in.SourcePermAuthKeyId,
		OperationID:          in.OperationId,
		TargetAuthPolicy:     in.TargetAuthPolicy,
		PublicUpdateType:     in.PublicUpdateType,
		PeerType:             in.PeerType,
		PeerID:               in.PeerId,
		PayloadSchemaVersion: in.PayloadSchemaVersion,
		Payload:              in.Payload,
		PayloadHash:          in.PayloadHash,
	})
	if err != nil {
		return nil, err
	}
	return userupdates.MakeTLUserPtsAppendResult(&userupdates.TLUserPtsAppendResult{
		UserId:         result.UserID,
		OperationId:    result.OperationID,
		Pts:            result.Pts,
		PtsCount:       result.PtsCount,
		AlreadyApplied: tg.ToBoolClazz(result.AlreadyApplied),
	}).ToUserPtsAppendResult(), nil
}
