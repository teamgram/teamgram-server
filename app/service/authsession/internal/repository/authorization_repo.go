package repository

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/zeromicro/go-zero/core/logx"
)

func authDataToAuthorization(data *cacheAuthData, current bool, country string, region string) *tg.Authorization {
	if data == nil || data.Client == nil {
		return nil
	}

	auth := tg.MakeTLAuthorization(&tg.TLAuthorization{
		// TODO: infer OfficialApp, Platform, and AppName from normalized client
		// metadata instead of treating lang_pack as the application name.
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

func (r *Repository) GetAuthorizationByAuthKeyId(ctx context.Context, authKeyId int64) (*tg.Authorization, error) {
	permAuthKeyId, err := r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
	if err != nil {
		return nil, err
	}
	return r.GetAuthorization(ctx, permAuthKeyId)
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

func (r *Repository) GetAuthorizationsByAuthKeyId(ctx context.Context, userId int64, excludeAuthKeyId int64) ([]*tg.Authorization, error) {
	excludePermAuthKeyId := int64(0)
	if excludeAuthKeyId != 0 {
		var err error
		excludePermAuthKeyId, err = r.GetPermAuthKeyIdByAuthKeyId(ctx, excludeAuthKeyId)
		if err != nil {
			return nil, err
		}
	}
	return r.GetAuthorizations(ctx, userId, excludePermAuthKeyId)
}

func (r *Repository) GetAuthorizations(ctx context.Context, userId int64, excludePermAuthKeyId int64) ([]*tg.Authorization, error) {
	rows, err := r.model.AuthUsersModel.SelectListByUserId(ctx, userId)
	if err != nil {
		return nil, wrapStorage(err)
	}
	if len(rows) == 0 {
		return []*tg.Authorization{}, nil
	}

	authKeyIds := make([]int64, 0, len(rows))
	for i := range rows {
		authKeyIds = append(authKeyIds, rows[i].AuthKeyId)
	}
	authRows, err := r.model.AuthsModel.FindListByAuthKeyIdList(ctx, authKeyIds...)
	if err != nil {
		return nil, wrapStorage(err)
	}
	authByKeyId := make(map[int64]*model.Auths, len(authRows))
	for i := range authRows {
		row := authRows[i]
		authByKeyId[row.AuthKeyId] = &row
	}

	authorizations := make([]*tg.Authorization, 0, len(rows))
	for i := range rows {
		authRow := authByKeyId[rows[i].AuthKeyId]
		if authRow == nil {
			logx.WithContext(ctx).Errorf("authsession.GetAuthorizations - auth data missing, user_id: %d, auth_key_id: %d",
				userId,
				rows[i].AuthKeyId)
			continue
		}
		data := authDataFromRows(rows[i].AuthKeyId, authRow, &rows[i])
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

func (r *Repository) ResetAuthorizationByAuthKeyId(ctx context.Context, userId int64, authKeyId int64, hash int64) ([]int64, error) {
	excludePermAuthKeyId := int64(0)
	if authKeyId != 0 {
		var err error
		excludePermAuthKeyId, err = r.GetPermAuthKeyIdByAuthKeyId(ctx, authKeyId)
		if err != nil {
			return nil, err
		}
	}
	keyIds, err := r.ResetAuthorization(ctx, userId, excludePermAuthKeyId, hash)
	if err != nil {
		return nil, err
	}
	return r.ExpandAuthKeyIds(ctx, keyIds)
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
