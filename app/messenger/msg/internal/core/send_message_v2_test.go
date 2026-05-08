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
	responsePayload := []byte(`{"schema_version":2,"pts":11,"pts_count":1,"event_type":"new_message","user_message_id":42}`)
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
	if short.Id != 42 || short.Pts != 11 || short.PtsCount != 1 || short.Date != 1_772_000_000 {
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
	responsePayload := []byte(`{"schema_version":2,"pts":15,"pts_count":1,"event_type":"new_message","user_message_id":45}`)
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

func TestMsgSendMessageV2ReceiverDispatchUsesBrokerDurableAck(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":17,"pts_count":1,"event_type":"new_message","user_message_id":47}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 9001,
			PeerSeq:            11,
			MessageDate:        1_772_000_070,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(9001, 1001),
			Status:              1,
			Pts:                 17,
			PtsCount:            1,
			CurrentPts:          17,
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

	got, err := core.MsgSendMessageV2(sendMessageRequest(1001, 1002, 9001, "broker ack"))
	if err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	if _, ok := got.ToUpdateShortSentMessage(); !ok {
		t.Fatalf("expected updateShortSentMessage, got %s", got.ClazzName())
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != payload.ReceiverOperationID(9001, 1002) {
		t.Fatalf("unexpected receiver operation: %+v", publisher.published)
	}
	if publisher.published.PeerID != 1001 || publisher.published.PayloadCodec != payload.PayloadCodecJSON {
		t.Fatalf("unexpected receiver route payload metadata: %+v", publisher.published)
	}
	if len(updatesClient.processedList) != 1 || updatesClient.processWithEffects != nil {
		t.Fatalf("sender path should use requester sync only, processed=%d with_effects=%+v", len(updatesClient.processedList), updatesClient.processWithEffects)
	}
	if repo.markReceiverAckedCalls != 1 {
		t.Fatalf("mark receiver acked calls = %d, want 1", repo.markReceiverAckedCalls)
	}

	publishErr := errors.New("broker unavailable")
	repo = &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 2, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        2,
			CanonicalMessageID: 9002,
			PeerSeq:            12,
			MessageDate:        1_772_000_071,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
	}
	updatesClient = &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(9002, 1001),
			Status:              1,
			Pts:                 18,
			PtsCount:            1,
			CurrentPts:          18,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher = &fakeReceiverPublisher{publishErr: publishErr}
	core = New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err = core.MsgSendMessageV2(sendMessageRequest(1001, 1002, 9001, "broker fail"))
	if err == nil {
		t.Fatalf("MsgSendMessageV2() error = nil, got=%+v", got)
	}
	if !errors.Is(err, msgpb.ErrReceiverBackpressure) {
		t.Fatalf("MsgSendMessageV2() error = %v, want ErrReceiverBackpressure", err)
	}
	if !errors.Is(err, publishErr) {
		t.Fatalf("MsgSendMessageV2() error = %v, want upstream publish error", err)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if repo.markReceiverAckedCalls != 0 {
		t.Fatalf("mark receiver acked calls = %d, want 0 after publish failure", repo.markReceiverAckedCalls)
	}
}

func TestMarshalSendRequestHashIgnoresClearDraftBeforeDate(t *testing.T) {
	_, firstHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "hello", 0, true, 9001, 1_778_160_035)
	if err != nil {
		t.Fatalf("marshalSendRequest(first) error = %v", err)
	}
	_, retryHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "hello", 0, true, 9001, 1_778_160_066)
	if err != nil {
		t.Fatalf("marshalSendRequest(retry) error = %v", err)
	}
	if string(firstHash) != string(retryHash) {
		t.Fatalf("request hash changed when only clear_draft_before_date changed: first=%x retry=%x", firstHash, retryHash)
	}

	_, changedTextHash, err := marshalSendRequest(1001, payload.PeerTypeUser, 1001, 77, "changed", 0, true, 9001, 1_778_160_066)
	if err != nil {
		t.Fatalf("marshalSendRequest(changed text) error = %v", err)
	}
	if string(firstHash) == string(changedTextHash) {
		t.Fatalf("request hash did not change when message text changed: hash=%x", firstHash)
	}
}

