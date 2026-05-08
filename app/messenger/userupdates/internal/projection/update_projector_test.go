package projection

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestProjectMessageEventNewMessageForDifference(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		PeerSeq:            9,
		MessageID:          101,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "hello",
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
	if message.Id != 101 || message.Message != "hello" {
		t.Fatalf("message = %+v", message)
	}
	update, ok := got.Update.(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", got.Update)
	}
	if update.Message != got.Message || update.Pts != 18 || update.PtsCount != 1 {
		t.Fatalf("update = %+v", update)
	}
}

func TestProjectMessageEventNewMessageForPushShortMessage(t *testing.T) {
	exclude := int64(9001)
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		PeerSeq:            9,
		MessageID:          101,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "hello",
		AuthKeyIdExclude:   &exclude,
	})

	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		UserID:        1001,
		Pts:           18,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1002,
		Payload:       body,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	update, ok := got.Updates.(*tg.TLUpdateShortMessage)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdateShortMessage", got.Updates)
	}
	if update.Id != 101 || update.UserId != 1002 || update.Message != "hello" || update.Pts != 18 || update.PtsCount != 1 {
		t.Fatalf("update = %+v", update)
	}
	if got.AuthKeyIDExclude == nil || *got.AuthKeyIDExclude != exclude {
		t.Fatalf("auth key exclude = %v, want %d", got.AuthKeyIDExclude, exclude)
	}
}

func TestProjectReadHistoryUsesSeqZeroForPush(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.OperationKindReadHistory,
		PeerSeq:              42,
		MessageID:            101,
		ReadMaxUserMessageID: 88,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		Date:                 1_772_000_000,
		Out:                  true,
	})

	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		UserID:        1001,
		Pts:           19,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1002,
		Payload:       body,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Updates)
	}
	if updates.Seq != 0 {
		t.Fatalf("updates seq = %d, want 0", updates.Seq)
	}
	readUpdate, ok := updates.Updates[0].(*tg.TLUpdateReadHistoryOutbox)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateReadHistoryOutbox", updates.Updates[0])
	}
	if readUpdate.MaxId != 88 {
		t.Fatalf("read max id = %d, want public id 88", readUpdate.MaxId)
	}
}

func TestProjectEditMessageUsesSeqZeroForPush(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.OperationKindEditMessage,
		CanonicalMessageID: 7001,
		PeerSeq:            9,
		MessageID:          101,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		EditDate:           1_772_000_100,
		MessageText:        "edited",
	})

	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		UserID:        1001,
		Pts:           20,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1002,
		Payload:       body,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Updates)
	}
	if updates.Seq != 0 {
		t.Fatalf("updates seq = %d, want 0", updates.Seq)
	}
	if _, ok := updates.Updates[0].(*tg.TLUpdateEditMessage); !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateEditMessage", updates.Updates[0])
	}
}

func TestProjectRejectsPayloadHashMismatch(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.EventKindNewMessage,
		PeerSeq:       9,
		MessageID:     101,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1002,
		Date:          1_772_000_000,
	})

	_, err := ProjectUserEvent(repository.UserEvent{
		Pts:                18,
		PtsCount:           1,
		EventType:          repository.EventTypeNewMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes([]byte("different")),
	}, ModeDifference)
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestProjectMessageEventV2UsesReplyPublicID(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   7001,
		PeerSeq:              9,
		MessageID:            101,
		ReplyToUserMessageID: 77,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		FromUserID:           1002,
		ToUserID:             1001,
		Date:                 1_772_000_000,
		MessageText:          "reply",
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
	reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("reply = %T, want *tg.TLMessageReplyHeader", message.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 77 {
		t.Fatalf("reply id = %v, want 77", reply.ReplyToMsgId)
	}
}

func TestProjectRejectsUnhydratedLegacyV1MessageID(t *testing.T) {
	body := mustMarshalLegacyMessageEventV1(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1002,
		ToUserID:           1001,
		Date:               1_772_000_000,
		MessageText:        "legacy",
	})

	_, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                18,
		PtsCount:           1,
		EventType:          repository.EventTypeNewMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersionV1,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("error = %v, want ErrUserupdatesStorage", err)
	}
}

func mustMarshalMessageEventV2(t *testing.T, event payload.MessageEventV2) []byte {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event: %v", err)
	}
	return body
}

func mustMarshalLegacyMessageEventV1(t *testing.T, event payload.MessageEventV1) []byte {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal legacy message event: %v", err)
	}
	return body
}
