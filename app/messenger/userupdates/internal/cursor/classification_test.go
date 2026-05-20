package cursor

import (
	"math"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestClassifyUpdateCommonPTSAllowlist(t *testing.T) {
	tests := []tg.UpdateClazz{
		tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{}),
		tg.MakeTLUpdateDeleteMessages(&tg.TLUpdateDeleteMessages{}),
		tg.MakeTLUpdateReadHistoryInbox(&tg.TLUpdateReadHistoryInbox{}),
		tg.MakeTLUpdateReadHistoryOutbox(&tg.TLUpdateReadHistoryOutbox{}),
		tg.MakeTLUpdateWebPage(&tg.TLUpdateWebPage{}),
		tg.MakeTLUpdateReadMessagesContents(&tg.TLUpdateReadMessagesContents{}),
		tg.MakeTLUpdateEditMessage(&tg.TLUpdateEditMessage{}),
		tg.MakeTLUpdateFolderPeers(&tg.TLUpdateFolderPeers{}),
		tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{}),
	}
	for _, update := range tests {
		if got := ClassifyUpdate(update); got != DeliveryClassUserPTS {
			t.Fatalf("%T classified as %s, want %s", update, got, DeliveryClassUserPTS)
		}
	}
}

func TestClassifyUpdateChannelPTSAllowlist(t *testing.T) {
	tests := []tg.UpdateClazz{
		tg.MakeTLUpdateChannelTooLong(&tg.TLUpdateChannelTooLong{}),
		tg.MakeTLUpdateNewChannelMessage(&tg.TLUpdateNewChannelMessage{}),
		tg.MakeTLUpdateReadChannelInbox(&tg.TLUpdateReadChannelInbox{}),
		tg.MakeTLUpdateDeleteChannelMessages(&tg.TLUpdateDeleteChannelMessages{}),
		tg.MakeTLUpdateEditChannelMessage(&tg.TLUpdateEditChannelMessage{}),
		tg.MakeTLUpdateChannelWebPage(&tg.TLUpdateChannelWebPage{}),
		tg.MakeTLUpdatePinnedChannelMessages(&tg.TLUpdatePinnedChannelMessages{}),
	}
	for _, update := range tests {
		if got := ClassifyUpdate(update); got != DeliveryClassChannelPTS {
			t.Fatalf("%T classified as %s, want %s", update, got, DeliveryClassChannelPTS)
		}
	}
}

func TestClassifyUpdateNonPTSDefaultsRealtimeOnly(t *testing.T) {
	got := ClassifyUpdate(tg.MakeTLUpdatePhoneCall(&tg.TLUpdatePhoneCall{}))
	if got != DeliveryClassRealtimeOnly {
		t.Fatalf("ClassifyUpdate(updatePhoneCall) = %s, want %s", got, DeliveryClassRealtimeOnly)
	}
}

func TestCheckedInt32(t *testing.T) {
	got, err := CheckedInt32(math.MaxInt32, "seq")
	if err != nil {
		t.Fatalf("CheckedInt32(MaxInt32) error = %v", err)
	}
	if got != math.MaxInt32 {
		t.Fatalf("CheckedInt32(MaxInt32) = %d", got)
	}
	if _, err := CheckedInt32(int64(math.MaxInt32)+1, "seq"); err == nil {
		t.Fatalf("CheckedInt32(MaxInt32+1) error = nil")
	}
	if _, err := CheckedInt32(int64(math.MinInt32)-1, "seq"); err == nil {
		t.Fatalf("CheckedInt32(MinInt32-1) error = nil")
	}
}
