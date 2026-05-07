package repository

import (
	"context"
	"testing"
	"time"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDialogSideEffectStore struct {
	rows      []DialogSideEffect
	completed []int64
}

func (s *fakeDialogSideEffectStore) ClaimDialogSideEffectsByKind(ctx context.Context, kind string, now time.Time, limit int32) ([]DialogSideEffect, error) {
	return s.rows, nil
}

func (s *fakeDialogSideEffectStore) MarkDialogSideEffectCompleted(ctx context.Context, sideEffectID int64) error {
	s.completed = append(s.completed, sideEffectID)
	return nil
}

func (s *fakeDialogSideEffectStore) MarkDialogSideEffectRetryableFailure(ctx context.Context, sideEffectID int64, errCode string, now time.Time) error {
	return nil
}

func (s *fakeDialogSideEffectStore) MarkDialogSideEffectBlocked(ctx context.Context, sideEffectID int64, errCode string) error {
	return nil
}

type fakeDialogSideEffectClient struct {
	got *dialogpb.TLDialogUpsertSavedDialogFromMessage
}

func (c *fakeDialogSideEffectClient) DialogUpsertSavedDialogFromMessage(ctx context.Context, in *dialogpb.TLDialogUpsertSavedDialogFromMessage) (*tg.Bool, error) {
	c.got = in
	return tg.BoolTrue, nil
}

func TestDialogSideEffectWorkerPublishesSavedDialogTop(t *testing.T) {
	sourceDate := int64(1710000000)
	store := &fakeDialogSideEffectStore{rows: []DialogSideEffect{{
		SideEffectID:             1001,
		Kind:                     DialogSideEffectKindUpsertSavedDialogFromMessage,
		UserID:                   2001,
		PeerType:                 1,
		PeerID:                   3001,
		SourceMessageDate:        sourceDate,
		SourcePeerSeq:            41,
		SourceCanonicalMessageID: 9001,
		Payload:                  []byte(`{"schema_version":1}`),
	}}}
	client := &fakeDialogSideEffectClient{}
	worker := NewDialogSideEffectWorker(store, client, DialogSideEffectWorkerOptions{BatchSize: 10})

	worker.drain(context.Background())

	if client.got == nil {
		t.Fatal("DialogUpsertSavedDialogFromMessage was not called")
	}
	if client.got.UserId != 2001 || client.got.PeerType != 1 || client.got.PeerId != 3001 {
		t.Fatalf("dialog upsert peer = user:%d type:%d id:%d", client.got.UserId, client.got.PeerType, client.got.PeerId)
	}
	if client.got.TopPeerSeq != 41 || client.got.TopCanonicalMessageId != 9001 || client.got.TopMessageDate != int32(sourceDate) {
		t.Fatalf("dialog upsert top = seq:%d canonical:%d date:%d", client.got.TopPeerSeq, client.got.TopCanonicalMessageId, client.got.TopMessageDate)
	}
	if len(store.completed) != 1 || store.completed[0] != 1001 {
		t.Fatalf("completed = %v, want [1001]", store.completed)
	}
}
