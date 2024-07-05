// Copyright 2022 Teamgram Authors
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

package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/marmota/pkg/container2/sets"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

// from outBox --> make inBox
func (d *Dao) makeMessageInBox(fromId int64, peer *mtproto.PeerUtil, toUserId int64, clientRandomId int64, dialogMessageId int64, message *mtproto.Message) *mtproto.MessageBox {
	mentioned := mtproto.CheckHasMention(message.Entities, toUserId)
	logx.Infof("insert to inbox: %#v, message: {%#v}", mentioned, message)

	did := mtproto.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
	// from outBox --> make inBox
	return &mtproto.MessageBox{
		UserId:            fromId,
		MessageId:         0,
		DialogId1:         did.A,
		DialogId2:         did.B,
		DialogMessageId:   dialogMessageId,
		RandomId:          clientRandomId,
		MessageFilterType: mtproto.GetMediaType(message),
	}
}

func (d *Dao) sendMessageToInbox(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, toUserId int64, dialogMessageId, clientRandomId int64, message2 *mtproto.Message) (*mtproto.MessageBox, error) {
	var (
		inBoxMsgId = d.IDGenClient2.NextMessageBoxId(ctx, toUserId)
		dialogId   = mtproto.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		date       = time.Now().Unix()
		message    = proto.Clone(message2).(*mtproto.Message)
	)

	if peer.PeerType == mtproto.PEER_USER {
		if dialogMessageId == 0 {
			dialogMessageId = d.IDGenClient2.NextId(ctx)
		}
	}

	// fix message
	message.Out = false
	message.Id = inBoxMsgId
	switch message.GetReplyTo().GetPredicateName() {
	case mtproto.Predicate_messageReplyHeader:
		if replyId, _ := d.MessagesDAO.SelectPeerUserMessage(ctx, toUserId, fromId, message.GetReplyTo().GetFixedReplyToMsgId()); replyId != nil {
			// message.ReplyToMsgId.Value = replyId.UserMessageBoxId
			if message.ReplyTo != nil {
				message.ReplyTo.ReplyToMsgId = replyId.UserMessageBoxId
				message.ReplyTo.ReplyToMsgId_INT32 = replyId.UserMessageBoxId
				message.ReplyTo.ReplyToMsgId_FLAGINT32 = mtproto.MakeFlagsInt32(replyId.UserMessageBoxId)
			}

			if peer.PeerType == mtproto.PEER_CHAT && replyId.SenderUserId == toUserId {
				message.Mentioned = true
				if message2.GetAction().GetPredicateName() != mtproto.Predicate_messageActionPinMessage {
					message.MediaUnread = true
				}
			}
		} else {
			// message.ReplyToMsgId.Value = 0
			message.ReplyTo = nil
		}
	case mtproto.Predicate_messageReplyStoryHeader:
		// do nothing
	default:
		// do nothing
	}

	if peer.PeerType == mtproto.PEER_CHAT {
		if !message.Mentioned {
			message.Mentioned = mtproto.CheckHasMention(message.Entities, toUserId)
			if message.Mentioned {
				message.MediaUnread = true
			}
		}
	} else if peer.PeerType == mtproto.PEER_USER {
		message.FromId = nil
		message.PeerId = mtproto.MakePeerUser(fromId)
	}

	if !message.MediaUnread {
		message.MediaUnread = mtproto.CheckHasMediaUnread(message)
	}

	if peer.PeerType == mtproto.PEER_CHAT {
		if message2.GetAction().GetPredicateName() == mtproto.Predicate_messageActionGroupCall {
			call := message2.GetAction()
			if call != nil && len(call.Users) > 0 {
				if ok := container2.ContainsInt64(call.Users, toUserId); ok {
					message.MediaUnread = true
					message.Mentioned = true
				}
			}
		}
	}

	mData, _ := jsonx.Marshal(message)
	// mType, mData := mtproto.EncodeMessage(message)
	inBox := &mtproto.MessageBox{
		UserId:            toUserId,
		SenderUserId:      fromId,
		PeerType:          peer.PeerType,
		PeerId:            peer.PeerId,
		MessageId:         inBoxMsgId,
		DialogId1:         dialogId.A,
		DialogId2:         dialogId.B,
		DialogMessageId:   dialogMessageId,
		RandomId:          clientRandomId,
		Pts:               0,
		PtsCount:          0,
		MessageFilterType: mtproto.GetMediaType(message),
		Message:           message,
		Mentioned:         message.Mentioned,
		MediaUnread:       message.MediaUnread,
	}

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// TODO(@benqi): do ignore

		// Pts:              pts,
		// PtsCount:         ptsCount,
		inBoxDO := &dataobject.MessagesDO{
			UserId:            inBox.UserId,
			UserMessageBoxId:  inBox.MessageId,
			DialogId1:         inBox.DialogId1,
			DialogId2:         inBox.DialogId2,
			SenderUserId:      fromId,
			PeerType:          peer.PeerType,
			PeerId:            inBox.PeerId,
			RandomId:          inBox.RandomId,
			DialogMessageId:   inBox.DialogMessageId,
			MessageData:       string(mData),
			MessageFilterType: inBox.MessageFilterType, // TODO(@benqi): message_type
			Message:           message.Message,
			Mentioned:         inBox.Mentioned,
			MediaUnread:       inBox.MediaUnread,
			Date2:             date,
			Deleted:           false,
		}

		_, _, result.Err = d.MessagesDAO.InsertOrReturnIdTx(tx, inBoxDO)
		if result.Err != nil {
			return
		}

		switch peer.PeerType {
		case mtproto.PEER_USER:
			var (
				lastInsertId int64
				rowsAffected int64
			)

			dialogDO := &dataobject.DialogsDO{
				UserId:           inBox.UserId,
				PeerType:         peer.PeerType,
				PeerId:           fromId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_USER, fromId),
				TopMessage:       inBoxMsgId,
				UnreadCount:      1,
				DraftMessageData: "null",
				Date2:            date,
			}

			lastInsertId, rowsAffected, result.Err = d.DialogsDAO.InsertOrUpdateTx(tx, dialogDO)
			logx.WithContext(ctx).Infof("lastInsertId:%d, rowsAffected: %d, result: %v, do: %v", lastInsertId, rowsAffected, result, dialogDO)
		case mtproto.PEER_CHAT:
			var (
				lastInsertId int64
				rowsAffected int64
			)

			dialogDO := &dataobject.DialogsDO{
				UserId:               inBox.UserId,
				PeerType:             peer.PeerType,
				PeerId:               peer.PeerId,
				PeerDialogId:         mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId),
				Pinned:               0,
				TopMessage:           inBoxMsgId,
				PinnedMsgId:          0,
				ReadInboxMaxId:       0,
				ReadOutboxMaxId:      0,
				UnreadCount:          1,
				UnreadMentionsCount:  0,
				UnreadReactionsCount: 0,
				UnreadMark:           false,
				DraftType:            0,
				DraftMessageData:     "null",
				FolderId:             0,
				FolderPinned:         0,
				HasScheduled:         false,
				TtlPeriod:            0,
				ThemeEmoticon:        "",
				Date2:                date,
			}
			if inBox.Mentioned {
				dialogDO.UnreadMentionsCount = 1
			}

			lastInsertId, rowsAffected, result.Err = d.DialogsDAO.InsertOrUpdateTx(tx, dialogDO)
			logx.WithContext(ctx).Infof("lastInsertId:%d, rowsAffected: %d, result: %v, do: %v", lastInsertId, rowsAffected, result, dialogDO)
			if result.Err != nil {
				return
			}
		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
		}

		for _, entity := range message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					d.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
						UserId:           inBox.UserId,
						PeerType:         peer.PeerType,
						PeerId:           peer.PeerId,
						HashTag:          entity.GetUrl(),
						HashTagMessageId: inBox.MessageId,
					})
				}
			}
		}
	})

	// TODO(@benqi): process duplicate

	if tR.Err != nil {
		return nil, tR.Err
	}

	inBox.Pts = d.IDGenClient2.NextPtsId(ctx, toUserId)
	inBox.PtsCount = 1

	return inBox, nil
}

