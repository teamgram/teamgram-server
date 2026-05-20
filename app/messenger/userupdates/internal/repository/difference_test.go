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
