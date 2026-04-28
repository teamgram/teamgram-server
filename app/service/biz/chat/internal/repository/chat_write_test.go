package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository/model"
)

func TestBackfillBulkInsertIDsAssignsSequentialIDs(t *testing.T) {
	rows := []*model.ChatParticipants{
		{UserId: 1},
		{UserId: 2},
		{UserId: 3},
	}
	if err := backfillBulkInsertIDs(rows, 41, 3); err != nil {
		t.Fatalf("backfillBulkInsertIDs error: %v", err)
	}
	for i, row := range rows {
		want := int64(41 + i)
		if row.Id != want {
			t.Fatalf("row %d id = %d, want %d", i, row.Id, want)
		}
	}
}

func TestBackfillBulkInsertIDsRejectsMissingInsertMetadata(t *testing.T) {
	rows := []*model.ChatParticipants{{UserId: 1}}
	if err := backfillBulkInsertIDs(rows, 0, 1); err == nil {
		t.Fatal("backfillBulkInsertIDs accepted zero last insert id")
	}
	if err := backfillBulkInsertIDs(rows, 10, 0); err == nil {
		t.Fatal("backfillBulkInsertIDs accepted wrong rows affected")
	}
}
