package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthsessionClient struct {
	queryReq *authsession.TLAuthsessionQueryAuthKey
	setReq   *authsession.TLAuthsessionSetAuthKey
	saltsReq *authsession.TLAuthsessionGetFutureSalts
	userReq  *authsession.TLAuthsessionGetUserId
	queryErr error
	setErr   error
	saltsErr error
	userErr  error
	key      *tg.AuthKeyInfo
	salts    *tg.FutureSalts
	userID   int64
}

func (f *fakeAuthsessionClient) AuthsessionQueryAuthKey(ctx context.Context, in *authsession.TLAuthsessionQueryAuthKey) (*tg.AuthKeyInfo, error) {
	f.queryReq = in
	return f.key, f.queryErr
}

func (f *fakeAuthsessionClient) AuthsessionSetAuthKey(ctx context.Context, in *authsession.TLAuthsessionSetAuthKey) (*tg.Bool, error) {
	f.setReq = in
	return tg.BoolTrue, f.setErr
}

func (f *fakeAuthsessionClient) AuthsessionGetFutureSalts(ctx context.Context, in *authsession.TLAuthsessionGetFutureSalts) (*tg.FutureSalts, error) {
	f.saltsReq = in
	return f.salts, f.saltsErr
}

func (f *fakeAuthsessionClient) AuthsessionGetUserId(ctx context.Context, in *authsession.TLAuthsessionGetUserId) (*tg.Int64, error) {
	f.userReq = in
	return tg.MakeInt64(f.userID), f.userErr
}

func TestRepositoryAuthKeyMethodsWrapAuthsessionClient(t *testing.T) {
	key := tg.MakeTLAuthKeyInfo(&tg.AuthKeyInfo{
		AuthKeyId: 123,
		AuthKey:   []byte("auth-key"),
	}).ToAuthKeyInfo()
	salt := tg.MakeTLFutureSalt(&tg.FutureSalt{
		ValidSince: 1,
		ValidUntil: 2,
		Salt:       3,
	}).ToFutureSalt()
	salts := tg.MakeTLFutureSalts(&tg.FutureSalts{
		ReqMsgId: 7,
		Now:      8,
		Salts:    []*tg.FutureSalt{salt},
	}).ToFutureSalts()
	fake := &fakeAuthsessionClient{key: key, salts: salts, userID: 1001}
	repo := &Repository{AuthsessionClient: fake}

	gotKey, err := repo.QueryAuthKey(context.Background(), 123)
	if err != nil {
		t.Fatalf("QueryAuthKey() error = %v", err)
	}
	if gotKey != key || fake.queryReq == nil || fake.queryReq.AuthKeyId != 123 {
		t.Fatalf("QueryAuthKey() = %v, request = %v", gotKey, fake.queryReq)
	}

	if err := repo.SetAuthKey(context.Background(), key, salt, 60); err != nil {
		t.Fatalf("SetAuthKey() error = %v", err)
	}
	if fake.setReq == nil || fake.setReq.AuthKey != key || fake.setReq.FutureSalt != salt || fake.setReq.ExpiresIn != 60 {
		t.Fatalf("SetAuthKey() request = %v", fake.setReq)
	}

	gotSalts, err := repo.GetFutureSalts(context.Background(), 123, 4)
	if err != nil {
		t.Fatalf("GetFutureSalts() error = %v", err)
	}
	if gotSalts != salts || fake.saltsReq == nil || fake.saltsReq.AuthKeyId != 123 || fake.saltsReq.Num != 4 {
		t.Fatalf("GetFutureSalts() = %v, request = %v", gotSalts, fake.saltsReq)
	}

	gotUserID, err := repo.GetUserId(context.Background(), 123)
	if err != nil {
		t.Fatalf("GetUserId() error = %v", err)
	}
	if gotUserID != 1001 || fake.userReq == nil || fake.userReq.AuthKeyId != 123 {
		t.Fatalf("GetUserId() = %d, request = %v", gotUserID, fake.userReq)
	}
}

func TestRepositoryAuthKeyMethodsRequireConfiguredClient(t *testing.T) {
	repo := &Repository{}
	_, err := repo.QueryAuthKey(context.Background(), 123)
	if err == nil {
		t.Fatal("QueryAuthKey() error is nil")
	}
}

func TestRepositoryAuthKeyMethodsWrapErrors(t *testing.T) {
	want := errors.New("downstream failed")
	repo := &Repository{AuthsessionClient: &fakeAuthsessionClient{queryErr: want}}
	_, err := repo.QueryAuthKey(context.Background(), 123)
	if !errors.Is(err, want) {
		t.Fatalf("QueryAuthKey() error = %v, want wrapping %v", err, want)
	}
}
