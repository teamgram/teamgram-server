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

// salt cache
package auth_session

import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/cache"
	_ "github.com/nebula-chat/chatengine/pkg/cache/redis"
)

const (
	kSaltTimeout = 30 * 60 // salt timeout
	// kCacheConfig = `{"conn":":"127.0.0.1:6039"}`
	// kAdapterName     = "memory"
	// kCacheConfig     = `{"interval":60}`
	kCacheSaltPrefix = "salts"
)

var cacheSalts *cacheSaltManager = nil

func init() {
	rand.Seed(time.Now().UnixNano())
	// initCacheSaltsManager(kAdapterName, kCacheConfig)
}

type cacheSaltManager struct {
	cache   cache.Cache
	timeout time.Duration // salt timeout
}

type cacheSaltsData struct {
	// LastSalt *mtproto.FutureSalt_Data   `json:"last_salt"`
	Salts []*mtproto.FutureSalt_Data `json:"salts"`
}

func initCacheSaltsManager(name, config string) error {
	c, err := cache.NewCache(name, config)
	if err != nil {
		glog.Error(err)
		return err
	}

	cacheSalts = &cacheSaltManager{cache: c, timeout: kSaltTimeout}
	return nil
}

func genCacheSaltKey(id int64) string {
	return fmt.Sprintf("%s_%d", kCacheSaltPrefix, id)
}

func GetOrNotInsertSaltList(keyId int64, size int32) ([]*mtproto.TLFutureSalt, error) {
	var (
		salts = make([]*mtproto.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		// ok = false
		saltsData []*mtproto.FutureSalt_Data
		cacheKey  = genCacheSaltKey(keyId)
		lastSalt  *mtproto.FutureSalt_Data
	)

	v := cacheSalts.cache.Get(cacheKey)
	if v != nil {
		if cacheData, ok := v.([]byte); ok {
			caches := &cacheSaltsData{}
			err := json.Unmarshal(cacheData, caches)
			if err != nil {
				glog.Error("unmarshal error - ", err)
			} else {
				hasLastSalt := false
				for idx, salt := range caches.Salts {
					if salt.ValidUntil >= date {
						if !hasLastSalt {
							if idx > 0 {
								lastSalt = caches.Salts[idx-1]
								// saltsData = append(saltsData, saltList[idx-1])
							}
							hasLastSalt = true
						}
						saltsData = append(saltsData, salt)
						if lastValidUntil < salt.ValidUntil {
							lastValidUntil = salt.ValidUntil
						}
					}
				}
				if !hasLastSalt {
					lastSalt = caches.Salts[len(caches.Salts)-1]
				}

				// check ValidUntil
				if lastSalt != nil && lastSalt.ValidUntil+300 < date {
					lastSalt = nil
				}
			}
		} else {
			glog.Error("invalid cache - ", cacheKey)
		}
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < size; i++ {
			salt := &mtproto.FutureSalt_Data{
				ValidSince: lastValidUntil,
				ValidUntil: lastValidUntil + kSaltTimeout,
				Salt:       rand.Int63(),
			}
			saltsData = append(saltsData, salt)
			lastValidUntil += kSaltTimeout
		}
	}

	for i := int32(0); i < size; i++ {
		salt := &mtproto.TLFutureSalt{
			Data2: saltsData[i],
		}
		salts = append(salts, salt)
	}

	var (
		salts2     []*mtproto.TLFutureSalt
		saltsData2 []*mtproto.FutureSalt_Data
	)

	if lastSalt != nil {
		salts2 = append(salts2, &mtproto.TLFutureSalt{Data2: lastSalt})
		saltsData2 = append(saltsData2, lastSalt)
	}

	salts2 = append(salts2, salts...)
	saltsData2 = append(saltsData2, saltsData...)

	if left > 0 {
		caches := &cacheSaltsData{
			Salts: saltsData2,
		}
		cacheData, _ := json.Marshal(caches)
		err := cacheSalts.cache.Put(cacheKey, cacheData, time.Duration(len(saltsData))*kSaltTimeout*time.Second)
		if err != nil {
			// glog.Error(err)
			return nil, err
		}
	}
	return salts2, nil
}

func PutSaltCacche(keyId int64, salt *mtproto.TLFutureSalt) error {
	cacheKey := genCacheSaltKey(keyId)
	caches := &cacheSaltsData{
		Salts: []*mtproto.FutureSalt_Data{salt.Data2},
	}
	cacheData, _ := json.Marshal(caches)
	return cacheSalts.cache.Put(cacheKey, cacheData, kSaltTimeout*time.Second)
}
