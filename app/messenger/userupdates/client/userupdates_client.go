/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package userupdatesclient

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates/userupdatesservice"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

	"github.com/cloudwego/kitex/client"
)

var _ *tg.Bool

type UserupdatesClient interface {
	UserupdatesProcessUserOperation(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error)
	UserupdatesGetOperationResult(ctx context.Context, in *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error)
	UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error)
	UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error)
}

type defaultUserupdatesClient struct {
	cli client.Client
	rpc userupdatesservice.Client
}

func NewUserupdatesClient(cli client.Client) UserupdatesClient {
	return &defaultUserupdatesClient{
		cli: cli,
		rpc: userupdatesservice.NewRPCUserupdatesClient(cli),
	}
}

// UserupdatesProcessUserOperation
// userupdates.processUserOperation operation:UserOperation = UserOperationResult;
func (m *defaultUserupdatesClient) UserupdatesProcessUserOperation(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperation) (*userupdates.UserOperationResult, error) {
	return m.rpc.UserupdatesProcessUserOperation(ctx, in)
}

// UserupdatesGetOperationResult
// userupdates.getOperationResult user_id:long operation_id:string payload_hash:bytes = UserOperationResult;
func (m *defaultUserupdatesClient) UserupdatesGetOperationResult(ctx context.Context, in *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error) {
	return m.rpc.UserupdatesGetOperationResult(ctx, in)
}

// UserupdatesGetState
// userupdates.getState user_id:long auth_key_id:long = UserState;
func (m *defaultUserupdatesClient) UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error) {
	return m.rpc.UserupdatesGetState(ctx, in)
}

// UserupdatesGetDifference
// userupdates.getDifference flags:# user_id:long auth_key_id:long pts:long pts_total_limit:flags.0?int date:flags.1?long = UserDifference;
func (m *defaultUserupdatesClient) UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error) {
	return m.rpc.UserupdatesGetDifference(ctx, in)
}
