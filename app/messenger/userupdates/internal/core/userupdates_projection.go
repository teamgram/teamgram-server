package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/envelope"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/projection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func applyResultToTL(in *repository.ApplyUserOperationResult) (*userupdates.UserOperationResult, error) {
	if in == nil {
		return nil, userupdates.ErrOperationTerminal
	}
	schemaVersion := in.ResponseSchemaVersion
	if schemaVersion == 0 {
		schemaVersion = payload.OperationResponseSchemaVersion
	}
	return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		UserId:                in.UserID,
		OperationId:           in.OperationID,
		Status:                repository.OperationResultStatusCompleted,
		Pts:                   in.Pts,
		PtsCount:              in.PtsCount,
		CurrentPts:            in.Pts,
		ResponseSchemaVersion: &schemaVersion,
		ResponsePayload:       in.ResponsePayload,
		ResponsePayloadHash:   in.ResponseHash,
	}).ToUserOperationResult(), nil
}

func operationResultToTL(in *repository.OperationResult) (*userupdates.UserOperationResult, error) {
	if in == nil {
		return nil, userupdates.ErrOperationTerminal
	}
	schemaVersion, err := operationResponseSchemaVersion(in.ResponseSchemaVersion, in.ResponsePayload)
	if err != nil {
		return nil, err
	}
	return userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		UserId:                in.UserID,
		OperationId:           in.OperationID,
		Status:                in.Status,
		Pts:                   in.Pts,
		PtsCount:              in.PtsCount,
		CurrentPts:            in.Pts,
		ResponseSchemaVersion: schemaVersion,
		ResponsePayload:       in.ResponsePayload,
		ResponsePayloadHash:   in.ResponseHash,
	}).ToUserOperationResult(), nil
}

func operationResponseSchemaVersion(storedSchemaVersion int32, responsePayload []byte) (*int32, error) {
	if len(responsePayload) == 0 {
		return nil, nil
	}
	schemaVersion := storedSchemaVersion
	if schemaVersion == 0 {
		var header struct {
			SchemaVersion int32 `json:"schema_version"`
		}
		if err := json.Unmarshal(responsePayload, &header); err != nil {
			return nil, fmt.Errorf("%w: decode operation response schema: %v", userupdates.ErrUserupdatesStorage, err)
		}
		schemaVersion = header.SchemaVersion
	}
	switch schemaVersion {
	case payload.OperationResponseSchemaVersionV1:
		var response payload.OperationResponseV1
		if err := json.Unmarshal(responsePayload, &response); err != nil {
			return nil, fmt.Errorf("%w: decode legacy operation response: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if response.SchemaVersion != payload.OperationResponseSchemaVersionV1 {
			return nil, fmt.Errorf("%w: operation response schema mismatch stored=%d payload=%d", userupdates.ErrUserupdatesStorage, schemaVersion, response.SchemaVersion)
		}
	case payload.OperationResponseSchemaVersion:
		var response payload.OperationResponseV2
		if err := json.Unmarshal(responsePayload, &response); err != nil {
			return nil, fmt.Errorf("%w: decode operation response: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if response.SchemaVersion != payload.OperationResponseSchemaVersion {
			return nil, fmt.Errorf("%w: operation response schema mismatch stored=%d payload=%d", userupdates.ErrUserupdatesStorage, schemaVersion, response.SchemaVersion)
		}
	case payload.OperationResponseSchemaVersionV3:
		var response payload.OperationResponseV3
		if err := json.Unmarshal(responsePayload, &response); err != nil {
			return nil, fmt.Errorf("%w: decode v3 operation response: %v", userupdates.ErrUserupdatesStorage, err)
		}
		if response.SchemaVersion != payload.OperationResponseSchemaVersionV3 {
			return nil, fmt.Errorf("%w: operation response schema mismatch stored=%d payload=%d", userupdates.ErrUserupdatesStorage, schemaVersion, response.SchemaVersion)
		}
	default:
		return nil, fmt.Errorf("%w: unsupported operation response schema=%d", userupdates.ErrUserupdatesStorage, schemaVersion)
	}
	v := schemaVersion
	return &v, nil
}

func stateToTL(in repository.UserState) *userupdates.UserState {
	return userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts:         in.Pts,
		Seq:         int32(in.Seq),
		Date:        in.Date,
		UnreadCount: in.UnreadCount,
	}).ToUserState()
}

