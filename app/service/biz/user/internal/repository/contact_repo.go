package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) AddContact(ctx context.Context, userID int64, contactUserID int64, firstName, lastName, phone string) (bool, error) {
	var mutual bool
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		currentDO, err := r.model.UserContactsModel.SelectContactTx(tx, userID, contactUserID)
		if err != nil {
			if !isNotFound(err) {
				return fmt.Errorf("select contact: %w", err)
			}
			currentDO = nil
		}
		if currentDO != nil {
			mutual = currentDO.Mutual
		}

		reverseDO, err := r.model.UserContactsModel.SelectContactTx(tx, contactUserID, userID)
		if err != nil {
			if !isNotFound(err) {
				return fmt.Errorf("select reverse contact: %w", err)
			}
			reverseDO = nil
		}
		if reverseDO != nil {
			mutual = true
		}

		if _, _, err := r.model.UserContactsModel.InsertOrUpdateTx(tx, &model.UserContacts{
			OwnerUserId:      userID,
			ContactUserId:    contactUserID,
			ContactPhone:     phone,
			ContactFirstName: firstName,
			ContactLastName:  lastName,
			Mutual:           mutual,
			Date2:            time.Now().Unix(),
		}); err != nil {
			return fmt.Errorf("insert contact: %w", err)
		}

		if reverseDO != nil && !reverseDO.Mutual {
			if _, err := r.model.UserContactsModel.UpdateMutualTx(tx, true, contactUserID, userID); err != nil {
				return fmt.Errorf("update reverse mutual: %w", err)
			}
		}

		return nil
	}); err != nil {
		return false, fmt.Errorf("%w: add contact %d/%d: %w", userpb.ErrUserStorage, userID, contactUserID, err)
	}
	if err := r.invalidateContactCaches(ctx, userID, contactUserID); err != nil {
		return false, err
	}
	return mutual, nil
}

func (r *Repository) DeleteContact(ctx context.Context, userID int64, contactUserID int64) error {
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if _, err := r.model.UserContactsModel.DeleteContactsTx(tx, userID, []int64{contactUserID}); err != nil {
			return fmt.Errorf("delete contact: %w", err)
		}
		if _, err := r.model.UserContactsModel.UpdateMutualTx(tx, false, contactUserID, userID); err != nil {
			return fmt.Errorf("update reverse mutual: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("%w: delete contact %d/%d: %w", userpb.ErrUserStorage, userID, contactUserID, err)
	}
	return r.invalidateContactCaches(ctx, userID, contactUserID)
}

func (r *Repository) CheckContact(ctx context.Context, userID int64, contactUserID int64) (bool, error) {
	contactDO, err := r.model.UserContactsModel.SelectContact(ctx, userID, contactUserID)
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("%w: check contact %d/%d: %w", userpb.ErrUserStorage, userID, contactUserID, err)
	}
	return contactDO != nil, nil
}

func (r *Repository) GetContact(ctx context.Context, userID int64, contactUserID int64) (tg.ContactDataClazz, error) {
	contactDO, err := r.model.UserContactsModel.SelectContact(ctx, userID, contactUserID)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrContactNotFound
		}
		return nil, fmt.Errorf("%w: get contact %d/%d: %w", userpb.ErrUserStorage, userID, contactUserID, err)
	}
	if contactDO == nil {
		return nil, userpb.ErrContactNotFound
	}
	return makeContactData(contactDO), nil
}

