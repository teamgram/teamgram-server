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
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UpdatesGetDifference
// updates.getDifference#19c2f763 flags:# pts:int pts_limit:flags.1?int pts_total_limit:flags.0?int date:int qts:int qts_limit:flags.2?int = updates.Difference;
func (c *UpdatesCore) UpdatesGetDifference(in *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	client := c.svcCtx.Repo.UserupdatesClient
	if client == nil {
		return nil, fmt.Errorf("updates.getDifference: userupdates client is nil")
	}
	diff, err := client.UserupdatesGetDifference(c.ctx, &userupdates.TLUserupdatesGetDifference{
		UserId:        c.MD.UserId,
		AuthKeyId:     c.MD.PermAuthKeyId,
		Pts:           int64(in.Pts),
		PtsTotalLimit: in.PtsTotalLimit,
		Date:          int64Ptr(int64(in.Date)),
	})
	if err != nil {
		return nil, err
	}
	return userDifferenceToUpdatesDifference(diff)
}

func userDifferenceToUpdatesDifference(diff *userupdates.UserDifference) (*tg.UpdatesDifference, error) {
	if diff == nil {
		return nil, fmt.Errorf("updates.getDifference: user difference is nil")
	}
	if empty, ok := diff.ToUserDifferenceEmpty(); ok {
		return tg.MakeTLUpdatesDifferenceEmpty(&tg.TLUpdatesDifferenceEmpty{
			Date: userStateDate(empty.State),
			Seq:  userStateSeq(empty.State),
		}).ToUpdatesDifference(), nil
	}
	if full, ok := diff.ToUserDifference(); ok {
		return tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
			NewMessages:          full.NewMessages,
			NewEncryptedMessages: []tg.EncryptedMessageClazz{},
			OtherUpdates:         full.OtherUpdates,
			Chats:                []tg.ChatClazz{},
			Users:                []tg.UserClazz{},
			State:                userStateToUpdatesState(full.State),
		}).ToUpdatesDifference(), nil
	}
	return nil, fmt.Errorf("updates.getDifference: unsupported user difference %s", diff.ClazzName())
}

func userStateToUpdatesState(state userupdates.UserStateClazz) tg.UpdatesStateClazz {
	if state == nil {
		return tg.MakeTLUpdatesState(&tg.TLUpdatesState{})
	}
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         int32(state.Pts),
		Qts:         state.Qts,
		Date:        state.Date,
		Seq:         state.Seq,
		UnreadCount: state.UnreadCount,
	})
}

func userStateDate(state userupdates.UserStateClazz) int32 {
	if state == nil {
		return 0
	}
	return state.Date
}

func userStateSeq(state userupdates.UserStateClazz) int32 {
	if state == nil {
		return 0
	}
	return state.Seq
}

func int64Ptr(v int64) *int64 {
	return &v
}
