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

// DialogToggleSavedDialogPin
// dialog.toggleSavedDialogPin user_id:long peer:PeerUtil pinned:Bool = Bool;
func (c *DialogCore) DialogToggleSavedDialogPin(in *dialog.TLDialogToggleSavedDialogPin) (*tg.Bool, error) {
	if in == nil || in.Peer == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	sourcePermAuthKeyID, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	operationID := deterministicOperationID("saved_dialog_pin", in.UserId, in.Peer.PeerType, in.Peer.PeerId, tg.FromBoolClazz(in.Pinned))
	if err := c.svcCtx.Repo.ToggleSavedDialogPin(c.ctx, repository.SavedDialogPinInput{
		UserID:              in.UserId,
		PeerType:            in.Peer.PeerType,
		PeerID:              in.Peer.PeerId,
		Pinned:              tg.FromBoolClazz(in.Pinned),
		SourcePermAuthKeyID: sourcePermAuthKeyID,
		OperationID:         operationID,
		OutboxID:            deterministicOutboxID(operationID, "actor"),
		EventType:           "dialog.savedDialogPinned",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
