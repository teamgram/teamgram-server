// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msgtransfer/msgtransfer"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (d *Dao) SendUserMessageV4(ctx context.Context, fromId, toId int64, outBox *msgtransfer.OutboxMessage, out bool) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{PeerType: mtproto.PEER_USER, PeerId: toId}
	return d.sendMessageToOutboxV4(ctx, fromId, peer, outBox, out)
}

func (d *Dao) sendMessageToOutboxV4(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, outboxMessage *msgtransfer.OutboxMessage, out bool) (*mtproto.MessageBox, error) {
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

func (d *Dao) SendMessageToOutboxV1(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, outMsgBox *mtproto.MessageBox) error {
	message := outMsgBox.Message
	mData, _ := jsonx.Marshal(outMsgBox.GetMessage())
	outBoxMsgId := outMsgBox.MessageId

	var (
		savedPeerUtil *mtproto.PeerUtil
	)

	if message.GetSavedPeerId() != nil {
		savedPeerUtil = mtproto.FromPeer(message.GetSavedPeerId())
	} else {
		savedPeerUtil = &mtproto.PeerUtil{PeerType: mtproto.PEER_EMPTY, PeerId: 0}
	}

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, _, err := d.MessagesDAO.InsertOrReturnIdTx(
			tx,
			&dataobject.MessagesDO{
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
			})
		if err != nil {
			result.Err = err
			return
		}

		for _, entity := range message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					_, _, err = d.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
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

	_, err := d.DialogClient.DialogInsertOrUpdateDialog(
		ctx,
		&dialog.TLDialogInsertOrUpdateDialog{
			UserId:          fromId,
			PeerType:        peer.PeerType,
			PeerId:          peer.PeerId,
			TopMessage:      &wrapperspb.Int32Value{Value: outBoxMsgId},
			ReadOutboxMaxId: nil,
			ReadInboxMaxId:  nil,
			UnreadCount:     &wrapperspb.Int32Value{Value: 0},
			UnreadMark:      false,
			PinnedMsgId:     nil,
			Date2:           &wrapperspb.Int64Value{Value: int64(outMsgBox.Message.Date)},
		})
	if err != nil {
		// return i
	}

	return tR.Err
}
