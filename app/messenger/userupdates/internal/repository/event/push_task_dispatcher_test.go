package event

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatprojection "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chatprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakePushAuthsession struct {
	userID int64
	keys   []int64
}

func (f *fakePushAuthsession) AuthsessionGetPermAuthKeyIds(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyIds) (*authsession.VectorLong, error) {
	f.userID = in.UserId
	return &authsession.VectorLong{Datas: f.keys}, nil
}

type fakePushGateway struct {
	requests []*gateway.TLGatewayPushUpdatesData
}

func (f *fakePushGateway) GatewayPushUpdatesData(ctx context.Context, in *gateway.TLGatewayPushUpdatesData) (*tg.Bool, error) {
	f.requests = append(f.requests, in)
	return tg.BoolTrue, nil
}

type fakePushUserProjector struct {
	in    *userpb.TLUserGetUserProjectionBundle
	out   *userpb.UserProjectionBundle
	err   error
	calls int
}

func (f *fakePushUserProjector) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	f.calls++
	f.in = in
	return f.out, f.err
}

type fakePushChatProjector struct {
	in    *chatpb.TLChatGetChatProjectionBundle
	out   *chatpb.ChatProjectionBundle
	err   error
	calls int
}

func (f *fakePushChatProjector) ChatGetChatProjectionBundle(_ context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
	f.calls++
	f.in = in
	return f.out, f.err
}

func TestPushTaskDispatcherRoutesMessageUpdateToUserAuthKeys(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1001,
		FromUserID:         1001,
		ToUserID:           2002,
		Date:               1777781234,
		Out:                false,
		MessageText:        "hello",
	})
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        1,
		UserID:        2002,
		Pts:           38,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "op",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{111, 222}}
	gatewayClient := &fakePushGateway{}
	userClient := &fakePushUserProjector{}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, userClient)

	err = dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{
		Topic:     "userupdates.push_tasks.v1",
		Partition: 3,
		Offset:    4,
		Value:     body,
	})
	if err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if auth.userID != 2002 {
		t.Fatalf("auth route user_id = %d, want 2002", auth.userID)
	}
	if len(gatewayClient.requests) != 2 {
		t.Fatalf("gateway push count = %d, want 2", len(gatewayClient.requests))
	}
	if userClient.calls != 1 {
		t.Fatalf("user projection calls = %d, want 1 for full message update", userClient.calls)
	}
	for i, req := range gatewayClient.requests {
		if req.PermAuthKeyId != []int64{111, 222}[i] {
			t.Fatalf("request %d perm_auth_key_id = %d", i, req.PermAuthKeyId)
		}
		updates, ok := req.Updates.(*tg.TLUpdates)
		if !ok {
			t.Fatalf("request %d updates = %T, want *tg.TLUpdates", i, req.Updates)
		}
		update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
		if !ok {
			t.Fatalf("request %d update = %T, want *tg.TLUpdateNewMessage", i, updates.Updates[0])
		}
		message, ok := update.Message.(*tg.TLMessage)
		if !ok {
			t.Fatalf("request %d message = %T, want *tg.TLMessage", i, update.Message)
		}
		if update.Pts != 38 || update.PtsCount != 1 {
			t.Fatalf("request %d update pts = %#v", i, update)
		}
		if message.Id != 9 || message.Message != "hello" || message.Out || message.Date != 1777781234 {
			t.Fatalf("request %d message = %#v", i, message)
		}
		peer, ok := message.PeerId.(*tg.TLPeerUser)
		if !ok || peer.UserId != 1001 {
			t.Fatalf("request %d message peer = %#v, want peerUser(1001)", i, message.PeerId)
		}
	}
}

