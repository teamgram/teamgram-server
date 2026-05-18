package core

import (
	"context"
	"errors"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type chatInvitesFakeUserClient struct {
	userclient.UserClient

	getMutableUsers func(context.Context, *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error)
	getProjection   func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
}

func (f *chatInvitesFakeUserClient) UserGetMutableUsers(ctx context.Context, in *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error) {
	return f.getMutableUsers(ctx, in)
}

func (f *chatInvitesFakeUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	return f.getProjection(ctx, in)
}

func TestMessagesCheckChatInviteValidatesHash(t *testing.T) {
	c := newChatInvitesCore(&chatInvitesFakeChatClient{}, 100)

	if _, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{}); err != tg.ErrInviteHashEmpty {
		t.Fatalf("empty hash error = %v, want %v", err, tg.ErrInviteHashEmpty)
	}
	if _, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{Hash: "bad"}); err != tg.ErrInviteHashInvalid {
		t.Fatalf("short hash error = %v, want %v", err, tg.ErrInviteHashInvalid)
	}
	if _, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{Hash: "bad bad bad bad bad!"}); err != tg.ErrInviteHashInvalid {
		t.Fatalf("malformed hash error = %v, want %v", err, tg.ErrInviteHashInvalid)
	}
}

func TestMessagesCheckChatInviteProjectsInvite(t *testing.T) {
	var gotUsersReq *userpb.TLUserGetUserProjectionBundle
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		checkChatInvite: func(context.Context, *chatpb.TLChatCheckChatInvite) (*chatpb.ChatInviteExt, error) {
			return chatpb.MakeTLChatInvite(&chatpb.TLChatInvite{
				RequestNeeded:     true,
				Title:             "chat",
				Photo:             tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{}),
				ParticipantsCount: 2,
				Participants:      []int64{200, 300, 200},
			}).ToChatInviteExt(), nil
		},
	}, &chatInvitesFakeUserClient{
		getProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			gotUsersReq = in
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
						tg.MakeTLUser(&tg.TLUser{Id: 200, FirstName: nonEmptyStringPtr("alice")}),
						tg.MakeTLUser(&tg.TLUser{Id: 300, FirstName: nonEmptyStringPtr("bob")}),
					}}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}, 100)

	r, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{Hash: "abcdefghijklmnopqrst"})
	if err != nil {
		t.Fatalf("MessagesCheckChatInvite error = %v", err)
	}
	invite, ok := r.ToChatInvite()
	if !ok || invite.Title != "chat" || !invite.RequestNeeded || len(invite.Participants) != 2 {
		t.Fatalf("MessagesCheckChatInvite = %#v, want chatInvite with two participants", r)
	}
	if gotUsersReq == nil || !sameInt64s(gotUsersReq.TargetUserIds, []int64{200, 300}) || !sameInt64s(gotUsersReq.ViewerUserIds, []int64{100}) {
		t.Fatalf("user request = %+v, want target=[200 300] viewer=[100]", gotUsersReq)
	}
}

func TestMessagesCheckChatInviteProjectsAlreadyAndPeek(t *testing.T) {
	tests := []struct {
		name string
		ext  *chatpb.ChatInviteExt
		want string
	}{
		{
			name: "already",
			ext: chatpb.MakeTLChatInviteAlready(&chatpb.TLChatInviteAlready{
				Chat: mutableChatWithPhotoForTest(42, 100, "chat"),
			}).ToChatInviteExt(),
			want: tg.ClazzName_chatInviteAlready,
		},
		{
			name: "peek",
			ext: chatpb.MakeTLChatInvitePeek(&chatpb.TLChatInvitePeek{
				Chat:    mutableChatWithPhotoForTest(42, 100, "chat"),
				Expires: 123,
			}).ToChatInviteExt(),
			want: tg.ClazzName_chatInvitePeek,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
				checkChatInvite: func(context.Context, *chatpb.TLChatCheckChatInvite) (*chatpb.ChatInviteExt, error) {
					return tt.ext, nil
				},
			}, &chatInvitesFakeUserClient{}, 100)

			r, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{Hash: "abcdefghijklmnopqrst"})
			if err != nil {
				t.Fatalf("MessagesCheckChatInvite error = %v", err)
			}
			if r.ClazzName() != tt.want {
				t.Fatalf("MessagesCheckChatInvite clazz = %s, want %s", r.ClazzName(), tt.want)
			}
			chat := projectedInviteChat(t, r)
			if _, ok := chat.Photo.(*tg.TLChatPhoto); !ok {
				t.Fatalf("projected invite chat photo = %T, want *tg.TLChatPhoto", chat.Photo)
			}
		})
	}
}

