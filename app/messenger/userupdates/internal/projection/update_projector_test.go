package projection

import (
	"encoding/json"
	"errors"
	"math"
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

func TestProjectDeleteMessagesForDifferenceUsesPublicIDs(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.OperationKindDeleteMessages,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		Date:                 1_772_000_000,
		DeleteUserMessageIDs: []int64{107, 108},
	})

	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                31,
		PtsCount:           1,
		EventType:          repository.EventTypeDeleteMessages,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent() error = %v", err)
	}
	update, ok := got.Update.(*tg.TLUpdateDeleteMessages)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateDeleteMessages", got.Update)
	}
	if update.Pts != 31 || update.PtsCount != 1 || len(update.Messages) != 2 || update.Messages[0] != 107 || update.Messages[1] != 108 {
		t.Fatalf("update = %+v", update)
	}
}

func TestProjectDeleteMessagesForPushUsesSeqZero(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.OperationKindDeleteMessages,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		Date:                 1_772_000_000,
		DeleteUserMessageIDs: []int64{107},
	})

	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		UserID:        1001,
		Pts:           32,
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
	if updates.Date != 1_772_000_000 {
		t.Fatalf("updates date = %d, want event date", updates.Date)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateDeleteMessages)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateDeleteMessages", updates.Updates[0])
	}
	if update.Pts != 32 || update.PtsCount != 1 || len(update.Messages) != 1 || update.Messages[0] != 107 {
		t.Fatalf("delete update = %+v", update)
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

func TestProjectUpdatePinnedMessageUsesPublicMessageIDForDifference(t *testing.T) {
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:       payload.MessageEventSchemaVersion,
		EventKind:           payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID:  7001,
		PeerSeq:             9,
		MessageID:           101,
		PinnedUserMessageID: 101,
		PeerType:            payload.PeerTypeUser,
		PeerID:              1002,
		FromUserID:          1001,
		ToUserID:            1002,
		Date:                1_772_000_000,
	})

	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                21,
		PtsCount:           1,
		EventType:          repository.EventTypeUpdatePinnedMessage,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent() error = %v", err)
	}
	update, ok := got.Update.(*tg.TLUpdatePinnedMessages)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdatePinnedMessages", got.Update)
	}
	if !update.Pinned || update.Pts != 21 || update.PtsCount != 1 {
		t.Fatalf("update = %+v", update)
	}
	if len(update.Messages) != 1 || update.Messages[0] != 101 {
		t.Fatalf("messages = %v, want public id 101", update.Messages)
	}
	if update.Messages[0] == 9 {
		t.Fatalf("messages used peer seq instead of public id: %v", update.Messages)
	}
}

func TestProjectUpdatePinnedMessageUsesPublicMessageIDForPush(t *testing.T) {
	exclude := int64(9001)
	body := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:       payload.MessageEventSchemaVersion,
		EventKind:           payload.OperationKindUpdatePinnedMessage,
		CanonicalMessageID:  7001,
		PeerSeq:             9,
		MessageID:           101,
		PinnedUserMessageID: 101,
		PeerType:            payload.PeerTypeUser,
		PeerID:              1002,
		FromUserID:          1001,
		ToUserID:            1002,
		Date:                1_772_000_000,
		AuthKeyIdExclude:    &exclude,
	})

	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		UserID:        1001,
		Pts:           22,
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
	update, ok := updates.Updates[0].(*tg.TLUpdatePinnedMessages)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdatePinnedMessages", updates.Updates[0])
	}
	if !update.Pinned || update.Pts != 22 || update.PtsCount != 1 {
		t.Fatalf("update = %+v", update)
	}
	if len(update.Messages) != 1 || update.Messages[0] != 101 {
		t.Fatalf("messages = %v, want public id 101", update.Messages)
	}
	if got.AuthKeyIDExclude == nil || *got.AuthKeyIDExclude != exclude {
		t.Fatalf("auth key exclude = %v, want %d", got.AuthKeyIDExclude, exclude)
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

func TestProjectMessageEventV1LegacyCompatibility(t *testing.T) {
	body := mustMarshalLegacyMessageEventV1(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		MessageID:          101,
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
		t.Fatalf("ProjectUserEvent(V1) error = %v, want ErrUserupdatesStorage for unhydrated legacy public id", err)
	}
}

func TestProjectMessageEventV2Compatibility(t *testing.T) {
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
		MessageText:        "v2",
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
		t.Fatalf("ProjectUserEvent(V2) error = %v", err)
	}
	if _, ok := got.Message.(*tg.TLMessage); !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", got.Message)
	}
}

