package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/chats/internal/svc"
	msgclient "github.com/teamgram/teamgram-server/v2/app/messenger/msg/client"
	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type chatsFakeChatClient struct {
	chatclient.ChatClient

	getMutable     func(context.Context, *chatpb.TLChatGetMutableChat) (*tg.MutableChat, error)
	editAbout      func(context.Context, *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error)
	editTitle      func(context.Context, *chatpb.TLChatEditChatTitle) (*tg.MutableChat, error)
	editDefault    func(context.Context, *chatpb.TLChatEditChatDefaultBannedRights) (*tg.MutableChat, error)
	deleteChat     func(context.Context, *chatpb.TLChatDeleteChat) (*tg.MutableChat, error)
	editAdmin      func(context.Context, *chatpb.TLChatEditChatAdmin) (*tg.MutableChat, error)
	addChatUser    func(context.Context, *chatpb.TLChatAddChatUser) (*tg.MutableChat, error)
	deleteChatUser func(context.Context, *chatpb.TLChatDeleteChatUser) (*tg.MutableChat, error)
	createChat     func(context.Context, *chatpb.TLChatCreateChat2) (*tg.MutableChat, error)
}

func (f *chatsFakeChatClient) ChatGetMutableChat(ctx context.Context, in *chatpb.TLChatGetMutableChat) (*tg.MutableChat, error) {
	return f.getMutable(ctx, in)
}

func (f *chatsFakeChatClient) ChatEditChatAbout(ctx context.Context, in *chatpb.TLChatEditChatAbout) (*tg.MutableChat, error) {
	return f.editAbout(ctx, in)
}

func (f *chatsFakeChatClient) ChatEditChatTitle(ctx context.Context, in *chatpb.TLChatEditChatTitle) (*tg.MutableChat, error) {
	return f.editTitle(ctx, in)
}

func (f *chatsFakeChatClient) ChatEditChatDefaultBannedRights(ctx context.Context, in *chatpb.TLChatEditChatDefaultBannedRights) (*tg.MutableChat, error) {
	return f.editDefault(ctx, in)
}

func (f *chatsFakeChatClient) ChatDeleteChat(ctx context.Context, in *chatpb.TLChatDeleteChat) (*tg.MutableChat, error) {
	return f.deleteChat(ctx, in)
}

func (f *chatsFakeChatClient) ChatEditChatAdmin(ctx context.Context, in *chatpb.TLChatEditChatAdmin) (*tg.MutableChat, error) {
	return f.editAdmin(ctx, in)
}

func (f *chatsFakeChatClient) ChatAddChatUser(ctx context.Context, in *chatpb.TLChatAddChatUser) (*tg.MutableChat, error) {
	return f.addChatUser(ctx, in)
}

func (f *chatsFakeChatClient) ChatDeleteChatUser(ctx context.Context, in *chatpb.TLChatDeleteChatUser) (*tg.MutableChat, error) {
	return f.deleteChatUser(ctx, in)
}

func (f *chatsFakeChatClient) ChatCreateChat2(ctx context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
	return f.createChat(ctx, in)
}

type chatsFakeMsgClient struct {
	msgclient.MsgClient

	sendMessageV2 func(context.Context, *msgpb.TLMsgSendMessageV2) (*tg.Updates, error)
}

func (f *chatsFakeMsgClient) MsgSendMessageV2(ctx context.Context, in *msgpb.TLMsgSendMessageV2) (*tg.Updates, error) {
	return f.sendMessageV2(ctx, in)
}

func newChatsCore(client chatclient.ChatClient, selfID int64) *ChatsCore {
	return newChatsCoreWithRepo(&repository.Repository{ChatClient: client}, selfID)
}

func newChatsCoreWithRepo(repo *repository.Repository, selfID int64) *ChatsCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: repo,
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID, PermAuthKeyId: 9001}
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

func TestMessagesEditChatTitleMapsRequestAndUpdates(t *testing.T) {
	var got *chatpb.TLChatEditChatTitle
	c := newChatsCore(&chatsFakeChatClient{
		editTitle: func(_ context.Context, in *chatpb.TLChatEditChatTitle) (*tg.MutableChat, error) {
			got = in
			return testMutableChat(in.ChatId, in.Title), nil
		},
	}, 100)

	r, err := c.MessagesEditChatTitle(&tg.TLMessagesEditChatTitle{ChatId: 42, Title: "new"})
	if err != nil {
		t.Fatalf("MessagesEditChatTitle error = %v", err)
	}
	if got == nil || got.ChatId != 42 || got.EditUserId != 100 || got.Title != "new" {
		t.Fatalf("request = %+v, want chat_id=42 edit_user_id=100 title=new", got)
	}
	assertUpdateChat(t, r, 42)
}

