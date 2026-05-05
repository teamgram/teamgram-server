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
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
)

// DialogEditPeerFolders
// dialog.editPeerFolders user_id:long peer_dialog_list:Vector<long> folder_id:int = Vector<DialogPinnedExt>;
func (c *DialogCore) DialogEditPeerFolders(in *dialog.TLDialogEditPeerFolders) (*dialog.VectorDialogPinnedExt, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	if len(in.OutboxIds) < len(in.PeerDialogList) {
		return nil, dialog.ErrOutboxUnavailable
	}
	out := &dialog.VectorDialogPinnedExt{Datas: make([]dialog.DialogPinnedExtClazz, 0, len(in.PeerDialogList))}
	for i, id := range in.PeerDialogList {
		peer, err := repository.SplitPeerDialogID(id)
		if err != nil {
			return nil, err
		}
		_, err = c.svcCtx.Repo.EditPeerFolders(c.ctx, repository.EditPeerFoldersInput{
			UserID:              in.UserId,
			PeerType:            peer.PeerType,
			PeerID:              peer.PeerID,
			NewFolderID:         in.FolderId,
			SourcePermAuthKeyID: in.SourcePermAuthKeyId,
			OperationID:         in.OperationId + ":" + fmt.Sprint(id),
			OutboxID:            in.OutboxIds[i],
			PublicUpdateType:    "updateFolderPeers",
			Payload:             []byte(`{"schema_version":1}`),
		})
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, dialog.MakeTLDialogPinnedExt(&dialog.TLDialogPinnedExt{
			Order:    int64(i + 1),
			PeerType: peer.PeerType,
			PeerId:   peer.PeerID,
		}))
	}
	return out, nil
}
