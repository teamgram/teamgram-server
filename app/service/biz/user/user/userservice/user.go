/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package userservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ *tg.Bool

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"/user.RPCUser/user.getLastSeens": kitex.NewMethodInfo(
		getLastSeensHandler,
		newGetLastSeensArgs,
		newGetLastSeensResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateLastSeen": kitex.NewMethodInfo(
		updateLastSeenHandler,
		newUpdateLastSeenArgs,
		newUpdateLastSeenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getLastSeen": kitex.NewMethodInfo(
		getLastSeenHandler,
		newGetLastSeenArgs,
		newGetLastSeenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getImmutableUser": kitex.NewMethodInfo(
		getImmutableUserHandler,
		newGetImmutableUserArgs,
		newGetImmutableUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getMutableUsers": kitex.NewMethodInfo(
		getMutableUsersHandler,
		newGetMutableUsersArgs,
		newGetMutableUsersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getImmutableUserByPhone": kitex.NewMethodInfo(
		getImmutableUserByPhoneHandler,
		newGetImmutableUserByPhoneArgs,
		newGetImmutableUserByPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getImmutableUserByToken": kitex.NewMethodInfo(
		getImmutableUserByTokenHandler,
		newGetImmutableUserByTokenArgs,
		newGetImmutableUserByTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setAccountDaysTTL": kitex.NewMethodInfo(
		setAccountDaysTTLHandler,
		newSetAccountDaysTTLArgs,
		newSetAccountDaysTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getAccountDaysTTL": kitex.NewMethodInfo(
		getAccountDaysTTLHandler,
		newGetAccountDaysTTLArgs,
		newGetAccountDaysTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getNotifySettings": kitex.NewMethodInfo(
		getNotifySettingsHandler,
		newGetNotifySettingsArgs,
		newGetNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getNotifySettingsList": kitex.NewMethodInfo(
		getNotifySettingsListHandler,
		newGetNotifySettingsListArgs,
		newGetNotifySettingsListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setNotifySettings": kitex.NewMethodInfo(
		setNotifySettingsHandler,
		newSetNotifySettingsArgs,
		newSetNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.resetNotifySettings": kitex.NewMethodInfo(
		resetNotifySettingsHandler,
		newResetNotifySettingsArgs,
		newResetNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getAllNotifySettings": kitex.NewMethodInfo(
		getAllNotifySettingsHandler,
		newGetAllNotifySettingsArgs,
		newGetAllNotifySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getGlobalPrivacySettings": kitex.NewMethodInfo(
		getGlobalPrivacySettingsHandler,
		newGetGlobalPrivacySettingsArgs,
		newGetGlobalPrivacySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setGlobalPrivacySettings": kitex.NewMethodInfo(
		setGlobalPrivacySettingsHandler,
		newSetGlobalPrivacySettingsArgs,
		newSetGlobalPrivacySettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getPrivacy": kitex.NewMethodInfo(
		getPrivacyHandler,
		newGetPrivacyArgs,
		newGetPrivacyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setPrivacy": kitex.NewMethodInfo(
		setPrivacyHandler,
		newSetPrivacyArgs,
		newSetPrivacyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.checkPrivacy": kitex.NewMethodInfo(
		checkPrivacyHandler,
		newCheckPrivacyArgs,
		newCheckPrivacyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.addPeerSettings": kitex.NewMethodInfo(
		addPeerSettingsHandler,
		newAddPeerSettingsArgs,
		newAddPeerSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getPeerSettings": kitex.NewMethodInfo(
		getPeerSettingsHandler,
		newGetPeerSettingsArgs,
		newGetPeerSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.deletePeerSettings": kitex.NewMethodInfo(
		deletePeerSettingsHandler,
		newDeletePeerSettingsArgs,
		newDeletePeerSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.changePhone": kitex.NewMethodInfo(
		changePhoneHandler,
		newChangePhoneArgs,
		newChangePhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.createNewUser": kitex.NewMethodInfo(
		createNewUserHandler,
		newCreateNewUserArgs,
		newCreateNewUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.deleteUser": kitex.NewMethodInfo(
		deleteUserHandler,
		newDeleteUserArgs,
		newDeleteUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.blockPeer": kitex.NewMethodInfo(
		blockPeerHandler,
		newBlockPeerArgs,
		newBlockPeerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.unBlockPeer": kitex.NewMethodInfo(
		unBlockPeerHandler,
		newUnBlockPeerArgs,
		newUnBlockPeerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.blockedByUser": kitex.NewMethodInfo(
		blockedByUserHandler,
		newBlockedByUserArgs,
		newBlockedByUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.isBlockedByUser": kitex.NewMethodInfo(
		isBlockedByUserHandler,
		newIsBlockedByUserArgs,
		newIsBlockedByUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.checkBlockUserList": kitex.NewMethodInfo(
		checkBlockUserListHandler,
		newCheckBlockUserListArgs,
		newCheckBlockUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getBlockedList": kitex.NewMethodInfo(
		getBlockedListHandler,
		newGetBlockedListArgs,
		newGetBlockedListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getContactSignUpNotification": kitex.NewMethodInfo(
		getContactSignUpNotificationHandler,
		newGetContactSignUpNotificationArgs,
		newGetContactSignUpNotificationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setContactSignUpNotification": kitex.NewMethodInfo(
		setContactSignUpNotificationHandler,
		newSetContactSignUpNotificationArgs,
		newSetContactSignUpNotificationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getContentSettings": kitex.NewMethodInfo(
		getContentSettingsHandler,
		newGetContentSettingsArgs,
		newGetContentSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setContentSettings": kitex.NewMethodInfo(
		setContentSettingsHandler,
		newSetContentSettingsArgs,
		newSetContentSettingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.deleteContact": kitex.NewMethodInfo(
		deleteContactHandler,
		newDeleteContactArgs,
		newDeleteContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getContactList": kitex.NewMethodInfo(
		getContactListHandler,
		newGetContactListArgs,
		newGetContactListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getContactIdList": kitex.NewMethodInfo(
		getContactIdListHandler,
		newGetContactIdListArgs,
		newGetContactIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getContact": kitex.NewMethodInfo(
		getContactHandler,
		newGetContactArgs,
		newGetContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.addContact": kitex.NewMethodInfo(
		addContactHandler,
		newAddContactArgs,
		newAddContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.checkContact": kitex.NewMethodInfo(
		checkContactHandler,
		newCheckContactArgs,
		newCheckContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getImportersByPhone": kitex.NewMethodInfo(
		getImportersByPhoneHandler,
		newGetImportersByPhoneArgs,
		newGetImportersByPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.deleteImportersByPhone": kitex.NewMethodInfo(
		deleteImportersByPhoneHandler,
		newDeleteImportersByPhoneArgs,
		newDeleteImportersByPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.importContacts": kitex.NewMethodInfo(
		importContactsHandler,
		newImportContactsArgs,
		newImportContactsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getCountryCode": kitex.NewMethodInfo(
		getCountryCodeHandler,
		newGetCountryCodeArgs,
		newGetCountryCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateAbout": kitex.NewMethodInfo(
		updateAboutHandler,
		newUpdateAboutArgs,
		newUpdateAboutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateFirstAndLastName": kitex.NewMethodInfo(
		updateFirstAndLastNameHandler,
		newUpdateFirstAndLastNameArgs,
		newUpdateFirstAndLastNameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateVerified": kitex.NewMethodInfo(
		updateVerifiedHandler,
		newUpdateVerifiedArgs,
		newUpdateVerifiedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateUsername": kitex.NewMethodInfo(
		updateUsernameHandler,
		newUpdateUsernameArgs,
		newUpdateUsernameResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateProfilePhoto": kitex.NewMethodInfo(
		updateProfilePhotoHandler,
		newUpdateProfilePhotoArgs,
		newUpdateProfilePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.deleteProfilePhotos": kitex.NewMethodInfo(
		deleteProfilePhotosHandler,
		newDeleteProfilePhotosArgs,
		newDeleteProfilePhotosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getProfilePhotos": kitex.NewMethodInfo(
		getProfilePhotosHandler,
		newGetProfilePhotosArgs,
		newGetProfilePhotosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setBotCommands": kitex.NewMethodInfo(
		setBotCommandsHandler,
		newSetBotCommandsArgs,
		newSetBotCommandsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.isBot": kitex.NewMethodInfo(
		isBotHandler,
		newIsBotArgs,
		newIsBotResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getBotInfo": kitex.NewMethodInfo(
		getBotInfoHandler,
		newGetBotInfoArgs,
		newGetBotInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.checkBots": kitex.NewMethodInfo(
		checkBotsHandler,
		newCheckBotsArgs,
		newCheckBotsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getFullUser": kitex.NewMethodInfo(
		getFullUserHandler,
		newGetFullUserArgs,
		newGetFullUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateEmojiStatus": kitex.NewMethodInfo(
		updateEmojiStatusHandler,
		newUpdateEmojiStatusArgs,
		newUpdateEmojiStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getUserDataById": kitex.NewMethodInfo(
		getUserDataByIdHandler,
		newGetUserDataByIdArgs,
		newGetUserDataByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getUserDataListByIdList": kitex.NewMethodInfo(
		getUserDataListByIdListHandler,
		newGetUserDataListByIdListArgs,
		newGetUserDataListByIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getUserDataByToken": kitex.NewMethodInfo(
		getUserDataByTokenHandler,
		newGetUserDataByTokenArgs,
		newGetUserDataByTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.search": kitex.NewMethodInfo(
		searchHandler,
		newSearchArgs,
		newSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateBotData": kitex.NewMethodInfo(
		updateBotDataHandler,
		newUpdateBotDataArgs,
		newUpdateBotDataResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getImmutableUserV2": kitex.NewMethodInfo(
		getImmutableUserV2Handler,
		newGetImmutableUserV2Args,
		newGetImmutableUserV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getMutableUsersV2": kitex.NewMethodInfo(
		getMutableUsersV2Handler,
		newGetMutableUsersV2Args,
		newGetMutableUsersV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.createNewTestUser": kitex.NewMethodInfo(
		createNewTestUserHandler,
		newCreateNewTestUserArgs,
		newCreateNewTestUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.editCloseFriends": kitex.NewMethodInfo(
		editCloseFriendsHandler,
		newEditCloseFriendsArgs,
		newEditCloseFriendsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setStoriesMaxId": kitex.NewMethodInfo(
		setStoriesMaxIdHandler,
		newSetStoriesMaxIdArgs,
		newSetStoriesMaxIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setColor": kitex.NewMethodInfo(
		setColorHandler,
		newSetColorArgs,
		newSetColorResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updateBirthday": kitex.NewMethodInfo(
		updateBirthdayHandler,
		newUpdateBirthdayArgs,
		newUpdateBirthdayResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getBirthdays": kitex.NewMethodInfo(
		getBirthdaysHandler,
		newGetBirthdaysArgs,
		newGetBirthdaysResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setStoriesHidden": kitex.NewMethodInfo(
		setStoriesHiddenHandler,
		newSetStoriesHiddenArgs,
		newSetStoriesHiddenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updatePersonalChannel": kitex.NewMethodInfo(
		updatePersonalChannelHandler,
		newUpdatePersonalChannelArgs,
		newUpdatePersonalChannelResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getUserIdByPhone": kitex.NewMethodInfo(
		getUserIdByPhoneHandler,
		newGetUserIdByPhoneArgs,
		newGetUserIdByPhoneResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.setAuthorizationTTL": kitex.NewMethodInfo(
		setAuthorizationTTLHandler,
		newSetAuthorizationTTLArgs,
		newSetAuthorizationTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getAuthorizationTTL": kitex.NewMethodInfo(
		getAuthorizationTTLHandler,
		newGetAuthorizationTTLArgs,
		newGetAuthorizationTTLResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.updatePremium": kitex.NewMethodInfo(
		updatePremiumHandler,
		newUpdatePremiumArgs,
		newUpdatePremiumResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"/user.RPCUser/user.getBotInfoV2": kitex.NewMethodInfo(
		getBotInfoV2Handler,
		newGetBotInfoV2Args,
		newGetBotInfoV2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

func init() {
	iface.RegisterKitexServiceInfo("RPCUser", userServiceServiceInfo)
	iface.RegisterKitexServiceInfoForClient("RPCUser", userServiceServiceInfoForClient)
	iface.RegisterKitexServiceInfoForStreamClient("RPCUser", userServiceServiceInfoForStreamClient)
}

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}

// NewServiceInfoForStreamClient creates a new ServiceInfo containing streaming methods
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCUser"
	handlerType := (*user.RPCUser)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func getLastSeensHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLastSeensArgs)
	realResult := result.(*GetLastSeensResult)
	success, err := handler.(user.RPCUser).UserGetLastSeens(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLastSeensArgs() interface{} {
	return &GetLastSeensArgs{}
}

func newGetLastSeensResult() interface{} {
	return &GetLastSeensResult{}
}

type GetLastSeensArgs struct {
	Req *user.TLUserGetLastSeens
}

func (p *GetLastSeensArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLastSeensArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLastSeensArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetLastSeens)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLastSeensArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLastSeensArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLastSeensArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetLastSeens)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLastSeensArgs_Req_DEFAULT *user.TLUserGetLastSeens

