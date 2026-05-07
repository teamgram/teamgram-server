package repository

import (
	"testing"
	"time"
)

func TestUnixTimeHelpers(t *testing.T) {
	original := time.Date(2026, 1, 1, 0, 2, 3, 456789000, time.UTC)

	if got := unixFromTime(original); got != original.Unix() {
		t.Fatalf("unixFromTime() = %d, want %d", got, original.Unix())
	}
	if got := unixOrZero(-1); got != 0 {
		t.Fatalf("unixOrZero(-1) = %d, want 0", got)
	}
	if got := timeFromUnixOrZero(original.Unix()); !got.Equal(time.Unix(original.Unix(), 0).UTC()) {
		t.Fatalf("timeFromUnixOrZero() = %v, want unix UTC", got)
	}
}
