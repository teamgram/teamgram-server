package projection

import (
	"bytes"
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
	if message.Out {
		t.Fatalf("message.Out = true, want false for receiver projection")
	}
	if message.FromId != nil {
		t.Fatalf("message.FromId = %#v, want nil for receiver projection", message.FromId)
	}
	peer, ok := message.PeerId.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1002 {
		t.Fatalf("message.PeerId = %#v, want peerUser(1002)", message.PeerId)
	}
	update, ok := got.Update.(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", got.Update)
	}
	if update.Message != got.Message || update.Pts != 18 || update.PtsCount != 1 {
		t.Fatalf("update = %+v", update)
	}
}

func TestProjectMessageEventNewMessageForPushFullMessage(t *testing.T) {
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
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Updates)
	}
	if len(updates.Updates) != 1 {
		t.Fatalf("updates len = %d, want 1", len(updates.Updates))
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	peer, ok := message.PeerId.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1002 {
		t.Fatalf("message peer = %#v, want peerUser(1002)", message.PeerId)
	}
	if message.Id != 101 || message.Message != "hello" || message.Out || update.Pts != 18 || update.PtsCount != 1 {
		t.Fatalf("message/update = %+v / %+v", message, update)
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

func TestProjectMessageEventV3ChatCreateServiceAction(t *testing.T) {
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
		ServiceAction: &payload.ServiceActionRefV1{
			SchemaVersion: payload.ServiceActionSchemaVersionV1,
			Kind:          payload.ServiceActionKindChatCreate,
			Title:         "new chat",
			Users:         []int64{102, 103},
		},
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
	if !ok || len(updates.Updates) != 1 {
		t.Fatalf("updates = %#v, want one TLUpdates item", got.Updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	service, ok := update.Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessageService", update.Message)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatCreate)
	if !ok || action.Title != "new chat" || len(action.Users) != 2 {
		t.Fatalf("action = %T %+v, want chat create title/users", service.Action, service.Action)
	}
}

func TestProjectMessageEventV3PhotoSizesSurviveJSONProjection(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 105,
		PeerSeq:            13,
		MessageID:          81,
		PeerType:           payload.PeerTypeChat,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "photo",
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV1,
			Kind:          "photo",
			ID:            777,
			AccessHash:    888,
			FileReference: []byte("1234567890123456789012345"),
			Date:          1700000000,
			DcID:          2,
			PhotoSizes: []payload.PhotoSizeRefV1{
				{Kind: "stripped", Type: "i", Bytes: []byte{0x01, 0x16, 0x28, 0xaa}},
				{Kind: "progressive", Type: "y", W: 1280, H: 394, Sizes: []int32{100, 200, 300}},
			},
		},
	})

	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             202,
		Pts:                24,
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
	media, ok := message.Media.(*tg.TLMessageMediaPhoto)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaPhoto", message.Media)
	}
	photo, ok := media.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("photo = %T, want *tg.TLPhoto", media.Photo)
	}
	if len(photo.Sizes) != 2 {
		t.Fatalf("photo sizes len = %d, want 2", len(photo.Sizes))
	}
	stripped, ok := photo.Sizes[0].(*tg.TLPhotoStrippedSize)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoStrippedSize", photo.Sizes[0])
	}
	if !bytes.Equal(stripped.Bytes, []byte{0x01, 0x16, 0x28, 0xaa}) {
		t.Fatalf("stripped bytes = %#v, want telegram stripped preview bytes", stripped.Bytes)
	}
	progressive, ok := photo.Sizes[1].(*tg.TLPhotoSizeProgressive)
	if !ok {
		t.Fatalf("projected size = %T, want TLPhotoSizeProgressive", photo.Sizes[1])
	}
	if progressive.Sizes[2] != 300 {
		t.Fatalf("progressive sizes = %#v, want final offset 300", progressive.Sizes)
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

func TestProjectMessageEventV3FullPushPreservesSilent(t *testing.T) {
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
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if !message.Silent {
		t.Fatalf("silent = false, want true")
	}
}

func TestProjectMessageEventV3FullPushPreservesEntities(t *testing.T) {
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
		MessageText:        "@alice see https://teamgram.io",
		Entities: []payload.MessageEntityV1{
			{Offset: 0, Length: 6, Kind: "mention"},
			{Offset: 11, Length: 19, Kind: "url"},
		},
	})
	got, err := ProjectPushTask(&payload.PushTaskKafkaMessageV1{
		Payload:  body,
		Pts:      22,
		PeerType: payload.PeerTypeUser,
		PeerID:   202,
	})
	if err != nil {
		t.Fatalf("ProjectPushTask() error = %v", err)
	}
	updates, ok := got.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", got.Updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if len(message.Entities) != 2 {
		t.Fatalf("entities len = %d, want 2", len(message.Entities))
	}
	if _, ok := message.Entities[0].(*tg.TLMessageEntityMention); !ok {
		t.Fatalf("entities[0] = %T, want mention", message.Entities[0])
	}
	if _, ok := message.Entities[1].(*tg.TLMessageEntityUrl); !ok {
		t.Fatalf("entities[1] = %T, want url", message.Entities[1])
	}
}

