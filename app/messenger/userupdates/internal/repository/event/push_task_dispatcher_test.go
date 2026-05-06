package event

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/interface/gateway/gateway"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
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
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient)

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
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient)

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
	dispatcher := NewPushTaskDispatcher(auth, gatewayClient)

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
