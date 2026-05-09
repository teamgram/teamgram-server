package core

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func buildMediaOutbox(in *tg.TLMessagesSendMedia, selfUserID, peerUserID int64, media tg.MessageMediaClazz, replyTo tg.MessageReplyHeaderClazz) (*msg.TLOutboxMessage, int32) {
	date := int32(time.Now().Unix())
	return buildMessageMediaOutbox(mediaOutboxInput{
		RandomId:    in.RandomId,
		Background:  in.Background,
		Silent:      in.Silent,
		Noforwards:  in.Noforwards,
		InvertMedia: in.InvertMedia,
		FromId:      selfUserID,
		PeerId:      peerUserID,
		ReplyTo:     replyTo,
		Date:        date,
		Message:     in.Message,
		Media:       media,
		Entities:    in.Entities,
	}), date
}

type mediaOutboxInput struct {
	RandomId    int64
	Background  bool
	Silent      bool
	Noforwards  bool
	InvertMedia bool
	FromId      int64
	PeerId      int64
	ReplyTo     tg.MessageReplyHeaderClazz
	Date        int32
	Message     string
	Media       tg.MessageMediaClazz
	Entities    []tg.MessageEntityClazz
	GroupedId   *int64
}

func buildMessageMediaOutbox(in mediaOutboxInput) *msg.TLOutboxMessage {
	return msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
		RandomId:   in.RandomId,
		Background: in.Background,
		Message: tg.MakeTLMessage(&tg.TLMessage{
			Out:         true,
			Silent:      in.Silent,
			Noforwards:  in.Noforwards,
			InvertMedia: in.InvertMedia,
			FromId:      tg.MakePeerUser(in.FromId),
			PeerId:      tg.MakePeerUser(in.PeerId),
			ReplyTo:     in.ReplyTo,
			Date:        in.Date,
			Message:     in.Message,
			Media:       in.Media,
			Entities:    in.Entities,
			GroupedId:   in.GroupedId,
		}),
	})
}

