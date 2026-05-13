package core

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
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

const createChatClientMsgID int64 = 7639238861095885056

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

	sendMessage func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error)
}

func (f *chatsFakeMsgClient) MsgSendMessage(ctx context.Context, in *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
	return f.sendMessage(ctx, in)
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

func withClientMsgID(c *ChatsCore, clientMsgID int64) *ChatsCore {
	c.MD.ClientMsgId = clientMsgID
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

func testCreatedMutableChat(id int64, title string) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                id,
			Creator:           100,
			Title:             title,
			ParticipantsCount: 3,
			Date:              10,
			Version:           2,
		}).ToImmutableChat(),
		ChatParticipants: []tg.ImmutableChatParticipantClazz{
			tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
				ChatId:          id,
				UserId:          100,
				State:           chatpb.ChatMemberStateNormal,
				ParticipantType: chatpb.ChatMemberCreator,
				InviterUserId:   100,
				InvitedAt:       1_772_000_000,
				Date:            1_772_000_000,
			}).ToImmutableChatParticipant(),
			tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
				ChatId:          id,
				UserId:          200,
				State:           chatpb.ChatMemberStateNormal,
				ParticipantType: chatpb.ChatMemberNormal,
				InviterUserId:   100,
				InvitedAt:       1_772_000_001,
				Date:            1_772_000_001,
			}).ToImmutableChatParticipant(),
			tg.MakeTLImmutableChatParticipant(&tg.TLImmutableChatParticipant{
				ChatId:          id,
				UserId:          300,
				State:           chatpb.ChatMemberStateNormal,
				ParticipantType: chatpb.ChatMemberNormal,
				InviterUserId:   100,
				InvitedAt:       1_772_000_002,
				Date:            1_772_000_002,
			}).ToImmutableChatParticipant(),
		},
	}).ToMutableChat()
}

func testMsgResponseUpdates() *tg.Updates {
	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{
			tg.MakeTLUpdateMessageID(&tg.TLUpdateMessageID{Id: 11, RandomId: 22}),
		},
		Date: 10,
		Seq:  1,
	}).ToUpdates()
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
				return testCreatedMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				return testMsgResponseUpdates(), nil
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
	if got.ClientMsgId != nil || got.OperationId != nil {
		t.Fatalf("client_msg_id/operation_id = %v/%v, want nil/nil by default", got.ClientMsgId, got.OperationId)
	}
}

func TestMessagesCreateChatSendsClientMsgIDOnlyWhenPresent(t *testing.T) {
	var got *chatpb.TLChatCreateChat2
	c := withClientMsgID(newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				got = in
				return testCreatedMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				return testMsgResponseUpdates(), nil
			},
		},
	}, 100), createChatClientMsgID)

	_, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
		},
		Title: "team",
	})
	if err != nil {
		t.Fatalf("MessagesCreateChat error = %v", err)
	}
	if got == nil {
		t.Fatalf("request is nil")
	}
	if got.ClientMsgId == nil || *got.ClientMsgId != createChatClientMsgID {
		t.Fatalf("client_msg_id = %v, want %d", got.ClientMsgId, createChatClientMsgID)
	}
	if got.OperationId != nil {
		t.Fatalf("operation_id = %v, want nil", got.OperationId)
	}
}

func TestMessagesCreateChatRejectsMissingMetadataBeforeSideEffects(t *testing.T) {
	called := false
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(context.Context, *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				called = true
				return testMutableChat(42, "team"), nil
			},
		},
	}, 100)
	c.MD = nil

	_, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
		},
		Title: "team",
	})
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("MessagesCreateChat error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
	if called {
		t.Fatalf("ChatCreateChat2 was called")
	}
}

func TestMessagesCreateChatRejectsMissingPermAuthKeyBeforeSideEffects(t *testing.T) {
	called := false
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(context.Context, *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				called = true
				return testMutableChat(42, "team"), nil
			},
		},
	}, 100)
	c.MD.PermAuthKeyId = 0

	_, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{
			tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
		},
		Title: "team",
	})
	if err != tg.ErrAuthKeyPermEmpty {
		t.Fatalf("MessagesCreateChat error = %v, want %v", err, tg.ErrAuthKeyPermEmpty)
	}
	if called {
		t.Fatalf("ChatCreateChat2 was called")
	}
}

