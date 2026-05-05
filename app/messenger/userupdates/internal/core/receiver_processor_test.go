package core

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestReceiverProcessorAppliesEnvelope(t *testing.T) {
	repo := &fakeReceiverProcessorRepository{}
	processor := NewReceiverProcessor(repo)
	envelope := payload.ReceiverOperationEnvelopeV1{
		UserID:       1002,
		BucketID:     11,
		PartitionID:  7,
		OperationID:  "v1:msg:2001:receiver:1002:in",
		OpType:       payload.OpTypeSendMessage,
		PeerType:     payload.PeerTypeUser,
		PeerID:       1001,
		PayloadCodec: payload.PayloadCodecJSON,
		Payload:      []byte(`{"schema_version":1}`),
		PayloadHash:  payload.HashBytes([]byte(`{"schema_version":1}`)),
	}

	if err := processor.Process(context.Background(), envelope); err != nil {
		t.Fatalf("Process() error = %v", err)
	}
	if repo.input.UserID != envelope.UserID ||
		repo.input.OperationID != envelope.OperationID ||
		!bytes.Equal(repo.input.PayloadHash, envelope.PayloadHash) ||
		repo.input.PartitionID != envelope.PartitionID {
		t.Fatalf("unexpected repository input: %+v", repo.input)
	}
}

func TestReceiverProcessorReturnsRepositoryError(t *testing.T) {
	repo := &fakeReceiverProcessorRepository{err: errors.New("apply failed")}
	processor := NewReceiverProcessor(repo)

	err := processor.Process(context.Background(), payload.ReceiverOperationEnvelopeV1{
		UserID:      1002,
		OperationID: "v1:msg:2001:receiver:1002:in",
	})
	if err == nil {
		t.Fatal("expected error")
	}
}

type fakeReceiverProcessorRepository struct {
	input repository.ApplyUserOperationInput
	err   error
}

func (f *fakeReceiverProcessorRepository) ApplyUserOperation(_ context.Context, in repository.ApplyUserOperationInput) (*repository.ApplyUserOperationResult, error) {
	f.input = in
	return &repository.ApplyUserOperationResult{UserID: in.UserID, OperationID: in.OperationID, Pts: 1, PtsCount: 1}, f.err
}

func (f *fakeReceiverProcessorRepository) GetOperationResult(context.Context, int64, string) (*repository.OperationResult, error) {
	return nil, nil
}

func (f *fakeReceiverProcessorRepository) GetState(context.Context, int64, int64) (*repository.UserState, error) {
	return nil, nil
}

func (f *fakeReceiverProcessorRepository) GetDifference(context.Context, repository.GetDifferenceInput) (*repository.GetDifferenceResult, error) {
	return nil, nil
}

func (f *fakeReceiverProcessorRepository) ListDialogProjections(context.Context, int64, repository.DialogProjectionCursor, int32) ([]repository.DialogProjection, error) {
	return nil, nil
}

func (f *fakeReceiverProcessorRepository) GetDialogProjectionsByPeers(context.Context, int64, []repository.DialogProjectionPeer) (map[repository.DialogProjectionPeer]repository.DialogProjection, error) {
	return nil, nil
}

func (f *fakeReceiverProcessorRepository) CountVisibleDialogs(context.Context, int64) (int32, error) {
	return 0, nil
}