func TestProjectMessageEventV3FullMessagePreservesEntities(t *testing.T) {
	body := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 105,
		PeerSeq:            13,
		MessageID:          81,
		PeerType:           payload.PeerTypeChat,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		MessageText:        "hello user",
		Entities: []payload.MessageEntityV1{
			{Offset: 6, Length: 4, Kind: "mention_name", UserID: 303},
		},
	})
	got, err := ProjectUserEvent(repository.UserEvent{
		UserID:             202,
		Pts:                23,
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
	if len(message.Entities) != 1 {
		t.Fatalf("entities len = %d, want 1", len(message.Entities))
	}
	mentionName, ok := message.Entities[0].(*tg.TLMessageEntityMentionName)
	if !ok {
		t.Fatalf("entity = %T, want mention name", message.Entities[0])
	}
	if mentionName.UserId != 303 || mentionName.Offset != 6 || mentionName.Length != 4 {
		t.Fatalf("mention name = %#v, want user 303 range 6:4", mentionName)
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
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV1,
			Kind:          "document",
			ID:            555,
			AccessHash:    666,
			FileReference: []byte("doc-ref"),
			Date:          1700000000,
			DcID:          4,
			Size:          98765,
			MimeType:      "application/pdf",
			DocumentAttributes: []payload.DocumentAttributeRefV1{
				{Kind: "filename", FileName: "report.pdf"},
			},
		},
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
	doc, ok := media.Document.(*tg.TLDocument)
	if !ok || doc.Id != 555 || doc.AccessHash != 666 || doc.MimeType != "application/pdf" || doc.Size2 != 98765 || doc.DcId != 4 {
		t.Fatalf("document = %+v ok=%v, want displayable document", media.Document, ok)
	}
	if len(doc.Attributes) != 1 {
		t.Fatalf("document attrs = %+v, want filename attr", doc.Attributes)
	}
	filename, ok := doc.Attributes[0].(*tg.TLDocumentAttributeFilename)
	if !ok || filename.FileName != "report.pdf" {
		t.Fatalf("document filename attr = %+v ok=%v, want report.pdf", doc.Attributes[0], ok)
	}
}

func TestProjectMessageEventV3DocumentMediaProjectsFullUploadedDocumentContract(t *testing.T) {
	videoStartTs := 1.25
	videoTimestamp := int32(7)
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
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV2,
			Kind:          "document",
			ID:            555,
			AccessHash:    666,
			FileReference: []byte("doc-ref"),
			Date:          1700000000,
			DcID:          4,
			Size:          98765,
			MimeType:      "video/mp4",
			DocumentVideoThumbs: []payload.VideoSizeRefV1{
				{Kind: "size", Type: "v", W: 320, H: 200, Size: 4567, VideoStartTs: &videoStartTs},
			},
			DocumentAttributes: []payload.DocumentAttributeRefV1{
				{Kind: "filename", FileName: "clip.mp4"},
				{Kind: "video", W: 1280, H: 720, DurationFloat: 3.5, SupportsStreaming: true, VideoStartTs: &videoStartTs},
				{Kind: "sticker", Alt: ":)", StickerSet: &payload.StickerSetRefV1{Kind: "id", ID: 1001, AccessHash: 2002}, Mask: true, MaskCoords: &payload.MaskCoordsRefV1{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}},
				{Kind: "custom_emoji", Alt: ":)", StickerSet: &payload.StickerSetRefV1{Kind: "id", ID: 3003, AccessHash: 4004}, Free: true, TextColor: true},
				{Kind: "has_stickers"},
			},
			DocumentMediaFlags: &payload.DocumentMediaFlagsV1{Video: true, Spoiler: true},
			VideoCover: &payload.PhotoRefV1{
				ID:            777,
				AccessHash:    888,
				FileReference: []byte("cover-ref"),
				Date:          1700000001,
				DcID:          5,
				Sizes: []payload.PhotoSizeRefV1{
					{Kind: "size", Type: "x", W: 640, H: 360, Size: 4321},
				},
			},
			VideoTimestamp: &videoTimestamp,
		},
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
	if !media.Video || !media.Spoiler {
		t.Fatalf("messageMediaDocument flags video=%v spoiler=%v, want true/true", media.Video, media.Spoiler)
	}
	doc, ok := media.Document.(*tg.TLDocument)
	if !ok {
		t.Fatalf("document = %T, want *tg.TLDocument", media.Document)
	}
	assertProjectedDocumentIdentity(t, doc)
	if len(doc.VideoThumbs) != 1 {
		t.Fatalf("VideoThumbs len = %d, want 1", len(doc.VideoThumbs))
	}
	videoThumb, ok := doc.VideoThumbs[0].(*tg.TLVideoSize)
	if !ok {
		t.Fatalf("VideoThumbs[0] = %T, want *tg.TLVideoSize", doc.VideoThumbs[0])
	}
	if videoThumb.Type != "v" || videoThumb.W != 320 || videoThumb.H != 200 || videoThumb.Size2 != 4567 {
		t.Fatalf("VideoThumbs[0] = %#v, want videoSize v 320x200/4567", videoThumb)
	}
	if videoThumb.VideoStartTs == nil || *videoThumb.VideoStartTs != videoStartTs {
		t.Fatalf("VideoThumbs[0].VideoStartTs = %v, want %v", videoThumb.VideoStartTs, videoStartTs)
	}
	assertProjectedDocumentAttributes(t, doc.Attributes, videoStartTs)
	if media.VideoTimestamp == nil || *media.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", media.VideoTimestamp, videoTimestamp)
	}
	videoCover, ok := media.VideoCover.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("VideoCover = %T, want *tg.TLPhoto", media.VideoCover)
	}
	assertProjectedVideoCover(t, videoCover)
}

