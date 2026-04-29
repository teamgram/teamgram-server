/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/core"
)

var _ *tg.Bool


// UserupdatesProcessUserOperation
// userupdates.processUserOperation operation:UserOperation = UserOperationResult;
func (s *Service) UserupdatesProcessUserOperation(ctx context.Context, request *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.processUserOperation - request: %s", request)

	r, err := c.UserupdatesProcessUserOperation(request)
	if err != nil {
        c.Logger.Errorf("userupdates.processUserOperation - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.processUserOperation - reply: %s", r)
	return r, err
}

// UserupdatesGetOperationResult
// userupdates.getOperationResult user_id:long operation_id:string payload_hash:bytes = UserOperationResult;
func (s *Service) UserupdatesGetOperationResult(ctx context.Context, request *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getOperationResult - request: %s", request)

	r, err := c.UserupdatesGetOperationResult(request)
	if err != nil {
        c.Logger.Errorf("userupdates.getOperationResult - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getOperationResult - reply: %s", r)
	return r, err
}

// UserupdatesGetState
// userupdates.getState user_id:long auth_key_id:long = UserState;
func (s *Service) UserupdatesGetState(ctx context.Context, request *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getState - request: %s", request)

	r, err := c.UserupdatesGetState(request)
	if err != nil {
        c.Logger.Errorf("userupdates.getState - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getState - reply: %s", r)
	return r, err
}

// UserupdatesGetDifference
// userupdates.getDifference flags:# user_id:long auth_key_id:long pts:long pts_total_limit:flags.0?int date:flags.1?long = UserDifference;
func (s *Service) UserupdatesGetDifference(ctx context.Context, request *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getDifference - request: %s", request)

	r, err := c.UserupdatesGetDifference(request)
	if err != nil {
        c.Logger.Errorf("userupdates.getDifference - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getDifference - reply: %s", r)
	return r, err
}

