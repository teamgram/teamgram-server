// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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

package core

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const NormalizedOutboxSchemaVersionV1 = 1

type normalizedOutboxMessage struct {
	SchemaVersion              int
	RandomID                   int64
	Background                 bool
	FromUserID                 int64
	PeerType                   int32
	PeerID                     int64
	Out                        bool
	Date                       int64
	MessageText                string
	Entities                   []payload.MessageEntityV1
	ReplyToUserMessageID       int64
	ReplyToCanonicalMessageID  int64
	MediaRef                   *payload.MediaRefV1
	Attrs                      payload.MessageAttrsV1
	ForwardRef                 *payload.ForwardRefV1
	ForwardSourceCanonicalID   int64
	ForwardSourceUserMessageID int64
}

type batchSideEffects struct {
	ClearDraft           bool
	SourcePermAuthKeyID  int64
	ClearDraftBeforeDate int32
}

type normalizeOutboxInput struct {
	Ctx          context.Context
	SenderUserID int64
	PeerType     int32
	PeerID       int64
	Outbox       *msg.TLOutboxMessage
	Repo         repository.MessageRepository
}

func normalizeOutboxMessage(in normalizeOutboxInput) (normalizedOutboxMessage, error) {
	if in.Outbox == nil || in.Outbox.Message == nil {
		return normalizedOutboxMessage{}, fmt.Errorf("%w: missing outbox message", msg.ErrSendStateConflict)
	}
	if in.Outbox.ScheduleDate != nil {
		return normalizedOutboxMessage{}, fmt.Errorf("%w: scheduled messages are not supported", msg.ErrSendStateConflict)
	}
	if in.Outbox.RandomId == 0 {
		return normalizedOutboxMessage{}, fmt.Errorf("%w: random_id is required", msg.ErrSendStateConflict)
	}
	message, ok := in.Outbox.Message.(*tg.TLMessage)
	if !ok {
		return normalizedOutboxMessage{}, fmt.Errorf("%w: unsupported outbox message shape", msg.ErrSendStateConflict)
	}
	if message.FromScheduled || message.ScheduleRepeatPeriod != nil {
		return normalizedOutboxMessage{}, fmt.Errorf("%w: scheduled messages are not supported", msg.ErrSendStateConflict)
	}
	mediaRef, err := normalizeMediaRef(message.Media)
	if err != nil {
		return normalizedOutboxMessage{}, err
	}
	replyTo, err := normalizeReplyTo(in, message)
	if err != nil {
		return normalizedOutboxMessage{}, err
	}
	forwardRef, forwardSource, err := normalizeForwardRef(in, message)
	if err != nil {
		return normalizedOutboxMessage{}, err
	}
	entities, err := normalizeEntities(message.Entities)
	if err != nil {
		return normalizedOutboxMessage{}, err
	}
	attrs := payload.MessageAttrsV1{
		SchemaVersion: payload.MessageAttrsSchemaVersionV1,
		Noforwards:    message.Noforwards,
		Silent:        message.Silent,
		InvertMedia:   message.InvertMedia,
	}
	if message.GroupedId != nil {
		attrs.GroupedID = *message.GroupedId
	}
	return normalizedOutboxMessage{
		SchemaVersion:              NormalizedOutboxSchemaVersionV1,
		RandomID:                   in.Outbox.RandomId,
		Background:                 in.Outbox.Background,
		FromUserID:                 in.SenderUserID,
		PeerType:                   in.PeerType,
		PeerID:                     in.PeerID,
		Out:                        true,
		MessageText:                message.Message,
		Entities:                   entities,
		ReplyToUserMessageID:       replyTo.UserMessageID,
		ReplyToCanonicalMessageID:  replyTo.CanonicalMessageID,
		MediaRef:                   mediaRef,
		Attrs:                      attrs,
		ForwardRef:                 forwardRef,
		ForwardSourceCanonicalID:   forwardSource.CanonicalMessageID,
		ForwardSourceUserMessageID: forwardSource.UserMessageID,
	}, nil
}

