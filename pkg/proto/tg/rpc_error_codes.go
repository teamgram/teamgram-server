// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

package tg

import (
	"github.com/teamgooo/teamgooo-server/pkg/proto/iface/ecode"
)

// https://core.telegram.org/api/errors
/*
Error handling
There will be errors when working with the API, and they must be correctly handled on the client.

An error is characterized by several parameters:

Error Code
Numerical value similar to HTTP status. Contains information on the type of error that occurred: for example, a data input error, privacy error, or server error. This is a required parameter.

Error Type
A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED. This is an optional parameter.

Error Constructors
There should be a way to handle errors that are returned in rpc_error constructors.

Below is a list of error codes and their meanings:
*/

const (
	/*
		303 SEE_OTHER
		The request must be repeated, but directed to a different data center.

		Examples of Errors:
			FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
			PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
			NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
			USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)

		In all these cases, the error description’s string literal contains the number of the data center (instead of the X) to which the repeated query must be sent.
		More information about redirects between data centers »
	*/
	ErrSeeOther = 303

	/*
		400 BAD_REQUEST
		The query contains errors. In the event that a request was created using a form and contains user generated data, the user should be notified that the data must be corrected before the query is repeated.

		Examples of Errors:
			FIRSTNAME_INVALID: The first name is invalid
			LASTNAME_INVALID: The last name is invalid
			PHONE_NUMBER_INVALID: The phone number is invalid
			PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
			PHONE_CODE_EMPTY: phone_code is missing
			PHONE_CODE_EXPIRED: The confirmation code has expired
			API_ID_INVALID: The api_id/api_hash combination is invalid
			PHONE_NUMBER_OCCUPIED: The phone number is already in use
			PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
			USERS_TOO_FEW: Not enough users (to create a chat, for example)
			USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
			TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
			FILE_PART_INVALID: The file part number is invalid
			FILE_PARTS_INVALID: The number of file parts is invalid
			FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage
			MD5_CHECKSUM_INVALID: The MD5 checksums do not match
			PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
			FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
			FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
			MSG_WAIT_FAILED: A request that must be completed before processing the current request returned an error
			MSG_WAIT_TIMEOUT: A request that must be completed before processing the current request didn't finish processing yet
	*/
	ErrBadRequest = 400

	/*
		401 UNAUTHORIZED
		There was an unauthorized attempt to use functionality available only to authorized users.

		Examples of Errors:
			AUTH_KEY_UNREGISTERED: The key is not registered in the system
			AUTH_KEY_INVALID: The key is invalid
			USER_DEACTIVATED: The user has been deleted/deactivated
			SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
			SESSION_EXPIRED: The authorization has expired
			AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	*/
	ErrUnauthorized = 401

	/*
		403 FORBIDDEN
		Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	*/
	ErrForbidden = 403

	/*
		404 NOT_FOUND
		An attempt to invoke a non-existent object, such as a method.
	*/
	ErrNotFound = 404

	/*
		406 NOT_ACCEPTABLE

		Similar to 400 BAD_REQUEST, but the app must display the error to the user a bit differently.
		Do not display any visible error to the user when receiving the rpc_error constructor: instead, wait for an updateServiceNotification update, and handle it as usual.
		Basically, an updateServiceNotification popup update will be emitted independently (ie NOT as an Updates constructor inside rpc_result but as a normal update) immediately after emission of a 406 rpc_error: the update will contain the actual localized error message to show to the user with a UI popup.

		An exception to this is the AUTH_KEY_DUPLICATED error, which is only emitted if any of the non-media DC detects that an authorized session is sending requests in parallel from two separate TCP connections, from the same or different IP addresses.
		Note that parallel connections are still allowed and actually recommended for media DCs.
		Also note that by session we mean a logged-in session identified by an authorization constructor, fetchable using account.getAuthorizations, not an MTProto session.

		If the client receives an AUTH_KEY_DUPLICATED error, the session was already invalidated by the server and the user must generate a new auth key and login again.
	*/
	ErrNotAcceptable = 406

	/*
		420 FLOOD
		The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded. For example, in an attempt to request a large number of text messages (SMS) for the same phone number.

		Error Example:
			FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	*/
	ErrFlood = 420

	/*
		500 INTERNAL
		An internal server error occurred while a request was being processed; for example, there was a disruption while accessing a database or file storage.

		If a client receives a 500 error, or you believe this error should not have occurred, please collect as much information as possible about the query and error and send it to the developers.
	*/
	ErrInternal = 500

	/*
		Other Error Codes
		If a server returns an error with a code other than the ones listed above, it may be considered the same as a 500 error and treated as an internal server error.
	*/

	// ErrTimeOut503
	// TODO(@benqi): ???, if ecode < 0, panic
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeOut503 = 5030000

	// ErrNotReturnClient ErrNotReturnClient
	ErrNotReturnClient = 700

	// ErrDatabase db error
	ErrDatabase = 600
)

// NewErrFileMigrateX
// 303 SEE_OTHER
//
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X |
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X |
//

// NewErrFileMigrateX
// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
func NewErrFileMigrateX(dc int32) error {
	return ecode.NewCodeErrorf(ErrSeeOther, "FILE_MIGRATE_%d", dc)
}

// NewErrPhoneMigrateX
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X. |
func NewErrPhoneMigrateX(dc int32) error {
	return ecode.NewCodeErrorf(ErrSeeOther, "PHONE_MIGRATE_%d", dc)
}

// NewErrNetworkMigrateX
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X. |
// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
func NewErrNetworkMigrateX(dc int32) error {
	return ecode.NewCodeErrorf(ErrSeeOther, "NETWORK_MIGRATE_%d", dc)
}

// NewErrUserMigrateX
// USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)
func NewErrUserMigrateX(dc int32) error {
	return ecode.NewCodeErrorf(ErrSeeOther, "USER_MIGRATE_%d", dc)
}

// NewErrFloodWaitX 420 FLOOD
//
// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
func NewErrFloodWaitX(second int32) error {
	return ecode.NewCodeErrorf(ErrFlood, "FLOOD_WAIT_%d", second)
}

// NewErrSlowModeWaitX
// | 420 | SLOWMODE_WAIT_X | Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat. |
func NewErrSlowModeWaitX(second int32) error {
	return ecode.NewCodeErrorf(ErrFlood, "SLOWMODE_WAIT_%d", second)
}

// NewErrTakeoutInitDelayX
// | 420 | TAKEOUT_INIT_DELAY_X | Wait X seconds before initing takeout. |
func NewErrTakeoutInitDelayX(second int32) error {
	return ecode.NewCodeErrorf(ErrFlood, "TAKEOUT_INIT_DELAY_%d", second)
}

// NewErr2faConfirmWaitX
// | 420 | 2FA_CONFIRM_WAIT_X | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in X seconds. |
func NewErr2faConfirmWaitX(second int32) error {
	return ecode.NewCodeErrorf(ErrFlood, "2FA_CONFIRM_WAIT_%d", second)
}

// 420 ErrFlood  = 420
var (
	// ErrP0nyFloodwait
	// | 420 | P0NY_FLOODWAIT |   |
	ErrP0nyFloodwait = ecode.NewCodeError(ErrFlood, "P0NY_FLOODWAIT")

	//// ErrSlowmodeWait_%d
	//// | 420 | SLOWMODE_WAIT_%d | Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat. |
	//ErrSlowmodeWait_%d = ecode.NewCodeError(420, "SLOWMODE_WAIT_%d")
	//
	//// ErrTakeoutInitDelay_%d
	//// | 420 | TAKEOUT_INIT_DELAY_%d | Wait %d seconds before initializing takeout. |
	//ErrTakeoutInitDelay_%d = ecode.NewCodeError(420, "TAKEOUT_INIT_DELAY_%d")
	//
	//// Err2faConfirmWait_%d
	//// | 420 | 2FA_CONFIRM_WAIT_%d | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in %d seconds. |
	//Err2faConfirmWait_%d = ecode.NewCodeError(420, "2FA_CONFIRM_WAIT_%d")
)

// 400 BAD_REQUEST

// NewErrFileReferenceX
// | 400 | FILE_REFERENCE_* | The file reference expired, it must be refreshed. |
func NewErrFileReferenceX(second int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "FILE_REFERENCE_%d", second)
}

// NewErrPasswordTooFreshX
// PASSWORD_TOO_FRESH_%d
// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
func NewErrPasswordTooFreshX(second int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "PASSWORD_TOO_FRESH_%d", second)
}

// NewErrEmailUnconfirmedX
// ErrEmailUnconfirmed_%d
// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
func NewErrEmailUnconfirmedX(len int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", len)
}

// NewErrSessionTooFreshX
// ErrSessionTooFresh_%d
// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
func NewErrSessionTooFreshX(second int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "SESSION_TOO_FRESH_%d", second)
}

// NewFilePartXMissing
// FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage. Try repeating the method call to resave the part.
func NewFilePartXMissing(x int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "FILE_PART_%d_MISSING", x)
}

// NewEmailUnconfirmedX
// 400	EMAIL_UNCONFIRMED_X	The provided email isn't confirmed, X is the length of the verification code that was just sent to the email: use account.verifyEmail to enter the received verification code and enable the recovery email.
func NewEmailUnconfirmedX(x int) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", x)
}

