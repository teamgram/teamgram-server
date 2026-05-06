package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/svc"
	dialogclient "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/client"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
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
}

func TestMessagesGetAllDraftsFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesGetAllDrafts(&tg.TLMessagesGetAllDrafts{})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesGetAllDrafts error = %v, want %v", err, tg.ErrInternalServerError)
	}
}

func TestMessagesClearAllDraftsFailsWhenDialogClientIsNotConfigured(t *testing.T) {
	c := newDraftsCoreForTest(&repository.Repository{}, 100)

	_, err := c.MessagesClearAllDrafts(&tg.TLMessagesClearAllDrafts{})
	if err != tg.ErrInternalServerError {
		t.Fatalf("MessagesClearAllDrafts error = %v, want %v", err, tg.ErrInternalServerError)
	}
}
