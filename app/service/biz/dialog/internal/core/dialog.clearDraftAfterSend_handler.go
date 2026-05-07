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

// DialogClearDraftAfterSend
// dialog.clearDraftAfterSend user_id:long peer_type:int peer_id:long clear_before_date:int source_perm_auth_key_id:long source_operation_id:string outbox_id:long = Bool;
func (c *DialogCore) DialogClearDraftAfterSend(in *dialog.TLDialogClearDraftAfterSend) (*tg.Bool, error) {
	if in == nil {
		return nil, dialog.ErrDialogInvalid
	}
	if in.SourceOperationId == "" {
		return nil, dialog.ErrOutboxUnavailable
	}
	_, err := c.svcCtx.Repo.ClearDraftAfterSend(c.ctx, repository.ClearDraftAfterSendInput{
		UserID:              in.UserId,
		PeerType:            in.PeerType,
		PeerID:              in.PeerId,
		ClearBeforeDate:     int64(in.ClearBeforeDate),
		SourcePermAuthKeyID: in.SourcePermAuthKeyId,
		OperationID:         in.SourceOperationId,
		OutboxID:            in.OutboxId,
		EventType:           "dialog.draftClearedAfterSend",
		Payload:             []byte(`{"schema_version":1}`),
	})
	if err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
