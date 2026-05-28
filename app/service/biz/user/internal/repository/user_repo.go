package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetImmutableUser(ctx context.Context, id int64, privacy bool, contacts ...int64) (*tg.ImmutableUser, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}

	userDO, err := r.model.UsersModel.FindOne(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get immutable user %d: %w", userpb.ErrUserStorage, id, err)
	}
	return r.immutableUserFor(ctx, userDO, privacy, contacts)
}

func (r *Repository) GetImmutableUserV2(ctx context.Context, id int64, privacy bool, contacts []int64) (*tg.ImmutableUser, error) {
	return r.GetImmutableUser(ctx, id, privacy, contacts...)
}

// GetMutableUsers returns existing users only; missing ids are omitted.
func (r *Repository) GetMutableUsers(ctx context.Context, ids []int64, privacy bool, contacts []int64) ([]tg.ImmutableUserClazz, error) {
	if len(ids) == 0 {
		return []tg.ImmutableUserClazz{}, nil
	}
	userDOList, err := r.model.UsersModel.SelectUsersByIdList(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%w: get mutable users: %w", userpb.ErrUserStorage, err)
	}

	users := make([]tg.ImmutableUserClazz, 0, len(userDOList))
	for i := range userDOList {
		userData, err := r.immutableUserFor(ctx, &userDOList[i], privacy, contacts)
		if err != nil {
			return nil, err
		}
		users = append(users, userData)
	}
	return users, nil
}

func (r *Repository) GetUserData(ctx context.Context, id int64) (tg.UserDataClazz, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.FindOne(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get user data %d: %w", userpb.ErrUserStorage, id, err)
	}
	return userDataFromModel(userDO), nil
}

func (r *Repository) GetUserDataByToken(ctx context.Context, token string) (tg.UserDataClazz, error) {
	id, err := r.GetUserIDByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.GetUserData(ctx, id)
}

func (r *Repository) GetUserDataList(ctx context.Context, ids []int64) ([]tg.UserDataClazz, error) {
	if len(ids) == 0 {
		return []tg.UserDataClazz{}, nil
	}
	userDOList, err := r.model.UsersModel.SelectUsersByIdList(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%w: get user data list: %w", userpb.ErrUserStorage, err)
	}
	users := make([]tg.UserDataClazz, 0, len(userDOList))
	for i := range userDOList {
		users = append(users, userDataFromModel(&userDOList[i]))
	}
	return users, nil
}

func (r *Repository) GetImmutableUserByPhone(ctx context.Context, phone string) (*tg.ImmutableUser, error) {
	userDO, err := r.model.UsersModel.SelectByPhoneNumber(ctx, phone)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get immutable user by phone %s: %w", userpb.ErrUserStorage, phone, err)
	}
	if userDO == nil {
		return nil, userpb.ErrUserNotFound
	}
	return immutableUserFromModel(userDO), nil
}

func (r *Repository) GetImmutableUserByToken(ctx context.Context, token string) (*tg.ImmutableUser, error) {
	id, err := r.GetUserIDByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return r.GetImmutableUser(ctx, id, false)
}

func (r *Repository) GetUserIDByToken(ctx context.Context, token string) (int64, error) {
	id, err := r.model.BotsModel.SelectByToken(ctx, token)
	if err != nil {
		if isNotFound(err) {
			return 0, userpb.ErrBotNotFound
		}
		return 0, fmt.Errorf("%w: get user id by token: %w", userpb.ErrUserStorage, err)
	}
	if id == 0 {
		return 0, userpb.ErrBotNotFound
	}
	return id, nil
}

func (r *Repository) IsBot(ctx context.Context, id int64) (bool, error) {
	if id == 0 {
		return false, userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.FindOne(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return false, userpb.ErrUserNotFound
		}
		return false, fmt.Errorf("%w: is bot %d: %w", userpb.ErrUserStorage, id, err)
	}
	return userDO.IsBot, nil
}

func (r *Repository) CheckBots(ctx context.Context, ids []int64) ([]int64, error) {
	bots, err := r.model.UsersModel.SelectBots(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%w: check bots: %w", userpb.ErrUserStorage, err)
	}
	return bots, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int64, phone, reason string) error {
	return r.execUserUpdate(ctx, id, "delete user", func() (int64, error) {
		return r.model.UsersModel.Delete(ctx, phone, reason, id)
	})
}