func (d *Dao) SendUserMessageToInbox(ctx context.Context, fromId, toId int64, dialogMessageId, clientRandomId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_USER,
		PeerId:   toId,
	}
	message.Out = false
	return d.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageId, clientRandomId, message)
}

func (d *Dao) SendChatMessageToInbox(ctx context.Context, fromId, chatId, toId int64, dialogMessageId, clientRandomId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_CHAT,
		PeerId:   chatId,
	}
	message.Out = false
	return d.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageId, clientRandomId, message)
}

func (d *Dao) SendUserMultiMessageToInbox(ctx context.Context, fromId, toId int64, inBoxList []*inbox.InboxMessageData) ([]*mtproto.MessageBox, error) {
	var (
		boxList = make([]*mtproto.MessageBox, 0, len(inBoxList))
	)

	for _, box := range inBoxList {
		peer := &mtproto.PeerUtil{
			PeerType: mtproto.PEER_USER,
			PeerId:   toId,
		}
		box.Message.Out = false
		inBox, _ := d.sendMessageToInbox(ctx, fromId, peer, toId, box.DialogMessageId, box.RandomId, box.Message)
		boxList = append(boxList, inBox)
	}

	return boxList, nil
}

func (d *Dao) SendChatMultiMessageToInbox(ctx context.Context, fromId, chatId, toId int64, inBoxList []*inbox.InboxMessageData) ([]*mtproto.MessageBox, error) {
	var (
		boxList = make([]*mtproto.MessageBox, 0, len(inBoxList))
	)
	for _, box := range inBoxList {
		peer := &mtproto.PeerUtil{
			PeerType: mtproto.PEER_CHAT,
			PeerId:   chatId,
		}
		box.Message.Out = false
		inBox, _ := d.sendMessageToInbox(ctx, fromId, peer, toId, box.DialogMessageId, box.RandomId, box.Message)
		boxList = append(boxList, inBox)
	}

	return boxList, nil
}

