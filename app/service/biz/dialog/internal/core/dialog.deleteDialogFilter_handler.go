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

// DialogDeleteDialogFilter
// dialog.deleteDialogFilter user_id:long id:int = Bool;
func (c *DialogCore) DialogDeleteDialogFilter(in *dialog.TLDialogDeleteDialogFilter) (*tg.Bool, error) {
	sourceAuth, err := c.sourcePermAuthKeyID()
	if err != nil {
		return nil, err
	}
	operationID := deterministicOperationID("delete_filter", in.UserId, in.Id)
	if err := c.svcCtx.Repo.DeleteDialogFilter(c.ctx, repository.DeleteDialogFilterInput{
		UserID:              in.UserId,
		DialogFilterID:      in.Id,
		SourcePermAuthKeyID: sourceAuth,
		OperationID:         operationID,
		OutboxID:            deterministicOutboxID(operationID, "filter"),
		EventType:           "dialog.filterDeleted",
	}); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
