package core

import (
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func makeDialogExt(record repository.DialogRecord) *dialogpb.DialogExt {
	folderID := record.FolderID
	ttlPeriod := record.TTLPeriod

	return dialogpb.MakeTLDialogExt(&dialogpb.TLDialogExt{
		Order: record.PeerDialogID,
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