func differenceToTL(in *repository.GetDifferenceResult) (*userupdates.UserDifference, error) {
	if in == nil {
		return userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
			State: stateToTL(repository.UserState{}),
		}).ToUserDifference(), nil
	}
	state := stateToTL(in.State)
	if len(in.Events) == 0 && len(in.AuthSeqEvents) == 0 {
		return userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
			State: state,
		}).ToUserDifference(), nil
	}

	newMessages := make([]tg.MessageClazz, 0, len(in.Events))
	otherUpdates := make([]tg.UpdateClazz, 0, len(in.Events)+len(in.AuthSeqEvents))
	for _, event := range in.Events {
		projected, err := projection.ProjectUserEvent(event, projection.ModeDifference)
		if err != nil {
			return nil, err
		}
		if projected.Message != nil {
			newMessages = append(newMessages, projected.Message)
		}
		if len(projected.OtherUpdates) > 0 {
			otherUpdates = append(otherUpdates, projected.OtherUpdates...)
		} else if projected.Update != nil {
			otherUpdates = append(otherUpdates, projected.Update)
		}
	}
	for _, event := range in.AuthSeqEvents {
		update, err := authSeqEventToTLUpdate(event)
		if err != nil {
			return nil, err
		}
		otherUpdates = append(otherUpdates, update)
	}
	return userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
		NewMessages:  newMessages,
		OtherUpdates: otherUpdates,
		State:        state,
	}).ToUserDifference(), nil
}

func messageViewToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	if view.MessageStatus != repository.MessageStatusLive {
		return nil, nil
	}
	switch view.ViewSchemaVersion {
	case payload.MessageEventSchemaVersionV1:
		return legacyMessageViewToTLMessage(view)
	case payload.MessageEventSchemaVersion:
		return currentMessageViewToTLMessage(view)
	case payload.MessageEventSchemaVersionV3:
		return messageViewV3ToTLMessage(view)
	case payload.MessageEventSchemaVersionV4:
		return messageViewV4ToTLMessage(view)
	default:
		return nil, fmt.Errorf("%w: unsupported message view schema=%d", userupdates.ErrUserupdatesStorage, view.ViewSchemaVersion)
	}
}

func legacyMessageViewToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	var messageEvent payload.MessageEventV1
	if err := json.Unmarshal(view.ViewPayload, &messageEvent); err != nil {
		return nil, fmt.Errorf("%w: decode legacy message view payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersionV1 {
		return nil, fmt.Errorf("%w: unsupported legacy message view event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.PeerType != view.PeerType || messageEvent.PeerID != view.PeerID || messageEvent.MessageID != view.PeerSeq {
		return nil, fmt.Errorf("%w: legacy message view payload mismatch", userupdates.ErrUserupdatesStorage)
	}
	if messageEvent.EventKind == payload.OperationKindEditMessage {
		return legacyEditMessageEventToTLMessage(messageEvent)
	}
	return legacyMessageEventToTLMessage(messageEvent)
}

func currentMessageViewToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	var messageEvent payload.MessageEventV2
	if err := json.Unmarshal(view.ViewPayload, &messageEvent); err != nil {
		return nil, fmt.Errorf("%w: decode message view payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersion {
		return nil, fmt.Errorf("%w: unsupported message view event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.PeerType != view.PeerType ||
		messageEvent.PeerID != view.PeerID ||
		messageEvent.PeerSeq != view.PeerSeq ||
		messageEvent.MessageID != view.UserMessageID {
		return nil, fmt.Errorf("%w: message view payload mismatch", userupdates.ErrUserupdatesStorage)
	}
	if messageEvent.EventKind == payload.OperationKindEditMessage {
		return currentEditMessageEventToTLMessage(messageEvent)
	}
	return currentMessageEventToTLMessage(messageEvent)
}

func messageViewV3ToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	var messageEvent payload.MessageEventV3
	if err := json.Unmarshal(view.ViewPayload, &messageEvent); err != nil {
		return nil, fmt.Errorf("%w: decode v3 message view payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersionV3 {
		return nil, fmt.Errorf("%w: unsupported v3 message view event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.PeerType != view.PeerType ||
		messageEvent.PeerID != view.PeerID ||
		messageEvent.PeerSeq != view.PeerSeq ||
		messageEvent.MessageID != view.UserMessageID {
		return nil, fmt.Errorf("%w: v3 message view payload mismatch", userupdates.ErrUserupdatesStorage)
	}
	if messageEvent.EventKind == payload.OperationKindEditMessage {
		return messageEventV3EditToTLMessage(messageEvent)
	}
	return messageEventV3ToTLMessage(messageEvent)
}

func messageViewV4ToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	var messageEvent payload.MessageEventV4
	if err := json.Unmarshal(view.ViewPayload, &messageEvent); err != nil {
		return nil, fmt.Errorf("%w: decode v4 message view payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersionV4 {
		return nil, fmt.Errorf("%w: unsupported v4 message view event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.MessageFact.PeerType != view.PeerType ||
		messageEvent.MessageFact.PeerID != view.PeerID ||
		messageEvent.MessageFact.PeerSeq != view.PeerSeq ||
		messageEvent.MessageID != view.UserMessageID ||
		messageEvent.MessageFact.CanonicalMessageID != view.CanonicalMessageID {
		return nil, fmt.Errorf("%w: v4 message view payload mismatch", userupdates.ErrUserupdatesStorage)
	}
	fact, err := payload.WrapFact(payload.FactKindNewMessage, messageEvent.MessageFact)
	if err != nil {
		return nil, fmt.Errorf("%w: wrap v4 message view fact: %v", userupdates.ErrUserupdatesStorage, err)
	}
	projected, err := projection.ProjectFacts([]payload.UpdateFactV1{fact}, projection.ViewerContext{
		UserID:           view.UserID,
		AuthKeyIDExclude: messageEvent.AuthKeyIdExclude,
	}, envelope.ModeDifference, 0, 0, view.UserMessageID)
	if err != nil {
		return nil, err
	}
	for _, update := range projected {
		newMessage, ok := update.Update.(*tg.TLUpdateNewMessage)
		if ok {
			return newMessage.Message, nil
		}
	}
	return nil, fmt.Errorf("%w: v4 message view did not project a new message", userupdates.ErrUserupdatesStorage)
}

func legacyMessageEventToTLMessage(messageEvent payload.MessageEventV1) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
	if err != nil {
		return nil, err
	}
	replyTo, err := replyHeaderFromPeerSeq(messageEvent.ReplyToPeerSeq)
	if err != nil {
		return nil, err
	}
	date, err := userupdatesDateInt32FromUnixSeconds(int64(messageEvent.Date), "message date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      messageEvent.Out,
		Id:       messageID,
		FromId:   messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:   peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		ReplyTo:  replyTo,
		Date:     date,
		Message:  messageEvent.MessageText,
		Entities: projectionMessageEntities(messageEvent.Entities),
	}), nil
}

func legacyEditMessageEventToTLMessage(messageEvent payload.MessageEventV1) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.OperationKindEditMessage {
		return nil, fmt.Errorf("%w: unsupported edit event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
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
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      messageEvent.Out,
		Id:       messageID,
		FromId:   messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:   peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		Date:     date,
		Message:  messageEvent.MessageText,
		Entities: projectionMessageEntities(messageEvent.Entities),
		EditDate: &editDate32,
	}), nil
}

func currentMessageEventToTLMessage(messageEvent payload.MessageEventV2) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
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
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      messageEvent.Out,
		Id:       messageID,
		FromId:   messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:   peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		ReplyTo:  replyTo,
		Date:     date,
		Message:  messageEvent.MessageText,
		Entities: projectionMessageEntities(messageEvent.Entities),
	}), nil
}

func currentEditMessageEventToTLMessage(messageEvent payload.MessageEventV2) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.OperationKindEditMessage {
		return nil, fmt.Errorf("%w: unsupported edit event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
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
	return tg.MakeTLMessage(&tg.TLMessage{
		Out:      messageEvent.Out,
		Id:       messageID,
		FromId:   messageFromPeer(messageEvent.Out, messageEvent.PeerType, messageEvent.FromUserID),
		PeerId:   peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		Date:     date,
		Message:  messageEvent.MessageText,
		Entities: projectionMessageEntities(messageEvent.Entities),
		EditDate: &editDate32,
	}), nil
}

func messageEventV3ToTLMessage(messageEvent payload.MessageEventV3) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
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
		Entities:    projectionMessageEntities(messageEvent.Entities),
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

func messageEventV3EditToTLMessage(messageEvent payload.MessageEventV3) (tg.MessageClazz, error) {
	if messageEvent.EventKind != payload.OperationKindEditMessage {
		return nil, fmt.Errorf("%w: unsupported edit event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
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
		Entities:    projectionMessageEntities(messageEvent.Entities),
		GroupedId:   messageGroupedID(messageEvent.Attrs),
		TtlPeriod:   messageTTLPeriod(messageEvent.MediaRef),
		EditDate:    &editDate32,
	}), nil
}

func projectionMessageEntities(entities []payload.MessageEntityV1) []tg.MessageEntityClazz {
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
			Photo:      viewMessagePhoto(media),
			TtlSeconds: ttl,
		})
	case "document":
		flags := viewDocumentMediaFlags(media)
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Spoiler:        flags.Spoiler,
			Video:          flags.Video,
			Round:          flags.Round,
			Voice:          flags.Voice,
			Document:       viewMessageDocument(media),
			VideoCover:     viewPhotoRef(media.VideoCover),
			VideoTimestamp: cloneInt32Ptr(media.VideoTimestamp),
			TtlSeconds:     ttl,
		})
	case "contact":
		return viewMessageContact(media)
	default:
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{})
	}
}