func resolveMessageMedia(ctx context.Context, mediaClient resolveMediaClient, ownerID int64, input tg.InputMediaClazz) (tg.MessageMediaClazz, error) {
	if input == nil || mediaClient == nil {
		return nil, tg.ErrMediaEmpty
	}

	inputMedia := &tg.InputMedia{Clazz: input}
	now := int32(time.Now().Unix())
	switch input.InputMediaClazzName() {
	case tg.ClazzName_inputMediaUploadedPhoto:
		uploadedPhoto, ok := inputMedia.ToInputMediaUploadedPhoto()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		photo, err := mediaClient.MediaUploadPhotoFile(ctx, &mediapb.TLMediaUploadPhotoFile{
			OwnerId:    ownerID,
			File:       uploadedPhoto.File,
			Stickers:   uploadedPhoto.Stickers,
			TtlSeconds: uploadedPhoto.TtlSeconds,
		})
		if err != nil {
			return nil, err
		}
		if photo == nil || photo.Clazz == nil {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      photo.Clazz,
			TtlSeconds: uploadedPhoto.TtlSeconds,
		}), nil
	case tg.ClazzName_inputMediaPhoto:
		mediaPhoto, ok := inputMedia.ToInputMediaPhoto()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		inputPhoto, ok := (&tg.InputPhoto{Clazz: mediaPhoto.Id}).ToInputPhoto()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		sizeList, err := mediaClient.MediaGetPhotoSizeList(ctx, &mediapb.TLMediaGetPhotoSizeList{SizeId: inputPhoto.Id})
		if err != nil {
			return nil, err
		}
		if sizeList == nil {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo: tg.MakeTLPhoto(&tg.TLPhoto{
				Id:         inputPhoto.Id,
				AccessHash: inputPhoto.AccessHash,
				Date:       now,
				Sizes:      sizeList.Sizes,
				DcId:       sizeList.DcId,
			}),
			TtlSeconds: mediaPhoto.TtlSeconds,
		}), nil
	case tg.ClazzName_inputMediaUploadedDocument:
		documentMedia, err := mediaClient.MediaUploadedDocumentMedia(ctx, &mediapb.TLMediaUploadedDocumentMedia{
			OwnerId: ownerID,
			Media:   input,
		})
		if err != nil {
			return nil, err
		}
		if documentMedia == nil || documentMedia.Clazz == nil {
			return nil, tg.ErrMediaEmpty
		}
		return documentMedia.Clazz, nil
	case tg.ClazzName_inputMediaDocument:
		mediaDocument, ok := inputMedia.ToInputMediaDocument()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		inputDocument, ok := (&tg.InputDocument{Clazz: mediaDocument.Id}).ToInputDocument()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		document, err := mediaClient.MediaGetDocument(ctx, &mediapb.TLMediaGetDocument{Id: inputDocument.Id})
		if err != nil {
			return nil, err
		}
		if document == nil || document.Clazz == nil {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Document:   document.Clazz,
			TtlSeconds: mediaDocument.TtlSeconds,
		}), nil
	case tg.ClazzName_inputMediaGeoPoint:
		geoPoint, ok := inputMedia.ToInputMediaGeoPoint()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaGeo(&tg.TLMessageMediaGeo{Geo: makeSendMediaGeoPoint(geoPoint.GeoPoint)}), nil
	case tg.ClazzName_inputMediaContact:
		contact, ok := inputMedia.ToInputMediaContact()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
			PhoneNumber: contact.PhoneNumber,
			FirstName:   contact.FirstName,
			LastName:    contact.LastName,
			Vcard:       contact.Vcard,
			UserId:      0,
		}), nil
	case tg.ClazzName_inputMediaVenue:
		venue, ok := inputMedia.ToInputMediaVenue()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaVenue(&tg.TLMessageMediaVenue{
			Geo:       makeSendMediaGeoPoint(venue.GeoPoint),
			Title:     venue.Title,
			Address:   venue.Address,
			Provider:  venue.Provider,
			VenueId:   venue.VenueId,
			VenueType: venue.VenueType,
		}), nil
	case tg.ClazzName_inputMediaGeoLive:
		geoLive, ok := inputMedia.ToInputMediaGeoLive()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaGeoLive(&tg.TLMessageMediaGeoLive{
			Geo:    makeSendMediaGeoPoint(geoLive.GeoPoint),
			Period: valueOrZeroInt32(geoLive.Period),
		}), nil
	case tg.ClazzName_inputMediaPoll:
		poll, ok := inputMedia.ToInputMediaPoll()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaPoll(&tg.TLMessageMediaPoll{Poll: poll.Poll}), nil
	case tg.ClazzName_inputMediaDice:
		dice, ok := inputMedia.ToInputMediaDice()
		if !ok {
			return nil, tg.ErrMediaEmpty
		}
		return makeSendMediaDice(dice.Emoticon), nil
	default:
		return nil, tg.ErrMediaEmpty
	}
}

func mapMediaResolveError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, tg.ErrMediaEmpty):
		return tg.ErrMediaEmpty
	case errors.Is(err, mediapb.ErrMediaInvalidArgument),
		errors.Is(err, mediapb.ErrMediaInvalidUploadedFile),
		errors.Is(err, mediapb.ErrMediaChecksumInvalid),
		errors.Is(err, mediapb.ErrPhotoNotFound),
		errors.Is(err, mediapb.ErrDocumentNotFound),
		errors.Is(err, mediapb.ErrMediaBlocked):
		return tg.ErrMediaEmpty
	default:
		return tg.ErrInternalServerError
	}
}

func makeSendMediaGeoPoint(input tg.InputGeoPointClazz) tg.GeoPointClazz {
	if inputPoint, ok := (&tg.InputGeoPoint{Clazz: input}).ToInputGeoPoint(); ok {
		return tg.MakeTLGeoPoint(&tg.TLGeoPoint{
			Long:           inputPoint.Long,
			Lat:            inputPoint.Lat,
			AccuracyRadius: inputPoint.AccuracyRadius,
		})
	}
	return tg.MakeTLGeoPointEmpty(&tg.TLGeoPointEmpty{})
}

func makeSendMediaDice(emoticon string) tg.MessageMediaClazz {
	maxValue := int32(6)
	if emoticon == "🏀" {
		maxValue = 5
	}
	return tg.MakeTLMessageMediaDice(&tg.TLMessageMediaDice{
		Value:    rand.Int31()%maxValue + 1,
		Emoticon: emoticon,
	})
}

func valueOrZeroInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}