func TestMsgSendMessageV2ReplyToPayloadUsesPublicIDResolution(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":16,"pts_count":1,"event_type":"new_message","user_message_id":43}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		sendState: &repository.SendState{SendStateID: 1, Status: repository.SendStateStatusInitialized},
		canonical: &repository.CanonicalMessageResult{
			SendStateID:        1,
			CanonicalMessageID: 8001,
			PeerSeq:            10,
			MessageDate:        1_772_000_060,
			RequestPayloadHash: payload.HashBytes([]byte("request")),
			CreatedNew:         true,
		},
		resolvedMessageID: &repository.ResolvedMessageID{
			UserID:             1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			UserMessageID:      42,
			PeerSeq:            7,
			CanonicalMessageID: 7001,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         payload.SenderOperationID(8001, 1001),
			Status:              1,
			Pts:                 16,
			PtsCount:            1,
			CurrentPts:          16,
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

	req := sendMessageRequest(1001, 1002, 9001, "reply body")
	replyToMsgID := int32(42)
	req.Message[0].Message.(*tg.TLMessage).ReplyTo = tg.MakeTLMessageReplyHeader(&tg.TLMessageReplyHeader{
		ReplyToMsgId: &replyToMsgID,
	})

	if _, err := core.MsgSendMessageV2(req); err != nil {
		t.Fatalf("MsgSendMessageV2() error = %v", err)
	}
	if repo.resolveInput.UserMessageID != 42 || repo.resolveInput.PeerSeq != 0 {
		t.Fatalf("reply resolver input = %+v, want public user_message_id 42", repo.resolveInput)
	}
	for name, body := range map[string][]byte{
		"sender":   updatesClient.processed.Payload,
		"receiver": publisher.published.Payload,
	} {
		var raw map[string]any
		if err := json.Unmarshal(body, &raw); err != nil {
			t.Fatalf("decode %s payload: %v", name, err)
		}
		if raw["reply_to_canonical_message_id"] != float64(7001) {
			t.Fatalf("%s reply_to_canonical_message_id = %v, want 7001; payload=%s", name, raw["reply_to_canonical_message_id"], string(body))
		}
		if _, ok := raw["reply_to_peer_seq"]; ok {
			t.Fatalf("%s payload leaked view peer_seq before projection: %s", name, string(body))
		}
		if _, ok := raw["reply_to_user_message_id"]; ok {
			t.Fatalf("%s payload must not include reply_to_user_message_id for retry hash compatibility: %s", name, string(body))
		}
	}
}

func TestMsgSendMessageV2RecoversSenderCommitFromUserUpdatesResult(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":12,"pts_count":1,"event_type":"new_message","user_message_id":44}`)
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
	responsePayload := []byte(`{"schema_version":2,"pts":13,"pts_count":1,"event_type":"new_message","user_message_id":46}`)
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
	responsePayload := []byte(`{"schema_version":2,"pts":14,"pts_count":1,"event_type":"new_message","user_message_id":48}`)
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
	if short.Pts != 14 || short.PtsCount != 1 || short.Id != 48 {
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
		historyCursorBounds: repository.HistoryCursorBounds{
			OffsetPeerSeq: 13,
			MaxPeerSeq:    12,
			MinPeerSeq:    11,
		},
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID:   9001,
				PeerSeq:              2,
				UserMessageID:        102,
				ReplyToUserMessageID: 101,
				FromUserID:           1001,
				Outgoing:             true,
				PeerType:             payload.PeerTypeUser,
				PeerID:               1002,
				MessageKind:          repository.MessageKindText,
				MessageText:          "second",
				MessageDate:          1_772_000_020,
			},
			{
				CanonicalMessageID: 9000,
				PeerSeq:            1,
				UserMessageID:      101,
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
		OffsetId:  103,
		AddOffset: -2,
		MaxId:     102,
		MinId:     101,
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
	if newest.Id != 102 || newest.Message != "second" || newest.Date != 1_772_000_020 || !newest.Out {
		t.Fatalf("unexpected newest message: %+v", newest)
	}
	reply, ok := newest.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok {
		t.Fatalf("newest ReplyTo = %T, want *tg.TLMessageReplyHeader", newest.ReplyTo)
	}
	if reply.ReplyToMsgId == nil || *reply.ReplyToMsgId != 101 {
		t.Fatalf("reply_to_msg_id = %v, want 101", reply.ReplyToMsgId)
	}
	if repo.historyInput.PeerType != payload.PeerTypeUser ||
		repo.historyInput.UserID != 1001 ||
		repo.historyInput.PeerID != 1002 ||
		repo.historyInput.AddOffset != -2 ||
		repo.historyInput.Limit != 20 ||
		!repo.historyInput.CursorsResolved ||
		repo.historyInput.ResolvedCursorBounds.OffsetPeerSeq != 13 ||
		repo.historyInput.ResolvedCursorBounds.MaxPeerSeq != 12 ||
		repo.historyInput.ResolvedCursorBounds.MinPeerSeq != 11 {
		t.Fatalf("unexpected history input: %+v", repo.historyInput)
	}
	if repo.resolveHistoryInput.offsetID != 103 || repo.resolveHistoryInput.maxID != 102 || repo.resolveHistoryInput.minID != 101 {
		t.Fatalf("history cursors were not resolved from public ids: %+v", repo.resolveHistoryInput)
	}
}

func TestMsgGetHistoryUsesViewerScopedOutgoingFlag(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9201,
				PeerSeq:            3,
				UserMessageID:      103,
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

func TestMsgGetHistoryMissingPublicCursorReturnsEmpty(t *testing.T) {
	repo := &fakeMsgRepository{
		historyCursorBounds: repository.HistoryCursorBounds{NoMatch: true},
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9202,
				PeerSeq:            4,
				UserMessageID:      104,
				FromUserID:         1002,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				MessageKind:        repository.MessageKindText,
				MessageText:        "must not leak across unresolved cursor",
				MessageDate:        1_772_000_040,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgGetHistory(&msgpb.TLMsgGetHistory{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		MaxId:    999,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgGetHistory() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("expected messages.messages, got %s", got.ClazzName())
	}
	if len(messages.Messages) != 0 {
		t.Fatalf("messages len = %d, want empty for unresolved positive public cursor", len(messages.Messages))
	}
	if !repo.historyInput.ResolvedCursorBounds.NoMatch || repo.resolveHistoryInput.maxID != 999 {
		t.Fatalf("history no-match bounds not propagated: input=%+v resolved=%+v", repo.historyInput, repo.resolveHistoryInput)
	}
}

func TestMsgGetHistoryReturnsNotModifiedForMatchingHash(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{
			{
				CanonicalMessageID: 9101,
				PeerSeq:            2,
				UserMessageID:      102,
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
				UserMessageID:      101,
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
		Hash:     pagination.HashInt64IDs([]int64{102, 101}),
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
		processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1002, 102, 9001),
			Status:      1,
			Pts:         15,
			PtsCount:    1,
			CurrentPts:  15,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 102}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      102,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     102,
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
	if len(updatesClient.processedList) != 0 {
		t.Fatalf("direct processed operations = %d, want 0", len(updatesClient.processedList))
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	readerOperation := updatesClient.processWithEffects.Operation
	if readerOperation == nil {
		t.Fatal("with-effects requester operation is nil")
	}
	if readerOperation.OperationId != readHistoryOperationID(1001, 1002, 102, 9001) {
		t.Fatalf("reader operation_id = %q", readerOperation.OperationId)
	}
	if readerOperation.AuthKeyIdExclude == nil || *readerOperation.AuthKeyIdExclude != 9001 {
		t.Fatalf("auth_key_id_exclude = %v, want 9001", readerOperation.AuthKeyIdExclude)
	}
	var readerOp payload.MessageOperationV1
	if err := json.Unmarshal(readerOperation.Payload, &readerOp); err != nil {
		t.Fatalf("decode read history payload: %v", err)
	}
	if readerOp.OperationKind != payload.OperationKindReadHistory || readerOp.PeerID != 1002 || readerOp.ReadInboxMaxPeerSeq != 2 || readerOp.ReadMaxUserMessageID != 102 || readerOp.ReadOutboxMaxPeerSeq != 0 || readerOp.Out {
		t.Fatalf("unexpected reader read history payload: %+v", readerOp)
	}
	if len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %d, want 1", len(updatesClient.processWithEffects.AffectedEffects))
	}
	affected := updatesClient.processWithEffects.AffectedEffects[0]
	if affected.RequesterUserId != 1001 || affected.DeliveryPolicy != int32(DeliveryPolicyDurableAsync) || affected.OperationKind != payload.OperationKindReadHistory {
		t.Fatalf("unexpected affected metadata: %+v", affected)
	}
	peerOperation := affected.Operation
	if peerOperation == nil {
		t.Fatal("affected peer operation is nil")
	}
	if peerOperation.UserId != 1002 || peerOperation.PeerId != 1001 {
		t.Fatalf("unexpected peer operation routing: %+v", peerOperation)
	}
	if peerOperation.OperationId != readHistoryOutboxOperationID(1002, 1001, 2) {
		t.Fatalf("peer read outbox operation_id = %q, want peer-seq scoped id", peerOperation.OperationId)
	}
	if peerOperation.OperationId == readHistoryOperationID(1002, 1001, 102, 0) {
		t.Fatalf("peer read outbox operation_id reused requester public max_id: %q", peerOperation.OperationId)
	}
	if peerOperation.AuthKeyIdExclude != nil {
		t.Fatalf("peer operation auth_key_id_exclude = %v, want nil", peerOperation.AuthKeyIdExclude)
	}
	var peerOp payload.MessageOperationV1
	if err := json.Unmarshal(peerOperation.Payload, &peerOp); err != nil {
		t.Fatalf("decode peer read outbox payload: %v", err)
	}
	if peerOp.OperationKind != payload.OperationKindReadHistory || peerOp.PeerID != 1001 || peerOp.ReadInboxMaxPeerSeq != 0 || peerOp.ReadOutboxMaxPeerSeq != 2 || peerOp.ReadMaxUserMessageID != 0 || !peerOp.Out {
		t.Fatalf("unexpected peer read outbox payload: %+v", peerOp)
	}
}

func TestMsgReadHistoryV2MissingPublicMaxIDNoopsWithoutPTS(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     102,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil || got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected messages = %+v, want no-op pts=0 pts_count=0", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil || len(updatesClient.processedList) != 0 {
		t.Fatalf("missing public max_id dispatched side effects: processed=%+v with_effects=%+v list=%d", updatesClient.processed, updatesClient.processWithEffects, len(updatesClient.processedList))
	}
	if repo.resolveInput.UserID != 1001 || repo.resolveInput.PeerType != payload.PeerTypeUser || repo.resolveInput.PeerID != 1002 || repo.resolveInput.UserMessageID != 102 {
		t.Fatalf("resolve input = %+v, want requester scoped public id", repo.resolveInput)
	}
}

func TestMsgReadHistoryV2SkipsSelfAffectedOutbox(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: readHistoryOperationID(1001, 1001, 2, 9001),
			Status:      1,
			Pts:         16,
			PtsCount:    1,
			CurrentPts:  16,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1001, userMessageID: 2}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1001,
				UserMessageID:      2,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1001,
		MaxId:     2,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil || got.Pts != 16 || got.PtsCount != 1 {
		t.Fatalf("affected messages = %+v, want pts=16 pts_count=1", got)
	}
	if updatesClient.processWithEffects != nil {
		t.Fatalf("with-effects call = %+v, want nil for self read history", updatesClient.processWithEffects)
	}
	if len(updatesClient.processedList) != 1 {
		t.Fatalf("direct processed operations = %d, want 1", len(updatesClient.processedList))
	}
}

func TestMsgReadHistoryV2ReturnsErrorWhenDurableEffectAcceptFails(t *testing.T) {
	acceptErr := errors.New("affected effect accept failed")
	updatesClient := &fakeUserUpdatesClient{
		processWithEffectsErr: acceptErr,
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 2}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      2,
				PeerSeq:            2,
				CanonicalMessageID: 7002,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     2,
	})
	if err == nil {
		t.Fatalf("MsgReadHistoryV2() error = nil, got = %+v", got)
	}
	if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want ErrSenderSyncFailed", err)
	}
	if !errors.Is(err, acceptErr) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want upstream accept error", err)
	}
	if updatesClient.processWithEffects == nil {
		t.Fatal("UserupdatesProcessUserOperationWithEffects was not called")
	}
	if len(updatesClient.processedList) != 0 {
		t.Fatalf("direct processed operations = %d, want 0", len(updatesClient.processedList))
	}
}

