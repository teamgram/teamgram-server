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
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func makeMessageBoxByDO(boxDO *dataobject.MessagesDO) *mtproto.MessageBox {
	// TODO(@benqi): check boxDO and dataDO
	// message, _ := mtproto.DecodeMessage(int(boxDO.MessageType), hack.Bytes(boxDO.MessageData))

	box := &mtproto.MessageBox{
		UserId:            boxDO.UserId,
		SenderUserId:      boxDO.SenderUserId,
		PeerType:          boxDO.PeerType,
		PeerId:            boxDO.PeerId,
		MessageId:         boxDO.UserMessageBoxId,
		DialogId1:         boxDO.DialogId1,
		DialogId2:         boxDO.DialogId2,
		DialogMessageId:   boxDO.DialogMessageId,
		RandomId:          boxDO.RandomId,
		Pts:               0,
		PtsCount:          0,
		MessageFilterType: boxDO.MessageFilterType,
		Message:           nil,
	}
	jsonx.UnmarshalFromString(boxDO.MessageData, &box.Message)
	box.Message = box.Message.FixData()
	return box
}

func (d *Dao) sendMessageToOutbox(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, outboxMessage *msg.OutboxMessage) (*mtproto.MessageBox, bool, error) {
	var (
		dialogId = mtproto.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		err      error
		message  = outboxMessage.Message
	)

	idList := d.IDGenClient2.GetNextIdList(
		ctx,
		idgen_client.MakeIDTypeNextId(),
		idgen_client.MakeIDTypeNgen(idgen_client.IDTypeMessageBox, fromId),
		idgen_client.MakeIDTypeNgen(idgen_client.IDTypePts, fromId))
	if len(idList) != 3 {
		err = mtproto.ErrInternalServerError
		return nil, false, err
	}

	dialogMessageId := idList[0].Id
	outBoxMsgId := int32(idList[1].Id)
	pts := int32(idList[2].Id)

	if dialogMessageId == 0 || outBoxMsgId == 0 || pts == 0 {
		err = mtproto.ErrInternalServerError
		return nil, false, err
	}

	// A, B
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		message.Out = true
		message.Id = outBoxMsgId
		// message.Mentioned = model.CheckHasMention(message.Entities, fromId)
		//if message.Mentioned {
		//	message.MediaUnread = true
		//}
		//if !message.MediaUnread {
		message.MediaUnread = mtproto.CheckHasMediaUnread(message)
		//}
		// Pts:              pts,
		// PtsCount:         ptsCount,
		// mType, mData := mtproto.EncodeMessage(message)
		mData, _ := jsonx.Marshal(message)
		outMsgBox := &mtproto.MessageBox{
			UserId:            fromId,
			SenderUserId:      fromId,
			PeerType:          peer.PeerType,
			PeerId:            peer.PeerId,
			MessageId:         outBoxMsgId,
			DialogId1:         dialogId.A,
			DialogId2:         dialogId.B,
			DialogMessageId:   dialogMessageId,
			RandomId:          outboxMessage.RandomId,
			Pts:               0,
			PtsCount:          0,
			MessageFilterType: mtproto.GetMediaType(message),
			Message:           message,
		}

		var (
			savedPeerUtil *mtproto.PeerUtil
		)
		if message.GetSavedPeerId() != nil {
			savedPeerUtil = mtproto.FromPeer(message.GetSavedPeerId())
		} else {
			savedPeerUtil = &mtproto.PeerUtil{PeerType: mtproto.PEER_EMPTY, PeerId: 0}
		}

		outBoxDO := &dataobject.MessagesDO{
			UserId:            outMsgBox.UserId,
			UserMessageBoxId:  outMsgBox.MessageId,
			DialogId1:         outMsgBox.DialogId1,
			DialogId2:         outMsgBox.DialogId2,
			DialogMessageId:   outMsgBox.DialogMessageId,
			SenderUserId:      outMsgBox.UserId,
			PeerType:          peer.PeerType,
			PeerId:            peer.PeerId,
			RandomId:          outMsgBox.RandomId,
			MessageFilterType: outMsgBox.MessageFilterType,
			MessageData:       hack.String(mData),
			Message:           message.Message,
			Mentioned:         false,
			MediaUnread:       message.MediaUnread,
			Date2:             int64(outMsgBox.Message.Date),
			SavedPeerType:     savedPeerUtil.PeerType,
			SavedPeerId:       savedPeerUtil.PeerId,
			Deleted:           false,
		}

		lastInsertId, rowsAffected, err := d.MessagesDAO.InsertOrReturnIdTx(tx, outBoxDO)
		if err != nil {
			result.Err = err
			return
		}

		if rowsAffected == 0 {
			// TODO(@benqi): random_id已经存在
			if lastInsertId > 0 {
				result.Data = lastInsertId
				return
			} else {
				result.Err = errors.New("insert error")
				return
			}
		}

		//outMsgBox.Pts = d.IDGenClient2.NextPtsId(ctx, fromId)
		//outMsgBox.PtsCount = 1
		// logx.WithContext(ctx).Infof("sendMessage - (pts: %d, pts_count: %d)", outMsgBox.Pts, outMsgBox.PtsCount)

		switch peer.PeerType {
		case mtproto.PEER_USER:
			dialogDO := &dataobject.DialogsDO{
				UserId:           fromId,
				PeerType:         peer.PeerType,
				PeerId:           peer.PeerId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_USER, peer.PeerId),
				TopMessage:       outBoxMsgId,
				UnreadCount:      0,
				DraftMessageData: "null",
				Date2:            int64(outMsgBox.Message.Date),
			}
			if dialogMessageId > 1 {
				//// if box_id > 1, then dialogs already created.
				//
				//// TODO(@benqi): unread_count and unread_mentions_count
				//cMap := map[string]interface{}{
				//	"top_message": dialogDO.TopMessage,
				//	"date2":       dialogDO.Date2,
				//	"unread_mark": 0,
				//}
				//
				//// TODO(@benqi): clear draft
				//if true {
				//	cMap["draft_message_data"] = "null"
				//	cMap["draft_type"] = 0
				//}

				rowsAffected, result.Err = d.DialogsDAO.UpdateOutboxDialogTx(
					tx,
					dialogDO.TopMessage,
					dialogDO.Date2,
					fromId,
					mtproto.PEER_USER,
					peer.PeerId)
				// log.Infof("rowsAffected = %d, %v", rowsAffected, dialogDO)
				if result.Err != nil {
					return
				}
				// again handle rowsAffected == 0
				if rowsAffected == 0 {
					_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
				}
			} else {
				_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
			}

			result.Data = outMsgBox
		case mtproto.PEER_CHAT:
			dialogDO := &dataobject.DialogsDO{
				UserId:           fromId,
				PeerType:         peer.PeerType,
				PeerId:           peer.PeerId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_CHAT, peer.PeerId),
				TopMessage:       outBoxMsgId,
				UnreadCount:      0,
				DraftMessageData: "null",
				Date2:            int64(outMsgBox.Message.Date),
			}

			if dialogMessageId > 1 {
				rowsAffected, result.Err = d.DialogsDAO.UpdateOutboxDialogTx(tx,
					dialogDO.TopMessage,
					dialogDO.Date2,
					fromId,
					mtproto.PEER_CHAT,
					peer.PeerId)
				// log.Infof("rowsAffected = %d, %v", rowsAffected, dialogDO)
				if result.Err != nil {
					return
				}
				// again handle rowsAffected == 0
				if rowsAffected == 0 {
					_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
				}
			} else {
				_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
			}

			result.Data = outMsgBox
		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
			logx.WithContext(ctx).Errorf("%v", result)
			return
		}

		for _, entity := range message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					d.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
						UserId:           outMsgBox.UserId,
						PeerType:         peer.PeerType,
						PeerId:           peer.PeerId,
						HashTag:          entity.GetUrl(),
						HashTagMessageId: outMsgBox.MessageId,
					})
				}
			}
		}
	})

	if tR.Err != nil {
		return nil, false, tR.Err
	}

	var outBox *mtproto.MessageBox

	switch tR.Data.(type) {
	case *mtproto.MessageBox:
		outBox = tR.Data.(*mtproto.MessageBox)
		outBox.Pts = pts // d.IDGenClient2.NextPtsId(ctx, fromId)
		outBox.PtsCount = 1

	case int64:
		if tR.Data.(int64) <= 0 {
			logx.WithContext(ctx).Error("unknown error")
			return nil, false, errors.New("fatal unknown error")
		}

		// dup, we'll recreate box
		do, err := d.MessagesDAO.SelectByRandomId(ctx, fromId, outboxMessage.RandomId)
		if err != nil {
			return nil, false, err
		}
		if do != nil {
			outBox = makeMessageBoxByDO(do)
			outBox.Pts = d.IDGenClient2.CurrentPtsId(ctx, outBox.UserId)
			outBox.PtsCount = 1
			return outBox, false, nil
		} else {
			logx.WithContext(ctx).Error("unknown error")
			return nil, false, errors.New("fatal unknown error")
		}
	default:
		logx.WithContext(ctx).Error("unknown error")
		return nil, false, errors.New("fatal unknown error")
	}

	return outBox, true, nil
}

