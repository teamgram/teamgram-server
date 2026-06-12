package core

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/phonenumber"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/zeromicro/go-zero/core/logx"
)

func buildMediaOutbox(in *tg.TLMessagesSendMedia, selfUserID int64, peer resolvedMessagePeer, media tg.MessageMediaClazz, replyTo tg.MessageReplyHeaderClazz) (*msg.TLOutboxMessage, int32) {
	date := int32(time.Now().Unix())
	return buildMessageMediaOutbox(mediaOutboxInput{
		RandomId:    in.RandomId,
		Background:  in.Background,
		Silent:      in.Silent,
		Noforwards:  in.Noforwards,
		InvertMedia: in.InvertMedia,
		FromId:      selfUserID,
		PeerType:    peer.PeerType,
		PeerId:      peer.PeerID,
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
	PeerType    int32
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
			PeerId:      messagePeerClazz(in.PeerType, in.PeerId),
			ReplyTo:     in.ReplyTo,
			Date:        in.Date,
			Message:     in.Message,
			Media:       in.Media,
			Entities:    remakeMessageTextEntities(in.Message, in.Entities, in.FromId, false),
			GroupedId:   in.GroupedId,
		}),
	})
}

func resolveMessageMedia(ctx context.Context, mediaClient resolveMediaClient, userClient userLookupClient, ownerID int64, input tg.InputMediaClazz) (tg.MessageMediaClazz, error) {
	if input == nil {
		return nil, tg.ErrMediaEmpty
	}

	inputMedia := &tg.InputMedia{Clazz: input}
	switch input.InputMediaClazzName() {
	case tg.ClazzName_inputMediaUploadedPhoto:
		if mediaClient == nil {
			return nil, tg.ErrMediaEmpty
		}
		uploadedPhoto, ok := inputMedia.ToInputMediaUploadedPhoto()
		if !ok || uploadedPhoto == nil {
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
		if mediaClient == nil {
			return nil, tg.ErrMediaEmpty
		}
		mediaPhoto, ok := inputMedia.ToInputMediaPhoto()
		if !ok || mediaPhoto == nil {
			return nil, tg.ErrMediaEmpty
		}
		inputPhoto, ok := (&tg.InputPhoto{Clazz: mediaPhoto.Id}).ToInputPhoto()
		if !ok || inputPhoto == nil {
			return nil, tg.ErrMediaEmpty
		}
		photo, err := mediaClient.MediaGetPhoto(ctx, &mediapb.TLMediaGetPhoto{PhotoId: inputPhoto.Id})
		if err != nil {
			return nil, err
		}
		if photo == nil || photo.Clazz == nil {
			return nil, tg.ErrMediaEmpty
		}
		return tg.MakeTLMessageMediaPhoto(&tg.TLMessageMediaPhoto{
			Photo:      photo.Clazz,
			TtlSeconds: mediaPhoto.TtlSeconds,
		}), nil
	case tg.ClazzName_inputMediaContact:
		contact, ok := inputMedia.ToInputMediaContact()
		if !ok || contact == nil {
			return nil, tg.ErrMediaEmpty
		}
		messageMedia := tg.MakeTLMessageMediaContact(&tg.TLMessageMediaContact{
			PhoneNumber: contact.PhoneNumber,
			FirstName:   contact.FirstName,
			LastName:    contact.LastName,
			Vcard:       contact.Vcard,
			UserId:      0,
		})
		phoneNumber, err := phonenumber.CheckAndGetPhoneNumber(contact.PhoneNumber)
		if err == nil && userClient != nil {
			userID, err := userClient.UserGetUserIdByPhone(ctx, &userpb.TLUserGetUserIdByPhone{Phone: phoneNumber})
			if err != nil {
				logx.WithContext(ctx).Errorf("messages.resolveMedia - user lookup by contact phone failed: err: %v", err)
			} else if userID != nil {
				messageMedia.UserId = userID.V
			}
		}
		return messageMedia, nil
	case tg.ClazzName_inputMediaUploadedDocument:
		if mediaClient == nil {
			return nil, tg.ErrMediaEmpty
		}
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
		if mediaClient == nil {
			return nil, tg.ErrMediaEmpty
		}
		mediaDocument, ok := inputMedia.ToInputMediaDocument()
		if !ok || mediaDocument == nil {
			return nil, tg.ErrMediaEmpty
		}
		inputDocument, ok := (&tg.InputDocument{Clazz: mediaDocument.Id}).ToInputDocument()
		if !ok || inputDocument == nil {
			return nil, tg.ErrMediaEmpty
		}
		document, err := mediaClient.MediaGetDocument(ctx, &mediapb.TLMediaGetDocument{Id: inputDocument.Id})
		if err != nil {
			return nil, err
		}
		if document == nil || document.Clazz == nil {
			return nil, tg.ErrMediaEmpty
		}
		videoCover, err := resolveInputPhotoForMessageMedia(ctx, mediaClient, mediaDocument.VideoCover)
		if err != nil {
			return nil, err
		}
		video, round, voice := inferMessageMediaDocumentFlags(document.Clazz)
		return tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Video:          video,
			Round:          round,
			Voice:          voice,
			Document:       document.Clazz,
			VideoCover:     videoCover,
			VideoTimestamp: mediaDocument.VideoTimestamp,
			TtlSeconds:     mediaDocument.TtlSeconds,
		}), nil
	default:
		return nil, tg.ErrMediaEmpty
	}
}

