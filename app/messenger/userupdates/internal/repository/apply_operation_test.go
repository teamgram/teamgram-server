package repository

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestApplyUserOperationBatchEmptyReturnsEmptyResult(t *testing.T) {
	repo := &Repository{}

	got, err := repo.ApplyUserOperationBatch(context.Background(), nil)
	if err != nil {
		t.Fatalf("ApplyUserOperationBatch(nil) error = %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("ApplyUserOperationBatch(nil) len = %d, want 0", len(got))
	}
}

func TestBuildEventAndResponseCarriesAuthKeyExclude(t *testing.T) {
	authKeyIDExclude := int64(9001)
	eventPayload, _, responsePayload, _, err := buildEventAndResponse(
		ApplyUserOperationInput{
			OperationID:      "op",
			AuthKeyIDExclude: &authKeyIDExclude,
		},
		messageOperationFromV1(payload.MessageOperationV1{
			SchemaVersion:        payload.MessageOperationSchemaVersion,
			OperationKind:        payload.OperationKindSendMessage,
			CanonicalMessageID:   7001,
			PeerType:             payload.PeerTypeUser,
			PeerID:               2002,
			PeerSeq:              9,
			UserMessageID:        101,
			FromUserID:           1001,
			ToUserID:             2002,
			Date:                 1777781234,
			Out:                  true,
			MessageText:          "hello",
			ReplyToUserMessageID: 77,
		}),
		38,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.AuthKeyIdExclude == nil || *event.AuthKeyIdExclude != authKeyIDExclude {
		t.Fatalf("auth_key_id_exclude = %v, want %d", event.AuthKeyIdExclude, authKeyIDExclude)
	}
	if event.MessageID != 101 || event.PeerSeq != 9 || event.ReplyToUserMessageID != 77 {
		t.Fatalf("event public/internal ids = %+v, want message_id=101 peer_seq=9 reply=77", event)
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal response payload: %v", err)
	}
	if response.UserMessageID != 101 {
		t.Fatalf("response user_message_id = %d, want 101", response.UserMessageID)
	}
}

func TestBuildEditEventAndResponseCarriesPublicMessageID(t *testing.T) {
	eventPayload, _, responsePayload, _, err := buildEventAndResponse(
		ApplyUserOperationInput{OperationID: "edit-op"},
		messageOperationFromV1(payload.MessageOperationV1{
			SchemaVersion:      payload.MessageOperationSchemaVersion,
			OperationKind:      payload.OperationKindEditMessage,
			CanonicalMessageID: 7001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             2002,
			PeerSeq:            9,
			UserMessageID:      101,
			FromUserID:         1001,
			ToUserID:           2002,
			Date:               1777781234,
			EditDate:           1777782234,
			EditVersion:        2,
			Out:                true,
			MessageText:        "edited",
		}),
		39,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV2
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.EventKind != payload.OperationKindEditMessage || event.MessageID != 101 || event.PeerSeq != 9 {
		t.Fatalf("edit event ids = %+v, want public message_id=101 and peer_seq=9", event)
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal response payload: %v", err)
	}
	if response.UserMessageID != 101 {
		t.Fatalf("edit response user_message_id = %d, want 101", response.UserMessageID)
	}
}

func TestBuildEventAndResponseV3CarriesMediaAttrsForward(t *testing.T) {
	eventPayload, _, responsePayload, _, err := buildEventAndResponse(
		ApplyUserOperationInput{OperationID: "v3-op"},
		messageOperationFromV3(payload.MessageOperationV3{
			SchemaVersion:      payload.MessageOperationSchemaVersionV3,
			OperationKind:      payload.OperationKindSendMessage,
			CanonicalMessageID: 7001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             2002,
			PeerSeq:            9,
			UserMessageID:      101,
			FromUserID:         1001,
			ToUserID:           2002,
			Date:               1777781234,
			Out:                true,
			MessageText:        "caption",
			MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 333},
			Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444},
			ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 3003, Date: 1777781000},
		}),
		40,
		1,
	)
	if err != nil {
		t.Fatalf("buildEventAndResponse() error = %v", err)
	}
	var event payload.MessageEventV3
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		t.Fatalf("unmarshal event payload: %v", err)
	}
	if event.SchemaVersion != payload.MessageEventSchemaVersionV3 || event.MediaRef == nil || event.Attrs == nil || event.ForwardRef == nil {
		t.Fatalf("V3 event lost media/attrs/forward: %+v", event)
	}
	var response payload.OperationResponseV2
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal response payload: %v", err)
	}
	if response.UserMessageID != 101 {
		t.Fatalf("response user_message_id = %d, want 101", response.UserMessageID)
	}
}