func (d *Dao) SendUserMessage(ctx context.Context, fromId, toId int64, outBox *msg.OutboxMessage) (*mtproto.MessageBox, bool, error) {
	peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_USER, PeerId: toId}
	return d.sendMessageToOutbox(ctx, fromId, peer, outBox)
}

func (d *Dao) SendUserMessageV2(ctx context.Context, fromId, toId int64, outBox *msg.OutboxMessage, out bool) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_USER, PeerId: toId}
	return d.sendMessageToOutboxV2(ctx, fromId, peer, outBox, out)
}

func (d *Dao) SendUserMultiMessage(ctx context.Context, fromId, toId int64, outBoxList []*msg.OutboxMessage) ([]*mtproto.MessageBox, error) {
	var (
		boxList []*mtproto.MessageBox
	)

	for _, msg := range outBoxList {
		peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_USER, PeerId: toId}
		outBox, _, _ := d.sendMessageToOutbox(ctx, fromId, peer, msg)
		boxList = append(boxList, outBox)
	}

	return boxList, nil
}

func (d *Dao) SendChatMessage(ctx context.Context, fromId, chatId int64, outBox *msg.OutboxMessage) (*mtproto.MessageBox, bool, error) {
	peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_CHAT, PeerId: chatId}
	return d.sendMessageToOutbox(ctx, fromId, peer, outBox)
}

