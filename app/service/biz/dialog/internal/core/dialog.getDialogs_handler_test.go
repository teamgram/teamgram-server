package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDialogRepo struct {
	listFn         func(context.Context, int64, bool, int32) ([]repository.DialogRecord, error)
	listPinnedFn   func(context.Context, int64, int32) ([]repository.DialogRecord, error)
	countFn        func(context.Context, int64, bool, int32) (int32, error)
	getFn          func(context.Context, int64, int32, int64) (*repository.DialogRecord, error)
	idsFn          func(context.Context, int64, []int64) ([]repository.DialogRecord, error)
	extrasFn       func(context.Context, int64, []repository.PeerRef) ([]repository.DialogExtrasRecord, error)
	filterListFn   func(context.Context, int64) ([]repository.DialogFilterRecord, error)
	filterGetFn    func(context.Context, int64, int32) (*repository.DialogFilterRecord, error)
	filterSlugFn   func(context.Context, int64, string) (*repository.DialogFilterRecord, error)
	filterSaveFn   func(context.Context, repository.SaveDialogFilterInput) (*repository.DialogFilterRecord, error)
	filterDelFn    func(context.Context, repository.DeleteDialogFilterInput) error
	filterOrderFn  func(context.Context, repository.ReorderDialogFiltersInput) error
	wallpaperFn    func(context.Context, repository.PeerWallpaperInput) error
	policyFn       func(context.Context, repository.PrivatePeerPolicyInput) (*repository.PrivatePeerPolicyResult, error)
	saveDraftFn    func(context.Context, repository.SaveDraftInput) (*repository.DraftMutationResult, error)
	clearDraftFn   func(context.Context, repository.ClearDraftInput) (*repository.DraftMutationResult, error)
	clearAfterFn   func(context.Context, repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error)
	clearAllFn     func(context.Context, repository.ClearAllDraftsInput) ([]repository.DraftMutationResult, error)
	listDraftsFn   func(context.Context, int64) ([]repository.DraftRecord, error)
	upsertSavedFn  func(context.Context, repository.SavedDialogTopInput) error
	listSavedFn    func(context.Context, int64, bool, int64, int32) ([]repository.SavedDialogRecord, error)
	pinnedSavedFn  func(context.Context, int64) ([]repository.SavedDialogRecord, error)
	toggleSavedFn  func(context.Context, repository.SavedDialogPinInput) error
	reorderSavedFn func(context.Context, repository.ReorderPinnedSavedDialogsInput) error
	togglePinFn    func(context.Context, repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error)
	reorderPinFn   func(context.Context, repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error)
	editFoldersFn  func(context.Context, repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error)
}

func (f fakeDialogRepo) ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]repository.DialogRecord, error) {
	return f.listFn(ctx, userID, excludePinned, folderID)
}

func (f fakeDialogRepo) ListPinnedDialogs(ctx context.Context, userID int64, folderID int32) ([]repository.DialogRecord, error) {
	return f.listPinnedFn(ctx, userID, folderID)
}

func (f fakeDialogRepo) CountDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) (int32, error) {
	return f.countFn(ctx, userID, excludePinned, folderID)
}

func (f fakeDialogRepo) GetDialogByPeer(ctx context.Context, userID int64, peerType int32, peerID int64) (*repository.DialogRecord, error) {
	return f.getFn(ctx, userID, peerType, peerID)
}

func (f fakeDialogRepo) ListDialogsByPeerDialogIDs(ctx context.Context, userID int64, ids []int64) ([]repository.DialogRecord, error) {
	return f.idsFn(ctx, userID, ids)
}

func (f fakeDialogRepo) BatchGetDialogExtras(ctx context.Context, userID int64, peers []repository.PeerRef) ([]repository.DialogExtrasRecord, error) {
	if f.extrasFn != nil {
		return f.extrasFn(ctx, userID, peers)
	}
	out := make([]repository.DialogExtrasRecord, 0, len(peers))
	for _, peer := range peers {
		out = append(out, repository.DialogExtrasRecord{PeerType: peer.PeerType, PeerID: peer.PeerID})
	}
	return out, nil
}

