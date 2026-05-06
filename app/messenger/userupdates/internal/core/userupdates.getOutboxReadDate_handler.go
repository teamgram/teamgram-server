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
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UserupdatesGetOutboxReadDate
// userupdates.getOutboxReadDate user_id:long peer_type:int peer_id:long msg_id:int = OutboxReadDate;
func (c *UserupdatesCore) UserupdatesGetOutboxReadDate(in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.UserId <= 0 || in.PeerId <= 0 || in.MsgId <= 0 {
		return nil, tg.ErrMessageIdInvalid
	}
	readRepo, ok := c.svcCtx.Repo.(interface {
		GetOutboxReadDate(context.Context, repository.OutboxReadDateInput) (int32, error)
	})
	if !ok {
		return nil, tg.ErrMessageNotReadYet
	}
	date, err := readRepo.GetOutboxReadDate(c.ctx, repository.OutboxReadDateInput{
		UserID:   in.UserId,
		PeerType: in.PeerType,
		PeerID:   in.PeerId,
		MsgID:    in.MsgId,
	})
	if err != nil {
		return nil, err
	}
	return tg.MakeTLOutboxReadDate(&tg.TLOutboxReadDate{
		Date: date,
	}).ToOutboxReadDate(), nil
}
