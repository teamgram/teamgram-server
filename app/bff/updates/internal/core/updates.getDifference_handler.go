// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

// UpdatesGetDifference
// updates.getDifference#19c2f763 flags:# pts:int pts_limit:flags.1?int pts_total_limit:flags.0?int date:int qts:int qts_limit:flags.2?int = updates.Difference;
func (c *UpdatesCore) UpdatesGetDifference(in *tg.TLUpdatesGetDifference) (*tg.UpdatesDifference, error) {
	pts := in.Pts
	if pts <= 0 {
		pts = 1
	}
	date := in.Date
	if date <= 0 {
		date = 10
	}
	if in.Pts < 1 {
		message := makePlaceholderBFFDifferenceMessage(pts, date)
		return tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
			NewMessages: []tg.MessageClazz{
				message,
			},
			NewEncryptedMessages: []tg.EncryptedMessageClazz{},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message:  message,
					Pts:      pts,
					PtsCount: 1,
				}),
			},
			Chats: []tg.ChatClazz{},
			Users: []tg.UserClazz{},
			State: makePlaceholderUpdatesState(pts, date),
		}).ToUpdatesDifference(), nil
	}

	return tg.MakeTLUpdatesDifferenceEmpty(&tg.TLUpdatesDifferenceEmpty{
		Date: date,
		Seq:  0,
	}).ToUpdatesDifference(), nil
}
