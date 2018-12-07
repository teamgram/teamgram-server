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

package messages

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	message2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
	"time"
)

func makeGeoPointByInput(geoPoint *mtproto.InputGeoPoint) *mtproto.GeoPoint {
	var geo = &mtproto.GeoPoint{Data2: &mtproto.GeoPoint_Data{}}
	switch geoPoint.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputGeoPointEmpty:
		geo.Constructor = mtproto.TLConstructor_CRC32_geoPointEmpty
	case mtproto.TLConstructor_CRC32_inputGeoPoint:
		geo.Data2.Lat = geoPoint.GetData2().Lat
		geo.Data2.Long = geoPoint.GetData2().Long
		geo.Constructor = mtproto.TLConstructor_CRC32_geoPoint
	}
	return geo
}

func (s *MessagesServiceImpl) makeMediaByInputMedia(authKeyId int64, media *mtproto.InputMedia) *mtproto.MessageMedia {
	var (
		now = int32(time.Now().Unix())
		// photoModel = model.GetPhotoModel()
		// uuid = helper.NextSnowflakeId()
	)

	switch media.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputMediaUploadedPhoto:
		uploadedPhoto := media.To_InputMediaUploadedPhoto()
		file := uploadedPhoto.GetFile()

		result, err := document_client.UploadPhotoFile(authKeyId, file)
		// , file.GetData2().GetId(), file.GetData2().GetParts(), file.GetData2().GetName(), file.GetData2().GetMd5Checksum())
		if err != nil {
			glog.Errorf("UploadPhoto error: %v, by %s", err, logger.JsonDebugData(media))
		}

		// fileData := mediaData.GetFile().GetData2()
		photo := &mtproto.TLPhoto{Data2: &mtproto.Photo_Data{
			Id:          result.PhotoId,
			HasStickers: len(uploadedPhoto.GetStickers()) > 0,
			AccessHash:  result.AccessHash, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        now,
			Sizes:       result.SizeList,
		}}

		messageMedia := &mtproto.TLMessageMediaPhoto{Data2: &mtproto.MessageMedia_Data{
			Photo_1: photo.To_Photo(),
			// Caption:    uploadedPhoto.GetCaption(),
			TtlSeconds: uploadedPhoto.GetTtlSeconds(),
		}}
		return messageMedia.To_MessageMedia()

	case mtproto.TLConstructor_CRC32_inputMediaPhoto:
		//inputPhotoEmpty#1cd7bf0d = InputPhoto;
		// inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
		//inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
		mediaPhoto := media.To_InputMediaPhoto()
		sizeList, _ := document_client.GetPhotoSizeList(mediaPhoto.GetId().GetData2().GetId())

		photo := &mtproto.TLPhoto{Data2: &mtproto.Photo_Data{
			Id:          mediaPhoto.GetId().GetData2().GetId(),
			HasStickers: false,
			AccessHash:  mediaPhoto.GetId().GetData2().GetAccessHash(),
			// result.AccessHash, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:  now,
			Sizes: sizeList,
		}}

		messageMedia := &mtproto.TLMessageMediaPhoto{Data2: &mtproto.MessageMedia_Data{
			Photo_1: photo.To_Photo(),
			// Caption:    mediaPhoto.GetCaption(),
			TtlSeconds: mediaPhoto.GetTtlSeconds(),
		}}
		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaGeoPoint:
		// messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
		messageMedia := &mtproto.TLMessageMediaGeo{Data2: &mtproto.MessageMedia_Data{
			Geo: makeGeoPointByInput(media.To_InputMediaGeoPoint().GetGeoPoint()),
		}}

		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaContact:
		// messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
		contact := media.To_InputMediaContact()

		messageMedia := &mtproto.TLMessageMediaContact{Data2: &mtproto.MessageMedia_Data{
			PhoneNumber: contact.GetPhoneNumber(),
			FirstName:   contact.GetFirstName(),
			LastName:    contact.GetLastName(),
			// UserId:      user.GetMyUserByPhoneNumber(contact.GetPhoneNumber()).GetId(),
		}}

		phoneNumber, err := base.CheckAndGetPhoneNumber(contact.GetPhoneNumber())
		if err == nil {
			contactUser := s.UserModel.GetMyUserByPhoneNumber(phoneNumber)
			if contactUser != nil {
				messageMedia.SetUserId(contactUser.Data2.Id)
			}
		}

		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaUploadedDocument:
		// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
		uploadedDocument := media.To_InputMediaUploadedDocument()
		messageMedia, _ := document_client.UploadedDocumentMedia(authKeyId, uploadedDocument)

		return messageMedia.To_MessageMedia()
		// id:InputDocument caption:string ttl_seconds:flags.0?int
	case mtproto.TLConstructor_CRC32_inputMediaDocument:
		// inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
		// document := media.To_InputMediaDocument()
		id := media.To_InputMediaDocument().GetId()
		document3, _ := document_client.GetDocumentById(id.GetData2().GetId(), id.GetData2().GetAccessHash())

		// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
		messageMedia := &mtproto.TLMessageMediaDocument{Data2: &mtproto.MessageMedia_Data{
			Document: document3,
			// Caption:    media.To_InputMediaDocument().GetCaption(),
			TtlSeconds: media.To_InputMediaDocument().GetTtlSeconds(),
		}}

		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaVenue:
		// inputMediaVenue#c13d1c11 geo_point:InputGeoPoint title:string address:string provider:string venue_id:string venue_type:string = InputMedia;
		venue := media.To_InputMediaVenue()

		// messageMediaVenue#2ec0533f geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string = MessageMedia;
		messageMedia := &mtproto.TLMessageMediaVenue{Data2: &mtproto.MessageMedia_Data{
			Geo:       makeGeoPointByInput(venue.GetGeoPoint()),
			Title:     venue.GetTitle(),
			Address:   venue.GetAddress(),
			Provider:  venue.GetProvider(),
			VenueId:   venue.GetVenueId(),
			VenueType: venue.GetVenueType(),
		}}

		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaGifExternal:
		// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;

		// TODO(@benqi): MessageMedia???
		return mtproto.NewTLMessageMediaUnsupported().To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaDocumentExternal:
		// inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;

		// TODO(@benqi): MessageMedia???
		return mtproto.NewTLMessageMediaUnsupported().To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaPhotoExternal:
		// inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;

		// TODO(@benqi): MessageMedia???
		return mtproto.NewTLMessageMediaUnsupported().To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaGame:
		// inputMediaGame#d33f43f3 id:InputGame = InputMedia;
		// game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
		//
		// inputGameID#32c3e77 id:long access_hash:long = InputGame;
		// inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;

		// TODO(@benqi): Not impl inputMediaGame
		return mtproto.NewTLMessageMediaUnsupported().To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaInvoice:
		// inputMediaInvoice#f4e096c3 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string provider_data:DataJSON start_param:string = InputMedia;

		// TODO(@benqi): Not impl inputMediaGame
		return mtproto.NewTLMessageMediaUnsupported().To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaGeoLive:
		// inputMediaGeoLive#7b1a118f geo_point:InputGeoPoint period:int = InputMedia;

		// inputMediaGeoLive#7b1a118f geo_point:InputGeoPoint period:int = InputMedia;
		messageMedia := &mtproto.TLMessageMediaGeoLive{Data2: &mtproto.MessageMedia_Data{
			Geo:    makeGeoPointByInput(media.To_InputMediaGeoLive().GetGeoPoint()),
			Period: media.To_InputMediaGeoLive().GetPeriod(),
		}}

		return messageMedia.To_MessageMedia()
	}

	return mtproto.NewTLMessageMediaEmpty().To_MessageMedia()
}

