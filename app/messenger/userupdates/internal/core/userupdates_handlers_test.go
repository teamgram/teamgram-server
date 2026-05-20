package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/projection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload/serviceaction"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mustServiceActionRef(t *testing.T, action tg.MessageActionClazz) *payload.ServiceActionRefV1 {
	t.Helper()
	ref, err := serviceaction.Encode(action)
	if err != nil {
		t.Fatalf("serviceaction.Encode() error = %v", err)
	}
	return ref
}

func TestMessageServiceActionNilRefReturnsNil(t *testing.T) {
	action, err := messageServiceAction(nil)
	if err != nil {
		t.Fatalf("messageServiceAction(nil) error = %v", err)
	}
	if action != nil {
		t.Fatalf("messageServiceAction(nil) action = %T, want nil", action)
	}
}

func TestMessageServiceActionDecodeFailureWrapsStorage(t *testing.T) {
	_, err := messageServiceAction(&payload.ServiceActionRefV1{
		SchemaVersion: payload.ServiceActionSchemaVersionV1,
		Codec:         payload.ServiceActionCodecTLBinary,
		Layer:         payload.ServiceActionLayer,
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("messageServiceAction() error = %v, want ErrUserupdatesStorage", err)
	}
}

func TestProcessUserOperationMapsTLToRepository(t *testing.T) {
	operationPayload := []byte(`{"schema_version":1,"operation_kind":"send_message"}`)
	operationHash := payload.HashBytes(operationPayload)
	responsePayload := []byte(`{"schema_version":1,"pts":12,"pts_count":1}`)
	responseHash := payload.HashBytes(responsePayload)

	repo := &fakeUserUpdatesRepository{
		applyResult: &repository.ApplyUserOperationResult{
			UserID:          1001,
			OperationID:     "v1:msg:2001:sender:1001:out",
			Pts:             12,
			PtsCount:        1,
			ResponsePayload: responsePayload,
			ResponseHash:    responseHash,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})
	authKeyIDExclude := int64(9001)

	got, err := core.UserupdatesProcessUserOperation(&userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               1001,
			BucketId:             77,
			PartitionId:          13,
			OperationId:          "v1:msg:2001:sender:1001:out",
			OpType:               repository.OpTypeSendMessage,
			PeerType:             payload.PeerTypeUser,
			PeerId:               1002,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         repository.PayloadCodecJSON,
			PayloadHash:          operationHash,
			Payload:              operationPayload,
			AuthKeyIdExclude:     &authKeyIDExclude,
		}),
	})
	if err != nil {
		t.Fatalf("ProcessUserOperation returned error: %v", err)
	}
	if got.Pts != 12 || got.PtsCount != 1 || got.CurrentPts != 12 {
		t.Fatalf("unexpected pts result: pts=%d pts_count=%d current_pts=%d", got.Pts, got.PtsCount, got.CurrentPts)
	}
	if got.ResponseSchemaVersion == nil || *got.ResponseSchemaVersion != payload.OperationResponseSchemaVersion {
		t.Fatalf("unexpected response schema version: %v", got.ResponseSchemaVersion)
	}
	if string(got.ResponsePayload) != string(responsePayload) {
		t.Fatalf("unexpected response payload: %s", string(got.ResponsePayload))
	}
	if !bytes.Equal(got.ResponsePayloadHash, responseHash) {
		t.Fatalf("unexpected response hash: %x", got.ResponsePayloadHash)
	}
	if repo.applyInput.UserID != 1001 ||
		repo.applyInput.OperationID != "v1:msg:2001:sender:1001:out" ||
		!bytes.Equal(repo.applyInput.PayloadHash, operationHash) ||
		repo.applyInput.BucketID != 77 ||
		repo.applyInput.PartitionID != 13 ||
		repo.applyInput.AuthKeyIDExclude == nil ||
		*repo.applyInput.AuthKeyIDExclude != authKeyIDExclude {
		t.Fatalf("unexpected repository input: %+v", repo.applyInput)
	}
}

func TestProcessUserOperationWakesPushOutboxNotifier(t *testing.T) {
	operationPayload := []byte(`{"schema_version":1,"operation_kind":"read_history"}`)
	operationHash := payload.HashBytes(operationPayload)
	responsePayload := []byte(`{"schema_version":1,"pts":13,"pts_count":1}`)
	responseHash := payload.HashBytes(responsePayload)
	notifier := &fakePushOutboxNotifier{}
	repo := &fakeUserUpdatesRepository{
		applyResult: &repository.ApplyUserOperationResult{
			UserID:          1002,
			OperationID:     "v1:dialog:read_history:user:1002:peer:1001:max:9:auth:0",
			Pts:             13,
			PtsCount:        1,
			ResponsePayload: responsePayload,
			ResponseHash:    responseHash,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, PushOutboxNotifier: notifier})

	if _, err := core.UserupdatesProcessUserOperation(&userupdates.TLUserupdatesProcessUserOperation{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               1002,
			BucketId:             77,
			PartitionId:          13,
			OperationId:          "v1:dialog:read_history:user:1002:peer:1001:max:9:auth:0",
			OpType:               repository.OpTypeSendMessage,
			PeerType:             payload.PeerTypeUser,
			PeerId:               1001,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         repository.PayloadCodecJSON,
			PayloadHash:          operationHash,
			Payload:              operationPayload,
		}),
	}); err != nil {
		t.Fatalf("ProcessUserOperation returned error: %v", err)
	}
	if notifier.wakes != 1 {
		t.Fatalf("notifier wakes = %d, want 1", notifier.wakes)
	}
}

func TestMessageEventV3ToTLMessagePreservesEntities(t *testing.T) {
	message, err := messageEventV3ToTLMessage(payload.MessageEventV3{
		SchemaVersion: payload.MessageEventSchemaVersionV3,
		EventKind:     payload.EventKindNewMessage,
		MessageID:     11,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1002,
		FromUserID:    1001,
		Date:          1700000000,
		MessageText:   "@alice",
		Entities: []payload.MessageEntityV1{
			{Offset: 0, Length: 6, Kind: "mention"},
		},
	})
	if err != nil {
		t.Fatalf("messageEventV3ToTLMessage() error = %v", err)
	}
	tlMessage, ok := message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", message)
	}
	if len(tlMessage.Entities) != 1 {
		t.Fatalf("entities len = %d, want 1", len(tlMessage.Entities))
	}
	if _, ok := tlMessage.Entities[0].(*tg.TLMessageEntityMention); !ok {
		t.Fatalf("entity = %T, want mention", tlMessage.Entities[0])
	}
}

func TestMessageEventV3ToTLMessageProjectsChatCreateServiceAction(t *testing.T) {
	message, err := messageEventV3ToTLMessage(payload.MessageEventV3{
		SchemaVersion: payload.MessageEventSchemaVersionV3,
		EventKind:     payload.EventKindNewMessage,
		MessageID:     11,
		PeerType:      payload.PeerTypeChat,
		PeerID:        55,
		FromUserID:    1001,
		Date:          1700000000,
		ServiceAction: mustServiceActionRef(t, tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
			Title: "new chat",
			Users: []int64{1002, 1003},
		})),
	})
	if err != nil {
		t.Fatalf("messageEventV3ToTLMessage() error = %v", err)
	}
	service, ok := message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessageService", message)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatCreate)
	if !ok || action.Title != "new chat" || len(action.Users) != 2 {
		t.Fatalf("action = %T %+v, want chat create title/users", service.Action, service.Action)
	}
}

