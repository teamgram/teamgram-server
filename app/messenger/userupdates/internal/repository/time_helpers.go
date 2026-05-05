package repository

import (
	"database/sql"
	"time"
)

const mysqlTimeLayout = "2006-01-02 15:04:05.000000"

func mysqlDate(unix int32) time.Time {
	return time.Unix(int64(unix), 0).UTC()
}

func mysqlNow() time.Time {
	return time.Now().UTC()
}

func mysqlZeroTime() time.Time {
	return time.Unix(0, 0).UTC()
}

func mysqlTimestamp(t time.Time) string {
	return t.Local().Format(mysqlTimeLayout)
}

func mysqlNullInvalid() sql.NullTime {
	return sql.NullTime{}
}

func mysqlNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t.UTC(), Valid: true}
}

func mysqlNullDate(unix int32) sql.NullTime {
	return mysqlNullTime(mysqlDate(unix))
}

func mysqlNullNow() sql.NullTime {
	return mysqlNullTime(mysqlNow())
}

func mysqlNullZeroTime() sql.NullTime {
	return mysqlNullTime(mysqlZeroTime())
}

func mysqlNullTimeString(t sql.NullTime) string {
	if !t.Valid {
		return ""
	}
	return mysqlTimestamp(t.Time)
}

func mysqlTimeOrZero(t time.Time) time.Time {
	if t.IsZero() {
		return mysqlZeroTime()
	}
	return t.UTC()
}

func mysqlTimeFromSentinel(t time.Time) time.Time {
	t = t.UTC()
	if t.IsZero() || t.Equal(mysqlZeroTime()) {
		return time.Time{}
	}
	return t
}
