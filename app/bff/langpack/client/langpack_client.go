/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package langpack_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type LangpackClient interface {
	LangpackGetLangPack(ctx context.Context, in *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error)
	LangpackGetStrings(ctx context.Context, in *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error)
	LangpackGetDifference(ctx context.Context, in *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error)
	LangpackGetLanguages(ctx context.Context, in *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error)
	LangpackGetLanguage(ctx context.Context, in *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error)
}

type defaultLangpackClient struct {
	cli zrpc.Client
}

func NewLangpackClient(cli zrpc.Client) LangpackClient {
	return &defaultLangpackClient{
		cli: cli,
	}
}

// LangpackGetLangPack
// langpack.getLangPack#f2f2330a lang_pack:string lang_code:string = LangPackDifference;
func (m *defaultLangpackClient) LangpackGetLangPack(ctx context.Context, in *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	client := mtproto.NewRPCLangpackClient(m.cli.Conn())
	return client.LangpackGetLangPack(ctx, in)
}

// LangpackGetStrings
// langpack.getStrings#efea3803 lang_pack:string lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (m *defaultLangpackClient) LangpackGetStrings(ctx context.Context, in *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	client := mtproto.NewRPCLangpackClient(m.cli.Conn())
	return client.LangpackGetStrings(ctx, in)
}

// LangpackGetDifference
// langpack.getDifference#cd984aa5 lang_pack:string lang_code:string from_version:int = LangPackDifference;
func (m *defaultLangpackClient) LangpackGetDifference(ctx context.Context, in *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	client := mtproto.NewRPCLangpackClient(m.cli.Conn())
	return client.LangpackGetDifference(ctx, in)
}

// LangpackGetLanguages
// langpack.getLanguages#42c6978f lang_pack:string = Vector<LangPackLanguage>;
func (m *defaultLangpackClient) LangpackGetLanguages(ctx context.Context, in *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	client := mtproto.NewRPCLangpackClient(m.cli.Conn())
	return client.LangpackGetLanguages(ctx, in)
}

// LangpackGetLanguage
// langpack.getLanguage#6a596502 lang_pack:string lang_code:string = LangPackLanguage;
func (m *defaultLangpackClient) LangpackGetLanguage(ctx context.Context, in *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error) {
	client := mtproto.NewRPCLangpackClient(m.cli.Conn())
	return client.LangpackGetLanguage(ctx, in)
}
