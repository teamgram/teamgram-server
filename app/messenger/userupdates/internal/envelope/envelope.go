package envelope

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type Mode int

const (
	ModeReply Mode = iota + 1
	ModeSenderStream
	ModeReceiverStream
	ModeDifference
	ModeQueryResponse
	ModeEphemeralPush
)

type MessageIDMapping struct {
	MessageID int32
	RandomID  int64
}

type Input struct {
	Mode            Mode
	Updates         []tg.UpdateClazz
	Users           []tg.UserClazz
	Chats           []tg.ChatClazz
	Date            int32
	Seq             int32
	MessageIDByID   map[int32]int64
	AllowShortReply bool
}

type PeerObjectProjector interface {
	ProjectUsers(ctx context.Context, viewerUserID int64, ids []int64) ([]tg.UserClazz, error)
	ProjectChats(ctx context.Context, viewerUserID int64, ids []int64) ([]tg.ChatClazz, error)
}

func BuildUpdates(in Input) (*tg.Updates, error) {
	updates := in.Updates

	if in.Mode == ModeReply {
		var err error
		updates, err = buildReplyUpdates(in)
		if err != nil {
			return nil, err
		}
	} else {
		for _, update := range updates {
			if updateMessageID, ok := update.(*tg.TLUpdateMessageID); ok {
				if updateMessageID == nil {
					return nil, fmt.Errorf("envelope: nil updateMessageID")
				}
				return nil, fmt.Errorf("envelope: updateMessageID is reply-only in mode %d", in.Mode)
			}
		}
	}

	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: updates,
		Users:   in.Users,
		Chats:   in.Chats,
		Date:    in.Date,
		Seq:     in.Seq,
	}).ToUpdates(), nil
}

func buildReplyUpdates(in Input) ([]tg.UpdateClazz, error) {
	newMessageIDs := make([]int32, 0)
	seenNewMessageIDs := make(map[int32]struct{})
	existingRandomIDs := make(map[int32]int64)

	for _, update := range in.Updates {
		switch u := update.(type) {
		case *tg.TLUpdateMessageID:
			if u == nil {
				return nil, fmt.Errorf("envelope: nil updateMessageID")
			}
			if randomID, ok := existingRandomIDs[u.Id]; ok && randomID != u.RandomId {
				return nil, fmt.Errorf("envelope: duplicate updateMessageID for message_id %d", u.Id)
			}
			existingRandomIDs[u.Id] = u.RandomId
		case *tg.TLUpdateNewMessage:
			messageID, ok := messageIDFromUpdateNewMessage(u)
			if !ok {
				return nil, fmt.Errorf("envelope: updateNewMessage missing message id")
			}
			if _, ok := seenNewMessageIDs[messageID]; ok {
				return nil, fmt.Errorf("envelope: duplicate updateNewMessage message_id %d", messageID)
			}
			newMessageIDs = append(newMessageIDs, messageID)
			seenNewMessageIDs[messageID] = struct{}{}
		}
	}

	messageIDUpdates := make([]tg.UpdateClazz, 0, len(newMessageIDs))
	for _, messageID := range newMessageIDs {
		randomID, ok := in.MessageIDByID[messageID]
		if !ok {
			return nil, fmt.Errorf("envelope: missing updateMessageID mapping for message_id %d", messageID)
		}
		if randomID == 0 {
			return nil, fmt.Errorf("envelope: zero random_id for updateNewMessage message_id %d", messageID)
		}
		if existingRandomID, ok := existingRandomIDs[messageID]; ok && existingRandomID != randomID {
			return nil, fmt.Errorf("envelope: updateMessageID random_id mismatch for message_id %d", messageID)
		}
		messageIDUpdates = append(messageIDUpdates, tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{
			Id:       messageID,
			RandomId: randomID,
		}))
	}

	replyUpdates := make([]tg.UpdateClazz, 0, len(messageIDUpdates)+len(in.Updates))
	replyUpdates = append(replyUpdates, messageIDUpdates...)
	for _, update := range in.Updates {
		if _, ok := update.(*tg.TLUpdateMessageID); ok {
			continue
		}
		replyUpdates = append(replyUpdates, update)
	}
	return replyUpdates, nil
}

func messageIDFromUpdateNewMessage(update *tg.TLUpdateNewMessage) (int32, bool) {
	if update == nil {
		return 0, false
	}

	switch message := update.Message.(type) {
	case *tg.TLMessage:
		if message == nil {
			return 0, false
		}
		return message.Id, true
	case *tg.TLMessageService:
		if message == nil {
			return 0, false
		}
		return message.Id, true
	default:
		return 0, false
	}
}
