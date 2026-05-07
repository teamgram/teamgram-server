package repository

import (
	"testing"
	"time"
)

func TestMysqlTimestampFormatsUTC(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatalf("load location: %v", err)
	}

	instant := time.Date(2026, 5, 6, 23, 0, 11, 0, loc)
	if got, want := mysqlTimestamp(instant), "2026-05-06 15:00:11.000000"; got != want {
		t.Fatalf("mysqlTimestamp() = %q, want %q", got, want)
	}
}
