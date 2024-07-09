// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// InboxEditMessageToInboxV2
// inbox.editMessageToInboxV2 flags:# user_id:long out:flags.0?true from_id:long fromAuthKeyId:long peer_type:int peer_id:long box:MessageBox users:flags.1?Vector<User> chats:flags.2?Vector<Chat> = Void;
func (c *InboxCore) InboxEditMessageToInboxV2(in *inbox.TLInboxEditMessageToInboxV2) (*mtproto.Void, error) {
	if in.Out {
		var (
			mData, _ = jsonx.Marshal(in.NewMessage.Message)
		)

		if _, err := c.svcCtx.Dao.MessagesDAO.UpdateEditMessage(c.ctx, string(mData), in.NewMessage.Message.Message, in.UserId, in.NewMessage.MessageId); err != nil {
			c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
			return nil, err
		}

		// HashTagsDAO
		for _, entity := range in.DstMessage.Message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					c.svcCtx.Dao.HashTagsDAO.DeleteHashTagMessageId(c.ctx, in.UserId, in.DstMessage.MessageId)
					break
				}
			}
		}

		// c.svcCtx.Dao.HashTagsDAO.DeleteHashTagMessageId(c.ctx, in.FromId, in.NewMessage.MessageId)
		for _, entity := range in.NewMessage.Message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					c.svcCtx.Dao.HashTagsDAO.InsertOrUpdate(c.ctx, &dataobject.HashTagsDO{
						UserId:           in.UserId,
						PeerType:         in.PeerType,
						PeerId:           in.PeerId,
						HashTag:          entity.GetUrl(),
						HashTagMessageId: in.NewMessage.MessageId,
					})
				}
			}
		}

		_, err := c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
			UserId:        in.UserId,
			PermAuthKeyId: in.FromAuthKeyId,
			Updates: mtproto.MakeUpdatesByUpdatesUsersChats(
				in.Users,
				in.Chats,
				mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
					Pts_INT32:       in.NewMessage.Pts,
					PtsCount:        in.NewMessage.PtsCount,
					Message_MESSAGE: in.NewMessage.Message,
				}).To_Update()),
		})
		if err != nil {
			c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
			return nil, err
		}
	} else {
		var (
			newMessage = in.NewMessage.Message
			dstMessage *mtproto.Message
			pts        int32
			ptsCount   int32 = 1
		)

		pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId)
		if pts == 0 {
			c.Logger.Errorf("NextPtsId error: %v", in.UserId)
			err := mtproto.ErrInternalServerError
			return nil, err
		}

		dstMessageDO, _ := c.svcCtx.Dao.MessagesDAO.SelectByMessageDataId(c.ctx, in.UserId, in.NewMessage.DialogMessageId)
		if dstMessageDO == nil {
			return mtproto.EmptyVoid, nil
		}

		// newMessage.MessageId = dstMessageDO.UserMessageBoxId

		err := jsonx.UnmarshalFromString(dstMessageDO.MessageData, &dstMessage)
		if err != nil {
			c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
			return nil, err
		}

		tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			// HashTagsDAO
			for _, entity := range dstMessage.Entities {
				if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
					if entity.GetUrl() != "" {
						c.svcCtx.Dao.HashTagsDAO.DeleteHashTagMessageIdTx(tx, in.UserId, dstMessage.Id)
						break
					}
				}
			}

			dstMessage.Message = newMessage.Message
			dstMessage.Media = newMessage.Media
			dstMessage.ReplyMarkup = newMessage.ReplyMarkup
			dstMessage.Entities = newMessage.Entities
			dstMessage.EditDate = newMessage.EditDate
			dstMessage.EditHide = newMessage.EditHide

			newMessage = dstMessage

			mData, _ := jsonx.Marshal(newMessage)
			_, err = c.svcCtx.Dao.MessagesDAO.UpdateEditMessageTx(tx, string(mData), newMessage.Message, in.UserId, newMessage.Id)
			if err != nil {
				c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
				result.Err = err
				return
			}

			// c.svcCtx.Dao.HashTagsDAO.DeleteHashTagMessageId(c.ctx, in.FromId, in.NewMessage.MessageId)
			for _, entity := range in.NewMessage.Message.GetEntities() {
				if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
					if entity.GetUrl() != "" {
						c.svcCtx.Dao.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
							UserId:           in.UserId,
							PeerType:         in.PeerType,
							PeerId:           in.PeerId,
							HashTag:          entity.GetUrl(),
							HashTagMessageId: newMessage.Id,
						})
					}
				}
			}
		})
		if tR.Err != nil {
			return nil, tR.Err
		}

		var (
			isBot = false
		)

		for _, u := range in.GetUsers() {
			if u.GetId() == in.UserId {
				isBot = u.GetBot()
				break
			}
		}

		pushUpdates := mtproto.MakeUpdatesByUpdatesUsersChats(
			in.Users,
			in.Chats,
			mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
				Pts_INT32:       pts,
				PtsCount:        ptsCount,
				Message_MESSAGE: newMessage,
			}).To_Update())

		if isBot {
			if c.svcCtx.Dao.BotSyncClient != nil {
				_, err = c.svcCtx.Dao.BotSyncClient.SyncPushBotUpdates(c.ctx, &sync.TLSyncPushBotUpdates{
					UserId:  in.UserId,
					Updates: pushUpdates,
				})
			} else {
				// TODO: log
			}
		} else {
			_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  in.UserId,
				Updates: pushUpdates,
			})
		}
		if err != nil {
			c.Logger.Errorf("inbox.editMessageToInboxV2 - error: %v", err)
			return nil, err
		}
	}

	return mtproto.EmptyVoid, nil
}
