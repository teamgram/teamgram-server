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

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/core"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
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

// UserupdatesProcessUserOperationWithEffects
// userupdates.processUserOperationWithEffects operation:UserOperation affected_effects:Vector<AffectedUserOperation> = UserOperationResult;
func (s *Service) UserupdatesProcessUserOperationWithEffects(ctx context.Context, request *userupdates.TLUserupdatesProcessUserOperationWithEffects) (*userupdates.UserOperationResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.processUserOperationWithEffects - request: %s", request)

	r, err := c.UserupdatesProcessUserOperationWithEffects(request)
	if err != nil {
		c.Logger.Errorf("userupdates.processUserOperationWithEffects - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.processUserOperationWithEffects - reply: %s", r)
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

// UserupdatesListDialogs
// userupdates.listDialogs user_id:long top_message_date:long top_peer_seq:long peer_type:int peer_id:long limit:int = DialogProjectionList;
func (s *Service) UserupdatesListDialogs(ctx context.Context, request *userupdates.TLUserupdatesListDialogs) (*userupdates.DialogProjectionList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.listDialogs - request: %s", request)

	r, err := c.UserupdatesListDialogs(request)
	if err != nil {
		c.Logger.Errorf("userupdates.listDialogs - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.listDialogs - reply: %s", r)
	return r, err
}

// UserupdatesGetDialogsByPeers
// userupdates.getDialogsByPeers user_id:long peers:Vector<DialogProjectionPeer> = Vector<DialogProjection>;
func (s *Service) UserupdatesGetDialogsByPeers(ctx context.Context, request *userupdates.TLUserupdatesGetDialogsByPeers) (*userupdates.VectorDialogProjection, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getDialogsByPeers - request: %s", request)

	r, err := c.UserupdatesGetDialogsByPeers(request)
	if err != nil {
		c.Logger.Errorf("userupdates.getDialogsByPeers - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getDialogsByPeers - reply: %s", r)
	return r, err
}

// UserupdatesGetDialogCount
// userupdates.getDialogCount user_id:long = Int32;
func (s *Service) UserupdatesGetDialogCount(ctx context.Context, request *userupdates.TLUserupdatesGetDialogCount) (*tg.Int32, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getDialogCount - request: %s", request)

	r, err := c.UserupdatesGetDialogCount(request)
	if err != nil {
		c.Logger.Errorf("userupdates.getDialogCount - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getDialogCount - reply: %s", r)
	return r, err
}

// UserupdatesGetMessageViewsByPeerSeqs
// userupdates.getMessageViewsByPeerSeqs user_id:long peers:Vector<MessageViewPeerSeq> = MessageViewList;
func (s *Service) UserupdatesGetMessageViewsByPeerSeqs(ctx context.Context, request *userupdates.TLUserupdatesGetMessageViewsByPeerSeqs) (*userupdates.MessageViewList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getMessageViewsByPeerSeqs - request: %s", request)

	r, err := c.UserupdatesGetMessageViewsByPeerSeqs(request)
	if err != nil {
		c.Logger.Errorf("userupdates.getMessageViewsByPeerSeqs - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getMessageViewsByPeerSeqs - reply: %s", r)
	return r, err
}

// UserupdatesGetOutboxReadDate
// userupdates.getOutboxReadDate user_id:long peer_type:int peer_id:long msg_id:int = OutboxReadDate;
func (s *Service) UserupdatesGetOutboxReadDate(ctx context.Context, request *userupdates.TLUserupdatesGetOutboxReadDate) (*tg.OutboxReadDate, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.getOutboxReadDate - request: %s", request)

	r, err := c.UserupdatesGetOutboxReadDate(request)
	if err != nil {
		c.Logger.Errorf("userupdates.getOutboxReadDate - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.getOutboxReadDate - reply: %s", r)
	return r, err
}

// UserupdatesAppendDialogAuthSeqSideEffect
// userupdates.appendDialogAuthSeqSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserAuthSeqAppendResult;
func (s *Service) UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, request *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.appendDialogAuthSeqSideEffect - request: %s", request)

	r, err := c.UserupdatesAppendDialogAuthSeqSideEffect(request)
	if err != nil {
		c.Logger.Errorf("userupdates.appendDialogAuthSeqSideEffect - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.appendDialogAuthSeqSideEffect - reply: %s", r)
	return r, err
}

// UserupdatesAppendDialogPtsSideEffect
// userupdates.appendDialogPtsSideEffect flags:# user_id:long source_perm_auth_key_id:long operation_id:string target_auth_policy:string public_update_type:string peer_type:int peer_id:long payload_schema_version:int payload:bytes payload_hash:bytes = UserPtsAppendResult;
func (s *Service) UserupdatesAppendDialogPtsSideEffect(ctx context.Context, request *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("userupdates.appendDialogPtsSideEffect - request: %s", request)

	r, err := c.UserupdatesAppendDialogPtsSideEffect(request)
	if err != nil {
		c.Logger.Errorf("userupdates.appendDialogPtsSideEffect - error: request: %s, err: %v", request, err)
		return nil, err
	}

	c.Logger.Debugf("userupdates.appendDialogPtsSideEffect - reply: %s", r)
	return r, err
}
