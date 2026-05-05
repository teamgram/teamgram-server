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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

// DialogGetPinnedDialogsV2
// dialog.getPinnedDialogsV2 user_id:long folder_id:int limit:int = Vector<DialogExtV2>;
func (c *DialogCore) DialogGetPinnedDialogsV2(in *dialog.TLDialogGetPinnedDialogsV2) (*dialog.VectorDialogExtV2, error) {
	records, err := c.svcCtx.Repo.ListPinnedDialogs(c.ctx, in.UserId, in.FolderId)
	if err != nil {
		return nil, err
	}
	if in.Limit > 0 && len(records) > int(in.Limit) {
		records = records[:in.Limit]
	}
	peers := make([]repository.PeerRef, 0, len(records))
	for _, record := range records {
		peers = append(peers, repository.PeerRef{PeerType: record.PeerType, PeerID: record.PeerID})
	}
	extras, err := c.svcCtx.Repo.BatchGetDialogExtras(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}
	return makeDialogExtV2Vector(records, extras), nil
}
