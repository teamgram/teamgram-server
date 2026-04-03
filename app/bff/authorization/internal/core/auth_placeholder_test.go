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

func TestAuthorizationBoolPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	if _, err := c.AccountSetAuthorizationTTL(&tg.TLAccountSetAuthorizationTTL{AuthorizationTtlDays: -1}); err != tg.ErrInputMethodInvalid {
		t.Fatalf("expected input method invalid, got %v", err)
	}

	setTTL, err := c.AccountSetAuthorizationTTL(&tg.TLAccountSetAuthorizationTTL{AuthorizationTtlDays: 30})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(setTTL) {
		t.Fatalf("expected set authorization ttl boolTrue, got %#v", setTTL)
	}

	resetAuths, err := c.AuthResetAuthorizations(&tg.TLAuthResetAuthorizations{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(resetAuths) {
		t.Fatalf("expected resetAuthorizations boolTrue, got %#v", resetAuths)
	}

	dropTempKeys, err := c.AuthDropTempAuthKeys(&tg.TLAuthDropTempAuthKeys{ExceptAuthKeys: []int64{1, 2}})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(dropTempKeys) {
		t.Fatalf("expected dropTempAuthKeys boolTrue, got %#v", dropTempKeys)
	}
}

func TestAuthFirebaseAndPaidPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	if _, err := c.AuthRequestFirebaseSms(&tg.TLAuthRequestFirebaseSms{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
	}); err != tg.ErrPhoneCodeHashEmpty {
		t.Fatalf("expected phone code hash empty, got %v", err)
	}

	requestFirebase, err := c.AuthRequestFirebaseSms(&tg.TLAuthRequestFirebaseSms{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(requestFirebase) {
		t.Fatalf("expected requestFirebaseSms boolTrue, got %#v", requestFirebase)
	}

	if _, err := c.AuthReportMissingCode(&tg.TLAuthReportMissingCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		Mnc:           "",
	}); err != tg.ErrInputMethodInvalid {
		t.Fatalf("expected input method invalid, got %v", err)
	}

	reportMissing, err := c.AuthReportMissingCode(&tg.TLAuthReportMissingCode{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		Mnc:           "460",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(reportMissing) {
		t.Fatalf("expected reportMissingCode boolTrue, got %#v", reportMissing)
	}

	paidAuth, err := c.AuthCheckPaidAuth(&tg.TLAuthCheckPaidAuth{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
		FormId:        1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if paidAuth == nil {
		t.Fatal("expected sentCode placeholder, got nil")
	}
	sentCode, ok := paidAuth.ToAuthSentCode()
	if !ok {
		t.Fatalf("expected auth.sentCode, got %T", paidAuth.Clazz)
	}
	if sentCode.PhoneCodeHash != "hash" {
		t.Fatalf("expected phone_code_hash=hash, got %q", sentCode.PhoneCodeHash)
	}
	if sentCode.Timeout == nil || *sentCode.Timeout != 60 {
		t.Fatalf("expected timeout=60, got %#v", sentCode.Timeout)
	}
}

func TestAuthPasswordRecoveryPlaceholders(t *testing.T) {
	md := &kitexmetadata.RpcMetadata{UserId: 42}
	ctx, err := kitexmetadata.RpcMetadataToOutgoing(context.Background(), md)
	if err != nil {
		t.Fatalf("attach rpc metadata: %v", err)
	}
	c := New(ctx, nil)

	if _, err := c.AuthCheckPassword(&tg.TLAuthCheckPassword{}); err != tg.ErrPasswordEmpty {
		t.Fatalf("expected password empty, got %v", err)
	}

	checkPassword, err := c.AuthCheckPassword(&tg.TLAuthCheckPassword{
		Password: tg.MakeTLInputCheckPasswordEmpty(&tg.TLInputCheckPasswordEmpty{}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	authFromPassword, ok := checkPassword.ToAuthAuthorization()
	if !ok {
		t.Fatalf("expected auth.authorization, got %T", checkPassword.Clazz)
	}
	userFromPassword, ok := authFromPassword.User.(*tg.TLUserEmpty)
	if !ok || userFromPassword.Id != 42 {
		t.Fatalf("expected userEmpty id=42, got %#v", authFromPassword.User)
	}

	recovery, err := c.AuthRequestPasswordRecovery(&tg.TLAuthRequestPasswordRecovery{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if recovery == nil || recovery.EmailPattern != "t***@example.com" {
		t.Fatalf("expected placeholder email pattern, got %#v", recovery)
	}

	if _, err := c.AuthRecoverPassword(&tg.TLAuthRecoverPassword{}); err != tg.ErrCodeEmpty {
		t.Fatalf("expected code empty, got %v", err)
	}

	recoverPassword, err := c.AuthRecoverPassword(&tg.TLAuthRecoverPassword{
		Code: "12345",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	authRecovered, ok := recoverPassword.ToAuthAuthorization()
	if !ok {
		t.Fatalf("expected auth.authorization, got %T", recoverPassword.Clazz)
	}
	userRecovered, ok := authRecovered.User.(*tg.TLUserEmpty)
	if !ok || userRecovered.Id != 42 {
		t.Fatalf("expected recovered userEmpty id=42, got %#v", authRecovered.User)
	}
}

func TestAuthEmailVerificationPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	if _, err := c.AccountSendVerifyEmailCode(&tg.TLAccountSendVerifyEmailCode{
		Purpose: tg.MakeTLEmailVerifyPurposeLoginSetup(&tg.TLEmailVerifyPurposeLoginSetup{
			PhoneNumber:   "+8613812345678",
			PhoneCodeHash: "hash",
		}),
		Email: "bad-email",
	}); err != tg.ErrEmailInvalid {
		t.Fatalf("expected email invalid, got %v", err)
	}

	sentEmailCode, err := c.AccountSendVerifyEmailCode(&tg.TLAccountSendVerifyEmailCode{
		Purpose: tg.MakeTLEmailVerifyPurposeLoginSetup(&tg.TLEmailVerifyPurposeLoginSetup{
			PhoneNumber:   "+8613812345678",
			PhoneCodeHash: "hash",
		}),
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if sentEmailCode == nil || sentEmailCode.EmailPattern != "t***@example.com" || sentEmailCode.Length != 6 {
		t.Fatalf("expected sentEmailCode placeholder, got %#v", sentEmailCode)
	}

	if _, err := c.AccountVerifyEmail(&tg.TLAccountVerifyEmail{
		Purpose: tg.MakeTLEmailVerifyPurposeLoginSetup(&tg.TLEmailVerifyPurposeLoginSetup{
			PhoneNumber:   "+8613812345678",
			PhoneCodeHash: "hash",
		}),
		Verification: tg.MakeTLEmailVerificationCode(&tg.TLEmailVerificationCode{}),
	}); err != tg.ErrCodeEmpty {
		t.Fatalf("expected code empty, got %v", err)
	}

	verifiedEmail, err := c.AccountVerifyEmail(&tg.TLAccountVerifyEmail{
		Purpose: tg.MakeTLEmailVerifyPurposeLoginSetup(&tg.TLEmailVerifyPurposeLoginSetup{
			PhoneNumber:   "+8613812345678",
			PhoneCodeHash: "hash",
		}),
		Verification: tg.MakeTLEmailVerificationCode(&tg.TLEmailVerificationCode{Code: "123456"}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	verified, ok := verifiedEmail.ToAccountEmailVerified()
	if !ok {
		t.Fatalf("expected account.emailVerified, got %T", verifiedEmail.Clazz)
	}
	if verified.Email != "t***@example.com" {
		t.Fatalf("expected placeholder verified email, got %#v", verified)
	}

	invalidated, err := c.AccountInvalidateSignInCodes(&tg.TLAccountInvalidateSignInCodes{
		Codes: []string{"one", "two"},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(invalidated) {
		t.Fatalf("expected invalidateSignInCodes boolTrue, got %#v", invalidated)
	}

	if _, err := c.AuthResetLoginEmail(&tg.TLAuthResetLoginEmail{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "",
	}); err != tg.ErrPhoneCodeHashEmpty {
		t.Fatalf("expected phone code hash empty, got %v", err)
	}

	resetLoginEmail, err := c.AuthResetLoginEmail(&tg.TLAuthResetLoginEmail{
		PhoneNumber:   "+8613812345678",
		PhoneCodeHash: "hash",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	sentCode, ok := resetLoginEmail.ToAuthSentCode()
	if !ok {
		t.Fatalf("expected auth.sentCode, got %T", resetLoginEmail.Clazz)
	}
	if sentCode.PhoneCodeHash != "hash" {
		t.Fatalf("expected phone_code_hash=hash, got %q", sentCode.PhoneCodeHash)
	}
	if sentCode.Timeout == nil || *sentCode.Timeout != 60 {
		t.Fatalf("expected timeout=60, got %#v", sentCode.Timeout)
	}
}

func TestAuthRecoveryResetPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	if _, err := c.AuthCheckRecoveryPassword(&tg.TLAuthCheckRecoveryPassword{}); err != tg.ErrCodeEmpty {
		t.Fatalf("expected code empty, got %v", err)
	}

	checked, err := c.AuthCheckRecoveryPassword(&tg.TLAuthCheckRecoveryPassword{Code: "123456"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(checked) {
		t.Fatalf("expected checkRecoveryPassword boolTrue, got %#v", checked)
	}

	resetPassword, err := c.AccountResetPassword(&tg.TLAccountResetPassword{})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, ok := resetPassword.ToAccountResetPasswordOk(); !ok {
		t.Fatalf("expected account.resetPasswordOk, got %T", resetPassword.Clazz)
	}
}

func TestAuthImportAndBindPlaceholders(t *testing.T) {
	md := &kitexmetadata.RpcMetadata{UserId: 99}
	ctx, err := kitexmetadata.RpcMetadataToOutgoing(context.Background(), md)
	if err != nil {
		t.Fatalf("attach rpc metadata: %v", err)
	}
	c := New(ctx, nil)

	if _, err := c.AuthBindTempAuthKey(&tg.TLAuthBindTempAuthKey{}); err != tg.ErrTempAuthKeyEmpty {
		t.Fatalf("expected temp auth key empty, got %v", err)
	}

	bound, err := c.AuthBindTempAuthKey(&tg.TLAuthBindTempAuthKey{
		PermAuthKeyId:    1,
		Nonce:            2,
		ExpiresAt:        3,
		EncryptedMessage: []byte{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(bound) {
		t.Fatalf("expected bindTempAuthKey boolTrue, got %#v", bound)
	}

	if _, err := c.AuthImportBotAuthorization(&tg.TLAuthImportBotAuthorization{
		ApiId:        0,
		ApiHash:      "hash",
		BotAuthToken: "token",
	}); err != tg.ErrApiIdInvalid {
		t.Fatalf("expected api id invalid, got %v", err)
	}

	if _, err := c.AuthImportBotAuthorization(&tg.TLAuthImportBotAuthorization{
		ApiId:        1,
		ApiHash:      "hash",
		BotAuthToken: "",
	}); err != tg.ErrAuthTokenInvalid {
		t.Fatalf("expected auth token invalid, got %v", err)
	}

	botAuth, err := c.AuthImportBotAuthorization(&tg.TLAuthImportBotAuthorization{
		ApiId:        1,
		ApiHash:      "hash",
		BotAuthToken: "token",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	botAuthorization, ok := botAuth.ToAuthAuthorization()
	if !ok {
		t.Fatalf("expected auth.authorization, got %T", botAuth.Clazz)
	}
	botUser, ok := botAuthorization.User.(*tg.TLUserEmpty)
	if !ok || botUser.Id != 99 {
		t.Fatalf("expected userEmpty id=99, got %#v", botAuthorization.User)
	}

	if _, err := c.AuthImportWebTokenAuthorization(&tg.TLAuthImportWebTokenAuthorization{
		ApiId:        0,
		ApiHash:      "hash",
		WebAuthToken: "token",
	}); err != tg.ErrApiIdInvalid {
		t.Fatalf("expected api id invalid, got %v", err)
	}

	if _, err := c.AuthImportWebTokenAuthorization(&tg.TLAuthImportWebTokenAuthorization{
		ApiId:        1,
		ApiHash:      "hash",
		WebAuthToken: "",
	}); err != tg.ErrAuthTokenInvalid {
		t.Fatalf("expected auth token invalid, got %v", err)
	}

	webAuth, err := c.AuthImportWebTokenAuthorization(&tg.TLAuthImportWebTokenAuthorization{
		ApiId:        1,
		ApiHash:      "hash",
		WebAuthToken: "token",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	webAuthorization, ok := webAuth.ToAuthAuthorization()
	if !ok {
		t.Fatalf("expected auth.authorization, got %T", webAuth.Clazz)
	}
	webUser, ok := webAuthorization.User.(*tg.TLUserEmpty)
	if !ok || webUser.Id != 99 {
		t.Fatalf("expected userEmpty id=99, got %#v", webAuthorization.User)
	}
}
