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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

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
		Photo:               projectChatPhoto(chat.Chat.Photo),
		ParticipantsCount:   chat.Chat.ParticipantsCount,
		Date:                int32(chat.Chat.Date),
		Version:             chat.Chat.Version,
		MigratedTo:          chat.Chat.MigratedTo,
		DefaultBannedRights: chat.Chat.DefaultBannedRights,
	})
}

func projectChatPhoto(photo tg.PhotoClazz) tg.ChatPhotoClazz {
	if p, ok := photo.(*tg.TLPhoto); ok {
		return tg.MakeTLChatPhoto(&tg.TLChatPhoto{
			PhotoId: p.Id,
			DcId:    p.DcId,
		})
	}

	return tg.MakeTLChatPhotoEmpty(&tg.TLChatPhotoEmpty{})
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