func (r *Repository) UpdateFirstAndLastName(ctx context.Context, id int64, firstName, lastName string) error {
	if id == 0 {
		return userpb.ErrUserNotFound
	}

	rowsAffected, err := r.model.UsersModel.UpdateFirstAndLastName(ctx, firstName, lastName, id)
	if err != nil {
		return fmt.Errorf("%w: update first and last name %d: %w", userpb.ErrUserStorage, id, err)
	}
	if rowsAffected == 0 {
		if err := r.verifyUserUpdateTarget(ctx, id, "update first and last name"); err != nil {
			return err
		}
	}

	if err := r.invalidateUserDataCache(ctx, id, "invalidate user name cache"); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAbout(ctx context.Context, id int64, about string) error {
	return r.execUserUpdate(ctx, id, "update about", func() (int64, error) {
		return r.model.UsersModel.UpdateAbout(ctx, about, id)
	})
}

func (r *Repository) UpdateBirthday(ctx context.Context, id int64, birthday tg.BirthdayClazz) error {
	var birthdayValue string
	if birthday != nil {
		birthdayValue = birthday.ToBirthdayString()
	}
	return r.execUserUpdate(ctx, id, "update birthday", func() (int64, error) {
		return r.model.UsersModel.UpdateBirthday(ctx, birthdayValue, id)
	})
}

func (r *Repository) UpdatePremium(ctx context.Context, id int64, premium bool, months int32) error {
	premiumExpireDate := int64(0)
	updateExpireDate := !premium
	if premium && months > 0 {
		premiumExpireDate = time.Now().AddDate(0, int(months), 0).Unix()
		updateExpireDate = true
	}
	return r.execUserUpdate(ctx, id, "update premium", func() (int64, error) {
		return r.model.UsersModel.UpdatePremium(ctx, premium, premiumExpireDate, updateExpireDate, id)
	})
}

func (r *Repository) UpdateVerified(ctx context.Context, id int64, verified bool) error {
	return r.execUserUpdate(ctx, id, "update verified", func() (int64, error) {
		return r.model.UsersModel.UpdateVerified(ctx, verified, id)
	})
}

func (r *Repository) SetColor(ctx context.Context, id int64, forProfile bool, color int32, backgroundEmojiID int64) error {
	if forProfile {
		return r.execUserUpdate(ctx, id, "set profile color", func() (int64, error) {
			return r.model.UsersModel.UpdateProfileColor(ctx, color, backgroundEmojiID, id)
		})
	}
	return r.execUserUpdate(ctx, id, "set color", func() (int64, error) {
		return r.model.UsersModel.UpdateColor(ctx, color, backgroundEmojiID, id)
	})
}

func (r *Repository) SetMainProfileTab(ctx context.Context, id int64, tab tg.ProfileTabClazz) error {
	return r.execUserUpdate(ctx, id, "set main profile tab", func() (int64, error) {
		return r.model.UsersModel.UpdateMainTab(ctx, profileTabToType(tab), id)
	})
}

func (r *Repository) UpdatePersonalChannel(ctx context.Context, id int64, channelID int64) error {
	return r.execUserUpdate(ctx, id, "update personal channel", func() (int64, error) {
		return r.model.UsersModel.UpdatePersonalChannelId(ctx, channelID, id)
	})
}

func (r *Repository) UpdateEmojiStatus(ctx context.Context, id, documentID int64, until int32) error {
	return r.execUserUpdate(ctx, id, "update emoji status", func() (int64, error) {
		return r.model.UsersModel.UpdateEmojiStatus(ctx, documentID, until, id)
	})
}

func (r *Repository) SetStoriesMaxID(ctx context.Context, id int64, storiesMaxID int32) error {
	return r.execUserUpdate(ctx, id, "set stories max id", func() (int64, error) {
		return r.model.UsersModel.UpdateStoriesMaxId(ctx, storiesMaxID, id)
	})
}

func (r *Repository) ChangePhone(ctx context.Context, id int64, phone string) error {
	if id == 0 {
		return userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.SelectByPhoneNumber(ctx, phone)
	if err != nil {
		if !isNotFound(err) {
			return fmt.Errorf("%w: change phone lookup %s: %w", userpb.ErrUserStorage, phone, err)
		}
		userDO = nil
	}
	if userDO != nil && userDO.Id != id {
		return userpb.ErrPhoneNumberInUse
	}
	rowsAffected, err := r.model.UsersModel.UpdatePhone(ctx, phone, id)
	if err != nil {
		if errors.Is(err, model.ErrDuplicatePhone) {
			return userpb.ErrPhoneNumberInUse
		}
		return fmt.Errorf("%w: change phone %d: %w", userpb.ErrUserStorage, id, err)
	}
	if rowsAffected == 0 {
		if err := r.verifyUserUpdateTarget(ctx, id, "change phone"); err != nil {
			return err
		}
	}
	if err := r.invalidateUserDataCache(ctx, id, "invalidate phone user cache"); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserIDByPhone(ctx context.Context, phone string) (int64, error) {
	userDO, err := r.model.UsersModel.SelectByPhoneNumber(ctx, phone)
	if err != nil {
		if isNotFound(err) {
			return 0, userpb.ErrUserNotFound
		}
		return 0, fmt.Errorf("%w: get user id by phone %s: %w", userpb.ErrUserStorage, phone, err)
	}
	if userDO == nil {
		return 0, userpb.ErrUserNotFound
	}
	return userDO.Id, nil
}

func (r *Repository) GetCountryCode(ctx context.Context, id int64) (string, error) {
	if id == 0 {
		return "", userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.SelectCountryCode(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return "", userpb.ErrUserNotFound
		}
		return "", fmt.Errorf("%w: get country code %d: %w", userpb.ErrUserStorage, id, err)
	}
	if userDO == nil {
		return "", userpb.ErrUserNotFound
	}
	return userDO.CountryCode, nil
}

func (r *Repository) SetAccountDaysTTL(ctx context.Context, id int64, ttl int32) error {
	return r.execUserUpdate(ctx, id, "set account days ttl", func() (int64, error) {
		return r.model.UsersModel.UpdateAccountDaysTTL(ctx, ttl, id)
	})
}

func (r *Repository) GetAccountDaysTTL(ctx context.Context, id int64) (*tg.AccountDaysTTL, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.SelectAccountDaysTTL(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get account days ttl %d: %w", userpb.ErrUserStorage, id, err)
	}
	if userDO == nil {
		return nil, userpb.ErrUserNotFound
	}
	return tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: userDO.AccountDaysTtl}).ToAccountDaysTTL(), nil
}

func (r *Repository) SetAuthorizationTTL(ctx context.Context, id int64, ttl int32) error {
	return r.execUserUpdate(ctx, id, "set authorization ttl", func() (int64, error) {
		return r.model.UsersModel.UpdateAuthorizationTTL(ctx, ttl, id)
	})
}

func (r *Repository) GetAuthorizationTTL(ctx context.Context, id int64) (*tg.AccountDaysTTL, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	userDO, err := r.model.UsersModel.SelectAuthorizationTTL(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get authorization ttl %d: %w", userpb.ErrUserStorage, id, err)
	}
	if userDO == nil {
		return nil, userpb.ErrUserNotFound
	}
	return tg.MakeTLAccountDaysTTL(&tg.TLAccountDaysTTL{Days: userDO.AuthorizationTtlDays}).ToAccountDaysTTL(), nil
}

func (r *Repository) SetDefaultHistoryTTL(ctx context.Context, id int64, ttl int32) error {
	if id == 0 {
		return userpb.ErrUserNotFound
	}
	_, _, err := r.model.DefaultHistoryTtlModel.InsertOrUpdate(ctx, &model.DefaultHistoryTtl{
		UserId: id,
		Period: ttl,
	})
	if err != nil {
		return fmt.Errorf("%w: set default history ttl %d: %w", userpb.ErrUserStorage, id, err)
	}
	return nil
}

func (r *Repository) GetDefaultHistoryTTL(ctx context.Context, id int64) (*tg.DefaultHistoryTTL, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	ttlDO, err := r.model.DefaultHistoryTtlModel.Select(ctx, id)
	if err != nil {
		if !isNotFound(err) {
			return nil, fmt.Errorf("%w: get default history ttl %d: %w", userpb.ErrUserStorage, id, err)
		}
		ttlDO = nil
	}
	period := int32(0)
	if ttlDO != nil {
		period = ttlDO.Period
	}
	return tg.MakeTLDefaultHistoryTTL(&tg.TLDefaultHistoryTTL{Period: period}).ToDefaultHistoryTTL(), nil
}

func (r *Repository) UpdateLastSeen(ctx context.Context, id, lastSeenAt int64, expires int32) error {
	if id == 0 {
		return userpb.ErrUserNotFound
	}
	_, _, err := r.model.UserPresencesModel.InsertOrUpdate(ctx, &model.UserPresences{
		UserId:     id,
		LastSeenAt: lastSeenAt,
		Expires:    expires,
	})
	if err != nil {
		return fmt.Errorf("%w: update last seen %d: %w", userpb.ErrUserStorage, id, err)
	}
	r.invalidateProjectionPresenceCache(ctx, id)
	return nil
}

func (r *Repository) GetLastSeen(ctx context.Context, id int64) (userpb.LastSeenDataClazz, error) {
	if id == 0 {
		return nil, userpb.ErrUserNotFound
	}
	presenceDO, err := r.model.UserPresencesModel.Select(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: get last seen %d: %w", userpb.ErrUserStorage, id, err)
	}
	if presenceDO == nil {
		return nil, userpb.ErrUserNotFound
	}
	return lastSeenDataFromModel(presenceDO), nil
}

func (r *Repository) GetLastSeens(ctx context.Context, ids []int64) ([]userpb.LastSeenDataClazz, error) {
	if len(ids) == 0 {
		return []userpb.LastSeenDataClazz{}, nil
	}
	presenceDOList, err := r.model.UserPresencesModel.SelectList(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%w: get last seens: %w", userpb.ErrUserStorage, err)
	}
	datas := make([]userpb.LastSeenDataClazz, 0, len(presenceDOList))
	for i := range presenceDOList {
		datas = append(datas, lastSeenDataFromModel(&presenceDOList[i]))
	}
	return datas, nil
}

func immutableUserFromModel(do *model.Users) tg.ImmutableUserClazz {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User:             userDataFromModel(do),
		LastSeenAt:       0,
		Contacts:         []tg.ContactDataClazz{},
		ReverseContacts:  []tg.ContactDataClazz{},
		KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
	}).ToImmutableUser()
}

