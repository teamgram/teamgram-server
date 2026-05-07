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

func (r *Repository) CheckUsername(ctx context.Context, username string) (*userpb.UsernameExist, error) {
	if !isValidUsername(username) {
		return nil, userpb.ErrUsernameInvalid
	}
	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if !isNotFound(err) {
			return nil, fmt.Errorf("%w: check username %s: %w", userpb.ErrUserStorage, username, err)
		}
		usernameDO = nil
	}
	if usernameDO == nil {
		return usernameNotExisted(), nil
	}
	return usernameExisted(), nil
}

func (r *Repository) CheckPeerUsername(ctx context.Context, peerType int32, peerID int64, username string) (*userpb.UsernameExist, error) {
	if !isValidUsername(username) {
		return nil, userpb.ErrUsernameInvalid
	}
	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if !isNotFound(err) {
			return nil, fmt.Errorf("%w: check peer username %s: %w", userpb.ErrUserStorage, username, err)
		}
		usernameDO = nil
	}
	if usernameDO == nil {
		return usernameNotExisted(), nil
	}
	if usernameDO.PeerType == peerType && usernameDO.PeerId == peerID {
		return usernameExistedIsMe(), nil
	}
	return usernameExistedNotMe(), nil
}

func (r *Repository) UpdateUsername(ctx context.Context, id int64, username string) error {
	return r.updateUsernameByPeer(ctx, tg.PEER_USER, id, username, true)
}

func (r *Repository) DeleteUsername(ctx context.Context, username string) (bool, error) {
	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if !isNotFound(err) {
			return false, fmt.Errorf("%w: delete username lookup %s: %w", userpb.ErrUserStorage, username, err)
		}
		usernameDO = nil
	}
	rows, err := r.model.UsernameModel.Delete(ctx, username)
	if err != nil {
		return false, fmt.Errorf("%w: delete username %s: %w", userpb.ErrUserStorage, username, err)
	}
	if rows > 0 && usernameDO != nil && usernameDO.PeerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, usernameDO.PeerId)
	}
	return rows > 0, nil
}

func (r *Repository) DeleteUsernameByPeer(ctx context.Context, peerType int32, peerID int64) error {
	rows, err := r.model.UsernameModel.DeleteByPeer(ctx, peerType, peerID)
	if err != nil {
		return fmt.Errorf("%w: delete username by peer %d/%d: %w", userpb.ErrUserStorage, peerType, peerID, err)
	}
	if rows == 0 {
		return userpb.ErrUsernameNotFound
	}
	if peerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, peerID)
	}
	return nil
}

func (r *Repository) GetAccountUsername(ctx context.Context, userID int64) (*userpb.UsernameData, error) {
	return r.getUsernameByPeer(ctx, tg.PEER_USER, userID)
}

func (r *Repository) GetChannelUsername(ctx context.Context, channelID int64) (*userpb.UsernameData, error) {
	return r.getUsernameByPeer(ctx, tg.PEER_CHANNEL, channelID)
}

func (r *Repository) ResolveUsername(ctx context.Context, username string) (*tg.Peer, error) {
	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUsernameNotFound
		}
		return nil, fmt.Errorf("%w: resolve username %s: %w", userpb.ErrUserStorage, username, err)
	}
	if usernameDO == nil {
		return nil, userpb.ErrUsernameNotFound
	}
	peer := peerFromUsername(usernameDO)
	if peer == nil {
		return nil, userpb.ErrUsernameNotFound
	}
	return &tg.Peer{Clazz: peer}, nil
}

func (r *Repository) GetListByUsernameList(ctx context.Context, names []string) ([]userpb.UsernameDataClazz, error) {
	usernameDOList, err := r.model.UsernameModel.SelectList(ctx, names)
	if err != nil {
		return nil, fmt.Errorf("%w: get username list: %w", userpb.ErrUserStorage, err)
	}
	return usernameDataListFromModels(usernameDOList), nil
}

func (r *Repository) SearchUsername(ctx context.Context, q string, excludedContacts []int64, limit int32) ([]userpb.UsernameDataClazz, error) {
	if len(q) < 3 || limit <= 0 {
		return []userpb.UsernameDataClazz{}, nil
	}
	if limit > 50 {
		limit = 50
	}
	if len(excludedContacts) == 0 {
		excludedContacts = []int64{0}
	}

	usernameDOList, err := r.model.UsernameModel.SearchByQueryNotIdList(ctx, q+"%", excludedContacts, limit)
	if err != nil {
		return nil, fmt.Errorf("%w: search username %s: %w", userpb.ErrUserStorage, q, err)
	}
	return usernameDataListFromModels(usernameDOList), nil
}

