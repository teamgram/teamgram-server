// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package auth

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/crypto"
	"github.com/nebula-chat/chatengine/pkg/random2"
	"time"
)

// TODO(@benqi): 当前测试环境code统一为"12345"
// TODO(@benqi): 限制同一个authKeyId
// TODO(@benqi): 使用redis

/**
  auth.codeTypeSms#72a3158c = auth.CodeType;
  auth.codeTypeCall#741cd3e3 = auth.CodeType;
  auth.codeTypeFlashCall#226ccefb = auth.CodeType;

  auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
  auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
  auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
  auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
*/

const (
	kCodeType_None      = 0
	kCodeType_App       = 1
	kCodeType_Sms       = 2
	kCodeType_Call      = 3
	kCodeType_FlashCall = 4
)

// dataType 实现 lazy create
const (
	kDBTypeNone   = 0
	kDBTypeCreate = 1
	kDBTypeLoad   = 2
	kDBTypeUpdate = 3
	kDBTypeDelete = 3
)

const (
	kCodeStateNone    = 0
	kCodeStateOk      = 1
	kCodeStateSent    = 2
	kCodeStateSignIn  = 3
	kCodeStateSignUp  = 4
	kCodeStateDeleted = -1
	kCodeStateTimeout = -2
)

//type sendCodeCallback interface {
//	SendCode(string, string, int) error
//}

// TODO(@benqi): Add phone region
type phoneCodeData struct {
	authKeyId        int64
	phoneNumber      string
	code             string
	codeHash         string
	codeExpired      int32
	sentCodeType     int
	flashCallPattern string
	nextCodeType     int
	state            int
	dataType         int // dataType: kDBTypeCreate, kDBTypeLoad
	tableId          int64
	// codeCallback     sendCodeCallback
	dao *authsDAO
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// 由params(phoneRegistered, allowFlashCall, currentNumber)确定sentType和nextType
func makeCodeType(phoneRegistered, allowFlashCall, currentNumber bool) (int, int) {

	//if phoneRegistered {
	//	// TODO(@benqi): check other session online
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeApp{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//} else {
	//	// TODO(@benqi): sentCodeTypeFlashCall and sentCodeTypeCall, nextType
	//	// telegramd, we only use sms
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeSms{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//
	//	// TODO(@benqi): nextType
	//	// authSentCode.SetNextType()
	//}

	sentCodeType := kCodeType_App
	nextCodeType := kCodeType_None
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int) *mtproto.Auth_CodeType {
	switch codeType {
	case kCodeType_Sms:
		return mtproto.NewTLAuthCodeTypeSms().To_Auth_CodeType()
	case kCodeType_Call:
		return mtproto.NewTLAuthCodeTypeCall().To_Auth_CodeType()
	case kCodeType_FlashCall:
		return mtproto.NewTLAuthCodeTypeFlashCall().To_Auth_CodeType()
	default:
		return nil
	}
}

func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case kCodeType_App:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeApp,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_Sms:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeSms,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_Call:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeCall,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_FlashCall:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeFlashCall,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length:  int32(codeLength),
				Pattern: pattern,
			},
		}
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		glog.Error("makeAuthSentCodeType - ", err)
		panic(err)
	}

	return
}

func (code *phoneCodeData) String() string {
	return fmt.Sprintf("{authKeyId: %d, phoneNumber: %s, codeHash: %s, state: %d}", code.authKeyId, code.phoneNumber, code.codeHash, code.state)
}

func (code *phoneCodeData) checkDataType(validType int) {
	// TODO(@benqi): panic
	if code.dataType != validType {
		glog.Fatal("invalid dataType")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//func (code *phoneCodeData) fromDO(do *dataobject.AuthPhoneTransactionsDO) {
//	// TODO(@benqi): 111
//	do := &dataobject.AuthPhoneTransactionsDO{
//		ApiId:           apiId,
//		ApiHash:         apiHash,
//		PhoneNumber:     code.phoneNumber,
//		Code:            code.code,
//		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
//		TransactionHash: code.codeHash,
//	}
//}
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////
//func (code *phoneCodeData) toDO() *dataobject.AuthPhoneTransactionsDO {
//	return nil
//}

//func (code *phoneCodeData) doSendCodeCallback() error {
//	if code.codeCallback != nil {
//		return code.codeCallback.SendCode(code.code, code.codeHash, code.sentCodeType)
//	}
//
//	// TODO(@benqi): 测试环境默认发送成功
//	return nil
//}

// auth.sendCode
func (code *phoneCodeData) DoSendCode(
	phoneRegistered,
	allowFlashCall,
	currentNumber bool,
	apiId int32,
	apiHash string,
	sendSmsF func(phoneNumber, code, codeHash string, sentCodeType int) error) error {

	code.checkDataType(kDBTypeCreate)

	// 使用最简单的办法，每次新建
	sentCodeType, nextCodeType := makeCodeType(phoneRegistered, allowFlashCall, currentNumber)
	// TODO(@benqi): gen rand number

	if sendSmsF == nil {
		code.code = "12345"
	} else {
		code.code = random2.RandomNumeric(5)
	}

	// code.codeHash = fmt.Sprintf("%20d", helper.NextSnowflakeId())
	code.codeHash = crypto.GenerateStringNonce(16)
	code.codeExpired = int32(time.Now().Unix() + 15*60)
	code.sentCodeType = sentCodeType
	code.nextCodeType = nextCodeType

	//err := code.doSendCodeCallback()
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}

	// sendSmsF
	// save
	do := &dataobject.AuthPhoneTransactionsDO{
		AuthKeyId:        code.authKeyId,
		PhoneNumber:      code.phoneNumber,
		Code:             code.code,
		CodeExpired:      code.codeExpired,
		TransactionHash:  code.codeHash,
		SentCodeType:     int8(code.sentCodeType),
		FlashCallPattern: code.flashCallPattern,
		NextCodeType:     int8(code.nextCodeType),
		State:            kCodeStateSent,
		ApiId:            apiId,
		ApiHash:          apiHash,
		CreatedTime:      time.Now().Unix(),
	}
	code.tableId = code.dao.AuthPhoneTransactionsDAO.Insert(do)
	//// TODO(@benqi):
	//lastCreatedAt := time.Unix(time.Now().Unix()-15*60, 0).Format("2006-01-02 15:04:05")
	//do := code.dao.AuthPhoneTransactionsDAO.SelectByPhoneAndApiIdAndHash(code.phoneNumber, apiId, apiHash, lastCreatedAt)
	//if do == nil {
	//} else {
	//	// TODO(@benqi): FLOOD_WAIT_X, too many attempts, please try later.
	//}

	if sendSmsF != nil {
		return sendSmsF(code.phoneNumber, code.code, code.codeHash, code.sentCodeType)
	}

	return nil
}

// auth.resendCode
func (code *phoneCodeData) DoReSendCode(sendSmsF func(phoneNumber, code, codeHash string, sentCodeType int) error) error {
	code.checkDataType(kDBTypeLoad)

	do := code.dao.AuthPhoneTransactionsDAO.SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(err)
		return err
	}

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}
	//
	//// TODO(@benqi): check phone code valid, only number etc.
	//if do.Code == "" {
	//	err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
	//	glog.Error(err)
	//	return err
	//}

	// check state invalid.
	if do.State != kCodeStateSent {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// code.dao.AuthPhoneTransactionsDAO.UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	code.code = do.Code
	// TODO(@benqi): load from db
	code.codeExpired = do.CodeExpired
	code.sentCodeType = int(do.SentCodeType)
	code.flashCallPattern = do.FlashCallPattern
	code.nextCodeType = int(do.NextCodeType)
	code.state = int(do.State)
	code.tableId = do.Id

	if sendSmsF != nil {
		return sendSmsF(code.phoneNumber, code.code, code.codeHash, code.sentCodeType)
	}

	return nil
}