func (r *Repository) immutableUserFor(ctx context.Context, do *model.Users, privacy bool, contactIDs []int64) (*tg.ImmutableUser, error) {
	lastSeenAt, err := r.immutableLastSeen(ctx, do.Id)
	if err != nil {
		return nil, err
	}
	contacts, reverseContacts, err := r.immutableContacts(ctx, do.Id, contactIDs)
	if err != nil {
		return nil, err
	}
	privacyRules, err := r.immutablePrivacyRules(ctx, do.Id, privacy)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User:             userDataFromModel(do),
		LastSeenAt:       lastSeenAt,
		Contacts:         contacts,
		ReverseContacts:  reverseContacts,
		KeysPrivacyRules: privacyRules,
	}).ToImmutableUser(), nil
}

func (r *Repository) immutableLastSeen(ctx context.Context, id int64) (int64, error) {
	presenceDO, err := r.model.UserPresencesModel.Select(ctx, id)
	if err != nil {
		if !isNotFound(err) {
			return 0, fmt.Errorf("%w: get immutable last seen %d: %w", userpb.ErrUserStorage, id, err)
		}
		presenceDO = nil
	}
	if presenceDO == nil {
		return 0, nil
	}
	return presenceDO.LastSeenAt, nil
}

func (r *Repository) immutableContacts(ctx context.Context, id int64, contactIDs []int64) ([]tg.ContactDataClazz, []tg.ContactDataClazz, error) {
	if len(contactIDs) == 0 {
		return []tg.ContactDataClazz{}, []tg.ContactDataClazz{}, nil
	}
	contacts := make([]tg.ContactDataClazz, 0, len(contactIDs))
	reverseContacts := make([]tg.ContactDataClazz, 0, len(contactIDs))
	for _, contactID := range contactIDs {
		if contactID == 0 {
			continue
		}
		contactDO, err := r.model.UserContactsModel.SelectContact(ctx, id, contactID)
		if err != nil {
			if !isNotFound(err) {
				return nil, nil, fmt.Errorf("%w: get immutable contact %d/%d: %w", userpb.ErrUserStorage, id, contactID, err)
			}
			contactDO = nil
		}
		if contactDO != nil {
			contacts = append(contacts, makeContactData(contactDO))
		}
		reverseContactDO, err := r.model.UserContactsModel.SelectContact(ctx, contactID, id)
		if err != nil {
			if !isNotFound(err) {
				return nil, nil, fmt.Errorf("%w: get immutable reverse contact %d/%d: %w", userpb.ErrUserStorage, contactID, id, err)
			}
			reverseContactDO = nil
		}
		if reverseContactDO != nil {
			reverseContacts = append(reverseContacts, makeContactData(reverseContactDO))
		}
	}
	return contacts, reverseContacts, nil
}

