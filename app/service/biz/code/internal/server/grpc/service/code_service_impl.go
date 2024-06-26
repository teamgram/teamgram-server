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
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/app/service/biz/code/internal/core"
)

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (s *Service) CodeCreatePhoneCode(ctx context.Context, request *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("code.createPhoneCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.CodeCreatePhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("code.createPhoneCode - reply: %s", r)
	return r, err
}

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (s *Service) CodeGetPhoneCode(ctx context.Context, request *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("code.getPhoneCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.CodeGetPhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("code.getPhoneCode - reply: %s", r)
	return r, err
}

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (s *Service) CodeDeletePhoneCode(ctx context.Context, request *code.TLCodeDeletePhoneCode) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("code.deletePhoneCode - metadata: %s, request: %s", c.MD, request)

	r, err := c.CodeDeletePhoneCode(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("code.deletePhoneCode - reply: %s", r)
	return r, err
}

// CodeUpdatePhoneCodeData
// code.updatePhoneCodeData auth_key_id:long phone:string phone_code_hash:string code_data:PhoneCodeTransaction = Bool;
func (s *Service) CodeUpdatePhoneCodeData(ctx context.Context, request *code.TLCodeUpdatePhoneCodeData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("code.updatePhoneCodeData - metadata: %s, request: %s", c.MD, request)

	r, err := c.CodeUpdatePhoneCodeData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("code.updatePhoneCodeData - reply: %s", r)
	return r, err
}