func (p *GetLastSeensArgs) GetReq() *user.TLUserGetLastSeens {
	if !p.IsSetReq() {
		return GetLastSeensArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLastSeensArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLastSeensResult struct {
	Success *user.VectorLastSeenData
}

var GetLastSeensResult_Success_DEFAULT *user.VectorLastSeenData

func (p *GetLastSeensResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLastSeensResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLastSeensResult) Unmarshal(in []byte) error {
	msg := new(user.VectorLastSeenData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastSeensResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLastSeensResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLastSeensResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorLastSeenData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastSeensResult) GetSuccess() *user.VectorLastSeenData {
	if !p.IsSetSuccess() {
		return GetLastSeensResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLastSeensResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorLastSeenData)
}

func (p *GetLastSeensResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLastSeensResult) GetResult() interface{} {
	return p.Success
}

func updateLastSeenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateLastSeenArgs)
	realResult := result.(*UpdateLastSeenResult)
	success, err := handler.(user.RPCUser).UserUpdateLastSeen(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateLastSeenArgs() interface{} {
	return &UpdateLastSeenArgs{}
}

func newUpdateLastSeenResult() interface{} {
	return &UpdateLastSeenResult{}
}

type UpdateLastSeenArgs struct {
	Req *user.TLUserUpdateLastSeen
}

func (p *UpdateLastSeenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateLastSeenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateLastSeenArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateLastSeen)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateLastSeenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateLastSeenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateLastSeenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateLastSeen)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateLastSeenArgs_Req_DEFAULT *user.TLUserUpdateLastSeen

func (p *UpdateLastSeenArgs) GetReq() *user.TLUserUpdateLastSeen {
	if !p.IsSetReq() {
		return UpdateLastSeenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateLastSeenArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateLastSeenResult struct {
	Success *tg.Bool
}

var UpdateLastSeenResult_Success_DEFAULT *tg.Bool

func (p *UpdateLastSeenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateLastSeenResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateLastSeenResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateLastSeenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateLastSeenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateLastSeenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateLastSeenResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateLastSeenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateLastSeenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateLastSeenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateLastSeenResult) GetResult() interface{} {
	return p.Success
}

func getLastSeenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetLastSeenArgs)
	realResult := result.(*GetLastSeenResult)
	success, err := handler.(user.RPCUser).UserGetLastSeen(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetLastSeenArgs() interface{} {
	return &GetLastSeenArgs{}
}

func newGetLastSeenResult() interface{} {
	return &GetLastSeenResult{}
}

type GetLastSeenArgs struct {
	Req *user.TLUserGetLastSeen
}

func (p *GetLastSeenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetLastSeenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetLastSeenArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetLastSeen)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetLastSeenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetLastSeenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetLastSeenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetLastSeen)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetLastSeenArgs_Req_DEFAULT *user.TLUserGetLastSeen

func (p *GetLastSeenArgs) GetReq() *user.TLUserGetLastSeen {
	if !p.IsSetReq() {
		return GetLastSeenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLastSeenArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetLastSeenResult struct {
	Success *user.LastSeenData
}

var GetLastSeenResult_Success_DEFAULT *user.LastSeenData

func (p *GetLastSeenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetLastSeenResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetLastSeenResult) Unmarshal(in []byte) error {
	msg := new(user.LastSeenData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastSeenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetLastSeenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetLastSeenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.LastSeenData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLastSeenResult) GetSuccess() *user.LastSeenData {
	if !p.IsSetSuccess() {
		return GetLastSeenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLastSeenResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LastSeenData)
}

func (p *GetLastSeenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLastSeenResult) GetResult() interface{} {
	return p.Success
}

func getImmutableUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetImmutableUserArgs)
	realResult := result.(*GetImmutableUserResult)
	success, err := handler.(user.RPCUser).UserGetImmutableUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetImmutableUserArgs() interface{} {
	return &GetImmutableUserArgs{}
}

func newGetImmutableUserResult() interface{} {
	return &GetImmutableUserResult{}
}

type GetImmutableUserArgs struct {
	Req *user.TLUserGetImmutableUser
}

func (p *GetImmutableUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetImmutableUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetImmutableUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetImmutableUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetImmutableUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetImmutableUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetImmutableUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetImmutableUserArgs_Req_DEFAULT *user.TLUserGetImmutableUser

func (p *GetImmutableUserArgs) GetReq() *user.TLUserGetImmutableUser {
	if !p.IsSetReq() {
		return GetImmutableUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImmutableUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetImmutableUserResult struct {
	Success *tg.ImmutableUser
}

var GetImmutableUserResult_Success_DEFAULT *tg.ImmutableUser

func (p *GetImmutableUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetImmutableUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetImmutableUserResult) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetImmutableUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetImmutableUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserResult) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return GetImmutableUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImmutableUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *GetImmutableUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImmutableUserResult) GetResult() interface{} {
	return p.Success
}

func getMutableUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMutableUsersArgs)
	realResult := result.(*GetMutableUsersResult)
	success, err := handler.(user.RPCUser).UserGetMutableUsers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMutableUsersArgs() interface{} {
	return &GetMutableUsersArgs{}
}

func newGetMutableUsersResult() interface{} {
	return &GetMutableUsersResult{}
}

type GetMutableUsersArgs struct {
	Req *user.TLUserGetMutableUsers
}

func (p *GetMutableUsersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMutableUsersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetMutableUsersArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetMutableUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMutableUsersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMutableUsersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMutableUsersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetMutableUsers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMutableUsersArgs_Req_DEFAULT *user.TLUserGetMutableUsers

func (p *GetMutableUsersArgs) GetReq() *user.TLUserGetMutableUsers {
	if !p.IsSetReq() {
		return GetMutableUsersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMutableUsersArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetMutableUsersResult struct {
	Success *user.VectorImmutableUser
}

var GetMutableUsersResult_Success_DEFAULT *user.VectorImmutableUser

func (p *GetMutableUsersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMutableUsersResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetMutableUsersResult) Unmarshal(in []byte) error {
	msg := new(user.VectorImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableUsersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMutableUsersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMutableUsersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableUsersResult) GetSuccess() *user.VectorImmutableUser {
	if !p.IsSetSuccess() {
		return GetMutableUsersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMutableUsersResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorImmutableUser)
}

func (p *GetMutableUsersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMutableUsersResult) GetResult() interface{} {
	return p.Success
}

func getImmutableUserByPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetImmutableUserByPhoneArgs)
	realResult := result.(*GetImmutableUserByPhoneResult)
	success, err := handler.(user.RPCUser).UserGetImmutableUserByPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetImmutableUserByPhoneArgs() interface{} {
	return &GetImmutableUserByPhoneArgs{}
}

func newGetImmutableUserByPhoneResult() interface{} {
	return &GetImmutableUserByPhoneResult{}
}

type GetImmutableUserByPhoneArgs struct {
	Req *user.TLUserGetImmutableUserByPhone
}

func (p *GetImmutableUserByPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetImmutableUserByPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetImmutableUserByPhoneArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetImmutableUserByPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetImmutableUserByPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetImmutableUserByPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetImmutableUserByPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetImmutableUserByPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetImmutableUserByPhoneArgs_Req_DEFAULT *user.TLUserGetImmutableUserByPhone

func (p *GetImmutableUserByPhoneArgs) GetReq() *user.TLUserGetImmutableUserByPhone {
	if !p.IsSetReq() {
		return GetImmutableUserByPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImmutableUserByPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetImmutableUserByPhoneResult struct {
	Success *tg.ImmutableUser
}

var GetImmutableUserByPhoneResult_Success_DEFAULT *tg.ImmutableUser

func (p *GetImmutableUserByPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetImmutableUserByPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetImmutableUserByPhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserByPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetImmutableUserByPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetImmutableUserByPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserByPhoneResult) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return GetImmutableUserByPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImmutableUserByPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *GetImmutableUserByPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImmutableUserByPhoneResult) GetResult() interface{} {
	return p.Success
}

func getImmutableUserByTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetImmutableUserByTokenArgs)
	realResult := result.(*GetImmutableUserByTokenResult)
	success, err := handler.(user.RPCUser).UserGetImmutableUserByToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetImmutableUserByTokenArgs() interface{} {
	return &GetImmutableUserByTokenArgs{}
}

func newGetImmutableUserByTokenResult() interface{} {
	return &GetImmutableUserByTokenResult{}
}

type GetImmutableUserByTokenArgs struct {
	Req *user.TLUserGetImmutableUserByToken
}

func (p *GetImmutableUserByTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetImmutableUserByTokenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetImmutableUserByTokenArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetImmutableUserByToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetImmutableUserByTokenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetImmutableUserByTokenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetImmutableUserByTokenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetImmutableUserByToken)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetImmutableUserByTokenArgs_Req_DEFAULT *user.TLUserGetImmutableUserByToken

func (p *GetImmutableUserByTokenArgs) GetReq() *user.TLUserGetImmutableUserByToken {
	if !p.IsSetReq() {
		return GetImmutableUserByTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImmutableUserByTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetImmutableUserByTokenResult struct {
	Success *tg.ImmutableUser
}

var GetImmutableUserByTokenResult_Success_DEFAULT *tg.ImmutableUser

func (p *GetImmutableUserByTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetImmutableUserByTokenResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetImmutableUserByTokenResult) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserByTokenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetImmutableUserByTokenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetImmutableUserByTokenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserByTokenResult) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return GetImmutableUserByTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImmutableUserByTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *GetImmutableUserByTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImmutableUserByTokenResult) GetResult() interface{} {
	return p.Success
}

func setAccountDaysTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetAccountDaysTTLArgs)
	realResult := result.(*SetAccountDaysTTLResult)
	success, err := handler.(user.RPCUser).UserSetAccountDaysTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetAccountDaysTTLArgs() interface{} {
	return &SetAccountDaysTTLArgs{}
}

func newSetAccountDaysTTLResult() interface{} {
	return &SetAccountDaysTTLResult{}
}

type SetAccountDaysTTLArgs struct {
	Req *user.TLUserSetAccountDaysTTL
}

func (p *SetAccountDaysTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetAccountDaysTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetAccountDaysTTLArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetAccountDaysTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetAccountDaysTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetAccountDaysTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetAccountDaysTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetAccountDaysTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetAccountDaysTTLArgs_Req_DEFAULT *user.TLUserSetAccountDaysTTL

func (p *SetAccountDaysTTLArgs) GetReq() *user.TLUserSetAccountDaysTTL {
	if !p.IsSetReq() {
		return SetAccountDaysTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetAccountDaysTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetAccountDaysTTLResult struct {
	Success *tg.Bool
}

var SetAccountDaysTTLResult_Success_DEFAULT *tg.Bool

func (p *SetAccountDaysTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetAccountDaysTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetAccountDaysTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAccountDaysTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetAccountDaysTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetAccountDaysTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAccountDaysTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetAccountDaysTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetAccountDaysTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetAccountDaysTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetAccountDaysTTLResult) GetResult() interface{} {
	return p.Success
}

func getAccountDaysTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAccountDaysTTLArgs)
	realResult := result.(*GetAccountDaysTTLResult)
	success, err := handler.(user.RPCUser).UserGetAccountDaysTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAccountDaysTTLArgs() interface{} {
	return &GetAccountDaysTTLArgs{}
}

func newGetAccountDaysTTLResult() interface{} {
	return &GetAccountDaysTTLResult{}
}

type GetAccountDaysTTLArgs struct {
	Req *user.TLUserGetAccountDaysTTL
}

func (p *GetAccountDaysTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAccountDaysTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAccountDaysTTLArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetAccountDaysTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAccountDaysTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAccountDaysTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAccountDaysTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetAccountDaysTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAccountDaysTTLArgs_Req_DEFAULT *user.TLUserGetAccountDaysTTL

func (p *GetAccountDaysTTLArgs) GetReq() *user.TLUserGetAccountDaysTTL {
	if !p.IsSetReq() {
		return GetAccountDaysTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAccountDaysTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAccountDaysTTLResult struct {
	Success *tg.AccountDaysTTL
}

var GetAccountDaysTTLResult_Success_DEFAULT *tg.AccountDaysTTL

func (p *GetAccountDaysTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAccountDaysTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAccountDaysTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountDaysTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAccountDaysTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAccountDaysTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAccountDaysTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountDaysTTL)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAccountDaysTTLResult) GetSuccess() *tg.AccountDaysTTL {
	if !p.IsSetSuccess() {
		return GetAccountDaysTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAccountDaysTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountDaysTTL)
}

func (p *GetAccountDaysTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAccountDaysTTLResult) GetResult() interface{} {
	return p.Success
}

func getNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetNotifySettingsArgs)
	realResult := result.(*GetNotifySettingsResult)
	success, err := handler.(user.RPCUser).UserGetNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetNotifySettingsArgs() interface{} {
	return &GetNotifySettingsArgs{}
}

func newGetNotifySettingsResult() interface{} {
	return &GetNotifySettingsResult{}
}

type GetNotifySettingsArgs struct {
	Req *user.TLUserGetNotifySettings
}

func (p *GetNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetNotifySettingsArgs_Req_DEFAULT *user.TLUserGetNotifySettings

func (p *GetNotifySettingsArgs) GetReq() *user.TLUserGetNotifySettings {
	if !p.IsSetReq() {
		return GetNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetNotifySettingsResult struct {
	Success *tg.PeerNotifySettings
}

var GetNotifySettingsResult_Success_DEFAULT *tg.PeerNotifySettings

func (p *GetNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.PeerNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PeerNotifySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotifySettingsResult) GetSuccess() *tg.PeerNotifySettings {
	if !p.IsSetSuccess() {
		return GetNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PeerNotifySettings)
}

func (p *GetNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func getNotifySettingsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetNotifySettingsListArgs)
	realResult := result.(*GetNotifySettingsListResult)
	success, err := handler.(user.RPCUser).UserGetNotifySettingsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetNotifySettingsListArgs() interface{} {
	return &GetNotifySettingsListArgs{}
}

func newGetNotifySettingsListResult() interface{} {
	return &GetNotifySettingsListResult{}
}

type GetNotifySettingsListArgs struct {
	Req *user.TLUserGetNotifySettingsList
}

func (p *GetNotifySettingsListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetNotifySettingsListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetNotifySettingsListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetNotifySettingsList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetNotifySettingsListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetNotifySettingsListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetNotifySettingsListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetNotifySettingsList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetNotifySettingsListArgs_Req_DEFAULT *user.TLUserGetNotifySettingsList

func (p *GetNotifySettingsListArgs) GetReq() *user.TLUserGetNotifySettingsList {
	if !p.IsSetReq() {
		return GetNotifySettingsListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNotifySettingsListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetNotifySettingsListResult struct {
	Success *user.VectorPeerPeerNotifySettings
}

var GetNotifySettingsListResult_Success_DEFAULT *user.VectorPeerPeerNotifySettings

func (p *GetNotifySettingsListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetNotifySettingsListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetNotifySettingsListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorPeerPeerNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotifySettingsListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetNotifySettingsListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetNotifySettingsListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorPeerPeerNotifySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotifySettingsListResult) GetSuccess() *user.VectorPeerPeerNotifySettings {
	if !p.IsSetSuccess() {
		return GetNotifySettingsListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNotifySettingsListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorPeerPeerNotifySettings)
}

func (p *GetNotifySettingsListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNotifySettingsListResult) GetResult() interface{} {
	return p.Success
}

func setNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetNotifySettingsArgs)
	realResult := result.(*SetNotifySettingsResult)
	success, err := handler.(user.RPCUser).UserSetNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetNotifySettingsArgs() interface{} {
	return &SetNotifySettingsArgs{}
}

func newSetNotifySettingsResult() interface{} {
	return &SetNotifySettingsResult{}
}

type SetNotifySettingsArgs struct {
	Req *user.TLUserSetNotifySettings
}

func (p *SetNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetNotifySettingsArgs_Req_DEFAULT *user.TLUserSetNotifySettings

func (p *SetNotifySettingsArgs) GetReq() *user.TLUserSetNotifySettings {
	if !p.IsSetReq() {
		return SetNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetNotifySettingsResult struct {
	Success *tg.Bool
}

var SetNotifySettingsResult_Success_DEFAULT *tg.Bool

func (p *SetNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetNotifySettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func resetNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ResetNotifySettingsArgs)
	realResult := result.(*ResetNotifySettingsResult)
	success, err := handler.(user.RPCUser).UserResetNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newResetNotifySettingsArgs() interface{} {
	return &ResetNotifySettingsArgs{}
}

func newResetNotifySettingsResult() interface{} {
	return &ResetNotifySettingsResult{}
}

type ResetNotifySettingsArgs struct {
	Req *user.TLUserResetNotifySettings
}

func (p *ResetNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ResetNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ResetNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserResetNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ResetNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ResetNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ResetNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserResetNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ResetNotifySettingsArgs_Req_DEFAULT *user.TLUserResetNotifySettings

func (p *ResetNotifySettingsArgs) GetReq() *user.TLUserResetNotifySettings {
	if !p.IsSetReq() {
		return ResetNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ResetNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ResetNotifySettingsResult struct {
	Success *tg.Bool
}

var ResetNotifySettingsResult_Success_DEFAULT *tg.Bool

func (p *ResetNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ResetNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ResetNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResetNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ResetNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ResetNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResetNotifySettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ResetNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ResetNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ResetNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ResetNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func getAllNotifySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAllNotifySettingsArgs)
	realResult := result.(*GetAllNotifySettingsResult)
	success, err := handler.(user.RPCUser).UserGetAllNotifySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAllNotifySettingsArgs() interface{} {
	return &GetAllNotifySettingsArgs{}
}

func newGetAllNotifySettingsResult() interface{} {
	return &GetAllNotifySettingsResult{}
}

type GetAllNotifySettingsArgs struct {
	Req *user.TLUserGetAllNotifySettings
}

func (p *GetAllNotifySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAllNotifySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAllNotifySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetAllNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAllNotifySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAllNotifySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAllNotifySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetAllNotifySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAllNotifySettingsArgs_Req_DEFAULT *user.TLUserGetAllNotifySettings

func (p *GetAllNotifySettingsArgs) GetReq() *user.TLUserGetAllNotifySettings {
	if !p.IsSetReq() {
		return GetAllNotifySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAllNotifySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAllNotifySettingsResult struct {
	Success *user.VectorPeerPeerNotifySettings
}

var GetAllNotifySettingsResult_Success_DEFAULT *user.VectorPeerPeerNotifySettings

func (p *GetAllNotifySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAllNotifySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAllNotifySettingsResult) Unmarshal(in []byte) error {
	msg := new(user.VectorPeerPeerNotifySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllNotifySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAllNotifySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAllNotifySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorPeerPeerNotifySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAllNotifySettingsResult) GetSuccess() *user.VectorPeerPeerNotifySettings {
	if !p.IsSetSuccess() {
		return GetAllNotifySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAllNotifySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorPeerPeerNotifySettings)
}

func (p *GetAllNotifySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAllNotifySettingsResult) GetResult() interface{} {
	return p.Success
}

func getGlobalPrivacySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetGlobalPrivacySettingsArgs)
	realResult := result.(*GetGlobalPrivacySettingsResult)
	success, err := handler.(user.RPCUser).UserGetGlobalPrivacySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetGlobalPrivacySettingsArgs() interface{} {
	return &GetGlobalPrivacySettingsArgs{}
}

func newGetGlobalPrivacySettingsResult() interface{} {
	return &GetGlobalPrivacySettingsResult{}
}

type GetGlobalPrivacySettingsArgs struct {
	Req *user.TLUserGetGlobalPrivacySettings
}

func (p *GetGlobalPrivacySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetGlobalPrivacySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetGlobalPrivacySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetGlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetGlobalPrivacySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetGlobalPrivacySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetGlobalPrivacySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetGlobalPrivacySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetGlobalPrivacySettingsArgs_Req_DEFAULT *user.TLUserGetGlobalPrivacySettings

func (p *GetGlobalPrivacySettingsArgs) GetReq() *user.TLUserGetGlobalPrivacySettings {
	if !p.IsSetReq() {
		return GetGlobalPrivacySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetGlobalPrivacySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetGlobalPrivacySettingsResult struct {
	Success *tg.GlobalPrivacySettings
}

var GetGlobalPrivacySettingsResult_Success_DEFAULT *tg.GlobalPrivacySettings

func (p *GetGlobalPrivacySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetGlobalPrivacySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetGlobalPrivacySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.GlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetGlobalPrivacySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetGlobalPrivacySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetGlobalPrivacySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.GlobalPrivacySettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetGlobalPrivacySettingsResult) GetSuccess() *tg.GlobalPrivacySettings {
	if !p.IsSetSuccess() {
		return GetGlobalPrivacySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetGlobalPrivacySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.GlobalPrivacySettings)
}

func (p *GetGlobalPrivacySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetGlobalPrivacySettingsResult) GetResult() interface{} {
	return p.Success
}

func setGlobalPrivacySettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetGlobalPrivacySettingsArgs)
	realResult := result.(*SetGlobalPrivacySettingsResult)
	success, err := handler.(user.RPCUser).UserSetGlobalPrivacySettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetGlobalPrivacySettingsArgs() interface{} {
	return &SetGlobalPrivacySettingsArgs{}
}

func newSetGlobalPrivacySettingsResult() interface{} {
	return &SetGlobalPrivacySettingsResult{}
}

type SetGlobalPrivacySettingsArgs struct {
	Req *user.TLUserSetGlobalPrivacySettings
}

func (p *SetGlobalPrivacySettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetGlobalPrivacySettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetGlobalPrivacySettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetGlobalPrivacySettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetGlobalPrivacySettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetGlobalPrivacySettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetGlobalPrivacySettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetGlobalPrivacySettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetGlobalPrivacySettingsArgs_Req_DEFAULT *user.TLUserSetGlobalPrivacySettings

func (p *SetGlobalPrivacySettingsArgs) GetReq() *user.TLUserSetGlobalPrivacySettings {
	if !p.IsSetReq() {
		return SetGlobalPrivacySettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetGlobalPrivacySettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetGlobalPrivacySettingsResult struct {
	Success *tg.Bool
}

var SetGlobalPrivacySettingsResult_Success_DEFAULT *tg.Bool

func (p *SetGlobalPrivacySettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetGlobalPrivacySettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetGlobalPrivacySettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetGlobalPrivacySettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetGlobalPrivacySettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetGlobalPrivacySettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetGlobalPrivacySettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetGlobalPrivacySettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetGlobalPrivacySettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetGlobalPrivacySettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetGlobalPrivacySettingsResult) GetResult() interface{} {
	return p.Success
}

func getPrivacyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPrivacyArgs)
	realResult := result.(*GetPrivacyResult)
	success, err := handler.(user.RPCUser).UserGetPrivacy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPrivacyArgs() interface{} {
	return &GetPrivacyArgs{}
}

func newGetPrivacyResult() interface{} {
	return &GetPrivacyResult{}
}

type GetPrivacyArgs struct {
	Req *user.TLUserGetPrivacy
}

func (p *GetPrivacyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPrivacyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPrivacyArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetPrivacy)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPrivacyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPrivacyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPrivacyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetPrivacy)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPrivacyArgs_Req_DEFAULT *user.TLUserGetPrivacy

func (p *GetPrivacyArgs) GetReq() *user.TLUserGetPrivacy {
	if !p.IsSetReq() {
		return GetPrivacyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPrivacyArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPrivacyResult struct {
	Success *user.VectorPrivacyRule
}

var GetPrivacyResult_Success_DEFAULT *user.VectorPrivacyRule

func (p *GetPrivacyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPrivacyResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPrivacyResult) Unmarshal(in []byte) error {
	msg := new(user.VectorPrivacyRule)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPrivacyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPrivacyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPrivacyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorPrivacyRule)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPrivacyResult) GetSuccess() *user.VectorPrivacyRule {
	if !p.IsSetSuccess() {
		return GetPrivacyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPrivacyResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorPrivacyRule)
}

func (p *GetPrivacyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPrivacyResult) GetResult() interface{} {
	return p.Success
}

func setPrivacyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetPrivacyArgs)
	realResult := result.(*SetPrivacyResult)
	success, err := handler.(user.RPCUser).UserSetPrivacy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetPrivacyArgs() interface{} {
	return &SetPrivacyArgs{}
}

func newSetPrivacyResult() interface{} {
	return &SetPrivacyResult{}
}

type SetPrivacyArgs struct {
	Req *user.TLUserSetPrivacy
}

func (p *SetPrivacyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetPrivacyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetPrivacyArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetPrivacy)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetPrivacyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetPrivacyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetPrivacyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetPrivacy)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetPrivacyArgs_Req_DEFAULT *user.TLUserSetPrivacy

func (p *SetPrivacyArgs) GetReq() *user.TLUserSetPrivacy {
	if !p.IsSetReq() {
		return SetPrivacyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetPrivacyArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetPrivacyResult struct {
	Success *tg.Bool
}

var SetPrivacyResult_Success_DEFAULT *tg.Bool

func (p *SetPrivacyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetPrivacyResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetPrivacyResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPrivacyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetPrivacyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetPrivacyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPrivacyResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetPrivacyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetPrivacyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetPrivacyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetPrivacyResult) GetResult() interface{} {
	return p.Success
}

func checkPrivacyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckPrivacyArgs)
	realResult := result.(*CheckPrivacyResult)
	success, err := handler.(user.RPCUser).UserCheckPrivacy(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckPrivacyArgs() interface{} {
	return &CheckPrivacyArgs{}
}

func newCheckPrivacyResult() interface{} {
	return &CheckPrivacyResult{}
}

type CheckPrivacyArgs struct {
	Req *user.TLUserCheckPrivacy
}

func (p *CheckPrivacyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckPrivacyArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckPrivacyArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCheckPrivacy)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckPrivacyArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckPrivacyArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckPrivacyArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCheckPrivacy)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckPrivacyArgs_Req_DEFAULT *user.TLUserCheckPrivacy

func (p *CheckPrivacyArgs) GetReq() *user.TLUserCheckPrivacy {
	if !p.IsSetReq() {
		return CheckPrivacyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckPrivacyArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckPrivacyResult struct {
	Success *tg.Bool
}

var CheckPrivacyResult_Success_DEFAULT *tg.Bool

func (p *CheckPrivacyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckPrivacyResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckPrivacyResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckPrivacyResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckPrivacyResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckPrivacyResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckPrivacyResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return CheckPrivacyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckPrivacyResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *CheckPrivacyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckPrivacyResult) GetResult() interface{} {
	return p.Success
}

func addPeerSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AddPeerSettingsArgs)
	realResult := result.(*AddPeerSettingsResult)
	success, err := handler.(user.RPCUser).UserAddPeerSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAddPeerSettingsArgs() interface{} {
	return &AddPeerSettingsArgs{}
}

func newAddPeerSettingsResult() interface{} {
	return &AddPeerSettingsResult{}
}

type AddPeerSettingsArgs struct {
	Req *user.TLUserAddPeerSettings
}

func (p *AddPeerSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddPeerSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AddPeerSettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserAddPeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AddPeerSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AddPeerSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AddPeerSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserAddPeerSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AddPeerSettingsArgs_Req_DEFAULT *user.TLUserAddPeerSettings

func (p *AddPeerSettingsArgs) GetReq() *user.TLUserAddPeerSettings {
	if !p.IsSetReq() {
		return AddPeerSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddPeerSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddPeerSettingsResult struct {
	Success *tg.Bool
}

var AddPeerSettingsResult_Success_DEFAULT *tg.Bool

func (p *AddPeerSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddPeerSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *AddPeerSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddPeerSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AddPeerSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AddPeerSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddPeerSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AddPeerSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddPeerSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AddPeerSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddPeerSettingsResult) GetResult() interface{} {
	return p.Success
}

func getPeerSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetPeerSettingsArgs)
	realResult := result.(*GetPeerSettingsResult)
	success, err := handler.(user.RPCUser).UserGetPeerSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetPeerSettingsArgs() interface{} {
	return &GetPeerSettingsArgs{}
}

func newGetPeerSettingsResult() interface{} {
	return &GetPeerSettingsResult{}
}

type GetPeerSettingsArgs struct {
	Req *user.TLUserGetPeerSettings
}

func (p *GetPeerSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPeerSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetPeerSettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetPeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetPeerSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetPeerSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetPeerSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetPeerSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetPeerSettingsArgs_Req_DEFAULT *user.TLUserGetPeerSettings

func (p *GetPeerSettingsArgs) GetReq() *user.TLUserGetPeerSettings {
	if !p.IsSetReq() {
		return GetPeerSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPeerSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPeerSettingsResult struct {
	Success *tg.PeerSettings
}

var GetPeerSettingsResult_Success_DEFAULT *tg.PeerSettings

func (p *GetPeerSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPeerSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetPeerSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.PeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetPeerSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetPeerSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.PeerSettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPeerSettingsResult) GetSuccess() *tg.PeerSettings {
	if !p.IsSetSuccess() {
		return GetPeerSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPeerSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.PeerSettings)
}

func (p *GetPeerSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPeerSettingsResult) GetResult() interface{} {
	return p.Success
}

func deletePeerSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeletePeerSettingsArgs)
	realResult := result.(*DeletePeerSettingsResult)
	success, err := handler.(user.RPCUser).UserDeletePeerSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeletePeerSettingsArgs() interface{} {
	return &DeletePeerSettingsArgs{}
}

func newDeletePeerSettingsResult() interface{} {
	return &DeletePeerSettingsResult{}
}

type DeletePeerSettingsArgs struct {
	Req *user.TLUserDeletePeerSettings
}

func (p *DeletePeerSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePeerSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeletePeerSettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserDeletePeerSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeletePeerSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeletePeerSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeletePeerSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserDeletePeerSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeletePeerSettingsArgs_Req_DEFAULT *user.TLUserDeletePeerSettings

func (p *DeletePeerSettingsArgs) GetReq() *user.TLUserDeletePeerSettings {
	if !p.IsSetReq() {
		return DeletePeerSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePeerSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeletePeerSettingsResult struct {
	Success *tg.Bool
}

var DeletePeerSettingsResult_Success_DEFAULT *tg.Bool

func (p *DeletePeerSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePeerSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeletePeerSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePeerSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeletePeerSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeletePeerSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePeerSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeletePeerSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePeerSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeletePeerSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePeerSettingsResult) GetResult() interface{} {
	return p.Success
}

func changePhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ChangePhoneArgs)
	realResult := result.(*ChangePhoneResult)
	success, err := handler.(user.RPCUser).UserChangePhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newChangePhoneArgs() interface{} {
	return &ChangePhoneArgs{}
}

func newChangePhoneResult() interface{} {
	return &ChangePhoneResult{}
}

type ChangePhoneArgs struct {
	Req *user.TLUserChangePhone
}

func (p *ChangePhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangePhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ChangePhoneArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserChangePhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ChangePhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ChangePhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ChangePhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserChangePhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ChangePhoneArgs_Req_DEFAULT *user.TLUserChangePhone

func (p *ChangePhoneArgs) GetReq() *user.TLUserChangePhone {
	if !p.IsSetReq() {
		return ChangePhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangePhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChangePhoneResult struct {
	Success *tg.Bool
}

var ChangePhoneResult_Success_DEFAULT *tg.Bool

func (p *ChangePhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangePhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *ChangePhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangePhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ChangePhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ChangePhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangePhoneResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ChangePhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangePhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ChangePhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangePhoneResult) GetResult() interface{} {
	return p.Success
}

func createNewUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreateNewUserArgs)
	realResult := result.(*CreateNewUserResult)
	success, err := handler.(user.RPCUser).UserCreateNewUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreateNewUserArgs() interface{} {
	return &CreateNewUserArgs{}
}

func newCreateNewUserResult() interface{} {
	return &CreateNewUserResult{}
}

type CreateNewUserArgs struct {
	Req *user.TLUserCreateNewUser
}

func (p *CreateNewUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateNewUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CreateNewUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCreateNewUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreateNewUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreateNewUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreateNewUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCreateNewUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreateNewUserArgs_Req_DEFAULT *user.TLUserCreateNewUser

func (p *CreateNewUserArgs) GetReq() *user.TLUserCreateNewUser {
	if !p.IsSetReq() {
		return CreateNewUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateNewUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateNewUserResult struct {
	Success *tg.ImmutableUser
}

var CreateNewUserResult_Success_DEFAULT *tg.ImmutableUser

func (p *CreateNewUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateNewUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *CreateNewUserResult) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateNewUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreateNewUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreateNewUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateNewUserResult) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return CreateNewUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateNewUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *CreateNewUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateNewUserResult) GetResult() interface{} {
	return p.Success
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteUserArgs)
	realResult := result.(*DeleteUserResult)
	success, err := handler.(user.RPCUser).UserDeleteUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteUserArgs() interface{} {
	return &DeleteUserArgs{}
}

func newDeleteUserResult() interface{} {
	return &DeleteUserResult{}
}

type DeleteUserArgs struct {
	Req *user.TLUserDeleteUser
}

func (p *DeleteUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserDeleteUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserDeleteUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteUserArgs_Req_DEFAULT *user.TLUserDeleteUser

func (p *DeleteUserArgs) GetReq() *user.TLUserDeleteUser {
	if !p.IsSetReq() {
		return DeleteUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteUserResult struct {
	Success *tg.Bool
}

var DeleteUserResult_Success_DEFAULT *tg.Bool

func (p *DeleteUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteUserResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteUserResult) GetResult() interface{} {
	return p.Success
}

func blockPeerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*BlockPeerArgs)
	realResult := result.(*BlockPeerResult)
	success, err := handler.(user.RPCUser).UserBlockPeer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newBlockPeerArgs() interface{} {
	return &BlockPeerArgs{}
}

func newBlockPeerResult() interface{} {
	return &BlockPeerResult{}
}

type BlockPeerArgs struct {
	Req *user.TLUserBlockPeer
}

func (p *BlockPeerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BlockPeerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *BlockPeerArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserBlockPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *BlockPeerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in BlockPeerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *BlockPeerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserBlockPeer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var BlockPeerArgs_Req_DEFAULT *user.TLUserBlockPeer

func (p *BlockPeerArgs) GetReq() *user.TLUserBlockPeer {
	if !p.IsSetReq() {
		return BlockPeerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BlockPeerArgs) IsSetReq() bool {
	return p.Req != nil
}

type BlockPeerResult struct {
	Success *tg.Bool
}

var BlockPeerResult_Success_DEFAULT *tg.Bool

func (p *BlockPeerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BlockPeerResult")
	}
	return json.Marshal(p.Success)
}

func (p *BlockPeerResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BlockPeerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in BlockPeerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *BlockPeerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BlockPeerResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return BlockPeerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BlockPeerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *BlockPeerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BlockPeerResult) GetResult() interface{} {
	return p.Success
}

func unBlockPeerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UnBlockPeerArgs)
	realResult := result.(*UnBlockPeerResult)
	success, err := handler.(user.RPCUser).UserUnBlockPeer(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUnBlockPeerArgs() interface{} {
	return &UnBlockPeerArgs{}
}

func newUnBlockPeerResult() interface{} {
	return &UnBlockPeerResult{}
}

type UnBlockPeerArgs struct {
	Req *user.TLUserUnBlockPeer
}

func (p *UnBlockPeerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UnBlockPeerArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UnBlockPeerArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUnBlockPeer)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UnBlockPeerArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UnBlockPeerArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UnBlockPeerArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUnBlockPeer)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UnBlockPeerArgs_Req_DEFAULT *user.TLUserUnBlockPeer

func (p *UnBlockPeerArgs) GetReq() *user.TLUserUnBlockPeer {
	if !p.IsSetReq() {
		return UnBlockPeerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UnBlockPeerArgs) IsSetReq() bool {
	return p.Req != nil
}

type UnBlockPeerResult struct {
	Success *tg.Bool
}

var UnBlockPeerResult_Success_DEFAULT *tg.Bool

func (p *UnBlockPeerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UnBlockPeerResult")
	}
	return json.Marshal(p.Success)
}

func (p *UnBlockPeerResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnBlockPeerResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UnBlockPeerResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UnBlockPeerResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UnBlockPeerResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UnBlockPeerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UnBlockPeerResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UnBlockPeerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UnBlockPeerResult) GetResult() interface{} {
	return p.Success
}

func blockedByUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*BlockedByUserArgs)
	realResult := result.(*BlockedByUserResult)
	success, err := handler.(user.RPCUser).UserBlockedByUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newBlockedByUserArgs() interface{} {
	return &BlockedByUserArgs{}
}

func newBlockedByUserResult() interface{} {
	return &BlockedByUserResult{}
}

type BlockedByUserArgs struct {
	Req *user.TLUserBlockedByUser
}

func (p *BlockedByUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in BlockedByUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *BlockedByUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserBlockedByUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *BlockedByUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in BlockedByUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *BlockedByUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserBlockedByUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var BlockedByUserArgs_Req_DEFAULT *user.TLUserBlockedByUser

func (p *BlockedByUserArgs) GetReq() *user.TLUserBlockedByUser {
	if !p.IsSetReq() {
		return BlockedByUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BlockedByUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type BlockedByUserResult struct {
	Success *tg.Bool
}

var BlockedByUserResult_Success_DEFAULT *tg.Bool

func (p *BlockedByUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in BlockedByUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *BlockedByUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BlockedByUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in BlockedByUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *BlockedByUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BlockedByUserResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return BlockedByUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BlockedByUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *BlockedByUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BlockedByUserResult) GetResult() interface{} {
	return p.Success
}

func isBlockedByUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*IsBlockedByUserArgs)
	realResult := result.(*IsBlockedByUserResult)
	success, err := handler.(user.RPCUser).UserIsBlockedByUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newIsBlockedByUserArgs() interface{} {
	return &IsBlockedByUserArgs{}
}

func newIsBlockedByUserResult() interface{} {
	return &IsBlockedByUserResult{}
}

type IsBlockedByUserArgs struct {
	Req *user.TLUserIsBlockedByUser
}

func (p *IsBlockedByUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsBlockedByUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *IsBlockedByUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserIsBlockedByUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *IsBlockedByUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in IsBlockedByUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *IsBlockedByUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserIsBlockedByUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var IsBlockedByUserArgs_Req_DEFAULT *user.TLUserIsBlockedByUser

func (p *IsBlockedByUserArgs) GetReq() *user.TLUserIsBlockedByUser {
	if !p.IsSetReq() {
		return IsBlockedByUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsBlockedByUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type IsBlockedByUserResult struct {
	Success *tg.Bool
}

var IsBlockedByUserResult_Success_DEFAULT *tg.Bool

func (p *IsBlockedByUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsBlockedByUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *IsBlockedByUserResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsBlockedByUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in IsBlockedByUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *IsBlockedByUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsBlockedByUserResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return IsBlockedByUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsBlockedByUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *IsBlockedByUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsBlockedByUserResult) GetResult() interface{} {
	return p.Success
}

func checkBlockUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckBlockUserListArgs)
	realResult := result.(*CheckBlockUserListResult)
	success, err := handler.(user.RPCUser).UserCheckBlockUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckBlockUserListArgs() interface{} {
	return &CheckBlockUserListArgs{}
}

func newCheckBlockUserListResult() interface{} {
	return &CheckBlockUserListResult{}
}

type CheckBlockUserListArgs struct {
	Req *user.TLUserCheckBlockUserList
}

func (p *CheckBlockUserListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckBlockUserListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckBlockUserListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCheckBlockUserList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckBlockUserListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckBlockUserListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckBlockUserListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCheckBlockUserList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckBlockUserListArgs_Req_DEFAULT *user.TLUserCheckBlockUserList

func (p *CheckBlockUserListArgs) GetReq() *user.TLUserCheckBlockUserList {
	if !p.IsSetReq() {
		return CheckBlockUserListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckBlockUserListArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckBlockUserListResult struct {
	Success *user.VectorLong
}

var CheckBlockUserListResult_Success_DEFAULT *user.VectorLong

func (p *CheckBlockUserListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckBlockUserListResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckBlockUserListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckBlockUserListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckBlockUserListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckBlockUserListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckBlockUserListResult) GetSuccess() *user.VectorLong {
	if !p.IsSetSuccess() {
		return CheckBlockUserListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckBlockUserListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorLong)
}

func (p *CheckBlockUserListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckBlockUserListResult) GetResult() interface{} {
	return p.Success
}

func getBlockedListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetBlockedListArgs)
	realResult := result.(*GetBlockedListResult)
	success, err := handler.(user.RPCUser).UserGetBlockedList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetBlockedListArgs() interface{} {
	return &GetBlockedListArgs{}
}

func newGetBlockedListResult() interface{} {
	return &GetBlockedListResult{}
}

type GetBlockedListArgs struct {
	Req *user.TLUserGetBlockedList
}

func (p *GetBlockedListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBlockedListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetBlockedListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetBlockedList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetBlockedListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetBlockedListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetBlockedListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetBlockedList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetBlockedListArgs_Req_DEFAULT *user.TLUserGetBlockedList

func (p *GetBlockedListArgs) GetReq() *user.TLUserGetBlockedList {
	if !p.IsSetReq() {
		return GetBlockedListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBlockedListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetBlockedListResult struct {
	Success *user.VectorPeerBlocked
}

var GetBlockedListResult_Success_DEFAULT *user.VectorPeerBlocked

func (p *GetBlockedListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBlockedListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetBlockedListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorPeerBlocked)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBlockedListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetBlockedListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetBlockedListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorPeerBlocked)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBlockedListResult) GetSuccess() *user.VectorPeerBlocked {
	if !p.IsSetSuccess() {
		return GetBlockedListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBlockedListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorPeerBlocked)
}

func (p *GetBlockedListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetBlockedListResult) GetResult() interface{} {
	return p.Success
}

func getContactSignUpNotificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetContactSignUpNotificationArgs)
	realResult := result.(*GetContactSignUpNotificationResult)
	success, err := handler.(user.RPCUser).UserGetContactSignUpNotification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetContactSignUpNotificationArgs() interface{} {
	return &GetContactSignUpNotificationArgs{}
}

func newGetContactSignUpNotificationResult() interface{} {
	return &GetContactSignUpNotificationResult{}
}

type GetContactSignUpNotificationArgs struct {
	Req *user.TLUserGetContactSignUpNotification
}

func (p *GetContactSignUpNotificationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetContactSignUpNotificationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetContactSignUpNotificationArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetContactSignUpNotification)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetContactSignUpNotificationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetContactSignUpNotificationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetContactSignUpNotificationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetContactSignUpNotification)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetContactSignUpNotificationArgs_Req_DEFAULT *user.TLUserGetContactSignUpNotification

func (p *GetContactSignUpNotificationArgs) GetReq() *user.TLUserGetContactSignUpNotification {
	if !p.IsSetReq() {
		return GetContactSignUpNotificationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContactSignUpNotificationArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetContactSignUpNotificationResult struct {
	Success *tg.Bool
}

var GetContactSignUpNotificationResult_Success_DEFAULT *tg.Bool

func (p *GetContactSignUpNotificationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetContactSignUpNotificationResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetContactSignUpNotificationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactSignUpNotificationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetContactSignUpNotificationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetContactSignUpNotificationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactSignUpNotificationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return GetContactSignUpNotificationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContactSignUpNotificationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *GetContactSignUpNotificationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContactSignUpNotificationResult) GetResult() interface{} {
	return p.Success
}

func setContactSignUpNotificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetContactSignUpNotificationArgs)
	realResult := result.(*SetContactSignUpNotificationResult)
	success, err := handler.(user.RPCUser).UserSetContactSignUpNotification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetContactSignUpNotificationArgs() interface{} {
	return &SetContactSignUpNotificationArgs{}
}

func newSetContactSignUpNotificationResult() interface{} {
	return &SetContactSignUpNotificationResult{}
}

type SetContactSignUpNotificationArgs struct {
	Req *user.TLUserSetContactSignUpNotification
}

func (p *SetContactSignUpNotificationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetContactSignUpNotificationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetContactSignUpNotificationArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetContactSignUpNotification)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetContactSignUpNotificationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetContactSignUpNotificationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetContactSignUpNotificationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetContactSignUpNotification)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetContactSignUpNotificationArgs_Req_DEFAULT *user.TLUserSetContactSignUpNotification

func (p *SetContactSignUpNotificationArgs) GetReq() *user.TLUserSetContactSignUpNotification {
	if !p.IsSetReq() {
		return SetContactSignUpNotificationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetContactSignUpNotificationArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetContactSignUpNotificationResult struct {
	Success *tg.Bool
}

var SetContactSignUpNotificationResult_Success_DEFAULT *tg.Bool

func (p *SetContactSignUpNotificationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetContactSignUpNotificationResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetContactSignUpNotificationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetContactSignUpNotificationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetContactSignUpNotificationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetContactSignUpNotificationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetContactSignUpNotificationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetContactSignUpNotificationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetContactSignUpNotificationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetContactSignUpNotificationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetContactSignUpNotificationResult) GetResult() interface{} {
	return p.Success
}

func getContentSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetContentSettingsArgs)
	realResult := result.(*GetContentSettingsResult)
	success, err := handler.(user.RPCUser).UserGetContentSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetContentSettingsArgs() interface{} {
	return &GetContentSettingsArgs{}
}

func newGetContentSettingsResult() interface{} {
	return &GetContentSettingsResult{}
}

type GetContentSettingsArgs struct {
	Req *user.TLUserGetContentSettings
}

func (p *GetContentSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetContentSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetContentSettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetContentSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetContentSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetContentSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetContentSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetContentSettingsArgs_Req_DEFAULT *user.TLUserGetContentSettings

func (p *GetContentSettingsArgs) GetReq() *user.TLUserGetContentSettings {
	if !p.IsSetReq() {
		return GetContentSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContentSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetContentSettingsResult struct {
	Success *tg.AccountContentSettings
}

var GetContentSettingsResult_Success_DEFAULT *tg.AccountContentSettings

func (p *GetContentSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetContentSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetContentSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContentSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetContentSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetContentSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountContentSettings)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContentSettingsResult) GetSuccess() *tg.AccountContentSettings {
	if !p.IsSetSuccess() {
		return GetContentSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContentSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountContentSettings)
}

func (p *GetContentSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContentSettingsResult) GetResult() interface{} {
	return p.Success
}

func setContentSettingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetContentSettingsArgs)
	realResult := result.(*SetContentSettingsResult)
	success, err := handler.(user.RPCUser).UserSetContentSettings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetContentSettingsArgs() interface{} {
	return &SetContentSettingsArgs{}
}

func newSetContentSettingsResult() interface{} {
	return &SetContentSettingsResult{}
}

type SetContentSettingsArgs struct {
	Req *user.TLUserSetContentSettings
}

func (p *SetContentSettingsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetContentSettingsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetContentSettingsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetContentSettings)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetContentSettingsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetContentSettingsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetContentSettingsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetContentSettings)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetContentSettingsArgs_Req_DEFAULT *user.TLUserSetContentSettings

func (p *SetContentSettingsArgs) GetReq() *user.TLUserSetContentSettings {
	if !p.IsSetReq() {
		return SetContentSettingsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetContentSettingsArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetContentSettingsResult struct {
	Success *tg.Bool
}

var SetContentSettingsResult_Success_DEFAULT *tg.Bool

func (p *SetContentSettingsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetContentSettingsResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetContentSettingsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetContentSettingsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetContentSettingsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetContentSettingsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetContentSettingsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetContentSettingsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetContentSettingsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetContentSettingsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetContentSettingsResult) GetResult() interface{} {
	return p.Success
}

func deleteContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteContactArgs)
	realResult := result.(*DeleteContactResult)
	success, err := handler.(user.RPCUser).UserDeleteContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteContactArgs() interface{} {
	return &DeleteContactArgs{}
}

func newDeleteContactResult() interface{} {
	return &DeleteContactResult{}
}

type DeleteContactArgs struct {
	Req *user.TLUserDeleteContact
}

func (p *DeleteContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteContactArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserDeleteContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserDeleteContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteContactArgs_Req_DEFAULT *user.TLUserDeleteContact

func (p *DeleteContactArgs) GetReq() *user.TLUserDeleteContact {
	if !p.IsSetReq() {
		return DeleteContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteContactResult struct {
	Success *tg.Bool
}

var DeleteContactResult_Success_DEFAULT *tg.Bool

func (p *DeleteContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteContactResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteContactResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteContactResult) GetResult() interface{} {
	return p.Success
}

func getContactListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetContactListArgs)
	realResult := result.(*GetContactListResult)
	success, err := handler.(user.RPCUser).UserGetContactList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetContactListArgs() interface{} {
	return &GetContactListArgs{}
}

func newGetContactListResult() interface{} {
	return &GetContactListResult{}
}

type GetContactListArgs struct {
	Req *user.TLUserGetContactList
}

func (p *GetContactListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetContactListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetContactListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetContactList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetContactListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetContactListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetContactListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetContactList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetContactListArgs_Req_DEFAULT *user.TLUserGetContactList

func (p *GetContactListArgs) GetReq() *user.TLUserGetContactList {
	if !p.IsSetReq() {
		return GetContactListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContactListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetContactListResult struct {
	Success *user.VectorContactData
}

var GetContactListResult_Success_DEFAULT *user.VectorContactData

func (p *GetContactListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetContactListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetContactListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorContactData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetContactListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetContactListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorContactData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactListResult) GetSuccess() *user.VectorContactData {
	if !p.IsSetSuccess() {
		return GetContactListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContactListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorContactData)
}

func (p *GetContactListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContactListResult) GetResult() interface{} {
	return p.Success
}

func getContactIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetContactIdListArgs)
	realResult := result.(*GetContactIdListResult)
	success, err := handler.(user.RPCUser).UserGetContactIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetContactIdListArgs() interface{} {
	return &GetContactIdListArgs{}
}

func newGetContactIdListResult() interface{} {
	return &GetContactIdListResult{}
}

type GetContactIdListArgs struct {
	Req *user.TLUserGetContactIdList
}

func (p *GetContactIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetContactIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetContactIdListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetContactIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetContactIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetContactIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetContactIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetContactIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetContactIdListArgs_Req_DEFAULT *user.TLUserGetContactIdList

func (p *GetContactIdListArgs) GetReq() *user.TLUserGetContactIdList {
	if !p.IsSetReq() {
		return GetContactIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContactIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetContactIdListResult struct {
	Success *user.VectorLong
}

var GetContactIdListResult_Success_DEFAULT *user.VectorLong

func (p *GetContactIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetContactIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetContactIdListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetContactIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetContactIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactIdListResult) GetSuccess() *user.VectorLong {
	if !p.IsSetSuccess() {
		return GetContactIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContactIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorLong)
}

func (p *GetContactIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContactIdListResult) GetResult() interface{} {
	return p.Success
}

func getContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetContactArgs)
	realResult := result.(*GetContactResult)
	success, err := handler.(user.RPCUser).UserGetContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetContactArgs() interface{} {
	return &GetContactArgs{}
}

func newGetContactResult() interface{} {
	return &GetContactResult{}
}

type GetContactArgs struct {
	Req *user.TLUserGetContact
}

func (p *GetContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetContactArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetContactArgs_Req_DEFAULT *user.TLUserGetContact

func (p *GetContactArgs) GetReq() *user.TLUserGetContact {
	if !p.IsSetReq() {
		return GetContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetContactResult struct {
	Success *tg.ContactData
}

var GetContactResult_Success_DEFAULT *tg.ContactData

func (p *GetContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetContactResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetContactResult) GetSuccess() *tg.ContactData {
	if !p.IsSetSuccess() {
		return GetContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactData)
}

func (p *GetContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContactResult) GetResult() interface{} {
	return p.Success
}

func addContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AddContactArgs)
	realResult := result.(*AddContactResult)
	success, err := handler.(user.RPCUser).UserAddContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAddContactArgs() interface{} {
	return &AddContactArgs{}
}

func newAddContactResult() interface{} {
	return &AddContactResult{}
}

type AddContactArgs struct {
	Req *user.TLUserAddContact
}

func (p *AddContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AddContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AddContactArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserAddContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AddContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AddContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AddContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserAddContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AddContactArgs_Req_DEFAULT *user.TLUserAddContact

func (p *AddContactArgs) GetReq() *user.TLUserAddContact {
	if !p.IsSetReq() {
		return AddContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type AddContactResult struct {
	Success *tg.Bool
}

var AddContactResult_Success_DEFAULT *tg.Bool

func (p *AddContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AddContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *AddContactResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AddContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AddContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddContactResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AddContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AddContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddContactResult) GetResult() interface{} {
	return p.Success
}

func checkContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckContactArgs)
	realResult := result.(*CheckContactResult)
	success, err := handler.(user.RPCUser).UserCheckContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckContactArgs() interface{} {
	return &CheckContactArgs{}
}

func newCheckContactResult() interface{} {
	return &CheckContactResult{}
}

type CheckContactArgs struct {
	Req *user.TLUserCheckContact
}

func (p *CheckContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckContactArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCheckContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCheckContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckContactArgs_Req_DEFAULT *user.TLUserCheckContact

func (p *CheckContactArgs) GetReq() *user.TLUserCheckContact {
	if !p.IsSetReq() {
		return CheckContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckContactResult struct {
	Success *tg.Bool
}

var CheckContactResult_Success_DEFAULT *tg.Bool

func (p *CheckContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckContactResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckContactResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return CheckContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *CheckContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckContactResult) GetResult() interface{} {
	return p.Success
}

func getImportersByPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetImportersByPhoneArgs)
	realResult := result.(*GetImportersByPhoneResult)
	success, err := handler.(user.RPCUser).UserGetImportersByPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetImportersByPhoneArgs() interface{} {
	return &GetImportersByPhoneArgs{}
}

func newGetImportersByPhoneResult() interface{} {
	return &GetImportersByPhoneResult{}
}

type GetImportersByPhoneArgs struct {
	Req *user.TLUserGetImportersByPhone
}

func (p *GetImportersByPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetImportersByPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetImportersByPhoneArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetImportersByPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetImportersByPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetImportersByPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetImportersByPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetImportersByPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetImportersByPhoneArgs_Req_DEFAULT *user.TLUserGetImportersByPhone

func (p *GetImportersByPhoneArgs) GetReq() *user.TLUserGetImportersByPhone {
	if !p.IsSetReq() {
		return GetImportersByPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImportersByPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetImportersByPhoneResult struct {
	Success *user.VectorInputContact
}

var GetImportersByPhoneResult_Success_DEFAULT *user.VectorInputContact

func (p *GetImportersByPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetImportersByPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetImportersByPhoneResult) Unmarshal(in []byte) error {
	msg := new(user.VectorInputContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImportersByPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetImportersByPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetImportersByPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorInputContact)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImportersByPhoneResult) GetSuccess() *user.VectorInputContact {
	if !p.IsSetSuccess() {
		return GetImportersByPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImportersByPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorInputContact)
}

func (p *GetImportersByPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImportersByPhoneResult) GetResult() interface{} {
	return p.Success
}

func deleteImportersByPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteImportersByPhoneArgs)
	realResult := result.(*DeleteImportersByPhoneResult)
	success, err := handler.(user.RPCUser).UserDeleteImportersByPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteImportersByPhoneArgs() interface{} {
	return &DeleteImportersByPhoneArgs{}
}

func newDeleteImportersByPhoneResult() interface{} {
	return &DeleteImportersByPhoneResult{}
}

type DeleteImportersByPhoneArgs struct {
	Req *user.TLUserDeleteImportersByPhone
}

func (p *DeleteImportersByPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteImportersByPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteImportersByPhoneArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserDeleteImportersByPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteImportersByPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteImportersByPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteImportersByPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserDeleteImportersByPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteImportersByPhoneArgs_Req_DEFAULT *user.TLUserDeleteImportersByPhone

func (p *DeleteImportersByPhoneArgs) GetReq() *user.TLUserDeleteImportersByPhone {
	if !p.IsSetReq() {
		return DeleteImportersByPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteImportersByPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteImportersByPhoneResult struct {
	Success *tg.Bool
}

var DeleteImportersByPhoneResult_Success_DEFAULT *tg.Bool

func (p *DeleteImportersByPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteImportersByPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteImportersByPhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteImportersByPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteImportersByPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteImportersByPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteImportersByPhoneResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return DeleteImportersByPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteImportersByPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *DeleteImportersByPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteImportersByPhoneResult) GetResult() interface{} {
	return p.Success
}

func importContactsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ImportContactsArgs)
	realResult := result.(*ImportContactsResult)
	success, err := handler.(user.RPCUser).UserImportContacts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newImportContactsArgs() interface{} {
	return &ImportContactsArgs{}
}

func newImportContactsResult() interface{} {
	return &ImportContactsResult{}
}

type ImportContactsArgs struct {
	Req *user.TLUserImportContacts
}

func (p *ImportContactsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ImportContactsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ImportContactsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserImportContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ImportContactsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ImportContactsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ImportContactsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserImportContacts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ImportContactsArgs_Req_DEFAULT *user.TLUserImportContacts

func (p *ImportContactsArgs) GetReq() *user.TLUserImportContacts {
	if !p.IsSetReq() {
		return ImportContactsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ImportContactsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ImportContactsResult struct {
	Success *user.UserImportedContacts
}

var ImportContactsResult_Success_DEFAULT *user.UserImportedContacts

func (p *ImportContactsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ImportContactsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ImportContactsResult) Unmarshal(in []byte) error {
	msg := new(user.UserImportedContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportContactsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ImportContactsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ImportContactsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.UserImportedContacts)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ImportContactsResult) GetSuccess() *user.UserImportedContacts {
	if !p.IsSetSuccess() {
		return ImportContactsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ImportContactsResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UserImportedContacts)
}

func (p *ImportContactsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ImportContactsResult) GetResult() interface{} {
	return p.Success
}

func getCountryCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetCountryCodeArgs)
	realResult := result.(*GetCountryCodeResult)
	success, err := handler.(user.RPCUser).UserGetCountryCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetCountryCodeArgs() interface{} {
	return &GetCountryCodeArgs{}
}

func newGetCountryCodeResult() interface{} {
	return &GetCountryCodeResult{}
}

type GetCountryCodeArgs struct {
	Req *user.TLUserGetCountryCode
}

func (p *GetCountryCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCountryCodeArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetCountryCodeArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetCountryCode)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetCountryCodeArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetCountryCodeArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetCountryCodeArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetCountryCode)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetCountryCodeArgs_Req_DEFAULT *user.TLUserGetCountryCode

func (p *GetCountryCodeArgs) GetReq() *user.TLUserGetCountryCode {
	if !p.IsSetReq() {
		return GetCountryCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCountryCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCountryCodeResult struct {
	Success *tg.String
}

var GetCountryCodeResult_Success_DEFAULT *tg.String

func (p *GetCountryCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCountryCodeResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetCountryCodeResult) Unmarshal(in []byte) error {
	msg := new(tg.String)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCountryCodeResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetCountryCodeResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetCountryCodeResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.String)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCountryCodeResult) GetSuccess() *tg.String {
	if !p.IsSetSuccess() {
		return GetCountryCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCountryCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.String)
}

func (p *GetCountryCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCountryCodeResult) GetResult() interface{} {
	return p.Success
}

func updateAboutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateAboutArgs)
	realResult := result.(*UpdateAboutResult)
	success, err := handler.(user.RPCUser).UserUpdateAbout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateAboutArgs() interface{} {
	return &UpdateAboutArgs{}
}

func newUpdateAboutResult() interface{} {
	return &UpdateAboutResult{}
}

type UpdateAboutArgs struct {
	Req *user.TLUserUpdateAbout
}

func (p *UpdateAboutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateAboutArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateAboutArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateAbout)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateAboutArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateAboutArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateAboutArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateAbout)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateAboutArgs_Req_DEFAULT *user.TLUserUpdateAbout

func (p *UpdateAboutArgs) GetReq() *user.TLUserUpdateAbout {
	if !p.IsSetReq() {
		return UpdateAboutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateAboutArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateAboutResult struct {
	Success *tg.Bool
}

var UpdateAboutResult_Success_DEFAULT *tg.Bool

func (p *UpdateAboutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateAboutResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateAboutResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateAboutResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateAboutResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateAboutResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateAboutResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateAboutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateAboutResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateAboutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateAboutResult) GetResult() interface{} {
	return p.Success
}

func updateFirstAndLastNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateFirstAndLastNameArgs)
	realResult := result.(*UpdateFirstAndLastNameResult)
	success, err := handler.(user.RPCUser).UserUpdateFirstAndLastName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateFirstAndLastNameArgs() interface{} {
	return &UpdateFirstAndLastNameArgs{}
}

func newUpdateFirstAndLastNameResult() interface{} {
	return &UpdateFirstAndLastNameResult{}
}

type UpdateFirstAndLastNameArgs struct {
	Req *user.TLUserUpdateFirstAndLastName
}

func (p *UpdateFirstAndLastNameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateFirstAndLastNameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateFirstAndLastNameArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateFirstAndLastName)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateFirstAndLastNameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateFirstAndLastNameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateFirstAndLastNameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateFirstAndLastName)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateFirstAndLastNameArgs_Req_DEFAULT *user.TLUserUpdateFirstAndLastName

func (p *UpdateFirstAndLastNameArgs) GetReq() *user.TLUserUpdateFirstAndLastName {
	if !p.IsSetReq() {
		return UpdateFirstAndLastNameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateFirstAndLastNameArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateFirstAndLastNameResult struct {
	Success *tg.Bool
}

var UpdateFirstAndLastNameResult_Success_DEFAULT *tg.Bool

func (p *UpdateFirstAndLastNameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateFirstAndLastNameResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateFirstAndLastNameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateFirstAndLastNameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateFirstAndLastNameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateFirstAndLastNameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateFirstAndLastNameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateFirstAndLastNameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateFirstAndLastNameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateFirstAndLastNameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateFirstAndLastNameResult) GetResult() interface{} {
	return p.Success
}

func updateVerifiedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateVerifiedArgs)
	realResult := result.(*UpdateVerifiedResult)
	success, err := handler.(user.RPCUser).UserUpdateVerified(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateVerifiedArgs() interface{} {
	return &UpdateVerifiedArgs{}
}

func newUpdateVerifiedResult() interface{} {
	return &UpdateVerifiedResult{}
}

type UpdateVerifiedArgs struct {
	Req *user.TLUserUpdateVerified
}

func (p *UpdateVerifiedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateVerifiedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateVerifiedArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateVerified)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateVerifiedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateVerifiedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateVerifiedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateVerified)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateVerifiedArgs_Req_DEFAULT *user.TLUserUpdateVerified

func (p *UpdateVerifiedArgs) GetReq() *user.TLUserUpdateVerified {
	if !p.IsSetReq() {
		return UpdateVerifiedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateVerifiedArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateVerifiedResult struct {
	Success *tg.Bool
}

var UpdateVerifiedResult_Success_DEFAULT *tg.Bool

func (p *UpdateVerifiedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateVerifiedResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateVerifiedResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateVerifiedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateVerifiedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateVerifiedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateVerifiedResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateVerifiedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateVerifiedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateVerifiedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateVerifiedResult) GetResult() interface{} {
	return p.Success
}

func updateUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateUsernameArgs)
	realResult := result.(*UpdateUsernameResult)
	success, err := handler.(user.RPCUser).UserUpdateUsername(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateUsernameArgs() interface{} {
	return &UpdateUsernameArgs{}
}

func newUpdateUsernameResult() interface{} {
	return &UpdateUsernameResult{}
}

type UpdateUsernameArgs struct {
	Req *user.TLUserUpdateUsername
}

func (p *UpdateUsernameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateUsernameArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateUsernameArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateUsername)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateUsernameArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateUsernameArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateUsernameArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateUsername)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateUsernameArgs_Req_DEFAULT *user.TLUserUpdateUsername

func (p *UpdateUsernameArgs) GetReq() *user.TLUserUpdateUsername {
	if !p.IsSetReq() {
		return UpdateUsernameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUsernameArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateUsernameResult struct {
	Success *tg.Bool
}

var UpdateUsernameResult_Success_DEFAULT *tg.Bool

func (p *UpdateUsernameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateUsernameResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateUsernameResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateUsernameResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateUsernameResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUsernameResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateUsernameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUsernameResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateUsernameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUsernameResult) GetResult() interface{} {
	return p.Success
}

func updateProfilePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateProfilePhotoArgs)
	realResult := result.(*UpdateProfilePhotoResult)
	success, err := handler.(user.RPCUser).UserUpdateProfilePhoto(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateProfilePhotoArgs() interface{} {
	return &UpdateProfilePhotoArgs{}
}

func newUpdateProfilePhotoResult() interface{} {
	return &UpdateProfilePhotoResult{}
}

type UpdateProfilePhotoArgs struct {
	Req *user.TLUserUpdateProfilePhoto
}

func (p *UpdateProfilePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateProfilePhotoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateProfilePhotoArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateProfilePhoto)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateProfilePhotoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateProfilePhotoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateProfilePhotoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateProfilePhoto)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateProfilePhotoArgs_Req_DEFAULT *user.TLUserUpdateProfilePhoto

func (p *UpdateProfilePhotoArgs) GetReq() *user.TLUserUpdateProfilePhoto {
	if !p.IsSetReq() {
		return UpdateProfilePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateProfilePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateProfilePhotoResult struct {
	Success *tg.Int64
}

var UpdateProfilePhotoResult_Success_DEFAULT *tg.Int64

func (p *UpdateProfilePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateProfilePhotoResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateProfilePhotoResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateProfilePhotoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateProfilePhotoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateProfilePhotoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateProfilePhotoResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return UpdateProfilePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateProfilePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *UpdateProfilePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateProfilePhotoResult) GetResult() interface{} {
	return p.Success
}

func deleteProfilePhotosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*DeleteProfilePhotosArgs)
	realResult := result.(*DeleteProfilePhotosResult)
	success, err := handler.(user.RPCUser).UserDeleteProfilePhotos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newDeleteProfilePhotosArgs() interface{} {
	return &DeleteProfilePhotosArgs{}
}

func newDeleteProfilePhotosResult() interface{} {
	return &DeleteProfilePhotosResult{}
}

type DeleteProfilePhotosArgs struct {
	Req *user.TLUserDeleteProfilePhotos
}

func (p *DeleteProfilePhotosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteProfilePhotosArgs")
	}
	return json.Marshal(p.Req)
}

func (p *DeleteProfilePhotosArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserDeleteProfilePhotos)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *DeleteProfilePhotosArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in DeleteProfilePhotosArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *DeleteProfilePhotosArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserDeleteProfilePhotos)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var DeleteProfilePhotosArgs_Req_DEFAULT *user.TLUserDeleteProfilePhotos

func (p *DeleteProfilePhotosArgs) GetReq() *user.TLUserDeleteProfilePhotos {
	if !p.IsSetReq() {
		return DeleteProfilePhotosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteProfilePhotosArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteProfilePhotosResult struct {
	Success *tg.Int64
}

var DeleteProfilePhotosResult_Success_DEFAULT *tg.Int64

func (p *DeleteProfilePhotosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteProfilePhotosResult")
	}
	return json.Marshal(p.Success)
}

func (p *DeleteProfilePhotosResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteProfilePhotosResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in DeleteProfilePhotosResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *DeleteProfilePhotosResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteProfilePhotosResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return DeleteProfilePhotosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteProfilePhotosResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *DeleteProfilePhotosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteProfilePhotosResult) GetResult() interface{} {
	return p.Success
}

func getProfilePhotosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetProfilePhotosArgs)
	realResult := result.(*GetProfilePhotosResult)
	success, err := handler.(user.RPCUser).UserGetProfilePhotos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetProfilePhotosArgs() interface{} {
	return &GetProfilePhotosArgs{}
}

func newGetProfilePhotosResult() interface{} {
	return &GetProfilePhotosResult{}
}

type GetProfilePhotosArgs struct {
	Req *user.TLUserGetProfilePhotos
}

func (p *GetProfilePhotosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetProfilePhotosArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetProfilePhotosArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetProfilePhotos)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetProfilePhotosArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetProfilePhotosArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetProfilePhotosArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetProfilePhotos)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetProfilePhotosArgs_Req_DEFAULT *user.TLUserGetProfilePhotos

func (p *GetProfilePhotosArgs) GetReq() *user.TLUserGetProfilePhotos {
	if !p.IsSetReq() {
		return GetProfilePhotosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProfilePhotosArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetProfilePhotosResult struct {
	Success *user.VectorLong
}

var GetProfilePhotosResult_Success_DEFAULT *user.VectorLong

func (p *GetProfilePhotosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetProfilePhotosResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetProfilePhotosResult) Unmarshal(in []byte) error {
	msg := new(user.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProfilePhotosResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetProfilePhotosResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetProfilePhotosResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProfilePhotosResult) GetSuccess() *user.VectorLong {
	if !p.IsSetSuccess() {
		return GetProfilePhotosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProfilePhotosResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorLong)
}

func (p *GetProfilePhotosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProfilePhotosResult) GetResult() interface{} {
	return p.Success
}

func setBotCommandsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetBotCommandsArgs)
	realResult := result.(*SetBotCommandsResult)
	success, err := handler.(user.RPCUser).UserSetBotCommands(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetBotCommandsArgs() interface{} {
	return &SetBotCommandsArgs{}
}

func newSetBotCommandsResult() interface{} {
	return &SetBotCommandsResult{}
}

type SetBotCommandsArgs struct {
	Req *user.TLUserSetBotCommands
}

func (p *SetBotCommandsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetBotCommandsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetBotCommandsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetBotCommands)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetBotCommandsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetBotCommandsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetBotCommandsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetBotCommands)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetBotCommandsArgs_Req_DEFAULT *user.TLUserSetBotCommands

func (p *SetBotCommandsArgs) GetReq() *user.TLUserSetBotCommands {
	if !p.IsSetReq() {
		return SetBotCommandsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetBotCommandsArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetBotCommandsResult struct {
	Success *tg.Bool
}

var SetBotCommandsResult_Success_DEFAULT *tg.Bool

func (p *SetBotCommandsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetBotCommandsResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetBotCommandsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetBotCommandsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetBotCommandsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetBotCommandsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetBotCommandsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetBotCommandsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetBotCommandsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetBotCommandsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetBotCommandsResult) GetResult() interface{} {
	return p.Success
}

func isBotHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*IsBotArgs)
	realResult := result.(*IsBotResult)
	success, err := handler.(user.RPCUser).UserIsBot(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newIsBotArgs() interface{} {
	return &IsBotArgs{}
}

func newIsBotResult() interface{} {
	return &IsBotResult{}
}

type IsBotArgs struct {
	Req *user.TLUserIsBot
}

func (p *IsBotArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsBotArgs")
	}
	return json.Marshal(p.Req)
}

func (p *IsBotArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserIsBot)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *IsBotArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in IsBotArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *IsBotArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserIsBot)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var IsBotArgs_Req_DEFAULT *user.TLUserIsBot

func (p *IsBotArgs) GetReq() *user.TLUserIsBot {
	if !p.IsSetReq() {
		return IsBotArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsBotArgs) IsSetReq() bool {
	return p.Req != nil
}

type IsBotResult struct {
	Success *tg.Bool
}

var IsBotResult_Success_DEFAULT *tg.Bool

func (p *IsBotResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsBotResult")
	}
	return json.Marshal(p.Success)
}

func (p *IsBotResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsBotResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in IsBotResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *IsBotResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsBotResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return IsBotResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsBotResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *IsBotResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsBotResult) GetResult() interface{} {
	return p.Success
}

func getBotInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetBotInfoArgs)
	realResult := result.(*GetBotInfoResult)
	success, err := handler.(user.RPCUser).UserGetBotInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetBotInfoArgs() interface{} {
	return &GetBotInfoArgs{}
}

func newGetBotInfoResult() interface{} {
	return &GetBotInfoResult{}
}

type GetBotInfoArgs struct {
	Req *user.TLUserGetBotInfo
}

func (p *GetBotInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBotInfoArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetBotInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetBotInfo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetBotInfoArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetBotInfoArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetBotInfoArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetBotInfo)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetBotInfoArgs_Req_DEFAULT *user.TLUserGetBotInfo

func (p *GetBotInfoArgs) GetReq() *user.TLUserGetBotInfo {
	if !p.IsSetReq() {
		return GetBotInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBotInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetBotInfoResult struct {
	Success *tg.BotInfo
}

var GetBotInfoResult_Success_DEFAULT *tg.BotInfo

func (p *GetBotInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBotInfoResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetBotInfoResult) Unmarshal(in []byte) error {
	msg := new(tg.BotInfo)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBotInfoResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetBotInfoResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetBotInfoResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.BotInfo)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBotInfoResult) GetSuccess() *tg.BotInfo {
	if !p.IsSetSuccess() {
		return GetBotInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBotInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.BotInfo)
}

func (p *GetBotInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetBotInfoResult) GetResult() interface{} {
	return p.Success
}

func checkBotsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CheckBotsArgs)
	realResult := result.(*CheckBotsResult)
	success, err := handler.(user.RPCUser).UserCheckBots(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCheckBotsArgs() interface{} {
	return &CheckBotsArgs{}
}

func newCheckBotsResult() interface{} {
	return &CheckBotsResult{}
}

type CheckBotsArgs struct {
	Req *user.TLUserCheckBots
}

func (p *CheckBotsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckBotsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CheckBotsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCheckBots)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CheckBotsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CheckBotsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CheckBotsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCheckBots)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CheckBotsArgs_Req_DEFAULT *user.TLUserCheckBots

func (p *CheckBotsArgs) GetReq() *user.TLUserCheckBots {
	if !p.IsSetReq() {
		return CheckBotsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckBotsArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckBotsResult struct {
	Success *user.VectorLong
}

var CheckBotsResult_Success_DEFAULT *user.VectorLong

func (p *CheckBotsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckBotsResult")
	}
	return json.Marshal(p.Success)
}

func (p *CheckBotsResult) Unmarshal(in []byte) error {
	msg := new(user.VectorLong)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckBotsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CheckBotsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CheckBotsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorLong)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckBotsResult) GetSuccess() *user.VectorLong {
	if !p.IsSetSuccess() {
		return CheckBotsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckBotsResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorLong)
}

func (p *CheckBotsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckBotsResult) GetResult() interface{} {
	return p.Success
}

func getFullUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetFullUserArgs)
	realResult := result.(*GetFullUserResult)
	success, err := handler.(user.RPCUser).UserGetFullUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetFullUserArgs() interface{} {
	return &GetFullUserArgs{}
}

func newGetFullUserResult() interface{} {
	return &GetFullUserResult{}
}

type GetFullUserArgs struct {
	Req *user.TLUserGetFullUser
}

func (p *GetFullUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFullUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetFullUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetFullUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetFullUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetFullUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetFullUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetFullUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetFullUserArgs_Req_DEFAULT *user.TLUserGetFullUser

func (p *GetFullUserArgs) GetReq() *user.TLUserGetFullUser {
	if !p.IsSetReq() {
		return GetFullUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFullUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFullUserResult struct {
	Success *tg.UsersUserFull
}

var GetFullUserResult_Success_DEFAULT *tg.UsersUserFull

func (p *GetFullUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFullUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetFullUserResult) Unmarshal(in []byte) error {
	msg := new(tg.UsersUserFull)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFullUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetFullUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetFullUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UsersUserFull)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFullUserResult) GetSuccess() *tg.UsersUserFull {
	if !p.IsSetSuccess() {
		return GetFullUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFullUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UsersUserFull)
}

func (p *GetFullUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFullUserResult) GetResult() interface{} {
	return p.Success
}

func updateEmojiStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateEmojiStatusArgs)
	realResult := result.(*UpdateEmojiStatusResult)
	success, err := handler.(user.RPCUser).UserUpdateEmojiStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateEmojiStatusArgs() interface{} {
	return &UpdateEmojiStatusArgs{}
}

func newUpdateEmojiStatusResult() interface{} {
	return &UpdateEmojiStatusResult{}
}

type UpdateEmojiStatusArgs struct {
	Req *user.TLUserUpdateEmojiStatus
}

func (p *UpdateEmojiStatusArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateEmojiStatusArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateEmojiStatusArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateEmojiStatus)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateEmojiStatusArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateEmojiStatusArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateEmojiStatusArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateEmojiStatus)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateEmojiStatusArgs_Req_DEFAULT *user.TLUserUpdateEmojiStatus

func (p *UpdateEmojiStatusArgs) GetReq() *user.TLUserUpdateEmojiStatus {
	if !p.IsSetReq() {
		return UpdateEmojiStatusArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateEmojiStatusArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateEmojiStatusResult struct {
	Success *tg.Bool
}

var UpdateEmojiStatusResult_Success_DEFAULT *tg.Bool

func (p *UpdateEmojiStatusResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateEmojiStatusResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateEmojiStatusResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateEmojiStatusResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateEmojiStatusResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateEmojiStatusResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateEmojiStatusResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateEmojiStatusResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateEmojiStatusResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateEmojiStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateEmojiStatusResult) GetResult() interface{} {
	return p.Success
}

func getUserDataByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserDataByIdArgs)
	realResult := result.(*GetUserDataByIdResult)
	success, err := handler.(user.RPCUser).UserGetUserDataById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserDataByIdArgs() interface{} {
	return &GetUserDataByIdArgs{}
}

func newGetUserDataByIdResult() interface{} {
	return &GetUserDataByIdResult{}
}

type GetUserDataByIdArgs struct {
	Req *user.TLUserGetUserDataById
}

func (p *GetUserDataByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserDataByIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserDataByIdArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetUserDataById)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserDataByIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserDataByIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserDataByIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetUserDataById)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserDataByIdArgs_Req_DEFAULT *user.TLUserGetUserDataById

func (p *GetUserDataByIdArgs) GetReq() *user.TLUserGetUserDataById {
	if !p.IsSetReq() {
		return GetUserDataByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserDataByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserDataByIdResult struct {
	Success *tg.UserData
}

var GetUserDataByIdResult_Success_DEFAULT *tg.UserData

func (p *GetUserDataByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserDataByIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserDataByIdResult) Unmarshal(in []byte) error {
	msg := new(tg.UserData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataByIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserDataByIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserDataByIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UserData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataByIdResult) GetSuccess() *tg.UserData {
	if !p.IsSetSuccess() {
		return GetUserDataByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserDataByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UserData)
}

func (p *GetUserDataByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserDataByIdResult) GetResult() interface{} {
	return p.Success
}

func getUserDataListByIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserDataListByIdListArgs)
	realResult := result.(*GetUserDataListByIdListResult)
	success, err := handler.(user.RPCUser).UserGetUserDataListByIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserDataListByIdListArgs() interface{} {
	return &GetUserDataListByIdListArgs{}
}

func newGetUserDataListByIdListResult() interface{} {
	return &GetUserDataListByIdListResult{}
}

type GetUserDataListByIdListArgs struct {
	Req *user.TLUserGetUserDataListByIdList
}

func (p *GetUserDataListByIdListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserDataListByIdListArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserDataListByIdListArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetUserDataListByIdList)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserDataListByIdListArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserDataListByIdListArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserDataListByIdListArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetUserDataListByIdList)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserDataListByIdListArgs_Req_DEFAULT *user.TLUserGetUserDataListByIdList

func (p *GetUserDataListByIdListArgs) GetReq() *user.TLUserGetUserDataListByIdList {
	if !p.IsSetReq() {
		return GetUserDataListByIdListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserDataListByIdListArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserDataListByIdListResult struct {
	Success *user.VectorUserData
}

var GetUserDataListByIdListResult_Success_DEFAULT *user.VectorUserData

func (p *GetUserDataListByIdListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserDataListByIdListResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserDataListByIdListResult) Unmarshal(in []byte) error {
	msg := new(user.VectorUserData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataListByIdListResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserDataListByIdListResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserDataListByIdListResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorUserData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataListByIdListResult) GetSuccess() *user.VectorUserData {
	if !p.IsSetSuccess() {
		return GetUserDataListByIdListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserDataListByIdListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorUserData)
}

func (p *GetUserDataListByIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserDataListByIdListResult) GetResult() interface{} {
	return p.Success
}

func getUserDataByTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserDataByTokenArgs)
	realResult := result.(*GetUserDataByTokenResult)
	success, err := handler.(user.RPCUser).UserGetUserDataByToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserDataByTokenArgs() interface{} {
	return &GetUserDataByTokenArgs{}
}

func newGetUserDataByTokenResult() interface{} {
	return &GetUserDataByTokenResult{}
}

type GetUserDataByTokenArgs struct {
	Req *user.TLUserGetUserDataByToken
}

func (p *GetUserDataByTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserDataByTokenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserDataByTokenArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetUserDataByToken)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserDataByTokenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserDataByTokenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserDataByTokenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetUserDataByToken)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserDataByTokenArgs_Req_DEFAULT *user.TLUserGetUserDataByToken

func (p *GetUserDataByTokenArgs) GetReq() *user.TLUserGetUserDataByToken {
	if !p.IsSetReq() {
		return GetUserDataByTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserDataByTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserDataByTokenResult struct {
	Success *tg.UserData
}

var GetUserDataByTokenResult_Success_DEFAULT *tg.UserData

func (p *GetUserDataByTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserDataByTokenResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserDataByTokenResult) Unmarshal(in []byte) error {
	msg := new(tg.UserData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataByTokenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserDataByTokenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserDataByTokenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.UserData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDataByTokenResult) GetSuccess() *tg.UserData {
	if !p.IsSetSuccess() {
		return GetUserDataByTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserDataByTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.UserData)
}

func (p *GetUserDataByTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserDataByTokenResult) GetResult() interface{} {
	return p.Success
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SearchArgs)
	realResult := result.(*SearchResult)
	success, err := handler.(user.RPCUser).UserSearch(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSearchArgs() interface{} {
	return &SearchArgs{}
}

func newSearchResult() interface{} {
	return &SearchResult{}
}

type SearchArgs struct {
	Req *user.TLUserSearch
}

func (p *SearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SearchArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSearch)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SearchArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SearchArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SearchArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SearchArgs_Req_DEFAULT *user.TLUserSearch

func (p *SearchArgs) GetReq() *user.TLUserSearch {
	if !p.IsSetReq() {
		return SearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type SearchResult struct {
	Success *user.UsersFound
}

var SearchResult_Success_DEFAULT *user.UsersFound

func (p *SearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *SearchResult) Unmarshal(in []byte) error {
	msg := new(user.UsersFound)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SearchResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SearchResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.UsersFound)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchResult) GetSuccess() *user.UsersFound {
	if !p.IsSetSuccess() {
		return SearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UsersFound)
}

func (p *SearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchResult) GetResult() interface{} {
	return p.Success
}

func updateBotDataHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateBotDataArgs)
	realResult := result.(*UpdateBotDataResult)
	success, err := handler.(user.RPCUser).UserUpdateBotData(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateBotDataArgs() interface{} {
	return &UpdateBotDataArgs{}
}

func newUpdateBotDataResult() interface{} {
	return &UpdateBotDataResult{}
}

type UpdateBotDataArgs struct {
	Req *user.TLUserUpdateBotData
}

func (p *UpdateBotDataArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateBotDataArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateBotDataArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateBotData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateBotDataArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateBotDataArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateBotDataArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateBotData)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateBotDataArgs_Req_DEFAULT *user.TLUserUpdateBotData

func (p *UpdateBotDataArgs) GetReq() *user.TLUserUpdateBotData {
	if !p.IsSetReq() {
		return UpdateBotDataArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateBotDataArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateBotDataResult struct {
	Success *tg.Bool
}

var UpdateBotDataResult_Success_DEFAULT *tg.Bool

func (p *UpdateBotDataResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateBotDataResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateBotDataResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBotDataResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateBotDataResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateBotDataResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBotDataResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateBotDataResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateBotDataResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateBotDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateBotDataResult) GetResult() interface{} {
	return p.Success
}

func getImmutableUserV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetImmutableUserV2Args)
	realResult := result.(*GetImmutableUserV2Result)
	success, err := handler.(user.RPCUser).UserGetImmutableUserV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetImmutableUserV2Args() interface{} {
	return &GetImmutableUserV2Args{}
}

func newGetImmutableUserV2Result() interface{} {
	return &GetImmutableUserV2Result{}
}

type GetImmutableUserV2Args struct {
	Req *user.TLUserGetImmutableUserV2
}

func (p *GetImmutableUserV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetImmutableUserV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetImmutableUserV2Args) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetImmutableUserV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetImmutableUserV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetImmutableUserV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetImmutableUserV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetImmutableUserV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetImmutableUserV2Args_Req_DEFAULT *user.TLUserGetImmutableUserV2

func (p *GetImmutableUserV2Args) GetReq() *user.TLUserGetImmutableUserV2 {
	if !p.IsSetReq() {
		return GetImmutableUserV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImmutableUserV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetImmutableUserV2Result struct {
	Success *tg.ImmutableUser
}

var GetImmutableUserV2Result_Success_DEFAULT *tg.ImmutableUser

func (p *GetImmutableUserV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetImmutableUserV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetImmutableUserV2Result) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetImmutableUserV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetImmutableUserV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImmutableUserV2Result) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return GetImmutableUserV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImmutableUserV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *GetImmutableUserV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImmutableUserV2Result) GetResult() interface{} {
	return p.Success
}

func getMutableUsersV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetMutableUsersV2Args)
	realResult := result.(*GetMutableUsersV2Result)
	success, err := handler.(user.RPCUser).UserGetMutableUsersV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetMutableUsersV2Args() interface{} {
	return &GetMutableUsersV2Args{}
}

func newGetMutableUsersV2Result() interface{} {
	return &GetMutableUsersV2Result{}
}

type GetMutableUsersV2Args struct {
	Req *user.TLUserGetMutableUsersV2
}

func (p *GetMutableUsersV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetMutableUsersV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetMutableUsersV2Args) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetMutableUsersV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetMutableUsersV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetMutableUsersV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetMutableUsersV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetMutableUsersV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetMutableUsersV2Args_Req_DEFAULT *user.TLUserGetMutableUsersV2

func (p *GetMutableUsersV2Args) GetReq() *user.TLUserGetMutableUsersV2 {
	if !p.IsSetReq() {
		return GetMutableUsersV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMutableUsersV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetMutableUsersV2Result struct {
	Success *tg.MutableUsers
}

var GetMutableUsersV2Result_Success_DEFAULT *tg.MutableUsers

func (p *GetMutableUsersV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetMutableUsersV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetMutableUsersV2Result) Unmarshal(in []byte) error {
	msg := new(tg.MutableUsers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableUsersV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetMutableUsersV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetMutableUsersV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.MutableUsers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMutableUsersV2Result) GetSuccess() *tg.MutableUsers {
	if !p.IsSetSuccess() {
		return GetMutableUsersV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMutableUsersV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*tg.MutableUsers)
}

func (p *GetMutableUsersV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMutableUsersV2Result) GetResult() interface{} {
	return p.Success
}

func createNewTestUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*CreateNewTestUserArgs)
	realResult := result.(*CreateNewTestUserResult)
	success, err := handler.(user.RPCUser).UserCreateNewTestUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newCreateNewTestUserArgs() interface{} {
	return &CreateNewTestUserArgs{}
}

func newCreateNewTestUserResult() interface{} {
	return &CreateNewTestUserResult{}
}

type CreateNewTestUserArgs struct {
	Req *user.TLUserCreateNewTestUser
}

func (p *CreateNewTestUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateNewTestUserArgs")
	}
	return json.Marshal(p.Req)
}

func (p *CreateNewTestUserArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserCreateNewTestUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *CreateNewTestUserArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in CreateNewTestUserArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *CreateNewTestUserArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserCreateNewTestUser)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var CreateNewTestUserArgs_Req_DEFAULT *user.TLUserCreateNewTestUser

func (p *CreateNewTestUserArgs) GetReq() *user.TLUserCreateNewTestUser {
	if !p.IsSetReq() {
		return CreateNewTestUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateNewTestUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateNewTestUserResult struct {
	Success *tg.ImmutableUser
}

var CreateNewTestUserResult_Success_DEFAULT *tg.ImmutableUser

func (p *CreateNewTestUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateNewTestUserResult")
	}
	return json.Marshal(p.Success)
}

func (p *CreateNewTestUserResult) Unmarshal(in []byte) error {
	msg := new(tg.ImmutableUser)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateNewTestUserResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in CreateNewTestUserResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *CreateNewTestUserResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ImmutableUser)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateNewTestUserResult) GetSuccess() *tg.ImmutableUser {
	if !p.IsSetSuccess() {
		return CreateNewTestUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateNewTestUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ImmutableUser)
}

func (p *CreateNewTestUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateNewTestUserResult) GetResult() interface{} {
	return p.Success
}

func editCloseFriendsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*EditCloseFriendsArgs)
	realResult := result.(*EditCloseFriendsResult)
	success, err := handler.(user.RPCUser).UserEditCloseFriends(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newEditCloseFriendsArgs() interface{} {
	return &EditCloseFriendsArgs{}
}

func newEditCloseFriendsResult() interface{} {
	return &EditCloseFriendsResult{}
}

type EditCloseFriendsArgs struct {
	Req *user.TLUserEditCloseFriends
}

func (p *EditCloseFriendsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in EditCloseFriendsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *EditCloseFriendsArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserEditCloseFriends)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *EditCloseFriendsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in EditCloseFriendsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *EditCloseFriendsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserEditCloseFriends)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var EditCloseFriendsArgs_Req_DEFAULT *user.TLUserEditCloseFriends

func (p *EditCloseFriendsArgs) GetReq() *user.TLUserEditCloseFriends {
	if !p.IsSetReq() {
		return EditCloseFriendsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EditCloseFriendsArgs) IsSetReq() bool {
	return p.Req != nil
}

type EditCloseFriendsResult struct {
	Success *tg.Bool
}

var EditCloseFriendsResult_Success_DEFAULT *tg.Bool

func (p *EditCloseFriendsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in EditCloseFriendsResult")
	}
	return json.Marshal(p.Success)
}

func (p *EditCloseFriendsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditCloseFriendsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in EditCloseFriendsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *EditCloseFriendsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EditCloseFriendsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return EditCloseFriendsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EditCloseFriendsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *EditCloseFriendsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EditCloseFriendsResult) GetResult() interface{} {
	return p.Success
}

func setStoriesMaxIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetStoriesMaxIdArgs)
	realResult := result.(*SetStoriesMaxIdResult)
	success, err := handler.(user.RPCUser).UserSetStoriesMaxId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetStoriesMaxIdArgs() interface{} {
	return &SetStoriesMaxIdArgs{}
}

func newSetStoriesMaxIdResult() interface{} {
	return &SetStoriesMaxIdResult{}
}

type SetStoriesMaxIdArgs struct {
	Req *user.TLUserSetStoriesMaxId
}

func (p *SetStoriesMaxIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetStoriesMaxIdArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetStoriesMaxIdArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetStoriesMaxId)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetStoriesMaxIdArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetStoriesMaxIdArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetStoriesMaxIdArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetStoriesMaxId)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetStoriesMaxIdArgs_Req_DEFAULT *user.TLUserSetStoriesMaxId

func (p *SetStoriesMaxIdArgs) GetReq() *user.TLUserSetStoriesMaxId {
	if !p.IsSetReq() {
		return SetStoriesMaxIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetStoriesMaxIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetStoriesMaxIdResult struct {
	Success *tg.Bool
}

var SetStoriesMaxIdResult_Success_DEFAULT *tg.Bool

func (p *SetStoriesMaxIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetStoriesMaxIdResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetStoriesMaxIdResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetStoriesMaxIdResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetStoriesMaxIdResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetStoriesMaxIdResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetStoriesMaxIdResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetStoriesMaxIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetStoriesMaxIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetStoriesMaxIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetStoriesMaxIdResult) GetResult() interface{} {
	return p.Success
}

func setColorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetColorArgs)
	realResult := result.(*SetColorResult)
	success, err := handler.(user.RPCUser).UserSetColor(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetColorArgs() interface{} {
	return &SetColorArgs{}
}

func newSetColorResult() interface{} {
	return &SetColorResult{}
}

type SetColorArgs struct {
	Req *user.TLUserSetColor
}

func (p *SetColorArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetColorArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetColorArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetColor)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetColorArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetColorArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetColorArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetColor)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetColorArgs_Req_DEFAULT *user.TLUserSetColor

func (p *SetColorArgs) GetReq() *user.TLUserSetColor {
	if !p.IsSetReq() {
		return SetColorArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetColorArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetColorResult struct {
	Success *tg.Bool
}

var SetColorResult_Success_DEFAULT *tg.Bool

func (p *SetColorResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetColorResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetColorResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetColorResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetColorResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetColorResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetColorResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetColorResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetColorResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetColorResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetColorResult) GetResult() interface{} {
	return p.Success
}

func updateBirthdayHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdateBirthdayArgs)
	realResult := result.(*UpdateBirthdayResult)
	success, err := handler.(user.RPCUser).UserUpdateBirthday(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdateBirthdayArgs() interface{} {
	return &UpdateBirthdayArgs{}
}

func newUpdateBirthdayResult() interface{} {
	return &UpdateBirthdayResult{}
}

type UpdateBirthdayArgs struct {
	Req *user.TLUserUpdateBirthday
}

func (p *UpdateBirthdayArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateBirthdayArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdateBirthdayArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdateBirthday)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdateBirthdayArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdateBirthdayArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdateBirthdayArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdateBirthday)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdateBirthdayArgs_Req_DEFAULT *user.TLUserUpdateBirthday

func (p *UpdateBirthdayArgs) GetReq() *user.TLUserUpdateBirthday {
	if !p.IsSetReq() {
		return UpdateBirthdayArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateBirthdayArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateBirthdayResult struct {
	Success *tg.Bool
}

var UpdateBirthdayResult_Success_DEFAULT *tg.Bool

func (p *UpdateBirthdayResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateBirthdayResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdateBirthdayResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBirthdayResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdateBirthdayResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdateBirthdayResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBirthdayResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdateBirthdayResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateBirthdayResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdateBirthdayResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateBirthdayResult) GetResult() interface{} {
	return p.Success
}

func getBirthdaysHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetBirthdaysArgs)
	realResult := result.(*GetBirthdaysResult)
	success, err := handler.(user.RPCUser).UserGetBirthdays(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetBirthdaysArgs() interface{} {
	return &GetBirthdaysArgs{}
}

func newGetBirthdaysResult() interface{} {
	return &GetBirthdaysResult{}
}

type GetBirthdaysArgs struct {
	Req *user.TLUserGetBirthdays
}

func (p *GetBirthdaysArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBirthdaysArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetBirthdaysArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetBirthdays)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetBirthdaysArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetBirthdaysArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetBirthdaysArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetBirthdays)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetBirthdaysArgs_Req_DEFAULT *user.TLUserGetBirthdays

func (p *GetBirthdaysArgs) GetReq() *user.TLUserGetBirthdays {
	if !p.IsSetReq() {
		return GetBirthdaysArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBirthdaysArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetBirthdaysResult struct {
	Success *user.VectorContactBirthday
}

var GetBirthdaysResult_Success_DEFAULT *user.VectorContactBirthday

func (p *GetBirthdaysResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBirthdaysResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetBirthdaysResult) Unmarshal(in []byte) error {
	msg := new(user.VectorContactBirthday)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBirthdaysResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetBirthdaysResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetBirthdaysResult) Decode(d *bin.Decoder) (err error) {
	msg := new(user.VectorContactBirthday)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBirthdaysResult) GetSuccess() *user.VectorContactBirthday {
	if !p.IsSetSuccess() {
		return GetBirthdaysResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBirthdaysResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.VectorContactBirthday)
}

func (p *GetBirthdaysResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetBirthdaysResult) GetResult() interface{} {
	return p.Success
}

func setStoriesHiddenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetStoriesHiddenArgs)
	realResult := result.(*SetStoriesHiddenResult)
	success, err := handler.(user.RPCUser).UserSetStoriesHidden(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetStoriesHiddenArgs() interface{} {
	return &SetStoriesHiddenArgs{}
}

func newSetStoriesHiddenResult() interface{} {
	return &SetStoriesHiddenResult{}
}

type SetStoriesHiddenArgs struct {
	Req *user.TLUserSetStoriesHidden
}

func (p *SetStoriesHiddenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetStoriesHiddenArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetStoriesHiddenArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetStoriesHidden)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetStoriesHiddenArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetStoriesHiddenArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetStoriesHiddenArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetStoriesHidden)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetStoriesHiddenArgs_Req_DEFAULT *user.TLUserSetStoriesHidden

func (p *SetStoriesHiddenArgs) GetReq() *user.TLUserSetStoriesHidden {
	if !p.IsSetReq() {
		return SetStoriesHiddenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetStoriesHiddenArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetStoriesHiddenResult struct {
	Success *tg.Bool
}

var SetStoriesHiddenResult_Success_DEFAULT *tg.Bool

func (p *SetStoriesHiddenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetStoriesHiddenResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetStoriesHiddenResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetStoriesHiddenResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetStoriesHiddenResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetStoriesHiddenResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetStoriesHiddenResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetStoriesHiddenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetStoriesHiddenResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetStoriesHiddenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetStoriesHiddenResult) GetResult() interface{} {
	return p.Success
}

func updatePersonalChannelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePersonalChannelArgs)
	realResult := result.(*UpdatePersonalChannelResult)
	success, err := handler.(user.RPCUser).UserUpdatePersonalChannel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePersonalChannelArgs() interface{} {
	return &UpdatePersonalChannelArgs{}
}

func newUpdatePersonalChannelResult() interface{} {
	return &UpdatePersonalChannelResult{}
}

type UpdatePersonalChannelArgs struct {
	Req *user.TLUserUpdatePersonalChannel
}

func (p *UpdatePersonalChannelArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePersonalChannelArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePersonalChannelArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdatePersonalChannel)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePersonalChannelArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePersonalChannelArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePersonalChannelArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdatePersonalChannel)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePersonalChannelArgs_Req_DEFAULT *user.TLUserUpdatePersonalChannel

func (p *UpdatePersonalChannelArgs) GetReq() *user.TLUserUpdatePersonalChannel {
	if !p.IsSetReq() {
		return UpdatePersonalChannelArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePersonalChannelArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePersonalChannelResult struct {
	Success *tg.Bool
}

var UpdatePersonalChannelResult_Success_DEFAULT *tg.Bool

func (p *UpdatePersonalChannelResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePersonalChannelResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePersonalChannelResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePersonalChannelResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePersonalChannelResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePersonalChannelResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePersonalChannelResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdatePersonalChannelResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePersonalChannelResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdatePersonalChannelResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePersonalChannelResult) GetResult() interface{} {
	return p.Success
}

func getUserIdByPhoneHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetUserIdByPhoneArgs)
	realResult := result.(*GetUserIdByPhoneResult)
	success, err := handler.(user.RPCUser).UserGetUserIdByPhone(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetUserIdByPhoneArgs() interface{} {
	return &GetUserIdByPhoneArgs{}
}

func newGetUserIdByPhoneResult() interface{} {
	return &GetUserIdByPhoneResult{}
}

type GetUserIdByPhoneArgs struct {
	Req *user.TLUserGetUserIdByPhone
}

func (p *GetUserIdByPhoneArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserIdByPhoneArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetUserIdByPhoneArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetUserIdByPhone)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetUserIdByPhoneArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetUserIdByPhoneArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetUserIdByPhoneArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetUserIdByPhone)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetUserIdByPhoneArgs_Req_DEFAULT *user.TLUserGetUserIdByPhone

func (p *GetUserIdByPhoneArgs) GetReq() *user.TLUserGetUserIdByPhone {
	if !p.IsSetReq() {
		return GetUserIdByPhoneArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserIdByPhoneArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserIdByPhoneResult struct {
	Success *tg.Int64
}

var GetUserIdByPhoneResult_Success_DEFAULT *tg.Int64

func (p *GetUserIdByPhoneResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserIdByPhoneResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetUserIdByPhoneResult) Unmarshal(in []byte) error {
	msg := new(tg.Int64)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserIdByPhoneResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetUserIdByPhoneResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetUserIdByPhoneResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Int64)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserIdByPhoneResult) GetSuccess() *tg.Int64 {
	if !p.IsSetSuccess() {
		return GetUserIdByPhoneResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserIdByPhoneResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Int64)
}

func (p *GetUserIdByPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserIdByPhoneResult) GetResult() interface{} {
	return p.Success
}

func setAuthorizationTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*SetAuthorizationTTLArgs)
	realResult := result.(*SetAuthorizationTTLResult)
	success, err := handler.(user.RPCUser).UserSetAuthorizationTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newSetAuthorizationTTLArgs() interface{} {
	return &SetAuthorizationTTLArgs{}
}

func newSetAuthorizationTTLResult() interface{} {
	return &SetAuthorizationTTLResult{}
}

type SetAuthorizationTTLArgs struct {
	Req *user.TLUserSetAuthorizationTTL
}

func (p *SetAuthorizationTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetAuthorizationTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *SetAuthorizationTTLArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserSetAuthorizationTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *SetAuthorizationTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in SetAuthorizationTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *SetAuthorizationTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserSetAuthorizationTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var SetAuthorizationTTLArgs_Req_DEFAULT *user.TLUserSetAuthorizationTTL

func (p *SetAuthorizationTTLArgs) GetReq() *user.TLUserSetAuthorizationTTL {
	if !p.IsSetReq() {
		return SetAuthorizationTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetAuthorizationTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type SetAuthorizationTTLResult struct {
	Success *tg.Bool
}

var SetAuthorizationTTLResult_Success_DEFAULT *tg.Bool

func (p *SetAuthorizationTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetAuthorizationTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *SetAuthorizationTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAuthorizationTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in SetAuthorizationTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *SetAuthorizationTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetAuthorizationTTLResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return SetAuthorizationTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetAuthorizationTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *SetAuthorizationTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetAuthorizationTTLResult) GetResult() interface{} {
	return p.Success
}

func getAuthorizationTTLHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetAuthorizationTTLArgs)
	realResult := result.(*GetAuthorizationTTLResult)
	success, err := handler.(user.RPCUser).UserGetAuthorizationTTL(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetAuthorizationTTLArgs() interface{} {
	return &GetAuthorizationTTLArgs{}
}

func newGetAuthorizationTTLResult() interface{} {
	return &GetAuthorizationTTLResult{}
}

type GetAuthorizationTTLArgs struct {
	Req *user.TLUserGetAuthorizationTTL
}

func (p *GetAuthorizationTTLArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAuthorizationTTLArgs")
	}
	return json.Marshal(p.Req)
}

func (p *GetAuthorizationTTLArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetAuthorizationTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetAuthorizationTTLArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetAuthorizationTTLArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetAuthorizationTTLArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetAuthorizationTTL)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetAuthorizationTTLArgs_Req_DEFAULT *user.TLUserGetAuthorizationTTL

func (p *GetAuthorizationTTLArgs) GetReq() *user.TLUserGetAuthorizationTTL {
	if !p.IsSetReq() {
		return GetAuthorizationTTLArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAuthorizationTTLArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAuthorizationTTLResult struct {
	Success *tg.AccountDaysTTL
}

var GetAuthorizationTTLResult_Success_DEFAULT *tg.AccountDaysTTL

func (p *GetAuthorizationTTLResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAuthorizationTTLResult")
	}
	return json.Marshal(p.Success)
}

func (p *GetAuthorizationTTLResult) Unmarshal(in []byte) error {
	msg := new(tg.AccountDaysTTL)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationTTLResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetAuthorizationTTLResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetAuthorizationTTLResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.AccountDaysTTL)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAuthorizationTTLResult) GetSuccess() *tg.AccountDaysTTL {
	if !p.IsSetSuccess() {
		return GetAuthorizationTTLResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAuthorizationTTLResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.AccountDaysTTL)
}

func (p *GetAuthorizationTTLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAuthorizationTTLResult) GetResult() interface{} {
	return p.Success
}

func updatePremiumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*UpdatePremiumArgs)
	realResult := result.(*UpdatePremiumResult)
	success, err := handler.(user.RPCUser).UserUpdatePremium(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newUpdatePremiumArgs() interface{} {
	return &UpdatePremiumArgs{}
}

func newUpdatePremiumResult() interface{} {
	return &UpdatePremiumResult{}
}

type UpdatePremiumArgs struct {
	Req *user.TLUserUpdatePremium
}

func (p *UpdatePremiumArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdatePremiumArgs")
	}
	return json.Marshal(p.Req)
}

func (p *UpdatePremiumArgs) Unmarshal(in []byte) error {
	msg := new(user.TLUserUpdatePremium)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *UpdatePremiumArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in UpdatePremiumArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *UpdatePremiumArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserUpdatePremium)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var UpdatePremiumArgs_Req_DEFAULT *user.TLUserUpdatePremium

func (p *UpdatePremiumArgs) GetReq() *user.TLUserUpdatePremium {
	if !p.IsSetReq() {
		return UpdatePremiumArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePremiumArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdatePremiumResult struct {
	Success *tg.Bool
}

var UpdatePremiumResult_Success_DEFAULT *tg.Bool

func (p *UpdatePremiumResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdatePremiumResult")
	}
	return json.Marshal(p.Success)
}

func (p *UpdatePremiumResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePremiumResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in UpdatePremiumResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *UpdatePremiumResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePremiumResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return UpdatePremiumResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePremiumResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *UpdatePremiumResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePremiumResult) GetResult() interface{} {
	return p.Success
}

func getBotInfoV2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*GetBotInfoV2Args)
	realResult := result.(*GetBotInfoV2Result)
	success, err := handler.(user.RPCUser).UserGetBotInfoV2(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newGetBotInfoV2Args() interface{} {
	return &GetBotInfoV2Args{}
}

func newGetBotInfoV2Result() interface{} {
	return &GetBotInfoV2Result{}
}

type GetBotInfoV2Args struct {
	Req *user.TLUserGetBotInfoV2
}

func (p *GetBotInfoV2Args) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBotInfoV2Args")
	}
	return json.Marshal(p.Req)
}

func (p *GetBotInfoV2Args) Unmarshal(in []byte) error {
	msg := new(user.TLUserGetBotInfoV2)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *GetBotInfoV2Args) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in GetBotInfoV2Args")
	}

	return p.Req.Encode(x, layer)
}

func (p *GetBotInfoV2Args) Decode(d *bin.Decoder) (err error) {
	msg := new(user.TLUserGetBotInfoV2)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var GetBotInfoV2Args_Req_DEFAULT *user.TLUserGetBotInfoV2

func (p *GetBotInfoV2Args) GetReq() *user.TLUserGetBotInfoV2 {
	if !p.IsSetReq() {
		return GetBotInfoV2Args_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBotInfoV2Args) IsSetReq() bool {
	return p.Req != nil
}

type GetBotInfoV2Result struct {
	Success *user.BotInfoData
}

var GetBotInfoV2Result_Success_DEFAULT *user.BotInfoData

func (p *GetBotInfoV2Result) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBotInfoV2Result")
	}
	return json.Marshal(p.Success)
}

func (p *GetBotInfoV2Result) Unmarshal(in []byte) error {
	msg := new(user.BotInfoData)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBotInfoV2Result) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in GetBotInfoV2Result")
	}

	return p.Success.Encode(x, layer)
}

func (p *GetBotInfoV2Result) Decode(d *bin.Decoder) (err error) {
	msg := new(user.BotInfoData)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBotInfoV2Result) GetSuccess() *user.BotInfoData {
	if !p.IsSetSuccess() {
		return GetBotInfoV2Result_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBotInfoV2Result) SetSuccess(x interface{}) {
	p.Success = x.(*user.BotInfoData)
}

func (p *GetBotInfoV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetBotInfoV2Result) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserGetLastSeens(ctx context.Context, req *user.TLUserGetLastSeens) (r *user.VectorLastSeenData, err error) {
	// var _args GetLastSeensArgs
	// _args.Req = req
	// var _result GetLastSeensResult

	_result := new(user.VectorLastSeenData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getLastSeens", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateLastSeen(ctx context.Context, req *user.TLUserUpdateLastSeen) (r *tg.Bool, err error) {
	// var _args UpdateLastSeenArgs
	// _args.Req = req
	// var _result UpdateLastSeenResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateLastSeen", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetLastSeen(ctx context.Context, req *user.TLUserGetLastSeen) (r *user.LastSeenData, err error) {
	// var _args GetLastSeenArgs
	// _args.Req = req
	// var _result GetLastSeenResult

	_result := new(user.LastSeenData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getLastSeen", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetImmutableUser(ctx context.Context, req *user.TLUserGetImmutableUser) (r *tg.ImmutableUser, err error) {
	// var _args GetImmutableUserArgs
	// _args.Req = req
	// var _result GetImmutableUserResult

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getImmutableUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetMutableUsers(ctx context.Context, req *user.TLUserGetMutableUsers) (r *user.VectorImmutableUser, err error) {
	// var _args GetMutableUsersArgs
	// _args.Req = req
	// var _result GetMutableUsersResult

	_result := new(user.VectorImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getMutableUsers", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetImmutableUserByPhone(ctx context.Context, req *user.TLUserGetImmutableUserByPhone) (r *tg.ImmutableUser, err error) {
	// var _args GetImmutableUserByPhoneArgs
	// _args.Req = req
	// var _result GetImmutableUserByPhoneResult

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getImmutableUserByPhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetImmutableUserByToken(ctx context.Context, req *user.TLUserGetImmutableUserByToken) (r *tg.ImmutableUser, err error) {
	// var _args GetImmutableUserByTokenArgs
	// _args.Req = req
	// var _result GetImmutableUserByTokenResult

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getImmutableUserByToken", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetAccountDaysTTL(ctx context.Context, req *user.TLUserSetAccountDaysTTL) (r *tg.Bool, err error) {
	// var _args SetAccountDaysTTLArgs
	// _args.Req = req
	// var _result SetAccountDaysTTLResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setAccountDaysTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetAccountDaysTTL(ctx context.Context, req *user.TLUserGetAccountDaysTTL) (r *tg.AccountDaysTTL, err error) {
	// var _args GetAccountDaysTTLArgs
	// _args.Req = req
	// var _result GetAccountDaysTTLResult

	_result := new(tg.AccountDaysTTL)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getAccountDaysTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetNotifySettings(ctx context.Context, req *user.TLUserGetNotifySettings) (r *tg.PeerNotifySettings, err error) {
	// var _args GetNotifySettingsArgs
	// _args.Req = req
	// var _result GetNotifySettingsResult

	_result := new(tg.PeerNotifySettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetNotifySettingsList(ctx context.Context, req *user.TLUserGetNotifySettingsList) (r *user.VectorPeerPeerNotifySettings, err error) {
	// var _args GetNotifySettingsListArgs
	// _args.Req = req
	// var _result GetNotifySettingsListResult

	_result := new(user.VectorPeerPeerNotifySettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getNotifySettingsList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetNotifySettings(ctx context.Context, req *user.TLUserSetNotifySettings) (r *tg.Bool, err error) {
	// var _args SetNotifySettingsArgs
	// _args.Req = req
	// var _result SetNotifySettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserResetNotifySettings(ctx context.Context, req *user.TLUserResetNotifySettings) (r *tg.Bool, err error) {
	// var _args ResetNotifySettingsArgs
	// _args.Req = req
	// var _result ResetNotifySettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.resetNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetAllNotifySettings(ctx context.Context, req *user.TLUserGetAllNotifySettings) (r *user.VectorPeerPeerNotifySettings, err error) {
	// var _args GetAllNotifySettingsArgs
	// _args.Req = req
	// var _result GetAllNotifySettingsResult

	_result := new(user.VectorPeerPeerNotifySettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getAllNotifySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetGlobalPrivacySettings(ctx context.Context, req *user.TLUserGetGlobalPrivacySettings) (r *tg.GlobalPrivacySettings, err error) {
	// var _args GetGlobalPrivacySettingsArgs
	// _args.Req = req
	// var _result GetGlobalPrivacySettingsResult

	_result := new(tg.GlobalPrivacySettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getGlobalPrivacySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetGlobalPrivacySettings(ctx context.Context, req *user.TLUserSetGlobalPrivacySettings) (r *tg.Bool, err error) {
	// var _args SetGlobalPrivacySettingsArgs
	// _args.Req = req
	// var _result SetGlobalPrivacySettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setGlobalPrivacySettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetPrivacy(ctx context.Context, req *user.TLUserGetPrivacy) (r *user.VectorPrivacyRule, err error) {
	// var _args GetPrivacyArgs
	// _args.Req = req
	// var _result GetPrivacyResult

	_result := new(user.VectorPrivacyRule)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getPrivacy", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetPrivacy(ctx context.Context, req *user.TLUserSetPrivacy) (r *tg.Bool, err error) {
	// var _args SetPrivacyArgs
	// _args.Req = req
	// var _result SetPrivacyResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setPrivacy", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCheckPrivacy(ctx context.Context, req *user.TLUserCheckPrivacy) (r *tg.Bool, err error) {
	// var _args CheckPrivacyArgs
	// _args.Req = req
	// var _result CheckPrivacyResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.checkPrivacy", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserAddPeerSettings(ctx context.Context, req *user.TLUserAddPeerSettings) (r *tg.Bool, err error) {
	// var _args AddPeerSettingsArgs
	// _args.Req = req
	// var _result AddPeerSettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.addPeerSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetPeerSettings(ctx context.Context, req *user.TLUserGetPeerSettings) (r *tg.PeerSettings, err error) {
	// var _args GetPeerSettingsArgs
	// _args.Req = req
	// var _result GetPeerSettingsResult

	_result := new(tg.PeerSettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getPeerSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserDeletePeerSettings(ctx context.Context, req *user.TLUserDeletePeerSettings) (r *tg.Bool, err error) {
	// var _args DeletePeerSettingsArgs
	// _args.Req = req
	// var _result DeletePeerSettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.deletePeerSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserChangePhone(ctx context.Context, req *user.TLUserChangePhone) (r *tg.Bool, err error) {
	// var _args ChangePhoneArgs
	// _args.Req = req
	// var _result ChangePhoneResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.changePhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCreateNewUser(ctx context.Context, req *user.TLUserCreateNewUser) (r *tg.ImmutableUser, err error) {
	// var _args CreateNewUserArgs
	// _args.Req = req
	// var _result CreateNewUserResult

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.createNewUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserDeleteUser(ctx context.Context, req *user.TLUserDeleteUser) (r *tg.Bool, err error) {
	// var _args DeleteUserArgs
	// _args.Req = req
	// var _result DeleteUserResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.deleteUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserBlockPeer(ctx context.Context, req *user.TLUserBlockPeer) (r *tg.Bool, err error) {
	// var _args BlockPeerArgs
	// _args.Req = req
	// var _result BlockPeerResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.blockPeer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUnBlockPeer(ctx context.Context, req *user.TLUserUnBlockPeer) (r *tg.Bool, err error) {
	// var _args UnBlockPeerArgs
	// _args.Req = req
	// var _result UnBlockPeerResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.unBlockPeer", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserBlockedByUser(ctx context.Context, req *user.TLUserBlockedByUser) (r *tg.Bool, err error) {
	// var _args BlockedByUserArgs
	// _args.Req = req
	// var _result BlockedByUserResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.blockedByUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserIsBlockedByUser(ctx context.Context, req *user.TLUserIsBlockedByUser) (r *tg.Bool, err error) {
	// var _args IsBlockedByUserArgs
	// _args.Req = req
	// var _result IsBlockedByUserResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.isBlockedByUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCheckBlockUserList(ctx context.Context, req *user.TLUserCheckBlockUserList) (r *user.VectorLong, err error) {
	// var _args CheckBlockUserListArgs
	// _args.Req = req
	// var _result CheckBlockUserListResult

	_result := new(user.VectorLong)

	if err = p.c.Call(ctx, "/user.RPCUser/user.checkBlockUserList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetBlockedList(ctx context.Context, req *user.TLUserGetBlockedList) (r *user.VectorPeerBlocked, err error) {
	// var _args GetBlockedListArgs
	// _args.Req = req
	// var _result GetBlockedListResult

	_result := new(user.VectorPeerBlocked)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getBlockedList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetContactSignUpNotification(ctx context.Context, req *user.TLUserGetContactSignUpNotification) (r *tg.Bool, err error) {
	// var _args GetContactSignUpNotificationArgs
	// _args.Req = req
	// var _result GetContactSignUpNotificationResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getContactSignUpNotification", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetContactSignUpNotification(ctx context.Context, req *user.TLUserSetContactSignUpNotification) (r *tg.Bool, err error) {
	// var _args SetContactSignUpNotificationArgs
	// _args.Req = req
	// var _result SetContactSignUpNotificationResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setContactSignUpNotification", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetContentSettings(ctx context.Context, req *user.TLUserGetContentSettings) (r *tg.AccountContentSettings, err error) {
	// var _args GetContentSettingsArgs
	// _args.Req = req
	// var _result GetContentSettingsResult

	_result := new(tg.AccountContentSettings)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getContentSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetContentSettings(ctx context.Context, req *user.TLUserSetContentSettings) (r *tg.Bool, err error) {
	// var _args SetContentSettingsArgs
	// _args.Req = req
	// var _result SetContentSettingsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setContentSettings", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserDeleteContact(ctx context.Context, req *user.TLUserDeleteContact) (r *tg.Bool, err error) {
	// var _args DeleteContactArgs
	// _args.Req = req
	// var _result DeleteContactResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.deleteContact", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetContactList(ctx context.Context, req *user.TLUserGetContactList) (r *user.VectorContactData, err error) {
	// var _args GetContactListArgs
	// _args.Req = req
	// var _result GetContactListResult

	_result := new(user.VectorContactData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getContactList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetContactIdList(ctx context.Context, req *user.TLUserGetContactIdList) (r *user.VectorLong, err error) {
	// var _args GetContactIdListArgs
	// _args.Req = req
	// var _result GetContactIdListResult

	_result := new(user.VectorLong)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getContactIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetContact(ctx context.Context, req *user.TLUserGetContact) (r *tg.ContactData, err error) {
	// var _args GetContactArgs
	// _args.Req = req
	// var _result GetContactResult

	_result := new(tg.ContactData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getContact", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserAddContact(ctx context.Context, req *user.TLUserAddContact) (r *tg.Bool, err error) {
	// var _args AddContactArgs
	// _args.Req = req
	// var _result AddContactResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.addContact", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCheckContact(ctx context.Context, req *user.TLUserCheckContact) (r *tg.Bool, err error) {
	// var _args CheckContactArgs
	// _args.Req = req
	// var _result CheckContactResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.checkContact", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetImportersByPhone(ctx context.Context, req *user.TLUserGetImportersByPhone) (r *user.VectorInputContact, err error) {
	// var _args GetImportersByPhoneArgs
	// _args.Req = req
	// var _result GetImportersByPhoneResult

	_result := new(user.VectorInputContact)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getImportersByPhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserDeleteImportersByPhone(ctx context.Context, req *user.TLUserDeleteImportersByPhone) (r *tg.Bool, err error) {
	// var _args DeleteImportersByPhoneArgs
	// _args.Req = req
	// var _result DeleteImportersByPhoneResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.deleteImportersByPhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserImportContacts(ctx context.Context, req *user.TLUserImportContacts) (r *user.UserImportedContacts, err error) {
	// var _args ImportContactsArgs
	// _args.Req = req
	// var _result ImportContactsResult

	_result := new(user.UserImportedContacts)

	if err = p.c.Call(ctx, "/user.RPCUser/user.importContacts", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetCountryCode(ctx context.Context, req *user.TLUserGetCountryCode) (r *tg.String, err error) {
	// var _args GetCountryCodeArgs
	// _args.Req = req
	// var _result GetCountryCodeResult

	_result := new(tg.String)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getCountryCode", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateAbout(ctx context.Context, req *user.TLUserUpdateAbout) (r *tg.Bool, err error) {
	// var _args UpdateAboutArgs
	// _args.Req = req
	// var _result UpdateAboutResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateAbout", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateFirstAndLastName(ctx context.Context, req *user.TLUserUpdateFirstAndLastName) (r *tg.Bool, err error) {
	// var _args UpdateFirstAndLastNameArgs
	// _args.Req = req
	// var _result UpdateFirstAndLastNameResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateFirstAndLastName", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateVerified(ctx context.Context, req *user.TLUserUpdateVerified) (r *tg.Bool, err error) {
	// var _args UpdateVerifiedArgs
	// _args.Req = req
	// var _result UpdateVerifiedResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateVerified", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateUsername(ctx context.Context, req *user.TLUserUpdateUsername) (r *tg.Bool, err error) {
	// var _args UpdateUsernameArgs
	// _args.Req = req
	// var _result UpdateUsernameResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateUsername", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateProfilePhoto(ctx context.Context, req *user.TLUserUpdateProfilePhoto) (r *tg.Int64, err error) {
	// var _args UpdateProfilePhotoArgs
	// _args.Req = req
	// var _result UpdateProfilePhotoResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateProfilePhoto", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserDeleteProfilePhotos(ctx context.Context, req *user.TLUserDeleteProfilePhotos) (r *tg.Int64, err error) {
	// var _args DeleteProfilePhotosArgs
	// _args.Req = req
	// var _result DeleteProfilePhotosResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/user.RPCUser/user.deleteProfilePhotos", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetProfilePhotos(ctx context.Context, req *user.TLUserGetProfilePhotos) (r *user.VectorLong, err error) {
	// var _args GetProfilePhotosArgs
	// _args.Req = req
	// var _result GetProfilePhotosResult

	_result := new(user.VectorLong)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getProfilePhotos", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetBotCommands(ctx context.Context, req *user.TLUserSetBotCommands) (r *tg.Bool, err error) {
	// var _args SetBotCommandsArgs
	// _args.Req = req
	// var _result SetBotCommandsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setBotCommands", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserIsBot(ctx context.Context, req *user.TLUserIsBot) (r *tg.Bool, err error) {
	// var _args IsBotArgs
	// _args.Req = req
	// var _result IsBotResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.isBot", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetBotInfo(ctx context.Context, req *user.TLUserGetBotInfo) (r *tg.BotInfo, err error) {
	// var _args GetBotInfoArgs
	// _args.Req = req
	// var _result GetBotInfoResult

	_result := new(tg.BotInfo)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getBotInfo", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCheckBots(ctx context.Context, req *user.TLUserCheckBots) (r *user.VectorLong, err error) {
	// var _args CheckBotsArgs
	// _args.Req = req
	// var _result CheckBotsResult

	_result := new(user.VectorLong)

	if err = p.c.Call(ctx, "/user.RPCUser/user.checkBots", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetFullUser(ctx context.Context, req *user.TLUserGetFullUser) (r *tg.UsersUserFull, err error) {
	// var _args GetFullUserArgs
	// _args.Req = req
	// var _result GetFullUserResult

	_result := new(tg.UsersUserFull)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getFullUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateEmojiStatus(ctx context.Context, req *user.TLUserUpdateEmojiStatus) (r *tg.Bool, err error) {
	// var _args UpdateEmojiStatusArgs
	// _args.Req = req
	// var _result UpdateEmojiStatusResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateEmojiStatus", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetUserDataById(ctx context.Context, req *user.TLUserGetUserDataById) (r *tg.UserData, err error) {
	// var _args GetUserDataByIdArgs
	// _args.Req = req
	// var _result GetUserDataByIdResult

	_result := new(tg.UserData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getUserDataById", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetUserDataListByIdList(ctx context.Context, req *user.TLUserGetUserDataListByIdList) (r *user.VectorUserData, err error) {
	// var _args GetUserDataListByIdListArgs
	// _args.Req = req
	// var _result GetUserDataListByIdListResult

	_result := new(user.VectorUserData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getUserDataListByIdList", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetUserDataByToken(ctx context.Context, req *user.TLUserGetUserDataByToken) (r *tg.UserData, err error) {
	// var _args GetUserDataByTokenArgs
	// _args.Req = req
	// var _result GetUserDataByTokenResult

	_result := new(tg.UserData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getUserDataByToken", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSearch(ctx context.Context, req *user.TLUserSearch) (r *user.UsersFound, err error) {
	// var _args SearchArgs
	// _args.Req = req
	// var _result SearchResult

	_result := new(user.UsersFound)

	if err = p.c.Call(ctx, "/user.RPCUser/user.search", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateBotData(ctx context.Context, req *user.TLUserUpdateBotData) (r *tg.Bool, err error) {
	// var _args UpdateBotDataArgs
	// _args.Req = req
	// var _result UpdateBotDataResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateBotData", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetImmutableUserV2(ctx context.Context, req *user.TLUserGetImmutableUserV2) (r *tg.ImmutableUser, err error) {
	// var _args GetImmutableUserV2Args
	// _args.Req = req
	// var _result GetImmutableUserV2Result

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getImmutableUserV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetMutableUsersV2(ctx context.Context, req *user.TLUserGetMutableUsersV2) (r *tg.MutableUsers, err error) {
	// var _args GetMutableUsersV2Args
	// _args.Req = req
	// var _result GetMutableUsersV2Result

	_result := new(tg.MutableUsers)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getMutableUsersV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserCreateNewTestUser(ctx context.Context, req *user.TLUserCreateNewTestUser) (r *tg.ImmutableUser, err error) {
	// var _args CreateNewTestUserArgs
	// _args.Req = req
	// var _result CreateNewTestUserResult

	_result := new(tg.ImmutableUser)

	if err = p.c.Call(ctx, "/user.RPCUser/user.createNewTestUser", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserEditCloseFriends(ctx context.Context, req *user.TLUserEditCloseFriends) (r *tg.Bool, err error) {
	// var _args EditCloseFriendsArgs
	// _args.Req = req
	// var _result EditCloseFriendsResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.editCloseFriends", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetStoriesMaxId(ctx context.Context, req *user.TLUserSetStoriesMaxId) (r *tg.Bool, err error) {
	// var _args SetStoriesMaxIdArgs
	// _args.Req = req
	// var _result SetStoriesMaxIdResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setStoriesMaxId", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetColor(ctx context.Context, req *user.TLUserSetColor) (r *tg.Bool, err error) {
	// var _args SetColorArgs
	// _args.Req = req
	// var _result SetColorResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setColor", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdateBirthday(ctx context.Context, req *user.TLUserUpdateBirthday) (r *tg.Bool, err error) {
	// var _args UpdateBirthdayArgs
	// _args.Req = req
	// var _result UpdateBirthdayResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updateBirthday", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetBirthdays(ctx context.Context, req *user.TLUserGetBirthdays) (r *user.VectorContactBirthday, err error) {
	// var _args GetBirthdaysArgs
	// _args.Req = req
	// var _result GetBirthdaysResult

	_result := new(user.VectorContactBirthday)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getBirthdays", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetStoriesHidden(ctx context.Context, req *user.TLUserSetStoriesHidden) (r *tg.Bool, err error) {
	// var _args SetStoriesHiddenArgs
	// _args.Req = req
	// var _result SetStoriesHiddenResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setStoriesHidden", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdatePersonalChannel(ctx context.Context, req *user.TLUserUpdatePersonalChannel) (r *tg.Bool, err error) {
	// var _args UpdatePersonalChannelArgs
	// _args.Req = req
	// var _result UpdatePersonalChannelResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updatePersonalChannel", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetUserIdByPhone(ctx context.Context, req *user.TLUserGetUserIdByPhone) (r *tg.Int64, err error) {
	// var _args GetUserIdByPhoneArgs
	// _args.Req = req
	// var _result GetUserIdByPhoneResult

	_result := new(tg.Int64)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getUserIdByPhone", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserSetAuthorizationTTL(ctx context.Context, req *user.TLUserSetAuthorizationTTL) (r *tg.Bool, err error) {
	// var _args SetAuthorizationTTLArgs
	// _args.Req = req
	// var _result SetAuthorizationTTLResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.setAuthorizationTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetAuthorizationTTL(ctx context.Context, req *user.TLUserGetAuthorizationTTL) (r *tg.AccountDaysTTL, err error) {
	// var _args GetAuthorizationTTLArgs
	// _args.Req = req
	// var _result GetAuthorizationTTLResult

	_result := new(tg.AccountDaysTTL)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getAuthorizationTTL", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserUpdatePremium(ctx context.Context, req *user.TLUserUpdatePremium) (r *tg.Bool, err error) {
	// var _args UpdatePremiumArgs
	// _args.Req = req
	// var _result UpdatePremiumResult

	_result := new(tg.Bool)

	if err = p.c.Call(ctx, "/user.RPCUser/user.updatePremium", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}

func (p *kClient) UserGetBotInfoV2(ctx context.Context, req *user.TLUserGetBotInfoV2) (r *user.BotInfoData, err error) {
	// var _args GetBotInfoV2Args
	// _args.Req = req
	// var _result GetBotInfoV2Result

	_result := new(user.BotInfoData)

	if err = p.c.Call(ctx, "/user.RPCUser/user.getBotInfoV2", req, _result); err != nil {
		return
	}

	// return _result.GetSuccess(), nil
	return _result, nil
}
