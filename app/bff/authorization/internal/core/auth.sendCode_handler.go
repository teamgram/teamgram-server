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

package core

import (
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/logic"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/model"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	statuspb "github.com/teamgram/teamgram-server/app/service/status/status"

	"google.golang.org/grpc/status"
)

/*
 Android client auth.sendCode#86aef0ec, handler error
 1.
	if (error->error_code == 303) {
		uint32_t migrateToDatacenterId = DEFAULT_DATACENTER_ID;

		static std::vector<std::string> migrateErrors;
		if (migrateErrors.empty()) {
			migrateErrors.push_back("NETWORK_MIGRATE_");
			migrateErrors.push_back("PHONE_MIGRATE_");
			migrateErrors.push_back("USER_MIGRATE_");
		}

		size_t count = migrateErrors.size();
		for (uint32_t a = 0; a < count; a++) {
			std::string &possibleError = migrateErrors[a];
			if (error->error_message.find(possibleError) != std::string::npos) {
				std::string num = error->error_message.substr(possibleError.size(), error->error_message.size() - possibleError.size());
				uint32_t val = (uint32_t) atoi(num.c_str());
				migrateToDatacenterId = val;
			}
		}

		if (migrateToDatacenterId != DEFAULT_DATACENTER_ID) {
			ignoreResult = true;
			moveToDatacenter(migrateToDatacenterId);
		}
	}

 2.
	if (error.text != null) {
		if (error.text.contains("PHONE_NUMBER_INVALID")) {
			needShowInvalidAlert(req.phone_number, false);
		} else if (error.text.contains("PHONE_NUMBER_FLOOD")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("PhoneNumberFlood", R.string.PhoneNumberFlood));
		} else if (error.text.contains("PHONE_NUMBER_BANNED")) {
			needShowInvalidAlert(req.phone_number, true);
		} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
		} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
		} else if (error.text.startsWith("FLOOD_WAIT")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
		} else if (error.code != -1000) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
		}
	}
*/
// func makeAuthSendCodeByLayer51(request *mtproto.TLAuthSendCodeLayer51) *mtproto.TLAuthSendCode {
//	return &mtproto.TLAuthSendCode{
//		AllowFlashcall: request.AllowFlashcall,
//		PhoneNumber:    request.PhoneNumber,
//		CurrentNumber:  request.CurrentNumber,
//		ApiId:          request.ApiId,
//		ApiHash:        request.ApiHash,
//	}
// }
//

/*
## Possible errors

|Code |	Type |	Description|
|:-:|:-:|:-:|
|400 |	API_ID_INVALID |	API ID invalid |
|400 |	API_ID_PUBLISHED_FLOOD |	This API id was published somewhere, you can't use it now |
|400 |	BOT_METHOD_INVALID |	This method can't be used by a bot |
|400 |	INPUT_REQUEST_TOO_LONG |	The request is too big |
|303 |	NETWORK_MIGRATE_X |	Repeat the query to data-center X |
|303 |	PHONE_MIGRATE_X |	Repeat the query to data-center X |
|400 |	PHONE_NUMBER_APP_SIGNUP_FORBIDDEN |	You can't sign up using this app |
|400 |	PHONE_NUMBER_BANNED |	The provided phone number is banned from telegram |
|400 |	PHONE_NUMBER_FLOOD |	You asked for the code too many times. |
|400 |	PHONE_NUMBER_INVALID |	Invalid phone number |
|406 |	PHONE_PASSWORD_FLOOD |	You have tried logging in too many times |
|400 |	PHONE_PASSWORD_PROTECTED |	This phone is password protected |
|400 |	SMS_CODE_CREATE_FAILED |	An error occurred while creating the SMS code |
*/

