// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package core

import (
	"fmt"
)

const (
	userKeyIdsPrefix = "user_online_keys"
	maxBatchUsers    = 500
)

// hsetAndExpireScript atomically sets a hash field and refreshes the key TTL.
// KEYS[1] = key, ARGV[1] = field, ARGV[2] = value, ARGV[3] = expire seconds
const hsetAndExpireScript = `
redis.call('HSET', KEYS[1], ARGV[1], ARGV[2])
redis.call('EXPIRE', KEYS[1], ARGV[3])
return 1
`

func getUserKey(id int64) string {
	return fmt.Sprintf("%s#%d", userKeyIdsPrefix, id)
}