func TestMessagesCreateChatSendsChatCreateServiceMessage(t *testing.T) {
	var sent *msgpb.TLMsgSendMessage
	c := withClientMsgID(newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				return testCreatedMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(_ context.Context, in *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				sent = in
				return testMsgResponseUpdates(), nil
			},
		},
	}, 100), createChatClientMsgID)

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
	if got, want := sent.Message[0].RandomId, expectedCreateChatServiceMessageRandomID(100, createChatClientMsgID); got != want {
		t.Fatalf("random_id = %d, want deterministic %d", got, want)
	}
	if sent.Message[0].RandomId == 0 {
		t.Fatalf("random_id = 0, want non-zero")
	}
	assertCreateChatAttachFact(t, sent, 42)
}

func TestMessagesCreateChatReturnsMsgEnvelopeWithoutAddingUpdateChat(t *testing.T) {
	msgUpdates := testMsgResponseUpdates()
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				return testCreatedMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				return msgUpdates, nil
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
	if r == nil || r.Updates != msgUpdates.Clazz {
		t.Fatalf("reply updates = %#v, want exact msg response envelope %#v", r, msgUpdates.Clazz)
	}
	updates, ok := r.Updates.(*tg.TLUpdates)
	if !ok {
		t.Fatalf("reply updates = %T, want updates", r.Updates)
	}
	for _, update := range updates.Updates {
		if _, ok := update.(*tg.TLUpdateChat); ok {
			t.Fatalf("reply contains updateChat: %+v", updates.Updates)
		}
	}
	if len(r.MissingInvitees) != 0 {
		t.Fatalf("missing invitees = %d, want 0", len(r.MissingInvitees))
	}
}

func TestMessagesCreateChatRejectsMalformedMutableChat(t *testing.T) {
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				return testMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				t.Fatalf("MsgSendMessage was called for malformed mutable chat")
				return nil, nil
			},
		},
	}, 100)

	_, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200})},
		Title: "team",
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesCreateChat error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesCreateChatRejectsNilMsgUpdates(t *testing.T) {
	c := newChatsCoreWithRepo(&repository.Repository{
		ChatClient: &chatsFakeChatClient{
			createChat: func(_ context.Context, in *chatpb.TLChatCreateChat2) (*tg.MutableChat, error) {
				return testCreatedMutableChat(42, in.Title), nil
			},
		},
		MsgClient: &chatsFakeMsgClient{
			sendMessage: func(context.Context, *msgpb.TLMsgSendMessage) (*tg.Updates, error) {
				return nil, nil
			},
		},
	}, 100)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("MessagesCreateChat panicked for nil msg updates: %v", r)
		}
	}()
	_, err := c.MessagesCreateChat(&tg.TLMessagesCreateChat{
		Users: []tg.InputUserClazz{tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200})},
		Title: "team",
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesCreateChat error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestCreateChatServiceMessageRandomIDZeroDigestFallback(t *testing.T) {
	if got := createChatRandomIDFromDigest(make([]byte, 8)); got != 1 {
		t.Fatalf("random_id fallback = %d, want 1", got)
	}
}

func TestNormalizeCreateChatServiceMessageRandomID(t *testing.T) {
	if got := normalizeCreateChatServiceMessageRandomID(0); got != 1 {
		t.Fatalf("normalized random_id = %d, want 1", got)
	}
	if got := normalizeCreateChatServiceMessageRandomID(99); got != 99 {
		t.Fatalf("normalized random_id = %d, want 99", got)
	}
}

