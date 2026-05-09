package repository

import (
	"context"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetFullUser(ctx context.Context, selfUserID int64, id int64) (*tg.UsersUserFull, error) {
	if selfUserID == 0 || id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	targetDO, err := r.model.UsersModel.FindOne(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get full user %d: %w", userpb.ErrUserStorage, id, err)
	}

	self := selfUserID == id
	blocked := false
	if !self {
		blocked, err = r.IsPeerBlocked(ctx, selfUserID, tg.PEER_USER, id)
		if err != nil {
			return nil, err
		}
	}

	settings, err := r.GetPeerSettings(ctx, selfUserID, tg.PEER_USER, id)
	if err != nil {
		return nil, err
	}
	notifySettings, err := r.GetNotifySettings(ctx, selfUserID, tg.PEER_USER, id)
	if err != nil {
		return nil, err
	}

	var botInfo tg.BotInfoClazz
	if targetDO.IsBot {
		info, err := r.GetBotInfo(ctx, id)
		if err != nil {
			return nil, err
		}
		botInfo = info
	}

	projectedUser, err := r.projectFullUserUser(ctx, selfUserID, id)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLUsersUserFull(&tg.TLUsersUserFull{
		FullUser: tg.MakeTLUserFull(&tg.TLUserFull{
			Blocked:          blocked,
			Id:               id,
			About:            stringPtr(targetDO.About),
			Settings:         settings,
			NotifySettings:   notifySettings,
			BotInfo:          botInfo,
			CommonChatsCount: 0,
		}),
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{projectedUser},
	}).ToUsersUserFull(), nil
}

func (r *Repository) projectFullUserUser(ctx context.Context, selfUserID int64, id int64) (tg.UserClazz, error) {
	bundle, err := r.GetUserProjectionBundle(ctx, []int64{selfUserID}, []int64{id}, false)
	if err != nil {
		return nil, err
	}
	return fullUserProjectedUserFromBundle(bundle, selfUserID, id)
}

func fullUserProjectedUserFromBundle(bundle *UserProjectionBundle, selfUserID int64, id int64) (tg.UserClazz, error) {
	if bundle == nil {
		return nil, userpb.ErrUserNotFound
	}
	for _, viewerUsers := range bundle.ViewerUsers {
		if viewerUsers.ViewerUserId != selfUserID {
			continue
		}
		for _, user := range viewerUsers.Users {
			if userID(user) == id {
				return user, nil
			}
		}
	}
	return nil, userpb.ErrUserNotFound
}

func userID(user tg.UserClazz) int64 {
	switch u := user.(type) {
	case *tg.TLUser:
		return u.Id
	case *tg.TLUserEmpty:
		return u.Id
	default:
		return 0
	}
}
