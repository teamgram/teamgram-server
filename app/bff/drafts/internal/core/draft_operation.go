package core

import (
	"fmt"
	"hash/fnv"
)

func draftOperationID(kind string, userID int64, peerType int32, peerID int64, token int64) string {
	return fmt.Sprintf("v1:dialog:draft:%s:user:%d:peer:%d:%d:token:%d", kind, userID, peerType, peerID, token)
}

func clearAllDraftsOperationID(userID int64, token int64) string {
	return fmt.Sprintf("v1:dialog:draft:clear_all:user:%d:token:%d", userID, token)
}

func draftOutboxID(operationID string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(operationID))
	id := int64(h.Sum64() & 0x7fffffffffffffff)
	if id == 0 {
		return 1
	}
	return id
}
