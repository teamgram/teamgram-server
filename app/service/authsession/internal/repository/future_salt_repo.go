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

package repository

import (
	"context"
	"math/rand"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

func (r *Repository) getOrNotInsertSaltList(ctx context.Context, keyId int64, size int32) ([]*tg.TLFutureSalt, error) {
	var (
		salts = make([]*tg.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		saltsData      []*tg.TLFutureSalt
		lastSalt       *tg.TLFutureSalt
	)

	saltList, err := r.FutureSaltsModel.GetSalts(ctx, keyId)
	if err != nil {
		return nil, err
	}

	if len(saltList) > 0 {
		hasLastSalt := false
		for idx, salt := range saltList {
			if salt.ValidSince >= date {
				if !hasLastSalt {
					if idx > 0 {
						lastSalt = saltList[idx-1]
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
			lastSalt = saltList[len(saltList)-1]
		}

		// check ValidUntil
		if lastSalt != nil && lastSalt.ValidUntil+300 < date {
			lastSalt = nil
		}
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < size; i++ {
			salt := tg.MakeTLFutureSalt(&tg.TLFutureSalt{
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
		salts2     []*tg.TLFutureSalt
		saltsData2 []*tg.TLFutureSalt
	)

	if lastSalt != nil {
		salts2 = append(salts2, lastSalt)
		saltsData2 = append(saltsData2, lastSalt)
	}

	salts2 = append(salts2, salts...)
	saltsData2 = append(saltsData2, saltsData...)

	if left > 0 {
		err = r.FutureSaltsModel.PutSalts(ctx, keyId, saltsData2, saltTimeout)
		if err != nil {
			return nil, err
		}
	}
	return salts2, nil
}

func (r *Repository) GetFutureSalts(ctx context.Context, authKeyId int64, num int32) (*tg.TLFutureSalts, error) {
	pSalts, err := r.getOrNotInsertSaltList(ctx, authKeyId, num)
	if err != nil {
		return nil, err
	}
	salts := tg.MakeTLFutureSalts(&tg.TLFutureSalts{
		ReqMsgId: 0,
		Now:      0,
		Salts:    pSalts,
	})

	return salts, nil
}

func (r *Repository) PutSaltCache(ctx context.Context, authKeyId int64, salt *tg.TLFutureSalt) error {
	return r.FutureSaltsModel.PutSalts(ctx, authKeyId, []*tg.TLFutureSalt{salt}, saltTimeout)
}
