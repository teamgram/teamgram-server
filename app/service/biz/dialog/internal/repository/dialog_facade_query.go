package repository

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

type DialogExtrasRecord struct {
	PeerType             int32
	PeerID               int64
	FolderID             int32
	MainPinnedOrder      int64
	FolderPinnedOrder    int64
	DraftPayload         []byte
	PrivateTTLPeriod     int32
	PrivateThemeEmoticon string
	WallpaperID          int64
	WallpaperOverridden  bool
}

func (r *Repository) BatchGetDialogExtras(ctx context.Context, userID int64, peers []PeerRef) ([]DialogExtrasRecord, error) {
	if len(peers) == 0 {
		return []DialogExtrasRecord{}, nil
	}
	if r == nil || r.model == nil {
		return nil, wrapReadStorage("batch get dialog extras", errors.New("dialog repository models not initialized"))
	}

	out := make([]DialogExtrasRecord, 0, len(peers))
	for _, peer := range peers {
		extras := DialogExtrasRecord{
			PeerType: peer.PeerType,
			PeerID:   peer.PeerID,
		}
		if pref, err := r.model.DialogPreferencesModel.SelectByUserPeer(ctx, userID, peer.PeerType, peer.PeerID); err == nil {
			extras.FolderID = pref.FolderId
			extras.MainPinnedOrder = pref.MainPinnedOrder
			extras.FolderPinnedOrder = pref.FolderPinnedOrder
		} else if !errors.Is(err, model.ErrNotFound) {
			return nil, wrapReadStorage("select dialog extras preferences", err)
		}
		if draft, err := r.model.DialogDraftsModel.SelectByUserPeer(ctx, userID, peer.PeerType, peer.PeerID); err == nil {
			extras.DraftPayload = draft.DraftPayload
		} else if !errors.Is(err, model.ErrNotFound) {
			return nil, wrapReadStorage("select dialog extras draft", err)
		}
		if visual, err := r.model.DialogVisualSettingsModel.SelectByUserPeer(ctx, userID, peer.PeerType, peer.PeerID); err == nil {
			extras.WallpaperID = visual.WallpaperId
			extras.WallpaperOverridden = visual.WallpaperOverridden
		} else if !errors.Is(err, model.ErrNotFound) {
			return nil, wrapReadStorage("select dialog extras visual settings", err)
		}
		if peer.PeerType == PeerTypeUser {
			scope, err := MakePrivatePairScope(userID, peer.PeerID)
			if err != nil {
				return nil, wrapReadStorage("make dialog extras private scope", err)
			}
			if policy, err := r.model.DialogPeerPolicyModel.SelectByScope(ctx, scope.ScopeType, scope.ScopeID); err == nil {
				extras.PrivateTTLPeriod = policy.TtlPeriod
				extras.PrivateThemeEmoticon = policy.ThemeEmoticon
			} else if !errors.Is(err, model.ErrNotFound) {
				return nil, wrapReadStorage("select dialog extras peer policy", err)
			}
		}
		out = append(out, extras)
	}
	return out, nil
}
