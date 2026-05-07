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
	"github.com/teamgram/teamgram-server/v2/pkg/paging"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesListDialogs
// userupdates.listDialogs user_id:long top_message_date:long top_peer_seq:long peer_type:int peer_id:long limit:int = DialogProjectionList;
func (c *UserupdatesCore) UserupdatesListDialogs(in *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error) {
	if in == nil || in.UserId == 0 {
		return nil, fmt.Errorf("%w: invalid list dialogs request", userupdates.ErrOperationTerminal)
	}
	limit := paging.NormalizeDialogLimit(in.Limit)
	projections, err := c.svcCtx.Repo.ListDialogProjections(c.ctx, in.UserId, repository.DialogProjectionCursor{
		TopMessageDate: in.TopMessageDate,
		TopPeerSeq:     in.TopPeerSeq,
		PeerType:       in.PeerType,
		PeerID:         in.PeerId,
	}, limit)
	if err != nil {
		return nil, err
	}

	out := &userupdates.TLDialogProjectionList{
		Projections: make([]userupdates.DialogProjectionClazz, 0, len(projections)),
		Exhausted:   tg.ToBoolClazz(len(projections) < int(limit)),
	}
	if limit == 0 {
		out.Exhausted = tg.BoolTrueClazz
	}
	for _, projection := range projections {
		out.Projections = append(out.Projections, dialogProjectionToTL(projection))
	}
	if len(projections) != 0 {
		last := projections[len(projections)-1]
		out.NextTopMessageDate = last.TopMessageDate
		out.NextTopPeerSeq = last.TopPeerSeq
		out.NextPeerType = last.PeerType
		out.NextPeerId = last.PeerID
	}
	return userupdates.MakeTLDialogProjectionList(out), nil
}

func dialogProjectionToTL(in repository.DialogProjection) userupdates.DialogProjectionClazz {
	return userupdates.MakeTLDialogProjection(&userupdates.TLDialogProjection{
		PeerType:                 in.PeerType,
		PeerId:                   in.PeerID,
		TopPeerSeq:               in.TopPeerSeq,
		TopCanonicalMessageId:    in.TopCanonicalMessageID,
		TopMessageDate:           in.TopMessageDate,
		TopMessageStatus:         in.TopMessageStatus,
		ReadInboxMaxPeerSeq:      in.ReadInboxMaxPeerSeq,
		ReadOutboxMaxPeerSeq:     in.ReadOutboxMaxPeerSeq,
		UnreadCount:              in.UnreadCount,
		UnreadMentionsCount:      in.UnreadMentionsCount,
		UnreadReactionsCount:     in.UnreadReactionsCount,
		UnreadMark:               in.UnreadMark,
		PinnedPeerSeq:            in.PinnedPeerSeq,
		PinnedCanonicalMessageId: in.PinnedCanonicalMessageID,
		HasScheduled:             in.HasScheduled,
		AvailableMinPeerSeq:      in.AvailableMinPeerSeq,
		LastPts:                  in.LastPTS,
		LastPtsAt:                in.LastPTSAt,
		DialogSchemaVersion:      in.DialogSchemaVersion,
		DialogPayload:            in.DialogPayload,
	})
}
