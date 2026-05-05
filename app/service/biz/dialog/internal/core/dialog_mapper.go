package core

import (
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func makeDialogExt(record repository.DialogRecord) *dialogpb.DialogExt {
	folderID := record.FolderID
	ttlPeriod := record.TTLPeriod
	order := record.Order
	if order == 0 {
		order = record.PeerDialogID
	}

	return dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
		Order: order,
		Dialog: tg.MakeTLDialog(&tg.TLDialog{
			Pinned:               record.Pinned != 0,
			UnreadMark:           record.UnreadMark,
			Peer:                 tg.MakePeerHelper(record.PeerType, record.PeerID),
			TopMessage:           record.TopMessage,
			ReadInboxMaxId:       record.ReadInboxMaxID,
			ReadOutboxMaxId:      record.ReadOutboxMaxID,
			UnreadCount:          record.UnreadCount,
			UnreadMentionsCount:  record.UnreadMentionsCount,
			UnreadReactionsCount: record.UnreadReactionsCount,
			NotifySettings:       tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
			FolderId:             &folderID,
			TtlPeriod:            &ttlPeriod,
		}),
		AvailableMinId:      0,
		Date:                record.Date,
		ThemeEmoticon:       record.ThemeEmoticon,
		TtlPeriod:           record.TTLPeriod,
		WallpaperId:         record.WallpaperID,
		WallpaperOverridden: record.WallpaperOverridden,
	})
}

func makeDialogExtVector(records []repository.DialogRecord) *dialogpb.VectorDialogExt {
	out := &dialogpb.VectorDialogExt{Datas: make([]dialogpb.DialogExtClazz, 0, len(records))}
	for _, record := range records {
		out.Datas = append(out.Datas, makeDialogExt(record))
	}
	return out
}

func makeDialogExtras(record repository.DialogExtrasRecord) *dialogpb.DialogExtras {
	var ttl *int32
	if record.PrivateTTLPeriod != 0 {
		ttl = &record.PrivateTTLPeriod
	}
	var theme *string
	if record.PrivateThemeEmoticon != "" {
		theme = &record.PrivateThemeEmoticon
	}
	var wallpaperID *int64
	if record.WallpaperID != 0 {
		wallpaperID = &record.WallpaperID
	}
	return dialogpb.MakeTLDialogExtras(&dialogpb.TLDialogExtras{
		PeerType:             record.PeerType,
		PeerId:               record.PeerID,
		FolderId:             record.FolderID,
		MainPinnedOrder:      record.MainPinnedOrder,
		FolderPinnedOrder:    record.FolderPinnedOrder,
		DraftPayload:         record.DraftPayload,
		PrivateTtlPeriod:     ttl,
		PrivateThemeEmoticon: theme,
		WallpaperId:          wallpaperID,
		WallpaperOverridden:  record.WallpaperOverridden,
	})
}

func makeDialogExtrasVector(records []repository.DialogExtrasRecord) *dialogpb.VectorDialogExtras {
	out := &dialogpb.VectorDialogExtras{Datas: make([]dialogpb.DialogExtrasClazz, 0, len(records))}
	for _, record := range records {
		out.Datas = append(out.Datas, makeDialogExtras(record))
	}
	return out
}

func makeDialogExtV2(record repository.DialogRecord, extras *dialogpb.DialogExtras) *dialogpb.DialogExtV2 {
	if extras == nil {
		extras = makeDialogExtras(repository.DialogExtrasRecord{
			PeerType:            record.PeerType,
			PeerID:              record.PeerID,
			FolderID:            record.FolderID,
			MainPinnedOrder:     record.Pinned,
			FolderPinnedOrder:   record.FolderPinned,
			WallpaperID:         record.WallpaperID,
			WallpaperOverridden: record.WallpaperOverridden,
		})
	}
	return dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
		PeerType:                 record.PeerType,
		PeerId:                   record.PeerID,
		TopPeerSeq:               int64(record.TopMessage),
		TopCanonicalMessageId:    int64(record.TopMessage),
		TopMessageDate:           record.Date,
		UnreadCount:              record.UnreadCount,
		UnreadMentionsCount:      record.UnreadMentionsCount,
		UnreadReactionsCount:     record.UnreadReactionsCount,
		UnreadMark:               record.UnreadMark,
		PinnedPeerSeq:            int64(record.PinnedMsgID),
		PinnedCanonicalMessageId: int64(record.PinnedMsgID),
		HasScheduled:             record.HasScheduled,
		AvailableMinPeerSeq:      0,
		FolderId:                 record.FolderID,
		MainPinnedOrder:          record.Pinned,
		FolderPinnedOrder:        record.FolderPinned,
		Extras:                   extras,
	})
}

func makeDialogExtV2Vector(records []repository.DialogRecord, extras []repository.DialogExtrasRecord) *dialogpb.VectorDialogExtV2 {
	extrasByPeer := make(map[repository.PeerRef]*dialogpb.DialogExtras, len(extras))
	for _, record := range extras {
		extrasByPeer[repository.PeerRef{PeerType: record.PeerType, PeerID: record.PeerID}] = makeDialogExtras(record)
	}
	out := &dialogpb.VectorDialogExtV2{Datas: make([]dialogpb.DialogExtV2Clazz, 0, len(records))}
	for _, record := range records {
		out.Datas = append(out.Datas, makeDialogExtV2(record, extrasByPeer[repository.PeerRef{PeerType: record.PeerType, PeerID: record.PeerID}]))
	}
	return out
}