func TestV4CreateChatReplyProjectsParticipantsMessageAndUsersChats(t *testing.T) {
	senderID := int64(1001)
	receiverID := int64(1002)
	chatID := int64(7001)
	clientRandomID := int64(9000001)
	chatFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        chatID,
		ActorUserID:   senderID,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: senderID, Role: "creator", Date: 1_772_000_000},
			{UserID: receiverID, Role: "member", InviterUserID: senderID, Date: 1_772_000_000},
		},
	})
	if err != nil {
		t.Fatalf("WrapFact(chat participants) error = %v", err)
	}
	op := messageOperationFromV4(payload.MessageOperationV4{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		OperationKind: payload.OperationKindSendMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:      payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID: 8001,
			PeerType:           payload.PeerTypeChat,
			PeerID:             chatID,
			PeerSeq:            1,
			SenderUserID:       senderID,
			ToUserID:           receiverID,
			ClientRandomID:     clientRandomID,
			Date:               1_772_000_000,
			ServiceAction: &payload.ServiceActionRefV1{
				SchemaVersion: payload.ServiceActionSchemaVersionV1,
				Kind:          payload.ServiceActionKindChatCreate,
				Title:         "team",
				Users:         []int64{senderID, receiverID},
			},
		},
		AttachFacts: []payload.UpdateFactV1{chatFact},
	})
	op.UserMessageID = 101
	op.Out = true
	repo := NewForTest(nil, nil, "")
	repo.SetPeerProjectionClients(
		&fakeUserProjectionClient{bundle: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
			ViewerUsers: []userpb.ViewerUsersClazz{
				userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{
					ViewerUserId: senderID,
					Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: senderID}),
						tg.MakeTLUser(&tg.TLUser{Id: receiverID}),
					},
				}),
			},
		})},
		&fakeChatProjectionClient{chats: &chatpb.VectorMutableChat{
			Datas: []tg.MutableChatClazz{
				tg.MakeTLMutableChat(&tg.TLMutableChat{
					Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
						Id:                chatID,
						Creator:           senderID,
						Title:             "team",
						ParticipantsCount: 2,
						Date:              1_772_000_000,
						Version:           1,
					}),
				}),
			},
		}},
	)

	_, _, responseSchemaVersion, responsePayload, _, err := buildEventAndResponsePayloads(context.Background(), repo, ApplyUserOperationInput{
		UserID:      senderID,
		OperationID: "v4-service",
	}, op, 22, 1)
	if err != nil {
		t.Fatalf("buildEventAndResponsePayloads(V4) error = %v", err)
	}
	if responseSchemaVersion != payload.OperationResponseSchemaVersionV3 {
		t.Fatalf("response schema version = %d, want V3", responseSchemaVersion)
	}
	var response payload.OperationResponseV3
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Fatalf("unmarshal V4 response: %v", err)
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(response.ReplyEnvelope))
	if err != nil {
		t.Fatalf("decode reply envelope: %v", err)
	}
	reply, ok := obj.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("reply envelope = %T, want *tg.TLUpdates", obj)
	}
	if len(reply.Updates) != 3 {
		t.Fatalf("reply update count = %d, want 3", len(reply.Updates))
	}
	messageID, ok := reply.Updates[0].(*tg.TLUpdateMessageID)
	if !ok {
		t.Fatalf("first reply update = %T, want *tg.TLUpdateMessageID", reply.Updates[0])
	}
	if _, ok := reply.Updates[1].(*tg.TLUpdateChatParticipants); !ok {
		t.Fatalf("second reply update = %T, want *tg.TLUpdateChatParticipants", reply.Updates[1])
	}
	newMessage, ok := reply.Updates[2].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("third reply update = %T, want *tg.TLUpdateNewMessage", reply.Updates[2])
	}
	serviceMessage, ok := newMessage.Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("third reply message = %T, want *tg.TLMessageService", newMessage.Message)
	}
	if _, ok := serviceMessage.Action.(*tg.TLMessageActionChatCreate); !ok {
		t.Fatalf("service action = %T, want *tg.TLMessageActionChatCreate", serviceMessage.Action)
	}
	if messageID.Id != serviceMessage.Id || messageID.RandomId != clientRandomID {
		t.Fatalf("updateMessageID = %+v, want id=%d random_id=%d", messageID, serviceMessage.Id, clientRandomID)
	}
	if !hasUserID(reply.Users, senderID) || !hasUserID(reply.Users, receiverID) {
		t.Fatalf("reply users = %#v, want creator and invitee", reply.Users)
	}
	if !hasChatID(reply.Chats, chatID) {
		t.Fatalf("reply chats = %#v, want chat %d", reply.Chats, chatID)
	}
}

