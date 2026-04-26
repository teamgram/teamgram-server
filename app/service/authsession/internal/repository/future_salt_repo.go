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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/xkv"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	saltTimeout = 30 * 60 // salt timeout
)

func toFutureSaltRecord(salt *tg.TLFutureSalt) *xkv.FutureSaltRecord {
	if salt == nil {
		return nil
	}

	return &xkv.FutureSaltRecord{
		ValidSince: salt.ValidSince,
		ValidUntil: salt.ValidUntil,
		Salt:       salt.Salt,
	}
}

func toFutureSaltRecords(salts []*tg.TLFutureSalt) []*xkv.FutureSaltRecord {
	records := make([]*xkv.FutureSaltRecord, 0, len(salts))
	for _, salt := range salts {
		if record := toFutureSaltRecord(salt); record != nil {
			records = append(records, record)
		}
	}
	return records
}

func toTLFutureSalt(record *xkv.FutureSaltRecord) *tg.TLFutureSalt {
	if record == nil {
		return nil
	}

	return tg.MakeTLFutureSalt(&tg.TLFutureSalt{
		ValidSince: record.ValidSince,
		ValidUntil: record.ValidUntil,
		Salt:       record.Salt,
	})
}

func toTLFutureSalts(records []*xkv.FutureSaltRecord) []*tg.TLFutureSalt {
	salts := make([]*tg.TLFutureSalt, 0, len(records))
	for _, record := range records {
		if salt := toTLFutureSalt(record); salt != nil {
			salts = append(salts, salt)
		}
	}
	return salts
}

func (r *Repository) getOrNotInsertSaltList(ctx context.Context, keyId int64, size int32) ([]*tg.TLFutureSalt, error) {
	var (
		salts = make([]*tg.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		saltsData      []*tg.TLFutureSalt
		lastSalt       *tg.TLFutureSalt
	)

	saltRecords, err := r.futureSaltsModel.GetSalts(ctx, keyId)
	if err != nil {
		return nil, err
	}
	saltList := toTLFutureSalts(saltRecords)

	if len(saltList) > 0 {
		hasLastSalt := false
		for idx, salt := range saltList {
			if salt.ValidUntil >= date {
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

		lastSalt = previousSaltForCompatibility(lastSalt, date)
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < left; i++ {
			saltValue, err := secureRandInt63()
			if err != nil {
				return nil, err
			}
			salt := tg.MakeTLFutureSalt(&tg.TLFutureSalt{
				ValidSince: lastValidUntil,
				ValidUntil: lastValidUntil + saltTimeout,
				Salt:       saltValue,
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
		err = r.futureSaltsModel.PutSalts(ctx, keyId, toFutureSaltRecords(saltsData2))
		if err != nil {
			return nil, err
		}
	}
	return salts2, nil
}

func previousSaltForCompatibility(lastSalt *tg.TLFutureSalt, now int32) *tg.TLFutureSalt {
	if lastSalt == nil {
		return nil
	}

	// MTProto clients can legitimately send messages from the previous salt
	// window during clock skew or transport delay. Keep one just-expired salt
	// for a short grace period, but drop older salts from the response/cache.
	if lastSalt.ValidUntil+300 < now {
		return nil
	}
	return lastSalt
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
	return r.futureSaltsModel.PutSalts(ctx, authKeyId, []*xkv.FutureSaltRecord{toFutureSaltRecord(salt)})
}
