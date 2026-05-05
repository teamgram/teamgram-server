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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

// DialogClearAllDrafts
// dialog.clearAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (c *DialogCore) DialogClearAllDrafts(in *dialog.TLDialogClearAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	cleared, err := c.svcCtx.Repo.ClearAllDrafts(c.ctx, repository.ClearAllDraftsInput{
		UserID:              in.UserId,
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		OutboxIDs:           in.OutboxIds,
	})
	if err != nil {
		return nil, err
	}
	out := &dialog.VectorPeerWithDraftMessage{Datas: make([]dialog.PeerWithDraftMessageClazz, 0, len(cleared))}
	for _, row := range cleared {
		peer, err := repository.SplitPeerDialogID(row.PeerDialogID)
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, dialog.MakeTLUpdateDraftMessage(&dialog.TLUpdateDraftMessage{
			Peer:  tgPeer(peer.PeerType, peer.PeerID),
			Draft: tgDraftEmpty(),
		}))
	}
	return out, nil
}
