package core

import (
	"bytes"
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func TestAppendDialogAuthSeqExpandsNotSourceAuthKey(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{
		AuthsessionClient: &fakeActivePermAuthKeySource{keys: []int64{11, 22, 33}},
	})

	got, visibility, err := core.expandAuthSeqTargets(1001, 22, "not_source_perm_auth_key")
	if err != nil {
		t.Fatalf("expandAuthSeqTargets() error = %v", err)
	}
	if visibility != repository.AuthSeqVisibilityNotSourcePermAuthKey {
		t.Fatalf("visibility = %q, want %q", visibility, repository.AuthSeqVisibilityNotSourcePermAuthKey)
	}
	want := []int64{11, 33}
	if len(got) != len(want) {
		t.Fatalf("targets = %#v, want %#v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("targets = %#v, want %#v", got, want)
		}
	}
}

func TestAppendDialogAuthSeqSideEffectPassesExpandedTargetsToRepository(t *testing.T) {
	body := []byte("encoded-tl-update")
	hash := payload.HashBytes(body)
	repo := &fakeUserUpdatesRepository{
		appendAuthSeqUpdateResult: &repository.AuthSeqUpdateAppendResult{
			UserID:      1001,
			OperationID: "op-auth-seq",
			Deliveries:  []repository.AuthSeqDeliveryEvent{{Seq: 7, Date: 1779234419}},
		},
	}
	core := New(context.Background(), &svc.ServiceContext{
		Repo:              repo,
		AuthsessionClient: &fakeActivePermAuthKeySource{keys: []int64{11, 22, 33}},
	})

	_, err := core.UserupdatesAppendDialogAuthSeqSideEffect(&userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
		UserId:               1001,
		SourcePermAuthKeyId:  22,
		OperationId:          "op-auth-seq",
		TargetAuthPolicy:     repository.AuthSeqVisibilityNotSourcePermAuthKey,
		PublicUpdateType:     "updatePeerSettings",
		PeerType:             payload.PeerTypeUser,
		PeerId:               2002,
		PayloadSchemaVersion: repository.AuthSeqLayer,
		Payload:              body,
		PayloadHash:          hash,
	})
	if err != nil {
		t.Fatalf("UserupdatesAppendDialogAuthSeqSideEffect() error = %v", err)
	}
	in := repo.appendAuthSeqUpdateInput
	if in.UserID != 1001 ||
		in.SourcePermAuthKeyID != 22 ||
		in.OperationID != "op-auth-seq" ||
		in.UpdateType != "updatePeerSettings" ||
		in.ReplayPolicy != repository.AuthSeqReplayPolicyDurableReplay ||
		in.VisibilityPolicy != repository.AuthSeqVisibilityNotSourcePermAuthKey ||
		in.Layer != repository.AuthSeqLayer ||
		!bytes.Equal(in.TLBytes, body) ||
		!bytes.Equal(in.PayloadHash, hash) {
		t.Fatalf("unexpected auth seq append input: %+v", in)
	}
	wantTargets := []int64{11, 33}
	if len(in.TargetPermAuthKeyIDs) != len(wantTargets) {
		t.Fatalf("targets = %#v, want %#v", in.TargetPermAuthKeyIDs, wantTargets)
	}
	for i := range wantTargets {
		if in.TargetPermAuthKeyIDs[i] != wantTargets[i] {
			t.Fatalf("targets = %#v, want %#v", in.TargetPermAuthKeyIDs, wantTargets)
		}
	}
}

var _ activePermAuthKeySource = (*fakeActivePermAuthKeySource)(nil)