func TestMessagesCheckChatInviteMapsEmbeddedProjectionError(t *testing.T) {
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		checkChatInvite: func(context.Context, *chatpb.TLChatCheckChatInvite) (*chatpb.ChatInviteExt, error) {
			chat := mutableChatForTest(42, 100, "chat")
			chat.Chat.Date = int64(1) << 40
			return chatpb.MakeTLChatInviteAlready(&chatpb.TLChatInviteAlready{
				Chat: chat,
			}).ToChatInviteExt(), nil
		},
	}, &chatInvitesFakeUserClient{}, 100)

	_, err := c.MessagesCheckChatInvite(&tg.TLMessagesCheckChatInvite{Hash: "abcdefghijklmnopqrst"})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesCheckChatInvite error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesGetExportedChatInvitesReturnsInvitesAndAdminUsers(t *testing.T) {
	invite := exportedInviteForTest("link", 200)
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		getExportedChatInvites: func(context.Context, *chatpb.TLChatGetExportedChatInvites) (*chatpb.VectorExportedChatInvite, error) {
			return &chatpb.VectorExportedChatInvite{Datas: []tg.ExportedChatInviteClazz{invite}}, nil
		},
	}, usersForTest(200), 100)

	r, err := c.MessagesGetExportedChatInvites(&tg.TLMessagesGetExportedChatInvites{
		Peer:    inputPeerChatForTest(42),
		AdminId: tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
		Limit:   10,
	})
	if err != nil {
		t.Fatalf("MessagesGetExportedChatInvites error = %v", err)
	}
	if r.Count != 1 || len(r.Invites) != 1 || len(r.Users) != 1 {
		t.Fatalf("MessagesGetExportedChatInvites = %+v, want count/invite/user", r)
	}
}

func TestMessagesGetExportedChatInviteReturnsInviteAndUsers(t *testing.T) {
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		getExportedChatInvite: func(context.Context, *chatpb.TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error) {
			return exportedInviteForTest("link", 200).ToExportedChatInvite(), nil
		},
	}, usersForTest(100, 200), 100)

	r, err := c.MessagesGetExportedChatInvite(&tg.TLMessagesGetExportedChatInvite{
		Peer: inputPeerChatForTest(42),
		Link: "link",
	})
	if err != nil {
		t.Fatalf("MessagesGetExportedChatInvite error = %v", err)
	}
	normal, ok := r.ToMessagesExportedChatInvite()
	if !ok || normal.Invite == nil || len(normal.Users) != 2 {
		t.Fatalf("MessagesGetExportedChatInvite = %#v, want invite with self/admin users", r)
	}
}

func TestMessagesEditExportedChatInviteProjectsNormalReplacedAndInvalidLen(t *testing.T) {
	oldInvite := exportedInviteForTest("old", 200)
	newInvite := exportedInviteForTest("new", 200)

	tests := []struct {
		name  string
		data  []tg.ExportedChatInviteClazz
		want  string
		error error
	}{
		{name: "normal", data: []tg.ExportedChatInviteClazz{oldInvite}, want: tg.ClazzName_messages_exportedChatInvite},
		{name: "replaced", data: []tg.ExportedChatInviteClazz{oldInvite, newInvite}, want: tg.ClazzName_messages_exportedChatInviteReplaced},
		{name: "invalid", data: nil, error: tg.ErrInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
				editExportedChatInvite: func(context.Context, *chatpb.TLChatEditExportedChatInvite) (*chatpb.VectorExportedChatInvite, error) {
					return &chatpb.VectorExportedChatInvite{Datas: tt.data}, nil
				},
			}, usersForTest(100, 200), 100)

			r, err := c.MessagesEditExportedChatInvite(&tg.TLMessagesEditExportedChatInvite{
				Peer: inputPeerChatForTest(42),
				Link: "old",
			})
			if err != tt.error {
				t.Fatalf("MessagesEditExportedChatInvite error = %v, want %v", err, tt.error)
			}
			if tt.error == nil && r.ClazzName() != tt.want {
				t.Fatalf("MessagesEditExportedChatInvite clazz = %s, want %s", r.ClazzName(), tt.want)
			}
		})
	}
}

