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

package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/proto/mtproto"

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
	if err = d.kv.Setex(key, string(b), len(salts)*saltTimeout); err != nil {
		logx.WithContext(ctx).Errorf("conn.SETEX(%s) error(%v)", key, err)
	}
	return
}

func (d *Dao) GetSalts(ctx context.Context, keyId int64) (salts []*mtproto.TLFutureSalt, err error) {
	var (
		key  = genCacheSaltKey(keyId)
		bBuf string
	)

	bBuf, err = d.kv.Get(key)
	if err != nil {
		logx.WithContext(ctx).Errorf("conn.Do(GET %s) error(%v)", key, err)
		return
	} else if bBuf == "" {
		return nil, nil
	}

	if err = json.Unmarshal(hack.Bytes(bBuf), &salts); err != nil {
		logx.WithContext(ctx).Errorf("getSalts json.Unmarshal(%s) error(%v)", bBuf, err)
		return nil, nil
	}
	return
}
