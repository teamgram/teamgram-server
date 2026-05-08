package repository

import (
	"context"
	"encoding/json"
	"errors"
	"math"
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
	}, nil)
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

func TestHistoryMessageRowToMessageResolvesLegacyReplyPeerSeq(t *testing.T) {
	event := payload.MessageEventV1{
		SchemaVersion:      1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 1001,
		MessageID:          42,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "reply",
		ReplyToPeerSeq:     9,
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	msg, err := historyMessageRowToMessage(model.HistoryMessageRow{
		CanonicalMessageID: 1001,
		PeerSeq:            17,
		UserMessageID:      42,
		FromUserID:         1002,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		MessageKind:        MessageKindText,
		MessageText:        "reply",
		MessageDate:        1_772_000_000,
		ViewPayload:        body,
	}, func(peerSeq int64) (int64, error) {
		if peerSeq != 9 {
			t.Fatalf("resolver peer_seq = %d, want 9", peerSeq)
		}
		return 84, nil
	})
	if err != nil {
		t.Fatalf("historyMessageRowToMessage error = %v", err)
	}
	if msg.ReplyToPeerSeq != 9 {
		t.Fatalf("ReplyToPeerSeq = %d, want internal fallback peer_seq 9", msg.ReplyToPeerSeq)
	}
	if msg.ReplyToUserMessageID != 84 {
		t.Fatalf("ReplyToUserMessageID = %d, want resolved public id 84", msg.ReplyToUserMessageID)
	}
}

func TestHistoryPeerSeqBoundsKeepInt64Values(t *testing.T) {
	peerSeq := int64(math.MaxInt32) + 10
	bounds := HistoryCursorBounds{
		OffsetPeerSeq: int64(math.MaxInt32) + 20,
		MaxPeerSeq:    int64(math.MaxInt32) + 20,
		MinPeerSeq:    int64(math.MaxInt32) + 1,
	}
	if !historyMessageWithinBounds(peerSeq, bounds) {
		t.Fatalf("peer_seq %d should be inside int64 bounds %+v", peerSeq, bounds)
	}
	if historyMessageWithinBounds(bounds.MaxPeerSeq, bounds) {
		t.Fatalf("peer_seq equal max bound should be excluded")
	}
	if historyMessageWithinBounds(bounds.MinPeerSeq, bounds) {
		t.Fatalf("peer_seq equal min bound should be excluded")
	}
	if got := historyOffsetMarker(bounds.OffsetPeerSeq); got.OffsetID <= 0 {
		t.Fatalf("historyOffsetMarker(%d).OffsetID = %d, want positive marker without int32 cast", bounds.OffsetPeerSeq, got.OffsetID)
	}
}

func TestHistoryPeerSeqBoundsNoMatchExcludesAllRows(t *testing.T) {
	bounds := HistoryCursorBounds{NoMatch: true}
	if historyMessageWithinBounds(1, bounds) {
		t.Fatalf("historyMessageWithinBounds() = true, want false for unresolved positive public cursor")
	}
}

func TestSearchHashTagOffsetUsesResolvedPeerSeq(t *testing.T) {
	publicOffsetID := int32(42)
	resolved := &ResolvedMessageID{
		UserMessageID: int64(publicOffsetID),
		PeerSeq:       9001,
	}
	if got := searchHashTagOffsetPeerSeq(publicOffsetID, resolved); got != 9001 {
		t.Fatalf("searchHashTagOffsetPeerSeq() = %d, want resolved peer_seq 9001", got)
	}
	if got := searchHashTagOffsetPeerSeq(publicOffsetID, nil); got != 0 {
		t.Fatalf("searchHashTagOffsetPeerSeq(nil) = %d, want 0", got)
	}
	if got := searchHashTagOffsetPeerSeq(0, resolved); got != 0 {
		t.Fatalf("searchHashTagOffsetPeerSeq(0) = %d, want 0", got)
	}
}
