package core

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/paging"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUserupdatesListDialogsReturnsProjectionDTO(t *testing.T) {
	repo := &fakeUserUpdatesRepository{
		dialogProjections: []repository.DialogProjection{
			{
				UserID:                   1001,
				PeerType:                 payload.PeerTypeUser,
				PeerID:                   2002,
				TopPeerSeq:               9,
				TopCanonicalMessageID:    7001,
				TopMessageDate:           "2026-05-05 10:00:00.000000",
				TopMessageStatus:         repository.MessageStatusLive,
				ReadInboxMaxPeerSeq:      5,
				ReadOutboxMaxPeerSeq:     6,
				UnreadCount:              3,
				UnreadMentionsCount:      1,
				UnreadReactionsCount:     2,
				UnreadMark:               true,
				PinnedPeerSeq:            4,
				PinnedCanonicalMessageID: 6001,
				HasScheduled:             true,
				AvailableMinPeerSeq:      2,
				LastPTS:                  18,
				LastPTSAt:                "2026-05-05 10:01:00.000000",
				DialogSchemaVersion:      1,
				DialogPayload:            []byte(`{"projection":true}`),
			},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: repo})

	got, err := core.UserupdatesListDialogs(&userupdates.TLUserupdatesListDialogs{
		UserId:         1001,
		TopMessageDate: 0,
		TopPeerSeq:     0,
		PeerType:       0,
		PeerId:         0,
		Limit:          10,
	})
	if err != nil {
		t.Fatalf("UserupdatesListDialogs() error = %v", err)
	}
	if repo.dialogListUserID != 1001 || repo.dialogListLimit != 10 {
		t.Fatalf("unexpected repository call: user_id=%d limit=%d", repo.dialogListUserID, repo.dialogListLimit)
	}
	if len(got.Projections) != 1 {
		t.Fatalf("projection count = %d, want 1", len(got.Projections))
	}
	projection := got.Projections[0]
	if projection.PeerType != payload.PeerTypeUser ||
		projection.PeerId != 2002 ||
		projection.TopPeerSeq != 9 ||
		projection.TopCanonicalMessageId != 7001 ||
		projection.UnreadCount != 3 ||
		!projection.UnreadMark ||
		!projection.HasScheduled ||
		!bytes.Equal(projection.DialogPayload, []byte(`{"projection":true}`)) {
		t.Fatalf("unexpected projection: %+v", projection)
	}
	if got.NextTopMessageDate == 0 || got.NextTopPeerSeq != 9 || got.NextPeerType != payload.PeerTypeUser || got.NextPeerId != 2002 || got.Exhausted != tg.BoolTrueClazz {
		t.Fatalf("unexpected next cursor/exhausted fields: %+v", got)
	}
}

func TestUserupdatesGetDialogsByPeersRejectsOversizedPeerList(t *testing.T) {
	peers := make([]userupdates.DialogProjectionPeerClazz, paging.DialogMaxHydratePeersPerRequest+1)
	for i := range peers {
		peers[i] = userupdates.MakeTLDialogProjectionPeer(&userupdates.TLDialogProjectionPeer{
			PeerType: payload.PeerTypeUser,
			PeerId:   int64(i + 1),
		})
	}
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeUserUpdatesRepository{}})

	_, err := core.UserupdatesGetDialogsByPeers(&userupdates.TLUserupdatesGetDialogsByPeers{
		UserId: 1001,
		Peers:  peers,
	})
	if !errors.Is(err, userupdates.ErrDialogQueryTooLarge) {
		t.Fatalf("UserupdatesGetDialogsByPeers() error = %v, want ErrDialogQueryTooLarge", err)
	}
}

func TestAppendDialogAuthSeqSideEffectRequiresOperationID(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{Repo: &fakeUserUpdatesRepository{}})

	_, err := core.UserupdatesAppendDialogAuthSeqSideEffect(&userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
		UserId:               1001,
		SourcePermAuthKeyId:  9001,
		OperationId:          "",
		TargetAuthPolicy:     "not_source_perm_auth_key",
		PublicUpdateType:     "draft_clear",
		PeerType:             payload.PeerTypeUser,
		PeerId:               2002,
		PayloadSchemaVersion: 1,
		Payload:              []byte(`{"ok":true}`),
		PayloadHash:          payload.HashBytes([]byte(`{"ok":true}`)),
	})
	if !errors.Is(err, userupdates.ErrOperationTerminal) {
		t.Fatalf("UserupdatesAppendDialogAuthSeqSideEffect() error = %v, want ErrOperationTerminal", err)
	}
}
