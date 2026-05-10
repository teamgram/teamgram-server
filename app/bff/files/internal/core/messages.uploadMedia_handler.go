// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"math/rand"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/phonenumber"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesUploadMedia
// messages.uploadMedia#14967978 flags:# business_connection_id:flags.0?string peer:InputPeer media:InputMedia = MessageMedia;
func (c *FilesCore) MessagesUploadMedia(in *tg.TLMessagesUploadMedia) (*tg.MessageMedia, error) {
	return c.makeMediaByInputMedia(in.Media)
}

func (c *FilesCore) makeMediaByInputMedia(media tg.InputMediaClazz) (*tg.MessageMedia, error) {
	if media == nil {
		return nil, tg.ErrMediaInvalid
	}

	inputMedia := &tg.InputMedia{Clazz: media}

	switch media.InputMediaClazzName() {
	case tg.ClazzName_inputMediaEmpty:
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaUploadedPhoto:
		uploadedPhoto, ok := inputMedia.ToInputMediaUploadedPhoto()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		photo, err := c.svcCtx.Repo.MediaClient.MediaUploadPhotoFile(c.ctx, &mediapb.TLMediaUploadPhotoFile{
			OwnerId:    c.MD.PermAuthKeyId,
			File:       uploadedPhoto.File,
			Stickers:   nil,
			TtlSeconds: nil,
		})
		if err != nil {
			return nil, err
		}
		if photo == nil || photo.Clazz == nil {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      photo.Clazz,
			TtlSeconds: uploadedPhoto.TtlSeconds,
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaPhoto:
		mediaPhoto, ok := inputMedia.ToInputMediaPhoto()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		inputPhoto, ok := (&tg.InputPhoto{Clazz: mediaPhoto.Id}).ToInputPhoto()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		photo, err := c.svcCtx.Repo.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
			PhotoId: inputPhoto.Id,
		})
		if err != nil {
			return nil, err
		}
		if photo == nil || photo.Clazz == nil {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      photo.Clazz,
			TtlSeconds: mediaPhoto.TtlSeconds,
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaGeoPoint:
		geoPoint, ok := inputMedia.ToInputMediaGeoPoint()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaGeo(&tg.TLMessageMediaGeo{
			Geo: makeGeoPointByInput(geoPoint.GeoPoint),
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaContact:
		contact, ok := inputMedia.ToInputMediaContact()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		messageMedia := tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
			PhoneNumber: contact.PhoneNumber,
			FirstName:   contact.FirstName,
			LastName:    contact.LastName,
			Vcard:       contact.Vcard,
			UserId:      0,
		}).ToMessageMedia()
		phoneNumber, err := phonenumber.CheckAndGetPhoneNumber(contact.PhoneNumber)
		if err == nil {
			userID, err := c.svcCtx.Repo.UserClient.UserGetUserIdByPhone(c.ctx, &userpb.TLUserGetUserIdByPhone{
				Phone: phoneNumber,
			})
			if err != nil {
				c.Logger.Errorf("messages.uploadMedia - user lookup by contact phone failed: %v", err)
			} else if userID != nil {
				messageMedia.Clazz.(*tg.TLMessageMediaContact).UserId = userID.V
			}
		}
		return messageMedia, nil
	case tg.ClazzName_inputMediaUploadedDocument:
		documentMedia, err := c.svcCtx.Repo.MediaClient.MediaUploadedDocumentMedia(c.ctx, &mediapb.TLMediaUploadedDocumentMedia{
			OwnerId: c.MD.PermAuthKeyId,
			Media:   media,
		})
		if err != nil {
			return nil, tg.ErrMediaInvalid
		}
		return documentMedia, nil
	case tg.ClazzName_inputMediaDocument:
		mediaDocument, ok := inputMedia.ToInputMediaDocument()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		inputDocument, ok := (&tg.InputDocument{Clazz: mediaDocument.Id}).ToInputDocument()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		document, err := c.svcCtx.Repo.MediaClient.MediaGetDocument(c.ctx, &mediapb.TLMediaGetDocument{
			Id: inputDocument.Id,
		})
		if err != nil {
			return nil, err
		}
		if document == nil || document.Clazz == nil {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Document:   document.Clazz,
			TtlSeconds: mediaDocument.TtlSeconds,
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaVenue:
		venue, ok := inputMedia.ToInputMediaVenue()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaVenue(&tg.TLMessageMediaVenue{
			Geo:       makeGeoPointByInput(venue.GeoPoint),
			Title:     venue.Title,
			Address:   venue.Address,
			Provider:  venue.Provider,
			VenueId:   venue.VenueId,
			VenueType: venue.VenueType,
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaPhotoExternal,
		tg.ClazzName_inputMediaDocumentExternal,
		tg.ClazzName_inputMediaGame,
		tg.ClazzName_inputMediaInvoice:
		return tg.MakeTLMessageMediaUnsupported(&tg.TLMessageMediaUnsupported{}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaGeoLive:
		geoLive, ok := inputMedia.ToInputMediaGeoLive()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaGeoLive(&tg.TLMessageMediaGeoLive{
			Geo:    makeGeoPointByInput(geoLive.GeoPoint),
			Period: valueOrZero(geoLive.Period),
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaPoll:
		poll, ok := inputMedia.ToInputMediaPoll()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		return tg.MakeTLMessageMediaPoll(&tg.TLMessageMediaPoll{
			Poll:    poll.Poll,
			Results: nil,
		}).ToMessageMedia(), nil
	case tg.ClazzName_inputMediaDice:
		dice, ok := inputMedia.ToInputMediaDice()
		if !ok {
			return nil, tg.ErrMediaInvalid
		}
		return makeDiceMessageMedia(dice.Emoticon), nil
	default:
		return nil, tg.ErrMediaInvalid
	}
}

func makeGeoPointByInput(input tg.InputGeoPointClazz) tg.GeoPointClazz {
	if input == nil {
		return tg.MakeTLGeoPointEmpty(&tg.TLGeoPointEmpty{})
	}

	if inputPoint, ok := (&tg.InputGeoPoint{Clazz: input}).ToInputGeoPoint(); ok {
		return tg.MakeTLGeoPoint(&tg.TLGeoPoint{
			Long:           inputPoint.Long,
			Lat:            inputPoint.Lat,
			AccuracyRadius: inputPoint.AccuracyRadius,
		})
	}

	return tg.MakeTLGeoPointEmpty(&tg.TLGeoPointEmpty{})
}

func makeDiceMessageMedia(emoticon string) *tg.MessageMedia {
	maxValue := int32(6)
	if emoticon == "🏀" {
		maxValue = 5
	}

	return tg.MakeTLMessageMediaDice(&tg.TLMessageMediaDice{
		Value:    rand.Int31()%maxValue + 1,
		Emoticon: emoticon,
	}).ToMessageMedia()
}

func valueOrZero(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}
