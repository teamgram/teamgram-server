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

package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
)

func (r *Repository) GetUserMessage(ctx context.Context, userID int64, userMessageID int64) (*UserMessageBox, error) {
	if userID <= 0 || userMessageID <= 0 {
		return nil, msg.ErrMsgIdInvalid
	}
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	row, err := r.models.CanonicalQueries.SelectUserMessageBoxByGlobalID(ctx, userID, userMessageID, MessageStatusLive)
	if err != nil {
		if isModelNotFound(err) {
			return nil, msg.ErrMsgIdInvalid
		}
		return nil, storageError("get user message", err)
	}
	return userMessageBoxFromRow(row), nil
}

func (r *Repository) GetUserMessageList(ctx context.Context, userID int64, ids []int64) ([]UserMessageBox, error) {
	if userID <= 0 {
		return nil, msg.ErrMsgIdInvalid
	}
	for _, id := range ids {
		if id <= 0 {
			return nil, msg.ErrMsgIdInvalid
		}
	}
	if len(ids) == 0 {
		return nil, nil
	}
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	out := make([]UserMessageBox, 0, len(ids))
	for _, id := range ids {
		row, err := r.models.CanonicalQueries.SelectUserMessageBoxByGlobalID(ctx, userID, id, MessageStatusLive)
		if err != nil {
			if isModelNotFound(err) {
				return nil, msg.ErrMsgIdInvalid
			}
			return nil, storageError("get user message list", err)
		}
		box := userMessageBoxFromRow(row)
		if box == nil {
			return nil, msg.ErrMsgIdInvalid
		}
		out = append(out, *box)
	}
	return out, nil
}

func (r *Repository) ResolveForwardSourceIdentity(ctx context.Context, lookup ForwardSourceLookup) (*ForwardSourceIdentity, error) {
	if lookup.UserID <= 0 || lookup.SourceUserMessageID <= 0 {
		return nil, msg.ErrMsgIdInvalid
	}
	if lookup.SourcePeerType == 0 || lookup.SourcePeerID <= 0 {
		return nil, msg.ErrMsgIdInvalid
	}
	box, err := r.GetUserMessage(ctx, lookup.UserID, lookup.SourceUserMessageID)
	if err != nil {
		return nil, err
	}
	if box.PeerType != lookup.SourcePeerType || box.PeerID != lookup.SourcePeerID {
		return nil, msg.ErrMsgIdInvalid
	}
	return &ForwardSourceIdentity{
		UserID:             box.UserID,
		UserMessageID:      box.UserMessageID,
		CanonicalMessageID: box.CanonicalMessageID,
	}, nil
}

func (r *Repository) RevalidateForwardSources(ctx context.Context, sources []ForwardSourceIdentity) error {
	if len(sources) == 0 {
		return nil
	}
	if _, err := r.requireDB(); err != nil {
		return err
	}
	seen := make(map[ForwardSourceIdentity]struct{}, len(sources))
	for _, source := range sources {
		if source.UserID <= 0 || source.UserMessageID <= 0 || source.CanonicalMessageID <= 0 {
			return msg.ErrMsgIdInvalid
		}
		if _, ok := seen[source]; ok {
			continue
		}
		seen[source] = struct{}{}
		row, err := r.models.CanonicalQueries.SelectForwardSourceIdentity(ctx, source.UserID, source.UserMessageID, MessageStatusLive)
		if err != nil {
			if isModelNotFound(err) {
				return msg.ErrMsgIdInvalid
			}
			return storageError("revalidate forward source", err)
		}
		if row.CanonicalMessageID != source.CanonicalMessageID {
			return msg.ErrMsgIdInvalid
		}
	}
	return nil
}

func userMessageBoxFromRow(row *model.UserMessageBoxRow) *UserMessageBox {
	if row == nil {
		return nil
	}
	return &UserMessageBox{
		UserID:             row.UserID,
		UserMessageID:      row.UserMessageID,
		CanonicalMessageID: row.CanonicalMessageID,
		PeerType:           row.PeerType,
		PeerID:             row.PeerID,
		PeerSeq:            row.PeerSeq,
		FromUserID:         row.FromUserID,
		Outgoing:           row.Outgoing,
		MessageText:        row.MessageText,
		MessageDate:        row.MessageDate,
		ViewPayload:        append([]byte(nil), row.ViewPayload...),
	}
}
