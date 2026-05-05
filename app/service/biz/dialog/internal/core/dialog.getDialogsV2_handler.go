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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogGetDialogsV2
// dialog.getDialogsV2 user_id:long cursor:DialogCursor exclude_pinned:Bool limit:int = DialogPage;
func (c *DialogCore) DialogGetDialogsV2(in *dialog.TLDialogGetDialogsV2) (*dialog.DialogPage, error) {
	folderID := int32(0)
	if in.Cursor != nil {
		folderID = in.Cursor.FolderId
	}
	records, err := c.svcCtx.Repo.ListDialogs(c.ctx, in.UserId, tg.FromBoolClazz(in.ExcludePinned), folderID)
	if err != nil {
		return nil, err
	}
	exhausted := true
	if in.Limit > 0 && len(records) > int(in.Limit) {
		records = records[:in.Limit]
		exhausted = false
	}
	peers := make([]repository.PeerRef, 0, len(records))
	for _, record := range records {
		peers = append(peers, repository.PeerRef{PeerType: record.PeerType, PeerID: record.PeerID})
	}
	extras, err := c.svcCtx.Repo.BatchGetDialogExtras(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}
	page := dialog.MakeTLDialogPage(&dialog.TLDialogPage{
		Dialogs:    makeDialogExtV2Vector(records, extras).Datas,
		NextCursor: dialog.MakeTLDialogCursor(&dialog.TLDialogCursor{FolderId: folderID}),
		Exhausted:  tg.ToBoolClazz(exhausted),
	})
	if len(records) > 0 {
		last := records[len(records)-1]
		page.NextCursor = dialog.MakeTLDialogCursor(&dialog.TLDialogCursor{
			FolderId:       folderID,
			Section:        "main",
			TopMessageDate: last.Date,
			TopPeerSeq:     int64(last.TopMessage),
			PeerType:       last.PeerType,
			PeerId:         last.PeerID,
		})
	}
	return page, nil
}
