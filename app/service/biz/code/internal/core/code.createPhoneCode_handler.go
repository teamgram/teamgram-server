/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"time"

	"github.com/teamgram/marmota/pkg/random2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"
)

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (c *CodeCore) CodeCreatePhoneCode(in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.svcCtx.Dao.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("getCachePhoneCode - error: %v", err)
		err = mtproto.ErrInternalServerError
		return nil, err
	}
	if codeData == nil {
		codeData = &code.PhoneCodeTransaction{
			AuthKeyId:             in.AuthKeyId,
			Phone:                 in.Phone,
			SessionId:             in.SessionId,
			PhoneNumberRegistered: in.PhoneNumberRegistered,
			PhoneCode:             random2.RandomNumeric(5),
			PhoneCodeHash:         crypto.GenerateStringNonce(16),
			PhoneCodeExpired:      int32(time.Now().Unix() + 3*60),
			SentCodeType:          in.SentCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          in.NextCodeType,
			State:                 code.CodeStateSend, // model.CodeStateSent
		}

	} else if in.SessionId != codeData.SessionId {
		codeData.State = code.CodeStateSend
		codeData.SessionId = in.SessionId
	}

	//switch codeData.State {
	//case model.CodeStateSend:
	//	codeData.State = model.CodeStateSent
	//case model.CodeStateSent:
	//	codeData.State = model.CodeStateSent
	//default:
	//	// codeData = newCodeData()
	//}
	//
	//if err = c.Dao.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData); err != nil {
	//	log.Errorf("putCachePhoneCode - error: %v", err)
	//	err = mtproto.ErrInternalServerError
	//	return
	//}

	return codeData, nil
}
