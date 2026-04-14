package sess

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/session/internal/dao"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	authsessionclient "github.com/teamgram/teamgram-server/app/service/authsession/client"
)

func TestRPCErrorForAuthStateUsesRestartForTransientStates(t *testing.T) {
	for _, state := range []int{
		mtproto.AuthStateNew,
		mtproto.AuthStateWaitInit,
		mtproto.AuthStateUnauthorized,
	} {
		err := rpcErrorForAuthState(state)
		if !errors.Is(err, mtproto.ErrAuthRestart) {
			t.Fatalf("state %d: got %v, want AUTH_RESTART", state, err)
		}
	}
}

func TestRPCErrorForAuthStateUsesAuthKeyUnregisteredForTerminalStates(t *testing.T) {
	for _, state := range []int{
		mtproto.AuthStateUnknown,
		mtproto.AuthStateLogout,
		mtproto.AuthStateDeleted,
	} {
		err := rpcErrorForAuthState(state)
		if !errors.Is(err, mtproto.ErrAuthKeyUnregistered) {
			t.Fatalf("state %d: got %v, want AUTH_KEY_UNREGISTERED", state, err)
		}
	}
}

func TestShouldRefreshAuthState(t *testing.T) {
	cases := []struct {
		state int
		want  bool
	}{
		{mtproto.AuthStateNew, true},
		{mtproto.AuthStateWaitInit, true},
		{mtproto.AuthStateUnauthorized, true},
		{mtproto.AuthStateNeedPassword, false},
		{mtproto.AuthStateNormal, false},
		{mtproto.AuthStateLogout, false},
	}

	for _, c := range cases {
		if got := shouldRefreshAuthState(c.state); got != c.want {
			t.Fatalf("state %d: got %v, want %v", c.state, got, c.want)
		}
	}
}

func TestRefreshAuthStateIfNeededPromotesWrapperToNormal(t *testing.T) {
	const (
		authKeyID = 1001
		userID    = 2002
	)

	fakeClient := &fakeAuthsessionClient{
		authStateData: authsession.MakeTLAuthKeyStateData(&authsession.AuthKeyStateData{
			AuthKeyId: authKeyID,
			KeyState:  mtproto.AuthStateNormal,
			UserId:    userID,
			Client: authsession.MakeTLClientSession(&authsession.ClientSession{
				AuthKeyId: authKeyID,
				Layer:     200,
			}).To_ClientSession(),
		}).To_AuthKeyStateData(),
	}
	manager := NewMainAuthWrapperManager(&dao.Dao{
		AuthsessionClient: fakeClient,
	})

	wrapper := NewMainAuthWrapper(authKeyID, 0, mtproto.AuthStateUnauthorized, nil, 0, manager)
	defer wrapper.Stop()

	wrapper.refreshAuthStateIfNeeded(context.Background())

	if wrapper.state != mtproto.AuthStateNormal {
		t.Fatalf("wrapper.state = %d, want %d", wrapper.state, mtproto.AuthStateNormal)
	}
	if wrapper.AuthUserId != userID {
		t.Fatalf("wrapper.AuthUserId = %d, want %d", wrapper.AuthUserId, userID)
	}
	if wrapper.client == nil || wrapper.client.GetAuthKeyId() != authKeyID {
		t.Fatalf("wrapper.client = %#v", wrapper.client)
	}
	if fakeClient.getAuthStateDataCalls != 1 {
		t.Fatalf("AuthsessionGetAuthStateData calls = %d, want 1", fakeClient.getAuthStateDataCalls)
	}
}

var _ authsessionclient.AuthsessionClient = (*fakeAuthsessionClient)(nil)

type fakeAuthsessionClient struct {
	authStateData         *authsession.AuthKeyStateData
	getAuthStateDataCalls int
}

func (f *fakeAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	f.getAuthStateDataCalls++
	return f.authStateData, nil
}

func (f *fakeAuthsessionClient) AuthsessionGetAuthorizations(context.Context, *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionResetAuthorization(context.Context, *authsession.TLAuthsessionResetAuthorization) (*authsession.Vector_Long, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetLayer(context.Context, *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetLangPack(context.Context, *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetClient(context.Context, *authsession.TLAuthsessionGetClient) (*mtproto.String, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetLangCode(context.Context, *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetUserId(context.Context, *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetPushSessionId(context.Context, *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetFutureSalts(context.Context, *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionQueryAuthKey(context.Context, *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionSetAuthKey(context.Context, *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionBindAuthKeyUser(context.Context, *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionUnbindAuthKeyUser(context.Context, *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetPermAuthKeyId(context.Context, *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionBindTempAuthKey(context.Context, *authsession.TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionSetClientSessionInfo(context.Context, *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionGetAuthorization(context.Context, *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionSetLayer(context.Context, *authsession.TLAuthsessionSetLayer) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionSetInitConnection(context.Context, *authsession.TLAuthsessionSetInitConnection) (*mtproto.Bool, error) {
	return nil, nil
}
func (f *fakeAuthsessionClient) AuthsessionSetAndroidPushSessionId(context.Context, *authsession.TLAuthsessionSetAndroidPushSessionId) (*mtproto.Bool, error) {
	return nil, nil
}