func TestMessageDocumentMediaInfersLegacyFlagsFromAttributes(t *testing.T) {
	media := messageMedia(&payload.MediaRefV1{
		Kind:     "document",
		MimeType: "video/mp4",
		DocumentAttributes: []payload.DocumentAttributeRefV1{
			{Kind: "video", RoundMessage: true},
			{Kind: "audio", Voice: true},
		},
	})
	documentMedia, ok := media.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", media)
	}
	if !documentMedia.Video || !documentMedia.Round || !documentMedia.Voice {
		t.Fatalf("document media flags = video:%v round:%v voice:%v, want all inferred", documentMedia.Video, documentMedia.Round, documentMedia.Voice)
	}
}

func TestMessageDocumentMediaDoesNotInferVideoForWebMStickerOrCustomEmoji(t *testing.T) {
	for _, tt := range []struct {
		name string
		attr payload.DocumentAttributeRefV1
	}{
		{name: "sticker", attr: payload.DocumentAttributeRefV1{Kind: "sticker"}},
		{name: "custom_emoji", attr: payload.DocumentAttributeRefV1{Kind: "custom_emoji"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			media := messageMedia(&payload.MediaRefV1{
				Kind:     "document",
				MimeType: "video/webm",
				DocumentAttributes: []payload.DocumentAttributeRefV1{
					{Kind: "video"},
					tt.attr,
				},
			})
			documentMedia, ok := media.(*tg.TLMessageMediaDocument)
			if !ok {
				t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", media)
			}
			if documentMedia.Video {
				t.Fatalf("document media Video = true, want false for video/webm %s", tt.name)
			}
		})
	}
}

func assertProjectedDocumentIdentity(t *testing.T, document *tg.TLDocument) {
	t.Helper()
	if document.Id != 555 ||
		document.AccessHash != 666 ||
		string(document.FileReference) != "doc-ref" ||
		document.Date != 1700000000 ||
		document.DcId != 4 ||
		document.MimeType != "video/mp4" ||
		document.Size2 != 98765 {
		t.Fatalf("document identity = %#v, want full uploaded document identity", document)
	}
}

