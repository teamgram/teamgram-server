package repository

import (
	"context"
	"errors"
	"testing"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func TestMapDialogRecord(t *testing.T) {
	row := testDialogRow()
	record := mapDialogRecord(row)

	if record.UserID != row.UserId {
		t.Fatalf("UserID = %d, want %d", record.UserID, row.UserId)
	}
	if record.PeerDialogID != row.PeerDialogId {
		t.Fatalf("PeerDialogID = %d, want %d", record.PeerDialogID, row.PeerDialogId)
	}
	if record.Date != row.Date2 {
		t.Fatalf("Date = %d, want %d", record.Date, row.Date2)
	}
	if record.WallpaperOverridden != row.WallpaperOverridden {
		t.Fatalf("WallpaperOverridden = %t, want %t", record.WallpaperOverridden, row.WallpaperOverridden)
	}
}

func TestMapDialogRecordsNilAndEmpty(t *testing.T) {
	if got := mapDialogRecords(nil); len(got) != 0 {
		t.Fatalf("mapDialogRecords(nil) length = %d, want 0", len(got))
	}
	if got := mapDialogRecords([]model.Dialogs{}); len(got) != 0 {
		t.Fatalf("mapDialogRecords(empty) length = %d, want 0", len(got))
	}
}

func TestOldMixedDialogReadMethodsAreDeprecated(t *testing.T) {
	repo := NewRepositoryForTest(&model.Models{})
	if got, err := repo.ListDialogs(context.Background(), 1, false, 3); !errors.Is(err, dialogpb.ErrDeprecatedMethod) || got != nil {
		t.Fatalf("ListDialogs = (%+v, %v), want deprecated nil", got, err)
	}
	if got, err := repo.CountDialogs(context.Background(), 1, false, 3); !errors.Is(err, dialogpb.ErrDeprecatedMethod) || got != 0 {
		t.Fatalf("CountDialogs = (%d, %v), want deprecated zero", got, err)
	}
	if got, err := repo.GetDialogByPeer(context.Background(), 1, 2, 3); !errors.Is(err, dialogpb.ErrDeprecatedMethod) || got != nil {
		t.Fatalf("GetDialogByPeer = (%+v, %v), want deprecated nil", got, err)
	}
	if got, err := repo.ListDialogsByPeerDialogIDs(context.Background(), 1, []int64{1}); !errors.Is(err, dialogpb.ErrDeprecatedMethod) || got != nil {
		t.Fatalf("ListDialogsByPeerDialogIDs = (%+v, %v), want deprecated nil", got, err)
	}
}

func TestCountDialogRecordsUsesListLength(t *testing.T) {
	records := []DialogRecord{
		{UserID: 1, PeerDialogID: 10},
		{UserID: 1, PeerDialogID: 11},
	}
	if got := countDialogRecords(records); got != 2 {
		t.Fatalf("countDialogRecords = %d, want 2", got)
	}
}

func TestFindDialogInRecords(t *testing.T) {
	records := []DialogRecord{
		{UserID: 1, PeerType: 2, PeerID: 20, PeerDialogID: 200},
		{UserID: 1, PeerType: 3, PeerID: 30, PeerDialogID: 300},
	}

	got, ok := findDialogInRecords(records, func(r DialogRecord) bool {
		return r.PeerType == 3 && r.PeerID == 30
	})
	if !ok {
		t.Fatal("findDialogInRecords did not find record")
	}
	if got.PeerDialogID != 300 {
		t.Fatalf("PeerDialogID = %d, want 300", got.PeerDialogID)
	}
}

func TestFindDialogInRecordsMiss(t *testing.T) {
	_, ok := findDialogInRecords([]DialogRecord{{PeerDialogID: 1}}, func(r DialogRecord) bool {
		return r.PeerDialogID == 2
	})
	if ok {
		t.Fatal("findDialogInRecords matched missing record")
	}
}

func TestWrapReadStoragePreservesCause(t *testing.T) {
	cause := errors.New("query failed")
	err := wrapReadStorage("select dialogs", cause)
	if err == nil {
		t.Fatal("wrapReadStorage returned nil")
	}
	if !errors.Is(err, cause) {
		t.Fatalf("errors.Is(%v, cause) = false", err)
	}
}

func testDialogRow() model.Dialogs {
	return model.Dialogs{
		UserId:               1,
		PeerType:             2,
		PeerId:               20,
		PeerDialogId:         200,
		Pinned:               9,
		TopMessage:           100,
		PinnedMsgId:          88,
		ReadInboxMaxId:       77,
		ReadOutboxMaxId:      66,
		UnreadCount:          5,
		UnreadMentionsCount:  4,
		UnreadReactionsCount: 3,
		UnreadMark:           true,
		DraftType:            1,
		DraftMessageData:     `{"schema_version":1}`,
		FolderId:             3,
		FolderPinned:         2,
		HasScheduled:         true,
		TtlPeriod:            3600,
		ThemeEmoticon:        "moon",
		WallpaperId:          123,
		WallpaperOverridden:  true,
		Date2:                1710000000,
	}
}
