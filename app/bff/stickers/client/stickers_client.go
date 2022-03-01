/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package stickers_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type StickersClient interface {
	MessagesGetStickers(ctx context.Context, in *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error)
	MessagesGetAllStickers(ctx context.Context, in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error)
	MessagesGetStickerSet(ctx context.Context, in *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error)
	MessagesInstallStickerSet(ctx context.Context, in *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error)
	MessagesUninstallStickerSet(ctx context.Context, in *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error)
	MessagesReorderStickerSets(ctx context.Context, in *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error)
	MessagesGetFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error)
	MessagesReadFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error)
	MessagesGetRecentStickers(ctx context.Context, in *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error)
	MessagesSaveRecentSticker(ctx context.Context, in *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error)
	MessagesClearRecentStickers(ctx context.Context, in *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error)
	MessagesGetArchivedStickers(ctx context.Context, in *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error)
	MessagesGetMaskStickers(ctx context.Context, in *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error)
	MessagesGetAttachedStickers(ctx context.Context, in *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error)
	MessagesGetFavedStickers(ctx context.Context, in *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error)
	MessagesFaveSticker(ctx context.Context, in *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error)
	MessagesSearchStickerSets(ctx context.Context, in *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error)
	MessagesToggleStickerSets(ctx context.Context, in *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error)
	MessagesGetOldFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error)
	StickersCreateStickerSet(ctx context.Context, in *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error)
	StickersRemoveStickerFromSet(ctx context.Context, in *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error)
	StickersChangeStickerPosition(ctx context.Context, in *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error)
	StickersAddStickerToSet(ctx context.Context, in *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error)
	StickersSetStickerSetThumb(ctx context.Context, in *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error)
	StickersCheckShortName(ctx context.Context, in *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error)
	StickersSuggestShortName(ctx context.Context, in *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error)
}

type defaultStickersClient struct {
	cli zrpc.Client
}

func NewStickersClient(cli zrpc.Client) StickersClient {
	return &defaultStickersClient{
		cli: cli,
	}
}

// MessagesGetStickers
// messages.getStickers#d5a5d3a1 emoticon:string hash:long = messages.Stickers;
func (m *defaultStickersClient) MessagesGetStickers(ctx context.Context, in *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetStickers(ctx, in)
}

// MessagesGetAllStickers
// messages.getAllStickers#b8a0a1a8 hash:long = messages.AllStickers;
func (m *defaultStickersClient) MessagesGetAllStickers(ctx context.Context, in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetAllStickers(ctx, in)
}

// MessagesGetStickerSet
// messages.getStickerSet#c8a0ec74 stickerset:InputStickerSet hash:int = messages.StickerSet;
func (m *defaultStickersClient) MessagesGetStickerSet(ctx context.Context, in *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetStickerSet(ctx, in)
}

// MessagesInstallStickerSet
// messages.installStickerSet#c78fe460 stickerset:InputStickerSet archived:Bool = messages.StickerSetInstallResult;
func (m *defaultStickersClient) MessagesInstallStickerSet(ctx context.Context, in *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesInstallStickerSet(ctx, in)
}

// MessagesUninstallStickerSet
// messages.uninstallStickerSet#f96e55de stickerset:InputStickerSet = Bool;
func (m *defaultStickersClient) MessagesUninstallStickerSet(ctx context.Context, in *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesUninstallStickerSet(ctx, in)
}

// MessagesReorderStickerSets
// messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;
func (m *defaultStickersClient) MessagesReorderStickerSets(ctx context.Context, in *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesReorderStickerSets(ctx, in)
}

// MessagesGetFeaturedStickers
// messages.getFeaturedStickers#64780b14 hash:long = messages.FeaturedStickers;
func (m *defaultStickersClient) MessagesGetFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetFeaturedStickers(ctx, in)
}

// MessagesReadFeaturedStickers
// messages.readFeaturedStickers#5b118126 id:Vector<long> = Bool;
func (m *defaultStickersClient) MessagesReadFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesReadFeaturedStickers(ctx, in)
}

// MessagesGetRecentStickers
// messages.getRecentStickers#9da9403b flags:# attached:flags.0?true hash:long = messages.RecentStickers;
func (m *defaultStickersClient) MessagesGetRecentStickers(ctx context.Context, in *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetRecentStickers(ctx, in)
}

// MessagesSaveRecentSticker
// messages.saveRecentSticker#392718f8 flags:# attached:flags.0?true id:InputDocument unsave:Bool = Bool;
func (m *defaultStickersClient) MessagesSaveRecentSticker(ctx context.Context, in *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesSaveRecentSticker(ctx, in)
}