func (d *Dao) DeleteInboxMessages(ctx context.Context, deleteUserId int64, peer *mtproto.PeerUtil, deleteMsgDataIds []int64, cb func(ctx context.Context, userId int64, idList []int32)) error {
	var (
		deletedDialogsMap = map[int64][]*dataobject.MessagesDO{}
	)

	doDeleteMessageF := func(v *dataobject.MessagesDO) {
		if v.UserId == deleteUserId {
			return
		}

		if v2, ok := deletedDialogsMap[v.UserId]; !ok {
			deletedDialogsMap[v.UserId] = []*dataobject.MessagesDO{v}
		} else {
			deletedDialogsMap[v.UserId] = append(v2, v)
		}
	}

	switch peer.PeerType {
	case mtproto.PEER_USER:
		_, err := d.MessagesDAO.SelectByMessageDataIdListWithCB(
			ctx,
			d.MessagesDAO.CalcTableName(peer.PeerId),
			deleteMsgDataIds,
			func(sz, i int, v *dataobject.MessagesDO) {
				doDeleteMessageF(v)
			})
		if err != nil {
			return err
		}
	case mtproto.PEER_CHAT:
		var (
			tables = sets.NewWithLength(1)
		)

		pUserIdList, _ := d.ChatClient.ChatGetChatParticipantIdList(ctx, &chatpb.TLChatGetChatParticipantIdList{
			ChatId: peer.PeerId,
		})

		for _, uId := range pUserIdList.GetDatas() {
			tables.Insert(d.MessagesDAO.CalcTableName(uId))
		}

		for tableName, _ := range tables {
			d.MessagesDAO.SelectByMessageDataIdListWithCB(
				ctx,
				tableName,
				deleteMsgDataIds,
				func(sz, i int, v *dataobject.MessagesDO) {
					doDeleteMessageF(v)
				})
		}
	}

	// TODO(@benqi): sort

	for userId, msgDOList := range deletedDialogsMap {
		var (
			// topMessage int32
			dialogId mtproto.DialogID
			msgIds   []int32
		)

		//if dlgDO == nil {
		//	dlgDO = &dataobject.DialogsDO{
		//		ReadInboxMaxId: math.MaxInt32,
		//		UnreadCount:    0,
		//		TopMessage:     0,
		//	}
		//	topMessage = dlgDO.TopMessage
		//}

		for i := 0; i < len(msgDOList); i++ {
			if dialogId.A == 0 && dialogId.B == 0 {
				dialogId.A = msgDOList[i].DialogId1
				dialogId.B = msgDOList[i].DialogId2
			}

			// check conversation peer_id
			if dialogId.A != msgDOList[i].DialogId1 && dialogId.B != msgDOList[i].DialogId2 {
				// dialogId
				err := mtproto.ErrMessageIdInvalid
				logx.WithContext(ctx).Errorf("deleteInboxMessages error: %v", err)
				// continue
				return err
			}
			msgIds = append(msgIds, msgDOList[i].UserMessageBoxId)
		}

		dlgDO, _ := d.DialogsDAO.SelectDialog(ctx, userId, msgDOList[0].PeerType, mtproto.GetPeerIdByDialogId(userId, dialogId))
		if dlgDO != nil {
			topMessage := dlgDO.TopMessage
			for i := 0; i < len(msgDOList); i++ {
				if msgDOList[i].UserMessageBoxId >= dlgDO.TopMessage {
					dlgDO.TopMessage -= 1
				}
				if msgDOList[i].UserMessageBoxId > dlgDO.ReadInboxMaxId {
					dlgDO.UnreadCount -= 1
				}
			}
			if !(topMessage == dlgDO.TopMessage ||
				dlgDO.TopMessage == msgDOList[len(msgDOList)-1].UserMessageBoxId) {

				dlgDO.TopMessage, _ = d.MessagesDAO.SelectDialogLastMessageIdNotIdList(ctx, userId, dialogId.A, dialogId.B, msgIds)
			}

			if dlgDO.UnreadCount < 0 {
				dlgDO.UnreadCount = 0
			}
		}

		tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			_, result.Err = d.MessagesDAO.DeleteMessagesByMessageIdListTx(tx, userId, msgIds)
			if result.Err != nil {
				return
			}
			if dlgDO != nil {
				_, result.Err = d.DialogsDAO.UpdateCustomMapTx(
					tx,
					map[string]interface{}{
						"top_message":  dlgDO.TopMessage,
						"unread_count": dlgDO.UnreadCount,
					},
					userId,
					dlgDO.PeerType,
					dlgDO.PeerId)
			}
		})
		if tR.Err != nil {
			return tR.Err
		}

		if cb != nil {
			cb(ctx, userId, msgIds)
		}
	}
	return nil
}