func (d *Dao) SendChatMessageV2(ctx context.Context, fromId, chatId int64, outBox *msg.OutboxMessage) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_CHAT, PeerId: chatId}
	return d.sendMessageToOutboxV2(ctx, fromId, peer, outBox, true)
}

func (d *Dao) SendChatMultiMessage(ctx context.Context, fromId, chatId int64, outBoxList []*msg.OutboxMessage) ([]*mtproto.MessageBox, error) {
	var (
		boxList []*mtproto.MessageBox
	)

	for _, msg := range outBoxList {
		peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_CHAT, PeerId: chatId}
		box, _, _ := d.sendMessageToOutbox(ctx, fromId, peer, msg)
		boxList = append(boxList, box)
	}

	return boxList, nil
}

func (d *Dao) DeleteMessages(ctx context.Context, userId int64, msgIds []int32) (*mtproto.PeerUtil, []int64, error) {

	var (
		topMessageIndex      int32
		dialogId             *mtproto.DialogID
		deletedMsgDataIdList = make([]int64, 0, len(msgIds))
	)

	msgDOList, err := d.MessagesDAO.SelectByMessageIdListWithCB(
		ctx,
		userId,
		msgIds,
		func(sz, i int, v *dataobject.MessagesDO) {
			if dialogId == nil {
				dialogId = &mtproto.DialogID{
					A: v.DialogId1,
					B: v.DialogId2,
				}
			}

			if dialogId.A != v.DialogId1 || dialogId.B != v.DialogId2 {
				dialogId.A = 0
				dialogId.B = 0
			}
		})
	if err != nil {
		// mtproto.ErrMsgIdInvalid
		return nil, nil, err
	} else if dialogId.IsZero() {
		return mtproto.MakePeerUtil(mtproto.PEER_EMPTY, 0), []int64{}, nil
	}

	// 会话里最后n条消息，检查是否需要修改会话信息
	topMessageDOList, err := d.MessagesDAO.SelectDialogLastMessageList(ctx, userId, dialogId.A, dialogId.B, int32(len(msgIds)+1))
	if err != nil {
		return nil, nil, err
	} else if len(topMessageDOList) == 0 {
		// return []int64{}, nil
	} else {
		topMessageIndex = math.MaxInt32
	}

	getLastTopMessage := func(topMessage2 int32) (int32, int64) {
		for i := 0; i < len(topMessageDOList); i++ {
			if topMessageDOList[i].UserMessageBoxId >= topMessage2 {
				continue
			}
			return topMessageDOList[i].UserMessageBoxId, topMessageDOList[i].Date2
		}
		return 0, 0
	}

	//_, err = d.DialogsDAO.SelectDialogsWithCB().SelectByPeer(ctx, userId, model.GetPeerIdByDialogId(userId, dialogId))
	//if err != nil {
	//	return nil, err
	//}

	for i := 0; i < len(msgDOList); i++ {
		topMessage, _ := getLastTopMessage(topMessageIndex)
		if topMessage == msgDOList[i].UserMessageBoxId {
			topMessageIndex = topMessage
		}
	}

	peer := dialogId.ToPeerUtil(userId)

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, result.Err = d.MessagesDAO.DeleteMessagesByMessageIdListTx(tx, userId, msgIds)
		if result.Err != nil {
			return
		}
		topMessage, _ := getLastTopMessage(topMessageIndex)
		d.DialogsDAO.UpdateOutboxDialogTx(
			tx,
			topMessage,
			time.Now().Unix(),
			userId,
			peer.PeerType,
			peer.PeerId)
	})
	if tR.Err != nil {
		return nil, nil, tR.Err
	}

	for i := 0; i < len(msgDOList); i++ {
		deletedMsgDataIdList = append(deletedMsgDataIdList, msgDOList[i].DialogMessageId)
	}
	return peer, deletedMsgDataIdList, nil
}

