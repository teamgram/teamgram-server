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

package core

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/app/service/media/media"
	"github.com/teamgram/teamgram-server/pkg/phonenumber"
	"math/rand"
	"time"
)

// MessagesUploadMedia
// messages.uploadMedia#519bc2b1 peer:InputPeer media:InputMedia = MessageMedia;
func (c *FilesCore) MessagesUploadMedia(in *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	rValue, err := c.makeMediaByInputMedia(in.GetMedia())
	if err != nil {
		c.Logger.Errorf("messages.uploadMedia - error: %v", err)
		return nil, err
	}

	return rValue, nil
}

func (c *FilesCore) makeMediaByInputMedia(media *mtproto.InputMedia) (messageMedia *mtproto.MessageMedia, err error) {
	var (
		now = int32(time.Now().Unix())
	)

	switch media.PredicateName {
	case mtproto.Predicate_inputMediaEmpty:
		// inputMediaEmpty#9664f57f = InputMedia;

		messageMedia = mtproto.MakeTLMessageMediaEmpty(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaUploadedPhoto:
		// inputMediaUploadedPhoto#1e287d04 flags:#
		//	file:InputFile
		//	stickers:flags.0?Vector<InputDocument>
		//	ttl_seconds:flags.1?int = InputMedia;

		var (
			photo *mtproto.Photo
		)
		photo, err = c.svcCtx.Dao.MediaClient.MediaUploadPhotoFile(c.ctx, &mediapb.TLMediaUploadPhotoFile{
			OwnerId:    c.MD.PermAuthKeyId,
			File:       media.File,
			Stickers:   nil,
			TtlSeconds: nil,
		})
		if err != nil {
			c.Logger.Errorf("UploadPhoto error: %v, by %s", err, media)
			return
		}

		messageMedia = mtproto.MakeTLMessageMediaPhoto(&mtproto.MessageMedia{
			Photo_FLAGPHOTO: photo,
			TtlSeconds:      media.TtlSeconds,
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaPhoto:
		// inputMediaPhoto#b3ba0635 flags:#
		//	id:InputPhoto
		//	ttl_seconds:flags.0?int = InputMedia;

		mediaPhoto := media.To_InputMediaPhoto()
		sizeList, _ := c.svcCtx.Dao.MediaClient.MediaGetPhotoSizeList(c.ctx, &mediapb.TLMediaGetPhotoSizeList{
			SizeId: mediaPhoto.GetId_INPUTPHOTO().GetId(),
		})

		photo := mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          mediaPhoto.GetId_INPUTPHOTO().GetId(),
			HasStickers: false,
			AccessHash:  mediaPhoto.GetId_INPUTPHOTO().GetAccessHash(),
			Date:        now,
			Sizes:       sizeList.Sizes,
			DcId:        sizeList.DcId,
		})

		messageMedia = mtproto.MakeTLMessageMediaPhoto(&mtproto.MessageMedia{
			Photo_FLAGPHOTO: photo.To_Photo(),
			TtlSeconds:      mediaPhoto.GetTtlSeconds(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaGeoPoint:
		// inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;

		messageMedia = mtproto.MakeTLMessageMediaGeo(&mtproto.MessageMedia{
			Geo: mtproto.MakeGeoPointByInput(media.To_InputMediaGeoPoint().GetGeoPoint()),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaContact:
		// inputMediaContact#f8ab7dfb
		//	phone_number:string
		//	first_name:string
		//	last_name:string
		//	vcard:string = InputMedia;

		contact := media.To_InputMediaContact()

		messageMedia = mtproto.MakeTLMessageMediaContact(&mtproto.MessageMedia{
			PhoneNumber: contact.GetPhoneNumber(),
			FirstName:   contact.GetFirstName(),
			LastName:    contact.GetLastName(),
			Vcard:       contact.GetVcard(),
			UserId:      0,
		}).To_MessageMedia()

		phoneNumber, err := phonenumber.CheckAndGetPhoneNumber(contact.GetPhoneNumber())
		if err == nil {
			contactUser, _ := c.svcCtx.Dao.UserClient.UserGetImmutableUserByPhone(c.ctx, &userpb.TLUserGetImmutableUserByPhone{
				Phone: phoneNumber,
			})
			if contactUser != nil {
				messageMedia.UserId = contactUser.Id()
			}
		}
	case mtproto.Predicate_inputMediaUploadedDocument:
		// inputMediaUploadedDocument#5b38c6c1 flags:#
		//	nosound_video:flags.3?true
		//	force_file:flags.4?true
		//	file:InputFile
		//	thumb:flags.2?InputFile
		//	mime_type:string
		//	attributes:Vector<DocumentAttribute>
		//	stickers:flags.0?Vector<InputDocument>
		//	ttl_seconds:flags.1?int = InputMedia;
		documentMedia, err2 := c.svcCtx.Dao.MediaClient.MediaUploadedDocumentMedia(c.ctx, &mediapb.TLMediaUploadedDocumentMedia{
			OwnerId: c.MD.PermAuthKeyId,
			Media:   media,
		})
		if err2 != nil {
			err = mtproto.ErrMediaInvalid
			return
		}
		messageMedia = documentMedia
	case mtproto.Predicate_inputMediaDocument:
		// inputMediaDocument#33473058 flags:#
		//	id:InputDocument
		//	ttl_seconds:flags.0?int
		//	query:flags.1?string = InputMedia;

		id := media.To_InputMediaDocument().GetId_INPUTDOCUMENT()
		document3, _ := c.svcCtx.Dao.MediaClient.MediaGetDocument(c.ctx, &mediapb.TLMediaGetDocument{
			Id: id.GetId(),
		})

		// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
		messageMedia = mtproto.MakeTLMessageMediaDocument(&mtproto.MessageMedia{
			Document: document3,
			// Caption:    media.To_InputMediaDocument().GetCaption(),
			TtlSeconds: media.To_InputMediaDocument().GetTtlSeconds(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaVenue:
		// inputMediaVenue#c13d1c11
		//	geo_point:InputGeoPoint
		//	title:string
		//	address:string
		//	provider:string
		//	venue_id:string
		//	venue_type:string = InputMedia;
		venue := media.To_InputMediaVenue()
		// messageMediaVenue#2ec0533f geo:GeoPoint title:string address:string provider:string venue_id:string venue_type:string = MessageMedia;
		messageMedia = mtproto.MakeTLMessageMediaVenue(&mtproto.MessageMedia{
			Geo:       mtproto.MakeGeoPointByInput(venue.GetGeoPoint()),
			Title:     venue.GetTitle(),
			Address:   venue.GetAddress(),
			Provider:  venue.GetProvider_STRING(),
			VenueId:   venue.GetVenueId(),
			VenueType: venue.GetVenueType(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaPhotoExternal:
		// inputMediaPhotoExternal#e5bbfe1a flags:# url:string ttl_seconds:flags.0?int = InputMedia;

		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaDocumentExternal:
		// TODO(@benqi): MessageMedia???
		// inputMediaDocumentExternal#fb52dc99 flags:# url:string ttl_seconds:flags.0?int = InputMedia;
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaGame:
		// inputMediaGame#d33f43f3 id:InputGame = InputMedia;

		// TODO(@benqi): Not impl inputMediaGame
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaInvoice:
		// inputMediaInvoice#d9799874 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string provider_data:DataJSON start_param:flags.1?string = InputMedia;

		// TODO(@benqi): Not impl inputMediaGame
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaGeoLive:
		// inputMediaGeoLive#971fa843 flags:# stopped:flags.0?true geo_point:InputGeoPoint heading:flags.2?int period:flags.1?int proximity_notification_radius:flags.3?int = InputMedia;

		messageMedia = mtproto.MakeTLMessageMediaGeoLive(&mtproto.MessageMedia{
			Geo:    mtproto.MakeGeoPointByInput(media.To_InputMediaGeoLive().GetGeoPoint()),
			Period: media.To_InputMediaGeoLive().GetPeriod().GetValue(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaPoll:
		// inputMediaPoll#f94e5f1 flags:# poll:Poll correct_answers:flags.0?Vector<bytes> solution:flags.1?string solution_entities:flags.1?Vector<MessageEntity> = InputMedia;
		messageMedia = mtproto.MakeTLMessageMediaPoll(&mtproto.MessageMedia{
			Poll:    media.Poll,
			Results: nil,
		}).To_MessageMedia()

	case mtproto.Predicate_inputMediaDice:
		// inputMediaDice#e66fbf7b emoticon:string = InputMedia;

		if media.Emoticon == "🎲" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else if media.Emoticon == "🎯" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else if media.Emoticon == "🏀" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%5 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		}

	default:
		err = mtproto.ErrMediaInvalid
	}

	return
}