func TestMessageEventV3ToTLMessageProjectsChatAddUserServiceAction(t *testing.T) {
	message, err := messageEventV3ToTLMessage(payload.MessageEventV3{
		SchemaVersion: payload.MessageEventSchemaVersionV3,
		EventKind:     payload.EventKindNewMessage,
		MessageID:     11,
		PeerType:      payload.PeerTypeChat,
		PeerID:        55,
		FromUserID:    1001,
		Date:          1700000000,
		ServiceAction: mustServiceActionRef(t, tg.MakeTLMessageActionChatAddUser(&tg.TLMessageActionChatAddUser{
			Users: []int64{1002},
		})),
	})
	if err != nil {
		t.Fatalf("messageEventV3ToTLMessage() error = %v", err)
	}
	service, ok := message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessageService", message)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatAddUser)
	if !ok || len(action.Users) != 1 || action.Users[0] != 1002 {
		t.Fatalf("action = %T %+v, want chat add user [1002]", service.Action, service.Action)
	}
}

func TestMessageEventV3ToTLMessageProjectsGroupCallServiceAction(t *testing.T) {
	duration := int32(42)
	message, err := messageEventV3ToTLMessage(payload.MessageEventV3{
		SchemaVersion: payload.MessageEventSchemaVersionV3,
		EventKind:     payload.EventKindNewMessage,
		MessageID:     11,
		PeerType:      payload.PeerTypeChat,
		PeerID:        55,
		FromUserID:    1001,
		Date:          1700000000,
		ServiceAction: mustServiceActionRef(t, tg.MakeTLMessageActionGroupCall(&tg.TLMessageActionGroupCall{
			Call: tg.MakeTLInputGroupCall(&tg.TLInputGroupCall{
				Id:         7001,
				AccessHash: 8002,
			}),
			Duration: &duration,
		})),
	})
	if err != nil {
		t.Fatalf("messageEventV3ToTLMessage() error = %v", err)
	}
	service, ok := message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessageService", message)
	}
	action, ok := service.Action.(*tg.TLMessageActionGroupCall)
	if !ok {
		t.Fatalf("action = %T, want *tg.TLMessageActionGroupCall", service.Action)
	}
	call, ok := action.Call.(*tg.TLInputGroupCall)
	if !ok {
		t.Fatalf("call = %T, want *tg.TLInputGroupCall", action.Call)
	}
	if call.Id != 7001 || call.AccessHash != 8002 {
		t.Fatalf("call = %+v, want 7001/8002", call)
	}
	if action.Duration == nil || *action.Duration != 42 {
		t.Fatalf("duration = %v, want 42", action.Duration)
	}
}

func TestProcessUserOperationWithEffectsMapsAffectedOutboxes(t *testing.T) {
	requestPayload := []byte(`{"schema_version":1,"operation_kind":"read_history","out":false}`)
	requestHash := payload.HashBytes(requestPayload)
	effectPayload := []byte(`{"schema_version":1,"operation_kind":"read_history","out":true}`)
	effectHash := payload.HashBytes(effectPayload)
	responsePayload := []byte(`{"schema_version":1,"pts":15,"pts_count":1}`)
	responseHash := payload.HashBytes(responsePayload)
	notifier := &fakePushOutboxNotifier{}
	repo := &fakeUserUpdatesRepository{
		applyResult: &repository.ApplyUserOperationResult{
			UserID:          1001,
			OperationID:     "v1:dialog:read_history:user:1001:peer:1002:max:9:auth:0",
			Pts:             15,
			PtsCount:        1,
			ResponsePayload: responsePayload,
			ResponseHash:    responseHash,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, PushOutboxNotifier: notifier})

	got, err := core.UserupdatesProcessUserOperationWithEffects(&userupdates.TLUserupdatesProcessUserOperationWithEffects{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               1001,
			BucketId:             7,
			PartitionId:          3,
			OperationId:          "v1:dialog:read_history:user:1001:peer:1002:max:9:auth:0",
			OpType:               repository.OpTypeSendMessage,
			PeerType:             payload.PeerTypeUser,
			PeerId:               1002,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         repository.PayloadCodecJSON,
			PayloadHash:          requestHash,
			Payload:              requestPayload,
		}),
		AffectedEffects: []userupdates.AffectedUserOperationClazz{
			userupdates.MakeTLAffectedUserOperation(&userupdates.TLAffectedUserOperation{
				RequesterUserId: 1001,
				DeliveryPolicy:  repository.DeliveryPolicyDurableAsync,
				OperationKind:   payload.OperationKindReadHistory,
				Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
					UserId:               1002,
					BucketId:             8,
					PartitionId:          4,
					OperationId:          "v1:dialog:read_history:user:1002:peer:1001:max:9:auth:0",
					OpType:               repository.OpTypeSendMessage,
					PeerType:             payload.PeerTypeUser,
					PeerId:               1001,
					PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
					PayloadCodec:         repository.PayloadCodecJSON,
					PayloadHash:          effectHash,
					Payload:              effectPayload,
				}),
			}),
		},
	})
	if err != nil {
		t.Fatalf("ProcessUserOperationWithEffects returned error: %v", err)
	}
	if got.Pts != 15 || got.PtsCount != 1 || got.CurrentPts != 15 {
		t.Fatalf("unexpected pts result: pts=%d pts_count=%d current_pts=%d", got.Pts, got.PtsCount, got.CurrentPts)
	}
	if repo.applyInput.UserID != 1001 ||
		repo.applyInput.OperationID != "v1:dialog:read_history:user:1001:peer:1002:max:9:auth:0" ||
		!bytes.Equal(repo.applyInput.PayloadHash, requestHash) ||
		repo.applyInput.BucketID != 7 ||
		repo.applyInput.PartitionID != 3 {
		t.Fatalf("unexpected requester repository input: %+v", repo.applyInput)
	}
	if len(repo.applyInput.AffectedOutboxes) != 1 {
		t.Fatalf("affected outbox count = %d, want 1", len(repo.applyInput.AffectedOutboxes))
	}
	gotOutbox := repo.applyInput.AffectedOutboxes[0]
	if gotOutbox.RequesterUserID != 1001 ||
		gotOutbox.TargetUserID != 1002 ||
		gotOutbox.TargetBucketID != 8 ||
		gotOutbox.TargetPartitionID != 4 ||
		gotOutbox.OperationID != "v1:dialog:read_history:user:1002:peer:1001:max:9:auth:0" ||
		gotOutbox.OperationKind != payload.OperationKindReadHistory ||
		gotOutbox.PeerType != payload.PeerTypeUser ||
		gotOutbox.PeerID != 1001 ||
		gotOutbox.PayloadCodec != repository.PayloadCodecJSON ||
		!bytes.Equal(gotOutbox.Payload, effectPayload) ||
		!bytes.Equal(gotOutbox.PayloadHash, effectHash) ||
		gotOutbox.DeliveryPolicy != repository.DeliveryPolicyDurableAsync {
		t.Fatalf("unexpected affected outbox: %+v", gotOutbox)
	}
	if notifier.wakes != 1 {
		t.Fatalf("notifier wakes = %d, want 1", notifier.wakes)
	}
}