//func (d *Dao) editOutboxMessage(ctx context.Context, fromId int32, peer *model.PeerUtil, toId int32, message *mtproto.Message) (box *model.MessageBox, err error) {
//	mType, mData := model.EncodeMessage(message)
//	if _, err = m.MessagesDAO.UpdateEditMessage(ctx, int8(mType), hack.String(mData), message.Message, fromId, message.Id); err != nil {
//		return
//	}
//	box = &model.MessageBox{
//		SelfUserId:        fromId,
//		SendUserId:        fromId,
//		PeerType:          peer.PeerType,
//		PeerId:            peer.PeerId,
//		MessageId:         message.Id,
//		DialogId:          model.MakeDialogId(fromId, peer.PeerType, peer.PeerId),
//		DialogMessageId:   0,
//		RandomId:          0,
//		Pts:               int32(idgen.NextPtsId(ctx, fromId)),
//		PtsCount:          1,
//		MessageFilterType: 0,
//		MessageBoxType:    0,
//		MessageType:       int8(mType),
//		Message:           message,
//	}
//	return
//}

func (d *Dao) EditUserOutboxMessage(ctx context.Context, fromId, toId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	return d.editOutboxMessage(ctx, fromId, mtproto.PEER_USER, toId, message)
}

func (d *Dao) EditChatOutboxMessage(ctx context.Context, fromId, toId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	return d.editOutboxMessage(ctx, fromId, mtproto.PEER_CHAT, toId, message)
}

func (d *Dao) editOutboxMessage(ctx context.Context, fromId int64, peerType int32, peerId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	var (
		mData, _ = jsonx.Marshal(message)
		did      = mtproto.MakeDialogId(fromId, peerType, peerId)
	)

	if _, err := d.MessagesDAO.UpdateEditMessage(ctx, string(mData), message.Message, fromId, message.Id); err != nil {
		return nil, err
	}

	d.HashTagsDAO.DeleteHashTagMessageId(ctx, fromId, message.Id)
	for _, entity := range message.GetEntities() {
		if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
			if entity.GetUrl() != "" {
				d.HashTagsDAO.InsertOrUpdate(ctx, &dataobject.HashTagsDO{
					UserId:           fromId,
					PeerType:         peerType,
					PeerId:           peerId,
					HashTag:          entity.GetUrl(),
					HashTagMessageId: message.Id,
				})
			}
		}
	}

	return &mtproto.MessageBox{
		UserId:            fromId,
		MessageId:         message.Id,
		SenderUserId:      fromId,
		PeerType:          peerType,
		PeerId:            peerId,
		RandomId:          0,
		DialogId1:         did.A,
		DialogId2:         did.B,
		DialogMessageId:   0,
		MessageFilterType: 0,
		Message:           message,
		Pts:               d.IDGenClient2.NextPtsId(ctx, fromId),
		PtsCount:          1,
		Mentioned:         false,
		MediaUnread:       false,
	}, nil
}

