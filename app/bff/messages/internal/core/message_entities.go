package core

import (
	"net/mail"
	"sort"
	"strings"

	"github.com/teamgram/teamgram-server/v2/pkg/mention"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"mvdan.cc/xurls/v2"
)

var messageURLMatcher = xurls.Relaxed()

func remakeMessageTextEntities(text string, existing []tg.MessageEntityClazz, fromID int64, hasBot bool) []tg.MessageEntityClazz {
	if text == "" && len(existing) == 0 {
		return nil
	}
	var (
		entities []tg.MessageEntityClazz
		idxList  []int
	)
	getIdxList := func() []int {
		if idxList == nil {
			idxList = mention.EncodeStringToUTF16Index(text)
		}
		return idxList
	}

	for _, span := range messageURLMatcher.FindAllStringIndex(text, -1) {
		offset, length := utf16EntityRange(getIdxList(), span[0], span[1])
		if isEmailEntity(text[span[0]:span[1]]) {
			entities = append(entities, tg.MakeTLMessageEntityEmail(&tg.TLMessageEntityEmail{
				Offset: offset,
				Length: length,
			}))
		} else {
			entities = append(entities, tg.MakeTLMessageEntityUrl(&tg.TLMessageEntityUrl{
				Offset: offset,
				Length: length,
			}))
		}
	}

	for _, entity := range existing {
		converted, handled := remakeInputMentionNameEntity(entity, fromID)
		if handled {
			if converted != nil {
				entities = append(entities, converted)
			}
			continue
		}
		entities = append(entities, entity)
	}

	for _, tag := range mention.GetTags('@', text, '(', ')') {
		offset, length := tagEntityRange(getIdxList(), tag)
		entities = append(entities, tg.MakeTLMessageEntityMention(&tg.TLMessageEntityMention{
			Offset: offset,
			Length: length,
		}))
	}

	for _, tag := range mention.GetTags('#', text) {
		offset, length := tagEntityRange(getIdxList(), tag)
		entities = append(entities, tg.MakeTLMessageEntityHashtag(&tg.TLMessageEntityHashtag{
			Offset: offset,
			Length: length,
		}))
	}

	if hasBot {
		for _, tag := range mention.GetTags('/', text) {
			offset, length := tagEntityRange(getIdxList(), tag)
			entities = append(entities, tg.MakeTLMessageEntityBotCommand(&tg.TLMessageEntityBotCommand{
				Offset: offset,
				Length: length,
			}))
		}
	}

	sort.SliceStable(entities, func(i, j int) bool {
		iOffset, iLength := entityRange(entities[i])
		jOffset, jLength := entityRange(entities[j])
		if iOffset != jOffset {
			return iOffset < jOffset
		}
		return iLength < jLength
	})
	return entities
}

func tagEntityRange(idxList []int, tag mention.Tag) (int32, int32) {
	return utf16EntityRange(idxList, tag.Index, tag.Index+len(tag.Tag)+1)
}

func utf16EntityRange(idxList []int, start int, end int) (int32, int32) {
	return int32(idxList[start]), int32(idxList[end] - idxList[start])
}

func isEmailEntity(value string) bool {
	if !strings.Contains(value, "@") || strings.Contains(value, "://") || strings.HasPrefix(strings.ToLower(value), "www.") {
		return false
	}
	_, err := mail.ParseAddress(value)
	return err == nil
}

func remakeInputMentionNameEntity(entity tg.MessageEntityClazz, fromID int64) (tg.MessageEntityClazz, bool) {
	input, ok := entity.(*tg.TLInputMessageEntityMentionName)
	if !ok {
		return nil, false
	}
	userID := inputMessageEntityMentionNameUserID(input.UserId, fromID)
	if userID == 0 {
		return nil, true
	}
	return tg.MakeTLMessageEntityMentionName(&tg.TLMessageEntityMentionName{
		Offset: input.Offset,
		Length: input.Length,
		UserId: userID,
	}), true
}

func inputMessageEntityMentionNameUserID(input tg.InputUserClazz, fromID int64) int64 {
	switch user := input.(type) {
	case *tg.TLInputUserSelf:
		return fromID
	case *tg.TLInputUser:
		return user.UserId
	case *tg.TLInputUserFromMessage:
		return user.UserId
	default:
		return 0
	}
}

func entityRange(entity tg.MessageEntityClazz) (int32, int32) {
	switch e := entity.(type) {
	case *tg.TLMessageEntityMention:
		return e.Offset, e.Length
	case *tg.TLMessageEntityHashtag:
		return e.Offset, e.Length
	case *tg.TLMessageEntityBotCommand:
		return e.Offset, e.Length
	case *tg.TLMessageEntityUrl:
		return e.Offset, e.Length
	case *tg.TLMessageEntityEmail:
		return e.Offset, e.Length
	case *tg.TLMessageEntityBold:
		return e.Offset, e.Length
	case *tg.TLMessageEntityItalic:
		return e.Offset, e.Length
	case *tg.TLMessageEntityCode:
		return e.Offset, e.Length
	case *tg.TLMessageEntityPre:
		return e.Offset, e.Length
	case *tg.TLMessageEntityTextUrl:
		return e.Offset, e.Length
	case *tg.TLMessageEntityMentionName:
		return e.Offset, e.Length
	case *tg.TLInputMessageEntityMentionName:
		return e.Offset, e.Length
	case *tg.TLMessageEntityPhone:
		return e.Offset, e.Length
	case *tg.TLMessageEntityCashtag:
		return e.Offset, e.Length
	case *tg.TLMessageEntityUnderline:
		return e.Offset, e.Length
	case *tg.TLMessageEntityStrike:
		return e.Offset, e.Length
	case *tg.TLMessageEntityBankCard:
		return e.Offset, e.Length
	case *tg.TLMessageEntitySpoiler:
		return e.Offset, e.Length
	case *tg.TLMessageEntityCustomEmoji:
		return e.Offset, e.Length
	case *tg.TLMessageEntityBlockquote:
		return e.Offset, e.Length
	default:
		return 0, 0
	}
}
