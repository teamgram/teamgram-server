package core

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

func projectAccountSelfImmutableUser(immutableUser tg.ImmutableUserClazz) tg.UserClazz {
	if immutableUser == nil || immutableUser.User == nil {
		return tg.MakeTLUserEmpty(&tg.TLUserEmpty{})
	}
	data := immutableUser.User
	if data.Deleted {
		return tg.MakeTLUser(&tg.TLUser{Id: data.Id, Deleted: true})
	}
	accessHash := data.AccessHash
	return tg.MakeTLUser(&tg.TLUser{
		Self:       true,
		Id:         data.Id,
		AccessHash: &accessHash,
		FirstName:  nonEmptyAccountStringPtr(data.FirstName),
		LastName:   nonEmptyAccountStringPtr(data.LastName),
		Username:   nonEmptyAccountStringPtr(data.Username),
		Phone:      nonEmptyAccountStringPtr(data.Phone),
		Status:     tg.MakeTLUserStatusEmpty(&tg.TLUserStatusEmpty{}),
		Verified:   data.Verified,
		Support:    data.Support,
		Fake:       data.Fake,
		Premium:    data.Premium,
	})
}

func nonEmptyAccountStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
