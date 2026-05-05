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

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// DialogCreateDialogFilter
// dialog.createDialogFilter user_id:long dialog_filter:DialogFilterExt = DialogFilterExt;
func (c *DialogCore) DialogCreateDialogFilter(in *dialog.TLDialogCreateDialogFilter) (*dialog.DialogFilterExt, error) {
	sourceAuth, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(in.DialogFilter)
	title := ""
	filterID := in.DialogFilter.Id
	order := in.DialogFilter.Order
	if filter := in.DialogFilter.DialogFilter; filter != nil {
		if f, ok := filter.(*tg.TLDialogFilter); ok && f.Title != nil {
			title = f.Title.Text
		}
	}
	operationID := deterministicOperationID("create_filter", in.UserId, filterID, in.DialogFilter.Slug, order)
	record, err := c.svcCtx.Repo.SaveDialogFilter(c.ctx, repository.SaveDialogFilterInput{
		UserID:              in.UserId,
		DialogFilterID:      filterID,
		Slug:                in.DialogFilter.Slug,
		Title:               title,
		OrderValue:          order,
		Enabled:             true,
		FilterSchemaVersion: 1,
		FilterPayload:       payload,
		SourcePermAuthKeyID: sourceAuth,
		OperationID:         operationID,
		OutboxID:            deterministicOutboxID(operationID, "filter"),
		EventType:           "dialog.filterUpdated",
	})
	if err != nil {
		return nil, err
	}
	return makeDialogFilterExt(*record), nil
}
