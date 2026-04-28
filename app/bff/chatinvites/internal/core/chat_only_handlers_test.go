package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/chatinvites/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type chatInvitesFakeChatClient struct {
	chatclient.ChatClient

	exportInvite  func(context.Context, *chatpb.TLChatExportChatInvite) (*tg.ExportedChatInvite, error)
	deleteInvite  func(context.Context, *chatpb.TLChatDeleteExportedChatInvite) (*tg.Bool, error)
	deleteRevoked func(context.Context, *chatpb.TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error)

	getAdminsWithInvites   func(context.Context, *chatpb.TLChatGetAdminsWithInvites) (*chatpb.VectorChatAdminWithInvites, error)
	getExportedChatInvite  func(context.Context, *chatpb.TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error)
	getExportedChatInvites func(context.Context, *chatpb.TLChatGetExportedChatInvites) (*chatpb.VectorExportedChatInvite, error)
	checkChatInvite        func(context.Context, *chatpb.TLChatCheckChatInvite) (*chatpb.ChatInviteExt, error)
	getChatInviteImporters func(context.Context, *chatpb.TLChatGetChatInviteImporters) (*chatpb.VectorChatInviteImporter, error)
	editExportedChatInvite func(context.Context, *chatpb.TLChatEditExportedChatInvite) (*chatpb.VectorExportedChatInvite, error)
}

