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
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/strings2"

	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/usernames/plugin"
	chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	syncclient "github.com/teamgram/teamgram-server/v2/app/messenger/sync/client"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// MinUsernameLen is the minimum length for a username.
	MinUsernameLen = 5
)

// Repository is the dependency container for BFF usernames business logic.
type Repository struct {
	UserClient userclient.UserClient
	ChatClient chatclient.ChatClient
	SyncClient syncclient.SyncClient
	Plugin     plugin.UsernamesPlugin
}

// NewRepository creates a new Repository. Clients are created only
// when the corresponding config section is populated.
func NewRepository(c config.Config) *Repository {
	r := &Repository{}
	if hasRPCClientConfig(c.UserClient) {
		r.UserClient = userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient))
	}
	if hasRPCClientConfig(c.ChatClient) {
		r.ChatClient = chatclient.NewChatClient(chatclient.MustNewKitexClient(c.ChatClient))
	}
	if hasRPCClientConfig(c.SyncClient) {
		r.SyncClient = syncclient.NewSyncClient(syncclient.MustNewKitexClient(c.SyncClient))
	}
	return r
}

// SetPlugin sets the enterprise plugin.
func (r *Repository) SetPlugin(p plugin.UsernamesPlugin) {
	r.Plugin = p
}

// CheckAccountUsername validates the username format and checks availability.
func (r *Repository) CheckAccountUsername(ctx context.Context, userId int64, username string) (*tg.Bool, error) {
	if len(username) < MinUsernameLen ||
		!strings2.IsAlNumString(username) ||
		isFirstCharNumber(username) {
		return nil, ErrUsernameInvalid
	}

	existed, err := r.UserClient.UserCheckAccountUsername(ctx, &userpb.TLUserCheckAccountUsername{
		UserId:   userId,
		Username: username,
	})
	if err != nil {
		return nil, fmt.Errorf("usernames repository: check account username: %w", err)
	}

	if _, ok := existed.Clazz.(*userpb.TLUsernameExistedNotMe); ok {
		return tg.ToBool(false), nil
	}

	return tg.ToBool(true), nil
}

// UpdateAccountUsername validates and updates the username for a user account.
func (r *Repository) UpdateAccountUsername(ctx context.Context, userId int64, newUsername string) (*tg.User, error) {
	me, err := r.UserClient.UserGetImmutableUser(ctx, &userpb.TLUserGetImmutableUser{
		Id: userId,
	})
	if err != nil {
		return nil, fmt.Errorf("usernames repository: get immutable user: %w", err)
	}

	if me == nil || me.User == nil {
		return nil, fmt.Errorf("usernames repository: get immutable user: returned nil user data")
	}

	oldUsername := userDataUsername(me.User)
	if newUsername == oldUsername {
		return buildSelfUser(me), nil
	}

	if newUsername != "" {
		if len(newUsername) < MinUsernameLen ||
			!strings2.IsAlNumString(newUsername) ||
			isFirstCharNumber(newUsername) {
			return nil, ErrUsernameInvalid
		}

		ok, err := r.UserClient.UserUpdateUsernameByUsername(ctx, &userpb.TLUserUpdateUsernameByUsername{
			PeerType: tg.PEER_USER,
			PeerId:   userId,
			Username: newUsername,
		})
		if err != nil {
			return nil, fmt.Errorf("usernames repository: update username by username: %w", err)
		}
		if !tg.FromBool(ok) {
			return nil, ErrUsernameOccupied
		}
	}

	if oldUsername != "" {
		if _, err := r.UserClient.UserDeleteUsername(ctx, &userpb.TLUserDeleteUsername{
			Username: oldUsername,
		}); err != nil {
			return nil, fmt.Errorf("usernames repository: delete old username: %w", err)
		}
	}

	if _, err := r.UserClient.UserUpdateUsername(ctx, &userpb.TLUserUpdateUsername{
		UserId:   userId,
		Username: newUsername,
	}); err != nil {
		return nil, fmt.Errorf("usernames repository: update username: %w", err)
	}

	// Update the in-memory copy for the sync push.
	me.User.Username = newUsername

	r.pushUpdateUserName(ctx, userId, me)

	return buildSelfUser(me), nil
}

func (r *Repository) pushUpdateUserName(ctx context.Context, userId int64, me *tg.ImmutableUser) {
	if r.SyncClient == nil {
		return
	}

	update := tg.MakeTLUpdateUserName(&tg.TLUpdateUserName{
		UserId:    userId,
		FirstName: userDataFirstName(me.User),
		LastName:  userDataLastName(me.User),
		Usernames: []tg.UsernameClazz{
			tg.MakeTLUsername(&tg.TLUsername{
				Editable: true,
				Active:   true,
				Username: userDataUsername(me.User),
			}),
		},
	})

	if _, err := r.SyncClient.SyncUpdatesNotMe(ctx, &syncpb.TLSyncUpdatesNotMe{
		UserId:  userId,
		Updates: tg.MakeTLUpdates(&tg.TLUpdates{
			Updates: []tg.UpdateClazz{update},
			Users:   []tg.UserClazz{buildSelfUser(me).Clazz},
			Date:    int32(time.Now().Unix()),
		}),
	}); err != nil {
		logx.Errorf("pushUpdateUserName sync failed for userId=%d: %v", userId, err)
	}
}