// MessagesClearRecentStickers
// messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool;
func (m *defaultStickersClient) MessagesClearRecentStickers(ctx context.Context, in *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesClearRecentStickers(ctx, in)
}

// MessagesGetArchivedStickers
// messages.getArchivedStickers#57f17692 flags:# masks:flags.0?true offset_id:long limit:int = messages.ArchivedStickers;
func (m *defaultStickersClient) MessagesGetArchivedStickers(ctx context.Context, in *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetArchivedStickers(ctx, in)
}

// MessagesGetMaskStickers
// messages.getMaskStickers#640f82b8 hash:long = messages.AllStickers;
func (m *defaultStickersClient) MessagesGetMaskStickers(ctx context.Context, in *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetMaskStickers(ctx, in)
}

// MessagesGetAttachedStickers
// messages.getAttachedStickers#cc5b67cc media:InputStickeredMedia = Vector<StickerSetCovered>;
func (m *defaultStickersClient) MessagesGetAttachedStickers(ctx context.Context, in *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetAttachedStickers(ctx, in)
}

// MessagesGetFavedStickers
// messages.getFavedStickers#4f1aaa9 hash:long = messages.FavedStickers;
func (m *defaultStickersClient) MessagesGetFavedStickers(ctx context.Context, in *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetFavedStickers(ctx, in)
}

// MessagesFaveSticker
// messages.faveSticker#b9ffc55b id:InputDocument unfave:Bool = Bool;
func (m *defaultStickersClient) MessagesFaveSticker(ctx context.Context, in *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesFaveSticker(ctx, in)
}

// MessagesSearchStickerSets
// messages.searchStickerSets#35705b8a flags:# exclude_featured:flags.0?true q:string hash:long = messages.FoundStickerSets;
func (m *defaultStickersClient) MessagesSearchStickerSets(ctx context.Context, in *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesSearchStickerSets(ctx, in)
}

// MessagesToggleStickerSets
// messages.toggleStickerSets#b5052fea flags:# uninstall:flags.0?true archive:flags.1?true unarchive:flags.2?true stickersets:Vector<InputStickerSet> = Bool;
func (m *defaultStickersClient) MessagesToggleStickerSets(ctx context.Context, in *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesToggleStickerSets(ctx, in)
}

// MessagesGetOldFeaturedStickers
// messages.getOldFeaturedStickers#7ed094a1 offset:int limit:int hash:long = messages.FeaturedStickers;
func (m *defaultStickersClient) MessagesGetOldFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.MessagesGetOldFeaturedStickers(ctx, in)
}

// StickersCreateStickerSet
// stickers.createStickerSet#9021ab67 flags:# masks:flags.0?true animated:flags.1?true user_id:InputUser title:string short_name:string thumb:flags.2?InputDocument stickers:Vector<InputStickerSetItem> software:flags.3?string = messages.StickerSet;
func (m *defaultStickersClient) StickersCreateStickerSet(ctx context.Context, in *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersCreateStickerSet(ctx, in)
}

// StickersRemoveStickerFromSet
// stickers.removeStickerFromSet#f7760f51 sticker:InputDocument = messages.StickerSet;
func (m *defaultStickersClient) StickersRemoveStickerFromSet(ctx context.Context, in *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersRemoveStickerFromSet(ctx, in)
}

// StickersChangeStickerPosition
// stickers.changeStickerPosition#ffb6d4ca sticker:InputDocument position:int = messages.StickerSet;
func (m *defaultStickersClient) StickersChangeStickerPosition(ctx context.Context, in *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersChangeStickerPosition(ctx, in)
}

// StickersAddStickerToSet
// stickers.addStickerToSet#8653febe stickerset:InputStickerSet sticker:InputStickerSetItem = messages.StickerSet;
func (m *defaultStickersClient) StickersAddStickerToSet(ctx context.Context, in *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersAddStickerToSet(ctx, in)
}

// StickersSetStickerSetThumb
// stickers.setStickerSetThumb#9a364e30 stickerset:InputStickerSet thumb:InputDocument = messages.StickerSet;
func (m *defaultStickersClient) StickersSetStickerSetThumb(ctx context.Context, in *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersSetStickerSetThumb(ctx, in)
}

// StickersCheckShortName
// stickers.checkShortName#284b3639 short_name:string = Bool;
func (m *defaultStickersClient) StickersCheckShortName(ctx context.Context, in *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersCheckShortName(ctx, in)
}

// StickersSuggestShortName
// stickers.suggestShortName#4dafc503 title:string = stickers.SuggestedShortName;
func (m *defaultStickersClient) StickersSuggestShortName(ctx context.Context, in *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	client := mtproto.NewRPCStickersClient(m.cli.Conn())
	return client.StickersSuggestShortName(ctx, in)
}
