package userprojection

import (
	"context"
	"errors"
	"reflect"
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

var _ MissingPolicy = MissingExplicitInput

func TestProjectUsersEmptyInputReturnsNil(t *testing.T) {
	client := &fakeUserClient{}

	got, err := ProjectUsers(context.Background(), client, 1001, nil, Options{Missing: MissingExplicitInput})
	if err != nil {
		t.Fatalf("ProjectUsers(empty) error = %v", err)
	}
	if got != nil {
		t.Fatalf("ProjectUsers(empty) = %#v, want nil", got)
	}
	if client.in != nil {
		t.Fatalf("client called for empty input: %#v", client.in)
	}
}

func TestProjectUsersReturnsViewerVector(t *testing.T) {
	projected := tg.MakeTLUser(&tg.TLUser{Id: 1002, FirstName: strPtr("Target")})
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{
				ViewerUserId: 1001,
				Users:        []tg.UserClazz{projected},
			}),
		},
	}).ToUserProjectionBundle()}

	got, err := ProjectUsers(context.Background(), client, 1001, []int64{1002}, Options{Missing: MissingExplicitInput})
	if err != nil {
		t.Fatalf("ProjectUsers() error = %v", err)
	}
	if client.in == nil {
		t.Fatal("projection request was not sent")
	}
	if !reflect.DeepEqual(client.in.ViewerUserIds, []int64{1001}) {
		t.Fatalf("ViewerUserIds = %#v, want [1001]", client.in.ViewerUserIds)
	}
	if !reflect.DeepEqual(client.in.TargetUserIds, []int64{1002}) {
		t.Fatalf("TargetUserIds = %#v, want [1002]", client.in.TargetUserIds)
	}
	if !reflect.DeepEqual(got, []tg.UserClazz{projected}) {
		t.Fatalf("ProjectUsers() = %#v, want projected user", got)
	}
}

func TestProjectUsersPreservesDuplicateTargetsForOwnerRPC(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001}),
		},
	}).ToUserProjectionBundle()}

	_, err := ProjectUsers(context.Background(), client, 1001, []int64{1002, 1002}, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("ProjectUsers() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetUserIds, []int64{1002, 1002}) {
		t.Fatalf("TargetUserIds = %#v, want duplicate targets preserved", client.in.TargetUserIds)
	}
}

func TestProjectUsersMapsErrors(t *testing.T) {
	tests := []struct {
		name    string
		client  Client
		opts    Options
		wantErr error
	}{
		{
			name:    "nil client",
			client:  nil,
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrClientNotConfigured,
		},
		{
			name: "invalid request",
			client: &fakeUserClient{
				err: userpb.ErrUserInvalidArgument,
			},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrInvalidRequest,
		},
		{
			name:    "nil bundle",
			client:  &fakeUserClient{},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrNilBundle,
		},
		{
			name: "explicit missing target",
			client: &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				MissingUserIds: []int64{1002},
			}).ToUserProjectionBundle()},
			opts:    Options{Missing: MissingExplicitInput},
			wantErr: ErrExplicitUserMissing,
		},
		{
			name: "missing viewer",
			client: &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 9999}),
				},
			}).ToUserProjectionBundle()},
			opts:    Options{Missing: MissingStoredReference},
			wantErr: ErrViewerProjectionMissing,
		},
		{
			name: "empty viewer with require non empty",
			client: &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
				ViewerUsers: []userpb.ViewerUsersClazz{
					userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001}),
				},
			}).ToUserProjectionBundle()},
			opts:    Options{Missing: MissingStoredReference, RequireNonEmpty: true},
			wantErr: ErrViewerProjectionEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProjectUsers(context.Background(), tt.client, 1001, []int64{1002}, tt.opts)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ProjectUsers() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectUsersStoredReferenceMissingReturnsProjectedSubset(t *testing.T) {
	projected := tg.MakeTLUser(&tg.TLUser{Id: 1001})
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		MissingUserIds: []int64{1002},
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{
				ViewerUserId: 1001,
				Users:        []tg.UserClazz{projected},
			}),
		},
	}).ToUserProjectionBundle()}

	got, err := ProjectUsers(context.Background(), client, 1001, []int64{1001, 1002}, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("ProjectUsers() error = %v", err)
	}
	if !reflect.DeepEqual(got, []tg.UserClazz{projected}) {
		t.Fatalf("ProjectUsers() = %#v, want projected subset", got)
	}
}

func TestFillUpdatesUsersReplacesLegacyUsersAndCollectsContactMediaUser(t *testing.T) {
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
		Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 9999})},
	})

	err := FillUpdatesUsers(context.Background(), client, 1001, updates.ToUpdates(), Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("FillUpdatesUsers() error = %v", err)
	}
	if !reflect.DeepEqual(client.in.TargetUserIds, []int64{1001, 1002, 1003}) {
		t.Fatalf("TargetUserIds = %#v, want message and contact media users", client.in.TargetUserIds)
	}
	if len(updates.Users) != 3 || updates.Users[0].(*tg.TLUser).Id != 1001 || updates.Users[2].(*tg.TLUser).Id != 1003 {
		t.Fatalf("updates.Users = %#v, want projected users", updates.Users)
	}
}

func TestFillDifferenceUsersReplacesLegacyUsers(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		ViewerUsers: []userpb.ViewerUsersClazz{
			userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{Id: 1001}),
				tg.MakeTLUser(&tg.TLUser{Id: 1002}),
			}}),
		},
	}).ToUserProjectionBundle()}
	diff := tg.MakeTLUpdatesDifference(&tg.TLUpdatesDifference{
		NewMessages: []tg.MessageClazz{tg.MakeTLMessage(&tg.TLMessage{
			FromId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}),
			PeerId: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1002}),
		})},
		Users: []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 9999})},
	}).ToUpdatesDifference()

	err := FillDifferenceUsers(context.Background(), client, 1001, diff, Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("FillDifferenceUsers() error = %v", err)
	}
	full, ok := diff.ToUpdatesDifference()
	if !ok {
		t.Fatalf("diff = %s, want updates.difference", diff.ClazzName())
	}
	if len(full.Users) != 2 || full.Users[0].(*tg.TLUser).Id != 1001 || full.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("diff.Users = %#v, want projected users", full.Users)
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

	err := FillMessagesMessagesUsers(context.Background(), client, 1001, messages.ToMessagesMessages(), Options{Missing: MissingStoredReference})
	if err != nil {
		t.Fatalf("FillMessagesMessagesUsers() error = %v", err)
	}
	if len(messages.Users) != 2 || messages.Users[0].(*tg.TLUser).Id != 1001 || messages.Users[1].(*tg.TLUser).Id != 1002 {
		t.Fatalf("messages.Users = %#v, want projected users", messages.Users)
	}
}

func strPtr(v string) *string { return &v }
