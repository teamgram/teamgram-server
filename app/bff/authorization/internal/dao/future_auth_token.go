// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package dao

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/hash"
)

const (
	futureAuthTokenTimeout = 7 * 60 * 60 * 24 // salt timeout
	futureAuthTokenPrefix  = "future_auth_token_"
)

func (d *Dao) PutFutureAuthToken(ctx context.Context, futureAuthToken []byte, authKeyId int64) error {
	k := futureAuthTokenPrefix + hash.Md5Hex(futureAuthToken)
	return d.kv.SetexCtx(ctx, k, strconv.FormatInt(authKeyId, 10), futureAuthTokenTimeout)
}

func (d *Dao) GetFutureAuthToken(ctx context.Context, futureAuthToken []byte) (int64, error) {
	k := futureAuthTokenPrefix + hash.Md5Hex(futureAuthToken)
	rV, err := d.kv.GetCtx(ctx, k)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(rV, 10, 64)
}

func (d *Dao) DelFutureAuthToken(ctx context.Context, futureAuthToken []byte) error {
	k := futureAuthTokenPrefix + hash.Md5Hex(futureAuthToken)
	_, err := d.kv.DelCtx(ctx, k)
	return err
}
