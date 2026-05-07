package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/svc"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type projectionRepoStub struct {
	inViewerIds []int64
	inTargetIds []int64
	inWithFacts bool
	result      *repository.UserProjectionBundle
	err         error
}

func (s *projectionRepoStub) GetUserProjectionBundle(ctx context.Context, viewerIds []int64, targetIds []int64, withFacts bool) (*repository.UserProjectionBundle, error) {
	s.inViewerIds = append([]int64(nil), viewerIds...)
	s.inTargetIds = append([]int64(nil), targetIds...)
	s.inWithFacts = withFacts
	return s.result, s.err
}

func TestUserGetUserProjectionBundleReturnsViewerUsers(t *testing.T) {
	firstName := "T"
	repo := &projectionRepoStub{result: &repository.UserProjectionBundle{
		Facts: []tg.ImmutableUserClazz{tg.MakeTLImmutableUser(&tg.TLImmutableUser{
			User: tg.MakeTLUserData(&tg.TLUserData{Id: 1001, AccessHash: 11, FirstName: "A", LastName: "B"}),
		})},
		ViewerUsers: []repository.ViewerUsers{{
			ViewerUserId: 1001,
			Users:        []tg.UserClazz{tg.MakeTLUser(&tg.TLUser{Id: 1002, FirstName: &firstName})},
		}},
		MissingUserIds: []int64{1003},
	}}
	core := New(context.Background(), &svc.ServiceContext{UserProjectionRepo: repo})

	got, err := core.UserGetUserProjectionBundle(&userpb.TLUserGetUserProjectionBundle{
		WithFacts:     true,
		ViewerUserIds: []int64{1001},
		TargetUserIds: []int64{1002, 1003},
	})
	if err != nil {
		t.Fatalf("UserGetUserProjectionBundle() error = %v", err)
	}
	if !repo.inWithFacts || len(repo.inViewerIds) != 1 || repo.inViewerIds[0] != 1001 || len(repo.inTargetIds) != 2 {
		t.Fatalf("repo input viewer=%v target=%v withFacts=%v", repo.inViewerIds, repo.inTargetIds, repo.inWithFacts)
	}
	if got == nil || len(got.ViewerUsers) != 1 || got.ViewerUsers[0].ViewerUserId != 1001 || len(got.ViewerUsers[0].Users) != 1 {
		t.Fatalf("projection bundle = %#v", got)
	}
	if len(got.MissingUserIds) != 1 || got.MissingUserIds[0] != 1003 {
		t.Fatalf("missing ids = %v", got.MissingUserIds)
	}
	if len(got.Facts) != 1 {
		t.Fatalf("facts = %v", got.Facts)
	}
}

func TestUserGetUserProjectionBundleMapsInvalidShape(t *testing.T) {
	core := New(context.Background(), &svc.ServiceContext{UserProjectionRepo: &projectionRepoStub{err: userpb.ErrUserInvalidArgument}})
	_, err := core.UserGetUserProjectionBundle(&userpb.TLUserGetUserProjectionBundle{})
	if !errors.Is(err, userpb.ErrUserInvalidArgument) {
		t.Fatalf("error = %v, want %v", err, userpb.ErrUserInvalidArgument)
	}
}

func TestUserGetMutableUsersV2UsesProjectionBundleFacts(t *testing.T) {
	repo := &projectionRepoStub{result: &repository.UserProjectionBundle{
		Facts: []tg.ImmutableUserClazz{tg.MakeTLImmutableUser(&tg.TLImmutableUser{
			User: tg.MakeTLUserData(&tg.TLUserData{Id: 1001, AccessHash: 11, FirstName: "A"}),
		})},
	}}
	core := New(context.Background(), &svc.ServiceContext{UserProjectionRepo: repo})

	got, err := core.UserGetMutableUsersV2(&userpb.TLUserGetMutableUsersV2{
		Id:      []int64{1001},
		Privacy: true,
		To:      []int64{2001},
	})
	if err != nil {
		t.Fatalf("UserGetMutableUsersV2() error = %v", err)
	}
	if len(got.Users) != 1 || len(repo.inViewerIds) != 1 || repo.inViewerIds[0] != 2001 || !repo.inWithFacts {
		t.Fatalf("mutable=%#v repo viewer=%v target=%v withFacts=%v", got, repo.inViewerIds, repo.inTargetIds, repo.inWithFacts)
	}
}