func TestProcessUserOperationWithEffectsRejectsUnsupportedPolicy(t *testing.T) {
	requestPayload := []byte(`{"schema_version":1,"operation_kind":"read_history","out":false}`)
	requestHash := payload.HashBytes(requestPayload)
	effectPayload := []byte(`{"schema_version":1,"operation_kind":"read_history","out":true}`)
	effectHash := payload.HashBytes(effectPayload)
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeUserUpdatesRepository{}})

	_, err := core.UserupdatesProcessUserOperationWithEffects(&userupdates.TLUserupdatesProcessUserOperationWithEffects{
		Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:               1001,
			BucketId:             7,
			PartitionId:          3,
			OperationId:          "requester",
			OpType:               repository.OpTypeSendMessage,
			PeerType:             payload.PeerTypeUser,
			PeerId:               1002,
			PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
			PayloadCodec:         repository.PayloadCodecJSON,
			PayloadHash:          requestHash,
			Payload:              requestPayload,
		}),
		AffectedEffects: []userupdates.AffectedUserOperationClazz{
			userupdates.MakeTLAffectedUserOperation(&userupdates.TLAffectedUserOperation{
				RequesterUserId: 1001,
				DeliveryPolicy:  repository.DeliveryPolicyBrokerDurableAck,
				OperationKind:   payload.OperationKindReadHistory,
				Operation: userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
					UserId:               1002,
					BucketId:             8,
					PartitionId:          4,
					OperationId:          "affected",
					OpType:               repository.OpTypeSendMessage,
					PeerType:             payload.PeerTypeUser,
					PeerId:               1001,
					PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
					PayloadCodec:         repository.PayloadCodecJSON,
					PayloadHash:          effectHash,
					Payload:              effectPayload,
				}),
			}),
		},
	})
	if !errors.Is(err, userupdates.ErrOperationTerminal) {
		t.Fatalf("expected ErrOperationTerminal, got %v", err)
	}
}

func TestUserupdatesProcessUserOperationBatchRejectsOversizedBeforeRepositoryApply(t *testing.T) {
	repo := &fakeUserUpdatesRepository{}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})
	operations := make([]userupdates.UserOperationClazz, MaxUserOperationBatchSize+1)
	for i := range operations {
		operations[i] = userupdates.MakeTLUserOperation(&userupdates.TLUserOperation{
			UserId:      1001,
			OperationId: "oversized",
		})
	}

	_, err := core.UserupdatesProcessUserOperationBatch(&userupdates.TLUserupdatesProcessUserOperationBatch{
		Operations: operations,
	})
	if !errors.Is(err, userupdates.ErrOperationTerminal) {
		t.Fatalf("UserupdatesProcessUserOperationBatch() error = %v, want ErrOperationTerminal", err)
	}
	if repo.applyBatchCalled {
		t.Fatal("repository batch apply called for oversized request")
	}
}

func TestGetOperationResultRejectsMismatchedPayloadHash(t *testing.T) {
	goodPayload := []byte(`{"good":true}`)
	badPayload := []byte(`{"good":false}`)
	repo := &fakeUserUpdatesRepository{
		operationResult: &repository.OperationResult{
			UserID:      1001,
			OperationID: "v1:msg:2001:receiver:1001:in",
			Status:      repository.OperationResultStatusCompleted,
			Pts:         8,
			PtsCount:    1,
			PayloadHash: payload.HashBytes(goodPayload),
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetOperationResult(&userupdates.TLUserupdatesGetOperationResult{
		UserId:      1001,
		OperationId: "v1:msg:2001:receiver:1001:in",
		PayloadHash: payload.HashBytes(badPayload),
	})
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("expected ErrOperationPayloadConflict, got %v", err)
	}
}

func TestGetOperationResultPreservesLegacyResponseSchema(t *testing.T) {
	responsePayload := mustMarshalOperationResponseV1(t, payload.OperationResponseV1{
		SchemaVersion: payload.OperationResponseSchemaVersionV1,
		OperationID:   "v1:msg:2001:receiver:1001:in",
		Pts:           8,
		PtsCount:      1,
		EventType:     payload.EventKindNewMessage,
	})
	responseHash := payload.HashBytes(responsePayload)
	repo := &fakeUserUpdatesRepository{
		operationResult: &repository.OperationResult{
			UserID:                1001,
			OperationID:           "v1:msg:2001:receiver:1001:in",
			Status:                repository.OperationResultStatusCompleted,
			Pts:                   8,
			PtsCount:              1,
			PayloadHash:           []byte("payload-hash"),
			ResponseSchemaVersion: payload.OperationResponseSchemaVersionV1,
			ResponsePayload:       responsePayload,
			ResponseHash:          responseHash,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetOperationResult(&userupdates.TLUserupdatesGetOperationResult{
		UserId:      1001,
		OperationId: "v1:msg:2001:receiver:1001:in",
		PayloadHash: []byte("payload-hash"),
	})
	if err != nil {
		t.Fatalf("GetOperationResult returned error: %v", err)
	}
	if got.ResponseSchemaVersion == nil || *got.ResponseSchemaVersion != payload.OperationResponseSchemaVersionV1 {
		t.Fatalf("response schema version = %v, want %d", got.ResponseSchemaVersion, payload.OperationResponseSchemaVersionV1)
	}
	if !bytes.Equal(got.ResponsePayload, responsePayload) || !bytes.Equal(got.ResponsePayloadHash, responseHash) {
		t.Fatalf("unexpected response payload/hash: payload=%s hash=%x", string(got.ResponsePayload), got.ResponsePayloadHash)
	}
}

func TestGetDifferenceBuildsVisibleMessageFromEventPayload(t *testing.T) {
	eventPayload := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   2001,
		PeerSeq:              9,
		MessageID:            9,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		FromUserID:           1002,
		ToUserID:             1001,
		Date:                 1_772_000_000,
		Out:                  false,
		MessageText:          "hello from event payload",
		ReplyToUserMessageID: 55,
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
			Events: []repository.UserEvent{
				{
					UserID:             1001,
					Pts:                18,
					PtsCount:           1,
					OperationID:        "v1:msg:2001:receiver:1001:in",
					EventType:          repository.EventTypeNewMessage,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1002,
					CanonicalMessageID: 2001,
					PeerSeq:            9,
					ActorUserID:        1002,
					EventSchemaVersion: payload.MessageEventSchemaVersion,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if diff.State == nil || diff.State.Pts != 18 {
		t.Fatalf("unexpected state: %#v", diff.State)
	}
	if len(diff.NewMessages) != 1 {
		t.Fatalf("expected one new message, got %d", len(diff.NewMessages))
	}
	message, ok := diff.NewMessages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected TLMessage, got %T", diff.NewMessages[0])
	}
	if message.Id != 9 || message.Message != "hello from event payload" || message.Out {
		t.Fatalf("unexpected message projection: %+v", message)
	}
	reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("reply header = %T, want *tg.TLMessageReplyHeader", message.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 55 {
		t.Fatalf("reply_to_msg_id = %v, want receiver-local 55", reply.ReplyToMsgId)
	}
	if len(diff.OtherUpdates) != 0 {
		t.Fatalf("other update count = %d, want no duplicate updateNewMessage", len(diff.OtherUpdates))
	}
}

func TestV4CreateChatDifferenceOmitsUpdateMessageID(t *testing.T) {
	chatFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        2002,
		ActorUserID:   1001,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 1001, Role: "creator", Date: 1_772_000_000},
			{UserID: 1002, Role: "member", InviterUserID: 1001, Date: 1_772_000_000},
		},
	})
	if err != nil {
		t.Fatalf("WrapFact(chat participants) error = %v", err)
	}
	eventPayload := mustMarshalMessageEventV4(t, payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:      payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID: 2001,
			PeerType:           payload.PeerTypeChat,
			PeerID:             2002,
			PeerSeq:            1,
			SenderUserID:       1001,
			Date:               1_772_000_000,
			ServiceAction: mustServiceActionRef(t, tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
				Title: "v4 chat",
				Users: []int64{1001, 1002},
			})),
		},
		AttachFacts: []payload.UpdateFactV1{chatFact},
		MessageID:   101,
		Pts:         18,
		PtsCount:    1,
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
			Events: []repository.UserEvent{
				{
					UserID:             1001,
					Pts:                18,
					PtsCount:           1,
					OperationID:        "v4-chat-create",
					EventType:          repository.EventTypeNewMessage,
					PeerType:           payload.PeerTypeChat,
					PeerID:             2002,
					CanonicalMessageID: 2001,
					PeerSeq:            1,
					ActorUserID:        1001,
					EventSchemaVersion: payload.MessageEventSchemaVersionV4,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 1 {
		t.Fatalf("new message count = %d, want 1", len(diff.NewMessages))
	}
	service, ok := diff.NewMessages[0].(*tg.TLMessageService)
	if !ok {
		t.Fatalf("new message = %T, want *tg.TLMessageService", diff.NewMessages[0])
	}
	if _, ok := service.Action.(*tg.TLMessageActionChatCreate); !ok {
		t.Fatalf("service action = %T, want *tg.TLMessageActionChatCreate", service.Action)
	}
	if len(diff.OtherUpdates) != 1 {
		t.Fatalf("other update count = %d, want only non-message companion updates", len(diff.OtherUpdates))
	}
	for i, update := range diff.OtherUpdates {
		if _, ok := update.(*tg.TLUpdateMessageID); ok {
			t.Fatalf("other_updates[%d] = *tg.TLUpdateMessageID, want difference to omit reply-only update", i)
		}
		if _, ok := update.(*tg.TLUpdateNewMessage); ok {
			t.Fatalf("other_updates[%d] = *tg.TLUpdateNewMessage, want message only in new_messages", i)
		}
	}
	if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateChatParticipants); !ok {
		t.Fatalf("first update = %T, want *tg.TLUpdateChatParticipants", diff.OtherUpdates[0])
	}
}

