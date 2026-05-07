package repository

import "fmt"

const (
	userDataCacheKeyPrefix             = "user_data.v3#"
	projectionFactsCacheKeyPrefix      = "user:facts:v1:"
	projectionPrivacyCacheKeyPrefix    = "user:privacy:v1:"
	projectionContactMapCacheKeyPrefix = "user:contact-map:v1:"
	projectionPresenceCacheKeyPrefix   = "user:presence:v1:"
)

func userDataCacheKey(userID int64) string {
	return fmt.Sprintf("%s%d", userDataCacheKeyPrefix, userID)
}

func projectionFactsCacheKey(userID int64) string {
	return fmt.Sprintf("%s%d", projectionFactsCacheKeyPrefix, userID)
}

func projectionPrivacyCacheKey(userID int64) string {
	return fmt.Sprintf("%s%d", projectionPrivacyCacheKeyPrefix, userID)
}

func projectionContactMapCacheKey(ownerUserID int64) string {
	return fmt.Sprintf("%s%d", projectionContactMapCacheKeyPrefix, ownerUserID)
}

func projectionPresenceCacheKey(userID int64) string {
	return fmt.Sprintf("%s%d", projectionPresenceCacheKeyPrefix, userID)
}