// auth.cancelCode
func (code *phoneCodeData) DoCancelCode() bool {
	code.dao.AuthPhoneTransactionsDAO.Delete(int8(kCodeStateDeleted), code.authKeyId, code.phoneNumber, code.codeHash)
	return true
}

func (code *phoneCodeData) DoSignIn(phoneCode string, phoneRegistered bool) error {
	defer func() {
		if code.tableId != 0 {
			// Update attempts
			code.dao.AuthPhoneTransactionsDAO.UpdateAttempts(code.tableId)
		}
	}()

	code.checkDataType(kDBTypeLoad)

	do := code.dao.AuthPhoneTransactionsDAO.SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(code, ", error: ", err)
		return err
	}
	code.tableId = do.Id

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	if do.State != kCodeStateSent && do.State != kCodeStateSignIn {
		glog.Info("error - state ", do.State)
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// code.dao.AuthPhoneTransactionsDAO.UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	// TODO(@benqi): check phone code valid, only number etc.
	if do.Code != phoneCode {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
		glog.Error(err)
		return err
	}

	// code.code = do.Code
	// TODO(@benqi): load from db
	// code.codeExpired = do.CodeExpired
	// code.sentCodeType = int(do.SentCodeType)
	// code.flashCallPattern = do.FlashCallPattern
	// code.nextCodeType = int(do.NextCodeType)

	// code.state = kCodeStateSignIn
	if phoneRegistered {
		code.state = kCodeStateOk
	} else {
		code.state = kCodeStateSignIn
	}

	// update state
	code.dao.AuthPhoneTransactionsDAO.UpdateState(int8(code.state), code.tableId)
	return nil
}

// TODO(@benqi): 合并DoSignUp和DoSignIn部分代码
func (code *phoneCodeData) DoSignUp(phoneCode string) error {
	defer func() {
		if code.tableId != 0 {
			// Update attempts
			code.dao.AuthPhoneTransactionsDAO.UpdateAttempts(code.tableId)
		}
	}()

	code.checkDataType(kDBTypeLoad)

	do := code.dao.AuthPhoneTransactionsDAO.SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(err)
		return err
	}
	code.tableId = do.Id

	// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	// TODO(@benqi): remote client error, state is Ok
	if do.State != kCodeStateSignIn && do.State != kCodeStateDeleted && do.State != kCodeStateSignUp {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// code.dao.AuthPhoneTransactionsDAO.UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	// TODO(@benqi): check phone code valid, only number etc.
	if do.Code != phoneCode {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
		glog.Error(err)
		return err
	}

	code.state = kCodeStateOk

	// update state
	code.dao.AuthPhoneTransactionsDAO.UpdateState(int8(code.state), code.tableId)
	return nil
}

// 如果手机号已经注册，检查是否有其他设备在线，有则使用sentCodeTypeApp
// 否则使用sentCodeTypeSms
// TODO(@benqi): 有则使用sentCodeTypeFlashCall和entCodeTypeCall？？
func (code *phoneCodeData) ToAuthSentCode(phoneRegistered bool) *mtproto.TLAuthSentCode {
	// TODO(@benqi): only use sms

	authSentCode := &mtproto.TLAuthSentCode{Data2: &mtproto.Auth_SentCode_Data{
		PhoneRegistered: phoneRegistered,
		Type:            makeAuthSentCodeType(code.sentCodeType, len(code.code), code.flashCallPattern),
		PhoneCodeHash:   code.codeHash,
		NextType:        makeAuthCodeType(code.nextCodeType),
		Timeout:         60, // TODO(@benqi): 默认60s
	}}
	return authSentCode
}
