/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package codeclient

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code/codeservice"

	"github.com/cloudwego/kitex/client"
)

type CodeClient interface {
	CodeCreatePhoneCode(ctx context.Context, in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error)
	CodeGetPhoneCode(ctx context.Context, in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error)
	CodeDeletePhoneCode(ctx context.Context, in *code.TLCodeDeletePhoneCode) (*tg.Bool, error)
	CodeUpdatePhoneCodeData(ctx context.Context, in *code.TLCodeUpdatePhoneCodeData) (*tg.Bool, error)
}

type defaultCodeClient struct {
	cli client.Client
}

func NewCodeClient(cli client.Client) CodeClient {
	return &defaultCodeClient{
		cli: cli,
	}
}

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (m *defaultCodeClient) CodeCreatePhoneCode(ctx context.Context, in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	cli := codeservice.NewRPCCodeClient(m.cli)
	return cli.CodeCreatePhoneCode(ctx, in)
}

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (m *defaultCodeClient) CodeGetPhoneCode(ctx context.Context, in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	cli := codeservice.NewRPCCodeClient(m.cli)
	return cli.CodeGetPhoneCode(ctx, in)
}

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (m *defaultCodeClient) CodeDeletePhoneCode(ctx context.Context, in *code.TLCodeDeletePhoneCode) (*tg.Bool, error) {
	cli := codeservice.NewRPCCodeClient(m.cli)
	return cli.CodeDeletePhoneCode(ctx, in)
}

// CodeUpdatePhoneCodeData
// code.updatePhoneCodeData auth_key_id:long phone:string phone_code_hash:string code_data:PhoneCodeTransaction = Bool;
func (m *defaultCodeClient) CodeUpdatePhoneCodeData(ctx context.Context, in *code.TLCodeUpdatePhoneCodeData) (*tg.Bool, error) {
	cli := codeservice.NewRPCCodeClient(m.cli)
	return cli.CodeUpdatePhoneCodeData(ctx, in)
}
