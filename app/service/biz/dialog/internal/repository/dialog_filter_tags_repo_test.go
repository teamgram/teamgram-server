package repository

import (
	"context"
	"testing"
	"time"
)

func TestDialogFilterTagsDefaultFalseAndToggleByUser(t *testing.T) {
	ctx := context.Background()
	repo := NewRepositoryWithDBForTest(openDialogIntegrationDB(t))
	userID := time.Now().UnixNano()%1_000_000_000 + 9001

	enabled, err := repo.GetDialogFilterTagsEnabled(ctx, userID)
	if err != nil {
		t.Fatalf("GetDialogFilterTagsEnabled() error = %v", err)
	}
	if enabled {
		t.Fatalf("GetDialogFilterTagsEnabled() = true, want false for missing row")
	}

	if err := repo.SetDialogFilterTagsEnabled(ctx, userID, true); err != nil {
		t.Fatalf("SetDialogFilterTagsEnabled(true) error = %v", err)
	}
	enabled, err = repo.GetDialogFilterTagsEnabled(ctx, userID)
	if err != nil {
		t.Fatalf("GetDialogFilterTagsEnabled() after enable error = %v", err)
	}
	if !enabled {
		t.Fatalf("GetDialogFilterTagsEnabled() = false, want true")
	}

	if err := repo.SetDialogFilterTagsEnabled(ctx, userID, false); err != nil {
		t.Fatalf("SetDialogFilterTagsEnabled(false) error = %v", err)
	}
	enabled, err = repo.GetDialogFilterTagsEnabled(ctx, userID)
	if err != nil {
		t.Fatalf("GetDialogFilterTagsEnabled() after disable error = %v", err)
	}
	if enabled {
		t.Fatalf("GetDialogFilterTagsEnabled() = true, want false after disable")
	}
}
