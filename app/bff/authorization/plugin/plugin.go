// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package plugin

import (
	"context"
)

type AuthorizationPlugin interface {
	OnAuthLogout(ctx context.Context, userId int64, keys ...int64) error
	OnAuthAction(ctx context.Context, authKeyId, msgId int64, clientIp string, phoneNumber string, actionType int, log string)
	CheckPhoneNumberBanned(ctx context.Context, phoneNumber string) (bool, error)
	CheckSessionPasswordNeeded(ctx context.Context, userId int64) bool
}