func TestMessagesGetAdminsWithInvitesReturnsAdminsAndUsers(t *testing.T) {
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		getAdminsWithInvites: func(context.Context, *chatpb.TLChatGetAdminsWithInvites) (*chatpb.VectorChatAdminWithInvites, error) {
			return &chatpb.VectorChatAdminWithInvites{Datas: []tg.ChatAdminWithInvitesClazz{
				tg.MakeTLChatAdminWithInvites(&tg.TLChatAdminWithInvites{AdminId: 200, InvitesCount: 1}),
			}}, nil
		},
	}, usersForTest(200), 100)

	r, err := c.MessagesGetAdminsWithInvites(&tg.TLMessagesGetAdminsWithInvites{Peer: inputPeerChatForTest(42)})
	if err != nil {
		t.Fatalf("MessagesGetAdminsWithInvites error = %v", err)
	}
	if len(r.Admins) != 1 || len(r.Users) != 1 {
		t.Fatalf("MessagesGetAdminsWithInvites = %+v, want admin/user", r)
	}
}

func TestMessagesGetChatInviteImportersReturnsImportersAndUsersWithDefaultLimit(t *testing.T) {
	var got *chatpb.TLChatGetChatInviteImporters
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		getChatInviteImporters: func(_ context.Context, in *chatpb.TLChatGetChatInviteImporters) (*chatpb.VectorChatInviteImporter, error) {
			got = in
			return &chatpb.VectorChatInviteImporter{Datas: []tg.ChatInviteImporterClazz{
				tg.MakeTLChatInviteImporter(&tg.TLChatInviteImporter{UserId: 200, Date: 1}),
			}}, nil
		},
	}, usersForTest(200), 100)

	r, err := c.MessagesGetChatInviteImporters(&tg.TLMessagesGetChatInviteImporters{
		Peer:       inputPeerChatForTest(42),
		OffsetUser: tg.MakeTLInputUserEmpty(&tg.TLInputUserEmpty{}),
	})
	if err != nil {
		t.Fatalf("MessagesGetChatInviteImporters error = %v", err)
	}
	if r.Count != 1 || len(r.Importers) != 1 || len(r.Users) != 1 {
		t.Fatalf("MessagesGetChatInviteImporters = %+v, want importer/user", r)
	}
	if got == nil || got.Limit != 50 {
		t.Fatalf("request = %+v, want default limit 50", got)
	}
}

func TestMessagesGetChatInviteImportersRejectsSearchWithLink(t *testing.T) {
	link := "link"
	q := "alice"
	c := newChatInvitesCore(&chatInvitesFakeChatClient{}, 100)

	_, err := c.MessagesGetChatInviteImporters(&tg.TLMessagesGetChatInviteImporters{
		Peer:       inputPeerChatForTest(42),
		Link:       &link,
		Q:          &q,
		OffsetUser: tg.MakeTLInputUserEmpty(&tg.TLInputUserEmpty{}),
	})
	if err != tg.ErrSearchWithLinkNotSupported {
		t.Fatalf("MessagesGetChatInviteImporters error = %v, want %v", err, tg.ErrSearchWithLinkNotSupported)
	}
}

