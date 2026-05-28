// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthPhoneRegistrationFlow(t *testing.T) {
	repo := newFakeAuthRepository()
	c := newTestAuthorizationCore(t, repo, 1001001, 0)

	sentCode, err := c.AuthSendCode(&tg.TLAuthSendCode{
		PhoneNumber: "+86 137 1111 2222",
		ApiId:       1,
		ApiHash:     "test-api-hash",
		Settings:    tg.MakeTLCodeSettings(&tg.TLCodeSettings{}),
	})
	if err != nil {
		t.Fatalf("AuthSendCode error = %v", err)
	}
	code, ok := sentCode.ToAuthSentCode()
	if !ok {
		t.Fatalf("AuthSendCode returned %s, want auth.sentCode", sentCode.ClazzName())
	}
	if code.PhoneCodeHash == "" {
		t.Fatal("AuthSendCode returned empty phone_code_hash")
	}
	if _, ok := code.Type.(*tg.TLAuthSentCodeTypeSms); !ok {
		t.Fatalf("AuthSendCode type = %T, want *tg.TLAuthSentCodeTypeSms", code.Type)
	}
	if err := sentCode.Validate(223); err != nil {
		t.Fatalf("AuthSendCode result Validate(223) error = %v", err)
	}

	phoneCode := startupAuthPhoneCode
	signIn, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613711112222",
		PhoneCodeHash: code.PhoneCodeHash,
		PhoneCode:     &phoneCode,
	})
	if err != nil {
		t.Fatalf("AuthSignIn unregistered error = %v", err)
	}
	if _, ok := signIn.ToAuthAuthorizationSignUpRequired(); !ok {
		t.Fatalf("AuthSignIn unregistered returned %s, want auth.authorizationSignUpRequired", signIn.ClazzName())
	}

	signUp, err := c.AuthSignUp(&tg.TLAuthSignUp{
		PhoneNumber:   "+8613711112222",
		PhoneCodeHash: code.PhoneCodeHash,
		FirstName:     "Test",
		LastName:      "User",
	})
	if err != nil {
		t.Fatalf("AuthSignUp error = %v", err)
	}
	authorization, ok := signUp.ToAuthAuthorization()
	if !ok {
		t.Fatalf("AuthSignUp returned %s, want auth.authorization", signUp.ClazzName())
	}
	if authorization.User == nil {
		t.Fatal("AuthSignUp returned nil user")
	}
	if err := signUp.Validate(223); err != nil {
		t.Fatalf("AuthSignUp result Validate(223) error = %v", err)
	}
	if repo.createUserCalls != 1 {
		t.Fatalf("create user calls = %d, want 1", repo.createUserCalls)
	}
	if repo.bindCalls != 1 {
		t.Fatalf("bind calls = %d, want 1", repo.bindCalls)
	}
	if repo.projectCalls != 1 {
		t.Fatalf("project calls = %d, want 1", repo.projectCalls)
	}
	if repo.boundAuthKeyID != 1001001 || repo.boundUserID != fakeAuthUserID {
		t.Fatalf("bound auth/user = %d/%d, want %d/%d", repo.boundAuthKeyID, repo.boundUserID, int64(1001001), int64(fakeAuthUserID))
	}

	c = newTestAuthorizationCore(t, repo, 1001001, fakeAuthUserID)
	loggedOut, err := c.AuthLogOut(&tg.TLAuthLogOut{})
	if err != nil {
		t.Fatalf("AuthLogOut error = %v", err)
	}
	if loggedOut == nil {
		t.Fatal("AuthLogOut returned nil")
	}
	if err := loggedOut.Validate(223); err != nil {
		t.Fatalf("AuthLogOut result Validate(223) error = %v", err)
	}
	if repo.unbindCalls != 1 {
		t.Fatalf("unbind calls = %d, want 1", repo.unbindCalls)
	}
}

func TestAuthSignInExistingUserBindsAuthKey(t *testing.T) {
	repo := newFakeAuthRepository()
	repo.usersByPhone["8613711112222"] = makeFakeImmutableUser(fakeAuthUserID, "8613711112222", "CN", "Test", "User")
	c := newTestAuthorizationCore(t, repo, 3003003, 0)

	sentCode, err := c.AuthSendCode(&tg.TLAuthSendCode{
		PhoneNumber: "+8613711112222",
		ApiId:       1,
		ApiHash:     "test-api-hash",
		Settings:    tg.MakeTLCodeSettings(&tg.TLCodeSettings{}),
	})
	if err != nil {
		t.Fatalf("AuthSendCode error = %v", err)
	}
	code, ok := sentCode.ToAuthSentCode()
	if !ok {
		t.Fatalf("AuthSendCode returned %s, want auth.sentCode", sentCode.ClazzName())
	}

	phoneCode := startupAuthPhoneCode
	got, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+86 137 1111 2222",
		PhoneCodeHash: code.PhoneCodeHash,
		PhoneCode:     &phoneCode,
	})
	if err != nil {
		t.Fatalf("AuthSignIn existing user error = %v", err)
	}
	authorization, ok := got.ToAuthAuthorization()
	if !ok {
		t.Fatalf("AuthSignIn existing user returned %s, want auth.authorization", got.ClazzName())
	}
	if authorization.User == nil {
		t.Fatal("AuthSignIn existing user returned nil user")
	}
	if repo.bindCalls != 1 {
		t.Fatalf("bind calls = %d, want 1", repo.bindCalls)
	}
	if repo.projectCalls != 1 {
		t.Fatalf("project calls = %d, want 1", repo.projectCalls)
	}
	if repo.boundAuthKeyID != 3003003 || repo.boundUserID != fakeAuthUserID {
		t.Fatalf("bound auth/user = %d/%d, want %d/%d", repo.boundAuthKeyID, repo.boundUserID, int64(3003003), int64(fakeAuthUserID))
	}
}

