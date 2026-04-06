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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	synctypes "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = Updates;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*tg.Updates, error) {
	if len(in.Message) == 0 {
		return nil, tg.ErrInputRequestInvalid
	}

	switch in.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
		outbox := in.Message[0].ToOutboxMessage()
		if outbox == nil {
			return nil, tg.ErrInputRequestInvalid
		}

		date := int32(time.Now().Unix())
		var entities []tg.MessageEntityClazz
		if outbox.Message != nil {
			if msg2, ok := outbox.Message.(*tg.TLMessage); ok {
				if msg2.Date != 0 {
					date = msg2.Date
				}
				entities = msg2.Entities
			}
		}

		messageID := c.nextMessageId(outbox.RandomId)
		sentUpdates := tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
			Out:      true,
			Id:       messageID,
			Pts:      1,
			PtsCount: 1,
			Date:     date,
			Entities: entities,
		}).ToUpdates()

		// Persist the sent message before firing side effects.
		c.persistSentMessage(in, outbox, messageID, date)

		// Forward to recipient inbox and push sync updates if clients are wired.
		c.pushSendMessageSideEffects(in, outbox, messageID, date)

		return sentUpdates, nil
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}
}

// pushSendMessageSideEffects pushes inbox delivery and sync updates.
// Errors are logged but do not fail the send response.
func (c *MsgCore) pushSendMessageSideEffects(in *msg.TLMsgSendMessageV2, outbox *msg.OutboxMessage, messageID int32, date int32) {
	if c.svcCtx == nil {
		return
	}

	sentMessage := tg.MakeTLMessage(&tg.TLMessage{
		Out: true,
		Id:  messageID,
		FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: in.UserId,
		}),
		PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: in.PeerId,
		}),
		Date:    date,
		Message: "",
	})
	if outbox.Message != nil {
		if m, ok := outbox.Message.(*tg.TLMessage); ok {
			sentMessage.Message = m.Message
			sentMessage.Entities = m.Entities
		}
	}

	// Push to recipient inbox.
	if c.svcCtx.InboxClient != nil && in.PeerId != in.UserId {
		boxList := []tg.MessageBoxClazz{
			&tg.TLMessageBox{
				MessageId: messageID,
				Pts:       0,
				PtsCount:  1,
				Message:   sentMessage,
			},
		}
		_, err := c.svcCtx.InboxClient.InboxSendUserMessageToInboxV2(c.ctx, &inbox.TLInboxSendUserMessageToInboxV2{
			UserId:        in.PeerId,
			Out:           false,
			FromId:        in.UserId,
			FromAuthKeyId: in.AuthKeyId,
			PeerType:      in.PeerType,
			PeerId:        in.PeerId,
			BoxList:       boxList,
		})
		if err != nil {
			c.Logger.Errorf("msg.sendMessageV2 - inbox push error: %v", err)
		}
	}

	// Push sync updates to the sender's other sessions.
	if c.svcCtx.SyncClient != nil {
		senderUpdates := tg.MakeTLUpdates(&tg.TLUpdates{
			Updates: []tg.UpdateClazz{
				tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
					Message:  sentMessage,
					Pts:      1,
					PtsCount: 1,
				}),
			},
			Users: []tg.UserClazz{},
			Chats: []tg.ChatClazz{},
			Date:  date,
			Seq:   0,
		})

		_, err := c.svcCtx.SyncClient.SyncUpdatesNotMe(c.ctx, &synctypes.TLSyncUpdatesNotMe{
			UserId:        in.UserId,
			PermAuthKeyId: in.AuthKeyId,
			Updates:       senderUpdates,
		})
		if err != nil {
			c.Logger.Errorf("msg.sendMessageV2 - sync push error: %v", err)
		}

		// Push to recipient via sync.pushUpdates.
		if in.PeerId != in.UserId {
			recipientUpdates := tg.MakeTLUpdates(&tg.TLUpdates{
				Updates: []tg.UpdateClazz{
					tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
						Message:  sentMessage,
						Pts:      1,
						PtsCount: 1,
					}),
				},
				Users: []tg.UserClazz{},
				Chats: []tg.ChatClazz{},
				Date:  date,
				Seq:   0,
			})

			_, err = c.svcCtx.SyncClient.SyncPushUpdates(c.ctx, &synctypes.TLSyncPushUpdates{
				UserId:  in.PeerId,
				Updates: recipientUpdates,
			})
			if err != nil {
				c.Logger.Errorf("msg.sendMessageV2 - recipient sync push error: %v", err)
			}
		}
	}
}

// persistSentMessage stores the sent message in the message repository.
// Errors are logged but do not fail the send response.
func (c *MsgCore) persistSentMessage(in *msg.TLMsgSendMessageV2, outbox *msg.OutboxMessage, messageID int32, date int32) {
	if c.svcCtx == nil || c.svcCtx.Repository == nil {
		return
	}

	sentMessage := tg.MakeTLMessage(&tg.TLMessage{
		Out: true,
		Id:  messageID,
		FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: in.UserId,
		}),
		PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{
			UserId: in.PeerId,
		}),
		Date:    date,
		Message: "",
	})
	if outbox.Message != nil {
		if m, ok := outbox.Message.(*tg.TLMessage); ok {
			sentMessage.Message = m.Message
			sentMessage.Entities = m.Entities
		}
	}

	box := &repository.MessageBoxDO{
		UserId:       in.UserId,
		MessageId:    messageID,
		SenderUserId: in.UserId,
		PeerType:     in.PeerType,
		PeerId:       in.PeerId,
		RandomId:     outbox.RandomId,
		Message:      sentMessage,
		Pts:          1,
		PtsCount:     1,
	}

	if err := c.svcCtx.Repository.Message.PutMessage(c.ctx, box); err != nil {
		c.Logger.Errorf("msg.sendMessageV2 - persist message error: %v", err)
	}
}

// nextMessageId returns a real message ID from IdgenClient when wired,
// otherwise falls back to a placeholder derived from randomID.
func (c *MsgCore) nextMessageId(randomID int64) int32 {
	if c.svcCtx != nil && c.svcCtx.IdgenClient != nil {
		resp, err := c.svcCtx.IdgenClient.IdgenNextId(c.ctx, &idgen.TLIdgenNextId{})
		if err == nil && resp != nil {
			if tlInt64, ok := resp.Clazz.(*tg.TLInt64); ok && tlInt64.V > 0 {
				id := int32(tlInt64.V % 0x7fffffff)
				if id > 0 {
					return id
				}
			}
		}
	}
	// Fallback: deterministic placeholder from randomID.
	if randomID < 0 {
		randomID = -randomID
	}
	id := int32(randomID % 0x7fffffff)
	if id == 0 {
		id = 1
	}
	return id
}