func viewMessageContact(media *payload.MediaRefV1) tg.MessageMediaClazz {
	return tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
		PhoneNumber: media.PhoneNumber,
		FirstName:   media.FirstName,
		LastName:    media.LastName,
		Vcard:       media.Vcard,
		UserId:      media.UserID,
	})
}

func viewMessagePhoto(media *payload.MediaRefV1) tg.PhotoClazz {
	if media.Date == 0 && media.DcID == 0 && len(media.PhotoSizes) == 0 {
		return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{Id: media.ID})
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		Id:            media.ID,
		AccessHash:    media.AccessHash,
		FileReference: append([]byte(nil), media.FileReference...),
		Date:          media.Date,
		Sizes:         viewPhotoSizes(media.PhotoSizes),
		DcId:          media.DcID,
	})
}

func viewMessageDocument(media *payload.MediaRefV1) tg.DocumentClazz {
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
		Thumbs:        viewPhotoSizes(media.DocumentThumbs),
		VideoThumbs:   viewVideoSizes(media.DocumentVideoThumbs),
		DcId:          media.DcID,
		Attributes:    viewDocumentAttributes(media.DocumentAttributes),
	})
}

func viewPhotoRef(photo *payload.PhotoRefV1) tg.PhotoClazz {
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
		Sizes:         viewPhotoSizes(photo.Sizes),
		VideoSizes:    viewVideoSizes(photo.VideoSizes),
		DcId:          photo.DcID,
	})
}

