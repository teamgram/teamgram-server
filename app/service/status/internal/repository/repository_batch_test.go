package repository

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/status/status"
)

func TestAssignUserSessionBatchResultPreservesInputOrder(t *testing.T) {
	users := []int64{10, 20, 30}
	result := &status.VectorUserSessionEntryList{
		Datas: make([]*status.TLUserSessionEntryList, len(users)),
	}

	assignUserSessionBatchResult(context.Background(), result,
		[]userSessionBatchEntry{
			{index: 2, userID: 30},
			{index: 0, userID: 10},
		},
		[]map[string]string{
			{},
			{},
		},
	)
	assignUserSessionBatchResult(context.Background(), result,
		[]userSessionBatchEntry{
			{index: 1, userID: 20},
		},
		[]map[string]string{
			{},
		},
	)

	if len(result.Datas) != len(users) {
		t.Fatalf("len(result.Datas) = %d, want %d", len(result.Datas), len(users))
	}
	for i, userID := range users {
		if result.Datas[i] == nil {
			t.Fatalf("result.Datas[%d] is nil", i)
		}
		if result.Datas[i].UserId != userID {
			t.Fatalf("result.Datas[%d].UserId = %d, want %d", i, result.Datas[i].UserId, userID)
		}
		if len(result.Datas[i].UserSessions) != 0 {
			t.Fatalf("len(result.Datas[%d].UserSessions) = %d, want 0", i, len(result.Datas[i].UserSessions))
		}
	}
}
