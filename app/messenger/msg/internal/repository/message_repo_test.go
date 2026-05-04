package repository

import "testing"

func TestHistoryPeerSeqWindowHonorsNegativeAddOffset(t *testing.T) {
	minSeq, maxSeq := historyPeerSeqWindow(ListHistoryMessagesInput{
		OffsetID:  1,
		AddOffset: -25,
	})
	if minSeq != 0 {
		t.Fatalf("minSeq = %d, want 0", minSeq)
	}
	if maxSeq <= 1 {
		t.Fatalf("maxSeq = %d, want a window newer than offset_id 1", maxSeq)
	}
}

func TestHistoryPeerSeqWindowKeepsPlainOffsetOlderThanOffsetID(t *testing.T) {
	minSeq, maxSeq := historyPeerSeqWindow(ListHistoryMessagesInput{
		OffsetID: 3,
	})
	if minSeq != 0 || maxSeq != 3 {
		t.Fatalf("window = (%d, %d), want (0, 3)", minSeq, maxSeq)
	}
}

func TestHistoryPeerSeqWindowAppliesMinAndMaxFilters(t *testing.T) {
	minSeq, maxSeq := historyPeerSeqWindow(ListHistoryMessagesInput{
		OffsetID:  10,
		AddOffset: -5,
		MaxID:     12,
		MinID:     4,
	})
	if minSeq != 4 || maxSeq != 12 {
		t.Fatalf("window = (%d, %d), want (4, 12)", minSeq, maxSeq)
	}
}