func (d *Dao) EditUserInboxMessage(ctx context.Context, fromId, peerId int64, message *mtproto.Message) (box *mtproto.MessageBox, err error) {
	var peerMsgDO *dataobject.MessagesDO

	peerMsgDO, err = d.MessagesDAO.SelectPeerUserMessage(ctx, peerId, fromId, message.Id)
	if err != nil {
		return
	} else if peerMsgDO == nil {
		return
	}

	// message.Id
	message.Out = false
	message.Id = peerMsgDO.UserMessageBoxId
	var (
		peerMessage *mtproto.Message
	)
	jsonx.UnmarshalFromString(peerMsgDO.MessageData, &peerMessage)
	// peerMessage, _ := mtproto.DecodeMessage(int(peerMsgDO.MessageType), []byte(peerMsgDO.MessageData))
	message.FromId = peerMessage.FromId
	message.PeerId = peerMessage.PeerId
	message.ReplyTo = peerMessage.ReplyTo
	mData, _ := jsonx.Marshal(message)
	if _, err = d.MessagesDAO.UpdateEditMessage(ctx, string(mData), message.Message, peerId, message.Id); err != nil {
		return
	}

	box = &mtproto.MessageBox{
		UserId:            peerId,
		SenderUserId:      0,
		PeerType:          mtproto.PEER_USER,
		PeerId:            peerId,
		MessageId:         message.Id,
		DialogId1:         0,
		DialogId2:         0,
		DialogMessageId:   0,
		RandomId:          0,
		Pts:               d.IDGenClient2.NextPtsId(ctx, peerId),
		PtsCount:          1,
		MessageFilterType: 0,
		Message:           message,
	}
	return
}

func (d *Dao) EditChatInboxMessage(ctx context.Context, fromId int64, peerChatId, toId int64, message *mtproto.Message) (box *mtproto.MessageBox, err error) {
	var peerMsgDO *dataobject.MessagesDO

	peerMsgDO, err = d.MessagesDAO.SelectPeerUserMessage(ctx, toId, fromId, message.Id)
	if err != nil {
		return
	} else if peerMsgDO == nil {
		return
	}

	// message.Id
	message.Out = false
	message.Id = peerMsgDO.UserMessageBoxId
	if message.GetReplyTo() != nil {
		var (
			peerMessage *mtproto.Message
		)
		// peerMessage, _ := mtproto.DecodeMessage(int(peerMsgDO.MessageType), []byte(peerMsgDO.MessageData))
		jsonx.UnmarshalFromString(peerMsgDO.MessageData, &peerMessage)
		message.ReplyTo = peerMessage.ReplyTo
	}

	mData, _ := jsonx.Marshal(message)
	if _, err = d.MessagesDAO.UpdateEditMessage(ctx, string(mData), message.Message, toId, message.Id); err != nil {
		return
	}

	box = &mtproto.MessageBox{
		UserId:            toId,
		SenderUserId:      0,
		PeerType:          mtproto.PEER_CHAT,
		PeerId:            peerChatId,
		MessageId:         message.Id,
		DialogId1:         0,
		DialogId2:         0,
		DialogMessageId:   0,
		RandomId:          0,
		Pts:               d.IDGenClient2.NextPtsId(ctx, toId),
		PtsCount:          1,
		MessageFilterType: 0,
		Message:           message,
	}
	return
}
