// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bff_proxy_client

import (
	"reflect"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	gNewAlgoSalt1 = []byte{0xEC, 0xF8, 0x73, 0x76, 0x65, 0xBC, 0x77, 0x5A}
	gNewAlgoSalt2 = []byte{0xBE, 0xDE, 0x48, 0x88, 0x8C, 0x0F, 0x42, 0xAC, 0x34, 0xFF, 0xD1, 0xD4, 0x93, 0x5D, 0x8B, 0x21}

	gNewAlgoP = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}

	gNewAlgoG = int32(3)

	// salt: 7D 04 B3 4B 94 82 8C 3D [8 BYTES],
	gNewSecureAlgoSalt = []byte{0x7D, 0x04, 0xB3, 0x4B, 0x94, 0x82, 0x8C, 0x3D}
)

var (
	gNewAlgo       *mtproto.PasswordKdfAlgo
	gNewSecureAlgo *mtproto.SecurePasswordKdfAlgo
)

func init() {
	gNewAlgo = mtproto.MakeTLPasswordKdfAlgoModPow(&mtproto.PasswordKdfAlgo{
		Salt1: gNewAlgoSalt1,
		Salt2: gNewAlgoSalt2,
		G:     gNewAlgoG,
		P:     gNewAlgoP,
	}).To_PasswordKdfAlgo()

	gNewSecureAlgo = mtproto.MakeTLSecurePasswordKdfAlgoPBKDF2(&mtproto.SecurePasswordKdfAlgo{
		Salt: gNewSecureAlgoSalt,
	}).To_SecurePasswordKdfAlgo()
}

