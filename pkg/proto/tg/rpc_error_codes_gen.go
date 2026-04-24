// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/xerr"
)

// https://core.telegram.org/api/errors
/*
There will be errors when working with the API, and they must be correctly handled on the client.
An error is characterized by several parameters:

#### Error Code

Numerical value similar to HTTP status. Contains information on the type of error that occurred: for example,
a data input error, privacy error, or server error. This is a required parameter.

#### Error Type

A string literal in the form of `/[A-Z_0-9]+/`, which summarizes the problem. For example,
`AUTH_KEY_UNREGISTERED`. This is an optional parameter.

#### Error Database

A full human-readable JSON list of RPC errors that can be returned by all methods in the API
can be found [here »](<https://corefork.telegram.org/file/400780400111/3/RF-b0LDHWpc.202549.json/9b83afb26f1ba2f8aa>),
what follows is a description of its fields:

  * `errors` \- All error messages and codes for each method (object).
    * Keys: Error codes as strings (numeric strings)
    * Values: All error messages for each method (object)
      * Keys: Error messages (string)
      * Values: An array of methods which may emit this error (array of strings, may be empty for errors that can be emitted by any method)
  * `descriptions` \- Descriptions for every error mentioned in `errors` (and a few other errors not related to a specific method)
    * Keys: Error messages
    * Values: Error descriptions
  * `user_only` \- The full list of methods that can only be used by users, **not** bots.
  * `bot_only` \- The full list of methods that can only be used by bots, **not** users.
  * `business_supported` \- The full list of methods that can be used by bots over a
  	[business connection with invokeWithBusinessConnection](<bots_connected-business-bots.md>).
  * `unauthed_allowed` \- The full list of methods that can be used by not yet logged in connections.

Error messages and error descriptions may contain `printf` placeholders in key positions,
for now only `%d` is used to map durations contained in error messages to error descriptions.

Example:

```
    {
        "errors": {
            "420": {
                "2FA_CONFIRM_WAIT_%d": [
                    "account.deleteAccount"
                ],
                "SLOWMODE_WAIT_%d": [
                    "messages.forwardMessages",
                    "messages.sendInlineBotResult",
                    "messages.sendMedia",
                    "messages.sendMessage",
                    "messages.sendMultiMedia"
                ]
            }
        },
        "descriptions": {
            "2FA_CONFIRM_WAIT_%d": "Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in %d seconds.",
            "SLOWMODE_WAIT_%d": "Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.",
            "FLOOD_WAIT_%d": "Please wait %d seconds before repeating the action."
        },
        "user_only": [
            "account.deleteAccount"
        ],
        "bot_only": [
            "messages.setInlineBotResults"
        ],
        "business_supported": [
            "messages.sendMessage"
        ],
        "unauthed_allowed": [
            "auth.sendCode"
        ]
    }
```

* * *

#### Error Constructors

There should be a way to handle errors that are returned in [rpc_error](../mtproto/service_messages.md#rpc-error) constructors.

Below is a list of error codes and their meanings:
*/

const (
	/*
		### 303 SEE_OTHER

		The request must be repeated, but directed to a different data center.

		#### Examples of Errors:

		  * FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
		  * PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
		  * NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
		  * USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)

		In all these cases, the error description's string literal contains the number of the data center (instead of the X) to which the repeated query must be sent.
		[More information about redirects between data centers »](<datacenter.md>)
	*/
	ErrSeeOther = 303

	/*
		### 400 BAD_REQUEST

		The query contains errors. In the event that a request was created using a form and contains user generated data,
		the user should be notified that the data must be corrected before the query is repeated.

		#### Examples of Errors:

		  * FIRSTNAME_INVALID: The first name is invalid
		  * LASTNAME_INVALID: The last name is invalid
		  * PHONE_NUMBER_INVALID: The phone number is invalid
		  * PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
		  * PHONE_CODE_EMPTY: phone_code is missing
		  * PHONE_CODE_EXPIRED: The confirmation code has expired
		  * API_ID_INVALID: The api_id/api_hash combination is invalid
		  * PHONE_NUMBER_OCCUPIED: The phone number is already in use
		  * PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
		  * USERS_TOO_FEW: Not enough users (to create a chat, for example)
		  * USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
		  * TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
		  * FILE_PART_INVALID: The file part number is invalid
		  * FILE_PARTS_INVALID: The number of file parts is invalid
		  * FILE_PART_X_MISSING: Part X (where X is a number) of the file is missing from storage
		  * MD5_CHECKSUM_INVALID: The MD5 checksums do not match
		  * PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
		  * FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
		  * FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
	*/
	ErrBadRequest = 400

	/*
		### 401 UNAUTHORIZED

		There was an unauthorized attempt to use functionality available only to authorized users.

		#### Examples of Errors:

		  * AUTH_KEY_UNREGISTERED: The key is not registered in the system
		  * AUTH_KEY_INVALID: The key is invalid
		  * USER_DEACTIVATED: The user has been deleted/deactivated
		  * SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
		  * SESSION_EXPIRED: The authorization has expired
		  * AUTH_KEY_PERM_EMPTY: The method is unavailable for temporary authorization key, not bound to permanent
	*/
	ErrUnauthorized = 401

	/*
		### 403 FORBIDDEN

		Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	*/
	ErrForbidden = 403

	/*
		### 404 NOT_FOUND

		An attempt to invoke a non-existent object, such as a method.
	*/
	ErrNotFound = 404

	/*
		### 406 NOT_ACCEPTABLE

		Similar to 400 BAD_REQUEST, but the app must display the error to the user a bit differently.
		Do not display any visible error to the user when receiving the `rpc_error` constructor: instead,
		wait for an [updateServiceNotification](../constructor/updateServiceNotification.md) update, and handle it as usual.
		Basically, an [updateServiceNotification](../constructor/updateServiceNotification.md)
		`popup` update will be emitted independently (ie NOT as an [Updates](../type/Updates.md)
		constructor inside `rpc_result` but as a normal update) immediately after emission of a 406 `rpc_error`:
		the update will contain the actual localized error message to show to the user with a UI popup.

		An exception to this is the `AUTH_KEY_DUPLICATED` error,
		which is only emitted if any of the non-media DC detects that an authorized session is sending requests in parallel from two separate TCP connections,
		from the same or different IP addresses.
		Note that parallel connections are still allowed and actually recommended for media DCs.
		Also note that by session we mean a logged-in session identified by an [authorization](../constructor/authorization.md) constructor,
		fetchable using [account.getAuthorizations](../methods/method/account.getAuthorizations.md), not an MTProto session.

		If the client receives an `AUTH_KEY_DUPLICATED` error, the session was already invalidated by the server and the user must generate a new auth key and login again.
	*/
	ErrNotAcceptable = 406

	/*
		### 420 FLOOD

		The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded. For example,
		in an attempt to request a large number of text messages (SMS) for the same phone number.

		#### Error Example:

		  * FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
		  * FLOOD_PREMIUM_WAIT_X: A wait of X seconds is required (where X is a number);
		  the user may also purchase a [Telegram Premium subscription](<premium.md>) to remove this limitation.
		  See [here »](<files.md>) for more info on how to handle this error.
	*/
	ErrFlood = 420

	/*
		### 500 INTERNAL

		An internal server error occurred while a request was being processed; for example,
		there was a disruption while accessing a database or file storage.

		If a client receives a 500 error, or you believe this error should not have occurred,
		please collect as much information as possible about the query and error and send it to the developers.
	*/
	ErrInternal = 500

	/*
		### Other Error Codes

		If a server returns an error with a code other than the ones listed above,
		it may be considered the same as a 500 error and treated as an internal server error.
	*/

	// ErrTimeOut503
	// TODO(@benqi): ???, if ecode < 0, panic
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeOut503 = 5030000

	// ErrNotReturnClient ErrNotReturnClient
	ErrNotReturnClient = 700

	// ErrDatabase db error
	ErrDatabase = 600

	// ErrRedis redis error
	ErrRedis = 601

	// ErrThirdParty error
	ErrThirdParty = 602
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
	return xerr.NewCodeErrorf(ErrSeeOther, "FILE_MIGRATE_%d", dc)
}

// NewErrPhoneMigrateX
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X. |
func NewErrPhoneMigrateX(dc int32) error {
	return xerr.NewCodeErrorf(ErrSeeOther, "PHONE_MIGRATE_%d", dc)
}

// NewErrNetworkMigrateX
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X. |
// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
func NewErrNetworkMigrateX(dc int32) error {
	return xerr.NewCodeErrorf(ErrSeeOther, "NETWORK_MIGRATE_%d", dc)
}

// NewErrUserMigrateX
// USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)
func NewErrUserMigrateX(dc int32) error {
	return xerr.NewCodeErrorf(ErrSeeOther, "USER_MIGRATE_%d", dc)
}

// NewErrFloodWaitX 420 FLOOD
//
// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
func NewErrFloodWaitX(second int32) error {
	return xerr.NewCodeErrorf(ErrFlood, "FLOOD_WAIT_%d", second)
}

// NewErrSlowModeWaitX
// | 420 | SLOWMODE_WAIT_X | Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat. |
func NewErrSlowModeWaitX(second int32) error {
	return xerr.NewCodeErrorf(ErrFlood, "SLOWMODE_WAIT_%d", second)
}

// NewErrTakeoutInitDelayX
// | 420 | TAKEOUT_INIT_DELAY_X | Wait X seconds before initing takeout. |
func NewErrTakeoutInitDelayX(second int32) error {
	return xerr.NewCodeErrorf(ErrFlood, "TAKEOUT_INIT_DELAY_%d", second)
}

// NewErr2faConfirmWaitX
// | 420 | 2FA_CONFIRM_WAIT_X | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in X seconds. |
func NewErr2faConfirmWaitX(second int32) error {
	return xerr.NewCodeErrorf(ErrFlood, "2FA_CONFIRM_WAIT_%d", second)
}

// 420 ErrFlood  = 420
var (
	// ErrP0nyFloodwait
	// | 420 | P0NY_FLOODWAIT |   |
	ErrP0nyFloodwait = xerr.NewCodeError(ErrFlood, "P0NY_FLOODWAIT")

	//// ErrSlowmodeWait_%d
	//// | 420 | SLOWMODE_WAIT_%d | Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat. |
	//ErrSlowmodeWait_%d = xerr.NewCodeError(420, "SLOWMODE_WAIT_%d")
	//
	//// ErrTakeoutInitDelay_%d
	//// | 420 | TAKEOUT_INIT_DELAY_%d | Wait %d seconds before initializing takeout. |
	//ErrTakeoutInitDelay_%d = xerr.NewCodeError(420, "TAKEOUT_INIT_DELAY_%d")
	//
	//// Err2faConfirmWait_%d
	//// | 420 | 2FA_CONFIRM_WAIT_%d | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in %d seconds. |
	//Err2faConfirmWait_%d = xerr.NewCodeError(420, "2FA_CONFIRM_WAIT_%d")
)

// 400 BAD_REQUEST

// NewErrFileReferenceX
// | 400 | FILE_REFERENCE_* | The file reference expired, it must be refreshed. |
func NewErrFileReferenceX(second int32) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "FILE_REFERENCE_%d", second)
}

// NewErrPasswordTooFreshX
// PASSWORD_TOO_FRESH_%d
// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
func NewErrPasswordTooFreshX(second int32) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "PASSWORD_TOO_FRESH_%d", second)
}

// NewErrEmailUnconfirmedX
// ErrEmailUnconfirmed_%d
// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
func NewErrEmailUnconfirmedX(len int32) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", len)
}

// NewErrSessionTooFreshX
// ErrSessionTooFresh_%d
// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
func NewErrSessionTooFreshX(second int32) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "SESSION_TOO_FRESH_%d", second)
}

// NewFilePartXMissing
// FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage. Try repeating the method call to resave the part.
func NewFilePartXMissing(x int32) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "FILE_PART_%d_MISSING", x)
}

// NewEmailUnconfirmedX
// 400	EMAIL_UNCONFIRMED_X	The provided email isn't confirmed, X is the length of the verification code that was just sent to the email: use account.verifyEmail to enter the received verification code and enable the recovery email.
func NewEmailUnconfirmedX(x int) error {
	return xerr.NewCodeErrorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", x)
}

