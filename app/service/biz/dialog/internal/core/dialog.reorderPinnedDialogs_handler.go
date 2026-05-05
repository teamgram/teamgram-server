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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogReorderPinnedDialogs
// dialog.reorderPinnedDialogs user_id:long force:Bool folder_id:int id_list:Vector<long> = Bool;
func (c *DialogCore) DialogReorderPinnedDialogs(in *dialog.TLDialogReorderPinnedDialogs) (*tg.Bool, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	peers := make([]repository.PeerRef, 0, len(in.IdList))
	for _, id := range in.IdList {
		peer, err := repository.SplitPeerDialogID(id)
		if err != nil {
			return nil, err
		}
		peers = append(peers, peer)
	}
	_, err := c.svcCtx.Repo.ReorderPinnedDialogs(c.ctx, repository.ReorderPinnedDialogsInput{
		UserID:              in.UserId,
		FolderID:            in.FolderId,
		PeerOrder:           peers,
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		OutboxID:            in.OutboxId,
		EventType:           "dialog.pinnedDialogsReordered",
		Payload:             []byte(`{"schema_version":1}`),
	})
	if err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
