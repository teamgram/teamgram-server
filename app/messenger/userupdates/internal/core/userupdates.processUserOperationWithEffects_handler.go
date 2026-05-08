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

const deliveryPolicyDurableAsync int32 = repository.DeliveryPolicyDurableAsync

// UserupdatesProcessUserOperationWithEffects
// userupdates.processUserOperationWithEffects operation:UserOperation affected_effects:Vector<AffectedUserOperation> = UserOperationResult;
func (c *UserupdatesCore) UserupdatesProcessUserOperationWithEffects(in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error) {
	if in == nil || in.Operation == nil {
		return nil, fmt.Errorf("%w: missing operation", userupdates.ErrOperationTerminal)
	}
	applyIn, err := userOperationToRepositoryInput(in.Operation)
	if err != nil {
		return nil, err
	}

	for _, effect := range in.AffectedEffects {
		if effect == nil || effect.Operation == nil {
			return nil, fmt.Errorf("%w: missing affected operation", userupdates.ErrOperationTerminal)
		}
		if effect.DeliveryPolicy != deliveryPolicyDurableAsync {
			return nil, fmt.Errorf("%w: unsupported affected delivery policy %d", userupdates.ErrOperationTerminal, effect.DeliveryPolicy)
		}

		op := effect.Operation
		applyIn.AffectedOutboxes = append(applyIn.AffectedOutboxes, repository.AffectedOutbox{
			RequesterUserID:   effect.RequesterUserId,
			TargetUserID:      op.UserId,
			TargetBucketID:    op.BucketId,
			TargetPartitionID: op.PartitionId,
			OperationID:       op.OperationId,
			OpType:            op.OpType,
			OperationKind:     effect.OperationKind,
			PeerType:          op.PeerType,
			PeerID:            op.PeerId,
			PayloadCodec:      op.PayloadCodec,
			Payload:           op.Payload,
			PayloadHash:       op.PayloadHash,
			DeliveryPolicy:    effect.DeliveryPolicy,
		})
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
