package core

import (
	"context"
	"testing"

	kitexmetadata "github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestAuthCancelCodePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	if _, err := c.AuthCancelCode(&tg.TLAuthCancelCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
	}); err != tg.ErrPhoneCodeHashEmpty {
		t.Fatalf("expected phone code hash empty, got %v", err)
	}

	result, err := c.AuthCancelCode(&tg.TLAuthCancelCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(result) {
		t.Fatalf("expected boolTrue, got %#v", result)
	}
}

func TestAuthLogOutReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.AuthLogOut(&tg.TLAuthLogOut{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected loggedOut placeholder, got nil")
	}
	if len(result.FutureAuthToken) != 0 {
		t.Fatalf("expected empty future auth token, got %#v", result.FutureAuthToken)
	}
}

func TestAuthExportImportAuthorizationPlaceholders(t *testing.T) {
	md := &kitexmetadata.RpcMetadata{AuthId: 77, UserId: 42}
	ctx, err := kitexmetadata.RpcMetadataToOutgoing(context.Background(), md)
	if err != nil {
		t.Fatalf("attach rpc metadata: %v", err)
	}
	c := New(ctx, nil)

	exported, err := c.AuthExportAuthorization(&tg.TLAuthExportAuthorization{DcId: 2})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if exported == nil || exported.Id != 77 {
		t.Fatalf("expected exported auth id=77, got %#v", exported)
	}
	if len(exported.Bytes) != 1 || exported.Bytes[0] != 2 {
		t.Fatalf("expected exported bytes=[2], got %#v", exported.Bytes)
	}

	imported, err := c.AuthImportAuthorization(&tg.TLAuthImportAuthorization{
		Id:    exported.Id,
		Bytes: exported.Bytes,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	auth, ok := imported.ToAuthAuthorization()
	if !ok {
		t.Fatalf("expected auth.authorization, got %T", imported.Clazz)
	}
	user, ok := auth.User.(*tg.TLUserEmpty)
	if !ok {
		t.Fatalf("expected userEmpty placeholder, got %T", auth.User)
	}
	if user.Id != 42 {
		t.Fatalf("expected imported placeholder user id=42, got %d", user.Id)
	}
	if len(auth.FutureAuthToken) != 1 || auth.FutureAuthToken[0] != 2 {
		t.Fatalf("expected imported future auth token=[2], got %#v", auth.FutureAuthToken)
	}
}
