package core

import (
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *ChatInvitesCore) fetchUserClazzes(ids []int64, selfID int64) ([]tg.UserClazz, error) {
	ids = uniqueInt64s(ids)
	if len(ids) == 0 {
		return []tg.UserClazz{}, nil
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: ids,
		To: []int64{selfID},
	})
	if err != nil {
		c.Logger.Errorf("chatinvites.fetchUserClazzes - user.getMutableUsers failed: ids: %v, self_id: %d, err: %v", ids, selfID, err)
		return nil, tg.ErrInternalServerError
	}

	byID := make(map[int64]tg.UserClazz, len(ids))
	if users != nil {
		for _, immutableUser := range users.Datas {
			user := projectImmutableUser(immutableUser)
			if id, ok := userID(user); ok {
				byID[id] = user
			}
		}
	}

	result := make([]tg.UserClazz, 0, len(ids))
	for _, id := range ids {
		if user := byID[id]; user != nil {
			result = append(result, user)
		} else {
			result = append(result, tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: id}))
		}
	}
	return result, nil
}

func projectImmutableUser(immutableUser tg.ImmutableUserClazz) tg.UserClazz {
	if immutableUser == nil || immutableUser.User == nil {
		return tg.MakeTLUserEmpty(&tg.TLUserEmpty{})
	}

	data := immutableUser.User
	if data == nil {
		return tg.MakeTLUserEmpty(&tg.TLUserEmpty{})
	}
	if data.Deleted {
		return tg.MakeTLUser(&tg.TLUser{
			Id:      data.Id,
			Deleted: true,
		})
	}

	accessHash := data.AccessHash
	return tg.MakeTLUser(&tg.TLUser{
		Id:                 data.Id,
		AccessHash:         &accessHash,
		FirstName:          nonEmptyStringPtr(data.FirstName),
		LastName:           nonEmptyStringPtr(data.LastName),
		Username:           nonEmptyStringPtr(data.Username),
		Phone:              nonEmptyStringPtr(data.Phone),
		Photo:              projectUserProfilePhoto(data.ProfilePhoto),
		Status:             tg.MakeTLUserStatusEmpty(&tg.TLUserStatusEmpty{}),
		Bot:                data.Bot != nil,
		Verified:           data.Verified,
		Restricted:         data.Restricted,
		Scam:               data.Scam,
		Fake:               data.Fake,
		Premium:            data.Premium,
		Support:            data.Support,
		RestrictionReason:  data.RestrictionReason,
		EmojiStatus:        data.EmojiStatus,
		StoriesUnavailable: data.StoriesUnavailable,
		Color:              data.Color,
		ProfileColor:       data.ProfileColor,
		StoriesMaxId:       recentStoryIDPtr(data.StoriesMaxId),
	})
}

func projectUserProfilePhoto(photo tg.PhotoClazz) tg.UserProfilePhotoClazz {
	if p, ok := photo.(*tg.TLPhoto); ok {
		return tg.MakeTLUserProfilePhoto(&tg.TLUserProfilePhoto{
			PhotoId: p.Id,
			DcId:    p.DcId,
		})
	}
	return nil
}

func userID(user tg.UserClazz) (int64, bool) {
	switch u := user.(type) {
	case *tg.TLUser:
		return u.Id, true
	case *tg.TLUserEmpty:
		return u.Id, true
	default:
		return 0, false
	}
}

func uniqueInt64s(ids []int64) []int64 {
	seen := make(map[int64]struct{}, len(ids))
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}

func nonEmptyStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func recentStoryIDPtr(id int32) tg.RecentStoryClazz {
	if id == 0 {
		return nil
	}
	return tg.MakeTLRecentStory(&tg.TLRecentStory{MaxId: &id})
}
