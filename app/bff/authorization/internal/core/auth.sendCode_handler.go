// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"errors"

	"github.com/teamgram/proto/v2/tg"
)

/***************************************
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
func (c *AuthorizationCore) AuthSendCode(in *tg.TLAuthSendCode) (*tg.AuthSentCode, error) {
	// TODO: not impl
	// c.Logger.Errorf("auth.sendCode blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, errors.New("auth.sendCode not implemented")
}
