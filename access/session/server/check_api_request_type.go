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

package server

import (
	"github.com/nebula-chat/chatengine/mtproto"
)

func checkRpcUpdatesType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLAccountRegisterDevice,
		*mtproto.TLAccountUnregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return false

	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadSaveBigFilePart:

		// upload connection
		return false

	case *mtproto.TLUploadGetFile,
		*mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetCdnFile,
		*mtproto.TLUploadReuploadCdnFile,
		*mtproto.TLUploadGetCdnFileHashes:

		// download
		return false

	case *mtproto.TLHelpGetConfig:
		// TODO(@benqi): 可能为TEMP，判断TEMP的规则：
		//  从android client代码看，此连接仅收到help.getConfig消息
	}

	return true
}

func checkRpcPushType(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLAccountRegisterDevice,
		*mtproto.TLAccountUnregisterDevice:
		// *mtproto.TLAccountRegisterDeviceLayer74,
		// *mtproto.TLAccountUnregisterDeviceLayer74:
		// push

		return true
	}
	return false
}

// TL_auth_exportedAuthorization
// TL_auth_exportAuthorization
// TL_auth_importAuthorization
// TL_auth_sendCode
// TL_auth_cancelCode
// TL_auth_resendCode
// TL_auth_resendCode
// TL_auth_signIn
// TL_auth_requestPasswordRecovery
// TL_auth_checkPassword
// TL_auth_recoverPassword
// TL_auth_signUp
// TL_auth_requestPasswordRecovery
// TL_auth_recoverPassword

// TL_help_getCdnConfig
// TL_help_getConfig

// TL_langpack_getLanguages
// TL_langpack_getDifference
// TL_langpack_getLangPack
// TL_langpack_getStrings
// TL_langpack_getStrings
// TL_langpack_getStrings

// TL_account_getPassword
// TL_account_deleteAccount
// TL_account_deleteAccount
// TL_account_getPassword
// TL_account_updatePasswordSettings
// TL_account_getPasswordSettings
func checkRpcWithoutLogin(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLAuthCheckedPhone,
		*mtproto.TLAuthSendCodeLayer51,
		*mtproto.TLAuthLogOut,
		*mtproto.TLAuthSendCode,
		*mtproto.TLAuthSignIn,
		*mtproto.TLAuthSignUp,
		*mtproto.TLAuthExportedAuthorization,
		*mtproto.TLAuthExportAuthorization,
		*mtproto.TLAuthImportAuthorization,
		*mtproto.TLAuthCancelCode,
		*mtproto.TLAuthResendCode,
		*mtproto.TLAuthRequestPasswordRecovery,
		*mtproto.TLAuthCheckPassword,
		*mtproto.TLAuthRecoverPassword:

		return true

	case *mtproto.TLHelpGetConfig,
		*mtproto.TLHelpGetCdnConfig,
		*mtproto.TLHelpGetNearestDc:

		return true

	case *mtproto.TLLangpackGetLanguages,
		*mtproto.TLLangpackGetDifference,
		*mtproto.TLLangpackGetLangPack,
		*mtproto.TLLangpackGetLangPackLayer71,
		*mtproto.TLLangpackGetStrings:

		return true

	case *mtproto.TLAccountGetPassword,
		*mtproto.TLAccountDeleteAccount,
		*mtproto.TLAccountUpdatePasswordSettings,
		*mtproto.TLAccountGetPasswordSettings:

		return true
	}

	return false
}

func checkRpcUploadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadSaveBigFilePart,
		*mtproto.TLUploadReuploadCdnFile:
		return true
	}
	return false
}

func checkRpcDownloadRequest(tl mtproto.TLObject) bool {
	switch tl.(type) {
	case *mtproto.TLUploadGetFile,
		*mtproto.TLUploadGetWebFile,
		*mtproto.TLUploadGetCdnFile,
		*mtproto.TLUploadGetCdnFileHashes:
		return true
	}
	return false
}
