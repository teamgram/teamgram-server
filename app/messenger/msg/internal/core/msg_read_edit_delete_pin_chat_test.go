package core

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/svc"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMsgReadHistoryV2AllowsPeerTypeChat(t *testing.T) {
	chatClient := &fakeMsgChatClient{}
	updatesClient := &fakeUserUpdatesClient{
		processResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Pts: 31, PtsCount: 1}),
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        &fakeMsgRepository{},
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	got, err := core.MsgReadHistoryV2(&msgpb.TLMsgReadHistoryV2{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		MaxId:     0,
	})
	if err != nil {
		t.Fatalf("MsgReadHistoryV2() error = %v", err)
	}
	if got == nil || got.Pts != 31 {
		t.Fatalf("affected = %+v, want pts 31", got)
	}
	if len(chatClient.accesses) != 1 || chatClient.accesses[0].AccessKind != chatpb.ChatAccessReadHistory {
		t.Fatalf("chat access checks = %+v", chatClient.accesses)
	}
	if updatesClient.processWithEffects != nil {
		t.Fatalf("chat read history should not create peer outbox effect: %+v", updatesClient.processWithEffects)
	}
}

func TestMsgUpdatePinnedMessageRequiresChatActionAndFansOut(t *testing.T) {
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	updatesClient := &fakeUserUpdatesClient{processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Pts: 32, PtsCount: 1})}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeChat, peerID: 55, userMessageID: 10}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeChat,
				PeerID:             55,
				UserMessageID:      10,
				PeerSeq:            4,
				CanonicalMessageID: 7001,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	_, err := core.MsgUpdatePinnedMessage(&msgpb.TLMsgUpdatePinnedMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		Id:        10,
	})
	if err != nil {
		t.Fatalf("MsgUpdatePinnedMessage() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionPinMessage {
		t.Fatalf("chat action checks = %+v, want pin_message", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("pin affected effects = %+v, want two chat receivers", updatesClient.processWithEffects)
	}
}

func TestMsgEditMessagePeerTypeChatFansOut(t *testing.T) {
	responsePayload := []byte(`{"schema_version":2,"pts":34,"pts_count":1,"event_type":"edit_message","user_message_id":10}`)
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	updatesClient := &fakeUserUpdatesClient{processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{
		Pts:                 34,
		PtsCount:            1,
		ResponsePayload:     responsePayload,
		ResponsePayloadHash: payload.HashBytes(responsePayload),
	})}
	repo := &fakeMsgRepository{
		resolveByUserMessageID: map[resolveMessageKey]*repository.ResolvedMessageID{
			{userID: 1001, peerType: payload.PeerTypeChat, peerID: 55, userMessageID: 10}: {
				UserID:             1001,
				PeerType:           payload.PeerTypeChat,
				PeerID:             55,
				UserMessageID:      10,
				PeerSeq:            4,
				CanonicalMessageID: 7001,
			},
		},
		editResult: &repository.EditMessageResult{
			CanonicalMessageID: 7001,
			PeerSeq:            4,
			EditVersion:        2,
			MessageDate:        1_772_000_000,
			EditDate:           1_772_000_010,
			FromUserID:         1001,
			MessageText:        "edited",
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, UserUpdates: updatesClient, Chat: chatClient})

	_, err := core.MsgEditMessage(&msgpb.TLMsgEditMessage{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		NewMessage: msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
			Message: testTLMessage("edited"),
		}),
		DstMessage: testMessageBox(1001, payload.PeerTypeChat, 55, 10),
	})
	if err != nil {
		t.Fatalf("MsgEditMessage() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionEditOwnMessage {
		t.Fatalf("chat action checks = %+v, want edit_own_message", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("edit affected effects = %+v, want two chat receivers", updatesClient.processWithEffects)
	}
	var op payload.MessageOperationV1
	if err := json.Unmarshal(updatesClient.processWithEffects.AffectedEffects[0].Operation.Payload, &op); err != nil {
		t.Fatalf("decode edit receiver payload: %v", err)
	}
	if op.PeerType != payload.PeerTypeChat || op.PeerID != 55 || op.ToUserID == 55 {
		t.Fatalf("receiver edit payload = %+v, want chat peer and user to_user_id", op)
	}
}

func TestMsgDeleteHistoryRevokePeerTypeChatFansOut(t *testing.T) {
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	updatesClient := &fakeUserUpdatesClient{processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Pts: 35, PtsCount: 1})}
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeMsgRepository{}, UserUpdates: updatesClient, Chat: chatClient})

	_, err := core.MsgDeleteHistory(&msgpb.TLMsgDeleteHistory{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
		Revoke:    true,
	})
	if err != nil {
		t.Fatalf("MsgDeleteHistory() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionDeleteRevoke {
		t.Fatalf("chat action checks = %+v, want delete_revoke", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("delete history affected effects = %+v, want two chat receivers", updatesClient.processWithEffects)
	}
}

func TestMsgDeleteMessagesPeerTypeZeroResolvesChatForRevoke(t *testing.T) {
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	updatesClient := &fakeUserUpdatesClient{processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Pts: 33, PtsCount: 1})}
	repo := &fakeMsgRepository{
		resolveDeleteByUserMessageID: map[int64]*repository.ResolvedMessageID{
			10: {
				UserID:             1001,
				PeerType:           payload.PeerTypeChat,
				PeerID:             55,
				UserMessageID:      10,
				PeerSeq:            4,
				CanonicalMessageID: 7001,
				Outgoing:           true,
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:        repo,
		UserUpdates: updatesClient,
		Chat:        chatClient,
	})

	_, err := core.MsgDeleteMessages(&msgpb.TLMsgDeleteMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  0,
		PeerId:    0,
		Revoke:    true,
		Id:        []int32{10},
	})
	if err != nil {
		t.Fatalf("MsgDeleteMessages() error = %v", err)
	}
	if len(chatClient.accesses) != 1 || chatClient.accesses[0].ChatId != 55 {
		t.Fatalf("chat access checks = %+v, want chat 55", chatClient.accesses)
	}
	if len(chatClient.actions) != 0 {
		t.Fatalf("own outgoing revoke should not require admin action, got %+v", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("delete affected effects = %+v, want two chat receivers", updatesClient.processWithEffects)
	}
}

