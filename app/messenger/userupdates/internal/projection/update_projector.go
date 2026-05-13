package projection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/eventtypes"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Mode int

const (
	ModeDifference Mode = 1
	ModePush       Mode = 2
)

type Result struct {
	Message          tg.MessageClazz
	Update           tg.UpdateClazz
	OtherUpdates     []tg.UpdateClazz
	Updates          tg.UpdatesClazz
	AuthKeyIDExclude *int64
}

func ProjectUserEvent(event eventtypes.UserEvent, mode Mode) (Result, error) {
	if event.EventSchemaVersion == payload.MessageEventSchemaVersionV4 {
		return projectMessageEventV4(event, mode)
	}
	messageEvent, err := decodeUserEventPayload(event)
	if err != nil {
		return Result{}, err
	}
	if event.EventType == eventtypes.EventTypeDialogPublicUpdate {
		update, err := dialogEventToTLUpdate(dialogEventFromMessageEvent(event, messageEvent), event.Pts, event.PtsCount)
		if err != nil {
			return Result{}, err
		}
		return Result{Update: update}, nil
	}
	return projectMessageEvent(messageEventProjectionInput{
		mode:       mode,
		pts:        event.Pts,
		ptsCount:   event.PtsCount,
		eventType:  event.EventType,
		peerType:   event.PeerType,
		peerID:     event.PeerID,
		message:    messageEvent,
		datePrefix: "event",
	})
}

func projectMessageEventV4(event eventtypes.UserEvent, mode Mode) (Result, error) {
	messageEvent, err := decodeUserEventPayloadV4(event)
	if err != nil {
		return Result{}, err
	}
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return Result{}, fmt.Errorf("%w: unsupported v4 event kind=%s", userupdates.ErrUserupdatesStorage, messageEvent.EventKind)
	}
	facts := make([]payload.UpdateFactV1, 0, len(messageEvent.AttachFacts)+1)
	facts = append(facts, messageEvent.AttachFacts...)
	messageFact, err := payload.WrapFact(payload.FactKindNewMessage, messageEvent.MessageFact)
	if err != nil {
		return Result{}, fmt.Errorf("%w: wrap v4 message fact: %v", userupdates.ErrUserupdatesStorage, err)
	}
	facts = append(facts, messageFact)
	projected, err := ProjectFacts(facts, ViewerContext{
		UserID:           event.UserID,
		AuthKeyIDExclude: messageEvent.AuthKeyIdExclude,
	}, factProjectionMode(mode), event.Pts, event.PtsCount, messageEvent.MessageID)
	if err != nil {
		return Result{}, err
	}
	updates := make([]tg.UpdateClazz, 0, len(projected))
	var message tg.MessageClazz
	for _, update := range projected {
		if update.Update == nil {
			continue
		}
		updates = append(updates, update.Update)
		if newMessage, ok := update.Update.(*tg.TLUpdateNewMessage); ok && message == nil {
			message = newMessage.Message
		}
	}
	var firstUpdate tg.UpdateClazz
	if len(updates) > 0 {
		firstUpdate = updates[0]
	}
	return Result{
		Message:          message,
		Update:           firstUpdate,
		OtherUpdates:     updates,
		AuthKeyIDExclude: messageEvent.AuthKeyIdExclude,
	}, nil
}

func factProjectionMode(mode Mode) envelope.Mode {
	if mode == ModePush {
		return envelope.ModeEphemeralPush
	}
	return envelope.ModeDifference
}

func ProjectPushTask(msg *payload.PushTaskKafkaMessageV1) (Result, error) {
	if msg == nil {
		return Result{}, fmt.Errorf("%w: push task is nil", userupdates.ErrUserupdatesStorage)
	}
	if detectMessageEventSchemaVersion(msg.Payload) == payload.MessageEventSchemaVersionV4 {
		return projectPushTaskV4(msg)
	}
	messageEvent, err := decodeMessageEventPayloadBytes(detectMessageEventSchemaVersion(msg.Payload), msg.Payload)
	if err != nil {
		return Result{}, err
	}
	return projectMessageEvent(messageEventProjectionInput{
		mode:       ModePush,
		pts:        msg.Pts,
		ptsCount:   1,
		peerType:   msg.PeerType,
		peerID:     msg.PeerID,
		message:    messageEvent,
		datePrefix: "push",
	})
}

func projectPushTaskV4(msg *payload.PushTaskKafkaMessageV1) (Result, error) {
	return ProjectPushTaskV4Updates(msg)
}

