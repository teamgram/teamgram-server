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
	UserupdatesProcessUserOperationWithEffects(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error)
	UserupdatesProcessUserOperationBatch(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperationBatch) (*userupdates.VectorUserOperationResult, error)
	UserupdatesGetOperationResult(ctx context.Context, in *userupdates.TLUserupdatesGetOperationResult) (*userupdates.UserOperationResult, error)
	UserupdatesGetState(ctx context.Context, in *userupdates.TLUserupdatesGetState) (*userupdates.UserState, error)
	UserupdatesGetDifference(ctx context.Context, in *userupdates.TLUserupdatesGetDifference) (*userupdates.UserDifference, error)
	UserupdatesListDialogs(ctx context.Context, in *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error)
	UserupdatesGetDialogsByPeers(ctx context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error)
	UserupdatesGetDialogCount(ctx context.Context, in *userupdates.TLUserupdatesGetDialogCount) (*tg.Int32, error)
	UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error)
	UserupdatesGetOutboxReadDate(ctx context.Context, in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error)
	UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error)
	UserupdatesAppendDialogPtsSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error)
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

// UserupdatesProcessUserOperationWithEffects
// userupdates.processUserOperationWithEffects operation:UserOperation affected_effects:Vector<AffectedUserOperation> = UserOperationResult;
func (m *defaultUserupdatesClient) UserupdatesProcessUserOperationWithEffects(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error) {
	return m.rpc.UserupdatesProcessUserOperationWithEffects(ctx, in)
}

// UserupdatesProcessUserOperationBatch
// userupdates.processUserOperationBatch operations:Vector<UserOperation> = Vector<UserOperationResult>;
func (m *defaultUserupdatesClient) UserupdatesProcessUserOperationBatch(ctx context.Context, in *userupdates.TLUserupdatesProcessUserOperationBatch) (*userupdates.VectorUserOperationResult, error) {
	return m.rpc.UserupdatesProcessUserOperationBatch(ctx, in)
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

// UserupdatesListDialogs
// userupdates.listDialogs user_id:long top_message_date:long top_peer_seq:long peer_type:int peer_id:long limit:int = DialogProjectionList;
func (m *defaultUserupdatesClient) UserupdatesListDialogs(ctx context.Context, in *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error) {
	return m.rpc.UserupdatesListDialogs(ctx, in)
}

// UserupdatesGetDialogsByPeers
// userupdates.getDialogsByPeers user_id:long peers:Vector<DialogProjectionPeer> = Vector<DialogProjection>;
func (m *defaultUserupdatesClient) UserupdatesGetDialogsByPeers(ctx context.Context, in *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	return m.rpc.UserupdatesGetDialogsByPeers(ctx, in)
}

// UserupdatesGetDialogCount
// userupdates.getDialogCount user_id:long = Int32;
func (m *defaultUserupdatesClient) UserupdatesGetDialogCount(ctx context.Context, in *userupdates.TLUserupdatesGetDialogCount) (*tg.Int32, error) {
	return m.rpc.UserupdatesGetDialogCount(ctx, in)
}

// UserupdatesGetMessageViewsByPeerSeqs
// userupdates.getMessageViewsByPeerSeqs user_id:long peers:Vector<MessageViewPeerSeq> = MessageViewList;
func (m *defaultUserupdatesClient) UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, in *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
	return m.rpc.UserupdatesGetMessageViewsByPeerSeqs(ctx, in)
}

// UserupdatesGetOutboxReadDate
// userupdates.getOutboxReadDate user_id:long peer_type:int peer_id:long msg_id:int = OutboxReadDate;
func (m *defaultUserupdatesClient) UserupdatesGetOutboxReadDate(ctx context.Context, in *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
	return m.rpc.UserupdatesGetOutboxReadDate(ctx, in)
}

// UserupdatesAppendDialogAuthSeqSideEffect
// userupdates.appendDialogAuthSeqSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserAuthSeqAppendResult;
func (m *defaultUserupdatesClient) UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error) {
	return m.rpc.UserupdatesAppendDialogAuthSeqSideEffect(ctx, in)
}

// UserupdatesAppendDialogPtsSideEffect
// userupdates.appendDialogPtsSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserPtsAppendResult;
func (m *defaultUserupdatesClient) UserupdatesAppendDialogPtsSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error) {
	return m.rpc.UserupdatesAppendDialogPtsSideEffect(ctx, in)
}
