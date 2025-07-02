/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/core"
)

var _ *tg.Bool

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (s *Service) AuthsessionGetAuthorizations(ctx context.Context, request *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getAuthorizations - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getAuthorizations - reply: %s", r)
	return r, err
}

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (s *Service) AuthsessionResetAuthorization(ctx context.Context, request *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.resetAuthorization - metadata: {}, request: %s", request)

	r, err := c.AuthsessionResetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.resetAuthorization - reply: %s", r)
	return r, err
}

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (s *Service) AuthsessionGetLayer(ctx context.Context, request *authsession.TLAuthsessionGetLayer) (*tg.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getLayer - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetLayer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getLayer - reply: %s", r)
	return r, err
}

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (s *Service) AuthsessionGetLangPack(ctx context.Context, request *authsession.TLAuthsessionGetLangPack) (*tg.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getLangPack - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetLangPack(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getLangPack - reply: %s", r)
	return r, err
}

// AuthsessionGetClient
// authsession.getClient auth_key_id:long = String;
func (s *Service) AuthsessionGetClient(ctx context.Context, request *authsession.TLAuthsessionGetClient) (*tg.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getClient - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetClient(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getClient - reply: %s", r)
	return r, err
}

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (s *Service) AuthsessionGetLangCode(ctx context.Context, request *authsession.TLAuthsessionGetLangCode) (*tg.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getLangCode - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetLangCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getLangCode - reply: %s", r)
	return r, err
}

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (s *Service) AuthsessionGetUserId(ctx context.Context, request *authsession.TLAuthsessionGetUserId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getUserId - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetUserId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getUserId - reply: %s", r)
	return r, err
}

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (s *Service) AuthsessionGetPushSessionId(ctx context.Context, request *authsession.TLAuthsessionGetPushSessionId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getPushSessionId - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetPushSessionId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getPushSessionId - reply: %s", r)
	return r, err
}

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (s *Service) AuthsessionGetFutureSalts(ctx context.Context, request *authsession.TLAuthsessionGetFutureSalts) (*tg.FutureSalts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getFutureSalts - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetFutureSalts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getFutureSalts - reply: %s", r)
	return r, err
}

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (s *Service) AuthsessionQueryAuthKey(ctx context.Context, request *authsession.TLAuthsessionQueryAuthKey) (*tg.AuthKeyInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.queryAuthKey - metadata: {}, request: %s", request)

	r, err := c.AuthsessionQueryAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.queryAuthKey - reply: %s", r)
	return r, err
}

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (s *Service) AuthsessionSetAuthKey(ctx context.Context, request *authsession.TLAuthsessionSetAuthKey) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.setAuthKey - metadata: {}, request: %s", request)

	r, err := c.AuthsessionSetAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.setAuthKey - reply: %s", r)
	return r, err
}

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (s *Service) AuthsessionBindAuthKeyUser(ctx context.Context, request *authsession.TLAuthsessionBindAuthKeyUser) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.bindAuthKeyUser - metadata: {}, request: %s", request)

	r, err := c.AuthsessionBindAuthKeyUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.bindAuthKeyUser - reply: %s", r)
	return r, err
}

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (s *Service) AuthsessionUnbindAuthKeyUser(ctx context.Context, request *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.unbindAuthKeyUser - metadata: {}, request: %s", request)

	r, err := c.AuthsessionUnbindAuthKeyUser(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.unbindAuthKeyUser - reply: %s", r)
	return r, err
}

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (s *Service) AuthsessionGetPermAuthKeyId(ctx context.Context, request *authsession.TLAuthsessionGetPermAuthKeyId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getPermAuthKeyId - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetPermAuthKeyId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getPermAuthKeyId - reply: %s", r)
	return r, err
}

// AuthsessionBindTempAuthKey
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (s *Service) AuthsessionBindTempAuthKey(ctx context.Context, request *authsession.TLAuthsessionBindTempAuthKey) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.bindTempAuthKey - metadata: {}, request: %s", request)

	r, err := c.AuthsessionBindTempAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.bindTempAuthKey - reply: %s", r)
	return r, err
}

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (s *Service) AuthsessionSetClientSessionInfo(ctx context.Context, request *authsession.TLAuthsessionSetClientSessionInfo) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.setClientSessionInfo - metadata: {}, request: %s", request)

	r, err := c.AuthsessionSetClientSessionInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.setClientSessionInfo - reply: %s", r)
	return r, err
}

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (s *Service) AuthsessionGetAuthorization(ctx context.Context, request *authsession.TLAuthsessionGetAuthorization) (*tg.Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getAuthorization - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getAuthorization - reply: %s", r)
	return r, err
}

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (s *Service) AuthsessionGetAuthStateData(ctx context.Context, request *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.getAuthStateData - metadata: {}, request: %s", request)

	r, err := c.AuthsessionGetAuthStateData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.getAuthStateData - reply: %s", r)
	return r, err
}

// AuthsessionSetLayer
// authsession.setLayer auth_key_id:long ip:string layer:int = Bool;
func (s *Service) AuthsessionSetLayer(ctx context.Context, request *authsession.TLAuthsessionSetLayer) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.setLayer - metadata: {}, request: %s", request)

	r, err := c.AuthsessionSetLayer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.setLayer - reply: %s", r)
	return r, err
}

// AuthsessionSetInitConnection
// authsession.setInitConnection auth_key_id:long ip:string api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = Bool;
func (s *Service) AuthsessionSetInitConnection(ctx context.Context, request *authsession.TLAuthsessionSetInitConnection) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.setInitConnection - metadata: {}, request: %s", request)

	r, err := c.AuthsessionSetInitConnection(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.setInitConnection - reply: %s", r)
	return r, err
}

// AuthsessionSetAndroidPushSessionId
// authsession.setAndroidPushSessionId user_id:long auth_key_id:long session_id:long = Bool;
func (s *Service) AuthsessionSetAndroidPushSessionId(ctx context.Context, request *authsession.TLAuthsessionSetAndroidPushSessionId) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("authsession.setAndroidPushSessionId - metadata: {}, request: %s", request)

	r, err := c.AuthsessionSetAndroidPushSessionId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("authsession.setAndroidPushSessionId - reply: %s", r)
	return r, err
}