func normalizeReplyTo(in normalizeOutboxInput, message *tg.TLMessage) (resolvedReplyToMessage, error) {
	if message == nil || message.ReplyTo == nil {
		return resolvedReplyToMessage{}, nil
	}
	replyHeader, ok := message.ReplyTo.(*tg.TLMessageReplyHeader)
	if !ok || replyHeader.ReplyToMsgId == nil || *replyHeader.ReplyToMsgId <= 0 {
		return resolvedReplyToMessage{}, msg.ErrReplyToInvalid
	}
	if in.Repo == nil {
		return resolvedReplyToMessage{}, fmt.Errorf("%w: reply resolver is required", msg.ErrSendStateConflict)
	}
	replyTo, err := in.Repo.ResolveMessageID(in.Ctx, in.SenderUserID, in.PeerType, in.PeerID, int64(*replyHeader.ReplyToMsgId))
	if err != nil {
		return resolvedReplyToMessage{}, err
	}
	if replyTo == nil || replyTo.CanonicalMessageID == 0 || replyTo.UserMessageID <= 0 {
		return resolvedReplyToMessage{}, msg.ErrReplyToInvalid
	}
	return resolvedReplyToMessage{
		CanonicalMessageID: replyTo.CanonicalMessageID,
		UserMessageID:      replyTo.UserMessageID,
	}, nil
}

func normalizeMediaRef(media tg.MessageMediaClazz) (*payload.MediaRefV1, error) {
	if media == nil {
		return nil, nil
	}
	switch m := media.(type) {
	case *tg.TLMessageMediaEmpty:
		return nil, nil
	case *tg.TLMessageMediaPhoto:
		ref := &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "photo"}
		if p, ok := m.Photo.(*tg.TLPhoto); ok && p != nil {
			ref.ID = p.Id
			ref.AccessHash = p.AccessHash
			ref.FileReference = append([]byte(nil), p.FileReference...)
			ref.Date = p.Date
			ref.DcID = p.DcId
			ref.PhotoSizes = normalizePhotoSizes(p.Sizes)
		} else if m.Photo != nil {
			ref.ID, ref.AccessHash, ref.FileReference = photoIdentity(m.Photo)
		}
		if m.TtlSeconds != nil {
			ref.TTLSeconds = *m.TtlSeconds
		}
		return ref, nil
	case *tg.TLMessageMediaDocument:
		ref := &payload.MediaRefV1{SchemaVersion: payload.MediaRefSchemaVersionV1, Kind: "document"}
		if d, ok := m.Document.(*tg.TLDocument); ok && d != nil {
			ref.ID = d.Id
			ref.AccessHash = d.AccessHash
			ref.FileReference = append([]byte(nil), d.FileReference...)
			ref.Date = d.Date
			ref.MimeType = d.MimeType
			ref.Size = d.Size2
			ref.DcID = d.DcId
			ref.DocumentThumbs = normalizePhotoSizes(d.Thumbs)
			ref.DocumentAttributes = normalizeDocumentAttributes(d.Attributes)
		} else if m.Document != nil {
			ref.ID, ref.AccessHash, ref.FileReference, ref.MimeType = documentIdentity(m.Document)
		}
		if m.TtlSeconds != nil {
			ref.TTLSeconds = *m.TtlSeconds
		}
		return ref, nil
	default:
		return nil, fmt.Errorf("%w: unsupported message media %T", msg.ErrSendStateConflict, media)
	}
}

func photoIdentity(photo tg.PhotoClazz) (int64, int64, []byte) {
	switch p := photo.(type) {
	case *tg.TLPhoto:
		return p.Id, p.AccessHash, append([]byte(nil), p.FileReference...)
	case *tg.TLPhotoEmpty:
		return p.Id, 0, nil
	default:
		return 0, 0, nil
	}
}

func documentIdentity(document tg.DocumentClazz) (int64, int64, []byte, string) {
	switch d := document.(type) {
	case *tg.TLDocument:
		return d.Id, d.AccessHash, append([]byte(nil), d.FileReference...), d.MimeType
	case *tg.TLDocumentEmpty:
		return d.Id, 0, nil, ""
	default:
		return 0, 0, nil, ""
	}
}