func (r *Repository) immutablePrivacyRules(ctx context.Context, id int64, privacy bool) ([]tg.PrivacyKeyRulesClazz, error) {
	if !privacy {
		return []tg.PrivacyKeyRulesClazz{}, nil
	}
	privacyDOList, err := r.model.UserPrivaciesModel.SelectPrivacyAll(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w: get immutable privacy %d: %w", userpb.ErrUserStorage, id, err)
	}
	rules := make([]tg.PrivacyKeyRulesClazz, 0, len(privacyDOList))
	for i := range privacyDOList {
		decodedRules, err := decodePrivacyRules(privacyDOList[i].Rules)
		if err != nil {
			return nil, fmt.Errorf("%w: decode immutable privacy %d/%d: %w", userpb.ErrUserStorage, id, privacyDOList[i].KeyType, err)
		}
		rules = append(rules, tg.MakeTLPrivacyKeyRules(&tg.TLPrivacyKeyRules{
			Key:   privacyDOList[i].KeyType,
			Rules: decodedRules,
		}).ToPrivacyKeyRules())
	}
	return rules, nil
}

func userDataFromModel(do *model.Users) tg.UserDataClazz {
	if do == nil {
		return nil
	}
	return tg.MakeTLUserData(&tg.TLUserData{
		Id:                do.Id,
		AccessHash:        do.AccessHash,
		UserType:          do.UserType,
		SceretKeyId:       do.SecretKeyId,
		FirstName:         do.FirstName,
		LastName:          do.LastName,
		Username:          do.Username,
		Phone:             do.Phone,
		CountryCode:       do.CountryCode,
		Verified:          do.Verified,
		Support:           do.Support,
		Scam:              do.Scam,
		Fake:              do.Fake,
		About:             stringPtr(do.About),
		Restricted:        do.Restricted,
		RestrictionReason: []tg.RestrictionReasonClazz{},
		Deleted:           do.Deleted,
		Premium:           do.Premium,
		EmojiStatus:       emojiStatusFromCacheDTO(do.EmojiStatusDocumentId, do.EmojiStatusUntil),
		Color:             peerColorFromCacheDTO(do.Color, do.ColorBackgroundEmojiId),
		ProfileColor:      peerColorFromCacheDTO(do.ProfileColor, do.ProfileColorBackgroundEmojiId),
		StoriesMaxId:      do.StoriesMaxId,
		Birthday:          do.Birthday,
		PersonalChannelId: do.PersonalChannelId,
	}).ToUserData()
}