func newTestAuthorizationCore(t *testing.T, repo *fakeAuthRepository, authKeyID int64, userID int64) *AuthorizationCore {
	t.Helper()

	ctx, err := metadata.RpcMetadataToOutgoing(context.Background(), &metadata.RpcMetadata{
		PermAuthKeyId: authKeyID,
		AuthId:        authKeyID,
		SessionId:     2002002,
		UserId:        userID,
		Layer:         223,
		ClientAddr:    "127.0.0.1",
	})
	if err != nil {
		t.Fatalf("metadata context error = %v", err)
	}

	svcCtx := svc.NewServiceContext(config.Config{})
	svcCtx.Repo = repo
	return New(ctx, svcCtx)
}

const fakeAuthUserID int64 = 424242

type fakeAuthRepository struct {
	usersByPhone map[string]*tg.ImmutableUser

	createUserCalls int
	bindCalls       int
	unbindCalls     int
	projectCalls    int

	boundAuthKeyID int64
	boundUserID    int64

	bindTempCalls            int
	bindTempPermAuthKeyID    int64
	bindTempNonce            int64
	bindTempExpiresAt        int32
	bindTempEncryptedMessage []byte
	bindTempErr              error
}

func newFakeAuthRepository() *fakeAuthRepository {
	return &fakeAuthRepository{
		usersByPhone: map[string]*tg.ImmutableUser{},
	}
}

func (r *fakeAuthRepository) GetUserByPhone(ctx context.Context, phone string) (*tg.ImmutableUser, error) {
	_ = ctx
	if user, ok := r.usersByPhone[phone]; ok {
		return user, nil
	}
	return nil, repository.ErrUserNotFound
}

func (r *fakeAuthRepository) CreateUser(ctx context.Context, secretKeyId int64, phone string, countryCode string, firstName string, lastName string) (*tg.ImmutableUser, error) {
	_ = ctx
	_ = secretKeyId
	r.createUserCalls++
	user := makeFakeImmutableUser(fakeAuthUserID, phone, countryCode, firstName, lastName)
	r.usersByPhone[phone] = user
	return user, nil
}

func (r *fakeAuthRepository) ProjectSelfUser(ctx context.Context, userId int64) (tg.UserClazz, error) {
	_ = ctx
	r.projectCalls++
	for _, user := range r.usersByPhone {
		if user == nil || user.User == nil || user.User.Id != userId {
			continue
		}
		data := user.User
		return tg.MakeTLUser(&tg.TLUser{
			Self:       true,
			Id:         data.Id,
			AccessHash: &data.AccessHash,
			FirstName:  stringPtr(data.FirstName),
			LastName:   stringPtr(data.LastName),
			Username:   optionalStringPtr(data.Username),
			Phone:      optionalStringPtr(data.Phone),
			Status:     tg.UserStatusEmptyClazz,
		}), nil
	}
	return nil, repository.ErrUserNotFound
}

func (r *fakeAuthRepository) BindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error {
	_ = ctx
	r.bindCalls++
	r.boundAuthKeyID = authKeyId
	r.boundUserID = userId
	return nil
}

func (r *fakeAuthRepository) UnbindAuthKeyUser(ctx context.Context, authKeyId int64, userId int64) error {
	_ = ctx
	r.unbindCalls++
	r.boundAuthKeyID = authKeyId
	r.boundUserID = userId
	return nil
}

func (r *fakeAuthRepository) SetAuthorizationTTL(ctx context.Context, userId int64, ttl int32) error {
	_ = ctx
	_ = userId
	_ = ttl
	return nil
}

func makeFakeImmutableUser(id int64, phone string, countryCode string, firstName string, lastName string) *tg.ImmutableUser {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{
			Id:          id,
			AccessHash:  id * 10,
			SceretKeyId: id * 100,
			FirstName:   firstName,
			LastName:    lastName,
			Username:    "testuser",
			Phone:       phone,
			CountryCode: countryCode,
		}).ToUserData(),
		KeysPrivacyRules: []tg.PrivacyKeyRulesClazz{},
	}).ToImmutableUser()
}

func stringPtr(v string) *string {
	return &v
}

func optionalStringPtr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
