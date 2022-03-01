/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/langpack/internal/core"
)

// LangpackGetLangPack
// langpack.getLangPack#f2f2330a lang_pack:string lang_code:string = LangPackDifference;
func (s *Service) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("langpack.getLangPack - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.LangpackGetLangPack(request)
	if err != nil {
		return nil, err
	}

	c.Infof("langpack.getLangPack - reply: %s", r.DebugString())
	return r, err
}

// LangpackGetStrings
// langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (s *Service) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("langpack.getStrings - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.LangpackGetStrings(request)
	if err != nil {
		return nil, err
	}

	c.Infof("langpack.getStrings - reply: %s", r.DebugString())
	return r, err
}

// LangpackGetDifference
// langpack.getDifference#cd984aa5 lang_pack:string lang_code:string from_version:int = LangPackDifference;
func (s *Service) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("langpack.getDifference - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.LangpackGetDifference(request)
	if err != nil {
		return nil, err
	}

	c.Infof("langpack.getDifference - reply: %s", r.DebugString())
	return r, err
}

// LangpackGetLanguages
// langpack.getLanguages#42c6978f lang_pack:string = Vector<LangPackLanguage>;
func (s *Service) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("langpack.getLanguages - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.LangpackGetLanguages(request)
	if err != nil {
		return nil, err
	}

	c.Infof("langpack.getLanguages - reply: %s", r.DebugString())
	return r, err
}

// LangpackGetLanguage
// langpack.getLanguage#6a596502 lang_pack:string lang_code:string = LangPackLanguage;
func (s *Service) LangpackGetLanguage(ctx context.Context, request *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("langpack.getLanguage - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.LangpackGetLanguage(request)
	if err != nil {
		return nil, err
	}

	c.Infof("langpack.getLanguage - reply: %s", r.DebugString())
	return r, err
}
