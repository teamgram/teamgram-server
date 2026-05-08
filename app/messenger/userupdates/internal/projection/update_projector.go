package projection

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

type Mode int

const (
	ModeDifference Mode = 1
	ModePush       Mode = 2
)

type Result struct {
	Message          tg.MessageClazz
	Update           tg.UpdateClazz
	Updates          tg.UpdatesClazz
	AuthKeyIDExclude *int64
}

func ProjectUserEvent(event repository.UserEvent, mode Mode) (Result, error) {
	messageEvent, err := decodeUserEventPayload(event)
	if err != nil {
		return Result{}, err
	}
	if event.EventType == repository.EventTypeDialogPublicUpdate {
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

func ProjectPushTask(msg *payload.PushTaskKafkaMessageV1) (Result, error) {
	if msg == nil {
		return Result{}, fmt.Errorf("%w: push task is nil", userupdates.ErrUserupdatesStorage)
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
	PinnedUserMessageID  int64
	AuthKeyIdExclude     *int64
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

func decodeUserEventPayload(event repository.UserEvent) (decodedMessageEvent, error) {
	if event.EventCodec != repository.PayloadCodecJSON {
		return decodedMessageEvent{}, fmt.Errorf("%w: unsupported event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if len(event.EventPayloadHash) != 0 && !bytes.Equal(payload.HashBytes(event.EventPayload), event.EventPayloadHash) {
		return decodedMessageEvent{}, fmt.Errorf("%w: event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}
	return decodeMessageEventPayloadBytes(event.EventSchemaVersion, event.EventPayload)
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
			PinnedUserMessageID:  next.PinnedUserMessageID,
			AuthKeyIdExclude:     next.AuthKeyIdExclude,
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
		date, err := userupdatesDateInt32FromUnixSeconds(int64(in.message.Date), in.datePrefix+" message date")
		if err != nil {
			return Result{}, err
		}
		replyTo, err := replyHeaderFromUserMessageID(in.message.ReplyToUserMessageID)
		if err != nil {
			return Result{}, err
		}
		if in.message.PeerType == payload.PeerTypeUser {
			return Result{
				Updates: tg.MakeTLUpdateShortMessage(&tg.TLUpdateShortMessage{
					Out:      in.message.Out,
					Id:       message.(*tg.TLMessage).Id,
					UserId:   shortMessageUserID(in.message),
					Message:  in.message.MessageText,
					Pts:      pts,
					PtsCount: in.ptsCount,
					Date:     date,
					ReplyTo:  replyTo,
				}),
				AuthKeyIDExclude: in.message.AuthKeyIdExclude,
			}, nil
		}
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
	date, err := userupdatesDateInt32FromUnixSeconds(int64(dateSeconds), "updates date")
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{update},
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

func shortMessageUserID(event decodedMessageEvent) int64 {
	if event.Out {
		return event.PeerID
	}
	return event.FromUserID
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

func dialogEventFromMessageEvent(event repository.UserEvent, messageEvent decodedMessageEvent) payload.DialogEventV1 {
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
