package repository

import (
	"context"
	"fmt"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxUserSearchLimit int32 = 50

func (r *Repository) SearchUsers(ctx context.Context, q string, excludedContacts []int64, limit int32) (*userpb.UsersFound, error) {
	dataMode := len(excludedContacts) == 0
	if len(q) < 3 || limit <= 0 {
		return emptyUsersFound(dataMode), nil
	}
	if limit > maxUserSearchLimit {
		limit = maxUserSearchLimit
	}

	excludedIDs := excludedContacts
	if len(excludedIDs) == 0 {
		excludedIDs = []int64{0}
	}

	idList, err := r.model.UsersModel.SearchByQueryString(ctx, q+"%", "%"+q+"%", excludedIDs, limit)
	if err != nil {
		return nil, fmt.Errorf("%w: search users: %w", userpb.ErrUserStorage, err)
	}
	if !dataMode {
		return userpb.MakeTLUsersIdFound(&userpb.TLUsersIdFound{
			IdList: idList,
		}).ToUsersFound(), nil
	}

	users, err := r.GetUserDataList(ctx, idList)
	if err != nil {
		return nil, err
	}
	return userpb.MakeTLUsersDataFound(&userpb.TLUsersDataFound{
		Count:      int32(len(idList)),
		Users:      users,
		NextOffset: "",
	}).ToUsersFound(), nil
}

func emptyUsersFound(dataMode bool) *userpb.UsersFound {
	if dataMode {
		return userpb.MakeTLUsersDataFound(&userpb.TLUsersDataFound{
			Count:      0,
			Users:      []tg.UserDataClazz{},
			NextOffset: "",
		}).ToUsersFound()
	}

	return userpb.MakeTLUsersIdFound(&userpb.TLUsersIdFound{
		IdList: []int64{},
	}).ToUsersFound()
}
