package repository

import (
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type documentKind int

const (
	documentKindPlain documentKind = iota
	documentKindImage
	documentKindVideo
	documentKindGifv
	documentKindAudio
	documentKindVoice
	documentKindRound
	documentKindSticker
	documentKindCustomEmoji
)

type messageRenderKind int

const (
	messageRenderKindFile messageRenderKind = iota
	messageRenderKindImage
	messageRenderKindVideo
	messageRenderKindAudio
	messageRenderKindVoice
	messageRenderKindRound
)

type requiredDocumentTransform int

const (
	requiredDocumentTransformNone requiredDocumentTransform = iota
	requiredDocumentTransformGifv
	requiredDocumentTransformMp4
)

type documentClassification struct {
	Kind              documentKind
	RenderKind        messageRenderKind
	RequiredTransform requiredDocumentTransform
	Video             bool
	Round             bool
	Voice             bool
}

func classifyUploadedDocument(uploaded *tg.TLInputMediaUploadedDocument) (documentClassification, error) {
	if uploaded == nil {
		return documentClassification{}, media.ErrMediaInvalidArgument
	}
	mimeType := strings.ToLower(strings.TrimSpace(uploaded.MimeType))
	attrs := documentAttributePresence(uploaded.Attributes)

	if attrs.customEmoji {
		if !isCustomEmojiMime(mimeType) || attrs.sticker || attrs.audio || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return newDocumentClassification(documentKindCustomEmoji, messageRenderKindFile, requiredDocumentTransformNone), nil
	}
	if attrs.sticker {
		if !isStickerMime(mimeType) || attrs.audio || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return newDocumentClassification(documentKindSticker, messageRenderKindFile, requiredDocumentTransformNone), nil
	}
	if attrs.voice {
		if !isAudioMime(mimeType) || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return newDocumentClassification(documentKindVoice, messageRenderKindFile, requiredDocumentTransformNone), nil
		}
		return newDocumentClassification(documentKindVoice, messageRenderKindVoice, requiredDocumentTransformNone), nil
	}
	if attrs.audio {
		if !isAudioMime(mimeType) || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return newDocumentClassification(documentKindAudio, messageRenderKindAudio, requiredDocumentTransformNone), nil
	}
	if mimeType == "image/gif" && attrs.animated {
		if attrs.video {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return newDocumentClassification(documentKindGifv, messageRenderKindFile, requiredDocumentTransformGifv), nil
		}
		return newDocumentClassification(documentKindGifv, messageRenderKindVideo, requiredDocumentTransformGifv), nil
	}
	if attrs.round {
		if mimeType != "video/mp4" || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return newDocumentClassification(documentKindRound, messageRenderKindFile, requiredDocumentTransformMp4), nil
		}
		return newDocumentClassification(documentKindRound, messageRenderKindRound, requiredDocumentTransformMp4), nil
	}
	if attrs.video || mimeType == "video/mp4" {
		if mimeType != "video/mp4" || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return newDocumentClassification(documentKindVideo, messageRenderKindFile, requiredDocumentTransformMp4), nil
		}
		return newDocumentClassification(documentKindVideo, messageRenderKindVideo, requiredDocumentTransformMp4), nil
	}
	if isImageDocumentMime(mimeType) && attrs.imageSize {
		if attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return newDocumentClassification(documentKindImage, messageRenderKindImage, requiredDocumentTransformNone), nil
	}
	if attrs.animated {
		return documentClassification{}, media.ErrMediaInvalidUploadedFile
	}
	return newDocumentClassification(documentKindPlain, messageRenderKindFile, requiredDocumentTransformNone), nil
}

func newDocumentClassification(kind documentKind, renderKind messageRenderKind, transform requiredDocumentTransform) documentClassification {
	classification := documentClassification{
		Kind:              kind,
		RenderKind:        renderKind,
		RequiredTransform: transform,
	}
	classification.Video, classification.Round, classification.Voice = documentMessageFlagsFromRenderKind(classification)
	return classification
}

func documentMessageFlagsFromRenderKind(classification documentClassification) (video, round, voice bool) {
	switch classification.RenderKind {
	case messageRenderKindVideo:
		return true, false, false
	case messageRenderKindRound:
		return true, true, false
	case messageRenderKindVoice:
		return false, false, true
	default:
		return false, false, false
	}
}

type documentAttributes struct {
	customEmoji bool
	sticker     bool
	audio       bool
	voice       bool
	video       bool
	round       bool
	imageSize   bool
	animated    bool
}

func documentAttributePresence(attrs []tg.DocumentAttributeClazz) documentAttributes {
	var out documentAttributes
	for _, attr := range attrs {
		switch typed := attr.(type) {
		case *tg.TLDocumentAttributeCustomEmoji:
			out.customEmoji = true
		case *tg.TLDocumentAttributeSticker:
			out.sticker = true
		case *tg.TLDocumentAttributeAudio:
			out.audio = true
			if typed.Voice {
				out.voice = true
			}
		case *tg.TLDocumentAttributeVideo:
			out.video = true
			if typed.RoundMessage {
				out.round = true
			}
		case *tg.TLDocumentAttributeImageSize:
			out.imageSize = true
		case *tg.TLDocumentAttributeAnimated:
			out.animated = true
		}
	}
	return out
}

func isCustomEmojiMime(mimeType string) bool {
	return mimeType == "image/webp" || mimeType == "application/x-tgsticker" || mimeType == "video/webm"
}

func isStickerMime(mimeType string) bool {
	return mimeType == "image/webp" || mimeType == "application/x-tgsticker" || mimeType == "video/webm"
}

func isAudioMime(mimeType string) bool {
	switch mimeType {
	case "audio/ogg", "audio/mpeg", "audio/mp4", "audio/x-m4a", "audio/flac":
		return true
	default:
		return false
	}
}

func isImageDocumentMime(mimeType string) bool {
	switch mimeType {
	case "image/jpeg", "image/png", "image/webp", "image/gif":
		return true
	default:
		return false
	}
}
