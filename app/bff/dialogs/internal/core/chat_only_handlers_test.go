package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/dialogs/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type dialogsFakeChatClient struct {
	chatclient.ChatClient

	setHistoryTTL       func(context.Context, *chatpb.TLChatSetHistoryTTL) (*tg.MutableChat, error)
	getChatListByIDList func(context.Context, *chatpb.TLChatGetChatListByIdList) (*chatpb.VectorMutableChat, error)
}

func (f *dialogsFakeChatClient) ChatSetHistoryTTL(ctx context.Context, in *chatpb.TLChatSetHistoryTTL) (*tg.MutableChat, error) {
	return f.setHistoryTTL(ctx, in)
}

func (f *dialogsFakeChatClient) ChatGetChatListByIdList(ctx context.Context, in *chatpb.TLChatGetChatListByIdList) (*chatpb.VectorMutableChat, error) {
	return f.getChatListByIDList(ctx, in)
}

func newDialogsCore(client chatclient.ChatClient, selfID int64) *DialogsCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{ChatClient: client},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
	return c
}

func testDialogsMutableChat(id int64) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                id,
			Creator:           100,
			Title:             "chat",
			ParticipantsCount: 3,
			Date:              10,
			Version:           2,
		}).ToImmutableChat(),
		ChatParticipants: []tg.ImmutableChatParticipantClazz{},
	}).ToMutableChat()
}

func TestMessagesSetHistoryTTLRejectsNonChatPeer(t *testing.T) {
	c := newDialogsCore(&dialogsFakeChatClient{}, 100)

	_, err := c.MessagesSetHistoryTTL(&tg.TLMessagesSetHistoryTTL{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("MessagesSetHistoryTTL error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
}

func TestMessagesSetHistoryTTLMapsChatError(t *testing.T) {
	c := newDialogsCore(&dialogsFakeChatClient{
		setHistoryTTL: func(context.Context, *chatpb.TLChatSetHistoryTTL) (*tg.MutableChat, error) {
			return nil, chatpb.ErrChatNotFound
		},
	}, 100)

	_, err := c.MessagesSetHistoryTTL(&tg.TLMessagesSetHistoryTTL{
		Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
	})
	if err != tg.ErrChatIdInvalid {
		t.Fatalf("MessagesSetHistoryTTL error = %v, want %v", err, tg.ErrChatIdInvalid)
	}
}

func TestMessagesSetHistoryTTLMapsRequestFieldsAndUpdates(t *testing.T) {
	var got *chatpb.TLChatSetHistoryTTL
	c := newDialogsCore(&dialogsFakeChatClient{
		setHistoryTTL: func(_ context.Context, in *chatpb.TLChatSetHistoryTTL) (*tg.MutableChat, error) {
			got = in
			return testDialogsMutableChat(in.ChatId), nil
		},
	}, 100)

	r, err := c.MessagesSetHistoryTTL(&tg.TLMessagesSetHistoryTTL{
		Peer:   tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		Period: 86400,
	})
	if err != nil {
		t.Fatalf("MessagesSetHistoryTTL error = %v", err)
	}
	if got == nil || got.SelfId != 100 || got.ChatId != 42 || got.TtlPeriod != 86400 {
		t.Fatalf("request = %+v, want self_id=100 chat_id=42 ttl_period=86400", got)
	}
	updates, ok := r.ToUpdates()
	if !ok {
		t.Fatalf("MessagesSetHistoryTTL returned %s, want updates", r.ClazzName())
	}
	if len(updates.Updates) != 1 || len(updates.Chats) != 1 {
		t.Fatalf("updates lens = updates:%d chats:%d, want 1/1", len(updates.Updates), len(updates.Chats))
	}
	if update, ok := (&tg.Update{Clazz: updates.Updates[0]}).ToUpdateChat(); !ok || update.ChatId != 42 {
		t.Fatalf("update = %+v, ok=%v, want updateChat chat_id=42", update, ok)
	}
}
