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
	"math"

	chatprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/chatprojection"
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// UpdatesGetDifference
// updates.getDifference#19c2f763 flags:# pts:int pts_limit:flags.1?int pts_total_limit:flags.0?int date:int qts:int qts_limit:flags.2?int = updates.Difference;
func (c *UpdatesCore) UpdatesGetDifference(in *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error) {
	userID, permAuthKeyID, err := c.requireUserAndPermAuthKey()
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	client := c.svcCtx.Repo.UserupdatesClient
	if client == nil {
		return nil, fmt.Errorf("updates.getDifference: userupdates client is nil")
	}
	diff, err := client.UserupdatesGetDifference(c.ctx, &userupdates.TLUserupdatesGetDifference{
		UserId:        userID,
		AuthKeyId:     permAuthKeyID,
		Pts:           int64(in.Pts),
		PtsTotalLimit: in.PtsTotalLimit,
		Date:          int64Ptr(int64(in.Date)),
	})
	if err != nil {
		return nil, err
	}
	publicDiff, err := userDifferenceToUpdatesDifference(diff)
	if err != nil {
		return nil, err
	}
	if err := userprojection.FillDifferenceUsers(c.ctx, c.svcCtx.Repo.UserClient, userID, publicDiff, userprojection.MissingStoredReference); err != nil {
		return nil, err
	}
	if err := chatprojection.FillDifferenceChats(c.ctx, c.svcCtx.Repo.ChatClient, userID, publicDiff, chatprojection.MissingStoredReference); err != nil {
		return nil, err
	}
	return publicDiff, nil
}

func (c *UpdatesCore) requireUserAndPermAuthKey() (int64, int64, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return 0, 0, tg.ErrUserIdInvalid
	}
	if c.MD.PermAuthKeyId == 0 {
		return 0, 0, tg.ErrAuthKeyPermEmpty
	}
	return c.MD.UserId, c.MD.PermAuthKeyId, nil
}

func userDifferenceToUpdatesDifference(diff *userupdates.UserDifference) (*tg.UpdatesDifference, error) {
	if diff == nil {
		return nil, fmt.Errorf("updates.getDifference: user difference is nil")
	}
	if empty, ok := diff.ToUserDifferenceEmpty(); ok {
		state, err := userStateToUpdatesState(empty.State)
		if err != nil {
			return nil, err
		}
		return tg.MakeTLUpdatesDifferenceEmpty(&tg.TLUpdatesDifferenceEmpty{
			Date: state.Date,
			Seq:  state.Seq,
		}).ToUpdatesDifference(), nil
	}
	if full, ok := diff.ToUserDifference(); ok {
		state, err := userStateToUpdatesState(full.State)
		if err != nil {
			return nil, err
		}
		newMessages, otherUpdates := mergeUpdateNewMessages(full.NewMessages, full.OtherUpdates)
		return tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
			NewMessages:          newMessages,
			NewEncryptedMessages: []tg.EncryptedMessageClazz{},
			OtherUpdates:         otherUpdates,
			Chats:                []tg.ChatClazz{},
			Users:                []tg.UserClazz{},
			State:                state,
		}).ToUpdatesDifference(), nil
	}
	if slice, ok := diff.ToUserDifferenceSlice(); ok {
		intermediateState, err := userStateToUpdatesState(slice.IntermediateState)
		if err != nil {
			return nil, err
		}
		newMessages, otherUpdates := mergeUpdateNewMessages(slice.NewMessages, slice.OtherUpdates)
		return tg.MakeTLUpdatesDifferenceSlice(&tg.TLUpdatesDifferenceSlice{
			NewMessages:          newMessages,
			NewEncryptedMessages: []tg.EncryptedMessageClazz{},
			OtherUpdates:         otherUpdates,
			Chats:                []tg.ChatClazz{},
			Users:                []tg.UserClazz{},
			IntermediateState:    intermediateState,
		}).ToUpdatesDifference(), nil
	}
	if tooLong, ok := diff.ToUserDifferenceTooLong(); ok {
		pts, err := checkedPublicInt32(tooLong.Pts, "updates.differenceTooLong.pts")
		if err != nil {
			return nil, err
		}
		return tg.MakeTLUpdatesDifferenceTooLong(&tg.TLUpdatesDifferenceTooLong{
			Pts: pts,
		}).ToUpdatesDifference(), nil
	}
	return nil, fmt.Errorf("updates.getDifference: unsupported user difference %s", diff.ClazzName())
}

func mergeUpdateNewMessages(messages []tg.MessageClazz, updates []tg.UpdateClazz) ([]tg.MessageClazz, []tg.UpdateClazz) {
	if len(updates) == 0 {
		return messages, updates
	}
	mergedMessages := append([]tg.MessageClazz{}, messages...)
	seenMessageIDs := make(map[int32]struct{}, len(messages)+len(updates))
	for _, message := range messages {
		if id, ok := updateMessageID(message); ok {
			seenMessageIDs[id] = struct{}{}
		}
	}
	otherUpdates := make([]tg.UpdateClazz, 0, len(updates))
	for _, update := range updates {
		if updateNewMessage, ok := update.(*tg.TLUpdateNewMessage); ok {
			if updateNewMessage.Message != nil {
				if id, ok := updateMessageID(updateNewMessage.Message); ok {
					if _, seen := seenMessageIDs[id]; seen {
						continue
					}
					seenMessageIDs[id] = struct{}{}
				}
				mergedMessages = append(mergedMessages, updateNewMessage.Message)
			}
			continue
		}
		otherUpdates = append(otherUpdates, update)
	}
	return mergedMessages, otherUpdates
}

func updateMessageID(message tg.MessageClazz) (int32, bool) {
	switch m := message.(type) {
	case *tg.TLMessage:
		if m == nil {
			return 0, false
		}
		return m.Id, true
	case *tg.TLMessageService:
		if m == nil {
			return 0, false
		}
		return m.Id, true
	default:
		return 0, false
	}
}

func userStateToUpdatesState(state userupdates.UserStateClazz) (*tg.UpdatesState, error) {
	if state == nil {
		return tg.MakeTLUpdatesState(&tg.TLUpdatesState{}).ToUpdatesState(), nil
	}
	pts, err := checkedPublicInt32(state.Pts, "updates.state.pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         pts,
		Qts:         state.Qts,
		Date:        state.Date,
		Seq:         state.Seq,
		UnreadCount: state.UnreadCount,
	}).ToUpdatesState(), nil
}

func checkedPublicInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%s out of int32 range: %d", field, v)
	}
	return int32(v), nil
}

func int64Ptr(v int64) *int64 {
	return &v
}
