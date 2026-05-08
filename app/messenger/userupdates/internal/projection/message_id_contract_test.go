package projection

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectionUsesPublicMessageIDWhenPeerSeqDiffers(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		PeerSeq:            7,
		MessageID:          55,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "public id contract",
	})

	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                18,
		PtsCount:           1,
		EventType:          repository.EventTypeNewMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent() error = %v", err)
	}
	message, ok := got.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", got.Message)
	}
	if message.Id != 55 || message.Id == int32(7) {
		t.Fatalf("TLMessage.Id = %d, want public id 55 and not peer_seq 7", message.Id)
	}
}

func TestProjectionReadHistoryMaxIDUsesReadMaxUserMessageID(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.OperationKindReadHistory,
		PeerSeq:              7,
		MessageID:            7,
		ReadMaxUserMessageID: 55,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		Date:                 1_772_000_000,
		Out:                  false,
	})

	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                19,
		PtsCount:           1,
		EventType:          repository.EventTypeReadHistory,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent() error = %v", err)
	}
	update, ok := got.Update.(*tg.TLUpdateReadHistoryInbox)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateReadHistoryInbox", got.Update)
	}
	if update.MaxId != 55 || update.MaxId == int32(7) {
		t.Fatalf("read history max_id = %d, want public id 55 and not peer_seq 7", update.MaxId)
	}
}
