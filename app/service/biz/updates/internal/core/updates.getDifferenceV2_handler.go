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

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/message/message"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/updates/updates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// UpdatesGetDifferenceV2
// updates.getDifferenceV2 flags:# auth_key_id:long user_id:long pts:int pts_total_limit:flags.0?int date:long = Difference;
func (c *UpdatesCore) UpdatesGetDifferenceV2(in *updates.TLUpdatesGetDifferenceV2) (*updates.Difference, error) {
	pts := in.Pts
	if pts <= 0 {
		pts = 1
	}

	date := int32(in.Date)
	if date <= 0 {
		date = 10
	}

	if in.Pts < 1 {
		message := c.makeDifferenceMessage(in.UserId, pts, date)
		return updates.MakeTLDifference(&updates.TLDifference{
			NewMessages: []tg.MessageClazz{
				message,
			},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message:  message,
					Pts:      pts,
					PtsCount: 1,
				}),
			},
			State: makePlaceholderUpdatesState(pts, date),
		}).ToDifference(), nil
	}

	// TODO: return real merged updates once the updates storage layer is wired.
	return updates.MakeTLDifferenceEmpty(&updates.TLDifferenceEmpty{
		State: makePlaceholderUpdatesState(pts, date),
	}).ToDifference(), nil
}

func (c *UpdatesCore) makeDifferenceMessage(userID int64, pts int32, date int32) tg.MessageClazz {
	if c != nil && c.svcCtx != nil && c.svcCtx.MessageClient != nil && userID != 0 {
		boxes, err := c.svcCtx.MessageClient.MessageGetHistoryMessages(c.ctx, &message.TLMessageGetHistoryMessages{
			UserId:   userID,
			PeerType: tg.PEER_USER,
			PeerId:   userID,
			MaxId:    pts,
			Limit:    1,
		})
		if err == nil && boxes != nil && len(boxes.Datas) > 0 {
			if box := boxes.Datas[0]; box != nil && box.Message != nil {
				return box.Message
			}
		}
	}

	return makePlaceholderDifferenceMessage(userID, pts, date)
}

func makePlaceholderUpdatesState(pts int32, date int32) *tg.UpdatesState {
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Pts:         pts,
		Qts:         0,
		Date:        date,
		Seq:         0,
		UnreadCount: 0,
	}).ToUpdatesState()
}

func makePlaceholderDifferenceMessage(userID int64, messageID int32, date int32) tg.MessageClazz {
	return tg.MakeTLMessage(&tg.TLMessage{
		Out: true,
		Id:  messageID,
		FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: userID,
		}),
		PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: userID,
		}),
		Date:    date,
		Message: "placeholder",
	})
}
