package tg

import (
	"math"
	"testing"
)

func TestDateInt32FromUnixSecondsAcceptsProtocolRange(t *testing.T) {
	got, err := DateInt32FromUnixSeconds(math.MaxInt32)
	if err != nil {
		t.Fatalf("DateInt32FromUnixSeconds(MaxInt32) error = %v", err)
	}
	if got != math.MaxInt32 {
		t.Fatalf("DateInt32FromUnixSeconds(MaxInt32) = %d, want %d", got, math.MaxInt32)
	}
}

func TestDateInt32FromUnixSecondsRejectsOverflow(t *testing.T) {
	_, err := DateInt32FromUnixSeconds(int64(math.MaxInt32) + 1)
	if err == nil {
		t.Fatal("DateInt32FromUnixSeconds(MaxInt32+1) error = nil, want overflow error")
	}
}

func TestDateInt32FromUnixSecondsRejectsNegative(t *testing.T) {
	_, err := DateInt32FromUnixSeconds(-1)
	if err == nil {
		t.Fatal("DateInt32FromUnixSeconds(-1) error = nil, want range error")
	}
}
