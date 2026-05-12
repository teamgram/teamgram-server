package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/messages/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type messagesFakeChatClient struct {
	chatclient.ChatClient

	toggleNoForwards   func(context.Context, *chatpb.TLChatToggleNoForwards) (*tg.MutableChat, error)
	checkMessageAction func(context.Context, *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error)
	checkChatAccess    func(context.Context, *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error)
}

func (f *messagesFakeChatClient) ChatToggleNoForwards(ctx context.Context, in *chatpb.TLChatToggleNoForwards) (*tg.MutableChat, error) {
	return f.toggleNoForwards(ctx, in)
}

func (f *messagesFakeChatClient) ChatCheckMessageAction(ctx context.Context, in *chatpb.TLChatCheckMessageAction) (*chatpb.MessageActionCheckResult, error) {
	return f.checkMessageAction(ctx, in)
}

func (f *messagesFakeChatClient) ChatCheckChatAccess(ctx context.Context, in *chatpb.TLChatCheckChatAccess) (*chatpb.ChatAccessCheckResult, error) {
	return f.checkChatAccess(ctx, in)
}

func newMessagesCore(client chatclient.ChatClient, selfID int64) *MessagesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{ChatClient: client},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
	return c
}

func testMessagesMutableChat(id int64) *tg.MutableChat {
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

func TestMessagesToggleNoForwardsRejectsNonChatPeer(t *testing.T) {
	c := newMessagesCore(&messagesFakeChatClient{}, 100)

	_, err := c.MessagesToggleNoForwards(&tg.TLMessagesToggleNoForwards{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("MessagesToggleNoForwards error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
}

func TestMessagesToggleNoForwardsMapsChatError(t *testing.T) {
	c := newMessagesCore(&messagesFakeChatClient{
		toggleNoForwards: func(context.Context, *chatpb.TLChatToggleNoForwards) (*tg.MutableChat, error) {
			return nil, chatpb.ErrChatNotFound
		},
	}, 100)

	_, err := c.MessagesToggleNoForwards(&tg.TLMessagesToggleNoForwards{
		Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
	})
	if err != tg.ErrChatIdInvalid {
		t.Fatalf("MessagesToggleNoForwards error = %v, want %v", err, tg.ErrChatIdInvalid)
	}
}

func TestMessagesToggleNoForwardsMapsRequestFieldsAndUpdates(t *testing.T) {
	var got *chatpb.TLChatToggleNoForwards
	c := newMessagesCore(&messagesFakeChatClient{
		toggleNoForwards: func(_ context.Context, in *chatpb.TLChatToggleNoForwards) (*tg.MutableChat, error) {
			got = in
			return testMessagesMutableChat(in.ChatId), nil
		},
	}, 100)

	r, err := c.MessagesToggleNoForwards(&tg.TLMessagesToggleNoForwards{
		Peer:    tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		Enabled: tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("MessagesToggleNoForwards error = %v", err)
	}
	if got == nil || got.ChatId != 42 || got.OperatorId != 100 || !tg.FromBoolClazz(got.Enabled) {
		t.Fatalf("request = %+v, want chat_id=42 operator_id=100 enabled=true", got)
	}
	updates, ok := r.ToUpdates()
	if !ok {
		t.Fatalf("MessagesToggleNoForwards returned %s, want updates", r.ClazzName())
	}
	if len(updates.Updates) != 1 || len(updates.Chats) != 1 {
		t.Fatalf("updates lens = updates:%d chats:%d, want 1/1", len(updates.Updates), len(updates.Chats))
	}
	if update, ok := (&tg.Update{Clazz: updates.Updates[0]}).ToUpdateChat(); !ok || update.ChatId != 42 {
		t.Fatalf("update = %+v, ok=%v, want updateChat chat_id=42", update, ok)
	}
}
