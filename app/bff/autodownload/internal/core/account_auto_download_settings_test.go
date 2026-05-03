package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountGetAutoDownloadSettingsReturnsMasterDefaults(t *testing.T) {
	core := &AutoDownloadCore{ctx: context.Background()}

	got, err := core.AccountGetAutoDownloadSettings(&tg.TLAccountGetAutoDownloadSettings{})
	if err != nil {
		t.Fatalf("AccountGetAutoDownloadSettings() error = %v", err)
	}
	if got == nil {
		t.Fatal("AccountGetAutoDownloadSettings() = nil")
	}

	assertAutoDownloadSettings(t, "low", got.Low, true, 512000, 512000)
	assertAutoDownloadSettings(t, "medium", got.Medium, false, 10485760, 1048576)
	assertAutoDownloadSettings(t, "high", got.High, false, 15728640, 3145728)
}

func TestAccountSaveAutoDownloadSettingsReturnsBoolTrue(t *testing.T) {
	core := &AutoDownloadCore{ctx: context.Background()}

	got, err := core.AccountSaveAutoDownloadSettings(&tg.TLAccountSaveAutoDownloadSettings{})
	if err != nil {
		t.Fatalf("AccountSaveAutoDownloadSettings() error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("AccountSaveAutoDownloadSettings() = %p, want %p", got, tg.BoolTrue)
	}
}

func assertAutoDownloadSettings(t *testing.T, name string, got *tg.AutoDownloadSettings, phonecallsLessData bool, videoSizeMax int64, fileSizeMax int64) {
	t.Helper()
	if got == nil {
		t.Fatalf("%s settings = nil", name)
	}
	if got.Disabled {
		t.Fatalf("%s.Disabled = true, want false", name)
	}
	if !got.VideoPreloadLarge {
		t.Fatalf("%s.VideoPreloadLarge = false, want true", name)
	}
	if !got.AudioPreloadNext {
		t.Fatalf("%s.AudioPreloadNext = false, want true", name)
	}
	if got.PhonecallsLessData != phonecallsLessData {
		t.Fatalf("%s.PhonecallsLessData = %v, want %v", name, got.PhonecallsLessData, phonecallsLessData)
	}
	if got.PhotoSizeMax != 1048576 {
		t.Fatalf("%s.PhotoSizeMax = %d, want 1048576", name, got.PhotoSizeMax)
	}
	if got.VideoSizeMax != videoSizeMax {
		t.Fatalf("%s.VideoSizeMax = %d, want %d", name, got.VideoSizeMax, videoSizeMax)
	}
	if got.FileSizeMax != fileSizeMax {
		t.Fatalf("%s.FileSizeMax = %d, want %d", name, got.FileSizeMax, fileSizeMax)
	}
	if got.StoriesPreload {
		t.Fatalf("%s.StoriesPreload = true, want false", name)
	}
	if got.VideoUploadMaxbitrate != 0 {
		t.Fatalf("%s.VideoUploadMaxbitrate = %d, want 0", name, got.VideoUploadMaxbitrate)
	}
	if got.SmallQueueActiveOperationsMax != 0 {
		t.Fatalf("%s.SmallQueueActiveOperationsMax = %d, want 0", name, got.SmallQueueActiveOperationsMax)
	}
	if got.LargeQueueActiveOperationsMax != 0 {
		t.Fatalf("%s.LargeQueueActiveOperationsMax = %d, want 0", name, got.LargeQueueActiveOperationsMax)
	}
}
