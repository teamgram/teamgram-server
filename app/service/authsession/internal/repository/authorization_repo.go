package repository

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func authDataToAuthorization(data *cacheAuthData, current bool, country string, region string) *tg.Authorization {
	if data == nil || data.Client == nil {
		return nil
	}

	auth := tg.MakeTLAuthorization(&tg.TLAuthorization{
		Current:         current,
		OfficialApp:     true,
		PasswordPending: false,
		DeviceModel:     data.Client.DeviceModel,
		Platform:        "",
		SystemVersion:   data.Client.SystemVersion,
		ApiId:           data.Client.ApiId,
		AppName:         data.Client.LangPack,
		AppVersion:      data.Client.AppVersion,
		Ip:              data.Client.Ip,
		Country:         country,
		Region:          region,
	})
	if data.BindUser != nil {
		auth.DateCreated = int32(data.BindUser.DateCreated)
		auth.DateActive = int32(data.BindUser.DateActivated)
		if !current {
			auth.Hash = data.BindUser.Hash
		}
	}
	return auth.ToAuthorization()
}

func (r *Repository) GetAuthorization(ctx context.Context, permAuthKeyId int64) (*tg.Authorization, error) {
	data, err := r.GetAuthData(ctx, permAuthKeyId)
	if err != nil {
		return nil, err
	}
	if data == nil || data.Client == nil {
		return nil, authsession.ErrAuthorizationNotFound
	}

	country, region := r.getCountryAndRegionByIP(ctx, data.Client.Ip)
	return authDataToAuthorization(data, true, country, region), nil
}

func (r *Repository) GetAuthorizations(ctx context.Context, userId int64, excludePermAuthKeyId int64) ([]*tg.Authorization, error) {
	rows, err := r.model.AuthUsersModel.SelectListByUserId(ctx, userId)
	if err != nil {
		return nil, wrapStorage(err)
	}
	if len(rows) == 0 {
		return []*tg.Authorization{}, nil
	}

	authorizations := make([]*tg.Authorization, 0, len(rows))
	for i := range rows {
		data, err := r.GetAuthData(ctx, rows[i].AuthKeyId)
		if err != nil {
			return nil, err
		}
		current := rows[i].AuthKeyId == excludePermAuthKeyId
		country, region := r.getCountryAndRegionByIP(ctx, data.Client.Ip)
		authorization := authDataToAuthorization(data, current, country, region)
		if authorization == nil {
			continue
		}
		if current {
			authorizations = append([]*tg.Authorization{authorization}, authorizations...)
		} else {
			authorizations = append(authorizations, authorization)
		}
	}

	return authorizations, nil
}

func (r *Repository) ResetAuthorization(ctx context.Context, userId int64, excludePermAuthKeyId int64, hash int64) ([]int64, error) {
	rows, err := r.model.AuthUsersModel.SelectListByUserId(ctx, userId)
	if err != nil {
		return nil, wrapStorage(err)
	}

	var (
		cacheKeys []string
		rowIds    []int64
		keyIds    []int64
	)
	for i := range rows {
		if hash == 0 {
			if excludePermAuthKeyId == rows[i].AuthKeyId {
				continue
			}
		} else if hash != rows[i].Hash || excludePermAuthKeyId == rows[i].AuthKeyId {
			continue
		}

		cacheKeys = append(cacheKeys, authDataCacheKey(rows[i].AuthKeyId))
		rowIds = append(rowIds, rows[i].Id)
		keyIds = append(keyIds, rows[i].AuthKeyId)
	}
	if len(keyIds) == 0 {
		return []int64{}, nil
	}

	_, _, err = r.CachedConn.Exec(ctx, func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
		rowsAffected, err := r.model.AuthUsersModel.DeleteByHashList(ctx, rowIds)
		return 0, rowsAffected, err
	}, cacheKeys...)
	if err != nil {
		return nil, wrapStorage(err)
	}

	return keyIds, nil
}
