// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package logic

import (
	"context"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/dao"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/model"
	"github.com/teamgram/teamgram-server/pkg/code"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	*dao.Dao
	code.VerifyCodeInterface
}

func NewAuthSignLogic(dao *dao.Dao, code2 code.VerifyCodeInterface) *AuthLogic {
	return &AuthLogic{
		Dao:                 dao,
		VerifyCodeInterface: code2,
	}
}

func (m *AuthLogic) DoAuthSendCode(
	ctx context.Context,
	authKeyId int64,
	sessionId int64,
	phoneNumber string,
	phoneRegistered,
	allowFlashCall,
	currentNumber bool,
	apiId int32,
	apiHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {

	sentCodeType, nextCodeType := model.MakeCodeType(phoneRegistered, allowFlashCall, currentNumber)
	if codeData, err = m.Dao.CreatePhoneCode(ctx,
		authKeyId,
		sessionId,
		phoneNumber,
		phoneRegistered,
		sentCodeType,
		nextCodeType,
		model.CodeStateSend); err != nil {
		return
	}

	if cb != nil {
		if err = cb(codeData); err != nil {
			return
		}
	}

	// TODO(@benqi): after sendSms success, save codeData
	m.Dao.UpdatePhoneCodeData(ctx, authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)

	//if codeData.State == model.CodeStateSend {
	//	//switch codeData.State {
	//	//case model.CodeStateSend:
	//	//	codeData.State = model.CodeStateSent
	//	//case model.CodeStateSent:
	//	//	codeData.State = model.CodeStateSent
	//	//default:
	//	//	// codeData = newCodeData()
	//	//}
	//	codeData.State = model.CodeStateSent
	//
	//	go func() {
	//		// 400	SMS_CODE_CREATE_FAILED	An error occurred while creating the SMS code
	//		if m.VerifyCodeInterface != nil {
	//			m.VerifyCodeInterface.SendSmsVerifyCode(context.Background(), phoneNumber, codeData.PhoneCode, codeData.PhoneCodeHash)
	//		}
	//
	//		// TODO(@benqi): after sendSms success, save codeData
	//		m.AuthCore.UpdatePhoneCodeData(context.Background(), authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)
	//	}()
	//}

	return
}

// DoAuthReSendCode
// auth.resendCode
func (m *AuthLogic) DoAuthReSendCode(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.Dao.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错
	//
	// 误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}
	//
	//// TODO(@benqi): check phone code valid, only number etc.
	//if do.Code == "" {
	//	err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
	//	log.Error(err.Error())
	//	return err
	//}

	// check state invalid.
	if codeData.State != model.CodeStateSent && codeData.State != model.CodeStateSignIn {
		err = mtproto.ErrInternalServerError
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		// TODO(@benqi): update timeout state?
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	if cb != nil {
		err = cb(codeData)
		if err != nil {
			return
		}
	}

	m.Dao.UpdatePhoneCodeData(context.Background(), authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)

	//go func() {
	//	if m.VerifyCodeInterface != nil {
	//		m.VerifyCodeInterface.SendSmsVerifyCode(context.Background(), phoneNumber, codeData.PhoneCode, codeData.PhoneCodeHash)
	//	}
	//
	//	// TODO(@benqi): after sendSms success, save codeData
	//	codeData.State = model.CodeStateSent
	//	m.AuthCore.UpdatePhoneCodeData(context.Background(), authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)
	//}()

	return
}

// DoAuthCancelCode
// auth.cancelCode
func (m *AuthLogic) DoAuthCancelCode(ctx context.Context, authKeyId int64, phoneNumber, phoneCodeHash string) error {
	return m.Dao.DeletePhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash)
}

func (m *AuthLogic) DoAuthSignIn(ctx context.Context,
	authKeyId int64,
	phoneNumber,
	phoneCode,
	phoneCodeHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.Dao.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	if codeData.State != model.CodeStateOk &&
		codeData.State != model.CodeStateSent &&
		codeData.State != model.CodeStateSignIn {
		logx.WithContext(ctx).Errorf("error - invalid codeData state: %v", codeData)
		err = mtproto.ErrInternalServerError
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		// TODO(@benqi): update timeout state?
		// code.dao.AuthPhoneTransactionsDAO.UpdateState(kCodeStateTimeout, do.Id)
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	// TODO(@benqi): check phone code valid, only number etc.
	//if phoneCode != "12345" && codeData.PhoneCode != phoneCode {
	//	err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID)
	//	return err
	//}

	if cb != nil {
		if err = cb(codeData); err != nil {
			err = mtproto.ErrPhoneCodeInvalid
			return
		}
	}
	//if m.VerifyCodeInterface != nil {
	//	err = m.VerifyCodeInterface.VerifySmsCode(ctx, codeData.PhoneCodeHash, phoneCode, codeData.PhoneCode)
	//	if err != nil {
	//		err = mtproto.ErrPhoneCodeInvalid
	//		return
	//	}
	//} else if phoneCode != "12345" {
	//	err = mtproto.ErrPhoneCodeInvalid
	//	return
	//}

	// code.state = kCodeStateSignIn
	if codeData.PhoneNumberRegistered {
		codeData.State = model.CodeStateOk
	} else {
		codeData.State = model.CodeStateSignIn
	}

	err = m.Dao.UpdatePhoneCodeData(ctx, authKeyId, phoneNumber, phoneCodeHash, codeData)
	return
}

// DoAuthSignUp
// TODO(@benqi): 合并DoSignUp和DoSignIn部分代码
func (m *AuthLogic) DoAuthSignUp(ctx context.Context, authKeyId int64, phoneNumber string, phoneCode *string, phoneCodeHash string) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.Dao.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	// TODO(@benqi): remote client error, state is Ok
	if codeData.State != model.CodeStateOk &&
		codeData.State != model.CodeStateSignIn &&
		codeData.State != model.CodeStateDeleted &&
		codeData.State != model.CodeStateSignUp {
		err = mtproto.ErrInputRequestInvalid
		logx.WithContext(ctx).Errorf("invalid code state(%d) - err: %v", codeData.State, err)
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		// TODO(@benqi): update timeout state?
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	// auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
	if phoneCode != nil {
		if err = m.VerifyCodeInterface.VerifySmsCode(ctx,
			codeData.PhoneCodeHash,
			*phoneCode,
			codeData.PhoneCodeExtraData); err != nil {
			return
		}
		//if m.VerifyCodeInterface != nil {
		//	err = m.VerifyCodeInterface.VerifySmsCode(ctx, codeData.PhoneCodeHash, *phoneCode, codeData.PhoneCode)
		//	if err != nil {
		//		log.Errorf("verifySmsCode error: %v", err)
		//		err = mtproto.ErrPhoneCodeInvalid
		//		return
		//	}
		//} else if *phoneCode != "12345" {
		//	log.Errorf("verifySmsCode error: %v", err)
		//	err = mtproto.ErrPhoneCodeInvalid
		//	return
		//}
	}

	codeData.State = model.CodeStateOk

	/*
		if !m.AuthCore.UpdatePhoneCodeData(ctx, authKeyId, phoneNumber, phoneCodeHash, codeData) {
			err = mtproto.ErrInternalServerError
			return
		}

		m.PhoneCodeTransaction = codeData
	*/
	return
}
