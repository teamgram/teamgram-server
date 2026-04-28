package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func (r *Repository) UpdateProfilePhoto(ctx context.Context, userID, photoID int64) (int64, error) {
	if userID == 0 {
		return 0, userpb.ErrUserNotFound
	}

	mainPhotoID := photoID
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		if photoID == 0 {
			currentPhotoID, err := r.model.UsersModel.SelectProfilePhotoTx(tx, userID)
			if err != nil {
				return fmt.Errorf("select current profile photo: %w", err)
			}
			if currentPhotoID > 0 {
				nextPhotoID, err := r.model.UserProfilePhotosModel.SelectNextTx(tx, userID, []int64{currentPhotoID})
				if err != nil {
					return fmt.Errorf("select next profile photo: %w", err)
				}
				if _, err := r.model.UserProfilePhotosModel.DeleteTx(tx, userID, []int64{currentPhotoID}); err != nil {
					return fmt.Errorf("delete current profile photo: %w", err)
				}
				mainPhotoID = nextPhotoID
			}
		} else {
			if _, _, err := r.model.UserProfilePhotosModel.InsertOrUpdateTx(tx, &model.UserProfilePhotos{
				UserId:  userID,
				PhotoId: photoID,
				Date2:   time.Now().Unix(),
			}); err != nil {
				return fmt.Errorf("insert profile photo: %w", err)
			}
		}

		rows, err := r.model.UsersModel.UpdateProfilePhotoTx(tx, mainPhotoID, userID)
		if err != nil {
			return fmt.Errorf("update user profile photo: %w", err)
		}
		if rows == 0 {
			return userpb.ErrUserNotFound
		}
		return nil
	}); err != nil {
		if err == userpb.ErrUserNotFound {
			return 0, err
		}
		return 0, fmt.Errorf("%w: update profile photo %d: %w", userpb.ErrUserStorage, userID, err)
	}

	if err := r.DelCache(ctx, userDataCacheKey(userID)); err != nil {
		return 0, fmt.Errorf("%w: invalidate user cache %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return mainPhotoID, nil
}

func (r *Repository) GetProfilePhotos(ctx context.Context, userID int64) ([]int64, error) {
	if userID == 0 {
		return nil, userpb.ErrUserNotFound
	}
	photoIDs, err := r.model.UserProfilePhotosModel.SelectList(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get profile photos %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return photoIDs, nil
}

func (r *Repository) DeleteProfilePhotos(ctx context.Context, userID int64, photoIDs []int64) (int64, error) {
	if userID == 0 {
		return 0, userpb.ErrUserNotFound
	}
	if len(photoIDs) == 0 {
		return 0, nil
	}

	nextMainPhotoID := int64(0)
	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		mainPhotoID, err := r.model.UsersModel.SelectProfilePhotoTx(tx, userID)
		if err != nil {
			return fmt.Errorf("select profile photo: %w", err)
		}

		nextMainPhotoID = mainPhotoID
		if containsInt64(photoIDs, mainPhotoID) {
			nextPhotoID, err := r.model.UserProfilePhotosModel.SelectNextTx(tx, userID, photoIDs)
			if err != nil {
				return fmt.Errorf("select next profile photo: %w", err)
			}
			nextMainPhotoID = nextPhotoID
		}

		if _, err := r.model.UserProfilePhotosModel.DeleteTx(tx, userID, photoIDs); err != nil {
			return fmt.Errorf("delete profile photos: %w", err)
		}
		if nextMainPhotoID != mainPhotoID {
			rows, err := r.model.UsersModel.UpdateProfilePhotoTx(tx, nextMainPhotoID, userID)
			if err != nil {
				return fmt.Errorf("update user profile photo: %w", err)
			}
			if rows == 0 {
				return userpb.ErrUserNotFound
			}
		}
		return nil
	}); err != nil {
		if err == userpb.ErrUserNotFound {
			return 0, err
		}
		return 0, fmt.Errorf("%w: delete profile photos %d: %w", userpb.ErrUserStorage, userID, err)
	}

	if err := r.DelCache(ctx, userDataCacheKey(userID)); err != nil {
		return 0, fmt.Errorf("%w: invalidate user cache %d: %w", userpb.ErrUserStorage, userID, err)
	}
	return nextMainPhotoID, nil
}

func containsInt64(values []int64, needle int64) bool {
	for _, value := range values {
		if value == needle {
			return true
		}
	}
	return false
}
