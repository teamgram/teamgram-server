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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSaveDraft
// messages.saveDraft#54ae308e flags:# no_webpage:flags.1?true invert_media:flags.6?true reply_to:flags.4?InputReplyTo peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> media:flags.5?InputMedia effect:flags.7?long suggested_post:flags.8?SuggestedPost = Bool;
func (c *DraftsCore) MessagesSaveDraft(in *tg.TLMessagesSaveDraft) (*tg.Bool, error) {
	var (
		peer                = tg.FromInputPeer2(c.MD.UserId, in.Peer)
		isDraftMessageEmpty = true
		date                = int32(time.Now().Unix())
	)

	if c.svcCtx == nil || c.svcCtx.Repo == nil || c.svcCtx.Repo.DialogClient == nil {
		return tg.BoolTrue, nil
	}

	if in.NoWebpage == true {
		isDraftMessageEmpty = false
	} else if in.ReplyTo != nil {
		isDraftMessageEmpty = false
	} else if in.Message != "" {
		isDraftMessageEmpty = false
	} else if in.Entities != nil {
		isDraftMessageEmpty = false
	}

	if isDraftMessageEmpty {
		if _, err := c.svcCtx.Repo.DialogClient.DialogClearDraftMessage(c.ctx, &repository.DialogClearDraft{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		}); err != nil {
			return nil, err
		}
	} else {
		draft := tg.MakeTLDraftMessage(&tg.TLDraftMessage{
			NoWebpage:   in.NoWebpage,
			InvertMedia: in.InvertMedia,
			ReplyTo:     in.ReplyTo,
			Message:     in.Message,
			Entities:    in.Entities,
			Media:       in.Media,
			Date:        date,
			Effect:      in.Effect,
		})

		if _, err := c.svcCtx.Repo.DialogClient.DialogSaveDraftMessage(c.ctx, &repository.DialogSaveDraft{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Message:  draft,
		}); err != nil {
			return nil, err
		}
	}

	// TODO: build syncUpdates with user/chat resolution and call SyncUpdatesNotMe.
	// PEER_CHANNEL case requires plugin (enterprise feature).

	return tg.BoolTrue, nil
}
