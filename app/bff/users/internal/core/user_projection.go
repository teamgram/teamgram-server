package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func projectImmutableUser(immutableUser tg.ImmutableUserClazz) tg.UserClazz {
	if immutableUser == nil || immutableUser.User == nil {
		return tg.MakeTLUserEmpty(&tg.TLUserEmpty{})
	}

	data := immutableUser.User
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
		Usernames:          usernameList(data.Username, false),
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

func projectSelfImmutableUser(immutableUser tg.ImmutableUserClazz) tg.UserClazz {
	user := projectImmutableUser(immutableUser)
	if full, ok := user.(*tg.TLUser); ok {
		full.Self = true
		markUsernamesEditable(full.Usernames)
	}
	return user
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

func nonEmptyStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func usernameList(username string, editable bool) []tg.UsernameClazz {
	if username == "" {
		return []tg.UsernameClazz{}
	}
	return []tg.UsernameClazz{
		tg.MakeTLUsername(&tg.TLUsername{
			Username: username,
			Active:   true,
			Editable: editable,
		}),
	}
}

func markUsernamesEditable(usernames []tg.UsernameClazz) {
	for _, username := range usernames {
		if username != nil {
			username.Editable = true
		}
	}
}

func recentStoryIDPtr(id int32) tg.RecentStoryClazz {
	if id == 0 {
		return nil
	}
	return tg.MakeTLRecentStory(&tg.TLRecentStory{MaxId: &id})
}
