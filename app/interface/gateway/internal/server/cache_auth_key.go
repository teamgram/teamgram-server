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

package server

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"
)

type CacheV struct {
	V *mtproto.AuthKeyInfo
}

func (c CacheV) Size() int {
	return 1
}

func (s *Server) GetAuthKey(authKeyId int64) *mtproto.AuthKeyInfo {
	var (
		cacheK = fmt.Sprintf("%d", authKeyId)
		value  *CacheV
	)

	if v, ok := s.cache.Get(cacheK); ok {
		value = v.(*CacheV)
	}

	if value == nil {
		return nil
	} else {
		return value.V
	}
}

func (s *Server) PutAuthKey(keyInfo *mtproto.AuthKeyInfo) {
	var (
		cacheK = fmt.Sprintf("%d", keyInfo.AuthKeyId)
	)

	// TODO: expires_in
	s.cache.Set(cacheK, &CacheV{V: keyInfo})
}