func (f fakeDialogRepo) ListDialogFilters(ctx context.Context, userID int64) ([]repository.DialogFilterRecord, error) {
	if f.filterListFn != nil {
		return f.filterListFn(ctx, userID)
	}
	return []repository.DialogFilterRecord{}, nil
}

func (f fakeDialogRepo) GetDialogFilter(ctx context.Context, userID int64, filterID int32) (*repository.DialogFilterRecord, error) {
	if f.filterGetFn != nil {
		return f.filterGetFn(ctx, userID, filterID)
	}
	return &repository.DialogFilterRecord{UserID: userID, DialogFilterID: filterID}, nil
}

func (f fakeDialogRepo) GetDialogFilterBySlug(ctx context.Context, userID int64, slug string) (*repository.DialogFilterRecord, error) {
	if f.filterSlugFn != nil {
		return f.filterSlugFn(ctx, userID, slug)
	}
	return &repository.DialogFilterRecord{UserID: userID, Slug: slug}, nil
}

func (f fakeDialogRepo) SaveDialogFilter(ctx context.Context, in repository.SaveDialogFilterInput) (*repository.DialogFilterRecord, error) {
	if f.filterSaveFn != nil {
		return f.filterSaveFn(ctx, in)
	}
	return &repository.DialogFilterRecord{UserID: in.UserID, DialogFilterID: in.DialogFilterID, Slug: in.Slug, Title: in.Title, OrderValue: in.OrderValue}, nil
}