func TestPushTaskDispatcherSkipsExcludedAuthKey(t *testing.T) {
	excludedAuthKeyID := int64(111)
	eventPayload, err := json.Marshal(payload.MessageEventV1{
		SchemaVersion:    payload.MessageEventSchemaVersion,
		EventKind:        payload.EventKindNewMessage,
		MessageID:        9,
		PeerType:         payload.PeerTypeUser,
		PeerID:           1001,
		FromUserID:       1001,
		ToUserID:         2002,
		Date:             1777781234,
		Out:              false,
		MessageText:      "hello",
		AuthKeyIdExclude: &excludedAuthKeyID,
	})
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        1,
		UserID:        2002,
		Pts:           38,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "op",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{111, 222}}
	gatewayClient := &fakePushGateway{}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, nil)

	err = dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body})
	if err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	if gatewayClient.requests[0].PermAuthKeyId != 222 {
		t.Fatalf("pushed perm_auth_key_id = %d, want 222", gatewayClient.requests[0].PermAuthKeyId)
	}
}

func TestPushTaskDispatcherRoutesWrappedNewMessageWithZeroSeq(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 7001,
		MessageID:          9,
		PeerType:           payload.PeerTypeChat,
		PeerID:             1001,
		FromUserID:         2002,
		ToUserID:           2002,
		Date:               1777781234,
		Out:                false,
		MessageText:        "hello",
	})
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        3,
		UserID:        2002,
		Pts:           40,
		PushType:      1,
		PeerType:      payload.PeerTypeChat,
		PeerID:        1001,
		OperationID:   "chat-message",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{444}}
	gatewayClient := &fakePushGateway{}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, nil)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if updates.Seq != 0 {
		t.Fatalf("updates seq = %d, want 0", updates.Seq)
	}
	if len(updates.Updates) != 1 {
		t.Fatalf("updates payload = %+v", updates)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateNewMessage", updates.Updates[0])
	}
	if update.Pts != 40 || update.PtsCount != 1 {
		t.Fatalf("update pts = %+v", update)
	}
}

func TestV4CreateChatPushOmitsUpdateMessageID(t *testing.T) {
	chatFact, err := payload.WrapFact(payload.FactKindChatParticipantsChanged, payload.ChatParticipantsChangedFactV1{
		SchemaVersion: payload.MessageOperationSchemaVersionV4,
		ChatID:        3001,
		ActorUserID:   1001,
		Version:       1,
		Participants: []payload.ChatParticipantFactV1{
			{UserID: 1001, Role: "creator", Date: 1777781234},
			{UserID: 2002, Role: "member", InviterUserID: 1001, Date: 1777781234},
		},
	})
	if err != nil {
		t.Fatalf("WrapFact(chat participants) error = %v", err)
	}
	eventPayload, err := json.Marshal(payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:      payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID: 7002,
			PeerType:           payload.PeerTypeChat,
			PeerID:             3001,
			PeerSeq:            1,
			SenderUserID:       1001,
			ToUserID:           2002,
			Date:               1777781234,
			ServiceAction: &payload.ServiceActionRefV1{
				SchemaVersion: payload.ServiceActionSchemaVersionV1,
				Kind:          payload.ServiceActionKindChatCreate,
				Title:         "v4 chat",
				Users:         []int64{1001, 2002},
			},
		},
		AttachFacts: []payload.UpdateFactV1{chatFact},
		MessageID:   10,
		Pts:         41,
		PtsCount:    1,
	})
	if err != nil {
		t.Fatalf("marshal V4 event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        4,
		UserID:        1001,
		Pts:           41,
		PushType:      1,
		PeerType:      payload.PeerTypeChat,
		PeerID:        3001,
		OperationID:   "v4-chat-create",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{555}}
	gatewayClient := &fakePushGateway{}
	userClient := &fakePushUserProjector{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 2002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	chatClient := &fakePushChatProjector{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001, Chats: []tg.ChatClazz{
				tg.MakeTLChat(&tg.TLChat{
					Id:                3001,
					Title:             "v4 chat",
					ParticipantsCount: 2,
					Date:              1777781234,
					Version:           1,
				}),
			}}),
		},
	}).ToChatProjectionBundle()}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, userClient, chatClient)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Updates) != 2 {
		t.Fatalf("updates len = %d, want 2", len(updates.Updates))
	}
	for i, update := range updates.Updates {
		if _, ok := update.(*tg.TLUpdateMessageID); ok {
			t.Fatalf("updates[%d] = *tg.TLUpdateMessageID, want push to omit reply-only update", i)
		}
	}
	if _, ok := updates.Updates[0].(*tg.TLUpdateChatParticipants); !ok {
		t.Fatalf("first update = %T, want *tg.TLUpdateChatParticipants", updates.Updates[0])
	}
	newMessage, ok := updates.Updates[1].(*tg.TLUpdateNewMessage)
	if !ok {
		t.Fatalf("second update = %T, want *tg.TLUpdateNewMessage", updates.Updates[1])
	}
	service, ok := newMessage.Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("new message = %T, want *tg.TLMessageService", newMessage.Message)
	}
	if service.Id != 10 || newMessage.Pts != 41 || newMessage.PtsCount != 1 {
		t.Fatalf("service/update = %+v/%+v", service, newMessage)
	}
	if _, ok := service.Action.(*tg.TLMessageActionChatCreate); !ok {
		t.Fatalf("service action = %T, want *tg.TLMessageActionChatCreate", service.Action)
	}
	if userClient.calls != 1 || userClient.in == nil || len(userClient.in.TargetUserIds) != 2 {
		t.Fatalf("user projection request = calls:%d in:%#v", userClient.calls, userClient.in)
	}
	if chatClient.calls != 1 || chatClient.in == nil || len(chatClient.in.ViewerUserIds) != 1 || chatClient.in.ViewerUserIds[0] != 1001 {
		t.Fatalf("chat projection viewers = calls:%d in:%#v, want [1001]", chatClient.calls, chatClient.in)
	}
	if len(chatClient.in.TargetChatIds) != 1 || chatClient.in.TargetChatIds[0] != 3001 {
		t.Fatalf("chat projection targets = %#v, want [3001]", chatClient.in.TargetChatIds)
	}
	if !hasPushUserID(updates.Users, 1001) || !hasPushUserID(updates.Users, 2002) {
		t.Fatalf("push users = %#v, want creator and invitee", updates.Users)
	}
	if !hasPushChatID(updates.Chats, 3001) {
		t.Fatalf("push chats = %#v, want chat 3001", updates.Chats)
	}
}

