package core

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgSendMessageV2SingleUserPublishesReceiverOperation(t *testing.T) {
	responsePayload := []byte(`{"schema_version":1,"pts":11,"pts_count":1}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 2001,
			PeerSeq:            5,
			MessageDate:        1_772_000_000,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(2001, 1001),
			Status:              1,
			Pts:                 11,
			PtsCount:            1,
			CurrentPts:          11,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessageV2(sendMessageRequest(1001, 1002, 9001, "hello"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	short, ok := got.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if short.Id != 5 || short.Pts != 11 || short.PtsCount != 1 || short.Date != 1_772_000_000 {
		t.Fatalf("unexpected short sent message: %+v", short)
	}
	if repo.markCanonicalCalls != 1 || repo.markSenderCalls != 1 || repo.markReceiverAckedCalls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected repo call counts: %+v", repo)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != payload.ReceiverOperationID(2001, 1002) {
		t.Fatalf("unexpected receiver operation: %+v", publisher.published)
	}
	var receiverOp payload.MessageOperationV1
	if err := json.Unmarshal(publisher.published.Payload, &receiverOp); err != nil {
		t.Fatalf("decode receiver payload: %v", err)
	}
	if receiverOp.Out || receiverOp.FromUserID != 1001 || receiverOp.ToUserID != 1002 || receiverOp.PeerID != 1001 {
		t.Fatalf("unexpected receiver payload: %+v", receiverOp)
	}
	if updatesClient.processed == nil || updatesClient.processed.UserId != 1001 {
		t.Fatalf("sender operation was not sent to userupdates: %+v", updatesClient.processed)
	}
	var senderOp payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender payload: %v", err)
	}
	if !senderOp.Out || senderOp.PeerID != 1002 || senderOp.MessageText != "hello" {
		t.Fatalf("unexpected sender payload: %+v", senderOp)
	}
}

func TestMsgSendMessageV2RecoversSenderCommitFromUserUpdatesResult(t *testing.T) {
	responsePayload := []byte(`{"schema_version":1,"pts":12,"pts_count":1}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 3001,
			PeerSeq:            6,
			MessageDate:        1_772_000_010,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
		markSenderErrs: []error{errors.New("temporary mysql failure"), nil},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(3001, 1001),
			Status:              1,
			Pts:                 12,
			PtsCount:            1,
			CurrentPts:          12,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
		getResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(3001, 1001),
			Status:              1,
			Pts:                 12,
			PtsCount:            1,
			CurrentPts:          12,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessageV2(sendMessageRequest(1001, 1002, 9001, "recover"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if repo.markSenderCalls != 2 {
		t.Fatalf("mark sender calls = %d, want 2", repo.markSenderCalls)
	}
	if updatesClient.getOperationResultCalls != 1 {
		t.Fatalf("get operation result calls = %d, want 1", updatesClient.getOperationResultCalls)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
}

type fakeMsgRepository struct {
	sendState              *repository.SendState
	canonical              *repository.CanonicalMessageResult
	markSenderErrs         []error
	markCanonicalCalls     int
	markSenderCalls        int
	markReceiverAckedCalls int
	markCompletedCalls     int
	markRetryableCalls     int
}

func (f *fakeMsgRepository) CreateOrLoadSendState(context.Context, repository.CreateSendStateInput) (*repository.SendState, error) {
	return f.sendState, nil
}

func (f *fakeMsgRepository) CreateOrGetByClientRandom(context.Context, repository.CreateCanonicalMessageInput) (*repository.CanonicalMessageResult, error) {
	return f.canonical, nil
}

func (f *fakeMsgRepository) MarkCanonicalCreated(context.Context, int64, int64, int64) error {
	f.markCanonicalCalls++
	return nil
}

func (f *fakeMsgRepository) MarkSenderCommitted(_ context.Context, _ repository.MarkSenderCommittedInput) error {
	f.markSenderCalls++
	if len(f.markSenderErrs) == 0 {
		return nil
	}
	err := f.markSenderErrs[0]
	f.markSenderErrs = f.markSenderErrs[1:]
	return err
}

func (f *fakeMsgRepository) MarkReceiverOpsAcked(context.Context, int64, int64) error {
	f.markReceiverAckedCalls++
	return nil
}

func (f *fakeMsgRepository) MarkCompleted(context.Context, int64) error {
	f.markCompletedCalls++
	return nil
}

func (f *fakeMsgRepository) MarkRetryableFailure(context.Context, repository.MarkRetryableFailureInput) error {
	f.markRetryableCalls++
	return nil
}

type fakeUserUpdatesClient struct {
	processed               *userupdates.TLUserOperation
	processResult           *userupdates.UserOperationResult
	getResult               *userupdates.UserOperationResult
	getOperationResultCalls int
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperation(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	f.processed = in.Operation
	return f.processResult, nil
}

func (f *fakeUserUpdatesClient) UserupdatesGetOperationResult(_ context.Context, _ *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	f.getOperationResultCalls++
	return f.getResult, nil
}

type fakeReceiverPublisher struct {
	calls     int
	published repository.ReceiverOperation
}

func (f *fakeReceiverPublisher) Publish(_ context.Context, op repository.ReceiverOperation) error {
	f.calls++
	f.published = op
	return nil
}

func sendMessageRequest(senderID, peerID, authKeyID int64, text string) *msgpb.TLMsgSendMessageV2 {
	return &msgpb.TLMsgSendMessageV2{
		UserId:    senderID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				RandomId: 77,
				Message:  tg.MakeTLMessage(&tg.TLMessage{Message: text}),
			}),
		},
	}
}

func mustHashBytes(t *testing.T, b []byte) []byte {
	t.Helper()
	return payload.HashBytes(b)
}
