package dao

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func (d *Dao) AddToPtsQueue(ctx context.Context, userId int64, pts, ptsCount int32, update *mtproto.Update) int32 {
	i, err := d.AddToPtsQueueE(ctx, userId, pts, ptsCount, update)
	if err != nil {
		logx.WithContext(ctx).Errorf("AddToPtsQueue - error: %v", err)
	}
	return i
}

func (d *Dao) AddToPtsQueueE(ctx context.Context, userId int64, pts, ptsCount int32, update *mtproto.Update) (int32, error) {
	updateData, err := jsonx.Marshal(update)
	if err != nil {
		return 0, err
	}

	do := &dataobject.UserPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: mtproto.GetUpdateType(update),
		UpdateData: string(updateData),
		Date2:      time.Now().Unix(),
	}

	i, _, err := d.UserPtsUpdatesDAO.Insert(ctx, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("AddToPtsQueue - insert into user_pts_updates error: %v, do: %v", err, do)
		return int32(i), err
	}
	return int32(i), nil
}

func (d *Dao) AddToPtsQueueTx(tx *sqlx.Tx, userId int64, pts, ptsCount int32, update *mtproto.Update) (int32, error) {
	updateData, err := jsonx.Marshal(update)
	if err != nil {
		return 0, err
	}

	do := &dataobject.UserPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: mtproto.GetUpdateType(update),
		UpdateData: string(updateData),
		Date2:      time.Now().Unix(),
	}

	i, _, err := d.UserPtsUpdatesDAO.InsertTx(tx, do)
	return int32(i), err
}

// CheckPtsContinuity checks whether a user's pts sequence is continuous
// starting from afterPts. Returns any gaps found as (expectedPts, actualPts) pairs.
func (d *Dao) CheckPtsContinuity(ctx context.Context, userId int64, afterPts int32) (gaps [][2]int32, err error) {
	rows, err := d.UserPtsUpdatesDAO.SelectByGtPts(ctx, userId, afterPts)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}

	expectedPts := afterPts
	for _, row := range rows {
		nextExpected := expectedPts + row.PtsCount
		if row.Pts != nextExpected {
			gaps = append(gaps, [2]int32{nextExpected, row.Pts})
			logx.WithContext(ctx).Errorf("pts gap detected - user_id: %d, expected_pts: %d, actual_pts: %d",
				userId, nextExpected, row.Pts)
		}
		expectedPts = row.Pts
	}

	return gaps, nil
}
