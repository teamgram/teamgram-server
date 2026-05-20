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
	"bytes"
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesAppendDialogAuthSeqSideEffect
// userupdates.appendDialogAuthSeqSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserAuthSeqAppendResult;
func (c *UserupdatesCore) UserupdatesAppendDialogAuthSeqSideEffect(in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error) {
	if in == nil {
		return nil, fmt.Errorf("%w: missing append auth seq side effect request", userupdates.ErrOperationTerminal)
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

	targetPermAuthKeyIDs, visibilityPolicy, err := c.expandAuthSeqTargets(in.UserId, in.SourcePermAuthKeyId, in.TargetAuthPolicy)
	if err != nil {
		return nil, err
	}
	appender, ok := c.svcCtx.Repo.(authSeqUpdateAppender)
	if !ok {
		return nil, fmt.Errorf("%w: auth seq repository append is unavailable", userupdates.ErrAuthSeqLedgerUnavailable)
	}
	result, err := appender.AppendAuthSeqUpdate(c.ctx, repository.AuthSeqUpdateAppendInput{
		UserID:               in.UserId,
		SourcePermAuthKeyID:  in.SourcePermAuthKeyId,
		TargetPermAuthKeyIDs: targetPermAuthKeyIDs,
		OperationID:          in.OperationId,
		UpdateType:           in.PublicUpdateType,
		ReplayPolicy:         repository.AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     visibilityPolicy,
		Layer:                in.PayloadSchemaVersion,
		TLBytes:              in.Payload,
		PayloadHash:          in.PayloadHash,
	})
	if err != nil {
		return nil, err
	}
	var seq int64
	var date int32
	if len(result.Deliveries) > 0 {
		seq = result.Deliveries[0].Seq
		date = result.Deliveries[0].Date
	}
	return userupdates.MakeTLUserAuthSeqAppendResult(&userupdates.TLUserAuthSeqAppendResult{
		UserId:         result.UserID,
		OperationId:    result.OperationID,
		Seq:            seq,
		Date:           date,
		AlreadyApplied: tg.ToBoolClazz(result.AlreadyApplied),
	}).ToUserAuthSeqAppendResult(), nil
}

type authSeqUpdateAppender interface {
	AppendAuthSeqUpdate(ctx context.Context, in repository.AuthSeqUpdateAppendInput) (*repository.AuthSeqUpdateAppendResult, error)
}

func (c *UserupdatesCore) expandAuthSeqTargets(userID, sourcePermAuthKeyID int64, policy string) ([]int64, string, error) {
	switch policy {
	case "", repository.AuthSeqVisibilityAllUserAuthKeys:
		if c == nil || c.svcCtx == nil || c.svcCtx.AuthsessionClient == nil {
			return nil, "", fmt.Errorf("%w: authsession client is nil", userupdates.ErrAuthSeqLedgerUnavailable)
		}
		keys, err := c.svcCtx.AuthsessionClient.AuthsessionGetPermAuthKeyIds(c.ctx, &authsession.TLAuthsessionGetPermAuthKeyIds{UserId: userID})
		if err != nil {
			return nil, "", fmt.Errorf("%w: get auth seq targets: %v", userupdates.ErrAuthSeqLedgerUnavailable, err)
		}
		if keys == nil {
			return []int64{}, repository.AuthSeqVisibilityAllUserAuthKeys, nil
		}
		return uniqueActivePermAuthKeyIDs(keys.Datas), repository.AuthSeqVisibilityAllUserAuthKeys, nil
	case repository.AuthSeqVisibilityNotSourcePermAuthKey:
		if c == nil || c.svcCtx == nil || c.svcCtx.AuthsessionClient == nil {
			return nil, "", fmt.Errorf("%w: authsession client is nil", userupdates.ErrAuthSeqLedgerUnavailable)
		}
		targets, err := resolveAuthSeqNotMeTargets(c.ctx, c.svcCtx.AuthsessionClient, userID, sourcePermAuthKeyID)
		if err != nil {
			return nil, "", fmt.Errorf("%w: get auth seq targets: %v", userupdates.ErrAuthSeqLedgerUnavailable, err)
		}
		return targets, repository.AuthSeqVisibilityNotSourcePermAuthKey, nil
	default:
		return nil, "", fmt.Errorf("%w: unsupported target auth policy=%s", userupdates.ErrOperationTerminal, policy)
	}
}

func uniqueActivePermAuthKeyIDs(keys []int64) []int64 {
	if len(keys) == 0 {
		return []int64{}
	}
	seen := make(map[int64]struct{}, len(keys))
	out := make([]int64, 0, len(keys))
	for _, keyID := range keys {
		if keyID == 0 {
			continue
		}
		if _, ok := seen[keyID]; ok {
			continue
		}
		seen[keyID] = struct{}{}
		out = append(out, keyID)
	}
	return out
}

type dialogSideEffectAppendInput struct {
	UserID              int64
	SourcePermAuthKeyID int64
	OperationID         string
	TargetAuthPolicy    string
	Payload             []byte
	PayloadHash         []byte
}

func validateDialogSideEffectAppend(in dialogSideEffectAppendInput) error {
	if in.UserID == 0 {
		return fmt.Errorf("%w: missing user_id", userupdates.ErrOperationTerminal)
	}
	if in.OperationID == "" {
		return fmt.Errorf("%w: missing operation_id", userupdates.ErrOperationTerminal)
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(in.Payload)) {
		return userupdates.ErrOperationPayloadConflict
	}
	if in.TargetAuthPolicy == "not_source_perm_auth_key" && in.SourcePermAuthKeyID == 0 {
		return fmt.Errorf("%w: missing source_perm_auth_key_id", userupdates.ErrOperationTerminal)
	}
	return nil
}
