package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func TestDialogPreferencesTogglePinWritesVersionAndOutbox(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	in := ToggleDialogPinInput{
		UserID:              base%1_000_000_000 + 101,
		PeerType:            PeerTypeUser,
		PeerID:              base%1_000_000_000 + 201,
		Pinned:              true,
		PinOrder:            1,
		SourcePermAuthKeyID: base%1_000_000_000 + 301,
		OperationID:         fmt.Sprintf("pref-pin-%d", base),
		OutboxID:            base%1_000_000_000 + 401,
		EventType:           "dialog.preferencePinned",
		Payload:             []byte(`{"schema_version":1}`),
	}
	got, err := repo.ToggleDialogPin(ctx, in)
	if err != nil {
		t.Fatalf("ToggleDialogPin() error = %v", err)
	}
	if got.AggregateVersion != 1 {
		t.Fatalf("AggregateVersion = %d, want 1", got.AggregateVersion)
	}
	if _, err := repo.model.DialogAuthSeqOutboxModel.SelectByUserOperation(ctx, in.UserID, in.OperationID); err != nil {
		t.Fatalf("SelectByUserOperation() error = %v", err)
	}
	retry := in
	retry.OutboxID++
	retry.PinOrder = 99
	if _, err := repo.ToggleDialogPin(ctx, retry); err != nil {
		t.Fatalf("ToggleDialogPin(retry) error = %v", err)
	}
	version, err := repo.model.DialogPreferenceVersionsModel.SelectOne(ctx, in.UserID, PreferenceScopeMainPin, 0)
	if err != nil {
		t.Fatalf("SelectOne() error = %v", err)
	}
	if version.AggregateVersion != 1 {
		t.Fatalf("AggregateVersion after retry = %d, want 1", version.AggregateVersion)
	}
	pref, err := repo.model.DialogPreferencesModel.SelectByUserPeer(ctx, in.UserID, in.PeerType, in.PeerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if pref.MainPinnedOrder != 1 {
		t.Fatalf("MainPinnedOrder after retry = %d, want 1", pref.MainPinnedOrder)
	}
}

func TestDialogPreferencesFolderPinWritesFolderPinOrder(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	in := ToggleDialogPinInput{
		UserID:              base%1_000_000_000 + 301,
		PeerType:            PeerTypeUser,
		PeerID:              base%1_000_000_000 + 401,
		FolderID:            7,
		Pinned:              true,
		PinOrder:            44,
		SourcePermAuthKeyID: base%1_000_000_000 + 501,
		OperationID:         fmt.Sprintf("pref-folder-pin-%d", base),
		OutboxID:            base%1_000_000_000 + 601,
		EventType:           "dialog.preferencePinned",
		Payload:             []byte(`{"schema_version":1}`),
	}
	if _, err := repo.ToggleDialogPin(ctx, in); err != nil {
		t.Fatalf("ToggleDialogPin() error = %v", err)
	}
	row, err := repo.model.DialogPreferencesModel.SelectByUserPeer(ctx, in.UserID, in.PeerType, in.PeerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.FolderPinnedOrder != 44 || row.MainPinnedOrder != 0 {
		t.Fatalf("pin orders = main:%d folder:%d, want main:0 folder:44", row.MainPinnedOrder, row.FolderPinnedOrder)
	}
}

func TestListPinnedDialogsHonorsPreferenceOrder(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 451
	firstPeerID := base%1_000_000_000 + 452
	secondPeerID := base%1_000_000_000 + 453
	firstDialogID, err := MakePeerDialogID(PeerTypeUser, firstPeerID)
	if err != nil {
		t.Fatalf("MakePeerDialogID(first) error = %v", err)
	}
	secondDialogID, err := MakePeerDialogID(PeerTypeUser, secondPeerID)
	if err != nil {
		t.Fatalf("MakePeerDialogID(second) error = %v", err)
	}
	for _, row := range []*model.Dialogs{
		{UserId: userID, PeerType: PeerTypeUser, PeerId: firstPeerID, PeerDialogId: firstDialogID, TopMessage: 11, Date2: 111, DraftMessageData: `{}`},
		{UserId: userID, PeerType: PeerTypeUser, PeerId: secondPeerID, PeerDialogId: secondDialogID, TopMessage: 22, Date2: 222, DraftMessageData: `{}`},
	} {
		if _, err := repo.model.DialogsModel.Insert2(ctx, row); err != nil {
			t.Fatalf("DialogsModel.Insert2() error = %v", err)
		}
	}
	for _, in := range []ToggleDialogPinInput{
		{
			UserID:              userID,
			PeerType:            PeerTypeUser,
			PeerID:              firstPeerID,
			Pinned:              true,
			PinOrder:            10,
			SourcePermAuthKeyID: base%1_000_000_000 + 454,
			OperationID:         fmt.Sprintf("list-pin-first-%d", base),
			OutboxID:            base%1_000_000_000 + 455,
			EventType:           "dialog.preferencePinned",
			Payload:             []byte(`{"schema_version":1}`),
		},
		{
			UserID:              userID,
			PeerType:            PeerTypeUser,
			PeerID:              secondPeerID,
			Pinned:              true,
			PinOrder:            20,
			SourcePermAuthKeyID: base%1_000_000_000 + 454,
			OperationID:         fmt.Sprintf("list-pin-second-%d", base),
			OutboxID:            base%1_000_000_000 + 456,
			EventType:           "dialog.preferencePinned",
			Payload:             []byte(`{"schema_version":1}`),
		},
	} {
		if _, err := repo.ToggleDialogPin(ctx, in); err != nil {
			t.Fatalf("ToggleDialogPin() error = %v", err)
		}
	}
	got, err := repo.ListPinnedDialogs(ctx, userID, 0)
	if err != nil {
		t.Fatalf("ListPinnedDialogs() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len(ListPinnedDialogs) = %d, want 2", len(got))
	}
	if got[0].PeerID != secondPeerID || got[0].Order != 20 || got[1].PeerID != firstPeerID || got[1].Order != 10 {
		t.Fatalf("ListPinnedDialogs = %+v, want second(order 20) then first(order 10)", got)
	}
}

func TestReorderPinnedDialogsClearsOmittedPins(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 551
	firstPeerID := base%1_000_000_000 + 552
	secondPeerID := base%1_000_000_000 + 553
	for i, peerID := range []int64{firstPeerID, secondPeerID} {
		in := ToggleDialogPinInput{
			UserID:              userID,
			PeerType:            PeerTypeUser,
			PeerID:              peerID,
			Pinned:              true,
			PinOrder:            int64(i + 1),
			SourcePermAuthKeyID: base%1_000_000_000 + 554,
			OperationID:         fmt.Sprintf("reorder-clear-pin-%d-%d", base, i),
			OutboxID:            base%1_000_000_000 + 555 + int64(i),
			EventType:           "dialog.preferencePinned",
			Payload:             []byte(`{"schema_version":1}`),
		}
		if _, err := repo.ToggleDialogPin(ctx, in); err != nil {
			t.Fatalf("ToggleDialogPin() error = %v", err)
		}
	}
	if _, err := repo.ReorderPinnedDialogs(ctx, ReorderPinnedDialogsInput{
		UserID:              userID,
		PeerOrder:           []PeerRef{{PeerType: PeerTypeUser, PeerID: secondPeerID}},
		SourcePermAuthKeyID: base%1_000_000_000 + 554,
		OperationID:         fmt.Sprintf("reorder-clear-%d", base),
		OutboxID:            base%1_000_000_000 + 557,
		EventType:           "dialog.pinnedDialogsReordered",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		t.Fatalf("ReorderPinnedDialogs() error = %v", err)
	}
	first, err := repo.model.DialogPreferencesModel.SelectByUserPeer(ctx, userID, PeerTypeUser, firstPeerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer(first) error = %v", err)
	}
	second, err := repo.model.DialogPreferencesModel.SelectByUserPeer(ctx, userID, PeerTypeUser, secondPeerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer(second) error = %v", err)
	}
	if first.MainPinnedOrder != 0 || second.MainPinnedOrder != 1 {
		t.Fatalf("pin orders = first:%d second:%d, want first cleared and second normalized to 1", first.MainPinnedOrder, second.MainPinnedOrder)
	}
}

func TestEditPeerFoldersIncrementsOldFolderVersion(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 651
	peerID := base%1_000_000_000 + 652
	first := EditPeerFoldersInput{
		UserID:              userID,
		PeerType:            PeerTypeUser,
		PeerID:              peerID,
		NewFolderID:         7,
		SourcePermAuthKeyID: base%1_000_000_000 + 653,
		OperationID:         fmt.Sprintf("folder-first-%d", base),
		OutboxID:            base%1_000_000_000 + 654,
		PublicUpdateType:    "updateFolderPeers",
		Payload:             []byte(`{"schema_version":1}`),
	}
	if _, err := repo.EditPeerFolders(ctx, first); err != nil {
		t.Fatalf("EditPeerFolders(first) error = %v", err)
	}
	second := first
	second.NewFolderID = 8
	second.OperationID = fmt.Sprintf("folder-second-%d", base)
	second.OutboxID++
	if _, err := repo.EditPeerFolders(ctx, second); err != nil {
		t.Fatalf("EditPeerFolders(second) error = %v", err)
	}
	oldVersion, err := repo.model.DialogPreferenceVersionsModel.SelectOne(ctx, userID, PreferenceScopeFolderMembership, 7)
	if err != nil {
		t.Fatalf("SelectOne(old folder) error = %v", err)
	}
	if oldVersion.AggregateVersion != 2 {
		t.Fatalf("old folder AggregateVersion = %d, want 2", oldVersion.AggregateVersion)
	}
}

func TestDialogDraftsClearAfterSendSkipsNewerDraft(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 501
	peerID := base%1_000_000_000 + 601
	draftDate := time.Now().UTC()
	save := SaveDraftInput{
		UserID:              userID,
		PeerType:            PeerTypeUser,
		PeerID:              peerID,
		DraftKind:           1,
		Message:             "new draft",
		EntitiesPayload:     []byte(`[]`),
		DraftPayload:        []byte(`{"schema_version":1}`),
		Date:                unixFromTime(draftDate),
		SourcePermAuthKeyID: base%1_000_000_000 + 701,
		OperationID:         fmt.Sprintf("draft-save-%d", base),
		OutboxID:            base%1_000_000_000 + 801,
		EventType:           "dialog.draftSaved",
	}
	if _, err := repo.SaveDraft(ctx, save); err != nil {
		t.Fatalf("SaveDraft() error = %v", err)
	}
	clear := ClearDraftAfterSendInput{
		UserID:              userID,
		PeerType:            PeerTypeUser,
		PeerID:              peerID,
		ClearBeforeDate:     unixFromTime(draftDate.Add(-time.Second)),
		SourcePermAuthKeyID: save.SourcePermAuthKeyID,
		OperationID:         fmt.Sprintf("draft-clear-%d", base),
		OutboxID:            base%1_000_000_000 + 802,
		EventType:           "dialog.draftCleared",
		Payload:             []byte(`{"schema_version":1}`),
	}
	got, err := repo.ClearDraftAfterSend(ctx, clear)
	if err != nil {
		t.Fatalf("ClearDraftAfterSend() error = %v", err)
	}
	if got.Cleared {
		t.Fatalf("ClearDraftAfterSend cleared newer draft")
	}
	row, err := repo.model.DialogDraftsModel.SelectByUserPeer(ctx, userID, PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.Message != "new draft" {
		t.Fatalf("draft message = %q, want new draft", row.Message)
	}
	if _, err := repo.model.DialogAuthSeqOutboxModel.SelectByUserOperation(ctx, userID, clear.OperationID); !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("SelectByUserOperation(clear newer) error = %v, want not found", err)
	}
}

func TestDialogDraftsSaveDefaultsNilPayloads(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 601
	peerID := base%1_000_000_000 + 701
	save := SaveDraftInput{
		UserID:              userID,
		PeerType:            PeerTypeUser,
		PeerID:              peerID,
		DraftKind:           1,
		Message:             "plain draft",
		Date:                unixFromTime(time.Now().UTC()),
		SourcePermAuthKeyID: base%1_000_000_000 + 801,
		OperationID:         fmt.Sprintf("draft-save-nil-payloads-%d", base),
		OutboxID:            base%1_000_000_000 + 901,
		EventType:           "dialog.draftSaved",
	}
	if _, err := repo.SaveDraft(ctx, save); err != nil {
		t.Fatalf("SaveDraft() error = %v", err)
	}
	row, err := repo.model.DialogDraftsModel.SelectByUserPeer(ctx, userID, PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.EntitiesPayload == nil || row.DraftPayload == nil {
		t.Fatalf("payloads = entities:%v draft:%v, want non-nil blobs", row.EntitiesPayload, row.DraftPayload)
	}
}

func TestDialogDraftsClearAllCreatesOneOutboxPerClearedPeer(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 701
	sourceAuthKeyID := base%1_000_000_000 + 702
	for i := int64(0); i < 2; i++ {
		if _, err := repo.SaveDraft(ctx, SaveDraftInput{
			UserID:              userID,
			PeerType:            PeerTypeUser,
			PeerID:              base%1_000_000_000 + 710 + i,
			DraftKind:           1,
			Message:             fmt.Sprintf("draft-%d", i),
			EntitiesPayload:     []byte(`[]`),
			DraftPayload:        []byte(`{"schema_version":1}`),
			Date:                unixFromTime(time.Now().UTC().Add(time.Duration(i) * time.Second)),
			SourcePermAuthKeyID: sourceAuthKeyID,
			OperationID:         fmt.Sprintf("draft-clear-all-save-%d-%d", base, i),
			OutboxID:            base%1_000_000_000 + 720 + i,
			EventType:           "dialog.draftSaved",
		}); err != nil {
			t.Fatalf("SaveDraft(%d) error = %v", i, err)
		}
	}
	clearOpID := fmt.Sprintf("draft-clear-all-%d", base)
	got, err := repo.ClearAllDrafts(ctx, ClearAllDraftsInput{
		UserID:              userID,
		SourcePermAuthKeyID: sourceAuthKeyID,
		OperationID:         clearOpID,
		OutboxIDs: []int64{
			base%1_000_000_000 + 730,
			base%1_000_000_000 + 731,
		},
	})
	if err != nil {
		t.Fatalf("ClearAllDrafts() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len(ClearAllDrafts) = %d, want 2", len(got))
	}
	active, err := repo.model.DialogDraftsModel.SelectActiveByUser(ctx, userID)
	if err != nil {
		t.Fatalf("SelectActiveByUser() error = %v", err)
	}
	if len(active) != 0 {
		t.Fatalf("active drafts = %d, want 0", len(active))
	}
	for _, result := range got {
		op := fmt.Sprintf("%s:peer:%d", clearOpID, result.PeerDialogID)
		var rows []struct {
			OperationID string `db:"operation_id"`
		}
		if err := repo.db.QueryRowsPartial(ctx, &rows, "SELECT operation_id FROM dialog_auth_seq_outbox WHERE user_id = ? AND peer_id != 0 AND operation_id = ?", userID, op); err != nil {
			t.Fatalf("query outbox for peer_dialog_id %d: %v", result.PeerDialogID, err)
		}
		if len(rows) != 1 {
			t.Fatalf("outbox rows for %s = %d, want 1", op, len(rows))
		}
	}
}

func TestDialogFiltersSlugLookupUsesExplicitSlug(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 901
	if _, _, err := repo.model.DialogFiltersModel.InsertOrUpdate(ctx, dialogModelFilter{
		UserId:              userID,
		DialogFilterId:      1,
		Slug:                "work",
		Title:               "Work",
		OrderValue:          1,
		Enabled:             true,
		FilterSchemaVersion: 1,
		FilterPayload:       []byte(`{"schema_version":1}`),
	}.toModel()); err != nil {
		t.Fatalf("InsertOrUpdate() error = %v", err)
	}
	row, err := repo.model.DialogFiltersModel.SelectBySlug(ctx, userID, "work")
	if err != nil {
		t.Fatalf("SelectBySlug() error = %v", err)
	}
	if row.Title != "Work" {
		t.Fatalf("Title = %q, want Work", row.Title)
	}
}

func TestSavedDialogsPinnedFalseForcesZeroOrder(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1001
	peerID := base%1_000_000_000 + 1002
	if err := repo.UpsertSavedDialogFromMessage(ctx, SavedDialogTopInput{
		UserID:                userID,
		PeerType:              PeerTypeUser,
		PeerID:                peerID,
		TopPeerSeq:            1,
		TopCanonicalMessageID: base%1_000_000_000 + 1003,
		TopMessageDate:        unixFromTime(time.Now().UTC()),
		Payload:               []byte(`{"schema_version":1}`),
	}); err != nil {
		t.Fatalf("UpsertSavedDialogFromMessage() error = %v", err)
	}
	if err := repo.ToggleSavedDialogPin(ctx, SavedDialogPinInput{
		UserID:              userID,
		PeerType:            PeerTypeUser,
		PeerID:              peerID,
		Pinned:              false,
		PinOrder:            99,
		SourcePermAuthKeyID: base%1_000_000_000 + 1004,
		OperationID:         fmt.Sprintf("saved-pin-%d", base),
		OutboxID:            base%1_000_000_000 + 1005,
		EventType:           "dialog.savedDialogPinned",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		t.Fatalf("ToggleSavedDialogPin() error = %v", err)
	}
	row, err := repo.model.SavedDialogsModel.Select(ctx, userID, PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("Select() error = %v", err)
	}
	if row.Pinned || row.PinOrder != 0 {
		t.Fatalf("saved pin = %t/%d, want false/0", row.Pinned, row.PinOrder)
	}
}

func TestSavedDialogsOlderMessageDoesNotOverwriteTopPayload(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1051
	peerID := base%1_000_000_000 + 1052
	newer := time.Now().UTC()
	if err := repo.UpsertSavedDialogFromMessage(ctx, SavedDialogTopInput{
		UserID:                userID,
		PeerType:              PeerTypeUser,
		PeerID:                peerID,
		TopPeerSeq:            2,
		TopCanonicalMessageID: base%1_000_000_000 + 1053,
		TopMessageDate:        unixFromTime(newer),
		Payload:               []byte(`{"message":"newer"}`),
	}); err != nil {
		t.Fatalf("UpsertSavedDialogFromMessage(newer) error = %v", err)
	}
	if err := repo.UpsertSavedDialogFromMessage(ctx, SavedDialogTopInput{
		UserID:                userID,
		PeerType:              PeerTypeUser,
		PeerID:                peerID,
		TopPeerSeq:            1,
		TopCanonicalMessageID: base%1_000_000_000 + 1054,
		TopMessageDate:        unixFromTime(newer.Add(-time.Minute)),
		Payload:               []byte(`{"message":"older"}`),
	}); err != nil {
		t.Fatalf("UpsertSavedDialogFromMessage(older) error = %v", err)
	}
	row, err := repo.model.SavedDialogsModel.Select(ctx, userID, PeerTypeUser, peerID)
	if err != nil {
		t.Fatalf("Select() error = %v", err)
	}
	if string(row.SavedPayload) != `{"message":"newer"}` {
		t.Fatalf("SavedPayload = %s, want newer payload", row.SavedPayload)
	}
}

func TestSavedDialogPinInvariant(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1071
	peerIDs := []int64{base%1_000_000_000 + 1072, base%1_000_000_000 + 1073}
	for i, peerID := range peerIDs {
		if err := repo.UpsertSavedDialogFromMessage(ctx, SavedDialogTopInput{
			UserID:                userID,
			PeerType:              PeerTypeUser,
			PeerID:                peerID,
			TopPeerSeq:            int64(i + 1),
			TopCanonicalMessageID: base%1_000_000_000 + int64(1074+i),
			TopMessageDate:        unixFromTime(time.Now().UTC().Add(time.Duration(i) * time.Second)),
			Payload:               []byte(`{"schema_version":1}`),
		}); err != nil {
			t.Fatalf("UpsertSavedDialogFromMessage(%d) error = %v", i, err)
		}
		if err := repo.ToggleSavedDialogPin(ctx, SavedDialogPinInput{
			UserID:              userID,
			PeerType:            PeerTypeUser,
			PeerID:              peerID,
			Pinned:              true,
			SourcePermAuthKeyID: base%1_000_000_000 + int64(1080+i),
			OperationID:         fmt.Sprintf("saved-pin-invariant-%d-%d", base, i),
			OutboxID:            base%1_000_000_000 + int64(1082+i),
			EventType:           "dialog.savedDialogPinned",
			Payload:             []byte(`{"schema_version":1}`),
		}); err != nil {
			t.Fatalf("ToggleSavedDialogPin(%d) error = %v", i, err)
		}
	}
	rows, err := repo.ListPinnedSavedDialogs(ctx, userID)
	if err != nil {
		t.Fatalf("ListPinnedSavedDialogs() error = %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("pinned len = %d, want 2", len(rows))
	}
	seen := map[int64]bool{}
	for _, row := range rows {
		if !row.Pinned || row.PinOrder <= 0 {
			t.Fatalf("saved pin = %t/%d, want true/positive", row.Pinned, row.PinOrder)
		}
		if seen[row.PinOrder] {
			t.Fatalf("duplicate pin_order %d in %+v", row.PinOrder, rows)
		}
		seen[row.PinOrder] = true
	}
}

func TestReorderPinnedSavedDialogsNoDuplicateOrder(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1091
	peerIDs := []int64{base%1_000_000_000 + 1092, base%1_000_000_000 + 1093, base%1_000_000_000 + 1094}
	for i, peerID := range peerIDs {
		if err := repo.UpsertSavedDialogFromMessage(ctx, SavedDialogTopInput{
			UserID:                userID,
			PeerType:              PeerTypeUser,
			PeerID:                peerID,
			TopPeerSeq:            int64(i + 1),
			TopCanonicalMessageID: base%1_000_000_000 + int64(1095+i),
			TopMessageDate:        unixFromTime(time.Now().UTC().Add(time.Duration(i) * time.Second)),
			Payload:               []byte(`{"schema_version":1}`),
		}); err != nil {
			t.Fatalf("UpsertSavedDialogFromMessage(%d) error = %v", i, err)
		}
	}
	if err := repo.ReorderPinnedSavedDialogs(ctx, ReorderPinnedSavedDialogsInput{
		UserID: userID,
		Order: []PeerRef{
			{PeerType: PeerTypeUser, PeerID: peerIDs[2]},
			{PeerType: PeerTypeUser, PeerID: peerIDs[0]},
		},
		SourcePermAuthKeyID: base%1_000_000_000 + 1100,
		OperationID:         fmt.Sprintf("saved-reorder-%d", base),
		OutboxID:            base%1_000_000_000 + 1101,
		EventType:           "dialog.pinnedSavedDialogs",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		t.Fatalf("ReorderPinnedSavedDialogs() error = %v", err)
	}
	rows, err := repo.ListPinnedSavedDialogs(ctx, userID)
	if err != nil {
		t.Fatalf("ListPinnedSavedDialogs() error = %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("pinned len = %d, want 2", len(rows))
	}
	for i, row := range rows {
		if row.PinOrder != int64(i+1) {
			t.Fatalf("row %d pin_order = %d, want %d", i, row.PinOrder, i+1)
		}
		if row.PeerID != []int64{peerIDs[2], peerIDs[0]}[i] {
			t.Fatalf("row %d peer_id = %d", i, row.PeerID)
		}
	}
}

func TestDialogPeerPolicyPrivatePairCanonical(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1101
	peerID := base%1_000_000_000 + 1102
	got, err := repo.SetPrivatePeerPolicy(ctx, PrivatePeerPolicyInput{
		UserID:              userID,
		PeerUserID:          peerID,
		TTLPeriod:           3600,
		SourcePermAuthKeyID: base%1_000_000_000 + 1103,
		OperationID:         fmt.Sprintf("policy-%d", base),
		ActorOutboxID:       base%1_000_000_000 + 1104,
		PeerOutboxID:        base%1_000_000_000 + 1105,
		DeliveryPath:        "userupdates_auth_seq",
		PublicUpdateType:    "updatePeerHistoryTTL",
		Payload:             []byte(`{"schema_version":1}`),
	})
	if err != nil {
		t.Fatalf("SetPrivatePeerPolicy() error = %v", err)
	}
	if got.Scope.PeerType != 0 || got.Scope.PeerID != 0 {
		t.Fatalf("private pair peer tuple = %+v, want neutral", got.Scope)
	}
}

func TestDialogVisualSettingsWallpaperUserScoped(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	in := PeerWallpaperInput{
		UserID:              base%1_000_000_000 + 1201,
		PeerType:            PeerTypeChat,
		PeerID:              base%1_000_000_000 + 1202,
		WallpaperID:         42,
		WallpaperOverridden: true,
		SourcePermAuthKeyID: base%1_000_000_000 + 1203,
		OperationID:         fmt.Sprintf("wallpaper-%d", base),
		OutboxID:            base%1_000_000_000 + 1204,
		EventType:           "dialog.wallpaperChanged",
		Payload:             []byte(`{"schema_version":1}`),
	}
	if err := repo.SetPeerWallpaper(ctx, in); err != nil {
		t.Fatalf("SetPeerWallpaper() error = %v", err)
	}
	row, err := repo.model.DialogVisualSettingsModel.SelectByUserPeer(ctx, in.UserID, in.PeerType, in.PeerID)
	if err != nil {
		t.Fatalf("SelectByUserPeer() error = %v", err)
	}
	if row.WallpaperId != 42 || !row.WallpaperOverridden {
		t.Fatalf("wallpaper row = %+v", row)
	}
}

func TestDialogAuthSeqOutboxRetriesAreIdempotentAndDetectConflict(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	in := PeerWallpaperInput{
		UserID:              base%1_000_000_000 + 1251,
		PeerType:            PeerTypeChat,
		PeerID:              base%1_000_000_000 + 1252,
		WallpaperID:         42,
		WallpaperOverridden: true,
		SourcePermAuthKeyID: base%1_000_000_000 + 1253,
		OperationID:         fmt.Sprintf("wallpaper-idempotent-%d", base),
		OutboxID:            base%1_000_000_000 + 1254,
		EventType:           "dialog.wallpaperChanged",
		Payload:             []byte(`{"schema_version":1}`),
	}
	if err := repo.SetPeerWallpaper(ctx, in); err != nil {
		t.Fatalf("SetPeerWallpaper(first) error = %v", err)
	}
	retry := in
	retry.OutboxID++
	if err := repo.SetPeerWallpaper(ctx, retry); err != nil {
		t.Fatalf("SetPeerWallpaper(retry) error = %v", err)
	}
	conflict := retry
	conflict.OutboxID++
	conflict.Payload = []byte(`{"schema_version":1,"changed":true}`)
	if err := repo.SetPeerWallpaper(ctx, conflict); !errors.Is(err, dialogpb.ErrPayloadConflict) {
		t.Fatalf("SetPeerWallpaper(conflict) error = %v, want ErrPayloadConflict", err)
	}
}

func TestDialogAuthSeqOutboxRejectsMissingSourceAuthKey(t *testing.T) {
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	err := repo.SetPeerWallpaper(context.Background(), PeerWallpaperInput{UserID: 1, PeerType: PeerTypeUser, PeerID: 2})
	if !errors.Is(err, dialogpb.ErrSourceAuthKeyRequired) {
		t.Fatalf("SetPeerWallpaper() error = %v, want ErrSourceAuthKeyRequired", err)
	}
}

func TestDialogPublicUpdateOutboxPrivatePairTargets(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	userID := base%1_000_000_000 + 1301
	peerID := base%1_000_000_000 + 1302
	op := fmt.Sprintf("policy-target-%d", base)
	if _, err := repo.SetPrivatePeerPolicy(ctx, PrivatePeerPolicyInput{
		UserID:              userID,
		PeerUserID:          peerID,
		ThemeEmoticon:       "moon",
		SourcePermAuthKeyID: base%1_000_000_000 + 1303,
		OperationID:         op,
		ActorOutboxID:       base%1_000_000_000 + 1304,
		PeerOutboxID:        base%1_000_000_000 + 1305,
		DeliveryPath:        "userupdates_pts",
		PublicUpdateType:    "messageActionSetChatTheme",
		Payload:             []byte(`{"schema_version":1}`),
	}); err != nil {
		t.Fatalf("SetPrivatePeerPolicy() error = %v", err)
	}
	for _, target := range []int64{userID, peerID} {
		if _, err := repo.model.DialogPublicUpdateOutboxModel.SelectByTargetOperation(ctx, target, fmt.Sprintf("%s:target:%d", op, target), "userupdates_pts", "messageActionSetChatTheme"); err != nil {
			t.Fatalf("SelectByTargetOperation(target=%d) error = %v", target, err)
		}
	}
}

func TestDialogPublicUpdateOutboxRetriesAreIdempotentAndDetectConflict(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	base := time.Now().UnixNano()
	in := PrivatePeerPolicyInput{
		UserID:              base%1_000_000_000 + 1351,
		PeerUserID:          base%1_000_000_000 + 1352,
		ThemeEmoticon:       "moon",
		SourcePermAuthKeyID: base%1_000_000_000 + 1353,
		OperationID:         fmt.Sprintf("policy-idempotent-%d", base),
		ActorOutboxID:       base%1_000_000_000 + 1354,
		PeerOutboxID:        base%1_000_000_000 + 1355,
		DeliveryPath:        "userupdates_pts",
		PublicUpdateType:    "messageActionSetChatTheme",
		Payload:             []byte(`{"schema_version":1}`),
	}
	if _, err := repo.SetPrivatePeerPolicy(ctx, in); err != nil {
		t.Fatalf("SetPrivatePeerPolicy(first) error = %v", err)
	}
	retry := in
	retry.ActorOutboxID += 10
	retry.PeerOutboxID += 10
	if _, err := repo.SetPrivatePeerPolicy(ctx, retry); err != nil {
		t.Fatalf("SetPrivatePeerPolicy(retry) error = %v", err)
	}
	conflict := retry
	conflict.ActorOutboxID += 10
	conflict.PeerOutboxID += 10
	conflict.Payload = []byte(`{"schema_version":1,"changed":true}`)
	if _, err := repo.SetPrivatePeerPolicy(ctx, conflict); !errors.Is(err, dialogpb.ErrPayloadConflict) {
		t.Fatalf("SetPrivatePeerPolicy(conflict) error = %v, want ErrPayloadConflict", err)
	}
}

type dialogModelFilter struct {
	UserId              int64
	DialogFilterId      int32
	Slug                string
	Title               string
	OrderValue          int64
	Enabled             bool
	FilterSchemaVersion int32
	FilterPayload       []byte
}

func (f dialogModelFilter) toModel() *model.DialogFilters {
	return &model.DialogFilters{
		UserId:              f.UserId,
		DialogFilterId:      f.DialogFilterId,
		Slug:                f.Slug,
		Title:               f.Title,
		OrderValue:          f.OrderValue,
		Enabled:             f.Enabled,
		Deleted:             false,
		FilterSchemaVersion: f.FilterSchemaVersion,
		FilterPayload:       f.FilterPayload,
	}
}

func openDialogIntegrationDB(t *testing.T) *sqlx.DB {
	t.Helper()
	if testing.Short() {
		t.Skip("integration test skipped in short mode")
	}
	dsn := os.Getenv("TEAMGRAM_TEST_MYSQL_DSN")
	explicit := dsn != ""
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	}
	db, err := sqlx.Open(&sqlx.Config{DSN: dsn})
	if err != nil {
		if explicit {
			t.Fatalf("open mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	if _, err := db.Exec(context.Background(), "SELECT 1"); err != nil {
		if explicit {
			t.Fatalf("ping mysql: %v", err)
		}
		t.Skipf("mysql unavailable: %v", err)
	}
	return db
}
