/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"

	status2 "google.golang.org/grpc/status"
)

/*
// android
enum HandshakeType {
    HandshakeTypePerm,
    HandshakeTypeTemp,
    HandshakeTypeMediaTemp,
    HandshakeTypeCurrent,
    HandshakeTypeAll
};
*/
/*
# auth.bindTempAuthKey
	Binds a temporary authorization key temp_auth_key_id to the permanent authorization key perm_auth_key_id.
	Each permanent key may only be bound to one temporary key at a time,
	binding a new temporary key overwrites the previous one.


## Possible errors
| Code | Type | Description |
| ---- | ---- | ----------- |
| 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message is incorrect |
| 400 | INPUT_REQUEST_TOO_LONG | The request is too big |
| 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another perm_auth_key_id |
| 400 | TEMP_AUTH_KEY_EMPTY | The request was not performed with a temporary authorization key |
| -503 | Timeout | Timeout while fetching data |

*/

/*
AuthsessionBindTempAuthKey
## android client source code:

```
       TL_auth_bindTempAuthKey *request = new TL_auth_bindTempAuthKey();
       request->initFunc = [&, request, connection](int64_t messageId) {
           TL_bind_auth_key_inner *inner = new TL_bind_auth_key_inner();
           inner->expires_at = ConnectionsManager::getInstance(currentDatacenter->instanceNum).getCurrentTime() + timeDifference + TEMP_AUTH_KEY_EXPIRE_TIME;
           inner->perm_auth_key_id = currentDatacenter->authKeyPermId;
           inner->temp_auth_key_id = authKeyTempPendingId;
           RAND_bytes((uint8_t *) &inner->nonce, 8);
           inner->temp_session_id = connection->getSessionId();

           NetworkMessage *networkMessage = new NetworkMessage();
           networkMessage->message = std::unique_ptr<TL_message>(new TL_message());
           networkMessage->message->msg_id = authKeyPendingMessageId = messageId;
           networkMessage->message->bytes = inner->getObjectSize();
           networkMessage->message->body = std::unique_ptr<TLObject>(inner);
           networkMessage->message->seqno = 0;

           std::vector<std::unique_ptr<NetworkMessage>> array;
           array.push_back(std::unique_ptr<NetworkMessage>(networkMessage));

           request->perm_auth_key_id = inner->perm_auth_key_id;
           request->nonce = inner->nonce;
           request->expires_at = inner->expires_at;
           request->encrypted_message = currentDatacenter->createRequestsData(array, nullptr, connection, true);
       };

       authKeyPendingRequestId = ConnectionsManager::getInstance(currentDatacenter->instanceNum).sendRequest(request, [&](TLObject *response, TL_error *error, int32_t networkType) {
           authKeyPendingMessageId = 0;
           authKeyPendingRequestId = 0;
           if (response != nullptr && typeid(*response) == typeid(TL_boolTrue)) {
               if (LOGS_ENABLED) DEBUG_D("account%u dc%u handshake: bind completed", currentDatacenter->instanceNum, currentDatacenter->datacenterId);
               ConnectionsManager::getInstance(currentDatacenter->instanceNum).scheduleTask([&] {
                   ByteArray *authKey = authKeyTempPending;
                   authKeyTempPending = nullptr;
                   delegate->onHandshakeComplete(this, authKeyTempPendingId, authKey, timeDifference);
               });
           } else if (error == nullptr || error->code != 400 || error->text.find("ENCRYPTED_MESSAGE_INVALID") == std::string::npos) {
               ConnectionsManager::getInstance(currentDatacenter->instanceNum).scheduleTask([&] {
                   beginHandshake(true);
               });
           }
       }, nullptr, RequestFlagWithoutLogin | RequestFlagEnableUnauthorized | RequestFlagUseUnboundKey, currentDatacenter->datacenterId, connection->getConnectionType(), true, 0);
   }
```
*/
// authsession.bindTempAuthKey perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (c *AuthsessionCore) AuthsessionBindTempAuthKey(in *authsession.TLAuthsessionBindTempAuthKey) (*mtproto.Bool, error) {
	// 400	ENCRYPTED_MESSAGE_INVALID	Encrypted message is incorrect
	// 400	INPUT_REQUEST_TOO_LONG	The request is too big
	// 400	TEMP_AUTH_KEY_ALREADY_BOUND	The passed temporary key is already bound to another perm_auth_key_id
	// 400	TEMP_AUTH_KEY_EMPTY	The request was not performed with a temporary authorization key
	// -503	Timeout	Timeout while fetching data
	//

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetPermAuthKeyId())
	if err != nil {
		c.Logger.Errorf("auth.bindTempAuthKey - error: %v", err)
		if status2.Convert(err).Message() == "AUTH_KEY_UNREGISTERED" {
			err = mtproto.ErrEncryptedMessageInvalid
		}

		return nil, err
	}

	permAuthKey := crypto.NewAuthKey(in.PermAuthKeyId, keyData.AuthKey)
	innerData, err := permAuthKey.AesIgeDecryptV1(in.EncryptedMessage[8:8+16], in.EncryptedMessage[8+16:])
	if err != nil {
		c.Logger.Errorf("auth.bindTempAuthKey - error: %v", err)
		return nil, mtproto.ErrEncryptedMessageInvalid
	}

	// 8+8+8+8

	//// 2. 反序列化出pqInnerData
	dbuf := mtproto.NewDecodeBuf(innerData[32:])
	o := dbuf.Object()
	if dbuf.GetError() != nil {
		c.Logger.Errorf("auth.bindTempAuthKey - error: %v", dbuf.GetError())
		return nil, mtproto.ErrEncryptedMessageInvalid
	} else if bindAuthKeyInner, ok := o.(*mtproto.TLBindAuthKeyInner); !ok {
		c.Logger.Errorf("auth.bindTempAuthKey - invalid innerData")
		return nil, mtproto.ErrEncryptedMessageInvalid
	} else {
		// bind_auth_key_inner#75a3f765 nonce:long temp_auth_key_id:long perm_auth_key_id:long temp_session_id:long expires_at:int = BindAuthKeyInner;
		// bind
		c.Logger.Infof("auth.bindTempAuthKey - bind_auth_key_inner: %s", bindAuthKeyInner)
		tempKeyData, err2 := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, bindAuthKeyInner.GetTempAuthKeyId())
		if err2 != nil {
			c.Logger.Errorf("auth.bindTempAuthKey - invalid innerData")
			return nil, mtproto.ErrEncryptedMessageInvalid
		}

		// TODO: tx wrapper
		// bindTemp
		c.svcCtx.Dao.UnsafeBindKeyIdV2(c.ctx,
			bindAuthKeyInner.GetPermAuthKeyId(),
			tempKeyData.AuthKeyType,
			bindAuthKeyInner.GetTempAuthKeyId())

		// TODO: expiredIn int32
		// bindPerm
		c.svcCtx.Dao.UnsafeBindKeyIdV2(c.ctx,
			bindAuthKeyInner.GetTempAuthKeyId(),
			mtproto.AuthKeyTypePerm,
			bindAuthKeyInner.GetPermAuthKeyId())
	}

	return mtproto.BoolTrue, nil
}
