// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"math/rand"
	"mvdan.cc/xurls/v2"
	"sort"
	"time"

	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/app/service/biz/username/username"
	mediapb "github.com/teamgram/teamgram-server/app/service/media/media"
	"github.com/teamgram/teamgram-server/pkg/mention"
	"github.com/teamgram/teamgram-server/pkg/phonenumber"
)

// draft
func (c *MessagesCore) doClearDraft(userId int64, authKeyId int64, peer *mtproto.PeerUtil) {
	var (
		hasClearDraft bool
	)

	c.svcCtx.Dao.DialogClient.DialogClearDraftMessage(c.ctx, &dialog.TLDialogClearDraftMessage{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
	})

	// ClearDraft
	if hasClearDraft {
		updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			Draft:     mtproto.MakeTLDraftMessageEmpty(nil).To_DraftMessage(),
		}).To_Update()

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
			c.ctx,
			&sync.TLSyncUpdatesNotMe{
				UserId:    userId,
				AuthKeyId: authKeyId,
				Updates:   mtproto.MakeUpdatesByUpdates(updateDraftMessage),
			})
	}
}

/*
## gifVideo

wubenqideMacBook-Pro:Desktop wubenqi$ ffmpeg -i facebook-ad-gif.gif.mp4
ffmpeg version 4.1.3 Copyright (c) 2000-2019 the FFmpeg developers
  built with Apple LLVM version 10.0.1 (clang-1001.0.46.4)
  configuration: --prefix=/usr/local/Cellar/ffmpeg/4.1.3_1 --enable-shared --enable-pthreads --enable-version3 --enable-hardcoded-tables --enable-avresample --cc=clang --host-cflags='-I/Library/Java/JavaVirtualMachines/adoptopenjdk-11.0.2.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/adoptopenjdk-11.0.2.jdk/Contents/Home/include/darwin' --host-ldflags= --enable-ffplay --enable-gnutls --enable-gpl --enable-libaom --enable-libbluray --enable-libmp3lame --enable-libopus --enable-librubberband --enable-libsnappy --enable-libtesseract --enable-libtheora --enable-libvorbis --enable-libvpx --enable-libx264 --enable-libx265 --enable-libxvid --enable-lzma --enable-libfontconfig --enable-libfreetype --enable-frei0r --enable-libass --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-librtmp --enable-libspeex --enable-videotoolbox --disable-libjack --disable-indev=jack --enable-libaom --enable-libsoxr
  libavutil      56. 22.100 / 56. 22.100
  libavcodec     58. 35.100 / 58. 35.100
  libavformat    58. 20.100 / 58. 20.100
  libavdevice    58.  5.100 / 58.  5.100
  libavfilter     7. 40.101 /  7. 40.101
  libavresample   4.  0.  0 /  4.  0.  0
  libswscale      5.  3.100 /  5.  3.100
  libswresample   3.  3.100 /  3.  3.100
  libpostproc    55.  3.100 / 55.  3.100
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'facebook-ad-gif.gif.mp4':
  Metadata:
    major_brand     : isom
    minor_version   : 512
    compatible_brands: isomiso2avc1mp41
    encoder         : Lavf58.20.100
  Duration: 00:00:03.92, start: 0.000000, bitrate: 271 kb/s
    Stream #0:0(und): Video: h264 (High) (avc1 / 0x31637661), yuv420p, 286x320, 268 kb/s, 14.29 fps, 14.29 tbr, 12800 tbn, 28.57 tbc (default)
    Metadata:
      handler_name    : VideoHandler

## mp4
wubenqideMacBook-Pro:Desktop wubenqi$ ffmpeg -i 111.gif.mp4
ffmpeg version 4.1.3 Copyright (c) 2000-2019 the FFmpeg developers
  built with Apple LLVM version 10.0.1 (clang-1001.0.46.4)
  configuration: --prefix=/usr/local/Cellar/ffmpeg/4.1.3_1 --enable-shared --enable-pthreads --enable-version3 --enable-hardcoded-tables --enable-avresample --cc=clang --host-cflags='-I/Library/Java/JavaVirtualMachines/adoptopenjdk-11.0.2.jdk/Contents/Home/include -I/Library/Java/JavaVirtualMachines/adoptopenjdk-11.0.2.jdk/Contents/Home/include/darwin' --host-ldflags= --enable-ffplay --enable-gnutls --enable-gpl --enable-libaom --enable-libbluray --enable-libmp3lame --enable-libopus --enable-librubberband --enable-libsnappy --enable-libtesseract --enable-libtheora --enable-libvorbis --enable-libvpx --enable-libx264 --enable-libx265 --enable-libxvid --enable-lzma --enable-libfontconfig --enable-libfreetype --enable-frei0r --enable-libass --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-librtmp --enable-libspeex --enable-videotoolbox --disable-libjack --disable-indev=jack --enable-libaom --enable-libsoxr
  libavutil      56. 22.100 / 56. 22.100
  libavcodec     58. 35.100 / 58. 35.100
  libavformat    58. 20.100 / 58. 20.100
  libavdevice    58.  5.100 / 58.  5.100
  libavfilter     7. 40.101 /  7. 40.101
  libavresample   4.  0.  0 /  4.  0.  0
  libswscale      5.  3.100 /  5.  3.100
  libswresample   3.  3.100 /  3.  3.100
  libpostproc    55.  3.100 / 55.  3.100
Input #0, mov,mp4,m4a,3gp,3g2,mj2, from '111.gif.mp4':
  Metadata:
    major_brand     : isom
    minor_version   : 512
    compatible_brands: isomiso2avc1mp41
    encoder         : Lavf58.20.100
  Duration: 00:00:14.19, start: 0.000000, bitrate: 429 kb/s
    Stream #0:0(und): Video: h264 (High) (avc1 / 0x31637661), yuv420p, 320x180, 289 kb/s, 29.97 fps, 29.97 tbr, 30k tbn, 59.94 tbc (default)
    Metadata:
      handler_name    : Core Media Video
    Stream #0:1(und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 131 kb/s (default)
    Metadata:
      handler_name    : Core Media Audio

## tdesktop 检查 isGifv

bool FFMpegReaderImplementation::isGifv() const {
	if (_hasAudioStream) {
		return false;
	}
	if (dataSize() > Storage::kMaxAnimationInMemory) {
		return false;
	}
	if (_codecContext->codec_id != AV_CODEC_ID_H264) {
		return false;
	}
	return true;
}
*/