func TestPushTaskDispatcherProjectChatsUsesProjectionBundle(t *testing.T) {
	chatClient := &fakePushChatProjector{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001, Chats: []tg.ChatClazz{
				tg.MakeTLChat(&tg.TLChat{Id: 3001, Title: "v4 chat"}),
			}}),
		},
	}).ToChatProjectionBundle()}
	dispatcher := NewPushTaskDispatcher(nil, nil, nil, chatClient)

	chats, err := dispatcher.ProjectChats(context.Background(), 1001, []int64{3001})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}
	if len(chats) != 1 || chats[0].(*tg.TLChat).Id != 3001 {
		t.Fatalf("ProjectChats() = %#v, want chat 3001", chats)
	}
	if chatClient.calls != 1 || chatClient.in == nil || len(chatClient.in.ViewerUserIds) != 1 || chatClient.in.ViewerUserIds[0] != 1001 {
		t.Fatalf("chat projection viewers = calls:%d in:%#v, want [1001]", chatClient.calls, chatClient.in)
	}
	if len(chatClient.in.TargetChatIds) != 1 || chatClient.in.TargetChatIds[0] != 3001 {
		t.Fatalf("chat projection targets = %#v, want [3001]", chatClient.in.TargetChatIds)
	}
}

func TestPushTaskDispatcherProjectChatsWrapsProjectionError(t *testing.T) {
	dispatcher := NewPushTaskDispatcher(nil, nil, nil, &fakePushChatProjector{err: chatprojection.ErrNilBundle})

	_, err := dispatcher.ProjectChats(context.Background(), 1001, []int64{3001})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
	if !errors.Is(err, chatprojection.ErrNilBundle) {
		t.Fatalf("ProjectChats() error = %v, want ErrNilBundle preserved", err)
	}
}

func TestPushTaskDispatcherProjectChatsAllowsEmptyStoredReferenceProjection(t *testing.T) {
	dispatcher := NewPushTaskDispatcher(nil, nil, nil, &fakePushChatProjector{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001}),
		},
		MissingChatIds: []int64{3001},
	}).ToChatProjectionBundle()})

	chats, err := dispatcher.ProjectChats(context.Background(), 1001, []int64{3001})
	if err != nil {
		t.Fatalf("ProjectChats() error = %v", err)
	}
	if len(chats) != 0 {
		t.Fatalf("ProjectChats() = %#v, want empty projection", chats)
	}
}

