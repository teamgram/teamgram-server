// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package repository

import (
	"context"
	"errors"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestCheckAccountUsername_InvalidFormat(t *testing.T) {
	repo := &Repository{
		UserClient: &stubUserClient{},
	}

	tests := []struct {
		name     string
		username string
	}{
		{"too short", "ab"},
		{"starts with number", "1abcde"},
		{"invalid chars", "ab@cde"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.CheckAccountUsername(context.Background(), 1, tt.username)
			if !errors.Is(err, ErrUsernameInvalid) {
				t.Errorf("expected ErrUsernameInvalid, got %v", err)
			}
		})
	}
}

func TestCheckAccountUsername_Occupied(t *testing.T) {
	repo := &Repository{
		UserClient: &stubUserClient{
			checkAccountUsernameFn: func(ctx context.Context, in *userpb.TLUserCheckAccountUsername) (*userpb.UsernameExist, error) {
				return userpb.MakeTLUsernameExistedNotMe(&userpb.TLUsernameExistedNotMe{}).ToUsernameExist(), nil
			},
		},
	}

	result, err := repo.CheckAccountUsername(context.Background(), 1, "validname")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tg.FromBool(result) {
		t.Error("expected false for occupied username")
	}
}

func TestCheckAccountUsername_Available(t *testing.T) {
	repo := &Repository{
		UserClient: &stubUserClient{
			checkAccountUsernameFn: func(ctx context.Context, in *userpb.TLUserCheckAccountUsername) (*userpb.UsernameExist, error) {
				// Return username does not exist (available)
				return userpb.MakeTLUsernameNotExisted(&userpb.TLUsernameNotExisted{}).ToUsernameExist(), nil
			},
		},
	}

	result, err := repo.CheckAccountUsername(context.Background(), 1, "validname")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tg.FromBool(result) {
		t.Error("expected true for available username")
	}
}

func TestCheckAccountUsername_AvailableIsMe(t *testing.T) {
	repo := &Repository{
		UserClient: &stubUserClient{
			checkAccountUsernameFn: func(ctx context.Context, in *userpb.TLUserCheckAccountUsername) (*userpb.UsernameExist, error) {
				// Username exists but belongs to me
				return userpb.MakeTLUsernameExistedIsMe(&userpb.TLUsernameExistedIsMe{}).ToUsernameExist(), nil
			},
		},
	}

	result, err := repo.CheckAccountUsername(context.Background(), 1, "validname")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tg.FromBool(result) {
		t.Error("expected true for username belonging to self")
	}
}

