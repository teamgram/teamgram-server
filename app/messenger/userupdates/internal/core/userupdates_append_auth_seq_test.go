package core

import (
	"bytes"
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
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

func TestAppendDialogAuthSeqExpandsAllPolicy(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{
		AuthsessionClient: &fakeActivePermAuthKeySource{keys: []int64{11, 22, 33}},
	})

	got, visibility, err := core.expandAuthSeqTargets(1001, 22, "all")
	if err != nil {
		t.Fatalf("expandAuthSeqTargets() error = %v", err)
	}
	if visibility != repository.AuthSeqVisibilityAllUserAuthKeys {
		t.Fatalf("visibility = %q, want %q", visibility, repository.AuthSeqVisibilityAllUserAuthKeys)
	}
	want := []int64{11, 22, 33}
	if len(got) != len(want) {
		t.Fatalf("targets = %#v, want %#v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("targets = %#v, want %#v", got, want)
		}
	}
}

func TestDialogAuthSeqTLUpdateBuildsUpdate(t *testing.T) {
	body := []byte(`{"schema_version":1,"event_kind":"draft_saved","peer_type":1,"peer_id":1002}`)
	got, err := dialogAuthSeqTLUpdate(&userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
		UserId:           1001,
		OperationId:      "op-dialog-tl",
		PublicUpdateType: "draft_saved",
		PeerType:         1,
		PeerId:           1002,
		Payload:          body,
		PayloadHash:      payload.HashBytes(body),
	})
	if err != nil {
		t.Fatalf("dialogAuthSeqTLUpdate() error = %v", err)
	}
	if _, ok := got.(*tg.TLUpdateDraftMessage); !ok {
		t.Fatalf("update = %T, want *tg.TLUpdateDraftMessage", got)
	}
	encoded, err := iface.EncodeObject(got, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	if len(encoded) == 0 {
		t.Fatalf("encoded TL update is empty")
	}
}

func TestAppendDialogAuthSeqSideEffectPassesExpandedTargetsToRepository(t *testing.T) {
	body := []byte(`{"schema_version":1,"event_kind":"draft_saved","peer_type":1,"peer_id":2002}`)
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
		PublicUpdateType:     "draft_saved",
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
		in.UpdateType != "draft_saved" ||
		in.ReplayPolicy != repository.AuthSeqReplayPolicyDurableReplay ||
		in.VisibilityPolicy != repository.AuthSeqVisibilityNotSourcePermAuthKey ||
		in.Layer != repository.AuthSeqLayer {
		t.Fatalf("unexpected auth seq append input: %+v", in)
	}
	wantUpdate, err := dialogAuthSeqTLUpdate(&userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
		PublicUpdateType: "draft_saved",
		PeerType:         payload.PeerTypeUser,
		PeerId:           2002,
		Payload:          body,
		PayloadHash:      hash,
	})
	if err != nil {
		t.Fatalf("dialogAuthSeqTLUpdate() error = %v", err)
	}
	wantTLBytes, err := iface.EncodeObject(wantUpdate, repository.AuthSeqLayer)
	if err != nil {
		t.Fatalf("EncodeObject() error = %v", err)
	}
	if !bytes.Equal(in.TLBytes, wantTLBytes) {
		t.Fatalf("TLBytes = %x, want %x", in.TLBytes, wantTLBytes)
	}
	if !bytes.Equal(in.PayloadHash, payload.HashBytes(wantTLBytes)) {
		t.Fatalf("PayloadHash = %x, want %x", in.PayloadHash, payload.HashBytes(wantTLBytes))
	}
	if bytes.Equal(in.TLBytes, body) || bytes.Equal(in.PayloadHash, hash) {
		t.Fatalf("auth seq append input used original JSON payload: %+v", in)
	}
	decoded, err := iface.DecodeObject(bin.NewDecoder(in.TLBytes))
	if err != nil {
		t.Fatalf("DecodeObject(TLBytes) error = %v", err)
	}
	if _, ok := decoded.(*tg.TLUpdateDraftMessage); !ok {
		t.Fatalf("decoded TLBytes = %T, want *tg.TLUpdateDraftMessage", decoded)
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