func TestPushTaskDispatcherProjectChatsWrapsMissingViewer(t *testing.T) {
	dispatcher := NewPushTaskDispatcher(nil, nil, nil, &fakePushChatProjector{out: chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
		ViewerChats: []chatpb.ViewerChatsClazz{
			chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 2002}),
		},
	}).ToChatProjectionBundle()})

	_, err := dispatcher.ProjectChats(context.Background(), 1001, []int64{3001})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) {
		t.Fatalf("ProjectChats() error = %v, want ErrUserupdatesStorage", err)
	}
	if !errors.Is(err, chatprojection.ErrViewerProjectionMissing) {
		t.Fatalf("ProjectChats() error = %v, want ErrViewerProjectionMissing preserved", err)
	}
}

func TestBatchMessagePushRoutesOneUpdatesEnvelope(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventBatchV1{
		SchemaVersion: payload.MessageEventSchemaVersionBatchV1,
		EventKind:     payload.EventKindNewMessage,
		Messages: []payload.MessageEventBatchItemV1{
			{
				MessageFact: payload.NewMessageFactV1{
					SchemaVersion:      1,
					CanonicalMessageID: 8001,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1001,
					SenderUserID:       1001,
					ToUserID:           2002,
					Date:               1777781234,
					MessageText:        "first",
				},
				MessageID: 10,
				Pts:       41,
				PtsCount:  1,
			},
			{
				MessageFact: payload.NewMessageFactV1{
					SchemaVersion:      1,
					CanonicalMessageID: 8002,
					PeerType:           payload.PeerTypeUser,
					PeerID:             1001,
					SenderUserID:       1001,
					ToUserID:           2002,
					Date:               1777781235,
					MessageText:        "second",
				},
				MessageID: 11,
				Pts:       42,
				PtsCount:  1,
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal batch event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        5,
		UserID:        2002,
		Pts:           42,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "v1:msgbatch:receiver:2002:in:8001:8002",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{555}}
	gatewayClient := &fakePushGateway{}
	userClient := &fakePushUserProjector{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 2002, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 2002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, userClient)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Updates) != 2 {
		t.Fatalf("updates len = %d, want 2", len(updates.Updates))
	}
	for i, update := range updates.Updates {
		newMessage, ok := update.(*tg.TLUpdateNewMessage)
		if !ok {
			t.Fatalf("updates[%d] = %T, want *tg.TLUpdateNewMessage", i, update)
		}
		if newMessage.Pts != int32(41+i) || newMessage.PtsCount != 1 {
			t.Fatalf("updates[%d] pts/count = %d/%d, want %d/1", i, newMessage.Pts, newMessage.PtsCount, 41+i)
		}
	}
}

func TestPushTaskDispatcherV4RejectsChannelDependencies(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventV4{
		SchemaVersion: payload.MessageEventSchemaVersionV4,
		EventKind:     payload.EventKindNewMessage,
		MessageFact: payload.NewMessageFactV1{
			SchemaVersion:      payload.MessageOperationSchemaVersionV4,
			CanonicalMessageID: 7004,
			PeerType:           payload.PeerTypeChannel,
			PeerID:             4001,
			PeerSeq:            1,
			SenderUserID:       1001,
			ToUserID:           2002,
			Date:               1777781234,
			MessageText:        "channel dependency",
		},
		MessageID: 11,
		Pts:       42,
		PtsCount:  1,
	})
	if err != nil {
		t.Fatalf("marshal V4 event payload: %v", err)
	}
	dispatcher := NewPushTaskDispatcher(nil, nil, &fakePushUserProjector{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 2002, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
			}}),
		},
	}).ToUserProjectionBundle()})

	_, _, err = dispatcher.pushTaskUpdatesV4(context.Background(), &payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        5,
		UserID:        2002,
		Pts:           42,
		PushType:      1,
		PeerType:      payload.PeerTypeChannel,
		PeerID:        4001,
		OperationID:   "v4-channel",
		Payload:       eventPayload,
	})
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) || !strings.Contains(err.Error(), "channel") {
		t.Fatalf("pushTaskUpdatesV4() error = %v, want ErrUserupdatesStorage channel rejection", err)
	}
}

