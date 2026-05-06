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
	"github.com/teamgram/teamgram-server/v2/pkg/pagination"
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
	if updatesClient.processed.AuthKeyIdExclude == nil || *updatesClient.processed.AuthKeyIdExclude != 9001 {
		t.Fatalf("sender operation auth_key_id_exclude = %v, want 9001", updatesClient.processed.AuthKeyIdExclude)
	}
	var senderOp payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender payload: %v", err)
	}
	if !senderOp.Out || senderOp.PeerID != 1002 || senderOp.FromUserID != 1001 || senderOp.ToUserID != 1002 || senderOp.MessageText != "hello" {
		t.Fatalf("unexpected sender payload: %+v", senderOp)
	}
}

func TestMsgSendMessageV2ClearDraftWritesSenderOperationPayload(t *testing.T) {
	responsePayload := []byte(`{"schema_version":1,"pts":15,"pts_count":1}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 6001,
			PeerSeq:            9,
			MessageDate:        1_772_000_050,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(6001, 1001),
			Status:              1,
			Pts:                 15,
			PtsCount:            1,
			CurrentPts:          15,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: &fakeReceiverPublisher{},
	})

	sourceAuth := int64(9001)
	clearBefore := int32(1_772_000_049)
	req := sendMessageRequest(1001, 1002, 9001, "hello")
	req.ClearDraft = true
	req.SourcePermAuthKeyId = &sourceAuth
	req.ClearDraftBeforeDate = &clearBefore

	if _, err := core.MsgSendMessageV2(req); err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	var senderOp payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender payload: %v", err)
	}
	if !senderOp.ClearDraft || senderOp.SourcePermAuthKeyID != sourceAuth || senderOp.ClearDraftBeforeDate != clearBefore {
		t.Fatalf("unexpected clear draft payload: %+v", senderOp)
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

func TestMsgSendMessageV2RetrySkipsCanonicalMarkWhenAlreadyCanonical(t *testing.T) {
	responsePayload := []byte(`{"schema_version":1,"pts":13,"pts_count":1}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{
			SendStateID:        1,
			Status:             repository.SendStateStatusCanonical,
			CanonicalMessageID: 4001,
			PeerSeq:            7,
		},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 4001,
			PeerSeq:            7,
			MessageDate:        1_772_000_030,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         false,
		},
		markCanonicalErr: errors.New("canonical mark should not be retried"),
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(4001, 1001),
			Status:              1,
			Pts:                 13,
			PtsCount:            1,
			CurrentPts:          13,
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

	got, err := core.MsgSendMessageV2(sendMessageRequest(1001, 1002, 9001, "retry"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if repo.markCanonicalCalls != 0 {
		t.Fatalf("mark canonical calls = %d, want 0", repo.markCanonicalCalls)
	}
	if repo.markSenderCalls != 1 || publisher.calls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected retry call counts: repo=%+v publisher_calls=%d", repo, publisher.calls)
	}
}

func TestMsgSendMessageV2SelfRetryUsesCommittedSenderState(t *testing.T) {
	responsePayload := []byte(`{"schema_version":1,"pts":14,"pts_count":1}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{
			SendStateID:             1,
			Status:                  repository.SendStateStatusSenderCommitted,
			CanonicalMessageID:      5001,
			PeerSeq:                 8,
			SenderOperationID:       payload.SenderOperationID(5001, 1001),
			SenderPTS:               14,
			SenderPTSCount:          1,
			SenderUpdatePayload:     responsePayload,
			SenderUpdatePayloadHash: responseHash,
		},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 5001,
			PeerSeq:            8,
			MessageDate:        1_772_000_040,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         false,
		},
	}
	updatesClient := &fakeUserUpdatesClient{}
	publisher := &fakeReceiverPublisher{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err := core.MsgSendMessageV2(sendMessageRequest(1001, 1001, 9001, "self retry"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	short, ok := got.ToUpdateShortSentMessage()
	if !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if short.Pts != 14 || short.PtsCount != 1 || short.Id != 8 {
		t.Fatalf("unexpected short sent message: %+v", short)
	}
	if updatesClient.processed != nil || repo.markSenderCalls != 0 || publisher.calls != 0 {
		t.Fatalf("unexpected retry side effects: processed=%+v repo=%+v publisher_calls=%d", updatesClient.processed, repo, publisher.calls)
	}
	if repo.markReceiverAckedCalls != 1 || repo.markCompletedCalls != 1 {
		t.Fatalf("unexpected completion calls: repo=%+v", repo)
	}
}

func TestMsgGetHistoryReturnsCanonicalTextMessages(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9001,
				PeerSeq:            2,
				FromUserID:         1001,
				Outgoing:           true,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "second",
				MessageDate:        1_772_000_020,
			},
			{
				CanonicalMessageID: 9000,
				PeerSeq:            1,
				FromUserID:         1001,
				Outgoing:           true,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "first",
				MessageDate:        1_772_000_010,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		OffsetId:  3,
		AddOffset: -2,
		Limit:     20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 2 {
		t.Fatalf("messages len = %d, want 2", len(messages.Messages))
	}
	newest, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if newest.Id != 2 || newest.Message != "second" || newest.Date != 1_772_000_020 || !newest.Out {
		t.Fatalf("unexpected newest message: %+v", newest)
	}
	if repo.historyInput.PeerType != payload.PeerTypeUser ||
		repo.historyInput.UserID != 1001 ||
		repo.historyInput.PeerID != 1002 ||
		repo.historyInput.OffsetID != 3 ||
		repo.historyInput.AddOffset != -2 ||
		repo.historyInput.Limit != 20 {
		t.Fatalf("unexpected history input: %+v", repo.historyInput)
	}
}

func TestMsgGetHistoryUsesViewerScopedOutgoingFlag(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9201,
				PeerSeq:            3,
				FromUserID:         1002,
				Outgoing:           false,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				MessageKind:        repository.MessageKindText,
				MessageText:        "saved from self as incoming projection",
				MessageDate:        1_772_000_030,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1002,
		PeerType: payload.PeerTypeUser,
		PeerId:   1001,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 1 {
		t.Fatalf("messages len = %d, want 1", len(messages.Messages))
	}
	message, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok {
		t.Fatalf("message[0] = %T, want *tg.TLMessage", messages.Messages[0])
	}
	if message.Out {
		t.Fatalf("message.Out = true, want false from viewer-scoped projection: %+v", message)
	}
}

func TestMsgGetHistoryReturnsNotModifiedForMatchingHash(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9101,
				PeerSeq:            2,
				FromUserID:         1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "second",
				MessageDate:        1_772_000_020,
			},
			{
				CanonicalMessageID: 9100,
				PeerSeq:            1,
				FromUserID:         1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				MessageKind:        repository.MessageKindText,
				MessageText:        "first",
				MessageDate:        1_772_000_010,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		Limit:    20,
		Hash:     pagination.HashInt64IDs([]int64{2, 1}),
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	if _, ok := got.ToMessagesMessagesNotModified(); !ok {
		t.Fatalf("MsgGetHistory() = %s, want messages.messagesNotModified", got.ClazzName())
	}
}

func TestMsgGetHistoryPassesViewerUserID(t *testing.T) {
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:    1003,
		AuthKeyId: 9003,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Limit:     30,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	if repo.historyInput.UserID != 1003 {
		t.Fatalf("history input user_id = %d, want viewer user_id 1003", repo.historyInput.UserID)
	}
}

func TestMsgReadHistoryV2ReturnsAffectedMessagesAck(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1002, 2, 9001),
			Status:      1,
			Pts:         15,
			PtsCount:    1,
			CurrentPts:  15,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        &fakeMsgRepository{},
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     2,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil {
		t.Fatal("MsgReadHistoryV2() returned nil")
	}
	if got.Pts != 15 || got.PtsCount != 1 {
		t.Fatalf("affected messages = %+v, want pts=15 pts_count=1", got)
	}
	if len(updatesClient.processedList) != 2 {
		t.Fatalf("processed operations = %d, want reader inbox and peer outbox updates", len(updatesClient.processedList))
	}
	readerOperation := updatesClient.processedList[0]
	if readerOperation.OperationId != readHistoryOperationID(1001, 1002, 2, 9001) {
		t.Fatalf("reader operation_id = %q", readerOperation.OperationId)
	}
	if readerOperation.AuthKeyIdExclude == nil || *readerOperation.AuthKeyIdExclude != 9001 {
		t.Fatalf("auth_key_id_exclude = %v, want 9001", readerOperation.AuthKeyIdExclude)
	}
	var readerOp payload.MessageOperationV1
	if err := json.Unmarshal(readerOperation.Payload, &readerOp); err != nil {
		t.Fatalf("decode read history payload: %v", err)
	}
	if readerOp.OperationKind != payload.OperationKindReadHistory || readerOp.PeerID != 1002 || readerOp.ReadInboxMaxPeerSeq != 2 || readerOp.ReadOutboxMaxPeerSeq != 0 || readerOp.Out {
		t.Fatalf("unexpected reader read history payload: %+v", readerOp)
	}
	peerOperation := updatesClient.processedList[1]
	if peerOperation.UserId != 1002 || peerOperation.PeerId != 1001 {
		t.Fatalf("unexpected peer operation routing: %+v", peerOperation)
	}
	if peerOperation.AuthKeyIdExclude != nil {
		t.Fatalf("peer operation auth_key_id_exclude = %v, want nil", peerOperation.AuthKeyIdExclude)
	}
	var peerOp payload.MessageOperationV1
	if err := json.Unmarshal(peerOperation.Payload, &peerOp); err != nil {
		t.Fatalf("decode peer read outbox payload: %v", err)
	}
	if peerOp.OperationKind != payload.OperationKindReadHistory || peerOp.PeerID != 1001 || peerOp.ReadInboxMaxPeerSeq != 0 || peerOp.ReadOutboxMaxPeerSeq != 2 || !peerOp.Out {
		t.Fatalf("unexpected peer read outbox payload: %+v", peerOp)
	}
}

func TestMsgUpdatePinnedMessageRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: updatePinnedOperationID(1001, 1002, 7, false, 9001),
			Status:      1,
			Pts:         21,
			PtsCount:    1,
			CurrentPts:  21,
		}),
	}
	repo := &fakeMsgRepository{
		canonicalByPeerSeq: &repository.CanonicalMessage{
			CanonicalMessageID: 7001,
			PeerSeq:            7,
			FromUserID:         1002,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			MessageText:        "pin me",
			MessageDate:        1_772_000_123,
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        7,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShort()
	if !ok {
		t.Fatalf("MsgUpdatePinnedMessage() = %s, want updateShort", got.ClazzName())
	}
	pinned, ok := (&tg.Update{Clazz: short.Update}).ToUpdatePinnedMessages()
	if !ok || !pinned.Pinned || len(pinned.Messages) != 1 || pinned.Messages[0] != 7 || pinned.Pts != 21 {
		t.Fatalf("pinned update = %+v ok=%v", pinned, ok)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode update pinned payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindUpdatePinnedMessage || op.PinnedPeerSeq != 7 || op.PinnedCanonicalMessageID != 7001 {
		t.Fatalf("unexpected update pinned payload: %+v", op)
	}
}

func TestMsgDeleteMessagesRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{7, 8}, false, 9001),
			Status:      1,
			Pts:         31,
			PtsCount:    1,
			CurrentPts:  31,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeMsgRepository{}, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{7, 8},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 31 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v", got)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode delete payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindDeleteMessages || len(op.DeletePeerSeqs) != 2 || op.DeletePeerSeqs[0] != 7 {
		t.Fatalf("unexpected delete payload: %+v", op)
	}
}

func TestMsgDeleteHistoryRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteHistoryOperationID(1001, 1002, 9, true, false, 9001),
			Status:      1,
			Pts:         32,
			PtsCount:    1,
			CurrentPts:  32,
		}),
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeMsgRepository{}, UserUpdates: updatesClient})

	got, err := core.MsgDeleteHistory(&msgpb.TLMsgDeleteHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		JustClear: true,
		MaxId:     9,
	})
	if err != nil {
		t.Fatalf("MsgDeleteHistory() error = %v", err)
	}
	if got.Pts != 32 || got.PtsCount != 1 {
		t.Fatalf("affected history = %+v", got)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode delete history payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindDeleteHistory || op.DeleteMaxPeerSeq != 9 || !op.JustClear {
		t.Fatalf("unexpected delete history payload: %+v", op)
	}
}

type fakeMsgRepository struct {
	sendState              *repository.SendState
	canonical              *repository.CanonicalMessageResult
	canonicalByPeerSeq     *repository.CanonicalMessage
	history                []repository.HistoryMessage
	historyInput           repository.ListHistoryMessagesInput
	markCanonicalErr       error
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

func (f *fakeMsgRepository) GetCanonicalMessageByPeerSeq(context.Context, int32, int64, int64) (*repository.CanonicalMessage, error) {
	if f.canonicalByPeerSeq == nil {
		return nil, msgpb.ErrSendStateConflict
	}
	return f.canonicalByPeerSeq, nil
}

func (f *fakeMsgRepository) ListHistoryMessages(_ context.Context, in repository.ListHistoryMessagesInput) ([]repository.HistoryMessage, error) {
	f.historyInput = in
	return f.history, nil
}

func (f *fakeMsgRepository) MarkCanonicalCreated(context.Context, int64, int64, int64) error {
	f.markCanonicalCalls++
	return f.markCanonicalErr
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
	processedList           []*userupdates.TLUserOperation
	processResult           *userupdates.UserOperationResult
	getResult               *userupdates.UserOperationResult
	getOperationResultCalls int
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperation(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	f.processed = in.Operation
	f.processedList = append(f.processedList, in.Operation)
	return f.processResult, nil
}

func (f *fakeUserUpdatesClient) UserupdatesGetOperationResult(_ context.Context, _ *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	f.getOperationResultCalls++
	return f.getResult, nil
}

type fakeReceiverPublisher struct {
	calls      int
	published  repository.ReceiverOperation
	publishErr error
}

func (f *fakeReceiverPublisher) Publish(_ context.Context, op repository.ReceiverOperation) (repository.KafkaAck, error) {
	f.calls++
	f.published = op
	if f.publishErr != nil {
		return repository.KafkaAck{}, f.publishErr
	}
	return repository.KafkaAck{Topic: "memory", Partition: op.PartitionID, Offset: 0}, nil
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
