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
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogInsertOrUpdateDialogFilter
// dialog.insertOrUpdateDialogFilter user_id:long id:int dialog_filter:DialogFilter = Bool;
func (c *DialogCore) DialogInsertOrUpdateDialogFilter(in *dialog.TLDialogInsertOrUpdateDialogFilter) (*tg.Bool, error) {
	sourceAuth, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(in.DialogFilter)
	title := ""
	if f, ok := in.DialogFilter.(*tg.TLDialogFilter); ok && f.Title != nil {
		title = f.Title.Text
	}
	operationID := deterministicOperationID("upsert_filter", in.UserId, in.Id, payloadDigest(payload))
	if _, err := c.svcCtx.Repo.SaveDialogFilter(c.ctx, repository.SaveDialogFilterInput{
		UserID:              in.UserId,
		DialogFilterID:      in.Id,
		Title:               title,
		OrderValue:          int64(in.Id),
		Enabled:             true,
		FilterSchemaVersion: 1,
		FilterPayload:       payload,
		SourcePermAuthKeyID: sourceAuth,
		OperationID:         operationID,
		OutboxID:            deterministicOutboxID(operationID, "filter"),
		EventType:           "dialog.filterUpdated",
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}

func payloadDigest(payload []byte) string {
	sum := sha256.Sum256(payload)
	return fmt.Sprintf("%x", sum[:])
}