func TestProjectMessageEventV3MediaGroupedForward(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 101,
		PeerSeq:            9,
		MessageID:          77,
		PeerType:           payload.PeerTypeChat,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "caption",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333},
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444},
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 303, Date: 1700000001},
	})
	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		Payload:  body,
		Pts:      19,
		PeerType: payload.PeerTypeChat,
		PeerID:   202,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok || len(updates.Updates) == 0 {
		t.Fatalf("updates = %#v, want non-empty TLUpdates", got.Updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if message.Media == nil || message.GroupedId == nil || *message.GroupedId != 444 || message.FwdFrom == nil {
		t.Fatalf("message media/grouped/forward = media:%T grouped:%v fwd:%T", message.Media, message.GroupedId, message.FwdFrom)
	}
}

func TestProjectMessageEventV3UserPeerFullPushWhenMediaGroupedForward(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 102,
		PeerSeq:            10,
		MessageID:          78,
		PeerType:           payload.PeerTypeUser,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "caption",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333},
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444},
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 303, Date: 1700000001},
	})
	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		Payload:  body,
		Pts:      20,
		PeerType: payload.PeerTypeUser,
		PeerID:   202,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok || len(updates.Updates) != 1 {
		t.Fatalf("updates = %#v, want single TLUpdates wrapper", got.Updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if message.Media == nil || message.GroupedId == nil || *message.GroupedId != 444 || message.FwdFrom == nil {
		t.Fatalf("message media/grouped/forward = media:%T grouped:%v fwd:%T", message.Media, message.GroupedId, message.FwdFrom)
	}
}

func TestProjectMessageEventV3ShortPushPreservesSilent(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 103,
		PeerSeq:            11,
		MessageID:          79,
		PeerType:           payload.PeerTypeUser,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		MessageText:        "silent",
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, Silent: true},
	})
	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		Payload:  body,
		Pts:      21,
		PeerType: payload.PeerTypeUser,
		PeerID:   202,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	update, ok := got.Updates.(*tg.TLUpdateShortMessage)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdateShortMessage", got.Updates)
	}
	if !update.Silent {
		t.Fatalf("silent = false, want true")
	}
}

func TestProjectMessageEventV3DocumentMedia(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 103,
		PeerSeq:            11,
		MessageID:          79,
		PeerType:           payload.PeerTypeChat,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "document",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 555},
	})
	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                21,
		PtsCount:           1,
		EventType:          repository.EventTypeNewMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersionV3,
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
	media, ok := message.Media.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", message.Media)
	}
	doc, ok := media.Document.(*tg.TLDocumentEmpty)
	if !ok || doc.Id != 555 {
		t.Fatalf("document = %+v ok=%v, want documentEmpty id=555", media.Document, ok)
	}
}

func TestProjectMessageEventV3RejectsOutOfRangeForwardDate(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 104,
		PeerSeq:            12,
		MessageID:          80,
		PeerType:           payload.PeerTypeUser,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "forward",
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 303, Date: int64(math.MaxInt32) + 1},
	})
	_, err := ProjectUserEvent(repository.UserEvent{
		UserID:             1001,
		Pts:                22,
		PtsCount:           1,
		EventType:          repository.EventTypeNewMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersionV3,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       body,
		EventPayloadHash:   payload.HashBytes(body),
	}, ModeDifference)
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("error = %v, want ErrUserupdatesStorage", err)
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

func mustMarshalMessageEventV3(t *testing.T, event payload.MessageEventV3) []byte {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal MessageEventV3: %v", err)
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
