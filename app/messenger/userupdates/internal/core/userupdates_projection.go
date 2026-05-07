package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func applyResultToTL(in *repository.ApplyUserOperationResult) (*userupdates.UserOperationResult, error) {
	if in == nil {
		return nil, userupdates.ErrOperationTerminal
	}
	schemaVersion := int32(payload.OperationResponseSchemaVersion)
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
	var schemaVersion *int32
	if len(in.ResponsePayload) != 0 {
		v := int32(payload.OperationResponseSchemaVersion)
		schemaVersion = &v
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

func stateToTL(in repository.UserState) *userupdates.UserState {
	return userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts:  in.Pts,
		Seq:  int32(in.Seq),
		Date: in.Date,
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
		message, update, err := eventToTLUpdate(event)
		if err != nil {
			return nil, err
		}
		if message != nil {
			newMessages = append(newMessages, message)
		}
		otherUpdates = append(otherUpdates, update)
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

func eventToTLUpdate(event repository.UserEvent) (tg.MessageClazz, tg.UpdateClazz, error) {
	if event.EventCodec != repository.PayloadCodecJSON || event.EventSchemaVersion != payload.MessageEventSchemaVersion {
		return nil, nil, fmt.Errorf("%w: unsupported event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if len(event.EventPayloadHash) != 0 && !bytes.Equal(payload.HashBytes(event.EventPayload), event.EventPayloadHash) {
		return nil, nil, fmt.Errorf("%w: event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}

	var messageEvent payload.MessageEventV1
	if err := json.Unmarshal(event.EventPayload, &messageEvent); err != nil {
		return nil, nil, fmt.Errorf("%w: decode event payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersion {
		return nil, nil, fmt.Errorf("%w: unsupported event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if event.EventType == repository.EventTypeDialogPublicUpdate {
		update, err := dialogEventToTLUpdate(dialogEventFromMessageEvent(event, messageEvent), event.Pts, event.PtsCount)
		return nil, update, err
	}
	if messageEvent.EventKind == payload.OperationKindReadHistory {
		update, err := readHistoryEventToTLUpdate(event, messageEvent)
		return nil, update, err
	}
	if messageEvent.EventKind == payload.OperationKindEditMessage {
		update, err := editMessageEventToTLUpdate(event, messageEvent)
		return nil, update, err
	}
	if messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}
	message, err := messageEventToTLMessage(messageEvent)
	if err != nil {
		return nil, nil, err
	}
	pts, err := int64ToInt32(event.Pts, "pts")
	if err != nil {
		return nil, nil, err
	}
	update := tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
		Message:  message,
		Pts:      pts,
		PtsCount: event.PtsCount,
	})
	return message, update, nil
}

func readHistoryEventToTLUpdate(event repository.UserEvent, messageEvent payload.MessageEventV1) (tg.UpdateClazz, error) {
	maxID, err := int64ToInt32(messageEvent.MessageID, "message id")
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(event.Pts, "pts")
	if err != nil {
		return nil, err
	}
	peer := peerFromEvent(messageEvent.PeerType, messageEvent.PeerID)
	if messageEvent.Out {
		return tg.MakeTLUpdateReadHistoryOutbox(&tg.TLUpdateReadHistoryOutbox{
			Peer:     peer,
			MaxId:    maxID,
			Pts:      pts,
			PtsCount: event.PtsCount,
		}), nil
	}
	return tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{
		Peer:             peer,
		MaxId:            maxID,
		StillUnreadCount: 0,
		Pts:              pts,
		PtsCount:         event.PtsCount,
	}), nil
}

func editMessageEventToTLUpdate(event repository.UserEvent, messageEvent payload.MessageEventV1) (tg.UpdateClazz, error) {
	message, err := editMessageEventToTLMessage(messageEvent)
	if err != nil {
		return nil, err
	}
	pts, err := int64ToInt32(event.Pts, "pts")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
		Message:  message,
		Pts:      pts,
		PtsCount: event.PtsCount,
	}), nil
}

func messageViewToTLMessage(view repository.MessageView) (tg.MessageClazz, error) {
	if view.MessageStatus != repository.MessageStatusLive {
		return nil, nil
	}
	if view.ViewSchemaVersion != payload.MessageEventSchemaVersion {
		return nil, fmt.Errorf("%w: unsupported message view schema=%d", userupdates.ErrUserupdatesStorage, view.ViewSchemaVersion)
	}
	var messageEvent payload.MessageEventV1
	if err := json.Unmarshal(view.ViewPayload, &messageEvent); err != nil {
		return nil, fmt.Errorf("%w: decode message view payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersion {
		return nil, fmt.Errorf("%w: unsupported message view event schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.SchemaVersion)
	}
	if messageEvent.PeerType != view.PeerType || messageEvent.PeerID != view.PeerID || messageEvent.MessageID != view.PeerSeq {
		return nil, fmt.Errorf("%w: message view payload mismatch", userupdates.ErrUserupdatesStorage)
	}
	if messageEvent.EventKind == payload.OperationKindEditMessage {
		return editMessageEventToTLMessage(messageEvent)
	}
	return messageEventToTLMessage(messageEvent)
}

func messageEventToTLMessage(messageEvent payload.MessageEventV1) (tg.MessageClazz, error) {
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
		Out:     messageEvent.Out,
		Id:      messageID,
		FromId:  peerFromUser(messageEvent.FromUserID),
		PeerId:  peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		ReplyTo: replyTo,
		Date:    date,
		Message: messageEvent.MessageText,
	}), nil
}

func editMessageEventToTLMessage(messageEvent payload.MessageEventV1) (tg.MessageClazz, error) {
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
		FromId:   peerFromUser(messageEvent.FromUserID),
		PeerId:   peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		Date:     date,
		Message:  messageEvent.MessageText,
		EditDate: &editDate32,
	}), nil
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