func TestEnsureExistingMessageViewMatchesOperationSupportsV3(t *testing.T) {
	event := payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		PeerSeq:            9,
		MessageID:          101,
		PeerType:           payload.PeerTypeUser,
		PeerID:             2002,
		FromUserID:         1001,
		ToUserID:           2002,
		Date:               1777781234,
		Out:                true,
		MessageText:        "caption",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 333},
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444},
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 3003, Date: 1777781000},
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal V3 event: %v", err)
	}
	row := &model.UserMessageViews{
		PeerType:           event.PeerType,
		PeerId:             event.PeerID,
		PeerSeq:            event.PeerSeq,
		CanonicalMessageId: event.CanonicalMessageID,
		UserMessageId:      event.MessageID,
		FromUserId:         event.FromUserID,
		Outgoing:           event.Out,
		MessageStatus:      MessageStatusLive,
		Date:               int64(event.Date),
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV3,
		ViewPayload:        body,
	}
	op := messageOperationFromV3(payload.MessageOperationV3{
		SchemaVersion:      payload.MessageOperationSchemaVersionV3,
		OperationKind:      payload.OperationKindSendMessage,
		CanonicalMessageID: event.CanonicalMessageID,
		PeerType:           event.PeerType,
		PeerID:             event.PeerID,
		PeerSeq:            event.PeerSeq,
		UserMessageID:      event.MessageID,
		FromUserID:         event.FromUserID,
		ToUserID:           event.ToUserID,
		Date:               event.Date,
		Out:                event.Out,
		MessageText:        event.MessageText,
		MediaRef:           event.MediaRef,
		Attrs:              event.Attrs,
		ForwardRef:         event.ForwardRef,
	})
	if err := ensureExistingMessageViewMatchesOperation(row, op); err != nil {
		t.Fatalf("ensureExistingMessageViewMatchesOperation() error = %v", err)
	}
	op.MediaRef = &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo", ID: 999}
	if err := ensureExistingMessageViewMatchesOperation(row, op); !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("conflict error = %v, want ErrOperationPayloadConflict", err)
	}
}

func TestExtractHashTagsNormalizesAndDeduplicates(t *testing.T) {
	got := extractHashTags("hello #Go #go #team_gram #中文 ok #")
	want := []string{"go", "team_gram", "中文"}
	if len(got) != len(want) {
		t.Fatalf("extractHashTags len = %d, want %d: %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("extractHashTags[%d] = %q, want %q; all=%#v", i, got[i], want[i], got)
		}
	}
}

func TestValidateAffectedOutboxRejectsPayloadHashMismatch(t *testing.T) {
	err := validateAffectedOutbox(AffectedOutbox{
		TargetUserID:   1001,
		OperationID:    "affected-op",
		OperationKind:  payload.OperationKindSendMessage,
		DeliveryPolicy: DeliveryPolicyDurableAsync,
		PayloadCodec:   PayloadCodecJSON,
		Payload:        []byte(`{"ok":true}`),
		PayloadHash:    []byte("wrong"),
	})
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("validateAffectedOutbox() error = %v, want ErrOperationPayloadConflict", err)
	}
}
