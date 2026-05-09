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
)

// UserupdatesProcessUserOperation
// userupdates.processUserOperation operation:UserOperation = UserOperationResult;
func (c *UserupdatesCore) UserupdatesProcessUserOperation(in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	if in == nil || in.Operation == nil {
		return nil, fmt.Errorf("%w: missing operation", userupdates.ErrOperationTerminal)
	}
	applyIn, err := operationToApplyInput(in.Operation)
	if err != nil {
		return nil, err
	}

	result, err := c.svcCtx.Repo.ApplyUserOperation(c.ctx, applyIn)
	if err != nil {
		return nil, err
	}
	if c.svcCtx.PushOutboxNotifier != nil {
		c.svcCtx.PushOutboxNotifier.Wake()
	}
	return applyResultToTL(result)
}

func operationToApplyInput(op *userupdates.TLUserOperation) (repository.ApplyUserOperationInput, error) {
	if op == nil {
		return repository.ApplyUserOperationInput{}, fmt.Errorf("%w: missing operation", userupdates.ErrOperationTerminal)
	}

	var dependencyPts []int64
	if op.DependencyPts != nil {
		dependencyPts = []int64{*op.DependencyPts}
	}

	return repository.ApplyUserOperationInput{
		UserID:           op.UserId,
		OperationID:      op.OperationId,
		OpType:           op.OpType,
		PeerType:         op.PeerType,
		PeerID:           op.PeerId,
		PayloadCodec:     op.PayloadCodec,
		Payload:          op.Payload,
		PayloadHash:      op.PayloadHash,
		BucketID:         op.BucketId,
		PartitionID:      op.PartitionId,
		DependencyPts:    dependencyPts,
		AuthKeyIDExclude: op.AuthKeyIdExclude,
	}, nil
}
