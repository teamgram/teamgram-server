/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authsessionclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession/authsessionservice"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type AuthsessionClient interface {
	AuthsessionGetAuthorizations(ctx context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error)
	AuthsessionResetAuthorization(ctx context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error)
	AuthsessionGetLayer(ctx context.Context, in *authsession.TLAuthsessionGetLayer) (*tg.Int32, error)
	AuthsessionGetLangPack(ctx context.Context, in *authsession.TLAuthsessionGetLangPack) (*tg.String, error)
	AuthsessionGetClient(ctx context.Context, in *authsession.TLAuthsessionGetClient) (*tg.String, error)
	AuthsessionGetLangCode(ctx context.Context, in *authsession.TLAuthsessionGetLangCode) (*tg.String, error)
	AuthsessionGetUserId(ctx context.Context, in *authsession.TLAuthsessionGetUserId) (*tg.Int64, error)
	AuthsessionGetPushSessionId(ctx context.Context, in *authsession.TLAuthsessionGetPushSessionId) (*tg.Int64, error)
	AuthsessionGetFutureSalts(ctx context.Context, in *authsession.TLAuthsessionGetFutureSalts) (*tg.FutureSalts, error)
	AuthsessionQueryAuthKey(ctx context.Context, in *authsession.TLAuthsessionQueryAuthKey) (*tg.AuthKeyInfo, error)
	AuthsessionSetAuthKey(ctx context.Context, in *authsession.TLAuthsessionSetAuthKey) (*tg.Bool, error)
	AuthsessionBindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionBindAuthKeyUser) (*tg.Int64, error)
	AuthsessionUnbindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error)
	AuthsessionGetPermAuthKeyId(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyId) (*tg.Int64, error)
	AuthsessionBindTempAuthKey(ctx context.Context, in *authsession.TLAuthsessionBindTempAuthKey) (*tg.Bool, error)
	AuthsessionSetClientSessionInfo(ctx context.Context, in *authsession.TLAuthsessionSetClientSessionInfo) (*tg.Bool, error)
	AuthsessionGetAuthorization(ctx context.Context, in *authsession.TLAuthsessionGetAuthorization) (*tg.Authorization, error)
	AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error)
	AuthsessionSetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) (*tg.Bool, error)
	AuthsessionSetInitConnection(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) (*tg.Bool, error)
	AuthsessionSetAndroidPushSessionId(ctx context.Context, in *authsession.TLAuthsessionSetAndroidPushSessionId) (*tg.Bool, error)
}

type defaultAuthsessionClient struct {
	cli client.Client
}

func NewAuthsessionClient(cli client.Client) AuthsessionClient {
	return &defaultAuthsessionClient{
		cli: cli,
	}
}

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (m *defaultAuthsessionClient) AuthsessionGetAuthorizations(ctx context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*tg.AccountAuthorizations, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetAuthorizations(ctx, in)
}

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (m *defaultAuthsessionClient) AuthsessionResetAuthorization(ctx context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.VectorLong, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionResetAuthorization(ctx, in)
}

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (m *defaultAuthsessionClient) AuthsessionGetLayer(ctx context.Context, in *authsession.TLAuthsessionGetLayer) (*tg.Int32, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetLayer(ctx, in)
}

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetLangPack(ctx context.Context, in *authsession.TLAuthsessionGetLangPack) (*tg.String, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetLangPack(ctx, in)
}

// AuthsessionGetClient
// authsession.getClient auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetClient(ctx context.Context, in *authsession.TLAuthsessionGetClient) (*tg.String, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetClient(ctx, in)
}

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetLangCode(ctx context.Context, in *authsession.TLAuthsessionGetLangCode) (*tg.String, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetLangCode(ctx, in)
}

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (m *defaultAuthsessionClient) AuthsessionGetUserId(ctx context.Context, in *authsession.TLAuthsessionGetUserId) (*tg.Int64, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetUserId(ctx, in)
}

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (m *defaultAuthsessionClient) AuthsessionGetPushSessionId(ctx context.Context, in *authsession.TLAuthsessionGetPushSessionId) (*tg.Int64, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetPushSessionId(ctx, in)
}

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (m *defaultAuthsessionClient) AuthsessionGetFutureSalts(ctx context.Context, in *authsession.TLAuthsessionGetFutureSalts) (*tg.FutureSalts, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetFutureSalts(ctx, in)
}

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (m *defaultAuthsessionClient) AuthsessionQueryAuthKey(ctx context.Context, in *authsession.TLAuthsessionQueryAuthKey) (*tg.AuthKeyInfo, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionQueryAuthKey(ctx, in)
}

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetAuthKey(ctx context.Context, in *authsession.TLAuthsessionSetAuthKey) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionSetAuthKey(ctx, in)
}

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (m *defaultAuthsessionClient) AuthsessionBindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionBindAuthKeyUser) (*tg.Int64, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionBindAuthKeyUser(ctx, in)
}

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (m *defaultAuthsessionClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionUnbindAuthKeyUser(ctx, in)
}

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (m *defaultAuthsessionClient) AuthsessionGetPermAuthKeyId(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyId) (*tg.Int64, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetPermAuthKeyId(ctx, in)
}

// AuthsessionBindTempAuthKey
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (m *defaultAuthsessionClient) AuthsessionBindTempAuthKey(ctx context.Context, in *authsession.TLAuthsessionBindTempAuthKey) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionBindTempAuthKey(ctx, in)
}

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetClientSessionInfo(ctx context.Context, in *authsession.TLAuthsessionSetClientSessionInfo) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionSetClientSessionInfo(ctx, in)
}

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (m *defaultAuthsessionClient) AuthsessionGetAuthorization(ctx context.Context, in *authsession.TLAuthsessionGetAuthorization) (*tg.Authorization, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetAuthorization(ctx, in)
}

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (m *defaultAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionGetAuthStateData(ctx, in)
}

// AuthsessionSetLayer
// authsession.setLayer auth_key_id:long ip:string layer:int = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionSetLayer(ctx, in)
}

// AuthsessionSetInitConnection
// authsession.setInitConnection auth_key_id:long ip:string api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetInitConnection(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionSetInitConnection(ctx, in)
}

// AuthsessionSetAndroidPushSessionId
// authsession.setAndroidPushSessionId user_id:long auth_key_id:long session_id:long = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetAndroidPushSessionId(ctx context.Context, in *authsession.TLAuthsessionSetAndroidPushSessionId) (*tg.Bool, error) {
	cli := authsessionservice.NewRPCAuthsessionClient(m.cli)
	return cli.AuthsessionSetAndroidPushSessionId(ctx, in)
}
