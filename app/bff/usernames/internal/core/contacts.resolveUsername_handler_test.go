package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/internal/svc"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type resolveUsernameUserClient struct {
	userclient.UserClient
	resolveUsername func(ctx context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error)
}

func (c *resolveUsernameUserClient) UserResolveUsername(ctx context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error) {
	return c.resolveUsername(ctx, in)
}

func TestContactsResolveUsernameMapsMissingUsernameToTGError(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			UserClient: &resolveUsernameUserClient{
				resolveUsername: func(_ context.Context, _ *userpb.TLUserResolveUsername) (*tg.Peer, error) {
					return nil, userpb.ErrUsernameNotFound
				},
			},
		},
	})
	core.MD = &metadata.RpcMetadata{UserId: 1001}

	_, err := core.ContactsResolveUsername(&tg.TLContactsResolveUsername{Username: "missing"})
	if err != tg.ErrUsernameNotOccupied {
		t.Fatalf("ContactsResolveUsername error = %v, want USERNAME_NOT_OCCUPIED", err)
	}
}

func TestContactsResolveUsernameMapsRemoteMissingUsernameToTGError(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{
			UserClient: &resolveUsernameUserClient{
				resolveUsername: func(_ context.Context, _ *userpb.TLUserResolveUsername) (*tg.Peer, error) {
					return nil, errors.New("remote or network error: biz error: user: username not found")
				},
			},
		},
	})
	core.MD = &metadata.RpcMetadata{UserId: 1001}

	_, err := core.ContactsResolveUsername(&tg.TLContactsResolveUsername{Username: "missing"})
	if err != tg.ErrUsernameNotOccupied {
		t.Fatalf("ContactsResolveUsername error = %v, want USERNAME_NOT_OCCUPIED", err)
	}
}