func (r *Repository) UpdateUsernameByPeer(ctx context.Context, peerType int32, peerID int64, username string) (bool, error) {
	if !isValidUsername(username) {
		return false, userpb.ErrUsernameInvalid
	}
	_, _, err := r.model.UsernameModel.Insert(ctx, &model.Username{
		Username: username,
		PeerType: peerType,
		PeerId:   peerID,
		Editable: true,
		Active:   true,
		Order2:   time.Now().Unix() << 32,
	})
	if err != nil {
		if sqlx.IsDuplicate(err) {
			return false, nil
		}
		return false, fmt.Errorf("%w: update username by peer %d/%d: %w", userpb.ErrUserStorage, peerType, peerID, err)
	}
	if peerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, peerID)
	}
	return true, nil
}

func (r *Repository) updateUsernameByPeer(ctx context.Context, peerType int32, peerID int64, username string, syncUserRow bool) error {
	if peerID == 0 {
		return userpb.ErrUserNotFound
	}
	if !isValidUsername(username) {
		return userpb.ErrUsernameInvalid
	}

	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if !isNotFound(err) {
			return fmt.Errorf("%w: update username lookup %s: %w", userpb.ErrUserStorage, username, err)
		}
		usernameDO = nil
	}
	if usernameDO != nil {
		if usernameDO.PeerType != peerType || usernameDO.PeerId != peerID || !usernameDO.Editable {
			return userpb.ErrUsernameInUse
		}
		if syncUserRow {
			return r.execUserUpdate(ctx, peerID, "update username", func() (int64, error) {
				return r.model.UsersModel.UpdateUsername(ctx, username, peerID)
			})
		}
		return nil
	}

	if err := r.db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModel := r.model.WithTx(tx)
		if _, err := txModel.UsernameModel.DeleteByPeer(peerType, peerID); err != nil {
			return fmt.Errorf("delete old username by peer %d/%d: %w", peerType, peerID, err)
		}
		if _, _, err := txModel.UsernameModel.Insert(&model.Username{
			Username: username,
			PeerType: peerType,
			PeerId:   peerID,
			Editable: true,
			Active:   true,
			Order2:   time.Now().Unix() << 32,
		}); err != nil {
			if sqlx.IsDuplicate(err) {
				return userpb.ErrUsernameInUse
			}
			return fmt.Errorf("insert username %s: %w", username, err)
		}
		if syncUserRow {
			rows, err := txModel.UsersModel.UpdateUsername(username, peerID)
			if err != nil {
				return fmt.Errorf("update user username %d: %w", peerID, err)
			}
			if rows == 0 {
				return userpb.ErrUserNotFound
			}
		}
		return nil
	}); err != nil {
		if err == userpb.ErrUserNotFound || err == userpb.ErrUsernameInUse {
			return err
		}
		return fmt.Errorf("%w: update username by peer %d/%d: %w", userpb.ErrUserStorage, peerType, peerID, err)
	}

	if syncUserRow {
		if err := r.invalidateUserDataCache(ctx, peerID, "invalidate username user cache"); err != nil {
			return err
		}
	} else if peerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, peerID)
	}
	return nil
}

func (r *Repository) ToggleUsername(ctx context.Context, peerType int32, peerID int64, username string, active bool) error {
	if err := r.requireUsernameOwner(ctx, peerType, peerID, username); err != nil {
		return err
	}
	rows, err := r.model.UsernameModel.Update(ctx, map[string]interface{}{"active": active}, username)
	if err != nil {
		return fmt.Errorf("%w: toggle username %s: %w", userpb.ErrUserStorage, username, err)
	}
	if rows == 0 {
		return userpb.ErrUsernameNotFound
	}
	if peerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, peerID)
	}
	return nil
}

func (r *Repository) ReorderUsernames(ctx context.Context, peerType int32, peerID int64, usernames []string) error {
	order := time.Now().Unix() << 32
	for i, username := range usernames {
		if err := r.requireUsernameOwner(ctx, peerType, peerID, username); err != nil {
			return err
		}
		rows, err := r.model.UsernameModel.Update(ctx, map[string]interface{}{"order2": order + int64(i)}, username)
		if err != nil {
			return fmt.Errorf("%w: reorder username %s: %w", userpb.ErrUserStorage, username, err)
		}
		if rows == 0 {
			return userpb.ErrUsernameNotFound
		}
	}
	if peerType == tg.PEER_USER {
		r.invalidateProjectionFactCache(ctx, peerID)
	}
	return nil
}