func TestPushTaskDispatcherRoutesReadHistoryOutboxUpdate(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.OperationKindReadHistory,
		MessageID:     42,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		Date:          1777781234,
		Out:           true,
	})
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        2,
		UserID:        2002,
		Pts:           39,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "read-outbox",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{333}}
	gatewayClient := &fakePushGateway{}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, nil)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Updates) != 1 {
		t.Fatalf("updates payload = %+v", updates)
	}
	if updates.Seq != 0 {
		t.Fatalf("updates seq = %d, want 0", updates.Seq)
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateReadHistoryOutbox)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateReadHistoryOutbox", updates.Updates[0])
	}
	peer, ok := update.Peer.(*tg.TLPeerUser)
	if !ok || peer.UserId != 1001 {
		t.Fatalf("peer = %+v ok=%v", update.Peer, ok)
	}
	if update.MaxId != 42 || update.Pts != 39 || update.PtsCount != 1 {
		t.Fatalf("read outbox update = %+v", update)
	}
}

func TestPushTaskDispatcherProjectsEditMessageUsers(t *testing.T) {
	eventPayload, err := json.Marshal(payload.MessageEventV1{
		SchemaVersion:      payload.MessageEventSchemaVersion,
		EventKind:          payload.OperationKindEditMessage,
		CanonicalMessageID: 7001,
		MessageID:          9,
		PeerType:           payload.PeerTypeUser,
		PeerID:             1001,
		FromUserID:         1001,
		ToUserID:           2002,
		Date:               1777781234,
		EditDate:           1777781334,
		Out:                false,
		MessageText:        "edited",
	})
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	body, err := payload.MarshalPushTaskKafkaMessage(payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        3,
		UserID:        2002,
		Pts:           40,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "edit",
		Payload:       eventPayload,
	})
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	auth := &fakePushAuthsession{keys: []int64{444}}
	gatewayClient := &fakePushGateway{}
	userClient := &fakePushUserProjector{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 2002, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001, FirstName: stringPtr("Sender")}),
				tg.MakeTLUser(&tg.TLUser{Id: 2002, FirstName: stringPtr("Receiver")}),
			}}),
		},
	}).ToUserProjectionBundle()}
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient, userClient)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Updates) != 1 {
		t.Fatalf("updates payload = %+v", updates)
	}
	if updates.Date != 1777781333 || updates.Seq != 0 || len(updates.Users) != 2 {
		t.Fatalf("updates envelope = %+v", updates)
	}
	if userClient.in == nil || len(userClient.in.ViewerUserIds) != 1 || userClient.in.ViewerUserIds[0] != 2002 {
		t.Fatalf("projection viewer request = %#v", userClient.in)
	}
	if len(userClient.in.TargetUserIds) != 2 || userClient.in.TargetUserIds[0] != 1001 || userClient.in.TargetUserIds[1] != 2002 {
		t.Fatalf("projection target request = %#v", userClient.in)
	}
	fromUser, ok := updates.Users[0].(*tg.TLUser)
	if !ok || fromUser.Id != 1001 || fromUser.FirstName == nil || *fromUser.FirstName != "Sender" {
		t.Fatalf("sender user = %+v", updates.Users[0])
	}
	toUser, ok := updates.Users[1].(*tg.TLUser)
	if !ok || toUser.Id != 2002 || toUser.FirstName == nil || *toUser.FirstName != "Receiver" {
		t.Fatalf("receiver user = %+v", updates.Users[1])
	}
	update, ok := updates.Updates[0].(*tg.TLUpdateEditMessage)
	if !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateEditMessage", updates.Updates[0])
	}
	if update.Pts != 40 || update.PtsCount != 1 {
		t.Fatalf("edit update = %+v", update)
	}
	message, ok := update.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", update.Message)
	}
	if message.Id != 9 || message.Message != "edited" || message.EditDate == nil || *message.EditDate != 1777781334 {
		t.Fatalf("edit message = %+v", message)
	}
}

