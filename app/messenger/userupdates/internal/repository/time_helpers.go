package repository

import (
	"strconv"
	"time"
)

func unixNow() int64 {
	return time.Now().UTC().Unix()
}

func unixOrZero(seconds int64) int64 {
	if seconds <= 0 {
		return 0
	}
	return seconds
}

func unixOrNow(seconds int64) int64 {
	if seconds > 0 {
		return seconds
	}
	return unixNow()
}

func unixOptionalString(seconds int64) string {
	if seconds <= 0 {
		return ""
	}
	return strconv.FormatInt(seconds, 10)
}

func unixFromTimeOrZero(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.UTC().Unix()
}

func unixTimeFromSentinel(seconds int64) time.Time {
	if seconds <= 0 {
		return time.Time{}
	}
	return time.Unix(seconds, 0).UTC()
}