// AuthSendCode
// auth.sendCode#a677244f phone_number:string api_id:int api_hash:string settings:CodeSettings = auth.SentCode;
func (c *AuthorizationCore) AuthSendCode(in *mtproto.TLAuthSendCode) (*mtproto.Auth_SentCode, error) {
	rValue, err := c.authSendCode(c.MD.PermAuthKeyId, c.MD.SessionId, in)
	if err != nil {
		c.Logger.Errorf("auth.sendCode - error: {%v}", err)
		return nil, err
	}

	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.PermAuthKeyId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			logic.GetActionType(in),
			"auth.sendCode")
	}

	return rValue, nil
}

func (c *AuthorizationCore) authSendCode(authKeyId, sessionId int64, request *mtproto.TLAuthSendCode) (reply *mtproto.Auth_SentCode, err error) {
	// 1. check api_id and api_hash
	if err = c.svcCtx.Dao.CheckApiIdAndHash(request.ApiId, request.ApiHash); err != nil {
		c.Logger.Errorf("invalid api: {api_id: %d, api_hash: %s}", request.ApiId, request.ApiHash)
		return
	}

	// 2. check allow_flashcall and current_number
	// if allow_flashcall is true then current_number is true
	/*
		var currentNumber = request.Settings.CurrentNumber
		if request.Settings.CurrentNumber {
			currentNumber = false
		} else {
			currentNumber = mtproto.FromBool(request.CurrentNumber)
		}
		// TODO(@benqi): check allow_flashcall rule
		if !currentNumber && request.GetAllowFlashcall() {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("current_number is true but allow_flashcall is false - %v", err)
			return nil, err
		}
	*/

	// 3. check number

	// client phone number format: "+86 111 1111 1111"
	phoneNumber, err := checkPhoneNumberInvalid(request.PhoneNumber)
	if err != nil {
		c.Logger.Errorf("check phone_number(%s) error - %v", request.PhoneNumber, err)
		err = mtproto.ErrPhoneNumberInvalid
		return
	}

	// 4. MIGRATE datacenter
	// 	303	NETWORK_MIGRATE_X	重复查询到数据中心X
	// 	303	PHONE_MIGRATE_X	重复查询到数据中心X
	//
	// TODO(@benqi): MIGRATE datacenter
	// android client:
	//  migrateErrors.push_back("NETWORK_MIGRATE_");
	//  migrateErrors.push_back("PHONE_MIGRATE_");
	//  migrateErrors.push_back("USER_MIGRATE_");
	//
	// https://core.telegram.org/api/datacenter
	// The auth.sendCode method is the basic entry point when registering a new user or authorizing an existing user.
	//   95% of all redirection cases to a different DC will occure when invoking this method.
	//
	// The client does not yet know which DC it will be associated with; therefore,
	//   it establishes an encrypted connection to a random address and sends its query to that address.
	// Having received a phone_number from a client,
	// 	 we can find out whether or not it is registered in the system.
	//   If it is, then, if necessary, instead of sending a text message,
	//   we request that it establish a connection with a different DC first (PHONE_MIGRATE_X error).
	// If we do not yet have a user with this number, we examine its IP-address.
	//   We can use it to identify the closest DC.
	//   Again, if necessary, we redirect the user to a different DC (NETWORK_MIGRATE_X error).
	//
	// if userDO == nil {
	//	// phone registered
	//	// TODO(@benqi): 由phoneNumber和ip优选
	// } else {
	//	// TODO(@benqi): 由userId优选
	// }

	// 5. Check INPUT_REQUEST_TOO_LONG
	// TODO(@benqi):
	// 	400	INPUT_REQUEST_TOO_LONG	The request is too big

	// 5. banned phone number
	if c.svcCtx.Plugin != nil {
		banned, _ := c.svcCtx.Plugin.CheckPhoneNumberBanned(c.ctx, phoneNumber)
		if banned {
			c.Logger.Errorf("{phone_number: %s} banned: %v", phoneNumber, err)
			err = mtproto.ErrPhoneNumberBanned
			return
		}
	}

	// 400	PHONE_NUMBER_FLOOD	You asked for the code too many times.
	// phone number flood
	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	// 6. check can do action
	actionType := logic.GetActionType(request)
	if err = c.svcCtx.Dao.CheckCanDoAction(c.ctx, authKeyId, phoneNumber, actionType); err != nil {
		c.Logger.Errorf("check can do action - %s: %v", phoneNumber, err)
		return
	}

	// logic
	// Always crated new phoneCode
	var (
		phoneRegistered = false
		user            *mtproto.ImmutableUser
	)

	if user, err = c.svcCtx.Dao.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
		Phone: phoneNumber,
	}); err != nil {
		if nErr, ok := status.FromError(err); ok {
			// mtproto.ErrPhoneNumberUnoccupied
			_ = nErr
			err = nil
		} else {
			c.Logger.Errorf("checkPhoneNumberExist error: %v", err)
			return
		}
		//st, ok := errors.Cause(err).(*status.Error)
		//if !ok {
		//	c.Logger.Errorf("checkPhoneNumberExist error: %v, type: %s", st, reflect.TypeOf(err))
		//	return
		//}
		//c.Logger.Errorf("status: {code: %d, message: %s}", st.Code(), st.Message())
		// phoneRegistered = true
		//if st.Code() != int(ecode.RequestErr) && st.Message() != "RequestErr" {
		//	// t.Fatalf("testECodeStatus must return code: -400, message: RequestErr get: code: %d, message: %s", st.Code(), st.Message())
		//}
		//
		//*ecode.Status
		//switch nErr := errors.Cause(err).(type) {
		//case ecode.Codes:
		//	phoneRegistered = true
		//default:
		//	c.Logger.Errorf("checkPhoneNumberExist error: %v", nErr)
		//	return
		//}
	} else {
		phoneRegistered = true
	}

	if phoneRegistered {
		// https://core.telegram.org/api/auth#future-auth-tokens
		// TODO:
		//  At all times, the future auth token database should contain at most 20 tokens:
		//  evict older tokens as new tokens are added to stay below this limit.
		for _, v := range request.Settings.GetLogoutTokens() {
			id, _ := c.svcCtx.Dao.GetFutureAuthToken(c.ctx, v)
			if id == user.Id() {
				// Bind authKeyId and userId
				c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
					AuthKeyId: c.MD.PermAuthKeyId,
					UserId:    user.User.Id,
				})

				// Del
				c.svcCtx.Dao.DelFutureAuthToken(c.ctx, v)

				// Check SESSION_PASSWORD_NEEDED
				if c.svcCtx.Plugin != nil {
					if c.svcCtx.Plugin.CheckSessionPasswordNeeded(c.ctx, user.User.Id) {
						// hack
						// err = mtproto.ErrSessionPasswordNeeded
						err = status.Error(mtproto.ErrUnauthorized, fmt.Sprintf("SESSION_PASSWORD_NEEDED_%d", user.Id()))
						c.Logger.Infof("auth.sendCode - future-auth-tokens, next step auth.checkPassword: %v", err)
						return nil, err
					}
				}

				return mtproto.MakeTLAuthSentCodeSuccess(&mtproto.Auth_SentCode{
					Authorization: mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
						SetupPasswordRequired: false,
						OtherwiseReloginDays:  nil,
						TmpSessions:           nil,
						FutureAuthToken:       nil,
						User:                  user.ToSelfUser(),
					}).To_Auth_Authorization(),
				}).To_Auth_SentCode(), nil
			}
		}
	}
	// phoneRegistered = user != nil

	// codeLogic := logic.NewAuthSignLogic(s.AuthCore)
	codeData, err2 := c.svcCtx.AuthLogic.DoAuthSendCode(c.ctx,
		authKeyId,
		sessionId,
		phoneNumber,
		phoneRegistered,
		request.Settings.AllowFlashcall,
		request.Settings.CurrentNumber,
		request.ApiId,
		request.ApiHash,
		func(codeData2 *model.PhoneCodeTransaction) error {
			if codeData2.State == model.CodeStateSent {
				c.Logger.Infof("codeSent")
				return nil
			}

			var (
				needSendSms = true
			)

			if phoneRegistered {
				if user.GetUser().GetUserType() == userpb.UserTypeTest {
					needSendSms = false
					codeData2.SentCodeType = model.SentCodeTypeApp
					codeData2.PhoneCode = "12345"
					codeData2.PhoneCodeExtraData = "12345"
					c.Logger.Infof("is test server: %v", codeData2)
				} else {
					if status, _ := c.svcCtx.StatusClient.StatusGetUserOnlineSessions(c.ctx, &statuspb.TLStatusGetUserOnlineSessions{
						UserId: user.User.Id,
					}); len(status.GetUserSessions()) > 0 {
						c.Logger.Infof("user online")
						needSendSms = false
						codeData2.SentCodeType = model.SentCodeTypeApp
						codeData2.PhoneCodeExtraData = codeData2.PhoneCode
					}
				}
				threading2.WrapperGoFunc(c.ctx, nil, func(ctx context.Context) {
					c.pushSignInMessage(ctx, user.Id(), codeData2.PhoneCode)
				})
			}

			if needSendSms {
				c.Logger.Infof("send code by sms")
				extraData, err2 := c.svcCtx.AuthLogic.VerifyCodeInterface.SendSmsVerifyCode(
					context.Background(),
					phoneNumber,
					codeData2.PhoneCode,
					codeData2.PhoneCodeHash)
				if err2 != nil {
					c.Logger.Errorf("send sms code error: %v", err2)
					return err2
				} else {
					// codeData2.SentCodeType = model.CodeTypeSms
					codeData2.SentCodeType = model.SentCodeTypeSms
					codeData2.PhoneCodeExtraData = extraData
				}
			}

			//if user.User.UserType == userpb.UserTypeTest {
			//	c.Logger.Infof("test user: %s, %s", phoneNumber, user)
			//	codeData2.SentCodeType = model.CodeTypeApp
			//	codeData2.PhoneCode = "12345"
			//	codeData2.PhoneCodeExtraData = "12345"
			//	go func() {
			//		// c.pushSignInMessage(context.Background(), user.Id, codeData2.PhoneCode)
			//	}()
			//} else {
			//	var (
			//		online = false
			//	)
			//
			//	if phoneRegistered {
			//		if status, _ := c.svcCtx.StatusClient.StatusGetUserOnlineSessions(c.ctx, &status.TLStatusGetUserOnlineSessions{
			//			UserId: user.User.Id,
			//		}); len(status.GetUserSessions()) > 0 {
			//			c.Logger.Infof("user online")
			//			online = true
			//
			//			codeData2.SentCodeType = model.CodeTypeApp
			//			codeData2.PhoneCodeExtraData = codeData2.PhoneCode
			//			go func() {
			//				// s.pushSignInMessage(context.Background(), user.Id, codeData2.PhoneCode)
			//			}()
			//		}
			//		// &&
			//	}
			//
			//	if !phoneRegistered || !online {
			//		c.Logger.Infof("send code by sms")
			//		if extraData, err := c.svcCtx.AuthLogic.VerifyCodeInterface.SendSmsVerifyCode(
			//			context.Background(),
			//			phoneNumber,
			//			codeData2.PhoneCode,
			//			codeData2.PhoneCodeHash); err != nil {
			//			return err
			//		} else {
			//			codeData2.SentCodeType = model.CodeTypeSms
			//			codeData2.PhoneCodeExtraData = extraData
			//		}
			//	}
			//}

			codeData2.NextCodeType = model.CodeTypeSms
			codeData2.State = model.CodeStateSent
			codeData2.PhoneNumberRegistered = phoneRegistered

			return nil
		})

	if err2 != nil {
		c.Logger.Errorf("auth.sendCode - error: %v", err2)
		err = err2
		return
	}

	reply = codeData.ToAuthSentCode()
	return
}
