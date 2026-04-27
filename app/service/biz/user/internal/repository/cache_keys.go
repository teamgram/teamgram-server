package repository

import "fmt"

const userDataCacheKeyPrefix = "user_data.v3#"

func userDataCacheKey(userID int64) string {
	return fmt.Sprintf("%s%d", userDataCacheKeyPrefix, userID)
}
