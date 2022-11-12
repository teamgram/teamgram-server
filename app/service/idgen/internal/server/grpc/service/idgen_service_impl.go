/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/app/service/idgen/internal/core"
)

// IdgenNextId
// idgen.nextId = Int64;
func (s *Service) IdgenNextId(ctx context.Context, request *idgen.TLIdgenNextId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.nextId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenNextId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.nextId - reply: %s", r.DebugString())
	return r, err
}

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (s *Service) IdgenNextIds(ctx context.Context, request *idgen.TLIdgenNextIds) (*idgen.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.nextIds - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenNextIds(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.nextIds - reply: %s", r.DebugString())
	return r, err
}

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (s *Service) IdgenGetCurrentSeqId(ctx context.Context, request *idgen.TLIdgenGetCurrentSeqId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getCurrentSeqId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenGetCurrentSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.getCurrentSeqId - reply: %s", r.DebugString())
	return r, err
}

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (s *Service) IdgenSetCurrentSeqId(ctx context.Context, request *idgen.TLIdgenSetCurrentSeqId) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.setCurrentSeqId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenSetCurrentSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.setCurrentSeqId - reply: %s", r.DebugString())
	return r, err
}

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (s *Service) IdgenGetNextSeqId(ctx context.Context, request *idgen.TLIdgenGetNextSeqId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getNextSeqId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenGetNextSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.getNextSeqId - reply: %s", r.DebugString())
	return r, err
}

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (s *Service) IdgenGetNextNSeqId(ctx context.Context, request *idgen.TLIdgenGetNextNSeqId) (*mtproto.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getNextNSeqId - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.IdgenGetNextNSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("idgen.getNextNSeqId - reply: %s", r.DebugString())
	return r, err
}
