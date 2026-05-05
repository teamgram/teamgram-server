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

// DialogToggleDialogPin
// dialog.toggleDialogPin user_id:long peer_type:int peer_id:long pinned:Bool = Int32;
func (c *DialogCore) DialogToggleDialogPin(in *dialog.TLDialogToggleDialogPin) (*tg.Int32, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	result, err := c.svcCtx.Repo.ToggleDialogPin(c.ctx, repository.ToggleDialogPinInput{
		UserID:              in.UserId,
		PeerType:            in.PeerType,
		PeerID:              in.PeerId,
		Pinned:              tg.FromBoolClazz(in.Pinned),
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.OperationId,
		OutboxID:            in.OutboxId,
		EventType:           "dialog.pinToggled",
		Payload:             []byte(`{"schema_version":1}`),
	})
	if err != nil {
		return nil, err
	}
	return tg.MakeInt32(int32(result.AggregateVersion)), nil
}
