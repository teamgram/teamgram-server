// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package user

/*
## Inline bot samples
Here are some sample inline bots, in case you’re curious to see one in action. Try any of these:
@gif – GIF search
@vid – Video search
@pic – Yandex image search
@bing – Bing image search
@wiki – Wikipedia search
@imdb – IMDB search
@bold – Make bold, italic or fixed sys text

## NEW
@youtube - Connect your account for personalized results
@music - Search and send classical music
@foursquare – Find and send venue addresses
@sticker – Find and send stickers based on emoji
*/
const (
	BotFatherId   = int64(6)
	BotFatherName = "BotFather"
)

const (
	BotGifId          = int64(101)
	BotGifName        = "gif"
	BotVidId          = int64(102)
	BotVidName        = "vid"
	BotPicId          = int64(103)
	BotPicName        = "pic"
	BotBingId         = int64(104)
	BotBingName       = "bing"
	BotWikiId         = int64(105)
	BotWikiName       = "wiki"
	BotImdbId         = int64(106)
	BotImdbName       = "imdb"
	BotBoldId         = int64(107)
	BotBoldName       = "bold"
	BotYoutubeId      = int64(108)
	BotYoutubeName    = "youtube"
	BotMusicId        = int64(109)
	BotMusicName      = "music"
	BotFoursquareId   = int64(110)
	BotFoursquareName = "foursquare"
	BotStickerId      = int64(111)
	BotStickerName    = "sticker"
)

var (
	botIdNameMap = map[int64]string{
		BotFatherId:     BotFatherName,
		BotGifId:        BotGifName,
		BotVidId:        BotVidName,
		BotPicId:        BotPicName,
		BotBingId:       BotBingName,
		BotWikiId:       BotWikiName,
		BotImdbId:       BotImdbName,
		BotBoldId:       BotBoldName,
		BotYoutubeId:    BotYoutubeName,
		BotMusicId:      BotMusicName,
		BotFoursquareId: BotFoursquareName,
		BotStickerId:    BotStickerName,
	}

	botNameIdMap = map[string]int64{
		BotFatherName:     BotFatherId,
		BotGifName:        BotGifId,
		BotVidName:        BotVidId,
		BotPicName:        BotPicId,
		BotBingName:       BotBingId,
		BotWikiName:       BotWikiId,
		BotImdbName:       BotImdbId,
		BotBoldName:       BotBoldId,
		BotYoutubeName:    BotYoutubeId,
		BotMusicName:      BotMusicId,
		BotFoursquareName: BotFoursquareId,
		BotStickerName:    BotStickerId,
	}
)

func GetBotNameById(id int64) (n string) {
	n, _ = botIdNameMap[id]
	return
}

func GetBotIdByName(n string) (id int64) {
	id, _ = botNameIdMap[n]
	return
}

func IsBotFather(id int64) bool {
	return id == BotFatherId
}

func IsBotBing(id int64) bool {
	return id == BotBingId
}

func IsBotPic(id int64) bool {
	return id == BotPicId
}

func IsBotGif(id int64) bool {
	return id == BotGifId
}

func IsBotFoursquare(id int64) bool {
	return id == BotFoursquareId
}
