package userprojection

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	bizuserproj "github.com/teamgram/teamgram-server/v2/app/service/biz/user/userprojection"
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

func TestFillUpdatesUsersIncludesContactMediaUser(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
				tg.MakeTLUser(&tg.TLUser{Id: 1003}),
			}}),
		},
	}).ToUserProjectionBundle()}
	updates := tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: []tg.UpdateClazz{tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message: tg.MakeTLMessage(&tg.TLMessage{
				FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
				PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
				Media:  tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{UserId: 1003}),
			}),
		})},
	})

	err := FillUpdatesUsers(context.Background(), client, 1001, updates.ToUpdates(), MissingStoredReference)
	if err != nil {
		t.Fatalf("FillUpdatesUsers() error = %v", err)
	}
	if len(client.in.TargetUserIds) != 3 || client.in.TargetUserIds[2] != 1003 {
		t.Fatalf("target user ids = %v, want contact media user included", client.in.TargetUserIds)
	}
	if len(updates.Users) != 3 {
		t.Fatalf("users = %#v, want 3 projected users", updates.Users)
	}
}

func TestFillMessagesMessagesUsersReplacesLegacyUsers(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	messages := tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
		Messages: []tg.MessageClazz{tg.MakeTLMessage(&tg.TLMessage{
			FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
			PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
		})},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 9999})},
	})

	err := FillMessagesMessagesUsers(context.Background(), client, 1001, messages.ToMessagesMessages(), MissingStoredReference)
	if err != nil {
		t.Fatalf("FillMessagesMessagesUsers() error = %v", err)
	}
	if len(messages.Users) != 2 || messages.Users[0].(*tg.TLUser).Id != 1001 || messages.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("users = %#v", messages.Users)
	}
}

func TestMissingPolicyMapsToPublicProjectionPolicy(t *testing.T) {
	tests := []struct {
		name string
		in   MissingPolicy
		want bizuserproj.MissingPolicy
	}{
		{name: "explicit input", in: MissingExplicitInput, want: bizuserproj.MissingExplicitInput},
		{name: "stored reference", in: MissingStoredReference, want: bizuserproj.MissingStoredReference},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := publicMissingPolicy(tt.in)
			if got != tt.want {
				t.Fatalf("publicMissingPolicy(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestProjectUsersEmptyInputPreservesBFFEmptySlice(t *testing.T) {
	client := &fakeUserClient{}

	got, err := ProjectUsers(context.Background(), client, 1001, nil, MissingExplicitInput)
	if err != nil {
		t.Fatalf("ProjectUsers(empty) error = %v", err)
	}
	if got == nil {
		t.Fatalf("ProjectUsers(empty) = nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("ProjectUsers(empty) = %#v, want empty slice", got)
	}
	if client.in != nil {
		t.Fatalf("client called for empty input: %#v", client.in)
	}
}

func TestProjectUsersMissingViewerPreservesBFFEmptySlice(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 9999}),
		},
	}).ToUserProjectionBundle()}

	got, err := ProjectUsers(context.Background(), client, 1001, []int64{1002}, MissingStoredReference)
	if err != nil {
		t.Fatalf("ProjectUsers(missing viewer) error = %v", err)
	}
	if got == nil {
		t.Fatalf("ProjectUsers(missing viewer) = nil, want empty non-nil slice")
	}
	if len(got) != 0 {
		t.Fatalf("ProjectUsers(missing viewer) = %#v, want empty slice", got)
	}
}

func strPtr(v string) *string { return &v }