func TestChatParticipantsChangedFactFromMutableChatValidatesExpectedParticipants(t *testing.T) {
	tests := []struct {
		name     string
		chat     func() *tg.MutableChat
		actorID  int64
		invitees []int64
	}{
		{
			name: "missing actor",
			chat: func() *tg.MutableChat {
				chat := testCreatedMutableChat(42, "team")
				chat.ChatParticipants = chat.ChatParticipants[1:]
				return chat
			},
			actorID:  100,
			invitees: []int64{200, 300},
		},
		{
			name: "missing invitee",
			chat: func() *tg.MutableChat {
				chat := testCreatedMutableChat(42, "team")
				chat.ChatParticipants = chat.ChatParticipants[:2]
				return chat
			},
			actorID:  100,
			invitees: []int64{200, 300},
		},
		{
			name: "duplicate participant",
			chat: func() *tg.MutableChat {
				chat := testCreatedMutableChat(42, "team")
				chat.ChatParticipants = append(chat.ChatParticipants, chat.ChatParticipants[1])
				return chat
			},
			actorID:  100,
			invitees: []int64{200, 300},
		},
		{
			name: "actor is not creator",
			chat: func() *tg.MutableChat {
				chat := testCreatedMutableChat(42, "team")
				chat.ChatParticipants[0].ParticipantType = chatpb.ChatMemberAdmin
				return chat
			},
			actorID:  100,
			invitees: []int64{200, 300},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := chatParticipantsChangedFactFromMutableChat(tt.chat(), tt.actorID, tt.invitees)
			if err == nil {
				t.Fatalf("chatParticipantsChangedFactFromMutableChat error = nil, want error")
			}
		})
	}
}

func TestChatParticipantsChangedFactFromMutableChatMapsAdminRole(t *testing.T) {
	chat := testCreatedMutableChat(42, "team")
	chat.ChatParticipants[2].ParticipantType = chatpb.ChatMemberAdmin

	fact, err := chatParticipantsChangedFactFromMutableChat(chat, 100, []int64{200, 300})
	if err != nil {
		t.Fatalf("chatParticipantsChangedFactFromMutableChat error = %v", err)
	}
	if got := fact.Participants[2].Role; got != "admin" {
		t.Fatalf("participant role = %q, want admin", got)
	}
}

func expectedCreateChatServiceMessageRandomID(actorUserID, clientMsgID int64) int64 {
	sum := sha256.Sum256([]byte(fmt.Sprintf("create_chat:%d:%d:service_message", actorUserID, clientMsgID)))
	return int64(binary.BigEndian.Uint64(sum[:8]) & uint64(^uint64(0)>>1))
}

func assertCreateChatAttachFact(t *testing.T, sent *msgpb.TLMsgSendMessage, chatID int64) {
	t.Helper()
	if sent == nil {
		t.Fatalf("send request is nil")
	}
	if len(sent.AttachFacts) != 1 {
		t.Fatalf("attach facts len = %d, want 1", len(sent.AttachFacts))
	}
	attach := sent.AttachFacts[0]
	if attach == nil {
		t.Fatalf("attach fact is nil")
	}
	if attach.Kind != payload.FactKindChatParticipantsChanged {
		t.Fatalf("attach kind = %q, want %q", attach.Kind, payload.FactKindChatParticipantsChanged)
	}

	var envelope payload.UpdateFactV1
	if err := json.Unmarshal(attach.Payload, &envelope); err != nil {
		t.Fatalf("Unmarshal attach payload error = %v", err)
	}
	if envelope.Kind != payload.FactKindChatParticipantsChanged {
		t.Fatalf("envelope kind = %q, want %q", envelope.Kind, payload.FactKindChatParticipantsChanged)
	}
	decoded, err := payload.DecodeUpdateFact(envelope)
	if err != nil {
		t.Fatalf("DecodeUpdateFact error = %v", err)
	}
	fact, ok := decoded.(payload.ChatParticipantsChangedFactV1)
	if !ok {
		t.Fatalf("decoded attach fact = %T, want ChatParticipantsChangedFactV1", decoded)
	}
	if fact.ChatID != chatID || fact.ActorUserID != 100 || fact.Version != 2 {
		t.Fatalf("chat fact = %+v, want chat_id=%d actor=100 version=2", fact, chatID)
	}
	want := []payload.ChatParticipantFactV1{
		{UserID: 100, Role: "creator", InviterUserID: 100, Date: 1_772_000_000},
		{UserID: 200, Role: "member", InviterUserID: 100, Date: 1_772_000_001},
		{UserID: 300, Role: "member", InviterUserID: 100, Date: 1_772_000_002},
	}
	if len(fact.Participants) != len(want) {
		t.Fatalf("participants len = %d, want %d: %+v", len(fact.Participants), len(want), fact.Participants)
	}
	for i := range want {
		if fact.Participants[i] != want[i] {
			t.Fatalf("participant[%d] = %+v, want %+v", i, fact.Participants[i], want[i])
		}
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
