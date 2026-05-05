package core

import (
	"context"
	"errors"
	"testing"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDialogRepo struct {
	listFn        func(context.Context, int64, bool, int32) ([]repository.DialogRecord, error)
	listPinnedFn  func(context.Context, int64, int32) ([]repository.DialogRecord, error)
	countFn       func(context.Context, int64, bool, int32) (int32, error)
	getFn         func(context.Context, int64, int32, int64) (*repository.DialogRecord, error)
	idsFn         func(context.Context, int64, []int64) ([]repository.DialogRecord, error)
	extrasFn      func(context.Context, int64, []repository.PeerRef) ([]repository.DialogExtrasRecord, error)
	saveDraftFn   func(context.Context, repository.SaveDraftInput) (*repository.DraftMutationResult, error)
	clearDraftFn  func(context.Context, repository.ClearDraftInput) (*repository.DraftMutationResult, error)
	clearAfterFn  func(context.Context, repository.ClearDraftAfterSendInput) (*repository.DraftMutationResult, error)
	clearAllFn    func(context.Context, repository.ClearAllDraftsInput) ([]repository.DraftMutationResult, error)
	listDraftsFn  func(context.Context, int64) ([]repository.DraftRecord, error)
	togglePinFn   func(context.Context, repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error)
	reorderPinFn  func(context.Context, repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error)
	editFoldersFn func(context.Context, repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error)
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

func (f fakeDialogRepo) ToggleDialogPin(ctx context.Context, in repository.ToggleDialogPinInput) (*repository.PreferenceMutationResult, error) {
	return f.togglePinFn(ctx, in)
}

func (f fakeDialogRepo) ReorderPinnedDialogs(ctx context.Context, in repository.ReorderPinnedDialogsInput) (*repository.PreferenceMutationResult, error) {
	return f.reorderPinFn(ctx, in)
}

func (f fakeDialogRepo) EditPeerFolders(ctx context.Context, in repository.EditPeerFoldersInput) (*repository.PreferenceMutationResult, error) {
	return f.editFoldersFn(ctx, in)
}

func TestDialogGetDialogsReturnsMappedVector(t *testing.T) {
	repo := fakeDialogRepo{
		listFn: func(_ context.Context, userID int64, excludePinned bool, folderID int32) ([]repository.DialogRecord, error) {
			if userID != 1001 || !excludePinned || folderID != 2 {
				t.Fatalf("request = userID:%d excludePinned:%t folderID:%d", userID, excludePinned, folderID)
			}
			return []repository.DialogRecord{testDialogRecord()}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetDialogs(&dialogpb.TLDialogGetDialogs{
		UserId:        1001,
		ExcludePinned: tg.BoolTrueClazz,
		FolderId:      2,
	})
	if err != nil {
		t.Fatalf("DialogGetDialogs error = %v", err)
	}
	if len(got.Datas) != 1 {
		t.Fatalf("len(Datas) = %d, want 1", len(got.Datas))
	}
	ext := got.Datas[0].ToDialogExt()
	if ext.Order != 17 {
		t.Fatalf("Order = %d, want 17", ext.Order)
	}
	if ext.ThemeEmoticon != "moon" {
		t.Fatalf("ThemeEmoticon = %q, want moon", ext.ThemeEmoticon)
	}
	if ext.Dialog.(*tg.TLDialog).TopMessage != 99 {
		t.Fatalf("TopMessage = %d, want 99", ext.Dialog.(*tg.TLDialog).TopMessage)
	}
}

func TestDialogGetPinnedDialogsHonorsRepositoryOrder(t *testing.T) {
	repo := fakeDialogRepo{
		listPinnedFn: func(_ context.Context, userID int64, folderID int32) ([]repository.DialogRecord, error) {
			if userID != 1001 || folderID != 1 {
				t.Fatalf("request = userID:%d folderID:%d", userID, folderID)
			}
			first := testDialogRecord()
			first.Order = 90
			second := testDialogRecord()
			second.PeerID = 2003
			second.PeerDialogID = 18
			second.Order = 80
			return []repository.DialogRecord{first, second}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetPinnedDialogs(&dialogpb.TLDialogGetPinnedDialogs{UserId: 1001, FolderId: 1})
	if err != nil {
		t.Fatalf("DialogGetPinnedDialogs error = %v", err)
	}
	if len(got.Datas) != 2 {
		t.Fatalf("len(Datas) = %d, want 2", len(got.Datas))
	}
	if got.Datas[0].Order != 90 || got.Datas[1].Order != 80 {
		t.Fatalf("orders = %d,%d; want 90,80", got.Datas[0].Order, got.Datas[1].Order)
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
	if got.Message != "draft" || got.Date.Unix() != 123 {
		t.Fatalf("draft mapping = message:%q date:%v", got.Message, got.Date)
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
	if got.ClearBeforeDate.Unix() != 123 {
		t.Fatalf("ClearBeforeDate = %v, want unix 123", got.ClearBeforeDate)
	}
}

func TestDialogGetDialogsReturnsEmptyVector(t *testing.T) {
	repo := fakeDialogRepo{
		listFn: func(context.Context, int64, bool, int32) ([]repository.DialogRecord, error) {
			return []repository.DialogRecord{}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetDialogs(&dialogpb.TLDialogGetDialogs{UserId: 1001, ExcludePinned: tg.BoolFalseClazz})
	if err != nil {
		t.Fatalf("DialogGetDialogs error = %v", err)
	}
	if got == nil || len(got.Datas) != 0 {
		t.Fatalf("DialogGetDialogs = %#v, want empty vector", got)
	}
}

func TestDialogGetDialogsPassesThroughRepositoryError(t *testing.T) {
	cause := errors.New("repository failed")
	repo := fakeDialogRepo{
		listFn: func(context.Context, int64, bool, int32) ([]repository.DialogRecord, error) {
			return nil, cause
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	_, err := core.DialogGetDialogs(&dialogpb.TLDialogGetDialogs{UserId: 1001, ExcludePinned: tg.BoolFalseClazz})
	if !errors.Is(err, cause) {
		t.Fatalf("error = %v, want cause", err)
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

func TestDialogGetDialogsCount(t *testing.T) {
	repo := fakeDialogRepo{
		countFn: func(_ context.Context, userID int64, excludePinned bool, folderID int32) (int32, error) {
			if userID != 1001 || excludePinned || folderID != 7 {
				t.Fatalf("request = userID:%d excludePinned:%t folderID:%d", userID, excludePinned, folderID)
			}
			return 4, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetDialogsCount(&dialogpb.TLDialogGetDialogsCount{
		UserId:        1001,
		ExcludePinned: tg.BoolFalseClazz,
		FolderId:      7,
	})
	if err != nil {
		t.Fatalf("DialogGetDialogsCount error = %v", err)
	}
	if got.V != 4 {
		t.Fatalf("count = %d, want 4", got.V)
	}
}

func TestDialogGetDialogById(t *testing.T) {
	record := testDialogRecord()
	repo := fakeDialogRepo{
		getFn: func(_ context.Context, userID int64, peerType int32, peerID int64) (*repository.DialogRecord, error) {
			if userID != 1001 || peerType != tg.PEER_USER || peerID != 2002 {
				t.Fatalf("request = userID:%d peerType:%d peerID:%d", userID, peerType, peerID)
			}
			return &record, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetDialogById(&dialogpb.TLDialogGetDialogById{
		UserId:   1001,
		PeerType: tg.PEER_USER,
		PeerId:   2002,
	})
	if err != nil {
		t.Fatalf("DialogGetDialogById error = %v", err)
	}
	if got.Order != 17 {
		t.Fatalf("Order = %d, want 17", got.Order)
	}
}

func TestDialogGetDialogsByIdList(t *testing.T) {
	repo := fakeDialogRepo{
		idsFn: func(_ context.Context, userID int64, ids []int64) ([]repository.DialogRecord, error) {
			if userID != 1001 {
				t.Fatalf("userID = %d, want 1001", userID)
			}
			if len(ids) != 2 || ids[0] != 17 || ids[1] != 18 {
				t.Fatalf("ids = %#v, want [17 18]", ids)
			}
			first := testDialogRecord()
			second := testDialogRecord()
			second.PeerDialogID = 18
			return []repository.DialogRecord{first, second}, nil
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.DialogGetDialogsByIdList(&dialogpb.TLDialogGetDialogsByIdList{
		UserId: 1001,
		IdList: []int64{17, 18},
	})
	if err != nil {
		t.Fatalf("DialogGetDialogsByIdList error = %v", err)
	}
	if len(got.Datas) != 2 {
		t.Fatalf("len(Datas) = %d, want 2", len(got.Datas))
	}
}