func TestMsgReadHistoryV2NilServiceContextReturnsSenderSyncFailed(t *testing.T) {
	core := New(context.Background(), nil)
	_, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     2,
	})
	if !errors.Is(err, msgpb.ErrSenderSyncFailed) {
		t.Fatalf("MsgReadHistoryV2() error = %v, want ErrSenderSyncFailed", err)
	}
}

func TestMsgUpdatePinnedMessageRoutesProjectionOperation(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: updatePinnedOperationID(1001, 1002, 107, false, 9001),
			Status:      1,
			Pts:         21,
			PtsCount:    1,
			CurrentPts:  21,
		}),
	}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_123,
			},
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
		Id:        107,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShort()
	if !ok {
		t.Fatalf("MsgUpdatePinnedMessage() = %s, want updateShort", got.ClazzName())
	}
	pinned, ok := (&tg.Update{Clazz: short.Update}).ToUpdatePinnedMessages()
	if !ok || !pinned.Pinned || len(pinned.Messages) != 1 || pinned.Messages[0] != 107 || pinned.Pts != 21 {
		t.Fatalf("pinned update = %+v ok=%v", pinned, ok)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode update pinned payload: %v", err)
	}
	if op.OperationKind != payload.OperationKindUpdatePinnedMessage || op.PinnedPeerSeq != 7 || op.PinnedUserMessageID != 107 || op.PinnedCanonicalMessageID != 7001 {
		t.Fatalf("unexpected update pinned payload: %+v", op)
	}
	if updatesClient.processed.CanonicalPeerSeq == nil || *updatesClient.processed.CanonicalPeerSeq != 7 {
		t.Fatalf("canonical peer seq = %v, want resolved peer_seq 7", updatesClient.processed.CanonicalPeerSeq)
	}
}

func TestMsgUpdatePinnedMessageRejectsZeroIDForPin(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        0,
		Unpin:     false,
	})
	if !errors.Is(err, msgpb.ErrSendStateConflict) {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v, want ErrSendStateConflict, got=%+v", err, got)
	}
	if updatesClient.processed != nil || len(updatesClient.processedList) != 0 || updatesClient.processWithEffects != nil {
		t.Fatalf("zero pin id dispatched operation: processed=%+v list=%d with_effects=%+v", updatesClient.processed, len(updatesClient.processedList), updatesClient.processWithEffects)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("zero pin id resolved message unexpectedly: %+v", repo.resolveInput)
	}
}

