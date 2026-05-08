package event

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
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
	if userClient.calls != 0 {
		t.Fatalf("user projection calls = %d, want 0 for updateShortMessage", userClient.calls)
	}
	for i, req := range gatewayClient.requests {
		if req.PermAuthKeyId != []int64{111, 222}[i] {
			t.Fatalf("request %d perm_auth_key_id = %d", i, req.PermAuthKeyId)
		}
		update, ok := req.Updates.(*tg.TLUpdateShortMessage)
		if !ok {
			t.Fatalf("request %d updates = %T, want *tg.TLUpdateShortMessage", i, req.Updates)
		}
		if update.Pts != 38 || update.PtsCount != 1 {
			t.Fatalf("request %d update pts = %#v", i, update)
		}
		if update.Id != 9 || update.UserId != 1001 || update.Message != "hello" || update.Out || update.Date != 1777781234 {
			t.Fatalf("request %d update = %#v", i, update)
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

func stringPtr(v string) *string { return &v }