// stubUserClient embeds the userclient.UserClient interface and overrides
// only the methods needed for testing.
type stubUserClient struct {
	userclient.UserClient
	checkAccountUsernameFn func(ctx context.Context, in *userpb.TLUserCheckAccountUsername) (*userpb.UsernameExist, error)
	getImmutableUserFn     func(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error)
	resolveUsernameFn      func(ctx context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error)
	projectionFn           func(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
}

func (s *stubUserClient) UserCheckAccountUsername(ctx context.Context, in *userpb.TLUserCheckAccountUsername) (*userpb.UsernameExist, error) {
	if s.checkAccountUsernameFn != nil {
		return s.checkAccountUsernameFn(ctx, in)
	}
	return userpb.MakeTLUsernameExistedNotMe(&userpb.TLUsernameExistedNotMe{}).ToUsernameExist(), nil
}

func (s *stubUserClient) UserGetImmutableUser(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	if s.getImmutableUserFn != nil {
		return s.getImmutableUserFn(ctx, in)
	}
	return nil, nil
}

func (s *stubUserClient) UserResolveUsername(ctx context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error) {
	if s.resolveUsernameFn != nil {
		return s.resolveUsernameFn(ctx, in)
	}
	return nil, nil
}

func (s *stubUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	if s.projectionFn != nil {
		return s.projectionFn(ctx, in)
	}
	return nil, nil
}

type stubChatClient struct {
	chatclient.ChatClient

	projectionFn func(ctx context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error)
}

func (s *stubChatClient) ChatGetChatProjectionBundle(ctx context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
	if s.projectionFn != nil {
		return s.projectionFn(ctx, in)
	}
	return nil, nil
}

func TestResolveUsernameProjectsUser(t *testing.T) {
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	repo := &Repository{
		UserClient: &stubUserClient{
			resolveUsernameFn: func(_ context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error) {
				if in.Username != "target" {
					t.Fatalf("resolve username = %q, want target", in.Username)
				}
				return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 2002}).ToPeer(), nil
			},
			projectionFn: func(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1001, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 1001, Self: true}),
							tg.MakeTLUser(&tg.TLUser{Id: 2002, Contact: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}

	got, err := repo.ResolveUsername(context.Background(), 1001, "target")
	if err != nil {
		t.Fatalf("ResolveUsername error = %v", err)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1001 ||
		len(gotProjection.TargetUserIds) != 2 || gotProjection.TargetUserIds[0] != 1001 || gotProjection.TargetUserIds[1] != 2002 {
		t.Fatalf("projection request = %+v, want viewer [1001] target [1001 2002]", gotProjection)
	}
	if len(got.Users) != 2 {
		t.Fatalf("users = %#v", got.Users)
	}
	self, ok := got.Users[0].(*tg.TLUser)
	if !ok || self.Id != 1001 || !self.Self {
		t.Fatalf("self user = %#v", got.Users[0])
	}
	target, ok := got.Users[1].(*tg.TLUser)
	if !ok || target.Id != 2002 || !target.Contact {
		t.Fatalf("target user = %#v", got.Users[1])
	}
}

func TestResolveUsernameProjectsChatWithPhoto(t *testing.T) {
	var gotProjection *chatpb.TLChatGetChatProjectionBundle
	repo := &Repository{
		UserClient: &stubUserClient{
			resolveUsernameFn: func(_ context.Context, in *userpb.TLUserResolveUsername) (*tg.Peer, error) {
				if in.Username != "teamchat" {
					t.Fatalf("resolve username = %q, want teamchat", in.Username)
				}
				return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: 3001}).ToPeer(), nil
			},
		},
		ChatClient: &stubChatClient{
			projectionFn: func(_ context.Context, in *chatpb.TLChatGetChatProjectionBundle) (*chatpb.ChatProjectionBundle, error) {
				gotProjection = in
				return chatpb.MakeTLChatProjectionBundle(&chatpb.TLChatProjectionBundle{
					ViewerChats: []chatpb.ViewerChatsClazz{
						chatpb.MakeTLViewerChats(&chatpb.TLViewerChats{ViewerUserId: 1001, Chats: []tg.ChatClazz{
							tg.MakeTLChat(&tg.TLChat{
								Id:    3001,
								Title: "chat",
								Photo: tg.MakeTLChatPhoto(&tg.TLChatPhoto{PhotoId: 7001, DcId: 4}),
							}),
						}}),
					},
				}).ToChatProjectionBundle(), nil
			},
		},
	}

	got, err := repo.ResolveUsername(context.Background(), 1001, "teamchat")
	if err != nil {
		t.Fatalf("ResolveUsername error = %v", err)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1001 ||
		len(gotProjection.TargetChatIds) != 1 || gotProjection.TargetChatIds[0] != 3001 {
		t.Fatalf("chat projection request = %+v, want viewer [1001] target [3001]", gotProjection)
	}
	if len(got.Chats) != 1 {
		t.Fatalf("chats = %#v, want one chat", got.Chats)
	}
	chat, ok := got.Chats[0].(*tg.TLChat)
	if !ok {
		t.Fatalf("resolved chat = %T, want *tg.TLChat", got.Chats[0])
	}
	if _, ok := chat.Photo.(*tg.TLChatPhoto); !ok {
		t.Fatalf("resolved chat photo = %T, want *tg.TLChatPhoto", chat.Photo)
	}
}

func TestResolveUsernameMapsUserNotFoundToNotOccupied(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "direct sentinel",
			err:  userpb.ErrUsernameNotFound,
		},
		{
			name: "kitex remote biz error string",
			err:  errors.New("remote or network error: biz error: user: username not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				UserClient: &stubUserClient{
					resolveUsernameFn: func(_ context.Context, _ *userpb.TLUserResolveUsername) (*tg.Peer, error) {
						return nil, tt.err
					},
				},
			}

			_, err := repo.ResolveUsername(context.Background(), 1001, "missing")
			if !errors.Is(err, ErrUsernameNotOccupied) {
				t.Fatalf("ResolveUsername error = %v, want ErrUsernameNotOccupied", err)
			}
		})
	}
}

func TestUpdateAccountUsername_NoChange(t *testing.T) {
	var gotProjection *userpb.TLUserGetUserProjectionBundle
	repo := &Repository{
		UserClient: &stubUserClient{
			getImmutableUserFn: func(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
				return &tg.ImmutableUser{
					User: &tg.TLUserData{
						Id:        1,
						Username:  "already_set",
						FirstName: "Test",
					},
				}, nil
			},
			projectionFn: func(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
				gotProjection = in
				return userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
					ViewerUsers: []userpb.ViewerUsersClazz{
						userpb.MakeTLViewerUsers(&userpb.TLViewerUsers{ViewerUserId: 1, Users: []tg.UserClazz{
							tg.MakeTLUser(&tg.TLUser{Id: 1, Self: true}),
						}}),
					},
				}).ToUserProjectionBundle(), nil
			},
		},
	}
	// newUsername equals oldUsername; no RPC should be called.
	user, err := repo.UpdateAccountUsername(context.Background(), 1, "already_set")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotProjection == nil || len(gotProjection.ViewerUserIds) != 1 || gotProjection.ViewerUserIds[0] != 1 ||
		len(gotProjection.TargetUserIds) != 1 || gotProjection.TargetUserIds[0] != 1 {
		t.Fatalf("projection request = %+v, want viewer/target 1", gotProjection)
	}
	if user == nil {
		t.Fatal("expected non-nil user for no-change update")
	}
}

func TestUpdateAccountUsername_InvalidFormat(t *testing.T) {
	repo := &Repository{
		UserClient: &stubUserClient{
			getImmutableUserFn: func(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
				return &tg.ImmutableUser{
					User: &tg.TLUserData{
						Id:        1,
						Username:  "original",
						FirstName: "Test",
					},
				}, nil
			},
		},
	}

	tests := []struct {
		name     string
		username string
	}{
		{"too short", "ab"},
		{"starts with number", "1abcde"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.UpdateAccountUsername(context.Background(), 1, tt.username)
			if !errors.Is(err, ErrUsernameInvalid) {
				t.Errorf("expected ErrUsernameInvalid, got %v", err)
			}
		})
	}
}