func (f fakeDialogRepo) DeleteDialogFilter(ctx context.Context, in repository.DeleteDialogFilterInput) error {
	if f.filterDelFn != nil {
		return f.filterDelFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) ReorderDialogFilters(ctx context.Context, in repository.ReorderDialogFiltersInput) error {
	if f.filterOrderFn != nil {
		return f.filterOrderFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) SetPeerWallpaper(ctx context.Context, in repository.PeerWallpaperInput) error {
	if f.wallpaperFn != nil {
		return f.wallpaperFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) SetPrivatePeerPolicy(ctx context.Context, in repository.PrivatePeerPolicyInput) (*repository.PrivatePeerPolicyResult, error) {
	if f.policyFn != nil {
		return f.policyFn(ctx, in)
	}
	return &repository.PrivatePeerPolicyResult{}, nil
}

func (f fakeDialogRepo) SaveDraft(ctx context.Context, in repository.SaveDraftInput) (*repository.DraftMutationResult, error) {
	return f.saveDraftFn(ctx, in)
}

func (f fakeDialogRepo) ClearDraft(ctx context.Context, in repository.ClearDraftInput) (*repository.DraftMutationResult, error) {
	return f.clearDraftFn(ctx, in)
}

func (f fakeDialogRepo) ClearDraftAfterSend(ctx context.Context, in repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error) {
	return f.clearAfterFn(ctx, in)
}

func (f fakeDialogRepo) ClearAllDrafts(ctx context.Context, in repository.ClearAllDraftsInput) ([]repository.DraftMutationResult, error) {
	return f.clearAllFn(ctx, in)
}

func (f fakeDialogRepo) ListActiveDrafts(ctx context.Context, userID int64) ([]repository.DraftRecord, error) {
	return f.listDraftsFn(ctx, userID)
}

func (f fakeDialogRepo) UpsertSavedDialogFromMessage(ctx context.Context, in repository.SavedDialogTopInput) error {
	if f.upsertSavedFn != nil {
		return f.upsertSavedFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) ListSavedDialogs(ctx context.Context, userID int64, excludePinned bool, offsetDate int64, limit int32) ([]repository.SavedDialogRecord, error) {
	if f.listSavedFn != nil {
		return f.listSavedFn(ctx, userID, excludePinned, offsetDate, limit)
	}
	return []repository.SavedDialogRecord{}, nil
}

func (f fakeDialogRepo) ListPinnedSavedDialogs(ctx context.Context, userID int64) ([]repository.SavedDialogRecord, error) {
	if f.pinnedSavedFn != nil {
		return f.pinnedSavedFn(ctx, userID)
	}
	return []repository.SavedDialogRecord{}, nil
}

func (f fakeDialogRepo) ToggleSavedDialogPin(ctx context.Context, in repository.SavedDialogPinInput) error {
	if f.toggleSavedFn != nil {
		return f.toggleSavedFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) ReorderPinnedSavedDialogs(ctx context.Context, in repository.ReorderPinnedSavedDialogsInput) error {
	if f.reorderSavedFn != nil {
		return f.reorderSavedFn(ctx, in)
	}
	return nil
}

func (f fakeDialogRepo) ToggleDialogPin(ctx context.Context, in repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error) {
	return f.togglePinFn(ctx, in)
}

func (f fakeDialogRepo) ReorderPinnedDialogs(ctx context.Context, in repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error) {
	return f.reorderPinFn(ctx, in)
}

func (f fakeDialogRepo) EditPeerFolders(ctx context.Context, in repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error) {
	return f.editFoldersFn(ctx, in)
}

type fakeDialogUserupdates struct {
	listFn     func(context.Context, *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error)
	byPeersFn  func(context.Context, *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error)
	countFn    func(context.Context, *userupdates.TLUserupdatesGetDialogCount) (*tg.Int32, error)
	listCalled bool
}

func (f *fakeDialogUserupdates) UserupdatesListDialogs(ctx context.Context, in *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error) {
	f.listCalled = true
	return f.listFn(ctx, in)
}

func (f *fakeDialogUserupdates) UserupdatesGetDialogsByPeers(ctx context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	return f.byPeersFn(ctx, in)
}

func (f *fakeDialogUserupdates) UserupdatesGetDialogCount(ctx context.Context, in *userupdates.TLUserupdatesGetDialogCount) (*tg.Int32, error) {
	return f.countFn(ctx, in)
}

func TestDialogGetDialogsReturnsMappedVector(t *testing.T) {
	repo := fakeDialogRepo{
		extrasFn: func(_ context.Context, userID int64, peers []repository.PeerRef) ([]repository.DialogExtrasRecord, error) {
			if userID != 1001 || len(peers) != 1 || peers[0].PeerID != 2002 {
				t.Fatalf("extras request = userID:%d peers:%+v", userID, peers)
			}
			return []repository.DialogExtrasRecord{{PeerType: tg.PEER_USER, PeerID: 2002, FolderID: 2, MainPinnedOrder: 17}}, nil
		},
	}
	updates := &fakeDialogUserupdates{
		listFn: func(_ context.Context, in *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error) {
			if in.UserId != 1001 || in.TopMessageDate != 10 || in.TopPeerSeq != 9 || in.PeerType != tg.PEER_USER || in.PeerId != 2002 || in.Limit != 20 {
				t.Fatalf("list request = %+v", in)
			}
			return userupdates.MakeTLDialogProjectionList(&userupdates.TLDialogProjectionList{
				Projections: []userupdates.DialogProjectionClazz{testDialogProjection()},
				Exhausted:   tg.BoolTrueClazz,
			}), nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, Userupdates: updates})

	got, err := core.DialogGetDialogsV2(&dialogpb.TLDialogGetDialogsV2{
		UserId:        1001,
		ExcludePinned: tg.BoolTrueClazz,
		Limit:         20,
		Cursor:        dialogpb.MakeTLDialogCursor(&dialogpb.TLDialogCursor{FolderId: 2, TopMessageDate: 10, TopPeerSeq: 9, PeerType: tg.PEER_USER, PeerId: 2002}),
	})
	if err != nil {
		t.Fatalf("DialogGetDialogsV2 error = %v", err)
	}
	if len(got.Dialogs) != 0 {
		t.Fatalf("excluded pinned dialog count = %d, want 0", len(got.Dialogs))
	}
	if !updates.listCalled {
		t.Fatal("userupdates list was not called")
	}
}

func TestDialogSaveDraftMessageCallsRepositoryWithSourceAuth(t *testing.T) {
	var got repository.SaveDraftInput
	repo := fakeDialogRepo{
		saveDraftFn: func(_ context.Context, in repository.SaveDraftInput) (*repository.DraftMutationResult, error) {
			got = in
			return &repository.DraftMutationResult{UserID: in.UserID}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.DialogSaveDraftMessage(&dialogpb.TLDialogSaveDraftMessage{
		UserId:              1001,
		PeerType:            repository.PeerTypeUser,
		PeerId:              2002,
		Message:             tg.MakeTLDraftMessage(&tg.TLDraftMessage{Message: "draft", Date: 123}),
		SourcePermAuthKeyId: 9001,
		OperationId:         "op-save",
		OutboxId:            7001,
	})
	if err != nil {
		t.Fatalf("DialogSaveDraftMessage error = %v", err)
	}
	if got.SourcePermAuthKeyID != 9001 || got.OperationID != "op-save" || got.OutboxID != 7001 {
		t.Fatalf("repository input = %+v", got)
	}
	if got.Message != "draft" || got.Date != 123 {
		t.Fatalf("draft mapping = message:%q date:%d", got.Message, got.Date)
	}
}

func TestDialogClearDraftAfterSendUsesSourceOperationID(t *testing.T) {
	var got repository.ClearDraftAfterSendInput
	repo := fakeDialogRepo{
		clearAfterFn: func(_ context.Context, in repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error) {
			got = in
			return &repository.DraftMutationResult{UserID: in.UserID}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.DialogClearDraftAfterSend(&dialogpb.TLDialogClearDraftAfterSend{
		UserId:              1001,
		PeerType:            repository.PeerTypeUser,
		PeerId:              2002,
		ClearBeforeDate:     123,
		SourcePermAuthKeyId: 9001,
		SourceOperationId:   "sender-op",
		OutboxId:            7001,
	})
	if err != nil {
		t.Fatalf("DialogClearDraftAfterSend error = %v", err)
	}
	if got.OperationID != "sender-op" {
		t.Fatalf("OperationID = %q, want sender-op", got.OperationID)
	}
	if got.SourcePermAuthKeyID != 9001 || got.OutboxID != 7001 {
		t.Fatalf("repository input = %+v", got)
	}
	if got.ClearBeforeDate != 123 {
		t.Fatalf("ClearBeforeDate = %d, want 123", got.ClearBeforeDate)
	}
}

func TestDialogGetPinnedDialogsHonorsRepositoryOrder(t *testing.T) {
	repo := fakeDialogRepo{
		listPinnedFn: func(_ context.Context, userID int64, folderID int32) ([]repository.DialogRecord, error) {
			if userID != 1001 || folderID != 1 {
				t.Fatalf("request = userID:%d folderID:%d", userID, folderID)
			}
			first := testDialogRecord()
			second := testDialogRecord()
			second.PeerID = 2003
			return []repository.DialogRecord{first, second}, nil
		},
	}
	updates := &fakeDialogUserupdates{
		byPeersFn: func(_ context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
			if len(in.Peers) != 2 || in.Peers[0].PeerId != 2002 || in.Peers[1].PeerId != 2003 {
				t.Fatalf("peers = %+v", in.Peers)
			}
			second := testDialogProjection()
			second.PeerId = 2003
			return &userupdates.VectorDialogProjection{Datas: []userupdates.DialogProjectionClazz{testDialogProjection(), second}}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo, Userupdates: updates})

	got, err := core.DialogGetPinnedDialogsV2(&dialogpb.TLDialogGetPinnedDialogsV2{UserId: 1001, FolderId: 1})
	if err != nil {
		t.Fatalf("DialogGetPinnedDialogsV2 error = %v", err)
	}
	if len(got.Datas) != 2 || got.Datas[0].PeerId != 2002 || got.Datas[1].PeerId != 2003 {
		t.Fatalf("pinned dialogs = %+v", got.Datas)
	}
}

func TestMakeDialogExtV2FromProjectionExposesPublicMessageIDs(t *testing.T) {
	projection := testDialogProjection()
	projection.TopPeerSeq = 99
	projection.TopUserMessageId = 42
	projection.ReadInboxMaxPeerSeq = 12
	projection.ReadInboxMaxUserMessageId = 43
	projection.ReadOutboxMaxPeerSeq = 13
	projection.ReadOutboxMaxUserMessageId = 44
	projection.PinnedPeerSeq = 33
	projection.PinnedUserMessageId = 45
	projection.AvailableMinPeerSeq = 2
	projection.AvailableMinUserMessageId = 41

	got := makeDialogExtV2FromProjection(projection, nil)

	if got.TopPeerSeq != 99 || got.TopUserMessageId != 42 {
		t.Fatalf("top ids = peer_seq:%d public:%d, want peer_seq 99 public 42", got.TopPeerSeq, got.TopUserMessageId)
	}
	if got.ReadInboxMaxPeerSeq != 12 || got.ReadInboxMaxUserMessageId != 43 ||
		got.ReadOutboxMaxPeerSeq != 13 || got.ReadOutboxMaxUserMessageId != 44 {
		t.Fatalf("read ids = %+v, want internal peer seqs with public mirrors", got)
	}
	if got.PinnedPeerSeq != 33 || got.PinnedUserMessageId != 45 ||
		got.AvailableMinPeerSeq != 2 || got.AvailableMinUserMessageId != 41 {
		t.Fatalf("pin/available ids = %+v, want internal peer seqs with public mirrors", got)
	}
}

func TestMakeLegacyDialogExtV2UsesRecordPublicIDsWithoutInventingPeerSeq(t *testing.T) {
	record := testDialogRecord()
	record.TopMessage = 42
	record.ReadInboxMaxID = 43
	record.ReadOutboxMaxID = 44
	record.PinnedMsgID = 45

	got := makeDialogExtV2(record, nil)

	if got.TopUserMessageId != 42 || got.ReadInboxMaxUserMessageId != 43 ||
		got.ReadOutboxMaxUserMessageId != 44 || got.PinnedUserMessageId != 45 {
		t.Fatalf("public ids = %+v, want record ids as public mirrors", got)
	}
	if got.TopPeerSeq != 42 || got.PinnedPeerSeq != 45 {
		t.Fatalf("legacy internal ids = top:%d pinned:%d, want existing record ids preserved", got.TopPeerSeq, got.PinnedPeerSeq)
	}
}

func testDialogRecord() repository.DialogRecord {
	return repository.DialogRecord{
		UserID:               1001,
		PeerType:             tg.PEER_USER,
		PeerID:               2002,
		PeerDialogID:         17,
		Pinned:               1,
		TopMessage:           99,
		PinnedMsgID:          33,
		ReadInboxMaxID:       12,
		ReadOutboxMaxID:      13,
		UnreadCount:          2,
		UnreadMentionsCount:  1,
		UnreadReactionsCount: 3,
		UnreadMark:           true,
		FolderID:             2,
		TTLPeriod:            3600,
		ThemeEmoticon:        "moon",
		WallpaperID:          42,
		WallpaperOverridden:  true,
		Date:                 1710000000,
	}
}

func testDialogProjection() *userupdates.TLDialogProjection {
	return userupdates.MakeTLDialogProjection(&userupdates.TLDialogProjection{
		PeerType:                   tg.PEER_USER,
		PeerId:                     2002,
		TopPeerSeq:                 99,
		TopUserMessageId:           42,
		TopCanonicalMessageId:      7001,
		TopMessageDate:             1710000000,
		ReadInboxMaxPeerSeq:        12,
		ReadInboxMaxUserMessageId:  43,
		ReadOutboxMaxPeerSeq:       13,
		ReadOutboxMaxUserMessageId: 44,
		UnreadCount:                2,
		UnreadMentionsCount:        1,
		UnreadReactionsCount:       3,
		UnreadMark:                 true,
		PinnedPeerSeq:              33,
		PinnedUserMessageId:        45,
		PinnedCanonicalMessageId:   33,
		HasScheduled:               true,
		AvailableMinPeerSeq:        1,
		AvailableMinUserMessageId:  41,
	})
}