func ProjectPushTaskV4Updates(msg *payload.PushTaskKafkaMessageV1) (Result, error) {
	if msg == nil {
		return Result{}, fmt.Errorf("%w: push task is nil", userupdates.ErrUserupdatesStorage)
	}
	var messageEvent payload.MessageEventV4
	if err := json.Unmarshal(msg.Payload, &messageEvent); err != nil {
		return Result{}, fmt.Errorf("%w: decode v4 push message event: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersionV4 {
		return Result{}, fmt.Errorf("%w: unsupported v4 push message event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return Result{}, fmt.Errorf("%w: unsupported v4 push event kind=%s", userupdates.ErrUserupdatesStorage, messageEvent.EventKind)
	}
	facts := make([]payload.UpdateFactV1, 0, len(messageEvent.AttachFacts)+1)
	facts = append(facts, messageEvent.AttachFacts...)
	messageFact, err := payload.WrapFact(payload.FactKindNewMessage, messageEvent.MessageFact)
	if err != nil {
		return Result{}, fmt.Errorf("%w: wrap v4 push message fact: %v", userupdates.ErrUserupdatesStorage, err)
	}
	facts = append(facts, messageFact)
	projected, err := ProjectFacts(facts, ViewerContext{
		UserID:           msg.UserID,
		AuthKeyIDExclude: messageEvent.AuthKeyIdExclude,
	}, envelope.ModeEphemeralPush, msg.Pts, 1, messageEvent.MessageID)
	if err != nil {
		return Result{}, err
	}
	updates := make([]tg.UpdateClazz, 0, len(projected))
	for _, item := range projected {
		if item.Update != nil {
			updates = append(updates, item.Update)
		}
	}
	return Result{
		OtherUpdates:     updates,
		AuthKeyIDExclude: messageEvent.AuthKeyIdExclude,
	}, nil
}

type messageEventProjectionInput struct {
	mode       Mode
	pts        int64
	ptsCount   int32
	eventType  int32
	peerType   int32
	peerID     int64
	message    decodedMessageEvent
	datePrefix string
}

type decodedMessageEvent struct {
	SchemaVersion        int
	EventKind            string
	CanonicalMessageID   int64
	PeerSeq              int64
	MessageID            int64
	PeerType             int32
	PeerID               int64
	FromUserID           int64
	ToUserID             int64
	Date                 int32
	EditDate             int32
	EditVersion          int32
	Out                  bool
	MessageText          string
	Entities             []payload.MessageEntityV1
	ReplyToUserMessageID int64
	ReadMaxUserMessageID int64
	DeleteUserMessageIDs []int64
	PinnedUserMessageID  int64
	AuthKeyIdExclude     *int64
	MediaRef             *payload.MediaRefV1
	Attrs                *payload.MessageAttrsV1
	ForwardRef           *payload.ForwardRefV1
	ServiceAction        *payload.ServiceActionRefV1
}

func projectMessageEvent(in messageEventProjectionInput) (Result, error) {
	switch in.message.EventKind {
	case payload.EventKindNewMessage:
		return projectNewMessage(in)
	case payload.OperationKindReadHistory:
		update, err := readHistoryUpdate(in)
		if err != nil {
			return Result{}, err
		}
		if in.mode == ModePush {
			updates, err := wrapPushUpdate(update, in.message.Date)
			if err != nil {
				return Result{}, err
			}
			return Result{Updates: updates, AuthKeyIDExclude: in.message.AuthKeyIdExclude}, nil
		}
		return Result{Update: update}, nil
	case payload.OperationKindDeleteMessages:
		update, err := deleteMessagesUpdate(in)
		if err != nil {
			return Result{}, err
		}
		if in.mode == ModePush {
			updates, err := wrapPushUpdate(update, in.message.Date)
			if err != nil {
				return Result{}, err
			}
			return Result{Updates: updates, AuthKeyIDExclude: in.message.AuthKeyIdExclude}, nil
		}
		return Result{Update: update}, nil
	case payload.OperationKindEditMessage:
		update, err := editMessageUpdate(in)
		if err != nil {
			return Result{}, err
		}
		if in.mode == ModePush {
			updateDate := int64(in.message.EditDate)
			if updateDate == 0 {
				updateDate = int64(in.message.Date)
			}
			updates, err := wrapPushUpdate(update, int32(updateDate-1))
			if err != nil {
				return Result{}, err
			}
			return Result{Updates: updates, AuthKeyIDExclude: in.message.AuthKeyIdExclude}, nil
		}
		return Result{Update: update}, nil
	case payload.OperationKindUpdatePinnedMessage:
		update, err := updatePinnedMessageUpdate(in)
		if err != nil {
			return Result{}, err
		}
		if in.mode == ModePush {
			updates, err := wrapPushUpdate(update, in.message.Date)
			if err != nil {
				return Result{}, err
			}
			return Result{Updates: updates, AuthKeyIDExclude: in.message.AuthKeyIdExclude}, nil
		}
		return Result{Update: update}, nil
	default:
		return Result{}, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, in.message.EventKind, in.message.SchemaVersion)
	}
}

func decodeUserEventPayload(event eventtypes.UserEvent) (decodedMessageEvent, error) {
	if event.EventCodec != eventtypes.PayloadCodecJSON {
		return decodedMessageEvent{}, fmt.Errorf("%w: unsupported event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if len(event.EventPayloadHash) != 0 && !bytes.Equal(payload.HashBytes(event.EventPayload), event.EventPayloadHash) {
		return decodedMessageEvent{}, fmt.Errorf("%w: event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}
	return decodeMessageEventPayloadBytes(event.EventSchemaVersion, event.EventPayload)
}

func decodeUserEventPayloadV4(event eventtypes.UserEvent) (payload.MessageEventV4, error) {
	if event.EventCodec != eventtypes.PayloadCodecJSON {
		return payload.MessageEventV4{}, fmt.Errorf("%w: unsupported event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if len(event.EventPayloadHash) != 0 && !bytes.Equal(payload.HashBytes(event.EventPayload), event.EventPayloadHash) {
		return payload.MessageEventV4{}, fmt.Errorf("%w: event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}
	var messageEvent payload.MessageEventV4
	if err := json.Unmarshal(event.EventPayload, &messageEvent); err != nil {
		return payload.MessageEventV4{}, fmt.Errorf("%w: decode v4 message event: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersionV4 {
		return payload.MessageEventV4{}, fmt.Errorf("%w: unsupported v4 message event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	return messageEvent, nil
}

func decodeMessageEventPayloadBytes(schemaVersion int32, body []byte) (decodedMessageEvent, error) {
	switch int(schemaVersion) {
	case payload.MessageEventSchemaVersionV1:
		var old payload.MessageEventV1
		if err := json.Unmarshal(body, &old); err != nil {
			return decodedMessageEvent{}, fmt.Errorf("%w: decode v1 message event: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if old.SchemaVersion != payload.MessageEventSchemaVersionV1 {
			return decodedMessageEvent{}, fmt.Errorf("%w: unsupported v1 message event schema=%d", userupdates.ErrUserupdatesStorage, old.SchemaVersion)
		}
		return decodedMessageEvent{
			SchemaVersion:      old.SchemaVersion,
			EventKind:          old.EventKind,
			CanonicalMessageID: old.CanonicalMessageID,
			PeerSeq:            old.MessageID,
			PeerType:           old.PeerType,
			PeerID:             old.PeerID,
			FromUserID:         old.FromUserID,
			ToUserID:           old.ToUserID,
			Date:               old.Date,
			EditDate:           old.EditDate,
			EditVersion:        old.EditVersion,
			Out:                old.Out,
			MessageText:        old.MessageText,
			Entities:           old.Entities,
			AuthKeyIdExclude:   old.AuthKeyIdExclude,
		}, nil
	case payload.MessageEventSchemaVersion:
		var next payload.MessageEventV2
		if err := json.Unmarshal(body, &next); err != nil {
			return decodedMessageEvent{}, fmt.Errorf("%w: decode v2 message event: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if next.SchemaVersion != payload.MessageEventSchemaVersion {
			return decodedMessageEvent{}, fmt.Errorf("%w: unsupported v2 message event schema=%d", userupdates.ErrUserupdatesStorage, next.SchemaVersion)
		}
		return decodedMessageEvent{
			SchemaVersion:        next.SchemaVersion,
			EventKind:            next.EventKind,
			CanonicalMessageID:   next.CanonicalMessageID,
			PeerSeq:              next.PeerSeq,
			MessageID:            next.MessageID,
			PeerType:             next.PeerType,
			PeerID:               next.PeerID,
			FromUserID:           next.FromUserID,
			ToUserID:             next.ToUserID,
			Date:                 next.Date,
			EditDate:             next.EditDate,
			EditVersion:          next.EditVersion,
			Out:                  next.Out,
			MessageText:          next.MessageText,
			Entities:             next.Entities,
			ReplyToUserMessageID: next.ReplyToUserMessageID,
			ReadMaxUserMessageID: next.ReadMaxUserMessageID,
			DeleteUserMessageIDs: append([]int64(nil), next.DeleteUserMessageIDs...),
			PinnedUserMessageID:  next.PinnedUserMessageID,
			AuthKeyIdExclude:     next.AuthKeyIdExclude,
		}, nil
	case payload.MessageEventSchemaVersionV3:
		var next payload.MessageEventV3
		if err := json.Unmarshal(body, &next); err != nil {
			return decodedMessageEvent{}, fmt.Errorf("%w: decode v3 message event: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if next.SchemaVersion != payload.MessageEventSchemaVersionV3 {
			return decodedMessageEvent{}, fmt.Errorf("%w: unsupported v3 message event schema=%d", userupdates.ErrUserupdatesStorage, next.SchemaVersion)
		}
		return decodedMessageEvent{
			SchemaVersion:        next.SchemaVersion,
			EventKind:            next.EventKind,
			CanonicalMessageID:   next.CanonicalMessageID,
			PeerSeq:              next.PeerSeq,
			MessageID:            next.MessageID,
			PeerType:             next.PeerType,
			PeerID:               next.PeerID,
			FromUserID:           next.FromUserID,
			ToUserID:             next.ToUserID,
			Date:                 next.Date,
			EditDate:             next.EditDate,
			EditVersion:          next.EditVersion,
			Out:                  next.Out,
			MessageText:          next.MessageText,
			Entities:             next.Entities,
			ReplyToUserMessageID: next.ReplyToUserMessageID,
			ReadMaxUserMessageID: next.ReadMaxUserMessageID,
			DeleteUserMessageIDs: append([]int64(nil), next.DeleteUserMessageIDs...),
			PinnedUserMessageID:  next.PinnedUserMessageID,
			AuthKeyIdExclude:     next.AuthKeyIdExclude,
			MediaRef:             next.MediaRef,
			Attrs:                next.Attrs,
			ForwardRef:           next.ForwardRef,
			ServiceAction:        next.ServiceAction,
		}, nil
	default:
		return decodedMessageEvent{}, fmt.Errorf("%w: unsupported message event schema=%d", userupdates.ErrUserupdatesStorage, schemaVersion)
	}
}

func detectMessageEventSchemaVersion(body []byte) int32 {
	var envelope struct {
		SchemaVersion int `json:"schema_version"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return 0
	}
	return int32(envelope.SchemaVersion)
}

func projectNewMessage(in messageEventProjectionInput) (Result, error) {
	message, err := messageEventToTLMessage(in.message)
	if err != nil {
		return Result{}, err
	}
	pts, err := int64ToInt32(in.pts, "pts")
	if err != nil {
		return Result{}, err
	}
	if in.mode == ModePush {
		updates, err := wrapPushUpdate(tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: in.ptsCount,
		}), in.message.Date)
		if err != nil {
			return Result{}, err
		}
		return Result{Updates: updates, AuthKeyIDExclude: in.message.AuthKeyIdExclude}, nil
	}
	return Result{
		Message: message,
		Update: tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  message,
			Pts:      pts,
			PtsCount: in.ptsCount,
		}),
	}, nil
}

func readHistoryUpdate(in messageEventProjectionInput) (tg.UpdateClazz, error) {
	maxUserMessageID := in.message.ReadMaxUserMessageID
	if maxUserMessageID == 0 {
		maxUserMessageID = in.message.MessageID
	}
	maxID, err := messageIDInt32(maxUserMessageID, "read max user message id")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(in.pts, "pts")
	if err != nil {
		return nil, err
	}
	peer := peerFromEvent(in.message.PeerType, in.message.PeerID)
	if in.message.Out {
		return tg.MakeTLUpdateReadHistoryOutbox(&tg.TLUpdateReadHistoryOutbox{
			Peer:     peer,
			MaxId:    maxID,
			Pts:      pts,
			PtsCount: in.ptsCount,
		}), nil
	}
	return tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{
		Peer:             peer,
		MaxId:            maxID,
		StillUnreadCount: 0,
		Pts:              pts,
		PtsCount:         in.ptsCount,
	}), nil
}

func deleteMessagesUpdate(in messageEventProjectionInput) (tg.UpdateClazz, error) {
	messages := make([]int32, 0, len(in.message.DeleteUserMessageIDs))
	for _, id := range in.message.DeleteUserMessageIDs {
		msgID, err := messageIDInt32(id, "delete user message id")
		if err != nil {
			return nil, err
		}
		if msgID > 0 {
			messages = append(messages, msgID)
		}
	}
	pts, err := int64ToInt32(in.pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateDeleteMessages(&tg.TLUpdateDeleteMessages{
		Messages: messages,
		Pts:      pts,
		PtsCount: in.ptsCount,
	}), nil
}

func editMessageUpdate(in messageEventProjectionInput) (tg.UpdateClazz, error) {
	message, err := editMessageEventToTLMessage(in.message)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(in.pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
		Message:  message,
		Pts:      pts,
		PtsCount: in.ptsCount,
	}), nil
}

func updatePinnedMessageUpdate(in messageEventProjectionInput) (tg.UpdateClazz, error) {
	pts, err := int64ToInt32(in.pts, "pts")
	if err != nil {
		return nil, err
	}
	messages := []int32(nil)
	if in.message.PinnedUserMessageID > 0 {
		msgID, err := messageIDInt32(in.message.PinnedUserMessageID, "pinned user message id")
		if err != nil {
			return nil, err
		}
		messages = []int32{msgID}
	}
	return tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{
		Pinned:   in.message.PinnedUserMessageID > 0,
		Peer:     peerFromEvent(in.message.PeerType, in.message.PeerID),
		Messages: messages,
		Pts:      pts,
		PtsCount: in.ptsCount,
	}), nil
}

func wrapPushUpdate(update tg.UpdateClazz, dateSeconds int32) (tg.UpdatesClazz, error) {
	return wrapPushUpdates([]tg.UpdateClazz{update}, dateSeconds)
}

func wrapPushUpdates(updates []tg.UpdateClazz, dateSeconds int32) (tg.UpdatesClazz, error) {
	date, err := userupdatesDateInt32FromUnixSeconds(int64(dateSeconds), "updates date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: updates,
		Users:   []tg.UserClazz{},
		Chats:   []tg.ChatClazz{},
		Date:    date,
		Seq:     0,
	}), nil
}

func messageEventToTLMessage(messageEvent decodedMessageEvent) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := messageIDInt32(messageEvent.MessageID, "message id")
	if err != nil {
		return nil, err
	}
	replyTo, err := replyHeaderFromUserMessageID(messageEvent.ReplyToUserMessageID)
	if err != nil {
		return nil, err
	}
	date, err := userupdatesDateInt32FromUnixSeconds(int64(messageEvent.Date), "message date")
	if err != nil {
		return nil, err
	}
	if messageEvent.ServiceAction != nil {
		action, err := messageServiceAction(messageEvent.ServiceAction)
		if err != nil {
			return nil, err
		}
		return tg.MakeTLMessageService(&tg.TLMessageService{
			Out:    messageEvent.Out,
			Silent: messageAttrsSilent(messageEvent.Attrs),
			Id:     messageID,
			FromId: messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
			PeerId: peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
			Date:   date,
			Action: action,
		}), nil
	}
	fwdFrom, err := messageForwardHeader(messageEvent.ForwardRef)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:         messageEvent.Out,
		Silent:      messageAttrsSilent(messageEvent.Attrs),
		Noforwards:  messageAttrsNoforwards(messageEvent.Attrs),
		InvertMedia: messageAttrsInvertMedia(messageEvent.Attrs),
		Id:          messageID,
		FromId:      messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:      peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		FwdFrom:     fwdFrom,
		ReplyTo:     replyTo,
		Date:        date,
		Message:     messageEvent.MessageText,
		Media:       messageMedia(messageEvent.MediaRef),
		Entities:    messageEntities(messageEvent.Entities),
		GroupedId:   messageGroupedID(messageEvent.Attrs),
		TtlPeriod:   messageTTLPeriod(messageEvent.MediaRef),
	}), nil
}

func messageServiceAction(ref *payload.ServiceActionRefV1) (tg.MessageActionClazz, error) {
	if ref == nil {
		return nil, nil
	}
	switch ref.Kind {
	case payload.ServiceActionKindChatCreate:
		return tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
			Title: ref.Title,
			Users: append([]int64(nil), ref.Users...),
		}), nil
	default:
		return nil, fmt.Errorf("%w: unsupported service action kind=%s schema=%d", userupdates.ErrUserupdatesStorage, ref.Kind, ref.SchemaVersion)
	}
}

func editMessageEventToTLMessage(messageEvent decodedMessageEvent) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.OperationKindEditMessage {
		return nil, fmt.Errorf("%w: unsupported edit event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := messageIDInt32(messageEvent.MessageID, "message id")
	if err != nil {
		return nil, err
	}
	editDate := messageEvent.EditDate
	if editDate == 0 {
		editDate = messageEvent.Date
	}
	date, err := userupdatesDateInt32FromUnixSeconds(int64(messageEvent.Date), "edit message date")
	if err != nil {
		return nil, err
	}
	editDate32, err := userupdatesDateInt32FromUnixSeconds(int64(editDate), "edit date")
	if err != nil {
		return nil, err
	}
	fwdFrom, err := messageForwardHeader(messageEvent.ForwardRef)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:         messageEvent.Out,
		Silent:      messageAttrsSilent(messageEvent.Attrs),
		Noforwards:  messageAttrsNoforwards(messageEvent.Attrs),
		InvertMedia: messageAttrsInvertMedia(messageEvent.Attrs),
		Id:          messageID,
		FromId:      messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:      peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		FwdFrom:     fwdFrom,
		Date:        date,
		Message:     messageEvent.MessageText,
		Media:       messageMedia(messageEvent.MediaRef),
		Entities:    messageEntities(messageEvent.Entities),
		GroupedId:   messageGroupedID(messageEvent.Attrs),
		TtlPeriod:   messageTTLPeriod(messageEvent.MediaRef),
		EditDate:    &editDate32,
	}), nil
}

func messageEntities(entities []payload.MessageEntityV1) []tg.MessageEntityClazz {
	if len(entities) == 0 {
		return nil
	}
	out := make([]tg.MessageEntityClazz, 0, len(entities))
	for _, entity := range entities {
		switch entity.Kind {
		case "mention":
			out = append(out, tg.MakeTLMessageEntityMention(&tg.TLMessageEntityMention{Offset: entity.Offset, Length: entity.Length}))
		case "hashtag":
			out = append(out, tg.MakeTLMessageEntityHashtag(&tg.TLMessageEntityHashtag{Offset: entity.Offset, Length: entity.Length}))
		case "bot_command":
			out = append(out, tg.MakeTLMessageEntityBotCommand(&tg.TLMessageEntityBotCommand{Offset: entity.Offset, Length: entity.Length}))
		case "url":
			out = append(out, tg.MakeTLMessageEntityUrl(&tg.TLMessageEntityUrl{Offset: entity.Offset, Length: entity.Length}))
		case "email":
			out = append(out, tg.MakeTLMessageEntityEmail(&tg.TLMessageEntityEmail{Offset: entity.Offset, Length: entity.Length}))
		case "bold":
			out = append(out, tg.MakeTLMessageEntityBold(&tg.TLMessageEntityBold{Offset: entity.Offset, Length: entity.Length}))
		case "italic":
			out = append(out, tg.MakeTLMessageEntityItalic(&tg.TLMessageEntityItalic{Offset: entity.Offset, Length: entity.Length}))
		case "code":
			out = append(out, tg.MakeTLMessageEntityCode(&tg.TLMessageEntityCode{Offset: entity.Offset, Length: entity.Length}))
		case "pre":
			out = append(out, tg.MakeTLMessageEntityPre(&tg.TLMessageEntityPre{Offset: entity.Offset, Length: entity.Length, Language: entity.URL}))
		case "text_url":
			out = append(out, tg.MakeTLMessageEntityTextUrl(&tg.TLMessageEntityTextUrl{Offset: entity.Offset, Length: entity.Length, Url: entity.URL}))
		case "mention_name":
			out = append(out, tg.MakeTLMessageEntityMentionName(&tg.TLMessageEntityMentionName{Offset: entity.Offset, Length: entity.Length, UserId: entity.UserID}))
		case "phone":
			out = append(out, tg.MakeTLMessageEntityPhone(&tg.TLMessageEntityPhone{Offset: entity.Offset, Length: entity.Length}))
		case "cashtag":
			out = append(out, tg.MakeTLMessageEntityCashtag(&tg.TLMessageEntityCashtag{Offset: entity.Offset, Length: entity.Length}))
		case "underline":
			out = append(out, tg.MakeTLMessageEntityUnderline(&tg.TLMessageEntityUnderline{Offset: entity.Offset, Length: entity.Length}))
		case "strike":
			out = append(out, tg.MakeTLMessageEntityStrike(&tg.TLMessageEntityStrike{Offset: entity.Offset, Length: entity.Length}))
		case "bank_card":
			out = append(out, tg.MakeTLMessageEntityBankCard(&tg.TLMessageEntityBankCard{Offset: entity.Offset, Length: entity.Length}))
		case "spoiler":
			out = append(out, tg.MakeTLMessageEntitySpoiler(&tg.TLMessageEntitySpoiler{Offset: entity.Offset, Length: entity.Length}))
		case "blockquote":
			out = append(out, tg.MakeTLMessageEntityBlockquote(&tg.TLMessageEntityBlockquote{Offset: entity.Offset, Length: entity.Length}))
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func messageMedia(media *payload.MediaRefV1) tg.MessageMediaClazz {
	if media == nil {
		return nil
	}
	ttl := messageTTLPeriod(media)
	switch media.Kind {
	case "photo":
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      messagePhoto(media),
			TtlSeconds: ttl,
		})
	case "document":
		flags := messageDocumentMediaFlags(media)
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Spoiler:        flags.Spoiler,
			Video:          flags.Video,
			Round:          flags.Round,
			Voice:          flags.Voice,
			Document:       messageDocument(media),
			VideoCover:     messagePhotoRef(media.VideoCover),
			VideoTimestamp: cloneInt32Ptr(media.VideoTimestamp),
			TtlSeconds:     ttl,
		})
	case "contact":
		return messageContact(media)
	default:
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{})
	}
}

func messageContact(media *payload.MediaRefV1) tg.MessageMediaClazz {
	return tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
		PhoneNumber: media.PhoneNumber,
		FirstName:   media.FirstName,
		LastName:    media.LastName,
		Vcard:       media.Vcard,
		UserId:      media.UserID,
	})
}

func messagePhoto(media *payload.MediaRefV1) tg.PhotoClazz {
	if media.Date == 0 && media.DcID == 0 && len(media.PhotoSizes) == 0 {
		return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: media.ID})
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		Id:            media.ID,
		AccessHash:    media.AccessHash,
		FileReference: append([]byte(nil), media.FileReference...),
		Date:          media.Date,
		Sizes:         messagePhotoSizes(media.PhotoSizes),
		DcId:          media.DcID,
	})
}

func messageDocument(media *payload.MediaRefV1) tg.DocumentClazz {
	if media.Date == 0 && media.DcID == 0 && media.Size == 0 && len(media.DocumentAttributes) == 0 {
		return tg.MakeTLDocumentEmpty(&tg.TLDocumentEmpty{Id: media.ID})
	}
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:            media.ID,
		AccessHash:    media.AccessHash,
		FileReference: append([]byte(nil), media.FileReference...),
		Date:          media.Date,
		MimeType:      media.MimeType,
		Size2:         media.Size,
		Thumbs:        messagePhotoSizes(media.DocumentThumbs),
		VideoThumbs:   messageVideoSizes(media.DocumentVideoThumbs),
		DcId:          media.DcID,
		Attributes:    messageDocumentAttributes(media.DocumentAttributes),
	})
}

func messagePhotoRef(photo *payload.PhotoRefV1) tg.PhotoClazz {
	if photo == nil {
		return nil
	}
	if photo.Date == 0 && photo.DcID == 0 && len(photo.Sizes) == 0 && len(photo.VideoSizes) == 0 {
		return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: photo.ID})
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		Id:            photo.ID,
		AccessHash:    photo.AccessHash,
		FileReference: append([]byte(nil), photo.FileReference...),
		Date:          photo.Date,
		Sizes:         messagePhotoSizes(photo.Sizes),
		VideoSizes:    messageVideoSizes(photo.VideoSizes),
		DcId:          photo.DcID,
	})
}

func messagePhotoSizes(sizes []payload.PhotoSizeRefV1) []tg.PhotoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]tg.PhotoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		switch size.Kind {
		case "empty":
			out = append(out, tg.MakeTLPhotoSizeEmpty(&tg.TLPhotoSizeEmpty{Type: size.Type}))
		case "size":
			out = append(out, tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: size.Type, W: size.W, H: size.H, Size2: size.Size}))
		case "cached":
			out = append(out, tg.MakeTLPhotoCachedSize(&tg.TLPhotoCachedSize{Type: size.Type, W: size.W, H: size.H, Bytes: append([]byte(nil), size.Bytes...)}))
		case "stripped":
			out = append(out, tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{Type: size.Type, Bytes: append([]byte(nil), size.Bytes...)}))
		case "progressive":
			out = append(out, tg.MakeTLPhotoSizeProgressive(&tg.TLPhotoSizeProgressive{Type: size.Type, W: size.W, H: size.H, Sizes: append([]int32(nil), size.Sizes...)}))
		case "path":
			out = append(out, tg.MakeTLPhotoPathSize(&tg.TLPhotoPathSize{Type: size.Type, Bytes: append([]byte(nil), size.Bytes...)}))
		}
	}
	return out
}

func messageVideoSizes(sizes []payload.VideoSizeRefV1) []tg.VideoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]tg.VideoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		switch size.Kind {
		case "size":
			out = append(out, tg.MakeTLVideoSize(&tg.TLVideoSize{
				Type:         size.Type,
				W:            size.W,
				H:            size.H,
				Size2:        size.Size,
				VideoStartTs: cloneFloat64Ptr(size.VideoStartTs),
			}))
		case "emoji_markup":
			out = append(out, tg.MakeTLVideoSizeEmojiMarkup(&tg.TLVideoSizeEmojiMarkup{
				EmojiId:          size.EmojiID,
				BackgroundColors: append([]int32(nil), size.BackgroundColors...),
			}))
		case "sticker_markup":
			out = append(out, tg.MakeTLVideoSizeStickerMarkup(&tg.TLVideoSizeStickerMarkup{
				Stickerset:       messageStickerSetRef(size.StickerSet),
				StickerId:        size.StickerID,
				BackgroundColors: append([]int32(nil), size.BackgroundColors...),
			}))
		}
	}
	return out
}

func messageDocumentAttributes(attrs []payload.DocumentAttributeRefV1) []tg.DocumentAttributeClazz {
	if len(attrs) == 0 {
		return nil
	}
	out := make([]tg.DocumentAttributeClazz, 0, len(attrs))
	for _, attr := range attrs {
		switch attr.Kind {
		case "filename":
			out = append(out, tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: attr.FileName}))
		case "image_size":
			out = append(out, tg.MakeTLDocumentAttributeImageSize(&tg.TLDocumentAttributeImageSize{W: attr.W, H: attr.H}))
		case "animated":
			out = append(out, tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}))
		case "video":
			out = append(out, tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{
				RoundMessage:      attr.RoundMessage,
				SupportsStreaming: attr.SupportsStreaming,
				Nosound:           attr.NoSound,
				Duration:          attr.DurationFloat,
				W:                 attr.W,
				H:                 attr.H,
				PreloadPrefixSize: attr.PreloadPrefixSize,
				VideoStartTs:      attr.VideoStartTs,
				VideoCodec:        attr.VideoCodec,
			}))
		case "audio":
			out = append(out, tg.MakeTLDocumentAttributeAudio(&tg.TLDocumentAttributeAudio{
				Voice:     attr.Voice,
				Duration:  attr.Duration,
				Title:     attr.Title,
				Performer: attr.Performer,
				Waveform:  append([]byte(nil), attr.Waveform...),
			}))
		case "sticker":
			out = append(out, tg.MakeTLDocumentAttributeSticker(&tg.TLDocumentAttributeSticker{
				Mask:       attr.Mask,
				Alt:        attr.Alt,
				Stickerset: messageStickerSetRef(attr.StickerSet),
				MaskCoords: messageMaskCoordsRef(attr.MaskCoords),
			}))
		case "custom_emoji":
			out = append(out, tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{
				Free:       attr.Free,
				TextColor:  attr.TextColor,
				Alt:        attr.Alt,
				Stickerset: messageStickerSetRef(attr.StickerSet),
			}))
		case "has_stickers":
			out = append(out, tg.MakeTLDocumentAttributeHasStickers(&tg.TLDocumentAttributeHasStickers{}))
		}
	}
	return out
}

func messageDocumentMediaFlags(media *payload.MediaRefV1) payload.DocumentMediaFlagsV1 {
	if media == nil {
		return payload.DocumentMediaFlagsV1{}
	}
	if media.DocumentMediaFlags != nil {
		return *media.DocumentMediaFlags
	}
	var flags payload.DocumentMediaFlagsV1
	for _, attr := range media.DocumentAttributes {
		switch attr.Kind {
		case "audio":
			flags.Voice = flags.Voice || attr.Voice
		case "video":
			flags.Round = flags.Round || attr.RoundMessage
			if !messageIsWebMStickerOrCustomEmoji(media) {
				flags.Video = true
			}
		}
	}
	return flags
}

func messageIsWebMStickerOrCustomEmoji(media *payload.MediaRefV1) bool {
	if media == nil || media.MimeType != "video/webm" {
		return false
	}
	for _, attr := range media.DocumentAttributes {
		if attr.Kind == "sticker" || attr.Kind == "custom_emoji" {
			return true
		}
	}
	return false
}

func messageStickerSetRef(ref *payload.StickerSetRefV1) tg.InputStickerSetClazz {
	if ref == nil {
		return tg.MakeTLInputStickerSetEmpty(&tg.TLInputStickerSetEmpty{})
	}
	switch ref.Kind {
	case "", "empty":
		return tg.MakeTLInputStickerSetEmpty(&tg.TLInputStickerSetEmpty{})
	case "id":
		return tg.MakeTLInputStickerSetID(&tg.TLInputStickerSetID{Id: ref.ID, AccessHash: ref.AccessHash})
	case "short_name":
		return tg.MakeTLInputStickerSetShortName(&tg.TLInputStickerSetShortName{ShortName: ref.ShortName})
	case "animated_emoji":
		return tg.MakeTLInputStickerSetAnimatedEmoji(&tg.TLInputStickerSetAnimatedEmoji{})
	case "dice":
		return tg.MakeTLInputStickerSetDice(&tg.TLInputStickerSetDice{Emoticon: ref.Emoticon})
	case "animated_emoji_animations":
		return tg.MakeTLInputStickerSetAnimatedEmojiAnimations(&tg.TLInputStickerSetAnimatedEmojiAnimations{})
	case "premium_gifts":
		return tg.MakeTLInputStickerSetPremiumGifts(&tg.TLInputStickerSetPremiumGifts{})
	case "emoji_generic_animations":
		return tg.MakeTLInputStickerSetEmojiGenericAnimations(&tg.TLInputStickerSetEmojiGenericAnimations{})
	case "emoji_default_statuses":
		return tg.MakeTLInputStickerSetEmojiDefaultStatuses(&tg.TLInputStickerSetEmojiDefaultStatuses{})
	case "emoji_default_topic_icons":
		return tg.MakeTLInputStickerSetEmojiDefaultTopicIcons(&tg.TLInputStickerSetEmojiDefaultTopicIcons{})
	case "emoji_channel_default_statuses":
		return tg.MakeTLInputStickerSetEmojiChannelDefaultStatuses(&tg.TLInputStickerSetEmojiChannelDefaultStatuses{})
	case "ton_gifts":
		return tg.MakeTLInputStickerSetTonGifts(&tg.TLInputStickerSetTonGifts{})
	default:
		return nil
	}
}

func messageMaskCoordsRef(ref *payload.MaskCoordsRefV1) tg.MaskCoordsClazz {
	if ref == nil {
		return nil
	}
	return tg.MakeTLMaskCoords(&tg.TLMaskCoords{N: ref.N, X: ref.X, Y: ref.Y, Zoom: ref.Zoom})
}

func cloneInt32Ptr(v *int32) *int32 {
	if v == nil {
		return nil
	}
	out := *v
	return &out
}

func cloneFloat64Ptr(v *float64) *float64 {
	if v == nil {
		return nil
	}
	out := *v
	return &out
}

func messageGroupedID(attrs *payload.MessageAttrsV1) *int64 {
	if attrs == nil || attrs.GroupedID == 0 {
		return nil
	}
	groupedID := attrs.GroupedID
	return &groupedID
}

func messageTTLPeriod(media *payload.MediaRefV1) *int32 {
	if media == nil || media.TTLSeconds == 0 {
		return nil
	}
	ttl := media.TTLSeconds
	return &ttl
}

func messageAttrsSilent(attrs *payload.MessageAttrsV1) bool {
	return attrs != nil && attrs.Silent
}

func messageAttrsNoforwards(attrs *payload.MessageAttrsV1) bool {
	return attrs != nil && attrs.Noforwards
}

func messageAttrsInvertMedia(attrs *payload.MessageAttrsV1) bool {
	return attrs != nil && attrs.InvertMedia
}

func messageForwardHeader(ref *payload.ForwardRefV1) (tg.MessageFwdHeaderClazz, error) {
	if ref == nil {
		return nil, nil
	}
	date, err := userupdatesDateInt32FromUnixSeconds(ref.Date, "forward date")
	if err != nil {
		return nil, err
	}
	fromName := stringPtr(ref.FromName)
	var sourceMessageID *int32
	if ref.SourceMessageID > 0 {
		v, err := int64ToInt32(ref.SourceMessageID, "forward source message id")
		if err != nil {
			return nil, err
		}
		sourceMessageID = &v
	}
	var savedFromMessageID *int32
	if ref.SavedFromMessageID > 0 {
		v, err := int64ToInt32(ref.SavedFromMessageID, "forward saved message id")
		if err != nil {
			return nil, err
		}
		savedFromMessageID = &v
	}
	return tg.MakeTLMessageFwdHeader(&tg.TLMessageFwdHeader{
		FromId:         forwardPeer(ref.FromUserID, ref.SourcePeerType, ref.SourcePeerID),
		FromName:       fromName,
		Date:           date,
		ChannelPost:    sourceMessageID,
		SavedFromPeer:  peerFromOptional(ref.SavedFromPeerType, ref.SavedFromPeerID),
		SavedFromMsgId: savedFromMessageID,
	}), nil
}

func forwardPeer(fromUserID int64, sourcePeerType int32, sourcePeerID int64) tg.PeerClazz {
	if sourcePeerType != payload.PeerTypeUser {
		if peer := peerFromOptional(sourcePeerType, sourcePeerID); peer != nil {
			return peer
		}
	}
	if fromUserID > 0 {
		return peerFromUser(fromUserID)
	}
	if peer := peerFromOptional(sourcePeerType, sourcePeerID); peer != nil {
		return peer
	}
	return nil
}

func peerFromOptional(peerType int32, peerID int64) tg.PeerClazz {
	if peerID == 0 {
		return nil
	}
	return peerFromEvent(peerType, peerID)
}

func stringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func replyHeaderFromUserMessageID(userMessageID int64) (tg.MessageReplyHeaderClazz, error) {
	if userMessageID <= 0 {
		return nil, nil
	}
	replyToMsgID, err := messageIDInt32(userMessageID, "reply user message id")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}

func messageIDInt32(v int64, field string) (int32, error) {
	if v <= 0 {
		return 0, fmt.Errorf("%w: missing public %s", userupdates.ErrUserupdatesStorage, field)
	}
	return int64ToInt32(v, field)
}

func int64ToInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %s out of int32 range", userupdates.ErrOperationTerminal, field)
	}
	return int32(v), nil
}

func userupdatesDateInt32FromUnixSeconds(seconds int64, field string) (int32, error) {
	date, err := tg.DateInt32FromUnixSeconds(seconds)
	if err != nil {
		return 0, fmt.Errorf("%w: convert %s: %v", userupdates.ErrUserupdatesStorage, field, err)
	}
	return date, nil
}

func peerFromUser(userID int64) tg.PeerClazz {
	if userID == 0 {
		return nil
	}
	return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: userID})
}

func messageFromPeer(out bool, peerType int32, fromUserID int64) tg.PeerClazz {
	if !out && peerType == payload.PeerTypeUser {
		return nil
	}
	return peerFromUser(fromUserID)
}

func peerFromEvent(peerType int32, peerID int64) tg.PeerClazz {
	switch peerType {
	case payload.PeerTypeUser:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	case payload.PeerTypeChat:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peerID})
	case payload.PeerTypeChannel:
		return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: peerID})
	default:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	}
}

func dialogEventFromMessageEvent(event eventtypes.UserEvent, messageEvent decodedMessageEvent) payload.DialogEventV1 {
	return payload.DialogEventV1{
		SchemaVersion:    payload.DialogEventSchemaVersion,
		EventKind:        messageEvent.EventKind,
		PublicUpdateType: messageEvent.EventKind,
		PeerType:         event.PeerType,
		PeerID:           event.PeerID,
	}
}

func dialogEventToTLUpdate(event payload.DialogEventV1, pts int64, ptsCount int32) (tg.UpdateClazz, error) {
	peer := peerFromEvent(event.PeerType, event.PeerID)
	dialogPeer := tg.MakeTLDialogPeer(&tg.TLDialogPeer{Peer: peer})
	switch event.EventKind {
	case payload.DialogEventDraftSaved:
		return tg.MakeTLUpdateDraftMessage(&tg.TLUpdateDraftMessage{Peer: peer, Draft: tg.MakeTLDraftMessage(&tg.TLDraftMessage{})}), nil
	case payload.DialogEventDraftCleared, payload.DialogEventDraftClearedAfterSend:
		return tg.MakeTLUpdateDraftMessage(&tg.TLUpdateDraftMessage{Peer: peer, Draft: tg.MakeTLDraftMessageEmpty(&tg.TLDraftMessageEmpty{})}), nil
	case payload.DialogEventPinToggled:
		pinned := true
		if event.Pinned != nil {
			pinned = *event.Pinned
		}
		return tg.MakeTLUpdateDialogPinned(&tg.TLUpdateDialogPinned{Pinned: pinned, FolderId: event.FolderID, Peer: dialogPeer}), nil
	case payload.DialogEventPinnedDialogsReordered:
		return tg.MakeTLUpdatePinnedDialogs(&tg.TLUpdatePinnedDialogs{FolderId: event.FolderID, Order: []tg.DialogPeerClazz{dialogPeer}}), nil
	case payload.DialogEventFolderPeersChanged:
		pts32, err := int64ToInt32(pts, "pts")
		if err != nil {
			return nil, err
		}
		folderID := int32(0)
		if event.FolderID != nil {
			folderID = *event.FolderID
		}
		return tg.MakeTLUpdateFolderPeers(&tg.TLUpdateFolderPeers{FolderPeers: []tg.FolderPeerClazz{tg.MakeTLFolderPeer(&tg.TLFolderPeer{Peer: peer, FolderId: folderID})}, Pts: pts32, PtsCount: ptsCount}), nil
	case payload.DialogEventFilterUpdated, payload.DialogEventFilterDeleted:
		return tg.MakeTLUpdateDialogFilter(&tg.TLUpdateDialogFilter{}), nil
	case payload.DialogEventFiltersOrderUpdated:
		return tg.MakeTLUpdateDialogFilterOrder(&tg.TLUpdateDialogFilterOrder{Order: []int32{}}), nil
	case payload.DialogEventWallpaperChanged:
		return tg.MakeTLUpdatePeerWallpaper(&tg.TLUpdatePeerWallpaper{Peer: peer}), nil
	case payload.DialogEventPrivatePeerHistoryTTL:
		return tg.MakeTLUpdatePeerHistoryTTL(&tg.TLUpdatePeerHistoryTTL{Peer: peer, TtlPeriod: event.TTLPeriod}), nil
	case payload.DialogEventSavedDialogPinned:
		pinned := true
		if event.Pinned != nil {
			pinned = *event.Pinned
		}
		return tg.MakeTLUpdateSavedDialogPinned(&tg.TLUpdateSavedDialogPinned{Pinned: pinned, Peer: dialogPeer}), nil
	case payload.DialogEventPinnedSavedDialogsChanged:
		return tg.MakeTLUpdatePinnedSavedDialogs(&tg.TLUpdatePinnedSavedDialogs{Order: []tg.DialogPeerClazz{dialogPeer}}), nil
	default:
		return nil, fmt.Errorf("%w: unsupported dialog event kind=%s", userupdates.ErrUserupdatesStorage, event.EventKind)
	}
}
