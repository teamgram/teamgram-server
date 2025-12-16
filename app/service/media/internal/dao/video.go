// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/media/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Dao) GetVideoSizeListList(ctx context.Context, idList []int64) (sizes map[int64][]*mtproto.VideoSize) {
	sizes = make(map[int64][]*mtproto.VideoSize)
	if len(idList) == 0 {
		return
	}

	_, _ = m.VideoSizesDAO.SelectListByVideoSizeIdListWithCB(
		ctx,
		idList,
		func(sz, i int, v *dataobject.VideoSizesDO) {
			szList, ok := sizes[v.VideoSizeId]
			if !ok {
				szList = []*mtproto.VideoSize{}
			}

			sz2 := getVideoSize(v)
			if sz2 != nil {
				szList = append(szList, sz2)
				// szList = append(szList, getVideoSize(v))
				sizes[v.VideoSizeId] = szList
			}
		})

	return
}

func (m *Dao) GetVideoSizeList(ctx context.Context, sizeId int64) (sizes []*mtproto.VideoSize) {
	sizes = make([]*mtproto.VideoSize, 0, 2)

	_, _ = m.VideoSizesDAO.SelectListByVideoSizeIdWithCB(
		ctx,
		sizeId,
		func(sz, i int, v *dataobject.VideoSizesDO) {
			sz2 := getVideoSize(v)
			if sz2 != nil {
				sizes = append(sizes, sz2)
			}
		})

	return
}

func getVideoSize(sz *dataobject.VideoSizesDO) *mtproto.VideoSize {
	var (
		videoSize *mtproto.VideoSize
	)

	switch sz.SizeType {
	case "e":
		err := jsonx.UnmarshalFromString(sz.FilePath, &videoSize)
		if err != nil {
			// TODO(@benqi): log
			// return nil
		}
	case "s":
		err := jsonx.UnmarshalFromString(sz.FilePath, &videoSize)
		if err != nil {
			// TODO(@benqi): log
			// return nil
		}
	default:
		videoSize = mtproto.MakeTLVideoSize(&mtproto.VideoSize{
			Type:         sz.SizeType,
			W:            sz.Width,
			H:            sz.Height,
			Size2:        sz.FileSize,
			VideoStartTs: nil,
		}).To_VideoSize()
		if sz.VideoStartTs > 0 {
			videoSize.VideoStartTs = &wrapperspb.DoubleValue{Value: sz.VideoStartTs}
		}
	}
	return videoSize
}

func (m *Dao) SaveVideoSizeV2(ctx context.Context, szId int64, szList []*mtproto.VideoSize) error {
	if len(szList) == 0 {
		return nil
	}

	for _, sz := range szList {
		var (
			szDO *dataobject.VideoSizesDO
		)

		switch sz.GetPredicateName() {
		case "videoSizeEmojiMarkup":
			data, _ := json.Marshal(sz)
			szDO = &dataobject.VideoSizesDO{
				VideoSizeId:  szId,
				SizeType:     "e",
				Width:        0,
				Height:       0,
				FileSize:     0,
				VideoStartTs: 0,
				FilePath:     string(data),
			}
		case "videoSizeStickerMarkup":
			data, _ := json.Marshal(sz)
			szDO = &dataobject.VideoSizesDO{
				VideoSizeId:  szId,
				SizeType:     "s",
				Width:        0,
				Height:       0,
				FileSize:     0,
				VideoStartTs: 0,
				FilePath:     string(data),
			}
		default:
			szDO = &dataobject.VideoSizesDO{
				VideoSizeId:  szId,
				SizeType:     sz.Type,
				Width:        sz.W,
				Height:       sz.H,
				FileSize:     sz.Size2,
				VideoStartTs: sz.GetVideoStartTs().GetValue(),
				FilePath:     fmt.Sprintf("%s/%d.dat", sz.Type, szId),
			}
		}
		if _, _, err := m.VideoSizesDAO.Insert(ctx, szDO); err != nil {
			return err
		}
	}

	return nil
}
