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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"reflect"
)

func checkRpcUpdatesType(tl mtproto.TLObject) bool {
	switch tl.(type) {
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

//func checkRpcPushType(tl mtproto.TLObject) bool {
//	switch tl.(type) {
//	case *mtproto.TLAccountRegisterDevice,
//		 *mtproto.TLAccountUnregisterDevice,
//		 *mtproto.TLAccountRegisterDeviceLayer71,
//		 *mtproto.TLAccountUnregisterDeviceLayer71:
//		return true
//	}
//	return false
//}

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
		*mtproto.TLLangpackGetLanguagesLayer70,
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

func getSessionType(method mtproto.TLObject) int {
	var sType = kSessionUnknown

	// TODO(@benqi): check sessionType temp and genericMedia
	switch method.(type) {
	case *mtproto.TLUploadGetCdnFile,
		*mtproto.TLUploadGetCdnFileHashes,
		*mtproto.TLUploadGetFile,
		*mtproto.TLUploadGetFileHashes,
		*mtproto.TLUploadGetWebFile:
		sType = kSessionDownload
	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadSaveBigFilePart,
		*mtproto.TLUploadReuploadCdnFile:
		sType = kSessionUpload
	default:
		sType = kSessionGeneric
	}

	return sType
}

func getSessionType2(object mtproto.TLObject, sessionType *int) {
	glog.Info("getSessionType2 - ", reflect.TypeOf(object))
	if *sessionType != kSessionUnknown {
		return
	}

	switch object.(type) {
	case *mtproto.TLMsgContainer:
		msgContainer, _ := object.(*mtproto.TLMsgContainer)
		for _, m2 := range msgContainer.Messages {
			getSessionType2(m2.Object, sessionType)
			if *sessionType != kSessionUnknown {
				break
			}
		}
	case *mtproto.TLGzipPacked:
		gzipPacked, _ := object.(*mtproto.TLGzipPacked)
		if gzipPacked.Obj == nil {
			return
		}
		getSessionType2(gzipPacked.Obj, sessionType)

	case *mtproto.TLMsgCopy:
		// not use in client
		// glog.Error("android client not use msg_copy: ", object)

	case *mtproto.TLMsgsAck,

		*mtproto.TLMsgsStateReq,
		*mtproto.TLMsgsStateInfo,
		*mtproto.TLMsgsAllInfo,

		*mtproto.TLMsgResendReq,
		*mtproto.TLMsgDetailedInfo,
		*mtproto.TLMsgNewDetailedInfo:
		// unknown

	case *mtproto.TLRpcDropAnswer:
		// unknown
		// *sessionType = kSessionGeneric

	case *mtproto.TLPing,
		*mtproto.TLPingDelayDisconnect:
		// unknown

	case *mtproto.TLGetFutureSalts:
		*sessionType = kSessionGeneric
	case *mtproto.TLDestroySession:
		*sessionType = kSessionGeneric

	//////////////////////////////////////////////////////////////
	case *mtproto.TLInvokeAfterMsg,
		*mtproto.TLInvokeAfterMsgs,
		*mtproto.TLInvokeWithLayer,
		// *mtproto.TLInvokeWithoutUpdates,
		*mtproto.TLInvokeWithMessagesRange,
		*mtproto.TLInvokeWithTakeout:

		*sessionType = kSessionGeneric

	case *mtproto.TLInvokeWithoutUpdates:
		// TODO(@benqi): move TLInvokeWithoutUpdates etc to parsed_manually_types.
		dbuf := mtproto.NewDecodeBuf(object.(*mtproto.TLInvokeWithoutUpdates).Query)
		q := dbuf.Object()
		getSessionType2(q, sessionType)

	//////////////////////////////////////////////////////////////
	case *mtproto.TLUploadGetCdnFile,
		// *mtproto.TLUploadGetCdnFileHashes,
		*mtproto.TLUploadGetFile,
		// *mtproto.TLUploadGetFileHashes,
		*mtproto.TLUploadGetWebFile:

		*sessionType = kSessionDownload

	//////////////////////////////////////////////////////////////
	case *mtproto.TLUploadSaveFilePart,
		*mtproto.TLUploadSaveBigFilePart:
		// *mtproto.TLUploadReuploadCdnFile:

		*sessionType = kSessionUpload
	default:
		*sessionType = kSessionGeneric
	}
}