var (
	// ErrMethodNotImpl
	// @benqi Add By NebulaChat, not impl the method
	// METHOD_NOT_IMPL: The method not impl
	ErrMethodNotImpl = ecode.NewCodeError(ErrBadRequest, "METHOD_NOT_IMPL")

	// ErrGroupCallParticipantInvalid GROUPCALL_PARTICIPANT_INVALID
	ErrGroupCallParticipantInvalid = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_PARTICIPANT_INVALID")

	// ErrCheckSumInvalid
	// MD5_CHECKSUM_INVALID: The file’s checksum did not match the md5_checksum parameter
	ErrCheckSumInvalid = ecode.NewCodeError(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrGroupCallInvalid GROUPCALL_INVALID
	ErrGroupCallInvalid = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrThemeSlugOccupied
	// THEME_SLUG_OCCUPIED
	ErrThemeSlugOccupied = ecode.NewCodeError(ErrBadRequest, "THEME_SLUG_OCCUPIED")

	// ErrShortnameOccupyFailed
	// | 400 | SHORTNAME_OCCUPY_FAILED | An internal error occurred. |
	ErrShortnameOccupyFailed = ecode.NewCodeError(ErrBadRequest, "SHORTNAME_OCCUPY_FAILED")

	// ErrThemeSlugInvalid
	// THEME_SLUG_INVALID 400 The input theme slug was not valid
	ErrThemeSlugInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_SLUG_INVALID")

	// ErrApiServerNeeded
	// | 400 | API_SERVER_NEEDED | This method be used by api server |
	ErrApiServerNeeded = ecode.NewCodeError(ErrBadRequest, "API_SERVER_NEEDED")

	// ErrInputConstructorInvalid
	// 400	INPUT_CONSTRUCTOR_INVALID	The provided constructor is invalid
	// ErrInputConstructorInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// ErrEncryptedChatIdInvalid
	// | 400 | ENCRYPTED_CHAT_ID_INVALID | The encrypted chat id is invalid |
	ErrEncryptedChatIdInvalid = ecode.NewCodeError(ErrBadRequest, "ENCRYPTED_CHAT_ID_INVALID")

	// ErrAuthTokenAccepted
	// 400 - AUTH_TOKEN_ALREADY_ACCEPTED, the authorization token was already used
	ErrAuthTokenAccepted = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrBotMethodInvalid
	// | 400 | BOT_METHOD_INVALID | This method can't be used by a bot |
	// ErrBotMethodInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_METHOD_INVALID")

	// ErrInputRequestInvalid
	// INPUT_REQUEST_INVALID: The method not impl
	ErrInputRequestInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_REQUEST_INVALID")

	// ErrEnterpriseIsBlocked ErrEnterpriseIsBlocked
	ErrEnterpriseIsBlocked = ecode.NewCodeError(ErrBadRequest, "ERR_ENTERPRISE_IS_BLOCKED")

	// ErrErrBadRequest ErrErrBadRequest
	ErrErrBadRequest = ecode.NewCodeError(ErrBadRequest, "ERR_BAD_REQUEST")

	// ErrTimeTooBig
	// 400	TIME_TOO_BIG
	ErrTimeTooBig = ecode.NewCodeError(ErrBadRequest, "TIME_TOO_BIG")

	// ErrUrlInvalid
	// | 400 | URL_INVALID | Invalid URL provided. |
	ErrUrlInvalid = ecode.NewCodeError(ErrBadRequest, "URL_INVALID")

	// ErrButtonUserPrivacyRestricted
	// | 400 | BUTTON_USER_PRIVACY_RESTRICTED | The privacy setting of the user specified in a [inputKeyboardButtonUserProfile](https://core.telegram.org/constructor/inputKeyboardButtonUserProfile) button do not allow creating such a button. |
	ErrButtonUserPrivacyRestricted = ecode.NewCodeError(ErrBadRequest, "BUTTON_USER_PRIVACY_RESTRICTED")

	// ErrEmojiNotModified
	// | 400 | EMOJI_NOT_MODIFIED | The theme wasn't changed. |
	ErrEmojiNotModified = ecode.NewCodeError(ErrBadRequest, "EMOJI_NOT_MODIFIED")

	// ErrPublicKeyRequired
	// | 400 | PUBLIC_KEY_REQUIRED | A public key is required. |
	ErrPublicKeyRequired = ecode.NewCodeError(ErrBadRequest, "PUBLIC_KEY_REQUIRED")

	// ErrGroupcallJoinMissing
	// | 400 | GROUPCALL_JOIN_MISSING | You haven't joined this group call. |
	ErrGroupcallJoinMissing = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_JOIN_MISSING")

	// ErrPhotoThumbUrlEmpty
	// | 400 | PHOTO_THUMB_URL_EMPTY | Photo thumbnail URL is empty. |
	ErrPhotoThumbUrlEmpty = ecode.NewCodeError(ErrBadRequest, "PHOTO_THUMB_URL_EMPTY")

	// ErrUserIdInvalid
	// | 400 | USER_ID_INVALID | The provided user ID is invalid. |
	ErrUserIdInvalid = ecode.NewCodeError(ErrBadRequest, "USER_ID_INVALID")

	// ErrButtonTextInvalid
	// | 400 | BUTTON_TEXT_INVALID | The specified button text is invalid. |
	ErrButtonTextInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_TEXT_INVALID")

	// ErrCallAlreadyDeclined
	// | 400 | CALL_ALREADY_DECLINED | The call was already declined. |
	ErrCallAlreadyDeclined = ecode.NewCodeError(ErrBadRequest, "CALL_ALREADY_DECLINED")

	// ErrFileReference_%dInvalid
	// | 400 | FILE_REFERENCE_%d_INVALID | The file reference of the media file at index %d in the passed media array is invalid. |
	// ErrFileReference_%dInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_REFERENCE_%d_INVALID")

	// ErrAdminIdInvalid
	// | 400 | ADMIN_ID_INVALID | The specified admin ID is invalid. |
	ErrAdminIdInvalid = ecode.NewCodeError(ErrBadRequest, "ADMIN_ID_INVALID")

	// ErrInvitesTooMuch
	// | 400 | INVITES_TOO_MUCH | The maximum number of per-folder invites specified by the `chatlist_invites_limit_default`/`chatlist_invites_limit_premium` [client configuration parameters &raquo;](https://core.telegram.org/api/config#chatlist-invites-limit-default) was reached. |
	ErrInvitesTooMuch = ecode.NewCodeError(ErrBadRequest, "INVITES_TOO_MUCH")

	// ErrStickerPngDimensions
	// | 400 | STICKER_PNG_DIMENSIONS | Sticker png dimensions invalid. |
	ErrStickerPngDimensions = ecode.NewCodeError(ErrBadRequest, "STICKER_PNG_DIMENSIONS")

	// ErrInlineResultExpired
	// | 400 | INLINE_RESULT_EXPIRED | The inline query expired. |
	ErrInlineResultExpired = ecode.NewCodeError(ErrBadRequest, "INLINE_RESULT_EXPIRED")

	// ErrMessageEmpty
	// | 400 | MESSAGE_EMPTY | The provided message is empty. |
	ErrMessageEmpty = ecode.NewCodeError(ErrBadRequest, "MESSAGE_EMPTY")

	// ErrReplyToUserInvalid
	// | 400 | REPLY_TO_USER_INVALID | The replied-to user is invalid. |
	ErrReplyToUserInvalid = ecode.NewCodeError(ErrBadRequest, "REPLY_TO_USER_INVALID")

	// ErrButtonUrlInvalid
	// | 400 | BUTTON_URL_INVALID | Button URL invalid. |
	ErrButtonUrlInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_URL_INVALID")

	// ErrEntityBoundsInvalid
	// | 400 | ENTITY_BOUNDS_INVALID | A specified [entity offset or length](https://core.telegram.org/api/entities#entity-length) is invalid, see [here &raquo;](https://core.telegram.org/api/entities#entity-length) for info on how to properly compute the entity offset/length. |
	ErrEntityBoundsInvalid = ecode.NewCodeError(ErrBadRequest, "ENTITY_BOUNDS_INVALID")

	// ErrHashtagInvalid
	// | 400 | HASHTAG_INVALID | The specified hashtag is invalid. |
	ErrHashtagInvalid = ecode.NewCodeError(ErrBadRequest, "HASHTAG_INVALID")

	// ErrInputFileInvalid
	// | 400 | INPUT_FILE_INVALID | The specified [InputFile](https://core.telegram.org/type/InputFile) is invalid. |
	ErrInputFileInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_FILE_INVALID")

	// ErrPhoneNumberBanned
	// | 400 | PHONE_NUMBER_BANNED | The provided phone number is banned from telegram. |
	ErrPhoneNumberBanned = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_BANNED")

	// ErrUsernameInvalid
	// | 400 | USERNAME_INVALID | The provided username is not valid. |
	ErrUsernameInvalid = ecode.NewCodeError(ErrBadRequest, "USERNAME_INVALID")

	// ErrPhotoContentUrlEmpty
	// | 400 | PHOTO_CONTENT_URL_EMPTY | Photo URL invalid. |
	ErrPhotoContentUrlEmpty = ecode.NewCodeError(ErrBadRequest, "PHOTO_CONTENT_URL_EMPTY")

	// ErrSearchWithLinkNotSupported
	// | 400 | SEARCH_WITH_LINK_NOT_SUPPORTED | You cannot provide a search query and an invite link at the same time. |
	ErrSearchWithLinkNotSupported = ecode.NewCodeError(ErrBadRequest, "SEARCH_WITH_LINK_NOT_SUPPORTED")

	// ErrBusinessRecipientsEmpty
	// | 400 | BUSINESS_RECIPIENTS_EMPTY | You didn't set any flag in inputBusinessBotRecipients, thus the bot cannot work with *any* peer. |
	ErrBusinessRecipientsEmpty = ecode.NewCodeError(ErrBadRequest, "BUSINESS_RECIPIENTS_EMPTY")

	// ErrChatTitleEmpty
	// | 400 | CHAT_TITLE_EMPTY | No chat title provided. |
	ErrChatTitleEmpty = ecode.NewCodeError(ErrBadRequest, "CHAT_TITLE_EMPTY")

	// ErrMsgIdInvalid
	// | 400 | MSG_ID_INVALID | Invalid message ID provided. |
	ErrMsgIdInvalid = ecode.NewCodeError(ErrBadRequest, "MSG_ID_INVALID")

	// ErrStickersEmpty
	// | 400 | STICKERS_EMPTY | No sticker provided. |
	ErrStickersEmpty = ecode.NewCodeError(ErrBadRequest, "STICKERS_EMPTY")

	// ErrCallPeerInvalid
	// | 400 | CALL_PEER_INVALID | The provided call peer object is invalid. |
	ErrCallPeerInvalid = ecode.NewCodeError(ErrBadRequest, "CALL_PEER_INVALID")

	// ErrForumEnabled
	// | 400 | FORUM_ENABLED | You can't execute the specified action because the group is a [forum](https://core.telegram.org/api/forum), disable forum functionality to continue. |
	ErrForumEnabled = ecode.NewCodeError(ErrBadRequest, "FORUM_ENABLED")

	// ErrGameBotInvalid
	// | 400 | GAME_BOT_INVALID | Bots can't send another bot's game. |
	ErrGameBotInvalid = ecode.NewCodeError(ErrBadRequest, "GAME_BOT_INVALID")

	// ErrBotMethodInvalid
	// | 400 | BOT_METHOD_INVALID | The specified method cannot be used by bots. |
	ErrBotMethodInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_METHOD_INVALID")

	// ErrHashInvalid
	// | 400 | HASH_INVALID | The provided hash is invalid. |
	ErrHashInvalid = ecode.NewCodeError(ErrBadRequest, "HASH_INVALID")

	// ErrMessageIdsEmpty
	// | 400 | MESSAGE_IDS_EMPTY | No message ids were provided. |
	ErrMessageIdsEmpty = ecode.NewCodeError(ErrBadRequest, "MESSAGE_IDS_EMPTY")

	// ErrParticipantIdInvalid
	// | 400 | PARTICIPANT_ID_INVALID | The specified participant ID is invalid. |
	ErrParticipantIdInvalid = ecode.NewCodeError(ErrBadRequest, "PARTICIPANT_ID_INVALID")

	// ErrStickerTgsNodoc
	// | 400 | STICKER_TGS_NODOC | You must send the animated sticker as a document. |
	ErrStickerTgsNodoc = ecode.NewCodeError(ErrBadRequest, "STICKER_TGS_NODOC")

	// ErrPeerFlood
	// | 400 | PEER_FLOOD | The current account is spamreported, you cannot execute this action, check @spambot for more info. |
	ErrPeerFlood = ecode.NewCodeError(ErrBadRequest, "PEER_FLOOD")

	// ErrPhotoInvalid
	// | 400 | PHOTO_INVALID | Photo invalid. |
	ErrPhotoInvalid = ecode.NewCodeError(ErrBadRequest, "PHOTO_INVALID")

	// Err400UserpicUploadRequired
	// | 400 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	Err400UserpicUploadRequired = ecode.NewCodeError(ErrBadRequest, "USERPIC_UPLOAD_REQUIRED")

	// ErrEntitiesTooLong
	// | 400 | ENTITIES_TOO_LONG | You provided too many styled message entities. |
	ErrEntitiesTooLong = ecode.NewCodeError(ErrBadRequest, "ENTITIES_TOO_LONG")

	// ErrPhoneCodeExpired
	// | 400 | PHONE_CODE_EXPIRED | The phone code you provided has expired. |
	ErrPhoneCodeExpired = ecode.NewCodeError(ErrBadRequest, "PHONE_CODE_EXPIRED")

	// ErrPhoneCodeHashEmpty
	// | 400 | PHONE_CODE_HASH_EMPTY | phone_code_hash is missing. |
	ErrPhoneCodeHashEmpty = ecode.NewCodeError(ErrBadRequest, "PHONE_CODE_HASH_EMPTY")

	// ErrVideoFileInvalid
	// | 400 | VIDEO_FILE_INVALID | The specified video file is invalid. |
	ErrVideoFileInvalid = ecode.NewCodeError(ErrBadRequest, "VIDEO_FILE_INVALID")

	// ErrChatInvitePermanent
	// | 400 | CHAT_INVITE_PERMANENT | You can't set an expiration date on permanent invite links. |
	ErrChatInvitePermanent = ecode.NewCodeError(ErrBadRequest, "CHAT_INVITE_PERMANENT")

	// ErrChatlistExcludeInvalid
	// | 400 | CHATLIST_EXCLUDE_INVALID | The specified `exclude_peers` are invalid. |
	ErrChatlistExcludeInvalid = ecode.NewCodeError(ErrBadRequest, "CHATLIST_EXCLUDE_INVALID")

	// ErrStartParamEmpty
	// | 400 | START_PARAM_EMPTY | The start parameter is empty. |
	ErrStartParamEmpty = ecode.NewCodeError(ErrBadRequest, "START_PARAM_EMPTY")

	// ErrAuthTokenExpired
	// | 400 | AUTH_TOKEN_EXPIRED | The authorization token has expired. |
	ErrAuthTokenExpired = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_EXPIRED")

	// ErrMediaTypeInvalid
	// | 400 | MEDIA_TYPE_INVALID | The specified media type cannot be used in stories. |
	ErrMediaTypeInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_TYPE_INVALID")

	// ErrPeerHistoryEmpty
	// | 400 | PEER_HISTORY_EMPTY | You can't pin an empty chat with a user. |
	ErrPeerHistoryEmpty = ecode.NewCodeError(ErrBadRequest, "PEER_HISTORY_EMPTY")

	// ErrTtlDaysInvalid
	// | 400 | TTL_DAYS_INVALID | The provided TTL is invalid. |
	ErrTtlDaysInvalid = ecode.NewCodeError(ErrBadRequest, "TTL_DAYS_INVALID")

	// ErrFilePartSizeChanged
	// | 400 | FILE_PART_SIZE_CHANGED | Provided file part size has changed. |
	ErrFilePartSizeChanged = ecode.NewCodeError(ErrBadRequest, "FILE_PART_SIZE_CHANGED")

	// ErrMessageNotReadYet
	// | 400 | MESSAGE_NOT_READ_YET | The specified message wasn't read yet. |
	ErrMessageNotReadYet = ecode.NewCodeError(ErrBadRequest, "MESSAGE_NOT_READ_YET")

	// ErrReplyToInvalid
	// | 400 | REPLY_TO_INVALID | The specified `reply_to` field is invalid. |
	ErrReplyToInvalid = ecode.NewCodeError(ErrBadRequest, "REPLY_TO_INVALID")

	// ErrTranscriptionFailed
	// | 400 | TRANSCRIPTION_FAILED | Audio transcription failed. |
	ErrTranscriptionFailed = ecode.NewCodeError(ErrBadRequest, "TRANSCRIPTION_FAILED")

	// ErrResultIdDuplicate
	// | 400 | RESULT_ID_DUPLICATE | You provided a duplicate result ID. |
	ErrResultIdDuplicate = ecode.NewCodeError(ErrBadRequest, "RESULT_ID_DUPLICATE")

	// ErrThemeFileInvalid
	// | 400 | THEME_FILE_INVALID | Invalid theme file provided. |
	ErrThemeFileInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_FILE_INVALID")

	// ErrToLangInvalid
	// | 400 | TO_LANG_INVALID | The specified destination language is invalid. |
	ErrToLangInvalid = ecode.NewCodeError(ErrBadRequest, "TO_LANG_INVALID")

	// Err400PhoneNumberInvalid
	// | 400 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	Err400PhoneNumberInvalid = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_INVALID")

	// ErrRangesInvalid
	// | 400 | RANGES_INVALID | Invalid range provided. |
	ErrRangesInvalid = ecode.NewCodeError(ErrBadRequest, "RANGES_INVALID")

	// Err400StickersetInvalid
	// | 400 | STICKERSET_INVALID | The provided sticker set is invalid. |
	Err400StickersetInvalid = ecode.NewCodeError(ErrBadRequest, "STICKERSET_INVALID")

	// ErrBotOnesideNotAvail
	// | 400 | BOT_ONESIDE_NOT_AVAIL | Bots can't pin messages in PM just for themselves. |
	ErrBotOnesideNotAvail = ecode.NewCodeError(ErrBadRequest, "BOT_ONESIDE_NOT_AVAIL")

	// ErrMessageTooOld
	// | 400 | MESSAGE_TOO_OLD | The message is too old, the requested information is not available. |
	ErrMessageTooOld = ecode.NewCodeError(ErrBadRequest, "MESSAGE_TOO_OLD")

	// ErrPasswordRecoveryNa
	// | 400 | PASSWORD_RECOVERY_NA | No email was set, can't recover password via email. |
	ErrPasswordRecoveryNa = ecode.NewCodeError(ErrBadRequest, "PASSWORD_RECOVERY_NA")

	// ErrAudioContentUrlEmpty
	// | 400 | AUDIO_CONTENT_URL_EMPTY | The remote URL specified in the content field is empty. |
	ErrAudioContentUrlEmpty = ecode.NewCodeError(ErrBadRequest, "AUDIO_CONTENT_URL_EMPTY")

	// ErrDateEmpty
	// | 400 | DATE_EMPTY | Date empty. |
	ErrDateEmpty = ecode.NewCodeError(ErrBadRequest, "DATE_EMPTY")

	// ErrStartParamInvalid
	// | 400 | START_PARAM_INVALID | Start parameter invalid. |
	ErrStartParamInvalid = ecode.NewCodeError(ErrBadRequest, "START_PARAM_INVALID")

	// ErrMessageNotModified
	// | 400 | MESSAGE_NOT_MODIFIED | The provided message data is identical to the previous message data, the message wasn't modified. |
	ErrMessageNotModified = ecode.NewCodeError(ErrBadRequest, "MESSAGE_NOT_MODIFIED")

	// ErrPollAnswersInvalid
	// | 400 | POLL_ANSWERS_INVALID | Invalid poll answers were provided. |
	ErrPollAnswersInvalid = ecode.NewCodeError(ErrBadRequest, "POLL_ANSWERS_INVALID")

	// ErrSendMessageMediaInvalid
	// | 400 | SEND_MESSAGE_MEDIA_INVALID | Invalid media provided. |
	ErrSendMessageMediaInvalid = ecode.NewCodeError(ErrBadRequest, "SEND_MESSAGE_MEDIA_INVALID")

	// Err400UserChannelsTooMuch
	// | 400 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	Err400UserChannelsTooMuch = ecode.NewCodeError(ErrBadRequest, "USER_CHANNELS_TOO_MUCH")

	// ErrChatDiscussionUnallowed
	// | 400 | CHAT_DISCUSSION_UNALLOWED | You can't enable forum topics in a discussion group linked to a channel. |
	ErrChatDiscussionUnallowed = ecode.NewCodeError(ErrBadRequest, "CHAT_DISCUSSION_UNALLOWED")

	// ErrConnectionApiIdInvalid
	// | 400 | CONNECTION_API_ID_INVALID | The provided API id is invalid. |
	ErrConnectionApiIdInvalid = ecode.NewCodeError(ErrBadRequest, "CONNECTION_API_ID_INVALID")

	// ErrGraphInvalidReload
	// | 400 | GRAPH_INVALID_RELOAD | Invalid graph token provided, please reload the stats and provide the updated token. |
	ErrGraphInvalidReload = ecode.NewCodeError(ErrBadRequest, "GRAPH_INVALID_RELOAD")

	// ErrHideRequesterMissing
	// | 400 | HIDE_REQUESTER_MISSING | The join request was missing or was already handled. |
	ErrHideRequesterMissing = ecode.NewCodeError(ErrBadRequest, "HIDE_REQUESTER_MISSING")

	// ErrImportIdInvalid
	// | 400 | IMPORT_ID_INVALID | The specified import ID is invalid. |
	ErrImportIdInvalid = ecode.NewCodeError(ErrBadRequest, "IMPORT_ID_INVALID")

	// ErrMediaInvalid
	// | 400 | MEDIA_INVALID | Media invalid. |
	ErrMediaInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_INVALID")

	// ErrOptionsTooMuch
	// | 400 | OPTIONS_TOO_MUCH | Too many options provided. |
	ErrOptionsTooMuch = ecode.NewCodeError(ErrBadRequest, "OPTIONS_TOO_MUCH")

	// ErrPaymentProviderInvalid
	// | 400 | PAYMENT_PROVIDER_INVALID | The specified payment provider is invalid. |
	ErrPaymentProviderInvalid = ecode.NewCodeError(ErrBadRequest, "PAYMENT_PROVIDER_INVALID")

	// ErrBusinessWorkHoursPeriodInvalid
	// | 400 | BUSINESS_WORK_HOURS_PERIOD_INVALID | The specified work hours are invalid, see [here &raquo;](https://core.telegram.org/api/business#opening-hours) for the exact requirements. |
	ErrBusinessWorkHoursPeriodInvalid = ecode.NewCodeError(ErrBadRequest, "BUSINESS_WORK_HOURS_PERIOD_INVALID")

	// ErrChannelTooBig
	// | 400 | CHANNEL_TOO_BIG | This channel has too many participants (>1000) to be deleted. |
	ErrChannelTooBig = ecode.NewCodeError(ErrBadRequest, "CHANNEL_TOO_BIG")

	// ErrEncryptionIdInvalid
	// | 400 | ENCRYPTION_ID_INVALID | The provided secret chat ID is invalid. |
	ErrEncryptionIdInvalid = ecode.NewCodeError(ErrBadRequest, "ENCRYPTION_ID_INVALID")

	// ErrStorySendFloodMonthly_%d
	// | 400 | STORY_SEND_FLOOD_MONTHLY_%d | You've hit the monthly story limit as specified by the [`stories_sent_monthly_limit_*` client configuration parameters](https://core.telegram.org/api/config#stories-sent-monthly-limit-default): wait for the specified number of seconds before posting a new story. |
	// ErrStorySendFloodMonthly_%d = ecode.NewCodeError(ErrBadRequest, "STORY_SEND_FLOOD_MONTHLY_%d")

	// ErrButtonDataInvalid
	// | 400 | BUTTON_DATA_INVALID | The data of one or more of the buttons you provided is invalid. |
	ErrButtonDataInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_DATA_INVALID")

	// ErrMediaPrevInvalid
	// | 400 | MEDIA_PREV_INVALID | Previous media invalid. |
	ErrMediaPrevInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_PREV_INVALID")

	// ErrPackShortNameInvalid
	// | 400 | PACK_SHORT_NAME_INVALID | Short pack name invalid. |
	ErrPackShortNameInvalid = ecode.NewCodeError(ErrBadRequest, "PACK_SHORT_NAME_INVALID")

	// ErrImportFormatUnrecognized
	// | 400 | IMPORT_FORMAT_UNRECOGNIZED | The specified chat export file was exported from an unsupported chat app. |
	ErrImportFormatUnrecognized = ecode.NewCodeError(ErrBadRequest, "IMPORT_FORMAT_UNRECOGNIZED")

	// ErrVideoTitleEmpty
	// | 400 | VIDEO_TITLE_EMPTY | The specified video title is empty. |
	ErrVideoTitleEmpty = ecode.NewCodeError(ErrBadRequest, "VIDEO_TITLE_EMPTY")

	// ErrConnectionLangPackInvalid
	// | 400 | CONNECTION_LANG_PACK_INVALID | The specified language pack is empty. |
	ErrConnectionLangPackInvalid = ecode.NewCodeError(ErrBadRequest, "CONNECTION_LANG_PACK_INVALID")

	// Err400ChatForwardsRestricted
	// | 400 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	Err400ChatForwardsRestricted = ecode.NewCodeError(ErrBadRequest, "CHAT_FORWARDS_RESTRICTED")

	// ErrEmailUnconfirmed
	// | 400 | EMAIL_UNCONFIRMED | Email unconfirmed. |
	ErrEmailUnconfirmed = ecode.NewCodeError(ErrBadRequest, "EMAIL_UNCONFIRMED")

	// ErrFilePartLengthInvalid
	// | 400 | FILE_PART_LENGTH_INVALID | The length of a file part is invalid. |
	ErrFilePartLengthInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_PART_LENGTH_INVALID")

	// ErrWebpushKeyInvalid
	// | 400 | WEBPUSH_KEY_INVALID | The specified web push elliptic curve Diffie-Hellman public key is invalid. |
	ErrWebpushKeyInvalid = ecode.NewCodeError(ErrBadRequest, "WEBPUSH_KEY_INVALID")

	// ErrChannelsTooMuch
	// | 400 | CHANNELS_TOO_MUCH | You have joined too many channels/supergroups. |
	ErrChannelsTooMuch = ecode.NewCodeError(ErrBadRequest, "CHANNELS_TOO_MUCH")

	// ErrFileReference_%dExpired
	// | 400 | FILE_REFERENCE_%d_EXPIRED | The file reference of the media file at index %d in the passed media array expired, it [must be refreshed](https://core.telegram.org/api/file_reference). |
	// ErrFileReference_%dExpired = ecode.NewCodeError(ErrBadRequest, "FILE_REFERENCE_%d_EXPIRED")

	// ErrPhoneNotOccupied
	// | 400 | PHONE_NOT_OCCUPIED | No user is associated to the specified phone number. |
	ErrPhoneNotOccupied = ecode.NewCodeError(ErrBadRequest, "PHONE_NOT_OCCUPIED")

	// ErrExpireDateInvalid
	// | 400 | EXPIRE_DATE_INVALID | The specified expiration date is invalid. |
	ErrExpireDateInvalid = ecode.NewCodeError(ErrBadRequest, "EXPIRE_DATE_INVALID")

	// ErrMethodInvalid
	// | 400 | METHOD_INVALID | The specified method is invalid. |
	ErrMethodInvalid = ecode.NewCodeError(ErrBadRequest, "METHOD_INVALID")

	// ErrPhoneNumberUnoccupied
	// | 400 | PHONE_NUMBER_UNOCCUPIED | The phone number is not yet being used. |
	ErrPhoneNumberUnoccupied = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_UNOCCUPIED")

	// ErrSettingsInvalid
	// | 400 | SETTINGS_INVALID | Invalid settings were provided. |
	ErrSettingsInvalid = ecode.NewCodeError(ErrBadRequest, "SETTINGS_INVALID")

	// ErrTempAuthKeyEmpty
	// | 400 | TEMP_AUTH_KEY_EMPTY | No temporary auth key provided. |
	ErrTempAuthKeyEmpty = ecode.NewCodeError(ErrBadRequest, "TEMP_AUTH_KEY_EMPTY")

	// ErrInputFetchError
	// | 400 | INPUT_FETCH_ERROR | An error occurred while parsing the provided TL constructor. |
	ErrInputFetchError = ecode.NewCodeError(ErrBadRequest, "INPUT_FETCH_ERROR")

	// ErrFileTokenInvalid
	// | 400 | FILE_TOKEN_INVALID | The master DC did not accept the `file_token` (e.g., the token has expired). Continue downloading the file from the master DC using upload.getFile. |
	ErrFileTokenInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_TOKEN_INVALID")

	// ErrFilterIdInvalid
	// | 400 | FILTER_ID_INVALID | The specified filter ID is invalid. |
	ErrFilterIdInvalid = ecode.NewCodeError(ErrBadRequest, "FILTER_ID_INVALID")

	// ErrPinRestricted
	// | 400 | PIN_RESTRICTED | You can't pin messages. |
	ErrPinRestricted = ecode.NewCodeError(ErrBadRequest, "PIN_RESTRICTED")

	// ErrPhoneHashExpired
	// | 400 | PHONE_HASH_EXPIRED | An invalid or expired `phone_code_hash` was provided. |
	ErrPhoneHashExpired = ecode.NewCodeError(ErrBadRequest, "PHONE_HASH_EXPIRED")

	// ErrChatAboutTooLong
	// | 400 | CHAT_ABOUT_TOO_LONG | Chat about too long. |
	ErrChatAboutTooLong = ecode.NewCodeError(ErrBadRequest, "CHAT_ABOUT_TOO_LONG")

	// ErrChatLinkExists
	// | 400 | CHAT_LINK_EXISTS | The chat is public, you can't hide the history to new users. |
	ErrChatLinkExists = ecode.NewCodeError(ErrBadRequest, "CHAT_LINK_EXISTS")

	// ErrConnectionLayerInvalid
	// | 400 | CONNECTION_LAYER_INVALID | Layer invalid. |
	ErrConnectionLayerInvalid = ecode.NewCodeError(ErrBadRequest, "CONNECTION_LAYER_INVALID")

	// ErrDataJsonInvalid
	// | 400 | DATA_JSON_INVALID | The provided JSON data is invalid. |
	ErrDataJsonInvalid = ecode.NewCodeError(ErrBadRequest, "DATA_JSON_INVALID")

	// ErrStickerFileInvalid
	// | 400 | STICKER_FILE_INVALID | Sticker file invalid. |
	ErrStickerFileInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_FILE_INVALID")

	// ErrStoryNotModified
	// | 400 | STORY_NOT_MODIFIED | The new story information you passed is equal to the previous story information, thus it wasn't modified. |
	ErrStoryNotModified = ecode.NewCodeError(ErrBadRequest, "STORY_NOT_MODIFIED")

	// ErrBotResponseTimeout
	// | 400 | BOT_RESPONSE_TIMEOUT | A timeout occurred while fetching data from the bot. |
	ErrBotResponseTimeout = ecode.NewCodeError(ErrBadRequest, "BOT_RESPONSE_TIMEOUT")

	// ErrFileTitleEmpty
	// | 400 | FILE_TITLE_EMPTY | An empty file title was specified. |
	ErrFileTitleEmpty = ecode.NewCodeError(ErrBadRequest, "FILE_TITLE_EMPTY")

	// ErrImportFormatDateInvalid
	// | 400 | IMPORT_FORMAT_DATE_INVALID | The date specified in the import file is invalid. |
	ErrImportFormatDateInvalid = ecode.NewCodeError(ErrBadRequest, "IMPORT_FORMAT_DATE_INVALID")

	// ErrScoreInvalid
	// | 400 | SCORE_INVALID | The specified game score is invalid. |
	ErrScoreInvalid = ecode.NewCodeError(ErrBadRequest, "SCORE_INVALID")

	// ErrSmsjobIdInvalid
	// | 400 | SMSJOB_ID_INVALID | The specified job ID is invalid. |
	ErrSmsjobIdInvalid = ecode.NewCodeError(ErrBadRequest, "SMSJOB_ID_INVALID")

	// ErrVideoPauseForbidden
	// | 400 | VIDEO_PAUSE_FORBIDDEN | You cannot pause the video stream. |
	ErrVideoPauseForbidden = ecode.NewCodeError(ErrBadRequest, "VIDEO_PAUSE_FORBIDDEN")

	// ErrFileContentTypeInvalid
	// | 400 | FILE_CONTENT_TYPE_INVALID | File content-type is invalid. |
	ErrFileContentTypeInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_CONTENT_TYPE_INVALID")

	// ErrImportTokenInvalid
	// | 400 | IMPORT_TOKEN_INVALID | The specified token is invalid. |
	ErrImportTokenInvalid = ecode.NewCodeError(ErrBadRequest, "IMPORT_TOKEN_INVALID")

	// ErrReactionsTooMany
	// | 400 | REACTIONS_TOO_MANY | The message already has exactly `reactions_uniq_max` reaction emojis, you can't react with a new emoji, see [the docs for more info &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrReactionsTooMany = ecode.NewCodeError(ErrBadRequest, "REACTIONS_TOO_MANY")

	// ErrBusinessWorkHoursEmpty
	// | 400 | BUSINESS_WORK_HOURS_EMPTY | No work hours were specified. |
	ErrBusinessWorkHoursEmpty = ecode.NewCodeError(ErrBadRequest, "BUSINESS_WORK_HOURS_EMPTY")

	// ErrContactIdInvalid
	// | 400 | CONTACT_ID_INVALID | The provided contact ID is invalid. |
	ErrContactIdInvalid = ecode.NewCodeError(ErrBadRequest, "CONTACT_ID_INVALID")

	// Err400TopicDeleted
	// | 400 | TOPIC_DELETED | The specified topic was deleted. |
	Err400TopicDeleted = ecode.NewCodeError(ErrBadRequest, "TOPIC_DELETED")

	// Err400UserIsBlocked
	// | 400 | USER_IS_BLOCKED | You were blocked by this user. |
	Err400UserIsBlocked = ecode.NewCodeError(ErrBadRequest, "USER_IS_BLOCKED")

	// ErrAudioTitleEmpty
	// | 400 | AUDIO_TITLE_EMPTY | An empty audio title was provided. |
	ErrAudioTitleEmpty = ecode.NewCodeError(ErrBadRequest, "AUDIO_TITLE_EMPTY")

	// ErrCodeHashInvalid
	// | 400 | CODE_HASH_INVALID | Code hash invalid. |
	ErrCodeHashInvalid = ecode.NewCodeError(ErrBadRequest, "CODE_HASH_INVALID")

	// ErrPeersListEmpty
	// | 400 | PEERS_LIST_EMPTY | The specified list of peers is empty. |
	ErrPeersListEmpty = ecode.NewCodeError(ErrBadRequest, "PEERS_LIST_EMPTY")

	// ErrEmailUnconfirmed_%d
	// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
	// ErrEmailUnconfirmed_%d = ecode.NewCodeError(ErrBadRequest, "EMAIL_UNCONFIRMED_%d")

	// ErrInputUserDeactivated
	// | 400 | INPUT_USER_DEACTIVATED | The specified user was deleted. |
	ErrInputUserDeactivated = ecode.NewCodeError(ErrBadRequest, "INPUT_USER_DEACTIVATED")

	// ErrPhotoIdInvalid
	// | 400 | PHOTO_ID_INVALID | Photo ID invalid. |
	ErrPhotoIdInvalid = ecode.NewCodeError(ErrBadRequest, "PHOTO_ID_INVALID")

	// ErrSmsCodeCreateFailed
	// | 400 | SMS_CODE_CREATE_FAILED | An error occurred while creating the SMS code. |
	ErrSmsCodeCreateFailed = ecode.NewCodeError(ErrBadRequest, "SMS_CODE_CREATE_FAILED")

	// ErrDocumentInvalid
	// | 400 | DOCUMENT_INVALID | The specified document is invalid. |
	ErrDocumentInvalid = ecode.NewCodeError(ErrBadRequest, "DOCUMENT_INVALID")

	// ErrFileIdInvalid
	// | 400 | FILE_ID_INVALID | The provided file id is invalid. |
	ErrFileIdInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_ID_INVALID")

	// Err400FreshChangeAdminsForbidden
	// | 400 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	Err400FreshChangeAdminsForbidden = ecode.NewCodeError(ErrBadRequest, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// ErrChannelParicipantMissing
	// | 400 | CHANNEL_PARICIPANT_MISSING | The current user is not in the channel. |
	ErrChannelParicipantMissing = ecode.NewCodeError(ErrBadRequest, "CHANNEL_PARICIPANT_MISSING")

	// ErrGraphOutdatedReload
	// | 400 | GRAPH_OUTDATED_RELOAD | The graph is outdated, please get a new async token using stats.getBroadcastStats. |
	ErrGraphOutdatedReload = ecode.NewCodeError(ErrBadRequest, "GRAPH_OUTDATED_RELOAD")

	// ErrInviteSlugEmpty
	// | 400 | INVITE_SLUG_EMPTY | The specified invite slug is empty. |
	ErrInviteSlugEmpty = ecode.NewCodeError(ErrBadRequest, "INVITE_SLUG_EMPTY")

	// ErrStickerGifDimensions
	// | 400 | STICKER_GIF_DIMENSIONS | The specified video sticker has invalid dimensions. |
	ErrStickerGifDimensions = ecode.NewCodeError(ErrBadRequest, "STICKER_GIF_DIMENSIONS")

	// ErrWallpaperFileInvalid
	// | 400 | WALLPAPER_FILE_INVALID | The specified wallpaper file is invalid. |
	ErrWallpaperFileInvalid = ecode.NewCodeError(ErrBadRequest, "WALLPAPER_FILE_INVALID")

	// Err400AddressInvalid
	// | 400 | ADDRESS_INVALID | The specified geopoint address is invalid. |
	Err400AddressInvalid = ecode.NewCodeError(ErrBadRequest, "ADDRESS_INVALID")

	// ErrBalanceTooLow
	// | 400 | BALANCE_TOO_LOW | The transaction cannot be completed because the current [Telegram Stars balance](https://core.telegram.org/api/stars) is too low. |
	ErrBalanceTooLow = ecode.NewCodeError(ErrBadRequest, "BALANCE_TOO_LOW")

	// ErrBoostNotModified
	// | 400 | BOOST_NOT_MODIFIED | You're already [boosting](https://core.telegram.org/api/boost) the specified channel. |
	ErrBoostNotModified = ecode.NewCodeError(ErrBadRequest, "BOOST_NOT_MODIFIED")

	// ErrTokenInvalid
	// | 400 | TOKEN_INVALID | The provided token is invalid. |
	ErrTokenInvalid = ecode.NewCodeError(ErrBadRequest, "TOKEN_INVALID")

	// ErrBroadcastRequired
	// | 400 | BROADCAST_REQUIRED | This method can only be called on a channel, please use stats.getMegagroupStats for supergroups. |
	ErrBroadcastRequired = ecode.NewCodeError(ErrBadRequest, "BROADCAST_REQUIRED")

	// ErrInputChatlistInvalid
	// | 400 | INPUT_CHATLIST_INVALID | The specified folder is invalid. |
	ErrInputChatlistInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_CHATLIST_INVALID")

	// ErrMultiMediaTooLong
	// | 400 | MULTI_MEDIA_TOO_LONG | Too many media files for album. |
	ErrMultiMediaTooLong = ecode.NewCodeError(ErrBadRequest, "MULTI_MEDIA_TOO_LONG")

	// ErrFolderIdEmpty
	// | 400 | FOLDER_ID_EMPTY | An empty folder ID was specified. |
	ErrFolderIdEmpty = ecode.NewCodeError(ErrBadRequest, "FOLDER_ID_EMPTY")

	// ErrFromMessageBotDisabled
	// | 400 | FROM_MESSAGE_BOT_DISABLED | Bots can't use fromMessage min constructors. |
	ErrFromMessageBotDisabled = ecode.NewCodeError(ErrBadRequest, "FROM_MESSAGE_BOT_DISABLED")

	// ErrInviteHashEmpty
	// | 400 | INVITE_HASH_EMPTY | The invite hash is empty. |
	ErrInviteHashEmpty = ecode.NewCodeError(ErrBadRequest, "INVITE_HASH_EMPTY")

	// ErrRevoteNotAllowed
	// | 400 | REVOTE_NOT_ALLOWED | You cannot change your vote. |
	ErrRevoteNotAllowed = ecode.NewCodeError(ErrBadRequest, "REVOTE_NOT_ALLOWED")

	// ErrSlotsEmpty
	// | 400 | SLOTS_EMPTY | The specified slot list is empty. |
	ErrSlotsEmpty = ecode.NewCodeError(ErrBadRequest, "SLOTS_EMPTY")

	// ErrArticleTitleEmpty
	// | 400 | ARTICLE_TITLE_EMPTY | The title of the article is empty. |
	ErrArticleTitleEmpty = ecode.NewCodeError(ErrBadRequest, "ARTICLE_TITLE_EMPTY")

	// ErrBotBusinessMissing
	// | 400 | BOT_BUSINESS_MISSING | The specified bot is not a business bot (the [user](https://core.telegram.org/constructor/user).`bot_business` flag is not set). |
	ErrBotBusinessMissing = ecode.NewCodeError(ErrBadRequest, "BOT_BUSINESS_MISSING")

	// ErrChargeAlreadyRefunded
	// | 400 | CHARGE_ALREADY_REFUNDED | The transaction was already refunded. |
	ErrChargeAlreadyRefunded = ecode.NewCodeError(ErrBadRequest, "CHARGE_ALREADY_REFUNDED")

	// ErrVideoContentTypeInvalid
	// | 400 | VIDEO_CONTENT_TYPE_INVALID | The video's content type is invalid. |
	ErrVideoContentTypeInvalid = ecode.NewCodeError(ErrBadRequest, "VIDEO_CONTENT_TYPE_INVALID")

	// ErrInputFetchFail
	// | 400 | INPUT_FETCH_FAIL | An error occurred while parsing the provided TL constructor. |
	ErrInputFetchFail = ecode.NewCodeError(ErrBadRequest, "INPUT_FETCH_FAIL")

	// ErrOptionInvalid
	// | 400 | OPTION_INVALID | Invalid option selected. |
	ErrOptionInvalid = ecode.NewCodeError(ErrBadRequest, "OPTION_INVALID")

	// ErrQuizAnswerMissing
	// | 400 | QUIZ_ANSWER_MISSING | You can forward a quiz while hiding the original author only after choosing an option in the quiz. |
	ErrQuizAnswerMissing = ecode.NewCodeError(ErrBadRequest, "QUIZ_ANSWER_MISSING")

	// ErrWcConvertUrlInvalid
	// | 400 | WC_CONVERT_URL_INVALID | WC convert URL invalid. |
	ErrWcConvertUrlInvalid = ecode.NewCodeError(ErrBadRequest, "WC_CONVERT_URL_INVALID")

	// ErrUsersTooMuch
	// | 400 | USERS_TOO_MUCH | The maximum number of users has been exceeded (to create a chat, for example). |
	ErrUsersTooMuch = ecode.NewCodeError(ErrBadRequest, "USERS_TOO_MUCH")

	// ErrAutoarchiveNotAvailable
	// | 400 | AUTOARCHIVE_NOT_AVAILABLE | The autoarchive setting is not available at this time: please check the value of the [autoarchive_setting_available field in client config &raquo;](https://core.telegram.org/api/config#client-configuration) before calling this method. |
	ErrAutoarchiveNotAvailable = ecode.NewCodeError(ErrBadRequest, "AUTOARCHIVE_NOT_AVAILABLE")

	// ErrCodeEmpty
	// | 400 | CODE_EMPTY | The provided code is empty. |
	ErrCodeEmpty = ecode.NewCodeError(ErrBadRequest, "CODE_EMPTY")

	// ErrDhGAInvalid
	// | 400 | DH_G_A_INVALID | g_a invalid. |
	ErrDhGAInvalid = ecode.NewCodeError(ErrBadRequest, "DH_G_A_INVALID")

	// ErrChatRevokeDateUnsupported
	// | 400 | CHAT_REVOKE_DATE_UNSUPPORTED | `min_date` and `max_date` are not available for using with non-user peers. |
	ErrChatRevokeDateUnsupported = ecode.NewCodeError(ErrBadRequest, "CHAT_REVOKE_DATE_UNSUPPORTED")

	// ErrQuizMultipleInvalid
	// | 400 | QUIZ_MULTIPLE_INVALID | Quizzes can't have the multiple_choice flag set! |
	ErrQuizMultipleInvalid = ecode.NewCodeError(ErrBadRequest, "QUIZ_MULTIPLE_INVALID")

	// ErrWebpageMediaEmpty
	// | 400 | WEBPAGE_MEDIA_EMPTY | Webpage media empty. |
	ErrWebpageMediaEmpty = ecode.NewCodeError(ErrBadRequest, "WEBPAGE_MEDIA_EMPTY")

	// ErrRsaDecryptFailed
	// | 400 | RSA_DECRYPT_FAILED | Internal RSA decryption failed. |
	ErrRsaDecryptFailed = ecode.NewCodeError(ErrBadRequest, "RSA_DECRYPT_FAILED")

	// ErrTakeoutInvalid
	// | 400 | TAKEOUT_INVALID | The specified takeout ID is invalid. |
	ErrTakeoutInvalid = ecode.NewCodeError(ErrBadRequest, "TAKEOUT_INVALID")

	// ErrUsageLimitInvalid
	// | 400 | USAGE_LIMIT_INVALID | The specified usage limit is invalid. |
	ErrUsageLimitInvalid = ecode.NewCodeError(ErrBadRequest, "USAGE_LIMIT_INVALID")

	// ErrWebpageCurlFailed
	// | 400 | WEBPAGE_CURL_FAILED | Failure while fetching the webpage with cURL. |
	ErrWebpageCurlFailed = ecode.NewCodeError(ErrBadRequest, "WEBPAGE_CURL_FAILED")

	// ErrBotDomainInvalid
	// | 400 | BOT_DOMAIN_INVALID | Bot domain invalid. |
	ErrBotDomainInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_DOMAIN_INVALID")

	// ErrButtonTypeInvalid
	// | 400 | BUTTON_TYPE_INVALID | The type of one or more of the buttons you provided is invalid. |
	ErrButtonTypeInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_TYPE_INVALID")

	// ErrChatAboutNotModified
	// | 400 | CHAT_ABOUT_NOT_MODIFIED | About text has not changed. |
	ErrChatAboutNotModified = ecode.NewCodeError(ErrBadRequest, "CHAT_ABOUT_NOT_MODIFIED")

	// ErrCallAlreadyAccepted
	// | 400 | CALL_ALREADY_ACCEPTED | The call was already accepted. |
	ErrCallAlreadyAccepted = ecode.NewCodeError(ErrBadRequest, "CALL_ALREADY_ACCEPTED")

	// ErrDcIdInvalid
	// | 400 | DC_ID_INVALID | The provided DC ID is invalid. |
	ErrDcIdInvalid = ecode.NewCodeError(ErrBadRequest, "DC_ID_INVALID")

	// ErrPrivacyValueInvalid
	// | 400 | PRIVACY_VALUE_INVALID | The specified privacy rule combination is invalid. |
	ErrPrivacyValueInvalid = ecode.NewCodeError(ErrBadRequest, "PRIVACY_VALUE_INVALID")

	// ErrEmailInvalid
	// | 400 | EMAIL_INVALID | The specified email is invalid. |
	ErrEmailInvalid = ecode.NewCodeError(ErrBadRequest, "EMAIL_INVALID")

	// ErrStorySendFloodWeekly_%d
	// | 400 | STORY_SEND_FLOOD_WEEKLY_%d | You've hit the weekly story limit as specified by the [`stories_sent_weekly_limit_*` client configuration parameters](https://core.telegram.org/api/config#stories-sent-weekly-limit-default): wait for the specified number of seconds before posting a new story. |
	// ErrStorySendFloodWeekly_%d = ecode.NewCodeError(ErrBadRequest, "STORY_SEND_FLOOD_WEEKLY_%d")

	// ErrTimezoneInvalid
	// | 400 | TIMEZONE_INVALID | The specified timezone does not exist. |
	ErrTimezoneInvalid = ecode.NewCodeError(ErrBadRequest, "TIMEZONE_INVALID")

	// ErrShortNameOccupied
	// | 400 | SHORT_NAME_OCCUPIED | The specified short name is already in use. |
	ErrShortNameOccupied = ecode.NewCodeError(ErrBadRequest, "SHORT_NAME_OCCUPIED")

	// ErrMediaCaptionTooLong
	// | 400 | MEDIA_CAPTION_TOO_LONG | The caption is too long. |
	ErrMediaCaptionTooLong = ecode.NewCodeError(ErrBadRequest, "MEDIA_CAPTION_TOO_LONG")

	// ErrPersistentTimestampInvalid
	// | 400 | PERSISTENT_TIMESTAMP_INVALID | Persistent timestamp invalid. |
	ErrPersistentTimestampInvalid = ecode.NewCodeError(ErrBadRequest, "PERSISTENT_TIMESTAMP_INVALID")

	// ErrScheduleTooMuch
	// | 400 | SCHEDULE_TOO_MUCH | There are too many scheduled messages. |
	ErrScheduleTooMuch = ecode.NewCodeError(ErrBadRequest, "SCHEDULE_TOO_MUCH")

	// ErrBoostPeerInvalid
	// | 400 | BOOST_PEER_INVALID | The specified `boost_peer` is invalid. |
	ErrBoostPeerInvalid = ecode.NewCodeError(ErrBadRequest, "BOOST_PEER_INVALID")

	// ErrQuoteTextInvalid
	// | 400 | QUOTE_TEXT_INVALID | The specified `reply_to`.`quote_text` field is invalid. |
	ErrQuoteTextInvalid = ecode.NewCodeError(ErrBadRequest, "QUOTE_TEXT_INVALID")

	// ErrContactAddMissing
	// | 400 | CONTACT_ADD_MISSING | Contact to add is missing. |
	ErrContactAddMissing = ecode.NewCodeError(ErrBadRequest, "CONTACT_ADD_MISSING")

	// ErrPhoneCodeInvalid
	// | 400 | PHONE_CODE_INVALID | The provided phone code is invalid. |
	ErrPhoneCodeInvalid = ecode.NewCodeError(ErrBadRequest, "PHONE_CODE_INVALID")

	// ErrTitleInvalid
	// | 400 | TITLE_INVALID | The specified stickerpack title is invalid. |
	ErrTitleInvalid = ecode.NewCodeError(ErrBadRequest, "TITLE_INVALID")

	// Err400BannedRightsInvalid
	// | 400 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	Err400BannedRightsInvalid = ecode.NewCodeError(ErrBadRequest, "BANNED_RIGHTS_INVALID")

	// ErrCustomReactionsTooMany
	// | 400 | CUSTOM_REACTIONS_TOO_MANY | Too many custom reactions were specified. |
	ErrCustomReactionsTooMany = ecode.NewCodeError(ErrBadRequest, "CUSTOM_REACTIONS_TOO_MANY")

	// ErrShortNameInvalid
	// | 400 | SHORT_NAME_INVALID | The specified short name is invalid. |
	ErrShortNameInvalid = ecode.NewCodeError(ErrBadRequest, "SHORT_NAME_INVALID")

	// ErrMessageIdInvalid
	// | 400 | MESSAGE_ID_INVALID | The provided message id is invalid. |
	ErrMessageIdInvalid = ecode.NewCodeError(ErrBadRequest, "MESSAGE_ID_INVALID")

	// ErrPeerIdNotSupported
	// | 400 | PEER_ID_NOT_SUPPORTED | The provided peer ID is not supported. |
	ErrPeerIdNotSupported = ecode.NewCodeError(ErrBadRequest, "PEER_ID_NOT_SUPPORTED")

	// ErrInputRequestTooLong
	// | 400 | INPUT_REQUEST_TOO_LONG | The request payload is too long. |
	ErrInputRequestTooLong = ecode.NewCodeError(ErrBadRequest, "INPUT_REQUEST_TOO_LONG")

	// ErrChatTooBig
	// | 400 | CHAT_TOO_BIG | This method is not available for groups with more than `chat_read_mark_size_threshold` members, [see client configuration &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrChatTooBig = ecode.NewCodeError(ErrBadRequest, "CHAT_TOO_BIG")

	// ErrFilePartsInvalid
	// | 400 | FILE_PARTS_INVALID | The number of file parts is invalid. |
	ErrFilePartsInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_PARTS_INVALID")

	// ErrMegagroupGeoRequired
	// | 400 | MEGAGROUP_GEO_REQUIRED | This method can only be invoked on a geogroup. |
	ErrMegagroupGeoRequired = ecode.NewCodeError(ErrBadRequest, "MEGAGROUP_GEO_REQUIRED")

	// ErrUserBannedInChannel
	// | 400 | USER_BANNED_IN_CHANNEL | You're banned from sending messages in supergroups/channels. |
	ErrUserBannedInChannel = ecode.NewCodeError(ErrBadRequest, "USER_BANNED_IN_CHANNEL")

	// ErrWebdocumentSizeTooBig
	// | 400 | WEBDOCUMENT_SIZE_TOO_BIG | Webdocument is too big! |
	ErrWebdocumentSizeTooBig = ecode.NewCodeError(ErrBadRequest, "WEBDOCUMENT_SIZE_TOO_BIG")

	// ErrChatIdEmpty
	// | 400 | CHAT_ID_EMPTY | The provided chat ID is empty. |
	ErrChatIdEmpty = ecode.NewCodeError(ErrBadRequest, "CHAT_ID_EMPTY")

	// ErrFromPeerInvalid
	// | 400 | FROM_PEER_INVALID | The specified from_id is invalid. |
	ErrFromPeerInvalid = ecode.NewCodeError(ErrBadRequest, "FROM_PEER_INVALID")

	// ErrGroupcallInvalid
	// | 400 | GROUPCALL_INVALID | The specified group call is invalid. |
	ErrGroupcallInvalid = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrFilePartSizeInvalid
	// | 400 | FILE_PART_SIZE_INVALID | The provided file part size is invalid. |
	ErrFilePartSizeInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_PART_SIZE_INVALID")

	// ErrUserIsBot
	// | 400 | USER_IS_BOT | Bots can't send messages to other bots. |
	ErrUserIsBot = ecode.NewCodeError(ErrBadRequest, "USER_IS_BOT")

	// ErrMaxDateInvalid
	// | 400 | MAX_DATE_INVALID | The specified maximum date is invalid. |
	ErrMaxDateInvalid = ecode.NewCodeError(ErrBadRequest, "MAX_DATE_INVALID")

	// ErrTopicTitleEmpty
	// | 400 | TOPIC_TITLE_EMPTY | The specified topic title is empty. |
	ErrTopicTitleEmpty = ecode.NewCodeError(ErrBadRequest, "TOPIC_TITLE_EMPTY")

	// ErrUserVolumeInvalid
	// | 400 | USER_VOLUME_INVALID | The specified user volume is invalid. |
	ErrUserVolumeInvalid = ecode.NewCodeError(ErrBadRequest, "USER_VOLUME_INVALID")

	// ErrQuizCorrectAnswersTooMuch
	// | 400 | QUIZ_CORRECT_ANSWERS_TOO_MUCH | You specified too many correct answers in a quiz, quizzes can only have one right answer! |
	ErrQuizCorrectAnswersTooMuch = ecode.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_TOO_MUCH")

	// ErrThemeTitleInvalid
	// | 400 | THEME_TITLE_INVALID | The specified theme title is invalid. |
	ErrThemeTitleInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_TITLE_INVALID")

	// ErrUserCreator
	// | 400 | USER_CREATOR | For channels.editAdmin: you've tried to edit the admin rights of the owner, but you're not the owner; for channels.leaveChannel: you can't leave this channel, because you're its creator. |
	ErrUserCreator = ecode.NewCodeError(ErrBadRequest, "USER_CREATOR")

	// ErrAuthTokenException
	// | 400 | AUTH_TOKEN_EXCEPTION | An error occurred while importing the auth token. |
	ErrAuthTokenException = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_EXCEPTION")

	// ErrAuthTokenInvalid
	// | 400 | AUTH_TOKEN_INVALID | The specified auth token is invalid. |
	ErrAuthTokenInvalid = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_INVALID")

	// ErrChannelsAdminPublicTooMuch
	// | 400 | CHANNELS_ADMIN_PUBLIC_TOO_MUCH | You're admin of too many public channels, make some channels private to change the username of this channel. |
	ErrChannelsAdminPublicTooMuch = ecode.NewCodeError(ErrBadRequest, "CHANNELS_ADMIN_PUBLIC_TOO_MUCH")

	// ErrInputTextTooLong
	// | 400 | INPUT_TEXT_TOO_LONG | The specified text is too long. |
	ErrInputTextTooLong = ecode.NewCodeError(ErrBadRequest, "INPUT_TEXT_TOO_LONG")

	// ErrSessionTooFresh_%d
	// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
	// ErrSessionTooFresh_%d = ecode.NewCodeError(ErrBadRequest, "SESSION_TOO_FRESH_%d")

	// ErrWallpaperInvalid
	// | 400 | WALLPAPER_INVALID | The specified wallpaper is invalid. |
	ErrWallpaperInvalid = ecode.NewCodeError(ErrBadRequest, "WALLPAPER_INVALID")

	// ErrYouBlockedUser
	// | 400 | YOU_BLOCKED_USER | You blocked this user. |
	ErrYouBlockedUser = ecode.NewCodeError(ErrBadRequest, "YOU_BLOCKED_USER")

	// ErrAccessTokenInvalid
	// | 400 | ACCESS_TOKEN_INVALID | Access token invalid. |
	ErrAccessTokenInvalid = ecode.NewCodeError(ErrBadRequest, "ACCESS_TOKEN_INVALID")

	// ErrCallProtocolFlagsInvalid
	// | 400 | CALL_PROTOCOL_FLAGS_INVALID | Call protocol flags invalid. |
	ErrCallProtocolFlagsInvalid = ecode.NewCodeError(ErrBadRequest, "CALL_PROTOCOL_FLAGS_INVALID")

	// ErrChatIdInvalid
	// | 400 | CHAT_ID_INVALID | The provided chat id is invalid. |
	ErrChatIdInvalid = ecode.NewCodeError(ErrBadRequest, "CHAT_ID_INVALID")

	// ErrWebpageNotFound
	// | 400 | WEBPAGE_NOT_FOUND | A preview for the specified webpage `url` could not be generated. |
	ErrWebpageNotFound = ecode.NewCodeError(ErrBadRequest, "WEBPAGE_NOT_FOUND")

	// ErrBotAppBotInvalid
	// | 400 | BOT_APP_BOT_INVALID | The bot_id passed in the inputBotAppShortName constructor is invalid. |
	ErrBotAppBotInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_APP_BOT_INVALID")

	// ErrStickerThumbTgsNotgs
	// | 400 | STICKER_THUMB_TGS_NOTGS | Incorrect stickerset TGS thumb file provided. |
	ErrStickerThumbTgsNotgs = ecode.NewCodeError(ErrBadRequest, "STICKER_THUMB_TGS_NOTGS")

	// ErrUsernamePurchaseAvailable
	// | 400 | USERNAME_PURCHASE_AVAILABLE | The specified username can be purchased on https://fragment.com. |
	ErrUsernamePurchaseAvailable = ecode.NewCodeError(ErrBadRequest, "USERNAME_PURCHASE_AVAILABLE")

	// ErrConnectionDeviceModelEmpty
	// | 400 | CONNECTION_DEVICE_MODEL_EMPTY | The specified device model is empty. |
	ErrConnectionDeviceModelEmpty = ecode.NewCodeError(ErrBadRequest, "CONNECTION_DEVICE_MODEL_EMPTY")

	// ErrBotWebviewDisabled
	// | 400 | BOT_WEBVIEW_DISABLED | A webview cannot be opened in the specified conditions: emitted for example if `from_bot_menu` or `url` are set and `peer` is not the chat with the bot. |
	ErrBotWebviewDisabled = ecode.NewCodeError(ErrBadRequest, "BOT_WEBVIEW_DISABLED")

	// ErrPhotoInvalidDimensions
	// | 400 | PHOTO_INVALID_DIMENSIONS | The photo dimensions are invalid. |
	ErrPhotoInvalidDimensions = ecode.NewCodeError(ErrBadRequest, "PHOTO_INVALID_DIMENSIONS")

	// ErrPollOptionDuplicate
	// | 400 | POLL_OPTION_DUPLICATE | Duplicate poll options provided. |
	ErrPollOptionDuplicate = ecode.NewCodeError(ErrBadRequest, "POLL_OPTION_DUPLICATE")

	// ErrInvoicePayloadInvalid
	// | 400 | INVOICE_PAYLOAD_INVALID | The specified invoice payload is invalid. |
	ErrInvoicePayloadInvalid = ecode.NewCodeError(ErrBadRequest, "INVOICE_PAYLOAD_INVALID")

	// ErrQuizCorrectAnswersEmpty
	// | 400 | QUIZ_CORRECT_ANSWERS_EMPTY | No correct quiz answer was specified. |
	ErrQuizCorrectAnswersEmpty = ecode.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_EMPTY")

	// ErrResultIdEmpty
	// | 400 | RESULT_ID_EMPTY | Result ID empty. |
	ErrResultIdEmpty = ecode.NewCodeError(ErrBadRequest, "RESULT_ID_EMPTY")

	// ErrStoryPeriodInvalid
	// | 400 | STORY_PERIOD_INVALID | The specified story period is invalid for this account. |
	ErrStoryPeriodInvalid = ecode.NewCodeError(ErrBadRequest, "STORY_PERIOD_INVALID")

	// Err400UserNotMutualContact
	// | 400 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	Err400UserNotMutualContact = ecode.NewCodeError(ErrBadRequest, "USER_NOT_MUTUAL_CONTACT")

	// ErrGiftSlugExpired
	// | 400 | GIFT_SLUG_EXPIRED | The specified gift slug has expired. |
	ErrGiftSlugExpired = ecode.NewCodeError(ErrBadRequest, "GIFT_SLUG_EXPIRED")

	// ErrParticipantsTooFew
	// | 400 | PARTICIPANTS_TOO_FEW | Not enough participants. |
	ErrParticipantsTooFew = ecode.NewCodeError(ErrBadRequest, "PARTICIPANTS_TOO_FEW")

	// ErrResultsTooMuch
	// | 400 | RESULTS_TOO_MUCH | Too many results were provided. |
	ErrResultsTooMuch = ecode.NewCodeError(ErrBadRequest, "RESULTS_TOO_MUCH")

	// ErrStickerMimeInvalid
	// | 400 | STICKER_MIME_INVALID | The specified sticker MIME type is invalid. |
	ErrStickerMimeInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_MIME_INVALID")

	// ErrStickerpackStickersTooMuch
	// | 400 | STICKERPACK_STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickerpackStickersTooMuch = ecode.NewCodeError(ErrBadRequest, "STICKERPACK_STICKERS_TOO_MUCH")

	// ErrEmoticonEmpty
	// | 400 | EMOTICON_EMPTY | The emoji is empty. |
	ErrEmoticonEmpty = ecode.NewCodeError(ErrBadRequest, "EMOTICON_EMPTY")

	// ErrEncryptionAlreadyDeclined
	// | 400 | ENCRYPTION_ALREADY_DECLINED | The secret chat was already declined. |
	ErrEncryptionAlreadyDeclined = ecode.NewCodeError(ErrBadRequest, "ENCRYPTION_ALREADY_DECLINED")

	// ErrResultTypeInvalid
	// | 400 | RESULT_TYPE_INVALID | Result type invalid. |
	ErrResultTypeInvalid = ecode.NewCodeError(ErrBadRequest, "RESULT_TYPE_INVALID")

	// Err400ChannelPrivate
	// | 400 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	Err400ChannelPrivate = ecode.NewCodeError(ErrBadRequest, "CHANNEL_PRIVATE")

	// ErrPhotoSaveFileInvalid
	// | 400 | PHOTO_SAVE_FILE_INVALID | Internal issues, try again later. |
	ErrPhotoSaveFileInvalid = ecode.NewCodeError(ErrBadRequest, "PHOTO_SAVE_FILE_INVALID")

	// ErrTaskAlreadyExists
	// | 400 | TASK_ALREADY_EXISTS | An email reset was already requested. |
	ErrTaskAlreadyExists = ecode.NewCodeError(ErrBadRequest, "TASK_ALREADY_EXISTS")

	// ErrCreateCallFailed
	// | 400 | CREATE_CALL_FAILED | An error occurred while creating the call. |
	ErrCreateCallFailed = ecode.NewCodeError(ErrBadRequest, "CREATE_CALL_FAILED")

	// ErrEmojiInvalid
	// | 400 | EMOJI_INVALID | The specified theme emoji is valid. |
	ErrEmojiInvalid = ecode.NewCodeError(ErrBadRequest, "EMOJI_INVALID")

	// ErrMediaVideoStoryMissing
	// | 400 | MEDIA_VIDEO_STORY_MISSING | A non-story video cannot be repubblished as a story (emitted when trying to resend a non-story video as a story using inputDocument). |
	ErrMediaVideoStoryMissing = ecode.NewCodeError(ErrBadRequest, "MEDIA_VIDEO_STORY_MISSING")

	// ErrSha256HashInvalid
	// | 400 | SHA256_HASH_INVALID | The provided SHA256 hash is invalid. |
	ErrSha256HashInvalid = ecode.NewCodeError(ErrBadRequest, "SHA256_HASH_INVALID")

	// ErrEncryptedMessageInvalid
	// | 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message invalid. |
	ErrEncryptedMessageInvalid = ecode.NewCodeError(ErrBadRequest, "ENCRYPTED_MESSAGE_INVALID")

	// ErrGroupcallSsrcDuplicateMuch
	// | 400 | GROUPCALL_SSRC_DUPLICATE_MUCH | The app needs to retry joining the group call with a new SSRC value. |
	ErrGroupcallSsrcDuplicateMuch = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_SSRC_DUPLICATE_MUCH")

	// ErrRaiseHandForbidden
	// | 400 | RAISE_HAND_FORBIDDEN | You cannot raise your hand. |
	ErrRaiseHandForbidden = ecode.NewCodeError(ErrBadRequest, "RAISE_HAND_FORBIDDEN")

	// ErrStickersetNotModified
	// | 400 | STICKERSET_NOT_MODIFIED | The passed stickerset information is equal to the current information. |
	ErrStickersetNotModified = ecode.NewCodeError(ErrBadRequest, "STICKERSET_NOT_MODIFIED")

	// ErrBotNotConnectedYet
	// | 400 | BOT_NOT_CONNECTED_YET | No [business bot](https://core.telegram.org/api/business#connected-bots) is connected to the currently logged in user. |
	ErrBotNotConnectedYet = ecode.NewCodeError(ErrBadRequest, "BOT_NOT_CONNECTED_YET")

	// ErrFileReferenceEmpty
	// | 400 | FILE_REFERENCE_EMPTY | An empty [file reference](https://core.telegram.org/api/file_reference) was specified. |
	ErrFileReferenceEmpty = ecode.NewCodeError(ErrBadRequest, "FILE_REFERENCE_EMPTY")

	// ErrUserAlreadyParticipant
	// | 400 | USER_ALREADY_PARTICIPANT | The user is already in the group. |
	ErrUserAlreadyParticipant = ecode.NewCodeError(ErrBadRequest, "USER_ALREADY_PARTICIPANT")

	// ErrWebpageUrlInvalid
	// | 400 | WEBPAGE_URL_INVALID | The specified webpage `url` is invalid. |
	ErrWebpageUrlInvalid = ecode.NewCodeError(ErrBadRequest, "WEBPAGE_URL_INVALID")

	// ErrChannelInvalid
	// | 400 | CHANNEL_INVALID | The provided channel is invalid. |
	ErrChannelInvalid = ecode.NewCodeError(ErrBadRequest, "CHANNEL_INVALID")

	// ErrSendMessageTypeInvalid
	// | 400 | SEND_MESSAGE_TYPE_INVALID | The message type is invalid. |
	ErrSendMessageTypeInvalid = ecode.NewCodeError(ErrBadRequest, "SEND_MESSAGE_TYPE_INVALID")

	// ErrStickerInvalid
	// | 400 | STICKER_INVALID | The provided sticker is invalid. |
	ErrStickerInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_INVALID")

	// ErrRandomIdInvalid
	// | 400 | RANDOM_ID_INVALID | A provided random ID is invalid. |
	ErrRandomIdInvalid = ecode.NewCodeError(ErrBadRequest, "RANDOM_ID_INVALID")

	// ErrConnectionAppVersionEmpty
	// | 400 | CONNECTION_APP_VERSION_EMPTY | App version is empty. |
	ErrConnectionAppVersionEmpty = ecode.NewCodeError(ErrBadRequest, "CONNECTION_APP_VERSION_EMPTY")

	// ErrFileReferenceExpired
	// | 400 | FILE_REFERENCE_EXPIRED | File reference expired, it must be refetched as described in [the documentation](https://core.telegram.org/api/file_reference). |
	ErrFileReferenceExpired = ecode.NewCodeError(ErrBadRequest, "FILE_REFERENCE_EXPIRED")

	// ErrInviteRequestSent
	// | 400 | INVITE_REQUEST_SENT | You have successfully requested to join this chat or channel. |
	ErrInviteRequestSent = ecode.NewCodeError(ErrBadRequest, "INVITE_REQUEST_SENT")

	// ErrChatlinkSlugEmpty
	// | 400 | CHATLINK_SLUG_EMPTY | The specified slug is empty. |
	ErrChatlinkSlugEmpty = ecode.NewCodeError(ErrBadRequest, "CHATLINK_SLUG_EMPTY")

	// ErrLangPackInvalid
	// | 400 | LANG_PACK_INVALID | The provided language pack is invalid. |
	ErrLangPackInvalid = ecode.NewCodeError(ErrBadRequest, "LANG_PACK_INVALID")

	// ErrResultIdInvalid
	// | 400 | RESULT_ID_INVALID | One of the specified result IDs is invalid. |
	ErrResultIdInvalid = ecode.NewCodeError(ErrBadRequest, "RESULT_ID_INVALID")

	// ErrThemeFormatInvalid
	// | 400 | THEME_FORMAT_INVALID | Invalid theme format provided. |
	ErrThemeFormatInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_FORMAT_INVALID")

	// ErrBotCommandInvalid
	// | 400 | BOT_COMMAND_INVALID | The specified command is invalid. |
	ErrBotCommandInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_COMMAND_INVALID")

	// ErrBotScoreNotModified
	// | 400 | BOT_SCORE_NOT_MODIFIED | The score wasn't modified. |
	ErrBotScoreNotModified = ecode.NewCodeError(ErrBadRequest, "BOT_SCORE_NOT_MODIFIED")

	// ErrBroadcastPublicVotersForbidden
	// | 400 | BROADCAST_PUBLIC_VOTERS_FORBIDDEN | You can't forward polls with public voters. |
	ErrBroadcastPublicVotersForbidden = ecode.NewCodeError(ErrBadRequest, "BROADCAST_PUBLIC_VOTERS_FORBIDDEN")

	// ErrNewSaltInvalid
	// | 400 | NEW_SALT_INVALID | The new salt is invalid. |
	ErrNewSaltInvalid = ecode.NewCodeError(ErrBadRequest, "NEW_SALT_INVALID")

	// ErrStickerEmojiInvalid
	// | 400 | STICKER_EMOJI_INVALID | Sticker emoji invalid. |
	ErrStickerEmojiInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_EMOJI_INVALID")

	// ErrColorInvalid
	// | 400 | COLOR_INVALID | The specified color palette ID was invalid. |
	ErrColorInvalid = ecode.NewCodeError(ErrBadRequest, "COLOR_INVALID")

	// Err400UserBotInvalid
	// | 400 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	Err400UserBotInvalid = ecode.NewCodeError(ErrBadRequest, "USER_BOT_INVALID")

	// Err400GroupcallForbidden
	// | 400 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	Err400GroupcallForbidden = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_FORBIDDEN")

	// ErrCdnMethodInvalid
	// | 400 | CDN_METHOD_INVALID | You can't call this method in a CDN DC. |
	ErrCdnMethodInvalid = ecode.NewCodeError(ErrBadRequest, "CDN_METHOD_INVALID")

	// ErrEmojiMarkupInvalid
	// | 400 | EMOJI_MARKUP_INVALID | The specified `video_emoji_markup` was invalid. |
	ErrEmojiMarkupInvalid = ecode.NewCodeError(ErrBadRequest, "EMOJI_MARKUP_INVALID")

	// ErrGeoPointInvalid
	// | 400 | GEO_POINT_INVALID | Invalid geoposition provided. |
	ErrGeoPointInvalid = ecode.NewCodeError(ErrBadRequest, "GEO_POINT_INVALID")

	// Err400ParticipantJoinMissing
	// | 400 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	Err400ParticipantJoinMissing = ecode.NewCodeError(ErrBadRequest, "PARTICIPANT_JOIN_MISSING")

	// ErrPersistentTimestampEmpty
	// | 400 | PERSISTENT_TIMESTAMP_EMPTY | Persistent timestamp empty. |
	ErrPersistentTimestampEmpty = ecode.NewCodeError(ErrBadRequest, "PERSISTENT_TIMESTAMP_EMPTY")

	// ErrTopicIdInvalid
	// | 400 | TOPIC_ID_INVALID | The specified topic ID is invalid. |
	ErrTopicIdInvalid = ecode.NewCodeError(ErrBadRequest, "TOPIC_ID_INVALID")

	// ErrFolderIdInvalid
	// | 400 | FOLDER_ID_INVALID | Invalid folder ID. |
	ErrFolderIdInvalid = ecode.NewCodeError(ErrBadRequest, "FOLDER_ID_INVALID")

	// ErrMegagroupRequired
	// | 400 | MEGAGROUP_REQUIRED | You can only use this method on a supergroup. |
	ErrMegagroupRequired = ecode.NewCodeError(ErrBadRequest, "MEGAGROUP_REQUIRED")

	// ErrOrderInvalid
	// | 400 | ORDER_INVALID | The specified username order is invalid. |
	ErrOrderInvalid = ecode.NewCodeError(ErrBadRequest, "ORDER_INVALID")

	// ErrRandomIdEmpty
	// | 400 | RANDOM_ID_EMPTY | Random ID empty. |
	ErrRandomIdEmpty = ecode.NewCodeError(ErrBadRequest, "RANDOM_ID_EMPTY")

	// ErrEncryptionDeclined
	// | 400 | ENCRYPTION_DECLINED | The secret chat was declined. |
	ErrEncryptionDeclined = ecode.NewCodeError(ErrBadRequest, "ENCRYPTION_DECLINED")

	// ErrPasswordRecoveryExpired
	// | 400 | PASSWORD_RECOVERY_EXPIRED | The recovery code has expired. |
	ErrPasswordRecoveryExpired = ecode.NewCodeError(ErrBadRequest, "PASSWORD_RECOVERY_EXPIRED")

	// ErrUsernameNotOccupied
	// | 400 | USERNAME_NOT_OCCUPIED | The provided username is not occupied. |
	ErrUsernameNotOccupied = ecode.NewCodeError(ErrBadRequest, "USERNAME_NOT_OCCUPIED")

	// ErrMediaEmpty
	// | 400 | MEDIA_EMPTY | The provided media object is invalid. |
	ErrMediaEmpty = ecode.NewCodeError(ErrBadRequest, "MEDIA_EMPTY")

	// ErrQuizCorrectAnswerInvalid
	// | 400 | QUIZ_CORRECT_ANSWER_INVALID | An invalid value was provided to the correct_answers field. |
	ErrQuizCorrectAnswerInvalid = ecode.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWER_INVALID")

	// ErrUserAdminInvalid
	// | 400 | USER_ADMIN_INVALID | You're not an admin. |
	ErrUserAdminInvalid = ecode.NewCodeError(ErrBadRequest, "USER_ADMIN_INVALID")

	// ErrWebpushAuthInvalid
	// | 400 | WEBPUSH_AUTH_INVALID | The specified web push authentication secret is invalid. |
	ErrWebpushAuthInvalid = ecode.NewCodeError(ErrBadRequest, "WEBPUSH_AUTH_INVALID")

	// ErrBotAlreadyDisabled
	// | 400 | BOT_ALREADY_DISABLED | The connected business bot was already disabled for the specified peer. |
	ErrBotAlreadyDisabled = ecode.NewCodeError(ErrBadRequest, "BOT_ALREADY_DISABLED")

	// ErrChannelIdInvalid
	// | 400 | CHANNEL_ID_INVALID | The specified supergroup ID is invalid. |
	ErrChannelIdInvalid = ecode.NewCodeError(ErrBadRequest, "CHANNEL_ID_INVALID")

	// Err400ChannelTooLarge
	// | 400 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	Err400ChannelTooLarge = ecode.NewCodeError(ErrBadRequest, "CHANNEL_TOO_LARGE")

	// ErrBotAppShortnameInvalid
	// | 400 | BOT_APP_SHORTNAME_INVALID | The specified bot app short name is invalid. |
	ErrBotAppShortnameInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_APP_SHORTNAME_INVALID")

	// ErrPasswordHashInvalid
	// | 400 | PASSWORD_HASH_INVALID | The provided password hash is invalid. |
	ErrPasswordHashInvalid = ecode.NewCodeError(ErrBadRequest, "PASSWORD_HASH_INVALID")

	// ErrStoriesNeverCreated
	// | 400 | STORIES_NEVER_CREATED | This peer hasn't ever posted any stories. |
	ErrStoriesNeverCreated = ecode.NewCodeError(ErrBadRequest, "STORIES_NEVER_CREATED")

	// ErrGifIdInvalid
	// | 400 | GIF_ID_INVALID | The provided GIF ID is invalid. |
	ErrGifIdInvalid = ecode.NewCodeError(ErrBadRequest, "GIF_ID_INVALID")

	// ErrReplyMarkupTooLong
	// | 400 | REPLY_MARKUP_TOO_LONG | The specified reply_markup is too long. |
	ErrReplyMarkupTooLong = ecode.NewCodeError(ErrBadRequest, "REPLY_MARKUP_TOO_LONG")

	// ErrBusinessPeerInvalid
	// | 400 | BUSINESS_PEER_INVALID | Messages can't be set to the specified peer through the current [business connection](https://core.telegram.org/api/business#connected-bots). |
	ErrBusinessPeerInvalid = ecode.NewCodeError(ErrBadRequest, "BUSINESS_PEER_INVALID")

	// ErrResetRequestMissing
	// | 400 | RESET_REQUEST_MISSING | No password reset is in progress. |
	ErrResetRequestMissing = ecode.NewCodeError(ErrBadRequest, "RESET_REQUEST_MISSING")

	// ErrInputConstructorInvalid
	// | 400 | INPUT_CONSTRUCTOR_INVALID | The specified TL constructor is invalid. |
	ErrInputConstructorInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// ErrFirstnameInvalid
	// | 400 | FIRSTNAME_INVALID | The first name is invalid. |
	ErrFirstnameInvalid = ecode.NewCodeError(ErrBadRequest, "FIRSTNAME_INVALID")

	// ErrLastnameInvalid
	// | 400 | LASTNAME_INVALID | The last name is invalid. |
	ErrLastnameInvalid = ecode.NewCodeError(ErrBadRequest, "LASTNAME_INVALID")

	// ErrRingtoneMimeInvalid
	// | 400 | RINGTONE_MIME_INVALID | The MIME type for the ringtone is invalid. |
	ErrRingtoneMimeInvalid = ecode.NewCodeError(ErrBadRequest, "RINGTONE_MIME_INVALID")

	// ErrWebdocumentMimeInvalid
	// | 400 | WEBDOCUMENT_MIME_INVALID | Invalid webdocument mime type provided. |
	ErrWebdocumentMimeInvalid = ecode.NewCodeError(ErrBadRequest, "WEBDOCUMENT_MIME_INVALID")

	// ErrWebpushTokenInvalid
	// | 400 | WEBPUSH_TOKEN_INVALID | The specified web push token is invalid. |
	ErrWebpushTokenInvalid = ecode.NewCodeError(ErrBadRequest, "WEBPUSH_TOKEN_INVALID")

	// ErrInputMethodInvalid
	// | 400 | INPUT_METHOD_INVALID | The specified method is invalid. |
	ErrInputMethodInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_METHOD_INVALID")

	// ErrPhotoCropSizeSmall
	// | 400 | PHOTO_CROP_SIZE_SMALL | Photo is too small. |
	ErrPhotoCropSizeSmall = ecode.NewCodeError(ErrBadRequest, "PHOTO_CROP_SIZE_SMALL")

	// ErrSwitchWebviewUrlInvalid
	// | 400 | SWITCH_WEBVIEW_URL_INVALID | The URL specified in switch_webview.url is invalid! |
	ErrSwitchWebviewUrlInvalid = ecode.NewCodeError(ErrBadRequest, "SWITCH_WEBVIEW_URL_INVALID")

	// ErrThemeMimeInvalid
	// | 400 | THEME_MIME_INVALID | The theme's MIME type is invalid. |
	ErrThemeMimeInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_MIME_INVALID")

	// ErrPrivacyTooLong
	// | 400 | PRIVACY_TOO_LONG | Too many privacy rules were specified, the current limit is 1000. |
	ErrPrivacyTooLong = ecode.NewCodeError(ErrBadRequest, "PRIVACY_TOO_LONG")

	// ErrSrpPasswordChanged
	// | 400 | SRP_PASSWORD_CHANGED | Password has changed. |
	ErrSrpPasswordChanged = ecode.NewCodeError(ErrBadRequest, "SRP_PASSWORD_CHANGED")

	// ErrUsersTooFew
	// | 400 | USERS_TOO_FEW | Not enough users (to create a chat, for example). |
	ErrUsersTooFew = ecode.NewCodeError(ErrBadRequest, "USERS_TOO_FEW")

	// ErrBotGroupsBlocked
	// | 400 | BOT_GROUPS_BLOCKED | This bot can't be added to groups. |
	ErrBotGroupsBlocked = ecode.NewCodeError(ErrBadRequest, "BOT_GROUPS_BLOCKED")

	// ErrEntityMentionUserInvalid
	// | 400 | ENTITY_MENTION_USER_INVALID | You mentioned an invalid user. |
	ErrEntityMentionUserInvalid = ecode.NewCodeError(ErrBadRequest, "ENTITY_MENTION_USER_INVALID")

	// ErrNewSettingsInvalid
	// | 400 | NEW_SETTINGS_INVALID | The new password settings are invalid. |
	ErrNewSettingsInvalid = ecode.NewCodeError(ErrBadRequest, "NEW_SETTINGS_INVALID")

	// ErrEmoticonStickerpackMissing
	// | 400 | EMOTICON_STICKERPACK_MISSING | inputStickerSetDice.emoji cannot be empty. |
	ErrEmoticonStickerpackMissing = ecode.NewCodeError(ErrBadRequest, "EMOTICON_STICKERPACK_MISSING")

	// ErrContactMissing
	// | 400 | CONTACT_MISSING | The specified user is not a contact. |
	ErrContactMissing = ecode.NewCodeError(ErrBadRequest, "CONTACT_MISSING")

	// ErrReplyMarkupGameEmpty
	// | 400 | REPLY_MARKUP_GAME_EMPTY | A game message is being edited, but the newly provided keyboard doesn't have a keyboardButtonGame button. |
	ErrReplyMarkupGameEmpty = ecode.NewCodeError(ErrBadRequest, "REPLY_MARKUP_GAME_EMPTY")

	// ErrButtonPosInvalid
	// | 400 | BUTTON_POS_INVALID | The position of one of the keyboard buttons is invalid (i.e. a Game or Pay button not in the first position, and so on...). |
	ErrButtonPosInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_POS_INVALID")

	// ErrUntilDateInvalid
	// | 400 | UNTIL_DATE_INVALID | Invalid until date provided. |
	ErrUntilDateInvalid = ecode.NewCodeError(ErrBadRequest, "UNTIL_DATE_INVALID")

	// ErrThemeInvalid
	// | 400 | THEME_INVALID | Invalid theme provided. |
	ErrThemeInvalid = ecode.NewCodeError(ErrBadRequest, "THEME_INVALID")

	// ErrUsernameNotModified
	// | 400 | USERNAME_NOT_MODIFIED | The username was not modified. |
	ErrUsernameNotModified = ecode.NewCodeError(ErrBadRequest, "USERNAME_NOT_MODIFIED")

	// ErrFilePartTooBig
	// | 400 | FILE_PART_TOO_BIG | The uploaded file part is too big. |
	ErrFilePartTooBig = ecode.NewCodeError(ErrBadRequest, "FILE_PART_TOO_BIG")

	// ErrMessageEditTimeExpired
	// | 400 | MESSAGE_EDIT_TIME_EXPIRED | You can't edit this message anymore, too much time has passed since its creation. |
	ErrMessageEditTimeExpired = ecode.NewCodeError(ErrBadRequest, "MESSAGE_EDIT_TIME_EXPIRED")

	// ErrReplyMessagesTooMuch
	// | 400 | REPLY_MESSAGES_TOO_MUCH | Each shortcut can contain a maximum of [appConfig.`quick_reply_messages_limit`](https://core.telegram.org/api/config#quick-reply-messages-limit) messages, the limit was reached. |
	ErrReplyMessagesTooMuch = ecode.NewCodeError(ErrBadRequest, "REPLY_MESSAGES_TOO_MUCH")

	// ErrButtonUserInvalid
	// | 400 | BUTTON_USER_INVALID | The `user_id` passed to inputKeyboardButtonUserProfile is invalid! |
	ErrButtonUserInvalid = ecode.NewCodeError(ErrBadRequest, "BUTTON_USER_INVALID")

	// ErrScheduleDateInvalid
	// | 400 | SCHEDULE_DATE_INVALID | Invalid schedule date provided. |
	ErrScheduleDateInvalid = ecode.NewCodeError(ErrBadRequest, "SCHEDULE_DATE_INVALID")

	// ErrWallpaperNotFound
	// | 400 | WALLPAPER_NOT_FOUND | The specified wallpaper could not be found. |
	ErrWallpaperNotFound = ecode.NewCodeError(ErrBadRequest, "WALLPAPER_NOT_FOUND")

	// ErrLinkNotModified
	// | 400 | LINK_NOT_MODIFIED | Discussion link not modified. |
	ErrLinkNotModified = ecode.NewCodeError(ErrBadRequest, "LINK_NOT_MODIFIED")

	// ErrRandomLengthInvalid
	// | 400 | RANDOM_LENGTH_INVALID | Random length invalid. |
	ErrRandomLengthInvalid = ecode.NewCodeError(ErrBadRequest, "RANDOM_LENGTH_INVALID")

	// ErrSecureSecretRequired
	// | 400 | SECURE_SECRET_REQUIRED | A secure secret is required. |
	ErrSecureSecretRequired = ecode.NewCodeError(ErrBadRequest, "SECURE_SECRET_REQUIRED")

	// ErrTopicCloseSeparately
	// | 400 | TOPIC_CLOSE_SEPARATELY | The `close` flag cannot be provided together with any of the other flags. |
	ErrTopicCloseSeparately = ecode.NewCodeError(ErrBadRequest, "TOPIC_CLOSE_SEPARATELY")

	// ErrInputLayerInvalid
	// | 400 | INPUT_LAYER_INVALID | The specified layer is invalid. |
	ErrInputLayerInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_LAYER_INVALID")

	// ErrCollectibleInvalid
	// | 400 | COLLECTIBLE_INVALID | The specified collectible is invalid. |
	ErrCollectibleInvalid = ecode.NewCodeError(ErrBadRequest, "COLLECTIBLE_INVALID")

	// ErrExportCardInvalid
	// | 400 | EXPORT_CARD_INVALID | Provided card is invalid. |
	ErrExportCardInvalid = ecode.NewCodeError(ErrBadRequest, "EXPORT_CARD_INVALID")

	// ErrGroupedMediaInvalid
	// | 400 | GROUPED_MEDIA_INVALID | Invalid grouped media. |
	ErrGroupedMediaInvalid = ecode.NewCodeError(ErrBadRequest, "GROUPED_MEDIA_INVALID")

	// ErrUsernamesActiveTooMuch
	// | 400 | USERNAMES_ACTIVE_TOO_MUCH | The maximum number of active usernames was reached. |
	ErrUsernamesActiveTooMuch = ecode.NewCodeError(ErrBadRequest, "USERNAMES_ACTIVE_TOO_MUCH")

	// ErrAdminRankEmojiNotAllowed
	// | 400 | ADMIN_RANK_EMOJI_NOT_ALLOWED | An admin rank cannot contain emojis. |
	ErrAdminRankEmojiNotAllowed = ecode.NewCodeError(ErrBadRequest, "ADMIN_RANK_EMOJI_NOT_ALLOWED")

	// ErrPackTitleInvalid
	// | 400 | PACK_TITLE_INVALID | The stickerpack title is invalid. |
	ErrPackTitleInvalid = ecode.NewCodeError(ErrBadRequest, "PACK_TITLE_INVALID")

	// ErrRingtoneInvalid
	// | 400 | RINGTONE_INVALID | The specified ringtone is invalid. |
	ErrRingtoneInvalid = ecode.NewCodeError(ErrBadRequest, "RINGTONE_INVALID")

	// ErrApiIdInvalid
	// | 400 | API_ID_INVALID | API ID invalid. |
	ErrApiIdInvalid = ecode.NewCodeError(ErrBadRequest, "API_ID_INVALID")

	// ErrSearchQueryEmpty
	// | 400 | SEARCH_QUERY_EMPTY | The search query is empty. |
	ErrSearchQueryEmpty = ecode.NewCodeError(ErrBadRequest, "SEARCH_QUERY_EMPTY")

	// ErrVenueIdInvalid
	// | 400 | VENUE_ID_INVALID | The specified venue ID is invalid. |
	ErrVenueIdInvalid = ecode.NewCodeError(ErrBadRequest, "VENUE_ID_INVALID")

	// ErrQueryTooShort
	// | 400 | QUERY_TOO_SHORT | The query string is too short. |
	ErrQueryTooShort = ecode.NewCodeError(ErrBadRequest, "QUERY_TOO_SHORT")

	// ErrConnectionSystemLangCodeEmpty
	// | 400 | CONNECTION_SYSTEM_LANG_CODE_EMPTY | The specified system language code is empty. |
	ErrConnectionSystemLangCodeEmpty = ecode.NewCodeError(ErrBadRequest, "CONNECTION_SYSTEM_LANG_CODE_EMPTY")

	// ErrBotGamesDisabled
	// | 400 | BOT_GAMES_DISABLED | Games can't be sent to channels. |
	ErrBotGamesDisabled = ecode.NewCodeError(ErrBadRequest, "BOT_GAMES_DISABLED")

	// ErrContactNameEmpty
	// | 400 | CONTACT_NAME_EMPTY | Contact name empty. |
	ErrContactNameEmpty = ecode.NewCodeError(ErrBadRequest, "CONTACT_NAME_EMPTY")

	// ErrImageProcessFailed
	// | 400 | IMAGE_PROCESS_FAILED | Failure while processing image. |
	ErrImageProcessFailed = ecode.NewCodeError(ErrBadRequest, "IMAGE_PROCESS_FAILED")

	// ErrStickerTgsNotgs
	// | 400 | STICKER_TGS_NOTGS | Invalid TGS sticker provided. |
	ErrStickerTgsNotgs = ecode.NewCodeError(ErrBadRequest, "STICKER_TGS_NOTGS")

	// Err400UserNotParticipant
	// | 400 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	Err400UserNotParticipant = ecode.NewCodeError(ErrBadRequest, "USER_NOT_PARTICIPANT")

	// ErrTmpPasswordInvalid
	// | 400 | TMP_PASSWORD_INVALID | The passed tmp_password is invalid. |
	ErrTmpPasswordInvalid = ecode.NewCodeError(ErrBadRequest, "TMP_PASSWORD_INVALID")

	// ErrBoostsEmpty
	// | 400 | BOOSTS_EMPTY | No boost slots were specified. |
	ErrBoostsEmpty = ecode.NewCodeError(ErrBadRequest, "BOOSTS_EMPTY")

	// ErrBotChannelsNa
	// | 400 | BOT_CHANNELS_NA | Bots can't edit admin privileges. |
	ErrBotChannelsNa = ecode.NewCodeError(ErrBadRequest, "BOT_CHANNELS_NA")

	// ErrFileReferenceInvalid
	// | 400 | FILE_REFERENCE_INVALID | The specified [file reference](https://core.telegram.org/api/file_reference) is invalid. |
	ErrFileReferenceInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_REFERENCE_INVALID")

	// ErrInviteSlugExpired
	// | 400 | INVITE_SLUG_EXPIRED | The specified chat folder link has expired. |
	ErrInviteSlugExpired = ecode.NewCodeError(ErrBadRequest, "INVITE_SLUG_EXPIRED")

	// Err400MsgWaitFailed
	// | 400 | MSG_WAIT_FAILED | A waiting call returned an error. |
	Err400MsgWaitFailed = ecode.NewCodeError(ErrBadRequest, "MSG_WAIT_FAILED")

	// ErrWebdocumentInvalid
	// | 400 | WEBDOCUMENT_INVALID | Invalid webdocument URL provided. |
	ErrWebdocumentInvalid = ecode.NewCodeError(ErrBadRequest, "WEBDOCUMENT_INVALID")

	// ErrTranslateReqQuotaExceeded
	// | 400 | TRANSLATE_REQ_QUOTA_EXCEEDED | Translation is currently unavailable due to a temporary server-side lack of resources. |
	ErrTranslateReqQuotaExceeded = ecode.NewCodeError(ErrBadRequest, "TRANSLATE_REQ_QUOTA_EXCEEDED")

	// ErrChatPublicRequired
	// | 400 | CHAT_PUBLIC_REQUIRED | You can only enable join requests in public groups. |
	ErrChatPublicRequired = ecode.NewCodeError(ErrBadRequest, "CHAT_PUBLIC_REQUIRED")

	// ErrPhonePasswordProtected
	// | 400 | PHONE_PASSWORD_PROTECTED | This phone is password protected. |
	ErrPhonePasswordProtected = ecode.NewCodeError(ErrBadRequest, "PHONE_PASSWORD_PROTECTED")

	// ErrReactionInvalid
	// | 400 | REACTION_INVALID | The specified reaction is invalid. |
	ErrReactionInvalid = ecode.NewCodeError(ErrBadRequest, "REACTION_INVALID")

	// ErrTopicHideSeparately
	// | 400 | TOPIC_HIDE_SEPARATELY | The `hide` flag cannot be provided together with any of the other flags. |
	ErrTopicHideSeparately = ecode.NewCodeError(ErrBadRequest, "TOPIC_HIDE_SEPARATELY")

	// ErrAuthTokenInvalidx
	// | 400 | AUTH_TOKEN_INVALIDX | The specified auth token is invalid. |
	ErrAuthTokenInvalidx = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_INVALIDX")

	// ErrOffsetInvalid
	// | 400 | OFFSET_INVALID | The provided offset is invalid. |
	ErrOffsetInvalid = ecode.NewCodeError(ErrBadRequest, "OFFSET_INVALID")

	// ErrReplyMarkupInvalid
	// | 400 | REPLY_MARKUP_INVALID | The provided reply markup is invalid. |
	ErrReplyMarkupInvalid = ecode.NewCodeError(ErrBadRequest, "REPLY_MARKUP_INVALID")

	// ErrRightsNotModified
	// | 400 | RIGHTS_NOT_MODIFIED | The new admin rights are equal to the old rights, no change was made. |
	ErrRightsNotModified = ecode.NewCodeError(ErrBadRequest, "RIGHTS_NOT_MODIFIED")

	// ErrConnectionNotInited
	// | 400 | CONNECTION_NOT_INITED | Please initialize the connection using initConnection before making queries. |
	ErrConnectionNotInited = ecode.NewCodeError(ErrBadRequest, "CONNECTION_NOT_INITED")

	// ErrBankCardNumberInvalid
	// | 400 | BANK_CARD_NUMBER_INVALID | The specified card number is invalid. |
	ErrBankCardNumberInvalid = ecode.NewCodeError(ErrBadRequest, "BANK_CARD_NUMBER_INVALID")

	// Err400CallOccupyFailed
	// | 400 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	Err400CallOccupyFailed = ecode.NewCodeError(ErrBadRequest, "CALL_OCCUPY_FAILED")

	// ErrContactReqMissing
	// | 400 | CONTACT_REQ_MISSING | Missing contact request. |
	ErrContactReqMissing = ecode.NewCodeError(ErrBadRequest, "CONTACT_REQ_MISSING")

	// ErrShortcutInvalid
	// | 400 | SHORTCUT_INVALID | The specified shortcut is invalid. |
	ErrShortcutInvalid = ecode.NewCodeError(ErrBadRequest, "SHORTCUT_INVALID")

	// ErrEmailVerifyExpired
	// | 400 | EMAIL_VERIFY_EXPIRED | The verification email has expired. |
	ErrEmailVerifyExpired = ecode.NewCodeError(ErrBadRequest, "EMAIL_VERIFY_EXPIRED")

	// ErrNotJoined
	// | 400 | NOT_JOINED | The current user hasn't joined the Peer-to-Peer Login Program. |
	ErrNotJoined = ecode.NewCodeError(ErrBadRequest, "NOT_JOINED")

	// ErrRequestTokenInvalid
	// | 400 | REQUEST_TOKEN_INVALID | The master DC did not accept the `request_token` from the CDN DC. Continue downloading the file from the master DC using upload.getFile. |
	ErrRequestTokenInvalid = ecode.NewCodeError(ErrBadRequest, "REQUEST_TOKEN_INVALID")

	// ErrPasswordRequired
	// | 400 | PASSWORD_REQUIRED | A [2FA password](https://core.telegram.org/api/srp) must be configured to use Telegram Passport. |
	ErrPasswordRequired = ecode.NewCodeError(ErrBadRequest, "PASSWORD_REQUIRED")

	// ErrTopicsEmpty
	// | 400 | TOPICS_EMPTY | You specified no topic IDs. |
	ErrTopicsEmpty = ecode.NewCodeError(ErrBadRequest, "TOPICS_EMPTY")

	// ErrUserBlocked
	// | 400 | USER_BLOCKED | User blocked. |
	ErrUserBlocked = ecode.NewCodeError(ErrBadRequest, "USER_BLOCKED")

	// ErrBotCommandDescriptionInvalid
	// | 400 | BOT_COMMAND_DESCRIPTION_INVALID | The specified command description is invalid. |
	ErrBotCommandDescriptionInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_COMMAND_DESCRIPTION_INVALID")

	// Err400NotEligible
	// | 400 | NOT_ELIGIBLE | The current user is not eligible to join the Peer-to-Peer Login Program. |
	Err400NotEligible = ecode.NewCodeError(ErrBadRequest, "NOT_ELIGIBLE")

	// ErrQueryIdInvalid
	// | 400 | QUERY_ID_INVALID | The query ID is invalid. |
	ErrQueryIdInvalid = ecode.NewCodeError(ErrBadRequest, "QUERY_ID_INVALID")

	// ErrTmpPasswordDisabled
	// | 400 | TMP_PASSWORD_DISABLED | The temporary password is disabled. |
	ErrTmpPasswordDisabled = ecode.NewCodeError(ErrBadRequest, "TMP_PASSWORD_DISABLED")

	// ErrGiftSlugInvalid
	// | 400 | GIFT_SLUG_INVALID | The specified slug is invalid. |
	ErrGiftSlugInvalid = ecode.NewCodeError(ErrBadRequest, "GIFT_SLUG_INVALID")

	// ErrMessageTooLong
	// | 400 | MESSAGE_TOO_LONG | The provided message is too long. |
	ErrMessageTooLong = ecode.NewCodeError(ErrBadRequest, "MESSAGE_TOO_LONG")

	// ErrStickerVideoNodoc
	// | 400 | STICKER_VIDEO_NODOC | You must send the video sticker as a document. |
	ErrStickerVideoNodoc = ecode.NewCodeError(ErrBadRequest, "STICKER_VIDEO_NODOC")

	// ErrEmailNotAllowed
	// | 400 | EMAIL_NOT_ALLOWED | The specified email cannot be used to complete the operation. |
	ErrEmailNotAllowed = ecode.NewCodeError(ErrBadRequest, "EMAIL_NOT_ALLOWED")

	// ErrMediaNewInvalid
	// | 400 | MEDIA_NEW_INVALID | The new media is invalid. |
	ErrMediaNewInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_NEW_INVALID")

	// ErrPasswordEmpty
	// | 400 | PASSWORD_EMPTY | The provided password is empty. |
	ErrPasswordEmpty = ecode.NewCodeError(ErrBadRequest, "PASSWORD_EMPTY")

	// ErrStickerThumbPngNopng
	// | 400 | STICKER_THUMB_PNG_NOPNG | Incorrect stickerset thumb file provided, PNG / WEBP expected. |
	ErrStickerThumbPngNopng = ecode.NewCodeError(ErrBadRequest, "STICKER_THUMB_PNG_NOPNG")

	// ErrAdExpired
	// | 400 | AD_EXPIRED | The ad has expired (too old or not found). |
	ErrAdExpired = ecode.NewCodeError(ErrBadRequest, "AD_EXPIRED")

	// ErrBoostsRequired
	// | 400 | BOOSTS_REQUIRED | The specified channel must first be [boosted by its users](https://core.telegram.org/api/boost) in order to perform this action. |
	ErrBoostsRequired = ecode.NewCodeError(ErrBadRequest, "BOOSTS_REQUIRED")

	// ErrBotAppInvalid
	// | 400 | BOT_APP_INVALID | The specified bot app is invalid. |
	ErrBotAppInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_APP_INVALID")

	// ErrPackShortNameOccupied
	// | 400 | PACK_SHORT_NAME_OCCUPIED | A stickerpack with this name already exists. |
	ErrPackShortNameOccupied = ecode.NewCodeError(ErrBadRequest, "PACK_SHORT_NAME_OCCUPIED")

	// ErrStickerDocumentInvalid
	// | 400 | STICKER_DOCUMENT_INVALID | The specified sticker document is invalid. |
	ErrStickerDocumentInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_DOCUMENT_INVALID")

	// Err400TakeoutRequired
	// | 400 | TAKEOUT_REQUIRED | A [takeout](https://core.telegram.org/api/takeout) session needs to be initialized first, [see here &raquo; for more info](https://core.telegram.org/api/takeout). |
	Err400TakeoutRequired = ecode.NewCodeError(ErrBadRequest, "TAKEOUT_REQUIRED")

	// ErrFilePart_%dMissing
	// | 400 | FILE_PART_%d_MISSING | Part %d of the file is missing from storage. Try repeating the method call to resave the part. |
	// ErrFilePart_%dMissing = ecode.NewCodeError(ErrBadRequest, "FILE_PART_%d_MISSING")

	// ErrLangCodeNotSupported
	// | 400 | LANG_CODE_NOT_SUPPORTED | The specified language code is not supported. |
	ErrLangCodeNotSupported = ecode.NewCodeError(ErrBadRequest, "LANG_CODE_NOT_SUPPORTED")

	// ErrLocationInvalid
	// | 400 | LOCATION_INVALID | The provided location is invalid. |
	ErrLocationInvalid = ecode.NewCodeError(ErrBadRequest, "LOCATION_INVALID")

	// ErrMsgTooOld
	// | 400 | MSG_TOO_OLD | [`chat_read_mark_expire_period` seconds](https://core.telegram.org/api/config#chat-read-mark-expire-period) have passed since the message was sent, read receipts were deleted. |
	ErrMsgTooOld = ecode.NewCodeError(ErrBadRequest, "MSG_TOO_OLD")

	// ErrPhoneNumberAppSignupForbidden
	// | 400 | PHONE_NUMBER_APP_SIGNUP_FORBIDDEN | You can't sign up using this app. |
	ErrPhoneNumberAppSignupForbidden = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_APP_SIGNUP_FORBIDDEN")

	// ErrSrpIdInvalid
	// | 400 | SRP_ID_INVALID | Invalid SRP ID provided. |
	ErrSrpIdInvalid = ecode.NewCodeError(ErrBadRequest, "SRP_ID_INVALID")

	// ErrAdminsTooMuch
	// | 400 | ADMINS_TOO_MUCH | There are too many admins. |
	ErrAdminsTooMuch = ecode.NewCodeError(ErrBadRequest, "ADMINS_TOO_MUCH")

	// ErrChatlinkSlugExpired
	// | 400 | CHATLINK_SLUG_EXPIRED | The specified [business chat link](https://core.telegram.org/api/business#business-chat-links) has expired. |
	ErrChatlinkSlugExpired = ecode.NewCodeError(ErrBadRequest, "CHATLINK_SLUG_EXPIRED")

	// Err400InviteHashExpired
	// | 400 | INVITE_HASH_EXPIRED | The invite link has expired. |
	Err400InviteHashExpired = ecode.NewCodeError(ErrBadRequest, "INVITE_HASH_EXPIRED")

	// Err400VoiceMessagesForbidden
	// | 400 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	Err400VoiceMessagesForbidden = ecode.NewCodeError(ErrBadRequest, "VOICE_MESSAGES_FORBIDDEN")

	// ErrFileMigrate_%d
	// | 400 | FILE_MIGRATE_%d | The file currently being accessed is stored in DC %d, please re-send the query to that DC. |
	// ErrFileMigrate_%d = ecode.NewCodeError(ErrBadRequest, "FILE_MIGRATE_%d")

	// ErrBroadcastIdInvalid
	// | 400 | BROADCAST_ID_INVALID | Broadcast ID invalid. |
	ErrBroadcastIdInvalid = ecode.NewCodeError(ErrBadRequest, "BROADCAST_ID_INVALID")

	// ErrStickersTooMuch
	// | 400 | STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickersTooMuch = ecode.NewCodeError(ErrBadRequest, "STICKERS_TOO_MUCH")

	// ErrReactionEmpty
	// | 400 | REACTION_EMPTY | Empty reaction provided. |
	ErrReactionEmpty = ecode.NewCodeError(ErrBadRequest, "REACTION_EMPTY")

	// ErrSlowmodeMultiMsgsDisabled
	// | 400 | SLOWMODE_MULTI_MSGS_DISABLED | Slowmode is enabled, you cannot forward multiple messages to this group. |
	ErrSlowmodeMultiMsgsDisabled = ecode.NewCodeError(ErrBadRequest, "SLOWMODE_MULTI_MSGS_DISABLED")

	// ErrVideoStopForbidden
	// | 400 | VIDEO_STOP_FORBIDDEN | You cannot stop the video stream. |
	ErrVideoStopForbidden = ecode.NewCodeError(ErrBadRequest, "VIDEO_STOP_FORBIDDEN")

	// ErrAuthTokenAlreadyAccepted
	// | 400 | AUTH_TOKEN_ALREADY_ACCEPTED | The specified auth token was already accepted. |
	ErrAuthTokenAlreadyAccepted = ecode.NewCodeError(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrChatRestricted
	// | 400 | CHAT_RESTRICTED | You can't send messages in this chat, you were restricted. |
	ErrChatRestricted = ecode.NewCodeError(ErrBadRequest, "CHAT_RESTRICTED")

	// ErrPhotoCropFileMissing
	// | 400 | PHOTO_CROP_FILE_MISSING | Photo crop file missing. |
	ErrPhotoCropFileMissing = ecode.NewCodeError(ErrBadRequest, "PHOTO_CROP_FILE_MISSING")

	// ErrApiIdPublishedFlood
	// | 400 | API_ID_PUBLISHED_FLOOD | This API id was published somewhere, you can't use it now. |
	ErrApiIdPublishedFlood = ecode.NewCodeError(ErrBadRequest, "API_ID_PUBLISHED_FLOOD")

	// ErrTokenTypeInvalid
	// | 400 | TOKEN_TYPE_INVALID | The specified token type is invalid. |
	ErrTokenTypeInvalid = ecode.NewCodeError(ErrBadRequest, "TOKEN_TYPE_INVALID")

	// Err400TopicClosed
	// | 400 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	Err400TopicClosed = ecode.NewCodeError(ErrBadRequest, "TOPIC_CLOSED")

	// ErrTtlMediaInvalid
	// | 400 | TTL_MEDIA_INVALID | Invalid media Time To Live was provided. |
	ErrTtlMediaInvalid = ecode.NewCodeError(ErrBadRequest, "TTL_MEDIA_INVALID")

	// ErrBirthdayInvalid
	// | 400 | BIRTHDAY_INVALID | An invalid age was specified, must be between 0 and 150 years. |
	ErrBirthdayInvalid = ecode.NewCodeError(ErrBadRequest, "BIRTHDAY_INVALID")

	// ErrReceiptEmpty
	// | 400 | RECEIPT_EMPTY | The specified receipt is empty. |
	ErrReceiptEmpty = ecode.NewCodeError(ErrBadRequest, "RECEIPT_EMPTY")

	// ErrSecondsInvalid
	// | 400 | SECONDS_INVALID | Invalid duration provided. |
	ErrSecondsInvalid = ecode.NewCodeError(ErrBadRequest, "SECONDS_INVALID")

	// ErrStickerVideoNowebm
	// | 400 | STICKER_VIDEO_NOWEBM | The specified video sticker is not in webm format. |
	ErrStickerVideoNowebm = ecode.NewCodeError(ErrBadRequest, "STICKER_VIDEO_NOWEBM")

	// ErrConnectionIdInvalid
	// | 400 | CONNECTION_ID_INVALID | The specified connection ID is invalid. |
	ErrConnectionIdInvalid = ecode.NewCodeError(ErrBadRequest, "CONNECTION_ID_INVALID")

	// ErrLangCodeInvalid
	// | 400 | LANG_CODE_INVALID | The specified language code is invalid. |
	ErrLangCodeInvalid = ecode.NewCodeError(ErrBadRequest, "LANG_CODE_INVALID")

	// ErrMediaGroupedInvalid
	// | 400 | MEDIA_GROUPED_INVALID | You tried to send media of different types in an album. |
	ErrMediaGroupedInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_GROUPED_INVALID")

	// ErrMegagroupIdInvalid
	// | 400 | MEGAGROUP_ID_INVALID | Invalid supergroup ID. |
	ErrMegagroupIdInvalid = ecode.NewCodeError(ErrBadRequest, "MEGAGROUP_ID_INVALID")

	// ErrAccessTokenExpired
	// | 400 | ACCESS_TOKEN_EXPIRED | Access token expired. |
	ErrAccessTokenExpired = ecode.NewCodeError(ErrBadRequest, "ACCESS_TOKEN_EXPIRED")

	// ErrAlbumPhotosTooMany
	// | 400 | ALBUM_PHOTOS_TOO_MANY | You have uploaded too many profile photos, delete some before retrying. |
	ErrAlbumPhotosTooMany = ecode.NewCodeError(ErrBadRequest, "ALBUM_PHOTOS_TOO_MANY")

	// ErrCollectibleNotFound
	// | 400 | COLLECTIBLE_NOT_FOUND | The specified collectible could not be found. |
	ErrCollectibleNotFound = ecode.NewCodeError(ErrBadRequest, "COLLECTIBLE_NOT_FOUND")

	// ErrStickerPngNopng
	// | 400 | STICKER_PNG_NOPNG | One of the specified stickers is not a valid PNG file. |
	ErrStickerPngNopng = ecode.NewCodeError(ErrBadRequest, "STICKER_PNG_NOPNG")

	// ErrAdminRightsEmpty
	// | 400 | ADMIN_RIGHTS_EMPTY | The chatAdminRights constructor passed in keyboardButtonRequestPeer.peer_type.user_admin_rights has no rights set (i.e. flags is 0). |
	ErrAdminRightsEmpty = ecode.NewCodeError(ErrBadRequest, "ADMIN_RIGHTS_EMPTY")

	// ErrNextOffsetInvalid
	// | 400 | NEXT_OFFSET_INVALID | The specified offset is longer than 64 bytes. |
	ErrNextOffsetInvalid = ecode.NewCodeError(ErrBadRequest, "NEXT_OFFSET_INVALID")

	// ErrStickerIdInvalid
	// | 400 | STICKER_ID_INVALID | The provided sticker ID is invalid. |
	ErrStickerIdInvalid = ecode.NewCodeError(ErrBadRequest, "STICKER_ID_INVALID")

	// ErrBotsTooMuch
	// | 400 | BOTS_TOO_MUCH | There are too many bots in this chat/channel. |
	ErrBotsTooMuch = ecode.NewCodeError(ErrBadRequest, "BOTS_TOO_MUCH")

	// ErrPhoneCodeEmpty
	// | 400 | PHONE_CODE_EMPTY | phone_code is missing. |
	ErrPhoneCodeEmpty = ecode.NewCodeError(ErrBadRequest, "PHONE_CODE_EMPTY")

	// ErrUsernameOccupied
	// | 400 | USERNAME_OCCUPIED | The provided username is already occupied. |
	ErrUsernameOccupied = ecode.NewCodeError(ErrBadRequest, "USERNAME_OCCUPIED")

	// Err400PremiumAccountRequired
	// | 400 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	Err400PremiumAccountRequired = ecode.NewCodeError(ErrBadRequest, "PREMIUM_ACCOUNT_REQUIRED")

	// ErrUserKicked
	// | 400 | USER_KICKED | This user was kicked from this supergroup/channel. |
	ErrUserKicked = ecode.NewCodeError(ErrBadRequest, "USER_KICKED")

	// ErrUserPublicMissing
	// | 400 | USER_PUBLIC_MISSING | Cannot generate a link to stories posted by a peer without a username. |
	ErrUserPublicMissing = ecode.NewCodeError(ErrBadRequest, "USER_PUBLIC_MISSING")

	// ErrWallpaperMimeInvalid
	// | 400 | WALLPAPER_MIME_INVALID | The specified wallpaper MIME type is invalid. |
	ErrWallpaperMimeInvalid = ecode.NewCodeError(ErrBadRequest, "WALLPAPER_MIME_INVALID")

	// Err400ChatInvalid
	// | 400 | CHAT_INVALID | Invalid chat. |
	Err400ChatInvalid = ecode.NewCodeError(ErrBadRequest, "CHAT_INVALID")

	// ErrInputFilterInvalid
	// | 400 | INPUT_FILTER_INVALID | The specified filter is invalid. |
	ErrInputFilterInvalid = ecode.NewCodeError(ErrBadRequest, "INPUT_FILTER_INVALID")

	// ErrPhotoFileMissing
	// | 400 | PHOTO_FILE_MISSING | Profile photo file missing. |
	ErrPhotoFileMissing = ecode.NewCodeError(ErrBadRequest, "PHOTO_FILE_MISSING")

	// ErrOffsetPeerIdInvalid
	// | 400 | OFFSET_PEER_ID_INVALID | The provided offset peer is invalid. |
	ErrOffsetPeerIdInvalid = ecode.NewCodeError(ErrBadRequest, "OFFSET_PEER_ID_INVALID")

	// ErrWebdocumentUrlEmpty
	// | 400 | WEBDOCUMENT_URL_EMPTY | The passed web document URL is empty. |
	ErrWebdocumentUrlEmpty = ecode.NewCodeError(ErrBadRequest, "WEBDOCUMENT_URL_EMPTY")

	// ErrChatlinksTooMuch
	// | 400 | CHATLINKS_TOO_MUCH | Too many [business chat links](https://core.telegram.org/api/business#business-chat-links) were created, please delete some older links. |
	ErrChatlinksTooMuch = ecode.NewCodeError(ErrBadRequest, "CHATLINKS_TOO_MUCH")

	// ErrSlugInvalid
	// | 400 | SLUG_INVALID | The specified invoice slug is invalid. |
	ErrSlugInvalid = ecode.NewCodeError(ErrBadRequest, "SLUG_INVALID")

	// ErrConnectionSystemEmpty
	// | 400 | CONNECTION_SYSTEM_EMPTY | The specified system version is empty. |
	ErrConnectionSystemEmpty = ecode.NewCodeError(ErrBadRequest, "CONNECTION_SYSTEM_EMPTY")

	// ErrPhoneNumberOccupied
	// | 400 | PHONE_NUMBER_OCCUPIED | The phone number is already in use. |
	ErrPhoneNumberOccupied = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_OCCUPIED")

	// ErrStoriesTooMuch
	// | 400 | STORIES_TOO_MUCH | You have hit the maximum active stories limit as specified by the [`story_expiring_limit_*` client configuration parameters](https://core.telegram.org/api/config#story-expiring-limit-default): you should buy a [Premium](https://core.telegram.org/api/premium) subscription, delete an active story, or wait for the oldest story to expire. |
	ErrStoriesTooMuch = ecode.NewCodeError(ErrBadRequest, "STORIES_TOO_MUCH")

	// ErrAboutTooLong
	// | 400 | ABOUT_TOO_LONG | About string too long. |
	ErrAboutTooLong = ecode.NewCodeError(ErrBadRequest, "ABOUT_TOO_LONG")

	// ErrEmailNotSetup
	// | 400 | EMAIL_NOT_SETUP | In order to change the login email with emailVerifyPurposeLoginChange, an existing login email must already be set using emailVerifyPurposeLoginSetup. |
	ErrEmailNotSetup = ecode.NewCodeError(ErrBadRequest, "EMAIL_NOT_SETUP")

	// ErrEncryptionAlreadyAccepted
	// | 400 | ENCRYPTION_ALREADY_ACCEPTED | Secret chat already accepted. |
	ErrEncryptionAlreadyAccepted = ecode.NewCodeError(ErrBadRequest, "ENCRYPTION_ALREADY_ACCEPTED")

	// ErrPasswordTooFresh_%d
	// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
	// ErrPasswordTooFresh_%d = ecode.NewCodeError(ErrBadRequest, "PASSWORD_TOO_FRESH_%d")

	// ErrReplyMarkupBuyEmpty
	// | 400 | REPLY_MARKUP_BUY_EMPTY | Reply markup for buy button empty. |
	ErrReplyMarkupBuyEmpty = ecode.NewCodeError(ErrBadRequest, "REPLY_MARKUP_BUY_EMPTY")

	// ErrUserBotRequired
	// | 400 | USER_BOT_REQUIRED | This method can only be called by a bot. |
	ErrUserBotRequired = ecode.NewCodeError(ErrBadRequest, "USER_BOT_REQUIRED")

	// ErrBotPaymentsDisabled
	// | 400 | BOT_PAYMENTS_DISABLED | Please enable bot payments in botfather before calling this method. |
	ErrBotPaymentsDisabled = ecode.NewCodeError(ErrBadRequest, "BOT_PAYMENTS_DISABLED")

	// ErrErrorTextEmpty
	// | 400 | ERROR_TEXT_EMPTY | The provided error message is empty. |
	ErrErrorTextEmpty = ecode.NewCodeError(ErrBadRequest, "ERROR_TEXT_EMPTY")

	// ErrGroupcallAlreadyDiscarded
	// | 400 | GROUPCALL_ALREADY_DISCARDED | The group call was already discarded. |
	ErrGroupcallAlreadyDiscarded = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_ALREADY_DISCARDED")

	// ErrMinDateInvalid
	// | 400 | MIN_DATE_INVALID | The specified minimum date is invalid. |
	ErrMinDateInvalid = ecode.NewCodeError(ErrBadRequest, "MIN_DATE_INVALID")

	// ErrScheduleDateTooLate
	// | 400 | SCHEDULE_DATE_TOO_LATE | You can't schedule a message this far in the future. |
	ErrScheduleDateTooLate = ecode.NewCodeError(ErrBadRequest, "SCHEDULE_DATE_TOO_LATE")

	// ErrUserAlreadyInvited
	// | 400 | USER_ALREADY_INVITED | You have already invited this user. |
	ErrUserAlreadyInvited = ecode.NewCodeError(ErrBadRequest, "USER_ALREADY_INVITED")

	// ErrFileEmtpy
	// | 400 | FILE_EMTPY | An empty file was provided. |
	ErrFileEmtpy = ecode.NewCodeError(ErrBadRequest, "FILE_EMTPY")

	// ErrInviteHashInvalid
	// | 400 | INVITE_HASH_INVALID | The invite hash is invalid. |
	ErrInviteHashInvalid = ecode.NewCodeError(ErrBadRequest, "INVITE_HASH_INVALID")

	// ErrPhotoExtInvalid
	// | 400 | PHOTO_EXT_INVALID | The extension of the photo is invalid. |
	ErrPhotoExtInvalid = ecode.NewCodeError(ErrBadRequest, "PHOTO_EXT_INVALID")

	// ErrPinnedDialogsTooMuch
	// | 400 | PINNED_DIALOGS_TOO_MUCH | Too many pinned dialogs. |
	ErrPinnedDialogsTooMuch = ecode.NewCodeError(ErrBadRequest, "PINNED_DIALOGS_TOO_MUCH")

	// ErrStickerVideoBig
	// | 400 | STICKER_VIDEO_BIG | The specified video sticker is too big. |
	ErrStickerVideoBig = ecode.NewCodeError(ErrBadRequest, "STICKER_VIDEO_BIG")

	// Err400ChatAdminRequired
	// | 400 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	Err400ChatAdminRequired = ecode.NewCodeError(ErrBadRequest, "CHAT_ADMIN_REQUIRED")

	// ErrCurrencyTotalAmountInvalid
	// | 400 | CURRENCY_TOTAL_AMOUNT_INVALID | The total amount of all prices is invalid. |
	ErrCurrencyTotalAmountInvalid = ecode.NewCodeError(ErrBadRequest, "CURRENCY_TOTAL_AMOUNT_INVALID")

	// ErrDataInvalid
	// | 400 | DATA_INVALID | Encrypted data invalid. |
	ErrDataInvalid = ecode.NewCodeError(ErrBadRequest, "DATA_INVALID")

	// ErrFilePartInvalid
	// | 400 | FILE_PART_INVALID | The file part number is invalid. |
	ErrFilePartInvalid = ecode.NewCodeError(ErrBadRequest, "FILE_PART_INVALID")

	// ErrGifContentTypeInvalid
	// | 400 | GIF_CONTENT_TYPE_INVALID | GIF content-type invalid. |
	ErrGifContentTypeInvalid = ecode.NewCodeError(ErrBadRequest, "GIF_CONTENT_TYPE_INVALID")

	// ErrPasswordMissing
	// | 400 | PASSWORD_MISSING | You must [enable 2FA](https://core.telegram.org/api/srp) before executing this operation. |
	ErrPasswordMissing = ecode.NewCodeError(ErrBadRequest, "PASSWORD_MISSING")

	// ErrPrivacyKeyInvalid
	// | 400 | PRIVACY_KEY_INVALID | The privacy key is invalid. |
	ErrPrivacyKeyInvalid = ecode.NewCodeError(ErrBadRequest, "PRIVACY_KEY_INVALID")

	// ErrChatNotModified
	// | 400 | CHAT_NOT_MODIFIED | No changes were made to chat information because the new information you passed is identical to the current information. |
	ErrChatNotModified = ecode.NewCodeError(ErrBadRequest, "CHAT_NOT_MODIFIED")

	// ErrEmoticonInvalid
	// | 400 | EMOTICON_INVALID | The specified emoji is invalid. |
	ErrEmoticonInvalid = ecode.NewCodeError(ErrBadRequest, "EMOTICON_INVALID")

	// ErrExternalUrlInvalid
	// | 400 | EXTERNAL_URL_INVALID | External URL invalid. |
	ErrExternalUrlInvalid = ecode.NewCodeError(ErrBadRequest, "EXTERNAL_URL_INVALID")

	// ErrFilterIncludeEmpty
	// | 400 | FILTER_INCLUDE_EMPTY | The include_peers vector of the filter is empty. |
	ErrFilterIncludeEmpty = ecode.NewCodeError(ErrBadRequest, "FILTER_INCLUDE_EMPTY")

	// ErrTopicNotModified
	// | 400 | TOPIC_NOT_MODIFIED | The updated topic info is equal to the current topic info, nothing was changed. |
	ErrTopicNotModified = ecode.NewCodeError(ErrBadRequest, "TOPIC_NOT_MODIFIED")

	// ErrMaxIdInvalid
	// | 400 | MAX_ID_INVALID | The provided max ID is invalid. |
	ErrMaxIdInvalid = ecode.NewCodeError(ErrBadRequest, "MAX_ID_INVALID")

	// ErrFilterNotSupported
	// | 400 | FILTER_NOT_SUPPORTED | The specified filter cannot be used in this context. |
	ErrFilterNotSupported = ecode.NewCodeError(ErrBadRequest, "FILTER_NOT_SUPPORTED")

	// ErrInviteForbiddenWithJoinas
	// | 400 | INVITE_FORBIDDEN_WITH_JOINAS | If the user has anonymously joined a group call as a channel, they can't invite other users to the group call because that would cause deanonymization, because the invite would be sent using the original user ID, not the anonymized channel ID. |
	ErrInviteForbiddenWithJoinas = ecode.NewCodeError(ErrBadRequest, "INVITE_FORBIDDEN_WITH_JOINAS")

	// ErrInviteRevokedMissing
	// | 400 | INVITE_REVOKED_MISSING | The specified invite link was already revoked or is invalid. |
	ErrInviteRevokedMissing = ecode.NewCodeError(ErrBadRequest, "INVITE_REVOKED_MISSING")

	// ErrPeerIdInvalid
	// | 400 | PEER_ID_INVALID | The provided peer id is invalid. |
	ErrPeerIdInvalid = ecode.NewCodeError(ErrBadRequest, "PEER_ID_INVALID")

	// ErrPollQuestionInvalid
	// | 400 | POLL_QUESTION_INVALID | One of the poll questions is not acceptable. |
	ErrPollQuestionInvalid = ecode.NewCodeError(ErrBadRequest, "POLL_QUESTION_INVALID")

	// ErrScheduleStatusPrivate
	// | 400 | SCHEDULE_STATUS_PRIVATE | Can't schedule until user is online, if the user's last seen timestamp is hidden by their privacy settings. |
	ErrScheduleStatusPrivate = ecode.NewCodeError(ErrBadRequest, "SCHEDULE_STATUS_PRIVATE")

	// ErrCodeInvalid
	// | 400 | CODE_INVALID | Code invalid. |
	ErrCodeInvalid = ecode.NewCodeError(ErrBadRequest, "CODE_INVALID")

	// ErrFormExpired
	// | 400 | FORM_EXPIRED | The form was generated more than 10 minutes ago and has expired, please re-generate it using [payments.getPaymentForm](https://core.telegram.org/method/payments.getPaymentForm) and pass the new `form_id`. |
	ErrFormExpired = ecode.NewCodeError(ErrBadRequest, "FORM_EXPIRED")

	// ErrMd5ChecksumInvalid
	// | 400 | MD5_CHECKSUM_INVALID | The MD5 checksums do not match. |
	ErrMd5ChecksumInvalid = ecode.NewCodeError(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrQuickRepliesTooMuch
	// | 400 | QUICK_REPLIES_TOO_MUCH | A maximum of [appConfig.`quick_replies_limit`](https://core.telegram.org/api/config#quick-replies-limit) shortcuts may be created, the limit was reached. |
	ErrQuickRepliesTooMuch = ecode.NewCodeError(ErrBadRequest, "QUICK_REPLIES_TOO_MUCH")

	// ErrTypesEmpty
	// | 400 | TYPES_EMPTY | No top peer type was provided. |
	ErrTypesEmpty = ecode.NewCodeError(ErrBadRequest, "TYPES_EMPTY")

	// ErrBotInlineDisabled
	// | 400 | BOT_INLINE_DISABLED | This bot can't be used in inline mode. |
	ErrBotInlineDisabled = ecode.NewCodeError(ErrBadRequest, "BOT_INLINE_DISABLED")

	// ErrGroupcallNotModified
	// | 400 | GROUPCALL_NOT_MODIFIED | Group call settings weren't modified. |
	ErrGroupcallNotModified = ecode.NewCodeError(ErrBadRequest, "GROUPCALL_NOT_MODIFIED")

	// ErrLimitInvalid
	// | 400 | LIMIT_INVALID | The provided limit is invalid. |
	ErrLimitInvalid = ecode.NewCodeError(ErrBadRequest, "LIMIT_INVALID")

	// ErrPhotoContentTypeInvalid
	// | 400 | PHOTO_CONTENT_TYPE_INVALID | Photo mime-type invalid. |
	ErrPhotoContentTypeInvalid = ecode.NewCodeError(ErrBadRequest, "PHOTO_CONTENT_TYPE_INVALID")

	// ErrMediaFileInvalid
	// | 400 | MEDIA_FILE_INVALID | The specified media file is invalid. |
	ErrMediaFileInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_FILE_INVALID")

	// ErrMediaTtlInvalid
	// | 400 | MEDIA_TTL_INVALID | The specified media TTL is invalid. |
	ErrMediaTtlInvalid = ecode.NewCodeError(ErrBadRequest, "MEDIA_TTL_INVALID")

	// ErrPhoneNumberFlood
	// | 400 | PHONE_NUMBER_FLOOD | You asked for the code too many times. |
	ErrPhoneNumberFlood = ecode.NewCodeError(ErrBadRequest, "PHONE_NUMBER_FLOOD")

	// Err400UserInvalid
	// | 400 | USER_INVALID | Invalid user provided. |
	Err400UserInvalid = ecode.NewCodeError(ErrBadRequest, "USER_INVALID")

	// ErrBotInvalid
	// | 400 | BOT_INVALID | This is not a valid bot. |
	ErrBotInvalid = ecode.NewCodeError(ErrBadRequest, "BOT_INVALID")

	// ErrChannelsAdminLocatedTooMuch
	// | 400 | CHANNELS_ADMIN_LOCATED_TOO_MUCH | The user has reached the limit of public geogroups. |
	ErrChannelsAdminLocatedTooMuch = ecode.NewCodeError(ErrBadRequest, "CHANNELS_ADMIN_LOCATED_TOO_MUCH")

	// ErrFilterTitleEmpty
	// | 400 | FILTER_TITLE_EMPTY | The title field of the filter is empty. |
	ErrFilterTitleEmpty = ecode.NewCodeError(ErrBadRequest, "FILTER_TITLE_EMPTY")

	// ErrStoryIdInvalid
	// | 400 | STORY_ID_INVALID | The specified story ID is invalid. |
	ErrStoryIdInvalid = ecode.NewCodeError(ErrBadRequest, "STORY_ID_INVALID")

	// ErrTokenEmpty
	// | 400 | TOKEN_EMPTY | The specified token is empty. |
	ErrTokenEmpty = ecode.NewCodeError(ErrBadRequest, "TOKEN_EMPTY")

	// ErrDataTooLong
	// | 400 | DATA_TOO_LONG | Data too long. |
	ErrDataTooLong = ecode.NewCodeError(ErrBadRequest, "DATA_TOO_LONG")

	// ErrParticipantVersionOutdated
	// | 400 | PARTICIPANT_VERSION_OUTDATED | The other participant does not use an up to date telegram client with support for calls. |
	ErrParticipantVersionOutdated = ecode.NewCodeError(ErrBadRequest, "PARTICIPANT_VERSION_OUTDATED")

	// ErrScheduleBotNotAllowed
	// | 400 | SCHEDULE_BOT_NOT_ALLOWED | Bots cannot schedule messages. |
	ErrScheduleBotNotAllowed = ecode.NewCodeError(ErrBadRequest, "SCHEDULE_BOT_NOT_ALLOWED")

	// ErrFilePartEmpty
	// | 400 | FILE_PART_EMPTY | The provided file part is empty. |
	ErrFilePartEmpty = ecode.NewCodeError(ErrBadRequest, "FILE_PART_EMPTY")

	// ErrSwitchPmTextEmpty
	// | 400 | SWITCH_PM_TEXT_EMPTY | The switch_pm.text field was empty. |
	ErrSwitchPmTextEmpty = ecode.NewCodeError(ErrBadRequest, "SWITCH_PM_TEXT_EMPTY")

	// ErrTempAuthKeyAlreadyBound
	// | 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another **perm_auth_key_id**. |
	ErrTempAuthKeyAlreadyBound = ecode.NewCodeError(ErrBadRequest, "TEMP_AUTH_KEY_ALREADY_BOUND")

	// ErrMessagePollClosed
	// | 400 | MESSAGE_POLL_CLOSED | Poll closed. |
	ErrMessagePollClosed = ecode.NewCodeError(ErrBadRequest, "MESSAGE_POLL_CLOSED")

	// ErrReplyMessageIdInvalid
	// | 400 | REPLY_MESSAGE_ID_INVALID | The specified reply-to message ID is invalid. |
	ErrReplyMessageIdInvalid = ecode.NewCodeError(ErrBadRequest, "REPLY_MESSAGE_ID_INVALID")

	// ErrInputTextEmpty
	// | 400 | INPUT_TEXT_EMPTY | The specified text is empty. |
	ErrInputTextEmpty = ecode.NewCodeError(ErrBadRequest, "INPUT_TEXT_EMPTY")

	// ErrLanguageInvalid
	// | 400 | LANGUAGE_INVALID | The specified lang_code is invalid. |
	ErrLanguageInvalid = ecode.NewCodeError(ErrBadRequest, "LANGUAGE_INVALID")

	// ErrMegagroupPrehistoryHidden
	// | 400 | MEGAGROUP_PREHISTORY_HIDDEN | Group with hidden history for new members can't be set as discussion groups. |
	ErrMegagroupPrehistoryHidden = ecode.NewCodeError(ErrBadRequest, "MEGAGROUP_PREHISTORY_HIDDEN")

	// ErrJoinAsPeerInvalid
	// | 400 | JOIN_AS_PEER_INVALID | The specified peer cannot be used to join a group call. |
	ErrJoinAsPeerInvalid = ecode.NewCodeError(ErrBadRequest, "JOIN_AS_PEER_INVALID")

	// ErrSendAsPeerInvalid
	// | 400 | SEND_AS_PEER_INVALID | You can't send messages as the specified peer. |
	ErrSendAsPeerInvalid = ecode.NewCodeError(ErrBadRequest, "SEND_AS_PEER_INVALID")

	// ErrWebdocumentUrlInvalid
	// | 400 | WEBDOCUMENT_URL_INVALID | The specified webdocument URL is invalid. |
	ErrWebdocumentUrlInvalid = ecode.NewCodeError(ErrBadRequest, "WEBDOCUMENT_URL_INVALID")

	// Err400ChatSendInlineForbidden
	// | 400 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	Err400ChatSendInlineForbidden = ecode.NewCodeError(ErrBadRequest, "CHAT_SEND_INLINE_FORBIDDEN")

	// ErrGraphExpiredReload
	// | 400 | GRAPH_EXPIRED_RELOAD | This graph has expired, please obtain a new graph token. |
	ErrGraphExpiredReload = ecode.NewCodeError(ErrBadRequest, "GRAPH_EXPIRED_RELOAD")

	// ErrImportFileInvalid
	// | 400 | IMPORT_FILE_INVALID | The specified chat export file is invalid. |
	ErrImportFileInvalid = ecode.NewCodeError(ErrBadRequest, "IMPORT_FILE_INVALID")

	// ErrQueryIdEmpty
	// | 400 | QUERY_ID_EMPTY | The query ID is empty. |
	ErrQueryIdEmpty = ecode.NewCodeError(ErrBadRequest, "QUERY_ID_EMPTY")

	// ErrStoryIdEmpty
	// | 400 | STORY_ID_EMPTY | You specified no story IDs. |
	ErrStoryIdEmpty = ecode.NewCodeError(ErrBadRequest, "STORY_ID_EMPTY")

	// ErrTtlPeriodInvalid
	// | 400 | TTL_PERIOD_INVALID | The specified TTL period is invalid. |
	ErrTtlPeriodInvalid = ecode.NewCodeError(ErrBadRequest, "TTL_PERIOD_INVALID")

	// ErrUserBot
	// | 400 | USER_BOT | Bots can only be admins in channels. |
	ErrUserBot = ecode.NewCodeError(ErrBadRequest, "USER_BOT")

	// ErrEmailHashExpired
	// | 400 | EMAIL_HASH_EXPIRED | Email hash expired. |
	ErrEmailHashExpired = ecode.NewCodeError(ErrBadRequest, "EMAIL_HASH_EXPIRED")

	// ErrNewSettingsEmpty
	// | 400 | NEW_SETTINGS_EMPTY | No password is set on the current account, and no new password was specified in `new_settings`. |
	ErrNewSettingsEmpty = ecode.NewCodeError(ErrBadRequest, "NEW_SETTINGS_EMPTY")

	// ErrPollAnswerInvalid
	// | 400 | POLL_ANSWER_INVALID | One of the poll answers is not acceptable. |
	ErrPollAnswerInvalid = ecode.NewCodeError(ErrBadRequest, "POLL_ANSWER_INVALID")

	// ErrMaxQtsInvalid
	// | 400 | MAX_QTS_INVALID | The specified max_qts is invalid. |
	ErrMaxQtsInvalid = ecode.NewCodeError(ErrBadRequest, "MAX_QTS_INVALID")

	// ErrPollOptionInvalid
	// | 400 | POLL_OPTION_INVALID | Invalid poll option provided. |
	ErrPollOptionInvalid = ecode.NewCodeError(ErrBadRequest, "POLL_OPTION_INVALID")

	// ErrStartParamTooLong
	// | 400 | START_PARAM_TOO_LONG | Start parameter is too long. |
	ErrStartParamTooLong = ecode.NewCodeError(ErrBadRequest, "START_PARAM_TOO_LONG")

	// ErrAuthBytesInvalid
	// | 400 | AUTH_BYTES_INVALID | The provided authorization is invalid. |
	ErrAuthBytesInvalid = ecode.NewCodeError(ErrBadRequest, "AUTH_BYTES_INVALID")

	// ErrChannelForumMissing
	// | 400 | CHANNEL_FORUM_MISSING | This supergroup is not a forum. |
	ErrChannelForumMissing = ecode.NewCodeError(ErrBadRequest, "CHANNEL_FORUM_MISSING")

	// ErrGeneralModifyIconForbidden
	// | 400 | GENERAL_MODIFY_ICON_FORBIDDEN | You can't modify the icon of the "General" topic. |
	ErrGeneralModifyIconForbidden = ecode.NewCodeError(ErrBadRequest, "GENERAL_MODIFY_ICON_FORBIDDEN")

	// ErrAdminRankInvalid
	// | 400 | ADMIN_RANK_INVALID | The specified admin rank is invalid. |
	ErrAdminRankInvalid = ecode.NewCodeError(ErrBadRequest, "ADMIN_RANK_INVALID")
)

// 401 UNAUTHORIZED
var (
	// ErrSessionPasswordNeeded
	// 401	SESSION_PASSWORD_NEEDED	The user has enabled 2FA, more steps are needed
	ErrSessionPasswordNeeded = ecode.NewCodeError(ErrUnauthorized, "SESSION_PASSWORD_NEEDED")

	// ErrAuthKeyUnregistered
	// AUTH_KEY_UNREGISTERED: The key is not registered in the system
	ErrAuthKeyUnregistered = ecode.NewCodeError(ErrUnauthorized, "AUTH_KEY_UNREGISTERED")

	// ErrAuthKeyInvalid
	// | 401 | AUTH_KEY_INVALID | Auth key invalid. |
	ErrAuthKeyInvalid = ecode.NewCodeError(ErrUnauthorized, "AUTH_KEY_INVALID")

	// ErrUserDeactivated
	// USER_DEACTIVATED: The user has been deleted/deactivated
	ErrUserDeactivated = ecode.NewCodeError(ErrUnauthorized, "USER_DEACTIVATED")

	// ErrSessionRevoked
	// SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	ErrSessionRevoked = ecode.NewCodeError(ErrUnauthorized, "SESSION_REVOKED")

	// ErrSessionExpired
	// SESSION_EXPIRED: The authorization has expired
	ErrSessionExpired = ecode.NewCodeError(ErrUnauthorized, "SESSION_EXPIRED")

	// ErrAuthKeyPermEmpty
	// | 401 | AUTH_KEY_PERM_EMPTY | The temporary auth key must be binded to the permanent auth key to use these methods. |
	ErrAuthKeyPermEmpty = ecode.NewCodeError(ErrUnauthorized, "AUTH_KEY_PERM_EMPTY")

	// ErrActiveUserRequired
	// ACTIVE_USER_REQUIRED	401	The method is only available to already activated users
	ErrActiveUserRequired = ecode.NewCodeError(ErrUnauthorized, "ACTIVE_USER_REQUIRED")
)

// 403 FORBIDDEN
var (

	// Err403PremiumAccountRequired
	// | 403 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	Err403PremiumAccountRequired = ecode.NewCodeError(ErrForbidden, "PREMIUM_ACCOUNT_REQUIRED")

	// ErrChatSendGameForbidden
	// | 403 | CHAT_SEND_GAME_FORBIDDEN | You can't send a game to this chat. |
	ErrChatSendGameForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_GAME_FORBIDDEN")

	// ErrChatSendVideosForbidden
	// | 403 | CHAT_SEND_VIDEOS_FORBIDDEN | You can't send videos in this chat. |
	ErrChatSendVideosForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_VIDEOS_FORBIDDEN")

	// ErrMessageAuthorRequired
	// | 403 | MESSAGE_AUTHOR_REQUIRED | Message author required. |
	ErrMessageAuthorRequired = ecode.NewCodeError(ErrForbidden, "MESSAGE_AUTHOR_REQUIRED")

	// ErrChatSendPollForbidden
	// | 403 | CHAT_SEND_POLL_FORBIDDEN | You can't send polls in this chat. |
	ErrChatSendPollForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_POLL_FORBIDDEN")

	// ErrMessageDeleteForbidden
	// | 403 | MESSAGE_DELETE_FORBIDDEN | You can't delete one of the messages you tried to delete, most likely because it is a service message. |
	ErrMessageDeleteForbidden = ecode.NewCodeError(ErrForbidden, "MESSAGE_DELETE_FORBIDDEN")

	// Err403UserBotInvalid
	// | 403 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	Err403UserBotInvalid = ecode.NewCodeError(ErrForbidden, "USER_BOT_INVALID")

	// ErrInlineBotRequired
	// | 403 | INLINE_BOT_REQUIRED | Only the inline bot can edit message. |
	ErrInlineBotRequired = ecode.NewCodeError(ErrForbidden, "INLINE_BOT_REQUIRED")

	// Err403UserChannelsTooMuch
	// | 403 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	Err403UserChannelsTooMuch = ecode.NewCodeError(ErrForbidden, "USER_CHANNELS_TOO_MUCH")

	// ErrUserDeleted
	// | 403 | USER_DELETED | You can't send this secret message because the other participant deleted their account. |
	ErrUserDeleted = ecode.NewCodeError(ErrForbidden, "USER_DELETED")

	// Err403UserNotParticipant
	// | 403 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	Err403UserNotParticipant = ecode.NewCodeError(ErrForbidden, "USER_NOT_PARTICIPANT")

	// Err403UserRestricted
	// | 403 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	Err403UserRestricted = ecode.NewCodeError(ErrForbidden, "USER_RESTRICTED")

	// ErrChannelPublicGroupNa
	// | 403 | CHANNEL_PUBLIC_GROUP_NA | channel/supergroup not available. |
	ErrChannelPublicGroupNa = ecode.NewCodeError(ErrForbidden, "CHANNEL_PUBLIC_GROUP_NA")

	// ErrChatActionForbidden
	// | 403 | CHAT_ACTION_FORBIDDEN | You cannot execute this action. |
	ErrChatActionForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_ACTION_FORBIDDEN")

	// Err403ChatSendInlineForbidden
	// | 403 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	Err403ChatSendInlineForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_INLINE_FORBIDDEN")

	// Err403NotEligible
	// | 403 | NOT_ELIGIBLE | The current user is not eligible to join the Peer-to-Peer Login Program. |
	Err403NotEligible = ecode.NewCodeError(ErrForbidden, "NOT_ELIGIBLE")

	// Err403ParticipantJoinMissing
	// | 403 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	Err403ParticipantJoinMissing = ecode.NewCodeError(ErrForbidden, "PARTICIPANT_JOIN_MISSING")

	// Err403UserInvalid
	// | 403 | USER_INVALID | Invalid user provided. |
	Err403UserInvalid = ecode.NewCodeError(ErrForbidden, "USER_INVALID")

	// ErrChatSendDocsForbidden
	// | 403 | CHAT_SEND_DOCS_FORBIDDEN | You can't send documents in this chat. |
	ErrChatSendDocsForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_DOCS_FORBIDDEN")

	// ErrChatSendGifsForbidden
	// | 403 | CHAT_SEND_GIFS_FORBIDDEN | You can't send gifs in this chat. |
	ErrChatSendGifsForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_GIFS_FORBIDDEN")

	// ErrChatSendPhotosForbidden
	// | 403 | CHAT_SEND_PHOTOS_FORBIDDEN | You can't send photos in this chat. |
	ErrChatSendPhotosForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_PHOTOS_FORBIDDEN")

	// ErrChatSendRoundvideosForbidden
	// | 403 | CHAT_SEND_ROUNDVIDEOS_FORBIDDEN | You can't send round videos to this chat. |
	ErrChatSendRoundvideosForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_ROUNDVIDEOS_FORBIDDEN")

	// ErrGroupcallAlreadyStarted
	// | 403 | GROUPCALL_ALREADY_STARTED | The groupcall has already started, you can join directly using [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrGroupcallAlreadyStarted = ecode.NewCodeError(ErrForbidden, "GROUPCALL_ALREADY_STARTED")

	// ErrRightForbidden
	// | 403 | RIGHT_FORBIDDEN | Your admin rights do not allow you to do this. |
	ErrRightForbidden = ecode.NewCodeError(ErrForbidden, "RIGHT_FORBIDDEN")

	// Err403UserNotMutualContact
	// | 403 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	Err403UserNotMutualContact = ecode.NewCodeError(ErrForbidden, "USER_NOT_MUTUAL_CONTACT")

	// ErrYourPrivacyRestricted
	// | 403 | YOUR_PRIVACY_RESTRICTED | You cannot fetch the read date of this message because you have disallowed other users to do so for *your* messages; to fix, allow other users to see *your* exact last online date OR purchase a [Telegram Premium](https://core.telegram.org/api/premium) subscription. |
	ErrYourPrivacyRestricted = ecode.NewCodeError(ErrForbidden, "YOUR_PRIVACY_RESTRICTED")

	// Err403ChatAdminRequired
	// | 403 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	Err403ChatAdminRequired = ecode.NewCodeError(ErrForbidden, "CHAT_ADMIN_REQUIRED")

	// ErrChatGuestSendForbidden
	// | 403 | CHAT_GUEST_SEND_FORBIDDEN | You join the discussion group before commenting, see [here &raquo;](https://core.telegram.org/api/discussion#requiring-users-to-join-the-group) for more info. |
	ErrChatGuestSendForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_GUEST_SEND_FORBIDDEN")

	// ErrChatSendMediaForbidden
	// | 403 | CHAT_SEND_MEDIA_FORBIDDEN | You can't send media in this chat. |
	ErrChatSendMediaForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_MEDIA_FORBIDDEN")

	// ErrChatSendAudiosForbidden
	// | 403 | CHAT_SEND_AUDIOS_FORBIDDEN | You can't send audio messages in this chat. |
	ErrChatSendAudiosForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_AUDIOS_FORBIDDEN")

	// ErrChatSendPlainForbidden
	// | 403 | CHAT_SEND_PLAIN_FORBIDDEN | You can't send non-media (text) messages in this chat. |
	ErrChatSendPlainForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_PLAIN_FORBIDDEN")

	// ErrChatSendStickersForbidden
	// | 403 | CHAT_SEND_STICKERS_FORBIDDEN | You can't send stickers in this chat. |
	ErrChatSendStickersForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_STICKERS_FORBIDDEN")

	// ErrChatSendVoicesForbidden
	// | 403 | CHAT_SEND_VOICES_FORBIDDEN | You can't send voice recordings in this chat. |
	ErrChatSendVoicesForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_SEND_VOICES_FORBIDDEN")

	// ErrSensitiveChangeForbidden
	// | 403 | SENSITIVE_CHANGE_FORBIDDEN | You can't change your sensitive content settings. |
	ErrSensitiveChangeForbidden = ecode.NewCodeError(ErrForbidden, "SENSITIVE_CHANGE_FORBIDDEN")

	// ErrAnonymousReactionsDisabled
	// | 403 | ANONYMOUS_REACTIONS_DISABLED | Sorry, anonymous administrators cannot leave reactions or participate in polls. |
	ErrAnonymousReactionsDisabled = ecode.NewCodeError(ErrForbidden, "ANONYMOUS_REACTIONS_DISABLED")

	// ErrBroadcastForbidden
	// | 403 | BROADCAST_FORBIDDEN | Channel poll voters and reactions cannot be fetched to prevent deanonymization. |
	ErrBroadcastForbidden = ecode.NewCodeError(ErrForbidden, "BROADCAST_FORBIDDEN")

	// ErrChatAdminInviteRequired
	// | 403 | CHAT_ADMIN_INVITE_REQUIRED | You do not have the rights to do this. |
	ErrChatAdminInviteRequired = ecode.NewCodeError(ErrForbidden, "CHAT_ADMIN_INVITE_REQUIRED")

	// ErrPublicChannelMissing
	// | 403 | PUBLIC_CHANNEL_MISSING | You can only export group call invite links for public chats or channels. |
	ErrPublicChannelMissing = ecode.NewCodeError(ErrForbidden, "PUBLIC_CHANNEL_MISSING")

	// Err403TakeoutRequired
	// | 403 | TAKEOUT_REQUIRED | A [takeout](https://core.telegram.org/api/takeout) session needs to be initialized first, [see here &raquo; for more info](https://core.telegram.org/api/takeout). |
	Err403TakeoutRequired = ecode.NewCodeError(ErrForbidden, "TAKEOUT_REQUIRED")

	// Err403UserIsBlocked
	// | 403 | USER_IS_BLOCKED | You were blocked by this user. |
	Err403UserIsBlocked = ecode.NewCodeError(ErrForbidden, "USER_IS_BLOCKED")

	// ErrUserPrivacyRestricted
	// | 403 | USER_PRIVACY_RESTRICTED | The user's privacy settings do not allow you to do this. |
	ErrUserPrivacyRestricted = ecode.NewCodeError(ErrForbidden, "USER_PRIVACY_RESTRICTED")

	// ErrChatForbidden
	// | 403 | CHAT_FORBIDDEN | This chat is not available to the current user. |
	ErrChatForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_FORBIDDEN")

	// Err403GroupcallForbidden
	// | 403 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	Err403GroupcallForbidden = ecode.NewCodeError(ErrForbidden, "GROUPCALL_FORBIDDEN")

	// ErrPollVoteRequired
	// | 403 | POLL_VOTE_REQUIRED | Cast a vote in the poll before calling this method. |
	ErrPollVoteRequired = ecode.NewCodeError(ErrForbidden, "POLL_VOTE_REQUIRED")

	// Err403PrivacyPremiumRequired
	// | 403 | PRIVACY_PREMIUM_REQUIRED | You need a [Telegram Premium subscription](https://core.telegram.org/api/premium) to send a message to this user. |
	Err403PrivacyPremiumRequired = ecode.NewCodeError(ErrForbidden, "PRIVACY_PREMIUM_REQUIRED")

	// ErrChatWriteForbidden
	// | 403 | CHAT_WRITE_FORBIDDEN | You can't write in this chat. |
	ErrChatWriteForbidden = ecode.NewCodeError(ErrForbidden, "CHAT_WRITE_FORBIDDEN")

	// ErrEditBotInviteForbidden
	// | 403 | EDIT_BOT_INVITE_FORBIDDEN | Normal users can't edit invites that were created by bots. |
	ErrEditBotInviteForbidden = ecode.NewCodeError(ErrForbidden, "EDIT_BOT_INVITE_FORBIDDEN")

	// Err403VoiceMessagesForbidden
	// | 403 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	Err403VoiceMessagesForbidden = ecode.NewCodeError(ErrForbidden, "VOICE_MESSAGES_FORBIDDEN")
)

// 406 NOT_ACCEPTABLE

// NewErrPreviousChatImportActiveWaitX
// ErrPreviousChatImportActiveWait_%dmin
// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
func NewErrPreviousChatImportActiveWaitX(minute int32) error {
	return ecode.NewCodeErrorf(ErrBadRequest, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%ddMIN", minute)
}

var (

	// Err406FreshChangeAdminsForbidden
	// | 406 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	Err406FreshChangeAdminsForbidden = ecode.NewCodeError(ErrNotAcceptable, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// Err406PhoneNumberInvalid
	// | 406 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	Err406PhoneNumberInvalid = ecode.NewCodeError(ErrNotAcceptable, "PHONE_NUMBER_INVALID")

	// ErrPremiumCurrentlyUnavailable
	// | 406 | PREMIUM_CURRENTLY_UNAVAILABLE | You cannot currently purchase a Premium subscription. |
	ErrPremiumCurrentlyUnavailable = ecode.NewCodeError(ErrNotAcceptable, "PREMIUM_CURRENTLY_UNAVAILABLE")

	// Err406TopicClosed
	// | 406 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	Err406TopicClosed = ecode.NewCodeError(ErrNotAcceptable, "TOPIC_CLOSED")

	// Err406TopicDeleted
	// | 406 | TOPIC_DELETED | The specified topic was deleted. |
	Err406TopicDeleted = ecode.NewCodeError(ErrNotAcceptable, "TOPIC_DELETED")

	// ErrFreshResetAuthorisationForbidden
	// | 406 | FRESH_RESET_AUTHORISATION_FORBIDDEN | You can't logout other sessions if less than 24 hours have passed since you logged on the current session. |
	ErrFreshResetAuthorisationForbidden = ecode.NewCodeError(ErrNotAcceptable, "FRESH_RESET_AUTHORISATION_FORBIDDEN")

	// Err406PrivacyPremiumRequired
	// | 406 | PRIVACY_PREMIUM_REQUIRED | You need a [Telegram Premium subscription](https://core.telegram.org/api/premium) to send a message to this user. |
	Err406PrivacyPremiumRequired = ecode.NewCodeError(ErrNotAcceptable, "PRIVACY_PREMIUM_REQUIRED")

	// Err406UserpicUploadRequired
	// | 406 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	Err406UserpicUploadRequired = ecode.NewCodeError(ErrNotAcceptable, "USERPIC_UPLOAD_REQUIRED")

	// ErrBusinessAddressActive
	// | 406 | BUSINESS_ADDRESS_ACTIVE | The user is currently advertising a [Business Location](https://core.telegram.org/api/business#location), the location may only be changed (or removed) using [account.updateBusinessLocation &raquo;](https://core.telegram.org/method/account.updateBusinessLocation).  . |
	ErrBusinessAddressActive = ecode.NewCodeError(ErrNotAcceptable, "BUSINESS_ADDRESS_ACTIVE")

	// Err406ChannelPrivate
	// | 406 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	Err406ChannelPrivate = ecode.NewCodeError(ErrNotAcceptable, "CHANNEL_PRIVATE")

	// ErrSendCodeUnavailable
	// | 406 | SEND_CODE_UNAVAILABLE | Returned when all available options for this type of number were already used (e.g. flash-call, then SMS, then this error might be returned to trigger a second resend). |
	ErrSendCodeUnavailable = ecode.NewCodeError(ErrNotAcceptable, "SEND_CODE_UNAVAILABLE")

	// ErrUpdateAppToLogin
	// | 406 | UPDATE_APP_TO_LOGIN | Please update your client to login. |
	ErrUpdateAppToLogin = ecode.NewCodeError(ErrNotAcceptable, "UPDATE_APP_TO_LOGIN")

	// Err406UserRestricted
	// | 406 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	Err406UserRestricted = ecode.NewCodeError(ErrNotAcceptable, "USER_RESTRICTED")

	// ErrUserpicPrivacyRequired
	// | 406 | USERPIC_PRIVACY_REQUIRED | You need to disable privacy settings for your profile picture in order to make your geolocation public. |
	ErrUserpicPrivacyRequired = ecode.NewCodeError(ErrNotAcceptable, "USERPIC_PRIVACY_REQUIRED")

	// ErrAuthKeyDuplicated
	// | 406 | AUTH_KEY_DUPLICATED | Concurrent usage of the current session from multiple connections was detected, the current session was invalidated by the server for security reasons! |
	ErrAuthKeyDuplicated = ecode.NewCodeError(ErrNotAcceptable, "AUTH_KEY_DUPLICATED")

	// ErrCallProtocolCompatLayerInvalid
	// | 406 | CALL_PROTOCOL_COMPAT_LAYER_INVALID | The other side of the call does not support any of the VoIP protocols supported by the local client, as specified by the `protocol.layer` and `protocol.library_versions` fields. |
	ErrCallProtocolCompatLayerInvalid = ecode.NewCodeError(ErrNotAcceptable, "CALL_PROTOCOL_COMPAT_LAYER_INVALID")

	// Err406ChannelTooLarge
	// | 406 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	Err406ChannelTooLarge = ecode.NewCodeError(ErrNotAcceptable, "CHANNEL_TOO_LARGE")

	// ErrFilerefUpgradeNeeded
	// | 406 | FILEREF_UPGRADE_NEEDED | The client has to be updated in order to support [file references](https://core.telegram.org/api/file_reference). |
	ErrFilerefUpgradeNeeded = ecode.NewCodeError(ErrNotAcceptable, "FILEREF_UPGRADE_NEEDED")

	// ErrFreshChangePhoneForbidden
	// | 406 | FRESH_CHANGE_PHONE_FORBIDDEN | You can't change phone number right after logging in, please wait at least 24 hours. |
	ErrFreshChangePhoneForbidden = ecode.NewCodeError(ErrNotAcceptable, "FRESH_CHANGE_PHONE_FORBIDDEN")

	// ErrPaymentUnsupported
	// | 406 | PAYMENT_UNSUPPORTED | A detailed description of the error will be received separately as described [here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrPaymentUnsupported = ecode.NewCodeError(ErrNotAcceptable, "PAYMENT_UNSUPPORTED")

	// ErrStickersetOwnerAnonymous
	// | 406 | STICKERSET_OWNER_ANONYMOUS | Provided stickerset can't be installed as group stickerset to prevent admin deanonymization. |
	ErrStickersetOwnerAnonymous = ecode.NewCodeError(ErrNotAcceptable, "STICKERSET_OWNER_ANONYMOUS")

	// Err406ChatForwardsRestricted
	// | 406 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	Err406ChatForwardsRestricted = ecode.NewCodeError(ErrNotAcceptable, "CHAT_FORWARDS_RESTRICTED")

	// Err406StickersetInvalid
	// | 406 | STICKERSET_INVALID | The provided sticker set is invalid. |
	Err406StickersetInvalid = ecode.NewCodeError(ErrNotAcceptable, "STICKERSET_INVALID")

	// Err406BannedRightsInvalid
	// | 406 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	Err406BannedRightsInvalid = ecode.NewCodeError(ErrNotAcceptable, "BANNED_RIGHTS_INVALID")

	// Err406InviteHashExpired
	// | 406 | INVITE_HASH_EXPIRED | The invite link has expired. |
	Err406InviteHashExpired = ecode.NewCodeError(ErrNotAcceptable, "INVITE_HASH_EXPIRED")

	// ErrPhonePasswordFlood
	// | 406 | PHONE_PASSWORD_FLOOD | You have tried logging in too many times. |
	ErrPhonePasswordFlood = ecode.NewCodeError(ErrNotAcceptable, "PHONE_PASSWORD_FLOOD")

	// ErrPreviousChatImportActiveWait_%dmin
	// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
	// ErrPreviousChatImportActiveWait_%dmin = ecode.NewCodeError(ErrNotAcceptable, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN")

)

// 500 InternalServerError
var (
	// StatusInternalServerError - StatusInternelServerError
	// StatusInternalServerError = ecode.NewCodeError(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrInternalServerError
	// | 500 | INTERNAL_SERVER_ERROR |  |
	ErrInternalServerError = ecode.NewCodeError(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrAuthRestart
	// | 500 | AUTH_RESTART | Restart the authorization process. |
	ErrAuthRestart = ecode.NewCodeError(ErrInternal, "AUTH_RESTART")

	// Err500CallOccupyFailed
	// | 500 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	Err500CallOccupyFailed = ecode.NewCodeError(ErrInternal, "CALL_OCCUPY_FAILED")

	// ErrCdnUploadTimeout
	// | 500 | CDN_UPLOAD_TIMEOUT | A server-side timeout occurred while reuploading the file to the CDN DC. |
	ErrCdnUploadTimeout = ecode.NewCodeError(ErrInternal, "CDN_UPLOAD_TIMEOUT")

	// ErrChatIdGenerateFailed
	// | 500 | CHAT_ID_GENERATE_FAILED | Failure while generating the chat ID. |
	ErrChatIdGenerateFailed = ecode.NewCodeError(ErrInternal, "CHAT_ID_GENERATE_FAILED")

	// Err500ChatInvalid
	// | 500 | CHAT_INVALID | Invalid chat. |
	Err500ChatInvalid = ecode.NewCodeError(ErrInternal, "CHAT_INVALID")

	// ErrPersistentTimestampOutdated
	// | 500 | PERSISTENT_TIMESTAMP_OUTDATED | Channel internal replication issues, try again later (treat this like an RPC_CALL_FAIL). |
	ErrPersistentTimestampOutdated = ecode.NewCodeError(ErrInternal, "PERSISTENT_TIMESTAMP_OUTDATED")

	// ErrRandomIdDuplicate
	// | 500 | RANDOM_ID_DUPLICATE | You provided a random ID that was already used. |
	ErrRandomIdDuplicate = ecode.NewCodeError(ErrInternal, "RANDOM_ID_DUPLICATE")

	// ErrSendMediaInvalid
	// | 500 | SEND_MEDIA_INVALID | The specified media is invalid. |
	ErrSendMediaInvalid = ecode.NewCodeError(ErrInternal, "SEND_MEDIA_INVALID")

	// Err500MsgWaitFailed
	// | 500 | MSG_WAIT_FAILED | A waiting call returned an error. |
	Err500MsgWaitFailed = ecode.NewCodeError(ErrInternal, "MSG_WAIT_FAILED")

	// ErrSignInFailed
	// | 500 | SIGN_IN_FAILED | Failure while signing in. |
	ErrSignInFailed = ecode.NewCodeError(ErrInternal, "SIGN_IN_FAILED")
)

// -503 Timeout
var (
	// StatusTimeout - StatusTimeout
	// StatusTimeout = status.New(ErrTimeOut503, "Timeout")

	// ErrTimeout
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeout = ecode.NewCodeError(ErrTimeOut503, "Timeout")
)

// -500
var (
// // ErrInvalid MsgResendReq Query
// // | -500 | Invalid msg_resend_req query | Invalid msg_resend_req query. |
// ErrInvalid MsgResendReq Query = ecode.NewCodeError(-500, "Invalid msg_resend_req query")
//
// // ErrInvalid MsgsAck Query
// // | -500 | Invalid msgs_ack query | Invalid msgs_ack query. |
// ErrInvalid MsgsAck Query = ecode.NewCodeError(-500, "Invalid msgs_ack query")
//
// // ErrInvalid MsgsStateReq Query
// // | -500 | Invalid msgs_state_req query | Invalid msgs_state_req query. |
// ErrInvalid MsgsStateReq Query = ecode.NewCodeError(-500, "Invalid msgs_state_req query")
)

// 700
var (
	// ErrPushRpcClient
	// db error
	// TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
	ErrPushRpcClient = ecode.NewCodeError(ErrNotReturnClient, "NOTRETURN_CLIENT")

	// ErrMigratedToChannel
	// MIGRATED_TO_CHANNEL
	ErrMigratedToChannel = ecode.NewCodeError(ErrNotReturnClient, "MIGRATED_TO_CHANNEL")
)

// StatusErrEqual is essentially a copy of testutils.StatusErrEqual(), to avoid a
// cyclic dependency.
/*
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return status1.Code() == status2.Code() && status1.Message() == status2.Message()
}
*/

// NewErrRedirectToX
// REDIRECT_TO_SERVER
// ErrRedirectToServer = status.Error(ErrNotReturnClient, "REDIRECT_TO_SERVER")
func NewErrRedirectToX(v string) error {
	return ecode.NewCodeErrorf(ErrNotReturnClient, "REDIRECT_TO_%s", v)
}
