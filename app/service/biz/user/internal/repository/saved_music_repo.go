package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
)

func (r *Repository) SaveMusic(ctx context.Context, userID, musicID int64, unsave bool) error {
	if userID == 0 {
		return userpb.ErrUserNotFound
	}

	var nextID int64
	var shouldUpdateUser bool
	if unsave {
		musicDOList, err := r.model.UserSavedMusicModel.SelectList(ctx, userID)
		if err != nil {
			return fmt.Errorf("%w: select saved music %d: %w", userpb.ErrUserStorage, userID, err)
		}

		if _, err := r.model.UserSavedMusicModel.Delete(ctx, userID, musicID); err != nil {
			return fmt.Errorf("%w: unsave music %d/%d: %w", userpb.ErrUserStorage, userID, musicID, err)
		}

		for i := range musicDOList {
			if musicDOList[i].SavedMusicId != musicID {
				continue
			}
			shouldUpdateUser = true
			if i == 0 {
				if len(musicDOList) > 1 {
					nextID = musicDOList[1].SavedMusicId
				}
			} else {
				nextID = musicDOList[0].SavedMusicId
			}
			break
		}
	} else {
		if _, _, err := r.model.UserSavedMusicModel.InsertOrUpdate(ctx, &model.UserSavedMusic{
			UserId:       userID,
			SavedMusicId: musicID,
			Order2:       time.Now().Unix() << 32,
		}); err != nil {
			return fmt.Errorf("%w: save music %d/%d: %w", userpb.ErrUserStorage, userID, musicID, err)
		}
		nextID = musicID
		shouldUpdateUser = true
	}

	if shouldUpdateUser {
		rows, err := r.model.UsersModel.UpdateSavedMusicId(ctx, nextID, userID)
		if err != nil {
			return fmt.Errorf("%w: update saved music id %d: %w", userpb.ErrUserStorage, userID, err)
		}
		if rows == 0 {
			return userpb.ErrUserNotFound
		}
		if err := r.DelCache(ctx, userDataCacheKey(userID)); err != nil {
			return fmt.Errorf("%w: invalidate user cache %d: %w", userpb.ErrUserStorage, userID, err)
		}
	}

	return nil
}

func (r *Repository) GetSavedMusicIDList(ctx context.Context, userID int64) ([]int64, error) {
	if userID == 0 {
		return nil, userpb.ErrUserNotFound
	}
	musicDOList, err := r.model.UserSavedMusicModel.SelectList(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: get saved music list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	ids := make([]int64, 0, len(musicDOList))
	for i := range musicDOList {
		ids = append(ids, musicDOList[i].SavedMusicId)
	}
	return ids, nil
}
