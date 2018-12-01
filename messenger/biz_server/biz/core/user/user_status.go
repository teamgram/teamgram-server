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
	"github.com/golang/glog"
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
func makeUserStatus(do *dataobject.UserPresencesDO, showStatus bool) *mtproto.UserStatus {
	now := time.Now().Unix()

	if showStatus {
		if now <= do.LastSeenAt+60 {
			status := &mtproto.TLUserStatusOnline{Data2: &mtproto.UserStatus_Data{
				Expires: int32(do.LastSeenAt + 60),
			}}
			return status.To_UserStatus()
		} else {
			status := &mtproto.TLUserStatusOffline{Data2: &mtproto.UserStatus_Data{
				WasOnline: int32(do.LastSeenAt),
			}}
			return status.To_UserStatus()
		}
	} else {
		if now - do.LastSeenAt >= 60*60*24*30 {
			return nil
		} else if now - do.LastSeenAt >= 60*60*24*7 {
			return mtproto.NewTLUserStatusLastMonth().To_UserStatus()
		} else if now - do.LastSeenAt >= 60*60*24*3 {
			return mtproto.NewTLUserStatusLastWeek().To_UserStatus()
		} else {
			return mtproto.NewTLUserStatusRecently().To_UserStatus()
		}
	}
}

/******

// https://telegram.org/faq#q-can-i-hide-my-last-seen-time

Q: Can I hide my ‘last seen’ time?

You can choose who sees this info in [Privacy and Security](!https://telegram.org/blog/privacy-revolution) settings.

Remember that you won‘t see Last Seen timestamps for people with whom you don’t share your own.
You will, however, see an approximate last seen value.
This keeps stalkers away but makes it possible to understand whether a person is reachable over Telegram.
There are four possible approximate values:

- Last seen recently — covers anything between 1 second and 2-3 days
- Last seen within a week — between 2-3 and seven days
- Last seen within a month — between 6-7 days and a month
- Last seen a long time ago — more than a month (this is also always shown to blocked users)

Q: Who can see me ‘online’?

The last seen rules apply to your online status as well.
People can only see you online if you're sharing your last seen status with them.
There is one exception to this: people will be able to see you online for a brief period
when you send them a message in a one-on-one chat or in a group where you both are members.

 */

func (m *UserModel) GetUserStatus(selfId int32, presencesDO *dataobject.UserPresencesDO, isContact, isBlocked bool) *mtproto.UserStatus {
	if isBlocked {
		return nil
	}

	if presencesDO == nil {
		return nil
	}

	showStatus := m.accountCallback.CheckShowStatus(presencesDO.UserId, selfId, isContact)
	glog.Info("showStatus: ", presencesDO, ", selfId: ", selfId, ", showStatus: ", showStatus)
	return makeUserStatus(presencesDO, showStatus)

	//do := m.dao.UserPresencesDAO.Select(userId)
	//if do == nil {
	//	return nil
	//	// return mtproto.NewTLUserStatusEmpty().To_UserStatus()
	//} else {
	//	return makeUserStatus(do, true)
	//}
}


func (m *UserModel) GetUserStatus2(selfId, userId int32, isContact, isBlocked bool) *mtproto.UserStatus {
	if isBlocked {
		return nil
	}

	do := m.dao.UserPresencesDAO.Select(userId)
	return m.GetUserStatus(selfId, do, isContact, isBlocked)
	//if do == nil {
	//	return nil
	//	// return mtproto.NewTLUserStatusEmpty().To_UserStatus()
	//}

	//showStatus := m.accountCallback.GetPrivacyShowStatus(userId, userId, isContact)
	//return makeUserStatus(do, showStatus)
}

//func (m *UserModel) GetUserStatusList(idList []int32) []*mtproto.UserStatus {
//	statusList := make([]*mtproto.UserStatus, 0, len(idList))
//	doList := m.dao.UserPresencesDAO.SelectList(idList)
//	if len(doList) == len(idList) {
//		for idx := 0; idx < len(doList); idx++ {
//			statusList = append(statusList, makeUserStatus(&doList[idx], true))
//		}
//	} else {
//		f := func(id int32) int {
//			for i := 0; i < len(doList); i++ {
//				if doList[i].UserId == id {
//					return i
//				}
//			}
//			return -1
//		}
//		for _, id := range idList {
//			idx := f(id)
//			if idx == -1 {
//				statusList = append(statusList, userStatusEmpty)
//			} else {
//				statusList = append(statusList, makeUserStatus(&doList[idx], true))
//			}
//		}
//	}
//	return statusList
//}

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