func TestPushTaskDispatcherProjectionFailureIsDegraded(t *testing.T) {
	body := mustPushTaskBody(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.OperationKindEditMessage,
		MessageID:     9,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		FromUserID:    1001,
		ToUserID:      2002,
		Date:          1777781234,
		EditDate:      1777781334,
		MessageText:   "edited",
	}, payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        4,
		UserID:        2002,
		Pts:           41,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "edit-failure",
	})
	gatewayClient := &fakePushGateway{}
	dispatcher := NewPushTaskDispatcher(
		&fakePushAuthsession{keys: []int64{555}},
		gatewayClient,
		&fakePushUserProjector{err: userpb.ErrUserStorage},
	)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Users) != 0 {
		t.Fatalf("users = %#v, want degraded empty users", updates.Users)
	}
}

func TestPushTaskDispatcherMissingStoredUserIsDegraded(t *testing.T) {
	body := mustPushTaskBody(t, payload.MessageEventV1{
		SchemaVersion: payload.MessageEventSchemaVersion,
		EventKind:     payload.OperationKindEditMessage,
		MessageID:     9,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		FromUserID:    1001,
		ToUserID:      2002,
		Date:          1777781234,
		EditDate:      1777781334,
		MessageText:   "edited",
	}, payload.PushTaskKafkaMessageV1{
		SchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskID:        5,
		UserID:        2002,
		Pts:           42,
		PushType:      1,
		PeerType:      payload.PeerTypeUser,
		PeerID:        1001,
		OperationID:   "edit-missing",
	})
	gatewayClient := &fakePushGateway{}
	dispatcher := NewPushTaskDispatcher(
		&fakePushAuthsession{keys: []int64{666}},
		gatewayClient,
		&fakePushUserProjector{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
			ViewerUsers: []userpb.ViewerUsersClazz{
				userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 2002, Users: []tg.UserClazz{
					tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				}}),
			},
			MissingUserIds: []int64{2002},
		}).ToUserProjectionBundle()},
	)

	if err := dispatcher.HandlePushTaskKafkaRecord(context.Background(), PushTaskKafkaRecord{Value: body}); err != nil {
		t.Fatalf("HandlePushTaskKafkaRecord() error = %v", err)
	}
	if len(gatewayClient.requests) != 1 {
		t.Fatalf("gateway push count = %d, want 1", len(gatewayClient.requests))
	}
	updates, ok := gatewayClient.requests[0].Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("updates = %T, want *tg.TLUpdates", gatewayClient.requests[0].Updates)
	}
	if len(updates.Users) != 1 || updates.Users[0].(*tg.TLUser).Id != 1001 {
		t.Fatalf("users = %#v", updates.Users)
	}
}

func mustPushTaskBody(t *testing.T, event payload.MessageEventV1, task payload.PushTaskKafkaMessageV1) []byte {
	t.Helper()
	eventPayload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event payload: %v", err)
	}
	task.Payload = eventPayload
	body, err := payload.MarshalPushTaskKafkaMessage(task)
	if err != nil {
		t.Fatalf("marshal push task: %v", err)
	}
	return body
}

func hasPushUserID(users []tg.UserClazz, id int64) bool {
	for _, user := range users {
		switch u := user.(type) {
		case *tg.TLUser:
			if u != nil && u.Id == id {
				return true
			}
		case *tg.TLUserEmpty:
			if u != nil && u.Id == id {
				return true
			}
		}
	}
	return false
}

func hasPushChatID(chats []tg.ChatClazz, id int64) bool {
	for _, chat := range chats {
		switch c := chat.(type) {
		case *tg.TLChat:
			if c != nil && c.Id == id {
				return true
			}
		case *tg.TLChatEmpty:
			if c != nil && c.Id == id {
				return true
			}
		case *tg.TLChatForbidden:
			if c != nil && c.Id == id {
				return true
			}
		}
	}
	return false
}

func stringPtr(v string) *string { return &v }