func (f *chatInvitesFakeChatClient) ChatExportChatInvite(ctx context.Context, in *chatpb.TLChatExportChatInvite) (*tg.ExportedChatInvite, error) {
	return f.exportInvite(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatDeleteExportedChatInvite(ctx context.Context, in *chatpb.TLChatDeleteExportedChatInvite) (*tg.Bool, error) {
	return f.deleteInvite(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatDeleteRevokedExportedChatInvites(ctx context.Context, in *chatpb.TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error) {
	return f.deleteRevoked(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatGetAdminsWithInvites(ctx context.Context, in *chatpb.TLChatGetAdminsWithInvites) (*chatpb.VectorChatAdminWithInvites, error) {
	return f.getAdminsWithInvites(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatGetExportedChatInvite(ctx context.Context, in *chatpb.TLChatGetExportedChatInvite) (*tg.ExportedChatInvite, error) {
	return f.getExportedChatInvite(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatGetExportedChatInvites(ctx context.Context, in *chatpb.TLChatGetExportedChatInvites) (*chatpb.VectorExportedChatInvite, error) {
	return f.getExportedChatInvites(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatCheckChatInvite(ctx context.Context, in *chatpb.TLChatCheckChatInvite) (*chatpb.ChatInviteExt, error) {
	return f.checkChatInvite(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatGetChatInviteImporters(ctx context.Context, in *chatpb.TLChatGetChatInviteImporters) (*chatpb.VectorChatInviteImporter, error) {
	return f.getChatInviteImporters(ctx, in)
}

func (f *chatInvitesFakeChatClient) ChatEditExportedChatInvite(ctx context.Context, in *chatpb.TLChatEditExportedChatInvite) (*chatpb.VectorExportedChatInvite, error) {
	return f.editExportedChatInvite(ctx, in)
}

func newChatInvitesCore(client chatclient.ChatClient, selfID int64) *ChatInvitesCore {
	return newChatInvitesCoreWithClients(client, nil, selfID)
}

func newChatInvitesCoreWithClients(client chatclient.ChatClient, userClient userclient.UserClient, selfID int64) *ChatInvitesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{ChatClient: client, UserClient: userClient},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID}
	return c
}

func TestMessagesExportChatInviteRejectsNonChatPeer(t *testing.T) {
	c := newChatInvitesCore(&chatInvitesFakeChatClient{}, 100)

	_, err := c.MessagesExportChatInvite(&tg.TLMessagesExportChatInvite{
		Peer: tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
	})
	if err != tg.Err400PeerIdInvalid {
		t.Fatalf("MessagesExportChatInvite error = %v, want %v", err, tg.Err400PeerIdInvalid)
	}
}

func TestMessagesExportChatInviteMapsChatError(t *testing.T) {
	c := newChatInvitesCore(&chatInvitesFakeChatClient{
		exportInvite: func(context.Context, *chatpb.TLChatExportChatInvite) (*tg.ExportedChatInvite, error) {
			return nil, chatpb.ErrChatAdminRequired
		},
	}, 100)

	_, err := c.MessagesExportChatInvite(&tg.TLMessagesExportChatInvite{
		Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
	})
	if err != tg.Err400ChatAdminRequired {
		t.Fatalf("MessagesExportChatInvite error = %v, want %v", err, tg.Err400ChatAdminRequired)
	}
}

func TestMessagesExportChatInviteMapsRequestFields(t *testing.T) {
	expireDate := int32(123)
	usageLimit := int32(7)
	title := "invite"
	var got *chatpb.TLChatExportChatInvite
	want := tg.MakeTLChatInviteExported(&tg.TLChatInviteExported{Link: "https://t.me/+hash"}).ToExportedChatInvite()
	c := newChatInvitesCore(&chatInvitesFakeChatClient{
		exportInvite: func(_ context.Context, in *chatpb.TLChatExportChatInvite) (*tg.ExportedChatInvite, error) {
			got = in
			return want, nil
		},
	}, 100)

	r, err := c.MessagesExportChatInvite(&tg.TLMessagesExportChatInvite{
		LegacyRevokePermanent: true,
		RequestNeeded:         true,
		Peer:                  tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		ExpireDate:            &expireDate,
		UsageLimit:            &usageLimit,
		Title:                 &title,
	})
	if err != nil {
		t.Fatalf("MessagesExportChatInvite error = %v", err)
	}
	if r != want {
		t.Fatalf("MessagesExportChatInvite = %v, want %v", r, want)
	}
	if got == nil || got.ChatId != 42 || got.AdminId != 100 || !got.LegacyRevokePermanent || !got.RequestNeeded ||
		got.ExpireDate != &expireDate || got.UsageLimit != &usageLimit || got.Title != &title {
		t.Fatalf("request = %+v, want mapped fields", got)
	}
}

func TestMessagesDeleteExportedChatInviteMapsRequestFields(t *testing.T) {
	var got *chatpb.TLChatDeleteExportedChatInvite
	c := newChatInvitesCore(&chatInvitesFakeChatClient{
		deleteInvite: func(_ context.Context, in *chatpb.TLChatDeleteExportedChatInvite) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, 100)

	r, err := c.MessagesDeleteExportedChatInvite(&tg.TLMessagesDeleteExportedChatInvite{
		Peer: tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		Link: "link",
	})
	if err != nil {
		t.Fatalf("MessagesDeleteExportedChatInvite error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("MessagesDeleteExportedChatInvite = %v, want BoolTrue", r)
	}
	if got == nil || got.SelfId != 100 || got.ChatId != 42 || got.Link != "link" {
		t.Fatalf("request = %+v, want self_id=100 chat_id=42 link=link", got)
	}
}

func TestMessagesDeleteRevokedExportedChatInvitesMapsRequestFields(t *testing.T) {
	var got *chatpb.TLChatDeleteRevokedExportedChatInvites
	c := newChatInvitesCore(&chatInvitesFakeChatClient{
		deleteRevoked: func(_ context.Context, in *chatpb.TLChatDeleteRevokedExportedChatInvites) (*tg.Bool, error) {
			got = in
			return tg.BoolTrue, nil
		},
	}, 100)

	r, err := c.MessagesDeleteRevokedExportedChatInvites(&tg.TLMessagesDeleteRevokedExportedChatInvites{
		Peer:    tg.MakeTLInputPeerChat(&tg.TLInputPeerChat{ChatId: 42}),
		AdminId: tg.MakeTLInputUser(&tg.TLInputUser{UserId: 200}),
	})
	if err != nil {
		t.Fatalf("MessagesDeleteRevokedExportedChatInvites error = %v", err)
	}
	if r != tg.BoolTrue {
		t.Fatalf("MessagesDeleteRevokedExportedChatInvites = %v, want BoolTrue", r)
	}
	if got == nil || got.SelfId != 100 || got.ChatId != 42 || got.AdminId != 200 {
		t.Fatalf("request = %+v, want self_id=100 chat_id=42 admin_id=200", got)
	}
}