func normalizePhotoSizes(sizes []tg.PhotoSizeClazz) []payload.PhotoSizeRefV1 {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]payload.PhotoSizeRefV1, 0, len(sizes))
	for _, size := range sizes {
		switch s := size.(type) {
		case *tg.TLPhotoSizeEmpty:
			out = append(out, payload.PhotoSizeRefV1{Kind: "empty", Type: s.Type})
		case *tg.TLPhotoSize:
			out = append(out, payload.PhotoSizeRefV1{Kind: "size", Type: s.Type, W: s.W, H: s.H, Size: s.Size2})
		case *tg.TLPhotoCachedSize:
			out = append(out, payload.PhotoSizeRefV1{Kind: "cached", Type: s.Type, W: s.W, H: s.H, Bytes: append([]byte(nil), s.Bytes...)})
		case *tg.TLPhotoStrippedSize:
			out = append(out, payload.PhotoSizeRefV1{Kind: "stripped", Type: s.Type, Bytes: append([]byte(nil), s.Bytes...)})
		case *tg.TLPhotoSizeProgressive:
			out = append(out, payload.PhotoSizeRefV1{Kind: "progressive", Type: s.Type, W: s.W, H: s.H, Sizes: append([]int32(nil), s.Sizes...)})
		}
	}
	return out
}

func normalizeDocumentAttributes(attrs []tg.DocumentAttributeClazz) []payload.DocumentAttributeRefV1 {
	if len(attrs) == 0 {
		return nil
	}
	out := make([]payload.DocumentAttributeRefV1, 0, len(attrs))
	for _, attr := range attrs {
		switch a := attr.(type) {
		case *tg.TLDocumentAttributeFilename:
			out = append(out, payload.DocumentAttributeRefV1{Kind: "filename", FileName: a.FileName})
		case *tg.TLDocumentAttributeImageSize:
			out = append(out, payload.DocumentAttributeRefV1{Kind: "image_size", W: a.W, H: a.H})
		case *tg.TLDocumentAttributeAnimated:
			out = append(out, payload.DocumentAttributeRefV1{Kind: "animated"})
		case *tg.TLDocumentAttributeVideo:
			out = append(out, payload.DocumentAttributeRefV1{
				Kind:              "video",
				DurationFloat:     a.Duration,
				W:                 a.W,
				H:                 a.H,
				RoundMessage:      a.RoundMessage,
				SupportsStreaming: a.SupportsStreaming,
				NoSound:           a.Nosound,
				PreloadPrefixSize: a.PreloadPrefixSize,
				VideoStartTs:      a.VideoStartTs,
				VideoCodec:        a.VideoCodec,
			})
		case *tg.TLDocumentAttributeAudio:
			out = append(out, payload.DocumentAttributeRefV1{
				Kind:      "audio",
				Duration:  a.Duration,
				Title:     a.Title,
				Performer: a.Performer,
				Waveform:  append([]byte(nil), a.Waveform...),
				Voice:     a.Voice,
			})
		}
	}
	return out
}

func normalizeForwardRef(in normalizeOutboxInput, message *tg.TLMessage) (*payload.ForwardRefV1, repository.ForwardSourceIdentity, error) {
	if message == nil || message.FwdFrom == nil {
		return nil, repository.ForwardSourceIdentity{}, nil
	}
	fwd := message.FwdFrom
	ref := &payload.ForwardRefV1{
		SchemaVersion: payload.ForwardRefSchemaVersionV1,
		Date:          int64(fwd.Date),
	}
	if fwd.FromName != nil {
		ref.FromName = *fwd.FromName
	}
	ref.FromUserID, ref.SourcePeerType, ref.SourcePeerID = forwardSourcePeer(fwd.FromId)
	if fwd.ChannelPost != nil {
		ref.SourceMessageID = int64(*fwd.ChannelPost)
	}
	savedPeerType, savedPeerID := peerIdentity(fwd.SavedFromPeer)
	ref.SavedFromPeerType = savedPeerType
	ref.SavedFromPeerID = savedPeerID
	if fwd.SavedFromMsgId != nil {
		ref.SavedFromMessageID = int64(*fwd.SavedFromMsgId)
		if ref.SourcePeerType == 0 && savedPeerType != 0 {
			ref.SourcePeerType = savedPeerType
			ref.SourcePeerID = savedPeerID
			ref.SourceMessageID = int64(*fwd.SavedFromMsgId)
		}
	}
	if savedPeerType == 0 || savedPeerID <= 0 || ref.SavedFromMessageID <= 0 {
		return nil, repository.ForwardSourceIdentity{}, msg.ErrMsgIdInvalid
	}
	if in.Repo == nil {
		return nil, repository.ForwardSourceIdentity{}, fmt.Errorf("%w: forward resolver is required", msg.ErrSendStateConflict)
	}
	ref.SourcePeerType = savedPeerType
	ref.SourcePeerID = savedPeerID
	ref.SourceMessageID = ref.SavedFromMessageID
	source, err := in.Repo.ResolveForwardSourceIdentity(in.Ctx, repository.ForwardSourceLookup{
		UserID:              in.SenderUserID,
		SourcePeerType:      savedPeerType,
		SourcePeerID:        savedPeerID,
		SourceUserMessageID: ref.SavedFromMessageID,
	})
	if err != nil {
		return nil, repository.ForwardSourceIdentity{}, err
	}
	if source == nil || source.CanonicalMessageID <= 0 {
		return nil, repository.ForwardSourceIdentity{}, msg.ErrMsgIdInvalid
	}
	return ref, *source, nil
}

