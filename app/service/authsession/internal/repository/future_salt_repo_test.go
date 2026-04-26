package repository

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/xkv"
)

type fakeFutureSaltsModel struct {
	salts []*xkv.FutureSaltRecord
	put   []*xkv.FutureSaltRecord
}

func (f *fakeFutureSaltsModel) PutSalts(ctx context.Context, keyId int64, salts []*xkv.FutureSaltRecord) error {
	f.put = salts
	return nil
}

func (f *fakeFutureSaltsModel) GetSalts(ctx context.Context, keyId int64) ([]*xkv.FutureSaltRecord, error) {
	return f.salts, nil
}

func (f *fakeFutureSaltsModel) DeleteSalts(ctx context.Context, keyId int64) error {
	f.salts = nil
	return nil
}

func TestGetOrNotInsertSaltListOnlyGeneratesMissingSalts(t *testing.T) {
	now := int32(time.Now().Unix())
	model := &fakeFutureSaltsModel{
		salts: []*xkv.FutureSaltRecord{
			{ValidSince: now, ValidUntil: now + saltTimeout, Salt: 1},
			{ValidSince: now + saltTimeout, ValidUntil: now + 2*saltTimeout, Salt: 2},
		},
	}
	repo := &Repository{futureSaltsModel: model}

	salts, err := repo.getOrNotInsertSaltList(context.Background(), 1001, 3)
	if err != nil {
		t.Fatalf("getOrNotInsertSaltList() error = %v", err)
	}
	if len(salts) != 3 {
		t.Fatalf("len(salts) = %d, want 3", len(salts))
	}
	if len(model.put) != 3 {
		t.Fatalf("cached salt count = %d, want exactly 3", len(model.put))
	}
}

func TestGetOrNotInsertSaltListKeepsRecentlyExpiredSaltForCompatibility(t *testing.T) {
	now := int32(time.Now().Unix())
	model := &fakeFutureSaltsModel{
		salts: []*xkv.FutureSaltRecord{
			{ValidSince: now - 2*saltTimeout, ValidUntil: now - 60, Salt: 99},
		},
	}
	repo := &Repository{futureSaltsModel: model}

	salts, err := repo.getOrNotInsertSaltList(context.Background(), 1001, 2)
	if err != nil {
		t.Fatalf("getOrNotInsertSaltList() error = %v", err)
	}
	if len(salts) != 3 {
		t.Fatalf("len(salts) = %d, want previous salt plus 2 future salts", len(salts))
	}
	if salts[0].Salt != 99 {
		t.Fatalf("first salt = %d, want recently expired compatibility salt", salts[0].Salt)
	}
}

func TestGetOrNotInsertSaltListDropsExpiredSaltOutsideCompatibilityWindow(t *testing.T) {
	now := int32(time.Now().Unix())
	model := &fakeFutureSaltsModel{
		salts: []*xkv.FutureSaltRecord{
			{ValidSince: now - 2*saltTimeout, ValidUntil: now - 301, Salt: 99},
		},
	}
	repo := &Repository{futureSaltsModel: model}

	salts, err := repo.getOrNotInsertSaltList(context.Background(), 1001, 2)
	if err != nil {
		t.Fatalf("getOrNotInsertSaltList() error = %v", err)
	}
	if len(salts) != 2 {
		t.Fatalf("len(salts) = %d, want only requested future salts", len(salts))
	}
	if salts[0].Salt == 99 {
		t.Fatal("expired salt outside compatibility window was returned")
	}
}
