package core

import (
	"context"

	"github.com/teamgram/proto/mtproto"
)

func (c *InboxCore) persistPtsUpdate(ctx context.Context, userId int64, update *mtproto.Update) {
	if update == nil {
		return
	}
	if update.Pts_INT32 <= 0 || update.PtsCount <= 0 {
		c.Logger.Errorf("persistPtsUpdate - invalid pts data, user_id: %d, pts: %d, pts_count: %d, update_type: %d",
			userId, update.Pts_INT32, update.PtsCount, mtproto.GetUpdateType(update))
		return
	}
	if _, err := c.svcCtx.Dao.AddToPtsQueueE(ctx, userId, update.Pts_INT32, update.PtsCount, update); err != nil {
		c.Logger.Errorf("persistPtsUpdate - AddToPtsQueueE error, user_id: %d, pts: %d, pts_count: %d, update_type: %d, err: %v",
			userId, update.Pts_INT32, update.PtsCount, mtproto.GetUpdateType(update), err)
	}
}
