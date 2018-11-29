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

package sticker

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

//public static final int TYPE_IMAGE = 0;
//public static final int TYPE_MASK = 1;
//public static final int TYPE_FAVE = 2;

//type Logic struct {
//
//}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;

// stickerSet#5585a139 flags:# archived:flags.1?true official:flags.2?true masks:flags.3?true installed_date:flags.0?int id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;

func makeStickerSet(do *dataobject.StickerSetsDO) *mtproto.StickerSet {
	sitckers := &mtproto.TLStickerSet{Data2: &mtproto.StickerSet_Data{
		// Installed:  true,
		Id:         do.StickerSetId,
		AccessHash: do.AccessHash,
		Title:      do.Title,
		ShortName:  do.ShortName,
		Hash:       do.Hash,
	}}
	return sitckers.To_StickerSet()
}

func (m *StickerModel) GetStickerSetList(hash int32) []*mtproto.StickerSet {
	//
	doList := m.dao.StickerSetsDAO.SelectAll()
	stickers := make([]*mtproto.StickerSet, len(doList))
	for i := 0; i < len(doList); i++ {
		stickers[i] = makeStickerSet(&doList[i])
	}
	return stickers
}

func (m *StickerModel) GetStickerSet(stickerset *mtproto.InputStickerSet) *mtproto.StickerSet {
	var (
		inputSet = stickerset.GetData2()
		set      *mtproto.StickerSet
	)

	switch stickerset.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputStickerSetID:
		do := m.dao.StickerSetsDAO.SelectByID(inputSet.GetId(), inputSet.GetAccessHash())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.TLConstructor_CRC32_inputStickerSetShortName:
		do := m.dao.StickerSetsDAO.SelectByShortName(inputSet.GetShortName())
		if do != nil {
			set = makeStickerSet(do)
		}
	case mtproto.TLConstructor_CRC32_inputStickerSetEmpty:
		glog.Error("stickerset is inputStickerSetEmpty")
	}

	return set
}

func (m *StickerModel) GetStickerPackList(setId int64) ([]*mtproto.StickerPack, []int64) {
	doList := m.dao.StickerPacksDAO.SelectBySetID(setId)
	packs := make([]*mtproto.StickerPack, len(doList))
	idList := make([]int64, len(doList))
	for i := 0; i < len(doList); i++ {
		packs[i] = &mtproto.StickerPack{
			Constructor: mtproto.TLConstructor_CRC32_stickerPack,
			Data2: &mtproto.StickerPack_Data{
				Emoticon:  doList[i].Emoticon,
				Documents: []int64{doList[i].DocumentId},
			},
		}
		idList[i] = doList[i].DocumentId
	}
	return packs, idList
}
