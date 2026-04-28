package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
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
	contact := self
	mutual := self
	blocked := false
	var contactDOForUser *model.UserContacts
	if !self {
		contactDO, err := r.model.UserContactsModel.SelectContact(ctx, selfUserID, id)
		if err != nil {
			if !isNotFound(err) {
				return nil, fmt.Errorf("%w: get full user contact %d/%d: %w", userpb.ErrUserStorage, selfUserID, id, err)
			}
			contactDO = nil
		}
		contactDOForUser = contactDO
		contact = contactDO != nil

		reverseDO, err := r.model.UserContactsModel.SelectContact(ctx, id, selfUserID)
		if err != nil {
			if !isNotFound(err) {
				return nil, fmt.Errorf("%w: get full user reverse contact %d/%d: %w", userpb.ErrUserStorage, id, selfUserID, err)
			}
			reverseDO = nil
		}
		mutual = reverseDO != nil && contact

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

	user := userFromModel(targetDO, self, contact, mutual, contactDOForUser)
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
		Users: []tg.UserClazz{user},
	}).ToUsersUserFull(), nil
}
