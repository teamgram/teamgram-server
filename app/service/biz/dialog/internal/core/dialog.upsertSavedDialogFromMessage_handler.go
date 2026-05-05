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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogUpsertSavedDialogFromMessage
// dialog.upsertSavedDialogFromMessage user_id:long peer_type:int peer_id:long top_peer_seq:long top_canonical_message_id:long top_message_date:int payload:bytes = Bool;
func (c *DialogCore) DialogUpsertSavedDialogFromMessage(in *dialog.TLDialogUpsertSavedDialogFromMessage) (*tg.Bool, error) {
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := c.svcCtx.Repo.UpsertSavedDialogFromMessage(c.ctx, repository.SavedDialogTopInput{
		UserID:                in.UserId,
		PeerType:              in.PeerType,
		PeerID:                in.PeerId,
		TopPeerSeq:            in.TopPeerSeq,
		TopCanonicalMessageID: in.TopCanonicalMessageId,
		TopMessageDate:        time.Unix(int64(in.TopMessageDate), 0).UTC(),
		Payload:               in.Payload,
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