func tgUsernameList(username string, editable bool) []tg.UsernameClazz {
	if username == "" {
		return []tg.UsernameClazz{}
	}
	return []tg.UsernameClazz{
		tg.MakeTLUsername(&tg.TLUsername{
			Username: username,
			Active:   true,
			Editable: editable,
		}),
	}
}

func lastSeenDataFromModel(do *model.UserPresences) userpb.LastSeenDataClazz {
	if do == nil {
		return nil
	}
	return userpb.MakeTLLastSeenData(&userpb.TLLastSeenData{
		UserId:     do.UserId,
		LastSeenAt: do.LastSeenAt,
		Expires:    do.Expires,
	}).ToLastSeenData()
}

func stringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func (r *Repository) execUserUpdate(ctx context.Context, id int64, op string, fn func() (int64, error)) error {
	if id == 0 {
		return userpb.ErrUserNotFound
	}
	rowsAffected, err := fn()
	if err != nil {
		return fmt.Errorf("%w: %s %d: %w", userpb.ErrUserStorage, op, id, err)
	}
	if rowsAffected == 0 {
		if err := r.verifyUserUpdateTarget(ctx, id, op); err != nil {
			return err
		}
	}
	if err := r.invalidateUserDataCache(ctx, id, "invalidate user cache"); err != nil {
		return err
	}
	return nil
}

func (r *Repository) verifyUserUpdateTarget(ctx context.Context, id int64, op string) error {
	_, err := r.model.UsersModel.FindOne(ctx, id)
	return userUpdateTargetError(op, id, err)
}

func userUpdateTargetError(op string, id int64, err error) error {
	if err == nil {
		return nil
	}
	if isNotFound(err) {
		return userpb.ErrUserNotFound
	}
	return fmt.Errorf("%w: %s verify user %d: %w", userpb.ErrUserStorage, op, id, err)
}

func profileTabToType(tab tg.ProfileTabClazz) int32 {
	switch tab.(type) {
	case *tg.TLProfileTabPosts:
		return 0
	case *tg.TLProfileTabGifts:
		return 1
	case *tg.TLProfileTabMedia:
		return 2
	case *tg.TLProfileTabFiles:
		return 3
	case *tg.TLProfileTabMusic:
		return 4
	case *tg.TLProfileTabVoice:
		return 5
	case *tg.TLProfileTabLinks:
		return 6
	case *tg.TLProfileTabGifs:
		return 7
	default:
		return 0
	}
}
