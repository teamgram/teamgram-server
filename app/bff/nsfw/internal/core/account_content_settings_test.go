package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/nsfw/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/nsfw/internal/svc"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeContentSettingsUserClient struct {
	userclient.UserClient
	getReq *userpb.TLUserGetContentSettings
	getRes *tg.AccountContentSettings
	getErr error
	setReq *userpb.TLUserSetContentSettings
	setRes *tg.Bool
	setErr error
}

func (f *fakeContentSettingsUserClient) UserGetContentSettings(ctx context.Context, in *userpb.TLUserGetContentSettings) (*tg.AccountContentSettings, error) {
	f.getReq = in
	return f.getRes, f.getErr
}

func (f *fakeContentSettingsUserClient) UserSetContentSettings(ctx context.Context, in *userpb.TLUserSetContentSettings) (*tg.Bool, error) {
	f.setReq = in
	return f.setRes, f.setErr
}

func newTestNsfwCore(userID int64, client userclient.UserClient) *NsfwCore {
	return &NsfwCore{
		ctx: context.Background(),
		svcCtx: &svc.ServiceContext{
			Repo: &repository.Repository{
				UserClient: client,
			},
		},
		MD: &metadata.RpcMetadata{
			UserId: userID,
		},
	}
}

func TestAccountGetContentSettings(t *testing.T) {
	expected := tg.MakeTLAccountContentSettings(&tg.TLAccountContentSettings{
		SensitiveEnabled:   true,
		SensitiveCanChange: true,
	}).ToAccountContentSettings()
	client := &fakeContentSettingsUserClient{getRes: expected}
	core := newTestNsfwCore(12345, client)

	got, err := core.AccountGetContentSettings(&tg.TLAccountGetContentSettings{})
	if err != nil {
		t.Fatalf("AccountGetContentSettings() error = %v", err)
	}
	if got != expected {
		t.Fatalf("AccountGetContentSettings() = %p, want %p", got, expected)
	}
	if client.getReq == nil || client.getReq.UserId != 12345 {
		t.Fatalf("UserGetContentSettings request = %#v, want user_id 12345", client.getReq)
	}
}

func TestAccountGetContentSettingsReturnsDownstreamError(t *testing.T) {
	downstreamErr := errors.New("downstream")
	client := &fakeContentSettingsUserClient{getErr: downstreamErr}
	core := newTestNsfwCore(12345, client)

	got, err := core.AccountGetContentSettings(&tg.TLAccountGetContentSettings{})
	if got != nil {
		t.Fatalf("AccountGetContentSettings() result = %#v, want nil", got)
	}
	if !errors.Is(err, downstreamErr) {
		t.Fatalf("AccountGetContentSettings() error = %v, want %v", err, downstreamErr)
	}
}

func TestAccountGetContentSettingsRequiresSelfID(t *testing.T) {
	core := newTestNsfwCore(0, &fakeContentSettingsUserClient{})

	got, err := core.AccountGetContentSettings(&tg.TLAccountGetContentSettings{})
	if got != nil {
		t.Fatalf("AccountGetContentSettings() result = %#v, want nil", got)
	}
	if err != tg.ErrUserIdInvalid {
		t.Fatalf("AccountGetContentSettings() error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}

func TestAccountGetContentSettingsRequiresUserClient(t *testing.T) {
	core := newTestNsfwCore(12345, nil)

	got, err := core.AccountGetContentSettings(&tg.TLAccountGetContentSettings{})
	if got != nil {
		t.Fatalf("AccountGetContentSettings() result = %#v, want nil", got)
	}
	if err == nil || err.Error() != "nsfw: user client is nil" {
		t.Fatalf("AccountGetContentSettings() error = %v, want nsfw user client error", err)
	}
}

func TestAccountSetContentSettings(t *testing.T) {
	expected := tg.BoolTrue
	client := &fakeContentSettingsUserClient{setRes: expected}
	core := newTestNsfwCore(12345, client)

	got, err := core.AccountSetContentSettings(&tg.TLAccountSetContentSettings{
		SensitiveEnabled: true,
	})
	if err != nil {
		t.Fatalf("AccountSetContentSettings() error = %v", err)
	}
	if got != expected {
		t.Fatalf("AccountSetContentSettings() = %p, want %p", got, expected)
	}
	if client.setReq == nil || client.setReq.UserId != 12345 || !client.setReq.SensitiveEnabled {
		t.Fatalf("UserSetContentSettings request = %#v, want user_id 12345 and sensitive_enabled true", client.setReq)
	}
}

func TestAccountSetContentSettingsRejectsNilInput(t *testing.T) {
	core := newTestNsfwCore(12345, &fakeContentSettingsUserClient{})

	got, err := core.AccountSetContentSettings(nil)
	if got != nil {
		t.Fatalf("AccountSetContentSettings() result = %#v, want nil", got)
	}
	if err != tg.ErrInputRequestInvalid {
		t.Fatalf("AccountSetContentSettings() error = %v, want %v", err, tg.ErrInputRequestInvalid)
	}
}