func (r *Repository) DeactivateAllChannelUsernames(ctx context.Context, channelID int64) error {
	if _, err := r.model.UsernameModel.DeleteByChannelId(ctx, channelID); err != nil {
		return fmt.Errorf("%w: deactivate channel usernames %d: %w", userpb.ErrUserStorage, channelID, err)
	}
	return nil
}

func (r *Repository) requireUsernameOwner(ctx context.Context, peerType int32, peerID int64, username string) error {
	usernameDO, err := r.model.UsernameModel.SelectByUsername(ctx, username)
	if err != nil {
		if isNotFound(err) {
			return userpb.ErrUsernameNotFound
		}
		return fmt.Errorf("%w: get username owner %s: %w", userpb.ErrUserStorage, username, err)
	}
	if usernameDO == nil || usernameDO.PeerType != peerType || usernameDO.PeerId != peerID {
		return userpb.ErrUsernameNotFound
	}
	return nil
}

func isValidUsername(username string) bool {
	if len(username) < 5 || len(username) > 32 {
		return false
	}
	if username[0] == '_' || username[len(username)-1] == '_' {
		return false
	}
	allDigits := true
	prevUnderscore := false
	for _, ch := range username {
		isDigit := ch >= '0' && ch <= '9'
		isLetter := ch >= 'a' && ch <= 'z'
		if !isDigit {
			allDigits = false
		}
		if ch == '_' {
			if prevUnderscore {
				return false
			}
			prevUnderscore = true
			continue
		}
		prevUnderscore = false
		if !isDigit && !isLetter {
			return false
		}
	}
	return !allDigits
}

func (r *Repository) getUsernameByPeer(ctx context.Context, peerType int32, peerID int64) (*userpb.UsernameData, error) {
	usernameDO, err := r.model.UsernameModel.SelectByPeer(ctx, peerType, peerID)
	if err != nil {
		if isNotFound(err) {
			return nil, userpb.ErrUsernameNotFound
		}
		return nil, fmt.Errorf("%w: get username by peer %d/%d: %w", userpb.ErrUserStorage, peerType, peerID, err)
	}
	if usernameDO == nil {
		return nil, userpb.ErrUsernameNotFound
	}
	usernameData := usernameDataFromModel(usernameDO)
	if usernameData == nil {
		return nil, userpb.ErrUsernameNotFound
	}
	return usernameData, nil
}

func usernameDataListFromModels(usernameDOList []model.Username) []userpb.UsernameDataClazz {
	datas := make([]userpb.UsernameDataClazz, 0, len(usernameDOList))
	for i := range usernameDOList {
		if data := usernameDataFromModel(&usernameDOList[i]); data != nil {
			datas = append(datas, data)
		}
	}
	return datas
}

func usernameDataFromModel(do *model.Username) userpb.UsernameDataClazz {
	if do == nil {
		return nil
	}
	peer := peerFromUsername(do)
	if peer == nil {
		return nil
	}
	return userpb.MakeTLUsernameData(&userpb.TLUsernameData{
		Username: do.Username,
		Peer:     peer,
		Editable: do.Editable,
		Active:   do.Active,
	}).ToUsernameData()
}

func peerFromUsername(do *model.Username) tg.PeerClazz {
	switch do.PeerType {
	case tg.PEER_USER:
		return tg.MakePeerUser(do.PeerId)
	case tg.PEER_CHANNEL:
		return tg.MakePeerChannel(do.PeerId)
	default:
		return nil
	}
}

func usernameNotExisted() *userpb.UsernameExist {
	return userpb.MakeTLUsernameNotExisted(&userpb.TLUsernameNotExisted{}).ToUsernameExist()
}

func usernameExisted() *userpb.UsernameExist {
	return userpb.MakeTLUsernameExisted(&userpb.TLUsernameExisted{}).ToUsernameExist()
}

func usernameExistedNotMe() *userpb.UsernameExist {
	return userpb.MakeTLUsernameExistedNotMe(&userpb.TLUsernameExistedNotMe{}).ToUsernameExist()
}

func usernameExistedIsMe() *userpb.UsernameExist {
	return userpb.MakeTLUsernameExistedIsMe(&userpb.TLUsernameExistedIsMe{}).ToUsernameExist()
}
