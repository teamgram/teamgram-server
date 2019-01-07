// Copyright (c) 2019-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package langpack

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// langpack.getStrings#2e1ee318 lang_code:string keys:Vector<string> = Vector<LangPackString>;
func (s *LangpackServiceImpl) LangpackGetStringsLayer68(ctx context.Context, request *mtproto.TLLangpackGetStringsLayer68) (*mtproto.Vector_LangPackString, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("langpack.getStringsLayer68 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	langPacks := queryLangPacks(request.GetLangCode())
	langpackStrings := &mtproto.Vector_LangPackString{}
	for _, s := range request.Keys {
		s2 := &mtproto.TLLangPackString{Data2: &mtproto.LangPackString_Data{
			Key:   s,
			Value: langPacks.Query(s),
		}}
		langpackStrings.Datas = append(langpackStrings.Datas, s2.To_LangPackString())
	}

	glog.Infof("langpack.getStrings#2e1ee318 - reply: %s", logger.JsonDebugData(langpackStrings))
	return langpackStrings, nil
}
