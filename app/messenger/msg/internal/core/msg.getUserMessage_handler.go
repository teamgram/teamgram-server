package core

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *MsgCore) MsgGetUserMessage(in *msg.TLMsgGetUserMessage) (*tg.MessageBox, error) {
	if in == nil {
		return nil, msg.ErrMsgIdInvalid
	}
	box, err := c.svcCtx.Repo.GetUserMessage(c.ctx, in.UserId, int64(in.Id))
	if err != nil {
		return nil, err
	}
	return messageBoxFromUserMessage(box)
}

func messageBoxFromUserMessage(box *repository.UserMessageBox) (*tg.MessageBox, error) {
	if box == nil {
		return nil, msg.ErrMsgIdInvalid
	}
	messageID, err := historyIDInt32(box.UserMessageID, "user message id")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(box.MessageDate, "user message date")
	if err != nil {
		return nil, err
	}
	message, err := userMessageBoxTLMessage(box, messageID, date)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageBox(&tg.TLMessageBox{
		UserId:          box.UserID,
		MessageId:       messageID,
		SenderUserId:    box.FromUserID,
		PeerType:        box.PeerType,
		PeerId:          box.PeerID,
		DialogId1:       box.UserID,
		DialogId2:       box.PeerID,
		DialogMessageId: box.UserMessageID,
		Message:         message,
		Reaction:        "",
	}).ToMessageBox(), nil
}

func userMessageBoxTLMessage(box *repository.UserMessageBox, messageID int32, date int32) (tg.MessageClazz, error) {
	if len(box.ViewPayload) == 0 {
		return tg.MakeTLMessage(&tg.TLMessage{
			Out:     box.Outgoing,
			Id:      messageID,
			FromId:  tg.MakePeerUser(box.FromUserID),
			PeerId:  sentMessagePeerFromOptional(box.PeerType, box.PeerID),
			Date:    date,
			Message: box.MessageText,
		}), nil
	}

	var envelope struct {
		SchemaVersion int `json:"schema_version"`
	}
	if err := json.Unmarshal(box.ViewPayload, &envelope); err != nil {
		return nil, fmt.Errorf("%w: decode user message view payload: %v", msg.ErrMsgStorage, err)
	}
	switch envelope.SchemaVersion {
	case payload.MessageEventSchemaVersionV1:
		var event payload.MessageEventV1
		if err := json.Unmarshal(box.ViewPayload, &event); err != nil {
			return nil, fmt.Errorf("%w: decode user message view payload v1: %v", msg.ErrMsgStorage, err)
		}
		return userMessageEventV1ToTLMessage(box, event, messageID)
	case payload.MessageEventSchemaVersion:
		var event payload.MessageEventV2
		if err := json.Unmarshal(box.ViewPayload, &event); err != nil {
			return nil, fmt.Errorf("%w: decode user message view payload v2: %v", msg.ErrMsgStorage, err)
		}
		return userMessageEventV2ToTLMessage(box, event)
	case payload.MessageEventSchemaVersionV3:
		var event payload.MessageEventV3
		if err := json.Unmarshal(box.ViewPayload, &event); err != nil {
			return nil, fmt.Errorf("%w: decode user message view payload v3: %v", msg.ErrMsgStorage, err)
		}
		return userMessageEventV3ToTLMessage(box, event)
	default:
		return nil, fmt.Errorf("%w: unsupported user message view schema=%d", msg.ErrMsgStorage, envelope.SchemaVersion)
	}
}

func userMessageEventV1ToTLMessage(box *repository.UserMessageBox, event payload.MessageEventV1, messageID int32) (tg.MessageClazz, error) {
	if event.SchemaVersion != payload.MessageEventSchemaVersionV1 ||
		event.PeerType != box.PeerType ||
		event.PeerID != box.PeerID ||
		event.CanonicalMessageID != box.CanonicalMessageID {
		return nil, fmt.Errorf("%w: user message view payload v1 mismatch", msg.ErrMsgStorage)
	}
	date, err := msgDateInt32FromUnixSeconds(int64(event.Date), "user message view date")
	if err != nil {
		return nil, err
	}
	editDatePtr, err := userMessageEditDate(event.EventKind, event.Date, event.EditDate)
	if err != nil {
		return nil, err
	}
	replyTo, err := userMessageReplyHeaderFromPeerSeq(event.ReplyToPeerSeq)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      event.Out,
		Id:       messageID,
		FromId:   tg.MakePeerUser(event.FromUserID),
		PeerId:   sentMessagePeerFromOptional(event.PeerType, event.PeerID),
		ReplyTo:  replyTo,
		Date:     date,
		Message:  event.MessageText,
		Entities: sentMessageEntities(event.Entities),
		EditDate: editDatePtr,
	}), nil
}