func TestBatchMessageDifferenceReturnsAllNewMessages(t *testing.T) {
	eventPayload := mustMarshalMessageEventBatchV1(t, payload.MessageEventBatchV1{
		SchemaVersion: payload.MessageEventSchemaVersionBatchV1,
		EventKind:     payload.EventKindNewMessage,
		Messages: []payload.MessageEventBatchItemV1{
			{
				MessageFact: payload.NewMessageFactV1{
					SchemaVersion:      1,
					CanonicalMessageID: 8001,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1002,
					SenderUserID:       1002,
					ToUserID:           1001,
					Date:               1_772_000_000,
					MessageText:        "first",
				},
				MessageID: 51,
				Pts:       18,
				PtsCount:  1,
			},
			{
				MessageFact: payload.NewMessageFactV1{
					SchemaVersion:      1,
					CanonicalMessageID: 8002,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1002,
					SenderUserID:       1002,
					ToUserID:           1001,
					Date:               1_772_000_001,
					MessageText:        "second",
				},
				MessageID: 52,
				Pts:       19,
				PtsCount:  1,
			},
		},
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 19},
			Events: []repository.UserEvent{
				{
					UserID:             1001,
					Pts:                19,
					PtsCount:           2,
					OperationID:        "v1:msgbatch:receiver:1001:in:8001:8002",
					EventType:          repository.EventTypeNewMessage,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1002,
					CanonicalMessageID: 8001,
					PeerSeq:            1,
					ActorUserID:        1002,
					EventSchemaVersion: payload.MessageEventSchemaVersionBatchV1,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 2 {
		t.Fatalf("new message count = %d, want 2", len(diff.NewMessages))
	}
	first, ok := diff.NewMessages[0].(*tg.TLMessage)
	if !ok || first.Id != 51 || first.Message != "first" {
		t.Fatalf("first new message = %#v", diff.NewMessages[0])
	}
	second, ok := diff.NewMessages[1].(*tg.TLMessage)
	if !ok || second.Id != 52 || second.Message != "second" {
		t.Fatalf("second new message = %#v", diff.NewMessages[1])
	}
	if len(diff.OtherUpdates) != 0 {
		t.Fatalf("other update count = %d, want no duplicate updateNewMessage", len(diff.OtherUpdates))
	}
}

func TestGetDifferenceBuildsReadHistoryOutboxUpdate(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.OperationKindReadHistory,
		MessageID:     66,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1571766987,
		Date:          1_778_029_828,
		Out:           true,
	})
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1571766986, Pts: 157},
			Events: []repository.UserEvent{
				{
					UserID:             1571766986,
					Pts:                157,
					PtsCount:           1,
					OperationID:        "read-outbox",
					EventType:          repository.EventTypeReadHistory,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1571766987,
					PeerSeq:            66,
					ActorUserID:        1571766987,
					EventSchemaVersion: payload.MessageEventSchemaVersion,
					EventCodec:         repository.PayloadCodecJSON,
					EventPayload:       eventPayload,
					EventPayloadHash:   payload.HashBytes(eventPayload),
				},
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1571766986,
		AuthKeyId:     9002,
		Pts:           156,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 0 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("unexpected difference lens: new=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
	update, ok := diff.OtherUpdates[0].(*tg.TLUpdateReadHistoryOutbox)
	if !ok {
		t.Fatalf("expected TLUpdateReadHistoryOutbox, got %T", diff.OtherUpdates[0])
	}
	peer, ok := update.Peer.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1571766987 {
		t.Fatalf("unexpected peer: %+v ok=%v", update.Peer, ok)
	}
	if update.MaxId != 66 || update.Pts != 157 || update.PtsCount != 1 {
		t.Fatalf("unexpected read outbox update: %+v", update)
	}
}

func TestGetMessageViewsByPeerSeqsBuildsMessagesFromViews(t *testing.T) {
	eventPayload := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:        payload.MessageEventSchemaVersion,
		EventKind:            payload.EventKindNewMessage,
		CanonicalMessageID:   2001,
		PeerSeq:              7,
		MessageID:            101,
		PeerType:             payload.PeerTypeUser,
		PeerID:               1002,
		FromUserID:           1001,
		ToUserID:             1002,
		Date:                 1_772_000_000,
		Out:                  true,
		MessageText:          "dialog top",
		ReplyToUserMessageID: 88,
	})
	peer := repository.MessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerID: 1002, PeerSeq: 7}
	repo := &fakeUserUpdatesRepository{
		messageViews: map[repository.MessageViewPeerSeq]repository.MessageView{
			peer: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				PeerSeq:            7,
				UserMessageID:      101,
				CanonicalMessageID: 2001,
				FromUserID:         1001,
				Outgoing:           true,
				MessageStatus:      repository.MessageStatusLive,
				ViewSchemaVersion:  payload.MessageEventSchemaVersion,
				ViewPayload:        eventPayload,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetMessageViewsByPeerSeqs(&userupdates.TLUserupdatesGetMessageViewsByPeerSeqs{
		UserId: 1001,
		Peers: []userupdates.MessageViewPeerSeqClazz{
			userupdates.MakeTLMessageViewPeerSeq(&userupdates.TLMessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerId: 1002, PeerSeq: 7}),
		},
	})
	if err != nil {
		t.Fatalf("GetMessageViewsByPeerSeqs returned error: %v", err)
	}
	if repo.messageViewUserID != 1001 || len(repo.messageViewPeers) != 1 || repo.messageViewPeers[0] != peer {
		t.Fatalf("unexpected repository message view request: user_id=%d peers=%+v", repo.messageViewUserID, repo.messageViewPeers)
	}
	if got == nil || len(got.Messages) != 1 {
		t.Fatalf("expected one message, got %+v", got)
	}
	message, ok := got.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("expected TLMessage, got %T", got.Messages[0])
	}
	if message.Id != 101 || message.Message != "dialog top" || !message.Out {
		t.Fatalf("unexpected message view projection: %+v", message)
	}
	reply, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("reply header = %T, want *tg.TLMessageReplyHeader", message.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 88 {
		t.Fatalf("reply_to_msg_id = %v, want public id 88", reply.ReplyToMsgId)
	}
}

