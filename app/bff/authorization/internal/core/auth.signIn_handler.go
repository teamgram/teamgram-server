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
	"fmt"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/logic"
	"github.com/teamgram/teamgram-server/app/bff/authorization/internal/model"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/pkg/code/conf"

	"google.golang.org/grpc/status"
)

/*
 1. PHONE_NUMBER_UNOCCUPIED ==> setPage(5, true, params, false);
 2. SESSION_PASSWORD_NEEDED ==> invoke rpc: TL_account_getPassword
 3. error:
	if (error.text.contains("PHONE_NUMBER_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
	} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
		onBackPressed();
		setPage(0, true, null, true);
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
	} else if (error.text.startsWith("FLOOD_WAIT")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	}
*/

// AuthSignIn
// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (c *AuthorizationCore) AuthSignIn(in *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	var (
		phoneCode     = in.GetPhoneCode_STRING()
		phoneCodeHash = in.PhoneCodeHash
	)

	if phoneCode == "" {
		phoneCode = in.GetPhoneCode_FLAGSTRING().GetValue()
	}

	if phoneCode == "" || phoneCodeHash == "" {
		err := mtproto.ErrPhoneCodeEmpty
		c.Logger.Errorf("auth.sendCode - error: %v", err)
		return nil, err
	}

	// 3. check number
	// client phone number format: "+86 111 1111 1111"
	phoneNumber, err := checkPhoneNumberInvalid(in.PhoneNumber)
	if err != nil {
		c.Logger.Errorf("check phone_number(%s) error - %v", in.PhoneNumber, err)
		err = mtproto.ErrPhoneNumberInvalid
		return nil, err
	}

	// 6. check can do action
	actionType := logic.GetActionType(in)
	if err = c.svcCtx.Dao.CheckCanDoAction(c.ctx, c.MD.PermAuthKeyId, phoneNumber, actionType); err != nil {
		c.Logger.Errorf("check can do action - %s: %v", phoneNumber, err)
		return nil, err
	}

	codeData, err2 := c.svcCtx.AuthLogic.DoAuthSignIn(c.ctx,
		c.MD.PermAuthKeyId,
		phoneNumber,
		phoneCode,
		phoneCodeHash,
		func(codeData2 *model.PhoneCodeTransaction) error {
			return c.svcCtx.AuthLogic.VerifyCodeInterface.VerifySmsCode(c.ctx,
				codeData2.PhoneCodeHash,
				phoneCode,
				codeData2.PhoneCodeExtraData)
		})

	if err2 != nil {
		c.Logger.Error(err2.Error())
		err = err2
		return nil, err
	}

	if c.svcCtx.Plugin != nil {
		c.svcCtx.Plugin.OnAuthAction(c.ctx,
			c.MD.PermAuthKeyId,
			c.MD.ClientMsgId,
			c.MD.ClientAddr,
			in.PhoneNumber,
			logic.GetActionType(in),
			"auth.signIn")
	}

	// signIn successful, check phoneRegistered.
	if !codeData.PhoneNumberRegistered {
		if c.MD.Layer >= 104 {
			//  not register, next step: auth.singIn
			return mtproto.MakeTLAuthAuthorizationSignUpRequired(&mtproto.Auth_Authorization{
				// TermsOfService: model.MakeTermOfService(),
				TermsOfService: nil,
			}).To_Auth_Authorization(), nil
		} else {
			c.Logger.Errorf("auth.signIn - not registered, next step auth.signIn, %v", err)
			err = mtproto.ErrPhoneNumberUnoccupied
			return nil, err
		}
	}

	// TODO(@benqi): err handle
	// do signIn...
	var (
		user *mtproto.ImmutableUser
	)

	user, err = c.svcCtx.Dao.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
		Phone: phoneNumber,
	})
	if err != nil {
		c.Logger.Errorf("user(%s) is err - %v", phoneNumber, err)
		return nil, err
	} else if user == nil {
		c.Logger.Errorf("user(%s) is nil", phoneNumber)
		err = mtproto.ErrInternalServerError
		return nil, err
	}

	// Bind authKeyId and userId
	c.svcCtx.Dao.AuthsessionClient.AuthsessionBindAuthKeyUser(c.ctx, &authsession.TLAuthsessionBindAuthKeyUser{
		AuthKeyId: c.MD.PermAuthKeyId,
		UserId:    user.User.Id,
	})

	// Check SESSION_PASSWORD_NEEDED
	if c.svcCtx.Plugin != nil {
		if c.svcCtx.Plugin.CheckSessionPasswordNeeded(c.ctx, user.User.Id) {
			// hack
			// err = mtproto.ErrSessionPasswordNeeded
			err = status.Error(mtproto.ErrUnauthorized, fmt.Sprintf("SESSION_PASSWORD_NEEDED_%d", user.Id()))
			c.Logger.Infof("auth.signIn - registered, next step auth.checkPassword: %v", err)
			return nil, err
		}
	}

	selfUser := user.ToSelfUser()

	c.svcCtx.AuthLogic.DeletePhoneCode(c.ctx, c.MD.PermAuthKeyId, in.PhoneNumber, phoneCodeHash)
	region, _ := c.svcCtx.Dao.GetCountryAndRegionByIp(c.MD.ClientAddr)

	var (
		now     = time.Now()
		signInN *mtproto.Update
	)

	if len(c.svcCtx.Config.SignInServiceNotification) == 0 {
		signInN = mtproto.MakeSignInServiceNotification(selfUser, c.MD.PermAuthKeyId, c.MD.Client, region, c.MD.ClientAddr)
	} else {
		signInN = mtproto.MakeTLUpdateServiceNotification(&mtproto.Update{
			Popup:          false,
			InboxDate:      mtproto.MakeFlagsInt32(int32(now.Unix())),
			Type:           fmt.Sprintf("auth%d_%d", c.MD.PermAuthKeyId, now.Unix()),
			Message_STRING: "",
			Media:          mtproto.MakeTLMessageMediaEmpty(nil).To_MessageMedia(),
			Entities:       nil,
		}).To_Update()

		builder := conf.ToMessageBuildHelper(
			c.svcCtx.Config.SignInServiceNotification,
			map[string]interface{}{
				"name":     mtproto.GetUserName(selfUser),
				"now":      now.UTC(),
				"client":   c.MD.Client,
				"region":   region,
				"clientIp": c.MD.ClientAddr,
			})
		signInN.Message_STRING, signInN.Entities = mtproto.MakeTextAndMessageEntities(builder)
	}

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:        user.Id(),
			PermAuthKeyId: c.MD.PermAuthKeyId,
			Updates:       mtproto.MakeUpdatesByUpdates(signInN),
		})

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		SetupPasswordRequired: false,
		OtherwiseReloginDays:  nil,
		TmpSessions:           nil,
		FutureAuthToken:       nil,
		User:                  selfUser,
	}).To_Auth_Authorization(), nil
}