func userMessageEventV2ToTLMessage(box *repository.UserMessageBox, event payload.MessageEventV2) (tg.MessageClazz, error) {
	if event.SchemaVersion != payload.MessageEventSchemaVersion ||
		event.PeerType != box.PeerType ||
		event.PeerID != box.PeerID ||
		event.PeerSeq != box.PeerSeq ||
		event.MessageID != box.UserMessageID ||
		event.CanonicalMessageID != box.CanonicalMessageID {
		return nil, fmt.Errorf("%w: user message view payload v2 mismatch", msg.ErrMsgStorage)
	}
	messageID, err := historyIDInt32(event.MessageID, "user message view id")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(int64(event.Date), "user message view date")
	if err != nil {
		return nil, err
	}
	editDatePtr, err := userMessageEditDate(event.EventKind, event.Date, event.EditDate)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      event.Out,
		Id:       messageID,
		FromId:   tg.MakePeerUser(event.FromUserID),
		PeerId:   sentMessagePeerFromOptional(event.PeerType, event.PeerID),
		ReplyTo:  sentMessageReplyHeader(event.ReplyToUserMessageID),
		Date:     date,
		Message:  event.MessageText,
		Entities: sentMessageEntities(event.Entities),
		EditDate: editDatePtr,
	}), nil
}

func userMessageEventV3ToTLMessage(box *repository.UserMessageBox, event payload.MessageEventV3) (tg.MessageClazz, error) {
	if event.SchemaVersion != payload.MessageEventSchemaVersionV3 ||
		event.PeerType != box.PeerType ||
		event.PeerID != box.PeerID ||
		event.PeerSeq != box.PeerSeq ||
		event.MessageID != box.UserMessageID ||
		event.CanonicalMessageID != box.CanonicalMessageID {
		return nil, fmt.Errorf("%w: user message view payload v3 mismatch", msg.ErrMsgStorage)
	}
	messageID, err := historyIDInt32(event.MessageID, "user message view id")
	if err != nil {
		return nil, err
	}
	date, err := msgDateInt32FromUnixSeconds(int64(event.Date), "user message view date")
	if err != nil {
		return nil, err
	}
	editDatePtr, err := userMessageEditDate(event.EventKind, event.Date, event.EditDate)
	if err != nil {
		return nil, err
	}
	fwdFrom, err := userMessageForwardHeader(event.ForwardRef)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:         event.Out,
		Silent:      event.Attrs != nil && event.Attrs.Silent,
		Noforwards:  event.Attrs != nil && event.Attrs.Noforwards,
		InvertMedia: event.Attrs != nil && event.Attrs.InvertMedia,
		Id:          messageID,
		FromId:      tg.MakePeerUser(event.FromUserID),
		PeerId:      sentMessagePeerFromOptional(event.PeerType, event.PeerID),
		FwdFrom:     fwdFrom,
		ReplyTo:     sentMessageReplyHeader(event.ReplyToUserMessageID),
		Date:        date,
		Message:     event.MessageText,
		Media:       sentMessageMedia(event.MediaRef),
		Entities:    sentMessageEntities(event.Entities),
		GroupedId:   sentMessageGroupedID(event.Attrs),
		TtlPeriod:   sentMessageTTLPeriod(event.MediaRef),
		EditDate:    editDatePtr,
	}), nil
}

func userMessageEditDate(eventKind string, date int32, editDate int32) (*int32, error) {
	switch eventKind {
	case payload.EventKindNewMessage:
		return nil, nil
	case payload.OperationKindEditMessage:
	default:
		return nil, fmt.Errorf("%w: unsupported user message view event kind=%s", msg.ErrMsgStorage, eventKind)
	}
	if editDate == 0 {
		editDate = date
	}
	v, err := msgDateInt32FromUnixSeconds(int64(editDate), "user message edit date")
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func userMessageForwardHeader(ref *payload.ForwardRefV1) (tg.MessageFwdHeaderClazz, error) {
	if ref == nil {
		return nil, nil
	}
	date, err := msgDateInt32FromUnixSeconds(ref.Date, "user message forward date")
	if err != nil {
		return nil, err
	}
	var sourceMessageID *int32
	if ref.SourceMessageID > 0 {
		v, err := historyIDInt32(ref.SourceMessageID, "user message forward source message id")
		if err != nil {
			return nil, err
		}
		sourceMessageID = &v
	}
	var savedFromMessageID *int32
	if ref.SavedFromMessageID > 0 {
		v, err := historyIDInt32(ref.SavedFromMessageID, "user message forward saved message id")
		if err != nil {
			return nil, err
		}
		savedFromMessageID = &v
	}
	return tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
		FromId:         sentMessageForwardPeer(ref),
		FromName:       stringPtr(ref.FromName),
		Date:           date,
		ChannelPost:    sourceMessageID,
		SavedFromPeer:  sentMessagePeerFromOptional(ref.SavedFromPeerType, ref.SavedFromPeerID),
		SavedFromMsgId: savedFromMessageID,
	}), nil
}

func userMessageReplyHeaderFromPeerSeq(peerSeq int64) (tg.MessageReplyHeaderClazz, error) {
	if peerSeq <= 0 {
		return nil, nil
	}
	replyToMsgID, err := historyIDInt32(peerSeq, "user message reply peer seq")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}
