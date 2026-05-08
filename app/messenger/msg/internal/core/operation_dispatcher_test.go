package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func TestOperationDispatcherRequesterSyncUsesOldRPCWithoutEffects(t *testing.T) {
	client := &operationDispatcherUserUpdatesClient{result: &userupdates.UserOperationResult{Pts: 11, PtsCount: 1}}
	core := New(context.Background(), &svc.ServiceContext{UserUpdates: client})

	envelope := testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync)
	got, err := core.dispatchRequesterSync(envelope, nil)
	if err != nil {
		t.Fatalf("dispatchRequesterSync() error = %v", err)
	}
	if got == nil || got.Pts != 11 {
		t.Fatalf("dispatchRequesterSync() result = %#v, want pts 11", got)
	}
	if client.processCalls != 1 || client.processWithEffectsCalls != 0 {
		t.Fatalf("calls process=%d with_effects=%d, want old rpc only", client.processCalls, client.processWithEffectsCalls)
	}
	if client.lastProcess.Operation.UserId != 1001 || client.lastProcess.Operation.OperationId != "requester" {
		t.Fatalf("old rpc operation = %#v", client.lastProcess.Operation)
	}
	assertTLUserOperationMatchesEnvelope(t, client.lastProcess.Operation, envelope)
}

func TestOperationDispatcherRequesterSyncWithDurableEffectUsesEffectsRPC(t *testing.T) {
	client := &operationDispatcherUserUpdatesClient{result: &userupdates.UserOperationResult{Pts: 12, PtsCount: 1}}
	core := New(context.Background(), &svc.ServiceContext{UserUpdates: client})

	got, err := core.dispatchRequesterSync(testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync), []OperationEnvelope{
		testOperationEnvelope(1002, "peer", DeliveryPolicyDurableAsync),
	})
	if err != nil {
		t.Fatalf("dispatchRequesterSync() error = %v", err)
	}
	if got == nil || got.Pts != 12 {
		t.Fatalf("dispatchRequesterSync() result = %#v, want pts 12", got)
	}
	if client.processCalls != 0 || client.processWithEffectsCalls != 1 {
		t.Fatalf("calls process=%d with_effects=%d, want effects rpc only", client.processCalls, client.processWithEffectsCalls)
	}
	if len(client.lastProcessWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects len = %d, want 1", len(client.lastProcessWithEffects.AffectedEffects))
	}
	effect := client.lastProcessWithEffects.AffectedEffects[0]
	if effect.RequesterUserId != 1001 || effect.DeliveryPolicy != int32(DeliveryPolicyDurableAsync) || effect.Operation.UserId != 1002 {
		t.Fatalf("affected effect = %#v", effect)
	}
	assertTLUserOperationMatchesEnvelope(t, effect.Operation, testOperationEnvelope(1002, "peer", DeliveryPolicyDurableAsync))
}

func TestOperationDispatcherSkipsSelfAffectedEffect(t *testing.T) {
	client := &operationDispatcherUserUpdatesClient{result: &userupdates.UserOperationResult{Pts: 13, PtsCount: 1}}
	core := New(context.Background(), &svc.ServiceContext{UserUpdates: client})

	_, err := core.dispatchRequesterSync(testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync), []OperationEnvelope{
		testOperationEnvelope(1001, "self", DeliveryPolicyDurableAsync),
	})
	if err != nil {
		t.Fatalf("dispatchRequesterSync() error = %v", err)
	}
	if client.processCalls != 1 || client.processWithEffectsCalls != 0 {
		t.Fatalf("calls process=%d with_effects=%d, want self effect skipped and old rpc", client.processCalls, client.processWithEffectsCalls)
	}
}

func TestOperationDispatcherBrokerDurableAckPublishesReceiverOperation(t *testing.T) {
	publisher := &operationDispatcherReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{ReceiverPublisher: publisher})

	envelope := testOperationEnvelope(1002, "receiver", DeliveryPolicyBrokerDurableAck)
	ack, err := core.dispatchBrokerDurableAck(envelope)
	if err != nil {
		t.Fatalf("dispatchBrokerDurableAck() error = %v", err)
	}
	if publisher.calls != 1 {
		t.Fatalf("publish calls = %d, want 1", publisher.calls)
	}
	if ack.Topic != "memory" || ack.Partition != publisher.published.PartitionID {
		t.Fatalf("ack = %#v published = %#v", ack, publisher.published)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != "receiver" {
		t.Fatalf("published op = %#v", publisher.published)
	}
	assertReceiverOperationMatchesEnvelope(t, publisher.published, envelope)
}

