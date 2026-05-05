package repository

import "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"

type DialogRecord struct {
	UserID               int64
	PeerType             int32
	PeerID               int64
	PeerDialogID         int64
	Order                int64
	Pinned               int64
	TopMessage           int32
	PinnedMsgID          int32
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	UnreadMentionsCount  int32
	UnreadReactionsCount int32
	UnreadMark           bool
	DraftType            int32
	DraftMessageData     string
	FolderID             int32
	FolderPinned         int64
	HasScheduled         bool
	TTLPeriod            int32
	ThemeEmoticon        string
	WallpaperID          int64
	WallpaperOverridden  bool
	Date                 int64
}

func mapDialogRecord(row model.Dialogs) DialogRecord {
	return DialogRecord{
		UserID:               row.UserId,
		PeerType:             row.PeerType,
		PeerID:               row.PeerId,
		PeerDialogID:         row.PeerDialogId,
		Pinned:               row.Pinned,
		TopMessage:           row.TopMessage,
		PinnedMsgID:          row.PinnedMsgId,
		ReadInboxMaxID:       row.ReadInboxMaxId,
		ReadOutboxMaxID:      row.ReadOutboxMaxId,
		UnreadCount:          row.UnreadCount,
		UnreadMentionsCount:  row.UnreadMentionsCount,
		UnreadReactionsCount: row.UnreadReactionsCount,
		UnreadMark:           row.UnreadMark,
		DraftType:            row.DraftType,
		DraftMessageData:     row.DraftMessageData,
		FolderID:             row.FolderId,
		FolderPinned:         row.FolderPinned,
		HasScheduled:         row.HasScheduled,
		TTLPeriod:            row.TtlPeriod,
		ThemeEmoticon:        row.ThemeEmoticon,
		WallpaperID:          row.WallpaperId,
		WallpaperOverridden:  row.WallpaperOverridden,
		Date:                 row.Date2,
	}
}

func mapDialogRecords(rows []model.Dialogs) []DialogRecord {
	if len(rows) == 0 {
		return []DialogRecord{}
	}
	out := make([]DialogRecord, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapDialogRecord(row))
	}
	return out
}