func TestMsgUpdatePinnedMessageAllowsZeroIDForUnpin(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: updatePinnedOperationID(1001, 1002, 0, true, 9001),
			Status:      1,
			Pts:         22,
			PtsCount:    1,
			CurrentPts:  22,
		}),
	}
	repo := &fakeMsgRepository{}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
	})

	got, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		Unpin:     true,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        0,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	short, ok := got.ToUpdateShort()
	if !ok {
		t.Fatalf("MsgUpdatePinnedMessage() = %s, want updateShort", got.ClazzName())
	}
	pinned, ok := (&tg.Update{Clazz: short.Update}).ToUpdatePinnedMessages()
	if !ok || pinned.Pinned || len(pinned.Messages) != 0 || pinned.Pts != 22 {
		t.Fatalf("unpin update = %+v ok=%v", pinned, ok)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("zero unpin id resolved message unexpectedly: %+v", repo.resolveInput)
	}
}

func TestMsgResolveDialogCursorTopMessageResolvesGlobalPublicID(t *testing.T) {
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			42: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             2002,
				UserMessageID:      42,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_000,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgResolveDialogCursorTopMessage(&msgpb.TLMsgResolveDialogCursorTopMessage{
		UserId:       1001,
		TopMessageId: 42,
	})
	if err != nil {
		t.Fatalf("MsgResolveDialogCursorTopMessage error = %v", err)
	}
	if got.Found != tg.BoolTrueClazz || got.PeerType != payload.PeerTypeUser || got.PeerId != 2002 ||
		got.PeerSeq != 7 || got.MessageDate != 1_772_000_000 {
		t.Fatalf("resolved cursor = %+v, want public id mapped to internal cursor", got)
	}
	if repo.resolveInput.UserID != 1001 || repo.resolveInput.UserMessageID != 42 {
		t.Fatalf("resolver input = %+v, want user/global public id", repo.resolveInput)
	}
}

func TestMsgResolveDialogCursorTopMessageReturnsNotFoundForUnknownPositiveID(t *testing.T) {
	repo := &fakeMsgRepository{resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgResolveDialogCursorTopMessage(&msgpb.TLMsgResolveDialogCursorTopMessage{
		UserId:       1001,
		TopMessageId: 404,
	})
	if err != nil {
		t.Fatalf("MsgResolveDialogCursorTopMessage error = %v", err)
	}
	if got.Found != tg.BoolFalseClazz || got.PeerSeq != 0 || got.PeerId != 0 {
		t.Fatalf("resolved cursor = %+v, want explicit not found", got)
	}
}

func TestMsgDeleteMessagesRoutesGlobalPublicIDsByRealPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107, 108}, false, 9001),
			Status:      1,
			Pts:         31,
			PtsCount:    1,
			CurrentPts:  31,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			108: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 108, PeerSeq: 8, CanonicalMessageID: 7008},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  0,
		PeerId:    0,
		Id:        []int32{107, 108},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 31 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v", got)
	}
	if updatesClient.processed == nil {
		t.Fatalf("delete operation was not dispatched")
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &op); err != nil {
		t.Fatalf("decode delete payload: %v", err)
	}
	if op.PeerID != 1002 || len(op.DeletePeerSeqs) != 2 || op.DeletePeerSeqs[0] != 7 || op.DeletePeerSeqs[1] != 8 {
		t.Fatalf("unexpected delete payload: %+v", op)
	}
	if len(op.DeleteUserMessageIDs) != 2 || op.DeleteUserMessageIDs[0] != 107 || op.DeleteUserMessageIDs[1] != 108 {
		t.Fatalf("unexpected public delete ids: %+v", op)
	}
}

func TestMsgDeleteMessagesRetryResolvesDeletedViewAndReturnsStoredResult(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001),
			Status:      1,
			Pts:         51,
			PtsCount:    1,
			CurrentPts:  51,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{},
		resolveDeleteByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{107},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 51 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v, want stored operation result", got)
	}
	if repo.resolveDeleteCalls != 1 {
		t.Fatalf("delete resolver calls = %d, want 1", repo.resolveDeleteCalls)
	}
	if updatesClient.processed == nil {
		t.Fatal("retry over deleted view should still dispatch requester operation")
	}
	if updatesClient.processed.OperationId != deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001) {
		t.Fatalf("operation id = %s", updatesClient.processed.OperationId)
	}
}

func TestBuildDeleteMessagesPayloadUsesStableDateAndHash(t *testing.T) {
	const deleteDate = int32(1_772_000_123)
	body1, hash1, err := buildDeleteMessagesPayload(1001, 1001, payload.PeerTypeUser, 1002, deleteDate, []int64{7}, []int64{107}, false)
	if err != nil {
		t.Fatalf("build payload 1: %v", err)
	}
	body2, hash2, err := buildDeleteMessagesPayload(1001, 1001, payload.PeerTypeUser, 1002, deleteDate, []int64{7}, []int64{107}, false)
	if err != nil {
		t.Fatalf("build payload 2: %v", err)
	}
	if string(body1) != string(body2) {
		t.Fatalf("delete payload changed across repeated builds:\n%s\n%s", body1, body2)
	}
	if string(hash1) != string(hash2) {
		t.Fatalf("delete payload hash changed: %x != %x", hash1, hash2)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(body1, &op); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	if op.Date != deleteDate {
		t.Fatalf("delete payload date = %d, want stable date %d", op.Date, deleteDate)
	}
}

func TestMsgDeleteMessagesGroupsGlobalIDsByPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResults: []*userupdates.UserOperationResult{
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001), Status: 1, Pts: 31, PtsCount: 1, CurrentPts: 31,
			}),
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1003, []int32{208}, false, 9001), Status: 1, Pts: 32, PtsCount: 1, CurrentPts: 32,
			}),
		},
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			208: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1003, UserMessageID: 208, PeerSeq: 2, CanonicalMessageID: 8002},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Id:        []int32{107, 208},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 32 || got.PtsCount != 2 {
		t.Fatalf("affected = %+v, want final pts 32 and two requester operations", got)
	}
	if len(updatesClient.processedOperations) != 2 {
		t.Fatalf("processed operations = %d, want 2", len(updatesClient.processedOperations))
	}
}

