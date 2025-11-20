// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

func (d *Dao) SendUserMessageToInbox(ctx context.Context, fromId, toId int64, dialogMessageId, clientRandomId int64, message *mtproto.Message) (*mtproto.MessageBox, error) {
	peer := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_USER,
		PeerId:   toId,
	}
	message.Out = false
	return d.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageId, clientRandomId, message)
}

func (d *Dao) sendMessageToInbox(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, toUserId int64, dialogMessageId, clientRandomId int64, message2 *mtproto.Message) (*mtproto.MessageBox, error) {
	var (
		inBoxMsgId = d.IDGenClient2.NextMessageBoxId(ctx, toUserId)
		dialogId   = mtproto.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		date       = time.Now().Unix()
		message    = proto.Clone(message2).(*mtproto.Message)

		dialogDO *dataobject.DialogsDO
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
			//var (
			//	lastInsertId int64
			//	rowsAffected int64
			//)

			dialogDO = &dataobject.DialogsDO{
				UserId:           inBox.UserId,
				PeerType:         peer.PeerType,
				PeerId:           fromId,
				PeerDialogId:     mtproto.MakePeerDialogId(mtproto.PEER_USER, fromId),
				TopMessage:       inBoxMsgId,
				UnreadCount:      1,
				DraftMessageData: "null",
				Date2:            date,
			}

		case mtproto.PEER_CHAT:
			dialogDO = &dataobject.DialogsDO{
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

		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
		}

		for _, entity := range message.GetEntities() {
			if entity.GetPredicateName() == mtproto.Predicate_messageEntityHashtag {
				if entity.GetUrl() != "" {
					_, _, _ = d.HashTagsDAO.InsertOrUpdateTx(tx, &dataobject.HashTagsDO{
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

	_, _, _ = d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			lastInsertId, rowsAffected, err := d.DialogsDAO.InsertOrUpdate(ctx, dialogDO)
			logx.WithContext(ctx).Infof("lastInsertId:%d, rowsAffected: %d, result: %v, do: %v", lastInsertId, rowsAffected, err, dialogDO)
			return 0, 0, err
		},
		dialog.GetDialogCacheKey(dialogDO.UserId, dialogDO.PeerDialogId))

	inBox.Pts = d.IDGenClient2.NextPtsId(ctx, toUserId)
	inBox.PtsCount = 1

	return inBox, nil
}
