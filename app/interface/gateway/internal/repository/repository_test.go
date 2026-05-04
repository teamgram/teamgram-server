package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeAuthsessionClient struct {
	queryReq  *authsession.TLAuthsessionQueryAuthKey
	setReq    *authsession.TLAuthsessionSetAuthKey
	saltsReq  *authsession.TLAuthsessionGetFutureSalts
	userReq   *authsession.TLAuthsessionGetUserId
	clientReq *authsession.TLAuthsessionSetClientSessionInfo
	layerReq  *authsession.TLAuthsessionSetLayer
	stateReq  *authsession.TLAuthsessionGetAuthStateData
	queryErr  error
	setErr    error
	saltsErr  error
	userErr   error
	clientErr error
	layerErr  error
	stateErr  error
	key       *tg.AuthKeyInfo
	salts     *tg.FutureSalts
	userID    int64
	stateData *authsession.AuthKeyStateData
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

func (f *fakeAuthsessionClient) AuthsessionSetClientSessionInfo(ctx context.Context, in *authsession.TLAuthsessionSetClientSessionInfo) (*tg.Bool, error) {
	f.clientReq = in
	return tg.BoolTrue, f.clientErr
}

func (f *fakeAuthsessionClient) AuthsessionSetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) (*tg.Bool, error) {
	f.layerReq = in
	return tg.BoolTrue, f.layerErr
}

func (f *fakeAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	f.stateReq = in
	return f.stateData, f.stateErr
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

func TestRepositoryClientSessionMethodsWrapAuthsessionClient(t *testing.T) {
	client := authsession.MakeTLClientSession(&authsession.TLClientSession{
		AuthKeyId:      1001,
		Ip:             "127.0.0.1",
		Layer:          223,
		ApiId:          2040,
		DeviceModel:    "tdesktop",
		SystemVersion:  "macOS",
		AppVersion:     "5.13",
		SystemLangCode: "en-US",
		LangPack:       "tdesktop",
		LangCode:       "en",
		Proxy:          "",
		Params:         "",
	}).ToClientSession()
	fake := &fakeAuthsessionClient{
		stateData: authsession.MakeTLAuthKeyStateData(&authsession.TLAuthKeyStateData{
			AuthKeyId: 1001,
			Client:    client,
		}).ToAuthKeyStateData(),
	}
	repo := &Repository{AuthsessionClient: fake}

	if err := repo.SetClientSessionInfo(context.Background(), client); err != nil {
		t.Fatalf("SetClientSessionInfo() error = %v", err)
	}
	if fake.clientReq == nil || fake.clientReq.Data != client {
		t.Fatalf("SetClientSessionInfo request = %#v", fake.clientReq)
	}

	if err := repo.SetLayer(context.Background(), 1001, "127.0.0.1", 223); err != nil {
		t.Fatalf("SetLayer() error = %v", err)
	}
	if fake.layerReq == nil || fake.layerReq.AuthKeyId != 1001 || fake.layerReq.Ip != "127.0.0.1" || fake.layerReq.Layer != 223 {
		t.Fatalf("SetLayer request = %#v", fake.layerReq)
	}

	got, err := repo.GetClientSession(context.Background(), 1001)
	if err != nil {
		t.Fatalf("GetClientSession() error = %v", err)
	}
	if got != client || fake.stateReq == nil || fake.stateReq.AuthKeyId != 1001 {
		t.Fatalf("GetClientSession() = %#v request=%#v", got, fake.stateReq)
	}
}

func TestRepositoryGetClientSessionEmptyResult(t *testing.T) {
	repo := &Repository{AuthsessionClient: &fakeAuthsessionClient{
		stateData: authsession.MakeTLAuthKeyStateData(&authsession.TLAuthKeyStateData{AuthKeyId: 1001}).ToAuthKeyStateData(),
	}}
	got, err := repo.GetClientSession(context.Background(), 1001)
	if err != nil {
		t.Fatalf("GetClientSession() error = %v", err)
	}
	if got != nil {
		t.Fatalf("GetClientSession() = %#v, want nil", got)
	}
}

func TestRepositoryGetClientSessionIgnoresUnboundTempAuthKey(t *testing.T) {
	repo := &Repository{AuthsessionClient: &fakeAuthsessionClient{
		stateErr: authsession.ErrPermAuthKeyEmpty,
	}}
	got, err := repo.GetClientSession(context.Background(), 1001)
	if err != nil {
		t.Fatalf("GetClientSession() error = %v", err)
	}
	if got != nil {
		t.Fatalf("GetClientSession() = %#v, want nil", got)
	}
}