func viewPhotoSizes(sizes []payload.PhotoSizeRefV1) []tg.PhotoSizeClazz {
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

func viewVideoSizes(sizes []payload.VideoSizeRefV1) []tg.VideoSizeClazz {
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
				Stickerset:       viewStickerSetRef(size.StickerSet),
				StickerId:        size.StickerID,
				BackgroundColors: append([]int32(nil), size.BackgroundColors...),
			}))
		}
	}
	return out
}

func viewDocumentAttributes(attrs []payload.DocumentAttributeRefV1) []tg.DocumentAttributeClazz {
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
				Stickerset: viewStickerSetRef(attr.StickerSet),
				MaskCoords: viewMaskCoordsRef(attr.MaskCoords),
			}))
		case "custom_emoji":
			out = append(out, tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{
				Free:       attr.Free,
				TextColor:  attr.TextColor,
				Alt:        attr.Alt,
				Stickerset: viewStickerSetRef(attr.StickerSet),
			}))
		case "has_stickers":
			out = append(out, tg.MakeTLDocumentAttributeHasStickers(&tg.TLDocumentAttributeHasStickers{}))
		}
	}
	return out
}

func viewDocumentMediaFlags(media *payload.MediaRefV1) payload.DocumentMediaFlagsV1 {
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
			if !viewIsWebMStickerOrCustomEmoji(media) {
				flags.Video = true
			}
		}
	}
	return flags
}

func viewIsWebMStickerOrCustomEmoji(media *payload.MediaRefV1) bool {
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

func viewStickerSetRef(ref *payload.StickerSetRefV1) tg.InputStickerSetClazz {
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

func viewMaskCoordsRef(ref *payload.MaskCoordsRefV1) tg.MaskCoordsClazz {
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

func replyHeaderFromPeerSeq(peerSeq int64) (tg.MessageReplyHeaderClazz, error) {
	if peerSeq <= 0 {
		return nil, nil
	}
	replyToMsgID, err := int64ToInt32(peerSeq, "reply peer seq")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}

func replyHeaderFromUserMessageID(userMessageID int64) (tg.MessageReplyHeaderClazz, error) {
	if userMessageID <= 0 {
		return nil, nil
	}
	replyToMsgID, err := int64ToInt32(userMessageID, "reply user message id")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{ReplyToMsgId: &replyToMsgID}), nil
}

func authSeqEventToTLUpdate(event repository.AuthSeqEvent) (tg.UpdateClazz, error) {
	if event.EventCodec != repository.PayloadCodecJSON || event.EventSchemaVersion <= 0 {
		return nil, fmt.Errorf("%w: unsupported auth seq event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if len(event.EventPayloadHash) != 0 && !bytes.Equal(payload.HashBytes(event.EventPayload), event.EventPayloadHash) {
		return nil, fmt.Errorf("%w: auth seq event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}
	dialogEvent := payload.DialogEventV1{
		SchemaVersion:    payload.DialogEventSchemaVersion,
		EventKind:        event.PublicUpdateType,
		PublicUpdateType: event.PublicUpdateType,
		PeerType:         event.PeerType,
		PeerID:           event.PeerID,
	}
	var decoded payload.DialogEventV1
	if len(event.EventPayload) != 0 && json.Unmarshal(event.EventPayload, &decoded) == nil && decoded.SchemaVersion != 0 {
		if decoded.EventKind != "" {
			dialogEvent.EventKind = decoded.EventKind
		}
		if decoded.PublicUpdateType != "" {
			dialogEvent.PublicUpdateType = decoded.PublicUpdateType
		}
		if decoded.PeerType != 0 {
			dialogEvent.PeerType = decoded.PeerType
		}
		if decoded.PeerID != 0 {
			dialogEvent.PeerID = decoded.PeerID
		}
		dialogEvent.FolderID = decoded.FolderID
		dialogEvent.Pinned = decoded.Pinned
		dialogEvent.TTLPeriod = decoded.TTLPeriod
	}
	return dialogEventToTLUpdate(dialogEvent, 0, 0)
}

func dialogEventFromMessageEvent(event repository.UserEvent, messageEvent payload.MessageEventV1) payload.DialogEventV1 {
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