func TestMsgUnpinAllMessagesRequiresChatActionAndFansOut(t *testing.T) {
	chatClient := &fakeMsgChatClient{memberIDs: []int64{1001, 1002, 1003}}
	updatesClient := &fakeUserUpdatesClient{processWithEffectsResult: userupdates.MakeTLUserOperationResult(&userupdates.TLUserOperationResult{Pts: 36, PtsCount: 1})}
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeMsgRepository{}, UserUpdates: updatesClient, Chat: chatClient})

	_, err := core.MsgUnpinAllMessages(&msgpb.TLMsgUnpinAllMessages{
		UserId:    1001,
		AuthKeyId: 9001,
		PeerType:  payload.PeerTypeChat,
		PeerId:    55,
	})
	if err != nil {
		t.Fatalf("MsgUnpinAllMessages() error = %v", err)
	}
	if len(chatClient.actions) != 1 || chatClient.actions[0].Action != chatpb.MessageActionUnpinAll {
		t.Fatalf("chat action checks = %+v, want unpin_all", chatClient.actions)
	}
	if updatesClient.processWithEffects == nil || len(updatesClient.processWithEffects.AffectedEffects) != 2 {
		t.Fatalf("unpin affected effects = %+v, want two chat receivers", updatesClient.processWithEffects)
	}
}

func testTLMessage(text string) tg.MessageClazz {
	return tg.MakeTLMessage(&tg.TLMessage{Message: text})
}

func testMessageBox(userID int64, peerType int32, peerID int64, messageID int32) tg.MessageBoxClazz {
	peer := tg.MakePeerUser(peerID)
	if peerType == payload.PeerTypeChat {
		peer = tg.MakePeerChat(peerID)
	}
	return tg.MakeTLMessageBox(&tg.TLMessageBox{
		UserId:    userID,
		MessageId: messageID,
		PeerType:  peerType,
		PeerId:    peerID,
		Message:   tg.MakeTLMessage(&tg.TLMessage{Id: messageID, PeerId: peer}),
	})
}
