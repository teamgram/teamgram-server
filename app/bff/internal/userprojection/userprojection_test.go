package userprojection

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUserClient struct {
	in  *userpb.TLUserGetUserProjectionBundle
	out *userpb.UserProjectionBundle
	err error
}

func (f *fakeUserClient) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	f.in = in
	return f.out, f.err
}

func TestProjectUsersReturnsViewerVector(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1002, FirstName: strPtr("Target")}),
			}}),
		},
	}).ToUserProjectionBundle()}

	got, err := ProjectUsers(context.Background(), client, 1001, []int64{1002}, MissingExplicitInput)
	if err != nil {
		t.Fatalf("ProjectUsers() error = %v", err)
	}
	if client.in == nil || len(client.in.ViewerUserIds) != 1 || client.in.ViewerUserIds[0] != 1001 || len(client.in.TargetUserIds) != 1 {
		t.Fatalf("request = %#v", client.in)
	}
	if len(got) != 1 {
		t.Fatalf("users = %#v", got)
	}
}

func TestProjectUsersMapsExplicitMissingToUserIdInvalid(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		MissingUserIds: []int64{1002},
	}).ToUserProjectionBundle()}

	_, err := ProjectUsers(context.Background(), client, 1001, []int64{1002}, MissingExplicitInput)
	if !errors.Is(err, tg.ErrUserIdInvalid) {
		t.Fatalf("error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}

func TestFillUpdatesUsersReplacesLegacyUsers(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	updates := tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{
				FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
				PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
			}),
		})},
		Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 9999})},
	})

	err := FillUpdatesUsers(context.Background(), client, 1001, updates.ToUpdates(), MissingStoredReference)
	if err != nil {
		t.Fatalf("FillUpdatesUsers() error = %v", err)
	}
	if len(updates.Users) != 2 || updates.Users[0].(*tg.TLUser).Id != 1001 || updates.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("users = %#v", updates.Users)
	}
}

func strPtr(v string) *string { return &v }
