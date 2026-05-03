package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAccountResetAuthorizationZeroHashReturnsBoolFalse(t *testing.T) {
	c := newAccountCoreForTest(nil, &fakeAccountAuthsessionClient{
		resetAuthorization: func(context.Context, *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
			t.Fatal("AuthsessionResetAuthorization should not be called for zero hash")
			return nil, nil
		},
	}, nil)

	got, err := c.AccountResetAuthorization(&tg.TLAccountResetAuthorization{Hash: 0})
	if err != nil {
		t.Fatalf("AccountResetAuthorization error = %v", err)
	}
	if got != tg.BoolFalse {
		t.Fatalf("result = %v, want BoolFalse", got)
	}
}

func TestAccountResetAuthorizationDelegatesToAuthsession(t *testing.T) {
	var gotReq *authsession.TLAuthsessionResetAuthorization
	c := newAccountCoreForTest(nil, &fakeAccountAuthsessionClient{
		resetAuthorization: func(_ context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
			gotReq = in
			return &authsession.VectorLong{Datas: []int64{111, 222}}, nil
		},
	}, nil)

	got, err := c.AccountResetAuthorization(&tg.TLAccountResetAuthorization{Hash: 77})
	if err != nil {
		t.Fatalf("AccountResetAuthorization error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("result = %v, want BoolTrue", got)
	}
	if gotReq == nil || gotReq.UserId != 1001 || gotReq.AuthKeyId != 9001 || gotReq.Hash != 77 {
		t.Fatalf("reset request = %+v, want user_id 1001 auth_key_id 9001 hash 77", gotReq)
	}
}

func TestAccountDeleteAccountDeletesUserAndAuthsessions(t *testing.T) {
	var deletedUsername string
	var deletedUser *userpb.TLUserDeleteUser
	var resetReq *authsession.TLAuthsessionResetAuthorization
	var unbindReq *authsession.TLAuthsessionUnbindAuthKeyUser
	c := newAccountCoreForTest(&fakeAccountUserClient{
		getUserDataByID: func(_ context.Context, in *userpb.TLUserGetUserDataById) (*tg.UserData, error) {
			if in.UserId != 1001 {
				t.Fatalf("get user data user_id = %d, want 1001", in.UserId)
			}
			return tg.MakeTLUserData(&tg.TLUserData{
				Id:       1001,
				Phone:    "8613000000000",
				Username: "self",
			}).ToUserData(), nil
		},
		deleteUsername: func(_ context.Context, in *userpb.TLUserDeleteUsername) (*tg.Bool, error) {
			deletedUsername = in.Username
			return tg.BoolTrue, nil
		},
		deleteUser: func(_ context.Context, in *userpb.TLUserDeleteUser) (*tg.Bool, error) {
			deletedUser = in
			return tg.BoolTrue, nil
		},
	}, &fakeAccountAuthsessionClient{
		resetAuthorization: func(_ context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
			resetReq = in
			return &authsession.VectorLong{Datas: []int64{111, 222}}, nil
		},
		unbindAuthKeyUser: func(_ context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error) {
			unbindReq = in
			return tg.BoolTrue, nil
		},
	}, nil)

	got, err := c.AccountDeleteAccount(&tg.TLAccountDeleteAccount{Reason: "test cleanup"})
	if err != nil {
		t.Fatalf("AccountDeleteAccount error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("result = %v, want BoolTrue", got)
	}
	if deletedUsername != "self" {
		t.Fatalf("deleted username = %q, want self", deletedUsername)
	}
	if deletedUser == nil || deletedUser.UserId != 1001 || deletedUser.Reason != "test cleanup" || deletedUser.Phone != "8613000000000" {
		t.Fatalf("delete user request = %+v, want user_id 1001 reason test cleanup phone 8613000000000", deletedUser)
	}
	if resetReq == nil || resetReq.UserId != 1001 || resetReq.AuthKeyId != 0 || resetReq.Hash != 0 {
		t.Fatalf("reset request = %+v, want user_id 1001 auth_key_id 0 hash 0", resetReq)
	}
	if unbindReq == nil || unbindReq.UserId != 1001 || unbindReq.AuthKeyId != 0 {
		t.Fatalf("unbind request = %+v, want user_id 1001 auth_key_id 0", unbindReq)
	}
}