var (
	// ErrMethodNotImpl
	// @benqi Add By NebulaChat, not impl the method
	// METHOD_NOT_IMPL: The method not impl
	ErrMethodNotImpl = xerr.NewCodeError(ErrBadRequest, "METHOD_NOT_IMPL")

	// ErrGroupCallParticipantInvalid GROUPCALL_PARTICIPANT_INVALID
	ErrGroupCallParticipantInvalid = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_PARTICIPANT_INVALID")

	// ErrCheckSumInvalid
	// MD5_CHECKSUM_INVALID: The file’s checksum did not match the md5_checksum parameter
	ErrCheckSumInvalid = xerr.NewCodeError(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrGroupCallInvalid GROUPCALL_INVALID
	ErrGroupCallInvalid = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrThemeSlugOccupied
	// THEME_SLUG_OCCUPIED
	ErrThemeSlugOccupied = xerr.NewCodeError(ErrBadRequest, "THEME_SLUG_OCCUPIED")

	// ErrShortnameOccupyFailed
	// | 400 | SHORTNAME_OCCUPY_FAILED | An internal error occurred. |
	ErrShortnameOccupyFailed = xerr.NewCodeError(ErrBadRequest, "SHORTNAME_OCCUPY_FAILED")

	// ErrThemeSlugInvalid
	// THEME_SLUG_INVALID 400 The input theme slug was not valid
	// ErrThemeSlugInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_SLUG_INVALID")

	// ErrApiServerNeeded
	// | 400 | API_SERVER_NEEDED | This method be used by api server |
	ErrApiServerNeeded = xerr.NewCodeError(ErrBadRequest, "API_SERVER_NEEDED")

	// ErrInputConstructorInvalid
	// 400	INPUT_CONSTRUCTOR_INVALID	The provided constructor is invalid
	// ErrInputConstructorInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// ErrEncryptedChatIdInvalid
	// | 400 | ENCRYPTED_CHAT_ID_INVALID | The encrypted chat id is invalid |
	ErrEncryptedChatIdInvalid = xerr.NewCodeError(ErrBadRequest, "ENCRYPTED_CHAT_ID_INVALID")

	// ErrAuthTokenAccepted
	// 400 - AUTH_TOKEN_ALREADY_ACCEPTED, the authorization token was already used
	ErrAuthTokenAccepted = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrBotMethodInvalid
	// | 400 | BOT_METHOD_INVALID | This method can't be used by a bot |
	// ErrBotMethodInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_METHOD_INVALID")

	// ErrInputRequestInvalid
	// INPUT_REQUEST_INVALID: The method not impl
	ErrInputRequestInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_REQUEST_INVALID")

	// ErrEnterpriseIsBlocked ErrEnterpriseIsBlocked
	ErrEnterpriseIsBlocked = xerr.NewCodeError(ErrBadRequest, "ERR_ENTERPRISE_IS_BLOCKED")

	// ErrErrBadRequest ErrErrBadRequest
	ErrErrBadRequest = xerr.NewCodeError(ErrBadRequest, "ERR_BAD_REQUEST")

	// ErrTimeTooBig
	// 400	TIME_TOO_BIG
	ErrTimeTooBig = xerr.NewCodeError(ErrBadRequest, "TIME_TOO_BIG")

	// Err400AddressInvalid
	// | 400 | ADDRESS_INVALID | The specified geopoint address is invalid. |
	Err400AddressInvalid = xerr.NewCodeError(ErrBadRequest, "ADDRESS_INVALID")

	// Err400BannedRightsInvalid
	// | 400 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	Err400BannedRightsInvalid = xerr.NewCodeError(ErrBadRequest, "BANNED_RIGHTS_INVALID")

	// Err400CallOccupyFailed
	// | 400 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	Err400CallOccupyFailed = xerr.NewCodeError(ErrBadRequest, "CALL_OCCUPY_FAILED")

	// Err400ChannelPrivate
	// | 400 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	Err400ChannelPrivate = xerr.NewCodeError(ErrBadRequest, "CHANNEL_PRIVATE")

	// Err400ChannelTooLarge
	// | 400 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	Err400ChannelTooLarge = xerr.NewCodeError(ErrBadRequest, "CHANNEL_TOO_LARGE")

	// Err400ChatAdminRequired
	// | 400 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	Err400ChatAdminRequired = xerr.NewCodeError(ErrBadRequest, "CHAT_ADMIN_REQUIRED")

	// Err400ChatForwardsRestricted
	// | 400 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	Err400ChatForwardsRestricted = xerr.NewCodeError(ErrBadRequest, "CHAT_FORWARDS_RESTRICTED")

	// Err400ChatInvalid
	// | 400 | CHAT_INVALID | Invalid chat. |
	Err400ChatInvalid = xerr.NewCodeError(ErrBadRequest, "CHAT_INVALID")

	// Err400ChatSendInlineForbidden
	// | 400 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	Err400ChatSendInlineForbidden = xerr.NewCodeError(ErrBadRequest, "CHAT_SEND_INLINE_FORBIDDEN")

	// Err400FreshChangeAdminsForbidden
	// | 400 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	Err400FreshChangeAdminsForbidden = xerr.NewCodeError(ErrBadRequest, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// Err400GroupcallForbidden
	// | 400 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	Err400GroupcallForbidden = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_FORBIDDEN")

	// Err400InviteHashExpired
	// | 400 | INVITE_HASH_EXPIRED | The invite link has expired. |
	Err400InviteHashExpired = xerr.NewCodeError(ErrBadRequest, "INVITE_HASH_EXPIRED")

	// Err400MethodInvalid
	// | 400 | METHOD_INVALID | The specified method is invalid. |
	Err400MethodInvalid = xerr.NewCodeError(ErrBadRequest, "METHOD_INVALID")

	// Err400MsgWaitFailed
	// | 400 | MSG_WAIT_FAILED | A waiting call returned an error. |
	Err400MsgWaitFailed = xerr.NewCodeError(ErrBadRequest, "MSG_WAIT_FAILED")

	// Err400NotEligible
	// | 400 | NOT_ELIGIBLE | The current user is not eligible to join the Peer-to-Peer Login Program. |
	Err400NotEligible = xerr.NewCodeError(ErrBadRequest, "NOT_ELIGIBLE")

	// Err400ParticipantJoinMissing
	// | 400 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	Err400ParticipantJoinMissing = xerr.NewCodeError(ErrBadRequest, "PARTICIPANT_JOIN_MISSING")

	// Err400PeerIdInvalid
	// | 400 | PEER_ID_INVALID | The provided peer id is invalid. |
	Err400PeerIdInvalid = xerr.NewCodeError(ErrBadRequest, "PEER_ID_INVALID")

	// Err400PhoneNumberInvalid
	// | 400 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	Err400PhoneNumberInvalid = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_INVALID")

	// Err400PremiumAccountRequired
	// | 400 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	Err400PremiumAccountRequired = xerr.NewCodeError(ErrBadRequest, "PREMIUM_ACCOUNT_REQUIRED")

	// Err400StickersetInvalid
	// | 400 | STICKERSET_INVALID | The provided sticker set is invalid. |
	Err400StickersetInvalid = xerr.NewCodeError(ErrBadRequest, "STICKERSET_INVALID")

	// Err400TakeoutRequired
	// | 400 | TAKEOUT_REQUIRED | A [takeout](https://core.telegram.org/api/takeout) session needs to be initialized first, [see here &raquo; for more info](https://core.telegram.org/api/takeout). |
	Err400TakeoutRequired = xerr.NewCodeError(ErrBadRequest, "TAKEOUT_REQUIRED")

	// Err400TopicClosed
	// | 400 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	Err400TopicClosed = xerr.NewCodeError(ErrBadRequest, "TOPIC_CLOSED")

	// Err400TopicDeleted
	// | 400 | TOPIC_DELETED | The specified topic was deleted. |
	Err400TopicDeleted = xerr.NewCodeError(ErrBadRequest, "TOPIC_DELETED")

	// Err400UserChannelsTooMuch
	// | 400 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	Err400UserChannelsTooMuch = xerr.NewCodeError(ErrBadRequest, "USER_CHANNELS_TOO_MUCH")

	// Err400UserInvalid
	// | 400 | USER_INVALID | Invalid user provided. |
	Err400UserInvalid = xerr.NewCodeError(ErrBadRequest, "USER_INVALID")

	// Err400UserIsBlocked
	// | 400 | USER_IS_BLOCKED | You were blocked by this user. |
	Err400UserIsBlocked = xerr.NewCodeError(ErrBadRequest, "USER_IS_BLOCKED")

	// Err400UserNotMutualContact
	// | 400 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	Err400UserNotMutualContact = xerr.NewCodeError(ErrBadRequest, "USER_NOT_MUTUAL_CONTACT")

	// Err400UserNotParticipant
	// | 400 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	Err400UserNotParticipant = xerr.NewCodeError(ErrBadRequest, "USER_NOT_PARTICIPANT")

	// Err400UserpicUploadRequired
	// | 400 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	Err400UserpicUploadRequired = xerr.NewCodeError(ErrBadRequest, "USERPIC_UPLOAD_REQUIRED")

	// Err400VoiceMessagesForbidden
	// | 400 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	Err400VoiceMessagesForbidden = xerr.NewCodeError(ErrBadRequest, "VOICE_MESSAGES_FORBIDDEN")

	// ErrAboutTooLong
	// | 400 | ABOUT_TOO_LONG | About string too long. |
	ErrAboutTooLong = xerr.NewCodeError(ErrBadRequest, "ABOUT_TOO_LONG")

	// ErrAccessTokenExpired
	// | 400 | ACCESS_TOKEN_EXPIRED | Access token expired. |
	ErrAccessTokenExpired = xerr.NewCodeError(ErrBadRequest, "ACCESS_TOKEN_EXPIRED")

	// ErrAccessTokenInvalid
	// | 400 | ACCESS_TOKEN_INVALID | Access token invalid. |
	ErrAccessTokenInvalid = xerr.NewCodeError(ErrBadRequest, "ACCESS_TOKEN_INVALID")

	// ErrAdExpired
	// | 400 | AD_EXPIRED | The ad has expired (too old or not found). |
	ErrAdExpired = xerr.NewCodeError(ErrBadRequest, "AD_EXPIRED")

	// ErrAdminIdInvalid
	// | 400 | ADMIN_ID_INVALID | The specified admin ID is invalid. |
	ErrAdminIdInvalid = xerr.NewCodeError(ErrBadRequest, "ADMIN_ID_INVALID")

	// ErrAdminRankEmojiNotAllowed
	// | 400 | ADMIN_RANK_EMOJI_NOT_ALLOWED | An admin rank cannot contain emojis. |
	ErrAdminRankEmojiNotAllowed = xerr.NewCodeError(ErrBadRequest, "ADMIN_RANK_EMOJI_NOT_ALLOWED")

	// ErrAdminRankInvalid
	// | 400 | ADMIN_RANK_INVALID | The specified admin rank is invalid. |
	ErrAdminRankInvalid = xerr.NewCodeError(ErrBadRequest, "ADMIN_RANK_INVALID")

	// ErrAdminRightsEmpty
	// | 400 | ADMIN_RIGHTS_EMPTY | The chatAdminRights constructor passed in keyboardButtonRequestPeer.peer_type.user_admin_rights has no rights set (i.e. flags is 0). |
	ErrAdminRightsEmpty = xerr.NewCodeError(ErrBadRequest, "ADMIN_RIGHTS_EMPTY")

	// ErrAdminsTooMuch
	// | 400 | ADMINS_TOO_MUCH | There are too many admins. |
	ErrAdminsTooMuch = xerr.NewCodeError(ErrBadRequest, "ADMINS_TOO_MUCH")

	// ErrAlbumPhotosTooMany
	// | 400 | ALBUM_PHOTOS_TOO_MANY | You have uploaded too many profile photos, delete some before retrying. |
	ErrAlbumPhotosTooMany = xerr.NewCodeError(ErrBadRequest, "ALBUM_PHOTOS_TOO_MANY")

	// ErrApiIdInvalid
	// | 400 | API_ID_INVALID | API ID invalid. |
	ErrApiIdInvalid = xerr.NewCodeError(ErrBadRequest, "API_ID_INVALID")

	// ErrApiIdPublishedFlood
	// | 400 | API_ID_PUBLISHED_FLOOD | This API id was published somewhere, you can't use it now. |
	ErrApiIdPublishedFlood = xerr.NewCodeError(ErrBadRequest, "API_ID_PUBLISHED_FLOOD")

	// ErrArticleTitleEmpty
	// | 400 | ARTICLE_TITLE_EMPTY | The title of the article is empty. |
	ErrArticleTitleEmpty = xerr.NewCodeError(ErrBadRequest, "ARTICLE_TITLE_EMPTY")

	// ErrAudioContentUrlEmpty
	// | 400 | AUDIO_CONTENT_URL_EMPTY | The remote URL specified in the content field is empty. |
	ErrAudioContentUrlEmpty = xerr.NewCodeError(ErrBadRequest, "AUDIO_CONTENT_URL_EMPTY")

	// ErrAudioTitleEmpty
	// | 400 | AUDIO_TITLE_EMPTY | An empty audio title was provided. |
	ErrAudioTitleEmpty = xerr.NewCodeError(ErrBadRequest, "AUDIO_TITLE_EMPTY")

	// ErrAuthBytesInvalid
	// | 400 | AUTH_BYTES_INVALID | The provided authorization is invalid. |
	ErrAuthBytesInvalid = xerr.NewCodeError(ErrBadRequest, "AUTH_BYTES_INVALID")

	// ErrAuthTokenAlreadyAccepted
	// | 400 | AUTH_TOKEN_ALREADY_ACCEPTED | The specified auth token was already accepted. |
	ErrAuthTokenAlreadyAccepted = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrAuthTokenException
	// | 400 | AUTH_TOKEN_EXCEPTION | An error occurred while importing the auth token. |
	ErrAuthTokenException = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_EXCEPTION")

	// ErrAuthTokenExpired
	// | 400 | AUTH_TOKEN_EXPIRED | The authorization token has expired. |
	ErrAuthTokenExpired = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_EXPIRED")

	// ErrAuthTokenInvalid
	// | 400 | AUTH_TOKEN_INVALID | The specified auth token is invalid. |
	ErrAuthTokenInvalid = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_INVALID")

	// ErrAuthTokenInvalidx
	// | 400 | AUTH_TOKEN_INVALIDX | The specified auth token is invalid. |
	ErrAuthTokenInvalidx = xerr.NewCodeError(ErrBadRequest, "AUTH_TOKEN_INVALIDX")

	// ErrAutoarchiveNotAvailable
	// | 400 | AUTOARCHIVE_NOT_AVAILABLE | The autoarchive setting is not available at this time: please check the value of the [autoarchive_setting_available field in client config &raquo;](https://core.telegram.org/api/config#client-configuration) before calling this method. |
	ErrAutoarchiveNotAvailable = xerr.NewCodeError(ErrBadRequest, "AUTOARCHIVE_NOT_AVAILABLE")

	// ErrBalanceTooLow
	// | 400 | BALANCE_TOO_LOW | The transaction cannot be completed because the current [Telegram Stars balance](https://core.telegram.org/api/stars) is too low. |
	ErrBalanceTooLow = xerr.NewCodeError(ErrBadRequest, "BALANCE_TOO_LOW")

	// ErrBankCardNumberInvalid
	// | 400 | BANK_CARD_NUMBER_INVALID | The specified card number is invalid. |
	ErrBankCardNumberInvalid = xerr.NewCodeError(ErrBadRequest, "BANK_CARD_NUMBER_INVALID")

	// ErrBirthdayInvalid
	// | 400 | BIRTHDAY_INVALID | An invalid age was specified, must be between 0 and 150 years. |
	ErrBirthdayInvalid = xerr.NewCodeError(ErrBadRequest, "BIRTHDAY_INVALID")

	// ErrBoostNotModified
	// | 400 | BOOST_NOT_MODIFIED | You're already [boosting](https://core.telegram.org/api/boost) the specified channel. |
	ErrBoostNotModified = xerr.NewCodeError(ErrBadRequest, "BOOST_NOT_MODIFIED")

	// ErrBoostPeerInvalid
	// | 400 | BOOST_PEER_INVALID | The specified `boost_peer` is invalid. |
	ErrBoostPeerInvalid = xerr.NewCodeError(ErrBadRequest, "BOOST_PEER_INVALID")

	// ErrBoostsEmpty
	// | 400 | BOOSTS_EMPTY | No boost slots were specified. |
	ErrBoostsEmpty = xerr.NewCodeError(ErrBadRequest, "BOOSTS_EMPTY")

	// ErrBoostsRequired
	// | 400 | BOOSTS_REQUIRED | The specified channel must first be [boosted by its users](https://core.telegram.org/api/boost) in order to perform this action. |
	ErrBoostsRequired = xerr.NewCodeError(ErrBadRequest, "BOOSTS_REQUIRED")

	// ErrBotAlreadyDisabled
	// | 400 | BOT_ALREADY_DISABLED | The connected business bot was already disabled for the specified peer. |
	ErrBotAlreadyDisabled = xerr.NewCodeError(ErrBadRequest, "BOT_ALREADY_DISABLED")

	// ErrBotAppBotInvalid
	// | 400 | BOT_APP_BOT_INVALID | The bot_id passed in the inputBotAppShortName constructor is invalid. |
	ErrBotAppBotInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_APP_BOT_INVALID")

	// ErrBotAppInvalid
	// | 400 | BOT_APP_INVALID | The specified bot app is invalid. |
	ErrBotAppInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_APP_INVALID")

	// ErrBotAppShortnameInvalid
	// | 400 | BOT_APP_SHORTNAME_INVALID | The specified bot app short name is invalid. |
	ErrBotAppShortnameInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_APP_SHORTNAME_INVALID")

	// ErrBotBusinessMissing
	// | 400 | BOT_BUSINESS_MISSING | The specified bot is not a business bot (the [user](https://core.telegram.org/constructor/user).`bot_business` flag is not set). |
	ErrBotBusinessMissing = xerr.NewCodeError(ErrBadRequest, "BOT_BUSINESS_MISSING")

	// ErrBotChannelsNa
	// | 400 | BOT_CHANNELS_NA | Bots can't edit admin privileges. |
	ErrBotChannelsNa = xerr.NewCodeError(ErrBadRequest, "BOT_CHANNELS_NA")

	// ErrBotCommandDescriptionInvalid
	// | 400 | BOT_COMMAND_DESCRIPTION_INVALID | The specified command description is invalid. |
	ErrBotCommandDescriptionInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_COMMAND_DESCRIPTION_INVALID")

	// ErrBotCommandInvalid
	// | 400 | BOT_COMMAND_INVALID | The specified command is invalid. |
	ErrBotCommandInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_COMMAND_INVALID")

	// ErrBotDomainInvalid
	// | 400 | BOT_DOMAIN_INVALID | Bot domain invalid. |
	ErrBotDomainInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_DOMAIN_INVALID")

	// ErrBotFallbackUnsupported
	// | 400 | BOT_FALLBACK_UNSUPPORTED | The fallback flag can't be set for bots. |
	ErrBotFallbackUnsupported = xerr.NewCodeError(ErrBadRequest, "BOT_FALLBACK_UNSUPPORTED")

	// ErrBotGamesDisabled
	// | 400 | BOT_GAMES_DISABLED | Games can't be sent to channels. |
	ErrBotGamesDisabled = xerr.NewCodeError(ErrBadRequest, "BOT_GAMES_DISABLED")

	// ErrBotGroupsBlocked
	// | 400 | BOT_GROUPS_BLOCKED | This bot can't be added to groups. |
	ErrBotGroupsBlocked = xerr.NewCodeError(ErrBadRequest, "BOT_GROUPS_BLOCKED")

	// ErrBotInlineDisabled
	// | 400 | BOT_INLINE_DISABLED | This bot can't be used in inline mode. |
	ErrBotInlineDisabled = xerr.NewCodeError(ErrBadRequest, "BOT_INLINE_DISABLED")

	// ErrBotInvalid
	// | 400 | BOT_INVALID | This is not a valid bot. |
	ErrBotInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_INVALID")

	// ErrBotInvoiceInvalid
	// | 400 | BOT_INVOICE_INVALID | The specified invoice is invalid. |
	ErrBotInvoiceInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_INVOICE_INVALID")

	// ErrBotMethodInvalid
	// | 400 | BOT_METHOD_INVALID | The specified method cannot be used by bots. |
	ErrBotMethodInvalid = xerr.NewCodeError(ErrBadRequest, "BOT_METHOD_INVALID")

	// ErrBotNotConnectedYet
	// | 400 | BOT_NOT_CONNECTED_YET | No [business bot](https://core.telegram.org/api/business#connected-bots) is connected to the currently logged in user. |
	ErrBotNotConnectedYet = xerr.NewCodeError(ErrBadRequest, "BOT_NOT_CONNECTED_YET")

	// ErrBotOnesideNotAvail
	// | 400 | BOT_ONESIDE_NOT_AVAIL | Bots can't pin messages in PM just for themselves. |
	ErrBotOnesideNotAvail = xerr.NewCodeError(ErrBadRequest, "BOT_ONESIDE_NOT_AVAIL")

	// ErrBotPaymentsDisabled
	// | 400 | BOT_PAYMENTS_DISABLED | Please enable bot payments in botfather before calling this method. |
	ErrBotPaymentsDisabled = xerr.NewCodeError(ErrBadRequest, "BOT_PAYMENTS_DISABLED")

	// ErrBotResponseTimeout
	// | 400 | BOT_RESPONSE_TIMEOUT | A timeout occurred while fetching data from the bot. |
	ErrBotResponseTimeout = xerr.NewCodeError(ErrBadRequest, "BOT_RESPONSE_TIMEOUT")

	// ErrBotScoreNotModified
	// | 400 | BOT_SCORE_NOT_MODIFIED | The score wasn't modified. |
	ErrBotScoreNotModified = xerr.NewCodeError(ErrBadRequest, "BOT_SCORE_NOT_MODIFIED")

	// ErrBotWebviewDisabled
	// | 400 | BOT_WEBVIEW_DISABLED | A webview cannot be opened in the specified conditions: emitted for example if `from_bot_menu` or `url` are set and `peer` is not the chat with the bot. |
	ErrBotWebviewDisabled = xerr.NewCodeError(ErrBadRequest, "BOT_WEBVIEW_DISABLED")

	// ErrBotsTooMuch
	// | 400 | BOTS_TOO_MUCH | There are too many bots in this chat/channel. |
	ErrBotsTooMuch = xerr.NewCodeError(ErrBadRequest, "BOTS_TOO_MUCH")

	// ErrBroadcastIdInvalid
	// | 400 | BROADCAST_ID_INVALID | Broadcast ID invalid. |
	ErrBroadcastIdInvalid = xerr.NewCodeError(ErrBadRequest, "BROADCAST_ID_INVALID")

	// ErrBroadcastPublicVotersForbidden
	// | 400 | BROADCAST_PUBLIC_VOTERS_FORBIDDEN | You can't forward polls with public voters. |
	ErrBroadcastPublicVotersForbidden = xerr.NewCodeError(ErrBadRequest, "BROADCAST_PUBLIC_VOTERS_FORBIDDEN")

	// ErrBroadcastRequired
	// | 400 | BROADCAST_REQUIRED | This method can only be called on a channel, please use stats.getMegagroupStats for supergroups. |
	ErrBroadcastRequired = xerr.NewCodeError(ErrBadRequest, "BROADCAST_REQUIRED")

	// ErrBusinessConnectionInvalid
	// | 400 | BUSINESS_CONNECTION_INVALID | The `connection_id` passed to the wrapping [invokeWithBusinessConnection](https://core.telegram.org/api/business) call is invalid. |
	ErrBusinessConnectionInvalid = xerr.NewCodeError(ErrBadRequest, "BUSINESS_CONNECTION_INVALID")

	// ErrBusinessConnectionNotAllowed
	// | 400 | BUSINESS_CONNECTION_NOT_ALLOWED | This method was invoked over a business connection using [invokeWithBusinessConnection](https://core.telegram.org/api/business#connected-bots), but either (1) we're a user, and users cannot invoke methods over a business connection; (2) we're a bot, but business mode was disabled in @botfather or (3); we're a bot, but this method cannot be invoked over a business connection. |
	ErrBusinessConnectionNotAllowed = xerr.NewCodeError(ErrBadRequest, "BUSINESS_CONNECTION_NOT_ALLOWED")

	// ErrBusinessPeerInvalid
	// | 400 | BUSINESS_PEER_INVALID | Messages can't be set to the specified peer through the current [business connection](https://core.telegram.org/api/business#connected-bots). |
	ErrBusinessPeerInvalid = xerr.NewCodeError(ErrBadRequest, "BUSINESS_PEER_INVALID")

	// ErrBusinessPeerUsageMissing
	// | 400 | BUSINESS_PEER_USAGE_MISSING | You cannot send a message to a user through a [business connection](https://core.telegram.org/api/business#connected-bots) if the user hasn't recently contacted us. |
	ErrBusinessPeerUsageMissing = xerr.NewCodeError(ErrBadRequest, "BUSINESS_PEER_USAGE_MISSING")

	// ErrBusinessRecipientsEmpty
	// | 400 | BUSINESS_RECIPIENTS_EMPTY | You didn't set any flag in inputBusinessBotRecipients, thus the bot cannot work with *any* peer. |
	ErrBusinessRecipientsEmpty = xerr.NewCodeError(ErrBadRequest, "BUSINESS_RECIPIENTS_EMPTY")

	// ErrBusinessWorkHoursEmpty
	// | 400 | BUSINESS_WORK_HOURS_EMPTY | No work hours were specified. |
	ErrBusinessWorkHoursEmpty = xerr.NewCodeError(ErrBadRequest, "BUSINESS_WORK_HOURS_EMPTY")

	// ErrBusinessWorkHoursPeriodInvalid
	// | 400 | BUSINESS_WORK_HOURS_PERIOD_INVALID | The specified work hours are invalid, see [here &raquo;](https://core.telegram.org/api/business#opening-hours) for the exact requirements. |
	ErrBusinessWorkHoursPeriodInvalid = xerr.NewCodeError(ErrBadRequest, "BUSINESS_WORK_HOURS_PERIOD_INVALID")

	// ErrButtonCopyTextInvalid
	// | 400 | BUTTON_COPY_TEXT_INVALID | The specified [keyboardButtonCopy](https://core.telegram.org/constructor/keyboardButtonCopy).`copy_text` is invalid. |
	ErrButtonCopyTextInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_COPY_TEXT_INVALID")

	// ErrButtonDataInvalid
	// | 400 | BUTTON_DATA_INVALID | The data of one or more of the buttons you provided is invalid. |
	ErrButtonDataInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_DATA_INVALID")

	// ErrButtonIdInvalid
	// | 400 | BUTTON_ID_INVALID | The specified button ID is invalid. |
	ErrButtonIdInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_ID_INVALID")

	// ErrButtonInvalid
	// | 400 | BUTTON_INVALID | The specified button is invalid. |
	ErrButtonInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_INVALID")

	// ErrButtonPosInvalid
	// | 400 | BUTTON_POS_INVALID | The position of one of the keyboard buttons is invalid (i.e. a Game or Pay button not in the first position, and so on...). |
	ErrButtonPosInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_POS_INVALID")

	// ErrButtonTextInvalid
	// | 400 | BUTTON_TEXT_INVALID | The specified button text is invalid. |
	ErrButtonTextInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_TEXT_INVALID")

	// ErrButtonTypeInvalid
	// | 400 | BUTTON_TYPE_INVALID | The type of one or more of the buttons you provided is invalid. |
	ErrButtonTypeInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_TYPE_INVALID")

	// ErrButtonUrlInvalid
	// | 400 | BUTTON_URL_INVALID | Button URL invalid. |
	ErrButtonUrlInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_URL_INVALID")

	// ErrButtonUserInvalid
	// | 400 | BUTTON_USER_INVALID | The `user_id` passed to inputKeyboardButtonUserProfile is invalid! |
	ErrButtonUserInvalid = xerr.NewCodeError(ErrBadRequest, "BUTTON_USER_INVALID")

	// ErrButtonUserPrivacyRestricted
	// | 400 | BUTTON_USER_PRIVACY_RESTRICTED | The privacy setting of the user specified in a [inputKeyboardButtonUserProfile](https://core.telegram.org/constructor/inputKeyboardButtonUserProfile) button do not allow creating such a button. |
	ErrButtonUserPrivacyRestricted = xerr.NewCodeError(ErrBadRequest, "BUTTON_USER_PRIVACY_RESTRICTED")

	// ErrCallAlreadyAccepted
	// | 400 | CALL_ALREADY_ACCEPTED | The call was already accepted. |
	ErrCallAlreadyAccepted = xerr.NewCodeError(ErrBadRequest, "CALL_ALREADY_ACCEPTED")

	// ErrCallAlreadyDeclined
	// | 400 | CALL_ALREADY_DECLINED | The call was already declined. |
	ErrCallAlreadyDeclined = xerr.NewCodeError(ErrBadRequest, "CALL_ALREADY_DECLINED")

	// ErrCallPeerInvalid
	// | 400 | CALL_PEER_INVALID | The provided call peer object is invalid. |
	ErrCallPeerInvalid = xerr.NewCodeError(ErrBadRequest, "CALL_PEER_INVALID")

	// ErrCallProtocolFlagsInvalid
	// | 400 | CALL_PROTOCOL_FLAGS_INVALID | Call protocol flags invalid. |
	ErrCallProtocolFlagsInvalid = xerr.NewCodeError(ErrBadRequest, "CALL_PROTOCOL_FLAGS_INVALID")

	// ErrCallProtocolLayerInvalid
	// | 400 | CALL_PROTOCOL_LAYER_INVALID | The specified protocol layer version range is invalid. |
	ErrCallProtocolLayerInvalid = xerr.NewCodeError(ErrBadRequest, "CALL_PROTOCOL_LAYER_INVALID")

	// ErrCdnMethodInvalid
	// | 400 | CDN_METHOD_INVALID | You can't call this method in a CDN DC. |
	ErrCdnMethodInvalid = xerr.NewCodeError(ErrBadRequest, "CDN_METHOD_INVALID")

	// ErrChannelForumMissing
	// | 400 | CHANNEL_FORUM_MISSING | This supergroup is not a forum. |
	ErrChannelForumMissing = xerr.NewCodeError(ErrBadRequest, "CHANNEL_FORUM_MISSING")

	// ErrChannelIdInvalid
	// | 400 | CHANNEL_ID_INVALID | The specified supergroup ID is invalid. |
	ErrChannelIdInvalid = xerr.NewCodeError(ErrBadRequest, "CHANNEL_ID_INVALID")

	// ErrChannelInvalid
	// | 400 | CHANNEL_INVALID | The provided channel is invalid. |
	ErrChannelInvalid = xerr.NewCodeError(ErrBadRequest, "CHANNEL_INVALID")

	// ErrChannelMonoforumUnsupported
	// | 400 | CHANNEL_MONOFORUM_UNSUPPORTED | [Monoforums](https://core.telegram.org/api/channel#monoforums) do not support this feature. |
	ErrChannelMonoforumUnsupported = xerr.NewCodeError(ErrBadRequest, "CHANNEL_MONOFORUM_UNSUPPORTED")

	// ErrChannelParicipantMissing
	// | 400 | CHANNEL_PARICIPANT_MISSING | The current user is not in the channel. |
	ErrChannelParicipantMissing = xerr.NewCodeError(ErrBadRequest, "CHANNEL_PARICIPANT_MISSING")

	// ErrChannelTooBig
	// | 400 | CHANNEL_TOO_BIG | This channel has too many participants (>1000) to be deleted. |
	ErrChannelTooBig = xerr.NewCodeError(ErrBadRequest, "CHANNEL_TOO_BIG")

	// ErrChannelsAdminLocatedTooMuch
	// | 400 | CHANNELS_ADMIN_LOCATED_TOO_MUCH | The user has reached the limit of public geogroups. |
	ErrChannelsAdminLocatedTooMuch = xerr.NewCodeError(ErrBadRequest, "CHANNELS_ADMIN_LOCATED_TOO_MUCH")

	// ErrChannelsAdminPublicTooMuch
	// | 400 | CHANNELS_ADMIN_PUBLIC_TOO_MUCH | You're admin of too many public channels, make some channels private to change the username of this channel. |
	ErrChannelsAdminPublicTooMuch = xerr.NewCodeError(ErrBadRequest, "CHANNELS_ADMIN_PUBLIC_TOO_MUCH")

	// ErrChannelsTooMuch
	// | 400 | CHANNELS_TOO_MUCH | You have joined too many channels/supergroups. |
	ErrChannelsTooMuch = xerr.NewCodeError(ErrBadRequest, "CHANNELS_TOO_MUCH")

	// ErrChargeAlreadyRefunded
	// | 400 | CHARGE_ALREADY_REFUNDED | The transaction was already refunded. |
	ErrChargeAlreadyRefunded = xerr.NewCodeError(ErrBadRequest, "CHARGE_ALREADY_REFUNDED")

	// ErrChargeIdEmpty
	// | 400 | CHARGE_ID_EMPTY | The specified charge_id is empty. |
	ErrChargeIdEmpty = xerr.NewCodeError(ErrBadRequest, "CHARGE_ID_EMPTY")

	// ErrChargeIdInvalid
	// | 400 | CHARGE_ID_INVALID | The specified charge_id is invalid. |
	ErrChargeIdInvalid = xerr.NewCodeError(ErrBadRequest, "CHARGE_ID_INVALID")

	// ErrChatAboutNotModified
	// | 400 | CHAT_ABOUT_NOT_MODIFIED | About text has not changed. |
	ErrChatAboutNotModified = xerr.NewCodeError(ErrBadRequest, "CHAT_ABOUT_NOT_MODIFIED")

	// ErrChatAboutTooLong
	// | 400 | CHAT_ABOUT_TOO_LONG | Chat about too long. |
	ErrChatAboutTooLong = xerr.NewCodeError(ErrBadRequest, "CHAT_ABOUT_TOO_LONG")

	// ErrChatDiscussionUnallowed
	// | 400 | CHAT_DISCUSSION_UNALLOWED | You can't enable forum topics in a discussion group linked to a channel. |
	ErrChatDiscussionUnallowed = xerr.NewCodeError(ErrBadRequest, "CHAT_DISCUSSION_UNALLOWED")

	// ErrChatIdEmpty
	// | 400 | CHAT_ID_EMPTY | The provided chat ID is empty. |
	ErrChatIdEmpty = xerr.NewCodeError(ErrBadRequest, "CHAT_ID_EMPTY")

	// ErrChatIdInvalid
	// | 400 | CHAT_ID_INVALID | The provided chat id is invalid. |
	ErrChatIdInvalid = xerr.NewCodeError(ErrBadRequest, "CHAT_ID_INVALID")

	// ErrChatInvitePermanent
	// | 400 | CHAT_INVITE_PERMANENT | You can't set an expiration date on permanent invite links. |
	ErrChatInvitePermanent = xerr.NewCodeError(ErrBadRequest, "CHAT_INVITE_PERMANENT")

	// ErrChatLinkExists
	// | 400 | CHAT_LINK_EXISTS | The chat is public, you can't hide the history to new users. |
	ErrChatLinkExists = xerr.NewCodeError(ErrBadRequest, "CHAT_LINK_EXISTS")

	// ErrChatMemberAddFailed
	// | 400 | CHAT_MEMBER_ADD_FAILED | Could not add participants. |
	ErrChatMemberAddFailed = xerr.NewCodeError(ErrBadRequest, "CHAT_MEMBER_ADD_FAILED")

	// ErrChatNotModified
	// | 400 | CHAT_NOT_MODIFIED | No changes were made to chat information because the new information you passed is identical to the current information. |
	ErrChatNotModified = xerr.NewCodeError(ErrBadRequest, "CHAT_NOT_MODIFIED")

	// ErrChatPublicRequired
	// | 400 | CHAT_PUBLIC_REQUIRED | You can only enable join requests in public groups. |
	ErrChatPublicRequired = xerr.NewCodeError(ErrBadRequest, "CHAT_PUBLIC_REQUIRED")

	// ErrChatRestricted
	// | 400 | CHAT_RESTRICTED | You can't send messages in this chat, you were restricted. |
	ErrChatRestricted = xerr.NewCodeError(ErrBadRequest, "CHAT_RESTRICTED")

	// ErrChatRevokeDateUnsupported
	// | 400 | CHAT_REVOKE_DATE_UNSUPPORTED | `min_date` and `max_date` are not available for using with non-user peers. |
	ErrChatRevokeDateUnsupported = xerr.NewCodeError(ErrBadRequest, "CHAT_REVOKE_DATE_UNSUPPORTED")

	// ErrChatTitleEmpty
	// | 400 | CHAT_TITLE_EMPTY | No chat title provided. |
	ErrChatTitleEmpty = xerr.NewCodeError(ErrBadRequest, "CHAT_TITLE_EMPTY")

	// ErrChatTooBig
	// | 400 | CHAT_TOO_BIG | This method is not available for groups with more than `chat_read_mark_size_threshold` members, [see client configuration &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrChatTooBig = xerr.NewCodeError(ErrBadRequest, "CHAT_TOO_BIG")

	// ErrChatlinkSlugEmpty
	// | 400 | CHATLINK_SLUG_EMPTY | The specified slug is empty. |
	ErrChatlinkSlugEmpty = xerr.NewCodeError(ErrBadRequest, "CHATLINK_SLUG_EMPTY")

	// ErrChatlinkSlugExpired
	// | 400 | CHATLINK_SLUG_EXPIRED | The specified [business chat link](https://core.telegram.org/api/business#business-chat-links) has expired. |
	ErrChatlinkSlugExpired = xerr.NewCodeError(ErrBadRequest, "CHATLINK_SLUG_EXPIRED")

	// ErrChatlinksTooMuch
	// | 400 | CHATLINKS_TOO_MUCH | Too many [business chat links](https://core.telegram.org/api/business#business-chat-links) were created, please delete some older links. |
	ErrChatlinksTooMuch = xerr.NewCodeError(ErrBadRequest, "CHATLINKS_TOO_MUCH")

	// ErrChatlistExcludeInvalid
	// | 400 | CHATLIST_EXCLUDE_INVALID | The specified `exclude_peers` are invalid. |
	ErrChatlistExcludeInvalid = xerr.NewCodeError(ErrBadRequest, "CHATLIST_EXCLUDE_INVALID")

	// ErrChatlistsTooMuch
	// | 400 | CHATLISTS_TOO_MUCH | You have created too many folder links, hitting the `chatlist_invites_limit_default`/`chatlist_invites_limit_premium` [limits &raquo;](https://core.telegram.org/api/config#chatlist-invites-limit-default). |
	ErrChatlistsTooMuch = xerr.NewCodeError(ErrBadRequest, "CHATLISTS_TOO_MUCH")

	// ErrCodeEmpty
	// | 400 | CODE_EMPTY | The provided code is empty. |
	ErrCodeEmpty = xerr.NewCodeError(ErrBadRequest, "CODE_EMPTY")

	// ErrCodeHashInvalid
	// | 400 | CODE_HASH_INVALID | Code hash invalid. |
	ErrCodeHashInvalid = xerr.NewCodeError(ErrBadRequest, "CODE_HASH_INVALID")

	// ErrCodeInvalid
	// | 400 | CODE_INVALID | Code invalid. |
	ErrCodeInvalid = xerr.NewCodeError(ErrBadRequest, "CODE_INVALID")

	// ErrCollectibleInvalid
	// | 400 | COLLECTIBLE_INVALID | The specified collectible is invalid. |
	ErrCollectibleInvalid = xerr.NewCodeError(ErrBadRequest, "COLLECTIBLE_INVALID")

	// ErrCollectibleNotFound
	// | 400 | COLLECTIBLE_NOT_FOUND | The specified collectible could not be found. |
	ErrCollectibleNotFound = xerr.NewCodeError(ErrBadRequest, "COLLECTIBLE_NOT_FOUND")

	// ErrColorInvalid
	// | 400 | COLOR_INVALID | The specified color palette ID was invalid. |
	ErrColorInvalid = xerr.NewCodeError(ErrBadRequest, "COLOR_INVALID")

	// ErrConnectionApiIdInvalid
	// | 400 | CONNECTION_API_ID_INVALID | The provided API id is invalid. |
	ErrConnectionApiIdInvalid = xerr.NewCodeError(ErrBadRequest, "CONNECTION_API_ID_INVALID")

	// ErrConnectionAppVersionEmpty
	// | 400 | CONNECTION_APP_VERSION_EMPTY | App version is empty. |
	ErrConnectionAppVersionEmpty = xerr.NewCodeError(ErrBadRequest, "CONNECTION_APP_VERSION_EMPTY")

	// ErrConnectionDeviceModelEmpty
	// | 400 | CONNECTION_DEVICE_MODEL_EMPTY | The specified device model is empty. |
	ErrConnectionDeviceModelEmpty = xerr.NewCodeError(ErrBadRequest, "CONNECTION_DEVICE_MODEL_EMPTY")

	// ErrConnectionIdInvalid
	// | 400 | CONNECTION_ID_INVALID | The specified connection ID is invalid. |
	ErrConnectionIdInvalid = xerr.NewCodeError(ErrBadRequest, "CONNECTION_ID_INVALID")

	// ErrConnectionLangPackInvalid
	// | 400 | CONNECTION_LANG_PACK_INVALID | The specified language pack is empty. |
	ErrConnectionLangPackInvalid = xerr.NewCodeError(ErrBadRequest, "CONNECTION_LANG_PACK_INVALID")

	// ErrConnectionLayerInvalid
	// | 400 | CONNECTION_LAYER_INVALID | Layer invalid. |
	ErrConnectionLayerInvalid = xerr.NewCodeError(ErrBadRequest, "CONNECTION_LAYER_INVALID")

	// ErrConnectionNotInited
	// | 400 | CONNECTION_NOT_INITED | Please initialize the connection using initConnection before making queries. |
	ErrConnectionNotInited = xerr.NewCodeError(ErrBadRequest, "CONNECTION_NOT_INITED")

	// ErrConnectionSystemEmpty
	// | 400 | CONNECTION_SYSTEM_EMPTY | The specified system version is empty. |
	ErrConnectionSystemEmpty = xerr.NewCodeError(ErrBadRequest, "CONNECTION_SYSTEM_EMPTY")

	// ErrConnectionSystemLangCodeEmpty
	// | 400 | CONNECTION_SYSTEM_LANG_CODE_EMPTY | The specified system language code is empty. |
	ErrConnectionSystemLangCodeEmpty = xerr.NewCodeError(ErrBadRequest, "CONNECTION_SYSTEM_LANG_CODE_EMPTY")

	// ErrContactAddMissing
	// | 400 | CONTACT_ADD_MISSING | Contact to add is missing. |
	ErrContactAddMissing = xerr.NewCodeError(ErrBadRequest, "CONTACT_ADD_MISSING")

	// ErrContactIdInvalid
	// | 400 | CONTACT_ID_INVALID | The provided contact ID is invalid. |
	ErrContactIdInvalid = xerr.NewCodeError(ErrBadRequest, "CONTACT_ID_INVALID")

	// ErrContactMissing
	// | 400 | CONTACT_MISSING | The specified user is not a contact. |
	ErrContactMissing = xerr.NewCodeError(ErrBadRequest, "CONTACT_MISSING")

	// ErrContactNameEmpty
	// | 400 | CONTACT_NAME_EMPTY | Contact name empty. |
	ErrContactNameEmpty = xerr.NewCodeError(ErrBadRequest, "CONTACT_NAME_EMPTY")

	// ErrContactReqMissing
	// | 400 | CONTACT_REQ_MISSING | Missing contact request. |
	ErrContactReqMissing = xerr.NewCodeError(ErrBadRequest, "CONTACT_REQ_MISSING")

	// ErrCreateCallFailed
	// | 400 | CREATE_CALL_FAILED | An error occurred while creating the call. |
	ErrCreateCallFailed = xerr.NewCodeError(ErrBadRequest, "CREATE_CALL_FAILED")

	// ErrCurrencyTotalAmountInvalid
	// | 400 | CURRENCY_TOTAL_AMOUNT_INVALID | The total amount of all prices is invalid. |
	ErrCurrencyTotalAmountInvalid = xerr.NewCodeError(ErrBadRequest, "CURRENCY_TOTAL_AMOUNT_INVALID")

	// ErrCustomReactionsTooMany
	// | 400 | CUSTOM_REACTIONS_TOO_MANY | Too many custom reactions were specified. |
	ErrCustomReactionsTooMany = xerr.NewCodeError(ErrBadRequest, "CUSTOM_REACTIONS_TOO_MANY")

	// ErrDataHashSizeInvalid
	// | 400 | DATA_HASH_SIZE_INVALID | The size of the specified secureValueErrorData.data_hash is invalid. |
	ErrDataHashSizeInvalid = xerr.NewCodeError(ErrBadRequest, "DATA_HASH_SIZE_INVALID")

	// ErrDataInvalid
	// | 400 | DATA_INVALID | Encrypted data invalid. |
	ErrDataInvalid = xerr.NewCodeError(ErrBadRequest, "DATA_INVALID")

	// ErrDataJsonInvalid
	// | 400 | DATA_JSON_INVALID | The provided JSON data is invalid. |
	ErrDataJsonInvalid = xerr.NewCodeError(ErrBadRequest, "DATA_JSON_INVALID")

	// ErrDataTooLong
	// | 400 | DATA_TOO_LONG | Data too long. |
	ErrDataTooLong = xerr.NewCodeError(ErrBadRequest, "DATA_TOO_LONG")

	// ErrDateEmpty
	// | 400 | DATE_EMPTY | Date empty. |
	ErrDateEmpty = xerr.NewCodeError(ErrBadRequest, "DATE_EMPTY")

	// ErrDcIdInvalid
	// | 400 | DC_ID_INVALID | The provided DC ID is invalid. |
	ErrDcIdInvalid = xerr.NewCodeError(ErrBadRequest, "DC_ID_INVALID")

	// ErrDhGAInvalid
	// | 400 | DH_G_A_INVALID | g_a invalid. |
	ErrDhGAInvalid = xerr.NewCodeError(ErrBadRequest, "DH_G_A_INVALID")

	// ErrDocumentInvalid
	// | 400 | DOCUMENT_INVALID | The specified document is invalid. |
	ErrDocumentInvalid = xerr.NewCodeError(ErrBadRequest, "DOCUMENT_INVALID")

	// ErrEffectIdInvalid
	// | 400 | EFFECT_ID_INVALID | The specified effect ID is invalid. |
	ErrEffectIdInvalid = xerr.NewCodeError(ErrBadRequest, "EFFECT_ID_INVALID")

	// ErrEmailHashExpired
	// | 400 | EMAIL_HASH_EXPIRED | Email hash expired. |
	ErrEmailHashExpired = xerr.NewCodeError(ErrBadRequest, "EMAIL_HASH_EXPIRED")

	// ErrEmailInvalid
	// | 400 | EMAIL_INVALID | The specified email is invalid. |
	ErrEmailInvalid = xerr.NewCodeError(ErrBadRequest, "EMAIL_INVALID")

	// ErrEmailNotAllowed
	// | 400 | EMAIL_NOT_ALLOWED | The specified email cannot be used to complete the operation. |
	ErrEmailNotAllowed = xerr.NewCodeError(ErrBadRequest, "EMAIL_NOT_ALLOWED")

	// ErrEmailNotSetup
	// | 400 | EMAIL_NOT_SETUP | In order to change the login email with emailVerifyPurposeLoginChange, an existing login email must already be set using emailVerifyPurposeLoginSetup. |
	ErrEmailNotSetup = xerr.NewCodeError(ErrBadRequest, "EMAIL_NOT_SETUP")

	// ErrEmailUnconfirmed
	// | 400 | EMAIL_UNCONFIRMED | Email unconfirmed. |
	ErrEmailUnconfirmed = xerr.NewCodeError(ErrBadRequest, "EMAIL_UNCONFIRMED")

	// ErrEmailUnconfirmed_%d
	// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
	// ErrEmailUnconfirmed_%d = xerr.NewCodeError(ErrBadRequest, "EMAIL_UNCONFIRMED_%d")

	// ErrEmailVerifyExpired
	// | 400 | EMAIL_VERIFY_EXPIRED | The verification email has expired. |
	ErrEmailVerifyExpired = xerr.NewCodeError(ErrBadRequest, "EMAIL_VERIFY_EXPIRED")

	// ErrEmojiInvalid
	// | 400 | EMOJI_INVALID | The specified theme emoji is valid. |
	ErrEmojiInvalid = xerr.NewCodeError(ErrBadRequest, "EMOJI_INVALID")

	// ErrEmojiMarkupInvalid
	// | 400 | EMOJI_MARKUP_INVALID | The specified `video_emoji_markup` was invalid. |
	ErrEmojiMarkupInvalid = xerr.NewCodeError(ErrBadRequest, "EMOJI_MARKUP_INVALID")

	// ErrEmojiNotModified
	// | 400 | EMOJI_NOT_MODIFIED | The theme wasn't changed. |
	ErrEmojiNotModified = xerr.NewCodeError(ErrBadRequest, "EMOJI_NOT_MODIFIED")

	// ErrEmoticonEmpty
	// | 400 | EMOTICON_EMPTY | The emoji is empty. |
	ErrEmoticonEmpty = xerr.NewCodeError(ErrBadRequest, "EMOTICON_EMPTY")

	// ErrEmoticonInvalid
	// | 400 | EMOTICON_INVALID | The specified emoji is invalid. |
	ErrEmoticonInvalid = xerr.NewCodeError(ErrBadRequest, "EMOTICON_INVALID")

	// ErrEmoticonStickerpackMissing
	// | 400 | EMOTICON_STICKERPACK_MISSING | inputStickerSetDice.emoji cannot be empty. |
	ErrEmoticonStickerpackMissing = xerr.NewCodeError(ErrBadRequest, "EMOTICON_STICKERPACK_MISSING")

	// ErrEncryptedMessageInvalid
	// | 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message invalid. |
	ErrEncryptedMessageInvalid = xerr.NewCodeError(ErrBadRequest, "ENCRYPTED_MESSAGE_INVALID")

	// ErrEncryptionAlreadyAccepted
	// | 400 | ENCRYPTION_ALREADY_ACCEPTED | Secret chat already accepted. |
	ErrEncryptionAlreadyAccepted = xerr.NewCodeError(ErrBadRequest, "ENCRYPTION_ALREADY_ACCEPTED")

	// ErrEncryptionAlreadyDeclined
	// | 400 | ENCRYPTION_ALREADY_DECLINED | The secret chat was already declined. |
	ErrEncryptionAlreadyDeclined = xerr.NewCodeError(ErrBadRequest, "ENCRYPTION_ALREADY_DECLINED")

	// ErrEncryptionDeclined
	// | 400 | ENCRYPTION_DECLINED | The secret chat was declined. |
	ErrEncryptionDeclined = xerr.NewCodeError(ErrBadRequest, "ENCRYPTION_DECLINED")

	// ErrEncryptionIdInvalid
	// | 400 | ENCRYPTION_ID_INVALID | The provided secret chat ID is invalid. |
	ErrEncryptionIdInvalid = xerr.NewCodeError(ErrBadRequest, "ENCRYPTION_ID_INVALID")

	// ErrEntitiesTooLong
	// | 400 | ENTITIES_TOO_LONG | You provided too many styled message entities. |
	ErrEntitiesTooLong = xerr.NewCodeError(ErrBadRequest, "ENTITIES_TOO_LONG")

	// ErrEntityBoundsInvalid
	// | 400 | ENTITY_BOUNDS_INVALID | A specified [entity offset or length](https://core.telegram.org/api/entities#entity-length) is invalid, see [here &raquo;](https://core.telegram.org/api/entities#entity-length) for info on how to properly compute the entity offset/length. |
	ErrEntityBoundsInvalid = xerr.NewCodeError(ErrBadRequest, "ENTITY_BOUNDS_INVALID")

	// ErrEntityMentionUserInvalid
	// | 400 | ENTITY_MENTION_USER_INVALID | You mentioned an invalid user. |
	ErrEntityMentionUserInvalid = xerr.NewCodeError(ErrBadRequest, "ENTITY_MENTION_USER_INVALID")

	// ErrErrorTextEmpty
	// | 400 | ERROR_TEXT_EMPTY | The provided error message is empty. |
	ErrErrorTextEmpty = xerr.NewCodeError(ErrBadRequest, "ERROR_TEXT_EMPTY")

	// ErrExpireDateInvalid
	// | 400 | EXPIRE_DATE_INVALID | The specified expiration date is invalid. |
	ErrExpireDateInvalid = xerr.NewCodeError(ErrBadRequest, "EXPIRE_DATE_INVALID")

	// ErrExpiresAtInvalid
	// | 400 | EXPIRES_AT_INVALID | The specified `expires_at` timestamp is invalid. |
	ErrExpiresAtInvalid = xerr.NewCodeError(ErrBadRequest, "EXPIRES_AT_INVALID")

	// ErrExportCardInvalid
	// | 400 | EXPORT_CARD_INVALID | Provided card is invalid. |
	ErrExportCardInvalid = xerr.NewCodeError(ErrBadRequest, "EXPORT_CARD_INVALID")

	// ErrExtendedMediaAmountInvalid
	// | 400 | EXTENDED_MEDIA_AMOUNT_INVALID | The specified `stars_amount` of the passed [inputMediaPaidMedia](https://core.telegram.org/constructor/inputMediaPaidMedia) is invalid. |
	ErrExtendedMediaAmountInvalid = xerr.NewCodeError(ErrBadRequest, "EXTENDED_MEDIA_AMOUNT_INVALID")

	// ErrExtendedMediaInvalid
	// | 400 | EXTENDED_MEDIA_INVALID | The specified paid media is invalid. |
	ErrExtendedMediaInvalid = xerr.NewCodeError(ErrBadRequest, "EXTENDED_MEDIA_INVALID")

	// ErrExternalUrlInvalid
	// | 400 | EXTERNAL_URL_INVALID | External URL invalid. |
	ErrExternalUrlInvalid = xerr.NewCodeError(ErrBadRequest, "EXTERNAL_URL_INVALID")

	// ErrFileContentTypeInvalid
	// | 400 | FILE_CONTENT_TYPE_INVALID | File content-type is invalid. |
	ErrFileContentTypeInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_CONTENT_TYPE_INVALID")

	// ErrFileEmtpy
	// | 400 | FILE_EMTPY | An empty file was provided. |
	ErrFileEmtpy = xerr.NewCodeError(ErrBadRequest, "FILE_EMTPY")

	// ErrFileIdInvalid
	// | 400 | FILE_ID_INVALID | The provided file id is invalid. |
	ErrFileIdInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_ID_INVALID")

	// ErrFileMigrate_%d
	// | 400 | FILE_MIGRATE_%d | The file currently being accessed is stored in DC %d, please re-send the query to that DC. |
	// ErrFileMigrate_%d = xerr.NewCodeError(ErrBadRequest, "FILE_MIGRATE_%d")

	// ErrFilePartEmpty
	// | 400 | FILE_PART_EMPTY | The provided file part is empty. |
	ErrFilePartEmpty = xerr.NewCodeError(ErrBadRequest, "FILE_PART_EMPTY")

	// ErrFilePartInvalid
	// | 400 | FILE_PART_INVALID | The file part number is invalid. |
	ErrFilePartInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_PART_INVALID")

	// ErrFilePartLengthInvalid
	// | 400 | FILE_PART_LENGTH_INVALID | The length of a file part is invalid. |
	ErrFilePartLengthInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_PART_LENGTH_INVALID")

	// ErrFilePartSizeChanged
	// | 400 | FILE_PART_SIZE_CHANGED | Provided file part size has changed. |
	ErrFilePartSizeChanged = xerr.NewCodeError(ErrBadRequest, "FILE_PART_SIZE_CHANGED")

	// ErrFilePartSizeInvalid
	// | 400 | FILE_PART_SIZE_INVALID | The provided file part size is invalid. |
	ErrFilePartSizeInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_PART_SIZE_INVALID")

	// ErrFilePartTooBig
	// | 400 | FILE_PART_TOO_BIG | The uploaded file part is too big. |
	ErrFilePartTooBig = xerr.NewCodeError(ErrBadRequest, "FILE_PART_TOO_BIG")

	// ErrFilePartTooSmall
	// | 400 | FILE_PART_TOO_SMALL | The size of the uploaded file part is too small, please see the documentation for the allowed sizes. |
	ErrFilePartTooSmall = xerr.NewCodeError(ErrBadRequest, "FILE_PART_TOO_SMALL")

	// ErrFilePart_%dMissing
	// | 400 | FILE_PART_%d_MISSING | Part %d of the file is missing from storage. Try repeating the method call to resave the part. |
	// ErrFilePart_%dMissing = xerr.NewCodeError(ErrBadRequest, "FILE_PART_%d_MISSING")

	// ErrFilePartsInvalid
	// | 400 | FILE_PARTS_INVALID | The number of file parts is invalid. |
	ErrFilePartsInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_PARTS_INVALID")

	// ErrFileReferenceEmpty
	// | 400 | FILE_REFERENCE_EMPTY | An empty [file reference](https://core.telegram.org/api/file-references) was specified. |
	ErrFileReferenceEmpty = xerr.NewCodeError(ErrBadRequest, "FILE_REFERENCE_EMPTY")

	// ErrFileReferenceExpired
	// | 400 | FILE_REFERENCE_EXPIRED | File reference expired, it must be refetched as described in [the documentation](https://core.telegram.org/api/file-references). |
	ErrFileReferenceExpired = xerr.NewCodeError(ErrBadRequest, "FILE_REFERENCE_EXPIRED")

	// ErrFileReferenceInvalid
	// | 400 | FILE_REFERENCE_INVALID | The specified [file reference](https://core.telegram.org/api/file-references) is invalid. |
	ErrFileReferenceInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_REFERENCE_INVALID")

	// ErrFileReference_%dExpired
	// | 400 | FILE_REFERENCE_%d_EXPIRED | The file reference of the media file at index %d in the passed media array expired, it [must be refreshed as specified in the documentation](https://core.telegram.org/api/file-references). . |
	// ErrFileReference_%dExpired = xerr.NewCodeError(ErrBadRequest, "FILE_REFERENCE_%d_EXPIRED")

	// ErrFileReference_%dInvalid
	// | 400 | FILE_REFERENCE_%d_INVALID | The [file reference](https://core.telegram.org/api/file-references) of the media file at index %d in the passed media array is invalid. |
	// ErrFileReference_%dInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_REFERENCE_%d_INVALID")

	// ErrFileTitleEmpty
	// | 400 | FILE_TITLE_EMPTY | An empty file title was specified. |
	ErrFileTitleEmpty = xerr.NewCodeError(ErrBadRequest, "FILE_TITLE_EMPTY")

	// ErrFileTokenInvalid
	// | 400 | FILE_TOKEN_INVALID | The master DC did not accept the `file_token` (e.g., the token has expired). Continue downloading the file from the master DC using upload.getFile. |
	ErrFileTokenInvalid = xerr.NewCodeError(ErrBadRequest, "FILE_TOKEN_INVALID")

	// ErrFilterIdInvalid
	// | 400 | FILTER_ID_INVALID | The specified filter ID is invalid. |
	ErrFilterIdInvalid = xerr.NewCodeError(ErrBadRequest, "FILTER_ID_INVALID")

	// ErrFilterIncludeEmpty
	// | 400 | FILTER_INCLUDE_EMPTY | The include_peers vector of the filter is empty. |
	ErrFilterIncludeEmpty = xerr.NewCodeError(ErrBadRequest, "FILTER_INCLUDE_EMPTY")

	// ErrFilterNotSupported
	// | 400 | FILTER_NOT_SUPPORTED | The specified filter cannot be used in this context. |
	ErrFilterNotSupported = xerr.NewCodeError(ErrBadRequest, "FILTER_NOT_SUPPORTED")

	// ErrFilterTitleEmpty
	// | 400 | FILTER_TITLE_EMPTY | The title field of the filter is empty. |
	ErrFilterTitleEmpty = xerr.NewCodeError(ErrBadRequest, "FILTER_TITLE_EMPTY")

	// ErrFirstnameInvalid
	// | 400 | FIRSTNAME_INVALID | The first name is invalid. |
	ErrFirstnameInvalid = xerr.NewCodeError(ErrBadRequest, "FIRSTNAME_INVALID")

	// ErrFolderIdEmpty
	// | 400 | FOLDER_ID_EMPTY | An empty folder ID was specified. |
	ErrFolderIdEmpty = xerr.NewCodeError(ErrBadRequest, "FOLDER_ID_EMPTY")

	// ErrFolderIdInvalid
	// | 400 | FOLDER_ID_INVALID | Invalid folder ID. |
	ErrFolderIdInvalid = xerr.NewCodeError(ErrBadRequest, "FOLDER_ID_INVALID")

	// ErrFormExpired
	// | 400 | FORM_EXPIRED | The form was generated more than 10 minutes ago and has expired, please re-generate it using [payments.getPaymentForm](https://core.telegram.org/method/payments.getPaymentForm) and pass the new `form_id`. |
	ErrFormExpired = xerr.NewCodeError(ErrBadRequest, "FORM_EXPIRED")

	// ErrFormIdEmpty
	// | 400 | FORM_ID_EMPTY | The specified form ID is empty. |
	ErrFormIdEmpty = xerr.NewCodeError(ErrBadRequest, "FORM_ID_EMPTY")

	// ErrFormSubmitDuplicate
	// | 400 | FORM_SUBMIT_DUPLICATE | The same payment form was already submitted.  . |
	ErrFormSubmitDuplicate = xerr.NewCodeError(ErrBadRequest, "FORM_SUBMIT_DUPLICATE")

	// ErrFormUnsupported
	// | 400 | FORM_UNSUPPORTED | Please update your client. |
	ErrFormUnsupported = xerr.NewCodeError(ErrBadRequest, "FORM_UNSUPPORTED")

	// ErrForumEnabled
	// | 400 | FORUM_ENABLED | You can't execute the specified action because the group is a [forum](https://core.telegram.org/api/forum), disable forum functionality to continue. |
	ErrForumEnabled = xerr.NewCodeError(ErrBadRequest, "FORUM_ENABLED")

	// ErrFromMessageBotDisabled
	// | 400 | FROM_MESSAGE_BOT_DISABLED | Bots can't use fromMessage min constructors. |
	ErrFromMessageBotDisabled = xerr.NewCodeError(ErrBadRequest, "FROM_MESSAGE_BOT_DISABLED")

	// ErrFromPeerInvalid
	// | 400 | FROM_PEER_INVALID | The specified from_id is invalid. |
	ErrFromPeerInvalid = xerr.NewCodeError(ErrBadRequest, "FROM_PEER_INVALID")

	// ErrFrozenParticipantMissing
	// | 400 | FROZEN_PARTICIPANT_MISSING | The current account is [frozen](https://core.telegram.org/api/auth#frozen-accounts), and cannot access the specified peer. |
	ErrFrozenParticipantMissing = xerr.NewCodeError(ErrBadRequest, "FROZEN_PARTICIPANT_MISSING")

	// ErrGameBotInvalid
	// | 400 | GAME_BOT_INVALID | Bots can't send another bot's game. |
	ErrGameBotInvalid = xerr.NewCodeError(ErrBadRequest, "GAME_BOT_INVALID")

	// ErrGeneralModifyIconForbidden
	// | 400 | GENERAL_MODIFY_ICON_FORBIDDEN | You can't modify the icon of the "General" topic. |
	ErrGeneralModifyIconForbidden = xerr.NewCodeError(ErrBadRequest, "GENERAL_MODIFY_ICON_FORBIDDEN")

	// ErrGeoPointInvalid
	// | 400 | GEO_POINT_INVALID | Invalid geoposition provided. |
	ErrGeoPointInvalid = xerr.NewCodeError(ErrBadRequest, "GEO_POINT_INVALID")

	// ErrGifContentTypeInvalid
	// | 400 | GIF_CONTENT_TYPE_INVALID | GIF content-type invalid. |
	ErrGifContentTypeInvalid = xerr.NewCodeError(ErrBadRequest, "GIF_CONTENT_TYPE_INVALID")

	// ErrGifIdInvalid
	// | 400 | GIF_ID_INVALID | The provided GIF ID is invalid. |
	ErrGifIdInvalid = xerr.NewCodeError(ErrBadRequest, "GIF_ID_INVALID")

	// ErrGiftMonthsInvalid
	// | 400 | GIFT_MONTHS_INVALID | The value passed in invoice.inputInvoicePremiumGiftStars.months is invalid. |
	ErrGiftMonthsInvalid = xerr.NewCodeError(ErrBadRequest, "GIFT_MONTHS_INVALID")

	// ErrGiftSlugExpired
	// | 400 | GIFT_SLUG_EXPIRED | The specified gift slug has expired. |
	ErrGiftSlugExpired = xerr.NewCodeError(ErrBadRequest, "GIFT_SLUG_EXPIRED")

	// ErrGiftSlugInvalid
	// | 400 | GIFT_SLUG_INVALID | The specified slug is invalid. |
	ErrGiftSlugInvalid = xerr.NewCodeError(ErrBadRequest, "GIFT_SLUG_INVALID")

	// ErrGiftStarsInvalid
	// | 400 | GIFT_STARS_INVALID | The specified amount of stars is invalid. |
	ErrGiftStarsInvalid = xerr.NewCodeError(ErrBadRequest, "GIFT_STARS_INVALID")

	// ErrGraphExpiredReload
	// | 400 | GRAPH_EXPIRED_RELOAD | This graph has expired, please obtain a new graph token. |
	ErrGraphExpiredReload = xerr.NewCodeError(ErrBadRequest, "GRAPH_EXPIRED_RELOAD")

	// ErrGraphInvalidReload
	// | 400 | GRAPH_INVALID_RELOAD | Invalid graph token provided, please reload the stats and provide the updated token. |
	ErrGraphInvalidReload = xerr.NewCodeError(ErrBadRequest, "GRAPH_INVALID_RELOAD")

	// ErrGraphOutdatedReload
	// | 400 | GRAPH_OUTDATED_RELOAD | The graph is outdated, please get a new async token using stats.getBroadcastStats. |
	ErrGraphOutdatedReload = xerr.NewCodeError(ErrBadRequest, "GRAPH_OUTDATED_RELOAD")

	// ErrGroupcallAlreadyDiscarded
	// | 400 | GROUPCALL_ALREADY_DISCARDED | The group call was already discarded. |
	ErrGroupcallAlreadyDiscarded = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_ALREADY_DISCARDED")

	// ErrGroupcallInvalid
	// | 400 | GROUPCALL_INVALID | The specified group call is invalid. |
	ErrGroupcallInvalid = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrGroupcallJoinMissing
	// | 400 | GROUPCALL_JOIN_MISSING | You haven't joined this group call. |
	ErrGroupcallJoinMissing = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_JOIN_MISSING")

	// ErrGroupcallNotModified
	// | 400 | GROUPCALL_NOT_MODIFIED | Group call settings weren't modified. |
	ErrGroupcallNotModified = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_NOT_MODIFIED")

	// ErrGroupcallSsrcDuplicateMuch
	// | 400 | GROUPCALL_SSRC_DUPLICATE_MUCH | The app needs to retry joining the group call with a new SSRC value. |
	ErrGroupcallSsrcDuplicateMuch = xerr.NewCodeError(ErrBadRequest, "GROUPCALL_SSRC_DUPLICATE_MUCH")

	// ErrGroupedMediaInvalid
	// | 400 | GROUPED_MEDIA_INVALID | Invalid grouped media. |
	ErrGroupedMediaInvalid = xerr.NewCodeError(ErrBadRequest, "GROUPED_MEDIA_INVALID")

	// ErrHashInvalid
	// | 400 | HASH_INVALID | The provided hash is invalid. |
	ErrHashInvalid = xerr.NewCodeError(ErrBadRequest, "HASH_INVALID")

	// ErrHashSizeInvalid
	// | 400 | HASH_SIZE_INVALID | The size of the specified secureValueError.hash is invalid. |
	ErrHashSizeInvalid = xerr.NewCodeError(ErrBadRequest, "HASH_SIZE_INVALID")

	// ErrHashtagInvalid
	// | 400 | HASHTAG_INVALID | The specified hashtag is invalid. |
	ErrHashtagInvalid = xerr.NewCodeError(ErrBadRequest, "HASHTAG_INVALID")

	// ErrHideRequesterMissing
	// | 400 | HIDE_REQUESTER_MISSING | The join request was missing or was already handled. |
	ErrHideRequesterMissing = xerr.NewCodeError(ErrBadRequest, "HIDE_REQUESTER_MISSING")

	// ErrIdExpired
	// | 400 | ID_EXPIRED | The passed prepared inline message ID has expired. |
	ErrIdExpired = xerr.NewCodeError(ErrBadRequest, "ID_EXPIRED")

	// ErrIdInvalid
	// | 400 | ID_INVALID | The passed ID is invalid. |
	ErrIdInvalid = xerr.NewCodeError(ErrBadRequest, "ID_INVALID")

	// ErrImageProcessFailed
	// | 400 | IMAGE_PROCESS_FAILED | Failure while processing image. |
	ErrImageProcessFailed = xerr.NewCodeError(ErrBadRequest, "IMAGE_PROCESS_FAILED")

	// ErrImportFileInvalid
	// | 400 | IMPORT_FILE_INVALID | The specified chat export file is invalid. |
	ErrImportFileInvalid = xerr.NewCodeError(ErrBadRequest, "IMPORT_FILE_INVALID")

	// ErrImportFormatDateInvalid
	// | 400 | IMPORT_FORMAT_DATE_INVALID | The date specified in the import file is invalid. |
	ErrImportFormatDateInvalid = xerr.NewCodeError(ErrBadRequest, "IMPORT_FORMAT_DATE_INVALID")

	// ErrImportFormatUnrecognized
	// | 400 | IMPORT_FORMAT_UNRECOGNIZED | The specified chat export file was exported from an unsupported chat app. |
	ErrImportFormatUnrecognized = xerr.NewCodeError(ErrBadRequest, "IMPORT_FORMAT_UNRECOGNIZED")

	// ErrImportIdInvalid
	// | 400 | IMPORT_ID_INVALID | The specified import ID is invalid. |
	ErrImportIdInvalid = xerr.NewCodeError(ErrBadRequest, "IMPORT_ID_INVALID")

	// ErrImportTokenInvalid
	// | 400 | IMPORT_TOKEN_INVALID | The specified token is invalid. |
	ErrImportTokenInvalid = xerr.NewCodeError(ErrBadRequest, "IMPORT_TOKEN_INVALID")

	// ErrInlineResultExpired
	// | 400 | INLINE_RESULT_EXPIRED | The inline query expired. |
	ErrInlineResultExpired = xerr.NewCodeError(ErrBadRequest, "INLINE_RESULT_EXPIRED")

	// ErrInputChatlistInvalid
	// | 400 | INPUT_CHATLIST_INVALID | The specified folder is invalid. |
	ErrInputChatlistInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_CHATLIST_INVALID")

	// ErrInputConstructorInvalid
	// | 400 | INPUT_CONSTRUCTOR_INVALID | The specified TL constructor is invalid. |
	ErrInputConstructorInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// ErrInputFetchError
	// | 400 | INPUT_FETCH_ERROR | An error occurred while parsing the provided TL constructor. |
	ErrInputFetchError = xerr.NewCodeError(ErrBadRequest, "INPUT_FETCH_ERROR")

	// ErrInputFetchFail
	// | 400 | INPUT_FETCH_FAIL | An error occurred while parsing the provided TL constructor. |
	ErrInputFetchFail = xerr.NewCodeError(ErrBadRequest, "INPUT_FETCH_FAIL")

	// ErrInputFileInvalid
	// | 400 | INPUT_FILE_INVALID | The specified [InputFile](https://core.telegram.org/type/InputFile) is invalid. |
	ErrInputFileInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_FILE_INVALID")

	// ErrInputFilterInvalid
	// | 400 | INPUT_FILTER_INVALID | The specified filter is invalid. |
	ErrInputFilterInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_FILTER_INVALID")

	// ErrInputLayerInvalid
	// | 400 | INPUT_LAYER_INVALID | The specified layer is invalid. |
	ErrInputLayerInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_LAYER_INVALID")

	// ErrInputMethodInvalid
	// | 400 | INPUT_METHOD_INVALID | The specified method is invalid. |
	ErrInputMethodInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_METHOD_INVALID")

	// ErrInputPeersEmpty
	// | 400 | INPUT_PEERS_EMPTY | The specified peer array is empty. |
	ErrInputPeersEmpty = xerr.NewCodeError(ErrBadRequest, "INPUT_PEERS_EMPTY")

	// ErrInputPurposeInvalid
	// | 400 | INPUT_PURPOSE_INVALID | The specified payment purpose is invalid. |
	ErrInputPurposeInvalid = xerr.NewCodeError(ErrBadRequest, "INPUT_PURPOSE_INVALID")

	// ErrInputRequestTooLong
	// | 400 | INPUT_REQUEST_TOO_LONG | The request payload is too long. |
	ErrInputRequestTooLong = xerr.NewCodeError(ErrBadRequest, "INPUT_REQUEST_TOO_LONG")

	// ErrInputTextEmpty
	// | 400 | INPUT_TEXT_EMPTY | The specified text is empty. |
	ErrInputTextEmpty = xerr.NewCodeError(ErrBadRequest, "INPUT_TEXT_EMPTY")

	// ErrInputTextTooLong
	// | 400 | INPUT_TEXT_TOO_LONG | The specified text is too long. |
	ErrInputTextTooLong = xerr.NewCodeError(ErrBadRequest, "INPUT_TEXT_TOO_LONG")

	// ErrInputUserDeactivated
	// | 400 | INPUT_USER_DEACTIVATED | The specified user was deleted. |
	ErrInputUserDeactivated = xerr.NewCodeError(ErrBadRequest, "INPUT_USER_DEACTIVATED")

	// ErrInviteForbiddenWithJoinas
	// | 400 | INVITE_FORBIDDEN_WITH_JOINAS | If the user has anonymously joined a group call as a channel, they can't invite other users to the group call because that would cause deanonymization, because the invite would be sent using the original user ID, not the anonymized channel ID. |
	ErrInviteForbiddenWithJoinas = xerr.NewCodeError(ErrBadRequest, "INVITE_FORBIDDEN_WITH_JOINAS")

	// ErrInviteHashEmpty
	// | 400 | INVITE_HASH_EMPTY | The invite hash is empty. |
	ErrInviteHashEmpty = xerr.NewCodeError(ErrBadRequest, "INVITE_HASH_EMPTY")

	// ErrInviteHashInvalid
	// | 400 | INVITE_HASH_INVALID | The invite hash is invalid. |
	ErrInviteHashInvalid = xerr.NewCodeError(ErrBadRequest, "INVITE_HASH_INVALID")

	// ErrInviteRequestSent
	// | 400 | INVITE_REQUEST_SENT | You have successfully requested to join this chat or channel. |
	ErrInviteRequestSent = xerr.NewCodeError(ErrBadRequest, "INVITE_REQUEST_SENT")

	// ErrInviteRevokedMissing
	// | 400 | INVITE_REVOKED_MISSING | The specified invite link was already revoked or is invalid. |
	ErrInviteRevokedMissing = xerr.NewCodeError(ErrBadRequest, "INVITE_REVOKED_MISSING")

	// ErrInviteSlugEmpty
	// | 400 | INVITE_SLUG_EMPTY | The specified invite slug is empty. |
	ErrInviteSlugEmpty = xerr.NewCodeError(ErrBadRequest, "INVITE_SLUG_EMPTY")

	// ErrInviteSlugExpired
	// | 400 | INVITE_SLUG_EXPIRED | The specified chat folder link has expired. |
	ErrInviteSlugExpired = xerr.NewCodeError(ErrBadRequest, "INVITE_SLUG_EXPIRED")

	// ErrInviteSlugInvalid
	// | 400 | INVITE_SLUG_INVALID | The specified invitation slug is invalid. |
	ErrInviteSlugInvalid = xerr.NewCodeError(ErrBadRequest, "INVITE_SLUG_INVALID")

	// ErrInvitesTooMuch
	// | 400 | INVITES_TOO_MUCH | The maximum number of per-folder invites specified by the `chatlist_invites_limit_default`/`chatlist_invites_limit_premium` [client configuration parameters &raquo;](https://core.telegram.org/api/config#chatlist-invites-limit-default) was reached. |
	ErrInvitesTooMuch = xerr.NewCodeError(ErrBadRequest, "INVITES_TOO_MUCH")

	// ErrInvoiceInvalid
	// | 400 | INVOICE_INVALID | The specified invoice is invalid. |
	ErrInvoiceInvalid = xerr.NewCodeError(ErrBadRequest, "INVOICE_INVALID")

	// ErrInvoicePayloadInvalid
	// | 400 | INVOICE_PAYLOAD_INVALID | The specified invoice payload is invalid. |
	ErrInvoicePayloadInvalid = xerr.NewCodeError(ErrBadRequest, "INVOICE_PAYLOAD_INVALID")

	// ErrJoinAsPeerInvalid
	// | 400 | JOIN_AS_PEER_INVALID | The specified peer cannot be used to join a group call. |
	ErrJoinAsPeerInvalid = xerr.NewCodeError(ErrBadRequest, "JOIN_AS_PEER_INVALID")

	// ErrLangCodeInvalid
	// | 400 | LANG_CODE_INVALID | The specified language code is invalid. |
	ErrLangCodeInvalid = xerr.NewCodeError(ErrBadRequest, "LANG_CODE_INVALID")

	// ErrLangCodeNotSupported
	// | 400 | LANG_CODE_NOT_SUPPORTED | The specified language code is not supported. |
	ErrLangCodeNotSupported = xerr.NewCodeError(ErrBadRequest, "LANG_CODE_NOT_SUPPORTED")

	// ErrLangPackInvalid
	// | 400 | LANG_PACK_INVALID | The provided language pack is invalid. |
	ErrLangPackInvalid = xerr.NewCodeError(ErrBadRequest, "LANG_PACK_INVALID")

	// ErrLanguageInvalid
	// | 400 | LANGUAGE_INVALID | The specified lang_code is invalid. |
	ErrLanguageInvalid = xerr.NewCodeError(ErrBadRequest, "LANGUAGE_INVALID")

	// ErrLastnameInvalid
	// | 400 | LASTNAME_INVALID | The last name is invalid. |
	ErrLastnameInvalid = xerr.NewCodeError(ErrBadRequest, "LASTNAME_INVALID")

	// ErrLimitInvalid
	// | 400 | LIMIT_INVALID | The provided limit is invalid. |
	ErrLimitInvalid = xerr.NewCodeError(ErrBadRequest, "LIMIT_INVALID")

	// ErrLinkNotModified
	// | 400 | LINK_NOT_MODIFIED | Discussion link not modified. |
	ErrLinkNotModified = xerr.NewCodeError(ErrBadRequest, "LINK_NOT_MODIFIED")

	// ErrLocationInvalid
	// | 400 | LOCATION_INVALID | The provided location is invalid. |
	ErrLocationInvalid = xerr.NewCodeError(ErrBadRequest, "LOCATION_INVALID")

	// ErrMaxDateInvalid
	// | 400 | MAX_DATE_INVALID | The specified maximum date is invalid. |
	ErrMaxDateInvalid = xerr.NewCodeError(ErrBadRequest, "MAX_DATE_INVALID")

	// ErrMaxIdInvalid
	// | 400 | MAX_ID_INVALID | The provided max ID is invalid. |
	ErrMaxIdInvalid = xerr.NewCodeError(ErrBadRequest, "MAX_ID_INVALID")

	// ErrMaxQtsInvalid
	// | 400 | MAX_QTS_INVALID | The specified max_qts is invalid. |
	ErrMaxQtsInvalid = xerr.NewCodeError(ErrBadRequest, "MAX_QTS_INVALID")

	// ErrMd5ChecksumInvalid
	// | 400 | MD5_CHECKSUM_INVALID | The MD5 checksums do not match. |
	ErrMd5ChecksumInvalid = xerr.NewCodeError(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrMediaAlreadyPaid
	// | 400 | MEDIA_ALREADY_PAID | You already paid for the specified media. |
	ErrMediaAlreadyPaid = xerr.NewCodeError(ErrBadRequest, "MEDIA_ALREADY_PAID")

	// ErrMediaCaptionTooLong
	// | 400 | MEDIA_CAPTION_TOO_LONG | The caption is too long. |
	ErrMediaCaptionTooLong = xerr.NewCodeError(ErrBadRequest, "MEDIA_CAPTION_TOO_LONG")

	// ErrMediaEmpty
	// | 400 | MEDIA_EMPTY | The provided media object is invalid. |
	ErrMediaEmpty = xerr.NewCodeError(ErrBadRequest, "MEDIA_EMPTY")

	// ErrMediaFileInvalid
	// | 400 | MEDIA_FILE_INVALID | The specified media file is invalid. |
	ErrMediaFileInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_FILE_INVALID")

	// ErrMediaGroupedInvalid
	// | 400 | MEDIA_GROUPED_INVALID | You tried to send media of different types in an album. |
	ErrMediaGroupedInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_GROUPED_INVALID")

	// ErrMediaInvalid
	// | 400 | MEDIA_INVALID | Media invalid. |
	ErrMediaInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_INVALID")

	// ErrMediaNewInvalid
	// | 400 | MEDIA_NEW_INVALID | The new media is invalid. |
	ErrMediaNewInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_NEW_INVALID")

	// ErrMediaPrevInvalid
	// | 400 | MEDIA_PREV_INVALID | Previous media invalid. |
	ErrMediaPrevInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_PREV_INVALID")

	// ErrMediaTtlInvalid
	// | 400 | MEDIA_TTL_INVALID | The specified media TTL is invalid. |
	ErrMediaTtlInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_TTL_INVALID")

	// ErrMediaTypeInvalid
	// | 400 | MEDIA_TYPE_INVALID | The specified media type cannot be used in stories. |
	ErrMediaTypeInvalid = xerr.NewCodeError(ErrBadRequest, "MEDIA_TYPE_INVALID")

	// ErrMediaVideoStoryMissing
	// | 400 | MEDIA_VIDEO_STORY_MISSING | A non-story video cannot be repubblished as a story (emitted when trying to resend a non-story video as a story using inputDocument). |
	ErrMediaVideoStoryMissing = xerr.NewCodeError(ErrBadRequest, "MEDIA_VIDEO_STORY_MISSING")

	// ErrMegagroupGeoRequired
	// | 400 | MEGAGROUP_GEO_REQUIRED | This method can only be invoked on a geogroup. |
	ErrMegagroupGeoRequired = xerr.NewCodeError(ErrBadRequest, "MEGAGROUP_GEO_REQUIRED")

	// ErrMegagroupIdInvalid
	// | 400 | MEGAGROUP_ID_INVALID | Invalid supergroup ID. |
	ErrMegagroupIdInvalid = xerr.NewCodeError(ErrBadRequest, "MEGAGROUP_ID_INVALID")

	// ErrMegagroupPrehistoryHidden
	// | 400 | MEGAGROUP_PREHISTORY_HIDDEN | Group with hidden history for new members can't be set as discussion groups. |
	ErrMegagroupPrehistoryHidden = xerr.NewCodeError(ErrBadRequest, "MEGAGROUP_PREHISTORY_HIDDEN")

	// ErrMegagroupRequired
	// | 400 | MEGAGROUP_REQUIRED | You can only use this method on a supergroup. |
	ErrMegagroupRequired = xerr.NewCodeError(ErrBadRequest, "MEGAGROUP_REQUIRED")

	// ErrMessageEditTimeExpired
	// | 400 | MESSAGE_EDIT_TIME_EXPIRED | You can't edit this message anymore, too much time has passed since its creation. |
	ErrMessageEditTimeExpired = xerr.NewCodeError(ErrBadRequest, "MESSAGE_EDIT_TIME_EXPIRED")

	// ErrMessageEmpty
	// | 400 | MESSAGE_EMPTY | The provided message is empty. |
	ErrMessageEmpty = xerr.NewCodeError(ErrBadRequest, "MESSAGE_EMPTY")

	// ErrMessageIdInvalid
	// | 400 | MESSAGE_ID_INVALID | The provided message id is invalid. |
	ErrMessageIdInvalid = xerr.NewCodeError(ErrBadRequest, "MESSAGE_ID_INVALID")

	// ErrMessageIdsEmpty
	// | 400 | MESSAGE_IDS_EMPTY | No message ids were provided. |
	ErrMessageIdsEmpty = xerr.NewCodeError(ErrBadRequest, "MESSAGE_IDS_EMPTY")

	// ErrMessageNotModified
	// | 400 | MESSAGE_NOT_MODIFIED | The provided message data is identical to the previous message data, the message wasn't modified. |
	ErrMessageNotModified = xerr.NewCodeError(ErrBadRequest, "MESSAGE_NOT_MODIFIED")

	// ErrMessageNotReadYet
	// | 400 | MESSAGE_NOT_READ_YET | The specified message wasn't read yet. |
	ErrMessageNotReadYet = xerr.NewCodeError(ErrBadRequest, "MESSAGE_NOT_READ_YET")

	// ErrMessagePollClosed
	// | 400 | MESSAGE_POLL_CLOSED | Poll closed. |
	ErrMessagePollClosed = xerr.NewCodeError(ErrBadRequest, "MESSAGE_POLL_CLOSED")

	// ErrMessageTooLong
	// | 400 | MESSAGE_TOO_LONG | The provided message is too long. |
	ErrMessageTooLong = xerr.NewCodeError(ErrBadRequest, "MESSAGE_TOO_LONG")

	// ErrMessageTooOld
	// | 400 | MESSAGE_TOO_OLD | The message is too old, the requested information is not available. |
	ErrMessageTooOld = xerr.NewCodeError(ErrBadRequest, "MESSAGE_TOO_OLD")

	// ErrMinDateInvalid
	// | 400 | MIN_DATE_INVALID | The specified minimum date is invalid. |
	ErrMinDateInvalid = xerr.NewCodeError(ErrBadRequest, "MIN_DATE_INVALID")

	// ErrMonthInvalid
	// | 400 | MONTH_INVALID | The number of months specified in inputInvoicePremiumGiftStars.months is invalid. |
	ErrMonthInvalid = xerr.NewCodeError(ErrBadRequest, "MONTH_INVALID")

	// ErrMsgIdInvalid
	// | 400 | MSG_ID_INVALID | Invalid message ID provided. |
	ErrMsgIdInvalid = xerr.NewCodeError(ErrBadRequest, "MSG_ID_INVALID")

	// ErrMsgTooOld
	// | 400 | MSG_TOO_OLD | [`chat_read_mark_expire_period` seconds](https://core.telegram.org/api/config#chat-read-mark-expire-period) have passed since the message was sent, read receipts were deleted. |
	ErrMsgTooOld = xerr.NewCodeError(ErrBadRequest, "MSG_TOO_OLD")

	// ErrMsgVoiceMissing
	// | 400 | MSG_VOICE_MISSING | The specified message is not a voice message. |
	ErrMsgVoiceMissing = xerr.NewCodeError(ErrBadRequest, "MSG_VOICE_MISSING")

	// ErrMultiMediaTooLong
	// | 400 | MULTI_MEDIA_TOO_LONG | Too many media files for album. |
	ErrMultiMediaTooLong = xerr.NewCodeError(ErrBadRequest, "MULTI_MEDIA_TOO_LONG")

	// ErrNewSaltInvalid
	// | 400 | NEW_SALT_INVALID | The new salt is invalid. |
	ErrNewSaltInvalid = xerr.NewCodeError(ErrBadRequest, "NEW_SALT_INVALID")

	// ErrNewSettingsEmpty
	// | 400 | NEW_SETTINGS_EMPTY | No password is set on the current account, and no new password was specified in `new_settings`. |
	ErrNewSettingsEmpty = xerr.NewCodeError(ErrBadRequest, "NEW_SETTINGS_EMPTY")

	// ErrNewSettingsInvalid
	// | 400 | NEW_SETTINGS_INVALID | The new password settings are invalid. |
	ErrNewSettingsInvalid = xerr.NewCodeError(ErrBadRequest, "NEW_SETTINGS_INVALID")

	// ErrNextOffsetInvalid
	// | 400 | NEXT_OFFSET_INVALID | The specified offset is longer than 64 bytes. |
	ErrNextOffsetInvalid = xerr.NewCodeError(ErrBadRequest, "NEXT_OFFSET_INVALID")

	// ErrNoPaymentNeeded
	// | 400 | NO_PAYMENT_NEEDED | The upgrade/transfer of the specified gift was already paid for or is free. |
	ErrNoPaymentNeeded = xerr.NewCodeError(ErrBadRequest, "NO_PAYMENT_NEEDED")

	// ErrNogeneralHideForbidden
	// | 400 | NOGENERAL_HIDE_FORBIDDEN | Only the "General" topic with `id=1` can be hidden. |
	ErrNogeneralHideForbidden = xerr.NewCodeError(ErrBadRequest, "NOGENERAL_HIDE_FORBIDDEN")

	// ErrNotJoined
	// | 400 | NOT_JOINED | The current user hasn't joined the Peer-to-Peer Login Program. |
	ErrNotJoined = xerr.NewCodeError(ErrBadRequest, "NOT_JOINED")

	// ErrOffsetInvalid
	// | 400 | OFFSET_INVALID | The provided offset is invalid. |
	ErrOffsetInvalid = xerr.NewCodeError(ErrBadRequest, "OFFSET_INVALID")

	// ErrOffsetPeerIdInvalid
	// | 400 | OFFSET_PEER_ID_INVALID | The provided offset peer is invalid. |
	ErrOffsetPeerIdInvalid = xerr.NewCodeError(ErrBadRequest, "OFFSET_PEER_ID_INVALID")

	// ErrOptionInvalid
	// | 400 | OPTION_INVALID | Invalid option selected. |
	ErrOptionInvalid = xerr.NewCodeError(ErrBadRequest, "OPTION_INVALID")

	// ErrOptionsTooMuch
	// | 400 | OPTIONS_TOO_MUCH | Too many options provided. |
	ErrOptionsTooMuch = xerr.NewCodeError(ErrBadRequest, "OPTIONS_TOO_MUCH")

	// ErrOrderInvalid
	// | 400 | ORDER_INVALID | The specified username order is invalid. |
	ErrOrderInvalid = xerr.NewCodeError(ErrBadRequest, "ORDER_INVALID")

	// ErrPackShortNameInvalid
	// | 400 | PACK_SHORT_NAME_INVALID | Short pack name invalid. |
	ErrPackShortNameInvalid = xerr.NewCodeError(ErrBadRequest, "PACK_SHORT_NAME_INVALID")

	// ErrPackShortNameOccupied
	// | 400 | PACK_SHORT_NAME_OCCUPIED | A stickerpack with this name already exists. |
	ErrPackShortNameOccupied = xerr.NewCodeError(ErrBadRequest, "PACK_SHORT_NAME_OCCUPIED")

	// ErrPackTitleInvalid
	// | 400 | PACK_TITLE_INVALID | The stickerpack title is invalid. |
	ErrPackTitleInvalid = xerr.NewCodeError(ErrBadRequest, "PACK_TITLE_INVALID")

	// ErrPackTypeInvalid
	// | 400 | PACK_TYPE_INVALID | The masks and emojis flags are mutually exclusive. |
	ErrPackTypeInvalid = xerr.NewCodeError(ErrBadRequest, "PACK_TYPE_INVALID")

	// ErrParentPeerInvalid
	// | 400 | PARENT_PEER_INVALID | The specified `parent_peer` is invalid. |
	ErrParentPeerInvalid = xerr.NewCodeError(ErrBadRequest, "PARENT_PEER_INVALID")

	// ErrParticipantIdInvalid
	// | 400 | PARTICIPANT_ID_INVALID | The specified participant ID is invalid. |
	ErrParticipantIdInvalid = xerr.NewCodeError(ErrBadRequest, "PARTICIPANT_ID_INVALID")

	// ErrParticipantVersionOutdated
	// | 400 | PARTICIPANT_VERSION_OUTDATED | The other participant does not use an up to date telegram client with support for calls. |
	ErrParticipantVersionOutdated = xerr.NewCodeError(ErrBadRequest, "PARTICIPANT_VERSION_OUTDATED")

	// ErrParticipantsTooFew
	// | 400 | PARTICIPANTS_TOO_FEW | Not enough participants. |
	ErrParticipantsTooFew = xerr.NewCodeError(ErrBadRequest, "PARTICIPANTS_TOO_FEW")

	// ErrPasswordEmpty
	// | 400 | PASSWORD_EMPTY | The provided password is empty. |
	ErrPasswordEmpty = xerr.NewCodeError(ErrBadRequest, "PASSWORD_EMPTY")

	// ErrPasswordHashInvalid
	// | 400 | PASSWORD_HASH_INVALID | The provided password hash is invalid. |
	ErrPasswordHashInvalid = xerr.NewCodeError(ErrBadRequest, "PASSWORD_HASH_INVALID")

	// ErrPasswordMissing
	// | 400 | PASSWORD_MISSING | You must [enable 2FA](https://core.telegram.org/api/srp) before executing this operation. |
	ErrPasswordMissing = xerr.NewCodeError(ErrBadRequest, "PASSWORD_MISSING")

	// ErrPasswordRecoveryExpired
	// | 400 | PASSWORD_RECOVERY_EXPIRED | The recovery code has expired. |
	ErrPasswordRecoveryExpired = xerr.NewCodeError(ErrBadRequest, "PASSWORD_RECOVERY_EXPIRED")

	// ErrPasswordRecoveryNa
	// | 400 | PASSWORD_RECOVERY_NA | No email was set, can't recover password via email. |
	ErrPasswordRecoveryNa = xerr.NewCodeError(ErrBadRequest, "PASSWORD_RECOVERY_NA")

	// ErrPasswordRequired
	// | 400 | PASSWORD_REQUIRED | A [2FA password](https://core.telegram.org/api/srp) must be configured to use Telegram Passport. |
	ErrPasswordRequired = xerr.NewCodeError(ErrBadRequest, "PASSWORD_REQUIRED")

	// ErrPasswordTooFresh_%d
	// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
	// ErrPasswordTooFresh_%d = xerr.NewCodeError(ErrBadRequest, "PASSWORD_TOO_FRESH_%d")

	// ErrPaymentCredentialsInvalid
	// | 400 | PAYMENT_CREDENTIALS_INVALID | The specified payment credentials are invalid. |
	ErrPaymentCredentialsInvalid = xerr.NewCodeError(ErrBadRequest, "PAYMENT_CREDENTIALS_INVALID")

	// ErrPaymentProviderInvalid
	// | 400 | PAYMENT_PROVIDER_INVALID | The specified payment provider is invalid. |
	ErrPaymentProviderInvalid = xerr.NewCodeError(ErrBadRequest, "PAYMENT_PROVIDER_INVALID")

	// ErrPaymentRequired
	// | 400 | PAYMENT_REQUIRED | Payment is required for this action, see [here &raquo;](https://core.telegram.org/api/gifts) for more info. |
	ErrPaymentRequired = xerr.NewCodeError(ErrBadRequest, "PAYMENT_REQUIRED")

	// ErrPeerFlood
	// | 400 | PEER_FLOOD | The current account is spamreported, you cannot execute this action, check @spambot for more info. |
	ErrPeerFlood = xerr.NewCodeError(ErrBadRequest, "PEER_FLOOD")

	// ErrPeerHistoryEmpty
	// | 400 | PEER_HISTORY_EMPTY | You can't pin an empty chat with a user. |
	ErrPeerHistoryEmpty = xerr.NewCodeError(ErrBadRequest, "PEER_HISTORY_EMPTY")

	// ErrPeerIdNotSupported
	// | 400 | PEER_ID_NOT_SUPPORTED | The provided peer ID is not supported. |
	ErrPeerIdNotSupported = xerr.NewCodeError(ErrBadRequest, "PEER_ID_NOT_SUPPORTED")

	// ErrPeerTypesInvalid
	// | 400 | PEER_TYPES_INVALID | The passed [keyboardButtonSwitchInline](https://core.telegram.org/constructor/keyboardButtonSwitchInline).`peer_types` field is invalid. |
	ErrPeerTypesInvalid = xerr.NewCodeError(ErrBadRequest, "PEER_TYPES_INVALID")

	// ErrPeersListEmpty
	// | 400 | PEERS_LIST_EMPTY | The specified list of peers is empty. |
	ErrPeersListEmpty = xerr.NewCodeError(ErrBadRequest, "PEERS_LIST_EMPTY")

	// ErrPersistentTimestampEmpty
	// | 400 | PERSISTENT_TIMESTAMP_EMPTY | Persistent timestamp empty. |
	ErrPersistentTimestampEmpty = xerr.NewCodeError(ErrBadRequest, "PERSISTENT_TIMESTAMP_EMPTY")

	// ErrPersistentTimestampInvalid
	// | 400 | PERSISTENT_TIMESTAMP_INVALID | Persistent timestamp invalid. |
	ErrPersistentTimestampInvalid = xerr.NewCodeError(ErrBadRequest, "PERSISTENT_TIMESTAMP_INVALID")

	// ErrPhoneCodeEmpty
	// | 400 | PHONE_CODE_EMPTY | phone_code is missing. |
	ErrPhoneCodeEmpty = xerr.NewCodeError(ErrBadRequest, "PHONE_CODE_EMPTY")

	// ErrPhoneCodeExpired
	// | 400 | PHONE_CODE_EXPIRED | The phone code you provided has expired. |
	ErrPhoneCodeExpired = xerr.NewCodeError(ErrBadRequest, "PHONE_CODE_EXPIRED")

	// ErrPhoneCodeHashEmpty
	// | 400 | PHONE_CODE_HASH_EMPTY | phone_code_hash is missing. |
	ErrPhoneCodeHashEmpty = xerr.NewCodeError(ErrBadRequest, "PHONE_CODE_HASH_EMPTY")

	// ErrPhoneCodeInvalid
	// | 400 | PHONE_CODE_INVALID | The provided phone code is invalid. |
	ErrPhoneCodeInvalid = xerr.NewCodeError(ErrBadRequest, "PHONE_CODE_INVALID")

	// ErrPhoneHashExpired
	// | 400 | PHONE_HASH_EXPIRED | An invalid or expired `phone_code_hash` was provided. |
	ErrPhoneHashExpired = xerr.NewCodeError(ErrBadRequest, "PHONE_HASH_EXPIRED")

	// ErrPhoneNotOccupied
	// | 400 | PHONE_NOT_OCCUPIED | No user is associated to the specified phone number. |
	ErrPhoneNotOccupied = xerr.NewCodeError(ErrBadRequest, "PHONE_NOT_OCCUPIED")

	// ErrPhoneNumberAppSignupForbidden
	// | 400 | PHONE_NUMBER_APP_SIGNUP_FORBIDDEN | You can't sign up using this app. |
	ErrPhoneNumberAppSignupForbidden = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_APP_SIGNUP_FORBIDDEN")

	// ErrPhoneNumberBanned
	// | 400 | PHONE_NUMBER_BANNED | The provided phone number is banned from telegram. |
	ErrPhoneNumberBanned = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_BANNED")

	// ErrPhoneNumberFlood
	// | 400 | PHONE_NUMBER_FLOOD | You asked for the code too many times. |
	ErrPhoneNumberFlood = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_FLOOD")

	// ErrPhoneNumberOccupied
	// | 400 | PHONE_NUMBER_OCCUPIED | The phone number is already in use. |
	ErrPhoneNumberOccupied = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_OCCUPIED")

	// ErrPhoneNumberUnoccupied
	// | 400 | PHONE_NUMBER_UNOCCUPIED | The phone number is not yet being used. |
	ErrPhoneNumberUnoccupied = xerr.NewCodeError(ErrBadRequest, "PHONE_NUMBER_UNOCCUPIED")

	// ErrPhonePasswordProtected
	// | 400 | PHONE_PASSWORD_PROTECTED | This phone is password protected. |
	ErrPhonePasswordProtected = xerr.NewCodeError(ErrBadRequest, "PHONE_PASSWORD_PROTECTED")

	// ErrPhotoContentTypeInvalid
	// | 400 | PHOTO_CONTENT_TYPE_INVALID | Photo mime-type invalid. |
	ErrPhotoContentTypeInvalid = xerr.NewCodeError(ErrBadRequest, "PHOTO_CONTENT_TYPE_INVALID")

	// ErrPhotoContentUrlEmpty
	// | 400 | PHOTO_CONTENT_URL_EMPTY | Photo URL invalid. |
	ErrPhotoContentUrlEmpty = xerr.NewCodeError(ErrBadRequest, "PHOTO_CONTENT_URL_EMPTY")

	// ErrPhotoCropFileMissing
	// | 400 | PHOTO_CROP_FILE_MISSING | Photo crop file missing. |
	ErrPhotoCropFileMissing = xerr.NewCodeError(ErrBadRequest, "PHOTO_CROP_FILE_MISSING")

	// ErrPhotoCropSizeSmall
	// | 400 | PHOTO_CROP_SIZE_SMALL | Photo is too small. |
	ErrPhotoCropSizeSmall = xerr.NewCodeError(ErrBadRequest, "PHOTO_CROP_SIZE_SMALL")

	// ErrPhotoExtInvalid
	// | 400 | PHOTO_EXT_INVALID | The extension of the photo is invalid. |
	ErrPhotoExtInvalid = xerr.NewCodeError(ErrBadRequest, "PHOTO_EXT_INVALID")

	// ErrPhotoFileMissing
	// | 400 | PHOTO_FILE_MISSING | Profile photo file missing. |
	ErrPhotoFileMissing = xerr.NewCodeError(ErrBadRequest, "PHOTO_FILE_MISSING")

	// ErrPhotoIdInvalid
	// | 400 | PHOTO_ID_INVALID | Photo ID invalid. |
	ErrPhotoIdInvalid = xerr.NewCodeError(ErrBadRequest, "PHOTO_ID_INVALID")

	// ErrPhotoInvalid
	// | 400 | PHOTO_INVALID | Photo invalid. |
	ErrPhotoInvalid = xerr.NewCodeError(ErrBadRequest, "PHOTO_INVALID")

	// ErrPhotoInvalidDimensions
	// | 400 | PHOTO_INVALID_DIMENSIONS | The photo dimensions are invalid. |
	ErrPhotoInvalidDimensions = xerr.NewCodeError(ErrBadRequest, "PHOTO_INVALID_DIMENSIONS")

	// ErrPhotoSaveFileInvalid
	// | 400 | PHOTO_SAVE_FILE_INVALID | Internal issues, try again later. |
	ErrPhotoSaveFileInvalid = xerr.NewCodeError(ErrBadRequest, "PHOTO_SAVE_FILE_INVALID")

	// ErrPhotoThumbUrlEmpty
	// | 400 | PHOTO_THUMB_URL_EMPTY | Photo thumbnail URL is empty. |
	ErrPhotoThumbUrlEmpty = xerr.NewCodeError(ErrBadRequest, "PHOTO_THUMB_URL_EMPTY")

	// ErrPinRestricted
	// | 400 | PIN_RESTRICTED | You can't pin messages. |
	ErrPinRestricted = xerr.NewCodeError(ErrBadRequest, "PIN_RESTRICTED")

	// ErrPinnedDialogsTooMuch
	// | 400 | PINNED_DIALOGS_TOO_MUCH | Too many pinned dialogs. |
	ErrPinnedDialogsTooMuch = xerr.NewCodeError(ErrBadRequest, "PINNED_DIALOGS_TOO_MUCH")

	// ErrPinnedTooMuch
	// | 400 | PINNED_TOO_MUCH | There are too many pinned topics, unpin some first. |
	ErrPinnedTooMuch = xerr.NewCodeError(ErrBadRequest, "PINNED_TOO_MUCH")

	// ErrPollAnswerInvalid
	// | 400 | POLL_ANSWER_INVALID | One of the poll answers is not acceptable. |
	ErrPollAnswerInvalid = xerr.NewCodeError(ErrBadRequest, "POLL_ANSWER_INVALID")

	// ErrPollAnswersInvalid
	// | 400 | POLL_ANSWERS_INVALID | Invalid poll answers were provided. |
	ErrPollAnswersInvalid = xerr.NewCodeError(ErrBadRequest, "POLL_ANSWERS_INVALID")

	// ErrPollOptionDuplicate
	// | 400 | POLL_OPTION_DUPLICATE | Duplicate poll options provided. |
	ErrPollOptionDuplicate = xerr.NewCodeError(ErrBadRequest, "POLL_OPTION_DUPLICATE")

	// ErrPollOptionInvalid
	// | 400 | POLL_OPTION_INVALID | Invalid poll option provided. |
	ErrPollOptionInvalid = xerr.NewCodeError(ErrBadRequest, "POLL_OPTION_INVALID")

	// ErrPollQuestionInvalid
	// | 400 | POLL_QUESTION_INVALID | One of the poll questions is not acceptable. |
	ErrPollQuestionInvalid = xerr.NewCodeError(ErrBadRequest, "POLL_QUESTION_INVALID")

	// ErrPricingChatInvalid
	// | 400 | PRICING_CHAT_INVALID | The pricing for the [subscription](https://core.telegram.org/api/subscriptions) is invalid, the maximum price is specified in the [`stars_subscription_amount_max` config key &raquo;](https://core.telegram.org/api/config#stars-subscription-amount-max). |
	ErrPricingChatInvalid = xerr.NewCodeError(ErrBadRequest, "PRICING_CHAT_INVALID")

	// ErrPrivacyKeyInvalid
	// | 400 | PRIVACY_KEY_INVALID | The privacy key is invalid. |
	ErrPrivacyKeyInvalid = xerr.NewCodeError(ErrBadRequest, "PRIVACY_KEY_INVALID")

	// ErrPrivacyTooLong
	// | 400 | PRIVACY_TOO_LONG | Too many privacy rules were specified, the current limit is 1000. |
	ErrPrivacyTooLong = xerr.NewCodeError(ErrBadRequest, "PRIVACY_TOO_LONG")

	// ErrPrivacyValueInvalid
	// | 400 | PRIVACY_VALUE_INVALID | The specified privacy rule combination is invalid. |
	ErrPrivacyValueInvalid = xerr.NewCodeError(ErrBadRequest, "PRIVACY_VALUE_INVALID")

	// ErrPublicKeyRequired
	// | 400 | PUBLIC_KEY_REQUIRED | A public key is required. |
	ErrPublicKeyRequired = xerr.NewCodeError(ErrBadRequest, "PUBLIC_KEY_REQUIRED")

	// ErrPurposeInvalid
	// | 400 | PURPOSE_INVALID | The specified payment purpose is invalid. |
	ErrPurposeInvalid = xerr.NewCodeError(ErrBadRequest, "PURPOSE_INVALID")

	// ErrQueryIdEmpty
	// | 400 | QUERY_ID_EMPTY | The query ID is empty. |
	ErrQueryIdEmpty = xerr.NewCodeError(ErrBadRequest, "QUERY_ID_EMPTY")

	// ErrQueryIdInvalid
	// | 400 | QUERY_ID_INVALID | The query ID is invalid. |
	ErrQueryIdInvalid = xerr.NewCodeError(ErrBadRequest, "QUERY_ID_INVALID")

	// ErrQueryTooShort
	// | 400 | QUERY_TOO_SHORT | The query string is too short. |
	ErrQueryTooShort = xerr.NewCodeError(ErrBadRequest, "QUERY_TOO_SHORT")

	// ErrQuickRepliesBotNotAllowed
	// | 400 | QUICK_REPLIES_BOT_NOT_ALLOWED | [Quick replies](https://core.telegram.org/api/business#quick-reply-shortcuts) cannot be used by bots. |
	ErrQuickRepliesBotNotAllowed = xerr.NewCodeError(ErrBadRequest, "QUICK_REPLIES_BOT_NOT_ALLOWED")

	// ErrQuickRepliesTooMuch
	// | 400 | QUICK_REPLIES_TOO_MUCH | A maximum of [appConfig.`quick_replies_limit`](https://core.telegram.org/api/config#quick-replies-limit) shortcuts may be created, the limit was reached. |
	ErrQuickRepliesTooMuch = xerr.NewCodeError(ErrBadRequest, "QUICK_REPLIES_TOO_MUCH")

	// ErrQuizAnswerMissing
	// | 400 | QUIZ_ANSWER_MISSING | You can forward a quiz while hiding the original author only after choosing an option in the quiz. |
	ErrQuizAnswerMissing = xerr.NewCodeError(ErrBadRequest, "QUIZ_ANSWER_MISSING")

	// ErrQuizCorrectAnswerInvalid
	// | 400 | QUIZ_CORRECT_ANSWER_INVALID | An invalid value was provided to the correct_answers field. |
	ErrQuizCorrectAnswerInvalid = xerr.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWER_INVALID")

	// ErrQuizCorrectAnswersEmpty
	// | 400 | QUIZ_CORRECT_ANSWERS_EMPTY | No correct quiz answer was specified. |
	ErrQuizCorrectAnswersEmpty = xerr.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_EMPTY")

	// ErrQuizCorrectAnswersTooMuch
	// | 400 | QUIZ_CORRECT_ANSWERS_TOO_MUCH | You specified too many correct answers in a quiz, quizzes can only have one right answer! |
	ErrQuizCorrectAnswersTooMuch = xerr.NewCodeError(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_TOO_MUCH")

	// ErrQuizMultipleInvalid
	// | 400 | QUIZ_MULTIPLE_INVALID | Quizzes can't have the multiple_choice flag set! |
	ErrQuizMultipleInvalid = xerr.NewCodeError(ErrBadRequest, "QUIZ_MULTIPLE_INVALID")

	// ErrQuoteTextInvalid
	// | 400 | QUOTE_TEXT_INVALID | The specified `reply_to`.`quote_text` field is invalid. |
	ErrQuoteTextInvalid = xerr.NewCodeError(ErrBadRequest, "QUOTE_TEXT_INVALID")

	// ErrRaiseHandForbidden
	// | 400 | RAISE_HAND_FORBIDDEN | You cannot raise your hand. |
	ErrRaiseHandForbidden = xerr.NewCodeError(ErrBadRequest, "RAISE_HAND_FORBIDDEN")

	// ErrRandomIdEmpty
	// | 400 | RANDOM_ID_EMPTY | Random ID empty. |
	ErrRandomIdEmpty = xerr.NewCodeError(ErrBadRequest, "RANDOM_ID_EMPTY")

	// ErrRandomIdExpired
	// | 400 | RANDOM_ID_EXPIRED | The specified `random_id` was expired (most likely it didn't follow the required `uint64_t random_id = (time() << 32) | ((uint64_t)random_uint32_t())` format, or the specified time is too far in the past). |
	ErrRandomIdExpired = xerr.NewCodeError(ErrBadRequest, "RANDOM_ID_EXPIRED")

	// ErrRandomIdInvalid
	// | 400 | RANDOM_ID_INVALID | A provided random ID is invalid. |
	ErrRandomIdInvalid = xerr.NewCodeError(ErrBadRequest, "RANDOM_ID_INVALID")

	// ErrRandomLengthInvalid
	// | 400 | RANDOM_LENGTH_INVALID | Random length invalid. |
	ErrRandomLengthInvalid = xerr.NewCodeError(ErrBadRequest, "RANDOM_LENGTH_INVALID")

	// ErrRangesInvalid
	// | 400 | RANGES_INVALID | Invalid range provided. |
	ErrRangesInvalid = xerr.NewCodeError(ErrBadRequest, "RANGES_INVALID")

	// ErrReactionEmpty
	// | 400 | REACTION_EMPTY | Empty reaction provided. |
	ErrReactionEmpty = xerr.NewCodeError(ErrBadRequest, "REACTION_EMPTY")

	// ErrReactionInvalid
	// | 400 | REACTION_INVALID | The specified reaction is invalid. |
	ErrReactionInvalid = xerr.NewCodeError(ErrBadRequest, "REACTION_INVALID")

	// ErrReactionsCountInvalid
	// | 400 | REACTIONS_COUNT_INVALID | The specified number of reactions is invalid. |
	ErrReactionsCountInvalid = xerr.NewCodeError(ErrBadRequest, "REACTIONS_COUNT_INVALID")

	// ErrReactionsTooMany
	// | 400 | REACTIONS_TOO_MANY | The message already has exactly `reactions_uniq_max` reaction emojis, you can't react with a new emoji, see [the docs for more info &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrReactionsTooMany = xerr.NewCodeError(ErrBadRequest, "REACTIONS_TOO_MANY")

	// ErrReceiptEmpty
	// | 400 | RECEIPT_EMPTY | The specified receipt is empty. |
	ErrReceiptEmpty = xerr.NewCodeError(ErrBadRequest, "RECEIPT_EMPTY")

	// ErrReplyMarkupBuyEmpty
	// | 400 | REPLY_MARKUP_BUY_EMPTY | Reply markup for buy button empty. |
	ErrReplyMarkupBuyEmpty = xerr.NewCodeError(ErrBadRequest, "REPLY_MARKUP_BUY_EMPTY")

	// ErrReplyMarkupGameEmpty
	// | 400 | REPLY_MARKUP_GAME_EMPTY | A game message is being edited, but the newly provided keyboard doesn't have a keyboardButtonGame button. |
	ErrReplyMarkupGameEmpty = xerr.NewCodeError(ErrBadRequest, "REPLY_MARKUP_GAME_EMPTY")

	// ErrReplyMarkupInvalid
	// | 400 | REPLY_MARKUP_INVALID | The provided reply markup is invalid. |
	ErrReplyMarkupInvalid = xerr.NewCodeError(ErrBadRequest, "REPLY_MARKUP_INVALID")

	// ErrReplyMarkupTooLong
	// | 400 | REPLY_MARKUP_TOO_LONG | The specified reply_markup is too long. |
	ErrReplyMarkupTooLong = xerr.NewCodeError(ErrBadRequest, "REPLY_MARKUP_TOO_LONG")

	// ErrReplyMessageIdInvalid
	// | 400 | REPLY_MESSAGE_ID_INVALID | The specified reply-to message ID is invalid. |
	ErrReplyMessageIdInvalid = xerr.NewCodeError(ErrBadRequest, "REPLY_MESSAGE_ID_INVALID")

	// ErrReplyMessagesTooMuch
	// | 400 | REPLY_MESSAGES_TOO_MUCH | Each shortcut can contain a maximum of [appConfig.`quick_reply_messages_limit`](https://core.telegram.org/api/config#quick-reply-messages-limit) messages, the limit was reached. |
	ErrReplyMessagesTooMuch = xerr.NewCodeError(ErrBadRequest, "REPLY_MESSAGES_TOO_MUCH")

	// ErrReplyToInvalid
	// | 400 | REPLY_TO_INVALID | The specified `reply_to` field is invalid. |
	ErrReplyToInvalid = xerr.NewCodeError(ErrBadRequest, "REPLY_TO_INVALID")

	// ErrReplyToMonoforumPeerInvalid
	// | 400 | REPLY_TO_MONOFORUM_PEER_INVALID | The specified inputReplyToMonoForum.monoforum_peer_id is invalid. |
	ErrReplyToMonoforumPeerInvalid = xerr.NewCodeError(ErrBadRequest, "REPLY_TO_MONOFORUM_PEER_INVALID")

	// ErrReplyToUserInvalid
	// | 400 | REPLY_TO_USER_INVALID | The replied-to user is invalid. |
	ErrReplyToUserInvalid = xerr.NewCodeError(ErrBadRequest, "REPLY_TO_USER_INVALID")

	// ErrRequestTokenInvalid
	// | 400 | REQUEST_TOKEN_INVALID | The master DC did not accept the `request_token` from the CDN DC. Continue downloading the file from the master DC using upload.getFile. |
	ErrRequestTokenInvalid = xerr.NewCodeError(ErrBadRequest, "REQUEST_TOKEN_INVALID")

	// ErrResetRequestMissing
	// | 400 | RESET_REQUEST_MISSING | No password reset is in progress. |
	ErrResetRequestMissing = xerr.NewCodeError(ErrBadRequest, "RESET_REQUEST_MISSING")

	// ErrResultIdDuplicate
	// | 400 | RESULT_ID_DUPLICATE | You provided a duplicate result ID. |
	ErrResultIdDuplicate = xerr.NewCodeError(ErrBadRequest, "RESULT_ID_DUPLICATE")

	// ErrResultIdEmpty
	// | 400 | RESULT_ID_EMPTY | Result ID empty. |
	ErrResultIdEmpty = xerr.NewCodeError(ErrBadRequest, "RESULT_ID_EMPTY")

	// ErrResultIdInvalid
	// | 400 | RESULT_ID_INVALID | One of the specified result IDs is invalid. |
	ErrResultIdInvalid = xerr.NewCodeError(ErrBadRequest, "RESULT_ID_INVALID")

	// ErrResultTypeInvalid
	// | 400 | RESULT_TYPE_INVALID | Result type invalid. |
	ErrResultTypeInvalid = xerr.NewCodeError(ErrBadRequest, "RESULT_TYPE_INVALID")

	// ErrResultsTooMuch
	// | 400 | RESULTS_TOO_MUCH | Too many results were provided. |
	ErrResultsTooMuch = xerr.NewCodeError(ErrBadRequest, "RESULTS_TOO_MUCH")

	// ErrRevoteNotAllowed
	// | 400 | REVOTE_NOT_ALLOWED | You cannot change your vote. |
	ErrRevoteNotAllowed = xerr.NewCodeError(ErrBadRequest, "REVOTE_NOT_ALLOWED")

	// ErrRightsNotModified
	// | 400 | RIGHTS_NOT_MODIFIED | The new admin rights are equal to the old rights, no change was made. |
	ErrRightsNotModified = xerr.NewCodeError(ErrBadRequest, "RIGHTS_NOT_MODIFIED")

	// ErrRingtoneInvalid
	// | 400 | RINGTONE_INVALID | The specified ringtone is invalid. |
	ErrRingtoneInvalid = xerr.NewCodeError(ErrBadRequest, "RINGTONE_INVALID")

	// ErrRingtoneMimeInvalid
	// | 400 | RINGTONE_MIME_INVALID | The MIME type for the ringtone is invalid. |
	ErrRingtoneMimeInvalid = xerr.NewCodeError(ErrBadRequest, "RINGTONE_MIME_INVALID")

	// ErrRsaDecryptFailed
	// | 400 | RSA_DECRYPT_FAILED | Internal RSA decryption failed. |
	ErrRsaDecryptFailed = xerr.NewCodeError(ErrBadRequest, "RSA_DECRYPT_FAILED")

	// ErrSavedIdEmpty
	// | 400 | SAVED_ID_EMPTY | The passed inputSavedStarGiftChat.saved_id is empty. |
	ErrSavedIdEmpty = xerr.NewCodeError(ErrBadRequest, "SAVED_ID_EMPTY")

	// ErrScheduleBotNotAllowed
	// | 400 | SCHEDULE_BOT_NOT_ALLOWED | Bots cannot schedule messages. |
	ErrScheduleBotNotAllowed = xerr.NewCodeError(ErrBadRequest, "SCHEDULE_BOT_NOT_ALLOWED")

	// ErrScheduleDateInvalid
	// | 400 | SCHEDULE_DATE_INVALID | Invalid schedule date provided. |
	ErrScheduleDateInvalid = xerr.NewCodeError(ErrBadRequest, "SCHEDULE_DATE_INVALID")

	// ErrScheduleDateTooLate
	// | 400 | SCHEDULE_DATE_TOO_LATE | You can't schedule a message this far in the future. |
	ErrScheduleDateTooLate = xerr.NewCodeError(ErrBadRequest, "SCHEDULE_DATE_TOO_LATE")

	// ErrScheduleStatusPrivate
	// | 400 | SCHEDULE_STATUS_PRIVATE | Can't schedule until user is online, if the user's last seen timestamp is hidden by their privacy settings. |
	ErrScheduleStatusPrivate = xerr.NewCodeError(ErrBadRequest, "SCHEDULE_STATUS_PRIVATE")

	// ErrScheduleTooMuch
	// | 400 | SCHEDULE_TOO_MUCH | There are too many scheduled messages. |
	ErrScheduleTooMuch = xerr.NewCodeError(ErrBadRequest, "SCHEDULE_TOO_MUCH")

	// ErrScoreInvalid
	// | 400 | SCORE_INVALID | The specified game score is invalid. |
	ErrScoreInvalid = xerr.NewCodeError(ErrBadRequest, "SCORE_INVALID")

	// ErrSearchQueryEmpty
	// | 400 | SEARCH_QUERY_EMPTY | The search query is empty. |
	ErrSearchQueryEmpty = xerr.NewCodeError(ErrBadRequest, "SEARCH_QUERY_EMPTY")

	// ErrSearchWithLinkNotSupported
	// | 400 | SEARCH_WITH_LINK_NOT_SUPPORTED | You cannot provide a search query and an invite link at the same time. |
	ErrSearchWithLinkNotSupported = xerr.NewCodeError(ErrBadRequest, "SEARCH_WITH_LINK_NOT_SUPPORTED")

	// ErrSecondsInvalid
	// | 400 | SECONDS_INVALID | Invalid duration provided. |
	ErrSecondsInvalid = xerr.NewCodeError(ErrBadRequest, "SECONDS_INVALID")

	// ErrSecureSecretRequired
	// | 400 | SECURE_SECRET_REQUIRED | A secure secret is required. |
	ErrSecureSecretRequired = xerr.NewCodeError(ErrBadRequest, "SECURE_SECRET_REQUIRED")

	// ErrSelfDeleteRestricted
	// | 400 | SELF_DELETE_RESTRICTED | Business bots can't delete messages just for the user, `revoke` **must** be set. |
	ErrSelfDeleteRestricted = xerr.NewCodeError(ErrBadRequest, "SELF_DELETE_RESTRICTED")

	// ErrSendAsPeerInvalid
	// | 400 | SEND_AS_PEER_INVALID | You can't send messages as the specified peer. |
	ErrSendAsPeerInvalid = xerr.NewCodeError(ErrBadRequest, "SEND_AS_PEER_INVALID")

	// ErrSendMessageGameInvalid
	// | 400 | SEND_MESSAGE_GAME_INVALID | An inputBotInlineMessageGame can only be contained in an inputBotInlineResultGame, not in an inputBotInlineResult/inputBotInlineResultPhoto/etc. |
	ErrSendMessageGameInvalid = xerr.NewCodeError(ErrBadRequest, "SEND_MESSAGE_GAME_INVALID")

	// ErrSendMessageMediaInvalid
	// | 400 | SEND_MESSAGE_MEDIA_INVALID | Invalid media provided. |
	ErrSendMessageMediaInvalid = xerr.NewCodeError(ErrBadRequest, "SEND_MESSAGE_MEDIA_INVALID")

	// ErrSendMessageTypeInvalid
	// | 400 | SEND_MESSAGE_TYPE_INVALID | The message type is invalid. |
	ErrSendMessageTypeInvalid = xerr.NewCodeError(ErrBadRequest, "SEND_MESSAGE_TYPE_INVALID")

	// ErrSessionTooFresh_%d
	// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
	// ErrSessionTooFresh_%d = xerr.NewCodeError(ErrBadRequest, "SESSION_TOO_FRESH_%d")

	// ErrSettingsInvalid
	// | 400 | SETTINGS_INVALID | Invalid settings were provided. |
	ErrSettingsInvalid = xerr.NewCodeError(ErrBadRequest, "SETTINGS_INVALID")

	// ErrSha256HashInvalid
	// | 400 | SHA256_HASH_INVALID | The provided SHA256 hash is invalid. |
	ErrSha256HashInvalid = xerr.NewCodeError(ErrBadRequest, "SHA256_HASH_INVALID")

	// ErrShortNameInvalid
	// | 400 | SHORT_NAME_INVALID | The specified short name is invalid. |
	ErrShortNameInvalid = xerr.NewCodeError(ErrBadRequest, "SHORT_NAME_INVALID")

	// ErrShortNameOccupied
	// | 400 | SHORT_NAME_OCCUPIED | The specified short name is already in use. |
	ErrShortNameOccupied = xerr.NewCodeError(ErrBadRequest, "SHORT_NAME_OCCUPIED")

	// ErrShortcutInvalid
	// | 400 | SHORTCUT_INVALID | The specified shortcut is invalid. |
	ErrShortcutInvalid = xerr.NewCodeError(ErrBadRequest, "SHORTCUT_INVALID")

	// ErrSlotsEmpty
	// | 400 | SLOTS_EMPTY | The specified slot list is empty. |
	ErrSlotsEmpty = xerr.NewCodeError(ErrBadRequest, "SLOTS_EMPTY")

	// ErrSlowmodeMultiMsgsDisabled
	// | 400 | SLOWMODE_MULTI_MSGS_DISABLED | Slowmode is enabled, you cannot forward multiple messages to this group. |
	ErrSlowmodeMultiMsgsDisabled = xerr.NewCodeError(ErrBadRequest, "SLOWMODE_MULTI_MSGS_DISABLED")

	// ErrSlugInvalid
	// | 400 | SLUG_INVALID | The specified invoice slug is invalid. |
	ErrSlugInvalid = xerr.NewCodeError(ErrBadRequest, "SLUG_INVALID")

	// ErrSmsCodeCreateFailed
	// | 400 | SMS_CODE_CREATE_FAILED | An error occurred while creating the SMS code. |
	ErrSmsCodeCreateFailed = xerr.NewCodeError(ErrBadRequest, "SMS_CODE_CREATE_FAILED")

	// ErrSmsjobIdInvalid
	// | 400 | SMSJOB_ID_INVALID | The specified job ID is invalid. |
	ErrSmsjobIdInvalid = xerr.NewCodeError(ErrBadRequest, "SMSJOB_ID_INVALID")

	// ErrSrpAInvalid
	// | 400 | SRP_A_INVALID | The specified inputCheckPasswordSRP.A value is invalid. |
	ErrSrpAInvalid = xerr.NewCodeError(ErrBadRequest, "SRP_A_INVALID")

	// ErrSrpIdInvalid
	// | 400 | SRP_ID_INVALID | Invalid SRP ID provided. |
	ErrSrpIdInvalid = xerr.NewCodeError(ErrBadRequest, "SRP_ID_INVALID")

	// ErrSrpPasswordChanged
	// | 400 | SRP_PASSWORD_CHANGED | Password has changed. |
	ErrSrpPasswordChanged = xerr.NewCodeError(ErrBadRequest, "SRP_PASSWORD_CHANGED")

	// ErrStargiftAlreadyConverted
	// | 400 | STARGIFT_ALREADY_CONVERTED | The specified star gift was already converted to Stars. |
	ErrStargiftAlreadyConverted = xerr.NewCodeError(ErrBadRequest, "STARGIFT_ALREADY_CONVERTED")

	// ErrStargiftAlreadyRefunded
	// | 400 | STARGIFT_ALREADY_REFUNDED | The specified star gift was already refunded. |
	ErrStargiftAlreadyRefunded = xerr.NewCodeError(ErrBadRequest, "STARGIFT_ALREADY_REFUNDED")

	// ErrStargiftAlreadyUpgraded
	// | 400 | STARGIFT_ALREADY_UPGRADED | The specified gift was already upgraded to a collectible gift. |
	ErrStargiftAlreadyUpgraded = xerr.NewCodeError(ErrBadRequest, "STARGIFT_ALREADY_UPGRADED")

	// ErrStargiftInvalid
	// | 400 | STARGIFT_INVALID | The passed gift is invalid. |
	ErrStargiftInvalid = xerr.NewCodeError(ErrBadRequest, "STARGIFT_INVALID")

	// ErrStargiftNotFound
	// | 400 | STARGIFT_NOT_FOUND | The specified gift was not found. |
	ErrStargiftNotFound = xerr.NewCodeError(ErrBadRequest, "STARGIFT_NOT_FOUND")

	// ErrStargiftOwnerInvalid
	// | 400 | STARGIFT_OWNER_INVALID | You cannot transfer or sell a gift owned by another user. |
	ErrStargiftOwnerInvalid = xerr.NewCodeError(ErrBadRequest, "STARGIFT_OWNER_INVALID")

	// ErrStargiftPeerInvalid
	// | 400 | STARGIFT_PEER_INVALID | The specified inputSavedStarGiftChat.peer is invalid. |
	ErrStargiftPeerInvalid = xerr.NewCodeError(ErrBadRequest, "STARGIFT_PEER_INVALID")

	// ErrStargiftResellCurrencyNotAllowed
	// | 400 | STARGIFT_RESELL_CURRENCY_NOT_ALLOWED | You can't buy the gift using the specified currency (i.e. trying to pay in Stars for TON gifts). |
	ErrStargiftResellCurrencyNotAllowed = xerr.NewCodeError(ErrBadRequest, "STARGIFT_RESELL_CURRENCY_NOT_ALLOWED")

	// ErrStargiftSlugInvalid
	// | 400 | STARGIFT_SLUG_INVALID | The specified gift slug is invalid. |
	ErrStargiftSlugInvalid = xerr.NewCodeError(ErrBadRequest, "STARGIFT_SLUG_INVALID")

	// ErrStargiftTransferTooEarly_%d
	// | 400 | STARGIFT_TRANSFER_TOO_EARLY_%d | You cannot transfer this gift yet, wait %d seconds. |
	// ErrStargiftTransferTooEarly_%d = xerr.NewCodeError(ErrBadRequest, "STARGIFT_TRANSFER_TOO_EARLY_%d")

	// ErrStargiftUpgradeUnavailable
	// | 400 | STARGIFT_UPGRADE_UNAVAILABLE | A received gift can only be upgraded to a collectible gift if the [messageActionStarGift](https://core.telegram.org/constructor/messageActionStarGift)/[savedStarGift](https://core.telegram.org/constructor/savedStarGift).`can_upgrade` flag is set. |
	ErrStargiftUpgradeUnavailable = xerr.NewCodeError(ErrBadRequest, "STARGIFT_UPGRADE_UNAVAILABLE")

	// ErrStargiftUsageLimited
	// | 400 | STARGIFT_USAGE_LIMITED | The gift is sold out. |
	ErrStargiftUsageLimited = xerr.NewCodeError(ErrBadRequest, "STARGIFT_USAGE_LIMITED")

	// ErrStargiftUserUsageLimited
	// | 400 | STARGIFT_USER_USAGE_LIMITED | You've reached the starGift.limited_per_user limit, you can't buy any more gifts of this type. |
	ErrStargiftUserUsageLimited = xerr.NewCodeError(ErrBadRequest, "STARGIFT_USER_USAGE_LIMITED")

	// ErrStarrefAwaitingEnd
	// | 400 | STARREF_AWAITING_END | The previous referral program was terminated less than 24 hours ago: further changes can be made after the date specified in userFull.starref_program.end_date. |
	ErrStarrefAwaitingEnd = xerr.NewCodeError(ErrBadRequest, "STARREF_AWAITING_END")

	// ErrStarrefExpired
	// | 400 | STARREF_EXPIRED | The specified referral link is invalid. |
	ErrStarrefExpired = xerr.NewCodeError(ErrBadRequest, "STARREF_EXPIRED")

	// ErrStarrefHashRevoked
	// | 400 | STARREF_HASH_REVOKED | The specified affiliate link was already revoked. |
	ErrStarrefHashRevoked = xerr.NewCodeError(ErrBadRequest, "STARREF_HASH_REVOKED")

	// ErrStarrefPermilleInvalid
	// | 400 | STARREF_PERMILLE_INVALID | The specified commission_permille is invalid: the minimum and maximum values for this parameter are contained in the [starref_min_commission_permille](https://core.telegram.org/api/config#starref-min-commission-permille) and [starref_max_commission_permille](https://core.telegram.org/api/config#starref-max-commission-permille) client configuration parameters. |
	ErrStarrefPermilleInvalid = xerr.NewCodeError(ErrBadRequest, "STARREF_PERMILLE_INVALID")

	// ErrStarrefPermilleTooLow
	// | 400 | STARREF_PERMILLE_TOO_LOW | The specified commission_permille is too low: the minimum and maximum values for this parameter are contained in the [starref_min_commission_permille](https://core.telegram.org/api/config#starref-min-commission-permille) and [starref_max_commission_permille](https://core.telegram.org/api/config#starref-max-commission-permille) client configuration parameters. |
	ErrStarrefPermilleTooLow = xerr.NewCodeError(ErrBadRequest, "STARREF_PERMILLE_TOO_LOW")

	// ErrStarsAmountInvalid
	// | 400 | STARS_AMOUNT_INVALID | The specified amount in stars is invalid. |
	ErrStarsAmountInvalid = xerr.NewCodeError(ErrBadRequest, "STARS_AMOUNT_INVALID")

	// ErrStarsInvoiceInvalid
	// | 400 | STARS_INVOICE_INVALID | The specified Telegram Star invoice is invalid. |
	ErrStarsInvoiceInvalid = xerr.NewCodeError(ErrBadRequest, "STARS_INVOICE_INVALID")

	// ErrStarsPaymentRequired
	// | 400 | STARS_PAYMENT_REQUIRED | To import this chat invite link, you must first [pay for the associated Telegram Star subscription &raquo;](https://core.telegram.org/api/subscriptions#channel-subscriptions). |
	ErrStarsPaymentRequired = xerr.NewCodeError(ErrBadRequest, "STARS_PAYMENT_REQUIRED")

	// ErrStartParamEmpty
	// | 400 | START_PARAM_EMPTY | The start parameter is empty. |
	ErrStartParamEmpty = xerr.NewCodeError(ErrBadRequest, "START_PARAM_EMPTY")

	// ErrStartParamInvalid
	// | 400 | START_PARAM_INVALID | Start parameter invalid. |
	ErrStartParamInvalid = xerr.NewCodeError(ErrBadRequest, "START_PARAM_INVALID")

	// ErrStartParamTooLong
	// | 400 | START_PARAM_TOO_LONG | Start parameter is too long. |
	ErrStartParamTooLong = xerr.NewCodeError(ErrBadRequest, "START_PARAM_TOO_LONG")

	// ErrStickerDocumentInvalid
	// | 400 | STICKER_DOCUMENT_INVALID | The specified sticker document is invalid. |
	ErrStickerDocumentInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_DOCUMENT_INVALID")

	// ErrStickerEmojiInvalid
	// | 400 | STICKER_EMOJI_INVALID | Sticker emoji invalid. |
	ErrStickerEmojiInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_EMOJI_INVALID")

	// ErrStickerFileInvalid
	// | 400 | STICKER_FILE_INVALID | Sticker file invalid. |
	ErrStickerFileInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_FILE_INVALID")

	// ErrStickerGifDimensions
	// | 400 | STICKER_GIF_DIMENSIONS | The specified video sticker has invalid dimensions. |
	ErrStickerGifDimensions = xerr.NewCodeError(ErrBadRequest, "STICKER_GIF_DIMENSIONS")

	// ErrStickerIdInvalid
	// | 400 | STICKER_ID_INVALID | The provided sticker ID is invalid. |
	ErrStickerIdInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_ID_INVALID")

	// ErrStickerInvalid
	// | 400 | STICKER_INVALID | The provided sticker is invalid. |
	ErrStickerInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_INVALID")

	// ErrStickerMimeInvalid
	// | 400 | STICKER_MIME_INVALID | The specified sticker MIME type is invalid. |
	ErrStickerMimeInvalid = xerr.NewCodeError(ErrBadRequest, "STICKER_MIME_INVALID")

	// ErrStickerPngDimensions
	// | 400 | STICKER_PNG_DIMENSIONS | Sticker png dimensions invalid. |
	ErrStickerPngDimensions = xerr.NewCodeError(ErrBadRequest, "STICKER_PNG_DIMENSIONS")

	// ErrStickerPngNopng
	// | 400 | STICKER_PNG_NOPNG | One of the specified stickers is not a valid PNG file. |
	ErrStickerPngNopng = xerr.NewCodeError(ErrBadRequest, "STICKER_PNG_NOPNG")

	// ErrStickerTgsNodoc
	// | 400 | STICKER_TGS_NODOC | You must send the animated sticker as a document. |
	ErrStickerTgsNodoc = xerr.NewCodeError(ErrBadRequest, "STICKER_TGS_NODOC")

	// ErrStickerTgsNotgs
	// | 400 | STICKER_TGS_NOTGS | Invalid TGS sticker provided. |
	ErrStickerTgsNotgs = xerr.NewCodeError(ErrBadRequest, "STICKER_TGS_NOTGS")

	// ErrStickerThumbPngNopng
	// | 400 | STICKER_THUMB_PNG_NOPNG | Incorrect stickerset thumb file provided, PNG / WEBP expected. |
	ErrStickerThumbPngNopng = xerr.NewCodeError(ErrBadRequest, "STICKER_THUMB_PNG_NOPNG")

	// ErrStickerThumbTgsNotgs
	// | 400 | STICKER_THUMB_TGS_NOTGS | Incorrect stickerset TGS thumb file provided. |
	ErrStickerThumbTgsNotgs = xerr.NewCodeError(ErrBadRequest, "STICKER_THUMB_TGS_NOTGS")

	// ErrStickerVideoBig
	// | 400 | STICKER_VIDEO_BIG | The specified video sticker is too big. |
	ErrStickerVideoBig = xerr.NewCodeError(ErrBadRequest, "STICKER_VIDEO_BIG")

	// ErrStickerVideoNodoc
	// | 400 | STICKER_VIDEO_NODOC | You must send the video sticker as a document. |
	ErrStickerVideoNodoc = xerr.NewCodeError(ErrBadRequest, "STICKER_VIDEO_NODOC")

	// ErrStickerVideoNowebm
	// | 400 | STICKER_VIDEO_NOWEBM | The specified video sticker is not in webm format. |
	ErrStickerVideoNowebm = xerr.NewCodeError(ErrBadRequest, "STICKER_VIDEO_NOWEBM")

	// ErrStickerpackStickersTooMuch
	// | 400 | STICKERPACK_STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickerpackStickersTooMuch = xerr.NewCodeError(ErrBadRequest, "STICKERPACK_STICKERS_TOO_MUCH")

	// ErrStickersEmpty
	// | 400 | STICKERS_EMPTY | No sticker provided. |
	ErrStickersEmpty = xerr.NewCodeError(ErrBadRequest, "STICKERS_EMPTY")

	// ErrStickersTooMuch
	// | 400 | STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickersTooMuch = xerr.NewCodeError(ErrBadRequest, "STICKERS_TOO_MUCH")

	// ErrStickersetNotModified
	// | 400 | STICKERSET_NOT_MODIFIED | The passed stickerset information is equal to the current information. |
	ErrStickersetNotModified = xerr.NewCodeError(ErrBadRequest, "STICKERSET_NOT_MODIFIED")

	// ErrStoriesNeverCreated
	// | 400 | STORIES_NEVER_CREATED | This peer hasn't ever posted any stories. |
	ErrStoriesNeverCreated = xerr.NewCodeError(ErrBadRequest, "STORIES_NEVER_CREATED")

	// ErrStoriesTooMuch
	// | 400 | STORIES_TOO_MUCH | You have hit the maximum active stories limit as specified by the [`story_expiring_limit_*` client configuration parameters](https://core.telegram.org/api/config#story-expiring-limit-default): you should buy a [Premium](https://core.telegram.org/api/premium) subscription, delete an active story, or wait for the oldest story to expire. |
	ErrStoriesTooMuch = xerr.NewCodeError(ErrBadRequest, "STORIES_TOO_MUCH")

	// ErrStoryIdEmpty
	// | 400 | STORY_ID_EMPTY | You specified no story IDs. |
	ErrStoryIdEmpty = xerr.NewCodeError(ErrBadRequest, "STORY_ID_EMPTY")

	// ErrStoryIdInvalid
	// | 400 | STORY_ID_INVALID | The specified story ID is invalid. |
	ErrStoryIdInvalid = xerr.NewCodeError(ErrBadRequest, "STORY_ID_INVALID")

	// ErrStoryNotModified
	// | 400 | STORY_NOT_MODIFIED | The new story information you passed is equal to the previous story information, thus it wasn't modified. |
	ErrStoryNotModified = xerr.NewCodeError(ErrBadRequest, "STORY_NOT_MODIFIED")

	// ErrStoryPeriodInvalid
	// | 400 | STORY_PERIOD_INVALID | The specified story period is invalid for this account. |
	ErrStoryPeriodInvalid = xerr.NewCodeError(ErrBadRequest, "STORY_PERIOD_INVALID")

	// ErrStorySendFloodMonthly_%d
	// | 400 | STORY_SEND_FLOOD_MONTHLY_%d | You've hit the monthly story limit as specified by the [`stories_sent_monthly_limit_*` client configuration parameters](https://core.telegram.org/api/config#stories-sent-monthly-limit-default): wait %d seconds before posting a new story. |
	// ErrStorySendFloodMonthly_%d = xerr.NewCodeError(ErrBadRequest, "STORY_SEND_FLOOD_MONTHLY_%d")

	// ErrStorySendFloodWeekly_%d
	// | 400 | STORY_SEND_FLOOD_WEEKLY_%d | You've hit the weekly story limit as specified by the [`stories_sent_weekly_limit_*` client configuration parameters](https://core.telegram.org/api/config#stories-sent-weekly-limit-default): wait for %d seconds before posting a new story. |
	// ErrStorySendFloodWeekly_%d = xerr.NewCodeError(ErrBadRequest, "STORY_SEND_FLOOD_WEEKLY_%d")

	// ErrSubscriptionExportMissing
	// | 400 | SUBSCRIPTION_EXPORT_MISSING | You cannot send a [bot subscription invoice](https://core.telegram.org/api/subscriptions#bot-subscriptions) directly, you may only create invoice links using [payments.exportInvoice](https://core.telegram.org/method/payments.exportInvoice). |
	ErrSubscriptionExportMissing = xerr.NewCodeError(ErrBadRequest, "SUBSCRIPTION_EXPORT_MISSING")

	// ErrSubscriptionIdInvalid
	// | 400 | SUBSCRIPTION_ID_INVALID | The specified subscription_id is invalid. |
	ErrSubscriptionIdInvalid = xerr.NewCodeError(ErrBadRequest, "SUBSCRIPTION_ID_INVALID")

	// ErrSubscriptionPeriodInvalid
	// | 400 | SUBSCRIPTION_PERIOD_INVALID | The specified subscription_pricing.period is invalid. |
	ErrSubscriptionPeriodInvalid = xerr.NewCodeError(ErrBadRequest, "SUBSCRIPTION_PERIOD_INVALID")

	// ErrSuggestedPostAmountInvalid
	// | 400 | SUGGESTED_POST_AMOUNT_INVALID | The specified price for the suggested post is invalid. |
	ErrSuggestedPostAmountInvalid = xerr.NewCodeError(ErrBadRequest, "SUGGESTED_POST_AMOUNT_INVALID")

	// ErrSuggestedPostPeerInvalid
	// | 400 | SUGGESTED_POST_PEER_INVALID | You cannot send suggested posts to non-[monoforum](https://core.telegram.org/api/monoforum) peers. |
	ErrSuggestedPostPeerInvalid = xerr.NewCodeError(ErrBadRequest, "SUGGESTED_POST_PEER_INVALID")

	// ErrSwitchPmTextEmpty
	// | 400 | SWITCH_PM_TEXT_EMPTY | The switch_pm.text field was empty. |
	ErrSwitchPmTextEmpty = xerr.NewCodeError(ErrBadRequest, "SWITCH_PM_TEXT_EMPTY")

	// ErrSwitchWebviewUrlInvalid
	// | 400 | SWITCH_WEBVIEW_URL_INVALID | The URL specified in switch_webview.url is invalid! |
	ErrSwitchWebviewUrlInvalid = xerr.NewCodeError(ErrBadRequest, "SWITCH_WEBVIEW_URL_INVALID")

	// ErrTakeoutInvalid
	// | 400 | TAKEOUT_INVALID | The specified takeout ID is invalid. |
	ErrTakeoutInvalid = xerr.NewCodeError(ErrBadRequest, "TAKEOUT_INVALID")

	// ErrTaskAlreadyExists
	// | 400 | TASK_ALREADY_EXISTS | An email reset was already requested. |
	ErrTaskAlreadyExists = xerr.NewCodeError(ErrBadRequest, "TASK_ALREADY_EXISTS")

	// ErrTempAuthKeyAlreadyBound
	// | 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another **perm_auth_key_id**. |
	ErrTempAuthKeyAlreadyBound = xerr.NewCodeError(ErrBadRequest, "TEMP_AUTH_KEY_ALREADY_BOUND")

	// ErrTempAuthKeyEmpty
	// | 400 | TEMP_AUTH_KEY_EMPTY | No temporary auth key provided. |
	ErrTempAuthKeyEmpty = xerr.NewCodeError(ErrBadRequest, "TEMP_AUTH_KEY_EMPTY")

	// ErrTermsUrlInvalid
	// | 400 | TERMS_URL_INVALID | The specified [invoice](https://core.telegram.org/constructor/invoice).`terms_url` is invalid. |
	ErrTermsUrlInvalid = xerr.NewCodeError(ErrBadRequest, "TERMS_URL_INVALID")

	// ErrThemeFileInvalid
	// | 400 | THEME_FILE_INVALID | Invalid theme file provided. |
	ErrThemeFileInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_FILE_INVALID")

	// ErrThemeFormatInvalid
	// | 400 | THEME_FORMAT_INVALID | Invalid theme format provided. |
	ErrThemeFormatInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_FORMAT_INVALID")

	// ErrThemeInvalid
	// | 400 | THEME_INVALID | Invalid theme provided. |
	ErrThemeInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_INVALID")

	// ErrThemeMimeInvalid
	// | 400 | THEME_MIME_INVALID | The theme's MIME type is invalid. |
	ErrThemeMimeInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_MIME_INVALID")

	// ErrThemeParamsInvalid
	// | 400 | THEME_PARAMS_INVALID | The specified `theme_params` field is invalid. |
	ErrThemeParamsInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_PARAMS_INVALID")

	// ErrThemeSlugInvalid
	// | 400 | THEME_SLUG_INVALID | The specified theme slug is invalid. |
	ErrThemeSlugInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_SLUG_INVALID")

	// ErrThemeTitleInvalid
	// | 400 | THEME_TITLE_INVALID | The specified theme title is invalid. |
	ErrThemeTitleInvalid = xerr.NewCodeError(ErrBadRequest, "THEME_TITLE_INVALID")

	// ErrTimezoneInvalid
	// | 400 | TIMEZONE_INVALID | The specified timezone does not exist. |
	ErrTimezoneInvalid = xerr.NewCodeError(ErrBadRequest, "TIMEZONE_INVALID")

	// ErrTitleInvalid
	// | 400 | TITLE_INVALID | The specified stickerpack title is invalid. |
	ErrTitleInvalid = xerr.NewCodeError(ErrBadRequest, "TITLE_INVALID")

	// ErrTmpPasswordDisabled
	// | 400 | TMP_PASSWORD_DISABLED | The temporary password is disabled. |
	ErrTmpPasswordDisabled = xerr.NewCodeError(ErrBadRequest, "TMP_PASSWORD_DISABLED")

	// ErrTmpPasswordInvalid
	// | 400 | TMP_PASSWORD_INVALID | The passed tmp_password is invalid. |
	ErrTmpPasswordInvalid = xerr.NewCodeError(ErrBadRequest, "TMP_PASSWORD_INVALID")

	// ErrToIdInvalid
	// | 400 | TO_ID_INVALID | The specified `to_id` of the passed inputInvoiceStarGiftResale or inputInvoiceStarGiftTransfer is invalid. |
	ErrToIdInvalid = xerr.NewCodeError(ErrBadRequest, "TO_ID_INVALID")

	// ErrToLangInvalid
	// | 400 | TO_LANG_INVALID | The specified destination language is invalid. |
	ErrToLangInvalid = xerr.NewCodeError(ErrBadRequest, "TO_LANG_INVALID")

	// ErrTodoItemDuplicate
	// | 400 | TODO_ITEM_DUPLICATE | Duplicate [checklist items](https://core.telegram.org/api/todo) detected. |
	ErrTodoItemDuplicate = xerr.NewCodeError(ErrBadRequest, "TODO_ITEM_DUPLICATE")

	// ErrTodoItemsEmpty
	// | 400 | TODO_ITEMS_EMPTY | A checklist was specified, but no [checklist items](https://core.telegram.org/api/todo) were passed. |
	ErrTodoItemsEmpty = xerr.NewCodeError(ErrBadRequest, "TODO_ITEMS_EMPTY")

	// ErrTodoNotModified
	// | 400 | TODO_NOT_MODIFIED | No todo items were specified, so no changes were made to the todo list. |
	ErrTodoNotModified = xerr.NewCodeError(ErrBadRequest, "TODO_NOT_MODIFIED")

	// ErrTokenEmpty
	// | 400 | TOKEN_EMPTY | The specified token is empty. |
	ErrTokenEmpty = xerr.NewCodeError(ErrBadRequest, "TOKEN_EMPTY")

	// ErrTokenInvalid
	// | 400 | TOKEN_INVALID | The provided token is invalid. |
	ErrTokenInvalid = xerr.NewCodeError(ErrBadRequest, "TOKEN_INVALID")

	// ErrTokenTypeInvalid
	// | 400 | TOKEN_TYPE_INVALID | The specified token type is invalid. |
	ErrTokenTypeInvalid = xerr.NewCodeError(ErrBadRequest, "TOKEN_TYPE_INVALID")

	// ErrTopicCloseSeparately
	// | 400 | TOPIC_CLOSE_SEPARATELY | The `close` flag cannot be provided together with any of the other flags. |
	ErrTopicCloseSeparately = xerr.NewCodeError(ErrBadRequest, "TOPIC_CLOSE_SEPARATELY")

	// ErrTopicHideSeparately
	// | 400 | TOPIC_HIDE_SEPARATELY | The `hide` flag cannot be provided together with any of the other flags. |
	ErrTopicHideSeparately = xerr.NewCodeError(ErrBadRequest, "TOPIC_HIDE_SEPARATELY")

	// ErrTopicIdInvalid
	// | 400 | TOPIC_ID_INVALID | The specified topic ID is invalid. |
	ErrTopicIdInvalid = xerr.NewCodeError(ErrBadRequest, "TOPIC_ID_INVALID")

	// ErrTopicNotModified
	// | 400 | TOPIC_NOT_MODIFIED | The updated topic info is equal to the current topic info, nothing was changed. |
	ErrTopicNotModified = xerr.NewCodeError(ErrBadRequest, "TOPIC_NOT_MODIFIED")

	// ErrTopicTitleEmpty
	// | 400 | TOPIC_TITLE_EMPTY | The specified topic title is empty. |
	ErrTopicTitleEmpty = xerr.NewCodeError(ErrBadRequest, "TOPIC_TITLE_EMPTY")

	// ErrTopicsEmpty
	// | 400 | TOPICS_EMPTY | You specified no topic IDs. |
	ErrTopicsEmpty = xerr.NewCodeError(ErrBadRequest, "TOPICS_EMPTY")

	// ErrTransactionIdInvalid
	// | 400 | TRANSACTION_ID_INVALID | The specified transaction ID is invalid. |
	ErrTransactionIdInvalid = xerr.NewCodeError(ErrBadRequest, "TRANSACTION_ID_INVALID")

	// ErrTranscriptionFailed
	// | 400 | TRANSCRIPTION_FAILED | Audio transcription failed. |
	ErrTranscriptionFailed = xerr.NewCodeError(ErrBadRequest, "TRANSCRIPTION_FAILED")

	// ErrTranslateReqQuotaExceeded
	// | 400 | TRANSLATE_REQ_QUOTA_EXCEEDED | Translation is currently unavailable due to a temporary server-side lack of resources. |
	ErrTranslateReqQuotaExceeded = xerr.NewCodeError(ErrBadRequest, "TRANSLATE_REQ_QUOTA_EXCEEDED")

	// ErrTtlDaysInvalid
	// | 400 | TTL_DAYS_INVALID | The provided TTL is invalid. |
	ErrTtlDaysInvalid = xerr.NewCodeError(ErrBadRequest, "TTL_DAYS_INVALID")

	// ErrTtlMediaInvalid
	// | 400 | TTL_MEDIA_INVALID | Invalid media Time To Live was provided. |
	ErrTtlMediaInvalid = xerr.NewCodeError(ErrBadRequest, "TTL_MEDIA_INVALID")

	// ErrTtlPeriodInvalid
	// | 400 | TTL_PERIOD_INVALID | The specified TTL period is invalid. |
	ErrTtlPeriodInvalid = xerr.NewCodeError(ErrBadRequest, "TTL_PERIOD_INVALID")

	// ErrTypesEmpty
	// | 400 | TYPES_EMPTY | No top peer type was provided. |
	ErrTypesEmpty = xerr.NewCodeError(ErrBadRequest, "TYPES_EMPTY")

	// ErrUnsupported
	// | 400 | UNSUPPORTED | `require_payment` cannot be *set* by users, only by monoforums: users must instead use the [inputPrivacyKeyNoPaidMessages](https://core.telegram.org/constructor/inputPrivacyKeyNoPaidMessages) privacy setting to remove a previously added exemption. |
	ErrUnsupported = xerr.NewCodeError(ErrBadRequest, "UNSUPPORTED")

	// ErrUntilDateInvalid
	// | 400 | UNTIL_DATE_INVALID | Invalid until date provided. |
	ErrUntilDateInvalid = xerr.NewCodeError(ErrBadRequest, "UNTIL_DATE_INVALID")

	// ErrUrlInvalid
	// | 400 | URL_INVALID | Invalid URL provided. |
	ErrUrlInvalid = xerr.NewCodeError(ErrBadRequest, "URL_INVALID")

	// ErrUsageLimitInvalid
	// | 400 | USAGE_LIMIT_INVALID | The specified usage limit is invalid. |
	ErrUsageLimitInvalid = xerr.NewCodeError(ErrBadRequest, "USAGE_LIMIT_INVALID")

	// ErrUserAdminInvalid
	// | 400 | USER_ADMIN_INVALID | You're not an admin. |
	ErrUserAdminInvalid = xerr.NewCodeError(ErrBadRequest, "USER_ADMIN_INVALID")

	// ErrUserAlreadyInvited
	// | 400 | USER_ALREADY_INVITED | You have already invited this user. |
	ErrUserAlreadyInvited = xerr.NewCodeError(ErrBadRequest, "USER_ALREADY_INVITED")

	// ErrUserAlreadyParticipant
	// | 400 | USER_ALREADY_PARTICIPANT | The user is already in the group. |
	ErrUserAlreadyParticipant = xerr.NewCodeError(ErrBadRequest, "USER_ALREADY_PARTICIPANT")

	// ErrUserBannedInChannel
	// | 400 | USER_BANNED_IN_CHANNEL | You're banned from sending messages in supergroups/channels. |
	ErrUserBannedInChannel = xerr.NewCodeError(ErrBadRequest, "USER_BANNED_IN_CHANNEL")

	// ErrUserBlocked
	// | 400 | USER_BLOCKED | User blocked. |
	ErrUserBlocked = xerr.NewCodeError(ErrBadRequest, "USER_BLOCKED")

	// ErrUserBot
	// | 400 | USER_BOT | Bots can only be admins in channels. |
	ErrUserBot = xerr.NewCodeError(ErrBadRequest, "USER_BOT")

	// ErrUserBotInvalid
	// | 400 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	ErrUserBotInvalid = xerr.NewCodeError(ErrBadRequest, "USER_BOT_INVALID")

	// ErrUserBotRequired
	// | 400 | USER_BOT_REQUIRED | This method can only be called by a bot. |
	ErrUserBotRequired = xerr.NewCodeError(ErrBadRequest, "USER_BOT_REQUIRED")

	// ErrUserCreator
	// | 400 | USER_CREATOR | For channels.editAdmin: you've tried to edit the admin rights of the owner, but you're not the owner; for channels.leaveChannel: you can't leave this channel, because you're its creator. |
	ErrUserCreator = xerr.NewCodeError(ErrBadRequest, "USER_CREATOR")

	// ErrUserGiftUnavailable
	// | 400 | USER_GIFT_UNAVAILABLE | Gifts are not available in the current region ([stars_gifts_enabled](https://core.telegram.org/api/config#stars-gifts-enabled) is equal to false). |
	ErrUserGiftUnavailable = xerr.NewCodeError(ErrBadRequest, "USER_GIFT_UNAVAILABLE")

	// ErrUserIdInvalid
	// | 400 | USER_ID_INVALID | The provided user ID is invalid. |
	ErrUserIdInvalid = xerr.NewCodeError(ErrBadRequest, "USER_ID_INVALID")

	// ErrUserIsBot
	// | 400 | USER_IS_BOT | Bots can't send messages to other bots. |
	ErrUserIsBot = xerr.NewCodeError(ErrBadRequest, "USER_IS_BOT")

	// ErrUserKicked
	// | 400 | USER_KICKED | This user was kicked from this supergroup/channel. |
	ErrUserKicked = xerr.NewCodeError(ErrBadRequest, "USER_KICKED")

	// ErrUserPublicMissing
	// | 400 | USER_PUBLIC_MISSING | Cannot generate a link to stories posted by a peer without a username. |
	ErrUserPublicMissing = xerr.NewCodeError(ErrBadRequest, "USER_PUBLIC_MISSING")

	// ErrUserVolumeInvalid
	// | 400 | USER_VOLUME_INVALID | The specified user volume is invalid. |
	ErrUserVolumeInvalid = xerr.NewCodeError(ErrBadRequest, "USER_VOLUME_INVALID")

	// ErrUsernameInvalid
	// | 400 | USERNAME_INVALID | The provided username is not valid. |
	ErrUsernameInvalid = xerr.NewCodeError(ErrBadRequest, "USERNAME_INVALID")

	// ErrUsernameNotModified
	// | 400 | USERNAME_NOT_MODIFIED | The username was not modified. |
	ErrUsernameNotModified = xerr.NewCodeError(ErrBadRequest, "USERNAME_NOT_MODIFIED")

	// ErrUsernameNotOccupied
	// | 400 | USERNAME_NOT_OCCUPIED | The provided username is not occupied. |
	ErrUsernameNotOccupied = xerr.NewCodeError(ErrBadRequest, "USERNAME_NOT_OCCUPIED")

	// ErrUsernameOccupied
	// | 400 | USERNAME_OCCUPIED | The provided username is already occupied. |
	ErrUsernameOccupied = xerr.NewCodeError(ErrBadRequest, "USERNAME_OCCUPIED")

	// ErrUsernamePurchaseAvailable
	// | 400 | USERNAME_PURCHASE_AVAILABLE | The specified username can be purchased on https://fragment.com. |
	ErrUsernamePurchaseAvailable = xerr.NewCodeError(ErrBadRequest, "USERNAME_PURCHASE_AVAILABLE")

	// ErrUsernamesActiveTooMuch
	// | 400 | USERNAMES_ACTIVE_TOO_MUCH | The maximum number of active usernames was reached. |
	ErrUsernamesActiveTooMuch = xerr.NewCodeError(ErrBadRequest, "USERNAMES_ACTIVE_TOO_MUCH")

	// ErrUsersTooFew
	// | 400 | USERS_TOO_FEW | Not enough users (to create a chat, for example). |
	ErrUsersTooFew = xerr.NewCodeError(ErrBadRequest, "USERS_TOO_FEW")

	// ErrUsersTooMuch
	// | 400 | USERS_TOO_MUCH | The maximum number of users has been exceeded (to create a chat, for example). |
	ErrUsersTooMuch = xerr.NewCodeError(ErrBadRequest, "USERS_TOO_MUCH")

	// ErrVenueIdInvalid
	// | 400 | VENUE_ID_INVALID | The specified venue ID is invalid. |
	ErrVenueIdInvalid = xerr.NewCodeError(ErrBadRequest, "VENUE_ID_INVALID")

	// ErrVideoContentTypeInvalid
	// | 400 | VIDEO_CONTENT_TYPE_INVALID | The video's content type is invalid. |
	ErrVideoContentTypeInvalid = xerr.NewCodeError(ErrBadRequest, "VIDEO_CONTENT_TYPE_INVALID")

	// ErrVideoFileInvalid
	// | 400 | VIDEO_FILE_INVALID | The specified video file is invalid. |
	ErrVideoFileInvalid = xerr.NewCodeError(ErrBadRequest, "VIDEO_FILE_INVALID")

	// ErrVideoPauseForbidden
	// | 400 | VIDEO_PAUSE_FORBIDDEN | You cannot pause the video stream. |
	ErrVideoPauseForbidden = xerr.NewCodeError(ErrBadRequest, "VIDEO_PAUSE_FORBIDDEN")

	// ErrVideoStopForbidden
	// | 400 | VIDEO_STOP_FORBIDDEN | You cannot stop the video stream. |
	ErrVideoStopForbidden = xerr.NewCodeError(ErrBadRequest, "VIDEO_STOP_FORBIDDEN")

	// ErrVideoTitleEmpty
	// | 400 | VIDEO_TITLE_EMPTY | The specified video title is empty. |
	ErrVideoTitleEmpty = xerr.NewCodeError(ErrBadRequest, "VIDEO_TITLE_EMPTY")

	// ErrWallpaperFileInvalid
	// | 400 | WALLPAPER_FILE_INVALID | The specified wallpaper file is invalid. |
	ErrWallpaperFileInvalid = xerr.NewCodeError(ErrBadRequest, "WALLPAPER_FILE_INVALID")

	// ErrWallpaperInvalid
	// | 400 | WALLPAPER_INVALID | The specified wallpaper is invalid. |
	ErrWallpaperInvalid = xerr.NewCodeError(ErrBadRequest, "WALLPAPER_INVALID")

	// ErrWallpaperMimeInvalid
	// | 400 | WALLPAPER_MIME_INVALID | The specified wallpaper MIME type is invalid. |
	ErrWallpaperMimeInvalid = xerr.NewCodeError(ErrBadRequest, "WALLPAPER_MIME_INVALID")

	// ErrWallpaperNotFound
	// | 400 | WALLPAPER_NOT_FOUND | The specified wallpaper could not be found. |
	ErrWallpaperNotFound = xerr.NewCodeError(ErrBadRequest, "WALLPAPER_NOT_FOUND")

	// ErrWcConvertUrlInvalid
	// | 400 | WC_CONVERT_URL_INVALID | WC convert URL invalid. |
	ErrWcConvertUrlInvalid = xerr.NewCodeError(ErrBadRequest, "WC_CONVERT_URL_INVALID")

	// ErrWebdocumentInvalid
	// | 400 | WEBDOCUMENT_INVALID | Invalid webdocument URL provided. |
	ErrWebdocumentInvalid = xerr.NewCodeError(ErrBadRequest, "WEBDOCUMENT_INVALID")

	// ErrWebdocumentMimeInvalid
	// | 400 | WEBDOCUMENT_MIME_INVALID | Invalid webdocument mime type provided. |
	ErrWebdocumentMimeInvalid = xerr.NewCodeError(ErrBadRequest, "WEBDOCUMENT_MIME_INVALID")

	// ErrWebdocumentSizeTooBig
	// | 400 | WEBDOCUMENT_SIZE_TOO_BIG | Webdocument is too big! |
	ErrWebdocumentSizeTooBig = xerr.NewCodeError(ErrBadRequest, "WEBDOCUMENT_SIZE_TOO_BIG")

	// ErrWebdocumentUrlEmpty
	// | 400 | WEBDOCUMENT_URL_EMPTY | The passed web document URL is empty. |
	ErrWebdocumentUrlEmpty = xerr.NewCodeError(ErrBadRequest, "WEBDOCUMENT_URL_EMPTY")

	// ErrWebdocumentUrlInvalid
	// | 400 | WEBDOCUMENT_URL_INVALID | The specified webdocument URL is invalid. |
	ErrWebdocumentUrlInvalid = xerr.NewCodeError(ErrBadRequest, "WEBDOCUMENT_URL_INVALID")

	// ErrWebpageCurlFailed
	// | 400 | WEBPAGE_CURL_FAILED | Failure while fetching the webpage with cURL. |
	ErrWebpageCurlFailed = xerr.NewCodeError(ErrBadRequest, "WEBPAGE_CURL_FAILED")

	// ErrWebpageMediaEmpty
	// | 400 | WEBPAGE_MEDIA_EMPTY | Webpage media empty. |
	ErrWebpageMediaEmpty = xerr.NewCodeError(ErrBadRequest, "WEBPAGE_MEDIA_EMPTY")

	// ErrWebpageNotFound
	// | 400 | WEBPAGE_NOT_FOUND | A preview for the specified webpage `url` could not be generated. |
	ErrWebpageNotFound = xerr.NewCodeError(ErrBadRequest, "WEBPAGE_NOT_FOUND")

	// ErrWebpageUrlInvalid
	// | 400 | WEBPAGE_URL_INVALID | The specified webpage `url` is invalid. |
	ErrWebpageUrlInvalid = xerr.NewCodeError(ErrBadRequest, "WEBPAGE_URL_INVALID")

	// ErrWebpushAuthInvalid
	// | 400 | WEBPUSH_AUTH_INVALID | The specified web push authentication secret is invalid. |
	ErrWebpushAuthInvalid = xerr.NewCodeError(ErrBadRequest, "WEBPUSH_AUTH_INVALID")

	// ErrWebpushKeyInvalid
	// | 400 | WEBPUSH_KEY_INVALID | The specified web push elliptic curve Diffie-Hellman public key is invalid. |
	ErrWebpushKeyInvalid = xerr.NewCodeError(ErrBadRequest, "WEBPUSH_KEY_INVALID")

	// ErrWebpushTokenInvalid
	// | 400 | WEBPUSH_TOKEN_INVALID | The specified web push token is invalid. |
	ErrWebpushTokenInvalid = xerr.NewCodeError(ErrBadRequest, "WEBPUSH_TOKEN_INVALID")

	// ErrYouBlockedUser
	// | 400 | YOU_BLOCKED_USER | You blocked this user. |
	ErrYouBlockedUser = xerr.NewCodeError(ErrBadRequest, "YOU_BLOCKED_USER")
)

// 401 UNAUTHORIZED
var (
	// ErrSessionPasswordNeeded
	// 401	SESSION_PASSWORD_NEEDED	The user has enabled 2FA, more steps are needed
	ErrSessionPasswordNeeded = xerr.NewCodeError(ErrUnauthorized, "SESSION_PASSWORD_NEEDED")

	// ErrAuthKeyUnregistered
	// AUTH_KEY_UNREGISTERED: The key is not registered in the system
	ErrAuthKeyUnregistered = xerr.NewCodeError(ErrUnauthorized, "AUTH_KEY_UNREGISTERED")

	// ErrAuthKeyInvalid
	// | 401 | AUTH_KEY_INVALID | Auth key invalid. |
	ErrAuthKeyInvalid = xerr.NewCodeError(ErrUnauthorized, "AUTH_KEY_INVALID")

	// ErrUserDeactivated
	// USER_DEACTIVATED: The user has been deleted/deactivated
	ErrUserDeactivated = xerr.NewCodeError(ErrUnauthorized, "USER_DEACTIVATED")

	// ErrSessionRevoked
	// SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	ErrSessionRevoked = xerr.NewCodeError(ErrUnauthorized, "SESSION_REVOKED")

	// ErrSessionExpired
	// SESSION_EXPIRED: The authorization has expired
	ErrSessionExpired = xerr.NewCodeError(ErrUnauthorized, "SESSION_EXPIRED")

	// ErrAuthKeyPermEmpty
	// | 401 | AUTH_KEY_PERM_EMPTY | The temporary auth key must be binded to the permanent auth key to use these methods. |
	ErrAuthKeyPermEmpty = xerr.NewCodeError(ErrUnauthorized, "AUTH_KEY_PERM_EMPTY")

	// ErrActiveUserRequired
	// ACTIVE_USER_REQUIRED	401	The method is only available to already activated users
	ErrActiveUserRequired = xerr.NewCodeError(ErrUnauthorized, "ACTIVE_USER_REQUIRED")
)

// 403 FORBIDDEN
var (

	// Err403ChatAdminRequired
	// | 403 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	Err403ChatAdminRequired = xerr.NewCodeError(ErrForbidden, "CHAT_ADMIN_REQUIRED")

	// Err403ChatSendInlineForbidden
	// | 403 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	Err403ChatSendInlineForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_INLINE_FORBIDDEN")

	// Err403GroupcallForbidden
	// | 403 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	Err403GroupcallForbidden = xerr.NewCodeError(ErrForbidden, "GROUPCALL_FORBIDDEN")

	// Err403NotEligible
	// | 403 | NOT_ELIGIBLE | The current user is not eligible to join the Peer-to-Peer Login Program. |
	Err403NotEligible = xerr.NewCodeError(ErrForbidden, "NOT_ELIGIBLE")

	// Err403ParticipantJoinMissing
	// | 403 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	Err403ParticipantJoinMissing = xerr.NewCodeError(ErrForbidden, "PARTICIPANT_JOIN_MISSING")

	// Err403PeerIdInvalid
	// | 403 | PEER_ID_INVALID | The provided peer id is invalid. |
	Err403PeerIdInvalid = xerr.NewCodeError(ErrForbidden, "PEER_ID_INVALID")

	// Err403PremiumAccountRequired
	// | 403 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	Err403PremiumAccountRequired = xerr.NewCodeError(ErrForbidden, "PREMIUM_ACCOUNT_REQUIRED")

	// Err403PrivacyPremiumRequired
	// | 403 | PRIVACY_PREMIUM_REQUIRED | You need a [Telegram Premium subscription](https://core.telegram.org/api/premium) to send a message to this user. |
	Err403PrivacyPremiumRequired = xerr.NewCodeError(ErrForbidden, "PRIVACY_PREMIUM_REQUIRED")

	// Err403TakeoutRequired
	// | 403 | TAKEOUT_REQUIRED | A [takeout](https://core.telegram.org/api/takeout) session needs to be initialized first, [see here &raquo; for more info](https://core.telegram.org/api/takeout). |
	Err403TakeoutRequired = xerr.NewCodeError(ErrForbidden, "TAKEOUT_REQUIRED")

	// Err403UserChannelsTooMuch
	// | 403 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	Err403UserChannelsTooMuch = xerr.NewCodeError(ErrForbidden, "USER_CHANNELS_TOO_MUCH")

	// Err403UserInvalid
	// | 403 | USER_INVALID | Invalid user provided. |
	Err403UserInvalid = xerr.NewCodeError(ErrForbidden, "USER_INVALID")

	// Err403UserIsBlocked
	// | 403 | USER_IS_BLOCKED | You were blocked by this user. |
	Err403UserIsBlocked = xerr.NewCodeError(ErrForbidden, "USER_IS_BLOCKED")

	// Err403UserNotMutualContact
	// | 403 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	Err403UserNotMutualContact = xerr.NewCodeError(ErrForbidden, "USER_NOT_MUTUAL_CONTACT")

	// Err403UserNotParticipant
	// | 403 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	Err403UserNotParticipant = xerr.NewCodeError(ErrForbidden, "USER_NOT_PARTICIPANT")

	// Err403UserRestricted
	// | 403 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	Err403UserRestricted = xerr.NewCodeError(ErrForbidden, "USER_RESTRICTED")

	// Err403VoiceMessagesForbidden
	// | 403 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	Err403VoiceMessagesForbidden = xerr.NewCodeError(ErrForbidden, "VOICE_MESSAGES_FORBIDDEN")

	// ErrAllowPaymentRequired_%d
	// | 403 | ALLOW_PAYMENT_REQUIRED_%d | This peer charges %d [Telegram Stars](https://core.telegram.org/api/stars) per message, but the `allow_paid_stars` was not set or its value is smaller than %d. |
	// ErrAllowPaymentRequired_%d = xerr.NewCodeError(ErrForbidden, "ALLOW_PAYMENT_REQUIRED_%d")

	// ErrAnonymousReactionsDisabled
	// | 403 | ANONYMOUS_REACTIONS_DISABLED | Sorry, anonymous administrators cannot leave reactions or participate in polls. |
	ErrAnonymousReactionsDisabled = xerr.NewCodeError(ErrForbidden, "ANONYMOUS_REACTIONS_DISABLED")

	// ErrBotAccessForbidden
	// | 403 | BOT_ACCESS_FORBIDDEN | The specified method *can* be used over a [business connection](https://core.telegram.org/api/bots/connected-business-bots) for some operations, but the specified query attempted an operation that is not allowed over a business connection. |
	ErrBotAccessForbidden = xerr.NewCodeError(ErrForbidden, "BOT_ACCESS_FORBIDDEN")

	// ErrBotVerifierForbidden
	// | 403 | BOT_VERIFIER_FORBIDDEN | This bot cannot assign [verification icons](https://core.telegram.org/api/bots/verification). |
	ErrBotVerifierForbidden = xerr.NewCodeError(ErrForbidden, "BOT_VERIFIER_FORBIDDEN")

	// ErrBroadcastForbidden
	// | 403 | BROADCAST_FORBIDDEN | Channel poll voters and reactions cannot be fetched to prevent deanonymization. |
	ErrBroadcastForbidden = xerr.NewCodeError(ErrForbidden, "BROADCAST_FORBIDDEN")

	// ErrChannelPublicGroupNa
	// | 403 | CHANNEL_PUBLIC_GROUP_NA | channel/supergroup not available. |
	ErrChannelPublicGroupNa = xerr.NewCodeError(ErrForbidden, "CHANNEL_PUBLIC_GROUP_NA")

	// ErrChatActionForbidden
	// | 403 | CHAT_ACTION_FORBIDDEN | You cannot execute this action. |
	ErrChatActionForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_ACTION_FORBIDDEN")

	// ErrChatAdminInviteRequired
	// | 403 | CHAT_ADMIN_INVITE_REQUIRED | You do not have the rights to do this. |
	ErrChatAdminInviteRequired = xerr.NewCodeError(ErrForbidden, "CHAT_ADMIN_INVITE_REQUIRED")

	// ErrChatForbidden
	// | 403 | CHAT_FORBIDDEN | This chat is not available to the current user. |
	ErrChatForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_FORBIDDEN")

	// ErrChatGuestSendForbidden
	// | 403 | CHAT_GUEST_SEND_FORBIDDEN | You join the discussion group before commenting, see [here &raquo;](https://core.telegram.org/api/discussion#requiring-users-to-join-the-group) for more info. |
	ErrChatGuestSendForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_GUEST_SEND_FORBIDDEN")

	// ErrChatSendAudiosForbidden
	// | 403 | CHAT_SEND_AUDIOS_FORBIDDEN | You can't send audio messages in this chat. |
	ErrChatSendAudiosForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_AUDIOS_FORBIDDEN")

	// ErrChatSendDocsForbidden
	// | 403 | CHAT_SEND_DOCS_FORBIDDEN | You can't send documents in this chat. |
	ErrChatSendDocsForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_DOCS_FORBIDDEN")

	// ErrChatSendGameForbidden
	// | 403 | CHAT_SEND_GAME_FORBIDDEN | You can't send a game to this chat. |
	ErrChatSendGameForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_GAME_FORBIDDEN")

	// ErrChatSendGifsForbidden
	// | 403 | CHAT_SEND_GIFS_FORBIDDEN | You can't send gifs in this chat. |
	ErrChatSendGifsForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_GIFS_FORBIDDEN")

	// ErrChatSendMediaForbidden
	// | 403 | CHAT_SEND_MEDIA_FORBIDDEN | You can't send media in this chat. |
	ErrChatSendMediaForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_MEDIA_FORBIDDEN")

	// ErrChatSendPhotosForbidden
	// | 403 | CHAT_SEND_PHOTOS_FORBIDDEN | You can't send photos in this chat. |
	ErrChatSendPhotosForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_PHOTOS_FORBIDDEN")

	// ErrChatSendPlainForbidden
	// | 403 | CHAT_SEND_PLAIN_FORBIDDEN | You can't send non-media (text) messages in this chat. |
	ErrChatSendPlainForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_PLAIN_FORBIDDEN")

	// ErrChatSendPollForbidden
	// | 403 | CHAT_SEND_POLL_FORBIDDEN | You can't send polls in this chat. |
	ErrChatSendPollForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_POLL_FORBIDDEN")

	// ErrChatSendRoundvideosForbidden
	// | 403 | CHAT_SEND_ROUNDVIDEOS_FORBIDDEN | You can't send round videos to this chat. |
	ErrChatSendRoundvideosForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_ROUNDVIDEOS_FORBIDDEN")

	// ErrChatSendStickersForbidden
	// | 403 | CHAT_SEND_STICKERS_FORBIDDEN | You can't send stickers in this chat. |
	ErrChatSendStickersForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_STICKERS_FORBIDDEN")

	// ErrChatSendVideosForbidden
	// | 403 | CHAT_SEND_VIDEOS_FORBIDDEN | You can't send videos in this chat. |
	ErrChatSendVideosForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_VIDEOS_FORBIDDEN")

	// ErrChatSendVoicesForbidden
	// | 403 | CHAT_SEND_VOICES_FORBIDDEN | You can't send voice recordings in this chat. |
	ErrChatSendVoicesForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_VOICES_FORBIDDEN")

	// ErrChatSendWebpageForbidden
	// | 403 | CHAT_SEND_WEBPAGE_FORBIDDEN | You can't send webpage previews to this chat. |
	ErrChatSendWebpageForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_SEND_WEBPAGE_FORBIDDEN")

	// ErrChatTypeInvalid
	// | 403 | CHAT_TYPE_INVALID | The specified user type is invalid. |
	ErrChatTypeInvalid = xerr.NewCodeError(ErrForbidden, "CHAT_TYPE_INVALID")

	// ErrChatWriteForbidden
	// | 403 | CHAT_WRITE_FORBIDDEN | You can't write in this chat. |
	ErrChatWriteForbidden = xerr.NewCodeError(ErrForbidden, "CHAT_WRITE_FORBIDDEN")

	// ErrEditBotInviteForbidden
	// | 403 | EDIT_BOT_INVITE_FORBIDDEN | Normal users can't edit invites that were created by bots. |
	ErrEditBotInviteForbidden = xerr.NewCodeError(ErrForbidden, "EDIT_BOT_INVITE_FORBIDDEN")

	// ErrGroupcallAlreadyStarted
	// | 403 | GROUPCALL_ALREADY_STARTED | The groupcall has already started, you can join directly using [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrGroupcallAlreadyStarted = xerr.NewCodeError(ErrForbidden, "GROUPCALL_ALREADY_STARTED")

	// ErrInlineBotRequired
	// | 403 | INLINE_BOT_REQUIRED | Only the inline bot can edit message. |
	ErrInlineBotRequired = xerr.NewCodeError(ErrForbidden, "INLINE_BOT_REQUIRED")

	// ErrMessageAuthorRequired
	// | 403 | MESSAGE_AUTHOR_REQUIRED | Message author required. |
	ErrMessageAuthorRequired = xerr.NewCodeError(ErrForbidden, "MESSAGE_AUTHOR_REQUIRED")

	// ErrMessageDeleteForbidden
	// | 403 | MESSAGE_DELETE_FORBIDDEN | You can't delete one of the messages you tried to delete, most likely because it is a service message. |
	ErrMessageDeleteForbidden = xerr.NewCodeError(ErrForbidden, "MESSAGE_DELETE_FORBIDDEN")

	// ErrPollVoteRequired
	// | 403 | POLL_VOTE_REQUIRED | Cast a vote in the poll before calling this method. |
	ErrPollVoteRequired = xerr.NewCodeError(ErrForbidden, "POLL_VOTE_REQUIRED")

	// ErrPublicChannelMissing
	// | 403 | PUBLIC_CHANNEL_MISSING | You can only export group call invite links for public chats or channels. |
	ErrPublicChannelMissing = xerr.NewCodeError(ErrForbidden, "PUBLIC_CHANNEL_MISSING")

	// ErrRightForbidden
	// | 403 | RIGHT_FORBIDDEN | Your admin rights do not allow you to do this. |
	ErrRightForbidden = xerr.NewCodeError(ErrForbidden, "RIGHT_FORBIDDEN")

	// ErrSensitiveChangeForbidden
	// | 403 | SENSITIVE_CHANGE_FORBIDDEN | You can't change your sensitive content settings. |
	ErrSensitiveChangeForbidden = xerr.NewCodeError(ErrForbidden, "SENSITIVE_CHANGE_FORBIDDEN")

	// ErrUserDeleted
	// | 403 | USER_DELETED | You can't send this secret message because the other participant deleted their account. |
	ErrUserDeleted = xerr.NewCodeError(ErrForbidden, "USER_DELETED")

	// ErrUserPermissionDenied
	// | 403 | USER_PERMISSION_DENIED | The user hasn't granted or has revoked the bot's access to change their emoji status using [bots.toggleUserEmojiStatusPermission](https://core.telegram.org/method/bots.toggleUserEmojiStatusPermission). |
	ErrUserPermissionDenied = xerr.NewCodeError(ErrForbidden, "USER_PERMISSION_DENIED")

	// ErrUserPrivacyRestricted
	// | 403 | USER_PRIVACY_RESTRICTED | The user's privacy settings do not allow you to do this. |
	ErrUserPrivacyRestricted = xerr.NewCodeError(ErrForbidden, "USER_PRIVACY_RESTRICTED")

	// ErrYourPrivacyRestricted
	// | 403 | YOUR_PRIVACY_RESTRICTED | You cannot fetch the read date of this message because you have disallowed other users to do so for *your* messages; to fix, allow other users to see *your* exact last online date OR purchase a [Telegram Premium](https://core.telegram.org/api/premium) subscription. |
	ErrYourPrivacyRestricted = xerr.NewCodeError(ErrForbidden, "YOUR_PRIVACY_RESTRICTED")
)

// 406 NOT_ACCEPTABLE

// NewErrPreviousChatImportActiveWaitX
// ErrPreviousChatImportActiveWait_%dmin
// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
func NewErrPreviousChatImportActiveWaitX(minute int32) error {
	return xerr.NewCodeErrorf(ErrNotAcceptable, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN", minute)
}

var (

	// Err406BannedRightsInvalid
	// | 406 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	Err406BannedRightsInvalid = xerr.NewCodeError(ErrNotAcceptable, "BANNED_RIGHTS_INVALID")

	// Err406ChannelPrivate
	// | 406 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	Err406ChannelPrivate = xerr.NewCodeError(ErrNotAcceptable, "CHANNEL_PRIVATE")

	// Err406ChannelTooLarge
	// | 406 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	Err406ChannelTooLarge = xerr.NewCodeError(ErrNotAcceptable, "CHANNEL_TOO_LARGE")

	// Err406ChatForwardsRestricted
	// | 406 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	Err406ChatForwardsRestricted = xerr.NewCodeError(ErrNotAcceptable, "CHAT_FORWARDS_RESTRICTED")

	// Err406FreshChangeAdminsForbidden
	// | 406 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	Err406FreshChangeAdminsForbidden = xerr.NewCodeError(ErrNotAcceptable, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// Err406InviteHashExpired
	// | 406 | INVITE_HASH_EXPIRED | The invite link has expired. |
	Err406InviteHashExpired = xerr.NewCodeError(ErrNotAcceptable, "INVITE_HASH_EXPIRED")

	// Err406PeerIdInvalid
	// | 406 | PEER_ID_INVALID | The provided peer id is invalid. |
	Err406PeerIdInvalid = xerr.NewCodeError(ErrNotAcceptable, "PEER_ID_INVALID")

	// Err406PhoneNumberInvalid
	// | 406 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	Err406PhoneNumberInvalid = xerr.NewCodeError(ErrNotAcceptable, "PHONE_NUMBER_INVALID")

	// Err406PrivacyPremiumRequired
	// | 406 | PRIVACY_PREMIUM_REQUIRED | You need a [Telegram Premium subscription](https://core.telegram.org/api/premium) to send a message to this user. |
	Err406PrivacyPremiumRequired = xerr.NewCodeError(ErrNotAcceptable, "PRIVACY_PREMIUM_REQUIRED")

	// Err406StickersetInvalid
	// | 406 | STICKERSET_INVALID | The provided sticker set is invalid. |
	Err406StickersetInvalid = xerr.NewCodeError(ErrNotAcceptable, "STICKERSET_INVALID")

	// Err406TopicClosed
	// | 406 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	Err406TopicClosed = xerr.NewCodeError(ErrNotAcceptable, "TOPIC_CLOSED")

	// Err406TopicDeleted
	// | 406 | TOPIC_DELETED | The specified topic was deleted. |
	Err406TopicDeleted = xerr.NewCodeError(ErrNotAcceptable, "TOPIC_DELETED")

	// Err406UserRestricted
	// | 406 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	Err406UserRestricted = xerr.NewCodeError(ErrNotAcceptable, "USER_RESTRICTED")

	// Err406UserpicUploadRequired
	// | 406 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	Err406UserpicUploadRequired = xerr.NewCodeError(ErrNotAcceptable, "USERPIC_UPLOAD_REQUIRED")

	// ErrAllowPaymentRequired
	// | 406 | ALLOW_PAYMENT_REQUIRED | This peer only accepts [paid messages &raquo;](https://core.telegram.org/api/paid-messages): this error is only emitted for older layers without paid messages support, so the client must be updated in order to use paid messages.  . |
	ErrAllowPaymentRequired = xerr.NewCodeError(ErrNotAcceptable, "ALLOW_PAYMENT_REQUIRED")

	// ErrApiGiftRestrictedUpdateApp
	// | 406 | API_GIFT_RESTRICTED_UPDATE_APP | Please update the app to access the gift API. |
	ErrApiGiftRestrictedUpdateApp = xerr.NewCodeError(ErrNotAcceptable, "API_GIFT_RESTRICTED_UPDATE_APP")

	// ErrAuthKeyDuplicated
	// | 406 | AUTH_KEY_DUPLICATED | Concurrent usage of the current session from multiple connections was detected, the current session was invalidated by the server for security reasons! |
	ErrAuthKeyDuplicated = xerr.NewCodeError(ErrNotAcceptable, "AUTH_KEY_DUPLICATED")

	// ErrBusinessAddressActive
	// | 406 | BUSINESS_ADDRESS_ACTIVE | The user is currently advertising a [Business Location](https://core.telegram.org/api/business#location), the location may only be changed (or removed) using [account.updateBusinessLocation &raquo;](https://core.telegram.org/method/account.updateBusinessLocation).  . |
	ErrBusinessAddressActive = xerr.NewCodeError(ErrNotAcceptable, "BUSINESS_ADDRESS_ACTIVE")

	// ErrCallProtocolCompatLayerInvalid
	// | 406 | CALL_PROTOCOL_COMPAT_LAYER_INVALID | The other side of the call does not support any of the VoIP protocols supported by the local client, as specified by the `protocol.layer` and `protocol.library_versions` fields. |
	ErrCallProtocolCompatLayerInvalid = xerr.NewCodeError(ErrNotAcceptable, "CALL_PROTOCOL_COMPAT_LAYER_INVALID")

	// ErrFilerefUpgradeNeeded
	// | 406 | FILEREF_UPGRADE_NEEDED | The client has to be updated in order to support [file references](https://core.telegram.org/api/file-references). |
	ErrFilerefUpgradeNeeded = xerr.NewCodeError(ErrNotAcceptable, "FILEREF_UPGRADE_NEEDED")

	// ErrFreshChangePhoneForbidden
	// | 406 | FRESH_CHANGE_PHONE_FORBIDDEN | You can't change phone number right after logging in, please wait at least 24 hours. |
	ErrFreshChangePhoneForbidden = xerr.NewCodeError(ErrNotAcceptable, "FRESH_CHANGE_PHONE_FORBIDDEN")

	// ErrFreshResetAuthorisationForbidden
	// | 406 | FRESH_RESET_AUTHORISATION_FORBIDDEN | You can't logout other sessions if less than 24 hours have passed since you logged on the current session. |
	ErrFreshResetAuthorisationForbidden = xerr.NewCodeError(ErrNotAcceptable, "FRESH_RESET_AUTHORISATION_FORBIDDEN")

	// ErrPaymentUnsupported
	// | 406 | PAYMENT_UNSUPPORTED | A detailed description of the error will be received separately as described [here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrPaymentUnsupported = xerr.NewCodeError(ErrNotAcceptable, "PAYMENT_UNSUPPORTED")

	// ErrPhonePasswordFlood
	// | 406 | PHONE_PASSWORD_FLOOD | You have tried logging in too many times. |
	ErrPhonePasswordFlood = xerr.NewCodeError(ErrNotAcceptable, "PHONE_PASSWORD_FLOOD")

	// ErrPrecheckoutFailed
	// | 406 | PRECHECKOUT_FAILED | Precheckout failed, a detailed and localized description for the error will be emitted via an [updateServiceNotification as specified here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrPrecheckoutFailed = xerr.NewCodeError(ErrNotAcceptable, "PRECHECKOUT_FAILED")

	// ErrPremiumCurrentlyUnavailable
	// | 406 | PREMIUM_CURRENTLY_UNAVAILABLE | You cannot currently purchase a Premium subscription. |
	ErrPremiumCurrentlyUnavailable = xerr.NewCodeError(ErrNotAcceptable, "PREMIUM_CURRENTLY_UNAVAILABLE")

	// ErrPreviousChatImportActiveWait_%dmin
	// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
	// ErrPreviousChatImportActiveWait_%dmin = xerr.NewCodeError(ErrNotAcceptable, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN")

	// ErrSendCodeUnavailable
	// | 406 | SEND_CODE_UNAVAILABLE | Returned when all available options for this type of number were already used (e.g. flash-call, then SMS, then this error might be returned to trigger a second resend). |
	ErrSendCodeUnavailable = xerr.NewCodeError(ErrNotAcceptable, "SEND_CODE_UNAVAILABLE")

	// ErrStargiftExportInProgress
	// | 406 | STARGIFT_EXPORT_IN_PROGRESS | A gift export is in progress, a detailed and localized description for the error will be emitted via an [updateServiceNotification as specified here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrStargiftExportInProgress = xerr.NewCodeError(ErrNotAcceptable, "STARGIFT_EXPORT_IN_PROGRESS")

	// ErrStarsFormAmountMismatch
	// | 406 | STARS_FORM_AMOUNT_MISMATCH | The form amount has changed, please fetch the new form using [payments.getPaymentForm](https://core.telegram.org/method/payments.getPaymentForm) and restart the process. |
	ErrStarsFormAmountMismatch = xerr.NewCodeError(ErrNotAcceptable, "STARS_FORM_AMOUNT_MISMATCH")

	// ErrStickersetOwnerAnonymous
	// | 406 | STICKERSET_OWNER_ANONYMOUS | Provided stickerset can't be installed as group stickerset to prevent admin deanonymization. |
	ErrStickersetOwnerAnonymous = xerr.NewCodeError(ErrNotAcceptable, "STICKERSET_OWNER_ANONYMOUS")

	// ErrTranslationsDisabled
	// | 406 | TRANSLATIONS_DISABLED | Translations are unavailable, a detailed and localized description for the error will be emitted via an [updateServiceNotification as specified here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrTranslationsDisabled = xerr.NewCodeError(ErrNotAcceptable, "TRANSLATIONS_DISABLED")

	// ErrUpdateAppToLogin
	// | 406 | UPDATE_APP_TO_LOGIN | Please update your client to login. |
	ErrUpdateAppToLogin = xerr.NewCodeError(ErrNotAcceptable, "UPDATE_APP_TO_LOGIN")

	// ErrUserpicPrivacyRequired
	// | 406 | USERPIC_PRIVACY_REQUIRED | You need to disable privacy settings for your profile picture in order to make your geolocation public. |
	ErrUserpicPrivacyRequired = xerr.NewCodeError(ErrNotAcceptable, "USERPIC_PRIVACY_REQUIRED")
)

// 500 InternalServerError
var (
	// StatusInternalServerError - StatusInternelServerError
	// StatusInternalServerError = xerr.NewCodeError(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrInternalServerError
	// | 500 | INTERNAL_SERVER_ERROR |  |
	ErrInternalServerError = xerr.NewCodeError(ErrInternal, "INTERNAL_SERVER_ERROR")

	// Err500CallOccupyFailed
	// | 500 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	Err500CallOccupyFailed = xerr.NewCodeError(ErrInternal, "CALL_OCCUPY_FAILED")

	// Err500ChatInvalid
	// | 500 | CHAT_INVALID | Invalid chat. |
	Err500ChatInvalid = xerr.NewCodeError(ErrInternal, "CHAT_INVALID")

	// Err500MsgWaitFailed
	// | 500 | MSG_WAIT_FAILED | A waiting call returned an error. |
	Err500MsgWaitFailed = xerr.NewCodeError(ErrInternal, "MSG_WAIT_FAILED")

	// ErrAuthKeyUnsynchronized
	// | 500 | AUTH_KEY_UNSYNCHRONIZED | Internal error, please repeat the method call. |
	ErrAuthKeyUnsynchronized = xerr.NewCodeError(ErrInternal, "AUTH_KEY_UNSYNCHRONIZED")

	// ErrAuthRestart
	// | 500 | AUTH_RESTART | Restart the authorization process. |
	ErrAuthRestart = xerr.NewCodeError(ErrInternal, "AUTH_RESTART")

	// ErrAuthRestart_%d
	// | 500 | AUTH_RESTART_%d | Internal error (debug info %d), please repeat the method call. |
	// ErrAuthRestart_%d = xerr.NewCodeError(ErrInternal, "AUTH_RESTART_%d")

	// ErrCdnUploadTimeout
	// | 500 | CDN_UPLOAD_TIMEOUT | A server-side timeout occurred while reuploading the file to the CDN DC. |
	ErrCdnUploadTimeout = xerr.NewCodeError(ErrInternal, "CDN_UPLOAD_TIMEOUT")

	// ErrChatIdGenerateFailed
	// | 500 | CHAT_ID_GENERATE_FAILED | Failure while generating the chat ID. |
	ErrChatIdGenerateFailed = xerr.NewCodeError(ErrInternal, "CHAT_ID_GENERATE_FAILED")

	// ErrPersistentTimestampOutdated
	// | 500 | PERSISTENT_TIMESTAMP_OUTDATED | Channel internal replication issues, try again later (treat this like an RPC_CALL_FAIL). |
	ErrPersistentTimestampOutdated = xerr.NewCodeError(ErrInternal, "PERSISTENT_TIMESTAMP_OUTDATED")

	// ErrRandomIdDuplicate
	// | 500 | RANDOM_ID_DUPLICATE | You provided a random ID that was already used. |
	ErrRandomIdDuplicate = xerr.NewCodeError(ErrInternal, "RANDOM_ID_DUPLICATE")

	// ErrSendMediaInvalid
	// | 500 | SEND_MEDIA_INVALID | The specified media is invalid. |
	ErrSendMediaInvalid = xerr.NewCodeError(ErrInternal, "SEND_MEDIA_INVALID")

	// ErrSignInFailed
	// | 500 | SIGN_IN_FAILED | Failure while signing in. |
	ErrSignInFailed = xerr.NewCodeError(ErrInternal, "SIGN_IN_FAILED")

	// ErrTranslateReqFailed
	// | 500 | TRANSLATE_REQ_FAILED | Translation failed, please try again later. |
	ErrTranslateReqFailed = xerr.NewCodeError(ErrInternal, "TRANSLATE_REQ_FAILED")

	// ErrTranslationTimeout
	// | 500 | TRANSLATION_TIMEOUT | A timeout occurred while translating the specified text. |
	ErrTranslationTimeout = xerr.NewCodeError(ErrInternal, "TRANSLATION_TIMEOUT")
)

// -503 Timeout
var (
	// StatusTimeout - StatusTimeout
	// StatusTimeout = status.New(ErrTimeOut503, "Timeout")

	// ErrTimeout
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeout = xerr.NewCodeError(ErrTimeOut503, "Timeout")
)

// -500
var (
// // ErrInvalid MsgResendReq Query
// // | -500 | Invalid msg_resend_req query | Invalid msg_resend_req query. |
// ErrInvalid MsgResendReq Query = xerr.NewCodeError(-500, "Invalid msg_resend_req query")
//
// // ErrInvalid MsgsAck Query
// // | -500 | Invalid msgs_ack query | Invalid msgs_ack query. |
// ErrInvalid MsgsAck Query = xerr.NewCodeError(-500, "Invalid msgs_ack query")
//
// // ErrInvalid MsgsStateReq Query
// // | -500 | Invalid msgs_state_req query | Invalid msgs_state_req query. |
// ErrInvalid MsgsStateReq Query = xerr.NewCodeError(-500, "Invalid msgs_state_req query")
)

// 700
var (
	// ErrPushRpcClient
	// db error
	// TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
	ErrPushRpcClient = xerr.NewCodeError(ErrNotReturnClient, "NOTRETURN_CLIENT")

	// ErrMigratedToChannel
	// MIGRATED_TO_CHANNEL
	ErrMigratedToChannel = xerr.NewCodeError(ErrNotReturnClient, "MIGRATED_TO_CHANNEL")
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