func inferMessageMediaDocumentFlags(document tg.DocumentClazz) (video, round, voice bool) {
	doc, ok := document.(*tg.TLDocument)
	if !ok || doc == nil {
		return false, false, false
	}
	webmStickerOrCustomEmoji := false
	if strings.EqualFold(doc.MimeType, "video/webm") {
		for _, attr := range doc.Attributes {
			switch attr.(type) {
			case *tg.TLDocumentAttributeSticker, *tg.TLDocumentAttributeCustomEmoji:
				webmStickerOrCustomEmoji = true
			}
		}
	}
	for _, attr := range doc.Attributes {
		switch a := attr.(type) {
		case *tg.TLDocumentAttributeVideo:
			if !webmStickerOrCustomEmoji {
				video = true
			}
			round = round || a.RoundMessage
		case *tg.TLDocumentAttributeAudio:
			voice = voice || a.Voice
		}
	}
	return video, round, voice
}

func resolveInputPhotoForMessageMedia(ctx context.Context, mediaClient resolveMediaClient, photo tg.InputPhotoClazz) (tg.PhotoClazz, error) {
	switch p := photo.(type) {
	case nil, *tg.TLInputPhotoEmpty:
		return nil, nil
	case *tg.TLInputPhoto:
		full, err := mediaClient.MediaGetPhoto(ctx, &mediapb.TLMediaGetPhoto{PhotoId: p.Id})
		if err != nil {
			return nil, err
		}
		fullPhoto, ok := full.ToPhoto()
		if !ok || fullPhoto.AccessHash != p.AccessHash {
			return nil, tg.ErrMediaEmpty
		}
		return fullPhoto, nil
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
		isRemoteMediaError(err, mediapb.ErrMediaInvalidArgument),
		errors.Is(err, mediapb.ErrMediaInvalidUploadedFile),
		isRemoteMediaError(err, mediapb.ErrMediaInvalidUploadedFile),
		errors.Is(err, mediapb.ErrMediaChecksumInvalid),
		isRemoteMediaError(err, mediapb.ErrMediaChecksumInvalid),
		errors.Is(err, mediapb.ErrPhotoNotFound),
		isRemoteMediaError(err, mediapb.ErrPhotoNotFound),
		errors.Is(err, mediapb.ErrDocumentNotFound),
		isRemoteMediaError(err, mediapb.ErrDocumentNotFound),
		errors.Is(err, mediapb.ErrMediaBlocked):
		return tg.ErrMediaEmpty
	default:
		return tg.ErrInternalServerError
	}
}

func isRemoteMediaError(err error, target error) bool {
	if err == nil || target == nil {
		return false
	}
	return strings.Contains(err.Error(), "biz error: "+target.Error())
}
