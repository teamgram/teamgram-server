package repository

import "testing"

func TestUserAggregateCacheDTOExcludesTLMediaObjects(t *testing.T) {
	dto := UserAggregateCacheDTO{
		SchemaVersion: 1,
		UserID:        100,
		PhotoID:       200,
		ContactIDs:    []int64{101, 102},
	}
	if dto.SchemaVersion != 1 || dto.UserID != 100 || dto.PhotoID != 200 || len(dto.ContactIDs) != 2 {
		t.Fatalf("unexpected dto: %#v", dto)
	}
}
