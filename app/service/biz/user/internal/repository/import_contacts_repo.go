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

func (r *Repository) ImportContacts(ctx context.Context, userID int64, contacts []tg.InputContactClazz) (*userpb.UserImportedContacts, error) {
	if _, err := r.GetUserData(ctx, userID); err != nil {
		return nil, err
	}

	phoneList := inputContactPhones(contacts)
	registeredUsers, err := r.model.UsersModel.SelectUsersByPhoneList(ctx, phoneList)
	if err != nil {
		return nil, fmt.Errorf("%w: import contacts lookup phones: %w", userpb.ErrUserStorage, err)
	}
	userByPhone := make(map[string]*model.Users, len(registeredUsers))
	registeredIDs := make([]int64, 0, len(registeredUsers))
	for i := range registeredUsers {
		userDO := registeredUsers[i]
		userByPhone[userDO.Phone] = &userDO
		registeredIDs = append(registeredIDs, userDO.Id)
	}

	currentContacts, err := r.currentContactMap(ctx, userID, registeredIDs)
	if err != nil {
		return nil, err
	}
	reverseContactIDs, err := r.model.UserContactsModel.SelectUserReverseContactIdList(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: import contacts reverse list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	reverseContacts := int64Set(reverseContactIDs)

	imported := make([]tg.ImportedContactClazz, 0, len(contacts))
	importedIDs := make([]int64, 0, len(contacts))
	updateIDs := make([]int64, 0, len(contacts))
	importedIDSet := make(map[int64]bool, len(contacts))
	updateIDSet := make(map[int64]bool, len(contacts))

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		for _, contact := range contacts {
			userDO := userByPhone[contact.Phone]
			if userDO == nil {
				if _, _, err := r.model.UnregisteredContactsModel.InsertOrUpdateTx(tx, &model.UnregisteredContacts{
					Phone:           contact.Phone,
					ImporterUserId:  userID,
					ImportFirstName: contact.FirstName,
					ImportLastName:  contact.LastName,
				}); err != nil {
					return fmt.Errorf("insert unregistered contact %s: %w", contact.Phone, err)
				}
				continue
			}

			contactDO := currentContacts[userDO.Id]
			mutual := reverseContacts[userDO.Id]
			if contactDO != nil {
				if _, err := r.model.UserContactsModel.UpdateContactNameTx(tx, contact.FirstName, contact.LastName, userID, userDO.Id); err != nil {
					return fmt.Errorf("update contact name %d/%d: %w", userID, userDO.Id, err)
				}
				if mutual && !contactDO.Mutual {
					if _, err := r.model.UserContactsModel.UpdateMutualTx(tx, true, userID, userDO.Id); err != nil {
						return fmt.Errorf("update contact mutual %d/%d: %w", userID, userDO.Id, err)
					}
				}
			} else {
				if mutual {
					if _, err := r.model.UserContactsModel.UpdateMutualTx(tx, true, userDO.Id, userID); err != nil {
						return fmt.Errorf("update reverse mutual %d/%d: %w", userDO.Id, userID, err)
					}
				} else if _, _, err := r.model.ImportedContactsModel.InsertOrUpdateTx(tx, &model.ImportedContacts{
					UserId:         userDO.Id,
					ImportedUserId: userID,
				}); err != nil {
					return fmt.Errorf("insert imported contact %d/%d: %w", userDO.Id, userID, err)
				}

				if _, _, err := r.model.UserContactsModel.InsertOrUpdateTx(tx, &model.UserContacts{
					OwnerUserId:      userID,
					ContactUserId:    userDO.Id,
					ContactPhone:     contact.Phone,
					ContactFirstName: contact.FirstName,
					ContactLastName:  contact.LastName,
					Mutual:           mutual,
					Date2:            time.Now().Unix(),
				}); err != nil {
					return fmt.Errorf("insert contact %d/%d: %w", userID, userDO.Id, err)
				}
			}

			imported = append(imported, tg.MakeTLImportedContact(&tg.TLImportedContact{
				UserId:   userDO.Id,
				ClientId: contact.ClientId,
			}).ToImportedContact())
			if !importedIDSet[userDO.Id] {
				importedIDs = append(importedIDs, userDO.Id)
				importedIDSet[userDO.Id] = true
			}
			if mutual && !updateIDSet[userDO.Id] {
				updateIDs = append(updateIDs, userDO.Id)
				updateIDSet[userDO.Id] = true
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("%w: import contacts %d: %w", userpb.ErrUserStorage, userID, err)
	}

	for _, id := range importedIDs {
		if err := r.invalidateContactCaches(ctx, userID, id); err != nil {
			return nil, err
		}
	}

	resultContacts, err := r.currentContactMap(ctx, userID, importedIDs)
	if err != nil {
		return nil, err
	}
	users, err := r.importContactUsers(ctx, userID, importedIDs, resultContacts, reverseContacts)
	if err != nil {
		return nil, err
	}
	return userpb.MakeTLUserImportedContacts(&userpb.TLUserImportedContacts{
		Imported:       imported,
		PopularInvites: []tg.PopularContactClazz{},
		RetryContacts:  []int64{},
		Users:          users,
		UpdateIdList:   updateIDs,
	}).ToUserImportedContacts(), nil
}

func (r *Repository) currentContactMap(ctx context.Context, userID int64, contactIDs []int64) (map[int64]*model.UserContacts, error) {
	contacts := make(map[int64]*model.UserContacts, len(contactIDs))
	if len(contactIDs) == 0 {
		return contacts, nil
	}
	contactList, err := r.model.UserContactsModel.SelectListByIdList(ctx, userID, contactIDs)
	if err != nil {
		return nil, fmt.Errorf("%w: import contacts current list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	for i := range contactList {
		contactDO := contactList[i]
		contacts[contactDO.ContactUserId] = &contactDO
	}
	return contacts, nil
}

func (r *Repository) importContactUsers(ctx context.Context, userID int64, importedIDs []int64, currentContacts map[int64]*model.UserContacts, reverseContacts map[int64]bool) ([]tg.UserClazz, error) {
	ids := append([]int64{userID}, importedIDs...)
	userDOList, err := r.model.UsersModel.SelectUsersByIdList(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%w: import contacts users: %w", userpb.ErrUserStorage, err)
	}
	users := make([]tg.UserClazz, 0, len(userDOList))
	for i := range userDOList {
		userDO := userDOList[i]
		contactDO := currentContacts[userDO.Id]
		users = append(users, userFromModel(&userDO, userDO.Id == userID, contactDO != nil, reverseContacts[userDO.Id], contactDO))
	}
	return users, nil
}

func inputContactPhones(contacts []tg.InputContactClazz) []string {
	phones := make([]string, 0, len(contacts))
	seen := make(map[string]bool, len(contacts))
	for _, contact := range contacts {
		if contact == nil || contact.Phone == "" || seen[contact.Phone] {
			continue
		}
		phones = append(phones, contact.Phone)
		seen[contact.Phone] = true
	}
	return phones
}

func int64Set(values []int64) map[int64]bool {
	set := make(map[int64]bool, len(values))
	for _, value := range values {
		set[value] = true
	}
	return set
}

func userFromModel(do *model.Users, self bool, contact bool, mutual bool, contactDO *model.UserContacts) tg.UserClazz {
	accessHash := do.AccessHash
	user := tg.MakeTLUser(&tg.TLUser{
		Self:              self,
		Contact:           contact,
		MutualContact:     mutual,
		Deleted:           do.Deleted,
		Bot:               do.IsBot,
		Verified:          do.Verified,
		Restricted:        do.Restricted,
		Support:           do.Support,
		Scam:              do.Scam,
		Fake:              do.Fake,
		Premium:           do.Premium,
		Id:                do.Id,
		AccessHash:        &accessHash,
		FirstName:         stringPtr(do.FirstName),
		LastName:          stringPtr(do.LastName),
		Username:          stringPtr(do.Username),
		Phone:             stringPtr(do.Phone),
		RestrictionReason: []tg.RestrictionReasonClazz{},
	})
	if contactDO != nil {
		user.CloseFriend = contactDO.CloseFriend
		user.StoriesHidden = contactDO.StoriesHidden
	}
	return user
}
