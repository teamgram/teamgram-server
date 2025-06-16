/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/internal/core"
)

// IdgenNextId
// idgen.nextId = Int64;
func (s *Service) IdgenNextId(ctx context.Context, request *idgen.TLIdgenNextId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.nextId - metadata: {}, request: %v", request)

	r, err := c.IdgenNextId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenNextIds
// idgen.nextIds num:int = Vector<long>;
func (s *Service) IdgenNextIds(ctx context.Context, request *idgen.TLIdgenNextIds) (*idgen.VectorLong, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.nextIds - metadata: {}, request: %v", request)

	r, err := c.IdgenNextIds(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenGetCurrentSeqId
// idgen.getCurrentSeqId key:string = Int64;
func (s *Service) IdgenGetCurrentSeqId(ctx context.Context, request *idgen.TLIdgenGetCurrentSeqId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getCurrentSeqId - metadata: {}, request: %v", request)

	r, err := c.IdgenGetCurrentSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenSetCurrentSeqId
// idgen.setCurrentSeqId key:string id:long = Bool;
func (s *Service) IdgenSetCurrentSeqId(ctx context.Context, request *idgen.TLIdgenSetCurrentSeqId) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.setCurrentSeqId - metadata: {}, request: %v", request)

	r, err := c.IdgenSetCurrentSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenGetNextSeqId
// idgen.getNextSeqId key:string = Int64;
func (s *Service) IdgenGetNextSeqId(ctx context.Context, request *idgen.TLIdgenGetNextSeqId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getNextSeqId - metadata: {}, request: %v", request)

	r, err := c.IdgenGetNextSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenGetNextNSeqId
// idgen.getNextNSeqId key:string n:int = Int64;
func (s *Service) IdgenGetNextNSeqId(ctx context.Context, request *idgen.TLIdgenGetNextNSeqId) (*tg.Int64, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getNextNSeqId - metadata: {}, request: %v", request)

	r, err := c.IdgenGetNextNSeqId(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenGetNextIdValList
// idgen.getNextIdValList id:Vector<InputId> = Vector<IdVal>;
func (s *Service) IdgenGetNextIdValList(ctx context.Context, request *idgen.TLIdgenGetNextIdValList) (*idgen.VectorIdVal, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getNextIdValList - metadata: {}, request: %v", request)

	r, err := c.IdgenGetNextIdValList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}

// IdgenGetCurrentSeqIdList
// idgen.getCurrentSeqIdList id:Vector<InputId> = Vector<IdVal>;
func (s *Service) IdgenGetCurrentSeqIdList(ctx context.Context, request *idgen.TLIdgenGetCurrentSeqIdList) (*idgen.VectorIdVal, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("idgen.getCurrentSeqIdList - metadata: {}, request: %v", request)

	r, err := c.IdgenGetCurrentSeqIdList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("echos.echo - reply: %v", r)
	return r, err
}
