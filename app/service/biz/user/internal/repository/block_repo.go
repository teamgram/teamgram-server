package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) BlockPeer(ctx context.Context, userID int64, peerType int32, peerID int64) error {
	_, _, err := r.model.UserPeerBlocksModel.InsertOrUpdate(ctx, &model.UserPeerBlocks{
		UserId:   userID,
		PeerType: peerType,
		PeerId:   peerID,
		Date:     time.Now().Unix(),
	})
	if err != nil {
		return fmt.Errorf("%w: block peer %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return nil
}

func (r *Repository) UnblockPeer(ctx context.Context, userID int64, peerType int32, peerID int64) error {
	_, err := r.model.UserPeerBlocksModel.Delete(ctx, userID, peerType, peerID)
	if err != nil {
		return fmt.Errorf("%w: unblock peer %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return nil
}

func (r *Repository) IsPeerBlocked(ctx context.Context, userID int64, peerType int32, peerID int64) (bool, error) {
	blockedDO, err := r.model.UserPeerBlocksModel.Select(ctx, userID, peerType, peerID)
	if err != nil {
		if isNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("%w: check blocked peer %d/%d/%d: %w", userpb.ErrUserStorage, userID, peerType, peerID, err)
	}
	return blockedDO != nil, nil
}

func (r *Repository) CheckBlockedUserList(ctx context.Context, userID int64, idList []int64) (*userpb.VectorLong, error) {
	blockedIDs, err := r.model.UserPeerBlocksModel.SelectListByIdList(ctx, userID, idList)
	if err != nil {
		return nil, fmt.Errorf("%w: check blocked user list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	if blockedIDs == nil {
		blockedIDs = []int64{}
	}
	return &userpb.VectorLong{Datas: blockedIDs}, nil
}

func (r *Repository) GetBlockedList(ctx context.Context, userID int64, offset, limit int32) (*userpb.VectorPeerBlocked, error) {
	blockedList, err := r.model.UserPeerBlocksModel.SelectListOffset(ctx, userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("%w: get blocked list %d: %w", userpb.ErrUserStorage, userID, err)
	}
	datas := make([]tg.PeerBlockedClazz, 0, len(blockedList))
	for i := range blockedList {
		blocked, ok := makePeerBlocked(&blockedList[i])
		if ok {
			datas = append(datas, blocked)
		}
	}
	return &userpb.VectorPeerBlocked{Datas: datas}, nil
}

func makePeerBlocked(blockedDO *model.UserPeerBlocks) (tg.PeerBlockedClazz, bool) {
	if blockedDO == nil {
		return nil, false
	}
	var peer tg.PeerClazz
	switch blockedDO.PeerType {
	case tg.PEER_USER:
		peer = tg.MakePeerUser(blockedDO.PeerId)
	case tg.PEER_CHAT:
		peer = tg.MakePeerChat(blockedDO.PeerId)
	case tg.PEER_CHANNEL:
		peer = tg.MakePeerChannel(blockedDO.PeerId)
	default:
		return nil, false
	}
	return tg.MakeTLPeerBlocked(&tg.TLPeerBlocked{
		PeerId: peer,
		Date:   int32(blockedDO.Date),
	}).ToPeerBlocked(), true
}