func TestGetMessageViewsByPeerSeqsOmitsFromIDForIncomingPrivateTopMessage(t *testing.T) {
	eventPayload := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2002,
		PeerSeq:            8,
		MessageID:          102,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1001,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_001,
		Out:                false,
		MessageText:        "incoming dialog top",
	})
	peer := repository.MessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerID: 1001, PeerSeq: 8}
	repo := &fakeUserUpdatesRepository{
		messageViews: map[repository.MessageViewPeerSeq]repository.MessageView{
			peer: {
				UserID:             1002,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				PeerSeq:            8,
				UserMessageID:      102,
				CanonicalMessageID: 2002,
				FromUserID:         1001,
				Outgoing:           false,
				MessageStatus:      repository.MessageStatusLive,
				ViewSchemaVersion:  payload.MessageEventSchemaVersionV3,
				ViewPayload:        eventPayload,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetMessageViewsByPeerSeqs(&userupdates.TLUserupdatesGetMessageViewsByPeerSeqs{
		UserId: 1002,
		Peers: []userupdates.MessageViewPeerSeqClazz{
			userupdates.MakeTLMessageViewPeerSeq(&userupdates.TLMessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerId: 1001, PeerSeq: 8}),
		},
	})
	if err != nil {
		t.Fatalf("GetMessageViewsByPeerSeqs returned error: %v", err)
	}
	if got == nil || len(got.Messages) != 1 {
		t.Fatalf("expected one message, got %+v", got)
	}
	message, ok := got.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", got.Messages[0])
	}
	if message.Out {
		t.Fatalf("message.Out = true, want false for incoming private top message")
	}
	if message.FromId != nil {
		t.Fatalf("message.FromId = %#v, want nil for incoming private top message", message.FromId)
	}
	peerID, ok := message.PeerId.(*tg.TLPeerUser)
	if !ok || peerID.UserId != 1001 {
		t.Fatalf("message.PeerId = %#v, want peerUser(1001)", message.PeerId)
	}
}

func TestMessageViewToTLMessageSupportsLegacyV1PeerSeqID(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersionV1,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2001,
		MessageID:          7,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		Out:                true,
		MessageText:        "legacy top",
		ReplyToPeerSeq:     5,
	})
	message, err := messageViewToTLMessage(repository.MessageView{
		UserID:             1001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		PeerSeq:            7,
		CanonicalMessageID: 2001,
		MessageStatus:      repository.MessageStatusLive,
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV1,
		ViewPayload:        eventPayload,
	})
	if err != nil {
		t.Fatalf("messageViewToTLMessage error = %v", err)
	}
	tlMessage, ok := message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", message)
	}
	if tlMessage.Id != 7 || tlMessage.Message != "legacy top" {
		t.Fatalf("unexpected legacy message: %+v", tlMessage)
	}
	reply, ok := tlMessage.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 5 {
		t.Fatalf("legacy reply header = %+v ok=%v, want peer seq 5", reply, ok)
	}
}

func TestEventToTLUpdateBuildsEditMessage(t *testing.T) {
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.OperationKindEditMessage,
		CanonicalMessageID: 2001,
		MessageID:          7,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		EditDate:           1_772_000_100,
		EditVersion:        1,
		Out:                true,
		MessageText:        "edited",
	})
	event := repository.UserEvent{
		UserID:             1001,
		Pts:                50,
		PtsCount:           1,
		EventType:          repository.EventTypeEditMessage,
		EventSchemaVersion: payload.MessageEventSchemaVersion,
		EventCodec:         repository.PayloadCodecJSON,
		EventPayload:       eventPayload,
		EventPayloadHash:   payload.HashBytes(eventPayload),
	}

	projected, err := projection.ProjectUserEvent(event, projection.ModeDifference)
	if err != nil {
		t.Fatalf("ProjectUserEvent error = %v", err)
	}
	if projected.Message != nil {
		t.Fatalf("message = %T, want nil for edit update", projected.Message)
	}
	edit, ok := projected.Update.(*tg.TLUpdateEditMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateEditMessage", projected.Update)
	}
	if edit.Pts != 50 || edit.PtsCount != 1 {
		t.Fatalf("unexpected edit update pts: %+v", edit)
	}
	tlMessage, ok := edit.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("edit message = %T, want *tg.TLMessage", edit.Message)
	}
	if tlMessage.Id != 7 || tlMessage.Message != "edited" || tlMessage.EditDate == nil || *tlMessage.EditDate != 1_772_000_100 {
		t.Fatalf("unexpected edit message: %+v", tlMessage)
	}
}

func TestMessageViewToTLMessageAcceptsEditPayload(t *testing.T) {
	eventPayload := mustMarshalMessageEventV2(t, payload.MessageEventV2{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.OperationKindEditMessage,
		CanonicalMessageID: 2001,
		PeerSeq:            7,
		MessageID:          101,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		EditDate:           1_772_000_100,
		Out:                true,
		MessageText:        "edited",
	})
	message, err := messageViewToTLMessage(repository.MessageView{
		UserID:             1001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		PeerSeq:            7,
		UserMessageID:      101,
		CanonicalMessageID: 2001,
		MessageStatus:      repository.MessageStatusLive,
		ViewSchemaVersion:  payload.MessageEventSchemaVersion,
		ViewPayload:        eventPayload,
	})
	if err != nil {
		t.Fatalf("messageViewToTLMessage error = %v", err)
	}
	tlMessage, ok := message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", message)
	}
	if tlMessage.Id != 101 || tlMessage.Message != "edited" || tlMessage.EditDate == nil || *tlMessage.EditDate != 1_772_000_100 {
		t.Fatalf("unexpected message: %+v", tlMessage)
	}
}

func TestMessageViewToTLMessageSupportsV3MediaAttrsForward(t *testing.T) {
	eventPayload := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2002,
		PeerSeq:            8,
		MessageID:          102,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		Out:                true,
		MessageText:        "caption",
		MediaRef:           &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document", ID: 333},
		Attrs:              &payload.MessageAttrsV1{SchemaVersion: payload.MessageAttrsSchemaVersionV1, GroupedID: 444, Noforwards: true},
		ForwardRef:         &payload.ForwardRefV1{SchemaVersion: payload.ForwardRefSchemaVersionV1, FromUserID: 3003, Date: 1_772_000_001},
	})
	message, err := messageViewToTLMessage(repository.MessageView{
		UserID:             1001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		PeerSeq:            8,
		UserMessageID:      102,
		CanonicalMessageID: 2002,
		MessageStatus:      repository.MessageStatusLive,
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV3,
		ViewPayload:        eventPayload,
	})
	if err != nil {
		t.Fatalf("messageViewToTLMessage error = %v", err)
	}
	tlMessage, ok := message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", message)
	}
	if tlMessage.Id != 102 || tlMessage.Message != "caption" || !tlMessage.Noforwards {
		t.Fatalf("unexpected V3 message: %+v", tlMessage)
	}
	if _, ok := tlMessage.Media.(*tg.TLMessageMediaDocument); !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", tlMessage.Media)
	}
	if tlMessage.GroupedId == nil || *tlMessage.GroupedId != 444 || tlMessage.FwdFrom == nil {
		t.Fatalf("grouped/forward = grouped:%v fwd:%T", tlMessage.GroupedId, tlMessage.FwdFrom)
	}
}