func (d *Dao) DeletePhoneCallHistory(ctx context.Context, userId int64) ([]int32, []int64, error) {
	var (
		dialogIdMap          = make(map[mtproto.DialogID][]int32)
		deletedIdList        = make([]int32, 0)
		deletedMsgDataIdList = make([]int64, 0)
	)

	_, err := d.MessagesDAO.SelectPhoneCallListWithCB(
		ctx,
		userId,
		mtproto.MEDIA_PHONE_CALL,
		math.MaxInt32,
		math.MaxInt32,
		func(sz, i int, v *dataobject.MessagesDO) {
			did := mtproto.DialogID{A: v.DialogId1, B: v.DialogId2}
			if v2, ok := dialogIdMap[did]; !ok {
				dialogIdMap[did] = []int32{v.UserMessageBoxId}
			} else {
				dialogIdMap[did] = append(v2, v.UserMessageBoxId)
			}
			deletedIdList = append(deletedIdList, v.UserMessageBoxId)
			deletedMsgDataIdList = append(deletedMsgDataIdList, v.DialogMessageId)
		},
	)
	if err != nil {
		logx.WithContext(ctx).Errorf("error: %v", err)
	}

	if len(dialogIdMap) == 0 {
		return []int32{}, []int64{}, nil
	}

	logx.WithContext(ctx).Infof("phone: %v", dialogIdMap)

	for dialogId, msgIds := range dialogIdMap {
		_, err = d.MessagesDAO.DeleteMessagesByMessageIdList(ctx, userId, msgIds)
		if err != nil {
			continue
		}
		lastMessageId, _ := d.MessagesDAO.SelectDialogLastMessageId(ctx, userId, dialogId.A, dialogId.B)
		d.DialogsDAO.UpdateCustomMap(ctx, map[string]interface{}{
			"top_message": lastMessageId,
		}, userId, mtproto.PEER_USER, mtproto.GetPeerIdByDialogId(userId, dialogId))
	}

	return deletedIdList, deletedMsgDataIdList, nil
}