func TestMsgDeleteMessagesRevokeEnqueuesPeerEffect(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, true, 9001),
			Status:      1,
			Pts:         33,
			PtsCount:    1,
			CurrentPts:  33,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	_, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Revoke:    true,
		Id:        []int32{107},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 1 {
		t.Fatalf("affected effects = %+v", updatesClient.processWithEffects)
	}
	effect := updatesClient.processWithEffects.AffectedEffects[0].ToAffectedUserOperation()
	if effect == nil || effect.Operation == nil || effect.Operation.UserId != 1002 || effect.Operation.PeerId != 1001 {
		t.Fatalf("peer effect = %+v", effect)
	}
	var peerOp payload.MessageOperationV1
	if err := json.Unmarshal(effect.Operation.Payload, &peerOp); err != nil {
		t.Fatalf("decode peer payload: %v", err)
	}
	if len(peerOp.DeleteUserMessageIDs) != 0 || len(peerOp.DeletePeerSeqs) != 1 || peerOp.DeletePeerSeqs[0] != 7 {
		t.Fatalf("peer delete payload = %+v", peerOp)
	}
}

func TestMsgDeleteMessagesRevokeSelfDoesNotEnqueuePeerEffect(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1001, []int32{3}, true, 9001), Status: 1, Pts: 34, PtsCount: 1, CurrentPts: 34,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			3: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1001, UserMessageID: 3, PeerSeq: 3, CanonicalMessageID: 7003},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})
	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{UserId: 1001, AuthKeyId: 9001, Revoke: true, Id: []int32{3}})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 34 || got.PtsCount != 1 {
		t.Fatalf("affected = %+v, want requester pts 34 and pts_count 1", got)
	}
	if updatesClient.processed == nil {
		t.Fatal("requester delete operation was not dispatched")
	}
	if updatesClient.processed.OperationId != deleteMessagesOperationID(1001, 1001, []int32{3}, true, 9001) {
		t.Fatalf("requester operation id = %s", updatesClient.processed.OperationId)
	}
	if updatesClient.processWithEffects != nil {
		t.Fatalf("self delete should not use with-effects dispatch: %+v", updatesClient.processWithEffects)
	}
}

func TestMsgDeleteMessagesRejectsPeerEmptyWithNonZeroPeerID(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:      1001,
			OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, false, 9001),
			Status:      1,
			Pts:         31,
			PtsCount:    1,
			CurrentPts:  31,
		}),
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  0,
		PeerId:    1002,
		Id:        []int32{107},
	})
	if err == nil {
		t.Fatalf("MsgDeleteMessages() error = nil, got = %+v", got)
	}
	if !errors.Is(err, msgpb.ErrSendStateConflict) {
		t.Fatalf("MsgDeleteMessages() error = %v, want ErrSendStateConflict", err)
	}
}

func TestMsgDeleteMessagesRevokeGroupsGlobalIDsByPeerUsesSequentialResults(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{
		processResults: []*userupdates.UserOperationResult{
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1002, []int32{107}, true, 9001), Status: 1, Pts: 41, PtsCount: 1, CurrentPts: 41,
			}),
			userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
				UserId: 1001, OperationId: deleteMessagesOperationID(1001, 1003, []int32{208}, true, 9001), Status: 1, Pts: 42, PtsCount: 1, CurrentPts: 42,
			}),
		},
	}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			107: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1002, UserMessageID: 107, PeerSeq: 7, CanonicalMessageID: 7007},
			208: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 1003, UserMessageID: 208, PeerSeq: 2, CanonicalMessageID: 8002},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		Revoke:    true,
		Id:        []int32{107, 208},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 42 || got.PtsCount != 2 {
		t.Fatalf("affected = %+v, want final pts 42 and two requester operations", got)
	}
	if len(updatesClient.processedOperations) != 2 {
		t.Fatalf("processed operations = %d, want 2", len(updatesClient.processedOperations))
	}
	if updatesClient.processedOperations[0].PeerId != 1002 || updatesClient.processedOperations[1].PeerId != 1003 {
		t.Fatalf("processed operation peers = %d,%d; want 1002,1003", updatesClient.processedOperations[0].PeerId, updatesClient.processedOperations[1].PeerId)
	}
}

func TestMsgDeleteMessagesIgnoresMissingPublicIDs(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{404},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected = %+v, want empty no-op", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("missing public ids should not dispatch operations: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
	}
}