func (r *Repository) GetContactList(ctx context.Context, userID int64) (*userpb.VectorContactData, error) {
	contactList, err := r.model.UserContactsModel.SelectUserContacts(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get contact list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	datas := make([]tg.ContactDataClazz, 0, len(contactList))
	for i := range contactList {
		datas = append(datas, makeContactData(&contactList[i]))
	}
	return &userpb.VectorContactData{Datas: datas}, nil
}

func (r *Repository) GetContactIDList(ctx context.Context, userID int64) (*userpb.VectorLong, error) {
	idList, err := r.model.UserContactsModel.SelectUserContactIdList(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get contact id list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	if idList == nil {
		idList = []int64{}
	}
	return &userpb.VectorLong{Datas: idList}, nil
}

func (r *Repository) EditCloseFriends(ctx context.Context, userID int64, idList []int64) error {
	contactList, err := r.model.UserContactsModel.SelectUserContacts(ctx, userID)
	if err != nil {
		return fmt.Errorf("%w: get close friend list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	currentIDs := make([]int64, 0, len(contactList))
	for i := range contactList {
		if contactList[i].CloseFriend {
			currentIDs = append(currentIDs, contactList[i].ContactUserId)
		}
	}
	if len(currentIDs) == 0 && len(idList) == 0 {
		return nil
	}
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if len(currentIDs) > 0 {
			if _, err := r.model.UserContactsModel.UpdateCloseFriendTx(tx, false, userID, currentIDs); err != nil {
				return fmt.Errorf("clear close friends: %w", err)
			}
		}
		if len(idList) > 0 {
			if _, err := r.model.UserContactsModel.UpdateCloseFriendTx(tx, true, userID, idList); err != nil {
				return fmt.Errorf("set close friends: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("%w: edit close friends %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return r.invalidateUserDataCache(ctx, userID, "invalidate close friends cache")
}

func (r *Repository) SetStoriesHidden(ctx context.Context, userID int64, contactUserID int64, hidden bool) error {
	if _, err := r.model.UserContactsModel.UpdateStoriesHidden(ctx, hidden, userID, contactUserID); err != nil {
		return fmt.Errorf("%w: set stories hidden %d/%d: %w", userpb.ErrUserStorage, userID, contactUserID, err)
	}
	return r.invalidateUserDataCache(ctx, userID, "invalidate stories hidden cache")
}

func (r *Repository) GetImportersByPhone(ctx context.Context, phone string) (*userpb.VectorInputContact, error) {
	importers, err := r.model.UnregisteredContactsModel.SelectImportersByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("%w: get importers by phone %s: %w", userpb.ErrUserStorage, phone, err)
	}
	datas := make([]tg.InputContactClazz, 0, len(importers))
	for i := range importers {
		datas = append(datas, tg.MakeTLInputPhoneContact(&tg.TLInputPhoneContact{
			ClientId:  importers[i].ImporterUserId,
			Phone:     "",
			FirstName: importers[i].ImportFirstName,
			LastName:  importers[i].ImportLastName,
		}).ToInputContact())
	}
	return &userpb.VectorInputContact{Datas: datas}, nil
}

func (r *Repository) DeleteImportersByPhone(ctx context.Context, phone string) error {
	if _, err := r.model.UnregisteredContactsModel.DeleteImportersByPhone(ctx, phone); err != nil {
		return fmt.Errorf("%w: delete importers by phone %s: %w", userpb.ErrUserStorage, phone, err)
	}
	return nil
}

func makeContactData(contactDO *model.UserContacts) tg.ContactDataClazz {
	firstName := contactDO.ContactFirstName
	lastName := contactDO.ContactLastName
	phone := contactDO.ContactPhone
	return tg.MakeTLContactData(&tg.TLContactData{
		UserId:        contactDO.OwnerUserId,
		ContactUserId: contactDO.ContactUserId,
		FirstName:     &firstName,
		LastName:      &lastName,
		MutualContact: contactDO.Mutual,
		Phone:         &phone,
		CloseFriend:   contactDO.CloseFriend,
		StoriesHidden: contactDO.StoriesHidden,
	}).ToContactData()
}

func (r *Repository) invalidateContactCaches(ctx context.Context, userID, contactUserID int64) error {
	if err := r.invalidateUserDataCache(ctx, userID, "invalidate contact owner cache"); err != nil {
		return err
	}
	if err := r.invalidateUserDataCache(ctx, contactUserID, "invalidate contact peer cache"); err != nil {
		return err
	}
	return nil
}

func (r *Repository) invalidateUserDataCache(ctx context.Context, userID int64, operation string) error {
	if err := r.DelCache(ctx, userDataCacheKey(userID)); err != nil {
		return fmt.Errorf("%w: %s %d: %w", userpb.ErrUserStorage, operation, userID, err)
	}
	return nil
}
