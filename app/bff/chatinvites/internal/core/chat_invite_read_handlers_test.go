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
}

func (f *chatInvitesFakeUserClient) UserGetMutableUsers(ctx context.Context, in *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error) {
	return f.getMutableUsers(ctx, in)
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
	var gotUsersReq *userpb.TLUserGetMutableUsers
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
		getMutableUsers: func(_ context.Context, in *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error) {
			gotUsersReq = in
			return &userpb.VectorImmutableUser{Datas: []tg.ImmutableUserClazz{
				immutableUserForTest(200, "alice"),
				immutableUserForTest(300, "bob"),
			}}, nil
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
	if gotUsersReq == nil || !sameInt64s(gotUsersReq.Id, []int64{200, 300}) || !sameInt64s(gotUsersReq.To, []int64{100}) {
		t.Fatalf("user request = %+v, want id=[200 300] to=[100]", gotUsersReq)
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
				Chat: mutableChatForTest(42, 100, "chat"),
			}).ToChatInviteExt(),
			want: tg.ClazzName_chatInviteAlready,
		},
		{
			name: "peek",
			ext: chatpb.MakeTLChatInvitePeek(&chatpb.TLChatInvitePeek{
				Chat:    mutableChatForTest(42, 100, "chat"),
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
		})
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
		getMutableUsers: func(context.Context, *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error) {
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
		getMutableUsers: func(context.Context, *userpb.TLUserGetMutableUsers) (*userpb.VectorImmutableUser, error) {
			users := make([]tg.ImmutableUserClazz, 0, len(ids))
			for _, id := range ids {
				users = append(users, immutableUserForTest(id, "user"))
			}
			return &userpb.VectorImmutableUser{Datas: users}, nil
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
