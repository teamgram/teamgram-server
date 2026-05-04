package core

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

type fakeActivePermAuthKeySource struct {
	gotUserID int64
	keys      []int64
	err       error
}

func (f *fakeActivePermAuthKeySource) AuthsessionGetPermAuthKeyIds(_ context.Context, in *authsession.TLAuthsessionGetPermAuthKeyIds) (*authsession.VectorLong, error) {
	f.gotUserID = in.UserId
	if f.err != nil {
		return nil, f.err
	}
	return &authsession.VectorLong{Datas: f.keys}, nil
}

func TestResolveAuthSeqNotMeTargetsUsesAuthsessionSnapshot(t *testing.T) {
	source := &fakeActivePermAuthKeySource{keys: []int64{9001, 9002, 0, 9002, 9003}}

	got, err := resolveAuthSeqNotMeTargets(context.Background(), source, 1001, 9001)
	if err != nil {
		t.Fatalf("resolveAuthSeqNotMeTargets() error = %v", err)
	}
	if source.gotUserID != 1001 {
		t.Fatalf("authsession user_id = %d, want 1001", source.gotUserID)
	}
	want := []int64{9002, 9003}
	if len(got) != len(want) {
		t.Fatalf("targets = %#v, want %#v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("targets = %#v, want %#v", got, want)
		}
	}
}

func TestResolveAuthSeqNotMeTargetsWrapsLookupFailure(t *testing.T) {
	source := &fakeActivePermAuthKeySource{err: errors.New("authsession unavailable")}

	_, err := resolveAuthSeqNotMeTargets(context.Background(), source, 1001, 9001)
	if !errors.Is(err, userupdates.ErrUserupdatesStorage) || !strings.Contains(err.Error(), "target_lookup") {
		t.Fatalf("resolveAuthSeqNotMeTargets() error = %v, want storage target_lookup", err)
	}
}
