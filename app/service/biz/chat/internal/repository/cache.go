package repository

import "fmt"

func chatAggregateCacheKey(chatID int64) string {
	return fmt.Sprintf("%s:%d", chatAggregateCacheKeyPrefix, chatID)
}

func chatParticipantCacheKey(chatID, userID int64) string {
	return fmt.Sprintf("%s:%d:%d", chatParticipantCacheKeyPrefix, chatID, userID)
}

func createChatFloodKey(userID int64) string {
	return fmt.Sprintf("%s:%d", createChatFloodKeyPrefix, userID)
}

func chatParticipantCacheKeys(chatID int64, userIDs []int64) []string {
	keys := make([]string, 0, len(userIDs))
	for _, userID := range userIDs {
		keys = append(keys, chatParticipantCacheKey(chatID, userID))
	}
	return keys
}

func chatAggregateAndParticipantCacheKeys(chatID int64, userIDs []int64) []string {
	keys := []string{chatAggregateCacheKey(chatID)}
	return append(keys, chatParticipantCacheKeys(chatID, userIDs)...)
}