func TestMessagesEditChatDefaultBannedRightsRejectsNonChatPeer(t *testing.T) {
	c := newChatsCore(&chatsFakeChatClient{}, 100)

	_, err := c.MessagesEditChatDefaultBannedRights(&tg.TLMessagesEditChatDefaultBannedRights{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("MessagesEditChatDefaultBannedRights error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
}

func TestMessagesDeleteChatMapsRequestAndErrors(t *testing.T) {
	var got *chatpb.TLChatDeleteChat
	c := newChatsCore(&chatsFakeChatClient{
		deleteChat: func(_ context.Context, in *chatpb.TLChatDeleteChat) (*tg.MutableChat, error) {
			got = in
			return nil, chatpb.ErrChatAdminRequired
		},
	}, 100)

	_, err := c.MessagesDeleteChat(&tg.TLMessagesDeleteChat{ChatId: 42})
	if err != tg.Err400ChatAdminRequired {
		t.Fatalf("MessagesDeleteChat error = %v, want %v", err, tg.Err400ChatAdminRequired)
	}
	if got == nil || got.ChatId != 42 || got.OperatorId != 100 {
		t.Fatalf("request = %+v, want chat_id=42 operator_id=100", got)
	}
}

func TestMessagesEditChatAdminRejectsInvalidUser(t *testing.T) {
	c := newChatsCore(&chatsFakeChatClient{}, 100)

	_, err := c.MessagesEditChatAdmin(&tg.TLMessagesEditChatAdmin{
		ChatId:  42,
		UserId:  tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}),
		IsAdmin: tg.BoolTrueClazz,
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("MessagesEditChatAdmin error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}

func TestMessagesAddAndDeleteChatUserMapInputUser(t *testing.T) {
	var addReq *chatpb.TLChatAddChatUser
	var deleteReq *chatpb.TLChatDeleteChatUser
	c := newChatsCore(&chatsFakeChatClient{
		addChatUser: func(_ context.Context, in *chatpb.TLChatAddChatUser) (*tg.MutableChat, error) {
			addReq = in
			return testMutableChat(in.ChatId, "chat"), nil
		},
		deleteChatUser: func(_ context.Context, in *chatpb.TLChatDeleteChatUser) (*tg.MutableChat, error) {
			deleteReq = in
			return testMutableChat(in.ChatId, "chat"), nil
		},
	}, 100)

	inUser := tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200})
	added, err := c.MessagesAddChatUser(&tg.TLMessagesAddChatUser{ChatId: 42, UserId: inUser})
	if err != nil {
		t.Fatalf("MessagesAddChatUser error = %v", err)
	}
	if added == nil || added.Updates == nil {
		t.Fatalf("MessagesAddChatUser = %+v, want updates", added)
	}
	if addReq == nil || addReq.ChatId != 42 || addReq.InviterId != 100 || addReq.UserId != 200 {
		t.Fatalf("add request = %+v, want chat_id=42 inviter_id=100 user_id=200", addReq)
	}

	updates, err := c.MessagesDeleteChatUser(&tg.TLMessagesDeleteChatUser{ChatId: 42, UserId: inUser})
	if err != nil {
		t.Fatalf("MessagesDeleteChatUser error = %v", err)
	}
	assertUpdateChat(t, updates, 42)
	if deleteReq == nil || deleteReq.ChatId != 42 || deleteReq.OperatorId != 100 || deleteReq.DeleteUserId != 200 {
		t.Fatalf("delete request = %+v, want chat_id=42 operator_id=100 delete_user_id=200", deleteReq)
	}
}

func TestMessagesCreateChatMapsUsersAndTitle(t *testing.T) {
	var got *chatpb.TLChatCreateChat2
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				got = in
				return testMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessageV2: func(context.Context, *msgpb.TLMsgSendMessageV2) (*tg.Updates, error) {
				return updatesWithChat(testMutableChat(42, "team"), 100), nil
			},
		},
	}, 100)

	r, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 300}),
		},
		Title: "team",
	})
	if err != nil {
		t.Fatalf("MessagesCreateChat error = %v", err)
	}
	if r == nil || r.Updates == nil {
		t.Fatalf("MessagesCreateChat = %+v, want updates", r)
	}
	if got == nil || got.CreatorId != 100 || got.Title != "team" || len(got.UserIdList) != 2 || got.UserIdList[0] != 200 || got.UserIdList[1] != 300 {
		t.Fatalf("request = %+v, want creator/title/users", got)
	}
}

func TestMessagesCreateChatSendsChatCreateServiceMessage(t *testing.T) {
	var sent *msgpb.TLMsgSendMessageV2
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				return testMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessageV2: func(_ context.Context, in *msgpb.TLMsgSendMessageV2) (*tg.Updates, error) {
				sent = in
				return updatesWithChat(testMutableChat(42, "team"), 100), nil
			},
		},
	}, 100)

	r, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 300}),
		},
		Title: "team",
	})
	if err != nil {
		t.Fatalf("MessagesCreateChat error = %v", err)
	}
	if r == nil || r.Updates == nil {
		t.Fatalf("MessagesCreateChat = %+v, want sent updates", r)
	}
	if sent == nil || sent.UserId != 100 || sent.AuthKeyId != 9001 || sent.PeerType != payload.PeerTypeChat || sent.PeerId != 42 || len(sent.Message) != 1 {
		t.Fatalf("send request = %+v, want one chat service message to chat 42", sent)
	}
	service, ok := sent.Message[0].Message.(*tg.TLMessageService)
	if !ok {
		t.Fatalf("outbox message = %T, want messageService", sent.Message[0].Message)
	}
	if peer, ok := service.PeerId.(*tg.TLPeerChat); !ok || peer.ChatId != 42 {
		t.Fatalf("service peer = %#v, want peerChat 42", service.PeerId)
	}
	action, ok := service.Action.(*tg.TLMessageActionChatCreate)
	if !ok {
		t.Fatalf("service action = %T, want messageActionChatCreate", service.Action)
	}
	if action.Title != "team" || len(action.Users) != 2 || action.Users[0] != 200 || action.Users[1] != 300 {
		t.Fatalf("chat create action = %+v, want title/users", action)
	}
}

func assertUpdateChat(t *testing.T, updates *tg.Updates, chatID int64) {
	t.Helper()
	if updates == nil {
		t.Fatalf("updates is nil")
	}
	data, ok := updates.ToUpdates()
	if !ok || len(data.Updates) != 1 {
		t.Fatalf("updates = %#v, want one update", updates)
	}
	update, ok := (&tg.Update{Clazz: data.Updates[0]}).ToUpdateChat()
	if !ok || update.ChatId != chatID {
		t.Fatalf("update = %#v, ok=%v, want updateChat %d", data.Updates[0], ok, chatID)
	}
}