// ResolveUsername resolves a username to a peer with full details.
func (r *Repository) ResolveUsername(ctx context.Context, selfId int64, username string) (*tg.ContactsResolvedPeer, error) {
	rName, err := r.UserClient.UserResolveUsername(ctx, &userpb.TLUserResolveUsername{
		Username: username,
	})
	if err != nil {
		return nil, fmt.Errorf("usernames repository: resolve username: %w", err)
	}

	resolvedPeer := tg.MakeTLContactsResolvedPeer(&tg.TLContactsResolvedPeer{
		Peer:  rName.Clazz,
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToContactsResolvedPeer()

	switch p := rName.Clazz.(type) {
	case *tg.TLPeerUser:
		mUsers, err := r.UserClient.UserGetMutableUsersV2(ctx, &userpb.TLUserGetMutableUsersV2{
			Id:      []int64{selfId, p.UserId},
			Privacy: true,
			HasTo:   true,
			To:      []int64{selfId},
		})
		if err != nil {
			return nil, fmt.Errorf("usernames repository: resolve username: get mutable users: %w", err)
		}
		if mUsers != nil {
			for _, u := range mUsers.Users {
				if u != nil {
					resolvedPeer.Users = append(resolvedPeer.Users, buildSelfUser(u).Clazz)
				}
			}
		}
	case *tg.TLPeerChat:
		chat, err := r.ChatClient.ChatGetChatBySelfId(ctx, &chatpb.TLChatGetChatBySelfId{
			SelfId: selfId,
			ChatId: p.ChatId,
		})
		if err != nil {
			return nil, fmt.Errorf("usernames repository: resolve username: get chat: %w", err)
		}
		if chat != nil && chat.Chat != nil {
			resolvedPeer.Chats = []tg.ChatClazz{
				projectMutableChat(chat, selfId),
			}
		}
	case *tg.TLPeerChannel:
		if r.Plugin != nil {
			resolvedPeer.Chats = r.Plugin.GetChannelListByIdList(ctx, selfId, p.ChannelId)
		} else {
			return nil, tg.ErrEnterpriseIsBlocked
		}
	}

	return resolvedPeer, nil
}

// --- helpers ---

func hasRPCClientConfig(c kitex.RpcClientConf) bool {
	return len(c.Endpoints) > 0 || len(c.Target) > 0 || c.HasEtcd()
}

func isFirstCharNumber(s string) bool {
	return len(s) > 0 && s[0] >= '0' && s[0] <= '9'
}

func userDataUsername(ud tg.UserDataClazz) string {
	if ud == nil {
		return ""
	}
	return ud.Username
}

func userDataFirstName(ud tg.UserDataClazz) string {
	if ud == nil {
		return ""
	}
	return ud.FirstName
}

func userDataLastName(ud tg.UserDataClazz) string {
	if ud == nil {
		return ""
	}
	return ud.LastName
}

// buildSelfUser builds a *User (wrapper) from an ImmutableUser, setting Self=true.
func buildSelfUser(me *tg.ImmutableUser) *tg.User {
	if me == nil || me.User == nil {
		return nil
	}
	ud := me.User
	return tg.MakeTLUser(&tg.TLUser{
		Self:       true,
		Id:         ud.Id,
		AccessHash: optionalInt64Ptr(ud.AccessHash),
		FirstName:  optionalStringPtr(ud.FirstName),
		LastName:   optionalStringPtr(ud.LastName),
		Username:   optionalStringPtr(ud.Username),
		Phone:      optionalStringPtr(ud.Phone),
		Verified:   ud.Verified,
		Support:    ud.Support,
		Fake:       ud.Fake,
		Premium:    ud.Premium,
		Status:     tg.UserStatusEmptyClazz,
	}).ToUser()
}

func optionalStringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func optionalInt64Ptr(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}

// projectMutableChat converts a MutableChat to a ChatClazz suitable for
// inclusion in a ContactsResolvedPeer response.
func projectMutableChat(chat *tg.MutableChat, selfID int64) tg.ChatClazz {
	if chat == nil || chat.Chat == nil {
		return nil
	}
	return tg.MakeTLChat(&tg.TLChat{
		Creator:             chat.Chat.Creator == selfID,
		Deactivated:         chat.Chat.Deactivated,
		CallActive:          chat.Chat.CallActive,
		CallNotEmpty:        chat.Chat.CallNotEmpty,
		Noforwards:          chat.Chat.Noforwards,
		Id:                  chat.Chat.Id,
		Title:               chat.Chat.Title,
		ParticipantsCount:   chat.Chat.ParticipantsCount,
		Date:                int32(chat.Chat.Date),
		Version:             chat.Chat.Version,
		MigratedTo:          chat.Chat.MigratedTo,
		DefaultBannedRights: chat.Chat.DefaultBannedRights,
	})
}
