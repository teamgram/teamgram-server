package core

import (
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
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

func makeDialogExtV2FromProjection(projection *userupdates.TLDialogProjection, extras *dialogpb.DialogExtras) *dialogpb.DialogExtV2 {
	if extras == nil {
		extras = makeDialogExtras(repository.DialogExtrasRecord{
			PeerType: projection.PeerType,
			PeerID:   projection.PeerId,
		})
	}
	return dialogpb.MakeTLDialogExtV2(&dialogpb.TLDialogExtV2{
		PeerType:                 projection.PeerType,
		PeerId:                   projection.PeerId,
		TopPeerSeq:               projection.TopPeerSeq,
		TopCanonicalMessageId:    projection.TopCanonicalMessageId,
		TopMessageDate:           projection.TopMessageDate,
		UnreadCount:              projection.UnreadCount,
		UnreadMentionsCount:      projection.UnreadMentionsCount,
		UnreadReactionsCount:     projection.UnreadReactionsCount,
		UnreadMark:               projection.UnreadMark,
		PinnedPeerSeq:            projection.PinnedPeerSeq,
		PinnedCanonicalMessageId: projection.PinnedCanonicalMessageId,
		HasScheduled:             projection.HasScheduled,
		AvailableMinPeerSeq:      projection.AvailableMinPeerSeq,
		FolderId:                 extras.FolderId,
		MainPinnedOrder:          extras.MainPinnedOrder,
		FolderPinnedOrder:        extras.FolderPinnedOrder,
		Extras:                   extras,
	})
}

func makeDialogExtV2VectorFromProjections(projections []*userupdates.TLDialogProjection, extras []repository.DialogExtrasRecord) *dialogpb.VectorDialogExtV2 {
	extrasByPeer := make(map[repository.PeerRef]*dialogpb.DialogExtras, len(extras))
	for _, record := range extras {
		extrasByPeer[repository.PeerRef{PeerType: record.PeerType, PeerID: record.PeerID}] = makeDialogExtras(record)
	}
	out := &dialogpb.VectorDialogExtV2{Datas: make([]dialogpb.DialogExtV2Clazz, 0, len(projections))}
	for _, projection := range projections {
		if projection == nil {
			continue
		}
		out.Datas = append(out.Datas, makeDialogExtV2FromProjection(projection, extrasByPeer[repository.PeerRef{PeerType: projection.PeerType, PeerID: projection.PeerId}]))
	}
	return out
}

func dialogProjectionPeersFromRefs(peers []repository.PeerRef) []userupdates.DialogProjectionPeerClazz {
	out := make([]userupdates.DialogProjectionPeerClazz, 0, len(peers))
	for _, peer := range peers {
		out = append(out, userupdates.MakeTLDialogProjectionPeer(&userupdates.TLDialogProjectionPeer{
			PeerType: peer.PeerType,
			PeerId:   peer.PeerID,
		}))
	}
	return out
}

func projectionPeerRef(projection *userupdates.TLDialogProjection) (repository.PeerRef, error) {
	if projection == nil {
		return repository.PeerRef{}, errors.New("nil dialog projection")
	}
	return repository.PeerRef{PeerType: projection.PeerType, PeerID: projection.PeerId}, nil
}

func makeDialogFilterExt(record repository.DialogFilterRecord) *dialogpb.DialogFilterExt {
	filter := tg.MakeTLDialogFilter(&tg.TLDialogFilter{
		Id: record.DialogFilterID,
		Title: tg.MakeTLTextWithEntities(&tg.TLTextWithEntities{
			Text:     record.Title,
			Entities: []tg.MessageEntityClazz{},
		}),
		PinnedPeers:  []tg.InputPeerClazz{},
		IncludePeers: []tg.InputPeerClazz{},
		ExcludePeers: []tg.InputPeerClazz{},
	})
	return dialogpb.MakeTLDialogFilterExt(&dialogpb.TLDialogFilterExt{
		Id:           record.DialogFilterID,
		Slug:         record.Slug,
		DialogFilter: filter,
		Order:        record.OrderValue,
	})
}

func makeDialogFilterExtVector(records []repository.DialogFilterRecord) *dialogpb.VectorDialogFilterExt {
	out := &dialogpb.VectorDialogFilterExt{Datas: make([]dialogpb.DialogFilterExtClazz, 0, len(records))}
	for _, record := range records {
		out.Datas = append(out.Datas, makeDialogFilterExt(record))
	}
	return out
}
