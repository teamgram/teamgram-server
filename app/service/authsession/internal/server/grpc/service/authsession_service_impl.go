/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/app/service/authsession/internal/core"
)

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (s *Service) AuthsessionGetAuthorizations(ctx context.Context, request *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getAuthorizations - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getAuthorizations - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (s *Service) AuthsessionResetAuthorization(ctx context.Context, request *authsession.TLAuthsessionResetAuthorization) (*authsession.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.resetAuthorization - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionResetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.resetAuthorization - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (s *Service) AuthsessionGetLayer(ctx context.Context, request *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getLayer - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetLayer(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getLayer - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (s *Service) AuthsessionGetLangPack(ctx context.Context, request *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getLangPack - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetLangPack(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getLangPack - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetClient
// authsession.getClient auth_key_id:long = String;
func (s *Service) AuthsessionGetClient(ctx context.Context, request *authsession.TLAuthsessionGetClient) (*mtproto.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getClient - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetClient(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getClient - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (s *Service) AuthsessionGetLangCode(ctx context.Context, request *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getLangCode - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetLangCode(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getLangCode - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (s *Service) AuthsessionGetUserId(ctx context.Context, request *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getUserId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetUserId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getUserId - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (s *Service) AuthsessionGetPushSessionId(ctx context.Context, request *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getPushSessionId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetPushSessionId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getPushSessionId - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (s *Service) AuthsessionGetFutureSalts(ctx context.Context, request *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getFutureSalts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetFutureSalts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getFutureSalts - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (s *Service) AuthsessionQueryAuthKey(ctx context.Context, request *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.queryAuthKey - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionQueryAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.queryAuthKey - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt = Bool;
func (s *Service) AuthsessionSetAuthKey(ctx context.Context, request *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.setAuthKey - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionSetAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.setAuthKey - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (s *Service) AuthsessionBindAuthKeyUser(ctx context.Context, request *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.bindAuthKeyUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionBindAuthKeyUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.bindAuthKeyUser - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (s *Service) AuthsessionUnbindAuthKeyUser(ctx context.Context, request *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.unbindAuthKeyUser - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionUnbindAuthKeyUser(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.unbindAuthKeyUser - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (s *Service) AuthsessionGetPermAuthKeyId(ctx context.Context, request *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getPermAuthKeyId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetPermAuthKeyId(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getPermAuthKeyId - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionBindTempAuthKey
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (s *Service) AuthsessionBindTempAuthKey(ctx context.Context, request *authsession.TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.bindTempAuthKey - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionBindTempAuthKey(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.bindTempAuthKey - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (s *Service) AuthsessionSetClientSessionInfo(ctx context.Context, request *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.setClientSessionInfo - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionSetClientSessionInfo(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.setClientSessionInfo - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (s *Service) AuthsessionGetAuthorization(ctx context.Context, request *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getAuthorization - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getAuthorization - reply: %s", r.DebugString())
	return r, err
}

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (s *Service) AuthsessionGetAuthStateData(ctx context.Context, request *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("authsession.getAuthStateData - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AuthsessionGetAuthStateData(request)
	if err != nil {
		return nil, err
	}

	c.Infof("authsession.getAuthStateData - reply: %s", r.DebugString())
	return r, err
}