func forwardSourcePeer(peer tg.PeerClazz) (fromUserID int64, sourcePeerType int32, sourcePeerID int64) {
	peerType, peerID := peerIdentity(peer)
	if peerType == payload.PeerTypeUser {
		return peerID, peerType, peerID
	}
	return 0, peerType, peerID
}

func peerIdentity(peer tg.PeerClazz) (int32, int64) {
	switch p := peer.(type) {
	case *tg.TLPeerUser:
		return payload.PeerTypeUser, p.UserId
	case *tg.TLPeerChat:
		return payload.PeerTypeChat, p.ChatId
	case *tg.TLPeerChannel:
		return payload.PeerTypeChannel, p.ChannelId
	default:
		return 0, 0
	}
}

func normalizeEntities(entities []tg.MessageEntityClazz) ([]payload.MessageEntityV1, error) {
	if len(entities) == 0 {
		return nil, nil
	}
	out := make([]payload.MessageEntityV1, 0, len(entities))
	for _, entity := range entities {
		if item, ok := normalizeEntity(entity); ok {
			out = append(out, item)
			continue
		}
		return nil, fmt.Errorf("%w: unsupported message entity %T", msg.ErrSendStateConflict, entity)
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

func normalizeEntity(entity tg.MessageEntityClazz) (payload.MessageEntityV1, bool) {
	switch e := entity.(type) {
	case *tg.TLMessageEntityMention:
		return entityV1(e.Offset, e.Length, "mention", ""), true
	case *tg.TLMessageEntityHashtag:
		return entityV1(e.Offset, e.Length, "hashtag", ""), true
	case *tg.TLMessageEntityBotCommand:
		return entityV1(e.Offset, e.Length, "bot_command", ""), true
	case *tg.TLMessageEntityUrl:
		return entityV1(e.Offset, e.Length, "url", ""), true
	case *tg.TLMessageEntityEmail:
		return entityV1(e.Offset, e.Length, "email", ""), true
	case *tg.TLMessageEntityBold:
		return entityV1(e.Offset, e.Length, "bold", ""), true
	case *tg.TLMessageEntityItalic:
		return entityV1(e.Offset, e.Length, "italic", ""), true
	case *tg.TLMessageEntityCode:
		return entityV1(e.Offset, e.Length, "code", ""), true
	case *tg.TLMessageEntityPre:
		return entityV1(e.Offset, e.Length, "pre", e.Language), true
	case *tg.TLMessageEntityTextUrl:
		return entityV1(e.Offset, e.Length, "text_url", e.Url), true
	case *tg.TLMessageEntityPhone:
		return entityV1(e.Offset, e.Length, "phone", ""), true
	case *tg.TLMessageEntityCashtag:
		return entityV1(e.Offset, e.Length, "cashtag", ""), true
	case *tg.TLMessageEntityUnderline:
		return entityV1(e.Offset, e.Length, "underline", ""), true
	case *tg.TLMessageEntityStrike:
		return entityV1(e.Offset, e.Length, "strike", ""), true
	case *tg.TLMessageEntityBankCard:
		return entityV1(e.Offset, e.Length, "bank_card", ""), true
	case *tg.TLMessageEntitySpoiler:
		return entityV1(e.Offset, e.Length, "spoiler", ""), true
	case *tg.TLMessageEntityBlockquote:
		return entityV1(e.Offset, e.Length, "blockquote", ""), true
	default:
		return payload.MessageEntityV1{}, false
	}
}

func entityV1(offset int32, length int32, kind string, url string) payload.MessageEntityV1 {
	return payload.MessageEntityV1{Offset: offset, Length: length, Kind: kind, URL: url}
}
