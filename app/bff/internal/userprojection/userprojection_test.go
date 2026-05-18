package userprojection

import (
	"context"
	"errors"
	"testing"

	public "github.com/teamgram/teamgram-server/v2/app/bff/projection/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUserClient struct {
	in  *userpb.TLUserGetUserProjectionBundle
	out *userpb.UserProjectionBundle
	err error
}

func (f *fakeUserClient) UserGetUserProjectionBundle(_ context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	f.in = in
	return f.out, f.err
}

func TestInternalAliasesPublicUserProjectionContract(t *testing.T) {
	if MissingExplicitInput != public.MissingExplicitInput {
		t.Fatalf("MissingExplicitInput alias drifted")
	}
	if MissingStoredReference != public.MissingStoredReference {
		t.Fatalf("MissingStoredReference alias drifted")
	}
}

func TestInternalProjectUsersDelegatesToPublicSemantics(t *testing.T) {
	client := &fakeUserClient{out: userpb.MakeTLUserProjectionBundle(&userpb.TLUserProjectionBundle{
		MissingUserIds: []int64{1002},
	}).ToUserProjectionBundle()}

	_, err := ProjectUsers(context.Background(), client, 1001, []int64{1002}, MissingExplicitInput)
	if !errors.Is(err, tg.ErrUserIdInvalid) {
		t.Fatalf("ProjectUsers() error = %v, want %v", err, tg.ErrUserIdInvalid)
	}
}
