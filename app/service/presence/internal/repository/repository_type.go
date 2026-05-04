// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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

package repository

import (
	"fmt"
	"strconv"
)

const (
	userKeyPrefix    = "presence:user"
	cleanupKeyPrefix = "presence:cleanup"
	MaxBatchUsers    = 200
)

const hsetAndExpireScript = `
redis.call('HSET', KEYS[1], ARGV[1], ARGV[2])
redis.call('EXPIRE', KEYS[1], ARGV[3])
return 1
`

func userKey(userID int64) string {
	return fmt.Sprintf("%s:%d", userKeyPrefix, userID)
}

func cleanupKey(userID int64) string {
	return fmt.Sprintf("%s:%d", cleanupKeyPrefix, userID)
}

func sessionField(authKeyID, sessionID int64) string {
	return strconv.FormatInt(authKeyID, 10) + ":" + strconv.FormatInt(sessionID, 10)
}