func (c *MessagesCore) makeMediaByInputMedia(media *mtproto.InputMedia) (messageMedia *mtproto.MessageMedia, err error) {
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
			OwnerId:    c.MD.AuthId,
			File:       media.File,
			Stickers:   nil,
			TtlSeconds: nil,
		})
		if err != nil {
			c.Logger.Errorf("UploadPhoto error: %v, by %s", err, media.DebugString())
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
			OwnerId: c.MD.AuthId,
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
			Provider:  venue.GetProvider(),
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

// TODO(@benqi): mention...
//
func (c *MessagesCore) fixMessageEntities(fromId int64, peer *mtproto.PeerUtil, noWebpage bool, message *mtproto.Message, hasBot bool) (*mtproto.Message, error) {
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
			})
			entities = append(entities, entityUrl.To_MessageEntity())
		}
	}

	if !noWebpage && firstUrl != "" {
		canEmbedLink := true
		// TODO(@benqi): check channel or super chat
		if peer.PeerType == mtproto.PEER_CHANNEL {
			// TODO(@benqi): enable check CanEmbedLinks
			// canEmbedLink = s.ChannelFacade.CanEmbedLinks(ctx, r.PeerId, r.From.Id)
		}

		if canEmbedLink {
			// TODO(@benqi): disable
			if c.svcCtx.Plugin != nil {
				ctx, _ := metadata.RpcMetadataToOutgoing(c.ctx, c.MD)
				webpage, _ := c.svcCtx.Plugin.GetWebpagePreview(ctx, firstUrl)
				if webpage != nil {
					message.Media = mtproto.MakeTLMessageMediaWebPage(&mtproto.MessageMedia{
						Webpage: webpage,
					}).To_MessageMedia()
				}
			}
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
			if v, _ := c.svcCtx.Dao.UsernameClient.UsernameResolveUsername(c.ctx, &username.TLUsernameResolveUsername{
				Username: tag.Tag,
			}); v != nil {
				if v.GetPredicateName() == mtproto.Predicate_peerUser {
					mention2.UserId_INT64 = v.UserId
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
			c.Logger.Infof("mention2: %v", mention2)
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

	if hasBot {
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
	return message, nil
}
