package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestDialogGetDialogsReturnsSinglePlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogs(&dialog.TLDialogGetDialogs{
		UserId:        1,
		ExcludePinned: tg.BoolTrueClazz,
		FolderId:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs vector, got nil")
	}
	if len(result.Datas) != 1 {
		t.Fatalf("expected single placeholder dialog, got %d items", len(result.Datas))
	}
	dialogExt, ok := result.Datas[0].(*dialog.TLDialogExt)
	if !ok {
		t.Fatalf("expected dialogExt placeholder, got %T", result.Datas[0])
	}
	placeholderDialog, ok := dialogExt.Dialog.(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected embedded dialog placeholder, got %T", dialogExt.Dialog)
	}
	if placeholderDialog.TopMessage != 10 {
		t.Fatalf("expected top_message=10, got %d", placeholderDialog.TopMessage)
	}
}

func TestDialogGetMyDialogsDataReturnsUserPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetMyDialogsData(&dialog.TLDialogGetMyDialogsData{
		UserId: 1,
		User:   true,
		Chat:   true,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected dialogs data, got nil")
	}

	simpleData, ok := result.Clazz.(*dialog.TLSimpleDialogsData)
	if !ok {
		t.Fatalf("expected simpleDialogsData placeholder, got %T", result.Clazz)
	}
	if len(simpleData.Users) != 1 || simpleData.Users[0] != 1 {
		t.Fatalf("expected user placeholder id=1, got %#v", simpleData.Users)
	}
	if len(simpleData.Chats) != 0 || len(simpleData.Channels) != 0 {
		t.Fatalf("expected empty chats/channels, got chats=%d channels=%d",
			len(simpleData.Chats), len(simpleData.Channels))
	}
}

func TestDialogGetDialogByIdReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogById(&dialog.TLDialogGetDialogById{
		UserId:   1,
		PeerType: tg.PEER_CHAT,
		PeerId:   42,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.Clazz == nil {
		t.Fatal("expected dialogExt placeholder, got nil")
	}
	dialogExt, ok := result.Clazz.(*dialog.TLDialogExt)
	if !ok {
		t.Fatalf("expected dialogExt, got %T", result.Clazz)
	}
	placeholderDialog, ok := dialogExt.Dialog.(*tg.TLDialog)
	if !ok {
		t.Fatalf("expected embedded dialog placeholder, got %T", dialogExt.Dialog)
	}
	peer, ok := placeholderDialog.Peer.(*tg.TLPeerChat)
	if !ok || peer.ChatId != 42 {
		t.Fatalf("expected peerChat(42), got %#v", placeholderDialog.Peer)
	}
}

func TestDialogGetDialogsByIdListReturnsPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsByIdList(&dialog.TLDialogGetDialogsByIdList{
		UserId: 1,
		IdList: []int64{11, 22},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected vector, got nil")
	}
	if len(result.Datas) != 2 {
		t.Fatalf("expected two placeholders, got %d", len(result.Datas))
	}
}

func TestDialogGetDialogsCountReturnsPlaceholderCount(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsCount(&dialog.TLDialogGetDialogsCount{
		UserId:        1,
		ExcludePinned: tg.BoolFalseClazz,
		FolderId:      0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 1 {
		t.Fatalf("expected placeholder count=1, got %#v", result)
	}
}

func TestDialogGetPinnedDialogsReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetPinnedDialogs(&dialog.TLDialogGetPinnedDialogs{
		UserId:   1,
		FolderId: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one pinned placeholder, got %#v", result)
	}
}

func TestDialogGetTopMessageReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetTopMessage(&dialog.TLDialogGetTopMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || result.V != 10 {
		t.Fatalf("expected top message=10, got %#v", result)
	}
}

func TestDialogPinnedMessagePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	getResult, err := c.DialogGetUserPinnedMessage(&dialog.TLDialogGetUserPinnedMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if getResult == nil || getResult.V != 1 {
		t.Fatalf("expected pinned placeholder id=1, got %#v", getResult)
	}

	updateResult, err := c.DialogUpdateUserPinnedMessage(&dialog.TLDialogUpdateUserPinnedMessage{
		UserId:      1,
		PeerType:    tg.PEER_USER,
		PeerId:      2,
		PinnedMsgId: 9,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(updateResult) {
		t.Fatalf("expected boolTrue placeholder, got %#v", updateResult)
	}
}

func TestDialogWritePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)
	unreadCount := int32(3)

	insertResult, err := c.DialogInsertOrUpdateDialog(&dialog.TLDialogInsertOrUpdateDialog{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(insertResult) {
		t.Fatalf("expected insert/update placeholder boolTrue, got %#v", insertResult)
	}

	deleteResult, err := c.DialogDeleteDialog(&dialog.TLDialogDeleteDialog{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(deleteResult) {
		t.Fatalf("expected delete placeholder boolTrue, got %#v", deleteResult)
	}

	unreadResult, err := c.DialogUpdateUnreadCount(&dialog.TLDialogUpdateUnreadCount{
		UserId:      1,
		PeerType:    tg.PEER_USER,
		PeerId:      2,
		UnreadCount: &unreadCount,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(unreadResult) {
		t.Fatalf("expected unread-count placeholder boolTrue, got %#v", unreadResult)
	}

	markResult, err := c.DialogMarkDialogUnread(&dialog.TLDialogMarkDialogUnread{
		UserId:     1,
		PeerType:   tg.PEER_USER,
		PeerId:     2,
		UnreadMark: tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(markResult) {
		t.Fatalf("expected unread-mark placeholder boolTrue, got %#v", markResult)
	}
}

func TestDialogGetDialogsByOffsetDateReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogsByOffsetDate(&dialog.TLDialogGetDialogsByOffsetDate{
		UserId:        1,
		ExcludePinned: tg.BoolFalseClazz,
		OffsetDate:    0,
		Limit:         10,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one offset-date placeholder, got %#v", result)
	}
}

func TestDialogGetDialogUnreadMarkListReturnsPlaceholder(t *testing.T) {
	c := New(context.Background(), nil)

	result, err := c.DialogGetDialogUnreadMarkList(&dialog.TLDialogGetDialogUnreadMarkList{
		UserId: 1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result == nil || len(result.Datas) != 1 {
		t.Fatalf("expected one unread-mark placeholder, got %#v", result)
	}
	peer, ok := result.Datas[0].(*tg.TLDialogPeer)
	if !ok {
		t.Fatalf("expected dialogPeer placeholder, got %T", result.Datas[0])
	}
	userPeer, ok := peer.Peer.(*tg.TLPeerUser)
	if !ok || userPeer.UserId != 1 {
		t.Fatalf("expected peerUser(1), got %#v", peer.Peer)
	}
}

func TestDialogPinnedAndSavedPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	pinResult, err := c.DialogToggleDialogPin(&dialog.TLDialogToggleDialogPin{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
		Pinned:   tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if pinResult == nil || pinResult.V != 1 {
		t.Fatalf("expected pinned rank placeholder=1, got %#v", pinResult)
	}

	reorderResult, err := c.DialogReorderPinnedDialogs(&dialog.TLDialogReorderPinnedDialogs{
		UserId:   1,
		Force:    tg.BoolTrueClazz,
		FolderId: 0,
		IdList:   []int64{2, 3},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(reorderResult) {
		t.Fatalf("expected reorder pinned boolTrue, got %#v", reorderResult)
	}

	savedResult, err := c.DialogGetSavedDialogs(&dialog.TLDialogGetSavedDialogs{
		UserId:   1,
		Limit:    10,
		OffsetId: 0,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if savedResult == nil || savedResult.Clazz == nil {
		t.Fatal("expected savedDialogList placeholder, got nil")
	}
	savedList, ok := savedResult.Clazz.(*dialog.TLSavedDialogList)
	if !ok || savedList.Count != 1 || len(savedList.Dialogs) != 1 {
		t.Fatalf("expected one saved dialog placeholder, got %#v", savedResult.Clazz)
	}

	pinnedSavedResult, err := c.DialogGetPinnedSavedDialogs(&dialog.TLDialogGetPinnedSavedDialogs{
		UserId: 1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	pinnedSavedList, ok := pinnedSavedResult.Clazz.(*dialog.TLSavedDialogList)
	if !ok || pinnedSavedList.Count != 1 {
		t.Fatalf("expected pinned saved dialog placeholder, got %#v", pinnedSavedResult.Clazz)
	}
	savedDialog, ok := pinnedSavedList.Dialogs[0].(*tg.TLSavedDialog)
	if !ok || !savedDialog.Pinned {
		t.Fatalf("expected pinned savedDialog placeholder, got %#v", pinnedSavedList.Dialogs[0])
	}

	toggleSavedResult, err := c.DialogToggleSavedDialogPin(&dialog.TLDialogToggleSavedDialogPin{
		UserId: 1,
		Peer: tg.MakeTLPeerUtil(&tg.TLPeerUtil{
			SelfId:   1,
			PeerType: tg.PEER_USER,
			PeerId:   2,
		}),
		Pinned: tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(toggleSavedResult) {
		t.Fatalf("expected toggle saved boolTrue, got %#v", toggleSavedResult)
	}

	reorderSavedResult, err := c.DialogReorderPinnedSavedDialogs(&dialog.TLDialogReorderPinnedSavedDialogs{
		UserId: 1,
		Force:  tg.BoolTrueClazz,
		Order: []tg.PeerUtilClazz{
			tg.MakeTLPeerUtil(&tg.TLPeerUtil{
				SelfId:   1,
				PeerType: tg.PEER_USER,
				PeerId:   2,
			}),
		},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(reorderSavedResult) {
		t.Fatalf("expected reorder saved boolTrue, got %#v", reorderSavedResult)
	}
}

func TestDialogThemeWallpaperAndTTLPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	readParticipants, err := c.DialogGetChannelMessageReadParticipants(&dialog.TLDialogGetChannelMessageReadParticipants{
		UserId:    1,
		ChannelId: 100,
		MsgId:     10,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(readParticipants.Datas) != 1 || readParticipants.Datas[0] != 1 {
		t.Fatalf("expected single read participant=1, got %#v", readParticipants.Datas)
	}

	ttlResult, err := c.DialogSetHistoryTTL(&dialog.TLDialogSetHistoryTTL{
		UserId:    1,
		PeerType:  tg.PEER_USER,
		PeerId:    2,
		TtlPeriod: 86400,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(ttlResult) {
		t.Fatalf("expected history ttl boolTrue, got %#v", ttlResult)
	}

	themeResult, err := c.DialogSetChatTheme(&dialog.TLDialogSetChatTheme{
		UserId:        1,
		PeerType:      tg.PEER_USER,
		PeerId:        2,
		ThemeEmoticon: "🙂",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(themeResult) {
		t.Fatalf("expected chat theme boolTrue, got %#v", themeResult)
	}

	wallpaperResult, err := c.DialogSetChatWallpaper(&dialog.TLDialogSetChatWallpaper{
		UserId:      1,
		PeerType:    tg.PEER_USER,
		PeerId:      2,
		WallpaperId: 9,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(wallpaperResult) {
		t.Fatalf("expected chat wallpaper boolTrue, got %#v", wallpaperResult)
	}
}

func TestDialogDraftPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	saveResult, err := c.DialogSaveDraftMessage(&dialog.TLDialogSaveDraftMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
		Message:  tg.MakeTLDraftMessageEmpty(&tg.TLDraftMessageEmpty{}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(saveResult) {
		t.Fatalf("expected save draft boolTrue, got %#v", saveResult)
	}

	clearResult, err := c.DialogClearDraftMessage(&dialog.TLDialogClearDraftMessage{
		UserId:   1,
		PeerType: tg.PEER_USER,
		PeerId:   2,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(clearResult) {
		t.Fatalf("expected clear draft boolTrue, got %#v", clearResult)
	}

	allDrafts, err := c.DialogGetAllDrafts(&dialog.TLDialogGetAllDrafts{UserId: 1})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if allDrafts == nil || len(allDrafts.Datas) != 1 {
		t.Fatalf("expected one draft placeholder, got %#v", allDrafts)
	}

	clearedDrafts, err := c.DialogClearAllDrafts(&dialog.TLDialogClearAllDrafts{UserId: 1})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if clearedDrafts == nil || len(clearedDrafts.Datas) != 1 {
		t.Fatalf("expected one cleared draft placeholder, got %#v", clearedDrafts)
	}
}

func TestDialogFolderAndFilterPlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	folderResult, err := c.DialogGetDialogFolder(&dialog.TLDialogGetDialogFolder{
		UserId:   1,
		FolderId: 1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if folderResult == nil || len(folderResult.Datas) != 1 {
		t.Fatalf("expected one folder placeholder, got %#v", folderResult)
	}

	editFoldersResult, err := c.DialogEditPeerFolders(&dialog.TLDialogEditPeerFolders{
		UserId:         1,
		PeerDialogList: []int64{2, 3},
		FolderId:       1,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if editFoldersResult == nil || len(editFoldersResult.Datas) != 2 {
		t.Fatalf("expected two pinned placeholders, got %#v", editFoldersResult)
	}

	filterResult, err := c.DialogGetDialogFilter(&dialog.TLDialogGetDialogFilter{
		UserId: 1,
		Id:     9,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if filterResult == nil || filterResult.Clazz == nil {
		t.Fatal("expected dialog filter placeholder, got nil")
	}
	filterExt, ok := filterResult.Clazz.(*dialog.TLDialogFilterExt)
	if !ok || filterExt.Id != 9 {
		t.Fatalf("expected dialog filter id=9, got %#v", filterResult.Clazz)
	}

	filtersResult, err := c.DialogGetDialogFilters(&dialog.TLDialogGetDialogFilters{UserId: 1})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if filtersResult == nil || len(filtersResult.Datas) != 1 {
		t.Fatalf("expected one filter placeholder, got %#v", filtersResult)
	}

	filterBySlugResult, err := c.DialogGetDialogFilterBySlug(&dialog.TLDialogGetDialogFilterBySlug{
		UserId: 1,
		Slug:   "demo",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	filterBySlug, ok := filterBySlugResult.Clazz.(*dialog.TLDialogFilterExt)
	if !ok || filterBySlug.Slug != "demo" {
		t.Fatalf("expected slug=demo placeholder, got %#v", filterBySlugResult.Clazz)
	}

	tagsResult, err := c.DialogGetDialogFilterTags(&dialog.TLDialogGetDialogFilterTags{UserId: 1})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(tagsResult) {
		t.Fatalf("expected get tags boolTrue, got %#v", tagsResult)
	}

	toggleTagsResult, err := c.DialogToggleDialogFilterTags(&dialog.TLDialogToggleDialogFilterTags{
		UserId:  1,
		Enabled: tg.BoolTrueClazz,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(toggleTagsResult) {
		t.Fatalf("expected toggle tags boolTrue, got %#v", toggleTagsResult)
	}
}

func TestDialogFilterWritePlaceholders(t *testing.T) {
	c := New(context.Background(), nil)

	insertResult, err := c.DialogInsertOrUpdateDialogFilter(&dialog.TLDialogInsertOrUpdateDialogFilter{
		UserId: 1,
		Id:     7,
		DialogFilter: tg.MakeTLDialogFilter(&tg.TLDialogFilter{
			Id: 7,
			Title: tg.MakeTLTextWithEntities(&tg.TLTextWithEntities{
				Text:     "placeholder",
				Entities: []tg.MessageEntityClazz{},
			}),
			PinnedPeers:  []tg.InputPeerClazz{},
			IncludePeers: []tg.InputPeerClazz{},
			ExcludePeers: []tg.InputPeerClazz{},
		}),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(insertResult) {
		t.Fatalf("expected insert/update filter boolTrue, got %#v", insertResult)
	}

	createResult, err := c.DialogCreateDialogFilter(&dialog.TLDialogCreateDialogFilter{
		UserId:       1,
		DialogFilter: makeDialogFilterExtPlaceholder(8, "placeholder-8"),
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	createdFilter, ok := createResult.Clazz.(*dialog.TLDialogFilterExt)
	if !ok || createdFilter.Id != 8 {
		t.Fatalf("expected created filter id=8, got %#v", createResult.Clazz)
	}

	orderResult, err := c.DialogUpdateDialogFiltersOrder(&dialog.TLDialogUpdateDialogFiltersOrder{
		UserId: 1,
		Order:  []int32{8, 7},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(orderResult) {
		t.Fatalf("expected update order boolTrue, got %#v", orderResult)
	}

	deleteResult, err := c.DialogDeleteDialogFilter(&dialog.TLDialogDeleteDialogFilter{
		UserId: 1,
		Id:     8,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !tg.FromBool(deleteResult) {
		t.Fatalf("expected delete filter boolTrue, got %#v", deleteResult)
	}
}