func TestOperationDispatcherRejectsUnsupportedAffectedPolicy(t *testing.T) {
	client := &operationDispatcherUserUpdatesClient{result: &userupdates.UserOperationResult{Pts: 14, PtsCount: 1}}
	core := New(context.Background(), &svc.ServiceContext{
		UserUpdates:       client,
		ReceiverPublisher: &operationDispatcherReceiverPublisher{},
	})

	_, err := core.dispatchRequesterSync(testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync), []OperationEnvelope{
		testOperationEnvelope(1002, "bad", DeliveryPolicyApplyAck),
	})
	if !errors.Is(err, msg.ErrSenderSyncFailed) {
		t.Fatalf("dispatchRequesterSync() error = %v, want ErrSenderSyncFailed", err)
	}

	_, err = core.dispatchBrokerDurableAck(testOperationEnvelope(1002, "bad-broker", DeliveryPolicyDurableAsync))
	if !errors.Is(err, msg.ErrReceiverBackpressure) {
		t.Fatalf("dispatchBrokerDurableAck() error = %v, want ErrReceiverBackpressure", err)
	}

	publishErr := errors.New("broker down")
	publisher := &operationDispatcherReceiverPublisher{publishErr: publishErr}
	core = New(context.Background(), &svc.ServiceContext{ReceiverPublisher: publisher})
	_, err = core.dispatchBrokerDurableAck(testOperationEnvelope(1002, "receiver", DeliveryPolicyBrokerDurableAck))
	if !errors.Is(err, msg.ErrReceiverBackpressure) {
		t.Fatalf("dispatchBrokerDurableAck() publish error = %v, want ErrReceiverBackpressure", err)
	}
	if !errors.Is(err, publishErr) {
		t.Fatalf("dispatchBrokerDurableAck() publish error = %v, want upstream error", err)
	}
}

func TestOperationDispatcherPreservesUpstreamSenderError(t *testing.T) {
	upstreamErr := errors.New("userupdates unavailable")
	client := &operationDispatcherUserUpdatesClient{err: upstreamErr}
	core := New(context.Background(), &svc.ServiceContext{UserUpdates: client})

	_, err := core.dispatchRequesterSync(testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync), nil)
	if !errors.Is(err, msg.ErrSenderSyncFailed) {
		t.Fatalf("dispatchRequesterSync() error = %v, want ErrSenderSyncFailed", err)
	}
	if !errors.Is(err, upstreamErr) {
		t.Fatalf("dispatchRequesterSync() error = %v, want upstream error", err)
	}
}

