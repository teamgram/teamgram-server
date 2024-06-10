// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package plugin

import (
	"context"
	"sort"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
	"github.com/teamgram/teamgram-server/pkg/mention"

	"mvdan.cc/xurls/v2"
)

func RemakeMessage(ctx context.Context, plugin MsgPlugin, message *mtproto.Message, fromId int64, noWebpage bool, hasBot func() bool) *mtproto.Message {
	var (
		entities mtproto.MessageEntitySlice
		idxList  []int
	)

	/*
		fmt.Println(*update.ChannelPost.Entities)
		// For each entity
		for _, e := range *update.ChannelPost.Entities {
			// Get the whole update Text
			str := update.ChannelPost.Text
			// Encode it into utf16
			utfEncodedString := utf16.Encode([]rune(str))
			// Decode just the piece of string I need
			runeString := utf16.Decode(utfEncodedString[e.Offset : e.Offset+e.Length])
			// Transform []rune into string
			str = string(runeString)
			fmt.Println(str)
		}
	*/
	getIdxList := func() []int {
		if len(idxList) == 0 {
			idxList = mention.EncodeStringToUTF16Index(message.Message)
		}
		return idxList
	}

	var firstUrl string
	rIndexes := xurls.Relaxed().FindAllStringIndex(message.Message, -1)
	if len(rIndexes) > 0 {
		if len(idxList) == 0 {
			getIdxList()
		}
		for idx, v := range rIndexes {
			if idx == 0 {
				firstUrl = message.Message[v[0]:v[1]]
			}
			entityUrl := mtproto.MakeTLMessageEntityUrl(&mtproto.MessageEntity{
				Offset: int32(idxList[v[0]]),
				Length: int32(idxList[v[1]] - idxList[v[0]]),
			}).To_MessageEntity()
			entities = append(entities, entityUrl)
		}
	}

	if !noWebpage && firstUrl != "" && plugin != nil {
		webpage, _ := plugin.GetWebpagePreview(ctx, firstUrl)
		if webpage != nil {
			message.Media = mtproto.MakeTLMessageMediaWebPage(&mtproto.MessageMedia{
				Webpage: webpage,
			}).To_MessageMedia()
		}
	}

	for _, entity := range message.Entities {
		switch entity.PredicateName {
		case mtproto.Predicate_inputMessageEntityMentionName:
			if entity.GetUserId_INPUTUSER().GetPredicateName() == mtproto.Predicate_inputUserSelf {
				entity.UserId_INPUTUSER.UserId = fromId
			}

			if entity.UserId_INPUTUSER.UserId != 0 {
				// TODO(@benqi): check user_id
				entityMentionName := mtproto.MakeTLMessageEntityMentionName(&mtproto.MessageEntity{
					Offset:       entity.Offset,
					Length:       entity.Length,
					UserId_INT64: entity.UserId_INPUTUSER.UserId,
				})

				entities = append(entities, entityMentionName.To_MessageEntity())
			}
			// }
		default:
			entities = append(entities, entity)
		}
	}

	tags := mention.GetTags('@', message.Message, '(', ')')
	if len(tags) > 0 {
		var nameList = make([]string, 0, len(tags))
		for _, tag := range tags {
			nameList = append(nameList, tag.Tag)
		}
		//names, _ := c.svcCtx.Dao.UsernameClient.UsernameGetListByUsernameList(c.ctx, &username.TLUsernameGetListByUsernameList{
		//	Names:                nameList,
		//})
		// c.Logger.Debugf("nameList: %v", names)

		for _, tag := range tags {
			if len(idxList) == 0 {
				getIdxList()
			}
			mention2 := mtproto.MakeTLMessageEntityMention(&mtproto.MessageEntity{
				Offset: int32(idxList[tag.Index]),
				Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
			}).To_MessageEntity()

			// stole field UserId_5
			if plugin != nil {
				if v, _ := plugin.UsernameResolveUsername(ctx, &username.TLUsernameResolveUsername{
					Username: tag.Tag,
				}); v != nil {
					if v.GetPredicateName() == mtproto.Predicate_peerUser {
						mention2.UserId_INT64 = v.UserId
					}
				}
			}
			//for _, v := range names.Datas {
			//	//
			//	if uname, ok := names[tag.Tag]; ok {
			//		if uname.PeerType == model.PEER_USER {
			//			mention2.UserId_INT32 = uname.PeerId
			//		}
			//	}
			//}
			entities = append(entities, mention2)
		}
	}

	tags = mention.GetTags('#', message.Message)
	for _, tag := range tags {
		if len(idxList) == 0 {
			getIdxList()
		}
		hashtag := mtproto.MakeTLMessageEntityHashtag(&mtproto.MessageEntity{
			Offset: int32(idxList[tag.Index]),
			Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
			Url:    "#" + tag.Tag, // NOTE: hack, steal url field
		}).To_MessageEntity()
		entities = append(entities, hashtag)
	}

	if hasBot != nil && hasBot() {
		tags = mention.GetTags('/', message.Message)
		for _, tag := range tags {
			if len(idxList) == 0 {
				getIdxList()
			}
			hashtag := mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: int32(idxList[tag.Index]),
				Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
			}).To_MessageEntity()
			entities = append(entities, hashtag)
		}
	}

	sort.Sort(entities)
	message.Entities = entities
	return message
}
