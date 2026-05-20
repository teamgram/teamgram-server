package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/svc"
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

var _ activePermAuthKeySource = (*fakeActivePermAuthKeySource)(nil)
