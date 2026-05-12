package core

import (
	"errors"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestCheckMessageActionSendTextRequiresActiveMember(t *testing.T) {
	core := newTestCore(&fakeReadRepo{
		mutableChat: mutableChatForMemberTests(10, 1,
			participantForMemberTests(10, 1, chatpb.ChatMemberCreator, chatpb.ChatMemberStateNormal, nil),
			participantForMemberTests(10, 2, chatpb.ChatMemberNormal, chatpb.ChatMemberStateLeft, nil)),
	})

	_, err := core.ChatCheckMessageAction(&chatpb.TLChatCheckMessageAction{
		SelfId: 2,
		ChatId: 10,
		Action: chatpb.MessageActionSendText,
	})
	if !errors.Is(err, chatpb.ErrUserNotParticipant) {
		t.Fatalf("ChatCheckMessageAction() error = %v, want ErrUserNotParticipant", err)
	}
}

func TestCheckMessageActionSendTextHonorsBannedRights(t *testing.T) {
	mChat := mutableChatForMemberTests(10, 1,
		participantForMemberTests(10, 1, chatpb.ChatMemberCreator, chatpb.ChatMemberStateNormal, nil),
		participantForMemberTests(10, 2, chatpb.ChatMemberNormal, chatpb.ChatMemberStateNormal, nil))
	mChat.Chat.DefaultBannedRights = tg.MakeTLChatBannedRights(&tg.TLChatBannedRights{SendMessages: true}).ToChatBannedRights()
	core := newTestCore(&fakeReadRepo{mutableChat: mChat})

	_, err := core.ChatCheckMessageAction(&chatpb.TLChatCheckMessageAction{
		SelfId: 2,
		ChatId: 10,
		Action: chatpb.MessageActionSendText,
	})
	if !errors.Is(err, chatpb.ErrChatWriteForbidden) {
		t.Fatalf("ChatCheckMessageAction() error = %v, want ErrChatWriteForbidden", err)
	}
}

func TestCheckMessageActionUnsupportedPollIsRejectedInChatService(t *testing.T) {
	core := newTestCore(&fakeReadRepo{
		mutableChat: mutableChatForMemberTests(10, 1,
			participantForMemberTests(10, 1, chatpb.ChatMemberCreator, chatpb.ChatMemberStateNormal, nil)),
	})

	_, err := core.ChatCheckMessageAction(&chatpb.TLChatCheckMessageAction{
		SelfId:    1,
		ChatId:    10,
		Action:    chatpb.MessageActionSendPoll,
		MediaKind: "poll",
	})
	if !errors.Is(err, chatpb.ErrMessageActionUnsupported) {
		t.Fatalf("ChatCheckMessageAction() error = %v, want ErrMessageActionUnsupported", err)
	}
}

func TestCheckMessageActionPinRequiresPinRight(t *testing.T) {
	core := newTestCore(&fakeReadRepo{
		mutableChat: mutableChatForMemberTests(10, 1,
			participantForMemberTests(10, 1, chatpb.ChatMemberCreator, chatpb.ChatMemberStateNormal, nil),
			participantForMemberTests(10, 2, chatpb.ChatMemberNormal, chatpb.ChatMemberStateNormal, nil)),
	})

	_, err := core.ChatCheckMessageAction(&chatpb.TLChatCheckMessageAction{
		SelfId: 2,
		ChatId: 10,
		Action: chatpb.MessageActionPinMessage,
	})
	if !errors.Is(err, chatpb.ErrChatAdminRequired) {
		t.Fatalf("ChatCheckMessageAction() error = %v, want ErrChatAdminRequired", err)
	}
}
