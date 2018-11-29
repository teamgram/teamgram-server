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

type authActionType int

const (
	kOPTypeSendCode 	authActionType = 1
	kOPTypeSignUp 		authActionType = 2
	kOPTypeSignIn 		authActionType = 3
	kOPTypeLogout 		authActionType = 4
	kOPTypeResendCode 	authActionType = 5
	kOPTypeCancelCode 	authActionType = 6
)

/*
	if (error.text != null) {
		if (error.text.contains("PHONE_NUMBER_INVALID")) {
			needShowInvalidAlert(req.phone_number, false);
		} else if (error.text.contains("PHONE_PASSWORD_FLOOD")) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
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
