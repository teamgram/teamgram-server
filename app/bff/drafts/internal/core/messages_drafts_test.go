package core

import (
	"context"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDraftDialogClient struct {
	dialogclient.DialogClient
	save     func(context.Context, *dialogpb.TLDialogSaveDraftMessage) (*tg.Bool, error)
	getAll   func(context.Context, *dialogpb.TLDialogGetAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error)
	clearAll func(context.Context, *dialogpb.TLDialogClearAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error)
}

func (f fakeDraftDialogClient) DialogSaveDraftMessage(ctx context.Context, in *dialogpb.TLDialogSaveDraftMessage) (*tg.Bool, error) {
	return f.save(ctx, in)
}

func (f fakeDraftDialogClient) DialogGetAllDrafts(ctx context.Context, in *dialogpb.TLDialogGetAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error) {
	return f.getAll(ctx, in)
}

func (f fakeDraftDialogClient) DialogClearAllDrafts(ctx context.Context, in *dialogpb.TLDialogClearAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error) {
	return f.clearAll(ctx, in)
}

type fakeDraftUserClient struct {
	userclient.UserClient
	getUserProjection func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
}

func (f fakeDraftUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	return f.getUserProjection(ctx, in)
}

type fakeDraftChatClient struct {
	chatclient.ChatClient
	getChatProjection func(context.Context, *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error)
}

func (f fakeDraftChatClient) ChatGetChatProjectionBundle(ctx context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
	return f.getChatProjection(ctx, in)
}

func newDraftsCoreForTest(repo *repository.Repository, selfID int64) *DraftsCore {
	c := New(context.Background(), &svc.ServiceContext{Repo: repo})
	c.MD = &metadata.RpcMetadata{UserId: selfID, PermAuthKeyId: 9001}
	return c
}

func TestMessagesSaveDraftFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesSaveDraft(&tg.TLMessagesSaveDraft{
		Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		Message: "draft",
	})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesSaveDraft error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesSaveDraftCarriesSourceAuthAndOutbox(t *testing.T) {
	var got *dialogpb.TLDialogSaveDraftMessage
	c := newDraftsCoreForTest(&repository.Repository{
		DialogClient: fakeDraftDialogClient{
			save: func(_ context.Context, in *dialogpb.TLDialogSaveDraftMessage) (*tg.Bool, error) {
				got = in
				return tg.BoolTrue, nil
			},
		},
	}, 100)

	_, err := c.MessagesSaveDraft(&tg.TLMessagesSaveDraft{
		Peer:    tg.MakeTLInputPeerUser(&tg.TLInputPeerUser{UserId: 200}),
		Message: "draft",
	})
	if err != nil {
		t.Fatalf("MessagesSaveDraft error = %v", err)
	}
	if got == nil {
		t.Fatal("dialog.saveDraftMessage was not called")
	}
	if got.SourcePermAuthKeyId != 9001 || got.OperationId == "" || got.OutboxId == 0 {
		t.Fatalf("dialog save draft request = %+v", got)
	}
	if got.PeerType != 1 || got.PeerId != 200 {
		t.Fatalf("dialog save draft peer = (%d,%d), want (1,200)", got.PeerType, got.PeerId)
	}
}

func TestMessagesSaveDraftMapsSelfPeerToUserDialogPeer(t *testing.T) {
	var got *dialogpb.TLDialogSaveDraftMessage
	c := newDraftsCoreForTest(&repository.Repository{
		DialogClient: fakeDraftDialogClient{
			save: func(_ context.Context, in *dialogpb.TLDialogSaveDraftMessage) (*tg.Bool, error) {
				got = in
				return tg.BoolTrue, nil
			},
		},
	}, 100)

	_, err := c.MessagesSaveDraft(&tg.TLMessagesSaveDraft{
		Peer:    tg.MakeTLInputPeerSelf(&tg.TLInputPeerSelf{}),
		Message: "self draft",
	})
	if err != nil {
		t.Fatalf("MessagesSaveDraft error = %v", err)
	}
	if got == nil {
		t.Fatal("dialog.saveDraftMessage was not called")
	}
	if got.PeerType != 1 || got.PeerId != 100 {
		t.Fatalf("dialog save draft peer = (%d,%d), want (1,100)", got.PeerType, got.PeerId)
	}
}

func TestMessagesGetAllDraftsFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesGetAllDrafts(&tg.TLMessagesGetAllDrafts{})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesGetAllDrafts error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesGetAllDraftsProjectsUsers(t *testing.T) {
	var gotUsers *userpb.TLUserGetUserProjectionBundle
	c := newDraftsCoreForTest(&repository.Repository{
		DialogClient: fakeDraftDialogClient{
			getAll: func(context.Context, *dialogpb.TLDialogGetAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error) {
				return &dialogpb.VectorPeerWithDraftMessage{Datas: []dialogpb.PeerWithDraftMessageClazz{
					dialogpb.MakeTLUpdateDraftMessage(&dialogpb.TLUpdateDraftMessage{
						Peer:  tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 200}),
						Draft: tg.MakeTLDraftMessage(&tg.TLDraftMessage{Message: "draft"}),
					}),
				}}, nil
			},
		},
		UserClient: fakeDraftUserClient{
			getUserProjection: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotUsers = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 100, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 200, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}, 100)

	got, err := c.MessagesGetAllDrafts(&tg.TLMessagesGetAllDrafts{})
	if err != nil {
		t.Fatalf("MessagesGetAllDrafts error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("got %s, want updates", got.ClazzName())
	}
	if gotUsers == nil || len(gotUsers.ViewerUserIds) != 1 || gotUsers.ViewerUserIds[0] != 100 || len(gotUsers.TargetUserIds) != 1 || gotUsers.TargetUserIds[0] != 200 {
		t.Fatalf("projection request = %+v, want viewer [100] target [200]", gotUsers)
	}
	if len(updates.Users) != 1 {
		t.Fatalf("users = %#v", updates.Users)
	}
	user, ok := updates.Users[0].(*tg.TLUser)
	if !ok || user.Id != 200 || !user.Contact {
		t.Fatalf("projected user = %#v", updates.Users[0])
	}
}

func TestMessagesGetAllDraftsProjectsChats(t *testing.T) {
	var (
		gotChats *chatpb.TLChatGetChatProjectionBundle
		calls    int
	)
	c := newDraftsCoreForTest(&repository.Repository{
		DialogClient: fakeDraftDialogClient{
			getAll: func(context.Context, *dialogpb.TLDialogGetAllDrafts) (*dialogpb.VectorPeerWithDraftMessage, error) {
				return &dialogpb.VectorPeerWithDraftMessage{Datas: []dialogpb.PeerWithDraftMessageClazz{
					dialogpb.MakeTLUpdateDraftMessage(&dialogpb.TLUpdateDraftMessage{
						Peer:  tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 300}),
						Draft: tg.MakeTLDraftMessage(&tg.TLDraftMessage{Message: "first"}),
					}),
					dialogpb.MakeTLUpdateDraftMessage(&dialogpb.TLUpdateDraftMessage{
						Peer:  tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 301}),
						Draft: tg.MakeTLDraftMessage(&tg.TLDraftMessage{Message: "second"}),
					}),
				}}, nil
			},
		},
		ChatClient: fakeDraftChatClient{
			getChatProjection: func(_ context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
				calls++
				gotChats = in
				return chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
					ViewerChats: []chatpb.ViewerChatsClazz{
						chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{
							ViewerUserId: 100,
							Chats: []tg.ChatClazz{
								tg.MakeTLChat(&tg.TLChat{Id: 300, Title: "first chat"}),
								tg.MakeTLChat(&tg.TLChat{Id: 301, Title: "second chat"}),
							},
						}),
					},
				}).ToChatProjectionBundle(), nil
			},
		},
	}, 100)

	got, err := c.MessagesGetAllDrafts(&tg.TLMessagesGetAllDrafts{})
	if err != nil {
		t.Fatalf("MessagesGetAllDrafts error = %v", err)
	}
	updates, ok := got.ToUpdates()
	if !ok {
		t.Fatalf("got %s, want updates", got.ClazzName())
	}
	if calls != 1 {
		t.Fatalf("ChatGetChatProjectionBundle calls = %d, want 1", calls)
	}
	if gotChats == nil || !reflect.DeepEqual(gotChats.ViewerUserIds, []int64{100}) {
		t.Fatalf("ChatGetChatProjectionBundle viewer request = %+v, want [100]", gotChats)
	}
	if !reflect.DeepEqual(gotChats.TargetChatIds, []int64{300, 301}) {
		t.Fatalf("ChatGetChatProjectionBundle target ids = %v, want [300 301]", gotChats.TargetChatIds)
	}
	if len(updates.Chats) != 2 {
		t.Fatalf("chats = %#v", updates.Chats)
	}
	first, ok := updates.Chats[0].(*tg.TLChat)
	if !ok || first.Id != 300 || first.Title != "first chat" {
		t.Fatalf("first projected chat = %#v", updates.Chats[0])
	}
	second, ok := updates.Chats[1].(*tg.TLChat)
	if !ok || second.Id != 301 || second.Title != "second chat" {
		t.Fatalf("second projected chat = %#v", updates.Chats[1])
	}
}

func TestMessagesClearAllDraftsFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesClearAllDrafts(&tg.TLMessagesClearAllDrafts{})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesClearAllDrafts error = %v, want %v", err, tg.ErrInternalServerError)
	}
}
