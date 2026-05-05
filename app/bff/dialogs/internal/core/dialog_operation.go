package core

import (
	"fmt"
	"hash/fnv"
	"time"
)

func dialogOperationID(kind string, userID int64, token int64) string {
	return fmt.Sprintf("v1:dialog:%s:user:%d:token:%d", kind, userID, token)
}

func dialogOutboxID(operationID string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(operationID))
	id := int64(h.Sum64() & 0x7fffffffffffffff)
	if id == 0 {
		return 1
	}
	return id
}

func dialogOperationToken() int64 {
	return time.Now().UnixNano()
}
