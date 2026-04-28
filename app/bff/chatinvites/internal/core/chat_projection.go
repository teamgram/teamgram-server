package core

import "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"

func selfID(md *metadata.RpcMetadata) int64 {
	if md == nil {
		return 0
	}
	return md.UserId
}
