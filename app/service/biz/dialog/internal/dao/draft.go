// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

func (d *Dao) GetNoCacheAllDraftIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		draftIdList []int64
	)

	_, err := d.DialogsDAO.SelectAllDraftsWithCB(
		ctx,
		userId,
		func(sz, i int, v *dataobject.DialogsDO) {
			if i == 0 {
				draftIdList = make([]int64, 0, sz)
			}
			draftIdList = append(draftIdList, v.Id)
		})
	if err != nil {
		return nil, err
	}

	if draftIdList == nil {
		draftIdList = make([]int64, 0)
	}

	return draftIdList, nil
}

func (d *Dao) GetAllDraftIdList(ctx context.Context, userId int64) ([]int64, error) {
	var (
		draftIdList []int64
	)

	err := d.CachedConn.QueryRow(
		ctx,
		&draftIdList,
		dialog.GetAllDraftIdListCacheKey(userId),
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			idList, err := d.GetNoCacheAllDraftIdList(ctx, userId)
			if err != nil {
				return err
			}
			*(v.(*[]int64)) = idList

			return nil
		})
	if err != nil {
		return nil, err
	}

	if draftIdList == nil {
		draftIdList = make([]int64, 0)
	}

	return draftIdList, nil
}

// GetDraftMessage returns draft message by id.
func (d *Dao) GetDraftMessage(ctx context.Context, userId, peerDialogId int64) (*mtproto.DraftMessage, error) {
	dlgExt, err := d.GetDialogByPeerDialogId(ctx, userId, peerDialogId)
	if err != nil {
		return nil, err
	}

	return dlgExt.GetDialog().GetDraft(), nil
}

func (d *Dao) GetAllDraftMessageList(ctx context.Context, userId int64) ([]*mtproto.DraftMessage, error) {
	idList, err := d.GetAllDraftIdList(ctx, userId)
	if err != nil {
		return nil, err
	}

	dlgExtList, err := d.GetDialogListByIdList(ctx, userId, idList)
	if err != nil {
		return nil, err
	}

	vList := make([]*mtproto.DraftMessage, 0, len(idList))
	for _, dlgExt := range dlgExtList {
		if dlgExt.GetDialog().GetDraft() != nil {
			vList = append(vList, dlgExt.GetDialog().GetDraft())
		}
	}

	return vList, nil
}
