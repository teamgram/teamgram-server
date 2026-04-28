package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type chatsFakeChatClient struct {
	chatclient.ChatClient

	getMutable func(context.Context, *chatpb.TLChatGetMutableChat) (*tg.MutableChat, error)
	editAbout  func(context.Context, *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error)
}

func (f *chatsFakeChatClient) ChatGetMutableChat(ctx context.Context, in *chatpb.TLChatGetMutableChat) (*tg.MutableChat, error) {
	return f.getMutable(ctx, in)
}

func (f *chatsFakeChatClient) ChatEditChatAbout(ctx context.Context, in *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error) {
	return f.editAbout(ctx, in)
}

func newChatsCore(client chatclient.ChatClient, selfID int64) *ChatsCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{ChatClient: client},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
	return c
}

func testMutableChat(id int64, title string) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                id,
			Creator:           100,
			Title:             title,
			ParticipantsCount: 3,
			Date:              10,
			Version:           2,
		}).ToImmutableChat(),
		ChatParticipants: []tg.ImmutableChatParticipantClazz{},
	}).ToMutableChat()
}

func TestMessagesEditChatAboutRejectsNonChatPeer(t *testing.T) {
	c := newChatsCore(&chatsFakeChatClient{}, 100)

	_, err := c.MessagesEditChatAbout(&tg.TLMessagesEditChatAbout{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("MessagesEditChatAbout error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
}

func TestMessagesEditChatAboutMapsChatError(t *testing.T) {
	c := newChatsCore(&chatsFakeChatClient{
		editAbout: func(context.Context, *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error) {
			return nil, chatpb.ErrChatNotFound
		},
	}, 100)

	_, err := c.MessagesEditChatAbout(&tg.TLMessagesEditChatAbout{
		Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
	})
	if err != tg.ErrChatIdInvalid {
		t.Fatalf("MessagesEditChatAbout error = %v, want %v", err, tg.ErrChatIdInvalid)
	}
}

func TestMessagesEditChatAboutMapsRequestFields(t *testing.T) {
	var got *chatpb.TLChatEditChatAbout
	c := newChatsCore(&chatsFakeChatClient{
		editAbout: func(_ context.Context, in *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error) {
			got = in
			return testMutableChat(in.ChatId, "chat"), nil
		},
	}, 100)

	r, err := c.MessagesEditChatAbout(&tg.TLMessagesEditChatAbout{
		Peer:  tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		About: "about",
	})
	if err != nil {
		t.Fatalf("MessagesEditChatAbout error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("MessagesEditChatAbout = %v, want BoolTrue", r)
	}
	if got == nil || got.ChatId != 42 || got.EditUserId != 100 || got.About != "about" {
		t.Fatalf("request = %+v, want chat_id=42 edit_user_id=100 about=about", got)
	}
}

func TestMessagesGetChatsSkipsMissingAndErrors(t *testing.T) {
	calls := make([]int64, 0, 3)
	c := newChatsCore(&chatsFakeChatClient{
		getMutable: func(_ context.Context, in *chatpb.TLChatGetMutableChat) (*tg.MutableChat, error) {
			calls = append(calls, in.ChatId)
			switch in.ChatId {
			case 1:
				return testMutableChat(1, "one"), nil
			case 2:
				return nil, errors.New("boom")
			default:
				return nil, nil
			}
		},
	}, 100)

	r, err := c.MessagesGetChats(&tg.TLMessagesGetChats{Id: []int64{1, 2, 3}})
	if err != nil {
		t.Fatalf("MessagesGetChats error = %v", err)
	}
	chats, ok := r.ToMessagesChats()
	if !ok {
		t.Fatalf("MessagesGetChats returned %s, want messages.chats", r.ClazzName())
	}
	if len(chats.Chats) != 1 {
		t.Fatalf("len(chats) = %d, want 1", len(chats.Chats))
	}
	if chat, ok := (&tg.Chat{Clazz: chats.Chats[0]}).ToChat(); !ok || chat.Id != 1 || chat.Title != "one" {
		t.Fatalf("projected chat = %+v, ok=%v", chat, ok)
	}
	if len(calls) != 3 || calls[0] != 1 || calls[1] != 2 || calls[2] != 3 {
		t.Fatalf("calls = %v, want [1 2 3]", calls)
	}
}
