// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/teamgram-server/app/service/biz/report/internal/dal/dataobject"
)

func (d *Dao) Report(ctx context.Context, userId int64, reportType, peerType int32, peerId, photoId, messageSenderUserId int64, messageId, reason int32, text string) (bool, error) {
	do := &dataobject.ReportsDO{
		UserId:              userId,
		ReportType:          reportType,
		PeerType:            peerType,
		PeerId:              peerId,
		ProfilePhotoId:      photoId,
		MessageSenderUserId: messageSenderUserId,
		MessageId:           messageId,
		Reason:              reason,
		Text:                text,
	}
	id, _, err := d.ReportsDAO.Insert(ctx, do)
	return id > 0, err
}

func (d *Dao) ReportIdList(ctx context.Context, userId int64, reportType, peerType int32, peerId, photoId, messageSenderUserId int64, messageIdList []int32, reason int32, text string) (bool, error) {
	bulkDOList := make([]*dataobject.ReportsDO, 0, len(messageIdList))

	for _, id := range messageIdList {
		bulkDOList = append(bulkDOList, &dataobject.ReportsDO{
			UserId:              userId,
			ReportType:          reportType,
			PeerType:            peerType,
			PeerId:              peerId,
			ProfilePhotoId:      photoId,
			MessageSenderUserId: messageSenderUserId,
			MessageId:           id,
			Reason:              reason,
			Text:                text,
		})
	}

	lastInsertId, _, err := d.ReportsDAO.InsertBulk(ctx, bulkDOList)
	return lastInsertId > 0, err
}