func (s *MessagesServiceImpl) makeOutboxMessageBySendMedia(authKeyId int64, fromId int32, peer *base.PeerUtil, request *mtproto.TLMessagesSendMedia) *mtproto.TLMessage {
	message := &mtproto.TLMessage{Data2: &mtproto.Message_Data{
		Out:          true,
		Silent:       request.GetSilent(),
		FromId:       fromId,
		ToId:         peer.ToPeer(),
		ReplyToMsgId: request.GetReplyToMsgId(),
		Media:        s.makeMediaByInputMedia(authKeyId, request.GetMedia()),
		ReplyMarkup:  request.GetReplyMarkup(),
		Date:         int32(time.Now().Unix()),
	}}

	// TODO(@benqi): check channel or super chat
	if peer.PeerType == base.PEER_CHANNEL {
		message.SetPost(true)
	}

	return message
}

func (s *MessagesServiceImpl) makeUpdatesByUpdateNewMessage(selfUserId int32, updateNew *mtproto.TLUpdateNewMessage) *mtproto.TLUpdates {
	userIdList, _, _ := message2.PickAllIDListByMessages([]*mtproto.Message{updateNew.GetMessage()})
	userList := s.UserModel.GetUserListByIdList(selfUserId, userIdList)
	return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateNew.To_Update()},
		Users:   userList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}

}

// messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;
func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.sendMedia#c8f16791 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): ???
	// request.NoWebpage
	// request.Background

	// peer
	var (
		peer *base.PeerUtil
		err  error
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMedia#c8f16791 - invalid peer", err)
		return nil, err
	}
	// TODO(@benqi): check user or channels's access_hash

	// peer = helper.FromInputPeer2(md.UserId, request.GetPeer())
	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{
			PeerType: base.PEER_USER,
			PeerId:   md.UserId,
		}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	// handle duplicateMessage
	hasDuplicateMessage, err := s.MessageModel.HasDuplicateMessage(md.UserId, request.GetRandomId())
	if err != nil {
		glog.Error("checkDuplicateMessage error - ", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := s.MessageModel.GetDuplicateMessage(md.UserId, request.GetRandomId())
		if err != nil {
			glog.Error("checkDuplicateMessage error - ", err)
			return nil, err
		}
		return upd, nil
	}

	// 1. draft
	if request.GetClearDraft() {
		s.DoClearDraft(md.UserId, md.AuthId, peer)
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 发件箱
	outboxMessage := s.makeOutboxMessageBySendMedia(md.AuthId, md.UserId, peer, request)

	resultCB := func(pts, ptsCount int32, outBox *message2.MessageBox2) (*mtproto.Updates, error) {
		updateNewMessage := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
			Message_1: outBox.ToMessage(md.UserId),
			Pts:       pts,
			PtsCount:  ptsCount,
		}}
		syncUpdates := 	s.makeUpdatesByUpdateNewMessage(md.UserId, updateNewMessage)

		updateMessageID := &mtproto.TLUpdateMessageID{Data2: &mtproto.Update_Data{
			Id_4:     outBox.MessageId,
			RandomId: request.GetRandomId(),
		}}
		updateList := []*mtproto.Update{updateMessageID.To_Update(), updateNewMessage.To_Update()}

		syncUpdates.SetUpdates(updateList)
		return syncUpdates.To_Updates(), nil
	}

	syncNotMeCB := func(pts, ptsCount int32, outBox *message2.MessageBox2) (int64, *mtproto.Updates, error) {
		updateNewMessage := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
			Message_1: outBox.ToMessage(md.UserId),
			Pts:       pts,
			PtsCount:  ptsCount,
		}}
		syncUpdates := 	s.makeUpdatesByUpdateNewMessage(md.UserId, updateNewMessage)
		return md.AuthId, syncUpdates.To_Updates(), nil
	}

	pushCB := func(pts, ptsCount int32, inBox *message2.MessageBox2) (*mtproto.Updates, error) {
		updateNewMessage := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
			Message_1: inBox.ToMessage(inBox.OwnerId),
			Pts:       pts,
			PtsCount:  ptsCount,
		}}
		pushUpdates := s.makeUpdatesByUpdateNewMessage(inBox.OwnerId, updateNewMessage)
		return pushUpdates.To_Updates(), nil
	}

	replyUpdates, err := s.MessageModel.SendMessage(
		md.UserId,
		peer,
		request.GetRandomId(),
		outboxMessage.To_Message(),
		resultCB,
		syncNotMeCB,
		pushCB)

	if replyUpdates != nil {
		// TODO(@benqi): if err
		s.MessageModel.PutDuplicateMessage(md.UserId, request.GetRandomId(), replyUpdates)
	} else {
		// TODO(@benqi): if err
	}

	glog.Infof("messages.sendMedia#c8f16791 - reply: %s", logger.JsonDebugData(replyUpdates))
	return replyUpdates, err
}
