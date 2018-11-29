// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package account

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
)

/*
	wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
	wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
*/

// account.getWallPapers#c04cfac2 = Vector<WallPaper>;
func (s *AccountServiceImpl) AccountGetWallPapers(ctx context.Context, request *mtproto.TLAccountGetWallPapers) (*mtproto.Vector_WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getWallPapers#c04cfac2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	//
	wallDataList := s.AccountModel.GetWallPaperList()

	walls := &mtproto.Vector_WallPaper{
		Datas: make([]*mtproto.WallPaper, 0, len(wallDataList)),
	}

	for _, wallData := range wallDataList {
		if wallData.Type == 0 {
			szList, _ := document_client.GetPhotoSizeList(wallData.PhotoId)
			wall := &mtproto.TLWallPaper{Data2: &mtproto.WallPaper_Data{
				Id:    wallData.Id,
				Title: wallData.Title,
				Sizes: szList,
				Color: wallData.Color,
			}}
			walls.Datas = append(walls.Datas, wall.To_WallPaper())
		} else {
			wall := &mtproto.TLWallPaperSolid{Data2: &mtproto.WallPaper_Data{
				Id:      wallData.Id,
				Title:   wallData.Title,
				Color:   wallData.Color,
				BgColor: wallData.BgColor,
			}}
			walls.Datas = append(walls.Datas, wall.To_WallPaper())
		}
	}

	glog.Infof("account.getWallPapers#c04cfac2 - reply: %s", logger.JsonDebugData(walls))
	return walls, nil
}
