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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (c *DraftsCore) MessagesClearAllDrafts(in *tg.TLMessagesClearAllDrafts) (*tg.Bool, error) {
	dialogClient, err := c.dialogClient()
	if err != nil {
		return nil, err
	}

	existing, err := dialogClient.DialogGetAllDrafts(c.ctx, &repository.DialogGetAllDrafts{
		UserId: c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}

	operationID := clearAllDraftsOperationID(c.MD.UserId, time.Now().UnixNano())
	var existingDrafts []dialogpb.PeerWithDraftMessageClazz
	if existing != nil {
		existingDrafts = existing.Datas
	}
	outboxIDs := make([]int64, 0, len(existingDrafts))
	for _, draft := range existingDrafts {
		if draft == nil {
			continue
		}
		peer := tg.FromPeer(draft.ToPeerWithDraftMessage().Peer)
		outboxIDs = append(outboxIDs, draftOutboxID(draftOperationID("clear_all", c.MD.UserId, peer.PeerType, peer.PeerId, int64(len(outboxIDs)+1))+"|"+operationID))
	}

	rValues, err := dialogClient.DialogClearAllDrafts(c.ctx, &repository.DialogClearAll{
		UserId:              c.MD.UserId,
		SourcePermAuthKeyId: c.MD.PermAuthKeyId,
		OperationId:         operationID,
		OutboxIds:           outboxIDs,
	})
	if err != nil {
		return nil, err
	}

	if len(rValues.Datas) == 0 {
		return tg.BoolTrue, nil
	}

	// TODO: for each cleared draft, build syncUpdates with user/chat
	// resolution and call SyncUpdatesNotMe. PEER_CHANNEL case requires
	// plugin (enterprise feature).

	return tg.BoolTrue, nil
}
