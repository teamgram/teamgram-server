// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package user

import (
	"time"

	"github.com/teamgram/proto/mtproto"
)

var (
	userStatusEmpty = mtproto.MakeTLUserStatusEmpty(nil).To_UserStatus()
)

func MakeUserStatusOnline() *mtproto.UserStatus {
	now := time.Now().Unix()
	status := mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
		Expires: int32(now + 60),
	}).To_UserStatus()
	return status
}

func MakeUserStatusOffline(lastSeenAt int64) *mtproto.UserStatus {
	return mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
		WasOnline: int32(lastSeenAt),
	}).To_UserStatus()
}

// MakeUserStatus
/**********************************************************************************
    // Android client source code
	//
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
func MakeUserStatus(lastSeenAt int64, allowTimestamp bool) *mtproto.UserStatus {
	now := time.Now().Unix()

	if allowTimestamp {
		if now <= lastSeenAt+60 {
			status := mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
				Expires: int32(lastSeenAt + 60),
			}).To_UserStatus()
			return status
		} else {
			status := mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
				WasOnline: int32(lastSeenAt),
			}).To_UserStatus()
			return status
		}
	} else {
		if now-lastSeenAt >= 60*60*24*30 {
			return nil
		} else if now-lastSeenAt >= 60*60*24*7 {
			return mtproto.MakeTLUserStatusLastMonth(nil).To_UserStatus()
		} else if now-lastSeenAt >= 60*60*24*3 {
			return mtproto.MakeTLUserStatusLastWeek(nil).To_UserStatus()
		} else {
			return mtproto.MakeTLUserStatusRecently(nil).To_UserStatus()
		}
	}
}
