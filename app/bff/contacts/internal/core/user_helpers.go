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

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

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

func (c *ContactsCore) projectUsers(ids []int64, missing userprojection.MissingPolicy) ([]tg.UserClazz, error) {
	return userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, ids, missing)
}

func projectedUserByID(users []tg.UserClazz, id int64) tg.UserClazz {
	for _, user := range users {
		if userID, ok := userID(user); ok && userID == id {
			return user
		}
	}
	return nil
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
