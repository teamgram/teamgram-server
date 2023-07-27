/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package authsession_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type AuthsessionClient interface {
	AuthsessionGetAuthorizations(ctx context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error)
	AuthsessionResetAuthorization(ctx context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.Vector_Long, error)
	AuthsessionGetLayer(ctx context.Context, in *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error)
	AuthsessionGetLangPack(ctx context.Context, in *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error)
	AuthsessionGetClient(ctx context.Context, in *authsession.TLAuthsessionGetClient) (*mtproto.String, error)
	AuthsessionGetLangCode(ctx context.Context, in *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error)
	AuthsessionGetUserId(ctx context.Context, in *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error)
	AuthsessionGetPushSessionId(ctx context.Context, in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error)
	AuthsessionGetFutureSalts(ctx context.Context, in *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error)
	AuthsessionQueryAuthKey(ctx context.Context, in *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error)
	AuthsessionSetAuthKey(ctx context.Context, in *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error)
	AuthsessionBindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error)
	AuthsessionUnbindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error)
	AuthsessionGetPermAuthKeyId(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error)
	AuthsessionBindTempAuthKey(ctx context.Context, in *authsession.TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error)
	AuthsessionSetClientSessionInfo(ctx context.Context, in *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error)
	AuthsessionGetAuthorization(ctx context.Context, in *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error)
	AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error)
	AuthsessionSetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) (*mtproto.Bool, error)
	AuthsessionSetInitConnection(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) (*mtproto.Bool, error)
}

type defaultAuthsessionClient struct {
	cli zrpc.Client
}

func NewAuthsessionClient(cli zrpc.Client) AuthsessionClient {
	return &defaultAuthsessionClient{
		cli: cli,
	}
}

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (m *defaultAuthsessionClient) AuthsessionGetAuthorizations(ctx context.Context, in *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetAuthorizations(ctx, in)
}

// AuthsessionResetAuthorization
// authsession.resetAuthorization user_id:long auth_key_id:long hash:long = Vector<long>;
func (m *defaultAuthsessionClient) AuthsessionResetAuthorization(ctx context.Context, in *authsession.TLAuthsessionResetAuthorization) (*authsession.Vector_Long, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionResetAuthorization(ctx, in)
}

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (m *defaultAuthsessionClient) AuthsessionGetLayer(ctx context.Context, in *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetLayer(ctx, in)
}

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetLangPack(ctx context.Context, in *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetLangPack(ctx, in)
}

// AuthsessionGetClient
// authsession.getClient auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetClient(ctx context.Context, in *authsession.TLAuthsessionGetClient) (*mtproto.String, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetClient(ctx, in)
}

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (m *defaultAuthsessionClient) AuthsessionGetLangCode(ctx context.Context, in *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetLangCode(ctx, in)
}

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (m *defaultAuthsessionClient) AuthsessionGetUserId(ctx context.Context, in *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetUserId(ctx, in)
}

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (m *defaultAuthsessionClient) AuthsessionGetPushSessionId(ctx context.Context, in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetPushSessionId(ctx, in)
}

// AuthsessionGetFutureSalts
// authsession.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (m *defaultAuthsessionClient) AuthsessionGetFutureSalts(ctx context.Context, in *authsession.TLAuthsessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetFutureSalts(ctx, in)
}

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (m *defaultAuthsessionClient) AuthsessionQueryAuthKey(ctx context.Context, in *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionQueryAuthKey(ctx, in)
}

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt expires_in:int = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetAuthKey(ctx context.Context, in *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionSetAuthKey(ctx, in)
}

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (m *defaultAuthsessionClient) AuthsessionBindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionBindAuthKeyUser(ctx, in)
}

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (m *defaultAuthsessionClient) AuthsessionUnbindAuthKeyUser(ctx context.Context, in *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionUnbindAuthKeyUser(ctx, in)
}

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (m *defaultAuthsessionClient) AuthsessionGetPermAuthKeyId(ctx context.Context, in *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetPermAuthKeyId(ctx, in)
}

// AuthsessionBindTempAuthKey
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (m *defaultAuthsessionClient) AuthsessionBindTempAuthKey(ctx context.Context, in *authsession.TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionBindTempAuthKey(ctx, in)
}

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetClientSessionInfo(ctx context.Context, in *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionSetClientSessionInfo(ctx, in)
}

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (m *defaultAuthsessionClient) AuthsessionGetAuthorization(ctx context.Context, in *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetAuthorization(ctx, in)
}

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (m *defaultAuthsessionClient) AuthsessionGetAuthStateData(ctx context.Context, in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionGetAuthStateData(ctx, in)
}

// AuthsessionSetLayer
// authsession.setLayer auth_key_id:long ip:string layer:int = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetLayer(ctx context.Context, in *authsession.TLAuthsessionSetLayer) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionSetLayer(ctx, in)
}

// AuthsessionSetInitConnection
// authsession.setInitConnection auth_key_id:long ip:string api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = Bool;
func (m *defaultAuthsessionClient) AuthsessionSetInitConnection(ctx context.Context, in *authsession.TLAuthsessionSetInitConnection) (*mtproto.Bool, error) {
	client := authsession.NewRPCAuthsessionClient(m.cli.Conn())
	return client.AuthsessionSetInitConnection(ctx, in)
}
