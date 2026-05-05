package core

import (
	"fmt"
	"hash/fnv"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
)

func (c *DialogCore) sourcePermAuthKeyID() (int64, error) {
	if c == nil || c.MD == nil || c.MD.PermAuthKeyId == 0 {
		return 0, dialogpb.ErrSourceAuthKeyRequired
	}
	return c.MD.PermAuthKeyId, nil
}

func deterministicOperationID(kind string, userID int64, parts ...interface{}) string {
	return fmt.Sprintf("v1:dialog:%s:user:%d:%v", kind, userID, parts)
}

func deterministicOutboxID(operationID string, salt string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(operationID))
	_, _ = h.Write([]byte(":"))
	_, _ = h.Write([]byte(salt))
	return int64(h.Sum64() & 0x7fffffffffffffff)
}