func TestMsgDeleteMessagesIgnoresResolvedIDsFromOtherPeer(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{
		resolveManyByUserMessageID: map[int64]*repository.ResolvedMessageID{
			207: {UserID: 1001, PeerType: payload.PeerTypeUser, PeerID: 2002, UserMessageID: 207, PeerSeq: 17, CanonicalMessageID: 7207},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		Id:        []int32{207},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 {
		t.Fatalf("affected = %+v, want no-op for other peer id", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("other peer public id dispatched delete: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
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
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 9}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      9,
				PeerSeq:            99,
				CanonicalMessageID: 7099,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

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
	if op.OperationKind != payload.OperationKindDeleteHistory || op.DeleteMaxPeerSeq != 99 || op.PeerSeq != 99 || !op.JustClear {
		t.Fatalf("unexpected delete history payload: %+v", op)
	}
	if updatesClient.processed.OperationId != deleteHistoryOperationID(1001, 1002, 99, true, false, 9001) {
		t.Fatalf("delete history operation_id = %q, want resolved peer-seq key", updatesClient.processed.OperationId)
	}
}

func TestMsgDeleteHistoryMissingPositiveMaxIDIsNoOp(t *testing.T) {
	updatesClient := &fakeUserUpdatesClient{}
	repo := &fakeMsgRepository{resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{}}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient})

	got, err := core.MsgDeleteHistory(&msgpb.TLMsgDeleteHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeUser,
		PeerId:    1002,
		MaxId:     404,
	})
	if err != nil {
		t.Fatalf("MsgDeleteHistory() error = %v", err)
	}
	if got.Pts != 0 || got.PtsCount != 0 || got.Offset != 0 {
		t.Fatalf("affected history = %+v, want empty no-op", got)
	}
	if updatesClient.processed != nil || updatesClient.processWithEffects != nil {
		t.Fatalf("missing max_id should not dispatch operations: processed=%+v with_effects=%+v", updatesClient.processed, updatesClient.processWithEffects)
	}
}

func TestMsgSearchHashtagDelegatesPublicOffsetToRepository(t *testing.T) {
	repo := &fakeMsgRepository{
		history: []repository.HistoryMessage{{
			UserMessageID: 707,
			PeerSeq:       7,
			FromUserID:    1001,
			PeerType:      payload.PeerTypeUser,
			PeerID:        1002,
			MessageKind:   repository.MessageKindText,
			MessageText:   "#tag result",
			MessageDate:   1_772_000_001,
		}},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgSearchHashtag(&msgpb.TLMsgSearchHashtag{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		HashTag:  "#tag",
		OffsetId: 909,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgSearchHashtag() error = %v", err)
	}
	if repo.searchInput.OffsetID != 909 {
		t.Fatalf("search offset id = %d, want public offset id 909 delegated to repository", repo.searchInput.OffsetID)
	}
	if repo.resolveInput.UserMessageID != 0 {
		t.Fatalf("handler resolved offset id directly: %+v", repo.resolveInput)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok || len(messages.Messages) != 1 {
		t.Fatalf("messages = %+v ok=%v", got, ok)
	}
	msg, ok := messages.Messages[0].(*tg.TLMessage)
	if !ok || msg.Id != 707 {
		t.Fatalf("search result message = %+v ok=%v, want public id 707", messages.Messages[0], ok)
	}
}

func TestMsgSearchHashtagMissingPublicOffsetReturnsEmpty(t *testing.T) {
	repo := &fakeMsgRepository{
		searchNoMatch: true,
		history: []repository.HistoryMessage{{
			UserMessageID: 707,
			PeerSeq:       7,
			FromUserID:    1001,
			PeerType:      payload.PeerTypeUser,
			PeerID:        1002,
			MessageKind:   repository.MessageKindText,
			MessageText:   "#tag first page",
			MessageDate:   1_772_000_001,
		}},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.MsgSearchHashtag(&msgpb.TLMsgSearchHashtag{
		UserId:   1001,
		PeerType: payload.PeerTypeUser,
		PeerId:   1002,
		HashTag:  "#tag",
		OffsetId: 909,
		Limit:    20,
	})
	if err != nil {
		t.Fatalf("MsgSearchHashtag() error = %v", err)
	}
	messages, ok := got.ToMessagesMessages()
	if !ok {
		t.Fatalf("messages = %s, want messages.messages", got.ClazzName())
	}
	if len(messages.Messages) != 0 {
		t.Fatalf("messages len = %d, want empty for missing positive offset", len(messages.Messages))
	}
}

func TestTask6DialogOperationIDsUseV2ResolvedIdentity(t *testing.T) {
	if got := readHistoryOperationID(1001, 1002, 102, 9001); got != "v2:dialog:read_history:user:1001:peer:1002:max_user:102:auth:9001" {
		t.Fatalf("readHistoryOperationID() = %q", got)
	}
	if got := readHistoryOutboxOperationID(1002, 1001, 2); got != "v2:dialog:read_history_outbox:user:1002:peer:1001:max_peer_seq:2" {
		t.Fatalf("readHistoryOutboxOperationID() = %q", got)
	}
	if got := deleteMessagesOperationID(1001, 1002, []int32{107, 108}, true, 9001); got != "v2:dialog:delete_messages:user:1001:peer:1002:ids:[107 108]:revoke:true:auth:9001" {
		t.Fatalf("deleteMessagesOperationID() = %q", got)
	}
	if got := deleteMessagesPeerSeqOperationID(1002, 1001, []int64{7, 8}, true); got != "v2:dialog:delete_messages:user:1002:peer:1001:peer_seqs:[7 8]:revoke:true" {
		t.Fatalf("deleteMessagesPeerSeqOperationID() = %q", got)
	}
	if got := deleteHistoryOperationID(1001, 1002, 99, true, false, 9001); got != "v2:dialog:delete_history:user:1001:peer:1002:max_peer_seq:99:clear:true:revoke:false:auth:9001" {
		t.Fatalf("deleteHistoryOperationID() = %q", got)
	}
	if got := updatePinnedOperationID(1001, 1002, 107, false, 9001); got != "v2:dialog:update_pinned:user:1001:peer:1002:id:107:unpin:false:auth:9001" {
		t.Fatalf("updatePinnedOperationID() = %q", got)
	}
}

func TestMsgEditMessageV2UpdatesCanonicalAndRoutesOperations(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":41,"pts_count":1,"user_message_id":107}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
				MessageDate:        1_772_000_010,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7001,
			PeerSeq:            7,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             4294968298,
			MessageText:        "edited",
			MessageDate:        1_772_000_010,
			EditDate:           1_772_000_100,
			EditVersion:        1,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7001, 1, 1001),
			Status:              1,
			Pts:                 41,
			PtsCount:            1,
			CurrentPts:          41,
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

	got, err := core.MsgEditMessageV2(editMessageRequest(1001, 1002, 9001, 107, "edited"))
	if err != nil {
		t.Fatalf("MsgEditMessageV2() error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("expected updates, got %s", got.ClazzName())
	}
	if updates.Date != 1_772_000_099 || updates.Seq != 0 || len(updates.Updates) != 1 || len(updates.Users) != 0 {
		t.Fatalf("unexpected edit updates envelope: %+v", updates)
	}
	if updates.Users == nil {
		t.Fatalf("msg edit response must use an empty users vector, not nil, so BFF replacement is deterministic")
	}
	edit, ok := updates.Updates[0].(*tg.TLUpdateEditMessage)
	if !ok {
		t.Fatalf("expected updateEditMessage, got %T", updates.Updates[0])
	}
	if edit.Pts != 41 || edit.PtsCount != 1 {
		t.Fatalf("unexpected edit update pts: %+v", edit)
	}
	editMessage, ok := edit.Message.(*tg.TLMessage)
	if !ok {
		t.Fatalf("edit message type = %T, want *tg.TLMessage", edit.Message)
	}
	if peer, ok := editMessage.PeerId.(*tg.TLPeerUser); !ok || peer.UserId != 1002 {
		t.Fatalf("edit response peer_id = %#v, want peerUser(1002)", editMessage.PeerId)
	}
	if editMessage.Id != 107 {
		t.Fatalf("edit response id = %d, want public user_message_id 107", editMessage.Id)
	}
	if repo.editInput.NewMessageText != "edited" || repo.editInput.ActorUserID != 1001 || repo.editInput.PeerSeq != 7 {
		t.Fatalf("unexpected edit input: %+v", repo.editInput)
	}
	if updatesClient.processed == nil || updatesClient.processed.OperationId != editMessageOperationID(7001, 1, 1001) {
		t.Fatalf("sender edit operation was not sent to userupdates: %+v", updatesClient.processed)
	}
	if updatesClient.processed.OperationId != "v2:msg:7001:edit:1:1001" {
		t.Fatalf("sender edit operation_id = %q, want v2 id", updatesClient.processed.OperationId)
	}
	var senderOp payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processed.Payload, &senderOp); err != nil {
		t.Fatalf("decode sender edit payload: %v", err)
	}
	if senderOp.OperationKind != payload.OperationKindEditMessage || senderOp.MessageText != "edited" || senderOp.PeerSeq != 7 || senderOp.UserMessageID != 107 || senderOp.Date != 1_772_000_010 || senderOp.EditDate != 1_772_000_100 || senderOp.EditVersion != 1 {
		t.Fatalf("unexpected sender edit payload: %+v", senderOp)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != editMessageOperationID(7001, 1, 1002) {
		t.Fatalf("unexpected receiver edit operation: %+v", publisher.published)
	}
	if publisher.published.OperationID != "v2:msg:7001:edit:1:1002" {
		t.Fatalf("receiver edit operation_id = %q, want v2 id", publisher.published.OperationID)
	}
}

func TestMsgEditMessageV2OperationIDIncludesEditVersion(t *testing.T) {
	first := editMessageOperationID(7001, 1, 1001)
	second := editMessageOperationID(7001, 2, 1001)
	if first == second {
		t.Fatalf("edit operation ids must differ across edit versions: %q", first)
	}
	if first != "v2:msg:7001:edit:1:1001" || second != "v2:msg:7001:edit:2:1001" {
		t.Fatalf("unexpected edit operation ids: first=%q second=%q", first, second)
	}
}

func TestMsgEditMessageV2ReceiverDispatchUsesBrokerDurableAck(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":42,"pts_count":1,"user_message_id":108}`)
	responseHash := mustHashBytes(t, responsePayload)
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 108}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      108,
				PeerSeq:            8,
				CanonicalMessageID: 7101,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7101,
			PeerSeq:            8,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			MessageText:        "edited by broker ack",
			MessageDate:        1_772_000_020,
			EditDate:           1_772_000_120,
			EditVersion:        2,
		},
	}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7101, 2, 1001),
			Status:              1,
			Pts:                 42,
			PtsCount:            1,
			CurrentPts:          42,
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

	got, err := core.MsgEditMessageV2(editMessageRequest(1001, 1002, 9001, 108, "edited by broker ack"))
	if err != nil {
		t.Fatalf("MsgEditMessageV2() error = %v", err)
	}
	if _, ok := got.ToUpdates(); !ok {
		t.Fatalf("expected updates, got %s", got.ClazzName())
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if publisher.published.UserID != 1002 || publisher.published.OperationID != editMessageOperationID(7101, 2, 1002) {
		t.Fatalf("unexpected receiver edit operation: %+v", publisher.published)
	}
	if publisher.published.PeerID != 1001 || publisher.published.PayloadCodec != payload.PayloadCodecJSON {
		t.Fatalf("unexpected receiver edit metadata: %+v", publisher.published)
	}
	if len(updatesClient.processedList) != 1 || updatesClient.processWithEffects != nil {
		t.Fatalf("edit sender path should use requester sync only, processed=%d with_effects=%+v", len(updatesClient.processedList), updatesClient.processWithEffects)
	}

	publishErr := errors.New("broker unavailable")
	repo = &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 109}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      109,
				PeerSeq:            9,
				CanonicalMessageID: 7201,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7201,
			PeerSeq:            9,
			FromUserID:         1001,
			PeerType:           payload.PeerTypeUser,
			PeerID:             1002,
			MessageText:        "edited fail",
			MessageDate:        1_772_000_021,
			EditDate:           1_772_000_121,
			EditVersion:        3,
		},
	}
	updatesClient = &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
			UserId:              1001,
			OperationId:         editMessageOperationID(7201, 3, 1001),
			Status:              1,
			Pts:                 43,
			PtsCount:            1,
			CurrentPts:          43,
			ResponsePayload:     responsePayload,
			ResponsePayloadHash: responseHash,
		}),
	}
	publisher = &fakeReceiverPublisher{publishErr: publishErr}
	core = New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		UserUpdates:       updatesClient,
		ReceiverPublisher: publisher,
	})

	got, err = core.MsgEditMessageV2(editMessageRequest(1001, 1002, 9001, 109, "edited fail"))
	if err == nil {
		t.Fatalf("MsgEditMessageV2() error = nil, got=%+v", got)
	}
	if !errors.Is(err, msgpb.ErrReceiverBackpressure) {
		t.Fatalf("MsgEditMessageV2() error = %v, want ErrReceiverBackpressure", err)
	}
	if !errors.Is(err, publishErr) {
		t.Fatalf("MsgEditMessageV2() error = %v, want upstream publish error", err)
	}
	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
}

