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
	clientPts := in.Pts
	if clientPts < 0 {
		clientPts = 0
	}
	clientDate := int32(in.Date)
	if clientDate <= 0 {
		clientDate = 10
	}

	// Use stored state if available, otherwise fall back to normalized client defaults.
	currentState, err := c.getUserUpdatesStateForDifference(in.UserId, clientPts, clientDate)
	if err != nil {
		return nil, err
	}

	// Completely fresh client (pts=0 and date=0 from original request) should always get catch-up.
	if in.Pts <= 0 && in.Date <= 0 {
		catchUpMessage := c.makeDifferenceMessage(in.UserId, clientPts)
		return updates.MakeTLDifference(&updates.TLDifference{
			NewMessages: []tg.MessageClazz{
				catchUpMessage,
			},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message:  catchUpMessage,
					Pts:      currentState.Pts,
					PtsCount: currentState.Pts - clientPts,
				}),
			},
			State: currentState,
		}).ToDifference(), nil
	}

	// For normal requests, compare with stored state.
	if clientPts < currentState.Pts {
		// Client is behind; return catch-up with messages from message service.
		catchUpMessage := c.makeDifferenceMessage(in.UserId, clientPts)
		return updates.MakeTLDifference(&updates.TLDifference{
			NewMessages: []tg.MessageClazz{
				catchUpMessage,
			},
			OtherUpdates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message:  catchUpMessage,
					Pts:      currentState.Pts,
					PtsCount: currentState.Pts - clientPts,
				}),
			},
			State: currentState,
		}).ToDifference(), nil
	}

	// Client is up-to-date; return empty difference with current state.
	return updates.MakeTLDifferenceEmpty(&updates.TLDifferenceEmpty{
		State: currentState,
	}).ToDifference(), nil
}

func (c *UpdatesCore) makeDifferenceMessage(userID int64, maxPts int32) tg.MessageClazz {
	if c != nil && c.svcCtx != nil && c.svcCtx.MessageClient != nil && userID != 0 {
		boxes, err := c.svcCtx.MessageClient.MessageGetHistoryMessages(c.ctx, &message.TLMessageGetHistoryMessages{
			UserId:   userID,
			PeerType: tg.PEER_USER,
			PeerId:   userID,
			MaxId:    maxPts,
			Limit:    1,
		})
		if err == nil && boxes != nil && len(boxes.Datas) > 0 {
			if box := boxes.Datas[0]; box != nil && box.Message != nil {
				return box.Message
			}
		}
	}

	// Fallback placeholder.
	return tg.MakeTLMessage(&tg.TLMessage{
		Out: true,
		Id:  maxPts,
		FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: userID,
		}),
		PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: userID,
		}),
		Date:    10,
		Message: "placeholder",
	})
}
