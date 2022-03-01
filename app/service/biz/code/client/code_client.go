/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package code_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type CodeClient interface {
	CodeCreatePhoneCode(ctx context.Context, in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error)
	CodeGetPhoneCode(ctx context.Context, in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error)
	CodeDeletePhoneCode(ctx context.Context, in *code.TLCodeDeletePhoneCode) (*mtproto.Bool, error)
	CodeUpdatePhoneCodeData(ctx context.Context, in *code.TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error)
}

type defaultCodeClient struct {
	cli zrpc.Client
}

func NewCodeClient(cli zrpc.Client) CodeClient {
	return &defaultCodeClient{
		cli: cli,
	}
}

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (m *defaultCodeClient) CodeCreatePhoneCode(ctx context.Context, in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	client := code.NewRPCCodeClient(m.cli.Conn())
	return client.CodeCreatePhoneCode(ctx, in)
}

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (m *defaultCodeClient) CodeGetPhoneCode(ctx context.Context, in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	client := code.NewRPCCodeClient(m.cli.Conn())
	return client.CodeGetPhoneCode(ctx, in)
}

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (m *defaultCodeClient) CodeDeletePhoneCode(ctx context.Context, in *code.TLCodeDeletePhoneCode) (*mtproto.Bool, error) {
	client := code.NewRPCCodeClient(m.cli.Conn())
	return client.CodeDeletePhoneCode(ctx, in)
}

// CodeUpdatePhoneCodeData
// code.updatePhoneCodeData auth_key_id:long phone:string phone_code_hash:string code_data:PhoneCodeTransaction = Bool;
func (m *defaultCodeClient) CodeUpdatePhoneCodeData(ctx context.Context, in *code.TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error) {
	client := code.NewRPCCodeClient(m.cli.Conn())
	return client.CodeUpdatePhoneCodeData(ctx, in)
}