func TestChatInviteReadHandlersRejectNonChatPeerAndMapChatError(t *testing.T) {
	c := newChatInvitesCore(&chatInvitesFakeChatClient{
		getAdminsWithInvites: func(context.Context, *chatpb.TLChatGetAdminsWithInvites) (*chatpb.VectorChatAdminWithInvites, error) {
			return nil, chatpb.ErrChatAdminRequired
		},
	}, 100)

	_, err := c.MessagesGetExportedChatInvites(&tg.TLMessagesGetExportedChatInvites{
		Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		AdminId: tg.MakeTLInputUserSelf(&tg.TLInputUserSelf{}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("non-chat peer error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}

	_, err = c.MessagesGetAdminsWithInvites(&tg.TLMessagesGetAdminsWithInvites{Peer: inputPeerChatForTest(42)})
	if err != tg.Err400ChatAdminRequired {
		t.Fatalf("chat error = %v, want %v", err, tg.Err400ChatAdminRequired)
	}
}

func TestChatInviteReadHandlersMapUserCompletionErrorToInternal(t *testing.T) {
	userErr := errors.New("user service unavailable")
	c := newChatInvitesCoreWithClients(&chatInvitesFakeChatClient{
		getAdminsWithInvites: func(context.Context, *chatpb.TLChatGetAdminsWithInvites) (*chatpb.VectorChatAdminWithInvites, error) {
			return &chatpb.VectorChatAdminWithInvites{Datas: []tg.ChatAdminWithInvitesClazz{
				tg.MakeTLChatAdminWithInvites(&tg.TLChatAdminWithInvites{AdminId: 200}),
			}}, nil
		},
	}, &chatInvitesFakeUserClient{
		getProjection: func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			return nil, userErr
		},
	}, 100)

	_, err := c.MessagesGetAdminsWithInvites(&tg.TLMessagesGetAdminsWithInvites{Peer: inputPeerChatForTest(42)})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesGetAdminsWithInvites error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func inputPeerChatForTest(chatID int64) tg.InputPeerClazz {
	return tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: chatID})
}

func usersForTest(ids ...int64) userclient.UserClient {
	return &chatInvitesFakeUserClient{
		getProjection: func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
			users := make([]tg.UserClazz, 0, len(ids))
			for _, id := range ids {
				users = append(users, tg.MakeTLUser(&tg.TLUser{Id: id, FirstName: nonEmptyStringPtr("user")}))
			}
			return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: users}),
				},
			}).ToUserProjectionBundle(), nil
		},
	}
}

func immutableUserForTest(id int64, firstName string) tg.ImmutableUserClazz {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{
			Id:         id,
			AccessHash: id * 10,
			FirstName:  firstName,
		}),
	}).ToImmutableUser()
}

func exportedInviteForTest(link string, adminID int64) *tg.TLChatInviteExported {
	return tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{
		Link:    link,
		AdminId: adminID,
		Date:    1,
	})
}

func mutableChatForTest(chatID, creatorID int64, title string) *tg.MutableChat {
	return tg.MakeTLMutableChat(&tg.TLMutableChat{
		Chat: tg.MakeTLImmutableChat(&tg.TLImmutableChat{
			Id:                chatID,
			Creator:           creatorID,
			Title:             title,
			ParticipantsCount: 1,
		}),
	}).ToMutableChat()
}

func mutableChatWithPhotoForTest(chatID, creatorID int64, title string) *tg.MutableChat {
	chat := mutableChatForTest(chatID, creatorID, title)
	chat.Chat.Photo = tg.MakeTLPhoto(&tg.TLPhoto{Id: 7001, DcId: 4})
	return chat
}

func projectedInviteChat(t *testing.T, invite *tg.ChatInvite) *tg.TLChat {
	t.Helper()
	if already, ok := invite.ToChatInviteAlready(); ok {
		chat, ok := already.Chat.(*tg.TLChat)
		if !ok {
			t.Fatalf("already chat = %T, want *tg.TLChat", already.Chat)
		}
		return chat
	}
	if peek, ok := invite.ToChatInvitePeek(); ok {
		chat, ok := peek.Chat.(*tg.TLChat)
		if !ok {
			t.Fatalf("peek chat = %T, want *tg.TLChat", peek.Chat)
		}
		return chat
	}
	t.Fatalf("invite = %s, want already or peek", invite.ClazzName())
	return nil
}

func sameInt64s(got, want []int64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
