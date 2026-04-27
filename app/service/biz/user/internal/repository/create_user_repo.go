package repository

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	userTypeRegular int32 = 2
	userTypeTest    int32 = 5

	defaultMinTestUserID = 10000000 - 1
	defaultMaxTestUserID = 20000000
)

func (r *Repository) CreateNewUser(ctx context.Context, secretKeyID int64, phone, countryCode, firstName, lastName string) (*tg.ImmutableUser, error) {
	now := time.Now().Unix()
	userDO := &model.Users{
		UserType:       userTypeRegular,
		AccessHash:     rand.Int63(),
		Phone:          phone,
		SecretKeyId:    secretKeyID,
		FirstName:      firstName,
		LastName:       lastName,
		CountryCode:    countryCode,
		AccountDaysTtl: 548,
	}
	id, _, err := r.model.UsersModel.Insert(ctx, userDO)
	if err != nil {
		if sqlx.IsDuplicate(err) {
			return nil, userpb.ErrPhoneNumberInUse
		}
		return nil, fmt.Errorf("%w: create user: %w", userpb.ErrUserStorage, err)
	}
	userDO.Id = id

	if err := r.UpdateLastSeen(ctx, userDO.Id, now, 300); err != nil {
		return nil, err
	}
	return immutableUserFromModelWithLastSeen(userDO, now), nil
}

func (r *Repository) CreateNewTestUser(ctx context.Context, secretKeyID, minID, maxID int64) (*tg.ImmutableUser, error) {
	now := time.Now().Unix()
	if minID == 0 {
		minID = defaultMinTestUserID
	}
	if maxID == 0 {
		maxID = defaultMaxTestUserID
	}

	id := minID + 1
	for retry := 0; retry <= 10; retry++ {
		userDO := newTestUser(id, secretKeyID)
		if _, _, err := r.model.UsersModel.InsertTestUser(ctx, userDO); err != nil {
			if !sqlx.IsDuplicate(err) {
				return nil, fmt.Errorf("%w: create test user %d: %w", userpb.ErrUserStorage, id, err)
			}

			nextDO, err := r.model.UsersModel.SelectNextTestUserId(ctx, maxID)
			if err != nil {
				return nil, fmt.Errorf("%w: select next test user id: %w", userpb.ErrUserStorage, err)
			}
			if nextDO == nil || nextDO.Id+1 >= maxID {
				return nil, fmt.Errorf("%w: next test user id not found", userpb.ErrUserStorage)
			}
			id = nextDO.Id + 1
			continue
		}

		return immutableUserFromModelWithLastSeen(userDO, now), nil
	}

	return nil, fmt.Errorf("%w: create test user retry exhausted", userpb.ErrUserStorage)
}

func newTestUser(id, secretKeyID int64) *model.Users {
	idText := strconv.FormatInt(id, 10)
	return &model.Users{
		Id:             id,
		UserType:       userTypeTest,
		AccessHash:     rand.Int63(),
		Phone:          "-" + idText,
		SecretKeyId:    secretKeyID,
		FirstName:      "t" + idText,
		LastName:       "",
		CountryCode:    "CN",
		AccountDaysTtl: 180,
	}
}

func immutableUserFromModelWithLastSeen(do *model.Users, lastSeenAt int64) *tg.ImmutableUser {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User:             userDataFromModel(do),
		LastSeenAt:       lastSeenAt,
		Contacts:         []tg.ContactDataClazz{},
		ReverseContacts:  []tg.ContactDataClazz{},
		KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
	}).ToImmutableUser()
}
