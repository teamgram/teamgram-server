package repository

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestStorageErrorPreservesCause(t *testing.T) {
	cause := errors.New("mysql unavailable")

	err := storageError("query", cause)

	if !errors.Is(err, msg.ErrMsgStorage) {
		t.Fatalf("storageError() does not wrap ErrMsgStorage: %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("storageError() does not preserve cause: %v", err)
	}
}

func TestMessageIDResolversAllowZeroIDWithoutStorage(t *testing.T) {
	repo := &Repository{}

	resolved, err := repo.ResolveMessageID(context.Background(), 1, payload.PeerTypeUser, 2, 0)
	if err != nil {
		t.Fatalf("ResolveMessageID zero error = %v", err)
	}
	if resolved != nil {
		t.Fatalf("ResolveMessageID zero = %+v, want nil", resolved)
	}

	publicID, err := repo.ResolvePeerSeqToUserMessageID(context.Background(), 1, payload.PeerTypeUser, 2, 0)
	if err != nil {
		t.Fatalf("ResolvePeerSeqToUserMessageID zero error = %v", err)
	}
	if publicID != 0 {
		t.Fatalf("ResolvePeerSeqToUserMessageID zero = %d, want 0", publicID)
	}

	bounds, err := repo.ResolveHistoryCursorIDs(context.Background(), 1, payload.PeerTypeUser, 2, 0, 0, 0)
	if err != nil {
		t.Fatalf("ResolveHistoryCursorIDs zero error = %v", err)
	}
	if bounds != (HistoryCursorBounds{}) {
		t.Fatalf("ResolveHistoryCursorIDs zero = %+v, want zero bounds", bounds)
	}
}

func TestHistoryMessageRowToMessageUsesPublicIDs(t *testing.T) {
	event := payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   1001,
		PeerSeq:              7,
		MessageID:            42,
		PeerType:             payload.PeerTypeUser,
		PeerID:               2002,
		FromUserID:           1002,
		ToUserID:             1001,
		Date:                 1_772_000_000,
		MessageText:          "reply",
		ReplyToUserMessageID: 21,
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	msg, err := historyMessageRowToMessage(model.HistoryMessageRow{
		CanonicalMessageID: 1001,
		PeerSeq:            7,
		UserMessageID:      42,
		FromUserID:         1002,
		Outgoing:           false,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		MessageKind:        MessageKindText,
		MessageText:        "reply",
		MessageDate:        1_772_000_000,
		ViewPayload:        body,
	})
	if err != nil {
		t.Fatalf("historyMessageRowToMessage error = %v", err)
	}
	if msg.UserMessageID != 42 {
		t.Fatalf("UserMessageID = %d, want 42", msg.UserMessageID)
	}
	if msg.PeerSeq != 7 {
		t.Fatalf("PeerSeq = %d, want internal peer seq 7", msg.PeerSeq)
	}
	if msg.ReplyToUserMessageID != 21 {
		t.Fatalf("ReplyToUserMessageID = %d, want 21", msg.ReplyToUserMessageID)
	}
	if msg.ReplyToPeerSeq != 0 {
		t.Fatalf("ReplyToPeerSeq = %d, want 0 for v2 payload", msg.ReplyToPeerSeq)
	}
}
