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
	"github.com/teamgram/teamgram-server/v2/pkg/code"
	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func newAuthorizationCoreForAuthTest(t *testing.T, md *kitexmetadata.RpcMetadata) (*AuthorizationCore, context.Context, *dao.Dao) {
	return newAuthorizationCoreForAuthTestWithVerifier(t, md, nil)
}

func newAuthorizationCoreForAuthTestWithVerifier(t *testing.T, md *kitexmetadata.RpcMetadata, verifier code.VerifyCodeInterface) (*AuthorizationCore, context.Context, *dao.Dao) {
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
		AuthLogic: logic.NewAuthSignLogic(d, verifier),
	}
	return New(ctx, svcCtx), ctx, d
}

func seedPhoneCodeTransaction(t *testing.T, ctx context.Context, d *dao.Dao, authKeyID, sessionID int64, registered bool, code, hash string, state int) string {
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
		State:                 state,
	})
	if err != nil {
		t.Fatalf("seed phone code: %v", err)
	}
	return phoneNumber
}

type fakeVerifyCode struct {
	sendExtraData string
	sendErr       error
	verifyErr     error
}

func (f *fakeVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, codeValue, codeHash string) (string, error) {
	_ = ctx
	_ = phoneNumber
	_ = codeValue
	_ = codeHash
	if f.sendErr != nil {
		return "", f.sendErr
	}
	return f.sendExtraData, nil
}

func (f *fakeVerifyCode) VerifySmsCode(ctx context.Context, codeHash, codeValue, extraData string) error {
	_ = ctx
	_ = codeHash
	_ = codeValue
	_ = extraData
	return f.verifyErr
}