func TestMessageViewToTLMessageSupportsV4SharedProjection(t *testing.T) {
	eventPayload := mustMarshalMessageEventV4(t, payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:        payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID:   2004,
			PeerType:             payload.PeerTypeChat,
			PeerID:               1002,
			PeerSeq:              10,
			SenderUserID:         1001,
			Date:                 1_772_000_000,
			ReplyToUserMessageID: 77,
			ServiceAction: mustServiceActionRef(t, tg.MakeTLMessageActionChatCreate(&tg.TLMessageActionChatCreate{
				Title: "v4 chat",
				Users: []int64{1001, 1002},
			})),
		},
		MessageID: 104,
	})
	message, err := messageViewToTLMessage(repository.MessageView{
		UserID:             1001,
		PeerType:           payload.PeerTypeChat,
		PeerID:             1002,
		PeerSeq:            10,
		UserMessageID:      104,
		CanonicalMessageID: 2004,
		MessageStatus:      repository.MessageStatusLive,
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV4,
		ViewPayload:        eventPayload,
	})
	if err != nil {
		t.Fatalf("messageViewToTLMessage error = %v", err)
	}
	service, ok := message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessageService", message)
	}
	if service.Id != 104 || !service.Out {
		t.Fatalf("service message = %+v, want id=104 outgoing", service)
	}
	if _, ok := service.Action.(*tg.TLMessageActionChatCreate); !ok {
		t.Fatalf("service action = %T, want *tg.TLMessageActionChatCreate", service.Action)
	}
	reply, ok := service.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok || reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 77 {
		t.Fatalf("reply header = %+v ok=%v, want reply_to_msg_id=77", reply, ok)
	}
}

func TestMessageViewToTLMessageSupportsV3ContactMedia(t *testing.T) {
	eventPayload := mustMarshalMessageEventV3(t, payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 2003,
		PeerSeq:            9,
		MessageID:          103,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		FromUserID:         1001,
		ToUserID:           1002,
		Date:               1_772_000_000,
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV1,
			Kind:          "contact",
			PhoneNumber:   "8613000000001",
			FirstName:     "13000000001",
			LastName:      "t2",
			UserID:        1571266964,
		},
	})
	message, err := messageViewToTLMessage(repository.MessageView{
		UserID:             1001,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1002,
		PeerSeq:            9,
		UserMessageID:      103,
		CanonicalMessageID: 2003,
		MessageStatus:      repository.MessageStatusLive,
		ViewSchemaVersion:  payload.MessageEventSchemaVersionV3,
		ViewPayload:        eventPayload,
	})
	if err != nil {
		t.Fatalf("messageViewToTLMessage error = %v", err)
	}
	tlMessage, ok := message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", message)
	}
	media, ok := tlMessage.Media.(*tg.TLMessageMediaContact)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaContact", tlMessage.Media)
	}
	if media.PhoneNumber != "8613000000001" || media.FirstName != "13000000001" || media.LastName != "t2" || media.UserId != 1571266964 {
		t.Fatalf("contact media = %+v, want shared contact fields", media)
	}
}

func TestGetStateReturnsRepositoryState(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		state: &repository.UserState{UserID: 1001, Pts: 55, UnreadCount: 3},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetState(&userupdates.TLUserupdatesGetState{UserId: 1001, AuthKeyId: 9001})
	if err != nil {
		t.Fatalf("GetState returned error: %v", err)
	}
	if got.Pts != 55 || got.UnreadCount != 3 {
		t.Fatalf("unexpected state: %+v", got)
	}
}

func TestGetDifferenceStateUnreadCountIsZero(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 55, Date: 1_772_000_030, UnreadCount: 3},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{UserId: 1001, AuthKeyId: 9001})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	diff, ok := got.ToUserDifferenceEmpty()
	if !ok {
		t.Fatalf("expected empty difference, got %s", got.ClazzName())
	}
	if diff.State.Date != 1_772_000_030 || diff.State.UnreadCount != 0 {
		t.Fatalf("difference state = %+v, want date carried and unread_count 0", diff.State)
	}
}

func TestGetStatePassesPermAuthKeyToRepository(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		state: &repository.UserState{UserID: 1001, Pts: 55},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetState(&userupdates.TLUserupdatesGetState{UserId: 1001, AuthKeyId: 9001})
	if err != nil {
		t.Fatalf("GetState returned error: %v", err)
	}
	if repo.stateUserID != 1001 || repo.statePermAuthKeyID != 9001 {
		t.Fatalf("unexpected repository state cursor input: user_id=%d perm_auth_key_id=%d", repo.stateUserID, repo.statePermAuthKeyID)
	}
}

func TestGetDifferenceCarriesNilDateAsPtsOnlyRequest(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:        1001,
		AuthKeyId:     9001,
		Pts:           17,
		PtsTotalLimit: int32Ptr(10),
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	if repo.differenceInput.UserID != 1001 || repo.differenceInput.PermAuthKeyID != 9001 || repo.differenceInput.Pts != 17 || repo.differenceInput.Limit != 10 {
		t.Fatalf("unexpected repository difference input: %+v", repo.differenceInput)
	}
	if repo.differenceInput.Date != nil {
		t.Fatalf("expected nil date, got %v", *repo.differenceInput.Date)
	}
}

