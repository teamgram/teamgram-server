package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthBindTempAuthKeyDelegatesToAuthsession(t *testing.T) {
	repo := newFakeAuthRepository()
	c := newTestAuthorizationCore(t, repo, 3003003, 0)
	req := &tg.TLAuthBindTempAuthKey{
		PermAuthKeyId:    1001001,
		Nonce:            2002002,
		ExpiresAt:        3003003,
		EncryptedMessage: []byte("encrypted-message"),
	}

	got, err := c.AuthBindTempAuthKey(req)
	if err != nil {
		t.Fatalf("AuthBindTempAuthKey() error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("AuthBindTempAuthKey() = %v, want boolTrue", got)
	}
	if repo.bindTempCalls != 1 {
		t.Fatalf("bind temp calls = %d, want 1", repo.bindTempCalls)
	}
	if repo.bindTempPermAuthKeyID != req.PermAuthKeyId ||
		repo.bindTempNonce != req.Nonce ||
		repo.bindTempExpiresAt != req.ExpiresAt ||
		string(repo.bindTempEncryptedMessage) != string(req.EncryptedMessage) {
		t.Fatalf("bound temp request = %d/%d/%d/%x, want %d/%d/%d/%x",
			repo.bindTempPermAuthKeyID,
			repo.bindTempNonce,
			repo.bindTempExpiresAt,
			repo.bindTempEncryptedMessage,
			req.PermAuthKeyId,
			req.Nonce,
			req.ExpiresAt,
			req.EncryptedMessage)
	}
}

func TestAuthBindTempAuthKeyMapsEncryptedMessageInvalid(t *testing.T) {
	repo := newFakeAuthRepository()
	repo.bindTempErr = authsession.ErrEncryptedMessageInvalid
	c := newTestAuthorizationCore(t, repo, 3003003, 0)

	_, err := c.AuthBindTempAuthKey(&tg.TLAuthBindTempAuthKey{
		PermAuthKeyId:    1001001,
		Nonce:            2002002,
		ExpiresAt:        3003003,
		EncryptedMessage: []byte("bad-message"),
	})
	if !errors.Is(err, tg.ErrEncryptedMessageInvalid) {
		t.Fatalf("AuthBindTempAuthKey() error = %v, want %v", err, tg.ErrEncryptedMessageInvalid)
	}
}

func (r *fakeAuthRepository) BindTempAuthKey(ctx context.Context, permAuthKeyID int64, nonce int64, expiresAt int32, encryptedMessage []byte) error {
	_ = ctx
	r.bindTempCalls++
	r.bindTempPermAuthKeyID = permAuthKeyID
	r.bindTempNonce = nonce
	r.bindTempExpiresAt = expiresAt
	r.bindTempEncryptedMessage = encryptedMessage
	return r.bindTempErr
}
