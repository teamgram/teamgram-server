/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/stickers/internal/core"
)

// MessagesGetStickers
// messages.getStickers#d5a5d3a1 emoticon:string hash:long = messages.Stickers;
func (s *Service) MessagesGetStickers(ctx context.Context, request *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAllStickers
// messages.getAllStickers#b8a0a1a8 hash:long = messages.AllStickers;
func (s *Service) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAllStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAllStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAllStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetStickerSet
// messages.getStickerSet#c8a0ec74 stickerset:InputStickerSet hash:int = messages.StickerSet;
func (s *Service) MessagesGetStickerSet(ctx context.Context, request *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getStickerSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getStickerSet - reply: %s", r.DebugString())
	return r, err
}

// MessagesInstallStickerSet
// messages.installStickerSet#c78fe460 stickerset:InputStickerSet archived:Bool = messages.StickerSetInstallResult;
func (s *Service) MessagesInstallStickerSet(ctx context.Context, request *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.installStickerSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesInstallStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.installStickerSet - reply: %s", r.DebugString())
	return r, err
}

// MessagesUninstallStickerSet
// messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;
func (s *Service) MessagesUninstallStickerSet(ctx context.Context, request *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.uninstallStickerSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesUninstallStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.uninstallStickerSet - reply: %s", r.DebugString())
	return r, err
}

// MessagesReorderStickerSets
// messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;
func (s *Service) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.reorderStickerSets - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReorderStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.reorderStickerSets - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetFeaturedStickers
// messages.getFeaturedStickers#64780b14 hash:long = messages.FeaturedStickers;
func (s *Service) MessagesGetFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getFeaturedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getFeaturedStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesReadFeaturedStickers
// messages.readFeaturedStickers#5b118126 id:Vector<long> = Bool;
func (s *Service) MessagesReadFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.readFeaturedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesReadFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.readFeaturedStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetRecentStickers
// messages.getRecentStickers#9da9403b flags:# attached:flags.0?true hash:long = messages.RecentStickers;
func (s *Service) MessagesGetRecentStickers(ctx context.Context, request *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getRecentStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetRecentStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getRecentStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesSaveRecentSticker
// messages.saveRecentSticker#392718f8 flags:# attached:flags.0?true id:InputDocument unsave:Bool = Bool;
func (s *Service) MessagesSaveRecentSticker(ctx context.Context, request *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.saveRecentSticker - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSaveRecentSticker(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.saveRecentSticker - reply: %s", r.DebugString())
	return r, err
}

// MessagesClearRecentStickers
// messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;
func (s *Service) MessagesClearRecentStickers(ctx context.Context, request *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.clearRecentStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesClearRecentStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.clearRecentStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetArchivedStickers
// messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;
func (s *Service) MessagesGetArchivedStickers(ctx context.Context, request *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getArchivedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetArchivedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getArchivedStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetMaskStickers
// messages.getMaskStickers#640f82b8 hash:long = messages.AllStickers;
func (s *Service) MessagesGetMaskStickers(ctx context.Context, request *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getMaskStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetMaskStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getMaskStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetAttachedStickers
// messages.getAttachedStickers#cc5b67cc media:InputStickeredMedia = Vector<StickerSetCovered>;
func (s *Service) MessagesGetAttachedStickers(ctx context.Context, request *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getAttachedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetAttachedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getAttachedStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetFavedStickers
// messages.getFavedStickers#4f1aaa9 hash:long = messages.FavedStickers;
func (s *Service) MessagesGetFavedStickers(ctx context.Context, request *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getFavedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetFavedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getFavedStickers - reply: %s", r.DebugString())
	return r, err
}

// MessagesFaveSticker
// messages.faveSticker#b9ffc55b id:InputDocument unfave:Bool = Bool;
func (s *Service) MessagesFaveSticker(ctx context.Context, request *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.faveSticker - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesFaveSticker(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.faveSticker - reply: %s", r.DebugString())
	return r, err
}

// MessagesSearchStickerSets
// messages.searchStickerSets#35705b8a flags:# exclude_featured:flags.0?true q:string hash:long = messages.FoundStickerSets;
func (s *Service) MessagesSearchStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.searchStickerSets - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesSearchStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.searchStickerSets - reply: %s", r.DebugString())
	return r, err
}

// MessagesToggleStickerSets
// messages.toggleStickerSets#b5052fea flags:# uninstall:flags.0?true archive:flags.1?true unarchive:flags.2?true stickersets:Vector<InputStickerSet> = Bool;
func (s *Service) MessagesToggleStickerSets(ctx context.Context, request *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.toggleStickerSets - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesToggleStickerSets(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.toggleStickerSets - reply: %s", r.DebugString())
	return r, err
}

// MessagesGetOldFeaturedStickers
// messages.getOldFeaturedStickers#7ed094a1 offset:int limit:int hash:long = messages.FeaturedStickers;
func (s *Service) MessagesGetOldFeaturedStickers(ctx context.Context, request *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("messages.getOldFeaturedStickers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.MessagesGetOldFeaturedStickers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("messages.getOldFeaturedStickers - reply: %s", r.DebugString())
	return r, err
}

// StickersCreateStickerSet
// stickers.createStickerSet#9021ab67 flags:# masks:flags.0?true animated:flags.1?true user_id:InputUser title:string short_name:string thumb:flags.2?InputDocument stickers:Vector<InputStickerSetItem> software:flags.3?string = messages.StickerSet;
func (s *Service) StickersCreateStickerSet(ctx context.Context, request *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.createStickerSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersCreateStickerSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.createStickerSet - reply: %s", r.DebugString())
	return r, err
}

// StickersRemoveStickerFromSet
// stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
func (s *Service) StickersRemoveStickerFromSet(ctx context.Context, request *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.removeStickerFromSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersRemoveStickerFromSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.removeStickerFromSet - reply: %s", r.DebugString())
	return r, err
}

// StickersChangeStickerPosition
// stickers.changeStickerPosition#ffb6d4ca sticker:InputDocument position:int = messages.StickerSet;
func (s *Service) StickersChangeStickerPosition(ctx context.Context, request *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.changeStickerPosition - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersChangeStickerPosition(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.changeStickerPosition - reply: %s", r.DebugString())
	return r, err
}

// StickersAddStickerToSet
// stickers.addStickerToSet#8653febe stickerset:InputStickerSet sticker:InputStickerSetItem = messages.StickerSet;
func (s *Service) StickersAddStickerToSet(ctx context.Context, request *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.addStickerToSet - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersAddStickerToSet(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.addStickerToSet - reply: %s", r.DebugString())
	return r, err
}

// StickersSetStickerSetThumb
// stickers.setStickerSetThumb#9a364e30 stickerset:InputStickerSet thumb:InputDocument = messages.StickerSet;
func (s *Service) StickersSetStickerSetThumb(ctx context.Context, request *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.setStickerSetThumb - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersSetStickerSetThumb(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.setStickerSetThumb - reply: %s", r.DebugString())
	return r, err
}

// StickersCheckShortName
// stickers.checkShortName#284b3639 short_name:string = Bool;
func (s *Service) StickersCheckShortName(ctx context.Context, request *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.checkShortName - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersCheckShortName(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.checkShortName - reply: %s", r.DebugString())
	return r, err
}

// StickersSuggestShortName
// stickers.suggestShortName#4dafc503 title:string = stickers.SuggestedShortName;
func (s *Service) StickersSuggestShortName(ctx context.Context, request *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("stickers.suggestShortName - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.StickersSuggestShortName(request)
	if err != nil {
		return nil, err
	}

	c.Infof("stickers.suggestShortName - reply: %s", r.DebugString())
	return r, err
}