func TestMsgEditMessageV2RejectsNonAuthor(t *testing.T) {
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeUser, peerID: 1002, userMessageID: 107}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeUser,
				PeerID:             1002,
				UserMessageID:      107,
				PeerSeq:            7,
				CanonicalMessageID: 7001,
			},
		},
		editErr: msgpb.ErrMessageAuthorRequired,
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: &fakeUserUpdatesClient{}, ReceiverPublisher: &fakeReceiverPublisher{}})

	_, err := core.MsgEditMessageV2(editMessageRequest(1001, 1002, 9001, 107, "edited"))
	if !errors.Is(err, msgpb.ErrMessageAuthorRequired) {
		t.Fatalf("MsgEditMessageV2 error = %v, want %v", err, msgpb.ErrMessageAuthorRequired)
	}
}

type fakeMsgRepository struct {
	sendState           *repository.SendState
	canonical           *repository.CanonicalMessageResult
	canonicalByPeerSeq  *repository.CanonicalMessage
	editResult          *repository.EditMessageResult
	editErr             error
	editInput           repository.EditCanonicalMessageInput
	history             []repository.HistoryMessage
	historyInput        repository.ListHistoryMessagesInput
	searchInput         repository.SearchHashTagMessagesInput
	searchNoMatch       bool
	historyCursorBounds repository.HistoryCursorBounds
	resolveHistoryInput struct {
		userID   int64
		peerType int32
		peerID   int64
		offsetID int32
		maxID    int32
		minID    int32
	}
	resolvedMessageID            *repository.ResolvedMessageID
	resolveByUserMessageID       map[resolveMessageKey]*repository.ResolvedMessageID
	resolveManyByUserMessageID   map[int64]*repository.ResolvedMessageID
	resolveDeleteByUserMessageID map[int64]*repository.ResolvedMessageID
	resolveInput                 repository.ResolvedMessageID
	resolveDeleteCalls           int
	markCanonicalErr             error
	markSenderErrs               []error
	markCanonicalCalls           int
	markSenderCalls              int
	markReceiverAckedCalls       int
	markCompletedCalls           int
	markRetryableCalls           int
}

