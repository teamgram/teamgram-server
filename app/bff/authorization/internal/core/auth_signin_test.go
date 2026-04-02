package core

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/dao"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/logic"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/bff/authorization/model"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestAuthSignInReturnsPhoneCodeEmptyWhenCodeMissing(t *testing.T) {
	c := New(context.Background(), nil)

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     nil,
	})
	if err != tg.ErrPhoneCodeEmpty {
		t.Fatalf("expected ErrPhoneCodeEmpty, got %v", err)
	}
}

func TestAuthSignInReturnsPhoneCodeEmptyWhenHashMissing(t *testing.T) {
	c := New(context.Background(), nil)
	code := "12345"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
		PhoneCode:     &code,
	})
	if err != tg.ErrPhoneCodeEmpty {
		t.Fatalf("expected ErrPhoneCodeEmpty, got %v", err)
	}
}

func TestAuthSignInReturnsPhoneNumberInvalidForBadPhone(t *testing.T) {
	c := New(context.Background(), nil)
	code := "12345"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "bad-phone",
		PhoneCodeHash: "hash",
		PhoneCode:     &code,
	})
	if err != tg.Err406PhoneNumberInvalid {
		t.Fatalf("expected Err406PhoneNumberInvalid, got %v", err)
	}
}

func TestAuthSignInReturnsSignUpRequiredForUnregisteredPhone(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForSignInTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})
	code := "12345"
	phoneNumber := seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, code, "hash")

	result, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     &code,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if _, ok := result.ToAuthAuthorizationSignUpRequired(); !ok {
		t.Fatalf("expected auth.authorizationSignUpRequired, got %T", result.Clazz)
	}

	codeData, err := d.GetPhoneCode(ctx, 101, phoneNumber, "hash")
	if err != nil {
		t.Fatalf("get phone code after sign in: %v", err)
	}
	if codeData.State != model.CodeStateSignIn {
		t.Fatalf("expected state %d, got %d", model.CodeStateSignIn, codeData.State)
	}
}

func TestAuthSignInReturnsPhoneCodeInvalidWhenCodeDoesNotMatch(t *testing.T) {
	c, ctx, d := newAuthorizationCoreForSignInTest(t, &kitexmetadata.RpcMetadata{
		PermAuthKeyId: 101,
		SessionId:     202,
		Layer:         181,
	})
	seedPhoneCodeTransaction(t, ctx, d, 101, 202, false, "12345", "hash")
	wrongCode := "54321"

	_, err := c.AuthSignIn(&tg.TLAuthSignIn{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		PhoneCode:     &wrongCode,
	})
	if err != tg.ErrPhoneCodeInvalid {
		t.Fatalf("expected ErrPhoneCodeInvalid, got %v", err)
	}
}

func seedPhoneCodeTransaction(t *testing.T, ctx context.Context, d *dao.Dao, authKeyID, sessionID int64, registered bool, code, hash string) string {
	t.Helper()

	_, phoneNumber, err := checkPhoneNumberInvalid("+8613812345678")
	if err != nil {
		t.Fatalf("normalize phone: %v", err)
	}

	err = d.UpdatePhoneCodeData(ctx, authKeyID, phoneNumber, hash, &model.PhoneCodeTransaction{
		AuthKeyId:             authKeyID,
		SessionId:             sessionID,
		PhoneNumber:           phoneNumber,
		PhoneNumberRegistered: registered,
		PhoneCode:             code,
		PhoneCodeHash:         hash,
		PhoneCodeExpired:      int32(time.Now().Add(time.Minute).Unix()),
		State:                 model.CodeStateSent,
	})
	if err != nil {
		t.Fatalf("seed phone code: %v", err)
	}
	return phoneNumber
}

func newAuthorizationCoreForSignInTest(t *testing.T, md *kitexmetadata.RpcMetadata) (*AuthorizationCore, context.Context, *dao.Dao) {
	t.Helper()

	redisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf("start miniredis: %v", err)
	}
	t.Cleanup(redisServer.Close)

	cfg := config.Config{
		KV: kv.KvConf{
			{
				RedisConf: redis.RedisConf{
					Host: redisServer.Addr(),
					Type: "node",
				},
				Weight: 100,
			},
		},
	}
	d := dao.NewWithKVStore(kv.NewStore(cfg.KV))
	ctx, err := kitexmetadata.RpcMetadataToOutgoing(context.Background(), md)
	if err != nil {
		t.Fatalf("attach rpc metadata: %v", err)
	}

	svcCtx := &svc.ServiceContext{
		Config:    cfg,
		Dao:       d,
		AuthLogic: logic.NewAuthSignLogic(d, nil),
	}
	return New(ctx, svcCtx), ctx, d
}
