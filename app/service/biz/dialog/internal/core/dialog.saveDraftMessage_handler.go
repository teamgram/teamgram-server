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
	"encoding/json"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogSaveDraftMessage
// dialog.saveDraftMessage user_id:long peer_type:int peer_id:long message:DraftMessage = Bool;
func (c *DialogCore) DialogSaveDraftMessage(in *dialog.TLDialogSaveDraftMessage) (*tg.Bool, error) {
	if in == nil || in.Message == nil {
		return nil, dialog.ErrDialogInvalid
	}
	payload, err := json.Marshal(in.Message)
	if err != nil {
		return nil, dialog.WrapDialogStorage("marshal draft payload", err)
	}
	var (
		draftKind       int32 = 1
		message         string
		entitiesPayload = []byte{}
		date            = time.Now().UTC()
	)
	if draft, ok := in.Message.(*tg.TLDraftMessage); ok {
		message = draft.Message
		date = time.Unix(int64(draft.Date), 0).UTC()
		if draft.Entities != nil {
			entitiesPayload, err = json.Marshal(draft.Entities)
			if err != nil {
				return nil, dialog.WrapDialogStorage("marshal draft entities", err)
			}
		}
	}
	if _, ok := in.Message.(*tg.TLDraftMessageEmpty); ok {
		draftKind = 0
	}

	_, err = c.svcCtx.Repo.SaveDraft(c.ctx, repository.SaveDraftInput{
		UserID:              in.UserId,
		PeerType:            in.PeerType,
		PeerID:              in.PeerId,
		DraftKind:           draftKind,
		Message:             message,
		EntitiesPayload:     entitiesPayload,
		DraftPayload:        payload,
		Date:                date,
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		OutboxID:            in.OutboxId,
		EventType:           "dialog.draftSaved",
	})
	if err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
