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
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogGetDialogsV2
// dialog.getDialogsV2 user_id:long cursor:DialogCursor exclude_pinned:Bool limit:int = DialogPage;
func (c *DialogCore) DialogGetDialogsV2(in *dialog.TLDialogGetDialogsV2) (*dialog.DialogPage, error) {
	if c.svcCtx.Userupdates == nil {
		return nil, dialog.WrapDialogStorage("userupdates.listDialogs", errors.New("userupdates client is not configured"))
	}
	folderID := int32(0)
	topMessageDate := int64(0)
	topPeerSeq := int64(0)
	peerType := int32(0)
	peerID := int64(0)
	if in.Cursor != nil {
		folderID = in.Cursor.FolderId
		topMessageDate = in.Cursor.TopMessageDate
		topPeerSeq = in.Cursor.TopPeerSeq
		peerType = in.Cursor.PeerType
		peerID = in.Cursor.PeerId
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 100
	}
	projectionPage, err := c.svcCtx.Userupdates.UserupdatesListDialogs(c.ctx, &userupdates.TLUserupdatesListDialogs{
		UserId:         in.UserId,
		TopMessageDate: topMessageDate,
		TopPeerSeq:     topPeerSeq,
		PeerType:       peerType,
		PeerId:         peerID,
		Limit:          limit,
	})
	if err != nil {
		return nil, dialog.WrapDialogStorage("userupdates.listDialogs", err)
	}
	records := projectionPage.Projections
	peers := make([]repository.PeerRef, 0, len(records))
	for _, projection := range records {
		if projection == nil {
			continue
		}
		peers = append(peers, repository.PeerRef{PeerType: projection.PeerType, PeerID: projection.PeerId})
	}
	extras, err := c.svcCtx.Repo.BatchGetDialogExtras(c.ctx, in.UserId, peers)
	if err != nil {
		return nil, err
	}
	if folderID != 0 || tg.FromBoolClazz(in.ExcludePinned) {
		extrasByPeer := make(map[repository.PeerRef]repository.DialogExtrasRecord, len(extras))
		for _, extra := range extras {
			extrasByPeer[repository.PeerRef{PeerType: extra.PeerType, PeerID: extra.PeerID}] = extra
		}
		filtered := make([]userupdates.DialogProjectionClazz, 0, len(records))
		for _, projection := range records {
			if projection == nil {
				continue
			}
			extra := extrasByPeer[repository.PeerRef{PeerType: projection.PeerType, PeerID: projection.PeerId}]
			if folderID != 0 && extra.FolderID != folderID {
				continue
			}
			if tg.FromBoolClazz(in.ExcludePinned) && (extra.MainPinnedOrder != 0 || extra.FolderPinnedOrder != 0) {
				continue
			}
			filtered = append(filtered, projection)
		}
		records = filtered
	}
	page := dialog.MakeTLDialogPage(&dialog.TLDialogPage{
		Dialogs:    makeDialogExtV2VectorFromProjections(records, extras).Datas,
		NextCursor: dialog.MakeTLDialogCursor(&dialog.TLDialogCursor{FolderId: folderID}),
		Exhausted:  projectionPage.Exhausted,
	})
	if projectionPage.NextTopMessageDate != 0 || projectionPage.NextTopPeerSeq != 0 || projectionPage.NextPeerId != 0 {
		page.NextCursor = dialog.MakeTLDialogCursor(&dialog.TLDialogCursor{
			FolderId:       folderID,
			Section:        "main",
			TopMessageDate: projectionPage.NextTopMessageDate,
			TopPeerSeq:     projectionPage.NextTopPeerSeq,
			PeerType:       projectionPage.NextPeerType,
			PeerId:         projectionPage.NextPeerId,
		})
	}
	return page, nil
}
