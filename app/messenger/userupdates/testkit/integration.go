//go:build integration

package testkit

import (
	"context"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type IDGenerator interface {
	NextID(ctx context.Context) (int64, error)
}

type Kit struct {
	repo *repository.Repository
}

func New(db *sqlx.DB, idgen IDGenerator, ownerInstance string) *Kit {
	return &Kit{repo: repository.NewForTest(db, idgen, ownerInstance)}
}

func (k *Kit) ClaimPartitionOwner(ctx context.Context, partitionID int32) (int64, error) {
	return k.repo.ClaimPartitionOwner(ctx, partitionID)
}

func (k *Kit) ProcessReceiverOperation(ctx context.Context, op payload.ReceiverOperationEnvelopeV1) error {
	return core.NewReceiverProcessor(k.repo).Process(ctx, op)
}

func (k *Kit) UserupdatesProcessUserOperation(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	return core.New(ctx, &svc.ServiceContext{Repo: k.repo}).UserupdatesProcessUserOperation(in)
}

func (k *Kit) UserupdatesGetOperationResult(ctx context.Context, in *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	return core.New(ctx, &svc.ServiceContext{Repo: k.repo}).UserupdatesGetOperationResult(in)
}

func (k *Kit) UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	return core.New(ctx, &svc.ServiceContext{Repo: k.repo}).UserupdatesGetState(in)
}

func (k *Kit) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	return core.New(ctx, &svc.ServiceContext{Repo: k.repo}).UserupdatesGetDifference(in)
}
