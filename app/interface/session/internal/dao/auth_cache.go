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
	"fmt"
	"strconv"
	"time"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"

	"github.com/zeromicro/go-zero/core/logx"
)

type cacheAuthValue struct {
	SaltList []*tg.TLFutureSalt
}

// Size Impl cache.Value interface
func (cv *cacheAuthValue) Size() int {
	return 1
}

func (d *Dao) getCacheValue(authKeyId int64) *cacheAuthValue {
	var (
		cacheK = strconv.FormatInt(authKeyId, 10)
	)

	if v, ok := d.cache.Get(cacheK); !ok {
		cv := new(cacheAuthValue)
		d.cache.Set(cacheK, cv)
		return cv
	} else {
		return v.(*cacheAuthValue)
	}
}

func (d *Dao) getFutureSaltList(ctx context.Context, authKeyId int64) ([]*tg.TLFutureSalt, bool) {
	var (
		cv   = d.getCacheValue(authKeyId)
		date = int32(time.Now().Unix())
	)

	if len(cv.SaltList) > 0 {
		futureSalts := cv.SaltList
		for i, salt := range futureSalts {
			if salt.ValidUntil >= date {
				if i > 0 {
					return futureSalts[i-1:], true
				} else {
					return futureSalts[i:], true
				}
			}
		}
	}

	futureSalts, err := d.AuthsessionClient.AuthsessionGetFutureSalts(ctx, &authsession.TLAuthsessionGetFutureSalts{
		AuthKeyId: authKeyId,
	})
	if err != nil {
		logx.WithContext(ctx).Error(err.Error())
		return nil, false
	}

	var (
		rB       = false
		saltList []*tg.TLFutureSalt
	)

	futureSalts.Match(func(futureSalts *tg.TLFutureSalts) interface{} {
		for i, salt := range futureSalts.Salts {
			if salt.ValidUntil >= date {
				if i > 0 {
					saltList = saltList[i-1:]
					cv.SaltList = saltList
					rB = true
					break
				} else {
					saltList = saltList[i:]
					cv.SaltList = saltList
					rB = true
					break
				}
			}
		}

		return nil
	})

	//saltList := futureSalts.GetSalts()
	//for i, salt := range saltList {
	//	if salt.Data2.ValidUntil >= date {
	//		if i > 0 {
	//			saltList = saltList[i-1:]
	//			cv.SaltList = saltList
	//			return saltList, true
	//		} else {
	//			saltList = saltList[i:]
	//			cv.SaltList = saltList
	//			return saltList, true
	//		}
	//	}
	//}

	return saltList, rB
}

func (d *Dao) GetOrFetchNewSalt(ctx context.Context, authKeyId int64) (salt, lastInvalidSalt *tg.TLFutureSalt, err error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	//TODO(@benqi): check len(cacheSalts) > 0
	if len(cacheSalts) < 2 {
		return nil, nil, fmt.Errorf("get salt error")
	} else {
		if cacheSalts[0].ValidUntil >= int32(time.Now().Unix()) {
			return cacheSalts[0], nil, nil
		} else {
			return cacheSalts[1], cacheSalts[0], nil
		}
	}
}

func (d *Dao) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) ([]*tg.TLFutureSalt, error) {
	cacheSalts, _ := d.getFutureSaltList(ctx, authKeyId)
	//TODO(@benqi): check len(cacheSalts) > 0

	return cacheSalts, nil
}
