// Copyright © 2025 The Teamgooo Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package core

import (
	"fmt"
	"strconv"
	"strings"
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

func getIdByUserKey(k string) int64 {
	a := strings.Split(k, "#")
	if len(a) < 2 {
		return 0
	}
	i, _ := strconv.ParseInt(a[1], 10, 64)

	return i
}
