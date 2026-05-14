package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestUserEventFromModelSkipsLegacyDialogPublicUpdateHydration(t *testing.T) {
	body, err := json.Marshal(payload.DialogEventV1{
		SchemaVersion:    payload.DialogEventSchemaVersion,
		EventKind:        payload.DialogEventDraftCleared,
		PublicUpdateType: payload.DialogEventDraftCleared,
		PeerType:         payload.PeerTypeUser,
		PeerID:           2002,
	})
	if err != nil {
		t.Fatalf("marshal dialog event: %v", err)
	}

	event, err := (*Repository)(nil).userEventFromModel(context.Background(), model.UserPtsEvents{
		UserId:             1001,
		Pts:                7,
		PtsCount:           1,
		OperationId:        "dialog-op",
		OpType:             OpTypeSendMessage,
		EventType:          EventTypeDialogPublicUpdate,
		PeerType:           payload.PeerTypeUser,
		PeerId:             2002,
		EventSchemaVersion: payload.MessageEventSchemaVersionV1,
		EventCodec:         PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	})
	if err != nil {
		t.Fatalf("userEventFromModel() error = %v", err)
	}
	if event.EventSchemaVersion != payload.MessageEventSchemaVersionV1 {
		t.Fatalf("event schema = %d, want legacy dialog schema v1", event.EventSchemaVersion)
	}
	if !bytes.Equal(event.EventPayload, body) {
		t.Fatalf("dialog payload was modified by message hydration")
	}
}

func TestNeedsLegacyMessageHydrationOnlyForMessageProjectionEvents(t *testing.T) {
	for _, eventType := range []int32{EventTypeNewMessage, EventTypeReadHistory, EventTypeUpdatePinnedMessage, EventTypeEditMessage} {
		if !needsLegacyMessageHydration(eventType) {
			t.Fatalf("event type %d should need legacy message hydration", eventType)
		}
	}
	for _, eventType := range []int32{EventTypeDialogPublicUpdate, EventTypeMarkDialogUnread, EventTypeScheduledMarker, EventTypeDeleteMessages, EventTypeDeleteHistory} {
		if needsLegacyMessageHydration(eventType) {
			t.Fatalf("event type %d should not need legacy message hydration", eventType)
		}
	}
}

func TestLatestDifferenceDateUsesNewestEventTime(t *testing.T) {
	v2Body, err := json.Marshal(payload.MessageEventV2{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.EventKindNewMessage,
		Date:          1_772_000_010,
		EditDate:      1_772_000_020,
	})
	if err != nil {
		t.Fatalf("marshal v2 event: %v", err)
	}
	v4Body, err := json.Marshal(payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion: payload.MessageOperationSchemaVersionV4,
			Date:          1_772_000_030,
		},
	})
	if err != nil {
		t.Fatalf("marshal v4 event: %v", err)
	}

	got, err := latestDifferenceDate([]UserEvent{
		{
			EventType:          EventTypeNewMessage,
			EventSchemaVersion: payload.MessageEventSchemaVersion,
			EventCodec:         PayloadCodecJSON,
			EventPayload:       v2Body,
			EventPayloadHash:   payload.HashBytes(v2Body),
		},
		{
			EventType:          EventTypeNewMessage,
			EventSchemaVersion: payload.MessageEventSchemaVersionV4,
			EventCodec:         PayloadCodecJSON,
			EventPayload:       v4Body,
			EventPayloadHash:   payload.HashBytes(v4Body),
		},
	}, []AuthSeqEvent{
		{Date: 1_772_000_040},
	}, 1_772_000_001)
	if err != nil {
		t.Fatalf("latestDifferenceDate() error = %v", err)
	}
	if got != 1_772_000_040 {
		t.Fatalf("latestDifferenceDate() = %d, want auth seq max date", got)
	}

	got, err = latestDifferenceDate([]UserEvent{
		{
			EventType:          EventTypeNewMessage,
			EventSchemaVersion: payload.MessageEventSchemaVersion,
			EventCodec:         PayloadCodecJSON,
			EventPayload:       v2Body,
			EventPayloadHash:   payload.HashBytes(v2Body),
		},
		{
			EventType:          EventTypeNewMessage,
			EventSchemaVersion: payload.MessageEventSchemaVersionV4,
			EventCodec:         PayloadCodecJSON,
			EventPayload:       v4Body,
			EventPayloadHash:   payload.HashBytes(v4Body),
		},
	}, nil, 1_772_000_001)
	if err != nil {
		t.Fatalf("latestDifferenceDate() without auth seq error = %v", err)
	}
	if got != 1_772_000_030 {
		t.Fatalf("latestDifferenceDate() = %d, want newest message event date", got)
	}
}
