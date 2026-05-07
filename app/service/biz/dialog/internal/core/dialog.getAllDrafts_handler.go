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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogGetAllDrafts
// dialog.getAllDrafts user_id:long = Vector<PeerWithDraftMessage>;
func (c *DialogCore) DialogGetAllDrafts(in *dialog.TLDialogGetAllDrafts) (*dialog.VectorPeerWithDraftMessage, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	drafts, err := c.svcCtx.Repo.ListActiveDrafts(c.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	out := &dialog.VectorPeerWithDraftMessage{Datas: make([]dialog.PeerWithDraftMessageClazz, 0, len(drafts))}
	for _, draft := range drafts {
		date, err := dialogDateInt32FromUnixSeconds(draft.Date, "draft date")
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, dialog.MakeTLUpdateDraftMessage(&dialog.TLUpdateDraftMessage{
			Peer: tgPeer(draft.PeerType, draft.PeerID),
			Draft: tg.MakeTLDraftMessage(&tg.TLDraftMessage{
				Message: draft.Message,
				Date:    date,
			}),
		}))
	}
	return out, nil
}