type resolveMessageKey struct {
	userID        int64
	peerType      int32
	peerID        int64
	userMessageID int64
}

func (f *fakeMsgRepository) CreateOrLoadSendState(context.Context, repository.CreateSendStateInput) (*repository.SendState, error) {
	return f.sendState, nil
}

func (f *fakeMsgRepository) CreateOrGetByClientRandom(context.Context, repository.CreateCanonicalMessageInput) (*repository.CanonicalMessageResult, error) {
	return f.canonical, nil
}

func (f *fakeMsgRepository) GetCanonicalMessageByPeerSeq(context.Context, int64, int32, int64, int64) (*repository.CanonicalMessage, error) {
	if f.canonicalByPeerSeq == nil {
		return nil, msgpb.ErrSendStateConflict
	}
	return f.canonicalByPeerSeq, nil
}

func (f *fakeMsgRepository) ListHistoryMessages(_ context.Context, in repository.ListHistoryMessagesInput) ([]repository.HistoryMessage, error) {
	f.historyInput = in
	if in.ResolvedCursorBounds.NoMatch {
		return []repository.HistoryMessage{}, nil
	}
	return f.history, nil
}

func (f *fakeMsgRepository) SearchHashTagMessages(_ context.Context, in repository.SearchHashTagMessagesInput) ([]repository.HistoryMessage, error) {
	f.searchInput = in
	if f.searchNoMatch && in.OffsetID > 0 {
		return []repository.HistoryMessage{}, nil
	}
	return f.history, nil
}

func (f *fakeMsgRepository) ResolveMessageID(_ context.Context, userID int64, peerType int32, peerID int64, userMessageID int64) (*repository.ResolvedMessageID, error) {
	f.resolveInput = repository.ResolvedMessageID{
		UserID:        userID,
		PeerType:      peerType,
		PeerID:        peerID,
		UserMessageID: userMessageID,
	}
	if f.resolveByUserMessageID != nil {
		return f.resolveByUserMessageID[resolveMessageKey{
			userID:        userID,
			peerType:      peerType,
			peerID:        peerID,
			userMessageID: userMessageID,
		}], nil
	}
	return f.resolvedMessageID, nil
}

func (f *fakeMsgRepository) ResolveMessageIDs(_ context.Context, userID int64, userMessageIDs []int64) ([]repository.ResolvedMessageID, error) {
	out := make([]repository.ResolvedMessageID, 0, len(userMessageIDs))
	for _, id := range userMessageIDs {
		f.resolveInput = repository.ResolvedMessageID{
			UserID:        userID,
			UserMessageID: id,
		}
		if f.resolveManyByUserMessageID != nil {
			if resolved := f.resolveManyByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolvedMessageID != nil && f.resolvedMessageID.UserMessageID == id {
			out = append(out, *f.resolvedMessageID)
		}
	}
	return out, nil
}

func (f *fakeMsgRepository) ResolveMessageIDsForDelete(_ context.Context, userID int64, userMessageIDs []int64) ([]repository.ResolvedMessageID, error) {
	f.resolveDeleteCalls++
	out := make([]repository.ResolvedMessageID, 0, len(userMessageIDs))
	for _, id := range userMessageIDs {
		f.resolveInput = repository.ResolvedMessageID{
			UserID:        userID,
			UserMessageID: id,
		}
		if f.resolveDeleteByUserMessageID != nil {
			if resolved := f.resolveDeleteByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolveManyByUserMessageID != nil {
			if resolved := f.resolveManyByUserMessageID[id]; resolved != nil {
				out = append(out, *resolved)
			}
			continue
		}
		if f.resolvedMessageID != nil && f.resolvedMessageID.UserMessageID == id {
			out = append(out, *f.resolvedMessageID)
		}
	}
	return out, nil
}

func (f *fakeMsgRepository) ResolveHistoryCursorIDs(_ context.Context, userID int64, peerType int32, peerID int64, offsetID int32, maxID int32, minID int32) (repository.HistoryCursorBounds, error) {
	f.resolveHistoryInput.userID = userID
	f.resolveHistoryInput.peerType = peerType
	f.resolveHistoryInput.peerID = peerID
	f.resolveHistoryInput.offsetID = offsetID
	f.resolveHistoryInput.maxID = maxID
	f.resolveHistoryInput.minID = minID
	return f.historyCursorBounds, nil
}

func (f *fakeMsgRepository) ResolvePeerSeqToUserMessageID(context.Context, int64, int32, int64, int64) (int64, error) {
	return 0, nil
}

func (f *fakeMsgRepository) EditCanonicalMessage(_ context.Context, in repository.EditCanonicalMessageInput) (*repository.EditMessageResult, error) {
	f.editInput = in
	if f.editErr != nil {
		return nil, f.editErr
	}
	return f.editResult, nil
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
	processed                *userupdates.TLUserOperation
	processedList            []*userupdates.TLUserOperation
	processedOperations      []*userupdates.TLUserOperation
	processResult            *userupdates.UserOperationResult
	processResults           []*userupdates.UserOperationResult
	processWithEffects       *userupdates.TLUserupdatesProcessUserOperationWithEffects
	processWithEffectsResult *userupdates.UserOperationResult
	processErr               error
	processWithEffectsErr    error
	getResult                *userupdates.UserOperationResult
	getOperationResultCalls  int
}

func (f *fakeUserUpdatesClient) nextProcessResult() *userupdates.UserOperationResult {
	if len(f.processResults) > 0 {
		result := f.processResults[0]
		f.processResults = f.processResults[1:]
		return result
	}
	return f.processResult
}

func (f *fakeUserUpdatesClient) UserupdatesProcessUserOperation(_ context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	f.processed = in.Operation
	f.processedList = append(f.processedList, in.Operation)
	f.processedOperations = append(f.processedOperations, in.Operation)
	if f.processErr != nil {
		return nil, f.processErr
	}
	return f.nextProcessResult(), nil
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

func editMessageRequest(userID, peerID, authKeyID int64, peerSeq int32, text string) *msgpb.TLMsgEditMessageV2 {
	return &msgpb.TLMsgEditMessageV2{
		UserId:    userID,
		AuthKeyId: authKeyID,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerID,
		NewMessage: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{Message: text}),
		}),
		DstMessage: tg.MakeTLMessageBox(&tg.TLMessageBox{
			MessageId: peerSeq,
		}),
	}
}

func mustHashBytes(t *testing.T, b []byte) []byte {
	t.Helper()
	return payload.HashBytes(b)
}
