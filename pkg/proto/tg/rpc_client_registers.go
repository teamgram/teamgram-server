/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2025-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package tg

import (
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func init() {
	// RPCAuthorization
	iface.RegisterRPCContextTuple("TLAuthSendCode", "/tg.RPCAuthorization/auth.sendCode", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAuthSignUp", "/tg.RPCAuthorization/auth.signUp", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthSignIn", "/tg.RPCAuthorization/auth.signIn", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthLogOut", "/tg.RPCAuthorization/auth.logOut", func() interface{} { return new(AuthLoggedOut) })
	iface.RegisterRPCContextTuple("TLAuthResetAuthorizations", "/tg.RPCAuthorization/auth.resetAuthorizations", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthExportAuthorization", "/tg.RPCAuthorization/auth.exportAuthorization", func() interface{} { return new(AuthExportedAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthImportAuthorization", "/tg.RPCAuthorization/auth.importAuthorization", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthBindTempAuthKey", "/tg.RPCAuthorization/auth.bindTempAuthKey", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthImportBotAuthorization", "/tg.RPCAuthorization/auth.importBotAuthorization", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthCheckPassword", "/tg.RPCAuthorization/auth.checkPassword", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthRequestPasswordRecovery", "/tg.RPCAuthorization/auth.requestPasswordRecovery", func() interface{} { return new(AuthPasswordRecovery) })
	iface.RegisterRPCContextTuple("TLAuthRecoverPassword", "/tg.RPCAuthorization/auth.recoverPassword", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthResendCode", "/tg.RPCAuthorization/auth.resendCode", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAuthCancelCode", "/tg.RPCAuthorization/auth.cancelCode", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthDropTempAuthKeys", "/tg.RPCAuthorization/auth.dropTempAuthKeys", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthCheckRecoveryPassword", "/tg.RPCAuthorization/auth.checkRecoveryPassword", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthImportWebTokenAuthorization", "/tg.RPCAuthorization/auth.importWebTokenAuthorization", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAuthRequestFirebaseSms", "/tg.RPCAuthorization/auth.requestFirebaseSms", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthResetLoginEmail", "/tg.RPCAuthorization/auth.resetLoginEmail", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAuthReportMissingCode", "/tg.RPCAuthorization/auth.reportMissingCode", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthCheckPaidAuth", "/tg.RPCAuthorization/auth.checkPaidAuth", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAccountSendVerifyEmailCode", "/tg.RPCAuthorization/account.sendVerifyEmailCode", func() interface{} { return new(AccountSentEmailCode) })
	iface.RegisterRPCContextTuple("TLAccountVerifyEmail", "/tg.RPCAuthorization/account.verifyEmail", func() interface{} { return new(AccountEmailVerified) })
	iface.RegisterRPCContextTuple("TLAccountResetPassword", "/tg.RPCAuthorization/account.resetPassword", func() interface{} { return new(AccountResetPasswordResult) })
	iface.RegisterRPCContextTuple("TLAccountSetAuthorizationTTL", "/tg.RPCAuthorization/account.setAuthorizationTTL", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountChangeAuthorizationSettings", "/tg.RPCAuthorization/account.changeAuthorizationSettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountInvalidateSignInCodes", "/tg.RPCAuthorization/account.invalidateSignInCodes", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAuthToggleBan", "/tg.RPCAuthorization/auth.toggleBan", func() interface{} { return new(PredefinedUser) })

	// RPCQrCode
	iface.RegisterRPCContextTuple("TLAuthExportLoginToken", "/tg.RPCQrCode/auth.exportLoginToken", func() interface{} { return new(AuthLoginToken) })
	iface.RegisterRPCContextTuple("TLAuthImportLoginToken", "/tg.RPCQrCode/auth.importLoginToken", func() interface{} { return new(AuthLoginToken) })
	iface.RegisterRPCContextTuple("TLAuthAcceptLoginToken", "/tg.RPCQrCode/auth.acceptLoginToken", func() interface{} { return new(Authorization) })

	// RPCPasskey
	iface.RegisterRPCContextTuple("TLAuthInitPasskeyLogin", "/tg.RPCPasskey/auth.initPasskeyLogin", func() interface{} { return new(AuthPasskeyLoginOptions) })
	iface.RegisterRPCContextTuple("TLAuthFinishPasskeyLogin", "/tg.RPCPasskey/auth.finishPasskeyLogin", func() interface{} { return new(AuthAuthorization) })
	iface.RegisterRPCContextTuple("TLAccountInitPasskeyRegistration", "/tg.RPCPasskey/account.initPasskeyRegistration", func() interface{} { return new(AccountPasskeyRegistrationOptions) })
	iface.RegisterRPCContextTuple("TLAccountRegisterPasskey", "/tg.RPCPasskey/account.registerPasskey", func() interface{} { return new(Passkey) })
	iface.RegisterRPCContextTuple("TLAccountGetPasskeys", "/tg.RPCPasskey/account.getPasskeys", func() interface{} { return new(AccountPasskeys) })
	iface.RegisterRPCContextTuple("TLAccountDeletePasskey", "/tg.RPCPasskey/account.deletePasskey", func() interface{} { return new(Bool) })

	// RPCNotification
	iface.RegisterRPCContextTuple("TLAccountRegisterDevice", "/tg.RPCNotification/account.registerDevice", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUnregisterDevice", "/tg.RPCNotification/account.unregisterDevice", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateNotifySettings", "/tg.RPCNotification/account.updateNotifySettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetNotifySettings", "/tg.RPCNotification/account.getNotifySettings", func() interface{} { return new(PeerNotifySettings) })
	iface.RegisterRPCContextTuple("TLAccountResetNotifySettings", "/tg.RPCNotification/account.resetNotifySettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateDeviceLocked", "/tg.RPCNotification/account.updateDeviceLocked", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetNotifyExceptions", "/tg.RPCNotification/account.getNotifyExceptions", func() interface{} { return new(Updates) })

	// RPCUserChannelProfiles
	iface.RegisterRPCContextTuple("TLAccountUpdateProfile", "/tg.RPCUserChannelProfiles/account.updateProfile", func() interface{} { return new(User) })
	iface.RegisterRPCContextTuple("TLAccountUpdateStatus", "/tg.RPCUserChannelProfiles/account.updateStatus", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateBirthday", "/tg.RPCUserChannelProfiles/account.updateBirthday", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdatePersonalChannel", "/tg.RPCUserChannelProfiles/account.updatePersonalChannel", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSetMainProfileTab", "/tg.RPCUserChannelProfiles/account.setMainProfileTab", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSaveMusic", "/tg.RPCUserChannelProfiles/account.saveMusic", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetSavedMusicIds", "/tg.RPCUserChannelProfiles/account.getSavedMusicIds", func() interface{} { return new(AccountSavedMusicIds) })
	iface.RegisterRPCContextTuple("TLUsersGetSavedMusic", "/tg.RPCUserChannelProfiles/users.getSavedMusic", func() interface{} { return new(UsersSavedMusic) })
	iface.RegisterRPCContextTuple("TLUsersGetSavedMusicByID", "/tg.RPCUserChannelProfiles/users.getSavedMusicByID", func() interface{} { return new(UsersSavedMusic) })
	iface.RegisterRPCContextTuple("TLUsersSuggestBirthday", "/tg.RPCUserChannelProfiles/users.suggestBirthday", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLContactsGetBirthdays", "/tg.RPCUserChannelProfiles/contacts.getBirthdays", func() interface{} { return new(ContactsContactBirthdays) })
	iface.RegisterRPCContextTuple("TLPhotosUpdateProfilePhoto", "/tg.RPCUserChannelProfiles/photos.updateProfilePhoto", func() interface{} { return new(PhotosPhoto) })
	iface.RegisterRPCContextTuple("TLPhotosUploadProfilePhoto", "/tg.RPCUserChannelProfiles/photos.uploadProfilePhoto", func() interface{} { return new(PhotosPhoto) })
	iface.RegisterRPCContextTuple("TLPhotosDeletePhotos", "/tg.RPCUserChannelProfiles/photos.deletePhotos", func() interface{} { return new(VectorLong) })
	iface.RegisterRPCContextTuple("TLPhotosGetUserPhotos", "/tg.RPCUserChannelProfiles/photos.getUserPhotos", func() interface{} { return new(PhotosPhotos) })
	iface.RegisterRPCContextTuple("TLPhotosUploadContactProfilePhoto", "/tg.RPCUserChannelProfiles/photos.uploadContactProfilePhoto", func() interface{} { return new(PhotosPhoto) })
	iface.RegisterRPCContextTuple("TLChannelsSetMainProfileTab", "/tg.RPCUserChannelProfiles/channels.setMainProfileTab", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateVerified", "/tg.RPCUserChannelProfiles/account.updateVerified", func() interface{} { return new(User) })

	// RPCWallpapers
	iface.RegisterRPCContextTuple("TLAccountGetWallPapers", "/tg.RPCWallpapers/account.getWallPapers", func() interface{} { return new(AccountWallPapers) })
	iface.RegisterRPCContextTuple("TLAccountGetWallPaper", "/tg.RPCWallpapers/account.getWallPaper", func() interface{} { return new(WallPaper) })
	iface.RegisterRPCContextTuple("TLAccountUploadWallPaper", "/tg.RPCWallpapers/account.uploadWallPaper", func() interface{} { return new(WallPaper) })
	iface.RegisterRPCContextTuple("TLAccountSaveWallPaper", "/tg.RPCWallpapers/account.saveWallPaper", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountInstallWallPaper", "/tg.RPCWallpapers/account.installWallPaper", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountResetWallPapers", "/tg.RPCWallpapers/account.resetWallPapers", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetMultiWallPapers", "/tg.RPCWallpapers/account.getMultiWallPapers", func() interface{} { return new(VectorWallPaper) })
	iface.RegisterRPCContextTuple("TLMessagesSetChatWallPaper", "/tg.RPCWallpapers/messages.setChatWallPaper", func() interface{} { return new(Updates) })

	// RPCReports
	iface.RegisterRPCContextTuple("TLAccountReportPeer", "/tg.RPCReports/account.reportPeer", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountReportProfilePhoto", "/tg.RPCReports/account.reportProfilePhoto", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReportSpam", "/tg.RPCReports/messages.reportSpam", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReport", "/tg.RPCReports/messages.report", func() interface{} { return new(ReportResult) })
	iface.RegisterRPCContextTuple("TLMessagesReportEncryptedSpam", "/tg.RPCReports/messages.reportEncryptedSpam", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsReportSpam", "/tg.RPCReports/channels.reportSpam", func() interface{} { return new(Bool) })

	// RPCUsernames
	iface.RegisterRPCContextTuple("TLAccountCheckUsername", "/tg.RPCUsernames/account.checkUsername", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateUsername", "/tg.RPCUsernames/account.updateUsername", func() interface{} { return new(User) })
	iface.RegisterRPCContextTuple("TLContactsResolveUsername", "/tg.RPCUsernames/contacts.resolveUsername", func() interface{} { return new(ContactsResolvedPeer) })
	iface.RegisterRPCContextTuple("TLChannelsCheckUsername", "/tg.RPCUsernames/channels.checkUsername", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsUpdateUsername", "/tg.RPCUsernames/channels.updateUsername", func() interface{} { return new(Bool) })

	// RPCPrivacySettings
	iface.RegisterRPCContextTuple("TLAccountGetPrivacy", "/tg.RPCPrivacySettings/account.getPrivacy", func() interface{} { return new(AccountPrivacyRules) })
	iface.RegisterRPCContextTuple("TLAccountSetPrivacy", "/tg.RPCPrivacySettings/account.setPrivacy", func() interface{} { return new(AccountPrivacyRules) })
	iface.RegisterRPCContextTuple("TLAccountGetGlobalPrivacySettings", "/tg.RPCPrivacySettings/account.getGlobalPrivacySettings", func() interface{} { return new(GlobalPrivacySettings) })
	iface.RegisterRPCContextTuple("TLAccountSetGlobalPrivacySettings", "/tg.RPCPrivacySettings/account.setGlobalPrivacySettings", func() interface{} { return new(GlobalPrivacySettings) })
	iface.RegisterRPCContextTuple("TLUsersGetRequirementsToContact", "/tg.RPCPrivacySettings/users.getRequirementsToContact", func() interface{} { return new(VectorRequirementToContact) })
	iface.RegisterRPCContextTuple("TLMessagesSetDefaultHistoryTTL", "/tg.RPCPrivacySettings/messages.setDefaultHistoryTTL", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetDefaultHistoryTTL", "/tg.RPCPrivacySettings/messages.getDefaultHistoryTTL", func() interface{} { return new(DefaultHistoryTTL) })

	// RPCAccount
	iface.RegisterRPCContextTuple("TLAccountDeleteAccount", "/tg.RPCAccount/account.deleteAccount", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetAccountTTL", "/tg.RPCAccount/account.getAccountTTL", func() interface{} { return new(AccountDaysTTL) })
	iface.RegisterRPCContextTuple("TLAccountSetAccountTTL", "/tg.RPCAccount/account.setAccountTTL", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSendChangePhoneCode", "/tg.RPCAccount/account.sendChangePhoneCode", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAccountChangePhone", "/tg.RPCAccount/account.changePhone", func() interface{} { return new(User) })
	iface.RegisterRPCContextTuple("TLAccountResetAuthorization", "/tg.RPCAccount/account.resetAuthorization", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSendConfirmPhoneCode", "/tg.RPCAccount/account.sendConfirmPhoneCode", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAccountConfirmPhone", "/tg.RPCAccount/account.confirmPhone", func() interface{} { return new(Bool) })

	// RPCPassport
	iface.RegisterRPCContextTuple("TLAccountGetAuthorizations", "/tg.RPCPassport/account.getAuthorizations", func() interface{} { return new(AccountAuthorizations) })
	iface.RegisterRPCContextTuple("TLAccountGetAllSecureValues", "/tg.RPCPassport/account.getAllSecureValues", func() interface{} { return new(VectorSecureValue) })
	iface.RegisterRPCContextTuple("TLAccountGetSecureValue", "/tg.RPCPassport/account.getSecureValue", func() interface{} { return new(VectorSecureValue) })
	iface.RegisterRPCContextTuple("TLAccountSaveSecureValue", "/tg.RPCPassport/account.saveSecureValue", func() interface{} { return new(SecureValue) })
	iface.RegisterRPCContextTuple("TLAccountDeleteSecureValue", "/tg.RPCPassport/account.deleteSecureValue", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetAuthorizationForm", "/tg.RPCPassport/account.getAuthorizationForm", func() interface{} { return new(AccountAuthorizationForm) })
	iface.RegisterRPCContextTuple("TLAccountAcceptAuthorization", "/tg.RPCPassport/account.acceptAuthorization", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSendVerifyPhoneCode", "/tg.RPCPassport/account.sendVerifyPhoneCode", func() interface{} { return new(AuthSentCode) })
	iface.RegisterRPCContextTuple("TLAccountVerifyPhone", "/tg.RPCPassport/account.verifyPhone", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLUsersSetSecureValueErrors", "/tg.RPCPassport/users.setSecureValueErrors", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLHelpGetPassportConfig", "/tg.RPCPassport/help.getPassportConfig", func() interface{} { return new(HelpPassportConfig) })

	// RPCTwoFa
	iface.RegisterRPCContextTuple("TLAccountGetPassword", "/tg.RPCTwoFa/account.getPassword", func() interface{} { return new(AccountPassword) })
	iface.RegisterRPCContextTuple("TLAccountGetPasswordSettings", "/tg.RPCTwoFa/account.getPasswordSettings", func() interface{} { return new(AccountPasswordSettings) })
	iface.RegisterRPCContextTuple("TLAccountUpdatePasswordSettings", "/tg.RPCTwoFa/account.updatePasswordSettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountConfirmPasswordEmail", "/tg.RPCTwoFa/account.confirmPasswordEmail", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountResendPasswordEmail", "/tg.RPCTwoFa/account.resendPasswordEmail", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountCancelPasswordEmail", "/tg.RPCTwoFa/account.cancelPasswordEmail", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountDeclinePasswordReset", "/tg.RPCTwoFa/account.declinePasswordReset", func() interface{} { return new(Bool) })

	// RPCPayments
	iface.RegisterRPCContextTuple("TLAccountGetTmpPassword", "/tg.RPCPayments/account.getTmpPassword", func() interface{} { return new(AccountTmpPassword) })
	iface.RegisterRPCContextTuple("TLMessagesSetBotShippingResults", "/tg.RPCPayments/messages.setBotShippingResults", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSetBotPrecheckoutResults", "/tg.RPCPayments/messages.setBotPrecheckoutResults", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsGetPaymentForm", "/tg.RPCPayments/payments.getPaymentForm", func() interface{} { return new(PaymentsPaymentForm) })
	iface.RegisterRPCContextTuple("TLPaymentsGetPaymentReceipt", "/tg.RPCPayments/payments.getPaymentReceipt", func() interface{} { return new(PaymentsPaymentReceipt) })
	iface.RegisterRPCContextTuple("TLPaymentsValidateRequestedInfo", "/tg.RPCPayments/payments.validateRequestedInfo", func() interface{} { return new(PaymentsValidatedRequestedInfo) })
	iface.RegisterRPCContextTuple("TLPaymentsSendPaymentForm", "/tg.RPCPayments/payments.sendPaymentForm", func() interface{} { return new(PaymentsPaymentResult) })
	iface.RegisterRPCContextTuple("TLPaymentsGetSavedInfo", "/tg.RPCPayments/payments.getSavedInfo", func() interface{} { return new(PaymentsSavedInfo) })
	iface.RegisterRPCContextTuple("TLPaymentsClearSavedInfo", "/tg.RPCPayments/payments.clearSavedInfo", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsGetBankCardData", "/tg.RPCPayments/payments.getBankCardData", func() interface{} { return new(PaymentsBankCardData) })
	iface.RegisterRPCContextTuple("TLPaymentsExportInvoice", "/tg.RPCPayments/payments.exportInvoice", func() interface{} { return new(PaymentsExportedInvoice) })

	// RPCSeamless
	iface.RegisterRPCContextTuple("TLAccountGetWebAuthorizations", "/tg.RPCSeamless/account.getWebAuthorizations", func() interface{} { return new(AccountWebAuthorizations) })
	iface.RegisterRPCContextTuple("TLAccountResetWebAuthorization", "/tg.RPCSeamless/account.resetWebAuthorization", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountResetWebAuthorizations", "/tg.RPCSeamless/account.resetWebAuthorizations", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesRequestUrlAuth", "/tg.RPCSeamless/messages.requestUrlAuth", func() interface{} { return new(UrlAuthResult) })
	iface.RegisterRPCContextTuple("TLMessagesAcceptUrlAuth", "/tg.RPCSeamless/messages.acceptUrlAuth", func() interface{} { return new(UrlAuthResult) })
	iface.RegisterRPCContextTuple("TLMessagesDeclineUrlAuth", "/tg.RPCSeamless/messages.declineUrlAuth", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesCheckUrlAuthMatchCode", "/tg.RPCSeamless/messages.checkUrlAuthMatchCode", func() interface{} { return new(Bool) })

	// RPCTakeout
	iface.RegisterRPCContextTuple("TLAccountInitTakeoutSession", "/tg.RPCTakeout/account.initTakeoutSession", func() interface{} { return new(AccountTakeout) })
	iface.RegisterRPCContextTuple("TLAccountFinishTakeoutSession", "/tg.RPCTakeout/account.finishTakeoutSession", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetSplitRanges", "/tg.RPCTakeout/messages.getSplitRanges", func() interface{} { return new(VectorMessageRange) })
	iface.RegisterRPCContextTuple("TLChannelsGetLeftChannels", "/tg.RPCTakeout/channels.getLeftChannels", func() interface{} { return new(MessagesChats) })

	// RPCContacts
	iface.RegisterRPCContextTuple("TLAccountGetContactSignUpNotification", "/tg.RPCContacts/account.getContactSignUpNotification", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountSetContactSignUpNotification", "/tg.RPCContacts/account.setContactSignUpNotification", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsGetContactIDs", "/tg.RPCContacts/contacts.getContactIDs", func() interface{} { return new(VectorInt) })
	iface.RegisterRPCContextTuple("TLContactsGetStatuses", "/tg.RPCContacts/contacts.getStatuses", func() interface{} { return new(VectorContactStatus) })
	iface.RegisterRPCContextTuple("TLContactsGetContacts", "/tg.RPCContacts/contacts.getContacts", func() interface{} { return new(ContactsContacts) })
	iface.RegisterRPCContextTuple("TLContactsImportContacts", "/tg.RPCContacts/contacts.importContacts", func() interface{} { return new(ContactsImportedContacts) })
	iface.RegisterRPCContextTuple("TLContactsDeleteContacts", "/tg.RPCContacts/contacts.deleteContacts", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLContactsDeleteByPhones", "/tg.RPCContacts/contacts.deleteByPhones", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsBlock", "/tg.RPCContacts/contacts.block", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsUnblock", "/tg.RPCContacts/contacts.unblock", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsGetBlocked", "/tg.RPCContacts/contacts.getBlocked", func() interface{} { return new(ContactsBlocked) })
	iface.RegisterRPCContextTuple("TLContactsSearch", "/tg.RPCContacts/contacts.search", func() interface{} { return new(ContactsFound) })
	iface.RegisterRPCContextTuple("TLContactsGetTopPeers", "/tg.RPCContacts/contacts.getTopPeers", func() interface{} { return new(ContactsTopPeers) })
	iface.RegisterRPCContextTuple("TLContactsResetTopPeerRating", "/tg.RPCContacts/contacts.resetTopPeerRating", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsResetSaved", "/tg.RPCContacts/contacts.resetSaved", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsGetSaved", "/tg.RPCContacts/contacts.getSaved", func() interface{} { return new(VectorSavedContact) })
	iface.RegisterRPCContextTuple("TLContactsToggleTopPeers", "/tg.RPCContacts/contacts.toggleTopPeers", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsAddContact", "/tg.RPCContacts/contacts.addContact", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLContactsAcceptContact", "/tg.RPCContacts/contacts.acceptContact", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLContactsGetLocated", "/tg.RPCContacts/contacts.getLocated", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLContactsEditCloseFriends", "/tg.RPCContacts/contacts.editCloseFriends", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsSetBlocked", "/tg.RPCContacts/contacts.setBlocked", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsUpdateContactNote", "/tg.RPCContacts/contacts.updateContactNote", func() interface{} { return new(Bool) })

	// RPCAutoDownload
	iface.RegisterRPCContextTuple("TLAccountGetAutoDownloadSettings", "/tg.RPCAutoDownload/account.getAutoDownloadSettings", func() interface{} { return new(AccountAutoDownloadSettings) })
	iface.RegisterRPCContextTuple("TLAccountSaveAutoDownloadSettings", "/tg.RPCAutoDownload/account.saveAutoDownloadSettings", func() interface{} { return new(Bool) })

	// RPCThemes
	iface.RegisterRPCContextTuple("TLAccountUploadTheme", "/tg.RPCThemes/account.uploadTheme", func() interface{} { return new(Document) })
	iface.RegisterRPCContextTuple("TLAccountCreateTheme", "/tg.RPCThemes/account.createTheme", func() interface{} { return new(Theme) })
	iface.RegisterRPCContextTuple("TLAccountUpdateTheme", "/tg.RPCThemes/account.updateTheme", func() interface{} { return new(Theme) })
	iface.RegisterRPCContextTuple("TLAccountSaveTheme", "/tg.RPCThemes/account.saveTheme", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountInstallTheme", "/tg.RPCThemes/account.installTheme", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetTheme", "/tg.RPCThemes/account.getTheme", func() interface{} { return new(Theme) })
	iface.RegisterRPCContextTuple("TLAccountGetThemes", "/tg.RPCThemes/account.getThemes", func() interface{} { return new(AccountThemes) })
	iface.RegisterRPCContextTuple("TLAccountGetChatThemes", "/tg.RPCThemes/account.getChatThemes", func() interface{} { return new(AccountThemes) })
	iface.RegisterRPCContextTuple("TLAccountGetUniqueGiftChatThemes", "/tg.RPCThemes/account.getUniqueGiftChatThemes", func() interface{} { return new(AccountChatThemes) })
	iface.RegisterRPCContextTuple("TLMessagesSetChatTheme", "/tg.RPCThemes/messages.setChatTheme", func() interface{} { return new(Updates) })

	// RPCNsfw
	iface.RegisterRPCContextTuple("TLAccountSetContentSettings", "/tg.RPCNsfw/account.setContentSettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetContentSettings", "/tg.RPCNsfw/account.getContentSettings", func() interface{} { return new(AccountContentSettings) })

	// RPCRingtone
	iface.RegisterRPCContextTuple("TLAccountGetSavedRingtones", "/tg.RPCRingtone/account.getSavedRingtones", func() interface{} { return new(AccountSavedRingtones) })
	iface.RegisterRPCContextTuple("TLAccountSaveRingtone", "/tg.RPCRingtone/account.saveRingtone", func() interface{} { return new(AccountSavedRingtone) })
	iface.RegisterRPCContextTuple("TLAccountUploadRingtone", "/tg.RPCRingtone/account.uploadRingtone", func() interface{} { return new(Document) })

	// RPCEmojiStatus
	iface.RegisterRPCContextTuple("TLAccountUpdateEmojiStatus", "/tg.RPCEmojiStatus/account.updateEmojiStatus", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetDefaultEmojiStatuses", "/tg.RPCEmojiStatus/account.getDefaultEmojiStatuses", func() interface{} { return new(AccountEmojiStatuses) })
	iface.RegisterRPCContextTuple("TLAccountGetRecentEmojiStatuses", "/tg.RPCEmojiStatus/account.getRecentEmojiStatuses", func() interface{} { return new(AccountEmojiStatuses) })
	iface.RegisterRPCContextTuple("TLAccountClearRecentEmojiStatuses", "/tg.RPCEmojiStatus/account.clearRecentEmojiStatuses", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetChannelDefaultEmojiStatuses", "/tg.RPCEmojiStatus/account.getChannelDefaultEmojiStatuses", func() interface{} { return new(AccountEmojiStatuses) })
	iface.RegisterRPCContextTuple("TLAccountGetChannelRestrictedStatusEmojis", "/tg.RPCEmojiStatus/account.getChannelRestrictedStatusEmojis", func() interface{} { return new(EmojiList) })
	iface.RegisterRPCContextTuple("TLAccountGetCollectibleEmojiStatuses", "/tg.RPCEmojiStatus/account.getCollectibleEmojiStatuses", func() interface{} { return new(AccountEmojiStatuses) })
	iface.RegisterRPCContextTuple("TLChannelsUpdateEmojiStatus", "/tg.RPCEmojiStatus/channels.updateEmojiStatus", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLBotsUpdateUserEmojiStatus", "/tg.RPCEmojiStatus/bots.updateUserEmojiStatus", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsToggleUserEmojiStatusPermission", "/tg.RPCEmojiStatus/bots.toggleUserEmojiStatusPermission", func() interface{} { return new(Bool) })

	// RPCFragment
	iface.RegisterRPCContextTuple("TLAccountReorderUsernames", "/tg.RPCFragment/account.reorderUsernames", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountToggleUsername", "/tg.RPCFragment/account.toggleUsername", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsReorderUsernames", "/tg.RPCFragment/channels.reorderUsernames", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsToggleUsername", "/tg.RPCFragment/channels.toggleUsername", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsDeactivateAllUsernames", "/tg.RPCFragment/channels.deactivateAllUsernames", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsReorderUsernames", "/tg.RPCFragment/bots.reorderUsernames", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsToggleUsername", "/tg.RPCFragment/bots.toggleUsername", func() interface{} { return new(Bool) })

	// RPCCustomEmojis
	iface.RegisterRPCContextTuple("TLAccountGetDefaultProfilePhotoEmojis", "/tg.RPCCustomEmojis/account.getDefaultProfilePhotoEmojis", func() interface{} { return new(EmojiList) })
	iface.RegisterRPCContextTuple("TLAccountGetDefaultGroupPhotoEmojis", "/tg.RPCCustomEmojis/account.getDefaultGroupPhotoEmojis", func() interface{} { return new(EmojiList) })
	iface.RegisterRPCContextTuple("TLMessagesGetCustomEmojiDocuments", "/tg.RPCCustomEmojis/messages.getCustomEmojiDocuments", func() interface{} { return new(VectorDocument) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiStickers", "/tg.RPCCustomEmojis/messages.getEmojiStickers", func() interface{} { return new(MessagesAllStickers) })
	iface.RegisterRPCContextTuple("TLMessagesGetFeaturedEmojiStickers", "/tg.RPCCustomEmojis/messages.getFeaturedEmojiStickers", func() interface{} { return new(MessagesFeaturedStickers) })
	iface.RegisterRPCContextTuple("TLMessagesSearchCustomEmoji", "/tg.RPCCustomEmojis/messages.searchCustomEmoji", func() interface{} { return new(EmojiList) })

	// RPCAutosave
	iface.RegisterRPCContextTuple("TLAccountGetAutoSaveSettings", "/tg.RPCAutosave/account.getAutoSaveSettings", func() interface{} { return new(AccountAutoSaveSettings) })
	iface.RegisterRPCContextTuple("TLAccountSaveAutoSaveSettings", "/tg.RPCAutosave/account.saveAutoSaveSettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountDeleteAutoSaveExceptions", "/tg.RPCAutosave/account.deleteAutoSaveExceptions", func() interface{} { return new(Bool) })

	// RPCAccentColors
	iface.RegisterRPCContextTuple("TLAccountUpdateColor", "/tg.RPCAccentColors/account.updateColor", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetDefaultBackgroundEmojis", "/tg.RPCAccentColors/account.getDefaultBackgroundEmojis", func() interface{} { return new(EmojiList) })
	iface.RegisterRPCContextTuple("TLHelpGetPeerColors", "/tg.RPCAccentColors/help.getPeerColors", func() interface{} { return new(HelpPeerColors) })
	iface.RegisterRPCContextTuple("TLHelpGetPeerProfileColors", "/tg.RPCAccentColors/help.getPeerProfileColors", func() interface{} { return new(HelpPeerColors) })
	iface.RegisterRPCContextTuple("TLChannelsUpdateColor", "/tg.RPCAccentColors/channels.updateColor", func() interface{} { return new(Updates) })

	// RPCBusinessOpeningHours
	iface.RegisterRPCContextTuple("TLAccountUpdateBusinessWorkHours", "/tg.RPCBusinessOpeningHours/account.updateBusinessWorkHours", func() interface{} { return new(Bool) })

	// RPCBusinessLocation
	iface.RegisterRPCContextTuple("TLAccountUpdateBusinessLocation", "/tg.RPCBusinessLocation/account.updateBusinessLocation", func() interface{} { return new(Bool) })

	// RPCBusinessGreeting
	iface.RegisterRPCContextTuple("TLAccountUpdateBusinessGreetingMessage", "/tg.RPCBusinessGreeting/account.updateBusinessGreetingMessage", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountUpdateBusinessAwayMessage", "/tg.RPCBusinessGreeting/account.updateBusinessAwayMessage", func() interface{} { return new(Bool) })

	// RPCBusinessConnectedBots
	iface.RegisterRPCContextTuple("TLAccountUpdateConnectedBot", "/tg.RPCBusinessConnectedBots/account.updateConnectedBot", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLAccountGetConnectedBots", "/tg.RPCBusinessConnectedBots/account.getConnectedBots", func() interface{} { return new(AccountConnectedBots) })
	iface.RegisterRPCContextTuple("TLAccountGetBotBusinessConnection", "/tg.RPCBusinessConnectedBots/account.getBotBusinessConnection", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLAccountToggleConnectedBotPaused", "/tg.RPCBusinessConnectedBots/account.toggleConnectedBotPaused", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountDisablePeerConnectedBot", "/tg.RPCBusinessConnectedBots/account.disablePeerConnectedBot", func() interface{} { return new(Bool) })

	// RPCBusinessIntro
	iface.RegisterRPCContextTuple("TLAccountUpdateBusinessIntro", "/tg.RPCBusinessIntro/account.updateBusinessIntro", func() interface{} { return new(Bool) })

	// RPCBusinessChatLinks
	iface.RegisterRPCContextTuple("TLAccountCreateBusinessChatLink", "/tg.RPCBusinessChatLinks/account.createBusinessChatLink", func() interface{} { return new(BusinessChatLink) })
	iface.RegisterRPCContextTuple("TLAccountEditBusinessChatLink", "/tg.RPCBusinessChatLinks/account.editBusinessChatLink", func() interface{} { return new(BusinessChatLink) })
	iface.RegisterRPCContextTuple("TLAccountDeleteBusinessChatLink", "/tg.RPCBusinessChatLinks/account.deleteBusinessChatLink", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLAccountGetBusinessChatLinks", "/tg.RPCBusinessChatLinks/account.getBusinessChatLinks", func() interface{} { return new(AccountBusinessChatLinks) })
	iface.RegisterRPCContextTuple("TLAccountResolveBusinessChatLink", "/tg.RPCBusinessChatLinks/account.resolveBusinessChatLink", func() interface{} { return new(AccountResolvedBusinessChatLinks) })

	// RPCSponsoredMessages
	iface.RegisterRPCContextTuple("TLAccountToggleSponsoredMessages", "/tg.RPCSponsoredMessages/account.toggleSponsoredMessages", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLContactsGetSponsoredPeers", "/tg.RPCSponsoredMessages/contacts.getSponsoredPeers", func() interface{} { return new(ContactsSponsoredPeers) })
	iface.RegisterRPCContextTuple("TLMessagesViewSponsoredMessage", "/tg.RPCSponsoredMessages/messages.viewSponsoredMessage", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesClickSponsoredMessage", "/tg.RPCSponsoredMessages/messages.clickSponsoredMessage", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReportSponsoredMessage", "/tg.RPCSponsoredMessages/messages.reportSponsoredMessage", func() interface{} { return new(ChannelsSponsoredMessageReportResult) })
	iface.RegisterRPCContextTuple("TLMessagesGetSponsoredMessages", "/tg.RPCSponsoredMessages/messages.getSponsoredMessages", func() interface{} { return new(MessagesSponsoredMessages) })
	iface.RegisterRPCContextTuple("TLChannelsRestrictSponsoredMessages", "/tg.RPCSponsoredMessages/channels.restrictSponsoredMessages", func() interface{} { return new(Updates) })

	// RPCReactionNotification
	iface.RegisterRPCContextTuple("TLAccountGetReactionsNotifySettings", "/tg.RPCReactionNotification/account.getReactionsNotifySettings", func() interface{} { return new(ReactionsNotifySettings) })
	iface.RegisterRPCContextTuple("TLAccountSetReactionsNotifySettings", "/tg.RPCReactionNotification/account.setReactionsNotifySettings", func() interface{} { return new(ReactionsNotifySettings) })

	// RPCPaidMessage
	iface.RegisterRPCContextTuple("TLAccountGetPaidMessagesRevenue", "/tg.RPCPaidMessage/account.getPaidMessagesRevenue", func() interface{} { return new(AccountPaidMessagesRevenue) })
	iface.RegisterRPCContextTuple("TLAccountToggleNoPaidMessagesException", "/tg.RPCPaidMessage/account.toggleNoPaidMessagesException", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsUpdatePaidMessagesPrice", "/tg.RPCPaidMessage/channels.updatePaidMessagesPrice", func() interface{} { return new(Updates) })

	// RPCUsers
	iface.RegisterRPCContextTuple("TLUsersGetUsers", "/tg.RPCUsers/users.getUsers", func() interface{} { return new(VectorUser) })
	iface.RegisterRPCContextTuple("TLUsersGetFullUser", "/tg.RPCUsers/users.getFullUser", func() interface{} { return new(UsersUserFull) })
	iface.RegisterRPCContextTuple("TLContactsResolvePhone", "/tg.RPCUsers/contacts.resolvePhone", func() interface{} { return new(ContactsResolvedPeer) })
	iface.RegisterRPCContextTuple("TLUsersGetMe", "/tg.RPCUsers/users.getMe", func() interface{} { return new(User) })

	// RPCMessageThreads
	iface.RegisterRPCContextTuple("TLContactsBlockFromReplies", "/tg.RPCMessageThreads/contacts.blockFromReplies", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetReplies", "/tg.RPCMessageThreads/messages.getReplies", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetDiscussionMessage", "/tg.RPCMessageThreads/messages.getDiscussionMessage", func() interface{} { return new(MessagesDiscussionMessage) })
	iface.RegisterRPCContextTuple("TLMessagesReadDiscussion", "/tg.RPCMessageThreads/messages.readDiscussion", func() interface{} { return new(Bool) })

	// RPCProfileLinks
	iface.RegisterRPCContextTuple("TLContactsExportContactToken", "/tg.RPCProfileLinks/contacts.exportContactToken", func() interface{} { return new(ExportedContactToken) })
	iface.RegisterRPCContextTuple("TLContactsImportContactToken", "/tg.RPCProfileLinks/contacts.importContactToken", func() interface{} { return new(User) })

	// RPCMessages
	iface.RegisterRPCContextTuple("TLMessagesGetMessages", "/tg.RPCMessages/messages.getMessages", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetHistory", "/tg.RPCMessages/messages.getHistory", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesSearch", "/tg.RPCMessages/messages.search", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesReadHistory", "/tg.RPCMessages/messages.readHistory", func() interface{} { return new(MessagesAffectedMessages) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteHistory", "/tg.RPCMessages/messages.deleteHistory", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteMessages", "/tg.RPCMessages/messages.deleteMessages", func() interface{} { return new(MessagesAffectedMessages) })
	iface.RegisterRPCContextTuple("TLMessagesReceivedMessages", "/tg.RPCMessages/messages.receivedMessages", func() interface{} { return new(VectorReceivedNotifyMessage) })
	iface.RegisterRPCContextTuple("TLMessagesSendMessage", "/tg.RPCMessages/messages.sendMessage", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesSendMedia", "/tg.RPCMessages/messages.sendMedia", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesForwardMessages", "/tg.RPCMessages/messages.forwardMessages", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesReadMessageContents", "/tg.RPCMessages/messages.readMessageContents", func() interface{} { return new(MessagesAffectedMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetMessagesViews", "/tg.RPCMessages/messages.getMessagesViews", func() interface{} { return new(MessagesMessageViews) })
	iface.RegisterRPCContextTuple("TLMessagesSearchGlobal", "/tg.RPCMessages/messages.searchGlobal", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetMessageEditData", "/tg.RPCMessages/messages.getMessageEditData", func() interface{} { return new(MessagesMessageEditData) })
	iface.RegisterRPCContextTuple("TLMessagesEditMessage", "/tg.RPCMessages/messages.editMessage", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetUnreadMentions", "/tg.RPCMessages/messages.getUnreadMentions", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesReadMentions", "/tg.RPCMessages/messages.readMentions", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLMessagesGetRecentLocations", "/tg.RPCMessages/messages.getRecentLocations", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesSendMultiMedia", "/tg.RPCMessages/messages.sendMultiMedia", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesUpdatePinnedMessage", "/tg.RPCMessages/messages.updatePinnedMessage", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetSearchCounters", "/tg.RPCMessages/messages.getSearchCounters", func() interface{} { return new(VectorMessagesSearchCounter) })
	iface.RegisterRPCContextTuple("TLMessagesUnpinAllMessages", "/tg.RPCMessages/messages.unpinAllMessages", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLMessagesGetSearchResultsCalendar", "/tg.RPCMessages/messages.getSearchResultsCalendar", func() interface{} { return new(MessagesSearchResultsCalendar) })
	iface.RegisterRPCContextTuple("TLMessagesGetSearchResultsPositions", "/tg.RPCMessages/messages.getSearchResultsPositions", func() interface{} { return new(MessagesSearchResultsPositions) })
	iface.RegisterRPCContextTuple("TLMessagesToggleNoForwards", "/tg.RPCMessages/messages.toggleNoForwards", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesSaveDefaultSendAs", "/tg.RPCMessages/messages.saveDefaultSendAs", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSearchSentMedia", "/tg.RPCMessages/messages.searchSentMedia", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetOutboxReadDate", "/tg.RPCMessages/messages.getOutboxReadDate", func() interface{} { return new(OutboxReadDate) })
	iface.RegisterRPCContextTuple("TLMessagesSummarizeText", "/tg.RPCMessages/messages.summarizeText", func() interface{} { return new(TextWithEntities) })
	iface.RegisterRPCContextTuple("TLChannelsGetSendAs", "/tg.RPCMessages/channels.getSendAs", func() interface{} { return new(ChannelsSendAsPeers) })
	iface.RegisterRPCContextTuple("TLChannelsSearchPosts", "/tg.RPCMessages/channels.searchPosts", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLChannelsCheckSearchPostsFlood", "/tg.RPCMessages/channels.checkSearchPostsFlood", func() interface{} { return new(SearchPostsFlood) })

	// RPCDialogs
	iface.RegisterRPCContextTuple("TLMessagesGetDialogs", "/tg.RPCDialogs/messages.getDialogs", func() interface{} { return new(MessagesDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesSetTyping", "/tg.RPCDialogs/messages.setTyping", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetPeerSettings", "/tg.RPCDialogs/messages.getPeerSettings", func() interface{} { return new(MessagesPeerSettings) })
	iface.RegisterRPCContextTuple("TLMessagesGetPeerDialogs", "/tg.RPCDialogs/messages.getPeerDialogs", func() interface{} { return new(MessagesPeerDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesToggleDialogPin", "/tg.RPCDialogs/messages.toggleDialogPin", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReorderPinnedDialogs", "/tg.RPCDialogs/messages.reorderPinnedDialogs", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetPinnedDialogs", "/tg.RPCDialogs/messages.getPinnedDialogs", func() interface{} { return new(MessagesPeerDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesSendScreenshotNotification", "/tg.RPCDialogs/messages.sendScreenshotNotification", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesMarkDialogUnread", "/tg.RPCDialogs/messages.markDialogUnread", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetDialogUnreadMarks", "/tg.RPCDialogs/messages.getDialogUnreadMarks", func() interface{} { return new(VectorDialogPeer) })
	iface.RegisterRPCContextTuple("TLMessagesGetOnlines", "/tg.RPCDialogs/messages.getOnlines", func() interface{} { return new(ChatOnlines) })
	iface.RegisterRPCContextTuple("TLMessagesHidePeerSettingsBar", "/tg.RPCDialogs/messages.hidePeerSettingsBar", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSetHistoryTTL", "/tg.RPCDialogs/messages.setHistoryTTL", func() interface{} { return new(Updates) })

	// RPCChats
	iface.RegisterRPCContextTuple("TLMessagesGetChats", "/tg.RPCChats/messages.getChats", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLMessagesGetFullChat", "/tg.RPCChats/messages.getFullChat", func() interface{} { return new(MessagesChatFull) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatTitle", "/tg.RPCChats/messages.editChatTitle", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatPhoto", "/tg.RPCChats/messages.editChatPhoto", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesAddChatUser", "/tg.RPCChats/messages.addChatUser", func() interface{} { return new(MessagesInvitedUsers) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteChatUser", "/tg.RPCChats/messages.deleteChatUser", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesCreateChat", "/tg.RPCChats/messages.createChat", func() interface{} { return new(MessagesInvitedUsers) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatAdmin", "/tg.RPCChats/messages.editChatAdmin", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesMigrateChat", "/tg.RPCChats/messages.migrateChat", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetCommonChats", "/tg.RPCChats/messages.getCommonChats", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatAbout", "/tg.RPCChats/messages.editChatAbout", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatDefaultBannedRights", "/tg.RPCChats/messages.editChatDefaultBannedRights", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteChat", "/tg.RPCChats/messages.deleteChat", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetMessageReadParticipants", "/tg.RPCChats/messages.getMessageReadParticipants", func() interface{} { return new(VectorReadParticipantDate) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatCreator", "/tg.RPCChats/messages.editChatCreator", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetFutureChatCreatorAfterLeave", "/tg.RPCChats/messages.getFutureChatCreatorAfterLeave", func() interface{} { return new(User) })
	iface.RegisterRPCContextTuple("TLMessagesEditChatParticipantRank", "/tg.RPCChats/messages.editChatParticipantRank", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsConvertToGigagroup", "/tg.RPCChats/channels.convertToGigagroup", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsSetEmojiStickers", "/tg.RPCChats/channels.setEmojiStickers", func() interface{} { return new(Bool) })

	// RPCSecretChats
	iface.RegisterRPCContextTuple("TLMessagesGetDhConfig", "/tg.RPCSecretChats/messages.getDhConfig", func() interface{} { return new(MessagesDhConfig) })
	iface.RegisterRPCContextTuple("TLMessagesRequestEncryption", "/tg.RPCSecretChats/messages.requestEncryption", func() interface{} { return new(EncryptedChat) })
	iface.RegisterRPCContextTuple("TLMessagesAcceptEncryption", "/tg.RPCSecretChats/messages.acceptEncryption", func() interface{} { return new(EncryptedChat) })
	iface.RegisterRPCContextTuple("TLMessagesDiscardEncryption", "/tg.RPCSecretChats/messages.discardEncryption", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSetEncryptedTyping", "/tg.RPCSecretChats/messages.setEncryptedTyping", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReadEncryptedHistory", "/tg.RPCSecretChats/messages.readEncryptedHistory", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSendEncrypted", "/tg.RPCSecretChats/messages.sendEncrypted", func() interface{} { return new(MessagesSentEncryptedMessage) })
	iface.RegisterRPCContextTuple("TLMessagesSendEncryptedFile", "/tg.RPCSecretChats/messages.sendEncryptedFile", func() interface{} { return new(MessagesSentEncryptedMessage) })
	iface.RegisterRPCContextTuple("TLMessagesSendEncryptedService", "/tg.RPCSecretChats/messages.sendEncryptedService", func() interface{} { return new(MessagesSentEncryptedMessage) })
	iface.RegisterRPCContextTuple("TLMessagesReceivedQueue", "/tg.RPCSecretChats/messages.receivedQueue", func() interface{} { return new(VectorLong) })

	// RPCStickers
	iface.RegisterRPCContextTuple("TLMessagesGetStickers", "/tg.RPCStickers/messages.getStickers", func() interface{} { return new(MessagesStickers) })
	iface.RegisterRPCContextTuple("TLMessagesGetAllStickers", "/tg.RPCStickers/messages.getAllStickers", func() interface{} { return new(MessagesAllStickers) })
	iface.RegisterRPCContextTuple("TLMessagesGetStickerSet", "/tg.RPCStickers/messages.getStickerSet", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLMessagesInstallStickerSet", "/tg.RPCStickers/messages.installStickerSet", func() interface{} { return new(MessagesStickerSetInstallResult) })
	iface.RegisterRPCContextTuple("TLMessagesUninstallStickerSet", "/tg.RPCStickers/messages.uninstallStickerSet", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReorderStickerSets", "/tg.RPCStickers/messages.reorderStickerSets", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetFeaturedStickers", "/tg.RPCStickers/messages.getFeaturedStickers", func() interface{} { return new(MessagesFeaturedStickers) })
	iface.RegisterRPCContextTuple("TLMessagesReadFeaturedStickers", "/tg.RPCStickers/messages.readFeaturedStickers", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetRecentStickers", "/tg.RPCStickers/messages.getRecentStickers", func() interface{} { return new(MessagesRecentStickers) })
	iface.RegisterRPCContextTuple("TLMessagesSaveRecentSticker", "/tg.RPCStickers/messages.saveRecentSticker", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesClearRecentStickers", "/tg.RPCStickers/messages.clearRecentStickers", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetArchivedStickers", "/tg.RPCStickers/messages.getArchivedStickers", func() interface{} { return new(MessagesArchivedStickers) })
	iface.RegisterRPCContextTuple("TLMessagesGetMaskStickers", "/tg.RPCStickers/messages.getMaskStickers", func() interface{} { return new(MessagesAllStickers) })
	iface.RegisterRPCContextTuple("TLMessagesGetAttachedStickers", "/tg.RPCStickers/messages.getAttachedStickers", func() interface{} { return new(VectorStickerSetCovered) })
	iface.RegisterRPCContextTuple("TLMessagesGetFavedStickers", "/tg.RPCStickers/messages.getFavedStickers", func() interface{} { return new(MessagesFavedStickers) })
	iface.RegisterRPCContextTuple("TLMessagesFaveSticker", "/tg.RPCStickers/messages.faveSticker", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSearchStickerSets", "/tg.RPCStickers/messages.searchStickerSets", func() interface{} { return new(MessagesFoundStickerSets) })
	iface.RegisterRPCContextTuple("TLMessagesToggleStickerSets", "/tg.RPCStickers/messages.toggleStickerSets", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetOldFeaturedStickers", "/tg.RPCStickers/messages.getOldFeaturedStickers", func() interface{} { return new(MessagesFeaturedStickers) })
	iface.RegisterRPCContextTuple("TLMessagesSearchEmojiStickerSets", "/tg.RPCStickers/messages.searchEmojiStickerSets", func() interface{} { return new(MessagesFoundStickerSets) })
	iface.RegisterRPCContextTuple("TLMessagesGetMyStickers", "/tg.RPCStickers/messages.getMyStickers", func() interface{} { return new(MessagesMyStickers) })
	iface.RegisterRPCContextTuple("TLMessagesSearchStickers", "/tg.RPCStickers/messages.searchStickers", func() interface{} { return new(MessagesFoundStickers) })
	iface.RegisterRPCContextTuple("TLStickersCreateStickerSet", "/tg.RPCStickers/stickers.createStickerSet", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersRemoveStickerFromSet", "/tg.RPCStickers/stickers.removeStickerFromSet", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersChangeStickerPosition", "/tg.RPCStickers/stickers.changeStickerPosition", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersAddStickerToSet", "/tg.RPCStickers/stickers.addStickerToSet", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersSetStickerSetThumb", "/tg.RPCStickers/stickers.setStickerSetThumb", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersCheckShortName", "/tg.RPCStickers/stickers.checkShortName", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStickersSuggestShortName", "/tg.RPCStickers/stickers.suggestShortName", func() interface{} { return new(StickersSuggestedShortName) })
	iface.RegisterRPCContextTuple("TLStickersChangeSticker", "/tg.RPCStickers/stickers.changeSticker", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersRenameStickerSet", "/tg.RPCStickers/stickers.renameStickerSet", func() interface{} { return new(MessagesStickerSet) })
	iface.RegisterRPCContextTuple("TLStickersDeleteStickerSet", "/tg.RPCStickers/stickers.deleteStickerSet", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStickersReplaceSticker", "/tg.RPCStickers/stickers.replaceSticker", func() interface{} { return new(MessagesStickerSet) })

	// RPCWebPage
	iface.RegisterRPCContextTuple("TLMessagesGetWebPagePreview", "/tg.RPCWebPage/messages.getWebPagePreview", func() interface{} { return new(MessagesWebPagePreview) })
	iface.RegisterRPCContextTuple("TLMessagesGetWebPage", "/tg.RPCWebPage/messages.getWebPage", func() interface{} { return new(MessagesWebPage) })

	// RPCChatInvites
	iface.RegisterRPCContextTuple("TLMessagesExportChatInvite", "/tg.RPCChatInvites/messages.exportChatInvite", func() interface{} { return new(ExportedChatInvite) })
	iface.RegisterRPCContextTuple("TLMessagesCheckChatInvite", "/tg.RPCChatInvites/messages.checkChatInvite", func() interface{} { return new(ChatInvite) })
	iface.RegisterRPCContextTuple("TLMessagesImportChatInvite", "/tg.RPCChatInvites/messages.importChatInvite", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetExportedChatInvites", "/tg.RPCChatInvites/messages.getExportedChatInvites", func() interface{} { return new(MessagesExportedChatInvites) })
	iface.RegisterRPCContextTuple("TLMessagesGetExportedChatInvite", "/tg.RPCChatInvites/messages.getExportedChatInvite", func() interface{} { return new(MessagesExportedChatInvite) })
	iface.RegisterRPCContextTuple("TLMessagesEditExportedChatInvite", "/tg.RPCChatInvites/messages.editExportedChatInvite", func() interface{} { return new(MessagesExportedChatInvite) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteRevokedExportedChatInvites", "/tg.RPCChatInvites/messages.deleteRevokedExportedChatInvites", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteExportedChatInvite", "/tg.RPCChatInvites/messages.deleteExportedChatInvite", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetAdminsWithInvites", "/tg.RPCChatInvites/messages.getAdminsWithInvites", func() interface{} { return new(MessagesChatAdminsWithInvites) })
	iface.RegisterRPCContextTuple("TLMessagesGetChatInviteImporters", "/tg.RPCChatInvites/messages.getChatInviteImporters", func() interface{} { return new(MessagesChatInviteImporters) })
	iface.RegisterRPCContextTuple("TLMessagesHideChatJoinRequest", "/tg.RPCChatInvites/messages.hideChatJoinRequest", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesHideAllChatJoinRequests", "/tg.RPCChatInvites/messages.hideAllChatJoinRequests", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsToggleJoinToSend", "/tg.RPCChatInvites/channels.toggleJoinToSend", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsToggleJoinRequest", "/tg.RPCChatInvites/channels.toggleJoinRequest", func() interface{} { return new(Updates) })

	// RPCDeepLinks
	iface.RegisterRPCContextTuple("TLMessagesStartBot", "/tg.RPCDeepLinks/messages.startBot", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLHelpGetRecentMeUrls", "/tg.RPCDeepLinks/help.getRecentMeUrls", func() interface{} { return new(HelpRecentMeUrls) })
	iface.RegisterRPCContextTuple("TLHelpGetDeepLinkInfo", "/tg.RPCDeepLinks/help.getDeepLinkInfo", func() interface{} { return new(HelpDeepLinkInfo) })

	// RPCFiles
	iface.RegisterRPCContextTuple("TLMessagesGetDocumentByHash", "/tg.RPCFiles/messages.getDocumentByHash", func() interface{} { return new(Document) })
	iface.RegisterRPCContextTuple("TLMessagesUploadMedia", "/tg.RPCFiles/messages.uploadMedia", func() interface{} { return new(MessageMedia) })
	iface.RegisterRPCContextTuple("TLMessagesUploadEncryptedFile", "/tg.RPCFiles/messages.uploadEncryptedFile", func() interface{} { return new(EncryptedFile) })
	iface.RegisterRPCContextTuple("TLUploadSaveFilePart", "/tg.RPCFiles/upload.saveFilePart", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLUploadGetFile", "/tg.RPCFiles/upload.getFile", func() interface{} { return new(UploadFile) })
	iface.RegisterRPCContextTuple("TLUploadSaveBigFilePart", "/tg.RPCFiles/upload.saveBigFilePart", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLUploadGetWebFile", "/tg.RPCFiles/upload.getWebFile", func() interface{} { return new(UploadWebFile) })
	iface.RegisterRPCContextTuple("TLUploadGetCdnFile", "/tg.RPCFiles/upload.getCdnFile", func() interface{} { return new(UploadCdnFile) })
	iface.RegisterRPCContextTuple("TLUploadReuploadCdnFile", "/tg.RPCFiles/upload.reuploadCdnFile", func() interface{} { return new(VectorFileHash) })
	iface.RegisterRPCContextTuple("TLUploadGetCdnFileHashes", "/tg.RPCFiles/upload.getCdnFileHashes", func() interface{} { return new(VectorFileHash) })
	iface.RegisterRPCContextTuple("TLUploadGetFileHashes", "/tg.RPCFiles/upload.getFileHashes", func() interface{} { return new(VectorFileHash) })
	iface.RegisterRPCContextTuple("TLHelpGetCdnConfig", "/tg.RPCFiles/help.getCdnConfig", func() interface{} { return new(CdnConfig) })

	// RPCGifs
	iface.RegisterRPCContextTuple("TLMessagesGetSavedGifs", "/tg.RPCGifs/messages.getSavedGifs", func() interface{} { return new(MessagesSavedGifs) })
	iface.RegisterRPCContextTuple("TLMessagesSaveGif", "/tg.RPCGifs/messages.saveGif", func() interface{} { return new(Bool) })

	// RPCInlineBot
	iface.RegisterRPCContextTuple("TLMessagesGetInlineBotResults", "/tg.RPCInlineBot/messages.getInlineBotResults", func() interface{} { return new(MessagesBotResults) })
	iface.RegisterRPCContextTuple("TLMessagesSetInlineBotResults", "/tg.RPCInlineBot/messages.setInlineBotResults", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSendInlineBotResult", "/tg.RPCInlineBot/messages.sendInlineBotResult", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesEditInlineBotMessage", "/tg.RPCInlineBot/messages.editInlineBotMessage", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetBotCallbackAnswer", "/tg.RPCInlineBot/messages.getBotCallbackAnswer", func() interface{} { return new(MessagesBotCallbackAnswer) })
	iface.RegisterRPCContextTuple("TLMessagesSetBotCallbackAnswer", "/tg.RPCInlineBot/messages.setBotCallbackAnswer", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSendBotRequestedPeer", "/tg.RPCInlineBot/messages.sendBotRequestedPeer", func() interface{} { return new(Updates) })

	// RPCDrafts
	iface.RegisterRPCContextTuple("TLMessagesSaveDraft", "/tg.RPCDrafts/messages.saveDraft", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetAllDrafts", "/tg.RPCDrafts/messages.getAllDrafts", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesClearAllDrafts", "/tg.RPCDrafts/messages.clearAllDrafts", func() interface{} { return new(Bool) })

	// RPCGames
	iface.RegisterRPCContextTuple("TLMessagesSetGameScore", "/tg.RPCGames/messages.setGameScore", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesSetInlineGameScore", "/tg.RPCGames/messages.setInlineGameScore", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetGameHighScores", "/tg.RPCGames/messages.getGameHighScores", func() interface{} { return new(MessagesHighScores) })
	iface.RegisterRPCContextTuple("TLMessagesGetInlineGameHighScores", "/tg.RPCGames/messages.getInlineGameHighScores", func() interface{} { return new(MessagesHighScores) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiGameInfo", "/tg.RPCGames/messages.getEmojiGameInfo", func() interface{} { return new(MessagesEmojiGameInfo) })

	// RPCPolls
	iface.RegisterRPCContextTuple("TLMessagesSendVote", "/tg.RPCPolls/messages.sendVote", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetPollResults", "/tg.RPCPolls/messages.getPollResults", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetPollVotes", "/tg.RPCPolls/messages.getPollVotes", func() interface{} { return new(MessagesVotesList) })

	// RPCEmoji
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiKeywords", "/tg.RPCEmoji/messages.getEmojiKeywords", func() interface{} { return new(EmojiKeywordsDifference) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiKeywordsDifference", "/tg.RPCEmoji/messages.getEmojiKeywordsDifference", func() interface{} { return new(EmojiKeywordsDifference) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiKeywordsLanguages", "/tg.RPCEmoji/messages.getEmojiKeywordsLanguages", func() interface{} { return new(VectorEmojiLanguage) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiURL", "/tg.RPCEmoji/messages.getEmojiURL", func() interface{} { return new(EmojiURL) })

	// RPCScheduledMessages
	iface.RegisterRPCContextTuple("TLMessagesGetScheduledHistory", "/tg.RPCScheduledMessages/messages.getScheduledHistory", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesGetScheduledMessages", "/tg.RPCScheduledMessages/messages.getScheduledMessages", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesSendScheduledMessages", "/tg.RPCScheduledMessages/messages.sendScheduledMessages", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteScheduledMessages", "/tg.RPCScheduledMessages/messages.deleteScheduledMessages", func() interface{} { return new(Updates) })

	// RPCFolders
	iface.RegisterRPCContextTuple("TLMessagesGetDialogFilters", "/tg.RPCFolders/messages.getDialogFilters", func() interface{} { return new(MessagesDialogFilters) })
	iface.RegisterRPCContextTuple("TLMessagesGetSuggestedDialogFilters", "/tg.RPCFolders/messages.getSuggestedDialogFilters", func() interface{} { return new(VectorDialogFilterSuggested) })
	iface.RegisterRPCContextTuple("TLMessagesUpdateDialogFilter", "/tg.RPCFolders/messages.updateDialogFilter", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesUpdateDialogFiltersOrder", "/tg.RPCFolders/messages.updateDialogFiltersOrder", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLFoldersEditPeerFolders", "/tg.RPCFolders/folders.editPeerFolders", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChatlistsExportChatlistInvite", "/tg.RPCFolders/chatlists.exportChatlistInvite", func() interface{} { return new(ChatlistsExportedChatlistInvite) })
	iface.RegisterRPCContextTuple("TLChatlistsDeleteExportedInvite", "/tg.RPCFolders/chatlists.deleteExportedInvite", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChatlistsEditExportedInvite", "/tg.RPCFolders/chatlists.editExportedInvite", func() interface{} { return new(ExportedChatlistInvite) })
	iface.RegisterRPCContextTuple("TLChatlistsGetExportedInvites", "/tg.RPCFolders/chatlists.getExportedInvites", func() interface{} { return new(ChatlistsExportedInvites) })
	iface.RegisterRPCContextTuple("TLChatlistsCheckChatlistInvite", "/tg.RPCFolders/chatlists.checkChatlistInvite", func() interface{} { return new(ChatlistsChatlistInvite) })
	iface.RegisterRPCContextTuple("TLChatlistsJoinChatlistInvite", "/tg.RPCFolders/chatlists.joinChatlistInvite", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChatlistsGetChatlistUpdates", "/tg.RPCFolders/chatlists.getChatlistUpdates", func() interface{} { return new(ChatlistsChatlistUpdates) })
	iface.RegisterRPCContextTuple("TLChatlistsJoinChatlistUpdates", "/tg.RPCFolders/chatlists.joinChatlistUpdates", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChatlistsHideChatlistUpdates", "/tg.RPCFolders/chatlists.hideChatlistUpdates", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChatlistsGetLeaveChatlistSuggestions", "/tg.RPCFolders/chatlists.getLeaveChatlistSuggestions", func() interface{} { return new(VectorPeer) })
	iface.RegisterRPCContextTuple("TLChatlistsLeaveChatlist", "/tg.RPCFolders/chatlists.leaveChatlist", func() interface{} { return new(Updates) })

	// RPCVoipCalls
	iface.RegisterRPCContextTuple("TLMessagesDeletePhoneCallHistory", "/tg.RPCVoipCalls/messages.deletePhoneCallHistory", func() interface{} { return new(MessagesAffectedFoundMessages) })
	iface.RegisterRPCContextTuple("TLPhoneGetCallConfig", "/tg.RPCVoipCalls/phone.getCallConfig", func() interface{} { return new(DataJSON) })
	iface.RegisterRPCContextTuple("TLPhoneRequestCall", "/tg.RPCVoipCalls/phone.requestCall", func() interface{} { return new(PhonePhoneCall) })
	iface.RegisterRPCContextTuple("TLPhoneAcceptCall", "/tg.RPCVoipCalls/phone.acceptCall", func() interface{} { return new(PhonePhoneCall) })
	iface.RegisterRPCContextTuple("TLPhoneConfirmCall", "/tg.RPCVoipCalls/phone.confirmCall", func() interface{} { return new(PhonePhoneCall) })
	iface.RegisterRPCContextTuple("TLPhoneReceivedCall", "/tg.RPCVoipCalls/phone.receivedCall", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPhoneDiscardCall", "/tg.RPCVoipCalls/phone.discardCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneSetCallRating", "/tg.RPCVoipCalls/phone.setCallRating", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneSaveCallDebug", "/tg.RPCVoipCalls/phone.saveCallDebug", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPhoneSendSignalingData", "/tg.RPCVoipCalls/phone.sendSignalingData", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPhoneSaveCallLog", "/tg.RPCVoipCalls/phone.saveCallLog", func() interface{} { return new(Bool) })

	// RPCImportedChats
	iface.RegisterRPCContextTuple("TLMessagesCheckHistoryImport", "/tg.RPCImportedChats/messages.checkHistoryImport", func() interface{} { return new(MessagesHistoryImportParsed) })
	iface.RegisterRPCContextTuple("TLMessagesInitHistoryImport", "/tg.RPCImportedChats/messages.initHistoryImport", func() interface{} { return new(MessagesHistoryImport) })
	iface.RegisterRPCContextTuple("TLMessagesUploadImportedMedia", "/tg.RPCImportedChats/messages.uploadImportedMedia", func() interface{} { return new(MessageMedia) })
	iface.RegisterRPCContextTuple("TLMessagesStartHistoryImport", "/tg.RPCImportedChats/messages.startHistoryImport", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesCheckHistoryImportPeer", "/tg.RPCImportedChats/messages.checkHistoryImportPeer", func() interface{} { return new(MessagesCheckedHistoryImportPeer) })

	// RPCReactions
	iface.RegisterRPCContextTuple("TLMessagesSendReaction", "/tg.RPCReactions/messages.sendReaction", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetMessagesReactions", "/tg.RPCReactions/messages.getMessagesReactions", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetMessageReactionsList", "/tg.RPCReactions/messages.getMessageReactionsList", func() interface{} { return new(MessagesMessageReactionsList) })
	iface.RegisterRPCContextTuple("TLMessagesSetChatAvailableReactions", "/tg.RPCReactions/messages.setChatAvailableReactions", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetAvailableReactions", "/tg.RPCReactions/messages.getAvailableReactions", func() interface{} { return new(MessagesAvailableReactions) })
	iface.RegisterRPCContextTuple("TLMessagesSetDefaultReaction", "/tg.RPCReactions/messages.setDefaultReaction", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetUnreadReactions", "/tg.RPCReactions/messages.getUnreadReactions", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesReadReactions", "/tg.RPCReactions/messages.readReactions", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLMessagesReportReaction", "/tg.RPCReactions/messages.reportReaction", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetTopReactions", "/tg.RPCReactions/messages.getTopReactions", func() interface{} { return new(MessagesReactions) })
	iface.RegisterRPCContextTuple("TLMessagesGetRecentReactions", "/tg.RPCReactions/messages.getRecentReactions", func() interface{} { return new(MessagesReactions) })
	iface.RegisterRPCContextTuple("TLMessagesClearRecentReactions", "/tg.RPCReactions/messages.clearRecentReactions", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesSendPaidReaction", "/tg.RPCReactions/messages.sendPaidReaction", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesTogglePaidReactionPrivacy", "/tg.RPCReactions/messages.togglePaidReactionPrivacy", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetPaidReactionPrivacy", "/tg.RPCReactions/messages.getPaidReactionPrivacy", func() interface{} { return new(Updates) })

	// RPCTranslation
	iface.RegisterRPCContextTuple("TLMessagesTranslateText", "/tg.RPCTranslation/messages.translateText", func() interface{} { return new(MessagesTranslatedText) })
	iface.RegisterRPCContextTuple("TLMessagesTogglePeerTranslations", "/tg.RPCTranslation/messages.togglePeerTranslations", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsToggleAutotranslation", "/tg.RPCTranslation/channels.toggleAutotranslation", func() interface{} { return new(Updates) })

	// RPCBotMenu
	iface.RegisterRPCContextTuple("TLMessagesGetAttachMenuBots", "/tg.RPCBotMenu/messages.getAttachMenuBots", func() interface{} { return new(AttachMenuBots) })
	iface.RegisterRPCContextTuple("TLMessagesGetAttachMenuBot", "/tg.RPCBotMenu/messages.getAttachMenuBot", func() interface{} { return new(AttachMenuBotsBot) })
	iface.RegisterRPCContextTuple("TLMessagesToggleBotInAttachMenu", "/tg.RPCBotMenu/messages.toggleBotInAttachMenu", func() interface{} { return new(Bool) })

	// RPCMiniBotApps
	iface.RegisterRPCContextTuple("TLMessagesRequestWebView", "/tg.RPCMiniBotApps/messages.requestWebView", func() interface{} { return new(WebViewResult) })
	iface.RegisterRPCContextTuple("TLMessagesProlongWebView", "/tg.RPCMiniBotApps/messages.prolongWebView", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesRequestSimpleWebView", "/tg.RPCMiniBotApps/messages.requestSimpleWebView", func() interface{} { return new(WebViewResult) })
	iface.RegisterRPCContextTuple("TLMessagesSendWebViewResultMessage", "/tg.RPCMiniBotApps/messages.sendWebViewResultMessage", func() interface{} { return new(WebViewMessageSent) })
	iface.RegisterRPCContextTuple("TLMessagesSendWebViewData", "/tg.RPCMiniBotApps/messages.sendWebViewData", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetBotApp", "/tg.RPCMiniBotApps/messages.getBotApp", func() interface{} { return new(MessagesBotApp) })
	iface.RegisterRPCContextTuple("TLMessagesRequestAppWebView", "/tg.RPCMiniBotApps/messages.requestAppWebView", func() interface{} { return new(WebViewResult) })
	iface.RegisterRPCContextTuple("TLBotsCanSendMessage", "/tg.RPCMiniBotApps/bots.canSendMessage", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsAllowSendMessage", "/tg.RPCMiniBotApps/bots.allowSendMessage", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLBotsInvokeWebViewCustomMethod", "/tg.RPCMiniBotApps/bots.invokeWebViewCustomMethod", func() interface{} { return new(DataJSON) })
	iface.RegisterRPCContextTuple("TLBotsCheckDownloadFileParams", "/tg.RPCMiniBotApps/bots.checkDownloadFileParams", func() interface{} { return new(Bool) })

	// RPCTranscription
	iface.RegisterRPCContextTuple("TLMessagesTranscribeAudio", "/tg.RPCTranscription/messages.transcribeAudio", func() interface{} { return new(MessagesTranscribedAudio) })
	iface.RegisterRPCContextTuple("TLMessagesRateTranscribedAudio", "/tg.RPCTranscription/messages.rateTranscribedAudio", func() interface{} { return new(Bool) })

	// RPCPaidMedia
	iface.RegisterRPCContextTuple("TLMessagesGetExtendedMedia", "/tg.RPCPaidMedia/messages.getExtendedMedia", func() interface{} { return new(Updates) })

	// RPCEmojiCategories
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiGroups", "/tg.RPCEmojiCategories/messages.getEmojiGroups", func() interface{} { return new(MessagesEmojiGroups) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiStatusGroups", "/tg.RPCEmojiCategories/messages.getEmojiStatusGroups", func() interface{} { return new(MessagesEmojiGroups) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiProfilePhotoGroups", "/tg.RPCEmojiCategories/messages.getEmojiProfilePhotoGroups", func() interface{} { return new(MessagesEmojiGroups) })
	iface.RegisterRPCContextTuple("TLMessagesGetEmojiStickerGroups", "/tg.RPCEmojiCategories/messages.getEmojiStickerGroups", func() interface{} { return new(MessagesEmojiGroups) })

	// RPCSavedMessageDialogs
	iface.RegisterRPCContextTuple("TLMessagesGetSavedDialogs", "/tg.RPCSavedMessageDialogs/messages.getSavedDialogs", func() interface{} { return new(MessagesSavedDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesGetSavedHistory", "/tg.RPCSavedMessageDialogs/messages.getSavedHistory", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteSavedHistory", "/tg.RPCSavedMessageDialogs/messages.deleteSavedHistory", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLMessagesGetPinnedSavedDialogs", "/tg.RPCSavedMessageDialogs/messages.getPinnedSavedDialogs", func() interface{} { return new(MessagesSavedDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesToggleSavedDialogPin", "/tg.RPCSavedMessageDialogs/messages.toggleSavedDialogPin", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesReorderPinnedSavedDialogs", "/tg.RPCSavedMessageDialogs/messages.reorderPinnedSavedDialogs", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetSavedDialogsByID", "/tg.RPCSavedMessageDialogs/messages.getSavedDialogsByID", func() interface{} { return new(MessagesSavedDialogs) })
	iface.RegisterRPCContextTuple("TLMessagesReadSavedHistory", "/tg.RPCSavedMessageDialogs/messages.readSavedHistory", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsGetMessageAuthor", "/tg.RPCSavedMessageDialogs/channels.getMessageAuthor", func() interface{} { return new(User) })

	// RPCSavedMessageTags
	iface.RegisterRPCContextTuple("TLMessagesGetSavedReactionTags", "/tg.RPCSavedMessageTags/messages.getSavedReactionTags", func() interface{} { return new(MessagesSavedReactionTags) })
	iface.RegisterRPCContextTuple("TLMessagesUpdateSavedReactionTag", "/tg.RPCSavedMessageTags/messages.updateSavedReactionTag", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetDefaultTagReactions", "/tg.RPCSavedMessageTags/messages.getDefaultTagReactions", func() interface{} { return new(MessagesReactions) })

	// RPCBusinessQuickReply
	iface.RegisterRPCContextTuple("TLMessagesGetQuickReplies", "/tg.RPCBusinessQuickReply/messages.getQuickReplies", func() interface{} { return new(MessagesQuickReplies) })
	iface.RegisterRPCContextTuple("TLMessagesReorderQuickReplies", "/tg.RPCBusinessQuickReply/messages.reorderQuickReplies", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesCheckQuickReplyShortcut", "/tg.RPCBusinessQuickReply/messages.checkQuickReplyShortcut", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesEditQuickReplyShortcut", "/tg.RPCBusinessQuickReply/messages.editQuickReplyShortcut", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteQuickReplyShortcut", "/tg.RPCBusinessQuickReply/messages.deleteQuickReplyShortcut", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLMessagesGetQuickReplyMessages", "/tg.RPCBusinessQuickReply/messages.getQuickReplyMessages", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLMessagesSendQuickReplyMessages", "/tg.RPCBusinessQuickReply/messages.sendQuickReplyMessages", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteQuickReplyMessages", "/tg.RPCBusinessQuickReply/messages.deleteQuickReplyMessages", func() interface{} { return new(Updates) })

	// RPCFolderTags
	iface.RegisterRPCContextTuple("TLMessagesToggleDialogFilterTags", "/tg.RPCFolderTags/messages.toggleDialogFilterTags", func() interface{} { return new(Bool) })

	// RPCMessageEffects
	iface.RegisterRPCContextTuple("TLMessagesGetAvailableEffects", "/tg.RPCMessageEffects/messages.getAvailableEffects", func() interface{} { return new(MessagesAvailableEffects) })

	// RPCFactChecks
	iface.RegisterRPCContextTuple("TLMessagesEditFactCheck", "/tg.RPCFactChecks/messages.editFactCheck", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteFactCheck", "/tg.RPCFactChecks/messages.deleteFactCheck", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesGetFactCheck", "/tg.RPCFactChecks/messages.getFactCheck", func() interface{} { return new(VectorFactCheck) })

	// RPCMainMiniBotApps
	iface.RegisterRPCContextTuple("TLMessagesRequestMainWebView", "/tg.RPCMainMiniBotApps/messages.requestMainWebView", func() interface{} { return new(WebViewResult) })
	iface.RegisterRPCContextTuple("TLBotsGetPopularAppBots", "/tg.RPCMainMiniBotApps/bots.getPopularAppBots", func() interface{} { return new(BotsPopularAppBots) })
	iface.RegisterRPCContextTuple("TLBotsAddPreviewMedia", "/tg.RPCMainMiniBotApps/bots.addPreviewMedia", func() interface{} { return new(BotPreviewMedia) })
	iface.RegisterRPCContextTuple("TLBotsEditPreviewMedia", "/tg.RPCMainMiniBotApps/bots.editPreviewMedia", func() interface{} { return new(BotPreviewMedia) })
	iface.RegisterRPCContextTuple("TLBotsDeletePreviewMedia", "/tg.RPCMainMiniBotApps/bots.deletePreviewMedia", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsReorderPreviewMedias", "/tg.RPCMainMiniBotApps/bots.reorderPreviewMedias", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsGetPreviewInfo", "/tg.RPCMainMiniBotApps/bots.getPreviewInfo", func() interface{} { return new(BotsPreviewInfo) })
	iface.RegisterRPCContextTuple("TLBotsGetPreviewMedias", "/tg.RPCMainMiniBotApps/bots.getPreviewMedias", func() interface{} { return new(VectorBotPreviewMedia) })

	// RPCPreparedInlineMessages
	iface.RegisterRPCContextTuple("TLMessagesSavePreparedInlineMessage", "/tg.RPCPreparedInlineMessages/messages.savePreparedInlineMessage", func() interface{} { return new(MessagesBotPreparedInlineMessage) })
	iface.RegisterRPCContextTuple("TLMessagesGetPreparedInlineMessage", "/tg.RPCPreparedInlineMessages/messages.getPreparedInlineMessage", func() interface{} { return new(MessagesPreparedInlineMessage) })

	// RPCGatewayVerificationMessages
	iface.RegisterRPCContextTuple("TLMessagesReportMessagesDelivery", "/tg.RPCGatewayVerificationMessages/messages.reportMessagesDelivery", func() interface{} { return new(Bool) })

	// RPCTodoLists
	iface.RegisterRPCContextTuple("TLMessagesToggleTodoCompleted", "/tg.RPCTodoLists/messages.toggleTodoCompleted", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesAppendTodoList", "/tg.RPCTodoLists/messages.appendTodoList", func() interface{} { return new(Updates) })

	// RPCSuggestedPosts
	iface.RegisterRPCContextTuple("TLMessagesToggleSuggestedPostApproval", "/tg.RPCSuggestedPosts/messages.toggleSuggestedPostApproval", func() interface{} { return new(Updates) })

	// RPCForums
	iface.RegisterRPCContextTuple("TLMessagesGetForumTopics", "/tg.RPCForums/messages.getForumTopics", func() interface{} { return new(MessagesForumTopics) })
	iface.RegisterRPCContextTuple("TLMessagesGetForumTopicsByID", "/tg.RPCForums/messages.getForumTopicsByID", func() interface{} { return new(MessagesForumTopics) })
	iface.RegisterRPCContextTuple("TLMessagesEditForumTopic", "/tg.RPCForums/messages.editForumTopic", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesUpdatePinnedForumTopic", "/tg.RPCForums/messages.updatePinnedForumTopic", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesReorderPinnedForumTopics", "/tg.RPCForums/messages.reorderPinnedForumTopics", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesCreateForumTopic", "/tg.RPCForums/messages.createForumTopic", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLMessagesDeleteTopicHistory", "/tg.RPCForums/messages.deleteTopicHistory", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLChannelsToggleForum", "/tg.RPCForums/channels.toggleForum", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsToggleViewForumAsMessages", "/tg.RPCForums/channels.toggleViewForumAsMessages", func() interface{} { return new(Updates) })

	// RPCUpdates
	iface.RegisterRPCContextTuple("TLUpdatesGetState", "/tg.RPCUpdates/updates.getState", func() interface{} { return new(UpdatesState) })
	iface.RegisterRPCContextTuple("TLUpdatesGetDifference", "/tg.RPCUpdates/updates.getDifference", func() interface{} { return new(UpdatesDifference) })
	iface.RegisterRPCContextTuple("TLUpdatesGetChannelDifference", "/tg.RPCUpdates/updates.getChannelDifference", func() interface{} { return new(UpdatesChannelDifference) })

	// RPCConfiguration
	iface.RegisterRPCContextTuple("TLHelpGetConfig", "/tg.RPCConfiguration/help.getConfig", func() interface{} { return new(Config) })
	iface.RegisterRPCContextTuple("TLHelpGetNearestDc", "/tg.RPCConfiguration/help.getNearestDc", func() interface{} { return new(NearestDc) })
	iface.RegisterRPCContextTuple("TLHelpGetAppUpdate", "/tg.RPCConfiguration/help.getAppUpdate", func() interface{} { return new(HelpAppUpdate) })
	iface.RegisterRPCContextTuple("TLHelpGetInviteText", "/tg.RPCConfiguration/help.getInviteText", func() interface{} { return new(HelpInviteText) })
	iface.RegisterRPCContextTuple("TLHelpGetSupport", "/tg.RPCConfiguration/help.getSupport", func() interface{} { return new(HelpSupport) })
	iface.RegisterRPCContextTuple("TLHelpGetAppConfig", "/tg.RPCConfiguration/help.getAppConfig", func() interface{} { return new(HelpAppConfig) })
	iface.RegisterRPCContextTuple("TLHelpGetSupportName", "/tg.RPCConfiguration/help.getSupportName", func() interface{} { return new(HelpSupportName) })
	iface.RegisterRPCContextTuple("TLHelpDismissSuggestion", "/tg.RPCConfiguration/help.dismissSuggestion", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLHelpGetCountriesList", "/tg.RPCConfiguration/help.getCountriesList", func() interface{} { return new(HelpCountriesList) })

	// RPCInternalBot
	iface.RegisterRPCContextTuple("TLHelpSetBotUpdatesStatus", "/tg.RPCInternalBot/help.setBotUpdatesStatus", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsSendCustomRequest", "/tg.RPCInternalBot/bots.sendCustomRequest", func() interface{} { return new(DataJSON) })
	iface.RegisterRPCContextTuple("TLBotsAnswerWebhookJSONQuery", "/tg.RPCInternalBot/bots.answerWebhookJSONQuery", func() interface{} { return new(Bool) })

	// RPCTos
	iface.RegisterRPCContextTuple("TLHelpGetTermsOfServiceUpdate", "/tg.RPCTos/help.getTermsOfServiceUpdate", func() interface{} { return new(HelpTermsOfServiceUpdate) })
	iface.RegisterRPCContextTuple("TLHelpAcceptTermsOfService", "/tg.RPCTos/help.acceptTermsOfService", func() interface{} { return new(Bool) })

	// RPCMiscellaneous
	iface.RegisterRPCContextTuple("TLHelpSaveAppLog", "/tg.RPCMiscellaneous/help.saveAppLog", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLHelpTest", "/tg.RPCMiscellaneous/help.test", func() interface{} { return new(Bool) })

	// RPCTsf
	iface.RegisterRPCContextTuple("TLHelpGetUserInfo", "/tg.RPCTsf/help.getUserInfo", func() interface{} { return new(HelpUserInfo) })
	iface.RegisterRPCContextTuple("TLHelpEditUserInfo", "/tg.RPCTsf/help.editUserInfo", func() interface{} { return new(HelpUserInfo) })

	// RPCPromoData
	iface.RegisterRPCContextTuple("TLHelpGetPromoData", "/tg.RPCPromoData/help.getPromoData", func() interface{} { return new(HelpPromoData) })
	iface.RegisterRPCContextTuple("TLHelpHidePromoData", "/tg.RPCPromoData/help.hidePromoData", func() interface{} { return new(Bool) })

	// RPCPremium
	iface.RegisterRPCContextTuple("TLHelpGetPremiumPromo", "/tg.RPCPremium/help.getPremiumPromo", func() interface{} { return new(HelpPremiumPromo) })
	iface.RegisterRPCContextTuple("TLPaymentsAssignAppStoreTransaction", "/tg.RPCPremium/payments.assignAppStoreTransaction", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsAssignPlayMarketTransaction", "/tg.RPCPremium/payments.assignPlayMarketTransaction", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsCanPurchaseStore", "/tg.RPCPremium/payments.canPurchaseStore", func() interface{} { return new(Bool) })

	// RPCTimezones
	iface.RegisterRPCContextTuple("TLHelpGetTimezonesList", "/tg.RPCTimezones/help.getTimezonesList", func() interface{} { return new(HelpTimezonesList) })

	// RPCChannels
	iface.RegisterRPCContextTuple("TLChannelsReadHistory", "/tg.RPCChannels/channels.readHistory", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsDeleteMessages", "/tg.RPCChannels/channels.deleteMessages", func() interface{} { return new(MessagesAffectedMessages) })
	iface.RegisterRPCContextTuple("TLChannelsGetMessages", "/tg.RPCChannels/channels.getMessages", func() interface{} { return new(MessagesMessages) })
	iface.RegisterRPCContextTuple("TLChannelsGetParticipants", "/tg.RPCChannels/channels.getParticipants", func() interface{} { return new(ChannelsChannelParticipants) })
	iface.RegisterRPCContextTuple("TLChannelsGetParticipant", "/tg.RPCChannels/channels.getParticipant", func() interface{} { return new(ChannelsChannelParticipant) })
	iface.RegisterRPCContextTuple("TLChannelsGetChannels", "/tg.RPCChannels/channels.getChannels", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLChannelsGetFullChannel", "/tg.RPCChannels/channels.getFullChannel", func() interface{} { return new(MessagesChatFull) })
	iface.RegisterRPCContextTuple("TLChannelsCreateChannel", "/tg.RPCChannels/channels.createChannel", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsEditAdmin", "/tg.RPCChannels/channels.editAdmin", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsEditTitle", "/tg.RPCChannels/channels.editTitle", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsEditPhoto", "/tg.RPCChannels/channels.editPhoto", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsJoinChannel", "/tg.RPCChannels/channels.joinChannel", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsLeaveChannel", "/tg.RPCChannels/channels.leaveChannel", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsInviteToChannel", "/tg.RPCChannels/channels.inviteToChannel", func() interface{} { return new(MessagesInvitedUsers) })
	iface.RegisterRPCContextTuple("TLChannelsDeleteChannel", "/tg.RPCChannels/channels.deleteChannel", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsExportMessageLink", "/tg.RPCChannels/channels.exportMessageLink", func() interface{} { return new(ExportedMessageLink) })
	iface.RegisterRPCContextTuple("TLChannelsToggleSignatures", "/tg.RPCChannels/channels.toggleSignatures", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsGetAdminedPublicChannels", "/tg.RPCChannels/channels.getAdminedPublicChannels", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLChannelsEditBanned", "/tg.RPCChannels/channels.editBanned", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsGetAdminLog", "/tg.RPCChannels/channels.getAdminLog", func() interface{} { return new(ChannelsAdminLogResults) })
	iface.RegisterRPCContextTuple("TLChannelsSetStickers", "/tg.RPCChannels/channels.setStickers", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsReadMessageContents", "/tg.RPCChannels/channels.readMessageContents", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsDeleteHistory", "/tg.RPCChannels/channels.deleteHistory", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsTogglePreHistoryHidden", "/tg.RPCChannels/channels.togglePreHistoryHidden", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsGetGroupsForDiscussion", "/tg.RPCChannels/channels.getGroupsForDiscussion", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLChannelsSetDiscussionGroup", "/tg.RPCChannels/channels.setDiscussionGroup", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsEditLocation", "/tg.RPCChannels/channels.editLocation", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLChannelsToggleSlowMode", "/tg.RPCChannels/channels.toggleSlowMode", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsGetInactiveChannels", "/tg.RPCChannels/channels.getInactiveChannels", func() interface{} { return new(MessagesInactiveChats) })
	iface.RegisterRPCContextTuple("TLChannelsDeleteParticipantHistory", "/tg.RPCChannels/channels.deleteParticipantHistory", func() interface{} { return new(MessagesAffectedHistory) })
	iface.RegisterRPCContextTuple("TLChannelsToggleParticipantsHidden", "/tg.RPCChannels/channels.toggleParticipantsHidden", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsEditCreator", "/tg.RPCChannels/channels.editCreator", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsGetFutureCreatorAfterLeave", "/tg.RPCChannels/channels.getFutureCreatorAfterLeave", func() interface{} { return new(User) })

	// RPCAntiSpam
	iface.RegisterRPCContextTuple("TLChannelsToggleAntiSpam", "/tg.RPCAntiSpam/channels.toggleAntiSpam", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLChannelsReportAntiSpamFalsePositive", "/tg.RPCAntiSpam/channels.reportAntiSpamFalsePositive", func() interface{} { return new(Bool) })

	// RPCChannelRecommendations
	iface.RegisterRPCContextTuple("TLChannelsGetChannelRecommendations", "/tg.RPCChannelRecommendations/channels.getChannelRecommendations", func() interface{} { return new(MessagesChats) })

	// RPCBoosts
	iface.RegisterRPCContextTuple("TLChannelsSetBoostsToUnblockRestrictions", "/tg.RPCBoosts/channels.setBoostsToUnblockRestrictions", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPremiumGetBoostsList", "/tg.RPCBoosts/premium.getBoostsList", func() interface{} { return new(PremiumBoostsList) })
	iface.RegisterRPCContextTuple("TLPremiumGetMyBoosts", "/tg.RPCBoosts/premium.getMyBoosts", func() interface{} { return new(PremiumMyBoosts) })
	iface.RegisterRPCContextTuple("TLPremiumApplyBoost", "/tg.RPCBoosts/premium.applyBoost", func() interface{} { return new(PremiumMyBoosts) })
	iface.RegisterRPCContextTuple("TLPremiumGetBoostsStatus", "/tg.RPCBoosts/premium.getBoostsStatus", func() interface{} { return new(PremiumBoostsStatus) })
	iface.RegisterRPCContextTuple("TLPremiumGetUserBoosts", "/tg.RPCBoosts/premium.getUserBoosts", func() interface{} { return new(PremiumBoostsList) })

	// RPCBots
	iface.RegisterRPCContextTuple("TLBotsSetBotCommands", "/tg.RPCBots/bots.setBotCommands", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsResetBotCommands", "/tg.RPCBots/bots.resetBotCommands", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsGetBotCommands", "/tg.RPCBots/bots.getBotCommands", func() interface{} { return new(VectorBotCommand) })
	iface.RegisterRPCContextTuple("TLBotsSetBotInfo", "/tg.RPCBots/bots.setBotInfo", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsGetBotInfo", "/tg.RPCBots/bots.getBotInfo", func() interface{} { return new(BotsBotInfo) })
	iface.RegisterRPCContextTuple("TLBotsGetAdminedBots", "/tg.RPCBots/bots.getAdminedBots", func() interface{} { return new(VectorUser) })

	// RPCBotMenuButton
	iface.RegisterRPCContextTuple("TLBotsSetBotMenuButton", "/tg.RPCBotMenuButton/bots.setBotMenuButton", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsGetBotMenuButton", "/tg.RPCBotMenuButton/bots.getBotMenuButton", func() interface{} { return new(BotMenuButton) })

	// RPCBotAdminRight
	iface.RegisterRPCContextTuple("TLBotsSetBotBroadcastDefaultAdminRights", "/tg.RPCBotAdminRight/bots.setBotBroadcastDefaultAdminRights", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsSetBotGroupDefaultAdminRights", "/tg.RPCBotAdminRight/bots.setBotGroupDefaultAdminRights", func() interface{} { return new(Bool) })

	// RPCAffiliatePrograms
	iface.RegisterRPCContextTuple("TLBotsUpdateStarRefProgram", "/tg.RPCAffiliatePrograms/bots.updateStarRefProgram", func() interface{} { return new(StarRefProgram) })
	iface.RegisterRPCContextTuple("TLPaymentsGetConnectedStarRefBots", "/tg.RPCAffiliatePrograms/payments.getConnectedStarRefBots", func() interface{} { return new(PaymentsConnectedStarRefBots) })
	iface.RegisterRPCContextTuple("TLPaymentsGetConnectedStarRefBot", "/tg.RPCAffiliatePrograms/payments.getConnectedStarRefBot", func() interface{} { return new(PaymentsConnectedStarRefBots) })
	iface.RegisterRPCContextTuple("TLPaymentsGetSuggestedStarRefBots", "/tg.RPCAffiliatePrograms/payments.getSuggestedStarRefBots", func() interface{} { return new(PaymentsSuggestedStarRefBots) })
	iface.RegisterRPCContextTuple("TLPaymentsConnectStarRefBot", "/tg.RPCAffiliatePrograms/payments.connectStarRefBot", func() interface{} { return new(PaymentsConnectedStarRefBots) })
	iface.RegisterRPCContextTuple("TLPaymentsEditConnectedStarRefBot", "/tg.RPCAffiliatePrograms/payments.editConnectedStarRefBot", func() interface{} { return new(PaymentsConnectedStarRefBots) })

	// RPCBotVerificationIcons
	iface.RegisterRPCContextTuple("TLBotsSetCustomVerification", "/tg.RPCBotVerificationIcons/bots.setCustomVerification", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLBotsGetBotRecommendations", "/tg.RPCBotVerificationIcons/bots.getBotRecommendations", func() interface{} { return new(UsersUsers) })

	// RPCGiveaways
	iface.RegisterRPCContextTuple("TLPaymentsGetPremiumGiftCodeOptions", "/tg.RPCGiveaways/payments.getPremiumGiftCodeOptions", func() interface{} { return new(VectorPremiumGiftCodeOption) })
	iface.RegisterRPCContextTuple("TLPaymentsGetGiveawayInfo", "/tg.RPCGiveaways/payments.getGiveawayInfo", func() interface{} { return new(PaymentsGiveawayInfo) })
	iface.RegisterRPCContextTuple("TLPaymentsLaunchPrepaidGiveaway", "/tg.RPCGiveaways/payments.launchPrepaidGiveaway", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsGiveawayOptions", "/tg.RPCGiveaways/payments.getStarsGiveawayOptions", func() interface{} { return new(VectorStarsGiveawayOption) })

	// RPCGiftCodes
	iface.RegisterRPCContextTuple("TLPaymentsCheckGiftCode", "/tg.RPCGiftCodes/payments.checkGiftCode", func() interface{} { return new(PaymentsCheckedGiftCode) })
	iface.RegisterRPCContextTuple("TLPaymentsApplyGiftCode", "/tg.RPCGiftCodes/payments.applyGiftCode", func() interface{} { return new(Updates) })

	// RPCStars
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsTopupOptions", "/tg.RPCStars/payments.getStarsTopupOptions", func() interface{} { return new(VectorStarsTopupOption) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsStatus", "/tg.RPCStars/payments.getStarsStatus", func() interface{} { return new(PaymentsStarsStatus) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsTransactions", "/tg.RPCStars/payments.getStarsTransactions", func() interface{} { return new(PaymentsStarsStatus) })
	iface.RegisterRPCContextTuple("TLPaymentsSendStarsForm", "/tg.RPCStars/payments.sendStarsForm", func() interface{} { return new(PaymentsPaymentResult) })
	iface.RegisterRPCContextTuple("TLPaymentsRefundStarsCharge", "/tg.RPCStars/payments.refundStarsCharge", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsRevenueStats", "/tg.RPCStars/payments.getStarsRevenueStats", func() interface{} { return new(PaymentsStarsRevenueStats) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsRevenueWithdrawalUrl", "/tg.RPCStars/payments.getStarsRevenueWithdrawalUrl", func() interface{} { return new(PaymentsStarsRevenueWithdrawalUrl) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsRevenueAdsAccountUrl", "/tg.RPCStars/payments.getStarsRevenueAdsAccountUrl", func() interface{} { return new(PaymentsStarsRevenueAdsAccountUrl) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsTransactionsByID", "/tg.RPCStars/payments.getStarsTransactionsByID", func() interface{} { return new(PaymentsStarsStatus) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsGiftOptions", "/tg.RPCStars/payments.getStarsGiftOptions", func() interface{} { return new(VectorStarsGiftOption) })

	// RPCStarSubscriptions
	iface.RegisterRPCContextTuple("TLPaymentsGetStarsSubscriptions", "/tg.RPCStarSubscriptions/payments.getStarsSubscriptions", func() interface{} { return new(PaymentsStarsStatus) })
	iface.RegisterRPCContextTuple("TLPaymentsChangeStarsSubscription", "/tg.RPCStarSubscriptions/payments.changeStarsSubscription", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsFulfillStarsSubscription", "/tg.RPCStarSubscriptions/payments.fulfillStarsSubscription", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsBotCancelStarsSubscription", "/tg.RPCStarSubscriptions/payments.botCancelStarsSubscription", func() interface{} { return new(Bool) })

	// RPCGifts
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGifts", "/tg.RPCGifts/payments.getStarGifts", func() interface{} { return new(PaymentsStarGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsSaveStarGift", "/tg.RPCGifts/payments.saveStarGift", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsConvertStarGift", "/tg.RPCGifts/payments.convertStarGift", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftUpgradePreview", "/tg.RPCGifts/payments.getStarGiftUpgradePreview", func() interface{} { return new(PaymentsStarGiftUpgradePreview) })
	iface.RegisterRPCContextTuple("TLPaymentsUpgradeStarGift", "/tg.RPCGifts/payments.upgradeStarGift", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsTransferStarGift", "/tg.RPCGifts/payments.transferStarGift", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsGetUniqueStarGift", "/tg.RPCGifts/payments.getUniqueStarGift", func() interface{} { return new(PaymentsUniqueStarGift) })
	iface.RegisterRPCContextTuple("TLPaymentsGetSavedStarGifts", "/tg.RPCGifts/payments.getSavedStarGifts", func() interface{} { return new(PaymentsSavedStarGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsGetSavedStarGift", "/tg.RPCGifts/payments.getSavedStarGift", func() interface{} { return new(PaymentsSavedStarGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftWithdrawalUrl", "/tg.RPCGifts/payments.getStarGiftWithdrawalUrl", func() interface{} { return new(PaymentsStarGiftWithdrawalUrl) })
	iface.RegisterRPCContextTuple("TLPaymentsToggleChatStarGiftNotifications", "/tg.RPCGifts/payments.toggleChatStarGiftNotifications", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsToggleStarGiftsPinnedToTop", "/tg.RPCGifts/payments.toggleStarGiftsPinnedToTop", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsGetResaleStarGifts", "/tg.RPCGifts/payments.getResaleStarGifts", func() interface{} { return new(PaymentsResaleStarGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsUpdateStarGiftPrice", "/tg.RPCGifts/payments.updateStarGiftPrice", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsGetUniqueStarGiftValueInfo", "/tg.RPCGifts/payments.getUniqueStarGiftValueInfo", func() interface{} { return new(PaymentsUniqueStarGiftValueInfo) })
	iface.RegisterRPCContextTuple("TLPaymentsCheckCanSendGift", "/tg.RPCGifts/payments.checkCanSendGift", func() interface{} { return new(PaymentsCheckCanSendGiftResult) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftAuctionState", "/tg.RPCGifts/payments.getStarGiftAuctionState", func() interface{} { return new(PaymentsStarGiftAuctionState) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftAuctionAcquiredGifts", "/tg.RPCGifts/payments.getStarGiftAuctionAcquiredGifts", func() interface{} { return new(PaymentsStarGiftAuctionAcquiredGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftActiveAuctions", "/tg.RPCGifts/payments.getStarGiftActiveAuctions", func() interface{} { return new(PaymentsStarGiftActiveAuctions) })
	iface.RegisterRPCContextTuple("TLPaymentsResolveStarGiftOffer", "/tg.RPCGifts/payments.resolveStarGiftOffer", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsSendStarGiftOffer", "/tg.RPCGifts/payments.sendStarGiftOffer", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftUpgradeAttributes", "/tg.RPCGifts/payments.getStarGiftUpgradeAttributes", func() interface{} { return new(PaymentsStarGiftUpgradeAttributes) })
	iface.RegisterRPCContextTuple("TLPaymentsGetCraftStarGifts", "/tg.RPCGifts/payments.getCraftStarGifts", func() interface{} { return new(PaymentsSavedStarGifts) })
	iface.RegisterRPCContextTuple("TLPaymentsCraftStarGift", "/tg.RPCGifts/payments.craftStarGift", func() interface{} { return new(Updates) })

	// RPCGiftCollections
	iface.RegisterRPCContextTuple("TLPaymentsCreateStarGiftCollection", "/tg.RPCGiftCollections/payments.createStarGiftCollection", func() interface{} { return new(StarGiftCollection) })
	iface.RegisterRPCContextTuple("TLPaymentsUpdateStarGiftCollection", "/tg.RPCGiftCollections/payments.updateStarGiftCollection", func() interface{} { return new(StarGiftCollection) })
	iface.RegisterRPCContextTuple("TLPaymentsReorderStarGiftCollections", "/tg.RPCGiftCollections/payments.reorderStarGiftCollections", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsDeleteStarGiftCollection", "/tg.RPCGiftCollections/payments.deleteStarGiftCollection", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPaymentsGetStarGiftCollections", "/tg.RPCGiftCollections/payments.getStarGiftCollections", func() interface{} { return new(PaymentsStarGiftCollections) })

	// RPCGroupCalls
	iface.RegisterRPCContextTuple("TLPhoneCreateGroupCall", "/tg.RPCGroupCalls/phone.createGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneJoinGroupCall", "/tg.RPCGroupCalls/phone.joinGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneLeaveGroupCall", "/tg.RPCGroupCalls/phone.leaveGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneInviteToGroupCall", "/tg.RPCGroupCalls/phone.inviteToGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneDiscardGroupCall", "/tg.RPCGroupCalls/phone.discardGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneToggleGroupCallSettings", "/tg.RPCGroupCalls/phone.toggleGroupCallSettings", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCall", "/tg.RPCGroupCalls/phone.getGroupCall", func() interface{} { return new(PhoneGroupCall) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupParticipants", "/tg.RPCGroupCalls/phone.getGroupParticipants", func() interface{} { return new(PhoneGroupParticipants) })
	iface.RegisterRPCContextTuple("TLPhoneCheckGroupCall", "/tg.RPCGroupCalls/phone.checkGroupCall", func() interface{} { return new(VectorInt) })
	iface.RegisterRPCContextTuple("TLPhoneToggleGroupCallRecord", "/tg.RPCGroupCalls/phone.toggleGroupCallRecord", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneEditGroupCallParticipant", "/tg.RPCGroupCalls/phone.editGroupCallParticipant", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneEditGroupCallTitle", "/tg.RPCGroupCalls/phone.editGroupCallTitle", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCallJoinAs", "/tg.RPCGroupCalls/phone.getGroupCallJoinAs", func() interface{} { return new(PhoneJoinAsPeers) })
	iface.RegisterRPCContextTuple("TLPhoneExportGroupCallInvite", "/tg.RPCGroupCalls/phone.exportGroupCallInvite", func() interface{} { return new(PhoneExportedGroupCallInvite) })
	iface.RegisterRPCContextTuple("TLPhoneToggleGroupCallStartSubscription", "/tg.RPCGroupCalls/phone.toggleGroupCallStartSubscription", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneStartScheduledGroupCall", "/tg.RPCGroupCalls/phone.startScheduledGroupCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneSaveDefaultGroupCallJoinAs", "/tg.RPCGroupCalls/phone.saveDefaultGroupCallJoinAs", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLPhoneJoinGroupCallPresentation", "/tg.RPCGroupCalls/phone.joinGroupCallPresentation", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneLeaveGroupCallPresentation", "/tg.RPCGroupCalls/phone.leaveGroupCallPresentation", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCallStreamChannels", "/tg.RPCGroupCalls/phone.getGroupCallStreamChannels", func() interface{} { return new(PhoneGroupCallStreamChannels) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCallStreamRtmpUrl", "/tg.RPCGroupCalls/phone.getGroupCallStreamRtmpUrl", func() interface{} { return new(PhoneGroupCallStreamRtmpUrl) })
	iface.RegisterRPCContextTuple("TLPhoneSendGroupCallMessage", "/tg.RPCGroupCalls/phone.sendGroupCallMessage", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneDeleteGroupCallMessages", "/tg.RPCGroupCalls/phone.deleteGroupCallMessages", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneDeleteGroupCallParticipantMessages", "/tg.RPCGroupCalls/phone.deleteGroupCallParticipantMessages", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCallStars", "/tg.RPCGroupCalls/phone.getGroupCallStars", func() interface{} { return new(PhoneGroupCallStars) })
	iface.RegisterRPCContextTuple("TLPhoneSaveDefaultSendAs", "/tg.RPCGroupCalls/phone.saveDefaultSendAs", func() interface{} { return new(Bool) })

	// RPCConferenceCalls
	iface.RegisterRPCContextTuple("TLPhoneCreateConferenceCall", "/tg.RPCConferenceCalls/phone.createConferenceCall", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneDeleteConferenceCallParticipants", "/tg.RPCConferenceCalls/phone.deleteConferenceCallParticipants", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneSendConferenceCallBroadcast", "/tg.RPCConferenceCalls/phone.sendConferenceCallBroadcast", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneInviteConferenceCallParticipant", "/tg.RPCConferenceCalls/phone.inviteConferenceCallParticipant", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneDeclineConferenceCallInvite", "/tg.RPCConferenceCalls/phone.declineConferenceCallInvite", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneGetGroupCallChainBlocks", "/tg.RPCConferenceCalls/phone.getGroupCallChainBlocks", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLPhoneSendGroupCallEncryptedMessage", "/tg.RPCConferenceCalls/phone.sendGroupCallEncryptedMessage", func() interface{} { return new(Bool) })

	// RPCLangpack
	iface.RegisterRPCContextTuple("TLLangpackGetLangPack", "/tg.RPCLangpack/langpack.getLangPack", func() interface{} { return new(LangPackDifference) })
	iface.RegisterRPCContextTuple("TLLangpackGetStrings", "/tg.RPCLangpack/langpack.getStrings", func() interface{} { return new(VectorLangPackString) })
	iface.RegisterRPCContextTuple("TLLangpackGetDifference", "/tg.RPCLangpack/langpack.getDifference", func() interface{} { return new(LangPackDifference) })
	iface.RegisterRPCContextTuple("TLLangpackGetLanguages", "/tg.RPCLangpack/langpack.getLanguages", func() interface{} { return new(VectorLangPackLanguage) })
	iface.RegisterRPCContextTuple("TLLangpackGetLanguage", "/tg.RPCLangpack/langpack.getLanguage", func() interface{} { return new(LangPackLanguage) })

	// RPCStatistics
	iface.RegisterRPCContextTuple("TLStatsGetBroadcastStats", "/tg.RPCStatistics/stats.getBroadcastStats", func() interface{} { return new(StatsBroadcastStats) })
	iface.RegisterRPCContextTuple("TLStatsLoadAsyncGraph", "/tg.RPCStatistics/stats.loadAsyncGraph", func() interface{} { return new(StatsGraph) })
	iface.RegisterRPCContextTuple("TLStatsGetMegagroupStats", "/tg.RPCStatistics/stats.getMegagroupStats", func() interface{} { return new(StatsMegagroupStats) })
	iface.RegisterRPCContextTuple("TLStatsGetMessagePublicForwards", "/tg.RPCStatistics/stats.getMessagePublicForwards", func() interface{} { return new(StatsPublicForwards) })
	iface.RegisterRPCContextTuple("TLStatsGetMessageStats", "/tg.RPCStatistics/stats.getMessageStats", func() interface{} { return new(StatsMessageStats) })
	iface.RegisterRPCContextTuple("TLStatsGetStoryStats", "/tg.RPCStatistics/stats.getStoryStats", func() interface{} { return new(StatsStoryStats) })
	iface.RegisterRPCContextTuple("TLStatsGetStoryPublicForwards", "/tg.RPCStatistics/stats.getStoryPublicForwards", func() interface{} { return new(StatsPublicForwards) })

	// RPCStories
	iface.RegisterRPCContextTuple("TLStoriesCanSendStory", "/tg.RPCStories/stories.canSendStory", func() interface{} { return new(StoriesCanSendStoryCount) })
	iface.RegisterRPCContextTuple("TLStoriesSendStory", "/tg.RPCStories/stories.sendStory", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLStoriesEditStory", "/tg.RPCStories/stories.editStory", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLStoriesDeleteStories", "/tg.RPCStories/stories.deleteStories", func() interface{} { return new(VectorInt) })
	iface.RegisterRPCContextTuple("TLStoriesTogglePinned", "/tg.RPCStories/stories.togglePinned", func() interface{} { return new(VectorInt) })
	iface.RegisterRPCContextTuple("TLStoriesGetAllStories", "/tg.RPCStories/stories.getAllStories", func() interface{} { return new(StoriesAllStories) })
	iface.RegisterRPCContextTuple("TLStoriesGetPinnedStories", "/tg.RPCStories/stories.getPinnedStories", func() interface{} { return new(StoriesStories) })
	iface.RegisterRPCContextTuple("TLStoriesGetStoriesArchive", "/tg.RPCStories/stories.getStoriesArchive", func() interface{} { return new(StoriesStories) })
	iface.RegisterRPCContextTuple("TLStoriesGetStoriesByID", "/tg.RPCStories/stories.getStoriesByID", func() interface{} { return new(StoriesStories) })
	iface.RegisterRPCContextTuple("TLStoriesToggleAllStoriesHidden", "/tg.RPCStories/stories.toggleAllStoriesHidden", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesReadStories", "/tg.RPCStories/stories.readStories", func() interface{} { return new(VectorInt) })
	iface.RegisterRPCContextTuple("TLStoriesIncrementStoryViews", "/tg.RPCStories/stories.incrementStoryViews", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesGetStoryViewsList", "/tg.RPCStories/stories.getStoryViewsList", func() interface{} { return new(StoriesStoryViewsList) })
	iface.RegisterRPCContextTuple("TLStoriesGetStoriesViews", "/tg.RPCStories/stories.getStoriesViews", func() interface{} { return new(StoriesStoryViews) })
	iface.RegisterRPCContextTuple("TLStoriesExportStoryLink", "/tg.RPCStories/stories.exportStoryLink", func() interface{} { return new(ExportedStoryLink) })
	iface.RegisterRPCContextTuple("TLStoriesReport", "/tg.RPCStories/stories.report", func() interface{} { return new(ReportResult) })
	iface.RegisterRPCContextTuple("TLStoriesActivateStealthMode", "/tg.RPCStories/stories.activateStealthMode", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLStoriesSendReaction", "/tg.RPCStories/stories.sendReaction", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLStoriesGetPeerStories", "/tg.RPCStories/stories.getPeerStories", func() interface{} { return new(StoriesPeerStories) })
	iface.RegisterRPCContextTuple("TLStoriesGetAllReadPeerStories", "/tg.RPCStories/stories.getAllReadPeerStories", func() interface{} { return new(Updates) })
	iface.RegisterRPCContextTuple("TLStoriesGetPeerMaxIDs", "/tg.RPCStories/stories.getPeerMaxIDs", func() interface{} { return new(VectorRecentStory) })
	iface.RegisterRPCContextTuple("TLStoriesGetChatsToSend", "/tg.RPCStories/stories.getChatsToSend", func() interface{} { return new(MessagesChats) })
	iface.RegisterRPCContextTuple("TLStoriesTogglePeerStoriesHidden", "/tg.RPCStories/stories.togglePeerStoriesHidden", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesGetStoryReactionsList", "/tg.RPCStories/stories.getStoryReactionsList", func() interface{} { return new(StoriesStoryReactionsList) })
	iface.RegisterRPCContextTuple("TLStoriesTogglePinnedToTop", "/tg.RPCStories/stories.togglePinnedToTop", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesSearchPosts", "/tg.RPCStories/stories.searchPosts", func() interface{} { return new(StoriesFoundStories) })
	iface.RegisterRPCContextTuple("TLStoriesCreateAlbum", "/tg.RPCStories/stories.createAlbum", func() interface{} { return new(StoryAlbum) })
	iface.RegisterRPCContextTuple("TLStoriesUpdateAlbum", "/tg.RPCStories/stories.updateAlbum", func() interface{} { return new(StoryAlbum) })
	iface.RegisterRPCContextTuple("TLStoriesReorderAlbums", "/tg.RPCStories/stories.reorderAlbums", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesDeleteAlbum", "/tg.RPCStories/stories.deleteAlbum", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLStoriesGetAlbums", "/tg.RPCStories/stories.getAlbums", func() interface{} { return new(StoriesAlbums) })
	iface.RegisterRPCContextTuple("TLStoriesGetAlbumStories", "/tg.RPCStories/stories.getAlbumStories", func() interface{} { return new(StoriesStories) })
	iface.RegisterRPCContextTuple("TLStoriesStartLive", "/tg.RPCStories/stories.startLive", func() interface{} { return new(Updates) })

	// RPCSmsjobs
	iface.RegisterRPCContextTuple("TLSmsjobsIsEligibleToJoin", "/tg.RPCSmsjobs/smsjobs.isEligibleToJoin", func() interface{} { return new(SmsjobsEligibilityToJoin) })
	iface.RegisterRPCContextTuple("TLSmsjobsJoin", "/tg.RPCSmsjobs/smsjobs.join", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLSmsjobsLeave", "/tg.RPCSmsjobs/smsjobs.leave", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLSmsjobsUpdateSettings", "/tg.RPCSmsjobs/smsjobs.updateSettings", func() interface{} { return new(Bool) })
	iface.RegisterRPCContextTuple("TLSmsjobsGetStatus", "/tg.RPCSmsjobs/smsjobs.getStatus", func() interface{} { return new(SmsjobsStatus) })
	iface.RegisterRPCContextTuple("TLSmsjobsGetSmsJob", "/tg.RPCSmsjobs/smsjobs.getSmsJob", func() interface{} { return new(SmsJob) })
	iface.RegisterRPCContextTuple("TLSmsjobsFinishJob", "/tg.RPCSmsjobs/smsjobs.finishJob", func() interface{} { return new(Bool) })

	// RPCFragmentCollectibles
	iface.RegisterRPCContextTuple("TLFragmentGetCollectibleInfo", "/tg.RPCFragmentCollectibles/fragment.getCollectibleInfo", func() interface{} { return new(FragmentCollectibleInfo) })

	// RPCTest
	iface.RegisterRPCContextTuple("TLTestParseInputAppEvent", "/tg.RPCTest/test.parseInputAppEvent", func() interface{} { return new(InputAppEvent) })

	// RPCPredefined
	iface.RegisterRPCContextTuple("TLPredefinedCreatePredefinedUser", "/tg.RPCPredefined/predefined.createPredefinedUser", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedUpdatePredefinedUsername", "/tg.RPCPredefined/predefined.updatePredefinedUsername", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedUpdatePredefinedProfile", "/tg.RPCPredefined/predefined.updatePredefinedProfile", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedUpdatePredefinedVerified", "/tg.RPCPredefined/predefined.updatePredefinedVerified", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedUpdatePredefinedCode", "/tg.RPCPredefined/predefined.updatePredefinedCode", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedGetPredefinedUser", "/tg.RPCPredefined/predefined.getPredefinedUser", func() interface{} { return new(PredefinedUser) })
	iface.RegisterRPCContextTuple("TLPredefinedGetPredefinedUsers", "/tg.RPCPredefined/predefined.getPredefinedUsers", func() interface{} { return new(VectorPredefinedUser) })

	// RPCBiz
	iface.RegisterRPCContextTuple("TLBizInvokeBizDataRaw", "/tg.RPCBiz/biz.invokeBizDataRaw", func() interface{} { return new(BizDataRaw) })

}