func (d *Dao) SendMessageToOutboxV1(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, outMsgBox *mtproto.MessageBox) error {
	message := outMsgBox.Message
	mData, _ := jsonx.Marshal(outMsgBox.GetMessage())
	outBoxMsgId := outMsgBox.MessageId
	dialogMessageId := outMsgBox.DialogMessageId

	var (
		savedPeerUtil *mtproto.PeerUtil
	)
	if message.GetSavedPeerId() != nil {
		savedPeerUtil = mtproto.FromPeer(message.GetSavedPeerId())
	} else {
		savedPeerUtil = &mtproto.PeerUtil{PeerType: mtproto.PEER_EMPTY, PeerId: 0}
	}

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		outBoxDO := &dataobject.MessagesDO{
			UserId:            outMsgBox.UserId,
			UserMessageBoxId:  outMsgBox.MessageId,
			DialogId1:         outMsgBox.DialogId1,
			DialogId2:         outMsgBox.DialogId2,
			DialogMessageId:   outMsgBox.DialogMessageId,
			SenderUserId:      outMsgBox.UserId,
			PeerType:          peer.PeerType,
			PeerId:            peer.PeerId,
			RandomId:          outMsgBox.RandomId,
			MessageFilterType: outMsgBox.MessageFilterType,
			MessageData:       hack.String(mData),
			Message:           message.GetMessage(),
			Mentioned:         false,
			MediaUnread:       message.GetMediaUnread(),
			Date2:             int64(outMsgBox.Message.Date),
			SavedPeerType:     savedPeerUtil.PeerType,
			SavedPeerId:       savedPeerUtil.PeerId,
			Deleted:           false,
		}

		lastInsertId, rowsAffected, err := d.MessagesDAO.InsertOrReturnIdTx(tx, outBoxDO)
		if err != nil {
			result.Err = err
			return
		}

		if rowsAffected == 0 {
			// TODO(@benqi): random_id已经存在
			if lastInsertId > 0 {
				result.Data = lastInsertId
				return
			} else {
				result.Err = errors.New("insert error")
				return
			}
		}

		switch peer.PeerType {
		case mtproto.PEER_USER:
			dialogDO := &dataobject.DialogsDO{
				UserId:           fromId,
				PeerType:         peer.PeerType,
				PeerId:           peer.PeerId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_USER, peer.PeerId),
				TopMessage:       outBoxMsgId,
				UnreadCount:      0,
				DraftMessageData: "null",
				Date2:            int64(outMsgBox.Message.Date),
			}
			if dialogMessageId > 1 {
				rowsAffected, result.Err = d.DialogsDAO.UpdateOutboxDialogTx(
					tx,
					dialogDO.TopMessage,
					dialogDO.Date2,
					fromId,
					mtproto.PEER_USER,
					peer.PeerId)
				// log.Infof("rowsAffected = %d, %v", rowsAffected, dialogDO)
				if result.Err != nil {
					return
				}
				// again handle rowsAffected == 0
				if rowsAffected == 0 {
					_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
				}
			} else {
				_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
			}
		case mtproto.PEER_CHAT:
			dialogDO := &dataobject.DialogsDO{
				UserId:           fromId,
				PeerType:         peer.PeerType,
				PeerId:           peer.PeerId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_CHAT, peer.PeerId),
				TopMessage:       outBoxMsgId,
				UnreadCount:      0,
				DraftMessageData: "null",
				Date2:            int64(outMsgBox.Message.Date),
			}

			if dialogMessageId > 1 {
				rowsAffected, result.Err = d.DialogsDAO.UpdateOutboxDialogTx(tx,
					dialogDO.TopMessage,
					dialogDO.Date2,
					fromId,
					mtproto.PEER_CHAT,
					peer.PeerId)
				// log.Infof("rowsAffected = %d, %v", rowsAffected, dialogDO)
				if result.Err != nil {
					return
				}
				// again handle rowsAffected == 0
				if rowsAffected == 0 {
					_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
				}
			} else {
				_, _, result.Err = d.DialogsDAO.InsertIgnoreTx(tx, dialogDO)
			}

			result.Data = outMsgBox
		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
			logx.WithContext(ctx).Errorf("%v", result)
			return
		}

		for _, entity := range message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					d.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
						UserId:           outMsgBox.UserId,
						PeerType:         peer.PeerType,
						PeerId:           peer.PeerId,
						HashTag:          entity.GetUrl(),
						HashTagMessageId: outMsgBox.MessageId,
					})
				}
			}
		}
	})

	return tR.Err
}

func (d *Dao) sendMessageToOutboxV2(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, outboxMessage *msg.OutboxMessage, out bool) (*mtproto.MessageBox, error) {
	var (
		dialogId        = mtproto.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		err             error
		message         = outboxMessage.Message
		outBoxMsgId     int32
		dialogMessageId int64
		pts             int32
	)

	if out {
		idList := d.IDGenClient2.GetNextIdList(
			ctx,
			idgen_client.MakeIDTypeNextId(),
			idgen_client.MakeIDTypeNgen(idgen_client.IDTypeMessageBox, fromId),
			idgen_client.MakeIDTypeNgen(idgen_client.IDTypePts, fromId))
		if len(idList) != 3 {
			err = mtproto.ErrInternalServerError
			return nil, err
		}

		dialogMessageId = idList[0].Id
		outBoxMsgId = int32(idList[1].Id)
		pts = int32(idList[2].Id)

		if dialogMessageId == 0 || outBoxMsgId == 0 || pts == 0 {
			logx.WithContext(ctx).Errorf("GetNextIdList error: %v", idList)
			err = mtproto.ErrInternalServerError
			return nil, err
		}
	} else {
		dialogMessageId = d.IDGenClient2.NextId(ctx)
		if dialogMessageId == 0 {
			err = mtproto.ErrInternalServerError
			logx.WithContext(ctx).Errorf("NextId error: %v", dialogMessageId)
			return nil, err

		}
	}

	message.Out = out
	message.Id = outBoxMsgId
	message.MediaUnread = mtproto.CheckHasMediaUnread(message)
	outMsgBox := mtproto.MakeTLMessageBox(&mtproto.MessageBox{
		UserId:            fromId,
		MessageId:         outBoxMsgId,
		SenderUserId:      fromId,
		PeerType:          peer.PeerType,
		PeerId:            peer.PeerId,
		RandomId:          outboxMessage.RandomId,
		DialogId1:         dialogId.A,
		DialogId2:         dialogId.B,
		DialogMessageId:   dialogMessageId,
		MessageFilterType: mtproto.GetMediaType(message),
		Message:           message,
		Mentioned:         false,
		MediaUnread:       false,
		Pinned:            false,
		Pts:               pts,
		PtsCount:          1,
		Views:             0,
		ReplyOwnerId:      0,
		Forwards:          0,
		Reaction:          "",
		CommentGroupId:    0,
		CommentGroupMsgId: 0,
		ReplyToMsgId:      0,
		ReplyToTopId:      0,
		TtlPeriod:         0,
		HasReaction:       false,
	}).To_MessageBox()

	return outMsgBox, nil
}

