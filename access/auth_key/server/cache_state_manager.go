/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// salt cache
package server

import (
	"fmt"
	"time"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/cache"
	"encoding/hex"
)

const (
	kSaltTimeout = 30 * 60 // salt timeout
	// kCacheConfig = `{"conn":":"127.0.0.1:6039"}`
	kAdapterName     = "memory"
	kCacheConfig     = `{"interval":60}`
	kCacheSaltPrefix = "salts"
)

var cacheStates *cacheStateManager = nil

func init() {
	// rand.Seed(time.Now().UnixNano())
	initCacheStateManager(kAdapterName, kCacheConfig)
}

type cacheStateManager struct {
	cache   cache.Cache
	timeout time.Duration // salt timeout
}

func initCacheStateManager(name, config string) error {
	if config == "" {
		config = kCacheConfig
	}

	c, err := cache.NewCache(name, kCacheConfig)
	if err != nil {
		glog.Error(err)
		return err
	}

	cacheStates = &cacheStateManager{cache: c, timeout: kSaltTimeout}
	return nil
}

func genCacheStateKey(nonce, serverNonce []byte) string {
	return fmt.Sprintf("%s_%s@%s", kCacheSaltPrefix, hex.EncodeToString(nonce), hex.EncodeToString(serverNonce))
}

func PutCacheState(nonce, serverNonce []byte, state *mtproto.HandshakeContext_Data) (error) {
	k := genCacheStateKey(nonce, serverNonce)
	glog.Info("put state key: (", k, ")")
	return cacheStates.cache.Put(genCacheStateKey(nonce, serverNonce), state, time.Duration(5*time.Second))
}

func GetCacheState(nonce, serverNonce []byte) (state *mtproto.HandshakeContext_Data) {
	k := genCacheStateKey(nonce, serverNonce)
	glog.Info("get state key: (", k, ")")
	v := cacheStates.cache.Get(genCacheStateKey(nonce, serverNonce))
	if v != nil {
		state = v.(*mtproto.HandshakeContext_Data)
	} else {
		glog.Warning("not found state by: (", k, ")")
	}
	return
}
