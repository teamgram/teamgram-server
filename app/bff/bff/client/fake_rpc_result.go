// Copyright 2022 Teamgooo Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bffproxyclient

import (
	"reflect"
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/crypto"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

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
	gNewAlgo       tg.PasswordKdfAlgoClazz
	gNewSecureAlgo tg.SecurePasswordKdfAlgoClazz
)

func init() {
	gNewAlgo = tg.MakeTLPasswordKdfAlgoModPow(&tg.TLPasswordKdfAlgoModPow{
		Salt1: gNewAlgoSalt1,
		Salt2: gNewAlgoSalt2,
		G:     gNewAlgoG,
		P:     gNewAlgoP,
	})

	gNewSecureAlgo = tg.MakeTLSecurePasswordKdfAlgoPBKDF2(&tg.TLSecurePasswordKdfAlgoPBKDF2{
		Salt: gNewSecureAlgoSalt,
	})
}

func fakeUpdatesState() *tg.TLUpdatesState {
	return tg.MakeTLUpdatesState(&tg.TLUpdatesState{
		Date: int32(time.Now().Unix()),
	})
}

func (c *BFFProxyClient2) TryReturnFakeRpcResult(object iface.TLObject) (iface.TLObject, error) {
	rt := reflect.TypeOf(object)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	switch rt.Name() {
	// langpack
	case "TLLangpackGetDifference":
		in := object.(*tg.TLLangpackGetDifference)
		return tg.MakeTLLangPackDifference(&tg.TLLangPackDifference{
			LangCode:    in.LangCode,
			FromVersion: in.FromVersion,
			Version:     in.FromVersion,
			Strings:     []tg.LangPackStringClazz{},
		}).ToLangPackDifference(), nil
	case "TLLangpackGetLangPack":
		in := object.(*tg.TLLangpackGetLangPack)
		return tg.MakeTLLangPackDifference(&tg.TLLangPackDifference{
			LangCode:    in.LangCode,
			FromVersion: 0,
			Version:     0,
			Strings:     []tg.LangPackStringClazz{},
		}).ToLangPackDifference(), nil
	case "TLLangpackGetLanguages":
		return &tg.VectorLangPackLanguage{
			Datas: []tg.LangPackLanguageClazz{},
		}, nil
	case "TLLangpackGetStrings":
		return &tg.VectorLangPackString{
			Datas: []tg.LangPackStringClazz{},
		}, nil

	// webpage
	case "TLMessagesGetWebPage":
		return tg.MakeTLWebPageEmpty(&tg.TLWebPageEmpty{
			Id: 0,
		}).ToWebPage(), nil
	case "TLMessagesGetWebPageView":
		return tg.MakeTLMessageMediaEmpty(&tg.TLMessageMediaEmpty{
			//
		}).ToMessageMedia(), nil

	// wallpaper
	case "TLAccountGetWallPapers":
		return tg.MakeTLAccountWallPapers(&tg.TLAccountWallPapers{
			Hash:       0,
			Wallpapers: []tg.WallPaperClazz{},
		}).ToAccountWallPapers(), nil

	// twofa
	case "TLAccountGetPassword":
		return tg.MakeTLAccountPassword(&tg.TLAccountPassword{
			HasRecovery:             false,
			HasSecureValues:         false,
			HasPassword:             false,
			CurrentAlgo:             nil,
			SrpB:                    nil,
			SrpId:                   nil,
			Hint:                    nil,
			EmailUnconfirmedPattern: nil,
			NewAlgo:                 gNewAlgo,
			NewSecureAlgo:           gNewSecureAlgo,
			SecureRandom:            crypto.RandomBytes(256),
		}).ToAccountPassword(), nil

	// tos
	case "TLHelpAcceptTermsOfService":
		return tg.BoolTrue, nil
	case "TLHelpGetTermsOfServiceUpdate":
		return tg.MakeTLHelpTermsOfServiceUpdateEmpty(&tg.TLHelpTermsOfServiceUpdateEmpty{
			Expires: int32(time.Now().Unix() + 3600),
		}).ToHelpTermsOfServiceUpdate(), nil

	// themes
	case "TLAccountGetThemes":
		return tg.MakeTLAccountThemes(&tg.TLAccountThemes{
			Hash:   0,
			Themes: []tg.ThemeClazz{},
		}).ToAccountThemes(), nil
	case "TLAccountGetChatThemes":
		return tg.MakeTLAccountThemes(&tg.TLAccountThemes{
			Hash:   0,
			Themes: []tg.ThemeClazz{},
		}).ToAccountThemes(), nil

	// tdesktop main screen startup placeholders
	case "TLAccountUpdateStatus":
		return tg.BoolTrue, nil
	case "TLUpdatesGetState":
		return fakeUpdatesState(), nil
	case "TLMessagesGetDialogs":
		return tg.MakeTLMessagesDialogs(&tg.TLMessagesDialogs{
			Dialogs:  []tg.DialogClazz{},
			Messages: []tg.MessageClazz{},
			Chats:    []tg.ChatClazz{},
			Users:    []tg.UserClazz{},
		}).ToMessagesDialogs(), nil
	case "TLMessagesGetPinnedDialogs":
		return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
			Dialogs:  []tg.DialogClazz{},
			Messages: []tg.MessageClazz{},
			Chats:    []tg.ChatClazz{},
			Users:    []tg.UserClazz{},
			State:    fakeUpdatesState(),
		}).ToMessagesPeerDialogs(), nil
	case "TLMessagesGetDialogFilters":
		return tg.MakeTLMessagesDialogFilters(&tg.TLMessagesDialogFilters{
			Filters: []tg.DialogFilterClazz{},
		}).ToMessagesDialogFilters(), nil
	case "TLHelpGetPeerColors", "TLHelpGetPeerProfileColors":
		return tg.MakeTLHelpPeerColors(&tg.TLHelpPeerColors{
			Hash:   0,
			Colors: []tg.HelpPeerColorOptionClazz{},
		}).ToHelpPeerColors(), nil
	case "TLMessagesGetAvailableEffects":
		return tg.MakeTLMessagesAvailableEffects(&tg.TLMessagesAvailableEffects{
			Hash:      0,
			Effects:   []tg.AvailableEffectClazz{},
			Documents: []tg.DocumentClazz{},
		}).ToMessagesAvailableEffects(), nil
	case "TLAccountGetDefaultEmojiStatuses":
		return tg.MakeTLAccountEmojiStatuses(&tg.TLAccountEmojiStatuses{
			Hash:     0,
			Statuses: []tg.EmojiStatusClazz{},
		}).ToAccountEmojiStatuses(), nil
	case "TLUsersGetFullUser":
		return tg.MakeTLUsersUserFull(&tg.TLUsersUserFull{
			FullUser: tg.MakeTLUserFull(&tg.TLUserFull{
				Settings:       tg.MakeTLPeerSettings(&tg.TLPeerSettings{}).ToPeerSettings(),
				NotifySettings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}).ToPeerNotifySettings(),
			}).ToUserFull(),
			Chats: []tg.ChatClazz{},
			Users: []tg.UserClazz{},
		}).ToUsersUserFull(), nil
	case "TLAccountGetNotifySettings":
		return tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}).ToPeerNotifySettings(), nil
	case "TLMessagesGetEmojiGroups", "TLMessagesGetEmojiStickerGroups":
		return tg.MakeTLMessagesEmojiGroups(&tg.TLMessagesEmojiGroups{
			Hash:   0,
			Groups: []tg.EmojiGroupClazz{},
		}).ToMessagesEmojiGroups(), nil
	case "TLMessagesGetAttachMenuBots":
		return tg.MakeTLAttachMenuBots(&tg.TLAttachMenuBots{
			Hash:  0,
			Bots:  []tg.AttachMenuBotClazz{},
			Users: []tg.UserClazz{},
		}).ToAttachMenuBots(), nil
	case "TLMessagesGetQuickReplies":
		return tg.MakeTLMessagesQuickReplies(&tg.TLMessagesQuickReplies{
			QuickReplies: []tg.QuickReplyClazz{},
			Messages:     []tg.MessageClazz{},
			Chats:        []tg.ChatClazz{},
			Users:        []tg.UserClazz{},
		}).ToMessagesQuickReplies(), nil
	case "TLStoriesGetAllStories":
		return tg.MakeTLStoriesAllStories(&tg.TLStoriesAllStories{
			Count:       0,
			State:       "",
			PeerStories: []tg.PeerStoriesClazz{},
			Chats:       []tg.ChatClazz{},
			Users:       []tg.UserClazz{},
			StealthMode: tg.MakeTLStoriesStealthMode(&tg.TLStoriesStealthMode{}).ToStoriesStealthMode(),
		}).ToStoriesAllStories(), nil
	case "TLStoriesGetStoriesArchive":
		return tg.MakeTLStoriesStories(&tg.TLStoriesStories{
			Count:       0,
			Stories:     []tg.StoryItemClazz{},
			PinnedToTop: []int32{},
			Chats:       []tg.ChatClazz{},
			Users:       []tg.UserClazz{},
		}).ToStoriesStories(), nil

	// stickers
	case "TLMessagesGetAllStickers":
		return tg.MakeTLMessagesAllStickers(&tg.TLMessagesAllStickers{
			Hash: 0,
			Sets: []tg.StickerSetClazz{},
		}).ToMessagesAllStickers(), nil
	case "TLMessagesGetStickerSet":
		return tg.MakeTLMessagesStickerSetNotModified(&tg.TLMessagesStickerSetNotModified{}).ToMessagesStickerSet(), nil
	case "TLMessagesGetFavedStickers":
		return tg.MakeTLMessagesFavedStickers(&tg.TLMessagesFavedStickers{
			Hash:     0,
			Packs:    []tg.StickerPackClazz{},
			Stickers: []tg.DocumentClazz{},
		}).ToMessagesFavedStickers(), nil
	case "TLMessagesGetRecentStickers":
		return tg.MakeTLMessagesRecentStickers(&tg.TLMessagesRecentStickers{
			Hash:     0,
			Packs:    []tg.StickerPackClazz{},
			Stickers: []tg.DocumentClazz{},
			Dates:    []int32{},
		}).ToMessagesRecentStickers(), nil
	case "TLMessagesGetStickers":
		return tg.MakeTLMessagesStickers(&tg.TLMessagesStickers{
			Hash:     0,
			Stickers: []tg.DocumentClazz{},
		}).ToMessagesStickers(), nil
	case "TLMessagesGetFeaturedStickers":
		return tg.MakeTLMessagesFeaturedStickers(&tg.TLMessagesFeaturedStickers{
			Hash:   0,
			Count:  0,
			Sets:   []tg.StickerSetCoveredClazz{},
			Unread: []int64{},
		}).ToMessagesFeaturedStickers(), nil

	// promodata
	case "TLHelpGetPromoData":
		return tg.MakeTLHelpPromoDataEmpty(&tg.TLHelpPromoDataEmpty{
			Expires: int32(time.Now().Unix() + 60*60),
		}).ToHelpPromoData(), nil

	// premium
	case "TLHelpGetPremiumPromo":
		return tg.MakeTLHelpPremiumPromo(&tg.TLHelpPremiumPromo{
			StatusText:     "Premium",
			StatusEntities: []tg.MessageEntityClazz{},
			VideoSections:  []string{},
			Videos:         []tg.DocumentClazz{},
			PeriodOptions:  []tg.PremiumSubscriptionOptionClazz{},
			Users:          []tg.UserClazz{},
		}).ToHelpPremiumPromo(), nil

	// reaction
	//case "TLReactionGetTopReactions":
	//	return tg.MakeTLMessagesReactions(&tg.TLMessagesReactions{
	//		Hash:      0,
	//		Reactions: []tg.ReactionClazz{},
	//	}).ToMessagesReactions(), nil
	//
	//case "TLMessagesGetRecentReactions":
	case "TLMessagesGetAvailableReactions":
		return tg.MakeTLMessagesAvailableReactions(&tg.TLMessagesAvailableReactions{
			Hash:      0,
			Reactions: []tg.AvailableReactionClazz{},
		}).ToMessagesAvailableReactions(), nil
	case "TLMessagesGetTopReactions":
		return tg.MakeTLMessagesReactions(&tg.TLMessagesReactions{
			Hash:      0,
			Reactions: []tg.ReactionClazz{},
		}).ToMessagesReactions(), nil
	case "TLMessagesGetRecentReactions":
		return tg.MakeTLMessagesReactions(&tg.TLMessagesReactions{
			Hash:      0,
			Reactions: []tg.ReactionClazz{},
		}).ToMessagesReactions(), nil
	case "TLMessagesGetDefaultTagReactions":
		return tg.MakeTLMessagesReactions(&tg.TLMessagesReactions{
			Hash:      0,
			Reactions: []tg.ReactionClazz{},
		}).ToMessagesReactions(), nil
	case "TLMessagesGetSavedReactionTags":
		return tg.MakeTLMessagesSavedReactionTags(&tg.TLMessagesSavedReactionTags{
			Tags: []tg.SavedReactionTagClazz{},
			Hash: 0,
		}).ToMessagesSavedReactionTags(), nil
	case "TLMessagesGetScheduledHistory":
		return tg.MakeTLMessagesMessages(&tg.TLMessagesMessages{
			Messages: []tg.MessageClazz{},
			Topics:   []tg.ForumTopicClazz{},
			Chats:    []tg.ChatClazz{},
			Users:    []tg.UserClazz{},
		}).ToMessagesMessages(), nil
	case "TLMessagesGetEmojiKeywordsLanguages":
		return &tg.VectorEmojiLanguage{
			Datas: []tg.EmojiLanguageClazz{},
		}, nil

	case "TLContactsGetContacts":
		return tg.MakeTLContactsContacts(&tg.TLContactsContacts{
			Contacts:   []tg.ContactClazz{},
			SavedCount: 0,
			Users:      []tg.UserClazz{},
		}).ToContactsContacts(), nil
	case "TLContactsGetTopPeers":
		return tg.MakeTLContactsTopPeers(&tg.TLContactsTopPeers{
			Categories: []tg.TopPeerCategoryPeersClazz{},
			Chats:      []tg.ChatClazz{},
			Users:      []tg.UserClazz{},
		}).ToContactsTopPeers(), nil

	case "TLPaymentsGetStarGiftActiveAuctions":
		return tg.MakeTLPaymentsStarGiftActiveAuctions(&tg.TLPaymentsStarGiftActiveAuctions{
			Auctions: []tg.StarGiftActiveAuctionStateClazz{},
			Users:    []tg.UserClazz{},
			Chats:    []tg.ChatClazz{},
		}).ToPaymentsStarGiftActiveAuctions(), nil

		//// gifs
		//case "TLMessagesGetSavedGifs":
		//	return mtproto.MakeTLMessagesSavedGifs(&mtproto.Messages_SavedGifs{
		//		Hash: 0,
		//		Gifs: []*mtproto.Document{},
		//	}).To_Messages_SavedGifs(), nil
		//case "TLMessagesSaveGif":
		//	return mtproto.BoolTrue, nil
		//
		//// promodata
		//case "TLHelpHidePromoData":
		//	return mtproto.BoolTrue, nil
		//
		//// emoji
		//case "TLMessagesGetEmojiKeywords":
		//	in := object.(*mtproto.TLMessagesGetEmojiKeywords)
		//	return mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
		//		LangCode:    in.LangCode,
		//		FromVersion: 0,
		//		Version:     0,
		//		Keywords:    []*mtproto.EmojiKeyword{},
		//	}).To_EmojiKeywordsDifference(), nil
		//case "TLMessagesGetEmojiKeywordsDifference":
		//	in := object.(*mtproto.TLMessagesGetEmojiKeywordsDifference)
		//	return mtproto.MakeTLEmojiKeywordsDifference(&mtproto.EmojiKeywordsDifference{
		//		LangCode:    in.LangCode,
		//		FromVersion: in.FromVersion,
		//		Version:     in.FromVersion,
		//		Keywords:    []*mtproto.EmojiKeyword{},
		//	}).To_EmojiKeywordsDifference(), nil
		//case "TLMessagesGetEmojiKeywordsLanguages":
		//	return &mtproto.Vector_EmojiLanguage{
		//		Datas: []*mtproto.EmojiLanguage{},
		//	}, nil
		//
		//// reports
		//case "TLAccountReportPeer":
		//	return mtproto.BoolTrue, nil
		//case "TLAccountReportProfilePhoto":
		//	return mtproto.BoolTrue, nil
		//case "TLChannelsReportSpam":
		//	return mtproto.BoolTrue, nil
		//case "TLMessagesReport":
		//	return mtproto.BoolTrue, nil
		//case "TLMessagesReportSpam":
		//	return mtproto.BoolTrue, nil
		//
		//// phone
		//case "TLPhoneGetCallConfig":
		//	return mtproto.MakeTLDataJSON(&mtproto.DataJSON{
		//		Data: "{}",
		//	}).To_DataJSON(), nil
		//
		//case "TLAccountGetAuthorizations":
		//	return mtproto.MakeTLAccountAuthorizations(&mtproto.Account_Authorizations{
		//		AuthorizationTtlDays: 0,
		//		Authorizations:       []*mtproto.Authorization{},
		//	}).To_Account_Authorizations(), nil
		//
		//case "TLAccountGetWebAuthorizations":
		//	return mtproto.MakeTLAccountWebAuthorizations(&mtproto.Account_WebAuthorizations{
		//		Authorizations: []*mtproto.WebAuthorization{},
		//		Users:          []*mtproto.User{},
		//	}).To_Account_WebAuthorizations(), nil
	}

	logx.Errorf("%s blocked, License key from https://teamgram.net required to unlock enterprise features.", rt.Name())
	return nil, tg.ErrEnterpriseIsBlocked
}