func TestOperationDispatcherNilServiceContextReturnsSentinel(t *testing.T) {
	core := New(context.Background(), nil)

	if _, err := core.dispatchRequesterSync(testOperationEnvelope(1001, "requester", DeliveryPolicyRequesterSync), nil); !errors.Is(err, msg.ErrSenderSyncFailed) {
		t.Fatalf("dispatchRequesterSync() error = %v, want ErrSenderSyncFailed", err)
	}
	if _, err := core.dispatchBrokerDurableAck(testOperationEnvelope(1002, "receiver", DeliveryPolicyBrokerDurableAck)); !errors.Is(err, msg.ErrReceiverBackpressure) {
		t.Fatalf("dispatchBrokerDurableAck() error = %v, want ErrReceiverBackpressure", err)
	}
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperationWithEffects(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error) {
	f.processWithEffects = in
	if f.processWithEffectsErr != nil {
		return nil, f.processWithEffectsErr
	}
	if f.processWithEffectsResult != nil {
		return f.processWithEffectsResult, nil
	}
	return f.processResult, nil
}

type operationDispatcherUserUpdatesClient struct {
	processCalls            int
	processWithEffectsCalls int
	lastProcess             *userupdates.TLUserupdatesProcessUserOperation
	lastProcessWithEffects  *userupdates.TLUserupdatesProcessUserOperationWithEffects
	result                  *userupdates.UserOperationResult
	err                     error
}

func (f *operationDispatcherUserUpdatesClient) UserupdatesProcessUserOperation(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	f.processCalls++
	f.lastProcess = in
	if f.err != nil {
		return nil, f.err
	}
	return f.result, nil
}

func (f *operationDispatcherUserUpdatesClient) UserupdatesProcessUserOperationWithEffects(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error) {
	f.processWithEffectsCalls++
	f.lastProcessWithEffects = in
	if f.err != nil {
		return nil, f.err
	}
	return f.result, nil
}

func (f *operationDispatcherUserUpdatesClient) UserupdatesGetOperationResult(context.Context, *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	return f.result, nil
}

type operationDispatcherReceiverPublisher struct {
	calls      int
	published  repository.ReceiverOperation
	publishErr error
}

func (p *operationDispatcherReceiverPublisher) Publish(_ context.Context, op repository.ReceiverOperation) (repository.KafkaAck, error) {
	p.calls++
	p.published = op
	if p.publishErr != nil {
		return repository.KafkaAck{}, p.publishErr
	}
	return repository.KafkaAck{Topic: "memory", Partition: op.PartitionID, Offset: 99}, nil
}

func testOperationEnvelope(userID int64, operationID string, policy DeliveryPolicy) OperationEnvelope {
	body := []byte(`{"schema_version":1}`)
	hash := payload.HashBytes(body)
	authKeyID := int64(7001)
	canonicalMessageID := int64(9001)
	canonicalPeerSeq := int64(3)
	canonicalDate := int64(1778160000)
	return OperationEnvelope{
		UserID:               userID,
		OperationID:          operationID,
		OpType:               99,
		OperationKind:        payload.OperationKindSendMessage,
		ActorUserID:          1001,
		AuthKeyID:            &authKeyID,
		AuthKeyIDExclude:     &authKeyID,
		PeerType:             payload.PeerTypeUser,
		PeerID:               2001,
		CanonicalMessageID:   &canonicalMessageID,
		CanonicalPeerSeq:     &canonicalPeerSeq,
		CanonicalDate:        &canonicalDate,
		PayloadSchemaVersion: payload.MessageOperationSchemaVersion,
		PayloadCodec:         payload.PayloadCodecJSON,
		PayloadHash:          hash,
		Payload:              body,
		DependencyPts:        []int64{41, 42},
		DeliveryPolicy:       policy,
	}
}

func assertTLUserOperationMatchesEnvelope(t *testing.T, got *userupdates.TLUserOperation, want OperationEnvelope) {
	t.Helper()
	if got == nil {
		t.Fatal("TL user operation is nil")
	}
	if got.OpType != want.OpType ||
		got.PayloadCodec != want.PayloadCodec ||
		got.PayloadSchemaVersion != want.PayloadSchemaVersion ||
		got.ActorUserId != want.ActorUserID ||
		got.PeerType != want.PeerType ||
		got.PeerId != want.PeerID {
		t.Fatalf("TL user operation = %#v, want envelope %#v", got, want)
	}
	if got.AuthKeyId == nil || *got.AuthKeyId != *want.AuthKeyID {
		t.Fatalf("auth_key_id = %#v, want %#v", got.AuthKeyId, want.AuthKeyID)
	}
	if got.AuthKeyIdExclude == nil || *got.AuthKeyIdExclude != *want.AuthKeyIDExclude {
		t.Fatalf("auth_key_id_exclude = %#v, want %#v", got.AuthKeyIdExclude, want.AuthKeyIDExclude)
	}
	if got.CanonicalDate == nil || *got.CanonicalDate != *want.CanonicalDate {
		t.Fatalf("canonical_date = %#v, want %#v", got.CanonicalDate, want.CanonicalDate)
	}
	if got.DependencyPts == nil || *got.DependencyPts != want.DependencyPts[0] {
		t.Fatalf("dependency_pts = %#v, want first %d", got.DependencyPts, want.DependencyPts[0])
	}
	if !bytesEqual(got.PayloadHash, want.PayloadHash) || !bytesEqual(got.Payload, want.Payload) {
		t.Fatalf("payload/hash mismatch got_hash=%v want_hash=%v got_payload=%s want_payload=%s", got.PayloadHash, want.PayloadHash, got.Payload, want.Payload)
	}
}

func assertReceiverOperationMatchesEnvelope(t *testing.T, got repository.ReceiverOperation, want OperationEnvelope) {
	t.Helper()
	route := payload.RouteUser(want.UserID)
	if got.OpType != want.OpType ||
		got.PayloadCodec != want.PayloadCodec ||
		got.PeerType != want.PeerType ||
		got.PeerID != want.PeerID ||
		got.BucketID != int32(route.BucketID) ||
		got.PartitionID != int32(route.ReceiverPartitionID) {
		t.Fatalf("receiver op = %#v, want envelope %#v", got, want)
	}
	if !bytesEqual(got.PayloadHash, want.PayloadHash) || !bytesEqual(got.Payload, want.Payload) {
		t.Fatalf("receiver payload/hash mismatch got_hash=%v want_hash=%v got_payload=%s want_payload=%s", got.PayloadHash, want.PayloadHash, got.Payload, want.Payload)
	}
	if len(got.DependencyPts) != len(want.DependencyPts) {
		t.Fatalf("receiver dependency pts = %+v, want %+v", got.DependencyPts, want.DependencyPts)
	}
	for i := range got.DependencyPts {
		if got.DependencyPts[i] != want.DependencyPts[i] {
			t.Fatalf("receiver dependency pts = %+v, want %+v", got.DependencyPts, want.DependencyPts)
		}
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