func TestGetDifferenceCarriesDateToRepository(t *testing.T) {
	date := int64(1_772_000_000)
	repo := &fakeUserUpdatesRepository{
		difference: &repository.GetDifferenceResult{
			State: repository.UserState{UserID: 1001, Pts: 18},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetDifference(&userupdates.TLUserupdatesGetDifference{
		UserId:    1001,
		AuthKeyId: 9001,
		Pts:       17,
		Date:      &date,
	})
	if err != nil {
		t.Fatalf("GetDifference returned error: %v", err)
	}
	if repo.differenceInput.Date == nil || *repo.differenceInput.Date != date {
		t.Fatalf("expected date %d, got %v", date, repo.differenceInput.Date)
	}
}

func TestDifferenceMapsDialogAuthSeqEvents(t *testing.T) {
	pinned := false
	folderID := int32(2)
	ttl := int32(86400)
	tests := []struct {
		name      string
		eventKind string
		event     payload.DialogEventV1
		wantType  any
	}{
		{
			name:      "draft saved",
			eventKind: payload.DialogEventDraftSaved,
			wantType:  &tg.TLUpdateDraftMessage{},
		},
		{
			name:      "dialog pin",
			eventKind: payload.DialogEventPinToggled,
			event:     payload.DialogEventV1{Pinned: &pinned, FolderID: &folderID},
			wantType:  &tg.TLUpdateDialogPinned{},
		},
		{
			name:      "pinned order",
			eventKind: payload.DialogEventPinnedDialogsReordered,
			event:     payload.DialogEventV1{FolderID: &folderID},
			wantType:  &tg.TLUpdatePinnedDialogs{},
		},
		{
			name:      "filter updated",
			eventKind: payload.DialogEventFilterUpdated,
			wantType:  &tg.TLUpdateDialogFilter{},
		},
		{
			name:      "filter order",
			eventKind: payload.DialogEventFiltersOrderUpdated,
			wantType:  &tg.TLUpdateDialogFilterOrder{},
		},
		{
			name:      "wallpaper",
			eventKind: payload.DialogEventWallpaperChanged,
			wantType:  &tg.TLUpdatePeerWallpaper{},
		},
		{
			name:      "private ttl",
			eventKind: payload.DialogEventPrivatePeerHistoryTTL,
			event:     payload.DialogEventV1{TTLPeriod: &ttl},
			wantType:  &tg.TLUpdatePeerHistoryTTL{},
		},
		{
			name:      "saved dialog pinned",
			eventKind: payload.DialogEventSavedDialogPinned,
			event:     payload.DialogEventV1{Pinned: &pinned},
			wantType:  &tg.TLUpdateSavedDialogPinned{},
		},
		{
			name:      "pinned saved order",
			eventKind: payload.DialogEventPinnedSavedDialogsChanged,
			wantType:  &tg.TLUpdatePinnedSavedDialogs{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialogEvent := tt.event
			dialogEvent.SchemaVersion = payload.DialogEventSchemaVersion
			dialogEvent.EventKind = tt.eventKind
			dialogEvent.PublicUpdateType = tt.eventKind
			dialogEvent.PeerType = payload.PeerTypeUser
			dialogEvent.PeerID = 1002
			update, err := dialogEventToTLUpdate(dialogEvent, 0, 0)
			if err != nil {
				t.Fatalf("dialogEventToTLUpdate() error = %v", err)
			}
			eventPayload, err := iface.EncodeObject(update, repository.AuthSeqLayer)
			if err != nil {
				t.Fatalf("EncodeObject() error = %v", err)
			}
			got, err := differenceToTL(&repository.GetDifferenceResult{
				State: repository.UserState{UserID: 1001, Pts: 18, Seq: 7, Date: 1_772_000_001},
				AuthSeqEvents: []repository.AuthSeqEvent{
					{
						UserID:             1001,
						Seq:                7,
						Date:               1_772_000_001,
						OperationID:        "v1:dialog:auth",
						PublicUpdateType:   tt.eventKind,
						PeerType:           payload.PeerTypeUser,
						PeerID:             1002,
						EventSchemaVersion: repository.AuthSeqLayer,
						EventCodec:         repository.AuthSeqCodecTLBinary,
						EventPayload:       eventPayload,
						EventPayloadHash:   payload.HashBytes(eventPayload),
					},
				},
			})
			if err != nil {
				t.Fatalf("differenceToTL returned error: %v", err)
			}
			diff, ok := got.ToUserDifference()
			if !ok {
				t.Fatalf("expected userDifference, got %s", got.ClazzName())
			}
			if diff.State == nil || diff.State.Pts != 18 || diff.State.Seq != 7 || diff.State.Date != 1_772_000_001 {
				t.Fatalf("unexpected state: %#v", diff.State)
			}
			if len(diff.OtherUpdates) != 1 {
				t.Fatalf("expected one auth-seq update, got %d", len(diff.OtherUpdates))
			}
			switch tt.wantType.(type) {
			case *tg.TLUpdateDraftMessage:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDraftMessage); !ok {
					t.Fatalf("expected TLUpdateDraftMessage, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogPinned:
				update, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogPinned)
				if !ok {
					t.Fatalf("expected TLUpdateDialogPinned, got %T", diff.OtherUpdates[0])
				}
				if update.Pinned != pinned || update.FolderId == nil || *update.FolderId != folderID {
					t.Fatalf("unexpected pinned update: %+v", update)
				}
			case *tg.TLUpdatePinnedDialogs:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePinnedDialogs); !ok {
					t.Fatalf("expected TLUpdatePinnedDialogs, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogFilter:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogFilter); !ok {
					t.Fatalf("expected TLUpdateDialogFilter, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdateDialogFilterOrder:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateDialogFilterOrder); !ok {
					t.Fatalf("expected TLUpdateDialogFilterOrder, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePeerWallpaper:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePeerWallpaper); !ok {
					t.Fatalf("expected TLUpdatePeerWallpaper, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePeerHistoryTTL:
				update, ok := diff.OtherUpdates[0].(*tg.TLUpdatePeerHistoryTTL)
				if !ok {
					t.Fatalf("expected TLUpdatePeerHistoryTTL, got %T", diff.OtherUpdates[0])
				}
				if update.TtlPeriod == nil || *update.TtlPeriod != ttl {
					t.Fatalf("unexpected ttl update: %+v", update)
				}
			case *tg.TLUpdateSavedDialogPinned:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdateSavedDialogPinned); !ok {
					t.Fatalf("expected TLUpdateSavedDialogPinned, got %T", diff.OtherUpdates[0])
				}
			case *tg.TLUpdatePinnedSavedDialogs:
				if _, ok := diff.OtherUpdates[0].(*tg.TLUpdatePinnedSavedDialogs); !ok {
					t.Fatalf("expected TLUpdatePinnedSavedDialogs, got %T", diff.OtherUpdates[0])
				}
			default:
				t.Fatalf("unsupported test type %T", tt.wantType)
			}
		})
	}
}

func TestDifferenceMapsFolderPeersAsPTSEvent(t *testing.T) {
	folderID := int32(3)
	eventPayload := mustMarshalMessageEvent(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.DialogEventFolderPeersChanged,
	})
	got, err := differenceToTL(&repository.GetDifferenceResult{
		State: repository.UserState{UserID: 1001, Pts: 21},
		Events: []repository.UserEvent{
			{
				UserID:             1001,
				Pts:                21,
				PtsCount:           1,
				OperationID:        "v1:dialog:folder",
				EventType:          repository.EventTypeDialogPublicUpdate,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				EventSchemaVersion: payload.MessageEventSchemaVersion,
				EventCodec:         repository.PayloadCodecJSON,
				EventPayload:       eventPayload,
				EventPayloadHash:   payload.HashBytes(eventPayload),
			},
		},
	})
	if err != nil {
		t.Fatalf("differenceToTL returned error: %v", err)
	}
	diff, ok := got.ToUserDifference()
	if !ok {
		t.Fatalf("expected userDifference, got %s", got.ClazzName())
	}
	if len(diff.NewMessages) != 0 || len(diff.OtherUpdates) != 1 {
		t.Fatalf("unexpected difference lens: new=%d updates=%d", len(diff.NewMessages), len(diff.OtherUpdates))
	}
	update, ok := diff.OtherUpdates[0].(*tg.TLUpdateFolderPeers)
	if !ok {
		t.Fatalf("expected TLUpdateFolderPeers, got %T", diff.OtherUpdates[0])
	}
	if update.Pts != 21 || update.PtsCount != 1 {
		t.Fatalf("unexpected pts update: %+v", update)
	}
	if len(update.FolderPeers) != 1 {
		t.Fatalf("expected one folder peer, got %d", len(update.FolderPeers))
	}
	fp := update.FolderPeers[0]
	if fp.FolderId != 0 && fp.FolderId != folderID {
		t.Fatalf("unexpected folder peer: %+v", fp)
	}
}

func TestDifferenceRejectsNonUpdateAuthSeqPayload(t *testing.T) {
	eventPayload, encodeErr := iface.EncodeObject(tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}), repository.AuthSeqLayer)
	if encodeErr != nil {
		t.Fatalf("EncodeObject() error = %v", encodeErr)
	}
	_, err := differenceToTL(&repository.GetDifferenceResult{
		State: repository.UserState{UserID: 1001, Pts: 18, Seq: 7, Date: 1_772_000_001},
		AuthSeqEvents: []repository.AuthSeqEvent{
			{
				UserID:             1001,
				Seq:                7,
				Date:               1_772_000_001,
				OperationID:        "v1:dialog:auth",
				PublicUpdateType:   "dialog.unknown",
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				EventSchemaVersion: repository.AuthSeqLayer,
				EventCodec:         repository.AuthSeqCodecTLBinary,
				EventPayload:       eventPayload,
				EventPayloadHash:   payload.HashBytes(eventPayload),
			},
		},
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("expected ErrUserupdatesStorage, got %v", err)
	}
}

func TestGetOutboxReadDateRoutesToRepository(t *testing.T) {
	repo := &fakeUserUpdatesRepository{outboxReadDate: 123456}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesGetOutboxReadDate(&userupdates.TLUserupdatesGetOutboxReadDate{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		MsgId:    9,
	})
	if err != nil {
		t.Fatalf("UserupdatesGetOutboxReadDate error = %v", err)
	}
	if got.Date != 123456 {
		t.Fatalf("date = %d, want 123456", got.Date)
	}
	if repo.outboxReadDateInput.UserID != 1001 || repo.outboxReadDateInput.PeerType != payload.PeerTypeUser ||
		repo.outboxReadDateInput.PeerID != 1002 || repo.outboxReadDateInput.MsgID != 9 {
		t.Fatalf("repository input = %+v", repo.outboxReadDateInput)
	}
}

func TestGetOutboxReadDateRejectsDateOverflow(t *testing.T) {
	repo := &fakeUserUpdatesRepository{outboxReadDate: int64(math.MaxInt32) + 1}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.UserupdatesGetOutboxReadDate(&userupdates.TLUserupdatesGetOutboxReadDate{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		MsgId:    9,
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("expected ErrUserupdatesStorage, got %v", err)
	}
}

type fakePushOutboxNotifier struct {
	wakes int
}

func (f *fakePushOutboxNotifier) Wake() {
	f.wakes++
}

type fakeUserUpdatesRepository struct {
	applyInput                repository.ApplyUserOperationInput
	applyResult               *repository.ApplyUserOperationResult
	applyBatchInputs          []repository.ApplyUserOperationInput
	applyBatchResults         []repository.ApplyUserOperationResult
	applyBatchCalled          bool
	operationResult           *repository.OperationResult
	stateUserID               int64
	statePermAuthKeyID        int64
	state                     *repository.UserState
	differenceInput           repository.GetDifferenceInput
	difference                *repository.GetDifferenceResult
	appendAuthSeqUpdateInput  repository.AuthSeqUpdateAppendInput
	appendAuthSeqUpdateResult *repository.AuthSeqUpdateAppendResult
	appendAuthSeqUpdateErr    error
	dialogListUserID          int64
	dialogListCursor          repository.DialogProjectionCursor
	dialogListLimit           int32
	dialogProjections         []repository.DialogProjection
	dialogPeerUserID          int64
	dialogPeers               []repository.DialogProjectionPeer
	dialogPeerMap             map[repository.DialogProjectionPeer]repository.DialogProjection
	messageViewUserID         int64
	messageViewPeers          []repository.MessageViewPeerSeq
	messageViews              map[repository.MessageViewPeerSeq]repository.MessageView
	dialogCountUserID         int64
	dialogCount               int32
	outboxReadDate            int64
	outboxReadDateInput       repository.OutboxReadDateInput
}

func (f *fakeUserUpdatesRepository) ApplyUserOperation(_ context.Context, in repository.ApplyUserOperationInput) (*repository.ApplyUserOperationResult, error) {
	f.applyInput = in
	return f.applyResult, nil
}

func (f *fakeUserUpdatesRepository) ApplyUserOperationBatch(_ context.Context, inputs []repository.ApplyUserOperationInput) ([]repository.ApplyUserOperationResult, error) {
	f.applyBatchCalled = true
	f.applyBatchInputs = inputs
	return f.applyBatchResults, nil
}

func (f *fakeUserUpdatesRepository) GetOperationResult(_ context.Context, _ int64, _ string) (*repository.OperationResult, error) {
	return f.operationResult, nil
}

func (f *fakeUserUpdatesRepository) GetState(_ context.Context, userID int64, permAuthKeyID int64) (*repository.UserState, error) {
	f.stateUserID = userID
	f.statePermAuthKeyID = permAuthKeyID
	return f.state, nil
}

func (f *fakeUserUpdatesRepository) GetDifference(_ context.Context, in repository.GetDifferenceInput) (*repository.GetDifferenceResult, error) {
	f.differenceInput = in
	return f.difference, nil
}

func (f *fakeUserUpdatesRepository) AppendAuthSeqUpdate(_ context.Context, in repository.AuthSeqUpdateAppendInput) (*repository.AuthSeqUpdateAppendResult, error) {
	f.appendAuthSeqUpdateInput = in
	if f.appendAuthSeqUpdateResult != nil || f.appendAuthSeqUpdateErr != nil {
		return f.appendAuthSeqUpdateResult, f.appendAuthSeqUpdateErr
	}
	return &repository.AuthSeqUpdateAppendResult{UserID: in.UserID, OperationID: in.OperationID}, nil
}

func (f *fakeUserUpdatesRepository) AppendDialogPtsSideEffect(context.Context, repository.DialogSideEffectAppendInput) (*repository.PtsAppendResult, error) {
	return nil, nil
}

func (f *fakeUserUpdatesRepository) ListDialogProjections(_ context.Context, userID int64, cursor repository.DialogProjectionCursor, limit int32) ([]repository.DialogProjection, error) {
	f.dialogListUserID = userID
	f.dialogListCursor = cursor
	f.dialogListLimit = limit
	return f.dialogProjections, nil
}

func (f *fakeUserUpdatesRepository) GetDialogProjectionsByPeers(_ context.Context, userID int64, peers []repository.DialogProjectionPeer) (map[repository.DialogProjectionPeer]repository.DialogProjection, error) {
	f.dialogPeerUserID = userID
	f.dialogPeers = peers
	return f.dialogPeerMap, nil
}

func (f *fakeUserUpdatesRepository) GetMessageViewsByPeerSeqs(_ context.Context, userID int64, peers []repository.MessageViewPeerSeq) (map[repository.MessageViewPeerSeq]repository.MessageView, error) {
	f.messageViewUserID = userID
	f.messageViewPeers = peers
	return f.messageViews, nil
}

func (f *fakeUserUpdatesRepository) CountVisibleDialogs(_ context.Context, userID int64) (int32, error) {
	f.dialogCountUserID = userID
	return f.dialogCount, nil
}

func (f *fakeUserUpdatesRepository) GetOutboxReadDate(_ context.Context, in repository.OutboxReadDateInput) (int64, error) {
	f.outboxReadDateInput = in
	return f.outboxReadDate, nil
}

func mustMarshalMessageEvent(t *testing.T, event payload.MessageEventV1) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event: %v", err)
	}
	return b
}

func mustMarshalOperationResponseV1(t *testing.T, response payload.OperationResponseV1) []byte {
	t.Helper()
	b, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("marshal operation response v1: %v", err)
	}
	return b
}

func mustMarshalMessageEventV2(t *testing.T, event payload.MessageEventV2) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event v2: %v", err)
	}
	return b
}

func mustMarshalMessageEventV3(t *testing.T, event payload.MessageEventV3) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event v3: %v", err)
	}
	return b
}

func mustMarshalMessageEventV4(t *testing.T, event payload.MessageEventV4) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event v4: %v", err)
	}
	return b
}

func mustMarshalMessageEventBatchV1(t *testing.T, event payload.MessageEventBatchV1) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal message event batch v1: %v", err)
	}
	return b
}

func mustMarshalDialogEvent(t *testing.T, event payload.DialogEventV1) []byte {
	t.Helper()
	b, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal dialog event: %v", err)
	}
	return b
}

func mustMarshalJSONMap(t *testing.T, m map[string]any) []byte {
	t.Helper()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal json map: %v", err)
	}
	return b
}

func int32Ptr(v int32) *int32 {
	return &v
}
