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

package core

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/mtproto"
	base2 "github.com/nebula-chat/chatengine/pkg/util"
	idgen "github.com/nebula-chat/chatengine/service/idgen/client"
)

const (
	TOKEN_TYPE_APNS         = 1
	TOKEN_TYPE_GCM          = 2
	TOKEN_TYPE_MPNS         = 3
	TOKEN_TYPE_SIMPLE_PUSH  = 4
	TOKEN_TYPE_UBUNTU_PHONE = 5
	TOKEN_TYPE_BLACKBERRY   = 6
	// Android里使用
	TOKEN_TYPE_INTERNAL_PUSH = 7
	// web
	TOKEN_TYPE_WEB_PUSH = 10
	TOKEN_TYPE_MAXSIZE  = 10
)

type CoreModel interface {
	InstallModel()
	RegisterCallback(cb interface{})
}

// type Instance func() Initializer
var uuidGen idgen.UUIDGen

func GetUUID() (uuid int64) {
	uuid, _ = uuidGen.GetUUID()
	return
}

var models []CoreModel

func RegisterCoreModel(model CoreModel) {
	models = append(models, model)
}

// 必须在mysql／redis等依赖安装完后才能执行
func InstallCoreModels(serverId int32, inited func()) []CoreModel {
	if inited != nil {
		inited()
	}

	uuidGen, _ = idgen.NewUUIDGen("snowflake", base2.Int32ToString(serverId))
	initSeqIDGen("cache")

	for _, m := range models {
		m.InstallModel()

		for _, m2 := range models {
			if m != m2 {
				m.RegisterCallback(m2)
			}
		}
	}

	return models
}

type AccountCallback interface {
	CheckAllowCalls(selfId, userId int32, isContact bool) bool
	CheckAllowChatInvites(selfId, userId int32, isContact bool) bool
	CheckShowStatus(selfId, userId int32, isContact bool) bool
	GetNotifySettings(selfUserId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings
}

type NotifySettingCallback interface {
	GetNotifySettings(selfUserId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings
}

type PhotoCallback interface {
	GetUserProfilePhoto(photoId int64) *mtproto.UserProfilePhoto
	GetChatPhoto(photoId int64) *mtproto.ChatPhoto
	GetPhoto(photoId int64) *mtproto.Photo
}

type ContactCallback interface {
	GetContactAndMutual(selfUserId, id int32) (bool, bool)
}

type DialogCallback interface {
	InsertOrUpdateDialog(userId, peerType, peerId, topMessage int32, hasMentioned, isInbox bool)
	InsertOrChannelUpdateDialog(userId, peerType, peerId int32)
}

type UsernameCallback interface {
	GetAccountUsername(userId int32) string
	GetChannelUsername(channelId int32) string
}

// dialog#e4def5db flags:#
// 	pinned:flags.2?true
// 	unread_mark:flags.3?true
// 	peer:Peer
// 	top_message:int
// 	read_inbox_max_id:int
// 	read_outbox_max_id:int
// 	unread_count:int
// 	unread_mentions_count:int
// 	notify_settings:PeerNotifySettings
// 	pts:flags.0?int
// 	draft:flags.1?DraftMessage = Dialog;
//
type ChannelCallback interface {
	GetTopMessageListByIdList(idList []int32) (topMessages map[int32]int32)
}

///*
//	// TODO(@benqi): chat notifySetting...
//	//if notifySettingFunc == nil {
//	//	notifySettings = &mtproto.PeerNotifySettings{
//	//		Constructor: mtproto.TLConstructor_CRC32_peerNotifySettings,
//	//		Data2: &mtproto.PeerNotifySettings_Data{
//	//			ShowPreviews: true,
//	//			Silent:       false,
//	//			MuteUntil:    0,
//	//			Sound:        "default",
//	//		},
//	//	}
//	//} else {
//	notifySettings := cb1(selfUserId, peer)
//	//}
// */
//type GetNotifySettingsCallback func (selfUserId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings
//type GetUserProfilePhotoCallback func (photoId int64) *mtproto.UserProfilePhoto
//type GetChatPhotoCallback func (photoId int64) *mtproto.ChatPhoto
//
///*
//	// TODO(@benqi):
//	if channelData.GetPhotoId() == 0 {
//		photoEmpty := &mtproto.TLPhotoEmpty{Data2: &mtproto.Photo_Data{
//			Id: 0,
//		}}
//		photo = photoEmpty.To_Photo()
//	} else {
//		//channelPhoto := &mtproto.TLPhoto{ Data2: &mtproto.Photo_Data{
//		//	Id:          channelData.channel.PhotoId,
//		//	HasStickers: false,
//		//	AccessHash:  channelData.channel.PhotoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
//		//	Date:        int32(time.Now().Unix()),
//		//	Sizes:       sizes,
//		//}}
//		photo = cb2(channelData.channel.PhotoId)
//		// channelPhoto.To_Photo()
//	}
// */
//type GetPhotoCallback func (photoId int64) *mtproto.Photo
//
////func MakeUserProfilePhoto(photoId int64, sizes []*mtproto.PhotoSize) *mtproto.UserProfilePhoto {
////	if len(sizes) == 0 {
////		return mtproto.NewTLUserProfilePhotoEmpty().To_UserProfilePhoto()
////	}
////
////	// TODO(@benqi): check PhotoSize is photoSizeEmpty
////	photo := &mtproto.TLUserProfilePhoto{Data2: &mtproto.UserProfilePhoto_Data{
////		PhotoId: photoId,
////		PhotoSmall: sizes[0].GetData2().GetLocation(),
////		PhotoBig: sizes[len(sizes)-1].GetData2().GetLocation(),
////	}}
////
////	return photo.To_UserProfilePhoto()
////}
////
////func MakeChatPhoto(sizes []*mtproto.PhotoSize) *mtproto.ChatPhoto {
////	if len(sizes) == 0 {
////		return mtproto.NewTLChatPhotoEmpty().To_ChatPhoto()
////	}
////
////	// TODO(@benqi): check PhotoSize is photoSizeEmpty
////	photo := &mtproto.TLChatPhoto{Data2: &mtproto.ChatPhoto_Data{
////		PhotoSmall: sizes[0].GetData2().GetLocation(),
////		PhotoBig: sizes[len(sizes)-1].GetData2().GetLocation(),
////	}}
////
////	return photo.To_ChatPhoto()
////}
//
//type CheckContactAndMutualCallback func (selfUserId, id int32) (bool, bool)
