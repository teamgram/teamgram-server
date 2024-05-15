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

package sess

import (
	"github.com/teamgram/proto/mtproto"
)

/*
	 By android source code - RequestFlagWithoutLogin:
		TL_account_getAuthorizationForm
		TL_account_getPasswordSettings
		TL_account_confirmPasswordEmail
		TL_account_updatePasswordSettings
		TL_account_getTmpPassword
		TL_account_getPassword
		TL_account_deleteAccount

		TL_auth_resendCode
		TL_auth_signIn
		TL_auth_cancelCode
		TL_auth_requestPasswordRecovery
		TL_auth_checkPassword
		TL_auth_checkRecoveryPassword
		TL_auth_recoverPassword
		TL_auth_signUp
		TL_auth_exportAuthorization
		TL_auth_importAuthorization
		TL_auth_bindTempAuthKey
		TL_auth_cancelCode

		TL_rpc_drop_answer
		TL_get_future_salts
		TL_ping

		TL_help_getNearestDc
		TL_help_getConfig
		TL_help_getCdnConfig

		TL_langpack_getLanguages
		TL_langpack_getLanguages
		TL_langpack_getDifference
		TL_langpack_getLangPack
		TL_langpack_getStrings
*/
func checkRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	// account
	case *mtproto.TLAccountGetPassword:
		return true

	// auth
	case *mtproto.TLAuthSendCode,
		*mtproto.TLAuthResendCode,
		*mtproto.TLAuthSignUp,
		*mtproto.TLAuthSignIn,
		*mtproto.TLAuthImportLoginToken,
		*mtproto.TLAuthExportedAuthorization,
		*mtproto.TLAuthExportAuthorization,
		*mtproto.TLAuthImportAuthorization,
		*mtproto.TLAuthCancelCode,
		// *mtproto.TLAuthRequestPasswordRecovery,	// TODO: before process, try fetch usrId
		// *mtproto.TLAuthRecoverPassword,			// TODO: before process, try fetch usrId
		*mtproto.TLAuthExportLoginToken,
		*mtproto.TLAuthAcceptLoginToken,
		*mtproto.TLAuthLogOut, // TODO: before process, try fetch usrId
		*mtproto.TLAuthBindTempAuthKey,
		*mtproto.TLAuthCheckPassword:

		return true

	// help
	case *mtproto.TLHelpGetConfig,
		*mtproto.TLHelpGetNearestDc,
		*mtproto.TLHelpGetAppUpdate,
		*mtproto.TLHelpGetCdnConfig,
		*mtproto.TLHelpGetAppConfig61E3F854,
		*mtproto.TLHelpGetAppConfig98914110:

		return true

	// langpack
	case *mtproto.TLLangpackGetLangPack,
		*mtproto.TLLangpackGetStrings,
		*mtproto.TLLangpackGetDifference,
		*mtproto.TLLangpackGetLanguages,
		*mtproto.TLLangpackGetLanguage:
		return true

	// TODO(@benqi): debug.
	case *mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetFile:
		return true

	// country
	case *mtproto.TLHelpGetCountriesList:
		return true

	case *mtproto.TLJsonObject:
		return true

	default:
		return false
	}
}
