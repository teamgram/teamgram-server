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

package core

import (
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func projectImmutableUser(immutableUser tg.ImmutableUserClazz) tg.UserClazz {
	if immutableUser == nil || immutableUser.User == nil {
		return tg.MakeTLUserEmpty(&tg.TLUserEmpty{})
	}

	data := immutableUser.User
	if data.Deleted {
		return tg.MakeTLUser(&tg.TLUser{Id: data.Id, Deleted: true})
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

func immutableUserByID(users []tg.ImmutableUserClazz, id int64) tg.ImmutableUserClazz {
	for _, immutableUser := range users {
		if immutableUser != nil && immutableUser.User != nil && immutableUser.User.Id == id {
			return immutableUser
		}
	}
	return nil
}

func projectUsersByIDs(users []tg.ImmutableUserClazz, ids []int64) []tg.UserClazz {
	byID := make(map[int64]tg.UserClazz, len(users))
	for _, immutableUser := range users {
		user := projectImmutableUser(immutableUser)
		if id, ok := userID(user); ok {
			byID[id] = user
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
	return result
}

func contactDatasToContacts(datas []tg.ContactDataClazz) []tg.ContactClazz {
	contacts := make([]tg.ContactClazz, 0, len(datas))
	for _, data := range datas {
		if data == nil {
			continue
		}
		contacts = append(contacts, tg.MakeTLContact(&tg.TLContact{
			UserId: data.ContactUserId,
			Mutual: tg.ToBoolClazz(data.MutualContact),
		}).ToContact())
	}
	return contacts
}

func makePeerSettings() tg.PeerSettingsClazz {
	return tg.MakeTLPeerSettings(&tg.TLPeerSettings{}).ToPeerSettings()
}

func makeUserStatus(lastSeenAt int64, allowTimestamp bool) tg.UserStatusClazz {
	now := time.Now().Unix()
	if allowTimestamp {
		if now <= lastSeenAt+60 {
			return tg.MakeTLUserStatusOnline(&tg.TLUserStatusOnline{Expires: int32(lastSeenAt + 60)})
		}
		return tg.MakeTLUserStatusOffline(&tg.TLUserStatusOffline{WasOnline: int32(lastSeenAt)})
	}
	if now-lastSeenAt >= 60*60*24*30 {
		return nil
	}
	if now-lastSeenAt >= 60*60*24*7 {
		return tg.MakeTLUserStatusLastMonth(&tg.TLUserStatusLastMonth{})
	}
	if now-lastSeenAt >= 60*60*24*3 {
		return tg.MakeTLUserStatusLastWeek(&tg.TLUserStatusLastWeek{})
	}
	return tg.MakeTLUserStatusRecently(&tg.TLUserStatusRecently{})
}