func (c *BFFProxyClient) TryReturnFakeRpcResult(object mtproto.TLObject) (mtproto.TLObject, error) {
	rt := reflect.TypeOf(object)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	switch rt.Name() {
	// langpack
	case "TLLangpackGetDifference":
		in := object.(*mtproto.TLLangpackGetDifference)
		return mtproto.MakeTLLangPackDifference(&mtproto.LangPackDifference{
			LangCode:    in.GetLangCode(),
			FromVersion: in.GetFromVersion(),
			Version:     in.GetFromVersion(),
			Strings:     []*mtproto.LangPackString{},
		}).To_LangPackDifference(), nil
	case "TLLangpackGetLangPack":
		in := object.(*mtproto.TLLangpackGetLangPack)
		return mtproto.MakeTLLangPackDifference(&mtproto.LangPackDifference{
			LangCode:    in.GetLangCode(),
			FromVersion: 0,
			Version:     0,
			Strings:     []*mtproto.LangPackString{},
		}).To_LangPackDifference(), nil
	case "TLLangpackGetLanguages":
		return &mtproto.Vector_LangPackLanguage{
			Datas: []*mtproto.LangPackLanguage{},
		}, nil
	case "TLLangpackGetStrings":
		return &mtproto.Vector_LangPackString{
			Datas: []*mtproto.LangPackString{},
		}, nil

	// webpage
	case "TLMessagesGetWebPage":
		return mtproto.MakeTLWebPageEmpty(&mtproto.WebPage{
			Id: 0,
		}).To_WebPage(), nil
	case "TLMessagesGetWebPageView":
		return mtproto.MakeTLMessageMediaEmpty(&mtproto.MessageMedia{
			//
		}).To_MessageMedia(), nil

	// wallpaper
	case "TLAccountGetWallPapers":
		return mtproto.MakeTLAccountWallPapers(&mtproto.Account_WallPapers{
			Hash:       0,
			Wallpapers: []*mtproto.WallPaper{},
		}).To_Account_WallPapers(), nil

	// twofa
	case "TLAccountGetPassword":
		return mtproto.MakeTLAccountPassword(&mtproto.Account_Password{
			HasRecovery:             false,
			HasSecureValues:         false,
			HasPassword:             false,
			CurrentAlgo:             nil,
			Srp_B:                   nil,
			SrpId:                   nil,
			Hint:                    nil,
			EmailUnconfirmedPattern: nil,
			NewAlgo:                 gNewAlgo,
			NewSecureAlgo:           gNewSecureAlgo,
			SecureRandom:            crypto.RandomBytes(256),
		}).To_Account_Password(), nil

	// tos
	case "TLHelpAcceptTermsOfService":
		return mtproto.BoolTrue, nil
	case "TLHelpGetTermsOfServiceUpdate":
		return mtproto.MakeTLHelpTermsOfServiceUpdateEmpty(&mtproto.Help_TermsOfServiceUpdate{
			Expires: int32(time.Now().Unix() + 3600),
		}).To_Help_TermsOfServiceUpdate(), nil

	// themes
	case "TLAccountGetThemes":
		return mtproto.MakeTLAccountThemes(&mtproto.Account_Themes{
			Hash:   0,
			Themes: []*mtproto.Theme{},
		}).To_Account_Themes(), nil
	case "TLAccountGetChatThemes":
		return mtproto.MakeTLAccountThemes(&mtproto.Account_Themes{
			Hash:   0,
			Themes: []*mtproto.Theme{},
		}).To_Account_Themes(), nil

	// stickers
	case "TLMessagesGetAllStickers":
		return mtproto.MakeTLMessagesAllStickers(&mtproto.Messages_AllStickers{
			Hash: 0,
			Sets: []*mtproto.StickerSet{},
		}).To_Messages_AllStickers(), nil
	case "TLMessagesGetArchivedStickers":
		return mtproto.MakeTLMessagesArchivedStickers(&mtproto.Messages_ArchivedStickers{
			Count: 0,
			Sets:  []*mtproto.StickerSetCovered{},
		}).To_Messages_ArchivedStickers(), nil
	case "TLMessagesGetFavedStickers":
		return mtproto.MakeTLMessagesFavedStickers(&mtproto.Messages_FavedStickers{
			Hash:     0,
			Packs:    []*mtproto.StickerPack{},
			Stickers: []*mtproto.Document{},
		}).To_Messages_FavedStickers(), nil
	case "TLMessagesGetMaskStickers":
		return mtproto.MakeTLMessagesAllStickers(&mtproto.Messages_AllStickers{
			Hash: 0,
			Sets: []*mtproto.StickerSet{},
		}).To_Messages_AllStickers(), nil
	case "TLMessagesGetOldFeaturedStickers":
		return mtproto.MakeTLMessagesFeaturedStickers(&mtproto.Messages_FeaturedStickers{
			Count:  0,
			Hash:   0,
			Sets:   []*mtproto.StickerSetCovered{},
			Unread: []int64{},
		}).To_Messages_FeaturedStickers(), nil
	case "TLMessagesGetRecentStickers":
		return mtproto.MakeTLMessagesRecentStickers(&mtproto.Messages_RecentStickers{
			Hash:     0,
			Packs:    []*mtproto.StickerPack{},
			Stickers: []*mtproto.Document{},
			Dates:    []int32{},
		}).To_Messages_RecentStickers(), nil
	case "TLMessagesGetStickers":
		return mtproto.MakeTLMessagesStickers(&mtproto.Messages_Stickers{
			Hash:     0,
			Stickers: []*mtproto.Document{},
		}).To_Messages_Stickers(), nil
	case "TLMessagesGetFeaturedStickers":
		return mtproto.MakeTLMessagesFeaturedStickers(&mtproto.Messages_FeaturedStickers{
			Count:  0,
			Hash:   0,
			Sets:   []*mtproto.StickerSetCovered{},
			Unread: []int64{},
		}).To_Messages_FeaturedStickers(), nil
	case "TLMessagesGetStickerSet":
		return nil, mtproto.ErrStickerIdInvalid

	// 	scheduledmessages
	case "TLMessagesGetScheduledMessages":
		return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages(), nil

	// reactions
	case "TLMessagesGetAvailableReactions":
		return mtproto.MakeTLMessagesAvailableReactions(&mtproto.Messages_AvailableReactions{
			Hash:      0,
			Reactions: []*mtproto.AvailableReaction{},
		}).To_Messages_AvailableReactions(), nil

	// folders
	case "TLMessagesGetDialogFilters":
		return &mtproto.Vector_DialogFilter{
			Datas: []*mtproto.DialogFilter{},
		}, nil

	// gifs
	case "TLMessagesGetSavedGifs":
		return mtproto.MakeTLMessagesSavedGifs(&mtproto.Messages_SavedGifs{
			Hash: 0,
			Gifs: []*mtproto.Document{},
		}).To_Messages_SavedGifs(), nil
	case "TLMessagesSaveGif":
		return mtproto.BoolTrue, nil

	// promodata
	case "TLHelpGetPromoData":
		return mtproto.MakeTLHelpPromoDataEmpty(&mtproto.Help_PromoData{
			Expires: int32(time.Now().Unix() + 60*60),
		}).To_Help_PromoData(), nil
	case "TLHelpHidePromoData":
		return mtproto.BoolTrue, nil

	// emoji
	case "TLMessagesGetEmojiKeywords":
		in := object.(*mtproto.TLMessagesGetEmojiKeywords)
		return mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
			LangCode:    in.LangCode,
			FromVersion: 0,
			Version:     0,
			Keywords:    []*mtproto.EmojiKeyword{},
		}).To_EmojiKeywordsDifference(), nil
	case "TLMessagesGetEmojiKeywordsDifference":
		in := object.(*mtproto.TLMessagesGetEmojiKeywordsDifference)
		return mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
			LangCode:    in.LangCode,
			FromVersion: in.FromVersion,
			Version:     in.FromVersion,
			Keywords:    []*mtproto.EmojiKeyword{},
		}).To_EmojiKeywordsDifference(), nil
	case "TLMessagesGetEmojiKeywordsLanguages":
		return &mtproto.Vector_EmojiLanguage{
			Datas: []*mtproto.EmojiLanguage{},
		}, nil

	// reports
	case "TLAccountReportPeer":
		return mtproto.BoolTrue, nil
	case "TLAccountReportProfilePhoto":
		return mtproto.BoolTrue, nil
	case "TLChannelsReportSpam":
		return mtproto.BoolTrue, nil
	case "TLMessagesReport":
		return mtproto.BoolTrue, nil
	case "TLMessagesReportSpam":
		return mtproto.BoolTrue, nil

	// phone
	case "TLPhoneGetCallConfig":
		return mtproto.MakeTLDataJSON(&mtproto.DataJSON{
			Data: "{}",
		}).To_DataJSON(), nil

	case "TLAccountGetAuthorizations":
		return mtproto.MakeTLAccountAuthorizations(&mtproto.Account_Authorizations{
			AuthorizationTtlDays: 0,
			Authorizations:       []*mtproto.Authorization{},
		}).To_Account_Authorizations(), nil

	case "TLAccountGetWebAuthorizations":
		return mtproto.MakeTLAccountWebAuthorizations(&mtproto.Account_WebAuthorizations{
			Authorizations: []*mtproto.WebAuthorization{},
			Users:          []*mtproto.User{},
		}).To_Account_WebAuthorizations(), nil
	}

	logx.Errorf("%s blocked, License key from https://teamgram.net required to unlock enterprise features.", rt.Name())
	return nil, mtproto.ErrEnterpriseIsBlocked
}