func TestProjectMessageEventV3ContactMedia(t *testing.T) {
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
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV1,
			Kind:          "contact",
			PhoneNumber:   "8613000000001",
			FirstName:     "13000000001",
			LastName:      "t2",
			UserID:        1571266964,
		},
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
	media, ok := message.Media.(*tg.TLMessageMediaContact)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaContact", message.Media)
	}
	if media.PhoneNumber != "8613000000001" || media.FirstName != "13000000001" || media.LastName != "t2" || media.UserId != 1571266964 {
		t.Fatalf("contact media = %+v, want shared contact fields", media)
	}
}

func assertProjectedDocumentAttributes(t *testing.T, attrs []tg.DocumentAttributeClazz, videoStartTs float64) {
	t.Helper()
	filename, hasFilename := findProjectionDocumentAttribute[*tg.TLDocumentAttributeFilename](attrs)
	video, hasVideo := findProjectionDocumentAttribute[*tg.TLDocumentAttributeVideo](attrs)
	sticker, hasSticker := findProjectionDocumentAttribute[*tg.TLDocumentAttributeSticker](attrs)
	customEmoji, hasCustomEmoji := findProjectionDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](attrs)
	_, hasStickers := findProjectionDocumentAttribute[*tg.TLDocumentAttributeHasStickers](attrs)
	if !hasFilename || !hasVideo || !hasSticker || !hasCustomEmoji || !hasStickers {
		t.Fatalf("document attrs = %#v, want filename/video/sticker/custom_emoji/has_stickers", attrs)
	}
	if filename.FileName != "clip.mp4" {
		t.Fatalf("filename attr FileName = %q, want clip.mp4", filename.FileName)
	}
	if video.Duration != 3.5 || video.W != 1280 || video.H != 720 || !video.SupportsStreaming {
		t.Fatalf("video attr = %#v, want duration/w/h/supports_streaming preserved", video)
	}
	if video.VideoStartTs == nil || *video.VideoStartTs != videoStartTs {
		t.Fatalf("video attr VideoStartTs = %v, want %v", video.VideoStartTs, videoStartTs)
	}
	stickerSet, ok := sticker.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || stickerSet.Id != 1001 || stickerSet.AccessHash != 2002 {
		t.Fatalf("sticker stickerset = %#v, want inputStickerSetID 1001/2002", sticker.Stickerset)
	}
	maskCoords := sticker.MaskCoords
	if maskCoords == nil || maskCoords.N != 1 || maskCoords.X != 0.5 || maskCoords.Y != 0.25 || maskCoords.Zoom != 1.5 {
		t.Fatalf("sticker mask coords = %#v, want exact TLMaskCoords", sticker.MaskCoords)
	}
	if sticker.Alt != ":)" || !sticker.Mask {
		t.Fatalf("sticker attr = %#v, want alt and mask preserved", sticker)
	}
	customStickerSet, ok := customEmoji.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || customStickerSet.Id != 3003 || customStickerSet.AccessHash != 4004 {
		t.Fatalf("custom emoji stickerset = %#v, want inputStickerSetID 3003/4004", customEmoji.Stickerset)
	}
	if customEmoji.Alt != ":)" || !customEmoji.Free || !customEmoji.TextColor {
		t.Fatalf("custom emoji attr = %#v, want alt/free/text_color preserved", customEmoji)
	}
}

func assertProjectedVideoCover(t *testing.T, cover *tg.TLPhoto) {
	t.Helper()
	if cover.Id != 777 || cover.AccessHash != 888 || string(cover.FileReference) != "cover-ref" || cover.Date != 1700000001 || cover.DcId != 5 {
		t.Fatalf("VideoCover = %#v, want full photo 777", cover)
	}
	if len(cover.Sizes) != 1 {
		t.Fatalf("VideoCover.Sizes len = %d, want 1", len(cover.Sizes))
	}
	size, ok := cover.Sizes[0].(*tg.TLPhotoSize)
	if !ok {
		t.Fatalf("VideoCover.Sizes[0] = %T, want *tg.TLPhotoSize", cover.Sizes[0])
	}
	if size.Type != "x" || size.W != 640 || size.H != 360 || size.Size2 != 4321 {
		t.Fatalf("VideoCover.Sizes[0] = %#v, want photoSize x 640x360/4321", size)
	}
}

func findProjectionDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) (T, bool) {
	for _, attr := range attrs {
		if got, ok := attr.(T); ok {
			return got, true
		}
	}
	var zero T
	return zero, false
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
