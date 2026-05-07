package repository

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func (r *Repository) ToggleDialogPin(ctx context.Context, in ToggleDialogPinInput) (*PreferenceMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	peerDialogID, err := MakePeerDialogID(in.PeerType, in.PeerID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	scope := PreferenceScopeMainPin
	if in.FolderID != 0 {
		scope = PreferenceScopeFolderPin
	}
	pinOrder := int64(0)
	if in.Pinned {
		pinOrder = in.PinOrder
		if pinOrder <= 0 {
			pinOrder = time.Now().UnixNano()
		}
	}
	var version int64
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		if in.FolderID == 0 {
			if err := upsertMainPinPreference(txModels, in.UserID, in.PeerType, in.PeerID, peerDialogID, pinOrder); err != nil {
				return err
			}
		} else {
			if err := upsertFolderPinPreference(txModels, in.UserID, in.PeerType, in.PeerID, peerDialogID, in.FolderID, pinOrder); err != nil {
				return err
			}
		}
		if err := incrementPreferenceVersion(txModels, in.UserID, scope, in.FolderID); err != nil {
			return err
		}
		row, err := txModels.DialogPreferenceVersionsModel.SelectOne(in.UserID, scope, in.FolderID)
		if err != nil {
			return storageError("select preference version", err)
		}
		version = row.AggregateVersion
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &PreferenceMutationResult{UserID: in.UserID, PeerDialogID: peerDialogID, AggregateVersion: version}, nil
}

func (r *Repository) ReorderPinnedDialogs(ctx context.Context, in ReorderPinnedDialogsInput) (*PreferenceMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	scope := PreferenceScopeMainPin
	if in.FolderID != 0 {
		scope = PreferenceScopeFolderPin
	}
	var version int64
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		peerDialogIDs := make([]int64, 0, len(in.PeerOrder))
		for i, peer := range in.PeerOrder {
			peerDialogID, err := MakePeerDialogID(peer.PeerType, peer.PeerID)
			if err != nil {
				return fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
			}
			peerDialogIDs = append(peerDialogIDs, peerDialogID)
			if in.FolderID == 0 {
				if err := upsertMainPinPreference(txModels, in.UserID, peer.PeerType, peer.PeerID, peerDialogID, int64(i+1)); err != nil {
					return err
				}
			} else {
				if err := upsertFolderPinPreference(txModels, in.UserID, peer.PeerType, peer.PeerID, peerDialogID, in.FolderID, int64(i+1)); err != nil {
					return err
				}
			}
		}
		if err := clearOmittedPinnedPreferences(txModels, in.UserID, in.FolderID, peerDialogIDs); err != nil {
			return err
		}
		if err := incrementPreferenceVersion(txModels, in.UserID, scope, in.FolderID); err != nil {
			return err
		}
		row, err := txModels.DialogPreferenceVersionsModel.SelectOne(in.UserID, scope, in.FolderID)
		if err != nil {
			return storageError("select preference version", err)
		}
		version = row.AggregateVersion
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &PreferenceMutationResult{UserID: in.UserID, AggregateVersion: version}, nil
}

func (r *Repository) EditPeerFolders(ctx context.Context, in EditPeerFoldersInput) (*PreferenceMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	peerDialogID, err := MakePeerDialogID(in.PeerType, in.PeerID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		oldFolderID := in.OldFolderID
		if row, err := txModels.DialogPreferencesModel.SelectByUserPeer(in.UserID, in.PeerType, in.PeerID); err == nil {
			oldFolderID = row.FolderId
		} else if !errors.Is(err, model.ErrNotFound) {
			return storageError("select old folder preference", err)
		}
		duplicate, err := insertPublicUpdateOutbox(txModels, publicUpdateOutboxInput{
			OutboxID:            in.OutboxID,
			SourceUserID:        in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			TargetUserID:        in.UserID,
			TargetAuthPolicy:    TargetAuthPolicyNotSourcePermAuthKey,
			OperationID:         in.OperationID,
			DeliveryPath:        "userupdates_pts",
			PublicUpdateType:    in.PublicUpdateType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		if err := upsertFolderMembershipPreference(txModels, in.UserID, in.PeerType, in.PeerID, peerDialogID, in.NewFolderID); err != nil {
			return err
		}
		if err := incrementPreferenceVersion(txModels, in.UserID, PreferenceScopeFolderMembership, in.NewFolderID); err != nil {
			return err
		}
		if oldFolderID != in.NewFolderID {
			if err := incrementPreferenceVersion(txModels, in.UserID, PreferenceScopeFolderMembership, oldFolderID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &PreferenceMutationResult{UserID: in.UserID, PeerDialogID: peerDialogID}, nil
}

func (r *Repository) SaveDraft(ctx context.Context, in SaveDraftInput) (*DraftMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	peerDialogID, err := MakePeerDialogID(in.PeerType, in.PeerID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	entitiesPayload := in.EntitiesPayload
	if entitiesPayload == nil {
		entitiesPayload = []byte{}
	}
	draftPayload := in.DraftPayload
	if draftPayload == nil {
		draftPayload = []byte{}
	}
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             draftPayload,
		})
		if err != nil || duplicate {
			return err
		}
		_, _, err = txModels.DialogDraftsModel.InsertOrUpdate(&model.DialogDrafts{
			UserId:                    in.UserID,
			PeerType:                  in.PeerType,
			PeerId:                    in.PeerID,
			PeerDialogId:              peerDialogID,
			DraftKind:                 in.DraftKind,
			Message:                   in.Message,
			EntitiesPayload:           entitiesPayload,
			ReplyToPeerSeq:            in.ReplyToPeerSeq,
			DraftPayloadSchemaVersion: 1,
			DraftPayload:              draftPayload,
			Date:                      unixOrZero(in.Date),
		})
		if err != nil {
			return storageError("save draft", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &DraftMutationResult{UserID: in.UserID, PeerDialogID: peerDialogID}, nil
}

func (r *Repository) ClearDraft(ctx context.Context, in ClearDraftInput) (*DraftMutationResult, error) {
	return r.ClearDraftAfterSend(ctx, ClearDraftAfterSendInput{
		UserID:              in.UserID,
		PeerType:            in.PeerType,
		PeerID:              in.PeerID,
		ClearBeforeDate:     unixNow() + int64((100*365*24*time.Hour)/time.Second),
		SourcePermAuthKeyID: in.SourcePermAuthKeyID,
		OperationID:         in.OperationID,
		OutboxID:            in.OutboxID,
		EventType:           in.EventType,
		Payload:             []byte(`{"schema_version":1}`),
	})
}

func (r *Repository) ClearDraftAfterSend(ctx context.Context, in ClearDraftAfterSendInput) (*DraftMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	peerDialogID, err := MakePeerDialogID(in.PeerType, in.PeerID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var affected int64
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		outbox := authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             in.Payload,
		}
		duplicate, err := authSeqOutboxDuplicateExists(txModels, outbox)
		if err != nil || duplicate {
			return err
		}
		if in.ClearBeforeDate > 0 {
			affected, err = txModels.DialogDraftsModel.ClearByUserPeerBeforeDate(
				[]byte{}, 1, []byte(`{"schema_version":1}`), 0, in.UserID, peerDialogID, in.ClearBeforeDate,
			)
			if err != nil {
				return storageError("clear draft after send", err)
			}
		}
		if affected == 0 {
			return nil
		}
		_, err = insertAuthSeqOutbox(txModels, outbox)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &DraftMutationResult{UserID: in.UserID, PeerDialogID: peerDialogID, Cleared: affected > 0}, nil
}

func (r *Repository) ClearAllDrafts(ctx context.Context, in ClearAllDraftsInput) ([]DraftMutationResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	if in.OperationID == "" {
		return nil, dialogpb.ErrOutboxUnavailable
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	var cleared []DraftMutationResult
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		drafts, err := txModels.DialogDraftsModel.SelectActiveByUser(in.UserID)
		if err != nil {
			return storageError("select active drafts", err)
		}
		if len(in.OutboxIDs) < len(drafts) {
			return dialogpb.ErrOutboxUnavailable
		}
		cleared = make([]DraftMutationResult, 0, len(drafts))
		for i := range drafts {
			draft := drafts[i]
			outbox := authSeqOutboxInput{
				OutboxID:            in.OutboxIDs[i],
				UserID:              in.UserID,
				SourcePermAuthKeyID: in.SourcePermAuthKeyID,
				OperationID:         fmt.Sprintf("%s:peer:%d", in.OperationID, draft.PeerDialogId),
				EventType:           "dialog.draftCleared",
				PeerType:            draft.PeerType,
				PeerID:              draft.PeerId,
				Payload:             []byte(`{"schema_version":1}`),
			}
			duplicate, err := authSeqOutboxDuplicateExists(txModels, outbox)
			if err != nil {
				return err
			}
			if duplicate {
				continue
			}
			affected, err := txModels.DialogDraftsModel.ClearByUserPeerBeforeDate(
				[]byte{}, 1, []byte(`{"schema_version":1}`), 0, in.UserID, draft.PeerDialogId, draft.Date,
			)
			if err != nil {
				return storageError("clear all drafts", err)
			}
			if affected == 0 {
				continue
			}
			if _, err := insertAuthSeqOutbox(txModels, outbox); err != nil {
				return err
			}
			cleared = append(cleared, DraftMutationResult{
				UserID:       in.UserID,
				PeerDialogID: draft.PeerDialogId,
				Cleared:      true,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cleared, nil
}

func (r *Repository) ListActiveDrafts(ctx context.Context, userID int64) ([]DraftRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	rows, err := models.DialogDraftsModel.SelectActiveByUser(ctx, userID)
	if err != nil {
		return nil, storageError("select active drafts", err)
	}
	out := make([]DraftRecord, 0, len(rows))
	for i := range rows {
		row := rows[i]
		out = append(out, DraftRecord{
			UserID:          row.UserId,
			PeerType:        row.PeerType,
			PeerID:          row.PeerId,
			PeerDialogID:    row.PeerDialogId,
			DraftKind:       row.DraftKind,
			Message:         row.Message,
			EntitiesPayload: row.EntitiesPayload,
			ReplyToPeerSeq:  row.ReplyToPeerSeq,
			DraftPayload:    row.DraftPayload,
			Date:            unixOrZero(row.Date),
		})
	}
	return out, nil
}

func (r *Repository) UpsertSavedDialogFromMessage(ctx context.Context, in SavedDialogTopInput) error {
	_, err := MakePeerDialogID(in.PeerType, in.PeerID)
	if err != nil {
		return fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	models, err := r.requireModels()
	if err != nil {
		return err
	}
	_, _, err = models.SavedDialogsModel.UpsertTopFromMessage(ctx, &model.SavedDialogs{
		UserId:                in.UserID,
		PeerType:              in.PeerType,
		PeerId:                in.PeerID,
		TopPeerSeq:            in.TopPeerSeq,
		TopCanonicalMessageId: in.TopCanonicalMessageID,
		TopMessageDate:        unixOrZero(in.TopMessageDate),
		SavedSchemaVersion:    1,
		SavedPayload:          in.Payload,
	})
	if err != nil {
		return storageError("upsert saved dialog", err)
	}
	return nil
}

func (r *Repository) ListSavedDialogs(ctx context.Context, userID int64, excludePinned bool, offsetDate int64, limit int32) ([]SavedDialogRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	if limit <= 0 {
		limit = 100
	}
	if offsetDate <= 0 {
		offsetDate = 253402300799
	}
	if !excludePinned {
		rows, err := models.SavedDialogsModel.SelectDialogs(ctx, userID, offsetDate, limit)
		if err != nil {
			return nil, storageError("select saved dialogs", err)
		}
		return makeSavedDialogRecords(rows)
	}
	rows, err := models.SavedDialogsModel.SelectUnpinnedDialogs(ctx, userID, offsetDate, limit)
	if err != nil {
		return nil, storageError("select unpinned saved dialogs", err)
	}
	return makeSavedDialogRecords(rows)
}

func (r *Repository) ListPinnedSavedDialogs(ctx context.Context, userID int64) ([]SavedDialogRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	rows, err := models.SavedDialogsModel.SelectPinnedDialogs(ctx, userID)
	if err != nil {
		return nil, storageError("select pinned saved dialogs", err)
	}
	return makeSavedDialogRecords(rows)
}

func (r *Repository) ToggleSavedDialogPin(ctx context.Context, in SavedDialogPinInput) error {
	if in.SourcePermAuthKeyID == 0 {
		return dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	return db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		pinOrder := in.PinOrder
		if !in.Pinned {
			pinOrder = 0
		} else {
			if pinOrder <= 0 {
				row, err := txModels.DialogRepositoryQueries.SelectSavedDialogNextPinOrder(in.UserID)
				if err != nil {
					return storageError("select next saved dialog pin order", err)
				}
				pinOrder = row.NextPinOrder
			}
			if _, err := txModels.SavedDialogsModel.ClearDuplicatePinOrder(in.UserID, pinOrder, in.PeerType, in.PeerID); err != nil {
				return storageError("clear duplicate saved dialog pin order", err)
			}
		}
		if _, err := txModels.SavedDialogsModel.UpdateUserPeerPinned(in.Pinned, pinOrder, in.UserID, in.PeerType, in.PeerID); err != nil {
			return storageError("toggle saved dialog pin", err)
		}
		return nil
	})
}

func makeSavedDialogRecords(rows []model.SavedDialogs) ([]SavedDialogRecord, error) {
	out := make([]SavedDialogRecord, 0, len(rows))
	for i := range rows {
		row := rows[i]
		out = append(out, SavedDialogRecord{
			UserID:                row.UserId,
			PeerType:              row.PeerType,
			PeerID:                row.PeerId,
			TopPeerSeq:            row.TopPeerSeq,
			TopCanonicalMessageID: row.TopCanonicalMessageId,
			TopMessageDate:        unixOrZero(row.TopMessageDate),
			Pinned:                row.Pinned,
			PinOrder:              row.PinOrder,
			SavedPayload:          row.SavedPayload,
		})
	}
	return out, nil
}

func (r *Repository) ReorderPinnedSavedDialogs(ctx context.Context, in ReorderPinnedSavedDialogsInput) error {
	if in.SourcePermAuthKeyID == 0 {
		return dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	return db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		if _, err := txModels.SavedDialogsModel.UpdateUserUnPinned(in.UserID); err != nil {
			return storageError("clear saved dialog pins", err)
		}
		for i, peer := range in.Order {
			if _, err := txModels.SavedDialogsModel.UpdateUserPeerPinned(true, int64(i+1), in.UserID, peer.PeerType, peer.PeerID); err != nil {
				return storageError("reorder saved dialog pin", err)
			}
		}
		return nil
	})
}

func (r *Repository) SetPrivatePeerPolicy(ctx context.Context, in PrivatePeerPolicyInput) (*PrivatePeerPolicyResult, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	scope, err := MakePrivatePairScope(in.UserID, in.PeerUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dialogpb.ErrInvalidPeer, err)
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicateCount := 0
		for _, target := range []struct {
			userID   int64
			outboxID int64
			policy   string
		}{
			{userID: in.UserID, outboxID: in.ActorOutboxID, policy: TargetAuthPolicyNotSourcePermAuthKey},
			{userID: in.PeerUserID, outboxID: in.PeerOutboxID, policy: TargetAuthPolicyAll},
		} {
			duplicate, err := insertPublicUpdateOutbox(txModels, publicUpdateOutboxInput{
				OutboxID:            target.outboxID,
				SourceUserID:        in.UserID,
				SourcePermAuthKeyID: in.SourcePermAuthKeyID,
				TargetUserID:        target.userID,
				TargetAuthPolicy:    target.policy,
				OperationID:         fmt.Sprintf("%s:target:%d", in.OperationID, target.userID),
				DeliveryPath:        in.DeliveryPath,
				PublicUpdateType:    in.PublicUpdateType,
				PeerType:            PeerTypeUser,
				PeerID:              in.PeerUserID,
				Payload:             in.Payload,
			})
			if err != nil {
				return err
			}
			if duplicate {
				duplicateCount++
			}
		}
		if duplicateCount == 2 {
			return nil
		}
		if _, _, err := txModels.DialogPeerPolicyModel.Upsert(&model.DialogPeerPolicy{
			ScopeType:     scope.ScopeType,
			ScopeId:       scope.ScopeID,
			PeerType:      scope.PeerType,
			PeerId:        scope.PeerID,
			TtlPeriod:     in.TTLPeriod,
			ThemeEmoticon: in.ThemeEmoticon,
			PolicyVersion: 1,
		}); err != nil {
			return storageError("upsert private peer policy", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &PrivatePeerPolicyResult{Scope: scope}, nil
}

func (r *Repository) SetPeerWallpaper(ctx context.Context, in PeerWallpaperInput) error {
	if in.SourcePermAuthKeyID == 0 {
		return dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	return db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		duplicate, err := insertAuthSeqOutbox(txModels, authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            in.PeerType,
			PeerID:              in.PeerID,
			Payload:             in.Payload,
		})
		if err != nil || duplicate {
			return err
		}
		if _, _, err := txModels.DialogVisualSettingsModel.Upsert(&model.DialogVisualSettings{
			UserId:              in.UserID,
			PeerType:            in.PeerType,
			PeerId:              in.PeerID,
			WallpaperId:         in.WallpaperID,
			WallpaperOverridden: in.WallpaperOverridden,
			VisualVersion:       1,
		}); err != nil {
			return storageError("set peer wallpaper", err)
		}
		return nil
	})
}

func upsertPreference(txModels *model.TxModels, userID int64, peerType int32, peerID int64, peerDialogID int64, folderID int32, mainPinOrder int64, folderPinOrder int64) error {
	_, _, err := txModels.DialogPreferencesModel.InsertOrUpdate(&model.DialogPreferences{
		UserId:             userID,
		PeerType:           peerType,
		PeerId:             peerID,
		PeerDialogId:       peerDialogID,
		FolderId:           folderID,
		MainPinnedOrder:    mainPinOrder,
		FolderPinnedOrder:  folderPinOrder,
		PreferencesVersion: 1,
	})
	if err != nil {
		return storageError("upsert dialog preference", err)
	}
	return nil
}

func upsertMainPinPreference(txModels *model.TxModels, userID int64, peerType int32, peerID int64, peerDialogID int64, pinOrder int64) error {
	_, _, err := txModels.DialogPreferencesModel.UpsertMainPin(&model.DialogPreferences{
		UserId:             userID,
		PeerType:           peerType,
		PeerId:             peerID,
		PeerDialogId:       peerDialogID,
		MainPinnedOrder:    pinOrder,
		PreferencesVersion: 1,
	})
	if err != nil {
		return storageError("upsert dialog main pin preference", err)
	}
	return nil
}

func upsertFolderPinPreference(txModels *model.TxModels, userID int64, peerType int32, peerID int64, peerDialogID int64, folderID int32, pinOrder int64) error {
	_, _, err := txModels.DialogPreferencesModel.UpsertFolderPin(&model.DialogPreferences{
		UserId:             userID,
		PeerType:           peerType,
		PeerId:             peerID,
		PeerDialogId:       peerDialogID,
		FolderId:           folderID,
		FolderPinnedOrder:  pinOrder,
		PreferencesVersion: 1,
	})
	if err != nil {
		return storageError("upsert dialog folder pin preference", err)
	}
	return nil
}

func upsertFolderMembershipPreference(txModels *model.TxModels, userID int64, peerType int32, peerID int64, peerDialogID int64, folderID int32) error {
	_, _, err := txModels.DialogPreferencesModel.UpsertFolderMembership(&model.DialogPreferences{
		UserId:             userID,
		PeerType:           peerType,
		PeerId:             peerID,
		PeerDialogId:       peerDialogID,
		FolderId:           folderID,
		PreferencesVersion: 1,
	})
	if err != nil {
		return storageError("upsert dialog folder membership preference", err)
	}
	return nil
}

func incrementPreferenceVersion(txModels *model.TxModels, userID int64, scope string, folderID int32) error {
	_, _, err := txModels.DialogPreferenceVersionsModel.Increment(&model.DialogPreferenceVersions{
		UserId:           userID,
		ScopeKind:        scope,
		FolderId:         folderID,
		AggregateVersion: 1,
	})
	if err != nil {
		return storageError("increment preference version", err)
	}
	return nil
}

type authSeqOutboxInput struct {
	OutboxID            int64
	UserID              int64
	SourcePermAuthKeyID int64
	OperationID         string
	EventType           string
	PeerType            int32
	PeerID              int64
	Payload             []byte
}

func insertAuthSeqOutbox(txModels *model.TxModels, in authSeqOutboxInput) (bool, error) {
	if in.OutboxID == 0 || in.OperationID == "" {
		return false, dialogpb.ErrOutboxUnavailable
	}
	if in.SourcePermAuthKeyID == 0 {
		return false, dialogpb.ErrSourceAuthKeyRequired
	}
	if len(in.Payload) == 0 {
		in.Payload = []byte(`{"schema_version":1}`)
	}
	row := &model.DialogAuthSeqOutbox{
		OutboxId:             in.OutboxID,
		UserId:               in.UserID,
		SourcePermAuthKeyId:  in.SourcePermAuthKeyID,
		TargetAuthPolicy:     TargetAuthPolicyNotSourcePermAuthKey,
		OperationId:          in.OperationID,
		EventType:            in.EventType,
		PeerType:             in.PeerType,
		PeerId:               in.PeerID,
		PayloadSchemaVersion: 1,
		Payload:              in.Payload,
		PayloadHash:          hashPayload(in.Payload),
		Status:               OutboxStatusPending,
		AttemptCount:         0,
		NextRetryAt:          unixNow(),
		LeaseOwner:           "",
		LeaseUntil:           0,
		LastErrorKind:        "",
		LastErrorMessage:     "",
	}
	_, rowsAffected, err := txModels.DialogAuthSeqOutboxModel.InsertIgnore(row)
	if err != nil {
		return false, storageError("insert dialog auth seq outbox", err)
	}
	if rowsAffected == 0 {
		existing, err := txModels.DialogAuthSeqOutboxModel.SelectByUserOperation(in.UserID, in.OperationID)
		if err != nil {
			return false, storageError("select dialog auth seq outbox duplicate", err)
		}
		if existing.SourcePermAuthKeyId != row.SourcePermAuthKeyId ||
			existing.TargetAuthPolicy != row.TargetAuthPolicy ||
			existing.EventType != row.EventType ||
			existing.PeerType != row.PeerType ||
			existing.PeerId != row.PeerId ||
			!bytes.Equal(existing.PayloadHash, row.PayloadHash) {
			return false, dialogpb.ErrPayloadConflict
		}
		return true, nil
	}
	return false, nil
}

func authSeqOutboxDuplicateExists(txModels *model.TxModels, in authSeqOutboxInput) (bool, error) {
	if in.OperationID == "" {
		return false, dialogpb.ErrOutboxUnavailable
	}
	if in.SourcePermAuthKeyID == 0 {
		return false, dialogpb.ErrSourceAuthKeyRequired
	}
	if len(in.Payload) == 0 {
		in.Payload = []byte(`{"schema_version":1}`)
	}
	existing, err := txModels.DialogAuthSeqOutboxModel.SelectByUserOperation(in.UserID, in.OperationID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return false, nil
		}
		return false, storageError("select dialog auth seq outbox duplicate", err)
	}
	if existing.SourcePermAuthKeyId != in.SourcePermAuthKeyID ||
		existing.TargetAuthPolicy != TargetAuthPolicyNotSourcePermAuthKey ||
		existing.EventType != in.EventType ||
		existing.PeerType != in.PeerType ||
		existing.PeerId != in.PeerID ||
		!bytes.Equal(existing.PayloadHash, hashPayload(in.Payload)) {
		return false, dialogpb.ErrPayloadConflict
	}
	return true, nil
}

type publicUpdateOutboxInput struct {
	OutboxID            int64
	SourceUserID        int64
	SourcePermAuthKeyID int64
	TargetUserID        int64
	TargetAuthPolicy    string
	OperationID         string
	DeliveryPath        string
	PublicUpdateType    string
	PeerType            int32
	PeerID              int64
	Payload             []byte
}

func insertPublicUpdateOutbox(txModels *model.TxModels, in publicUpdateOutboxInput) (bool, error) {
	if in.OutboxID == 0 || in.OperationID == "" {
		return false, dialogpb.ErrOutboxUnavailable
	}
	if in.TargetAuthPolicy == TargetAuthPolicyNotSourcePermAuthKey && in.SourcePermAuthKeyID == 0 {
		return false, dialogpb.ErrSourceAuthKeyRequired
	}
	if len(in.Payload) == 0 {
		in.Payload = []byte(`{"schema_version":1}`)
	}
	row := &model.DialogPublicUpdateOutbox{
		OutboxId:             in.OutboxID,
		SourceUserId:         in.SourceUserID,
		SourcePermAuthKeyId:  in.SourcePermAuthKeyID,
		TargetUserId:         in.TargetUserID,
		TargetAuthPolicy:     in.TargetAuthPolicy,
		OperationId:          in.OperationID,
		DeliveryPath:         in.DeliveryPath,
		PublicUpdateType:     in.PublicUpdateType,
		PeerType:             in.PeerType,
		PeerId:               in.PeerID,
		PayloadSchemaVersion: 1,
		Payload:              in.Payload,
		PayloadHash:          hashPayload(in.Payload),
		Status:               OutboxStatusPending,
		AttemptCount:         0,
		NextRetryAt:          unixNow(),
		LeaseOwner:           "",
		LeaseUntil:           0,
		PublishedPts:         0,
		PublishedPtsCount:    0,
		PublishedSeq:         0,
		PublishedDate:        0,
		LastErrorKind:        "",
		LastErrorMessage:     "",
	}
	_, rowsAffected, err := txModels.DialogPublicUpdateOutboxModel.InsertIgnore(row)
	if err != nil {
		return false, storageError("insert dialog public update outbox", err)
	}
	if rowsAffected == 0 {
		existing, err := txModels.DialogPublicUpdateOutboxModel.SelectByTargetOperation(in.TargetUserID, in.OperationID, in.DeliveryPath, in.PublicUpdateType)
		if err != nil {
			return false, storageError("select dialog public update outbox duplicate", err)
		}
		if existing.SourceUserId != row.SourceUserId ||
			existing.SourcePermAuthKeyId != row.SourcePermAuthKeyId ||
			existing.TargetAuthPolicy != row.TargetAuthPolicy ||
			existing.PeerType != row.PeerType ||
			existing.PeerId != row.PeerId ||
			!bytes.Equal(existing.PayloadHash, row.PayloadHash) {
			return false, dialogpb.ErrPayloadConflict
		}
		return true, nil
	}
	return false, nil
}

func hashPayload(b []byte) []byte {
	sum := sha256.Sum256(b)
	out := make([]byte, len(sum))
	copy(out, sum[:])
	return out
}

func clearOmittedPinnedPreferences(txModels *model.TxModels, userID int64, folderID int32, keepPeerDialogIDs []int64) error {
	if folderID != 0 {
		if len(keepPeerDialogIDs) == 0 {
			if _, err := txModels.DialogPreferencesModel.ClearFolderPinned(userID, folderID); err != nil {
				return storageError("clear omitted folder_pinned_order", err)
			}
			return nil
		}
		if _, err := txModels.DialogPreferencesModel.ClearFolderPinnedExcept(userID, folderID, keepPeerDialogIDs); err != nil {
			return storageError("clear omitted folder_pinned_order", err)
		}
		return nil
	}
	if len(keepPeerDialogIDs) == 0 {
		if _, err := txModels.DialogPreferencesModel.ClearMainPinned(userID); err != nil {
			return storageError("clear omitted main_pinned_order", err)
		}
		return nil
	}
	if _, err := txModels.DialogPreferencesModel.ClearMainPinnedExcept(userID, keepPeerDialogIDs); err != nil {
		return storageError("clear omitted main_pinned_order", err)
	}
	return nil
}

func unixNow() int64 {
	return time.Now().UTC().Unix()
}

func unixOrZero(seconds int64) int64 {
	if seconds <= 0 {
		return 0
	}
	return seconds
}

func unixFromTime(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.UTC().Unix()
}

func timeFromUnixOrZero(seconds int64) time.Time {
	if seconds <= 0 {
		return time.Time{}
	}
	return time.Unix(seconds, 0).UTC()
}
