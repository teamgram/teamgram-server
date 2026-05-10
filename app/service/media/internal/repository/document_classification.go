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
		return documentClassification{Kind: documentKindCustomEmoji, RenderKind: messageRenderKindFile}, nil
	}
	if attrs.sticker {
		if !isStickerMime(mimeType) || attrs.audio || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return documentClassification{Kind: documentKindSticker, RenderKind: messageRenderKindFile}, nil
	}
	if attrs.voice {
		if !isAudioMime(mimeType) || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return documentClassification{Kind: documentKindVoice, RenderKind: messageRenderKindFile}, nil
		}
		return documentClassification{Kind: documentKindVoice, RenderKind: messageRenderKindVoice, Voice: true}, nil
	}
	if attrs.audio {
		if !isAudioMime(mimeType) || attrs.video || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return documentClassification{Kind: documentKindAudio, RenderKind: messageRenderKindAudio}, nil
	}
	if mimeType == "image/gif" && attrs.animated {
		if attrs.video {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return documentClassification{Kind: documentKindGifv, RenderKind: messageRenderKindFile, RequiredTransform: requiredDocumentTransformGifv}, nil
		}
		return documentClassification{Kind: documentKindGifv, RenderKind: messageRenderKindVideo, RequiredTransform: requiredDocumentTransformGifv, Video: true}, nil
	}
	if attrs.round {
		if mimeType != "video/mp4" || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return documentClassification{Kind: documentKindRound, RenderKind: messageRenderKindFile, RequiredTransform: requiredDocumentTransformMp4}, nil
		}
		return documentClassification{Kind: documentKindRound, RenderKind: messageRenderKindRound, RequiredTransform: requiredDocumentTransformMp4, Video: true, Round: true}, nil
	}
	if attrs.video || mimeType == "video/mp4" {
		if mimeType != "video/mp4" || attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		if uploaded.ForceFile {
			return documentClassification{Kind: documentKindVideo, RenderKind: messageRenderKindFile, RequiredTransform: requiredDocumentTransformMp4}, nil
		}
		return documentClassification{Kind: documentKindVideo, RenderKind: messageRenderKindVideo, RequiredTransform: requiredDocumentTransformMp4, Video: true}, nil
	}
	if isImageDocumentMime(mimeType) && attrs.imageSize {
		if attrs.animated {
			return documentClassification{}, media.ErrMediaInvalidUploadedFile
		}
		return documentClassification{Kind: documentKindImage, RenderKind: messageRenderKindImage}, nil
	}
	if attrs.animated {
		return documentClassification{}, media.ErrMediaInvalidUploadedFile
	}
	return documentClassification{Kind: documentKindPlain, RenderKind: messageRenderKindFile}, nil
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

func hasDocumentAttributeCustomEmoji(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).customEmoji
}

func hasDocumentAttributeSticker(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).sticker
}

func hasDocumentAttributeAudio(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).audio
}

func hasDocumentAttributeVoice(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).voice
}

func hasDocumentAttributeVideo(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).video
}

func hasDocumentAttributeRoundMessage(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).round
}

func hasDocumentAttributeImageSize(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).imageSize
}

func hasDocumentAttributeAnimated(attrs []tg.DocumentAttributeClazz) bool {
	return documentAttributePresence(attrs).animated
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
