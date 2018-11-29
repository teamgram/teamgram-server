// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package user

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"time"
)

var (
	userStatusEmpty = mtproto.NewTLUserStatusEmpty().To_UserStatus()
)

func makeUserStatusOnline() *mtproto.UserStatus {
	now := time.Now().Unix()
	status := &mtproto.UserStatus{
		Constructor: mtproto.TLConstructor_CRC32_userStatusOnline,
		Data2: &mtproto.UserStatus_Data{
			Expires: int32(now + 60),
		},
	}
	return status
}

/*
    public static String formatUserStatus(int currentAccount, TLRPC.User user) {
        if (user != null && user.status != null && user.status.expires == 0) {
            if (user.status instanceof TLRPC.TL_userStatusRecently) {
                user.status.expires = -100;
            } else if (user.status instanceof TLRPC.TL_userStatusLastWeek) {
                user.status.expires = -101;
            } else if (user.status instanceof TLRPC.TL_userStatusLastMonth) {
                user.status.expires = -102;
            }
        }
        if (user != null && user.status != null && user.status.expires <= 0) {
            if (MessagesController.getInstance(currentAccount).onlinePrivacy.containsKey(user.id)) {
                return getString("Online", R.string.Online);
            }
        }
        if (user == null || user.status == null || user.status.expires == 0 || UserObject.isDeleted(user) || user instanceof TLRPC.TL_userEmpty) {
            return getString("ALongTimeAgo", R.string.ALongTimeAgo);
        } else {
            int currentTime = ConnectionsManager.getInstance(currentAccount).getCurrentTime();
            if (user.status.expires > currentTime) {
                return getString("Online", R.string.Online);
            } else {
                if (user.status.expires == -1) {
                    return getString("Invisible", R.string.Invisible);
                } else if (user.status.expires == -100) {
                    return getString("Lately", R.string.Lately);
                } else if (user.status.expires == -101) {
                    return getString("WithinAWeek", R.string.WithinAWeek);
                } else if (user.status.expires == -102) {
                    return getString("WithinAMonth", R.string.WithinAMonth);
                }  else {
                    return formatDateOnline(user.status.expires);
                }
            }
        }
    }

	// 60*60*24*7   week
	// 60*60*24*30  month
 */
func makeUserStatus(do *dataobject.UserPresencesDO) *mtproto.UserStatus {
	now := time.Now().Unix()

	if now <= do.LastSeenAt+5*60 {
		status := &mtproto.TLUserStatusOnline{Data2: &mtproto.UserStatus_Data{
			Expires: int32(do.LastSeenAt + 5*30),
		}}
		return status.To_UserStatus()
	} else {
		// TODO(@benqi): gen userStatusRecently, userStatusLastWeek, userStatusLastMonth
		status := &mtproto.TLUserStatusOffline{Data2: &mtproto.UserStatus_Data{
			WasOnline: int32(do.LastSeenAt),
		}}
		return status.To_UserStatus()
	}
}

func (m *UserModel) GetUserStatus(userId int32) *mtproto.UserStatus {
	do := m.dao.UserPresencesDAO.Select(userId)
	if do == nil {
		return mtproto.NewTLUserStatusEmpty().To_UserStatus()
	} else {
		return makeUserStatus(do)
	}
}

func (m *UserModel) GetUserStatusList(idList []int32) []*mtproto.UserStatus {
	statusList := make([]*mtproto.UserStatus, 0, len(idList))
	doList := m.dao.UserPresencesDAO.SelectList(idList)
	if len(doList) == len(idList) {
		for idx := 0; idx < len(doList); idx++ {
			statusList = append(statusList, makeUserStatus(&doList[idx]))
		}
	} else {
		f := func(id int32) int {
			for i := 0; i < len(doList); i++ {
				if doList[i].UserId == id {
					return i
				}
			}
			return -1
		}
		for _, id := range idList {
			idx := f(id)
			if idx == -1 {
				statusList = append(statusList, userStatusEmpty)
			} else {
				statusList = append(statusList, makeUserStatus(&doList[idx]))
			}
		}
	}
	return statusList
}

func (m *UserModel) UpdateUserStatus(userId int32, lastSeenAt int64) {
	if userId > 0 {
		userPresencesDO := &dataobject.UserPresencesDO{
			UserId:            userId,
			LastSeenAt:        lastSeenAt,
			LastSeenAuthKeyId: 0,
			LastSeenIp:        "",
		}
		m.dao.UserPresencesDAO.InsertOrUpdate(userPresencesDO)
	}
}

func (m *UserModel) DeleteUser(userId int32, reason string) bool {
	affected := m.dao.UsersDAO.Delete(reason, userId)
	return affected == 1
}
