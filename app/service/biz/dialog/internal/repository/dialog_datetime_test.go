package repository

import (
	"testing"
	"time"
)

func TestMysqlDateTimeForBindPreservesUTCWallClock(t *testing.T) {
	original := time.Date(2026, 1, 1, 0, 2, 3, 456789000, time.UTC)

	bound := mysqlDateTimeForBind(original)
	got := bound.In(mysqlDriverLocation()).Format("2006-01-02 15:04:05.000000")
	want := mysqlTimestamp(original)
	if got != want {
		t.Fatalf("mysqlDateTimeForBind formatted wall-clock = %q, want %q", got, want)
	}

	read := time.Date(2026, 1, 1, 0, 2, 3, 456789000, mysqlDriverLocation())
	gotUTC := mysqlDateTimeToUTC(read)
	if gotUTC.Location() != time.UTC {
		t.Fatalf("mysqlDateTimeToUTC location = %v, want UTC", gotUTC.Location())
	}
	if gotUTC.Format("2006-01-02 15:04:05.000000") != want {
		t.Fatalf("mysqlDateTimeToUTC wall-clock = %q, want %q", gotUTC.Format("2006-01-02 15:04:05.000000"), want)
	}
}
