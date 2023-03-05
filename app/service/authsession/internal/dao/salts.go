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

/**
 ## Android client source code, gen salts:
 ### handshake:
   - client received TL_server_DH_params_ok:
	```
	handshakeServerSalt = new TL_future_salt();
	handshakeServerSalt->valid_since = currentTime + timeDifference - 5;
	handshakeServerSalt->valid_until = handshakeServerSalt->valid_since + 30 * 60;
	for (int32_t a = 7; a >= 0; a--) {
		handshakeServerSalt->salt <<= 8;
		handshakeServerSalt->salt |= (authNewNonce->bytes[a] ^ authServerNonce->bytes[a]);
	}
	```

   - client received TL_dh_gen_ok:
	```
	std::unique_ptr<TL_future_salt> salt = std::unique_ptr<TL_future_salt>(handshakeServerSalt);
	addServerSalt(salt);
	handshakeServerSalt = nullptr;
	```

 ### received TL_new_session_created:
	```
	std::unique_ptr<TL_future_salt> salt = std::unique_ptr<TL_future_salt>(new TL_future_salt());
	salt->valid_until = salt->valid_since = getCurrentTime();
	salt->valid_until += 30 * 60;
	salt->salt = response->server_salt;
	datacenter->addServerSalt(salt);

	```

 ### send TL_get_future_salts request:
   - rpc request:
	```
    requestingSaltsForDc.push_back(datacenter->getDatacenterId());
    TL_get_future_salts *request = new TL_get_future_salts();
    request->num = 32;
    sendRequest(request, [&, datacenter](TLObject *response, TL_error *error, int32_t networkType) {
        std::vector<uint32_t>::iterator iter = std::find(requestingSaltsForDc.begin(), requestingSaltsForDc.end(), datacenter->getDatacenterId());
        if (iter != requestingSaltsForDc.end()) {
            requestingSaltsForDc.erase(iter);
        }
        if (error == nullptr) {
            TL_future_salts *res = (TL_future_salts *) response;
            datacenter->mergeServerSalts(res->salts);
            saveConfig();
        }
    }, nullptr, RequestFlagWithoutLogin | RequestFlagEnableUnauthorized, datacenter->getDatacenterId(), ConnectionTypeGeneric, true);

	```

  - rpc response:
	```
	TL_future_salts *response = (TL_future_salts *) message;
	int64_t requestMid = response->req_msg_id;
	for (requestsIter iter = runningRequests.begin(); iter != runningRequests.end(); iter++) {
		Request *request = iter->get();
		if (request->respondsToMessageId(requestMid)) {
			request->onComplete(response, nullptr, connection->currentNetworkType);
			request->completed = true;
			runningRequests.erase(iter);
			break;
		}
	}
	```

  - received TL_bad_server_salt:
	```
	datacenter->clearServerSalts();

	std::unique_ptr<TL_future_salt> salt = std::unique_ptr<TL_future_salt>(new TL_future_salt());
	salt->valid_until = salt->valid_since = getCurrentTime();
	salt->valid_until += 30 * 60;
	salt->salt = messageSalt;
	datacenter->addServerSalt(salt);
	```
*/

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	cacheSaltPrefix = "salts"
)

func genCacheSaltKey(id int64) string {
	return fmt.Sprintf("%s_%d", cacheSaltPrefix, id)
}

func (d *Dao) PutSalts(ctx context.Context, keyId int64, salts []*mtproto.TLFutureSalt) (err error) {
	var (
		b   []byte
		key = genCacheSaltKey(keyId)
	)

	if b, err = json.Marshal(salts); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", key, err)
		return
	}

	// 误差 500
	if err = d.kv.SetexCtx(ctx, key, string(b), len(salts)*saltTimeout); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", key, err)
	}
	return
}

func (d *Dao) GetSalts(ctx context.Context, keyId int64) (salts []*mtproto.TLFutureSalt, err error) {
	var (
		key  = genCacheSaltKey(keyId)
		bBuf string
	)

	bBuf, err = d.kv.GetCtx(ctx, key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(GET %s) error(%v)", key, err)
		return
	} else if bBuf == "" {
		return nil, nil
	}

	if err = jsonx.UnmarshalFromString(bBuf, &salts); err != nil {
		logx.WithContext(ctx).Errorf("getSalts jsonx.UnmarshalFromString(%s) error(%v)", bBuf, err)
		return nil, nil
	}
	return
}

func (d *Dao) getOrNotInsertSaltList(ctx context.Context, keyId int64, size int32) ([]*mtproto.TLFutureSalt, error) {
	var (
		salts = make([]*mtproto.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		saltsData      []*mtproto.TLFutureSalt
		lastSalt       *mtproto.TLFutureSalt
	)

	saltList, err := d.GetSalts(ctx, keyId)
	if err != nil {
		return nil, err
	}

	if len(saltList) > 0 {
		hasLastSalt := false
		for idx, salt := range saltList {
			if salt.GetValidUntil() >= date {
				if !hasLastSalt {
					if idx > 0 {
						lastSalt = saltList[idx-1]
					}
					hasLastSalt = true
				}
				saltsData = append(saltsData, salt)
				if lastValidUntil < salt.GetValidUntil() {
					lastValidUntil = salt.GetValidUntil()
				}
			}
		}
		if !hasLastSalt {
			lastSalt = saltList[len(saltList)-1]
		}

		// check ValidUntil
		if lastSalt != nil && lastSalt.GetValidUntil()+300 < date {
			lastSalt = nil
		}
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < size; i++ {
			salt := mtproto.MakeTLFutureSalt(&mtproto.FutureSalt{
				ValidSince: lastValidUntil,
				ValidUntil: lastValidUntil + saltTimeout,
				Salt:       rand.Int63(),
			})
			saltsData = append(saltsData, salt)
			lastValidUntil += saltTimeout
		}
	}

	for i := int32(0); i < size; i++ {
		salts = append(salts, saltsData[i])
	}

	var (
		salts2     []*mtproto.TLFutureSalt
		saltsData2 []*mtproto.TLFutureSalt
	)

	if lastSalt != nil {
		salts2 = append(salts2, lastSalt)
		saltsData2 = append(saltsData2, lastSalt)
	}

	salts2 = append(salts2, salts...)
	saltsData2 = append(saltsData2, saltsData...)

	if left > 0 {
		err = d.PutSalts(ctx, keyId, saltsData2)
		if err != nil {
			return nil, err
		}
	}
	return salts2, nil
}

func (d *Dao) PutSaltCache(ctx context.Context, keyId int64, salt *mtproto.TLFutureSalt) error {
	return d.PutSalts(ctx, keyId, []*mtproto.TLFutureSalt{salt})
}

func (d *Dao) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*mtproto.TLFutureSalts, error) {
	pSalts, err := d.getOrNotInsertSaltList(ctx, authKeyId, num)
	if err != nil {
		return nil, err
	}
	salts := &mtproto.TLFutureSalts{Data2: &mtproto.FutureSalts{
		ReqMsgId: 0,
		Now:      0,
		Salts:    pSalts,
	}}
	return salts, nil
}
