package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/svc"
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

func TestMessagesClearAllDraftsFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesClearAllDrafts(&tg.TLMessagesClearAllDrafts{})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesClearAllDrafts error = %v, want %v", err, tg.ErrInternalServerError)
	}
}
