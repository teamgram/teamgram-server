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
	listFn  func(context.Context, int64, bool, int32) ([]repository.DialogRecord, error)
	countFn func(context.Context, int64, bool, int32) (int32, error)
	getFn   func(context.Context, int64, int32, int64) (*repository.DialogRecord, error)
	idsFn   func(context.Context, int64, []int64) ([]repository.DialogRecord, error)
}

func (f fakeDialogRepo) ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]repository.DialogRecord, error) {
	return f.listFn(ctx, userID, excludePinned, folderID)
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
