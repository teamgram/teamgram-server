package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesReadHistoryAllowsInputPeerChat(t *testing.T) {
	var got *msgpb.TLMsgReadHistoryV2
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: allowAllMessagesChatClient(),
		MsgClient: &messagesFakeMsgClient{readHistoryV2: func(_ context.Context, in *msgpb.TLMsgReadHistoryV2) (*tg.MessagesAffectedMessages, error) {
			got = in
			return tg.MakeTLMessagesAffectedMessages(&tg.TLMessagesAffectedMessages{Pts: 1, PtsCount: 1}).ToMessagesAffectedMessages(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesReadHistory(&tg.TLMessagesReadHistory{Peer: inputPeerChat(55), MaxId: 10})
	if err != nil {
		t.Fatalf("MessagesReadHistory() error = %v", err)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 {
		t.Fatalf("msg request = %+v, want chat 55", got)
	}
}

func TestMessagesEditMessageAllowsInputPeerChat(t *testing.T) {
	var got *msgpb.TLMsgEditMessage
	text := "edited"
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: allowAllMessagesChatClient(),
		MsgClient: &messagesFakeMsgClient{editMessage: func(_ context.Context, in *msgpb.TLMsgEditMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesEditMessage(&tg.TLMessagesEditMessage{Peer: inputPeerChat(55), Id: 10, Message: &text})
	if err != nil {
		t.Fatalf("MessagesEditMessage() error = %v", err)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 || got.DstMessage.PeerType != payload.PeerTypeChat {
		t.Fatalf("msg request = %+v, want chat 55", got)
	}
}

func TestMessagesDeleteHistoryAllowsInputPeerChat(t *testing.T) {
	var got *msgpb.TLMsgDeleteHistory
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: allowAllMessagesChatClient(),
		MsgClient: &messagesFakeMsgClient{deleteHistory: func(_ context.Context, in *msgpb.TLMsgDeleteHistory) (*tg.MessagesAffectedHistory, error) {
			got = in
			return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{Pts: 2, PtsCount: 1}).ToMessagesAffectedHistory(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesDeleteHistory(&tg.TLMessagesDeleteHistory{Peer: inputPeerChat(55), MaxId: 10, Revoke: true})
	if err != nil {
		t.Fatalf("MessagesDeleteHistory() error = %v", err)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 || !got.Revoke {
		t.Fatalf("msg request = %+v, want chat 55 revoke", got)
	}
}

func TestMessagesUpdatePinnedMessageAllowsInputPeerChat(t *testing.T) {
	var got *msgpb.TLMsgUpdatePinnedMessage
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: allowAllMessagesChatClient(),
		MsgClient: &messagesFakeMsgClient{updatePinnedMessage: func(_ context.Context, in *msgpb.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
			got = in
			return testUpdates(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesUpdatePinnedMessage(&tg.TLMessagesUpdatePinnedMessage{Peer: inputPeerChat(55), Id: 10})
	if err != nil {
		t.Fatalf("MessagesUpdatePinnedMessage() error = %v", err)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 || got.Id != 10 {
		t.Fatalf("msg request = %+v, want chat 55", got)
	}
}

func TestMessagesUnpinAllMessagesAllowsInputPeerChat(t *testing.T) {
	var got *msgpb.TLMsgUnpinAllMessages
	core := newMessagesCoreWithRepo(&repository.Repository{
		ChatClient: allowAllMessagesChatClient(),
		MsgClient: &messagesFakeMsgClient{unpinAllMessages: func(_ context.Context, in *msgpb.TLMsgUnpinAllMessages) (*tg.MessagesAffectedHistory, error) {
			got = in
			return tg.MakeTLMessagesAffectedHistory(&tg.TLMessagesAffectedHistory{Pts: 3, PtsCount: 1}).ToMessagesAffectedHistory(), nil
		}},
	}, 1001, 9001)

	_, err := core.MessagesUnpinAllMessages(&tg.TLMessagesUnpinAllMessages{Peer: inputPeerChat(55)})
	if err != nil {
		t.Fatalf("MessagesUnpinAllMessages() error = %v", err)
	}
	if got == nil || got.PeerType != payload.PeerTypeChat || got.PeerId != 55 {
		t.Fatalf("msg request = %+v, want chat 55", got)
	}
}

func allowAllMessagesChatClient() *messagesFakeChatClient {
	return &messagesFakeChatClient{
		checkChatAccess: func(_ context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
			return chatpb.MakeTLChatAccessCheckResult(&chatpb.TLChatAccessCheckResult{
				SelfId: in.SelfId, ChatId: in.ChatId, AccessKind: in.AccessKind,
			}).ToChatAccessCheckResult(), nil
		},
		checkMessageAction: func(_ context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
			return chatpb.MakeTLMessageActionCheckResult(&chatpb.TLMessageActionCheckResult{
				SelfId: in.SelfId, ChatId: in.ChatId, Action: in.Action, MediaKind: in.MediaKind,
			}).ToMessageActionCheckResult(), nil
		},
	}
}
