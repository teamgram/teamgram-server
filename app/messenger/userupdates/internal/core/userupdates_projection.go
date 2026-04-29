package core

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func operationHashHex(hash []byte) string {
	return hex.EncodeToString(hash)
}

func hexPayload(hash string) ([]byte, error) {
	if hash == "" {
		return nil, nil
	}
	b, err := hex.DecodeString(hash)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid repository payload hash", userupdates.ErrUserupdatesStorage)
	}
	return b, nil
}

func applyResultToTL(in *repository.ApplyUserOperationResult) (*userupdates.UserOperationResult, error) {
	if in == nil {
		return nil, userupdates.ErrOperationTerminal
	}
	hash, err := hexPayload(in.ResponseHash)
	if err != nil {
		return nil, err
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
		ResponsePayloadHash:   hash,
	}).ToUserOperationResult(), nil
}

func operationResultToTL(in *repository.OperationResult) (*userupdates.UserOperationResult, error) {
	if in == nil {
		return nil, userupdates.ErrOperationTerminal
	}
	hash, err := hexPayload(in.ResponseHash)
	if err != nil {
		return nil, err
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
		ResponsePayloadHash:   hash,
	}).ToUserOperationResult(), nil
}

func stateToTL(in repository.UserState) *userupdates.UserState {
	return userupdates.MakeTLUserState(&userupdates.TLUserState{
		Pts: in.Pts,
	}).ToUserState()
}

func differenceToTL(in *repository.GetDifferenceResult) (*userupdates.UserDifference, error) {
	if in == nil {
		return userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
			State: stateToTL(repository.UserState{}),
		}).ToUserDifference(), nil
	}
	state := stateToTL(in.State)
	if len(in.Events) == 0 {
		return userupdates.MakeTLUserDifferenceEmpty(&userupdates.TLUserDifferenceEmpty{
			State: state,
		}).ToUserDifference(), nil
	}

	newMessages := make([]tg.MessageClazz, 0, len(in.Events))
	otherUpdates := make([]tg.UpdateClazz, 0, len(in.Events))
	for _, event := range in.Events {
		message, update, err := eventToTLUpdate(event)
		if err != nil {
			return nil, err
		}
		newMessages = append(newMessages, message)
		otherUpdates = append(otherUpdates, update)
	}
	return userupdates.MakeTLUserDifference(&userupdates.TLUserDifference{
		NewMessages:  newMessages,
		OtherUpdates: otherUpdates,
		State:        state,
	}).ToUserDifference(), nil
}

func eventToTLUpdate(event repository.UserEvent) (*tg.TLMessage, *tg.TLUpdateNewMessage, error) {
	if event.EventCodec != repository.PayloadCodecJSON || event.EventSchemaVersion != payload.MessageEventSchemaVersion {
		return nil, nil, fmt.Errorf("%w: unsupported event codec=%d schema=%d", userupdates.ErrUserupdatesStorage, event.EventCodec, event.EventSchemaVersion)
	}
	if event.EventPayloadHash != "" && payload.HashBytes(event.EventPayload) != event.EventPayloadHash {
		return nil, nil, fmt.Errorf("%w: event payload hash mismatch", userupdates.ErrUserupdatesStorage)
	}

	var messageEvent payload.MessageEventV1
	if err := json.Unmarshal(event.EventPayload, &messageEvent); err != nil {
		return nil, nil, fmt.Errorf("%w: decode event payload: %v", userupdates.ErrUserupdatesStorage, err)
	}
	if messageEvent.SchemaVersion != payload.MessageEventSchemaVersion || messageEvent.EventKind != payload.EventKindNewMessage {
		return nil, nil, fmt.Errorf("%w: unsupported event kind=%s schema=%d", userupdates.ErrUserupdatesStorage, messageEvent.EventKind, messageEvent.SchemaVersion)
	}

	messageID, err := int64ToInt32(messageEvent.MessageID, "message id")
	if err != nil {
		return nil, nil, err
	}
	pts, err := int64ToInt32(event.Pts, "pts")
	if err != nil {
		return nil, nil, err
	}
	message := tg.MakeTLMessage(&tg.TLMessage{
		Out:     messageEvent.Out,
		Id:      messageID,
		FromId:  peerFromUser(messageEvent.FromUserID),
		PeerId:  peerFromEvent(messageEvent.PeerType, messageEvent.PeerID),
		Date:    messageEvent.Date,
		Message: messageEvent.MessageText,
	})
	update := tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
		Message:  message,
		Pts:      pts,
		PtsCount: event.PtsCount,
	})
	return message, update, nil
}

func int64ToInt32(v int64, field string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("%w: %s out of int32 range", userupdates.ErrOperationTerminal, field)
	}
	return int32(v), nil
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
