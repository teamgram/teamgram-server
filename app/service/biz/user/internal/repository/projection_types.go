package repository

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

type ProjectionConfig struct {
	SQLInChunkSize         int
	MaxViewerUserIds       int
	MaxTargetUserIds       int
	MaxProjectionPairs     int
	ContactMapCacheEnabled bool
	ContactMapMaxEntries   int
}

type UserProjectionBundle struct {
	Facts          []tg.ImmutableUserClazz
	ViewerUsers    []ViewerUsers
	MissingUserIds []int64
}

type ViewerUsers struct {
	ViewerUserId int64
	Users        []tg.UserClazz
}
