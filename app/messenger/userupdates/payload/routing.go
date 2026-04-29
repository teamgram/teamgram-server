package payload

import (
	"encoding/binary"
	"hash/fnv"
)

type UserRoute struct {
	BucketID            int
	ReceiverPartitionID int
	PushPartitionID     int
}

func RouteUser(userID int64) UserRoute {
	bucketID := hashUserID(userID) % BucketCount
	return UserRoute{
		BucketID:            bucketID,
		ReceiverPartitionID: bucketID % ReceiverPartitionCount,
		PushPartitionID:     bucketID % PushPartitionCount,
	}
}

func hashUserID(userID int64) int {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(userID))
	h := fnv.New64a()
	_, _ = h.Write(b[:])
	return int(h.Sum64() % uint64(BucketCount))
}