func (d *Dao) EditUserOutboxMessageV2(ctx context.Context, fromId, toId int64, newMessage *msg.OutboxMessage, dstMessage *mtproto.MessageBox) (*mtproto.MessageBox, error) {
	return d.editOutboxMessageV2(ctx, fromId, mtproto.PEER_USER, toId, newMessage, dstMessage)
}

func (d *Dao) EditChatOutboxMessageV2(ctx context.Context, fromId, toId int64, newMessage *msg.OutboxMessage, dstMessage *mtproto.MessageBox) (*mtproto.MessageBox, error) {
	return d.editOutboxMessageV2(ctx, fromId, mtproto.PEER_CHAT, toId, newMessage, dstMessage)
}

func (d *Dao) editOutboxMessageV2(ctx context.Context, fromId int64, peerType int32, peerId int64, newMessage *msg.OutboxMessage, dstMessage *mtproto.MessageBox) (*mtproto.MessageBox, error) {
	var (
		pts            = d.IDGenClient2.NextPtsId(ctx, fromId)
		ptsCount int32 = 1
		message        = newMessage.Message
	)

	if pts == 0 {
		logx.WithContext(ctx).Errorf("NextPtsId error: %v", fromId)
		err := mtproto.ErrInternalServerError
		return nil, err
	}
	//if _, err := d.MessagesDAO.UpdateEditMessage(ctx, string(mData), message.Message, fromId, message.Id); err != nil {
	//	return nil, err
	//}

	//// d.HashTagsDAO.DeleteHashTagMessageId(ctx, fromId, message.Id)
	//for _, entity := range message.GetEntities() {
	//	if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
	//		if entity.GetUrl() != "" {
	//			d.HashTagsDAO.InsertOrUpdate(ctx, &dataobject.HashTagsDO{
	//				UserId:           fromId,
	//				PeerType:         peerType,
	//				PeerId:           peerId,
	//				HashTag:          entity.GetUrl(),
	//				HashTagMessageId: dstMessage.MessageId,
	//			})
	//		}
	//	}
	//}

	return mtproto.MakeTLMessageBox(&mtproto.MessageBox{
		UserId:            dstMessage.UserId,
		MessageId:         dstMessage.MessageId,
		SenderUserId:      dstMessage.SenderUserId,
		PeerType:          dstMessage.PeerType,
		PeerId:            dstMessage.PeerId,
		RandomId:          dstMessage.RandomId,
		DialogId1:         dstMessage.DialogId1,
		DialogId2:         dstMessage.DialogId2,
		DialogMessageId:   dstMessage.DialogMessageId,
		MessageFilterType: dstMessage.MessageFilterType,
		Message:           message,
		Mentioned:         dstMessage.Mentioned,
		MediaUnread:       dstMessage.MediaUnread,
		Pinned:            dstMessage.Pinned,
		Pts:               pts,
		PtsCount:          ptsCount,
		Views:             dstMessage.Views,
		ReplyOwnerId:      dstMessage.ReplyOwnerId,
		Forwards:          dstMessage.Forwards,
		Reaction:          dstMessage.Reaction,
		CommentGroupId:    dstMessage.CommentGroupId,
		CommentGroupMsgId: dstMessage.CommentGroupMsgId,
		ReplyToMsgId:      dstMessage.ReplyToMsgId,
		ReplyToTopId:      dstMessage.ReplyToTopId,
		TtlPeriod:         dstMessage.TtlPeriod,
		HasReaction:       dstMessage.HasReaction,
	}).To_MessageBox(), nil
}
